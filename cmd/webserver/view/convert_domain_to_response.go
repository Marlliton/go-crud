package view

import (
	"github.com/Marlliton/go-crud/cmd/webserver/controller/dto/response"
	"github.com/Marlliton/go-crud/internal/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
