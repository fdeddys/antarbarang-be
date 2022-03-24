package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/services"
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
