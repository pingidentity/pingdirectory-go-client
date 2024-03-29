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

// checks if the AddSimpleSearchReferenceCriteriaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSimpleSearchReferenceCriteriaRequest{}

// AddSimpleSearchReferenceCriteriaRequest struct for AddSimpleSearchReferenceCriteriaRequest
type AddSimpleSearchReferenceCriteriaRequest struct {
	Schemas []EnumsimpleSearchReferenceCriteriaSchemaUrn `json:"schemas"`
	// Specifies a request criteria object that must match the associated request for references included in this Simple Search Reference Criteria.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Specifies the OID of a control that must be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain all of those controls.
	AllIncludedReferenceControl []string `json:"allIncludedReferenceControl,omitempty"`
	// Specifies the OID of a control that may be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain at least one of those controls.
	AnyIncludedReferenceControl []string `json:"anyIncludedReferenceControl,omitempty"`
	// Specifies the OID of a control that should not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain at least one of those controls (that is, it may contain zero or more of those controls, but not all of them).
	NotAllIncludedReferenceControl []string `json:"notAllIncludedReferenceControl,omitempty"`
	// Specifies the OID of a control that must not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain any of those controls.
	NoneIncludedReferenceControl []string `json:"noneIncludedReferenceControl,omitempty"`
	// A description for this Search Reference Criteria
	Description *string `json:"description,omitempty"`
	// Name of the new Search Reference Criteria
	CriteriaName string `json:"criteriaName"`
}

// NewAddSimpleSearchReferenceCriteriaRequest instantiates a new AddSimpleSearchReferenceCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSimpleSearchReferenceCriteriaRequest(schemas []EnumsimpleSearchReferenceCriteriaSchemaUrn, criteriaName string) *AddSimpleSearchReferenceCriteriaRequest {
	this := AddSimpleSearchReferenceCriteriaRequest{}
	this.Schemas = schemas
	this.CriteriaName = criteriaName
	return &this
}

// NewAddSimpleSearchReferenceCriteriaRequestWithDefaults instantiates a new AddSimpleSearchReferenceCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSimpleSearchReferenceCriteriaRequestWithDefaults() *AddSimpleSearchReferenceCriteriaRequest {
	this := AddSimpleSearchReferenceCriteriaRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddSimpleSearchReferenceCriteriaRequest) GetSchemas() []EnumsimpleSearchReferenceCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumsimpleSearchReferenceCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetSchemasOk() ([]EnumsimpleSearchReferenceCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddSimpleSearchReferenceCriteriaRequest) SetSchemas(v []EnumsimpleSearchReferenceCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddSimpleSearchReferenceCriteriaRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetAllIncludedReferenceControl returns the AllIncludedReferenceControl field value if set, zero value otherwise.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetAllIncludedReferenceControl() []string {
	if o == nil || IsNil(o.AllIncludedReferenceControl) {
		var ret []string
		return ret
	}
	return o.AllIncludedReferenceControl
}

// GetAllIncludedReferenceControlOk returns a tuple with the AllIncludedReferenceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetAllIncludedReferenceControlOk() ([]string, bool) {
	if o == nil || IsNil(o.AllIncludedReferenceControl) {
		return nil, false
	}
	return o.AllIncludedReferenceControl, true
}

// HasAllIncludedReferenceControl returns a boolean if a field has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) HasAllIncludedReferenceControl() bool {
	if o != nil && !IsNil(o.AllIncludedReferenceControl) {
		return true
	}

	return false
}

// SetAllIncludedReferenceControl gets a reference to the given []string and assigns it to the AllIncludedReferenceControl field.
func (o *AddSimpleSearchReferenceCriteriaRequest) SetAllIncludedReferenceControl(v []string) {
	o.AllIncludedReferenceControl = v
}

// GetAnyIncludedReferenceControl returns the AnyIncludedReferenceControl field value if set, zero value otherwise.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetAnyIncludedReferenceControl() []string {
	if o == nil || IsNil(o.AnyIncludedReferenceControl) {
		var ret []string
		return ret
	}
	return o.AnyIncludedReferenceControl
}

// GetAnyIncludedReferenceControlOk returns a tuple with the AnyIncludedReferenceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetAnyIncludedReferenceControlOk() ([]string, bool) {
	if o == nil || IsNil(o.AnyIncludedReferenceControl) {
		return nil, false
	}
	return o.AnyIncludedReferenceControl, true
}

// HasAnyIncludedReferenceControl returns a boolean if a field has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) HasAnyIncludedReferenceControl() bool {
	if o != nil && !IsNil(o.AnyIncludedReferenceControl) {
		return true
	}

	return false
}

// SetAnyIncludedReferenceControl gets a reference to the given []string and assigns it to the AnyIncludedReferenceControl field.
func (o *AddSimpleSearchReferenceCriteriaRequest) SetAnyIncludedReferenceControl(v []string) {
	o.AnyIncludedReferenceControl = v
}

// GetNotAllIncludedReferenceControl returns the NotAllIncludedReferenceControl field value if set, zero value otherwise.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetNotAllIncludedReferenceControl() []string {
	if o == nil || IsNil(o.NotAllIncludedReferenceControl) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedReferenceControl
}

// GetNotAllIncludedReferenceControlOk returns a tuple with the NotAllIncludedReferenceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetNotAllIncludedReferenceControlOk() ([]string, bool) {
	if o == nil || IsNil(o.NotAllIncludedReferenceControl) {
		return nil, false
	}
	return o.NotAllIncludedReferenceControl, true
}

// HasNotAllIncludedReferenceControl returns a boolean if a field has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) HasNotAllIncludedReferenceControl() bool {
	if o != nil && !IsNil(o.NotAllIncludedReferenceControl) {
		return true
	}

	return false
}

// SetNotAllIncludedReferenceControl gets a reference to the given []string and assigns it to the NotAllIncludedReferenceControl field.
func (o *AddSimpleSearchReferenceCriteriaRequest) SetNotAllIncludedReferenceControl(v []string) {
	o.NotAllIncludedReferenceControl = v
}

// GetNoneIncludedReferenceControl returns the NoneIncludedReferenceControl field value if set, zero value otherwise.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetNoneIncludedReferenceControl() []string {
	if o == nil || IsNil(o.NoneIncludedReferenceControl) {
		var ret []string
		return ret
	}
	return o.NoneIncludedReferenceControl
}

// GetNoneIncludedReferenceControlOk returns a tuple with the NoneIncludedReferenceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetNoneIncludedReferenceControlOk() ([]string, bool) {
	if o == nil || IsNil(o.NoneIncludedReferenceControl) {
		return nil, false
	}
	return o.NoneIncludedReferenceControl, true
}

// HasNoneIncludedReferenceControl returns a boolean if a field has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) HasNoneIncludedReferenceControl() bool {
	if o != nil && !IsNil(o.NoneIncludedReferenceControl) {
		return true
	}

	return false
}

// SetNoneIncludedReferenceControl gets a reference to the given []string and assigns it to the NoneIncludedReferenceControl field.
func (o *AddSimpleSearchReferenceCriteriaRequest) SetNoneIncludedReferenceControl(v []string) {
	o.NoneIncludedReferenceControl = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddSimpleSearchReferenceCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddSimpleSearchReferenceCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddSimpleSearchReferenceCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddSimpleSearchReferenceCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

func (o AddSimpleSearchReferenceCriteriaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSimpleSearchReferenceCriteriaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.AllIncludedReferenceControl) {
		toSerialize["allIncludedReferenceControl"] = o.AllIncludedReferenceControl
	}
	if !IsNil(o.AnyIncludedReferenceControl) {
		toSerialize["anyIncludedReferenceControl"] = o.AnyIncludedReferenceControl
	}
	if !IsNil(o.NotAllIncludedReferenceControl) {
		toSerialize["notAllIncludedReferenceControl"] = o.NotAllIncludedReferenceControl
	}
	if !IsNil(o.NoneIncludedReferenceControl) {
		toSerialize["noneIncludedReferenceControl"] = o.NoneIncludedReferenceControl
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["criteriaName"] = o.CriteriaName
	return toSerialize, nil
}

type NullableAddSimpleSearchReferenceCriteriaRequest struct {
	value *AddSimpleSearchReferenceCriteriaRequest
	isSet bool
}

func (v NullableAddSimpleSearchReferenceCriteriaRequest) Get() *AddSimpleSearchReferenceCriteriaRequest {
	return v.value
}

func (v *NullableAddSimpleSearchReferenceCriteriaRequest) Set(val *AddSimpleSearchReferenceCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSimpleSearchReferenceCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSimpleSearchReferenceCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSimpleSearchReferenceCriteriaRequest(val *AddSimpleSearchReferenceCriteriaRequest) *NullableAddSimpleSearchReferenceCriteriaRequest {
	return &NullableAddSimpleSearchReferenceCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddSimpleSearchReferenceCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSimpleSearchReferenceCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
