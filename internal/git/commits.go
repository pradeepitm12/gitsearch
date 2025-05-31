package git

type CommitsSearcher struct {
	Token string
}

func NewCommitsSearcher(token string) *CommitsSearcher {
	return &CommitsSearcher{Token: token}
}

func (s *CommitsSearcher) Search(term, user string) ([]Result, error) {
	return genericSearch(
		s.Token,
		"commits",
		"application/vnd.github+json",
		term,
		user,
		ExtractStandardResults,
	)
}
