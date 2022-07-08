# github.com/willsussman/libwtf
## Original: github.com/mitchallen/go-lib

## Usage

### Initialize your module

```
$ go mod init example.com/my-libwtf-demo
```

### Get the libwtf module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/willsussman/libwtf@v0.1.0
```

```go
package main

import (
    "fmt"

    "github.com/willsussman/libwtf"
)

func main() {
    fmt.Println(lib.Add(2,3))
    fmt.Println(lib.Subtract(2,3))
}
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v0.1.0
$ git push origin --tags
```


