package sqlquery

import (
	"carda/internal/carda/connector"
	"carda/internal/carda/converter"
)

var (
	db = connector.Connect()
)

func InsertUser(name, email string) error {
	query := `call add_user_with_email($1, $2);`
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUsersAndContacts() ([]converter.UserModel, error) {
	query := `SELECT usr.user_id, usr.username, ct.contact_value FROM users usr JOIN contacts ct on ct.user_id = usr.user_id;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	usersdata := []converter.UserModel{}
	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		usersdata = append(usersdata, converter.UserModel{Id: id, Name: name, Email: email})
	}
	defer rows.Close()
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
