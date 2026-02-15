export type HttpMethod = "GET" | "POST" | "PUT" | "PATCH" | "DELETE";

export interface HttpRequestConfig {
  url: string;
  method: HttpMethod;
  headers: Record<string, string>;
  body?: unknown;
  timeout?: number;
  retries?: number;
}
