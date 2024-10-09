package model

import (
	"encoding/json"
	"errors"
)

// NOTE: Aqui é onde seria o nosso core com regras de negócio.
// WARN: Acredito não ser necessário exportar as structs do domínio (verificar isso mais tarde)

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (u *userDomain) GetEmail() string {
	return u.Email
}
func (u *userDomain) GetName() string {
	return u.Name
}
func (u *userDomain) GetPassword() string {
	return u.Password
}
func (u *userDomain) GetAge() int8 {
	return u.Age
}

func (u *userDomain) SetID(id string) {
	u.ID = id
}

func (u *userDomain) GetJSONValue() (string, error) {
	json, err := json.Marshal(u)
	if err != nil {
		return "", errors.New("Error trying convert userDomain")
	}

	return string(json), nil
}
