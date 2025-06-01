package main

import (
	"github.com/pradeepitm12/gitsearch/internal/git"
	"github.com/pradeepitm12/gitsearch/internal/middleware"
	"golang.org/x/time/rate"
	"log"
	"net"

	"github.com/pradeepitm12/gitsearch/internal/config"
	"github.com/pradeepitm12/gitsearch/internal/service"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg := config.Load()
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	limiter := rate.NewLimiter(10, 20)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.RateLimit(limiter)),
	)
	client := git.NewClient(cfg.GitHubToken)
	pb.RegisterGithubSearchServiceServer(grpcServer, service.NewServer(client))

	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
