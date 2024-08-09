package converter

import (
	"log"
	"strconv"
)

type UserView struct {
	Id    string
	Name  string
	Email string
}

type UserModel struct {
	Id    int
	Name  string
	Email string
}

func UserViewToModel(userview *UserView) UserModel {
	id, err := strconv.Atoi(userview.Id)
	if err != nil {
		log.Fatal(err)
	}
	return UserModel{Id: id, Email: userview.Email, Name: userview.Name}
}

func UserModelToView(usermodel *UserModel) UserView {
	id := strconv.Itoa(usermodel.Id)
	return UserView{Id: id, Email: usermodel.Email, Name: usermodel.Name}
}
