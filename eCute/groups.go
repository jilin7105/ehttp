package eCute

type Middleware func()

type GroupRouter struct {
	eh          *Eh
	prefix      string
	middlewares []handlerFunc
}

func (group *GroupRouter) Group(prefix string) *GroupRouter {
	engine := group.eh
	newGroup := &GroupRouter{
		prefix: group.prefix + prefix,
		eh:     engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (groupRouter *GroupRouter) Use(middleware ...handlerFunc) {
	groupRouter.middlewares = append(groupRouter.middlewares, middleware...)
}

func (groupRouter *GroupRouter) GET(pattern string, h handlerFunc) {
	groupRouter.addRouter("GET", pattern, h)
}

func (groupRouter *GroupRouter) POST(pattern string, h handlerFunc) {

	groupRouter.addRouter("POST", pattern, h)
}

func (groupRouter *GroupRouter) addRouter(method, pattern string, h handlerFunc) {
	pattern = groupRouter.prefix + pattern
	groupRouter.eh.router.addRouter(method, pattern, h)
}
