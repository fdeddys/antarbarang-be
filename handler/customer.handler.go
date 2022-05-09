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

var customerService = new(services.CustomerService)

func CustomerCreateHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("ok"))
	// fmt.Println(w, "catg : ", vars["category"])

	var customer model.Customer

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &customer)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct " + err.Error()))
		return
	}
	res := customerService.CreateCustomer(customer)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))

}

func GetCustomerByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = err.Error()
		return
	}
	resp := customerService.GetCustomerByID(id)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")

	w.WriteHeader(http.StatusOK)
	w.Write(result)
	// fmt.Println(w, "catg : ", result)
}

func GetCustomerBySellerIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["sellerId"], 10, 64)
	if err != nil {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = err.Error()
		return
	}
	resp := customerService.GetCustomerBySellerId(id)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")

	w.WriteHeader(http.StatusOK)
	w.Write(result)
	// fmt.Println(w, "catg : ", result)
}

func GetCustomerByNamaHandler(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &customer)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct " + err.Error()))
		return
	}

	resp := customerService.GetCustomerByNama(customer.Nama)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")

	w.WriteHeader(http.StatusOK)
	w.Write(result)
	// fmt.Println(w, "catg : ", result)
}

func CustomerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &customer)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct " + err.Error()))
		return
	}
	res := customerService.UpdateCustomer(customer)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func CustomerUpdateStatusHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customerId := vars["customer-id"]
	statusActive := vars["active"]
	var active bool

	if customerId == "" {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = "code tidak boleh kosong"
		return
	}

	seller, err := strconv.ParseInt(customerId, 10, 64)
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

	res := customerService.UpdateStatusCustomerActive(seller, active)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func GetCustomerPageHandler(w http.ResponseWriter, r *http.Request) {
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

	res := customerService.SearchCustomerPage(searchRequestDto, page, count)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
