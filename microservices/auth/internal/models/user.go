package models

type UserLoggedIn struct {
	UserId       string
	Login        string
	AccessToken  string
	RefreshToken string
}

type UserRegistered struct {
	UserId       string
	Login        string
	AccessToken  string
	RefreshToken string
}
