package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
)

var passwordRule = []validation.Rule{
	validation.Required,
	validation.Length(8, 32),
	validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a RegisterRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 64)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a LoginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

type RefreshRequest struct {
	Token string `json:"token"`
}

func (a RefreshRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(
			&a.Token,
			validation.Required,
			validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
		),
	)
}

type ExceedMessageRequest struct {
	Command  string `json:"command"`
	Header   string `json:"header"`
	RFB_MSG  string `json:"rfb___msg"`
	Email    string `json:"email"`
	Password string `json:"password"`
	TRAILER  string `json:"trailer"`
}

type ProfileRequest struct {
	UserId            string                   `json:"user_id"`
	Age               int64                    `json:"Age"`
	AgeFilterMax      int64                    `json:"age_filter_max"`
	AgeFilterMin      int64                    `json:"age_filter_min"`
	InterestedIn      []string                 `json:"interested_in"`
	Gender            string                   `json:"gender"`
	PreferredLocation []string                 `json:"preferred_location"`
	Swipes            []map[string]interface{} `json:"swipes"`
}

type SwipeRequest struct {
	UserId     string `json:"user_id"`
	ProfileId  int64  `json:"profile_id"`
	Preference int64  `json:"preference"`
}

type Filters struct {
	Age    string `json:"age" query:"age"`
	Gender string `json:"gender" query:"gender"`
}
