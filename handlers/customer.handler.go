package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/services"
	"github.com/gorilla/mux"
)

var customerService = new(services.CustomerService)

func CustomerCreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
	fmt.Println(w, "catg : ", vars["category"])
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
	vars := mux.Vars(r)

	nama := vars["nama"]
	if nama == "" {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = ""
		return
	}
	resp := customerService.GetCustomerByNama(nama)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")

	w.WriteHeader(http.StatusOK)
	w.Write(result)
	// fmt.Println(w, "catg : ", result)
}
