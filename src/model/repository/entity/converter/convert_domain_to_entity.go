package converter

import (
	"github.com/Marlliton/go-crud/src/model"
	"github.com/Marlliton/go-crud/src/model/repository/entity"
)

func ConvetDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Name:     domain.GetName(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}
