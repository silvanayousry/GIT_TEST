# GIT_TEST
exploring git and github
learning basics of GIt
## learning basics of html 
making a simple website to learn more 
<<<<<<< HEAD
* tutorial 
* learning
* html

`` sudo apt update ``


# checkeg
server health check implementation in golang (go)

### Installing 
```console
    go get -u  github.com/egirna/checkeg                                 
```

### Usage
**Import the package**
```go
    import "github.com/egirna/checkeg"                                    
 ```


## URL Health Check

#example
``` go 
 package main
 import (
	"fmt"
	"time"
	
	"github.com/egirna/checkeg"
 )

 func main() {
	
	fmt.Println(checkeg.CheckUrL("https://golangbot.com/go-packages/", []int{200, 300}))

```



## TCP Dial Check
