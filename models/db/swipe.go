package models

//
//import (
//	"github.com/golang-jwt/jwt/v4"
//	"github.com/kamva/mgm/v3"
//)
//
//type Swipe struct {
//	mgm.DefaultModel  `bson:",inline"`
//	UserId            string   `json:"user_id" bson:"user_id"`
//
//}
//
//type SwipeClaims struct {
//	jwt.RegisteredClaims
//	Email string `json:"email"`
//	Type  string `json:"type"`
//}
//
//func NewSwipe(swipe struct{}) *Swipe {
//	return &Swipe{
//		UserId:            userId,
//	}
//}
//
//func (model *Swipe) CollectionName() string {
//	return "swipe"
//}
