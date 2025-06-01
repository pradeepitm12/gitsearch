# gitsearch

gitsearch is a gRPC based search service used to search terms on github.

- You can search for code, commits, issues, repositories, topics, and users
- you can pass user filter with  "code" || "commits" || "issues".
- Example to search
```
grpcurl -plaintext -d '{"search_term":"jwt","user":"torvalds","type":"code"}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"fix bug","user":"golang","type":"commits"}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"memory leak","type":"issues"}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"grpc","type":"repositories"}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"devops","type":"topics"}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"torvalds","type":"users"}' localhost:50051 gitsearch.GithubSearchService/Search
```
This gRPC takes care of large result and implement pagination.
- Example to search with paginations and filter with users
```
grpcurl -plaintext -d '{"search_term":"jwt","user":"torvalds","type":"code","page":1,"per_page":10}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"fix bug","user":"golang","type":"commits","page":2,"per_page":20}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"memory leak","type":"issues","page":1,"per_page":15}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"grpc","type":"repositories","page":1,"per_page":5}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"devops","type":"topics","page":1,"per_page":5}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"torvalds","type":"users","page":1,"per_page":10}' localhost:50051 gitsearch.GithubSearchService/Search
```

### To build docker image 
```
docker build -t gitsearch:latest .
```

### To Run the image
```
docker run -p 50051:50051 -e GITHUB_TOKEN=$GITHUB_TOKEN gitsearch:latest
```