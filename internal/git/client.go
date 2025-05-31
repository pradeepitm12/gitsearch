package git

import (
	"fmt"
)

const (
	CODE    = "code"
	COMMITS = "commits"
	ISSUES  = "issues"
	REPO    = "repositories"
	TOPICS  = "topics"
	USERS   = "users"
)

type Searcher interface {
	Search(term string, user string, page, perPage int) ([]Result, error)
}

type Result struct {
	FileURL string
	Repo    string
}

func NewSearcher(searchType, token string) (Searcher, error) {
	switch searchType {
	case CODE:
		return NewCodeSearcher(token), nil
	case COMMITS:
		return NewCommitsSearcher(token), nil
	case ISSUES:
		return NewIssuesSearcher(token), nil
	case REPO:
		return NewRepoSearcher(token), nil
	case TOPICS:
		return NewTopicsSearcher(token), nil
	case USERS:
		return NewUsersSearcher(token), nil
	default:
		return nil, fmt.Errorf("unsupported search type: %s", searchType)
	}
}
