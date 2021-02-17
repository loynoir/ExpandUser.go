# ExpandUser.go 
python os.path.expanduser equivalent in golang

# Install
```sh
go get -u github.com/loynoir/ExpandUser.go
```

# Example

```go

package main

import (
    "fmt"
    "github.com/loynoir/ExpandUser.go"
)

func main() {
    x, _ := ExpandUser.ExpandUser("~/~/foo/bar/")
    fmt.Println(x)
}
```
