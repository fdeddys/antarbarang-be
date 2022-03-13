package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/services"
)

var driverService = new(services.DriverService)

func DriverCreateHandler(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &driver)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := driverService.CreateDriver(driver)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func DriverHandler(w http.ResponseWriter, r *http.Request) {

	driver := driverService.GetDriverByID()
	result, _ := json.Marshal(driver)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
