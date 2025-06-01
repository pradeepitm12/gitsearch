FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gitsearch ./cmd/server

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/gitsearch /gitsearch

ENTRYPOINT ["/gitsearch"]
