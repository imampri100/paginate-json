package filtergojsonq

import (
	"fmt"
	"strings"
)

type GojsonqConjunction string

const (
	OrGojsConj  GojsonqConjunction = "or"
	AndGojsConj GojsonqConjunction = "and"
)

func GenGojsonqConjunction(input string) (gjc GojsonqConjunction, err error) {
	switch {
	case OrGojsConj.IsMatch(input):
		{
			gjc = OrGojsConj
			break
		}
	case AndGojsConj.IsMatch(input):
		{
			gjc = AndGojsConj
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

func (gjc GojsonqConjunction) String() string {
	return string(gjc)
}

func (gjc GojsonqConjunction) IsMatch(input string) (isMatch bool) {
	input = strings.TrimSpace(input)
	isMatch = strings.EqualFold(input, gjc.String())
	return
}
