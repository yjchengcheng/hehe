package myserver

import (
	"hello-go/mycontroller"
	"net/http"
	"reflect"
)

type MyServer struct {
	routers       map[string]mycontroller.ControllerInterface
	methodRouters []MethodRouter
}

type MethodRouter struct {
	path   string
	method string
	fn     string
	c      mycontroller.ControllerInterface
}

func (m *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, v := range m.methodRouters {
		if v.path == r.URL.String() && v.method == r.Method {

			v.c.Init(w, r)
			vc := reflect.ValueOf(v.c)
			method := vc.MethodByName(v.fn)
			in := []reflect.Value{}
			method.Call(in)

		}
	}

}

func (m *MyServer) Route(path string, c mycontroller.ControllerInterface) {

	if m.routers == nil {
		m.routers = make(map[string]mycontroller.ControllerInterface)
	}

	m.routers[path] = c
}

func (m *MyServer) RoutePath(path string, method string, fn string, c mycontroller.ControllerInterface) {
	if m.methodRouters == nil {
		m.methodRouters = make([]MethodRouter, 0)
	}

	m.methodRouters = append(m.methodRouters, MethodRouter{
		path:   path,
		method: method,
		fn:     fn,
		c:      c,
	})

}
