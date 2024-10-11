package converter

import (
	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/internal/model/repository/entity"
)

func ConvetDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Name:     domain.GetName(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Age:      domain.GetAge(),
	}
}
