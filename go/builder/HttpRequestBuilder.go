package main

import (
	"errors"
	"time"
)

type HttpRequestBuilder struct {
	config HttpRequestConfig
}

func NewHttpRequestBuilder(url string) *HttpRequestBuilder {
	return &HttpRequestBuilder{
		config: HttpRequestConfig{
			URL:     url,
			Method:  GET,
			Headers: make(map[string]string),
			Timeout: 5 * time.Second,
			Retries: 0,
		},
	}
}

func (b *HttpRequestBuilder) SetURL(url string) *HttpRequestBuilder {
	b.config.URL = url
	return b
}

func (b *HttpRequestBuilder) SetMethod(method HttpMethod) *HttpRequestBuilder {
	b.config.Method = method
	return b
}

func (b *HttpRequestBuilder) SetHeader(key, value string) *HttpRequestBuilder {
	b.config.Headers[key] = value
	return b
}

func (b *HttpRequestBuilder) SetBody(body any) *HttpRequestBuilder {
	b.config.Body = body
	return b
}

func (b *HttpRequestBuilder) SetTimeout(timeout time.Duration) *HttpRequestBuilder {
	b.config.Timeout = timeout
	return b
}

func (b *HttpRequestBuilder) SetRetries(retries int) *HttpRequestBuilder {
	b.config.Retries = retries
	return b
}

func (b *HttpRequestBuilder) Build() (*HttpRequest, error) {
	if b.config.URL == "" {
		return nil, errors.New("URL is required")
	}
	return &HttpRequest{config: b.config}, nil
}
