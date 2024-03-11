package db

import "github.com/google/uuid"

type CreateUserParams struct {
	UUID     uuid.UUID `bson:"uuid"`
	UserName string
	Password string
}
