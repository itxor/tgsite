build:
	go build -o ./bin ./cmd/subscriber/nats_subscribers.go 
	go build -o ./bin ./cmd/tgparser/main.go
