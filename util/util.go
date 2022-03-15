package util

import "time"

func GetCurrTimeUnix() int64 {

	return time.Now().Unix()
}
