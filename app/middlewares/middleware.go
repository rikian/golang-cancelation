package middlewares

import (
	"net/http"

	"go/cancel/app/router/delete"
	"go/cancel/app/router/get"
	"go/cancel/app/router/post"
	"go/cancel/app/router/put"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	// check authorization
	// http method GET
	if r.Method == http.MethodGet {
		get.Handler(w, r)
		return
	}

	// http method POST
	if r.Method == http.MethodPost {
		post.Handler(w, r)
		return
	}

	// http method DELETE
	if r.Method == http.MethodDelete {
		delete.Handler(w, r)
		return
	}

	// http method PUT
	if r.Method == http.MethodPut {
		put.Handler(w, r)
		return
	}
}
