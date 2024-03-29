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

// checks if the AddAggregatePassThroughAuthenticationHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddAggregatePassThroughAuthenticationHandlerRequest{}

// AddAggregatePassThroughAuthenticationHandlerRequest struct for AddAggregatePassThroughAuthenticationHandlerRequest
type AddAggregatePassThroughAuthenticationHandlerRequest struct {
	Schemas []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn `json:"schemas"`
	// The set of subordinate pass-through authentication handlers that may be used to perform the authentication processing. Handlers will be invoked in order until one is found for which the bind operation matches the associated criteria and either succeeds or fails in a manner that should not be ignored.
	SubordinatePassThroughAuthenticationHandler []string                                                        `json:"subordinatePassThroughAuthenticationHandler"`
	ContinueOnFailureType                       []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp `json:"continueOnFailureType,omitempty"`
	// A description for this Pass Through Authentication Handler
	Description *string `json:"description,omitempty"`
	// The base DNs for the local users whose authentication attempts may be passed through to the external authentication service.
	IncludedLocalEntryBaseDN []string `json:"includedLocalEntryBaseDN,omitempty"`
	// A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Name of the new Pass Through Authentication Handler
	HandlerName string `json:"handlerName"`
}

// NewAddAggregatePassThroughAuthenticationHandlerRequest instantiates a new AddAggregatePassThroughAuthenticationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddAggregatePassThroughAuthenticationHandlerRequest(schemas []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn, subordinatePassThroughAuthenticationHandler []string, handlerName string) *AddAggregatePassThroughAuthenticationHandlerRequest {
	this := AddAggregatePassThroughAuthenticationHandlerRequest{}
	this.Schemas = schemas
	this.SubordinatePassThroughAuthenticationHandler = subordinatePassThroughAuthenticationHandler
	this.HandlerName = handlerName
	return &this
}

// NewAddAggregatePassThroughAuthenticationHandlerRequestWithDefaults instantiates a new AddAggregatePassThroughAuthenticationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddAggregatePassThroughAuthenticationHandlerRequestWithDefaults() *AddAggregatePassThroughAuthenticationHandlerRequest {
	this := AddAggregatePassThroughAuthenticationHandlerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetSchemas() []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetSchemasOk() ([]EnumaggregatePassThroughAuthenticationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetSchemas(v []EnumaggregatePassThroughAuthenticationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetSubordinatePassThroughAuthenticationHandler returns the SubordinatePassThroughAuthenticationHandler field value
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetSubordinatePassThroughAuthenticationHandler() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubordinatePassThroughAuthenticationHandler
}

// GetSubordinatePassThroughAuthenticationHandlerOk returns a tuple with the SubordinatePassThroughAuthenticationHandler field value
// and a boolean to check if the value has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetSubordinatePassThroughAuthenticationHandlerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubordinatePassThroughAuthenticationHandler, true
}

// SetSubordinatePassThroughAuthenticationHandler sets field value
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetSubordinatePassThroughAuthenticationHandler(v []string) {
	o.SubordinatePassThroughAuthenticationHandler = v
}

// GetContinueOnFailureType returns the ContinueOnFailureType field value if set, zero value otherwise.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetContinueOnFailureType() []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp {
	if o == nil || IsNil(o.ContinueOnFailureType) {
		var ret []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp
		return ret
	}
	return o.ContinueOnFailureType
}

// GetContinueOnFailureTypeOk returns a tuple with the ContinueOnFailureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetContinueOnFailureTypeOk() ([]EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp, bool) {
	if o == nil || IsNil(o.ContinueOnFailureType) {
		return nil, false
	}
	return o.ContinueOnFailureType, true
}

// HasContinueOnFailureType returns a boolean if a field has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasContinueOnFailureType() bool {
	if o != nil && !IsNil(o.ContinueOnFailureType) {
		return true
	}

	return false
}

// SetContinueOnFailureType gets a reference to the given []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp and assigns it to the ContinueOnFailureType field.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetContinueOnFailureType(v []EnumpassThroughAuthenticationHandlerContinueOnFailureTypeProp) {
	o.ContinueOnFailureType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field value if set, zero value otherwise.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDN() []string {
	if o == nil || IsNil(o.IncludedLocalEntryBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludedLocalEntryBaseDN
}

// GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetIncludedLocalEntryBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedLocalEntryBaseDN) {
		return nil, false
	}
	return o.IncludedLocalEntryBaseDN, true
}

// HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasIncludedLocalEntryBaseDN() bool {
	if o != nil && !IsNil(o.IncludedLocalEntryBaseDN) {
		return true
	}

	return false
}

// SetIncludedLocalEntryBaseDN gets a reference to the given []string and assigns it to the IncludedLocalEntryBaseDN field.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetIncludedLocalEntryBaseDN(v []string) {
	o.IncludedLocalEntryBaseDN = v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetConnectionCriteria() string {
	if o == nil || IsNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasConnectionCriteria() bool {
	if o != nil && !IsNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetHandlerName returns the HandlerName field value
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddAggregatePassThroughAuthenticationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddAggregatePassThroughAuthenticationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddAggregatePassThroughAuthenticationHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["subordinatePassThroughAuthenticationHandler"] = o.SubordinatePassThroughAuthenticationHandler
	if !IsNil(o.ContinueOnFailureType) {
		toSerialize["continueOnFailureType"] = o.ContinueOnFailureType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IncludedLocalEntryBaseDN) {
		toSerialize["includedLocalEntryBaseDN"] = o.IncludedLocalEntryBaseDN
	}
	if !IsNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	toSerialize["handlerName"] = o.HandlerName
	return toSerialize, nil
}

type NullableAddAggregatePassThroughAuthenticationHandlerRequest struct {
	value *AddAggregatePassThroughAuthenticationHandlerRequest
	isSet bool
}

func (v NullableAddAggregatePassThroughAuthenticationHandlerRequest) Get() *AddAggregatePassThroughAuthenticationHandlerRequest {
	return v.value
}

func (v *NullableAddAggregatePassThroughAuthenticationHandlerRequest) Set(val *AddAggregatePassThroughAuthenticationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddAggregatePassThroughAuthenticationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddAggregatePassThroughAuthenticationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddAggregatePassThroughAuthenticationHandlerRequest(val *AddAggregatePassThroughAuthenticationHandlerRequest) *NullableAddAggregatePassThroughAuthenticationHandlerRequest {
	return &NullableAddAggregatePassThroughAuthenticationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddAggregatePassThroughAuthenticationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddAggregatePassThroughAuthenticationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
