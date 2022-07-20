# github.com/willsussman/libwtf
### Tutorial: github.com/mitchallen/go-lib

## Usage

### Initialize your module

```
$ go mod init example.com/my-libwtf-demo
```

### Get the libwtf module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/willsussman/libwtf@v0.1.19
```

```go
package main

import (
    wtf "github.com/willsussman/libwtf"
    "fmt"
)

func main() {

    // act as an emitter

    attributes := []wtf.Attribute{
        // {Key: "firstname", Value: "Hari"},
        // {Key: "lastname", Value: "Balakrishnan"},
    }

    record := wtf.MakeRecord(255, attributes)
    ok := wtf.Emit(record)
    if ok != 0 {
        fmt.Println("Emission failed")
    }

    // act as a collector/analyzer

    dag := wtf.WheresTheFault(attributes)

    fmt.Printf("dag=%+v\n", dag)

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


