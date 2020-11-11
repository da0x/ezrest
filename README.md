# ezrest
Welcome to this simple and minimalist rest client for golang. This package aims
to give a minimal interface for calling RESTful services. Some headers are applied
by default which can be modified.
## Installation
To install the library simply run:
```
$ go get github.com/da0x/ezrest
```
To import the library, use:
```
import "github.com/da0x/ezrest"
```
## Get
Example:
```
var o struct {
    UserId int `json:"user_id"`
    Name string `json:"name"`
}
code, err := ezrest.Get("https://example.com/json", &o)
if err != nil {
    log.Fatalln(err)
}
println("response code:", code)
println("response object:", o)
```
## Request Headers
By default, this library will send default headers to indicate that the request content
type is json and to close the connection.
You can reset or clear out the request headers by directly manipulating or replacing ezrest.RequestHeaders.
Example
```
headers := ezrest.DefaultHeaders()
headers["new-header"] = "value"
ezrest.RequestHeaders = headers
...
_, err := ezrest.Get("...", &o)
```
## Maintainer
Justin Gehr
