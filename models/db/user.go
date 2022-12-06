package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/kamva/mgm/v3"
)

const (
	RoleUser = "user"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password" bson:"password"`
	Name             string `json:"name" bson:"name"`
	Gender           string `json:"gender" bson:"gender"`
	Age              int    `json:"age" bson:"age"`
	Location         string `json:"location" bson:"location"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Type  string `json:"type"`
}

func NewUser(email string, password string, name string, gender string, age int, location string) *User {
	return &User{
		Email:    email,
		Password: password,
		Name:     name,
		Gender:   gender,
		Age:      age,
		Location: location,
	}
}

func (model *User) CollectionName() string {
	return "users"
}
