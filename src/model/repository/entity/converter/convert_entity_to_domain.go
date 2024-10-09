package converter

import (
	"github.com/Marlliton/go-crud/src/model"
	"github.com/Marlliton/go-crud/src/model/repository/entity"
)

func ConvetEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(entity.Email, entity.Name, entity.Password, entity.Age)

	domain.SetID(entity.ID)

	return domain
}
