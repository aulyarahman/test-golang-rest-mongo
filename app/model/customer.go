package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Customer struct {
		Id        primitive.ObjectID `json:"id,omitempty" bson:"id"`
		UserName  string             `json:"user_name" bson:"user_name"`
		Email     string             `json:"email" bson:"email"`
		Password  string             `json:"password" bson:"password"`
		CreatedAt int32              `json:"created_at" bson:"created_at"`
	}

	ListCustomer struct {
		Data []Customer `json:"data"`
	}

	CustomerById struct {
		Id string `json:"id" bson:"id"`
	}
)
