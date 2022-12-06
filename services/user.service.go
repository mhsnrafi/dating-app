package services

import (
	"errors"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/kamva/mgm/v3"
	"github.com/sethvargo/go-password/password"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	db "mitte-task/models/db"
	"time"
)

// CreateUser create a user record
func CreateUser() (*db.User, error) {

	gender := faker.Gender()
	name := faker.Name()
	email := faker.Email()
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	age := rand.Intn(30-18) + 18

	plainPassword, err := password.Generate(32, 10, 10, false, false)
	if err != nil {
		return nil, errors.New("cannot generate a password")
	}
	fmt.Println(plainPassword)

	pass, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("cannot generate hashed password")
	}

	locations := make([]string, 0)
	locations = append(locations,
		"Berlin",
		"Hannover",
		"Hamburg",
		"Munich",
		"Cologne",
		"Frankfurt",
		"Nuremberg",
		"Kiel",
		"Stuttgard",
		"Dusseldorf")

	rand.Seed(time.Now().Unix())
	location := locations[rand.Intn(len(locations))]
	user := db.NewUser(email, string(pass), name, gender, age, location)
	err = mgm.Coll(user).Create(user)
	if err != nil {
		return nil, errors.New("cannot create new user")
	}
	user.Password = plainPassword

	return user, nil
}

// FindUserById find user by id
func FindUserById(userId primitive.ObjectID) (*db.User, error) {
	user := &db.User{}
	err := mgm.Coll(user).FindByID(userId, user)
	if err != nil {
		return nil, errors.New("cannot find user")
	}

	return user, nil
}

// FindUserByEmail find user by email
func FindUserByEmail(email string) (*db.User, error) {
	user := &db.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		return nil, errors.New("cannot find user")
	}

	return user, nil
}

// CheckUserMail search user by email, return error if someone uses
func CheckUserMail(email string) error {
	user := &db.User{}
	userCollection := mgm.Coll(user)
	err := userCollection.First(bson.M{"email": email}, user)
	if err == nil {
		return errors.New("email is already in use")
	}

	return nil
}
