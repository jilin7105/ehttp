# ehttp


```shell
#使用
go get github.com/jilin7105/ehttp
```

```go
package main

import (
	"github.com/jilin7105/ehttp/eCute"
	"log"
	"net/http"
)

func main() {
    //初始化  eCute.New() 两者的区别 , Default 自带 请求信息 和 错误处理中间件
	eh := eCute.Default()
	//注册路由
	eh.GET("/hello", hello)   
	eh.GET("/hello/you", helloYou)
	eh.GET("/hai/*about", hai) //通配符 * 仅能存在于路由最后
	eh.GET("/test/:hai/about", hait)  //通配符 : 可以在路由中使用
	eh.GET("/test/:hai/test", hait)

	v1 := eh.Group("/v1") //支持分组功能
	v1.Use(func(c *eCute.C) {  //支持中间件功能
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
	
	err := eh.Run(":9090") //启动服务
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


```