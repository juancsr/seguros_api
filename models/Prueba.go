package models

import (
	"falebella_chile/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Prueba collection
type Prueba struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

var collectionName = "prueba"

//GetAllPruebas obtener todas las pruebas
func GetAllPruebas() ([]*Prueba, error) {
	filter := bson.D{{}}
	defer db.CloseConnection()
	return filterPruebas(filter)
}

func filterPruebas(filter interface{}) ([]*Prueba, error) {
	ctx := db.CTX
	collection := db.GetCollection(collectionName)

	// A slice of pruebas for storing the decoded documents
	var pruebas []*Prueba

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return pruebas, err
	}

	for cur.Next(ctx) {
		var p Prueba
		err := cur.Decode(&p)
		if err != nil {
			return pruebas, err
		}

		pruebas = append(pruebas, &p)
	}

	if err := cur.Err(); err != nil {
		return pruebas, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(pruebas) == 0 {
		return pruebas, mongo.ErrNoDocuments
	}

	return pruebas, nil
}
