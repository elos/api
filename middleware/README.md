api/middleware
--------------

Package `middleware` provides the custom middleware used by the Elos HTTP API.

### What is Middleware (For the Beginner)

Middleware is how we refer to the logic to which we subject requests in between their acceptance and response.

Examples are easiest. Consider that we want to write `Hello World` in text to a HTTP client in response to their
`GET` request to the `/hello` endpoint. In go, we write a handler, and also a mutexer for handling the path. But the
handler would be like:

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World"))
}
```

So now we have a function than can handle a HTTP request from a client, and actually write a response. But perhaps
we also want to add logging to this funciton. We want to know who is calling what when. Well we could just add logging:

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
    log.Print("helloHandler was hit")
    w.Write([]byte("Hello World"))
}
```

But now if we want to write a goodbye handler, we duplicate the logging logic.
```go
func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
    log.Print("goodbye was hit")
    w.Write([]byte("Goodbye"))
}
```

The logging done by different engineers implementing different endpoints may not be consistent. We want consistent well-formatted information across all the handlers. Also logging actually has nothing fundamentally to do with handling the requests. So we break it out

So we have middleware. Handlers that see the request before the main one. Perhaps a `LogRequest` middleware that we wrap every handler in so that it logs that the `/hello` and `/goodbye` routes are getting hit.

