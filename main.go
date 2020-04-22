package main

import (
	"fmt"
	"hello-go/mycontroller"
	"net/http"

	"hello-go/myserver.go"
)

func main() {
	run()
}

func run() {
	var server myserver.MyServer

	var c FuckController

	server.RoutePath("/", "GET", "Fuck", &c)

	http.ListenAndServe("127.0.0.1:8080", &server)
}

type FuckController struct {
	mycontroller.BaseController
	a int
	b string
}

func (f *FuckController) Fuck() {
	fmt.Println("fuck invoked")

	f.W.Write([]byte("fuck invoked"))
}
