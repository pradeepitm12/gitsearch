package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	pb "github.com/pradeepitm12/gitsearch/proto/gitsearch"
)

type GitSearchServer struct {
	pb.UnimplementedGithubSearchServiceServer
	client *http.Client
	token  string
}

func NewGitSearchServer() *GitSearchServer {
	return &GitSearchServer{
		client: &http.Client{},
		token:  os.Getenv("GITHUB_TOKEN"),
	}
}

func (s *GitSearchServer) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	q := req.GetSearchTerm()
	if req.GetUser() != "" {
		q += fmt.Sprintf("+user:%s", req.GetUser())
	}
	url := fmt.Sprintf("https://api.github.com/search/code?q=%s", q)

	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Accept", "application/vnd.github.v3+json")

	if s.token != "" {
		httpReq.Header.Set("Authorization", "token "+s.token)
	} else {
		fmt.Println("No token set, sending unauthenticated request.")
	}

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	var body struct {
		Items []struct {
			HTMLURL    string `json:"html_url"`
			Repository struct {
				FullName string `json:"full_name"`
			} `json:"repository"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	results := []*pb.Result{}
	for _, item := range body.Items {
		results = append(results, &pb.Result{
			FileUrl: item.HTMLURL,
			Repo:    item.Repository.FullName,
		})
	}

	return &pb.SearchResponse{Results: results}, nil
}
