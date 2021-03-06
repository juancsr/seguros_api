package main

import (
	controller "falebella_chile/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Seguros API")
}

func apiRequest() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered", r)
		}
	}()
	http.HandleFunc("/", home)
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/cotizar_plan_seguros", controller.Cotizar).Methods("POST")
	fmt.Println("Server is running in 3000 port...")
	log.Fatal(http.ListenAndServe(":3000", route))
}

func main() {
	fmt.Println("Starting server....")
	apiRequest()
}
