package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type emailServer struct {
	pb.EmailGeneratorServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error while listening: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEmailGeneratorServer(grpcServer, &emailServer{})
	log.Printf("Server started at port: %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error while starting server: %s", err)
	}

}

//single client client request single server response

func (e *emailServer) UnaryConnection(ctx context.Context, req *pb.Name) (*pb.Email, error) {
	log.Printf("Getting req with name: %s", req.Name)
	return &pb.Email{
		Email: req.Name + "@gmail.com",
	}, nil
}

// Single client request server streaming

func (e *emailServer) ServerStreamConnection(req *pb.NameList, stream pb.EmailGenerator_ServerStreamConnectionServer) error {
	log.Printf("Got request with name: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.Email{
			Email: name + "@gmail.com",
		}

		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}

// Client Streaming Server Single Response

func (e *emailServer) ClientStreamConnection(stream pb.EmailGenerator_ClientStreamConnectionServer) error {
	var message []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.EmailList{EmailList: message})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %s", req.Name)
		message = append(message, req.Name+"@gmail.com")
	}
}

// client streaming server streaming - bidirectional streaming

func (e *emailServer) BidirectionalStreamConnection(stream pb.EmailGenerator_BidirectionalStreamConnectionServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %s", req.Name)
		res := &pb.Email{
			Email: req.Name + "@gmail.com",
		}

		if err := stream.Send(res); err != nil {
			return err
		}

	}
}
