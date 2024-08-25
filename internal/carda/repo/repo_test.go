package repo

import (
	"fmt"
	"log"
	"testing"
)

func TestInsertTestUsers(testing *testing.T) {
	user_names := [10]string{"Alice", "Bob", "Charlie", "Deelan", "Edan", "Fred", "George", "Ian", "Jeronimo", "Hank"}
	for _, user_name := range user_names {
		err := InsertUser(user_name, user_name+"@gmail.com")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func TestQuery(t *testing.T) {
	rows, err := GetUsersAndContacts()
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range rows {
		log.Printf("id=%s, name=%s, email=%s\n", row.Id, row.Name, row.Email)
	}
}

func TestDeleteUser(t *testing.T) {
	rows, err := GetUsersAndContacts()
	if err != nil {
		log.Fatal(err)
	}
	id := rows[0].Id
	err = DeleteUserById(id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delete user with id %s", id)
}
