package git

type RepoSearcher struct {
	Token string
}

func NewRepoSearcher(token string) *RepoSearcher {
	return &RepoSearcher{Token: token}
}

func (s *RepoSearcher) Search(term, user string) ([]Result, error) {
	return genericSearch(
		s.Token,
		"repositories",
		"application/vnd.github+json",
		term,
		user,
		ExtractStandardResults,
	)
}
