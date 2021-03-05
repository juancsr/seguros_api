package main

import (
	"fmt"
	"log"
	"time"

	"falebella_chile/models"
)

func main() {
	cotizacion := models.Cotizacion{
		Rut:                "12.345678-9",
		FechaNacimiento:    time.Now(),
		CorreoElectronico:  "ejemplo@dominio.com",
		Telefono:           "941234567",
		PrefijoNumTelefono: "+57",
	}
	models.RegistrarCotizacion(&cotizacion)
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
