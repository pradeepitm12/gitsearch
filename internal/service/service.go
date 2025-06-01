package service

import (
	"context"
	"github.com/pradeepitm12/gitsearch/internal/git"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	defaultPage    = 1
	defaultPerPage = 30
)

type SearchStrategy interface {
	Search(ctx context.Context, client git.GitClient, query string, page, perPage int) ([]*pb.Result, error)
}

type Server struct {
	pb.UnimplementedGithubSearchServiceServer
	client   git.GitClient
	registry map[string]SearchStrategy
}

func NewServer(client git.GitClient) *Server {
	return &Server{
		client: client,
		registry: map[string]SearchStrategy{
			"code":         &CodeStrategy{},
			"commits":      &CommitsStrategy{},
			"issues":       &IssuesStrategy{},
			"repositories": &RepositoriesStrategy{},
			"topics":       &TopicsStrategy{},
			"users":        &UsersStrategy{},
		},
	}
}

func (s *Server) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	if req.User != "" && !supportsUserScope(req.Type) {
		return nil, status.Errorf(codes.InvalidArgument, "user filter is not supported for search type %q", req.Type)
	}

	query := req.SearchTerm
	if req.User != "" {
		query += " user:" + req.User
	}
	strategy, ok := s.registry[req.Type]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "unsupported type: %s", req.Type)
	}

	page := int(req.Page)
	if page <= 0 {
		page = defaultPage
	}

	perPage := int(req.PerPage)
	if perPage <= 0 || perPage > 100 {
		perPage = defaultPerPage
	}

	results, err := strategy.Search(ctx, s.client, query, page, perPage)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "search failed: %v", err)
	}

	return &pb.SearchResponse{Results: results}, nil
}

func supportsUserScope(t string) bool {
	return t == "code" || t == "commits" || t == "issues"
}
