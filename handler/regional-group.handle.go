package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/services"
)

var regionalGroupService = new(services.RegionalGroupService)

func RegionalGroupCreateHandler(w http.ResponseWriter, r *http.Request) {
	var regionalGroup model.RegionalGroup

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &regionalGroup)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := regionalGroupService.CreateRegionalGroup(regionalGroup)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func RegionalGroupUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var regionalGroup model.RegionalGroup

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("Body ", r.Body)
	err := json.Unmarshal(dataBodyReq, &regionalGroup)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct " + err.Error()))
		return
	}
	res := regionalGroupService.UpdateRegionalGroup(regionalGroup)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func RegionalGroupAllHandler(w http.ResponseWriter, r *http.Request) {

	resp := regionalGroupService.GetAllRegionalGroup()
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
