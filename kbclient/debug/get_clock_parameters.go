// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetClockParams creates a new GetClockParams object
// with the default values initialized.
func NewGetClockParams() *GetClockParams {
	var ()
	return &GetClockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClockParamsWithTimeout creates a new GetClockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClockParamsWithTimeout(timeout time.Duration) *GetClockParams {
	var ()
	return &GetClockParams{

		timeout: timeout,
	}
}

// NewGetClockParamsWithContext creates a new GetClockParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClockParamsWithContext(ctx context.Context) *GetClockParams {
	var ()
	return &GetClockParams{

		Context: ctx,
	}
}

// NewGetClockParamsWithHTTPClient creates a new GetClockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClockParamsWithHTTPClient(client *http.Client) *GetClockParams {
	var ()
	return &GetClockParams{
		HTTPClient: client,
	}
}

/*GetClockParams contains all the parameters to send to the API endpoint
for the get clock operation typically these are written to a http.Request
*/
type GetClockParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get clock params
func (o *GetClockParams) WithTimeout(timeout time.Duration) *GetClockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get clock params
func (o *GetClockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get clock params
func (o *GetClockParams) WithContext(ctx context.Context) *GetClockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get clock params
func (o *GetClockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get clock params
func (o *GetClockParams) WithHTTPClient(client *http.Client) *GetClockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get clock params
func (o *GetClockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get clock params
func (o *GetClockParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetClockParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get clock params
func (o *GetClockParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get clock params
func (o *GetClockParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetClockParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get clock params
func (o *GetClockParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WriteToRequest writes these params to a swagger request
func (o *GetClockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
