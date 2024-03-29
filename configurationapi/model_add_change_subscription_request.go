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

// checks if the AddChangeSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddChangeSubscriptionRequest{}

// AddChangeSubscriptionRequest struct for AddChangeSubscriptionRequest
type AddChangeSubscriptionRequest struct {
	Schemas []EnumchangeSubscriptionSchemaUrn `json:"schemas,omitempty"`
	// A description for this Change Subscription
	Description *string `json:"description,omitempty"`
	// Specifies a set of connection criteria that must match the client connection associated with an operation in order for that operation to be processed by a change subscription handler.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// Specifies a set of request criteria that must match the request associated with an operation in order for that operation to be processed by a change subscription handler.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Specifies a set of result criteria that must match the result associated with an operation in order for that operation to be processed by a change subscription handler.
	ResultCriteria *string `json:"resultCriteria,omitempty"`
	// Specifies a timestamp that provides an expiration time for this change subscription. If an expiration time is provided, then the change subscription will not be active after that time has passed.
	ExpirationTime *string `json:"expirationTime,omitempty"`
	// Name of the new Change Subscription
	SubscriptionName string `json:"subscriptionName"`
}

// NewAddChangeSubscriptionRequest instantiates a new AddChangeSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddChangeSubscriptionRequest(subscriptionName string) *AddChangeSubscriptionRequest {
	this := AddChangeSubscriptionRequest{}
	this.SubscriptionName = subscriptionName
	return &this
}

// NewAddChangeSubscriptionRequestWithDefaults instantiates a new AddChangeSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddChangeSubscriptionRequestWithDefaults() *AddChangeSubscriptionRequest {
	this := AddChangeSubscriptionRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddChangeSubscriptionRequest) GetSchemas() []EnumchangeSubscriptionSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumchangeSubscriptionSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddChangeSubscriptionRequest) GetSchemasOk() ([]EnumchangeSubscriptionSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddChangeSubscriptionRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumchangeSubscriptionSchemaUrn and assigns it to the Schemas field.
func (o *AddChangeSubscriptionRequest) SetSchemas(v []EnumchangeSubscriptionSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddChangeSubscriptionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddChangeSubscriptionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddChangeSubscriptionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddChangeSubscriptionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddChangeSubscriptionRequest) GetConnectionCriteria() string {
	if o == nil || IsNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddChangeSubscriptionRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddChangeSubscriptionRequest) HasConnectionCriteria() bool {
	if o != nil && !IsNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddChangeSubscriptionRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddChangeSubscriptionRequest) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddChangeSubscriptionRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddChangeSubscriptionRequest) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddChangeSubscriptionRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetResultCriteria returns the ResultCriteria field value if set, zero value otherwise.
func (o *AddChangeSubscriptionRequest) GetResultCriteria() string {
	if o == nil || IsNil(o.ResultCriteria) {
		var ret string
		return ret
	}
	return *o.ResultCriteria
}

// GetResultCriteriaOk returns a tuple with the ResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddChangeSubscriptionRequest) GetResultCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCriteria) {
		return nil, false
	}
	return o.ResultCriteria, true
}

// HasResultCriteria returns a boolean if a field has been set.
func (o *AddChangeSubscriptionRequest) HasResultCriteria() bool {
	if o != nil && !IsNil(o.ResultCriteria) {
		return true
	}

	return false
}

// SetResultCriteria gets a reference to the given string and assigns it to the ResultCriteria field.
func (o *AddChangeSubscriptionRequest) SetResultCriteria(v string) {
	o.ResultCriteria = &v
}

// GetExpirationTime returns the ExpirationTime field value if set, zero value otherwise.
func (o *AddChangeSubscriptionRequest) GetExpirationTime() string {
	if o == nil || IsNil(o.ExpirationTime) {
		var ret string
		return ret
	}
	return *o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddChangeSubscriptionRequest) GetExpirationTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationTime) {
		return nil, false
	}
	return o.ExpirationTime, true
}

// HasExpirationTime returns a boolean if a field has been set.
func (o *AddChangeSubscriptionRequest) HasExpirationTime() bool {
	if o != nil && !IsNil(o.ExpirationTime) {
		return true
	}

	return false
}

// SetExpirationTime gets a reference to the given string and assigns it to the ExpirationTime field.
func (o *AddChangeSubscriptionRequest) SetExpirationTime(v string) {
	o.ExpirationTime = &v
}

// GetSubscriptionName returns the SubscriptionName field value
func (o *AddChangeSubscriptionRequest) GetSubscriptionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionName
}

// GetSubscriptionNameOk returns a tuple with the SubscriptionName field value
// and a boolean to check if the value has been set.
func (o *AddChangeSubscriptionRequest) GetSubscriptionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionName, true
}

// SetSubscriptionName sets field value
func (o *AddChangeSubscriptionRequest) SetSubscriptionName(v string) {
	o.SubscriptionName = v
}

func (o AddChangeSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddChangeSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.ResultCriteria) {
		toSerialize["resultCriteria"] = o.ResultCriteria
	}
	if !IsNil(o.ExpirationTime) {
		toSerialize["expirationTime"] = o.ExpirationTime
	}
	toSerialize["subscriptionName"] = o.SubscriptionName
	return toSerialize, nil
}

type NullableAddChangeSubscriptionRequest struct {
	value *AddChangeSubscriptionRequest
	isSet bool
}

func (v NullableAddChangeSubscriptionRequest) Get() *AddChangeSubscriptionRequest {
	return v.value
}

func (v *NullableAddChangeSubscriptionRequest) Set(val *AddChangeSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddChangeSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddChangeSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddChangeSubscriptionRequest(val *AddChangeSubscriptionRequest) *NullableAddChangeSubscriptionRequest {
	return &NullableAddChangeSubscriptionRequest{value: val, isSet: true}
}

func (v NullableAddChangeSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddChangeSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
