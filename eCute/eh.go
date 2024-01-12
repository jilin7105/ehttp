package eCute

import (
	"net/http"
	"strings"
)

type handlerFunc func(*C)

type Eh struct {
	router *Router
	*GroupRouter
	groups []*GroupRouter
}

func Default() *Eh {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}
func New() *Eh {
	engine := &Eh{router: NewRouter()}
	engine.GroupRouter = &GroupRouter{eh: engine}

	engine.groups = []*GroupRouter{engine.GroupRouter}
	return engine
}

func (eh *Eh) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []handlerFunc
	//增加中间件
	for _, group := range eh.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlerFuncs = middlewares
	eh.router.handle(c)
}

func (eh *Eh) Run(addr string) error {
	return http.ListenAndServe(addr, eh)
}
