package model

type UserDomainInterface interface {
	GetName() string
	GetPassword() string
	GetAge() int8
	GetEmail() string
	GetID() string

	SetID(string)

	EncryptPassword()
}

// NOTE: função construtora para criar/retornar um UserDomainInterface
func NewUserDomain(email, name, password string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		name:     name,
		password: password,
		age:      age,
	}
}

func NewUserUpdateDomain(name string, age int8) UserDomainInterface {
	return &userDomain{
		name: name,
		age:  age,
	}
}
