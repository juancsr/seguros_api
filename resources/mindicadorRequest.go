package resources

import (
	"encoding/json"
	"falebella_chile/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// GetUfValue Request al api de midicador para obtener el valor del uf actual
func GetUfValue() (models.MindicadorCL, error) {
	var apiPath strings.Builder
	now := time.Now()

	apiPath.WriteString("https://mindicador.cl/api/uf/")
	apiPath.WriteString(now.Format("02-01-2006"))

	resp, err := http.Get(apiPath.String())
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var indicador models.MindicadorCL
	json.Unmarshal(body, &indicador)
	return indicador, err
}
