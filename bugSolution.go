```go
package main

import "fmt"

func main() {
    var m map[string]int
    m = make(map[string]int) // Initialize the map
    m["key"] = 10
    fmt.Println(m["key"])

    //Alternative initialization
    m2 := map[string]int{"key": 20}
    fmt.Println(m2["key"])
}
```