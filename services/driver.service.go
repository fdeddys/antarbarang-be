package services

import (
	"fmt"
	"time"

	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/model"
)

type DriverService struct {
}

func (d *DriverService) CreateDriver() dto.ContentResponse {

	return dto.ContentResponse{}
}

func (d *DriverService) GetDriverByID() dto.ContentResponse {
	t := time.Now().Unix()
	return dto.ContentResponse{
		ErrCode: "00",
		ErrDesc: "OK",
		Contents: model.Driver{
			ID:           0,
			Picture:      "",
			Address:      "",
			Hp:           "",
			Name:         "",
			Status:       0,
			LastUpdateBy: fmt.Sprintf("%v", time.Unix(0, time.Now().UnixNano())),
			LastUpdate:   fmt.Sprintf("%v", time.Unix(t, 0)),
			Code:         "",
		},
	}
}
