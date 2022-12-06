package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/kamva/mgm/v3"
)

type Profile struct {
	mgm.DefaultModel  `bson:",inline"`
	UserId            string                   `json:"user_id" bson:"user_id"`
	Age               int64                    `json:"age" bson:"age"`
	AgeFilterMax      int64                    `json:"age_filter_max" bson:"age_filter_max"`
	AgeFilterMin      int64                    `json:"age_filter_min" bson:"age_filter_min"`
	InterestedIn      []string                 `json:"interested_in" bson:"interested_in"`
	Gender            string                   `json:"gender" bson:"gender"`
	PreferredLocation []string                 `json:"preferred_location" bson:"preferred_location"`
	Swipe             []map[string]interface{} `json:"swipe" bson:"swipe"`
}

type Swipe struct {
	UserId     string `json:"user_id"`
	Preference string `json:"preference"`
}

type ProfileClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
	Type  string `json:"type"`
}

func NewProfile(userId string, age int64, minAge int64, maxAge int64, interestIn []string, gender string, preferredLocation []string, swipe []map[string]interface{}) *Profile {
	return &Profile{
		UserId:            userId,
		Age:               age,
		AgeFilterMin:      minAge,
		AgeFilterMax:      maxAge,
		InterestedIn:      interestIn,
		Gender:            gender,
		PreferredLocation: preferredLocation,
		Swipe:             swipe,
	}
}

func (model *Profile) CollectionName() string {
	return "profile"
}
