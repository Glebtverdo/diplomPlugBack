package restTransport

import (
	"diplomPlugService/internal/models"
	"diplomPlugService/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllRequests(w http.ResponseWriter, r *http.Request) {
	arr, err := services.GetAllRequests()
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	str, err := json.Marshal(arr)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	w.Write(str)
}

func createNewRequest(w http.ResponseWriter, r *http.Request) {
	var obj models.RequestBody
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&obj)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	err = services.CreateNewRequest(obj)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	w.Write([]byte("sucess"))
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := services.DeleteRequest(id) //todo errorsHandler
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	w.Write([]byte("sucess"))
}

func changeRequest(w http.ResponseWriter, r *http.Request) {
	var obj models.Request
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&obj)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	err = services.ChangeRequest(obj)
	if err != nil {
		errorHandler(err, w)
		return
	}
	w.Write([]byte("success"))
}

func initRequestsRouter(router *mux.Router) {
	router.HandleFunc("/requests/get_all", getAllRequests).Methods("GET")
	router.HandleFunc("/requests/create", createNewRequest).Methods("Post")
	router.HandleFunc("/requests/delete/{id:[0-9]+}", deleteRequest).Methods("Delete")
	// router.HandleFunc("/requests/change", changeRequest).Methods("Put")
	router.HandleFunc("/requests/get_users_request", changeRequest).Methods("GET")
}
