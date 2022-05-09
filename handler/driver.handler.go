package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/services"
	"github.com/gorilla/mux"
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

func GetDriverByIdHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = err.Error()
		return
	}
	resp := driverService.GetDriverByID(id)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetDriverByCodeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	kode := vars["code"]
	if kode == "" {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = "code tidak boleh kosong"
		return
	}
	resp := driverService.GetDriverByKode(kode)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func LoginDriverHandler(w http.ResponseWriter, r *http.Request) {

	var loginRequest dto.LoginRequestDto

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &loginRequest)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}

	resp := driverService.LoginDriverByKode(loginRequest.Kode, loginRequest.Password)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DriverUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &driver)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct " + err.Error()))
		return
	}
	res := driverService.UpdateDriver(driver)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func DriverUpdateStatusHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	driverId := vars["driver-id"]
	statusActive := vars["active"]
	var active bool

	if driverId == "" {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = "id tidak boleh kosong"
		return
	}

	seller, err := strconv.ParseInt(driverId, 10, 64)
	if err != nil {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = "code tidak valid"
		return
	}

	if statusActive == "1" {
		active = true
	} else {
		active = false
	}

	res := driverService.UpdateStatusDriverActive(seller, active)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func DriverChangePasswordHandler(w http.ResponseWriter, r *http.Request) {

	var driver model.Driver

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &driver)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct " + err.Error()))
		return
	}

	res := driverService.ChangePasswordDriver(driver)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func GetDriverPageHandler(w http.ResponseWriter, r *http.Request) {
	var searchRequestDto dto.SearchRequestDto

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &searchRequestDto)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	vars := mux.Vars(r)

	pageStr := vars["page"]
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		fmt.Println("Error Page must be number", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error Page must be number "))
		return
	}

	countStr := vars["count"]
	count, errCount := strconv.Atoi(countStr)
	if errCount != nil {
		fmt.Println("Error Count must be number", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error Count must be number "))
		return
	}

	res := driverService.SearchDriverPage(searchRequestDto, page, count)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
