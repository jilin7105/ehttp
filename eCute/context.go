package eCute

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}
type C struct {
	ctx context.Context
	// request
	request *http.Request
	// response
	responsewriter http.ResponseWriter
	// params
	Params map[string]string

	// request info
	Path   string
	Method string
	// response info
	StatusCode int
	// middleware
	handlerFuncs []handlerFunc
	// index
	index int
}

func newContext(w http.ResponseWriter, r *http.Request) *C {
	return &C{
		request:        r,
		responsewriter: w,
		Params:         make(map[string]string),
		Path:           r.URL.Path,
		Method:         r.Method,
		StatusCode:     http.StatusOK,
		ctx:            context.Background(),
		index:          -1,
	}
}

func (c *C) PostForm(key string) string {
	return c.request.PostFormValue(key)
}

func (c *C) Query(key string) string {
	return c.request.URL.Query().Get(key)
}

func (c *C) SetHeader(key, value string) {
	c.responsewriter.Header().Set(key, value)
}

func (c *C) Status(code int) {
	c.StatusCode = code
}

func (c *C) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *C) JSON(code int, data interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.responsewriter)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.responsewriter, err.Error(), 500)
	}
}

func (c *C) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.responsewriter.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *C) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.responsewriter.Write([]byte(html))
}

func (c *C) Data(code int, data []byte) {
	c.Status(code)
	c.responsewriter.Write(data)
}

//洋葱模型
func (c *C) Next() {
	c.index++
	s := len(c.handlerFuncs)
	for ; c.index < s; c.index++ {
		c.handlerFuncs[c.index](c)
	}
}
