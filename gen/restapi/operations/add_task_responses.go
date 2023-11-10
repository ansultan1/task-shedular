// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/ansultan1/task-scheduler/gen/models"
)

// AddTaskCreatedCode is the HTTP code returned for type AddTaskCreated
const AddTaskCreatedCode int = 201

/*
AddTaskCreated task added

swagger:response addTaskCreated
*/
type AddTaskCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Task `json:"body,omitempty"`
}

// NewAddTaskCreated creates AddTaskCreated with default headers values
func NewAddTaskCreated() *AddTaskCreated {

	return &AddTaskCreated{}
}

// WithPayload adds the payload to the add task created response
func (o *AddTaskCreated) WithPayload(payload *models.Task) *AddTaskCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add task created response
func (o *AddTaskCreated) SetPayload(payload *models.Task) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddTaskCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddTaskConflictCode is the HTTP code returned for type AddTaskConflict
const AddTaskConflictCode int = 409

/*
AddTaskConflict task already exists

swagger:response addTaskConflict
*/
type AddTaskConflict struct {
}

// NewAddTaskConflict creates AddTaskConflict with default headers values
func NewAddTaskConflict() *AddTaskConflict {

	return &AddTaskConflict{}
}

// WriteResponse to the client
func (o *AddTaskConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(409)
}

// AddTaskInternalServerErrorCode is the HTTP code returned for type AddTaskInternalServerError
const AddTaskInternalServerErrorCode int = 500

/*
AddTaskInternalServerError internal server error

swagger:response addTaskInternalServerError
*/
type AddTaskInternalServerError struct {
}

// NewAddTaskInternalServerError creates AddTaskInternalServerError with default headers values
func NewAddTaskInternalServerError() *AddTaskInternalServerError {

	return &AddTaskInternalServerError{}
}

// WriteResponse to the client
func (o *AddTaskInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
