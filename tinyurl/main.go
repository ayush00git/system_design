package main

import (
	"fmt"
	"net/http"
	"log"

	"tinyurl/helpers"
)

func main() {
	uri := helpers.GetEnv()

	helpers.ConnectToMongo(uri);

	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Hey this is the home route")
		}),
	}

	log.Fatal(srv.ListenAndServe())
}