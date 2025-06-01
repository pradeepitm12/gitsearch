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

func TestSearch_Code(t *testing.T) {
	mockClient := new(mock.MockClient)
	srv := NewServer(mockClient)

	mockCodeResults := []*github.CodeResult{
		{
			HTMLURL:    github.String("https://github.com/repo1/file1.go"),
			Repository: &github.Repository{FullName: github.String("owner/repo1")},
		},
	}
	mockClient.On("SearchCode", tmock.Anything, "grpc user:golang").Return(mockCodeResults, nil)

	req := &pb.SearchRequest{
		Type:       "code",
		SearchTerm: "grpc",
		User:       "golang",
	}

	resp, err := srv.Search(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Results, 1)
	assert.Equal(t, "https://github.com/repo1/file1.go", resp.Results[0].FileUrl)
	assert.Equal(t, "owner/repo1", resp.Results[0].Repo)

	mockClient.AssertExpectations(t)
}
