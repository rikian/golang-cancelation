package app

import (
	"go/cancel/app/config"
	"go/cancel/app/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Listener(address string) {
	runDb()
	runServer(address)
}

func runDb() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbConn, status := config.ConnectDB(os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	if status {
		log.Fatal(dbConn)
		return
	}

	log.Println(dbConn)
}

func runServer(address string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { middlewares.Middleware(w, r) })

	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Print("server listen and serve " + address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
