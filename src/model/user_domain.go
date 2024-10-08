package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/Marlliton/go-crud/src/configuration/rest_err"
)

// NOTE: Aqui é onde seria o nosso core com regras de negócio.

// WARN: Acredito não ser necessário exportar as structs do domínio (verificar isso mais tarde)

type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

// NOTE: Vamos usar uma função construtora para retornar um UserDomainInterface
func NewUserDomain(email, name, password string, age int8) UserDomainInterface {
	return &userDomain{email, name, password, age}
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr                 // NOTE: Podemos omitir o nome dos  parametros, (string, userDomain) == (id string, user userDomain)
	FindUser(id string) (*userDomain, *rest_err.RestErr) // NOTE: Aqui já nomeamos o parametro
	DeleteUser(id string) *rest_err.RestErr
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))

	u.Password = hex.EncodeToString(hash.Sum(nil))
}
