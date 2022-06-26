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

var regionalService = new(services.RegionalService)

func RegionalCreateHandler(w http.ResponseWriter, r *http.Request) {
	var regional model.Regional

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &regional)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := regionalService.CreateRegional(regional)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func RegionalUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var regional model.Regional

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &regional)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct " + err.Error()))
		return
	}
	res := regionalService.UpdateRegional(regional)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func RegionalByGroupRegionalIdHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["groupRegionalId"])
	if err != nil {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = err.Error()
		return
	}
	resp := regionalService.GetRegionalByRegionalGroupId(id)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetRegionalPageHandler(w http.ResponseWriter, r *http.Request) {
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

	res := regionalService.SearchRegionalPage(searchRequestDto, page, count)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
