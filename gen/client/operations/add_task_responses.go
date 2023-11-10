// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ansultan1/task-scheduler/gen/models"
)

// AddTaskReader is a Reader for the AddTask structure.
type AddTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddTaskCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewAddTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddTaskInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /task] addTask", response, response.Code())
	}
}

// NewAddTaskCreated creates a AddTaskCreated with default headers values
func NewAddTaskCreated() *AddTaskCreated {
	return &AddTaskCreated{}
}

/*
AddTaskCreated describes a response with status code 201, with default header values.

task added
*/
type AddTaskCreated struct {
	Payload *models.Task
}

// IsSuccess returns true when this add task created response has a 2xx status code
func (o *AddTaskCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add task created response has a 3xx status code
func (o *AddTaskCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add task created response has a 4xx status code
func (o *AddTaskCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add task created response has a 5xx status code
func (o *AddTaskCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add task created response a status code equal to that given
func (o *AddTaskCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add task created response
func (o *AddTaskCreated) Code() int {
	return 201
}

func (o *AddTaskCreated) Error() string {
	return fmt.Sprintf("[POST /task][%d] addTaskCreated  %+v", 201, o.Payload)
}

func (o *AddTaskCreated) String() string {
	return fmt.Sprintf("[POST /task][%d] addTaskCreated  %+v", 201, o.Payload)
}

func (o *AddTaskCreated) GetPayload() *models.Task {
	return o.Payload
}

func (o *AddTaskCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddTaskConflict creates a AddTaskConflict with default headers values
func NewAddTaskConflict() *AddTaskConflict {
	return &AddTaskConflict{}
}

/*
AddTaskConflict describes a response with status code 409, with default header values.

task already exists
*/
type AddTaskConflict struct {
}

// IsSuccess returns true when this add task conflict response has a 2xx status code
func (o *AddTaskConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add task conflict response has a 3xx status code
func (o *AddTaskConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add task conflict response has a 4xx status code
func (o *AddTaskConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add task conflict response has a 5xx status code
func (o *AddTaskConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add task conflict response a status code equal to that given
func (o *AddTaskConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add task conflict response
func (o *AddTaskConflict) Code() int {
	return 409
}

func (o *AddTaskConflict) Error() string {
	return fmt.Sprintf("[POST /task][%d] addTaskConflict ", 409)
}

func (o *AddTaskConflict) String() string {
	return fmt.Sprintf("[POST /task][%d] addTaskConflict ", 409)
}

func (o *AddTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddTaskInternalServerError creates a AddTaskInternalServerError with default headers values
func NewAddTaskInternalServerError() *AddTaskInternalServerError {
	return &AddTaskInternalServerError{}
}

/*
AddTaskInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type AddTaskInternalServerError struct {
}

// IsSuccess returns true when this add task internal server error response has a 2xx status code
func (o *AddTaskInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add task internal server error response has a 3xx status code
func (o *AddTaskInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add task internal server error response has a 4xx status code
func (o *AddTaskInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add task internal server error response has a 5xx status code
func (o *AddTaskInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add task internal server error response a status code equal to that given
func (o *AddTaskInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add task internal server error response
func (o *AddTaskInternalServerError) Code() int {
	return 500
}

func (o *AddTaskInternalServerError) Error() string {
	return fmt.Sprintf("[POST /task][%d] addTaskInternalServerError ", 500)
}

func (o *AddTaskInternalServerError) String() string {
	return fmt.Sprintf("[POST /task][%d] addTaskInternalServerError ", 500)
}

func (o *AddTaskInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
