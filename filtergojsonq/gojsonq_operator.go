package filtergojsonq

type GojsonqOperator string

const (
	IsGojsOpr             GojsonqOperator = "="
	IsNotGojsOpr          GojsonqOperator = "!="
	GtGojsOpr             GojsonqOperator = ">"
	GteGojsOpr            GojsonqOperator = ">="
	LtGojsOpr             GojsonqOperator = "<"
	LteGojsOpr            GojsonqOperator = "<="
	InGojsOpr             GojsonqOperator = "in"
	NotInGojsOpr          GojsonqOperator = "notIn"
	StrictContainsGojsOpr GojsonqOperator = "strictContains"
	ContainsGojsOpr       GojsonqOperator = "contains"
)

func (gjo GojsonqOperator) MustSliceValue() (mustSlice bool) {
	switch gjo {
	case IsGojsOpr,
		IsNotGojsOpr,
		GtGojsOpr,
		GteGojsOpr,
		LtGojsOpr,
		LteGojsOpr,
		StrictContainsGojsOpr,
		ContainsGojsOpr:
		{
			mustSlice = false
			break
		}
	case InGojsOpr,
		NotInGojsOpr:
		{
			mustSlice = true
			break
		}
	}

	return
}

func (gjo GojsonqOperator) String() string {
	return string(gjo)
}
