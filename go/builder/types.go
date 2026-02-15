package main

import (
	"time"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	PATCH  HttpMethod = "PATCH"
	DELETE HttpMethod = "DELETE"
)

type HttpRequestConfig struct {
	URL     string
	Method  HttpMethod
	Headers map[string]string
	Body    any
	Timeout time.Duration
	Retries int
}
