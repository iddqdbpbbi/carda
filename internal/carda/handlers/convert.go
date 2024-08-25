package handlers

import (
	"carda/internal/carda/shared"
)

type view struct {
	Id    string
	Name  string
	Email string
}

func dTOToView(dto *shared.User) view {
	return view{Id: dto.Id, Email: dto.Email, Name: dto.Name}
}

func dTOToViewS(models []*shared.User) (veiws []*view) {
	veiws = make([]*view, 0, len(models))
	for _, model := range models {
		view := dTOToView(model)
		veiws = append(veiws, &view)
	}
	return
}
