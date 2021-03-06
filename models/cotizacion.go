package models

import (
	"falebella_chile/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Cotizacion representa una colección de mongodb
type Cotizacion struct {
	ID                 primitive.ObjectID `bson:"_id" json:"id" binding:"required"`
	Rut                string             `bson:"rut,omitempty" json:"rut" binding:"required"`
	FechaNacimiento    time.Time          `bson:"fechaNacimiento" json:"fecha_nacimiento"`
	CorreoElectronico  string             `bson:"correoElectronico" json:"correo_electronico"`
	Telefono           string             `bson:"telefono" json:"telefono"`
	PrefijoNumTelefono string             `bson:"prefijoNumTelefono" json:"prefijo_num_telefono"`
	cotizado           bool               `bson:"cotizado"`
}

const cotizacionCollectionName = "cotizacion"

// RegistrarCotizacion registra en BD
func RegistrarCotizacion(cotizacion *Cotizacion) error {
	ctx := db.CTX
	collection := db.GetCollection(cotizacionCollectionName)
	cotizacion.ID = primitive.NewObjectID()
	cotizacion.cotizado = false
	_, err := collection.InsertOne(ctx, cotizacion)
	return err
}
