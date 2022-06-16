package util

import "time"

func GetCurrTimeUnix() int64 {

	return time.Now().Unix()
}

func GetCurrDate() time.Time {

	return time.Now()
}

func DateTimeUnixToString(intTime int64) string {
	t := time.Unix(intTime, 0)
	layout := "02-Jan-2006- 15:04:05"

	return t.Format(layout)
}

func DateUnixToString(intTime int64) string {

	layout := "02-Jan-2006"
	// tim, _ := time.Parse(layout, "Mon, 23 Dec 2019 18:52:45 GMT")

	t := time.Unix(intTime, 0)

	return t.UTC().Format(layout)
}
