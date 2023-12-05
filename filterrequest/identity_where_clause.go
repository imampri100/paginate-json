package filterrequest

import (
	"errors"
	"fmt"
	"strings"
)

type ListIdentityWhereClause []IdentityWhereClause

func GenListIdentityWhereClause(input ListRawWhereClause) (liw *ListIdentityWhereClause, err error) {
	liw = new(ListIdentityWhereClause)

	for idxRaw := range input {
		itemRaw := input[idxRaw]

		iwc := GenIdentityWhereClause(itemRaw)
		if iwc.Error() != nil {
			err = iwc.Error()
			break
		}

		liw.appendData(iwc.Me())
	}

	return
}

func (liw ListIdentityWhereClause) Me() ListIdentityWhereClause {
	return liw
}

func (liw *ListIdentityWhereClause) appendData(identityWhereClause IdentityWhereClause) {
	(*liw) = append((*liw), identityWhereClause)
}

type IdentityWhereClause struct {
	rawWhereClause RawWhereClause
	length         int
	maxIndex       int
	firstIndex     FirstIndexClause
	secondIndex    SecondIndexClause
	thirdIndex     ThirdIndexClause
	err            error
}

func GenIdentityWhereClause(rawWhereClause RawWhereClause) (iwc *IdentityWhereClause) {
	iwc = new(IdentityWhereClause)

	// RawWhereClause
	iwc.setRawWhereClause(rawWhereClause)

	// length
	length := iwc.genLength()
	iwc.setLength(length)

	// maxIndex
	maxIndex, errMaxIndex := iwc.genMaxIndex()
	if errMaxIndex != nil {
		iwc.setErr(errMaxIndex)
		return
	}
	iwc.setMaxIndex(maxIndex)

	// firstIndex
	firstIndex, errFirstIndex := iwc.genFirstIndex()
	if errFirstIndex != nil {
		iwc.setErr(errFirstIndex)
		return
	}
	iwc.setFirstIndex(firstIndex)

	// secondIndex
	secondIndex, errSecondIndex := iwc.genSecondIndex()
	if errSecondIndex != nil {
		iwc.setErr(errSecondIndex)
		return
	}
	iwc.setSecondIndex(secondIndex)

	// thirdIndex
	thirdIndex, errThirdIndex := iwc.genThirdIndex()
	if errThirdIndex != nil {
		iwc.setErr(errThirdIndex)
		return
	}
	iwc.setThirdIndex(thirdIndex)

	// validate pattern
	errValidate := iwc.validatePattern()
	if errValidate != nil {
		iwc.setErr(errValidate)
		return
	}

	return
}

func (iwc IdentityWhereClause) Me() IdentityWhereClause {
	return iwc
}

func (iwc IdentityWhereClause) RawWhereClause() (rawWhereClause RawWhereClause) {
	rawWhereClause = iwc.getRawWhereClause()
	return
}

func (iwc IdentityWhereClause) Length() (length int) {
	length = iwc.getLength()
	return
}

func (iwc IdentityWhereClause) MaxIndex() (maxIndex int) {
	maxIndex = iwc.getMaxIndex()
	return
}

func (iwc IdentityWhereClause) FirstIndex() (firstIndex FirstIndexClause) {
	firstIndex = iwc.getFirstIndex()
	return
}

func (iwc IdentityWhereClause) HasFirstIndex() (hasFirstIndex bool) {
	hasFirstIndex = iwc.canGenFirstIndex()
	return
}

func (iwc IdentityWhereClause) SecondIndex() (secondIndex SecondIndexClause) {
	secondIndex = iwc.secondIndex
	return
}

func (iwc IdentityWhereClause) HasSecondIndex() (hasSecondIndex bool) {
	hasSecondIndex = iwc.canGenSecondIndex()
	return
}

func (iwc IdentityWhereClause) ThirdIndex() (thirdIndex ThirdIndexClause) {
	thirdIndex = iwc.thirdIndex
	return
}

func (iwc IdentityWhereClause) HasThirdIndex() (hasThirdIndex bool) {
	hasThirdIndex = iwc.canGenThirdIndex()
	return
}

func (iwc IdentityWhereClause) Error() (err error) {
	err = iwc.getErr()
	return
}

func (iwc *IdentityWhereClause) setRawWhereClause(rawWhereClause RawWhereClause) {
	iwc.rawWhereClause = rawWhereClause
}

func (iwc IdentityWhereClause) getRawWhereClause() (rawWhereClause RawWhereClause) {
	rawWhereClause = iwc.rawWhereClause
	return
}

func (iwc *IdentityWhereClause) setLength(length int) {
	iwc.length = length
}

func (iwc IdentityWhereClause) getLength() (length int) {
	length = iwc.length
	return
}

func (iwc IdentityWhereClause) genLength() (length int) {
	raw := iwc.getRawWhereClause()
	length = len(raw.RawData())
	return
}

func (iwc *IdentityWhereClause) setMaxIndex(maxIndex int) {
	iwc.maxIndex = maxIndex
}

func (iwc IdentityWhereClause) getMaxIndex() (maxIndex int) {
	maxIndex = iwc.maxIndex
	return
}

func (iwc IdentityWhereClause) genMaxIndex() (maxIndex int, err error) {
	length := iwc.getLength()
	if length == 0 {
		err = errors.New("cannot gen max index. Data length is 0")
		return
	}

	maxIndex = length - 1
	return
}

func (iwc *IdentityWhereClause) setFirstIndex(firstIndex FirstIndexClause) {
	iwc.firstIndex = firstIndex
}

func (iwc IdentityWhereClause) getFirstIndex() (firstIndex FirstIndexClause) {
	firstIndex = iwc.firstIndex
	return
}

func (iwc IdentityWhereClause) canGenFirstIndex() (canGen bool) {
	maxIndex := iwc.getMaxIndex()
	canGen = maxIndex >= 0
	return
}

func (iwc IdentityWhereClause) genFirstIndex() (firstIndex FirstIndexClause, err error) {
	canGen := iwc.canGenFirstIndex()
	if !canGen {
		return
	}

	raw := iwc.getRawWhereClause().RawData()
	firstIndexRaw := raw[0]

	fic := GenFirstIndexClause(firstIndexRaw)
	if fic.Err() != nil {
		err = fic.Err()
		return
	}

	firstIndex = fic.Me()
	return
}

func (iwc *IdentityWhereClause) setSecondIndex(secondIndex SecondIndexClause) {
	iwc.secondIndex = secondIndex
}

func (iwc IdentityWhereClause) getSecondIndex() (secondIndex SecondIndexClause) {
	secondIndex = iwc.secondIndex
	return
}

func (iwc IdentityWhereClause) canGenSecondIndex() (canGen bool) {
	maxIndex := iwc.getMaxIndex()
	canGen = maxIndex >= 1
	return
}

func (iwc IdentityWhereClause) genSecondIndex() (secondIndex SecondIndexClause, err error) {
	canGen := iwc.canGenSecondIndex()
	if !canGen {
		return
	}

	rawData := iwc.getRawWhereClause().RawData()
	secondIndexRaw := rawData[1]

	sic := GenSecondIndexClause(secondIndexRaw)
	if sic.Err() != nil {
		err = sic.Err()
		return
	}

	secondIndex = sic.Me()
	return
}

func (iwc *IdentityWhereClause) setThirdIndex(thirdIndex ThirdIndexClause) {
	iwc.thirdIndex = thirdIndex
}

func (iwc IdentityWhereClause) canGenThirdIndex() (canGen bool) {
	maxIndex := iwc.getMaxIndex()
	canGen = maxIndex >= 2
	return
}

func (iwc IdentityWhereClause) getThirdIndex() (thirdIndex ThirdIndexClause) {
	thirdIndex = iwc.thirdIndex
	return
}

func (iwc IdentityWhereClause) genThirdIndex() (thirdIndex ThirdIndexClause, err error) {
	canGen := iwc.canGenThirdIndex()
	if !canGen {
		return
	}

	rawData := iwc.getRawWhereClause().RawData()
	thirdIndexRaw := rawData[2]

	tic := GenThirdIndexClause(thirdIndexRaw)
	if tic.Err() != nil {
		err = tic.Err()
		return
	}

	thirdIndex = tic.Me()
	return
}

/*
validatePattern

Available pattern:
  - first idx: conjunction. Reflect: [conjunction]
  - first idx: column, second idx: value. Reflect: [column, value]. Value must be single value.
  - first idx: column, second idx: operator, third idx: value. Reflect: [column, operator, value]. Value can be single value or multiple value.
*/
func (iwc IdentityWhereClause) validatePattern() (err error) {
	errValidateMaxIdx := iwc.validateMaxIdx()
	if errValidateMaxIdx != nil {
		err = errValidateMaxIdx
		return
	}

	errValidateFirstIdx := iwc.validatePatternFirstIdx()
	if errValidateFirstIdx != nil {
		err = errValidateFirstIdx
		return
	}

	errValidateSecondIdx := iwc.validatePatternSecondIdx()
	if errValidateSecondIdx != nil {
		err = errValidateSecondIdx
		return
	}

	errValidateThirdIdx := iwc.validatePatternThirdIdx()
	if errValidateThirdIdx != nil {
		err = errValidateThirdIdx
		return
	}

	return
}

func (iwc IdentityWhereClause) validateMaxIdx() (err error) {
	canGenFirstIdx := iwc.canGenFirstIndex()
	if canGenFirstIdx {
		return
	}

	err = errors.New("validate pattern: max index is below 0. Make sure the filter contains minimum 1 item")
	return
}

/*
validatePatternFirstIdx

if first index == conjunction, then the second index must not be available
*/
func (iwc IdentityWhereClause) validatePatternFirstIdx() (err error) {
	firstIdx := iwc.getFirstIndex()
	if firstIdx.Identity() != ConjunctionFirstIndex {
		return
	}

	canGenSecondIdx := iwc.canGenSecondIndex()
	if !canGenSecondIdx {
		return
	}

	err = fmt.Errorf("validate pattern first idx: conjunction of filter ('AND' | 'OR'), can only includes 1 item")
	return
}

func (iwc IdentityWhereClause) validatePatternSecondIdx() (err error) {
	errValidateOpr := iwc.validateOperatorSecondIdx()
	if errValidateOpr != nil {
		err = errValidateOpr
		return
	}

	errValidateVal := iwc.validateValueSecondIdx()
	if errValidateVal != nil {
		err = errValidateVal
		return
	}

	return
}

/*
validateOperatorSecondIdx

if second index == operator, then the third index must be available
*/
func (iwc IdentityWhereClause) validateOperatorSecondIdx() (err error) {
	secondIdx := iwc.getSecondIndex()
	if secondIdx.Identity() != OperatorSecondIndex {
		return
	}

	canGenThirdIdx := iwc.canGenThirdIndex()
	if canGenThirdIdx {
		return
	}

	err = fmt.Errorf("validate pattern second idx: you must define value in third index when using operator in an item")
	return
}

/*
validateValueSecondIdx

if second index == value, then value must not be multiple values / slice
*/
func (iwc IdentityWhereClause) validateValueSecondIdx() (err error) {
	secondIdx := iwc.getSecondIndex()
	if secondIdx.Identity() != ValueSecondIndex {
		return
	}

	valKind := secondIdx.ValueKind()
	if !valKind.IsMultiple() {
		return
	}

	listSingleValKind := []string{
		StringValKind.String(),
		// IntValKind.String(),
		Float64ValKind.String(),
	}
	listSingleStr := strings.Join(listSingleValKind, ", ")
	err = fmt.Errorf("validate value second idx: value in second index of item must be single, such as: %s", listSingleStr)
	return
}

func (iwc IdentityWhereClause) validatePatternThirdIdx() (err error) {
	secondIdx := iwc.getSecondIndex()

	// check operator
	opr := secondIdx.Operator()
	isOprSlice := opr.MustSliceValue()

	// check value kind
	thirdIdx := iwc.getThirdIndex()
	valKind := thirdIdx.ValueKind()
	isValKindMultiple := valKind.IsMultiple()

	// compare operator with value kind
	if isOprSlice == isValKindMultiple {
		return
	}

	err = fmt.Errorf("validate pattern third idx: cannot use operator %s for value %+v", opr.String(), thirdIdx)
	return
}

func (iwc *IdentityWhereClause) setErr(err error) {
	iwc.err = err
}

func (iwc IdentityWhereClause) getErr() (err error) {
	err = iwc.err
	return
}
