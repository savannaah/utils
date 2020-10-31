package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/savannaah/utils/sxint32"
	"github.com/savannaah/utils/sxstring"
	"math"
	"reflect"
)

// Base64Encode takes in a string and returns a base 64 encoded string
func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

// Base64Decode takes in a base 64 encoded string and returns the //actual string or an error of it fails to decode the string
func Base64Decode(src string) (string, error) {
	if len(src) == 0 {
		return "", fmt.Errorf("cannot decode empty string, occurred in utils package")
	}
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Contains(list interface{}, value interface{}) bool {
	arr := reflect.ValueOf(list)

	if arr.Kind() != reflect.Array {
		panic("invalid data-type, occurred in utils package")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == value {
			return true
		}
	}
	return false
}

func Identical(a, b interface{}) bool {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(a), reflect.TypeOf(b)) {
	case "*string:*string":
		return sxstring.Indentical(a.(*string), b.(*string))
	case "*int32:*int32":
		return sxint32.Indentical(a.(*int32), b.(*int32))
	default:
		panic("invalid data-type, occurred in utils package")
	}
	return false
}

func Equal(a, b interface{}) bool {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(a),reflect.TypeOf(b)) {
	case "[]string:[]string":
		return sxstring.Equal(a.([]string), b.([]string))
	case "[]int32:[]int32":
		return sxint32.Equal(a.([]int32), b.([]int32))
	default:
		panic("invalid data-type, occurred in utils package")
	}
	return false
}

func Unique(a interface{}) interface{} {
	switch fmt.Sprintf("%s", reflect.TypeOf(a)) {
	case "[]string":
		return sxstring.Unique(a.([]string))
	case "[]int32":
		return sxint32.Unique(a.([]int32))
	default:
		panic("invalid data-type, occurred in utils package")
	}
	return nil
}

//returns missing elements in b compared to a
func Missing(a, b interface{}) interface{} {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(a), reflect.TypeOf(b)) {
	case "[]string:[]string":
		return sxstring.Missing(a.([]string), b.([]string))
	case "[]int32:[]int32":
		return sxint32.Missing(a.([]int32), b.([]int32))
	default:
		panic("invalid data-type, occurred in utils package")
	}
	return nil
}

//returns uncommon elements
func Unmatched(a, b interface{}) interface{} {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(a), reflect.TypeOf(b)) {
	case "[]string:[]string":
		return sxstring.Unmatched(a.([]string), b.([]string))
	case "[]int32:[]int32":
		return sxint32.Unmatched(a.([]int32), b.([]int32))
	default:
		panic("invalid data-type, occurred in utils package")
	}
	return nil
}

func CheckDecimalPlaces(place int, value float64) bool {
	valueF := value * math.Pow(10.0, float64(place))
	extra := valueF - float64(int(valueF))

	return extra == 0
}
