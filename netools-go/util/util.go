package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1
	}
	return i
}

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func GetStrTimestamp() string {
	return Int64ToString(GetTimestamp())
}
func GetStrTimestampByInt(t int64) string {
	return Int64ToString(t)
}

func IsBlank(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

func GetRandomString(length int) (res string) {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	cLen := len(charset)
	for i := 0; i < length; i++ {
		res += string(charset[rand.Intn(cLen)])
	}
	return
}

func GetFileExtension(filename string) (ext string) {
	ls := strings.Split(filename, ".")
	if len(ls) == 0 {
		return ""
	}
	return ls[len(ls)-1]
}
