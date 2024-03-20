github.com/DennisUden/GoLib

# Usage #

## Initialize your module ##
```go mod init example.com/my-golib-demo```

## Get the GoLib module ##
Note that you need to include the v in the version tag \
```go get github.com/DennisUden/GoLib@v.0.1.4``` 

```go
package main

import (
    "fmt"
    "github.com/DennisUden/GoLib"
)

func main() {
    // a is the number which needs to be rounded and b the number of decimals
    fmt.Prinln(GoLib.Round(a, b)) 
}
```
