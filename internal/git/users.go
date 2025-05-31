package git

type UsersSearcher struct {
	Token string
}

func NewUsersSearcher(token string) *UsersSearcher {
	return &UsersSearcher{Token: token}
}

func (s *UsersSearcher) Search(term, user string, page, perPage int) ([]Result, error) {
	return genericSearch(
		s.Token,
		"users",
		"application/vnd.github+json",
		term,
		user,
		page,
		perPage,
		ExtractStandardResults,
	)
}
