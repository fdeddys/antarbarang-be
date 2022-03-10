package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CustomerCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
	fmt.Println(w, "catg : ", vars["category"])
}

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
	fmt.Println(w, "catg : ", vars["category"])
}
