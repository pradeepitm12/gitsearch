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

func TestSearch_Issues(t *testing.T) {
	mockClient := new(mock.MockClient)
	srv := NewServer(mockClient)

	mockIssueResults := []*github.Issue{
		{
			HTMLURL:    github.String("https://github.com/repo3/issue1"),
			Repository: &github.Repository{FullName: github.String("owner/repo3")},
		},
	}
	mockClient.On("SearchIssues", tmock.Anything, "memory leak user:golang").Return(mockIssueResults, nil)

	req := &pb.SearchRequest{
		Type:       "issues",
		SearchTerm: "memory leak",
		User:       "golang",
	}

	resp, err := srv.Search(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Results, 1)
	assert.Equal(t, "https://github.com/repo3/issue1", resp.Results[0].FileUrl)
	assert.Equal(t, "owner/repo3", resp.Results[0].Repo)

	mockClient.AssertExpectations(t)
}
