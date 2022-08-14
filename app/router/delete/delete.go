package delete

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.Copy(ioutil.Discard, r.Body)
	if err != nil {
		http.Error(w, "body not allowed", http.StatusBadRequest)
		log.Print(body)
		return
	}
	defer r.Body.Close()
	w.Write([]byte("post handler ready"))
}
