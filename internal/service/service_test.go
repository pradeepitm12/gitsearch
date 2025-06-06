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

func TestServer_Search_Code(t *testing.T) {
	ctx := context.Background()

	mockClient := new(mock.MockClient)

	mockClient.
		On("SearchCode", tmock.Anything, "grpc user:golang").
		Return([]*github.CodeResult{
			{
				HTMLURL: github.String("https://github.com/golang/go/blob/main/main.go"),
				Repository: &github.Repository{
					FullName: github.String("golang/go"),
				},
			},
		}, nil)

	s := NewServer(mockClient)

	resp, err := s.Search(ctx, &pb.SearchRequest{
		Type:       "code",
		SearchTerm: "grpc",
		User:       "golang",
	})

	assert.NoError(t, err)
	assert.Len(t, resp.Results, 1)
	assert.Equal(t, "https://github.com/golang/go/blob/main/main.go", resp.Results[0].FileUrl)
	assert.Equal(t, "golang/go", resp.Results[0].Repo)

	// Assert that expectations were met
	mockClient.AssertExpectations(t)
}
