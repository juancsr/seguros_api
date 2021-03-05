package models

import (
	"falebella_chile/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Cotizacion representa una colección de mongodb
type Cotizacion struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Rut                string             `bson:"rut"`
	FechaNacimiento    time.Time          `bson:"fechaNacimiento"`
	CorreoElectronico  string             `bson:"correoElectronico"`
	Telefono           string             `bson:"telefono"`
	PrefijoNumTelefono string             `bson:"prefijoNumTelefono"`
}

const cotizacionCollectionName = "cotizacion"

// RegistrarCotizacion registra en BD
func RegistrarCotizacion(cotizacion *Cotizacion) error {
	ctx := db.CTX
	collection := db.GetCollection(cotizacionCollectionName)
	_, err := collection.InsertOne(ctx, cotizacion)
	defer db.CloseConnection()
	return err
}
