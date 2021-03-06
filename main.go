// Seguros API REST
//
// Seguros API REST Prueba Falabella Chile
//
//     Schemes: http, https
//     Host: localhost:3000
//     Version: 1.0
//     basePath: /
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	controller "falebella_chile/controllers"
	"falebella_chile/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	fmt.Println("* Init application *")
	planes, _ := models.GetAllPlanSeguros()
	if len(planes) == 0 {
		log.Println("Registrando planes...")
		registrarPlanes()
	}
}

func registrarPlanes() {
	aseguradora := models.Aseguradora{
		Nombre:   "MetLife",
		Contacto: "metlife@metlife.cl",
		Ciudad:   "Santiago",
		Pais:     "Chile",
	}
	mesesPromocion := []int8{3, 8}
	promocion := models.Promocion{
		Nombre:          "Promoción",
		Codigo:          "ERSFGGERT21434v",
		MesCuotasGratis: mesesPromocion,
		Descripcion:     "2 Cuotas Gratis",
	}

	plan := models.PlanSeguro{
		Nombre:           "PLAN Metlife",
		Aseguradora:      &aseguradora,
		MontoPesos:       11755,
		CapitalAsegurado: 500,
		Promocion:        &promocion,
	}
	models.RegistrarPlanSeguro(&plan)
	plan.MontoPesos = 23223
	plan.CapitalAsegurado = 1000
	models.RegistrarPlanSeguro(&plan)
	plan.MontoPesos = 34978
	plan.CapitalAsegurado = 1500
	models.RegistrarPlanSeguro(&plan)
}

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
