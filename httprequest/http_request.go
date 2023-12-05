package httprequest

type BuilderHttpRequest struct {
	Config IConfigRequest
}

type HttpRequest struct {
	builder        BuilderHttpRequest
	httpRawRequest interface{}
	generic        *HttpRequestGeneric
	parser         *HttpRequestParser
	err            error
}

func NewHttpRequest(builder BuilderHttpRequest) (hrq *HttpRequest) {
	hrq = new(HttpRequest)

	// builder
	hrq.setBuilder(builder)

	return
}

func (hrq HttpRequest) Me() HttpRequest {
	return hrq
}

func (hrq *HttpRequest) Query(httpRawRequest interface{}) (result IQuery) {
	// clearErr first
	hrq.clearErr()

	// httpRawRequest
	hrq.setHttpRawRequest(httpRawRequest)

	// generic
	generic, errGeneric := hrq.genGeneric()
	if errGeneric != nil {
		hrq.setErr(errGeneric)
		return
	}
	hrq.setGeneric(generic)

	// parser
	parser, errParser := hrq.genParser()
	if errParser != nil {
		hrq.setErr(errParser)
		return
	}
	hrq.setParser(parser)

	// parseRequest
	resParse, errParse := hrq.parseRequest()
	if errParser != nil {
		hrq.setErr(errParse)
		return
	}

	result = resParse
	return
}

func (hrq HttpRequest) Err() (err error) {
	err = hrq.getErr()
	return
}

func (hrq *HttpRequest) setBuilder(builder BuilderHttpRequest) {
	hrq.builder = builder
}

func (hrq HttpRequest) getBuilder() (builder BuilderHttpRequest) {
	builder = hrq.builder
	return
}

func (hpu *HttpRequest) setHttpRawRequest(httpRawRequest interface{}) {
	hpu.httpRawRequest = httpRawRequest
}

func (hpu HttpRequest) getHttpRawRequest() interface{} {
	return hpu.httpRawRequest
}

func (hrq *HttpRequest) setGeneric(generic *HttpRequestGeneric) {
	hrq.generic = generic
}

func (hrq HttpRequest) getGeneric() (generic *HttpRequestGeneric) {
	generic = hrq.generic
	return
}

func (hrq HttpRequest) genGeneric() (generic *HttpRequestGeneric, err error) {
	httpReq := hrq.getHttpRawRequest()

	builderHrg := BuilderHttpRequestGeneric{
		HttpRequestRaw: httpReq,
	}

	hrg := NewHttpRequestGeneric(builderHrg)
	if hrg.Err() != nil {
		err = hrg.Err()
		return
	}

	generic = hrg
	return
}

func (hrq *HttpRequest) setParser(parser *HttpRequestParser) {
	hrq.parser = parser
}

func (hrq HttpRequest) getParser() (parser *HttpRequestParser) {
	parser = hrq.parser
	return
}

func (hrq HttpRequest) genParser() (parser *HttpRequestParser, err error) {
	builder := hrq.getBuilder()
	config := builder.Config

	builderHrp := BuilderHttpRequestParser{
		Config: config,
	}

	hrp := NewHttpRequestParser(builderHrp)
	if hrp.Err() != nil {
		err = hrp.Err()
		return
	}

	parser = hrp
	return
}

func (hrq HttpRequest) parseRequest() (query IQuery, err error) {
	parser := hrq.getParser()
	generic := hrq.getGeneric()

	parser.Parse(generic)
	if parser.Err() != nil {
		err = parser.Err()
		return
	}

	query = parser
	return
}

func (hrq *HttpRequest) setErr(err error) {
	hrq.err = err
}

func (hrq *HttpRequest) clearErr() {
	hrq.err = nil
}

func (hrq HttpRequest) getErr() (err error) {
	err = hrq.err
	return
}
