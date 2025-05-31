package service

import (
	"context"

	"github.com/pradeepitm12/gitsearch/internal/git"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGithubSearchServiceServer
	token string
}

func New(token string) pb.GithubSearchServiceServer {
	return &server{token: token}
}

func (s *server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	searcher, err := git.NewSearcher(req.Type, s.token)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid search type: %v", err)
	}
	page := req.GetPage()
	perPage := req.GetPerPage()
	if page <= 0 {
		page = 1
	}
	if perPage <= 0 || perPage > 100 {
		perPage = 30
	}

	results, err := searcher.Search(req.SearchTerm, req.User, int(page), int(perPage))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "search failed: %v", err)
	}

	var response pb.SearchResponse
	for _, r := range results {
		response.Results = append(response.Results, &pb.Result{
			FileUrl: r.FileURL,
			Repo:    r.Repo,
		})
	}
	
	return &response, nil
}
