package paginatejson

import (
	"errors"
	"paginatejson/filtergojsonq"
	"paginatejson/filterrequest"
)

func Filtering(listIface []interface{}, filter string, configRequest IConfigRequest) (result []interface{}, err error) {
	builderWhereClauseColection := filterrequest.BuilderWhereClauseCollection{
		ConfigRequest: configRequest,
	}

	wcc := filterrequest.NewWhereClauseCollection(builderWhereClauseColection, filter)
	if wcc.Err() != nil {
		err = wcc.Err()
		return
	}

	listWhereClause := wcc.ListWhereClause()

	gojsonqFilterReq := convertRequestToGojsonqListFilter(listWhereClause)

	builderGojsonq := filtergojsonq.BuilderFilterGojsonq{
		FromInterface:     listIface,
		ListFilterRequest: gojsonqFilterReq,
		JsonSerializer:    configRequest,
	}

	fgj := filtergojsonq.NewFilterGojsonq(builderGojsonq)
	if fgj.Err() != nil {
		err = fgj.Err()
		return
	}

	tempRes := fgj.Result()

	resultListIface, ok := tempRes.([]interface{})
	if !ok {
		err = errors.New("Filtering: gojsonq filter didn't return same result as request")
		return
	}

	result = resultListIface
	return
}

func convertRequestToGojsonqListFilter(listWhere filterrequest.ListWhereClause) (listResult filtergojsonq.ListFilterRequest) {
	for idxWhere := range listWhere {
		itemWhere := listWhere[idxWhere]

		newResult := filtergojsonq.FilterRequest{}
		newResult.Conjunction = convertRequestToGojsonqConjunction(itemWhere.Conjunction())
		newResult.Column = itemWhere.Column()
		newResult.Operator = convertRequestToGojsonqOperator(itemWhere.Operator())
		newResult.Value = filtergojsonq.GojsonqValue(itemWhere.Value())
		newResult.ValueKind = convertRequestToGojsonqValueKind(itemWhere.ValueKind())

		listResult = append(listResult, newResult)
	}

	return
}

func convertRequestToGojsonqConjunction(roc filterrequest.RequestConjunction) (result filtergojsonq.GojsonqConjunction) {
	switch roc {
	case filterrequest.AndReqConj:
		{
			result = filtergojsonq.AndGojsConj
			break
		}
	case filterrequest.OrReqConj:
		{
			result = filtergojsonq.OrGojsConj
			break
		}
	}
	return
}

func convertRequestToGojsonqOperator(rop filterrequest.RequestOperator) (result filtergojsonq.GojsonqOperator) {
	switch rop {
	case filterrequest.IsReqOpr:
		{
			result = filtergojsonq.IsGojsOpr
			break
		}
	case filterrequest.IsNotReqOpr:
		{
			result = filtergojsonq.IsNotGojsOpr
			break
		}
	case filterrequest.GtReqOpr:
		{
			result = filtergojsonq.GtGojsOpr
			break
		}
	case filterrequest.GteReqOpr:
		{
			result = filtergojsonq.GteGojsOpr
			break
		}
	case filterrequest.LtReqOpr:
		{
			result = filtergojsonq.LtGojsOpr
			break
		}
	case filterrequest.LteReqOpr:
		{
			result = filtergojsonq.LteGojsOpr
			break
		}
	case filterrequest.InReqOpr:
		{
			result = filtergojsonq.InGojsOpr
			break
		}
	case filterrequest.NotInReqOpr:
		{
			result = filtergojsonq.NotInGojsOpr
			break
		}
	case filterrequest.LikeReqOpr:
		{
			result = filtergojsonq.StrictContainsGojsOpr
			break
		}
	case filterrequest.IlikeReqOpr:
		{
			result = filtergojsonq.ContainsGojsOpr
			break
		}
	}
	return
}

func convertRequestToGojsonqValueKind(rov filterrequest.ValueKind) (result filtergojsonq.ValueKind) {
	switch rov {
	case filterrequest.StringValKind:
		{
			result = filtergojsonq.StringValKind
			break
		}
	// case filterrequest.IntValKind:
	// 	{
	// 		result = filtergojsonq.IntValKind
	// 		break
	// 	}
	case filterrequest.Float64ValKind:
		{
			result = filtergojsonq.Float64ValKind
			break
		}
	case filterrequest.SliceStringValKind:
		{
			result = filtergojsonq.SliceStringValKind
			break
		}
	// case filterrequest.SliceIntValKind:
	// 	{
	// 		result = filtergojsonq.SliceIntValKind
	// 		break
	// 	}
	case filterrequest.SliceFloat64ValKind:
		{
			result = filtergojsonq.SliceFloat64ValKind
			break
		}
	case filterrequest.NilValKind:
		{
			result = filtergojsonq.NilValKind
			break
		}
	}
	return
}
