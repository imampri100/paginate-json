package filterrequest

type ActualPosition struct {
	wrapper PositionWhereClause
	current PositionWhereClause
}

func GenActualPosition(wrapper PositionWhereClause, current PositionWhereClause) (aps *ActualPosition) {
	aps = new(ActualPosition)
	aps.setWrapper(wrapper)
	aps.setCurrent(current)
	return
}

func (aps ActualPosition) Me() ActualPosition {
	return aps
}

func (aps *ActualPosition) ReplaceWrapper(wrapper PositionWhereClause) {
	aps.setWrapper(wrapper)
}

func (aps ActualPosition) Wrapper() (wrapper PositionWhereClause) {
	wrapper = aps.getWrapper()
	return
}

func (aps *ActualPosition) ReplaceCurrent(current PositionWhereClause) {
	aps.setCurrent(current)
}

func (aps ActualPosition) Current() (current PositionWhereClause) {
	current = aps.getCurrent()
	return
}

func (aps *ActualPosition) setWrapper(wrapper PositionWhereClause) {
	aps.wrapper = wrapper
}

func (aps ActualPosition) getWrapper() (wrapper PositionWhereClause) {
	wrapper = aps.wrapper
	return
}

func (aps *ActualPosition) setCurrent(current PositionWhereClause) {
	aps.current = current
}

func (aps ActualPosition) getCurrent() (current PositionWhereClause) {
	current = aps.current
	return
}
