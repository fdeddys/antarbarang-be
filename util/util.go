package util

import "time"

func GetCurrTimeUnix() int64 {

	return time.Now().Unix()
}

func DateUnixToString(intTime int64) string {
	t := time.Unix(intTime, 0)
	layout := "2006-01-02 15:04:05"

	return t.Format(layout)
}
