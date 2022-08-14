package middlewares

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Middleware(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.Copy(ioutil.Discard, r.Body)
	if err != nil {
		http.Error(w, "body not allowed", http.StatusBadRequest)
		log.Print(body)
		return
	}
	defer r.Body.Close()

	// check auth etc

	select {
	case <-time.After(5 * time.Second):
		w.Write([]byte("hello world"))
		log.Print("send response succesfully...")
	case <-ctx.Done():
		log.Printf("%v Request success canceled ... ", r.Method)
	}
}
