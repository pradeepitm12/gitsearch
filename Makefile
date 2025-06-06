PROTO_DIR=proto/gitsearch
PROTO_FILES=$(PROTO_DIR)/gitsearch.proto
OUT_DIR=$(PROTO_DIR)

GOPKG=github.com/pradeepitm12/gitsearch/proto/gitsearch

.PHONY: proto-doc

proto-doc:
	protoc --doc_out=docs \
	       --doc_opt=html,index.html \
	       proto/gitsearch/gitsearch.proto



.PHONY: proto

proto:
	protoc --go_out=. --go-grpc_out=. \
	       --go_opt=paths=source_relative \
	       --go-grpc_opt=paths=source_relative \
	       $(PROTO_FILES)

.PHONY:  build

build:
	go build -o bin/server ./cmd/server

.PHONY: clean

clean:
	rm -rf bin/

.PHONY: test

test:
	go test -v ./...
