package main

import (
	"fmt"
	"net/http"
	"log"

	"tinyurl/internals/db"
)

func main() {
	uri := db.GetEnv()

	db.ConnectToMongo(uri);

	srv := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Hey this is the home route")
		}),
	}

	log.Fatal(srv.ListenAndServe())
}