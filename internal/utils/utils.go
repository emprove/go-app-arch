package utils

import (
	"fmt"
	"math"
	"strconv"
)

const (
	StorageStaticPath = "/storage"
)

func GetResourceStorageUrl(path string, base string) string {
	if path == "" {
		return ""
	}

	return base + StorageStaticPath + "/" + path
}

func RoundPrice(price float64) float64 {
	return math.Round(price)
}

func AdjustDecimals(num float64, precision int) (float64, error) {
	formatString := fmt.Sprintf("%%.%df", precision)
	strNum := fmt.Sprintf(formatString, num)
	return strconv.ParseFloat(strNum, 64)
}

func GetOffset(page int, limit int) int {
	offset := (limit * page) - limit
	return offset
}

func Filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
