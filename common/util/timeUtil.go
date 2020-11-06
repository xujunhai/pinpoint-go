package util

import (
	"math/rand"
	"time"
)

func GetNowMSec() int64 {
	return time.Now().UnixNano() / 1000 / 1000
}

func GetRandomChar(num int) string {
	//0-9a-zA-Z
	count := 10 + 26 + 26

	str := ""
	for i := 0; i < num; i++ {
		index := rand.Intn(count)
		if index < 10 {
			str += string('0' + index)
		} else if index < 36 {
			str += string('a' + index - 10)
		} else {
			str += string('A' + index - 36)
		}
	}
	return str
}
