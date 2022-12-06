package utils

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertStringToObjectId(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return [12]byte{}, errors.New("failed to convert user id to object id")

	}
	return objID, err
}

func Contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func SubSlice(s1 []string, s2 []string) bool {
	for _, e := range s2 {
		if Contains(s1, e) {
			return true
		}
	}
	return false
}
