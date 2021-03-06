package goformclient

import (
	"net/http"
	"encoding/json"
)

type Response struct{
	status string
	statusCode int
	headers http.Header
	body []byte
}

func (r *Response) Status() string {
	return r.status
}
func (r *Response) StatusCode() int {
	return r.statusCode
}
func (r *Response) Header() http.Header {
	return r.headers
}
func (r *Response) Bytes() []byte {
	return r.body
}
func (r *Response) String() string {
	return string(r.status)
}
func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
