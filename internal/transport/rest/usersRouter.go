package rest

import (
	"diplomPlugService/internal/models"
	"diplomPlugService/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	arr, err := services.GetAllUsers()
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

func createNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserBody
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&user)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	err = services.CreateNewUser(user)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	w.Write([]byte("sucess"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := services.DeleteUser(id)
	if err != nil {
		errorHandler(err, w)
		return
	}
	w.Write([]byte("sucess"))
}

func changeUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserBody
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&user)
	if err != nil {
		errorHandler(err, w)
		return
	}
	// err = services.ChangeUser(user)
	// if err != nil {
	// 	errorHandler(err, w)
	// 	return
	// }
	// w.Write([]byte("success"))
}

func login(w http.ResponseWriter, r *http.Request) {
	var user models.UserLoginStruct
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&user)
	if err != nil {
		errorHandler(err, w) //todo errorsHandler
		return
	}
	tokenPair, err := services.Login(user)
	if err != nil {
		errorHandler(err, w)
		return
	}
	str, err := json.Marshal(tokenPair)
	if err != nil {
		errorHandler(err, w)
		return
	}
	w.Write(str)
}

func logout(w http.ResponseWriter, r *http.Request) {

}

func initUsersRouter(router *mux.Router) {
	router.HandleFunc("/users/get_all", getAllUsers).Methods("GET")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("GET")
	router.HandleFunc("/users/create", createNewUser).Methods("Post")
	router.HandleFunc("/users/delete/{id:[0-9]+}", deleteUser).Methods("Delete")
	router.HandleFunc("/users/change", changeUser).Methods("Put")
}
