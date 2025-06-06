package service

import (
	"context"

	"github.com/pradeepitm12/gitsearch/internal/git"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
)

type usersStrategy struct{}

func (c *usersStrategy) Search(ctx context.Context, client git.GitClient, query string, page, perPage int) ([]*pb.Result, error) {
	results, err := client.SearchUsers(ctx, query, page, perPage)
	if err != nil {
		return nil, err
	}

	var pbResults []*pb.Result
	for _, r := range results {
		pbResults = append(pbResults, &pb.Result{
			FileUrl: r.GetHTMLURL(),
			Repo:    r.GetLogin(),
		})
	}

	return pbResults, nil
}
