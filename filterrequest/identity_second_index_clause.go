package filterrequest

type SecondIndexClause struct {
	secondRawData interface{}
	identity      SecondIndexIdentity
	operator      RequestOperator
	valueKind     ValueKind
	value         RequestValue
	err           error
}

func GenSecondIndexClause(secondRawData interface{}) (sic *SecondIndexClause) {
	sic = new(SecondIndexClause)

	// secondRawData
	sic.setSecondRawData(secondRawData)

	// identity, operator
	identity, operator := sic.genIdentity()
	sic.setIdentity(identity)
	sic.setOperator(operator)

	// valueKind
	valueKind, errValKind := sic.genValueKind()
	if errValKind != nil {
		sic.setErr(errValKind)
		return
	}
	sic.setValueKind(valueKind)

	// value
	val, errVal := sic.genValue()
	if errVal != nil {
		sic.setErr(errVal)
		return
	}
	sic.setValue(val)

	return
}

func (sic SecondIndexClause) Me() SecondIndexClause {
	return sic
}

func (sic SecondIndexClause) Identity() (identity SecondIndexIdentity) {
	identity = sic.getIdentity()
	return
}

func (sic SecondIndexClause) Operator() (operator RequestOperator) {
	operator = sic.getOperator()
	return
}

func (sic SecondIndexClause) ValueKind() (valueKind ValueKind) {
	valueKind = sic.getValueKind()
	return
}

func (sic SecondIndexClause) Value() (value RequestValue) {
	value = sic.getValue()
	return
}

func (sic SecondIndexClause) Err() (err error) {
	err = sic.getErr()
	return
}

func (sic *SecondIndexClause) setSecondRawData(secondRawData interface{}) {
	sic.secondRawData = secondRawData
}

func (sic SecondIndexClause) getSecondRawData() (secondRawData interface{}) {
	secondRawData = sic.secondRawData
	return
}

func (sic *SecondIndexClause) setIdentity(identity SecondIndexIdentity) {
	sic.identity = identity
}

func (sic SecondIndexClause) getIdentity() (identity SecondIndexIdentity) {
	identity = sic.identity
	return
}

func (sic SecondIndexClause) genIdentity() (identity SecondIndexIdentity, operator RequestOperator) {
	secondRawData := sic.getSecondRawData()

	inputStr, ok := secondRawData.(string)
	if !ok {
		identity = ValueSecondIndex
		return
	}

	opr, errGen := GenRequestOperator(inputStr)
	if errGen != nil {
		identity = ValueSecondIndex
		return
	}

	identity = OperatorSecondIndex
	operator = opr
	return
}

func (sic *SecondIndexClause) setOperator(operator RequestOperator) {
	sic.operator = operator
}

func (sic SecondIndexClause) getOperator() (operator RequestOperator) {
	operator = sic.operator
	return
}

func (sic *SecondIndexClause) setValueKind(valueKind ValueKind) {
	sic.valueKind = valueKind
}

func (sic SecondIndexClause) getValueKind() (valueKind ValueKind) {
	valueKind = sic.valueKind
	return
}

func (sic SecondIndexClause) canGenValueKind() (canGen bool) {
	identity := sic.getIdentity()
	canGen = identity == ValueSecondIndex
	return
}

func (sic SecondIndexClause) genValueKind() (valueKind ValueKind, err error) {
	canGen := sic.canGenValueKind()
	if !canGen {
		return
	}

	secondRawData := sic.getSecondRawData()

	valKind, errValKind := GenValueKind(secondRawData)
	if errValKind != nil {
		err = errValKind
		return
	}

	valueKind = valKind
	return
}

func (sic *SecondIndexClause) setValue(value RequestValue) {
	sic.value = value
}

func (sic SecondIndexClause) getValue() (value RequestValue) {
	value = sic.value
	return
}

func (sic SecondIndexClause) canGenValue() (canGen bool) {
	canGen = sic.canGenValueKind()
	return
}

func (sic SecondIndexClause) genValue() (value RequestValue, err error) {
	canGen := sic.canGenValue()
	if !canGen {
		return
	}

	valKind := sic.getValueKind()
	secondRawData := sic.getSecondRawData()

	reqVal, errReqVal := GenRequestValue(secondRawData, valKind)
	if errReqVal != nil {
		err = errReqVal
		return
	}

	value = reqVal
	return
}

func (sic *SecondIndexClause) setErr(err error) {
	sic.err = err
}

func (sic SecondIndexClause) getErr() (err error) {
	err = sic.err
	return
}

type SecondIndexIdentity string

const (
	OperatorSecondIndex SecondIndexIdentity = "operator"
	ValueSecondIndex    SecondIndexIdentity = "value"
)

func (sic SecondIndexIdentity) String() string {
	return string(sic)
}
