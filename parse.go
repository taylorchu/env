package env

import (
	"os"
	"strconv"
	"time"
)

func String(key string, fallback string) string {
	value, found := os.LookupEnv(key)
	if !found {
		value = fallback
	}
	return value
}

func Duration(key string, fallback string) time.Duration {
	d, _ := time.ParseDuration(String(key, fallback))
	return d
}

func Bool(key string, fallback string) bool {
	b, _ := strconv.ParseBool(String(key, fallback))
	return b
}

func Float64(key string, fallback string) float64 {
	f, _ := strconv.ParseFloat(String(key, fallback), 64)
	return f
}

func Int64(key string, fallback string) int64 {
	i, _ := strconv.ParseInt(String(key, fallback), 10, 64)
	return i
}

func Uint64(key string, fallback string) uint64 {
	i, _ := strconv.ParseUint(String(key, fallback), 10, 64)
	return i
}
