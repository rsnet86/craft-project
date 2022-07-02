package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	dbPsql = CreatePostgres()
	redisClient = CreateRedis()

	router := mux.NewRouter()
	fmt.Println("Router Created")
	router.HandleFunc("/url/shortening", CreateTinyURl).Methods("POST")
	router.HandleFunc("/url/shortening", GetTinyURl).Methods("GET")
	router.HandleFunc("/url/shortening", DeleteTinyURl).Methods("DELETE")

	fmt.Println("Listening at port number 8080")
	http.ListenAndServe(":8080", router)
	fmt.Println("testing the connection")
}
