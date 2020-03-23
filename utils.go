package utils

import (
	"fmt"
	"github.com/savannaah/utils/sxint32"
	"github.com/savannaah/utils/sxstring"
	"math"
	"reflect"
)

func Contains(list interface{}, value interface{}) bool {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(value), reflect.TypeOf(list)) {
	case "string:[]string":
		return sxstring.Contains(list.([]string), value.(string))
	case "int32:[]int32":
		return sxint32.Contains(list.([]int32), value.(int32))
	default:
		return false
	}
}

func Identical(a, b interface{}) bool {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(a), reflect.TypeOf(b)) {
	case "*string:*string":
		return sxstring.Indentical(a.(*string), b.(*string))
	case "*int32:*int32":
		return sxint32.Indentical(a.(*int32), b.(*int32))
	default:
		return false
	}
}

func Equal(a, b interface{}) bool {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(a),reflect.TypeOf(b)) {
	case "[]string:[]string":
		return sxstring.Equal(a.([]string), b.([]string))
	case "[]int32:[]int32":
		return sxint32.Equal(a.([]int32), b.([]int32))
	default:
		return false
	}
}

func Unique(a interface{}) interface{} {
	switch fmt.Sprintf("%s", reflect.TypeOf(a)) {
	case "[]string":
		return sxstring.Unique(a.([]string))
	case "[]int32":
		return sxint32.Unique(a.([]int32))
	default:
		return nil
	}
}

func Missing(a, b interface{}) interface{} {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(a), reflect.TypeOf(b)) {
	case "[]string:[]string":
		return sxstring.Missing(a.([]string), b.([]string))
	case "[]int32:[]int32":
		return sxint32.Missing(a.([]int32), b.([]int32))
	default:
		return nil
	}
}

func Unmatched(a, b interface{}) interface{} {
	switch fmt.Sprintf("%s:%s", reflect.TypeOf(a), reflect.TypeOf(b)) {
	case "[]string:[]string":
		return sxstring.Unmatched(a.([]string), b.([]string))
	case "[]int32:[]int32":
		return sxint32.Unmatched(a.([]int32), b.([]int32))
	default:
		return nil
	}
}

func CheckDecimalPlaces(place int, value float64) bool {
	valueF := value * float64(math.Pow(10.0, float64(place)))
	extra := valueF - float64(int(valueF))

	return extra == 0
}
