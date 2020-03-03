package mgm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IDField struct contain model's ID field.
type IDField struct {
	MongoID primitive.ObjectID `json:"mongo_id" bson:"_id,omitempty"`
}

// DateFields struct contain `created_at` and `updated_at`
// fields that autofill on insert/update model.
type DateFields struct {
	MongoCreatedAt time.Time `json:"mongo_created_at" bson:"_created_at"`
	MongoUpdatedAt time.Time `json:"mongo_updated_at" bson:"_updated_at"`
}

// PrepareID method prepare id value to using it as id in filtering,...
// e.g convert hex-string id value to bson.ObjectId
func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method return model's id
func (f *IDField) GetID() interface{} {
	return f.MongoID
}

// SetID set id value of model's id field.
func (f *IDField) SetID(id interface{}) {
	f.MongoID = id.(primitive.ObjectID)
}

//--------------------------------
// DateField methods
//--------------------------------

// Creating hook used here to set `created_at` field
// value on inserting new model into database.
func (f *DateFields) Creating() error {
	f.MongoCreatedAt = time.Now().UTC()
	return nil
}

// Saving hook used here to set `updated_at` field value
// on create/update model.
func (f *DateFields) Saving() error {
	f.MongoUpdatedAt = time.Now().UTC()
	return nil
}
