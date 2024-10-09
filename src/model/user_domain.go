package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
)

// NOTE: Aqui é onde seria o nosso core com regras de negócio.

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetAge() int8
	GetEmail() string

	SetID(string)

	GetJSONValue() (string, error)
	EncryptPassword()
}

// WARN: Acredito não ser necessário exportar as structs do domínio (verificar isso mais tarde)

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int8
}

// NOTE: Vamos usar uma função construtora para retornar um UserDomainInterface
func NewUserDomain(email, name, password string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Name:     name,
		Password: password,
		Age:      age,
	}
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

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))

	u.Password = hex.EncodeToString(hash.Sum(nil))
}
