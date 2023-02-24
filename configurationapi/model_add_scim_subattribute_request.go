/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddScimSubattributeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddScimSubattributeRequest{}

// AddScimSubattributeRequest struct for AddScimSubattributeRequest
type AddScimSubattributeRequest struct {
	// Name of the new SCIM Subattribute
	SubattributeName string                          `json:"subattributeName"`
	Schemas          []EnumscimSubattributeSchemaUrn `json:"schemas,omitempty"`
	// A description for this SCIM Subattribute
	Description *string                       `json:"description,omitempty"`
	Type        *EnumscimSubattributeTypeProp `json:"type,omitempty"`
	// Specifies whether this sub-attribute is required.
	Required *bool `json:"required,omitempty"`
	// Specifies whether the sub-attribute values are case sensitive.
	CaseExact *bool `json:"caseExact,omitempty"`
	// Specifies whether this attribute may have multiple values.
	MultiValued *bool `json:"multiValued,omitempty"`
	// Specifies the suggested canonical type values for the sub-attribute.
	CanonicalValue []string                            `json:"canonicalValue,omitempty"`
	Mutability     *EnumscimSubattributeMutabilityProp `json:"mutability,omitempty"`
	Returned       *EnumscimSubattributeReturnedProp   `json:"returned,omitempty"`
	// Specifies the SCIM resource types that may be referenced. This property is only applicable for sub-attributes that are of type 'reference'. Valid values are: A SCIM resource type (e.g., 'User' or 'Group'), 'external' - indicating the resource is an external resource (e.g., such as a photo), or 'uri' - indicating that the reference is to a service endpoint or an identifier (such as a schema urn).
	ReferenceType []string `json:"referenceType,omitempty"`
}

// NewAddScimSubattributeRequest instantiates a new AddScimSubattributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddScimSubattributeRequest(subattributeName string) *AddScimSubattributeRequest {
	this := AddScimSubattributeRequest{}
	this.SubattributeName = subattributeName
	return &this
}

// NewAddScimSubattributeRequestWithDefaults instantiates a new AddScimSubattributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddScimSubattributeRequestWithDefaults() *AddScimSubattributeRequest {
	this := AddScimSubattributeRequest{}
	return &this
}

// GetSubattributeName returns the SubattributeName field value
func (o *AddScimSubattributeRequest) GetSubattributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubattributeName
}

// GetSubattributeNameOk returns a tuple with the SubattributeName field value
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetSubattributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubattributeName, true
}

// SetSubattributeName sets field value
func (o *AddScimSubattributeRequest) SetSubattributeName(v string) {
	o.SubattributeName = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetSchemas() []EnumscimSubattributeSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumscimSubattributeSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetSchemasOk() ([]EnumscimSubattributeSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumscimSubattributeSchemaUrn and assigns it to the Schemas field.
func (o *AddScimSubattributeRequest) SetSchemas(v []EnumscimSubattributeSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddScimSubattributeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetType() EnumscimSubattributeTypeProp {
	if o == nil || IsNil(o.Type) {
		var ret EnumscimSubattributeTypeProp
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetTypeOk() (*EnumscimSubattributeTypeProp, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumscimSubattributeTypeProp and assigns it to the Type field.
func (o *AddScimSubattributeRequest) SetType(v EnumscimSubattributeTypeProp) {
	o.Type = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *AddScimSubattributeRequest) SetRequired(v bool) {
	o.Required = &v
}

// GetCaseExact returns the CaseExact field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetCaseExact() bool {
	if o == nil || IsNil(o.CaseExact) {
		var ret bool
		return ret
	}
	return *o.CaseExact
}

// GetCaseExactOk returns a tuple with the CaseExact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetCaseExactOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseExact) {
		return nil, false
	}
	return o.CaseExact, true
}

// HasCaseExact returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasCaseExact() bool {
	if o != nil && !IsNil(o.CaseExact) {
		return true
	}

	return false
}

// SetCaseExact gets a reference to the given bool and assigns it to the CaseExact field.
func (o *AddScimSubattributeRequest) SetCaseExact(v bool) {
	o.CaseExact = &v
}

// GetMultiValued returns the MultiValued field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetMultiValued() bool {
	if o == nil || IsNil(o.MultiValued) {
		var ret bool
		return ret
	}
	return *o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetMultiValuedOk() (*bool, bool) {
	if o == nil || IsNil(o.MultiValued) {
		return nil, false
	}
	return o.MultiValued, true
}

// HasMultiValued returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasMultiValued() bool {
	if o != nil && !IsNil(o.MultiValued) {
		return true
	}

	return false
}

// SetMultiValued gets a reference to the given bool and assigns it to the MultiValued field.
func (o *AddScimSubattributeRequest) SetMultiValued(v bool) {
	o.MultiValued = &v
}

// GetCanonicalValue returns the CanonicalValue field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetCanonicalValue() []string {
	if o == nil || IsNil(o.CanonicalValue) {
		var ret []string
		return ret
	}
	return o.CanonicalValue
}

// GetCanonicalValueOk returns a tuple with the CanonicalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetCanonicalValueOk() ([]string, bool) {
	if o == nil || IsNil(o.CanonicalValue) {
		return nil, false
	}
	return o.CanonicalValue, true
}

// HasCanonicalValue returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasCanonicalValue() bool {
	if o != nil && !IsNil(o.CanonicalValue) {
		return true
	}

	return false
}

// SetCanonicalValue gets a reference to the given []string and assigns it to the CanonicalValue field.
func (o *AddScimSubattributeRequest) SetCanonicalValue(v []string) {
	o.CanonicalValue = v
}

// GetMutability returns the Mutability field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetMutability() EnumscimSubattributeMutabilityProp {
	if o == nil || IsNil(o.Mutability) {
		var ret EnumscimSubattributeMutabilityProp
		return ret
	}
	return *o.Mutability
}

// GetMutabilityOk returns a tuple with the Mutability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetMutabilityOk() (*EnumscimSubattributeMutabilityProp, bool) {
	if o == nil || IsNil(o.Mutability) {
		return nil, false
	}
	return o.Mutability, true
}

// HasMutability returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasMutability() bool {
	if o != nil && !IsNil(o.Mutability) {
		return true
	}

	return false
}

// SetMutability gets a reference to the given EnumscimSubattributeMutabilityProp and assigns it to the Mutability field.
func (o *AddScimSubattributeRequest) SetMutability(v EnumscimSubattributeMutabilityProp) {
	o.Mutability = &v
}

// GetReturned returns the Returned field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetReturned() EnumscimSubattributeReturnedProp {
	if o == nil || IsNil(o.Returned) {
		var ret EnumscimSubattributeReturnedProp
		return ret
	}
	return *o.Returned
}

// GetReturnedOk returns a tuple with the Returned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetReturnedOk() (*EnumscimSubattributeReturnedProp, bool) {
	if o == nil || IsNil(o.Returned) {
		return nil, false
	}
	return o.Returned, true
}

// HasReturned returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasReturned() bool {
	if o != nil && !IsNil(o.Returned) {
		return true
	}

	return false
}

// SetReturned gets a reference to the given EnumscimSubattributeReturnedProp and assigns it to the Returned field.
func (o *AddScimSubattributeRequest) SetReturned(v EnumscimSubattributeReturnedProp) {
	o.Returned = &v
}

// GetReferenceType returns the ReferenceType field value if set, zero value otherwise.
func (o *AddScimSubattributeRequest) GetReferenceType() []string {
	if o == nil || IsNil(o.ReferenceType) {
		var ret []string
		return ret
	}
	return o.ReferenceType
}

// GetReferenceTypeOk returns a tuple with the ReferenceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddScimSubattributeRequest) GetReferenceTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.ReferenceType) {
		return nil, false
	}
	return o.ReferenceType, true
}

// HasReferenceType returns a boolean if a field has been set.
func (o *AddScimSubattributeRequest) HasReferenceType() bool {
	if o != nil && !IsNil(o.ReferenceType) {
		return true
	}

	return false
}

// SetReferenceType gets a reference to the given []string and assigns it to the ReferenceType field.
func (o *AddScimSubattributeRequest) SetReferenceType(v []string) {
	o.ReferenceType = v
}

func (o AddScimSubattributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddScimSubattributeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subattributeName"] = o.SubattributeName
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.CaseExact) {
		toSerialize["caseExact"] = o.CaseExact
	}
	if !IsNil(o.MultiValued) {
		toSerialize["multiValued"] = o.MultiValued
	}
	if !IsNil(o.CanonicalValue) {
		toSerialize["canonicalValue"] = o.CanonicalValue
	}
	if !IsNil(o.Mutability) {
		toSerialize["mutability"] = o.Mutability
	}
	if !IsNil(o.Returned) {
		toSerialize["returned"] = o.Returned
	}
	if !IsNil(o.ReferenceType) {
		toSerialize["referenceType"] = o.ReferenceType
	}
	return toSerialize, nil
}

type NullableAddScimSubattributeRequest struct {
	value *AddScimSubattributeRequest
	isSet bool
}

func (v NullableAddScimSubattributeRequest) Get() *AddScimSubattributeRequest {
	return v.value
}

func (v *NullableAddScimSubattributeRequest) Set(val *AddScimSubattributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddScimSubattributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddScimSubattributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddScimSubattributeRequest(val *AddScimSubattributeRequest) *NullableAddScimSubattributeRequest {
	return &NullableAddScimSubattributeRequest{value: val, isSet: true}
}

func (v NullableAddScimSubattributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddScimSubattributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
