package api

const (
	HeaderRequestId string = "X-Request-Id"

	CtxRequestId     string = "request.id"
	CtxRequestMethod string = "request.method"
	CtxRequestHost   string = "request.host"
	CtxContentLength string = "request.content_length"
	CtxTlsVersion    string = "request.tls_version"
	CtxRequestPath   string = "request.path"
	CtxRequestQuery  string = "request.query"
	CtxLogger        string = "logger"
)
