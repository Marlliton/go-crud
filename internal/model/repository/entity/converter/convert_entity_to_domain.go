package converter

import (
	"github.com/Marlliton/go-crud/internal/model"
	"github.com/Marlliton/go-crud/internal/model/repository/entity"
)

func ConvetEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.Email, entity.Name, entity.Password, entity.Age)

	domain.SetID(entity.ID)

	return domain
}
