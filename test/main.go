package main

import (
	"ebasehttp/eCute"
	"log"
	"net/http"
)

func main() {

	eh := eCute.Default()
	eh.GET("/hello", hello)
	eh.GET("/hello/you", helloYou)
	eh.GET("/hai/*about", hai)
	eh.GET("/test/:hai/about", hait)
	eh.GET("/test/:hai/test", hait)

	v1 := eh.Group("/v1")
	v1.Use(func(c *eCute.C) {
		log.Println("v1.mid")
	})

	v1.GET("/test/a", func(c *eCute.C) {
		c.JSON(http.StatusOK, eCute.H{
			"a": 2,
		})
	})
	v1.GET("/test/b", func(c *eCute.C) {
		c.String(http.StatusOK, "2222", nil)
	})
	err := eh.Run(":9090")
	if err != nil {
		return
	}
	return
}

func hello(ctx *eCute.C) {
	ctx.String(200, "Hello, World!")
}
func helloYou(ctx *eCute.C) {

	ctx.JSON(http.StatusOK, eCute.H{
		"name": ctx.Query("name"),
		"age":  ctx.Query("age"),
	})
}

func hai(ctx *eCute.C) {

	ctx.JSON(http.StatusOK, eCute.H{
		"about": ctx.Param("about"),
	})
}
func hait(ctx *eCute.C) {

	ctx.JSON(http.StatusOK, eCute.H{
		"hai": ctx.Param("hai"),
	})
}
