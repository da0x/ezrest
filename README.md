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
code, err := ezrest.Get("https://example.com/json", ezrest.DefaultHeaders(), &o)
if err != nil {
    log.Fatalln(err)
}
println("response code:", code)
println("response object:", o)
```
## Post
Example:
```
var request struct {
    Name string `json:"name"`
}
var response struct {
    Id int `json:"id"`
    Name string `json:"name"`
}
code, err := ezrest.Post("https://example.com/json", ezrest.DefaultHeaders(), request, &response)
if err != nil {
    log.Fatalln(err)
}
println("response code:", code)
println("response object:", response)
```
## Request Headers
Request headers have to be explicitly defined. You can use ezrest.DefaultHeaders as shown above or
you can add your own. Here's an example:
```
headers := ezrest.DefaultHeaders()
headers["new-header"] = "value"
code, err := ezrest.Get("https://example.com/json", headers, &o)
if err != nil {
    log.Fatalln(err)
}
```
## Maintainer
Daher Alfawares
