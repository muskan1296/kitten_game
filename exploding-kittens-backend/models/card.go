package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Card struct {
    ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Type  string             `bson:"type" json:"type"`
    Value string             `bson:"value,omitempty" json:"value"`
}
