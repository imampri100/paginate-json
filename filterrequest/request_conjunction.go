package filterrequest

import (
	"fmt"
	"strings"
)

type RequestConjunction string

const (
	OrReqConj  RequestConjunction = "or"
	AndReqConj RequestConjunction = "and"
)

func GenRequestConjunction(input string) (rcj RequestConjunction, err error) {
	switch {
	case OrReqConj.IsMatch(input):
		{
			rcj = OrReqConj
			break
		}
	case AndReqConj.IsMatch(input):
		{
			rcj = AndReqConj
			break
		}
	default:
		{
			err = fmt.Errorf("conjunction %s is unknown", input)
			break
		}
	}
	return
}

func (rcj RequestConjunction) String() string {
	return string(rcj)
}

func (rcj RequestConjunction) IsMatch(input string) (isMatch bool) {
	input = strings.TrimSpace(input)
	isMatch = strings.EqualFold(input, rcj.String())
	return
}
