FROM golang 1.22 alpine AS Builder
COPY go.mod go.sum
RUN cmd/main.go
