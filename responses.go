package responses

import "net/http"

// crete a struct to hold the response
type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"data"`
	Error   error       `json:"error"`
	header  http.Header `json:"-"`
}

// Return Header
func (r Response) Header() http.Header {
	return r.header
}

// Response status OK
func StatusOK(payload interface{}) Response {
	return Response{
		Status:  http.StatusOK,
		Payload: payload,
	}
}

// Response status Created
func StatusCreated(payload interface{}) Response {
	return Response{
		Status:  http.StatusCreated,
		Payload: payload,
	}
}

// Response status Accepted
func StatusAccepted(payload interface{}) Response {
	return Response{
		Status:  http.StatusAccepted,
		Payload: payload,
	}
}

// Response status NoContent
func StatusNoContent() Response {
	return Response{
		Status: http.StatusNoContent,
	}
}

// Response status BadRequest
func StatusBadRequest(err error) Response {
	return Response{
		Status:  http.StatusBadRequest,
		Error:   err,
		Payload: ErrorBadRequest,
	}
}

// Response status Unauthorized
func StatusUnauthorized(err error) Response {
	return Response{
		Status:  http.StatusUnauthorized,
		Error:   err,
		Payload: ErrorUnauthorized,
	}
}

// Response status Forbidden
func StatusForbidden(err error) Response {
	return Response{
		Status:  http.StatusForbidden,
		Error:   err,
		Payload: ErrorForbidden,
	}
}

// Response status NotFound
func StatusNotFound(err error) Response {
	return Response{
		Status:  http.StatusNotFound,
		Error:   err,
		Payload: ErrorNotFound,
	}
}

// Response status Conflict
func StatusConflict(payload interface{}, err error) Response {
	return Response{
		Status:  http.StatusConflict,
		Error:   err,
		Payload: ErrorConflict,
	}
}

// Response status InternalServerError
func StatusInternalServerError(err error) Response {
	return Response{
		Status: http.StatusInternalServerError,
		Error:  err,
	}
}

// Response status NotImplemented
func StatusNotImplemented(err error) Response {
	return Response{
		Status: http.StatusNotImplemented,
		Error:  err,
	}
}

// Validate error
func (r Response) Validate() error {
	if r.Status >= http.StatusBadRequest && r.Error != nil {
		return r.Error
	}

	return nil
}
