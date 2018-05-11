package main

import (
	"log"
	"net"

	pb "grpc-labs/proto/foo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type Repository interface {
	Create(foo *pb.Foo) (*pb.Foo, error)
	GetAll() []*pb.Foo
}

type FooRepository struct {
	foos []*pb.Foo
}

func (repo *FooRepository) Create(foo *pb.Foo) (*pb.Foo, error) {
	updated := append(repo.foos, foo)
	repo.foos = updated
	return foo, nil
}

func (repo *FooRepository) GetAll() []*pb.Foo {
	return repo.foos
}

type service struct {
	repo Repository
}

func (s *service) CreateFoo(ctx context.Context, req *pb.Foo) (*pb.FooResponse, error) {
	foo, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	log.Println("Create call: CreateFoo")
	return &pb.FooResponse{Created: true, Foo: foo}, nil
}

func (s *service) GetFoo(ctx context.Context, req *pb.FooRequest) (*pb.FooResponse, error) {
	foos := s.repo.GetAll()
	log.Println("Client call: GetFoo")
	return &pb.FooResponse{Foos: foos}, nil
}

func main() {
	repo := &FooRepository{}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFooServiceServer(s, &service{repo})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
