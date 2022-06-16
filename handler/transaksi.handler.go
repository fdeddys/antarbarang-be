package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/services"
	"github.com/gorilla/mux"
)

var transaksiService = new(services.TransaksiService)

func NewTransaksiHandler(w http.ResponseWriter, r *http.Request) {
	var transaksi model.Transaksi

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &transaksi)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := transaksiService.CreateNew(transaksi)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func OnProccessHandler(w http.ResponseWriter, r *http.Request) {
	var transaksi model.Transaksi

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &transaksi)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := transaksiService.OnProccess(transaksi)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func OnTheWayHandler(w http.ResponseWriter, r *http.Request) {
	var transaksi model.Transaksi

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &transaksi)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := transaksiService.OnTheWay(transaksi)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func DoneProcessHandler(w http.ResponseWriter, r *http.Request) {
	var transaksi model.Transaksi

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &transaksi)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := transaksiService.DoneProcess(transaksi)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func GetTransaksiPageHandler(w http.ResponseWriter, r *http.Request) {
	var searchRequestDto dto.SearchTransaksiRequestDto

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

	res := transaksiService.SearchTransaksiPage(searchRequestDto, page, count)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func GetTransaksiAntarHandlerByDriverByTanggalAntar(w http.ResponseWriter, r *http.Request) {
	var searchRequestDto dto.SearchTransaksiRequestDto

	dataBodyReq, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(dataBodyReq, &searchRequestDto)

	if err != nil {
		fmt.Println("Error Struct", err.Error())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Error struct "))
		return
	}
	res := transaksiService.SearchTransaksiByDriverByTanggalAntar(searchRequestDto)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func GetTransaksiByTglAntarPageHandler(w http.ResponseWriter, r *http.Request) {
	var searchRequestDto dto.SearchTransaksiRequestDto

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

	res := transaksiService.SearchTransaksiByTglAntarPage(searchRequestDto, page, count)
	result, _ := json.Marshal(res)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
