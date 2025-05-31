package git

type IssuesSearcher struct {
	Token string
}

func NewIssuesSearcher(token string) *IssuesSearcher {
	return &IssuesSearcher{Token: token}
}

func (s *IssuesSearcher) Search(term, user string, page, perPage int) ([]Result, error) {
	return genericSearch(
		s.Token,
		"issues",
		"application/vnd.github+json",
		term,
		user,
		page,
		perPage,
		ExtractStandardResults,
	)
}
