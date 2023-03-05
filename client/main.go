package main

import (
	"context"
	pb "grpc/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not able to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewEmailGeneratorClient(conn)
	names := &pb.NameList{
		Names: []string{"ramasamy", "krishnamoorthy", "maarthandam"},
	}
	//unaryConnection(client)
	//serverStreamConnection(client, names)
	//clientStreamingConnection(client, names)
	bidirectionalStreamingConnection(client, names)
}

// single client request single server response

func unaryConnection(client pb.EmailGeneratorClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UnaryConnection(ctx, &pb.Name{
		Name: "ramasamy",
	})
	if err != nil {
		log.Fatalf("Error with unary connection: %v", err)
	}

	log.Println(res)
}

// single client request server streaming

func serverStreamConnection(client pb.EmailGeneratorClient, names *pb.NameList) {
	log.Printf("Server streaming started")
	stream, err := client.ServerStreamConnection(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not stream from server: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming: %v", err)
		}
		log.Println(message)
	}
	log.Println("streaming finished")
}

// client streaming server single response

func clientStreamingConnection(client pb.EmailGeneratorClient, names *pb.NameList) {
	log.Println("Client streaming started")
	stream, err := client.ClientStreamConnection(context.Background())
	if err != nil {
		log.Fatalf("Error while starting stream: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.Name{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending stream: %v", err)
		}
		log.Printf("Sent request with name: %s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving: %v", err)
	}

	log.Printf("%s", res.EmailList)
}

// client streaming server streaming
func bidirectionalStreamingConnection(client pb.EmailGeneratorClient, name *pb.NameList) {
	log.Println("bidirectional streaming started")
	stream, err := client.BidirectionalStreamConnection(context.Background())
	if err != nil {
		log.Fatalf("error while streaming: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving stream: %v", err)
			}
			log.Println(message.Email)
		}
		close(waitc)
	}()

	for _, name := range name.Names {
		req := &pb.Name{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		// time.Sleep(2 * time.Second)

	}

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
