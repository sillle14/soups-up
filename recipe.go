package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// TODO: This could maybe work with a full model using omitempty

type RecipeForm struct {
	Name         string `json:"name" form:"name"`
	Instructions string `json:"instructions" form:"instructions"`
}

type Recipe struct {
	Id           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" form:"name"`
	Instructions string             `json:"instructions" form:"instructions"`
}
