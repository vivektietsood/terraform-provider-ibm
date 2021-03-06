// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetSharesIDReader is a Reader for the GetSharesID structure.
type GetSharesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSharesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSharesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSharesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSharesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSharesIDOK creates a GetSharesIDOK with default headers values
func NewGetSharesIDOK() *GetSharesIDOK {
	return &GetSharesIDOK{}
}

/*GetSharesIDOK handles this case with default header values.

dummy
*/
type GetSharesIDOK struct {
	Payload *models.Share
}

func (o *GetSharesIDOK) Error() string {
	return fmt.Sprintf("[GET /shares/{id}][%d] getSharesIdOK  %+v", 200, o.Payload)
}

func (o *GetSharesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Share)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSharesIDNotFound creates a GetSharesIDNotFound with default headers values
func NewGetSharesIDNotFound() *GetSharesIDNotFound {
	return &GetSharesIDNotFound{}
}

/*GetSharesIDNotFound handles this case with default header values.

error
*/
type GetSharesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetSharesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /shares/{id}][%d] getSharesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSharesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSharesIDInternalServerError creates a GetSharesIDInternalServerError with default headers values
func NewGetSharesIDInternalServerError() *GetSharesIDInternalServerError {
	return &GetSharesIDInternalServerError{}
}

/*GetSharesIDInternalServerError handles this case with default header values.

error
*/
type GetSharesIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetSharesIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /shares/{id}][%d] getSharesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSharesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
