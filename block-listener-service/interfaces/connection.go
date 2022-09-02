package api

import "context"

type Connection interface {
	Connect(context.Context) error
	Close() error
	GetClient() any
}
