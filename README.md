# GoO1 notification library

### Usage
- To install run `go get github.com/elvis-chuks/go01`

- Code example
```go
...
import "github.com/elvis-chuks/go01"
...

messages := []string{"hi", "ho"}
response := go01.NotifyClient("http://localhost:5000/api/v1/", messages)
fmt.Println(response)
```