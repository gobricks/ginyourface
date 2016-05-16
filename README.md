[![Build Status](https://travis-ci.org/gobricks/ginyourface.svg?branch=master)](https://travis-ci.org/gobricks/ginyourface)
[![Go Report Card](https://goreportcard.com/badge/github.com/gobricks/ginyourface)](https://goreportcard.com/report/github.com/gobricks/ginyourface)


# ginyourface

[Gin](https://github.com/gin-gonic/gin) middleware for [facecontrol](https://github.com/gobricks/facecontrol).

# Basic example

Create file `main.go` and paste the following code into it:

``` go
    package main

    import (
        "github.com/gin-gonic/gin"

        "github.com/gobricks/ginyourface"
    )

    func main() {
        r := gin.New()

        // now every request will be validated through Facecontrol service
        // and userPayload will be returned if user has valid token
        r.Use(ginyourface.Facecontrol())
        
        r.GET("/personal", func(c *gin.Context) {
            userData := c.MustGet("userPayload").(interface{})
            c.String(http.StatusOK, "Hello %s", userData["username"])
        })

        r.Run(":8080")
    }
```

# Build and run

```
$ go build main.go
$ FC_HOST="https://facecontrol.mysite.com" FC_LOGIN_PAGE="https://login.mysite.com" FC_SESSION_COOKIE="sessid" ./main
```