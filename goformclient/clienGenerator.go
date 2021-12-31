package goformclient

import (
        "net/http"
        //"fmt"
        "time"
)

type clientGenerator struct {

	 headers http.Header
	 maxIdleConns int
         connectionTimeout  time.Duration
         responseTimeout  time.Duration
         disableTimeouts bool

}

type ClientGenerator  interface {
	 /*Http Header*/
        SetHeaders(headers http.Header) ClientGenerator
        /*Timoute Configuration*/
        DisableTimeouts(disableTimeouts bool) ClientGenerator
        SetConnectionTimeout(timeout time.Duration) ClientGenerator
        SetRequestTimeout(timeout time.Duration) ClientGenerator
        SetMaxIdleConns(i int)ClientGenerator
	Generate() Client
}

func NewGenerator() ClientGenerator {
        generator := &clientGenerator{}
        return generator
}

func (c *clientGenerator) Generate() Client {
	client := httpClient {
		headers:             c.headers,
		maxIdleConns:	     c.maxIdleConns,
		connectionTimeout:   c.connectionTimeout,
		responseTimeout:     c.responseTimeout,
		disableTimeouts:     c.disableTimeouts,

	}

	return &client

}
func (c *clientGenerator) SetHeaders(headers http.Header) ClientGenerator {
     c.headers = headers
     return c
}

func (c *clientGenerator) SetConnectionTimeout(timeout time.Duration) ClientGenerator {
     c.connectionTimeout = timeout
     return c
}

func (c *clientGenerator) SetRequestTimeout(timeout time.Duration) ClientGenerator {
     c.responseTimeout = timeout
     return c
}

func (c *clientGenerator) SetMaxIdleConns(i int) ClientGenerator{
     c.maxIdleConns = i
     return c
}
func (c *clientGenerator)DisableTimeouts(disable bool) ClientGenerator {

     c.disableTimeouts = disable
     return c
}

