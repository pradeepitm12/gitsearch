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

func TestSearch_Topics(t *testing.T) {
	mockClient := new(mock.MockClient)
	srv := NewServer(mockClient)

	mockTopicResults := []*github.TopicResult{
		{
			Name: github.String("devops"),
		},
	}
	mockClient.On("SearchTopics", tmock.Anything, "devops").Return(mockTopicResults, nil)

	req := &pb.SearchRequest{
		Type:       "topics",
		SearchTerm: "devops",
	}

	resp, err := srv.Search(context.Background(), req)
	assert.NoError(t, err)
	assert.Len(t, resp.Results, 1)
	assert.Equal(t, "https://github.com/topics/devops", resp.Results[0].FileUrl)
	assert.Equal(t, "n/a", resp.Results[0].Repo)

	mockClient.AssertExpectations(t)
}
