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

// checks if the TraditionalWorkQueueResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TraditionalWorkQueueResponse{}

// TraditionalWorkQueueResponse struct for TraditionalWorkQueueResponse
type TraditionalWorkQueueResponse struct {
	Schemas []EnumtraditionalWorkQueueSchemaUrn `json:"schemas"`
	// Specifies the number of worker threads to be used for processing operations placed in the queue.
	NumWorkerThreads int64 `json:"numWorkerThreads"`
	// Specifies the maximum number of queued operations that can be in the work queue at any given time.
	MaxWorkQueueCapacity                          *int64                                             `json:"maxWorkQueueCapacity,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewTraditionalWorkQueueResponse instantiates a new TraditionalWorkQueueResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraditionalWorkQueueResponse(schemas []EnumtraditionalWorkQueueSchemaUrn, numWorkerThreads int64) *TraditionalWorkQueueResponse {
	this := TraditionalWorkQueueResponse{}
	this.Schemas = schemas
	this.NumWorkerThreads = numWorkerThreads
	return &this
}

// NewTraditionalWorkQueueResponseWithDefaults instantiates a new TraditionalWorkQueueResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraditionalWorkQueueResponseWithDefaults() *TraditionalWorkQueueResponse {
	this := TraditionalWorkQueueResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *TraditionalWorkQueueResponse) GetSchemas() []EnumtraditionalWorkQueueSchemaUrn {
	if o == nil {
		var ret []EnumtraditionalWorkQueueSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *TraditionalWorkQueueResponse) GetSchemasOk() ([]EnumtraditionalWorkQueueSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *TraditionalWorkQueueResponse) SetSchemas(v []EnumtraditionalWorkQueueSchemaUrn) {
	o.Schemas = v
}

// GetNumWorkerThreads returns the NumWorkerThreads field value
func (o *TraditionalWorkQueueResponse) GetNumWorkerThreads() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NumWorkerThreads
}

// GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field value
// and a boolean to check if the value has been set.
func (o *TraditionalWorkQueueResponse) GetNumWorkerThreadsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumWorkerThreads, true
}

// SetNumWorkerThreads sets field value
func (o *TraditionalWorkQueueResponse) SetNumWorkerThreads(v int64) {
	o.NumWorkerThreads = v
}

// GetMaxWorkQueueCapacity returns the MaxWorkQueueCapacity field value if set, zero value otherwise.
func (o *TraditionalWorkQueueResponse) GetMaxWorkQueueCapacity() int64 {
	if o == nil || IsNil(o.MaxWorkQueueCapacity) {
		var ret int64
		return ret
	}
	return *o.MaxWorkQueueCapacity
}

// GetMaxWorkQueueCapacityOk returns a tuple with the MaxWorkQueueCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraditionalWorkQueueResponse) GetMaxWorkQueueCapacityOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxWorkQueueCapacity) {
		return nil, false
	}
	return o.MaxWorkQueueCapacity, true
}

// HasMaxWorkQueueCapacity returns a boolean if a field has been set.
func (o *TraditionalWorkQueueResponse) HasMaxWorkQueueCapacity() bool {
	if o != nil && !IsNil(o.MaxWorkQueueCapacity) {
		return true
	}

	return false
}

// SetMaxWorkQueueCapacity gets a reference to the given int64 and assigns it to the MaxWorkQueueCapacity field.
func (o *TraditionalWorkQueueResponse) SetMaxWorkQueueCapacity(v int64) {
	o.MaxWorkQueueCapacity = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *TraditionalWorkQueueResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraditionalWorkQueueResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *TraditionalWorkQueueResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *TraditionalWorkQueueResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *TraditionalWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraditionalWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *TraditionalWorkQueueResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *TraditionalWorkQueueResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o TraditionalWorkQueueResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TraditionalWorkQueueResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["numWorkerThreads"] = o.NumWorkerThreads
	if !IsNil(o.MaxWorkQueueCapacity) {
		toSerialize["maxWorkQueueCapacity"] = o.MaxWorkQueueCapacity
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableTraditionalWorkQueueResponse struct {
	value *TraditionalWorkQueueResponse
	isSet bool
}

func (v NullableTraditionalWorkQueueResponse) Get() *TraditionalWorkQueueResponse {
	return v.value
}

func (v *NullableTraditionalWorkQueueResponse) Set(val *TraditionalWorkQueueResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTraditionalWorkQueueResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTraditionalWorkQueueResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraditionalWorkQueueResponse(val *TraditionalWorkQueueResponse) *NullableTraditionalWorkQueueResponse {
	return &NullableTraditionalWorkQueueResponse{value: val, isSet: true}
}

func (v NullableTraditionalWorkQueueResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraditionalWorkQueueResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
