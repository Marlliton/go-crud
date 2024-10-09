package model

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetAge() int8
	GetEmail() string

	SetID(string)

	GetJSONValue() (string, error)
	EncryptPassword()
}

// NOTE: função construtora para criar/retornar um UserDomainInterface
func NewUserDomain(email, name, password string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Name:     name,
		Password: password,
		Age:      age,
	}
}
