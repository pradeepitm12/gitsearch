package main

import (
	"github.com/pradeepitm12/gitsearch/internal/service"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	srv := service.NewGitSearchServer() // You'll define this

	pb.RegisterGithubSearchServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)
	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
