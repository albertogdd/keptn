package models

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExpandedStage expanded stage
//
// swagger:model ExpandedStage
type ExpandedStage struct {

	// last event context
	LastEventContext *EventContext `json:"lastEventContext,omitempty"`

	// services
	Services []*ExpandedService `json:"services"`

	// Stage name
	StageName string `json:"stageName,omitempty"`
}

// Validate validates this expanded stage
func (m *ExpandedStage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastEventContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExpandedStage) validateLastEventContext(formats strfmt.Registry) error {

	if swag.IsZero(m.LastEventContext) { // not required
		return nil
	}

	if m.LastEventContext != nil {
		if err := m.LastEventContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastEventContext")
			}
			return err
		}
	}

	return nil
}

func (m *ExpandedStage) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExpandedStage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpandedStage) UnmarshalBinary(b []byte) error {
	var res ExpandedStage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
