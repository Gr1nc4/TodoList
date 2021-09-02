package main

import (
	"net/http"
	"todo/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/todolist", controllers.Todolist).Methods("GET")
	rtr.HandleFunc("/save_item", controllers.Save_item).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}/done_item", controllers.Done_item).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}/remove_item", controllers.Remove_item).Methods("POST")
	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}
func main() {
	handleFunc()
}
