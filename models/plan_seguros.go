package models

import (
	"falebella_chile/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Aseguradora atributo de la colección Aseguradora
type Aseguradora struct {
	Nombre   string `bson:"nombre" json:"nombre"`
	Contacto string `bson:"contacto" json:"contacto"`
	Ciudad   string `bson:"ciudad" json:"ciudad"`
	Pais     string `bson:"pais" pais:"pais"`
}

//Promocion atributo de la colección PlanSeguro
type Promocion struct {
	Nombre          string `bson:"nombre" json:"nombre"`
	Codigo          string `bson:"codigo" json:"codigo"`
	MesCuotasGratis []int8 `bson:"mes_cuota_gratis" json:"mes_cuota_gratis"`
	Descripcion     string `bson:"descripcion" json:"desripcion"`
}

//PlanSeguro representa una colección de mongodb
type PlanSeguro struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre           string             `bson:"nombre" json:"nombre"`
	Aseguradora      *Aseguradora       `bson:"aseguradora" json:"aseguradora"`
	MontoPesos       float64            `bson:"monto_pesos" json:"monto_pesos"`
	CapitalAsegurado float64            `bson:"capital_asegurado" json:"capital_asegurado"`
	Promocion        *Promocion         `bson:"promocion" json:"promocion"`
	MontoUf          float64            `json:"monto_uf"`
}

const plansegurosCollectionName = "plan_seguro"

//GetAllPlanSeguros obtener todas las planesSeguro
func GetAllPlanSeguros() ([]*PlanSeguro, error) {
	filter := bson.D{{}}
	// defer db.CloseConnection()
	return filterPlanSeguros(filter)
}

func filterPlanSeguros(filter interface{}) ([]*PlanSeguro, error) {
	ctx := db.CTX
	collection := db.GetCollection(plansegurosCollectionName)

	var planesSeguro []*PlanSeguro

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return planesSeguro, err
	}

	for cur.Next(ctx) {
		var p PlanSeguro
		err := cur.Decode(&p)
		if err != nil {
			return planesSeguro, err
		}

		planesSeguro = append(planesSeguro, &p)
	}

	if err := cur.Err(); err != nil {
		return planesSeguro, err
	}

	cur.Close(ctx)

	if len(planesSeguro) == 0 {
		return planesSeguro, mongo.ErrNoDocuments
	}

	return planesSeguro, nil
}

//RegistrarPlanSeguro registra planes de seguro en la bd
func RegistrarPlanSeguro(planSeguro *PlanSeguro) error {
	ctx := db.CTX
	collection := db.GetCollection(plansegurosCollectionName)
	planSeguro.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, planSeguro)
	return err
}
