package eCute

import (
	"log"
	"time"
)

func Logger() handlerFunc {
	return func(c *C) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.request.RequestURI, time.Since(t))
	}
}
