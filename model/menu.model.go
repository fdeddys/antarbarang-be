package model

// Menu ...
type Menu struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	ParentID    int64  `json:"parentId"`
	Icon        string `json:"icon"`
	Status      int    `json:"status"`
}
