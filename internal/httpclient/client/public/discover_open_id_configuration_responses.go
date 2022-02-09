// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// DiscoverOpenIDConfigurationReader is a Reader for the DiscoverOpenIDConfiguration structure.
type DiscoverOpenIDConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscoverOpenIDConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiscoverOpenIDConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDiscoverOpenIDConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDiscoverOpenIDConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDiscoverOpenIDConfigurationOK creates a DiscoverOpenIDConfigurationOK with default headers values
func NewDiscoverOpenIDConfigurationOK() *DiscoverOpenIDConfigurationOK {
	return &DiscoverOpenIDConfigurationOK{}
}

/* DiscoverOpenIDConfigurationOK describes a response with status code 200, with default header values.

wellKnown
*/
type DiscoverOpenIDConfigurationOK struct {
	Payload *models.WellKnown
}

func (o *DiscoverOpenIDConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIdConfigurationOK  %+v", 200, o.Payload)
}
func (o *DiscoverOpenIDConfigurationOK) GetPayload() *models.WellKnown {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WellKnown)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverOpenIDConfigurationUnauthorized creates a DiscoverOpenIDConfigurationUnauthorized with default headers values
func NewDiscoverOpenIDConfigurationUnauthorized() *DiscoverOpenIDConfigurationUnauthorized {
	return &DiscoverOpenIDConfigurationUnauthorized{}
}

/* DiscoverOpenIDConfigurationUnauthorized describes a response with status code 401, with default header values.

jsonError
*/
type DiscoverOpenIDConfigurationUnauthorized struct {
	Payload *models.JSONError
}

func (o *DiscoverOpenIDConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIdConfigurationUnauthorized  %+v", 401, o.Payload)
}
func (o *DiscoverOpenIDConfigurationUnauthorized) GetPayload() *models.JSONError {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiscoverOpenIDConfigurationInternalServerError creates a DiscoverOpenIDConfigurationInternalServerError with default headers values
func NewDiscoverOpenIDConfigurationInternalServerError() *DiscoverOpenIDConfigurationInternalServerError {
	return &DiscoverOpenIDConfigurationInternalServerError{}
}

/* DiscoverOpenIDConfigurationInternalServerError describes a response with status code 500, with default header values.

jsonError
*/
type DiscoverOpenIDConfigurationInternalServerError struct {
	Payload *models.JSONError
}

func (o *DiscoverOpenIDConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /.well-known/openid-configuration][%d] discoverOpenIdConfigurationInternalServerError  %+v", 500, o.Payload)
}
func (o *DiscoverOpenIDConfigurationInternalServerError) GetPayload() *models.JSONError {
	return o.Payload
}

func (o *DiscoverOpenIDConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
