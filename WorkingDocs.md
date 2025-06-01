# To generate proto files
    protoc \
    --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/gitsearch/gitsearch.proto

# ToDo
    Write a gRPC client in Go
    Add logging 
    Add unit tests and mock GitHub API
    Better error handelling


# New Testing
    grpcurl -plaintext -d '{"search_term":"jwt","user":"torvalds","type":"code"}' localhost:50051 gitsearch.GithubSearchService/Search
    grpcurl -plaintext -d '{"search_term":"fix bug","user":"golang","type":"commits"}' localhost:50051 gitsearch.GithubSearchService/Search
    grpcurl -plaintext -d '{"search_term":"memory leak","type":"issues"}' localhost:50051 gitsearch.GithubSearchService/Search
    grpcurl -plaintext -d '{"search_term":"grpc","type":"repositories"}' localhost:50051 gitsearch.GithubSearchService/Search
    grpcurl -plaintext -d '{"search_term":"devops","type":"topics"}' localhost:50051 gitsearch.GithubSearchService/Search
    grpcurl -plaintext -d '{"search_term":"torvalds","type":"users"}' localhost:50051 gitsearch.GithubSearchService/Search

With pageination
grpcurl -plaintext -d '{"type": "code","search_term": "grpc","user": "grpc","page": 1,"per_page": 5}' localhost:50051 gitsearch.GithubSearchService/Search


# ToDo,
paging          -- Done
testing         -- Done
user filter     -- Done
logging 
error handeling
Document


grpcurl -plaintext -d '{"search_term":"jwt","user":"torvalds","type":"code","page":1,"per_page":10}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"fix bug","user":"golang","type":"commits","page":2,"per_page":20}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"memory leak","type":"issues","page":1,"per_page":15}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"grpc","type":"repositories","page":1,"per_page":5}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"devops","type":"topics","page":1,"per_page":5}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"torvalds","type":"users","page":1,"per_page":10}' localhost:50051 gitsearch.GithubSearchService/Search


grpcurl -plaintext -d '{"search_term":"jwt","user":"torvalds","type":"code","page":1,"per_page":10}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"fix bug","user":"golang","type":"commits","page":2,"per_page":20}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"memory leak","type":"issues","page":1,"per_page":15}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"grpc","type":"repositories","page":1,"per_page":5}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"devops","type":"topics","page":1,"per_page":5}' localhost:50051 gitsearch.GithubSearchService/Search
grpcurl -plaintext -d '{"search_term":"torvalds","type":"users","page":1,"per_page":10}' localhost:50051 gitsearch.GithubSearchService/Search
