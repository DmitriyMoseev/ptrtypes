ptrtypes - constructors of pointers to Go's builtin types
=========================================================


Example
-------

```go
package main

import (
    "fmt"

    "github.com/DmitriyMoseev/ptrtypes"
)

func main() {
    var ptr *int64
    ptr = ptrtypes.NewInt64(777)
    fmt.Println("ptr points to", *ptr)
}
```
