package git

import "encoding/json"

type TopicsSearcher struct {
	Token string
}

func NewTopicsSearcher(token string) *TopicsSearcher {
	return &TopicsSearcher{Token: token}
}

func (s *TopicsSearcher) Search(term, user string, page, perPage int) ([]Result, error) {
	return genericSearch(
		s.Token,
		"topics",
		"application/vnd.github+json",
		term,
		user,
		page,
		perPage,
		extractTopics,
	)
}

func extractTopics(data []byte) ([]Result, error) {
	var raw struct {
		Items []struct {
			Name string `json:"name"` // e.g. "devops"
			URL  string `json:"url"`  // e.g. "https://api.github.com/topics/devops"
		} `json:"items"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	var results []Result
	for _, item := range raw.Items {
		results = append(results, Result{
			FileURL: "https://github.com/topics/" + item.Name,
			Repo:    item.Name,
		})
	}

	return results, nil
}
