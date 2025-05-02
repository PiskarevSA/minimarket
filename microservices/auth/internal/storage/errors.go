package storage

type StorageError struct{ Msg string }

func (e *StorageError) Error() string {
	return e.Msg
}

var ErrUserNotFound = &StorageError{"user not found"}
var ErrLoginAlreadyInUse = &StorageError{"login already in use"}
