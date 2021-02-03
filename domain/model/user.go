package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct 
{
	ID    string `json: "id" valid: "uuid"`
	Name  string `json: "name" valid: "notnull"`
	Email string `json: "email" valid: "notnull"`
	CreatedAt time.Time	`json: "created_at" valid:"-"` 
}

func NewUser( id string, name string, email string) 
{
	user = User
	{
		ID:   id,
		Name: name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.createdAt = time.Now()

	err := user.isValid()

	if err != nil
	{
		return &user, nil
	}
}