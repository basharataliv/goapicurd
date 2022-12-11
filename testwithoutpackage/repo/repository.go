package repo

import (
	"encoding/json"
	"log"
	con "testwithoutpackage/repo/dbCon"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id   int    `json:"ID"`
	Name string `json:"Name"`
	City string `json:"City"`
}

func CreateEvent(reqBody []byte) *Employee {

	db := con.DbConn()
	var newEvent Employee

	json.Unmarshal(reqBody, &newEvent)
	name := newEvent.Name
	city := newEvent.City
	insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(name, city)
	log.Println("INSERT: Name: " + name + " | City: " + city)

	defer db.Close()
	return &newEvent
}

func GetOneEvent(eventID string) *Employee {

	db := con.DbConn()
	//nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", eventID)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	return &emp

}

func GetAllEvents() *[]Employee {

	db := con.DbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		res = append(res, emp)
	}
	return &res
}

func UpdateEvent(id string, reqBody []byte) {
	db := con.DbConn()
	var updatedEvent Employee
	json.Unmarshal(reqBody, &updatedEvent)
	name := updatedEvent.Name
	city := updatedEvent.City

	insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(name, city, id)
	log.Println("UPDATE: Name: " + name + " | City: " + city)

	defer db.Close()
}

func DeleteEvent(id string) {
	db := con.DbConn()
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(id)
	log.Println("DELETE")
	defer db.Close()
}
