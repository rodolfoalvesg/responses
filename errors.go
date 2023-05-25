package responses

import "fmt"

var (
	ErrorBadRequest          = BadRequestError{Title: "Bad Request", Description: "The request could not be understood by the server due to malformed syntax."}
	ErrorUnauthorized        = UnauthorizedError{Title: "Unauthorized", Description: "The request requires user authentication."}
	ErrorNotFound            = NotFoundError{Title: "Not Found", Description: "The server has not found anything matching the Request-URI."}
	ErrorInternalServerError = InternalServerError{Title: "Internal Server Error", Description: "The server encountered an unexpected condition which prevented it from fulfilling the request."}
	ErrorForbidden           = ForbiddenError{Title: "Forbidden", Description: "The server understood the request, but is refusing to fulfill it."}
	ErrorConflict            = ConflictError{Title: "Conflict", Description: "The request could not be completed due to a conflict with the current state of the resource."}
)

type BadRequestError struct {
	Type        string `json:"type" default:"bad_request"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UnauthorizedError struct {
	Type        string `json:"type" default:"unauthorized"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type NotFoundError struct {
	Type        string `json:"type" default:"not_found"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type InternalServerError struct {
	Type        string `json:"type" default:"internal_server_error"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ForbiddenError struct {
	Type        string `json:"type" default:"forbidden"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ConflictError struct {
	Type        string `json:"type" default:"conflict"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ValidationError struct {
	Param string
	Err   error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Param, e.Err.Error())
}
