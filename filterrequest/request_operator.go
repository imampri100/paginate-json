package filterrequest

import (
	"fmt"
	"strings"
)

type RequestOperator string

const (
	_                           = ""
	IsReqOpr    RequestOperator = "="
	IsNotReqOpr RequestOperator = "!="
	GtReqOpr    RequestOperator = ">"
	GteReqOpr   RequestOperator = ">="
	LtReqOpr    RequestOperator = "<"
	LteReqOpr   RequestOperator = "<="
	InReqOpr    RequestOperator = "in"
	NotInReqOpr RequestOperator = "not in"
	LikeReqOpr  RequestOperator = "like"
	IlikeReqOpr RequestOperator = "ilike"
	// NotLikeReqOpr RequestOperator = "not like"; Not defined in GoJsonq
)

func GenRequestOperator(input string) (result RequestOperator, err error) {
	switch {
	case IsReqOpr.IsMatch(input):
		{
			result = IsReqOpr
			break
		}
	case IsNotReqOpr.IsMatch(input):
		{
			result = IsNotReqOpr
			break
		}
	case GtReqOpr.IsMatch(input):
		{
			result = GtReqOpr
			break
		}
	case GteReqOpr.IsMatch(input):
		{
			result = GteReqOpr
			break
		}
	case LtReqOpr.IsMatch(input):
		{
			result = LtReqOpr
			break
		}
	case LteReqOpr.IsMatch(input):
		{
			result = LteReqOpr
			break
		}
	case InReqOpr.IsMatch(input):
		{
			result = InReqOpr
			break
		}
	case NotInReqOpr.IsMatch(input):
		{
			result = NotInReqOpr
			break
		}
	case LikeReqOpr.IsMatch(input):
		{
			result = LikeReqOpr
			break
		}
	case IlikeReqOpr.IsMatch(input):
		{
			result = IlikeReqOpr
			break
		}
	default:
		{
			err = fmt.Errorf("operator %s is unknown", input)
			break
		}
	}
	return
}

func (rop RequestOperator) IsMatch(input string) (isMatch bool) {
	input = strings.TrimSpace(input)
	isMatch = strings.EqualFold(input, rop.String())
	return
}

func (rop RequestOperator) String() string {
	return string(rop)
}

func (rop RequestOperator) MustSliceValue() (mustSlice bool) {
	switch rop {
	case IsReqOpr,
		IsNotReqOpr,
		GtReqOpr,
		GteReqOpr,
		LtReqOpr,
		LteReqOpr,
		LikeReqOpr,
		IlikeReqOpr:
		{
			mustSlice = false
			break
		}
	case InReqOpr,
		NotInReqOpr:
		{
			mustSlice = true
			break
		}
	}

	return
}
