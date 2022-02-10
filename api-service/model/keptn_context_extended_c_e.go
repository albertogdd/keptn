package model

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KeptnContextExtendedCE keptn context extended c e
//
// swagger:model keptnContextExtendedCE
type KeptnContextExtendedCE struct {

	// contenttype
	Contenttype string `json:"contenttype,omitempty"`

	// data
	// Required: true
	Data interface{} `json:"data"`

	// extensions
	Extensions interface{} `json:"extensions,omitempty"`

	// gitcommitid
	Gitcommitid string `json:"gitcommitid,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// shkeptncontext
	Shkeptncontext string `json:"shkeptncontext,omitempty"`

	// shkeptnversion
	Shkeptnversion string `json:"shkeptnversion,omitempty"`

	// source
	// Required: true
	Source *string `json:"source"`

	// specversion
	Specversion string `json:"specversion,omitempty"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// triggeredid
	Triggeredid string `json:"triggeredid,omitempty"`

	// type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this keptn context extended c e
func (m *KeptnContextExtendedCE) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeptnContextExtendedCE) validateData(formats strfmt.Registry) error {

	if m.Data == nil {
		return errors.Required("data", "body", nil)
	}

	return nil
}

func (m *KeptnContextExtendedCE) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *KeptnContextExtendedCE) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KeptnContextExtendedCE) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this keptn context extended c e based on context it is used
func (m *KeptnContextExtendedCE) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KeptnContextExtendedCE) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeptnContextExtendedCE) UnmarshalBinary(b []byte) error {
	var res KeptnContextExtendedCE
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
