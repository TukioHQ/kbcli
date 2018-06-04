// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewChargebackReversalPaymentParams creates a new ChargebackReversalPaymentParams object
// with the default values initialized.
func NewChargebackReversalPaymentParams() *ChargebackReversalPaymentParams {
	var ()
	return &ChargebackReversalPaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChargebackReversalPaymentParamsWithTimeout creates a new ChargebackReversalPaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChargebackReversalPaymentParamsWithTimeout(timeout time.Duration) *ChargebackReversalPaymentParams {
	var ()
	return &ChargebackReversalPaymentParams{

		timeout: timeout,
	}
}

// NewChargebackReversalPaymentParamsWithContext creates a new ChargebackReversalPaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewChargebackReversalPaymentParamsWithContext(ctx context.Context) *ChargebackReversalPaymentParams {
	var ()
	return &ChargebackReversalPaymentParams{

		Context: ctx,
	}
}

// NewChargebackReversalPaymentParamsWithHTTPClient creates a new ChargebackReversalPaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChargebackReversalPaymentParamsWithHTTPClient(client *http.Client) *ChargebackReversalPaymentParams {
	var ()
	return &ChargebackReversalPaymentParams{
		HTTPClient: client,
	}
}

/*ChargebackReversalPaymentParams contains all the parameters to send to the API endpoint
for the chargeback reversal payment operation typically these are written to a http.Request
*/
type ChargebackReversalPaymentParams struct {

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
	Body *kbmodel.PaymentTransaction
	/*ControlPluginName*/
	ControlPluginName []string
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithTimeout(timeout time.Duration) *ChargebackReversalPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithContext(ctx context.Context) *ChargebackReversalPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithHTTPClient(client *http.Client) *ChargebackReversalPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithXKillbillAPIKey(xKillbillAPIKey string) *ChargebackReversalPaymentParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithXKillbillAPISecret(xKillbillAPISecret string) *ChargebackReversalPaymentParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithXKillbillComment(xKillbillComment *string) *ChargebackReversalPaymentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ChargebackReversalPaymentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithXKillbillReason(xKillbillReason *string) *ChargebackReversalPaymentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithBody(body *kbmodel.PaymentTransaction) *ChargebackReversalPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithControlPluginName(controlPluginName []string) *ChargebackReversalPaymentParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentID adds the paymentID to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithPaymentID(paymentID strfmt.UUID) *ChargebackReversalPaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) WithPluginProperty(pluginProperty []string) *ChargebackReversalPaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the chargeback reversal payment params
func (o *ChargebackReversalPaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *ChargebackReversalPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesControlPluginName := o.ControlPluginName

	joinedControlPluginName := swag.JoinByFormat(valuesControlPluginName, "multi")
	// query array param controlPluginName
	if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
		return err
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
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