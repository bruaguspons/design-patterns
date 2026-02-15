import { HttpRequest } from "./HttpRequest";
import { HttpMethod, HttpRequestConfig } from "./types.d";

export class HttpRequestBuilder {
  private config: Partial<HttpRequestConfig> = {
    headers: {},
  };

  constructor(url: string) {
    this.config.url = url;
  }

  public setUrl(url: string): this {
    this.config.url = url;
    return this;
  }

  public setMethod(method: HttpMethod): this {
    this.config.method = method;
    return this;
  }

  public setHeader(key: string, value: string): this {
    this.config.headers = {
      ...this.config.headers,
      [key]: value,
    };
    return this;
  }

  public setBody(body: unknown): this {
    this.config.body = body;
    return this;
  }

  public setTimeout(ms: number): this {
    this.config.timeout = ms;
    return this;
  }

  public setRetries(count: number): this {
    this.config.retries = count;
    return this;
  }

  public build(): HttpRequest {
    if (!this.config.url) {
      throw new Error("URL is required");
    }

    return new HttpRequest({
      url: this.config.url,
      method: this.config.method ?? "GET",
      headers: this.config.headers ?? {},
      body: this.config.body,
      timeout: this.config.timeout ?? 5000,
      retries: this.config.retries ?? 0,
    });
  }
}
