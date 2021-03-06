package controller

import (
	"encoding/json"
	"falebella_chile/models"
	"falebella_chile/resources"
	"io/ioutil"
	"log"
	"net/http"
)

// Cotizar recibe la información de una cotización y retorna una lista de planes de seguro
// swagger:operation GET /cotizar_plan_seguros users
// cotizar un plan de seguros
// ---
// produces:
// - application/json
// parameters:
// - name: string
// responses:
//   '200':
//     description: []Planes
func Cotizar(w http.ResponseWriter, r *http.Request) {
	var cotizacion models.Cotizacion
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &cotizacion)
	registrarCotizacion(&cotizacion)

	planes, err := models.GetAllPlanSeguros()
	if err != nil {
		log.Fatalln(err)
	}
	asignarMontoUf(planes)

	json.NewEncoder(w).Encode(planes)
}

//registrarCotizacion Registra la cotización en la bd
func registrarCotizacion(cotizacion *models.Cotizacion) {
	err := models.RegistrarCotizacion(cotizacion)
	if err != nil {
		log.Fatalln(err)
	}
}

//asignarMontoUf asigna el valor de uf a cada plan de seguro
func asignarMontoUf(planes []*models.PlanSeguro) {
	ufValue := getCurrentUfValue()
	for _, plan := range planes {
		plan.MontoUf = plan.MontoPesos / ufValue
	}
}

//getCurrentUfValue Valor actual de UF en Pesos
func getCurrentUfValue() float64 {
	resp, err := resources.GetUfValue()
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Serie[0].Valor
}
