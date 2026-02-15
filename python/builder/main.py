from HttpRequestBuilder import HttpRequestBuilder

if __name__ == "__main__":
    request = (
        HttpRequestBuilder("https://api.example.com/users")
        .set_method("POST")
        .set_header("Content-Type", "application/json")
        .set_body({"name": "Juan"})
        .set_timeout(3)   # 3 segundos
        .set_retries(2)   # hasta 2 reintentos
        .build()
    )

    data = request.send()
    print(data)
