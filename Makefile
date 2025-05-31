PROTO_DIR=proto/gitsearch
PROTO_FILES=$(PROTO_DIR)/gitsearch.proto
OUT_DIR=$(PROTO_DIR)

GOPKG=github.com/pradeepitm12/gitsearch/proto/gitsearch

.PHONY: all build proto clean

all: proto build

proto:
	protoc --go_out=. --go-grpc_out=. \
	       --go_opt=paths=source_relative \
	       --go-grpc_opt=paths=source_relative \
	       $(PROTO_FILES)

build:
	go build -o bin/server ./cmd/server

clean:
	rm -rf $(OUT_DIR)/*.pb.go bin/
