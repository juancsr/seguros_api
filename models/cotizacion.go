package models

import (
	"falebella_chile/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cotizacion struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Rut                string             `bson:"rut"`
	FechaNacimiento    time.Time          `bson:"fechaNacimiento"`
	CorreoElectronico  string             `bson:"correoElectronico"`
	Telefono           string             `bson:"telefono"`
	PrefijoNumTelefono string             `bson:"prefijoNumTelefono"`
}

var COLLECTION_NAME = "cotizacion"

// RegistrarCotizacion
func RegistrarCotizacion(cotizacion *Cotizacion) error {
	ctx := db.CTX
	collection := db.GetCollection(collectionName)
	_, err := collection.InsertOne(ctx, cotizacion)
	return err
}
