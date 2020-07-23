package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"

	_ "github.com/tattwei46/go-swagger/statik"
)

var swaggerPrefix = "/swaggerui/"

func swaggerMiddleware() gin.HandlerFunc {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	sh := http.StripPrefix(swaggerPrefix, staticServer)

	return func(c *gin.Context) {
		requestPath := c.Request.URL.Path
		if !strings.Contains(requestPath, swaggerPrefix) {
			c.Next()
		} else {
			sh.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
