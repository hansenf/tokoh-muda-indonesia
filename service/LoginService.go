package service

import "encoding/json"

type LoginService interface {
	LoginUser(nohp json.Number, password string) bool
}
type loginInformation struct {
	nohp     json.Number
	password string
}

func StaticLoginService() LoginService {

	return &loginInformation{
		nohp:     "1234",
		password: "rafli",
	}
}
func (info *loginInformation) LoginUser(nohp json.Number, password string) bool {
	return info.nohp == nohp && info.password == password
}
