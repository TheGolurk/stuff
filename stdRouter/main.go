package main

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i < 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

type App struct {
	UserHandler *UserHandler
}

func (h *App) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	if head == "user" {
		h.UserHandler.ServeHTTP(res, req)
		return
	}
	http.Error(res, "Not Found", http.StatusNotFound)
}

type UserHandler struct{}

func (h *UserHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = ShiftPath(req.URL.Path)
	id, err := strconv.Atoi(head)
	if err != nil {
		http.Error(res, fmt.Sprintf("Invalid user id %q", head), http.StatusBadRequest)
		return
	}
	switch req.Method {
	case "GET":
		h.handleGet(id)
	case "Post":
		h.handlePost(id)
	default:
		http.Error(res, "Only Post And Get methods are allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	a := &App{
		UserHandler: new(UserHandler),
	}
	http.ListenAndServe(":3000", a)
}
