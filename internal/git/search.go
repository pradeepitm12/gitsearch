package git

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func genericSearch(token, endpoint, accept, term, user string, page, perPage int, extract func([]byte) ([]Result, error)) ([]Result, error) {
	q := url.QueryEscape(term)
	if user != "" {
		q += "+user:" + url.QueryEscape(user)
	}

	apiURL := fmt.Sprintf("https://api.github.com/search/%s?q=%s&page=%d&per_page=%d", endpoint, q, page, perPage)
	req, _ := http.NewRequest("GET", apiURL, nil)

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Accept", accept)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API error: %d %s", resp.StatusCode, resp.Status)
	}

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return extract(raw)
}
