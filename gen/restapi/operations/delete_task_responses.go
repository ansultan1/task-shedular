// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteTaskNoContentCode is the HTTP code returned for type DeleteTaskNoContent
const DeleteTaskNoContentCode int = 204

/*
DeleteTaskNoContent task deleted

swagger:response deleteTaskNoContent
*/
type DeleteTaskNoContent struct {
}

// NewDeleteTaskNoContent creates DeleteTaskNoContent with default headers values
func NewDeleteTaskNoContent() *DeleteTaskNoContent {

	return &DeleteTaskNoContent{}
}

// WriteResponse to the client
func (o *DeleteTaskNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteTaskNotFoundCode is the HTTP code returned for type DeleteTaskNotFound
const DeleteTaskNotFoundCode int = 404

/*
DeleteTaskNotFound task not found

swagger:response deleteTaskNotFound
*/
type DeleteTaskNotFound struct {
}

// NewDeleteTaskNotFound creates DeleteTaskNotFound with default headers values
func NewDeleteTaskNotFound() *DeleteTaskNotFound {

	return &DeleteTaskNotFound{}
}

// WriteResponse to the client
func (o *DeleteTaskNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteTaskInternalServerErrorCode is the HTTP code returned for type DeleteTaskInternalServerError
const DeleteTaskInternalServerErrorCode int = 500

/*
DeleteTaskInternalServerError internal server error

swagger:response deleteTaskInternalServerError
*/
type DeleteTaskInternalServerError struct {
}

// NewDeleteTaskInternalServerError creates DeleteTaskInternalServerError with default headers values
func NewDeleteTaskInternalServerError() *DeleteTaskInternalServerError {

	return &DeleteTaskInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteTaskInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
