package service

import (
	"context"
	"github.com/pradeepitm12/gitsearch/internal/git"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
)

type RepositoriesStrategy struct{}

func (c *RepositoriesStrategy) Search(ctx context.Context, client git.GitClient, query string, page, perPage int) ([]*pb.Result, error) {
	results, err := client.SearchRepos(ctx, query, page, perPage)
	if err != nil {
		return nil, err
	}

	var pbResults []*pb.Result
	for _, r := range results {

		repoName := ""
		if r.FullName != nil {
			repoName = *r.FullName
		} else if r.Name != nil {
			repoName = *r.Name
		}

		pbResults = append(pbResults, &pb.Result{
			FileUrl: r.GetHTMLURL(),
			Repo:    repoName,
		})
	}
	return pbResults, nil
}
