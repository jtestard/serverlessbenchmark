package function

import (
	"fmt"
	"net/http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	msg := fmt.Sprintf("Hello world, input was: %s", string(req.Body))

	return handler.Response{
		Body:       []byte(msg),
		StatusCode: http.StatusOK,
	}, err
}
