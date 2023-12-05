package filterrequest

import (
	"fmt"
	"paginatejson/lib"
)

type BuilderListWhereClause struct {
	ConfigRequest IConfigRequest
	Input         ListConjunctionWhere
}

type ListWhereClause []WhereClause

/*
input possibilities
- slice

condition possibilties
- first idx: conjunction
- first idx: column, second idx: value
- first idx: column, second idx: operator, third idx: value
*/
func GenListWhereClause(builder BuilderListWhereClause) (lwc *ListWhereClause, err error) {
	lwc = new(ListWhereClause)

	configRequest := builder.ConfigRequest
	input := builder.Input

	for idxConjWhe := range input {
		itemConjWhe := input[idxConjWhe]

		builderWcl := BuilderWhereClause{
			ConfigRequest: configRequest,
		}

		wcl := GenWhereClause(builderWcl, itemConjWhe)
		if wcl.Err() != nil {
			err = wcl.Err()
			break
		}

		lwc.appendData(wcl.Me())
	}

	return
}

func (lwc ListWhereClause) Me() ListWhereClause {
	return lwc
}

func (lwc *ListWhereClause) appendData(whereClause WhereClause) {
	(*lwc) = append((*lwc), whereClause)
}

type BuilderWhereClause struct {
	ConfigRequest IConfigRequest
}

// WhereClause struct
type WhereClause struct {
	builder            BuilderWhereClause
	defaultConjunction RequestConjunction
	conjunctionWhere   ConjunctionWhere
	conjunction        RequestConjunction
	column             string
	operator           RequestOperator
	value              RequestValue
	valueKind          ValueKind
	err                error
}

func GenWhereClause(builder BuilderWhereClause, conjunctionWhere ConjunctionWhere) (wcl *WhereClause) {
	wcl = new(WhereClause)

	// builder
	wcl.setBuilder(builder)

	// defaultConjunction
	defaultConjunction, errDefConj := wcl.genDefaultConjunction()
	if errDefConj != nil {
		wcl.setErr(errDefConj)
		return
	}
	wcl.setDefaultConjunction(defaultConjunction)

	// conjunctionWhere
	wcl.setConjunctionWhere(conjunctionWhere)

	// conjunction
	conj := wcl.genConjunction()
	wcl.setConjunction(conj)

	// column
	col := wcl.genColumn()
	wcl.setColumn(col)

	// operator
	opr := wcl.genOperator()
	wcl.setOperator(opr)

	// value
	val := wcl.genValue()
	wcl.setValue(val)

	// value kind
	valKind := wcl.genValueKind()
	wcl.setValueKind(valKind)

	return
}

func (wcl WhereClause) Me() WhereClause {
	return wcl
}

func (wcl WhereClause) ConjunctionWhere() (conjunctionWhere ConjunctionWhere) {
	conjunctionWhere = wcl.getConjunctionWhere()
	return
}

func (wcl WhereClause) Conjunction() (conjunction RequestConjunction) {
	conjunction = wcl.getConjunction()
	return
}

func (wcl WhereClause) Column() (column string) {
	column = wcl.getColumn()
	return
}

func (wcl WhereClause) Operator() (operator RequestOperator) {
	operator = wcl.getOperator()
	return
}

func (wcl WhereClause) Value() (value RequestValue) {
	value = wcl.getValue()
	return
}

func (wcl WhereClause) ValueKind() (valueKind ValueKind) {
	valueKind = wcl.getValueKind()
	return
}

func (wcl WhereClause) IsEmptyConjunction() (isEmpty bool) {
	isEmpty = lib.IsEmptyStr(wcl.getConjunction().String())
	return
}

func (wcl WhereClause) IsEmptyOperator() (isEmpty bool) {
	isEmpty = lib.IsEmptyStr(wcl.getOperator().String())
	return
}

func (wra WhereClause) Err() error {
	return wra.getErr()
}

func (wra *WhereClause) setBuilder(builder BuilderWhereClause) {
	wra.builder = builder
}

func (wra WhereClause) getBuilder() BuilderWhereClause {
	return wra.builder
}

func (wcl *WhereClause) setDefaultConjunction(defaultConjunction RequestConjunction) {
	wcl.defaultConjunction = defaultConjunction
}

func (wcl WhereClause) getDefaultConjunction() RequestConjunction {
	return wcl.defaultConjunction
}

func (wcl WhereClause) genDefaultConjunction() (defaultConjunction RequestConjunction, err error) {
	builder := wcl.getBuilder()
	configRequest := builder.ConfigRequest
	configDefConjunction := configRequest.DefaultConjunction()

	defConjunction, errDefConjunction := GenRequestConjunction(configDefConjunction)
	if errDefConjunction != nil {
		err = fmt.Errorf("GenRequestConjunction(): default conjunction %s is unknown", configDefConjunction)
		return
	}

	defaultConjunction = defConjunction
	return
}

func (wcl *WhereClause) setConjunctionWhere(conjunctionWhere ConjunctionWhere) {
	wcl.conjunctionWhere = conjunctionWhere
}

func (wcl WhereClause) getConjunctionWhere() (conjunctionWhere ConjunctionWhere) {
	conjunctionWhere = wcl.conjunctionWhere
	return
}

func (wcl *WhereClause) setConjunction(conjunction RequestConjunction) {
	wcl.conjunction = conjunction
}

func (wcl WhereClause) getConjunction() (conjunction RequestConjunction) {
	conjunction = wcl.conjunction
	return
}

func (wcl WhereClause) genConjunction() (conjunction RequestConjunction) {
	defaultConjunction := wcl.getDefaultConjunction()

	// if there is any error, conjunction value will set as default
	conjunction = defaultConjunction

	conjWhe := wcl.getConjunctionWhere()
	conj := conjWhe.ConjunctionIdentity()

	hasFirstIdx := conj.HasFirstIndex()
	if !hasFirstIdx {
		return
	}

	firstIdx := conj.FirstIndex()
	hasConjunction := firstIdx.Identity() == ConjunctionFirstIndex
	if !hasConjunction {
		return
	}

	rawData := conj.RawWhereClause().RawData()
	firstIdxRaw := rawData[0]

	firstIdxData, ok := firstIdxRaw.(string)
	if !ok {
		return
	}

	conjunction, err := GenRequestConjunction(firstIdxData)
	if err != nil {
		return
	}

	return
}

func (wcl *WhereClause) setColumn(column string) {
	wcl.column = column
}

func (wcl WhereClause) getColumn() (column string) {
	column = wcl.column
	return
}

func (wcl WhereClause) genColumn() (column string) {
	conjWhe := wcl.getConjunctionWhere()
	where := conjWhe.WhereIdentity()

	hasFirstIdx := where.HasFirstIndex()
	if !hasFirstIdx {
		return
	}

	firstIdx := where.FirstIndex()
	if firstIdx.Identity() != ColumnFirstIndex {
		return
	}

	raw := where.RawWhereClause().RawData()
	firstIdxRaw := raw[0]
	firstIdxData, ok := firstIdxRaw.(string)
	if !ok {
		return
	}

	column = firstIdxData
	return
}

func (wcl *WhereClause) setOperator(operator RequestOperator) {
	wcl.operator = operator
}

func (wcl WhereClause) getOperator() (operator RequestOperator) {
	operator = wcl.operator
	return
}

func (wcl WhereClause) genOperator() (operator RequestOperator) {
	operator = IsReqOpr

	conjWhe := wcl.getConjunctionWhere()
	where := conjWhe.WhereIdentity()

	hasSecondIdx := where.HasSecondIndex()
	if !hasSecondIdx {
		return
	}

	secondIdx := where.SecondIndex()
	if secondIdx.Identity() != OperatorSecondIndex {
		return
	}

	raw := where.RawWhereClause().RawData()
	secondIdxRaw := raw[1]
	secondIdxData, ok := secondIdxRaw.(string)
	if !ok {
		return
	}

	opr, err := GenRequestOperator(secondIdxData)
	if err != nil {
		return
	}

	operator = opr
	return
}

func (wcl *WhereClause) setValue(value RequestValue) {
	wcl.value = value
}

func (wcl WhereClause) getValue() (value RequestValue) {
	value = wcl.value
	return
}

/*
genValue

match value in third index or second index
*/
func (wcl WhereClause) genValue() (value RequestValue) {
	valThird, isMatchValThird := wcl.genValueThirdIndex()
	if isMatchValThird {
		value = valThird
		return
	}

	valSecond, isMatchValSecond := wcl.genValueSecondIndex()
	if isMatchValSecond {
		value = valSecond
		return
	}

	return
}

func (wcl WhereClause) genValueSecondIndex() (value RequestValue, isMatch bool) {
	conjWhe := wcl.getConjunctionWhere()
	where := conjWhe.WhereIdentity()

	hasSecondIdx := where.HasSecondIndex()
	if !hasSecondIdx {
		return
	}

	secondIdx := where.SecondIndex()
	if secondIdx.Identity() != ValueSecondIndex {
		return
	}

	isMatch = true
	value = secondIdx.Value()
	return
}

func (wcl WhereClause) genValueThirdIndex() (value RequestValue, isMatch bool) {
	conjWhe := wcl.getConjunctionWhere()
	where := conjWhe.WhereIdentity()

	hasThirdIdx := where.HasThirdIndex()
	if !hasThirdIdx {
		return
	}

	thirdIdx := where.ThirdIndex()
	if thirdIdx.Identity() != ValueThirdIndex {
		return
	}

	isMatch = true
	value = thirdIdx.Value()
	return
}

func (wcl *WhereClause) setValueKind(valueKind ValueKind) {
	wcl.valueKind = valueKind
}

func (wcl WhereClause) getValueKind() (valueKind ValueKind) {
	valueKind = wcl.valueKind
	return
}

/*
genValueKind

match value kind in third index or second index
*/
func (wcl WhereClause) genValueKind() (value ValueKind) {
	valThird, isMatchValThird := wcl.genValueKindThirdIndex()
	if isMatchValThird {
		value = valThird
		return
	}

	valSecond, isMatchValSecond := wcl.genValueKindSecondIndex()
	if isMatchValSecond {
		value = valSecond
		return
	}

	return
}

func (wcl WhereClause) genValueKindSecondIndex() (value ValueKind, isMatch bool) {
	conjWhe := wcl.getConjunctionWhere()
	where := conjWhe.WhereIdentity()

	hasSecondIdx := where.HasSecondIndex()
	if !hasSecondIdx {
		return
	}

	secondIdx := where.SecondIndex()
	if secondIdx.Identity() != ValueSecondIndex {
		return
	}

	isMatch = true
	value = secondIdx.ValueKind()
	return
}

func (wcl WhereClause) genValueKindThirdIndex() (value ValueKind, isMatch bool) {
	conjWhe := wcl.getConjunctionWhere()
	where := conjWhe.WhereIdentity()

	hasThirdIdx := where.HasThirdIndex()
	if !hasThirdIdx {
		return
	}

	thirdIdx := where.ThirdIndex()
	if thirdIdx.Identity() != ValueThirdIndex {
		return
	}

	isMatch = true
	value = thirdIdx.ValueKind()
	return
}

func (wra *WhereClause) setErr(err error) {
	wra.err = err
}

func (wra WhereClause) getErr() error {
	return wra.err
}
