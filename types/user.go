package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var (
	MinFirstNameLength = 2
	MinLastNameLength  = 2
	MinPasswordLength  = 6
)

func (params CreateUserParams) Validate() map[string]string {
	errors := make(map[string]string)
	if len(params.FirstName) < MinFirstNameLength {
		errors["firstName"] = fmt.Sprintf("first name must be at least %d characters", MinFirstNameLength)
	}
	if len(params.LastName) < MinLastNameLength {
		errors["lastName"] = fmt.Sprintf("last name must be at least %d characters", MinLastNameLength)
	}
	if len(params.Password) < MinPasswordLength {
		errors["password"] = fmt.Sprintf("password must be at least %d characters", MinPasswordLength)
	}
	if !isValidEmail(params.Email) {
		errors["email"] = "email is invalid"
	}
	return errors
}

func isValidEmail(email string) bool {
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Match the email against the pattern
	match, _ := regexp.MatchString(pattern, email)

	return match
}

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (params UpdateUserParams) ToBson() (bson.M, map[string]string) {
	bsonParams := bson.M{}
	errors := make(map[string]string)
	if params.FirstName != "" && len(params.FirstName) < MinFirstNameLength {
		errors["firstName"] = fmt.Sprintf("first name must be at least %d characters", MinFirstNameLength)
	}
	if params.LastName != "" && len(params.LastName) < MinLastNameLength {
		errors["lastName"] = fmt.Sprintf("last name must be at least %d characters", MinLastNameLength)
	}

	if len(errors) > 0 {
		return nil, errors
	}

	if len(params.FirstName) > 0 {
		bsonParams["firstName"] = params.FirstName
	}
	if len(params.LastName) > 0 {
		bsonParams["lastName"] = params.LastName
	}

	return bsonParams, nil
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
	IsAdmin   bool               `bson:"isAdmin" json:"-" default:"false"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err

	}
	return &User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  string(hasedPassword),
	}, nil
}

func IsValidPassword(hashedpassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	return err == nil
}
