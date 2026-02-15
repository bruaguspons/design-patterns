from HttpRequestConfig import HttpRequestConfig
from typing import Any
import requests
import json

class HttpRequest:
    def __init__(self, config: HttpRequestConfig):
        self._config = config

    def send(self) -> Any:
        attempt = 0

        while True:
            try:
                response = requests.request(
                    method=self._config.method,
                    url=self._config.url,
                    headers=self._config.headers,
                    data=json.dumps(self._config.body)
                    if self._config.body is not None
                    else None,
                    timeout=self._config.timeout,
                )

                if not response.ok:
                    # Retry solo si es error 5xx
                    if 500 <= response.status_code < 600 and attempt < self._config.retries:
                        attempt += 1
                        continue
                    raise Exception(f"HTTP error: {response.status_code}")

                return response.json()

            except requests.Timeout:
                if attempt < self._config.retries:
                    attempt += 1
                    continue
                raise Exception("Request timeout")

            except Exception:
                if attempt < self._config.retries:
                    attempt += 1
                    continue
                raise
