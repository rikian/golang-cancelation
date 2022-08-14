package controllers

import (
	"context"
	"encoding/json"
	"go/cancel/app/entity"
	"go/cancel/app/repository/products"
	"log"
	"net/http"
	"time"
)

func InsertProducts(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context = r.Context()
	var jsonProduct *entity.Product
	err := json.NewDecoder(r.Body).Decode(&jsonProduct)
	select {
	case <-time.After(1 * time.Second):
		if err != nil {
			log.Print(err)
			w.Write([]byte("invalid format"))
			return
		}
		var product = products.InitProduct()
		err = product.InsertProduct(jsonProduct)
		if err != nil {
			w.Write([]byte("failed insert data " + err.Error()))
			return
		}
		w.Write([]byte("success insert product"))
		log.Print("send response succesfully...")
	case <-ctx.Done():
		log.Printf("%v Request success canceled ... ", r.Method)
	}
}
