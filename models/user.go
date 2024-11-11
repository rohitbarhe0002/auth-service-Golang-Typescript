package models

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password,omitempty" bson:"-"`
	PasswordHash string `json:"-" bson:"password_hash"`    
}
