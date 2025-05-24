package objects

import "github.com/google/uuid"

type UserID uuid.UUID

var NullUserID = UserID(uuid.Nil)

func NewUserID(userID uuid.UUID) UserID {
	return UserID(userID)
}

func (u UserID) UUID() uuid.UUID {
	return uuid.UUID(u)
}

func (u UserID) Equal(userID UserID) bool {
	return u == userID
}
