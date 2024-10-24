package typefmt

import (
	"encoding/json"
	"slices"
	"strconv"
	"strings"
)

func StrToIntSlice(slc []string) []int {
	values := make([]int, len(slc))
	for i := range values {
		v, err := strconv.Atoi(slc[i])
		if err != nil {
			panic(err)
		}
		values[i] = int(v)
	}
	return values
}

func StrToInt(str string) (int, bool) {
	if str == "" {
		return 0, false
	}

	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return v, true
}

func StrToNilInt(str string) *int {
	var v *int
	if vParsed, err := strconv.Atoi(str); err != nil {
		v = nil
	} else {
		v = &vParsed
	}
	return v
}

func StrToNilStr(str string) *string {
	var v *string
	if str != "" {
		v = &str
	} else {
		v = nil
	}
	return v
}

func StrToNilBool(str string) *bool {
	var v *bool
	if vParsed, err := strconv.ParseBool(str); err != nil {
		v = nil
	} else {
		v = &vParsed
	}
	return v
}

func IntSliceToCommaString(slc []int) string {
	s := make([]string, len(slc))
	for i, id := range slc {
		s[i] = strconv.Itoa(id)
	}
	return strings.Join(s, ",")
}

func UniqueIntSlice(slc []int) []int {
	slices.Sort(slc)
	slc = slices.Compact(slc)
	return slc
}

func UniqueStrSlice(slc []string) []string {
	slices.Sort(slc)
	slc = slices.Compact(slc)
	return slc
}

func JsonArrayToIntSlice(str string) []int {
	var s []int
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		panic("json unmarshal failed")
	}
	return s
}
