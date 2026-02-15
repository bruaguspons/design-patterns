package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HttpRequest struct {
	config HttpRequestConfig
}

func (r *HttpRequest) Send(result any) error {
	attempt := 0

	for {
		ctx, cancel := context.WithTimeout(context.Background(), r.config.Timeout)
		defer cancel()

		var bodyBytes []byte
		if r.config.Body != nil {
			var err error
			bodyBytes, err = json.Marshal(r.config.Body)
			if err != nil {
				return err
			}
		}

		req, err := http.NewRequestWithContext(
			ctx,
			string(r.config.Method),
			r.config.URL,
			bytes.NewBuffer(bodyBytes),
		)
		if err != nil {
			return err
		}

		for k, v := range r.config.Headers {
			req.Header.Set(k, v)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			if attempt < r.config.Retries {
				attempt++
				continue
			}
			if errors.Is(err, context.DeadlineExceeded) {
				return errors.New("request timeout")
			}
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			if attempt < r.config.Retries {
				attempt++
				continue
			}
			return fmt.Errorf("http error: %d", resp.StatusCode)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("http error: %d", resp.StatusCode)
		}

		return json.NewDecoder(resp.Body).Decode(result)
	}
}
