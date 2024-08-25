package handlers

import (
	"carda/internal/carda/repo"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Form struct {
	Usersdata []*view
}

type FormData struct {
	Name  string
	Email string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := parseWrapper(
		"web/carda/templates/index.html",
		"web/carda/templates/tabs/tabs.html",
	)
	tmpl.Execute(w, nil)
}

func UsersContactsFormHandler(w http.ResponseWriter, r *http.Request) {
	usersview := getUsers()
	tmpl := parseWrapper("web/carda/templates/tabs/tabs.html")
	tmpl.ExecuteTemplate(w, "userscontactsform", Form{Usersdata: usersview})
}

func AddUserFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := parseWrapper("web/carda/templates/tabs/tabs.html")
	tmpl.ExecuteTemplate(w, "adduserform", nil)
}

func SubmitUserHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	data := FormData{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}
	err := repo.InsertUser(data.Name, data.Email)
	if err != nil {
		log.Fatal("err with inserting user ", err)
	}
	response := `<p>Thank you, ` + data.Name + `! Your email ` + data.Email + ` has been received.</p>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(response))
}
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 || parts[2] == "" {
		log.Fatal("incorrect argument in path")
	}
	userid := parts[2]
	err := repo.DeleteUserById(userid)
	if err != nil {
		log.Fatal(err)
	}
}

func getUsers() (userview []*view) {
	usersmodel, err := repo.GetUsersAndContacts()
	if err != nil {
		log.Fatal(err)
	}
	userview = dTOToViewS(usersmodel)
	return userview
}

func parseWrapper(files ...string) *template.Template {
	return template.Must(template.ParseFiles(files...))
}
