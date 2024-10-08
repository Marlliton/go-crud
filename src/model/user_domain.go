package model

import (
	"crypto/md5"
	"encoding/hex"
)

// NOTE: Aqui é onde seria o nosso core com regras de negócio.

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetAge() int8
	GetEmail() string
	EncryptPassword()
}

// WARN: Acredito não ser necessário exportar as structs do domínio (verificar isso mais tarde)

type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

// NOTE: Vamos usar uma função construtora para retornar um UserDomainInterface
func NewUserDomain(email, name, password string, age int8) UserDomainInterface {
	return &userDomain{email, name, password, age}
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

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.password))

	u.password = hex.EncodeToString(hash.Sum(nil))
}
