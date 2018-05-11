package main

import (
	"log"

	pb "grpc-labs/proto/foo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func mockFoo() *pb.Foo {
	return &pb.Foo{
		Id:      1,
		Message: "Hello",
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewFooServiceClient(conn)

	r, err := client.CreateFoo(context.Background(), mockFoo())
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetFoo(context.Background(), &pb.FooRequest{})
	if err != nil {
		log.Fatalf("Could not list topics: %v", err)
	}
	for _, v := range getAll.Foos {
		log.Println(v)
	}
}
