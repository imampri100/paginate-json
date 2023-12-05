package filterrequest

type ListConjunctionWhere []ConjunctionWhere

func GenListConjunctionWhere(listIdentity ListIdentityWhereClause) (lcw *ListConjunctionWhere) {
	lcw = new(ListConjunctionWhere)

	var conjunction IdentityWhereClause

	for idxId := range listIdentity {
		itemId := listIdentity[idxId]

		if itemId.Length() == 0 {
			continue
		}

		if itemId.FirstIndex().Identity() == ConjunctionFirstIndex {
			conjunction = itemId
			continue
		}

		cwh := GenConjunctionWhere(conjunction, itemId)

		lcw.appendData(cwh.Me())

		// reset conjunction
		conjunction = IdentityWhereClause{}
	}

	return
}

func (lwc ListConjunctionWhere) Me() ListConjunctionWhere {
	return lwc
}

func (lwc *ListConjunctionWhere) appendData(conjunctionWhere ConjunctionWhere) {
	(*lwc) = append((*lwc), conjunctionWhere)
}

type ConjunctionWhere struct {
	conjunctionIdentity IdentityWhereClause
	whereIdentity       IdentityWhereClause
}

func GenConjunctionWhere(conjunctionIdentity, whereIdentity IdentityWhereClause) (cwh *ConjunctionWhere) {
	cwh = new(ConjunctionWhere)
	cwh.setConjunctionIdentity(conjunctionIdentity)
	cwh.setWhereIdentity(whereIdentity)
	return
}

func (cwh ConjunctionWhere) Me() ConjunctionWhere {
	return cwh
}

func (cwh ConjunctionWhere) ConjunctionIdentity() (conjunctionIdentity IdentityWhereClause) {
	conjunctionIdentity = cwh.getConjunctionIdentity()
	return
}

func (cwh ConjunctionWhere) WhereIdentity() (whereIdentity IdentityWhereClause) {
	whereIdentity = cwh.getWhereIdentity()
	return
}

func (cwh *ConjunctionWhere) setConjunctionIdentity(conjunctionIdentity IdentityWhereClause) {
	cwh.conjunctionIdentity = conjunctionIdentity
}

func (cwh ConjunctionWhere) getConjunctionIdentity() (conjunctionIdentity IdentityWhereClause) {
	conjunctionIdentity = cwh.conjunctionIdentity
	return
}

func (cwh *ConjunctionWhere) setWhereIdentity(whereIdentity IdentityWhereClause) {
	cwh.whereIdentity = whereIdentity
}

func (cwh ConjunctionWhere) getWhereIdentity() (whereIdentity IdentityWhereClause) {
	whereIdentity = cwh.whereIdentity
	return
}
