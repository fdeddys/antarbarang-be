package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/services"
)

var menuService = new(services.MenuService)

func GetMenuByUsernameHandler(w http.ResponseWriter, r *http.Request) {

	username := dto.CurrUser

	fmt.Println("Get user list by user name : ", dto.CurrUser)
	resp := menuService.GetMenuByUser(username)
	result, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application-json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
