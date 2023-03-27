package middlewares

import (
	"math/rand"
	"time"
)

func RandomAtoZWith0to9(count int) string {
	str := "abcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	var returnStr string
	for i := 1; i <= count; i++ {
		returnStr += (string([]rune(str)[rand.Intn(36)]))
	}
	return returnStr
}
