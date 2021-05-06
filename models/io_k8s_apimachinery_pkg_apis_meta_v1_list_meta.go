// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IoK8sApimachineryPkgApisMetaV1ListMeta ListMeta describes metadata that synthetic resources must have, including lists and various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
//
// swagger:model io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta
type IoK8sApimachineryPkgApisMetaV1ListMeta struct {

	// continue may be set if the user set a limit on the number of items returned, and indicates that the server has more data available. The value is opaque and may be used to issue another request to the endpoint that served this list to retrieve the next set of available objects. Continuing a consistent list may not be possible if the server configuration has changed or more than a few minutes have passed. The resourceVersion field returned when using this continue value will be identical to the value in the first response, unless you have received this token from an error message.
	Continue string `json:"continue,omitempty"`

	// remainingItemCount is the number of subsequent items in the list which are not included in this list response. If the list request contained label or field selectors, then the number of remaining items is unknown and the field will be left unset and omitted during serialization. If the list is complete (either because it is not chunking or because this is the last chunk), then there are no more remaining items and this field will be left unset and omitted during serialization. Servers older than v1.15 do not set this field. The intended use of the remainingItemCount is *estimating* the size of a collection. Clients should not rely on the remainingItemCount to be set or to be exact.
	//
	// This field is alpha and can be changed or removed without notice.
	RemainingItemCount int64 `json:"remainingItemCount,omitempty"`

	// String that identifies the server's internal version of this object that can be used by clients to determine when objects have changed. Value must be treated as opaque by clients and passed unmodified back to the server. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// selfLink is a URL representing this object. Populated by the system. Read-only.
	SelfLink string `json:"selfLink,omitempty"`
}

// Validate validates this io k8s apimachinery pkg apis meta v1 list meta
func (m *IoK8sApimachineryPkgApisMetaV1ListMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this io k8s apimachinery pkg apis meta v1 list meta based on context it is used
func (m *IoK8sApimachineryPkgApisMetaV1ListMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1ListMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IoK8sApimachineryPkgApisMetaV1ListMeta) UnmarshalBinary(b []byte) error {
	var res IoK8sApimachineryPkgApisMetaV1ListMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
