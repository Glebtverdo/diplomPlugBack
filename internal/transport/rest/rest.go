package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer() {
	router := mux.NewRouter()
	router.Schemes("https")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hellow world"))
	})
	// initObjectsRouter(router)
	// initUsersRouter(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
