package paginatejson

import (
	"errors"
	"paginatejson/lib"
	"regexp"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

func Fielding(listIface []interface{}, field string) (result []interface{}, err error) {
	fields := strings.Split(field, ",")
	if len(fields) == 0 {
		return
	}

	re := regexp.MustCompile(`[^A-z0-9_\.,]+`)

	columns := []string{}

	for idxField := range fields {
		itemField := fields[idxField]

		fieldByte := re.ReplaceAll([]byte(itemField), []byte(""))
		field := string(fieldByte)
		if lib.IsEmptyStr(field) {
			continue
		}

		columns = append(columns, field)
	}

	res, errField := fieldingJsonString(listIface, columns)
	if errField != nil {
		err = errField
		return
	}

	errMerge := lib.Merge(res, &result)
	if errMerge != nil {
		err = errors.New("gojsonq: reflect result is not equal with reflect input")
		return
	}

	return
}

func fieldingJsonString(listIface []interface{}, columns []string) (result interface{}, err error) {
	bte, errMarshal := lib.JSONMarshal(listIface)
	if errMarshal != nil {
		err = errMarshal
		return
	}

	jsonStr := string(bte)

	newJson := gojsonq.New().JSONString(jsonStr)
	if newJson.Count() == 0 {
		return
	}

	newJson.Select(columns...)

	res := newJson.Get()
	if newJson.Error() != nil {
		err = newJson.Error()
		return
	}

	result = res
	return
}
