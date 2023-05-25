package responses

import "net/http"

// crete a struct to hold the response
type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"data"`
	Error   error       `json:"error"`
	header  http.Header `json:"-"`
}

// Status codes
const (
	OK                  = 200
	Created             = 201
	Accepted            = 202
	NoContent           = 204
	BadRequest          = 400
	Unauthorized        = 401
	Forbidden           = 403
	NotFound            = 404
	Conflict            = 409
	InternalServerError = 500
	NotImplemented      = 501
)

// Return Header
func (r Response) Header() http.Header {
	return r.header
}

// Response status OK
func StatusOK(payload interface{}) Response {
	return Response{
		Status:  OK,
		Payload: payload,
	}
}

// Response status Created
func StatusCreated(payload interface{}) Response {
	return Response{
		Status:  Created,
		Payload: payload,
	}
}

// Response status Accepted
func StatusAccepted(payload interface{}) Response {
	return Response{
		Status:  Accepted,
		Payload: payload,
	}
}

// Response status NoContent
func StatusNoContent() Response {
	return Response{
		Status: NoContent,
	}
}

// Response status BadRequest
func StatusBadRequest(payload interface{}, err error) Response {
	return Response{
		Status:  BadRequest,
		Error:   err,
		Payload: payload,
	}
}

// Response status Unauthorized
func StatusUnauthorized(payload interface{}, err error) Response {
	return Response{
		Status:  Unauthorized,
		Error:   err,
		Payload: payload,
	}
}

// Response status Forbidden
func StatusForbidden(payload interface{}, err error) Response {
	return Response{
		Status:  Forbidden,
		Error:   err,
		Payload: payload,
	}
}

// Response status NotFound
func StatusNotFound(payload interface{}, err error) Response {
	return Response{
		Status:  NotFound,
		Error:   err,
		Payload: payload,
	}
}

// Response status Conflict
func StatusConflict(payload interface{}, err error) Response {
	return Response{
		Status:  Conflict,
		Error:   err,
		Payload: payload,
	}
}

// Response status InternalServerError
func StatusInternalServerError(err error) Response {
	return Response{
		Status: InternalServerError,
		Error:  err,
	}
}

// Response status NotImplemented
func StatusNotImplemented(err error) Response {
	return Response{
		Status: NotImplemented,
		Error:  err,
	}
}

// Validate error
func (r Response) Validate() error {
	if r.Status >= BadRequest && r.Error != nil {
		return r.Error
	}

	return nil
}
