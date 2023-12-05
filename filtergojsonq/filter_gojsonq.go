package filtergojsonq

import (
	"github.com/thedevsaddam/gojsonq/v2"
)

type IJsonSerializer interface {
	JsonMarshal() func(v interface{}) ([]byte, error)
	JsonUnmarshal() func(data []byte, v interface{}) error
}

type BuilderFilterGojsonq struct {
	FromInterface     []interface{}
	ListFilterRequest ListFilterRequest
	JsonSerializer    IJsonSerializer
}

type ListFilterRequest []FilterRequest

type FilterRequest struct {
	Conjunction GojsonqConjunction `json:"conjunction,omitempty"`
	Column      string             `json:"column,omitempty"`
	Operator    GojsonqOperator    `json:"operator,omitempty"`
	Value       GojsonqValue       `json:"value,omitempty"`
	ValueKind   ValueKind          `json:"value_kind,omitempty"`
}

type FilterGojsonq struct {
	builder  BuilderFilterGojsonq
	fromJson string
	result   interface{}
	err      error
}

func NewFilterGojsonq(builder BuilderFilterGojsonq) (fgj *FilterGojsonq) {
	fgj = new(FilterGojsonq)

	// builder
	fgj.setBuilder(builder)

	// fromJson
	fromJson, errFromJson := fgj.genFromJson()
	if errFromJson != nil {
		fgj.setErr(errFromJson)
		return
	}
	fgj.setFromJson(fromJson)

	// result
	result, errResult := fgj.genResult()
	if errResult != nil {
		fgj.setErr(errResult)
		return
	}
	fgj.setResult(result)

	return
}

func (fgj FilterGojsonq) Result() (result interface{}) {
	result = fgj.getResult()
	return
}

func (fgj FilterGojsonq) Err() (err error) {
	err = fgj.getErr()
	return
}

func (fgj *FilterGojsonq) setBuilder(builder BuilderFilterGojsonq) {
	fgj.builder = builder
}

func (fgj FilterGojsonq) getBuilder() (builder BuilderFilterGojsonq) {
	builder = fgj.builder
	return
}

func (fgj *FilterGojsonq) setFromJson(fromJson string) {
	fgj.fromJson = fromJson
}

func (fgj FilterGojsonq) getFromJson() (fromJson string) {
	fromJson = fgj.fromJson
	return
}

func (fgj FilterGojsonq) genFromJson() (fromJson string, err error) {
	builder := fgj.getBuilder()
	fromIface := builder.FromInterface
	jsonSerializer := builder.JsonSerializer
	jsonMarshal := jsonSerializer.JsonMarshal()

	bte, errMarshal := jsonMarshal(fromIface)
	if errMarshal != nil {
		err = errMarshal
		return
	}

	fromJson = string(bte)
	return
}

func (fgj *FilterGojsonq) setResult(result interface{}) {
	fgj.result = result
}

func (fgj FilterGojsonq) getResult() (result interface{}) {
	result = fgj.result
	return
}

func (fgj FilterGojsonq) genResult() (result interface{}, err error) {
	builder := fgj.getBuilder()
	rawIface := builder.FromInterface
	listFilterReq := builder.ListFilterRequest

	if len(listFilterReq) == 0 {
		result = rawIface
		return
	}

	fromJson := fgj.getFromJson()
	newJson := gojsonq.New().JSONString(fromJson)

	for idxReq := range listFilterReq {
		itemReq := listFilterReq[idxReq]
		fgj.genQuery(newJson, idxReq, itemReq)
	}

	res := newJson.Get()
	if newJson.Error() != nil {
		err = newJson.Error()
		return
	}

	result = res
	return
}

func (fgj FilterGojsonq) genQuery(newJson *gojsonq.JSONQ, idxReq int, itemReq FilterRequest) {
	conjunction := itemReq.Conjunction
	column := itemReq.Column
	operator := itemReq.Operator
	value := itemReq.Value
	valueKind := itemReq.ValueKind

	// get real value type
	var finalVal interface{}

switchValKind:
	switch valueKind {
	case SliceFloat64ValKind:
		{
			finalVal = value.SliceFloat64
			break switchValKind
		}
	case SliceStringValKind:
		{
			finalVal = value.SliceString
			break switchValKind
		}
	case Float64ValKind:
		{
			finalVal = value.Float64
			break switchValKind
		}
	case StringValKind:
		{
			finalVal = value.String
			break switchValKind
		}
	case NilValKind:
		{
			finalVal = nil
			break switchValKind
		}
	}

	// generate query
	isFirstIdx := idxReq == 0
	conjunctionAnd := conjunction == AndGojsConj
	if isFirstIdx || conjunctionAnd {
		newJson.Where(column, operator.String(), finalVal)
		return
	}

	newJson.OrWhere(column, operator.String(), finalVal)
}

func (fgj *FilterGojsonq) setErr(err error) {
	fgj.err = err
}

func (fgj *FilterGojsonq) clearErr() {
	fgj.err = nil
}

func (fgj FilterGojsonq) getErr() (err error) {
	err = fgj.err
	return
}
