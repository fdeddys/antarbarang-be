package util

import "time"

func GetCurrTimeUnix() int64 {

	return time.Now().Unix()
}

func DateTimeUnixToString(intTime int64) string {
	t := time.Unix(intTime, 0)
	layout := "02-Jan-2006- 15:04:05"

	return t.Format(layout)
}

func DateUnixToString(intTime int64) string {
	t := time.Unix(intTime, 0)
	layout := "02-Jan-2006"

	return t.Format(layout)
}
