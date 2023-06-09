package responses

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusOK(t *testing.T) {
	t.Parallel()

	want := Response{
		Status:  http.StatusOK,
		Payload: "OK",
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	got := StatusOK("OK", want.Headers)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Payload, got.Payload)
}

func TestStatusCreated(t *testing.T) {
	t.Parallel()

	want := Response{
		Status:  http.StatusCreated,
		Payload: "Created",
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	got := StatusCreated("Created", want.Headers)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Payload, got.Payload)
}

func TestStatusAccepted(t *testing.T) {
	t.Parallel()

	want := Response{
		Status:  http.StatusAccepted,
		Payload: "Accepted",
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	got := StatusAccepted("Accepted", want.Headers)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Payload, got.Payload)
}

func TestStatusNoContent(t *testing.T) {
	t.Parallel()

	want := Response{
		Status: http.StatusNoContent,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	got := StatusNoContent(want.Headers)

	require.Equal(t, want.Status, got.Status)
}

func TestStatusBadRequest(t *testing.T) {
	t.Parallel()

	errBadRequest := errors.New("Bad Request")

	want := Response{
		Status:  http.StatusBadRequest,
		Payload: ErrorBadRequest,
		Error:   errBadRequest,
	}

	got := StatusBadRequest(errBadRequest)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Payload, got.Payload)
	assert.Equal(t, want.Error, got.Error)
}

func TestStatusUnauthorized(t *testing.T) {
	t.Parallel()

	errUnauthorized := errors.New("Unauthorized")

	want := Response{
		Status:  http.StatusUnauthorized,
		Payload: ErrorUnauthorized,
		Error:   errUnauthorized,
	}

	got := StatusUnauthorized(errUnauthorized)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Payload, got.Payload)
	assert.Equal(t, want.Error, got.Error)
}

func TestStatusForbidden(t *testing.T) {
	t.Parallel()

	errForbidden := errors.New("Forbidden")

	want := Response{
		Status:  http.StatusForbidden,
		Payload: ErrorForbidden,
		Error:   errForbidden,
	}

	got := StatusForbidden(errForbidden)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Payload, got.Payload)
	assert.Equal(t, want.Error, got.Error)
}

func TestStatusNotFound(t *testing.T) {
	t.Parallel()

	errNotFound := errors.New("Not Found")

	want := Response{
		Status:  http.StatusNotFound,
		Payload: ErrorNotFound,
		Error:   errNotFound,
	}

	got := StatusNotFound(errNotFound)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Payload, got.Payload)
	assert.Equal(t, want.Error, got.Error)
}

func TestStatusInternalServerError(t *testing.T) {
	t.Parallel()

	errInternalServerError := errors.New("Internal Server Error")

	want := Response{
		Status: http.StatusInternalServerError,
		Error:  errInternalServerError,
	}

	got := StatusInternalServerError(errInternalServerError)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Error, got.Error)
}

func TestStatusNotImplemented(t *testing.T) {
	t.Parallel()

	errNotImplemented := errors.New("Not Implemented")

	want := Response{
		Status: http.StatusNotImplemented,
		Error:  errNotImplemented,
	}

	got := StatusNotImplemented(errNotImplemented)

	require.Equal(t, want.Status, got.Status)
	assert.Equal(t, want.Error, got.Error)
}
