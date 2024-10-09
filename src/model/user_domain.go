package model

import (
	"encoding/json"
	"errors"
)

// NOTE: Aqui é onde seria o nosso core com regras de negócio.
// WARN: Acredito não ser necessário exportar as structs do domínio (verificar isso mais tarde)

type userDomain struct {
	id       string
	email    string
	password string
	name     string
	age      int8
}

func (u *userDomain) GetEmail() string {
	return u.email
}
func (u *userDomain) GetName() string {
	return u.name
}
func (u *userDomain) GetPassword() string {
	return u.password
}
func (u *userDomain) GetAge() int8 {
	return u.age
}
func (u *userDomain) GetID() string {
	return u.id
}

func (u *userDomain) SetID(id string) {
	u.id = id
}

func (u *userDomain) GetJSONValue() (string, error) {
	json, err := json.Marshal(u)
	if err != nil {
		return "", errors.New("Error trying convert userDomain")
	}

	return string(json), nil
}
