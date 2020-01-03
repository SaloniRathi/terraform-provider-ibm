// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewGetFloatingIpsIDParams creates a new GetFloatingIpsIDParams object
// with the default values initialized.
func NewGetFloatingIpsIDParams() *GetFloatingIpsIDParams {
	var ()
	return &GetFloatingIpsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFloatingIpsIDParamsWithTimeout creates a new GetFloatingIpsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFloatingIpsIDParamsWithTimeout(timeout time.Duration) *GetFloatingIpsIDParams {
	var ()
	return &GetFloatingIpsIDParams{

		timeout: timeout,
	}
}

// NewGetFloatingIpsIDParamsWithContext creates a new GetFloatingIpsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFloatingIpsIDParamsWithContext(ctx context.Context) *GetFloatingIpsIDParams {
	var ()
	return &GetFloatingIpsIDParams{

		Context: ctx,
	}
}

// NewGetFloatingIpsIDParamsWithHTTPClient creates a new GetFloatingIpsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFloatingIpsIDParamsWithHTTPClient(client *http.Client) *GetFloatingIpsIDParams {
	var ()
	return &GetFloatingIpsIDParams{
		HTTPClient: client,
	}
}

/*GetFloatingIpsIDParams contains all the parameters to send to the API endpoint
for the get floating ips ID operation typically these are written to a http.Request
*/
type GetFloatingIpsIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The floating ip identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get floating ips ID params
func (o *GetFloatingIpsIDParams) WithTimeout(timeout time.Duration) *GetFloatingIpsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get floating ips ID params
func (o *GetFloatingIpsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get floating ips ID params
func (o *GetFloatingIpsIDParams) WithContext(ctx context.Context) *GetFloatingIpsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get floating ips ID params
func (o *GetFloatingIpsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get floating ips ID params
func (o *GetFloatingIpsIDParams) WithHTTPClient(client *http.Client) *GetFloatingIpsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get floating ips ID params
func (o *GetFloatingIpsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get floating ips ID params
func (o *GetFloatingIpsIDParams) WithGeneration(generation int64) *GetFloatingIpsIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get floating ips ID params
func (o *GetFloatingIpsIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get floating ips ID params
func (o *GetFloatingIpsIDParams) WithID(id string) *GetFloatingIpsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get floating ips ID params
func (o *GetFloatingIpsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get floating ips ID params
func (o *GetFloatingIpsIDParams) WithVersion(version string) *GetFloatingIpsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get floating ips ID params
func (o *GetFloatingIpsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetFloatingIpsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
