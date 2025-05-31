package git

type CodeSearcher struct {
	Token string
}

func NewCodeSearcher(token string) *CodeSearcher {
	return &CodeSearcher{Token: token}
}

func (s *CodeSearcher) Search(term, user string) ([]Result, error) {
	return genericSearch(
		s.Token,
		"code",
		"application/vnd.github+json",
		term,
		user,
		ExtractStandardResults,
	)
}
