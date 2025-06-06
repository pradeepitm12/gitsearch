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

func TestSearch_Users(t *testing.T) {
	mockClient := new(mock.MockClient)
	srv := NewServer(mockClient)

	mockUserResults := []*github.User{
		{
			HTMLURL: github.String("https://github.com/torvalds"),
			Login:   github.String("torvalds"),
		},
	}
	mockClient.On("SearchUsers", tmock.Anything, "torvalds").Return(mockUserResults, nil)

	req := &pb.SearchRequest{
		Type:       "users",
		SearchTerm: "torvalds",
		User:       "", // no user filter allowed
	}

	resp, err := srv.Search(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Results, 1)
	assert.Equal(t, "https://github.com/torvalds", resp.Results[0].FileUrl)
	assert.Equal(t, "torvalds", resp.Results[0].Repo) // users have no repo

	mockClient.AssertExpectations(t)
}
