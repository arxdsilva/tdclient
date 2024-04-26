package adapters

import (
	"context"
	"io"
)

// Http represents the API querier abstraction
type Http interface {
	DoRequest(ctx context.Context, method, path string, reader io.Reader, headers map[string]string) (resp []byte, status int, err error)
}
