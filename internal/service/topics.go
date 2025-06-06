package service

import (
	"context"
	"fmt"

	"github.com/pradeepitm12/gitsearch/internal/git"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
)

type topicsStrategy struct{}

func (c *topicsStrategy) Search(ctx context.Context, client git.GitClient, query string, page, perPage int) ([]*pb.Result, error) {
	results, err := client.SearchTopics(ctx, query, page, perPage)
	if err != nil {
		return nil, err
	}

	var pbResults []*pb.Result
	for _, r := range results {
		pbResults = append(pbResults, &pb.Result{
			FileUrl: fmt.Sprintf("https://github.com/topics/%s", r.GetName()),
			Repo:    "n/a",
		})
	}

	return pbResults, nil
}
