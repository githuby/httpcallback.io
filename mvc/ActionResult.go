package mvc

import (
	"net/http"
)

type HttpStatusCode int

// Represents a result of a controller method.
// The action can write the actual result to
// an http response stream.
//
// There are diffent types of ActionResults:
// - JsonResult
// - HttpStatusResult
// - etc...
type ActionResult interface {
	// Write the actual result to the http response stream.
	WriteResponse(http.ResponseWriter)
}

func (statusCode HttpStatusCode) WriteResponse(response http.ResponseWriter) {
	response.WriteHeader(int(statusCode))
}
