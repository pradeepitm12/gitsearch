package git

type CodeSearcher struct {
	Token string
}

func NewCodeSearcher(token string) *CodeSearcher {
	return &CodeSearcher{Token: token}
}

func (s *CodeSearcher) Search(term, user string, page, perPage int) ([]Result, error) {
	return genericSearch(
		s.Token,
		"code",
		"application/vnd.github+json",
		term,
		user,
		page,
		perPage,
		ExtractStandardResults,
	)
}
