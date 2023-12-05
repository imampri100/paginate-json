package filterrequest

import (
	"fmt"
	"paginatejson/lib"
)

type ListPositionWhereClause []PositionWhereClause

type PositionWhereClause struct {
	slice  []interface{}
	level  int
	index  int
	length int
}

func NewPositionWhereClause() (pwc *PositionWhereClause) {
	pwc = new(PositionWhereClause)
	return
}

func (pwc PositionWhereClause) Me() PositionWhereClause {
	return pwc
}

func (pwc PositionWhereClause) Print() (result string) {
	bte, err := lib.JSONMarshal(pwc)
	if err != nil {
		result = fmt.Sprintf("PositionWhereClause: %s", err.Error())
		return
	}

	result = string(bte)
	return
}

func (pwc PositionWhereClause) Slice() (slice []interface{}) {
	slice = pwc.getSlice()
	return
}

func (pwc PositionWhereClause) Level() (level int) {
	level = pwc.getLevel()
	return
}

func (pwc PositionWhereClause) Index() (index int) {
	index = pwc.getIndex()
	return
}

func (pwc PositionWhereClause) Length() (length int) {
	length = pwc.getLength()
	return
}

func (pwc *PositionWhereClause) ReplaceSlice(slice []interface{}) {
	pwc.setSlice(slice)
}

func (pwc *PositionWhereClause) IncrementLevel() {
	newLevel := pwc.genIncrementLevel()
	pwc.setLevel(newLevel)
}

func (pwc *PositionWhereClause) SetLevel(level int) {
	pwc.setLevel(level)
}

func (pwc *PositionWhereClause) ResetIndex() {
	pwc.setIndex(0)
}

func (pwc *PositionWhereClause) SetIndex(index int) {
	pwc.setIndex(index)
}

func (pwc *PositionWhereClause) IncrementIndex() {
	newIndex := pwc.genIncrementIndex()
	pwc.setIndex(newIndex)
}

func (pwc *PositionWhereClause) SetLength(length int) {
	pwc.setLength(length)
}

func (pwc *PositionWhereClause) setSlice(slice []interface{}) {
	pwc.slice = slice
}

func (pwc PositionWhereClause) getSlice() (slice []interface{}) {
	slice = pwc.slice
	return
}

func (pwc *PositionWhereClause) setLevel(level int) {
	pwc.level = level
}

func (pwc PositionWhereClause) getLevel() (level int) {
	level = pwc.level
	return
}

func (pwc PositionWhereClause) genIncrementLevel() (level int) {
	oldLevel := pwc.getLevel()
	newLevel := oldLevel + 1
	level = newLevel
	return
}

func (pwc *PositionWhereClause) setIndex(index int) {
	pwc.index = index
}

func (pwc PositionWhereClause) getIndex() (index int) {
	index = pwc.index
	return
}

func (pwc PositionWhereClause) genIncrementIndex() (index int) {
	oldIndex := pwc.getIndex()
	newIndex := oldIndex + 1
	index = newIndex
	return
}

func (pwc *PositionWhereClause) setLength(length int) {
	pwc.length = length
}

func (pwc PositionWhereClause) getLength() (length int) {
	length = pwc.length
	return
}
