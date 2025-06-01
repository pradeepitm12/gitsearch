package git

import (
	"context"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

type GitClient interface {
	SearchCode(ctx context.Context, query string, page, perPage int) ([]*github.CodeResult, error)
	SearchIssues(ctx context.Context, query string, page, perPage int) ([]*github.Issue, error)
	SearchCommits(ctx context.Context, query string, page, perPage int) ([]*github.CommitResult, error)
	SearchRepos(ctx context.Context, query string, page, perPage int) ([]*github.Repository, error)
	SearchTopics(ctx context.Context, query string, page, perPage int) ([]*github.TopicResult, error)
	SearchUsers(ctx context.Context, query string, page, perPage int) ([]*github.User, error)
}

type Client struct {
	client *github.Client
}

func NewClient(token string) *Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(context.Background(), ts)

	return &Client{client: github.NewClient(tc)}
}

func (c *Client) SearchCode(ctx context.Context, query string, page, perPage int) ([]*github.CodeResult, error) {
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	res, _, err := c.client.Search.Code(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	return res.CodeResults, nil
}

func (c *Client) SearchIssues(ctx context.Context, query string, page, perPage int) ([]*github.Issue, error) {
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	res, _, err := c.client.Search.Issues(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	return res.Issues, nil
}

func (c *Client) SearchCommits(ctx context.Context, query string, page, perPage int) ([]*github.CommitResult, error) {
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	res, _, err := c.client.Search.Commits(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	return res.Commits, nil
}

func (c *Client) SearchRepos(ctx context.Context, query string, page, perPage int) ([]*github.Repository, error) {
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	res, _, err := c.client.Search.Repositories(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	return res.Repositories, nil
}

func (c *Client) SearchTopics(ctx context.Context, query string, page, perPage int) ([]*github.TopicResult, error) {
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	res, _, err := c.client.Search.Topics(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	return res.Topics, nil
}

func (c *Client) SearchUsers(ctx context.Context, query string, page, perPage int) ([]*github.User, error) {
	opts := &github.SearchOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: perPage,
		},
	}
	res, _, err := c.client.Search.Users(ctx, query, opts)
	if err != nil {
		return nil, err
	}
	return res.Users, nil
}
