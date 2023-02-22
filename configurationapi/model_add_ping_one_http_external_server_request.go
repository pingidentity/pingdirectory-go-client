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

// AddPingOneHttpExternalServerRequest struct for AddPingOneHttpExternalServerRequest
type AddPingOneHttpExternalServerRequest struct {
	// Name of the new External Server
	ServerName                 string                                            `json:"serverName"`
	Schemas                    []EnumpingOneHttpExternalServerSchemaUrn          `json:"schemas"`
	HostnameVerificationMethod *EnumexternalServerHostnameVerificationMethodProp `json:"hostnameVerificationMethod,omitempty"`
	// The trust manager provider to use for HTTPS connection-level security.
	TrustManagerProvider *string `json:"trustManagerProvider,omitempty"`
	// Specifies the maximum length of time to wait for a connection to be established before aborting a request to PingOne.
	ConnectTimeout *string `json:"connectTimeout,omitempty"`
	// Specifies the maximum length of time to wait for response data to be read from an established connection before aborting a request to PingOne.
	ResponseTimeout *string `json:"responseTimeout,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewAddPingOneHttpExternalServerRequest instantiates a new AddPingOneHttpExternalServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPingOneHttpExternalServerRequest(serverName string, schemas []EnumpingOneHttpExternalServerSchemaUrn) *AddPingOneHttpExternalServerRequest {
	this := AddPingOneHttpExternalServerRequest{}
	this.ServerName = serverName
	this.Schemas = schemas
	return &this
}

// NewAddPingOneHttpExternalServerRequestWithDefaults instantiates a new AddPingOneHttpExternalServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPingOneHttpExternalServerRequestWithDefaults() *AddPingOneHttpExternalServerRequest {
	this := AddPingOneHttpExternalServerRequest{}
	return &this
}

// GetServerName returns the ServerName field value
func (o *AddPingOneHttpExternalServerRequest) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *AddPingOneHttpExternalServerRequest) SetServerName(v string) {
	o.ServerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPingOneHttpExternalServerRequest) GetSchemas() []EnumpingOneHttpExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumpingOneHttpExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetSchemasOk() ([]EnumpingOneHttpExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPingOneHttpExternalServerRequest) SetSchemas(v []EnumpingOneHttpExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetHostnameVerificationMethod returns the HostnameVerificationMethod field value if set, zero value otherwise.
func (o *AddPingOneHttpExternalServerRequest) GetHostnameVerificationMethod() EnumexternalServerHostnameVerificationMethodProp {
	if o == nil || isNil(o.HostnameVerificationMethod) {
		var ret EnumexternalServerHostnameVerificationMethodProp
		return ret
	}
	return *o.HostnameVerificationMethod
}

// GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetHostnameVerificationMethodOk() (*EnumexternalServerHostnameVerificationMethodProp, bool) {
	if o == nil || isNil(o.HostnameVerificationMethod) {
		return nil, false
	}
	return o.HostnameVerificationMethod, true
}

// HasHostnameVerificationMethod returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasHostnameVerificationMethod() bool {
	if o != nil && !isNil(o.HostnameVerificationMethod) {
		return true
	}

	return false
}

// SetHostnameVerificationMethod gets a reference to the given EnumexternalServerHostnameVerificationMethodProp and assigns it to the HostnameVerificationMethod field.
func (o *AddPingOneHttpExternalServerRequest) SetHostnameVerificationMethod(v EnumexternalServerHostnameVerificationMethodProp) {
	o.HostnameVerificationMethod = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *AddPingOneHttpExternalServerRequest) GetTrustManagerProvider() string {
	if o == nil || isNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || isNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasTrustManagerProvider() bool {
	if o != nil && !isNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *AddPingOneHttpExternalServerRequest) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *AddPingOneHttpExternalServerRequest) GetConnectTimeout() string {
	if o == nil || isNil(o.ConnectTimeout) {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasConnectTimeout() bool {
	if o != nil && !isNil(o.ConnectTimeout) {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given string and assigns it to the ConnectTimeout field.
func (o *AddPingOneHttpExternalServerRequest) SetConnectTimeout(v string) {
	o.ConnectTimeout = &v
}

// GetResponseTimeout returns the ResponseTimeout field value if set, zero value otherwise.
func (o *AddPingOneHttpExternalServerRequest) GetResponseTimeout() string {
	if o == nil || isNil(o.ResponseTimeout) {
		var ret string
		return ret
	}
	return *o.ResponseTimeout
}

// GetResponseTimeoutOk returns a tuple with the ResponseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetResponseTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.ResponseTimeout) {
		return nil, false
	}
	return o.ResponseTimeout, true
}

// HasResponseTimeout returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasResponseTimeout() bool {
	if o != nil && !isNil(o.ResponseTimeout) {
		return true
	}

	return false
}

// SetResponseTimeout gets a reference to the given string and assigns it to the ResponseTimeout field.
func (o *AddPingOneHttpExternalServerRequest) SetResponseTimeout(v string) {
	o.ResponseTimeout = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPingOneHttpExternalServerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPingOneHttpExternalServerRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddPingOneHttpExternalServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serverName"] = o.ServerName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.HostnameVerificationMethod) {
		toSerialize["hostnameVerificationMethod"] = o.HostnameVerificationMethod
	}
	if !isNil(o.TrustManagerProvider) {
		toSerialize["trustManagerProvider"] = o.TrustManagerProvider
	}
	if !isNil(o.ConnectTimeout) {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if !isNil(o.ResponseTimeout) {
		toSerialize["responseTimeout"] = o.ResponseTimeout
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAddPingOneHttpExternalServerRequest struct {
	value *AddPingOneHttpExternalServerRequest
	isSet bool
}

func (v NullableAddPingOneHttpExternalServerRequest) Get() *AddPingOneHttpExternalServerRequest {
	return v.value
}

func (v *NullableAddPingOneHttpExternalServerRequest) Set(val *AddPingOneHttpExternalServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPingOneHttpExternalServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPingOneHttpExternalServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPingOneHttpExternalServerRequest(val *AddPingOneHttpExternalServerRequest) *NullableAddPingOneHttpExternalServerRequest {
	return &NullableAddPingOneHttpExternalServerRequest{value: val, isSet: true}
}

func (v NullableAddPingOneHttpExternalServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPingOneHttpExternalServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}