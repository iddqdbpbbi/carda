package repo

import (
	"carda/internal/carda/dbconnect"
	"carda/internal/carda/shared"
)

var (
	db = dbconnect.Connect()
)

func InsertUser(name, email string) error {
	query := `call add_user_with_email($1, $2);`
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUsersAndContacts() (usersdata []*shared.User, err error) {
	query := `SELECT usr.user_id, usr.username, ct.contact_value FROM users usr JOIN contacts ct on ct.user_id = usr.user_id;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		modeluser := (model{Id: id, Name: name, Email: email})
		user := modeluser.modelToDTO()
		usersdata = append(usersdata, &user)
	}
	return usersdata, nil
}

func DeleteUserById(id string) error {
	query := `delete from users usr where usr.user_id = $1;`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
