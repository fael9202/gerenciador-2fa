package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TOTPAccount struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Name      string            `bson:"name" json:"name"`
	Issuer    string            `bson:"issuer" json:"issuer"`
	Secret    string            `bson:"secret" json:"secret"`
	CreatedAt time.Time         `bson:"created_at" json:"created_at"`
} 