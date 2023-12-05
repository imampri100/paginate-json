package filterrequest

import (
	"fmt"
	"paginatejson/lib"
	"reflect"
	"strings"
)

type ValueKind string

const (
	_             ValueKind = ""
	StringValKind ValueKind = "string"
	// IntValKind          ValueKind = "int"
	Float64ValKind     ValueKind = "float64"
	SliceStringValKind ValueKind = "[]string"
	// SliceIntValKind     ValueKind = "[]int"
	SliceFloat64ValKind ValueKind = "[]float64"
	NilValKind          ValueKind = "nil"
)

func GenValueKind(input interface{}) (vkd ValueKind, err error) {
	switch {
	case NilValKind.IsMatch(input):
		{
			vkd = NilValKind
			break
		}
	case SliceFloat64ValKind.IsMatch(input):
		{
			vkd = SliceFloat64ValKind
			break
		}
	// case SliceIntValKind.IsMatch(input):
	// 	{
	// 		vkd = SliceIntValKind
	// 		break
	// 	}
	case SliceStringValKind.IsMatch(input):
		{
			vkd = SliceStringValKind
			break
		}
	// case IntValKind.IsMatch(input):
	// 	{
	// 		vkd = IntValKind
	// 		break
	// 	}
	case Float64ValKind.IsMatch(input):
		{
			vkd = Float64ValKind
			break
		}
	case StringValKind.IsMatch(input):
		{
			vkd = StringValKind
			break
		}
	default:
		{
			listVals := []string{
				StringValKind.String(),
				// IntValKind.String(),
				Float64ValKind.String(),
				SliceStringValKind.String(),
				// SliceIntValKind.String(),
				Float64ValKind.String(),
				NilValKind.String(),
			}
			listValsStr := strings.Join(listVals, ",")

			kind := reflect.TypeOf(input)

			err = fmt.Errorf("gen request value kind: cannot handling value %+v type of %s. You can only use one of value type, such as: %s", input, kind, listValsStr)
			break
		}
	}
	return
}

func (vkd ValueKind) IsMatch(input interface{}) (isMatch bool) {
	if input == nil {
		isMatch = vkd == NilValKind
		return
	}

	switch vkd {
	case SliceFloat64ValKind:
		{
			sliceFloat64 := []float64{}
			err := lib.Merge(input, &sliceFloat64)
			isMatch = err == nil
			break
		}
	// case SliceIntValKind:
	// 	{
	// 		sliceInt := []int{}
	// 		err := lib.Merge(input, &sliceInt)
	// 		isMatch = err == nil
	// 		break
	// 	}
	case SliceStringValKind:
		{
			sliceString := []string{}
			err := lib.Merge(input, &sliceString)
			isMatch = err == nil
			break
		}
	case Float64ValKind:
		{
			float64 := float64(0)
			err := lib.Merge(input, &float64)
			isMatch = err == nil
			break
		}
	// case IntValKind:
	// 	{
	// 		int := int(0)
	// 		err := lib.Merge(input, &int)
	// 		isMatch = err == nil
	// 		break
	// 	}
	case StringValKind:
		{
			string := ""
			err := lib.Merge(input, &string)
			isMatch = err == nil
			break
		}
	}

	return
}

func (vkd ValueKind) String() string {
	return string(vkd)
}

func (vkd ValueKind) IsMultiple() (isMultiple bool) {
	switch vkd {
	case SliceStringValKind,
		// SliceIntValKind,
		SliceFloat64ValKind:
		{
			isMultiple = true
			break
		}
	default:
		{
			isMultiple = false
			break
		}
	}
	return
}
