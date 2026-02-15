import { HttpRequestConfig } from "./types.d";

export class HttpRequest {
  constructor(private readonly config: HttpRequestConfig) {}

  private async fetchWithTimeout(controller: AbortController): Promise<Response> {
    const timeoutId = setTimeout(() => {
      controller.abort();
    }, this.config.timeout);

    try {
      return await fetch(this.config.url, {
        method: this.config.method,
        headers: this.config.headers,
        body: this.config.body
          ? JSON.stringify(this.config.body)
          : undefined,
        signal: controller.signal,
      });
    } finally {
      clearTimeout(timeoutId);
    }
  }

  public async send<T = unknown>(): Promise<T> {
    const maxRetries = this.config.retries ?? 0;
    let attempt = 0;

    while (true) {
      const controller = new AbortController();

      try {
        const response = await this.fetchWithTimeout(controller);

        if (!response.ok) {
          // Retry solo si es error 5xx
          if (response.status >= 500 && attempt < maxRetries) {
            attempt++;
            continue;
          }
          throw new Error(`HTTP error: ${response.status}`);
        }

        return response.json() as Promise<T>;
      } catch (error) {
        if (attempt < maxRetries) {
          attempt++;
          continue;
        }

        if (error instanceof DOMException && error.name === "AbortError") {
          throw new Error("Request timeout");
        }

        throw error;
      }
    }
  }
}
