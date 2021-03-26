# simple-request
Simple lib to do a simple request in a simple way

#How to Use
Import the lib and call the request as simple as it need to be
```Go
package main
import (
    "fmt"
    "github.com/zzzep/simplerequest"
)
func main() {
    c, b, e := simplerequest.Get("http://example.com")
    if e != nil {
        fmt.Println("StatusCode:", c, "Error:", e)
    }
    fmt.Println("StatusCode:", c, "Payload Response:", b)
}
```
