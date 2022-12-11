package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testwithoutpackage/repo"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent *repo.Employee
	if r.Method == "POST" {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		}
		newEvent = repo.CreateEvent(reqBody)
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(*newEvent)
}

func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	emp := repo.GetOneEvent(eventID)
	json.NewEncoder(w).Encode(*emp)

}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {

	res := repo.GetAllEvents()

	json.NewEncoder(w).Encode(*res)
}
func UpdateEvent(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	repo.UpdateEvent(id, reqBody)
}
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	repo.DeleteEvent(id)
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
