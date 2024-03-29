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

// checks if the AddAggregateRequestCriteriaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAggregateRequestCriteriaRequest{}

// AddAggregateRequestCriteriaRequest struct for AddAggregateRequestCriteriaRequest
type AddAggregateRequestCriteriaRequest struct {
	Schemas []EnumaggregateRequestCriteriaSchemaUrn `json:"schemas"`
	// Specifies a request criteria object that must match the associated operation request in order to match the aggregate request criteria. If one or more all-included request criteria objects are provided, then an operation request must match all of them in order to match the aggregate request criteria.
	AllIncludedRequestCriteria []string `json:"allIncludedRequestCriteria,omitempty"`
	// Specifies a request criteria object that may match the associated operation request in order to the this aggregate request criteria. If one or more any-included request criteria objects are provided, then an operation request must match at least one of them in order to match the aggregate request criteria.
	AnyIncludedRequestCriteria []string `json:"anyIncludedRequestCriteria,omitempty"`
	// Specifies a request criteria object that should not match the associated operation request in order to match the aggregate request criteria. If one or more not-all-included request criteria objects are provided, then an operation request must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate request criteria.
	NotAllIncludedRequestCriteria []string `json:"notAllIncludedRequestCriteria,omitempty"`
	// Specifies a request criteria object that must not match the associated operation request in order to match the aggregate request criteria. If one or more none-included request criteria objects are provided, then an operation request must not match any of them in order to match the aggregate request criteria.
	NoneIncludedRequestCriteria []string `json:"noneIncludedRequestCriteria,omitempty"`
	// A description for this Request Criteria
	Description *string `json:"description,omitempty"`
	// Name of the new Request Criteria
	CriteriaName string `json:"criteriaName"`
}

// NewAddAggregateRequestCriteriaRequest instantiates a new AddAggregateRequestCriteriaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAggregateRequestCriteriaRequest(schemas []EnumaggregateRequestCriteriaSchemaUrn, criteriaName string) *AddAggregateRequestCriteriaRequest {
	this := AddAggregateRequestCriteriaRequest{}
	this.Schemas = schemas
	this.CriteriaName = criteriaName
	return &this
}

// NewAddAggregateRequestCriteriaRequestWithDefaults instantiates a new AddAggregateRequestCriteriaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAggregateRequestCriteriaRequestWithDefaults() *AddAggregateRequestCriteriaRequest {
	this := AddAggregateRequestCriteriaRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddAggregateRequestCriteriaRequest) GetSchemas() []EnumaggregateRequestCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumaggregateRequestCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAggregateRequestCriteriaRequest) GetSchemasOk() ([]EnumaggregateRequestCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAggregateRequestCriteriaRequest) SetSchemas(v []EnumaggregateRequestCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetAllIncludedRequestCriteria returns the AllIncludedRequestCriteria field value if set, zero value otherwise.
func (o *AddAggregateRequestCriteriaRequest) GetAllIncludedRequestCriteria() []string {
	if o == nil || IsNil(o.AllIncludedRequestCriteria) {
		var ret []string
		return ret
	}
	return o.AllIncludedRequestCriteria
}

// GetAllIncludedRequestCriteriaOk returns a tuple with the AllIncludedRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregateRequestCriteriaRequest) GetAllIncludedRequestCriteriaOk() ([]string, bool) {
	if o == nil || IsNil(o.AllIncludedRequestCriteria) {
		return nil, false
	}
	return o.AllIncludedRequestCriteria, true
}

// HasAllIncludedRequestCriteria returns a boolean if a field has been set.
func (o *AddAggregateRequestCriteriaRequest) HasAllIncludedRequestCriteria() bool {
	if o != nil && !IsNil(o.AllIncludedRequestCriteria) {
		return true
	}

	return false
}

// SetAllIncludedRequestCriteria gets a reference to the given []string and assigns it to the AllIncludedRequestCriteria field.
func (o *AddAggregateRequestCriteriaRequest) SetAllIncludedRequestCriteria(v []string) {
	o.AllIncludedRequestCriteria = v
}

// GetAnyIncludedRequestCriteria returns the AnyIncludedRequestCriteria field value if set, zero value otherwise.
func (o *AddAggregateRequestCriteriaRequest) GetAnyIncludedRequestCriteria() []string {
	if o == nil || IsNil(o.AnyIncludedRequestCriteria) {
		var ret []string
		return ret
	}
	return o.AnyIncludedRequestCriteria
}

// GetAnyIncludedRequestCriteriaOk returns a tuple with the AnyIncludedRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregateRequestCriteriaRequest) GetAnyIncludedRequestCriteriaOk() ([]string, bool) {
	if o == nil || IsNil(o.AnyIncludedRequestCriteria) {
		return nil, false
	}
	return o.AnyIncludedRequestCriteria, true
}

// HasAnyIncludedRequestCriteria returns a boolean if a field has been set.
func (o *AddAggregateRequestCriteriaRequest) HasAnyIncludedRequestCriteria() bool {
	if o != nil && !IsNil(o.AnyIncludedRequestCriteria) {
		return true
	}

	return false
}

// SetAnyIncludedRequestCriteria gets a reference to the given []string and assigns it to the AnyIncludedRequestCriteria field.
func (o *AddAggregateRequestCriteriaRequest) SetAnyIncludedRequestCriteria(v []string) {
	o.AnyIncludedRequestCriteria = v
}

// GetNotAllIncludedRequestCriteria returns the NotAllIncludedRequestCriteria field value if set, zero value otherwise.
func (o *AddAggregateRequestCriteriaRequest) GetNotAllIncludedRequestCriteria() []string {
	if o == nil || IsNil(o.NotAllIncludedRequestCriteria) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedRequestCriteria
}

// GetNotAllIncludedRequestCriteriaOk returns a tuple with the NotAllIncludedRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregateRequestCriteriaRequest) GetNotAllIncludedRequestCriteriaOk() ([]string, bool) {
	if o == nil || IsNil(o.NotAllIncludedRequestCriteria) {
		return nil, false
	}
	return o.NotAllIncludedRequestCriteria, true
}

// HasNotAllIncludedRequestCriteria returns a boolean if a field has been set.
func (o *AddAggregateRequestCriteriaRequest) HasNotAllIncludedRequestCriteria() bool {
	if o != nil && !IsNil(o.NotAllIncludedRequestCriteria) {
		return true
	}

	return false
}

// SetNotAllIncludedRequestCriteria gets a reference to the given []string and assigns it to the NotAllIncludedRequestCriteria field.
func (o *AddAggregateRequestCriteriaRequest) SetNotAllIncludedRequestCriteria(v []string) {
	o.NotAllIncludedRequestCriteria = v
}

// GetNoneIncludedRequestCriteria returns the NoneIncludedRequestCriteria field value if set, zero value otherwise.
func (o *AddAggregateRequestCriteriaRequest) GetNoneIncludedRequestCriteria() []string {
	if o == nil || IsNil(o.NoneIncludedRequestCriteria) {
		var ret []string
		return ret
	}
	return o.NoneIncludedRequestCriteria
}

// GetNoneIncludedRequestCriteriaOk returns a tuple with the NoneIncludedRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregateRequestCriteriaRequest) GetNoneIncludedRequestCriteriaOk() ([]string, bool) {
	if o == nil || IsNil(o.NoneIncludedRequestCriteria) {
		return nil, false
	}
	return o.NoneIncludedRequestCriteria, true
}

// HasNoneIncludedRequestCriteria returns a boolean if a field has been set.
func (o *AddAggregateRequestCriteriaRequest) HasNoneIncludedRequestCriteria() bool {
	if o != nil && !IsNil(o.NoneIncludedRequestCriteria) {
		return true
	}

	return false
}

// SetNoneIncludedRequestCriteria gets a reference to the given []string and assigns it to the NoneIncludedRequestCriteria field.
func (o *AddAggregateRequestCriteriaRequest) SetNoneIncludedRequestCriteria(v []string) {
	o.NoneIncludedRequestCriteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAggregateRequestCriteriaRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregateRequestCriteriaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAggregateRequestCriteriaRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAggregateRequestCriteriaRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCriteriaName returns the CriteriaName field value
func (o *AddAggregateRequestCriteriaRequest) GetCriteriaName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CriteriaName
}

// GetCriteriaNameOk returns a tuple with the CriteriaName field value
// and a boolean to check if the value has been set.
func (o *AddAggregateRequestCriteriaRequest) GetCriteriaNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CriteriaName, true
}

// SetCriteriaName sets field value
func (o *AddAggregateRequestCriteriaRequest) SetCriteriaName(v string) {
	o.CriteriaName = v
}

func (o AddAggregateRequestCriteriaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAggregateRequestCriteriaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.AllIncludedRequestCriteria) {
		toSerialize["allIncludedRequestCriteria"] = o.AllIncludedRequestCriteria
	}
	if !IsNil(o.AnyIncludedRequestCriteria) {
		toSerialize["anyIncludedRequestCriteria"] = o.AnyIncludedRequestCriteria
	}
	if !IsNil(o.NotAllIncludedRequestCriteria) {
		toSerialize["notAllIncludedRequestCriteria"] = o.NotAllIncludedRequestCriteria
	}
	if !IsNil(o.NoneIncludedRequestCriteria) {
		toSerialize["noneIncludedRequestCriteria"] = o.NoneIncludedRequestCriteria
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["criteriaName"] = o.CriteriaName
	return toSerialize, nil
}

type NullableAddAggregateRequestCriteriaRequest struct {
	value *AddAggregateRequestCriteriaRequest
	isSet bool
}

func (v NullableAddAggregateRequestCriteriaRequest) Get() *AddAggregateRequestCriteriaRequest {
	return v.value
}

func (v *NullableAddAggregateRequestCriteriaRequest) Set(val *AddAggregateRequestCriteriaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAggregateRequestCriteriaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAggregateRequestCriteriaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAggregateRequestCriteriaRequest(val *AddAggregateRequestCriteriaRequest) *NullableAddAggregateRequestCriteriaRequest {
	return &NullableAddAggregateRequestCriteriaRequest{value: val, isSet: true}
}

func (v NullableAddAggregateRequestCriteriaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAggregateRequestCriteriaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
