package models

//Login credential
type LoginCredentials struct {
	Nohp     string `form:"nohp"`
	Password string `form:"password"`
}
