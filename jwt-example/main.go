package main

import (
	"jwt-example/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/revoke", handlers.RevokeJWT)
	http.HandleFunc("/rotate", handlers.RotateJWT)

	println("Listening on :9060")
	log.Fatal(http.ListenAndServe(":9060", nil))
}
