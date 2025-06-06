package main

import (
	"context"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"github.com/pradeepitm12/gitsearch/internal/config"
	"github.com/pradeepitm12/gitsearch/internal/git"
	"github.com/pradeepitm12/gitsearch/internal/middleware"
	"github.com/pradeepitm12/gitsearch/internal/service"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
)

func main() {
	limiter := rate.NewLimiter(10, 20)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.RateLimit(limiter)),
	)
	cfg := config.Load()
	ctx := context.Background()
	client := git.NewClient(ctx, cfg.GitHubToken)
	pb.RegisterGithubSearchServiceServer(grpcServer, service.NewServer(client))

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
