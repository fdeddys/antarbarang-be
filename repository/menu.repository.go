package repository

import (
	"fmt"

	"com.ddabadi.antarbarang/database"
	"com.ddabadi.antarbarang/model"
)

// GetUserMenus ...
func GetUserMenus(user string) ([]model.Menu, error) {

	db := database.GetConn
	// defer db.Close()

	sqlStatement := `
		select d.id, d.name, d.description, link, icon, parent_id, d.status
		from admins a join
			m_role_user b on (a.id = b.user_id) join
			m_role_menu c on(b.role_id = c.role_id) join
			m_menus d on(c.menu_id = d.id)
		where a.nama = ? and d.status = 1 and c.status = 1
		group by d.id, a.nama
		order by d.ordering;
	`
	var menus []model.Menu

	datas, err := db().
		Query(sqlStatement, user)

	if err != nil {
		fmt.Println("Error query", err.Error())
		return menus, err
	}

	for datas.Next() {
		// fmt.Println("menu ", menus[0].ID)
		var menu model.Menu

		err = datas.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Description,
			&menu.Link,
			&menu.Icon,
			&menu.ParentID,
			&menu.Status,
		)
		menus = append(menus, menu)
	}
	return menus, nil

}
