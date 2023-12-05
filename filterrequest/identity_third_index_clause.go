package filterrequest

type ThirdIndexClause struct {
	thirdRawData interface{}
	identity     ThirdIndexIdentity
	value        RequestValue
	valueKind    ValueKind
	err          error
}

func GenThirdIndexClause(thirdRawData interface{}) (sic *ThirdIndexClause) {
	sic = new(ThirdIndexClause)

	// thirdRawData
	sic.setThirdRawData(thirdRawData)

	// identity
	identity := sic.genIdentity()
	sic.setIdentity(identity)

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

func (sic ThirdIndexClause) Me() ThirdIndexClause {
	return sic
}

func (sic ThirdIndexClause) Identity() (identity ThirdIndexIdentity) {
	identity = sic.getIdentity()
	return
}

func (sic ThirdIndexClause) ValueKind() (valueKind ValueKind) {
	valueKind = sic.getValueKind()
	return
}

func (sic ThirdIndexClause) Value() (value RequestValue) {
	value = sic.getValue()
	return
}

func (sic ThirdIndexClause) Err() (err error) {
	err = sic.getErr()
	return
}

func (sic *ThirdIndexClause) setThirdRawData(thirdRawData interface{}) {
	sic.thirdRawData = thirdRawData
}

func (sic ThirdIndexClause) getThirdRawData() (thirdRawData interface{}) {
	thirdRawData = sic.thirdRawData
	return
}

func (sic *ThirdIndexClause) setIdentity(identity ThirdIndexIdentity) {
	sic.identity = identity
}

func (sic ThirdIndexClause) getIdentity() (identity ThirdIndexIdentity) {
	identity = sic.identity
	return
}

func (sic ThirdIndexClause) genIdentity() (identity ThirdIndexIdentity) {
	identity = ValueThirdIndex
	return
}

func (sic *ThirdIndexClause) setValueKind(valueKind ValueKind) {
	sic.valueKind = valueKind
}

func (sic ThirdIndexClause) getValueKind() (valueKind ValueKind) {
	valueKind = sic.valueKind
	return
}

func (sic ThirdIndexClause) canGenValueKind() (canGen bool) {
	identity := sic.getIdentity()
	canGen = identity == ValueThirdIndex
	return
}

func (sic ThirdIndexClause) genValueKind() (valueKind ValueKind, err error) {
	canGen := sic.canGenValueKind()
	if !canGen {
		return
	}

	thirdRawData := sic.getThirdRawData()

	valKind, errValKind := GenValueKind(thirdRawData)
	if errValKind != nil {
		err = errValKind
		return
	}

	valueKind = valKind
	return
}

func (sic *ThirdIndexClause) setValue(value RequestValue) {
	sic.value = value
}

func (sic ThirdIndexClause) getValue() (value RequestValue) {
	value = sic.value
	return
}

func (sic ThirdIndexClause) canGenValue() (canGen bool) {
	canGen = sic.canGenValueKind()
	return
}

func (sic ThirdIndexClause) genValue() (value RequestValue, err error) {
	canGen := sic.canGenValue()
	if !canGen {
		return
	}

	valKind := sic.getValueKind()
	thirdRawData := sic.getThirdRawData()

	reqVal, errReqVal := GenRequestValue(thirdRawData, valKind)
	if errReqVal != nil {
		err = errReqVal
		return
	}

	value = reqVal
	return
}

func (sic *ThirdIndexClause) setErr(err error) {
	sic.err = err
}

func (sic ThirdIndexClause) getErr() (err error) {
	err = sic.err
	return
}

type ThirdIndexIdentity string

const (
	ValueThirdIndex ThirdIndexIdentity = "value"
)

func (sic ThirdIndexIdentity) String() string {
	return string(sic)
}
