package paginatejson

import (
	"errors"
	"paginatejson/lib"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

type SortOrder struct {
	Column    string
	Direction Direction
}

type Direction string

const (
	Ascending  Direction = "asc"
	Descending Direction = "desc"
)

func (d Direction) String() string {
	return string(d)
}

func Sorting(listOutstd []interface{}, sort, order string) (result []interface{}, err error) {
	finalSorts := []SortOrder{}

	sorts := strings.Split(sort, ",")
	for idxSort := range sorts {
		itemSort := sorts[idxSort]

		if lib.IsEmptyStr(itemSort) {
			continue
		}

		so := SortOrder{
			Column:    itemSort,
			Direction: Ascending,
		}

		if strings.EqualFold(order, Descending.String()) {
			so.Direction = Descending
		}

		if len(itemSort) >= 2 {
			if itemSort[:1] == "-" {
				so.Column = string(itemSort[1:])
				so.Direction = Descending
			}
		}

		finalSorts = append(finalSorts, so)
	}

	res, errSortJson := sortingJsonString(listOutstd, finalSorts)
	if errSortJson != nil {
		err = errSortJson
		return
	}

	errMerge := lib.Merge(res, &result)
	if errMerge != nil {
		err = errors.New("gojsonq: reflect result is not equal with reflect input")
		return
	}

	return
}

func sortingJsonString(input interface{}, sorts []SortOrder) (result interface{}, err error) {
	bte, errMarshal := lib.JSONMarshal(input)
	if errMarshal != nil {
		err = errMarshal
		return
	}

	jsonStr := string(bte)

	newJson := gojsonq.New().JSONString(jsonStr)
	if newJson.Count() == 0 {
		return
	}

	for idxSort := range sorts {
		itemSort := sorts[idxSort]
		newJson = newJson.SortBy(itemSort.Column, itemSort.Direction.String())
	}

	res := newJson.Get()
	if newJson.Error() != nil {
		err = newJson.Error()
		return
	}

	result = res
	return
}
