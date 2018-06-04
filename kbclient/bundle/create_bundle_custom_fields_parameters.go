// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewCreateBundleCustomFieldsParams creates a new CreateBundleCustomFieldsParams object
// with the default values initialized.
func NewCreateBundleCustomFieldsParams() *CreateBundleCustomFieldsParams {
	var ()
	return &CreateBundleCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBundleCustomFieldsParamsWithTimeout creates a new CreateBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBundleCustomFieldsParamsWithTimeout(timeout time.Duration) *CreateBundleCustomFieldsParams {
	var ()
	return &CreateBundleCustomFieldsParams{

		timeout: timeout,
	}
}

// NewCreateBundleCustomFieldsParamsWithContext creates a new CreateBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBundleCustomFieldsParamsWithContext(ctx context.Context) *CreateBundleCustomFieldsParams {
	var ()
	return &CreateBundleCustomFieldsParams{

		Context: ctx,
	}
}

// NewCreateBundleCustomFieldsParamsWithHTTPClient creates a new CreateBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBundleCustomFieldsParamsWithHTTPClient(client *http.Client) *CreateBundleCustomFieldsParams {
	var ()
	return &CreateBundleCustomFieldsParams{
		HTTPClient: client,
	}
}

/*CreateBundleCustomFieldsParams contains all the parameters to send to the API endpoint
for the create bundle custom fields operation typically these are written to a http.Request
*/
type CreateBundleCustomFieldsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body []*kbmodel.CustomField
	/*BundleID*/
	BundleID strfmt.UUID

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithTimeout(timeout time.Duration) *CreateBundleCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithContext(ctx context.Context) *CreateBundleCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithHTTPClient(client *http.Client) *CreateBundleCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateBundleCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateBundleCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreateBundleCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateBundleCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreateBundleCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreateBundleCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithBundleID adds the bundleID to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) WithBundleID(bundleID strfmt.UUID) *CreateBundleCustomFieldsParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the create bundle custom fields params
func (o *CreateBundleCustomFieldsParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBundleCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}

	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
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