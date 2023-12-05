package filterrequest

import "reflect"

func IsFirstIdx(idx int) bool {
	return idx == 0
}

func IsMiddleIdx(idx, len int) bool {
	return !IsFirstIdx(idx) && !IsLastIdx(idx, len)
}

func IsLastIdx(idx, len int) bool {
	return idx == len-1
}

func IsSlice(input interface{}) (isSlice bool) {
	isSlice = reflect.TypeOf(input).Kind() == reflect.Slice
	return
}
