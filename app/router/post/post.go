package post

import (
	"go/cancel/app/controllers"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/products/" {
		controllers.InsertProducts(w, r)
		return
	}

	w.Write([]byte("url not allowed"))
}
