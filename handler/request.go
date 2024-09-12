package handler

import (
	"fmt"
)

func requiredFieldErr(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (c *CreateOpeningRequest) Validate() error {

	if c.Role == "" {
		return requiredFieldErr("role", "string")
	}
	if c.Company == "" {
		return requiredFieldErr("company", "string")
	}

	if c.Location == "" {
		return requiredFieldErr("location", "string")
	}

	if c.Remote == nil {
		return requiredFieldErr("remote", "bool")
	}

	if c.Link == "" {
		return requiredFieldErr("link", "string")
	}

	if c.Salary == 0 {
		return requiredFieldErr("salary", "int64")
	}

	return nil
}
