package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type UserResponse struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
}
