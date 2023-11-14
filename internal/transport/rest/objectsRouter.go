package rest

import (
	"diplomPlugService/internal/models"
	"diplomPlugService/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllObjs(w http.ResponseWriter, r *http.Request) {
	arr, err := services.GetAllObjs()
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

func createNewObj(w http.ResponseWriter, r *http.Request) {
	var obj models.ObjectBody
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&obj)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	err = services.CreateNewObject(obj)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	w.Write([]byte("sucess"))
}

func deleteObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := services.DeleteObject(id) //todo errorsHandler
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	w.Write([]byte("sucess"))
}

func changeObject(w http.ResponseWriter, r *http.Request) {
	var obj models.Object
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&obj)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	err = services.ChangeObj(obj)
	if err != nil {
		errorHandler(err, w)
		return
	}
	w.Write([]byte("success"))
}

func initObjectsRouter(router *mux.Router) {
	router.HandleFunc("/objects/get_all", getAllObjs).Methods("GET")
	router.HandleFunc("/objects/create", createNewObj).Methods("Post")
	router.HandleFunc("/objects/delete/{id:[0-9]+}", deleteObject).Methods("Delete")
	router.HandleFunc("/objects/change", changeObject).Methods("Put")
}
