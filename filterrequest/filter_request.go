package filterrequest

import (
	"errors"
	"fmt"
	"paginatejson/lib"
)

type BuilderWhereClauseCollection struct {
	ConfigRequest IConfigRequest
}

type WhereClauseCollection struct {
	builder              BuilderWhereClauseCollection
	filterRequestString  string
	filterRequestList    []interface{}           // reflection: [["key", "operator", "value"],["OR"], [["key", "operator", "value"]]]
	listFilterOneLevel   ListRawWhereClause      // reflection: [RawWhereClause, RawWhereClause, RawWhereClause]
	listFilterIdentity   ListIdentityWhereClause // reflection: [IdentityWhereClause, IdentifyWhereClause, IdentityWhereClause]
	listConjunctionWhere ListConjunctionWhere
	listWhereClause      ListWhereClause
	err                  error
}

func NewWhereClauseCollection(builder BuilderWhereClauseCollection, filterReq string) (wcc *WhereClauseCollection) {
	wcc = new(WhereClauseCollection)

	// builder
	wcc.setBuilder(builder)

	// filterRequestString
	wcc.setFilterRequestString(filterReq)

	// filterRequestList
	filterReqList, isEmpty, errfilterReqList := wcc.genFilterRequestList()
	if errfilterReqList != nil {
		wcc.setErr(errfilterReqList)
		return
	} else if isEmpty {
		return
	}
	wcc.setFilterRequestList(filterReqList)

	// listFilterOneLevel
	listFilterOneLevel, errListOne := wcc.genListFilterOneLevel()
	if errListOne != nil {
		wcc.setErr(errListOne)
		return
	}
	wcc.setListFilterOneLevel(listFilterOneLevel)

	// listFilterIdentity
	listFilterIdentity, errListIdentity := wcc.genListFilterIdentity()
	if errListIdentity != nil {
		wcc.setErr(errListIdentity)
		return
	}
	wcc.setListFilterIdentity(listFilterIdentity)

	// listConjunctionWhere
	listConjunctionWhere := wcc.genListConjunctionWhere()
	wcc.setListConjunctionWhere(listConjunctionWhere)

	// listWhereClause
	listWhereClause, errListWhereClause := wcc.genListWhereClause()
	if errListWhereClause != nil {
		wcc.setErr(errListWhereClause)
		return
	}
	wcc.setListWhereClause(listWhereClause)

	return
}

func (wcc WhereClauseCollection) ListWhereClause() (listWhereClause ListWhereClause) {
	listWhereClause = wcc.getListWhereClause()
	return
}

func (wcc WhereClauseCollection) Err() (err error) {
	err = wcc.getErr()
	return
}

func (wal *WhereClauseCollection) setBuilder(builder BuilderWhereClauseCollection) {
	wal.builder = builder
}

func (wal WhereClauseCollection) getBuilder() BuilderWhereClauseCollection {
	return wal.builder
}

func (wcc *WhereClauseCollection) setFilterRequestString(filterRequestString string) {
	wcc.filterRequestString = filterRequestString
}

func (wcc WhereClauseCollection) getFilterRequestString() (filterRequestString string) {
	filterRequestString = wcc.filterRequestString
	return
}

func (wcc *WhereClauseCollection) setFilterRequestList(filterRequestList []interface{}) {
	wcc.filterRequestList = filterRequestList
}

func (wcc WhereClauseCollection) getFilterRequestList() (filterRequestList []interface{}) {
	filterRequestList = wcc.filterRequestList
	return
}

func (wcc WhereClauseCollection) genFilterRequestList() (filterRequestList []interface{}, isEmpty bool, err error) {
	filterReq := wcc.getFilterRequestString()

	if lib.IsEmptyStr(filterReq) {
		isEmpty = true
		return
	}

	iface := []interface{}{}
	errUnmarshal := lib.JSONUnmarshal([]byte(filterReq), &iface)
	if errUnmarshal != nil {
		err = errUnmarshal
		return
	}

	if len(iface) == 0 {
		err = errors.New("filter request must contains minimum 1 array of where clause")
		return
	}

	filterRequestList = iface
	return
}

func (wcc *WhereClauseCollection) setListFilterOneLevel(listFilterOneLevel ListRawWhereClause) {
	wcc.listFilterOneLevel = listFilterOneLevel
}

func (wcc WhereClauseCollection) getListFilterOneLevel() (listFilterOneLevel ListRawWhereClause) {
	listFilterOneLevel = wcc.listFilterOneLevel
	return
}

func (wcc WhereClauseCollection) genListFilterOneLevel() (listFilterOneLevel ListRawWhereClause, err error) {
	listIface := wcc.getFilterRequestList()

	wrapper := NewPositionWhereClause()

	newList, errList := GenListRawWhereClause(listIface, wrapper.Me())
	if errList != nil {
		err = errList
		return
	}

	listFilterOneLevel = newList.Me()
	return
}

func (wcc *WhereClauseCollection) setListFilterIdentity(listFilterIdentity ListIdentityWhereClause) {
	wcc.listFilterIdentity = listFilterIdentity
}

func (wcc WhereClauseCollection) getListFilterIdentity() (listFilterIdentity ListIdentityWhereClause) {
	listFilterIdentity = wcc.listFilterIdentity
	return
}

func (wcc WhereClauseCollection) genListFilterIdentity() (listFilterIdentity ListIdentityWhereClause, err error) {
	listRaw := wcc.getListFilterOneLevel()

	newList, errList := GenListIdentityWhereClause(listRaw)
	if errList != nil {
		err = errList
		return
	}

	listFilterIdentity = newList.Me()
	return
}

func (wcc *WhereClauseCollection) setListConjunctionWhere(listConjunctionWhere ListConjunctionWhere) {
	wcc.listConjunctionWhere = listConjunctionWhere
}

func (wcc WhereClauseCollection) getListConjunctionWhere() (listConjunctionWhere ListConjunctionWhere) {
	listConjunctionWhere = wcc.listConjunctionWhere
	return
}

func (wcc WhereClauseCollection) genListConjunctionWhere() (listConjunctionWhere ListConjunctionWhere) {
	listIdentity := wcc.getListFilterIdentity()
	listConjunctionWhere = GenListConjunctionWhere(listIdentity).Me()
	return
}

func (wcc *WhereClauseCollection) setListWhereClause(listWhereClause ListWhereClause) {
	wcc.listWhereClause = listWhereClause
}

func (wcc WhereClauseCollection) getListWhereClause() (listWhereClause ListWhereClause) {
	listWhereClause = wcc.listWhereClause
	return
}

func (wcc WhereClauseCollection) genListWhereClause() (listWhereClause ListWhereClause, err error) {
	builder := wcc.getBuilder()
	configRequest := builder.ConfigRequest

	listConjWhe := wcc.getListConjunctionWhere()

	builderLwc := BuilderListWhereClause{
		ConfigRequest: configRequest,
		Input:         listConjWhe,
	}

	lwc, errLwc := GenListWhereClause(builderLwc)
	if errLwc != nil {
		err = fmt.Errorf("GenListWhereClause().errLwc: %s", errLwc)
		return
	}

	listWhereClause = lwc.Me()
	return
}

func (wcc *WhereClauseCollection) setErr(err error) {
	wcc.err = err
}

func (wcc WhereClauseCollection) getErr() (err error) {
	err = wcc.err
	return
}
