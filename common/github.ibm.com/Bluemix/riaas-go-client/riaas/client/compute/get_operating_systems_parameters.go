// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOperatingSystemsParams creates a new GetOperatingSystemsParams object
// with the default values initialized.
func NewGetOperatingSystemsParams() *GetOperatingSystemsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetOperatingSystemsParams{
		Limit: &limitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOperatingSystemsParamsWithTimeout creates a new GetOperatingSystemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOperatingSystemsParamsWithTimeout(timeout time.Duration) *GetOperatingSystemsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetOperatingSystemsParams{
		Limit: &limitDefault,

		timeout: timeout,
	}
}

// NewGetOperatingSystemsParamsWithContext creates a new GetOperatingSystemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOperatingSystemsParamsWithContext(ctx context.Context) *GetOperatingSystemsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetOperatingSystemsParams{
		Limit: &limitDefault,

		Context: ctx,
	}
}

// NewGetOperatingSystemsParamsWithHTTPClient creates a new GetOperatingSystemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOperatingSystemsParamsWithHTTPClient(client *http.Client) *GetOperatingSystemsParams {
	var (
		limitDefault = int32(50)
	)
	return &GetOperatingSystemsParams{
		Limit:      &limitDefault,
		HTTPClient: client,
	}
}

/*GetOperatingSystemsParams contains all the parameters to send to the API endpoint
for the get operating systems operation typically these are written to a http.Request
*/
type GetOperatingSystemsParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Limit
	  The number of resources to return on a page

	*/
	Limit *int32
	/*Start
	  A server-supplied token determining what resource to start the page on

	*/
	Start *string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get operating systems params
func (o *GetOperatingSystemsParams) WithTimeout(timeout time.Duration) *GetOperatingSystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get operating systems params
func (o *GetOperatingSystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get operating systems params
func (o *GetOperatingSystemsParams) WithContext(ctx context.Context) *GetOperatingSystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get operating systems params
func (o *GetOperatingSystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get operating systems params
func (o *GetOperatingSystemsParams) WithHTTPClient(client *http.Client) *GetOperatingSystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get operating systems params
func (o *GetOperatingSystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get operating systems params
func (o *GetOperatingSystemsParams) WithGeneration(generation int64) *GetOperatingSystemsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get operating systems params
func (o *GetOperatingSystemsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithLimit adds the limit to the get operating systems params
func (o *GetOperatingSystemsParams) WithLimit(limit *int32) *GetOperatingSystemsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get operating systems params
func (o *GetOperatingSystemsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithStart adds the start to the get operating systems params
func (o *GetOperatingSystemsParams) WithStart(start *string) *GetOperatingSystemsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get operating systems params
func (o *GetOperatingSystemsParams) SetStart(start *string) {
	o.Start = start
}

// WithVersion adds the version to the get operating systems params
func (o *GetOperatingSystemsParams) WithVersion(version string) *GetOperatingSystemsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get operating systems params
func (o *GetOperatingSystemsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetOperatingSystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart string
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}