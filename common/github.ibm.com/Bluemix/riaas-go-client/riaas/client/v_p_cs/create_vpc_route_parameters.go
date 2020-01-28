// Code generated by go-swagger; DO NOT EDIT.

package v_p_cs

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

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// NewCreateVpcRouteParams creates a new CreateVpcRouteParams object
// with the default values initialized.
func NewCreateVpcRouteParams() *CreateVpcRouteParams {
	var ()
	return &CreateVpcRouteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVpcRouteParamsWithTimeout creates a new CreateVpcRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVpcRouteParamsWithTimeout(timeout time.Duration) *CreateVpcRouteParams {
	var ()
	return &CreateVpcRouteParams{

		timeout: timeout,
	}
}

// NewCreateVpcRouteParamsWithContext creates a new CreateVpcRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVpcRouteParamsWithContext(ctx context.Context) *CreateVpcRouteParams {
	var ()
	return &CreateVpcRouteParams{

		Context: ctx,
	}
}

// NewCreateVpcRouteParamsWithHTTPClient creates a new CreateVpcRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVpcRouteParamsWithHTTPClient(client *http.Client) *CreateVpcRouteParams {
	var ()
	return &CreateVpcRouteParams{
		HTTPClient: client,
	}
}

/*CreateVpcRouteParams contains all the parameters to send to the API endpoint
for the create vpc route operation typically these are written to a http.Request
*/
type CreateVpcRouteParams struct {

	/*RouteTemplate*/
	RouteTemplate *models.RouteTemplate
	/*Generation
	  The infrastructure generation for the request. For the API behavior documented here, use `1` or `2`.

	*/
	Generation int64
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string
	/*VpcID
	  The VPC identifier

	*/
	VpcID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create vpc route params
func (o *CreateVpcRouteParams) WithTimeout(timeout time.Duration) *CreateVpcRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create vpc route params
func (o *CreateVpcRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create vpc route params
func (o *CreateVpcRouteParams) WithContext(ctx context.Context) *CreateVpcRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create vpc route params
func (o *CreateVpcRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create vpc route params
func (o *CreateVpcRouteParams) WithHTTPClient(client *http.Client) *CreateVpcRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create vpc route params
func (o *CreateVpcRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRouteTemplate adds the routeTemplate to the create vpc route params
func (o *CreateVpcRouteParams) WithRouteTemplate(routeTemplate *models.RouteTemplate) *CreateVpcRouteParams {
	o.SetRouteTemplate(routeTemplate)
	return o
}

// SetRouteTemplate adds the routeTemplate to the create vpc route params
func (o *CreateVpcRouteParams) SetRouteTemplate(routeTemplate *models.RouteTemplate) {
	o.RouteTemplate = routeTemplate
}

// WithGeneration adds the generation to the create vpc route params
func (o *CreateVpcRouteParams) WithGeneration(generation int64) *CreateVpcRouteParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the create vpc route params
func (o *CreateVpcRouteParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithVersion adds the version to the create vpc route params
func (o *CreateVpcRouteParams) WithVersion(version string) *CreateVpcRouteParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the create vpc route params
func (o *CreateVpcRouteParams) SetVersion(version string) {
	o.Version = version
}

// WithVpcID adds the vpcID to the create vpc route params
func (o *CreateVpcRouteParams) WithVpcID(vpcID string) *CreateVpcRouteParams {
	o.SetVpcID(vpcID)
	return o
}

// SetVpcID adds the vpcId to the create vpc route params
func (o *CreateVpcRouteParams) SetVpcID(vpcID string) {
	o.VpcID = vpcID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVpcRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RouteTemplate != nil {
		if err := r.SetBodyParam(o.RouteTemplate); err != nil {
			return err
		}
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
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

	// path param vpc_id
	if err := r.SetPathParam("vpc_id", o.VpcID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}