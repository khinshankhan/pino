package server

import (
	"fmt"
	"net/http"
)

type (
	Method   string
	Endpoint struct {
		method Method
		path   string
		action func(w http.ResponseWriter, r *http.Request)
	}
)

const (
	GET    Method = "GET"
	POST   Method = "POST"
	DELETE Method = "DELETE"
	PUT    Method = "PUT"
)

func (m Method) Check() error {
	valid := []Method{GET, POST, DELETE, PUT}
	for _, v := range valid {
		if m != v {
			return fmt.Errorf("invalid method %s used", string(m))
		}
	}
	return nil
}
