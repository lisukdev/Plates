package user

import "github.com/google/uuid"

type AuthorizedUser struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}
