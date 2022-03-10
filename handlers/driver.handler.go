package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"com.ddabadi.antarbarang/services"
	"github.com/gorilla/mux"
)

var driverService = new(services.DriverService)

func DriverCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
	fmt.Println(w, "driver : ", vars["driver"])
}

func DriverHandler(w http.ResponseWriter, r *http.Request) {

	driver := driverService.GetDriverByID()
	result, _ := json.Marshal(driver)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
