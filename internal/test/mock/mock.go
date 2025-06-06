package mock

import (
	"context"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) SearchCode(ctx context.Context, query string, page, perPage int) ([]*github.CodeResult, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*github.CodeResult), args.Error(1)
}

func (m *MockClient) SearchCommits(ctx context.Context, query string, page, perPage int) ([]*github.CommitResult, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*github.CommitResult), args.Error(1)
}

func (m *MockClient) SearchIssues(ctx context.Context, query string, page, perPage int) ([]*github.Issue, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*github.Issue), args.Error(1)
}

func (m *MockClient) SearchRepos(ctx context.Context, query string, page, perPage int) ([]*github.Repository, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*github.Repository), args.Error(1)
}

func (m *MockClient) SearchTopics(ctx context.Context, query string, page, perPage int) ([]*github.TopicResult, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*github.TopicResult), args.Error(1)
}

func (m *MockClient) SearchUsers(ctx context.Context, query string, page, perPage int) ([]*github.User, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*github.User), args.Error(1)
}
