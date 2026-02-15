import { HttpRequestBuilder } from "./RequestBuilder";

const request = new HttpRequestBuilder("https://api.example.com/users")
  .setMethod("POST")
  .setHeader("Content-Type", "application/json")
  .setBody({ name: "Juan" })
  .setTimeout(3000)   // 3 segundos
  .setRetries(2)      // hasta 2 reintentos
  .build();

const data = await request.send<{ id: number; name: string }>();
