package filterrequest

import (
	"paginatejson/lib"
)

type RequestValue struct {
	String string
	// Int          int
	Float64     float64
	SliceString []string
	// SliceInt     []int
	SliceFloat64 []float64
}

func GenRequestValue(input interface{}, kind ValueKind) (rv RequestValue, err error) {
	switch kind {
	case StringValKind:
		{
			newVal := ""
			lib.Merge(input, &newVal)
			rv.String = newVal
			break
		}
	// case IntValKind:
	// 	{
	// 		newVal, ok := input.(int)
	// 		if !ok {
	// 			err = fmt.Errorf("gen request value: cannot cast value %+v with kind %s", input, kind.String())
	// 			break
	// 		}

	// 		rv.Int = newVal
	// 		break
	// 	}
	case Float64ValKind:
		{
			newVal := float64(0)
			lib.Merge(input, &newVal)
			rv.Float64 = newVal
			break
		}
	case SliceStringValKind:
		{
			newVal := []string{}
			lib.Merge(input, &newVal)
			rv.SliceString = newVal
			break
		}
	// case SliceIntValKind:
	// 	{
	// 		newVal, ok := input.([]int)
	// 		if !ok {
	// 			err = fmt.Errorf("gen request value: cannot cast value %+v with kind %s", input, kind.String())
	// 			break
	// 		}

	// 		rv.SliceInt = newVal
	// 		break
	// 	}
	case SliceFloat64ValKind:
		{
			newVal := []float64{}
			lib.Merge(input, &newVal)
			rv.SliceFloat64 = newVal
			break
		}
	}

	return
}
