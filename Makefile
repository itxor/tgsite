hello:
	echo "Hello"

build:
	go build -o bin/main cmd/tgparser/main.go
	go build -o bin/subscribers cmd/subscriber/nats_subscribers.go

run-main:
	./bin/main

run-subs:
	./bin/subscribers