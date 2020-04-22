package mycontroller

import (
	"net/http"
)

type BaseController struct {
	W    http.ResponseWriter
	R    *http.Request
	Data map[string]interface{}
}

func (b *BaseController) Init(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			b.W.WriteHeader(500)
			b.W.Write([]byte("HTTP 500"))
		}
	}()

	b.W = w
	b.R = r

	b.Data = make(map[string]interface{})

}

func (b *BaseController) Get() {
	panic("err")
	b.W.Write([]byte(" GET ----  fuck you"))
}

func (b *BaseController) Post() {
	b.W.Write([]byte(" POST ----  fuck you"))
}
