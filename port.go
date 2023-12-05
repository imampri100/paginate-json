package paginatejson

type IRequestHandler interface {
	HttpRequest(req interface{}) (resp IResponseHandler)
	ManualRequest(req PageRequest) (resp IResponseHandler)
}

type IResponseHandler interface {
	Response() (page Page)
}

type IParserHandler interface {
	PageRequest() PageRequest
}
