package nats

import "github.com/nats-io/nats.go"

type NatsClientInterface interface {
	Connect() (func(), error)
	GetConnect() *nats.EncodedConn
}
