package httputils

const (
	InternalServerErrMsg      = "Unable to process the request due to an internal error. Please contact the system administrator"
	InvalidPayloadErrMsg      = "Payload is not valid."
	ServerStartupMsg          = "Starting the API server..."
	ServerStartupSuccessMsg   = "The API server has started and is listening on %s"
	ServerUrlMsg              = "API server url: %s"
	ServerAPIDocsMsg          = "API server swagger docs url: %s"
	ServerStartupErrMsg       = "Unable to run the web server"
	NotImplementedYetMsg      = "Not Implemented"
	AuthorizationErrMsg       = "Insufficient privileges"
	ApiHealthPath             = "/health"
	ContentTypeHeaderKey      = "Content-Type"
	AcceptHeaderKey           = "Accept"
	AuthorizationHeaderKey    = "Authorization"
	ApplicationJsonMIMEType   = "application/json"
	TextPlainMIMEType         = "text/plain"
	XwwwFromUrlencodeMIMEType = "application/x-www-form-urlencoded"
	SwaggerPath               = "/swagger/*any"
	SwaggerSpecPathFormat     = "%s/swagger/doc.json"
	SwaggerUriPathFormat      = "%s/swagger/index.html"
)

type RestMsg struct {
	Message string `json:"message"`
}

type RestErrMsg struct {
	Error string `json:"error"`
}
