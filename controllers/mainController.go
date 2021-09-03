package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"todo/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func dbConnect(dbQuery string) *sql.Rows {
	db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/todo_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query(dbQuery)
	if err != nil {
		panic(err)
	}
	return res
}

func Todolist(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/todolist.html")
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	res := dbConnect("SELECT * FROM `todo_item`")

	var items = []models.TodoThings{}
	for res.Next() {
		var item models.TodoThings
		err = res.Scan(&item.Id, &item.Item, &item.Status)
		if err != nil {
			panic(err)
		}

		items = append(items, item)
	}
	var reversItemes = []models.TodoThings{}
	for i := len(items) - 1; i >= 0; i-- {
		reversItemes = append(reversItemes, items[i])
	}

	var completeReverse = []models.TodoThings{}
	t.ExecuteTemplate(w, "todolist", sortItemsByStatus(reversItemes, completeReverse))
}

func sortItemsByStatus(inArray, outArray []models.TodoThings) *[]models.TodoThings {
	for i := 0; i < len(inArray); i++ {
		if !inArray[i].Status {
			outArray = append(outArray, inArray[i])
		}
	}

	for i := 0; i < len(inArray); i++ {
		if inArray[i].Status {
			outArray = append(outArray, inArray[i])
		}
	}
	return &outArray
}

func Save_item(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item != "" {
		insert := dbConnect(fmt.Sprintf("INSERT INTO `todo_item` (`item`, `status`) VALUES('%s', '0')", item))
		defer insert.Close()

		http.Redirect(w, r, "/todolist", http.StatusSeeOther)
	} else {
		fmt.Fprintf(w, "Пустое значение")
	}
}

func Remove_item(w http.ResponseWriter, r *http.Request) {
	status := r.FormValue("status")
	vars := mux.Vars(r)
	if status == "Не выполнена" {
		fmt.Print("нельзя удалят ьсука")
	}
	del := dbConnect(fmt.Sprintf("DELETE FROM `todo_item` WHERE `id` = '%s'", vars["id"]))

	defer del.Close()

	http.Redirect(w, r, "/todolist", http.StatusSeeOther)

}

func Done_item(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	upd := dbConnect(fmt.Sprintf("UPDATE `todo_item` SET `status`='1' WHERE `id` = '%s'", vars["id"]))

	defer upd.Close()

	http.Redirect(w, r, "/todolist", http.StatusSeeOther)
}
