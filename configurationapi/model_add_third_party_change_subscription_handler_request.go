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

// checks if the AddThirdPartyChangeSubscriptionHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyChangeSubscriptionHandlerRequest{}

// AddThirdPartyChangeSubscriptionHandlerRequest struct for AddThirdPartyChangeSubscriptionHandlerRequest
type AddThirdPartyChangeSubscriptionHandlerRequest struct {
	Schemas []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Change Subscription Handler.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Change Subscription Handler. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Change Subscription Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether this change subscription handler is enabled within the server.
	Enabled bool `json:"enabled"`
	// The set of change subscriptions for which this change subscription handler should be notified. If no values are provided then it will be notified for all change subscriptions defined in the server.
	ChangeSubscription []string `json:"changeSubscription,omitempty"`
	// Name of the new Change Subscription Handler
	HandlerName string `json:"handlerName"`
}

// NewAddThirdPartyChangeSubscriptionHandlerRequest instantiates a new AddThirdPartyChangeSubscriptionHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyChangeSubscriptionHandlerRequest(schemas []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, extensionClass string, enabled bool, handlerName string) *AddThirdPartyChangeSubscriptionHandlerRequest {
	this := AddThirdPartyChangeSubscriptionHandlerRequest{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.HandlerName = handlerName
	return &this
}

// NewAddThirdPartyChangeSubscriptionHandlerRequestWithDefaults instantiates a new AddThirdPartyChangeSubscriptionHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyChangeSubscriptionHandlerRequestWithDefaults() *AddThirdPartyChangeSubscriptionHandlerRequest {
	this := AddThirdPartyChangeSubscriptionHandlerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetSchemas() []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetSchemasOk() ([]EnumthirdPartyChangeSubscriptionHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) SetSchemas(v []EnumthirdPartyChangeSubscriptionHandlerSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetChangeSubscription returns the ChangeSubscription field value if set, zero value otherwise.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetChangeSubscription() []string {
	if o == nil || IsNil(o.ChangeSubscription) {
		var ret []string
		return ret
	}
	return o.ChangeSubscription
}

// GetChangeSubscriptionOk returns a tuple with the ChangeSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetChangeSubscriptionOk() ([]string, bool) {
	if o == nil || IsNil(o.ChangeSubscription) {
		return nil, false
	}
	return o.ChangeSubscription, true
}

// HasChangeSubscription returns a boolean if a field has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) HasChangeSubscription() bool {
	if o != nil && !IsNil(o.ChangeSubscription) {
		return true
	}

	return false
}

// SetChangeSubscription gets a reference to the given []string and assigns it to the ChangeSubscription field.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) SetChangeSubscription(v []string) {
	o.ChangeSubscription = v
}

// GetHandlerName returns the HandlerName field value
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddThirdPartyChangeSubscriptionHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddThirdPartyChangeSubscriptionHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyChangeSubscriptionHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.ChangeSubscription) {
		toSerialize["changeSubscription"] = o.ChangeSubscription
	}
	toSerialize["handlerName"] = o.HandlerName
	return toSerialize, nil
}

type NullableAddThirdPartyChangeSubscriptionHandlerRequest struct {
	value *AddThirdPartyChangeSubscriptionHandlerRequest
	isSet bool
}

func (v NullableAddThirdPartyChangeSubscriptionHandlerRequest) Get() *AddThirdPartyChangeSubscriptionHandlerRequest {
	return v.value
}

func (v *NullableAddThirdPartyChangeSubscriptionHandlerRequest) Set(val *AddThirdPartyChangeSubscriptionHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyChangeSubscriptionHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyChangeSubscriptionHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyChangeSubscriptionHandlerRequest(val *AddThirdPartyChangeSubscriptionHandlerRequest) *NullableAddThirdPartyChangeSubscriptionHandlerRequest {
	return &NullableAddThirdPartyChangeSubscriptionHandlerRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyChangeSubscriptionHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyChangeSubscriptionHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
