package routes

import (
	"carda/internal/carda/handlers"
	"fmt"
	"net/http"
)

func SetupRoutes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/carda/static/"))))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/adduserform", handlers.AddUserFormHandler)
	http.HandleFunc("/userscontactsform", handlers.UsersContactsFormHandler)
	http.HandleFunc("/submituserdata", handlers.SubmitUserHandler)
	http.HandleFunc("/deleteuser/", handlers.DeleteUserHandler)

}

func StartRouting() {
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
