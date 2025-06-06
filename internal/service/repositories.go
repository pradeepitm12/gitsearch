package service

import (
	"cmp"
	"context"

	"github.com/pradeepitm12/gitsearch/internal/git"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
)

type repositoriesStrategy struct{}

func (c *repositoriesStrategy) Search(ctx context.Context, client git.GitClient, query string, page, perPage int) ([]*pb.Result, error) {
	results, err := client.SearchRepos(ctx, query, page, perPage)
	if err != nil {
		return nil, err
	}

	var pbResults []*pb.Result
	for _, r := range results {
		repoName := cmp.Or(r.GetFullName(), r.GetName())
		pbResults = append(pbResults, &pb.Result{
			FileUrl: r.GetHTMLURL(),
			Repo:    repoName,
		})
	}
	return pbResults, nil
}
