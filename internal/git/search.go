package git

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func genericSearch(token, endpoint, accept, term, user string, extract func([]byte) ([]Result, error)) ([]Result, error) {
	q := url.QueryEscape(term)
	if user != "" {
		q += "+user:" + user
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/search/%s?q=%s", endpoint, q), nil)
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
