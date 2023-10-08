package env

import (
	"os"
	"strconv"
)

func ReadStr(name string, fallback string) string {
	val, d := os.LookupEnv(name)
	if !d {
		return fallback
	}
	return val
}

func ReadBool(name string, fallback bool) bool {
	str := ReadStr(name, "")
	b, err := strconv.ParseBool(str)
	if err != nil {
		return fallback
	}
	return b
}
