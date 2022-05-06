package handler

import (
	"encoding/json"
	"net/http"

	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/services"
	"github.com/gorilla/mux"
)

var parameterService = new(services.ParameterService)

func ParamByNameHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	paramname := vars["paramname"]
	if paramname == "" {
		var res dto.ContentResponse
		res.ErrCode = constanta.ERR_CODE_04
		res.ErrDesc = constanta.ERR_CODE_04_PARAM_QUERY_STRING
		res.Contents = "param name tidak boleh kosong"
		return
	}
	resp := parameterService.GetByName(paramname)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
