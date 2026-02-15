from HttpRequestConfig import HttpMethod, HttpRequestConfig
from HttpRequest import HttpRequest
from typing import Any

class HttpRequestBuilder:
    def __init__(self, url: str):
        self._config = HttpRequestConfig(url=url)

    def set_url(self, url: str) -> "HttpRequestBuilder":
        self._config.url = url
        return self

    def set_method(self, method: HttpMethod) -> "HttpRequestBuilder":
        self._config.method = method
        return self

    def set_header(self, key: str, value: str) -> "HttpRequestBuilder":
        self._config.headers[key] = value
        return self

    def set_body(self, body: Any) -> "HttpRequestBuilder":
        self._config.body = body
        return self

    def set_timeout(self, seconds: int) -> "HttpRequestBuilder":
        self._config.timeout = seconds
        return self

    def set_retries(self, count: int) -> "HttpRequestBuilder":
        self._config.retries = count
        return self

    def build(self) -> HttpRequest:
        if not self._config.url:
            raise ValueError("URL is required")

        return HttpRequest(self._config)
