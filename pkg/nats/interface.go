package nats

import "github.com/nats-io/nats.go"

type NatsClientInterface interface {
	Connect() (func(), error)
	Publish(subject string, value interface{}) error
	GetConnect() *nats.EncodedConn
}
