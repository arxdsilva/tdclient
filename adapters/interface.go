package adapters

import "io"

// Http represents the API querier abstraction
type Http interface {
	// todo: DoRequestWithContext
	DoRequest(method, path string, reader io.Reader, headers map[string]string) (resp []byte, status int, err error)
}
