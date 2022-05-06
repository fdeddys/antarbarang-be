package services

import (
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"
)

// MenuService ...
type MenuService struct {
}

// GetMenuByUser ...
func (h MenuService) GetMenuByUser(user string) []model.Menu {
	var res []model.Menu
	// var err error
	res, _ = repository.GetUserMenus(user)

	return res
}
