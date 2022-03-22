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

var sellerService = new(services.SellerService)

func SaveSellerHandler(w http.ResponseWriter, r *http.Request) {
	var seller model.Seller

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &seller)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := sellerService.CreateSeller(seller)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func GetSellerByIDHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = err.Error()
		return
	}
	resp := sellerService.GetSellerByID(id)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetSellerByCodeHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	kode := vars["code"]
	if kode == "" {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = "code tidak boleh kosong"
		return
	}
	resp := sellerService.GetSellerByKode(kode)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func LoginSellerHandler(w http.ResponseWriter, r *http.Request) {

	var loginRequest dto.LoginRequestDto

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &loginRequest)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}

	resp := sellerService.LoginSellerByKode(loginRequest.Kode, loginRequest.Password)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
