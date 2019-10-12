package main

import (
	"broz3r.com/HelloGo_Rest-API/controllers"
	"broz3r.com/HelloGo_Rest-API/net"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/helloworld", controllers.HelloWorld).Methods("GET")
	router.Use(net.Logger)

	port := "9000"

	log.Printf("Listen on the port : %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Print(err)
	}
}
