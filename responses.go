package responses

import "net/http"

// crete a struct to hold the response
type Response struct {
	Status  int
	Payload interface{}
	Error   error
	Headers map[string]string
}

// Response status OK
func StatusOK(payload interface{}, headers map[string]string) Response {
	return Response{
		Status:  http.StatusOK,
		Payload: payload,
		Headers: headers,
	}
}

// Response status Created
func StatusCreated(payload interface{}, headers map[string]string) Response {
	return Response{
		Status:  http.StatusCreated,
		Payload: payload,
		Headers: headers,
	}
}

// Response status Accepted
func StatusAccepted(payload interface{}, header map[string]string) Response {
	return Response{
		Status:  http.StatusAccepted,
		Payload: payload,
		Headers: header,
	}
}

// Response status NoContent
func StatusNoContent(header map[string]string) Response {
	return Response{
		Status:  http.StatusNoContent,
		Headers: header,
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
