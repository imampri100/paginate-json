package filterrequest

import (
	"fmt"
	"reflect"
)

type FirstIndexClause struct {
	firstRawData interface{}
	identity     FirstIndexIdentity
	conjunction  RequestConjunction
	column       string
	err          error
}

func GenFirstIndexClause(firstRawData interface{}) (fic *FirstIndexClause) {
	fic = new(FirstIndexClause)

	// firstRawData
	fic.setFirstRawData(firstRawData)

	// identity, conjunction, column
	identity, conj, column, errIdentity := fic.genIdentity()
	if errIdentity != nil {
		fic.setErr(errIdentity)
		return
	}
	fic.setIdentity(identity)
	fic.setConjunction(conj)
	fic.setColumn(column)

	return
}

func (fic FirstIndexClause) Me() FirstIndexClause {
	return fic
}

func (fic FirstIndexClause) Identity() (identity FirstIndexIdentity) {
	identity = fic.getIdentity()
	return
}

func (fic FirstIndexClause) Conjunction() (conjunction RequestConjunction) {
	conjunction = fic.getConjunction()
	return
}

func (fic FirstIndexClause) Column() (column string) {
	column = fic.getColumn()
	return
}

func (fic FirstIndexClause) Err() (err error) {
	err = fic.getErr()
	return
}

func (fic *FirstIndexClause) setFirstRawData(firstRawData interface{}) {
	fic.firstRawData = firstRawData
}

func (fic FirstIndexClause) getFirstRawData() (firstRawData interface{}) {
	firstRawData = fic.firstRawData
	return
}

func (fic *FirstIndexClause) setIdentity(identity FirstIndexIdentity) {
	fic.identity = identity
}

func (fic FirstIndexClause) getIdentity() (identity FirstIndexIdentity) {
	identity = fic.identity
	return
}

func (fic FirstIndexClause) genIdentity() (identity FirstIndexIdentity, conjunction RequestConjunction, column string, err error) {
	firstRawData := fic.getFirstRawData()

	inputStr, ok := firstRawData.(string)
	if !ok {
		kind := reflect.TypeOf(firstRawData).Kind()
		err = fmt.Errorf("first index of filter must be string of conjunction or column. Got %+v, kind %s", inputStr, kind)
		return
	}

	conj, errGen := GenRequestConjunction(inputStr)
	if errGen == nil {
		identity = ConjunctionFirstIndex
		conjunction = conj
		return
	}

	identity = ColumnFirstIndex
	column = inputStr
	return
}

func (fic *FirstIndexClause) setConjunction(conjunction RequestConjunction) {
	fic.conjunction = conjunction
}

func (fic FirstIndexClause) getConjunction() (conjunction RequestConjunction) {
	conjunction = fic.conjunction
	return
}

func (fic *FirstIndexClause) setColumn(column string) {
	fic.column = column
}

func (fic FirstIndexClause) getColumn() (column string) {
	column = fic.column
	return
}

func (fic *FirstIndexClause) setErr(err error) {
	fic.err = err
}

func (fic FirstIndexClause) getErr() (err error) {
	err = fic.err
	return
}

type FirstIndexIdentity string

const (
	ConjunctionFirstIndex FirstIndexIdentity = "conjunction"
	ColumnFirstIndex      FirstIndexIdentity = "column"
)

func (fic FirstIndexIdentity) String() string {
	return string(fic)
}
