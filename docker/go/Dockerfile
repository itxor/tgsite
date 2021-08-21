FROM golang:1.17-alpine

ENTRYPOINT go clean -modcache && go build cmd/main.go && tail -f /dev/null