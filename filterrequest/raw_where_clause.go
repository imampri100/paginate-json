package filterrequest

import "fmt"

type ListRawWhereClause []RawWhereClause

/*
input possibilities
- slice
*/
func GenListRawWhereClause(listIface []interface{}, wrapperPosition PositionWhereClause) (lwc *ListRawWhereClause, err error) {
	lwc = new(ListRawWhereClause)

	// assign value
	wrapperPosition.ReplaceSlice(listIface)
	wrapperPosition.ResetIndex()
	wrapperPosition.SetLength(len(listIface))

	for idxIface := range listIface {
		itemIface := listIface[idxIface]

		if !IsSlice(itemIface) {
			newRaw := GenRawWhereClause(listIface, wrapperPosition)
			lwc.appendData(newRaw.Me())
			break
		}

		listIface, ok := itemIface.([]interface{})
		if !ok {
			err = fmt.Errorf("filter is not list interface in position: %s", wrapperPosition.Print())
			return
		}

		// increment level
		wrapperPosition.IncrementLevel()

		newLwc, errGen := GenListRawWhereClause(listIface, wrapperPosition)
		if errGen != nil {
			err = errGen
			break
		}

		lwc.appendData(newLwc.Me()...)

		// increment wrapper position index
		wrapperPosition.IncrementIndex()
	}

	return
}

func (lwc ListRawWhereClause) Me() ListRawWhereClause {
	return lwc
}

func (lwc *ListRawWhereClause) appendData(listRaw ...RawWhereClause) {
	(*lwc) = append((*lwc), listRaw...)
}

type RawWhereClause struct {
	rawData  []interface{}
	position PositionWhereClause
}

/*
input possibilities
- slice
- conjunction (string)
- column (string)

condition possibilties
- first idx: slice, second idx: slice, n... idx: slice
- first idx: conjunction
- first idx: column, second idx: operator, third idx: value
- first idx: column, second idx: value
*/
func GenRawWhereClause(rawData []interface{}, position PositionWhereClause) (rwc *RawWhereClause) {
	rwc = new(RawWhereClause)
	rwc.setRawData(rawData)
	rwc.setPosition(position)
	return
}

func (rwc RawWhereClause) Me() RawWhereClause {
	return rwc
}

func (rwc RawWhereClause) RawData() (rawData []interface{}) {
	rawData = rwc.getRawData()
	return
}

func (rwc RawWhereClause) Position() (position PositionWhereClause) {
	position = rwc.getPosition()
	return
}

func (rwc *RawWhereClause) setRawData(rawData []interface{}) {
	rwc.rawData = rawData
}

func (rwc RawWhereClause) getRawData() (rawData []interface{}) {
	rawData = rwc.rawData
	return
}

func (rwc *RawWhereClause) setPosition(position PositionWhereClause) {
	rwc.position = position
}

func (rwc RawWhereClause) getPosition() (position PositionWhereClause) {
	position = rwc.position
	return
}
