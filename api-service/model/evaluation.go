package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Evaluation evaluation
//
// swagger:model evaluation
type Evaluation struct {

	// Evaluation end timestamp
	End string `json:"end,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// Evaluation start timestamp
	Start string `json:"start,omitempty"`

	// Evaluation timeframe
	Timeframe string `json:"timeframe,omitempty"`
}

// Validate validates this evaluation
func (m *Evaluation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Evaluation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Evaluation) UnmarshalBinary(b []byte) error {
	var res Evaluation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
