package models

import "time"

type MindicadorCLSerie struct {
	Fecha time.Time `json:"fecha"`
	Valor float64   `json:"valor"`
}

type MindicadorCL struct {
	Version      string              `json:"version"`
	Autor        string              `json:"autor"`
	Codigo       string              `json:"codigo"`
	Nombre       string              `json:"nombre"`
	UnidadMedida string              `json:"unidad_medida"`
	Serie        []MindicadorCLSerie `json:"serie"`
}
