package rest

import (
	"diplomPlugService/internal/transport/midlewares"
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
	initObjectsRouter(router)
	initUsersRouter(router)
	router.Use(midlewares.CheckAuthorization)
	log.Fatal(http.ListenAndServe(":8000", router))
}
