package model

type User struct{
	UserID string `json:"user_id,omitempty" bson:"user_id"`
	Name string `json:"name,omitempty" bson:"name"`
	Department string `json:"department,omitempty" bson:"department"`
}