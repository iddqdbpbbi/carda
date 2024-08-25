package repo

import (
	"carda/internal/carda/shared"
	"strconv"
)

type model struct {
	Id    int
	Name  string
	Email string
}

func (model model) modelToDTO() shared.User {
	id := strconv.Itoa(model.Id)
	return shared.User{Id: id, Email: model.Email, Name: model.Name}
}
