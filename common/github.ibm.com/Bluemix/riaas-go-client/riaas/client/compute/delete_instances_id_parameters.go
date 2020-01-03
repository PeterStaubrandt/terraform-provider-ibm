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

// NewDeleteInstancesIDParams creates a new DeleteInstancesIDParams object
// with the default values initialized.
func NewDeleteInstancesIDParams() *DeleteInstancesIDParams {
	var ()
	return &DeleteInstancesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstancesIDParamsWithTimeout creates a new DeleteInstancesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstancesIDParamsWithTimeout(timeout time.Duration) *DeleteInstancesIDParams {
	var ()
	return &DeleteInstancesIDParams{

		timeout: timeout,
	}
}

// NewDeleteInstancesIDParamsWithContext creates a new DeleteInstancesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstancesIDParamsWithContext(ctx context.Context) *DeleteInstancesIDParams {
	var ()
	return &DeleteInstancesIDParams{

		Context: ctx,
	}
}

// NewDeleteInstancesIDParamsWithHTTPClient creates a new DeleteInstancesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstancesIDParamsWithHTTPClient(client *http.Client) *DeleteInstancesIDParams {
	var ()
	return &DeleteInstancesIDParams{
		HTTPClient: client,
	}
}

/*DeleteInstancesIDParams contains all the parameters to send to the API endpoint
for the delete instances ID operation typically these are written to a http.Request
*/
type DeleteInstancesIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The instance identifier

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

// WithTimeout adds the timeout to the delete instances ID params
func (o *DeleteInstancesIDParams) WithTimeout(timeout time.Duration) *DeleteInstancesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instances ID params
func (o *DeleteInstancesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instances ID params
func (o *DeleteInstancesIDParams) WithContext(ctx context.Context) *DeleteInstancesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instances ID params
func (o *DeleteInstancesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instances ID params
func (o *DeleteInstancesIDParams) WithHTTPClient(client *http.Client) *DeleteInstancesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instances ID params
func (o *DeleteInstancesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete instances ID params
func (o *DeleteInstancesIDParams) WithGeneration(generation int64) *DeleteInstancesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete instances ID params
func (o *DeleteInstancesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the delete instances ID params
func (o *DeleteInstancesIDParams) WithID(id string) *DeleteInstancesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete instances ID params
func (o *DeleteInstancesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete instances ID params
func (o *DeleteInstancesIDParams) WithVersion(version string) *DeleteInstancesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete instances ID params
func (o *DeleteInstancesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstancesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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