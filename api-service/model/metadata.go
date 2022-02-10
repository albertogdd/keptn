package model

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Metadata metadata
//
// swagger:model metadata
type Metadata struct {

	// bridgeversion
	Bridgeversion string `json:"bridgeversion,omitempty"`

	// keptnlabel
	Keptnlabel string `json:"keptnlabel,omitempty"`

	// keptnservices
	Keptnservices interface{} `json:"keptnservices,omitempty"`

	// keptnversion
	Keptnversion string `json:"keptnversion,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// shipyardversion
	Shipyardversion string `json:"shipyardversion,omitempty"`
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metadata based on context it is used
func (m *Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metadata) UnmarshalBinary(b []byte) error {
	var res Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
