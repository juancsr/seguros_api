package main

import (
	controller "falebella_chile/controllers"
	"falebella_chile/models"
	"fmt"
	"log"
	"net/http"
	"time"

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

func GetAllPlanesSeguro() {
	planes, err := models.GetAllPlanSeguros()
	if err != nil {
		log.Fatal("error: ", err)
	}

	for k, v := range planes {
		fmt.Println(k, " : ", v)
	}
}

func insertarCotizacion() {
	cotizacion := models.Cotizacion{
		Rut:                "12.345678-9",
		FechaNacimiento:    time.Now(),
		CorreoElectronico:  "ejemplo@dominio.com",
		Telefono:           "941234567",
		PrefijoNumTelefono: "+57",
	}
	err := models.RegistrarCotizacion(&cotizacion)
	if err != nil {
		log.Fatal("error: ", err)
	}
}

func pruebas() {
	pruebas, err := models.GetAllPruebas()

	if err != nil {
		log.Fatal("error:", err)
	}

	for k, v := range pruebas {
		fmt.Printf("%d: %s", k, v.Name)
	}
}
