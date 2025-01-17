package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portString:=os.Getenv("PORT")

	if(portString == ""){
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	err := srv.ListenAndServe()

	log.Printf("Server starting on port: %v\n", portString)
	if err != nil{
		log.Fatal(err)
	}
}