package objects

import "github.com/google/uuid"

type UserId uuid.UUID

var NullUserId = UserId(uuid.Nil)

func NewUserId(userId uuid.UUID) UserId {
	return UserId(userId)
}

func (u UserId) Uuid() uuid.UUID {
	return uuid.UUID(u)
}

func (u UserId) Equal(userId UserId) bool {
	return u == userId
}
