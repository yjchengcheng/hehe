package mycontroller

import "net/http"

type ControllerInterface interface {
	Init(w http.ResponseWriter, r *http.Request)
	Get()
	Post()
}
