# simple-request
Simple lib to do a simple request in a simple way

#How to Use
Import the lib and call the request as simple as it need to be
```Go
package main
import (
    "fmt"
    sr "github.com/zzzep/simplerequest"
)
func main() {
    c, b := sr.Get("http://example.com")
    fmt.Println("StatusCode:", c, "Payload Response:", b)
}
```
