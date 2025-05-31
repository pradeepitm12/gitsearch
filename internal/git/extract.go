package git

import (
	"encoding/json"
)

func ExtractStandardResults(data []byte) ([]Result, error) {
	var result struct {
		Items []struct {
			HTMLURL string `json:"html_url"`
			Repo    struct {
				FullName string `json:"full_name"`
			} `json:"repository"`
		}
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	var output []Result
	for _, item := range result.Items {
		output = append(output, Result{
			FileURL: item.HTMLURL,
			Repo:    item.Repo.FullName,
		})
	}
	return output, nil
}
