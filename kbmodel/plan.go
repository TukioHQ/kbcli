// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Plan plan
// swagger:model Plan
type Plan struct {

	// billing period
	BillingPeriod PlanBillingPeriodEnum `json:"billingPeriod,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// phases
	Phases []*Phase `json:"phases"`

	// pretty name
	PrettyName string `json:"prettyName,omitempty"`
}

// Validate validates this plan
func (m *Plan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPeriod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhases(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var planTypeBillingPeriodPropEnum []interface{}

func init() {
	var res []PlanBillingPeriodEnum
	if err := json.Unmarshal([]byte(`["DAILY","WEEKLY","BIWEEKLY","THIRTY_DAYS","SIXTY_DAYS","NINETY_DAYS","MONTHLY","BIMESTRIAL","QUARTERLY","TRIANNUAL","BIANNUAL","ANNUAL","BIENNIAL","NO_BILLING_PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		planTypeBillingPeriodPropEnum = append(planTypeBillingPeriodPropEnum, v)
	}
}

type PlanBillingPeriodEnum string

const (

	// PlanBillingPeriodDAILY captures enum value "DAILY"
	PlanBillingPeriodDAILY PlanBillingPeriodEnum = "DAILY"

	// PlanBillingPeriodWEEKLY captures enum value "WEEKLY"
	PlanBillingPeriodWEEKLY PlanBillingPeriodEnum = "WEEKLY"

	// PlanBillingPeriodBIWEEKLY captures enum value "BIWEEKLY"
	PlanBillingPeriodBIWEEKLY PlanBillingPeriodEnum = "BIWEEKLY"

	// PlanBillingPeriodTHIRTYDAYS captures enum value "THIRTY_DAYS"
	PlanBillingPeriodTHIRTYDAYS PlanBillingPeriodEnum = "THIRTY_DAYS"

	// PlanBillingPeriodSIXTYDAYS captures enum value "SIXTY_DAYS"
	PlanBillingPeriodSIXTYDAYS PlanBillingPeriodEnum = "SIXTY_DAYS"

	// PlanBillingPeriodNINETYDAYS captures enum value "NINETY_DAYS"
	PlanBillingPeriodNINETYDAYS PlanBillingPeriodEnum = "NINETY_DAYS"

	// PlanBillingPeriodMONTHLY captures enum value "MONTHLY"
	PlanBillingPeriodMONTHLY PlanBillingPeriodEnum = "MONTHLY"

	// PlanBillingPeriodBIMESTRIAL captures enum value "BIMESTRIAL"
	PlanBillingPeriodBIMESTRIAL PlanBillingPeriodEnum = "BIMESTRIAL"

	// PlanBillingPeriodQUARTERLY captures enum value "QUARTERLY"
	PlanBillingPeriodQUARTERLY PlanBillingPeriodEnum = "QUARTERLY"

	// PlanBillingPeriodTRIANNUAL captures enum value "TRIANNUAL"
	PlanBillingPeriodTRIANNUAL PlanBillingPeriodEnum = "TRIANNUAL"

	// PlanBillingPeriodBIANNUAL captures enum value "BIANNUAL"
	PlanBillingPeriodBIANNUAL PlanBillingPeriodEnum = "BIANNUAL"

	// PlanBillingPeriodANNUAL captures enum value "ANNUAL"
	PlanBillingPeriodANNUAL PlanBillingPeriodEnum = "ANNUAL"

	// PlanBillingPeriodBIENNIAL captures enum value "BIENNIAL"
	PlanBillingPeriodBIENNIAL PlanBillingPeriodEnum = "BIENNIAL"

	// PlanBillingPeriodNOBILLINGPERIOD captures enum value "NO_BILLING_PERIOD"
	PlanBillingPeriodNOBILLINGPERIOD PlanBillingPeriodEnum = "NO_BILLING_PERIOD"
)

var PlanBillingPeriodEnumValues = []string{
	"DAILY",
	"WEEKLY",
	"BIWEEKLY",
	"THIRTY_DAYS",
	"SIXTY_DAYS",
	"NINETY_DAYS",
	"MONTHLY",
	"BIMESTRIAL",
	"QUARTERLY",
	"TRIANNUAL",
	"BIANNUAL",
	"ANNUAL",
	"BIENNIAL",
	"NO_BILLING_PERIOD",
}

func (e PlanBillingPeriodEnum) IsValid() bool {
	for _, v := range PlanBillingPeriodEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *Plan) validateBillingPeriodEnum(path, location string, value PlanBillingPeriodEnum) error {
	if err := validate.Enum(path, location, value, planTypeBillingPeriodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Plan) validateBillingPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingPeriod) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingPeriodEnum("billingPeriod", "body", m.BillingPeriod); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validatePhases(formats strfmt.Registry) error {

	if swag.IsZero(m.Phases) { // not required
		return nil
	}

	for i := 0; i < len(m.Phases); i++ {

		if swag.IsZero(m.Phases[i]) { // not required
			continue
		}

		if m.Phases[i] != nil {

			if err := m.Phases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phases" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plan) UnmarshalBinary(b []byte) error {
	var res Plan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
