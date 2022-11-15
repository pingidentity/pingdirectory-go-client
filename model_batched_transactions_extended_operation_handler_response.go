/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BatchedTransactionsExtendedOperationHandlerResponse struct for BatchedTransactionsExtendedOperationHandlerResponse
type BatchedTransactionsExtendedOperationHandlerResponse struct {
	Schemas []EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn `json:"schemas"`
	// A description for this Extended Operation Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server).
	Enabled bool `json:"enabled"`
}

// NewBatchedTransactionsExtendedOperationHandlerResponse instantiates a new BatchedTransactionsExtendedOperationHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchedTransactionsExtendedOperationHandlerResponse(schemas []EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn, enabled bool) *BatchedTransactionsExtendedOperationHandlerResponse {
	this := BatchedTransactionsExtendedOperationHandlerResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewBatchedTransactionsExtendedOperationHandlerResponseWithDefaults instantiates a new BatchedTransactionsExtendedOperationHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchedTransactionsExtendedOperationHandlerResponseWithDefaults() *BatchedTransactionsExtendedOperationHandlerResponse {
	this := BatchedTransactionsExtendedOperationHandlerResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetSchemas() []EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetSchemasOk() ([]EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *BatchedTransactionsExtendedOperationHandlerResponse) SetSchemas(v []EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BatchedTransactionsExtendedOperationHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BatchedTransactionsExtendedOperationHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *BatchedTransactionsExtendedOperationHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o BatchedTransactionsExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableBatchedTransactionsExtendedOperationHandlerResponse struct {
	value *BatchedTransactionsExtendedOperationHandlerResponse
	isSet bool
}

func (v NullableBatchedTransactionsExtendedOperationHandlerResponse) Get() *BatchedTransactionsExtendedOperationHandlerResponse {
	return v.value
}

func (v *NullableBatchedTransactionsExtendedOperationHandlerResponse) Set(val *BatchedTransactionsExtendedOperationHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchedTransactionsExtendedOperationHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchedTransactionsExtendedOperationHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchedTransactionsExtendedOperationHandlerResponse(val *BatchedTransactionsExtendedOperationHandlerResponse) *NullableBatchedTransactionsExtendedOperationHandlerResponse {
	return &NullableBatchedTransactionsExtendedOperationHandlerResponse{value: val, isSet: true}
}

func (v NullableBatchedTransactionsExtendedOperationHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchedTransactionsExtendedOperationHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


