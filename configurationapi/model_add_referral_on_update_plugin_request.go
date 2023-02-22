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

// AddReferralOnUpdatePluginRequest struct for AddReferralOnUpdatePluginRequest
type AddReferralOnUpdatePluginRequest struct {
	// Name of the new Plugin
	PluginName string                                `json:"pluginName"`
	Schemas    []EnumreferralOnUpdatePluginSchemaUrn `json:"schemas"`
	PluginType []EnumpluginPluginTypeProp            `json:"pluginType"`
	// Specifies the base URL to use for the referrals generated by this plugin. It should include only the scheme, address, and port to use to communicate with the target server (e.g., \"ldap://server.example.com:389/\").
	ReferralBaseURL []string `json:"referralBaseURL"`
	// Specifies a base DN for requests for which to send referrals in response to update operations.
	BaseDN []string `json:"baseDN,omitempty"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddReferralOnUpdatePluginRequest instantiates a new AddReferralOnUpdatePluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddReferralOnUpdatePluginRequest(pluginName string, schemas []EnumreferralOnUpdatePluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, referralBaseURL []string, enabled bool) *AddReferralOnUpdatePluginRequest {
	this := AddReferralOnUpdatePluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.PluginType = pluginType
	this.ReferralBaseURL = referralBaseURL
	this.Enabled = enabled
	return &this
}

// NewAddReferralOnUpdatePluginRequestWithDefaults instantiates a new AddReferralOnUpdatePluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddReferralOnUpdatePluginRequestWithDefaults() *AddReferralOnUpdatePluginRequest {
	this := AddReferralOnUpdatePluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddReferralOnUpdatePluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddReferralOnUpdatePluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddReferralOnUpdatePluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddReferralOnUpdatePluginRequest) GetSchemas() []EnumreferralOnUpdatePluginSchemaUrn {
	if o == nil {
		var ret []EnumreferralOnUpdatePluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddReferralOnUpdatePluginRequest) GetSchemasOk() ([]EnumreferralOnUpdatePluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddReferralOnUpdatePluginRequest) SetSchemas(v []EnumreferralOnUpdatePluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value
func (o *AddReferralOnUpdatePluginRequest) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *AddReferralOnUpdatePluginRequest) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *AddReferralOnUpdatePluginRequest) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetReferralBaseURL returns the ReferralBaseURL field value
func (o *AddReferralOnUpdatePluginRequest) GetReferralBaseURL() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReferralBaseURL
}

// GetReferralBaseURLOk returns a tuple with the ReferralBaseURL field value
// and a boolean to check if the value has been set.
func (o *AddReferralOnUpdatePluginRequest) GetReferralBaseURLOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferralBaseURL, true
}

// SetReferralBaseURL sets field value
func (o *AddReferralOnUpdatePluginRequest) SetReferralBaseURL(v []string) {
	o.ReferralBaseURL = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddReferralOnUpdatePluginRequest) GetBaseDN() []string {
	if o == nil || isNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferralOnUpdatePluginRequest) GetBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddReferralOnUpdatePluginRequest) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *AddReferralOnUpdatePluginRequest) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddReferralOnUpdatePluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferralOnUpdatePluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddReferralOnUpdatePluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddReferralOnUpdatePluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddReferralOnUpdatePluginRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferralOnUpdatePluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddReferralOnUpdatePluginRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddReferralOnUpdatePluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddReferralOnUpdatePluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddReferralOnUpdatePluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddReferralOnUpdatePluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddReferralOnUpdatePluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pluginName"] = o.PluginName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["pluginType"] = o.PluginType
	}
	if true {
		toSerialize["referralBaseURL"] = o.ReferralBaseURL
	}
	if !isNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !isNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddReferralOnUpdatePluginRequest struct {
	value *AddReferralOnUpdatePluginRequest
	isSet bool
}

func (v NullableAddReferralOnUpdatePluginRequest) Get() *AddReferralOnUpdatePluginRequest {
	return v.value
}

func (v *NullableAddReferralOnUpdatePluginRequest) Set(val *AddReferralOnUpdatePluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddReferralOnUpdatePluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddReferralOnUpdatePluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddReferralOnUpdatePluginRequest(val *AddReferralOnUpdatePluginRequest) *NullableAddReferralOnUpdatePluginRequest {
	return &NullableAddReferralOnUpdatePluginRequest{value: val, isSet: true}
}

func (v NullableAddReferralOnUpdatePluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddReferralOnUpdatePluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}