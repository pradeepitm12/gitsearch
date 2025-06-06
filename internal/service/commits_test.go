package service

import (
	"context"
	"testing"

	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"

	"github.com/pradeepitm12/gitsearch/internal/test/mock"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
)

func TestSearch_Commits(t *testing.T) {
	mockClient := new(mock.MockClient)
	srv := NewServer(mockClient)

	mockCommitResults := []*github.CommitResult{
		{
			HTMLURL:    github.String("https://github.com/repo2/commit1"),
			Repository: &github.Repository{FullName: github.String("owner/repo2")},
		},
	}
	mockClient.On("SearchCommits", tmock.Anything, "fix bug user:golang").Return(mockCommitResults, nil)

	req := &pb.SearchRequest{
		Type:       "commits",
		SearchTerm: "fix bug",
		User:       "golang",
	}

	resp, err := srv.Search(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Results, 1)
	assert.Equal(t, "https://github.com/repo2/commit1", resp.Results[0].FileUrl)
	assert.Equal(t, "owner/repo2", resp.Results[0].Repo)

	mockClient.AssertExpectations(t)
}
