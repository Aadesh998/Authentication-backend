package models

type User struct {
	Name          string
	Email         string
	Password      string
	Verified      bool
	VerifiedToken string
}
