package service

import (
	"context"
	"github.com/google/go-github/v50/github"
	"github.com/pradeepitm12/gitsearch/internal/test/mock"
	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
	"github.com/stretchr/testify/assert"
	tmock "github.com/stretchr/testify/mock"
	"testing"
)

func TestSearch_Repositories(t *testing.T) {
	mockClient := new(mock.MockClient)
	srv := NewServer(mockClient)

	mockRepoResults := []*github.Repository{
		{
			HTMLURL:  github.String("https://github.com/repo4"),
			FullName: github.String("owner/repo4"),
		},
	}
	mockClient.On("SearchRepos", tmock.Anything, "grpc").Return(mockRepoResults, nil)

	req := &pb.SearchRequest{
		Type:       "repositories",
		SearchTerm: "grpc",
		User:       "golang", // This should trigger validation error, so test below will catch that
	}

	_, err := srv.Search(context.Background(), req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user filter is not supported")

	req.User = ""

	resp, err := srv.Search(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Results, 1)
	assert.Equal(t, "https://github.com/repo4", resp.Results[0].FileUrl)
	assert.Equal(t, "owner/repo4", resp.Results[0].Repo)

	mockClient.AssertExpectations(t)
}
