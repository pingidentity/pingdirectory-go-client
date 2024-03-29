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

// checks if the AddPingOneHttpExternalServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPingOneHttpExternalServerRequest{}

// AddPingOneHttpExternalServerRequest struct for AddPingOneHttpExternalServerRequest
type AddPingOneHttpExternalServerRequest struct {
	Schemas                    []EnumpingOneHttpExternalServerSchemaUrn                     `json:"schemas"`
	HostnameVerificationMethod *EnumexternalServerPingOneHttpHostnameVerificationMethodProp `json:"hostnameVerificationMethod,omitempty"`
	// The trust manager provider to use for HTTPS connection-level security.
	TrustManagerProvider *string `json:"trustManagerProvider,omitempty"`
	// Specifies the maximum length of time to wait for a connection to be established before aborting a request to PingOne.
	ConnectTimeout *string `json:"connectTimeout,omitempty"`
	// Specifies the maximum length of time to wait for response data to be read from an established connection before aborting a request to PingOne.
	ResponseTimeout *string `json:"responseTimeout,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
	// Name of the new External Server
	ServerName string `json:"serverName"`
}

// NewAddPingOneHttpExternalServerRequest instantiates a new AddPingOneHttpExternalServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPingOneHttpExternalServerRequest(schemas []EnumpingOneHttpExternalServerSchemaUrn, serverName string) *AddPingOneHttpExternalServerRequest {
	this := AddPingOneHttpExternalServerRequest{}
	this.Schemas = schemas
	this.ServerName = serverName
	return &this
}

// NewAddPingOneHttpExternalServerRequestWithDefaults instantiates a new AddPingOneHttpExternalServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPingOneHttpExternalServerRequestWithDefaults() *AddPingOneHttpExternalServerRequest {
	this := AddPingOneHttpExternalServerRequest{}
	return &this
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
func (o *AddPingOneHttpExternalServerRequest) GetHostnameVerificationMethod() EnumexternalServerPingOneHttpHostnameVerificationMethodProp {
	if o == nil || IsNil(o.HostnameVerificationMethod) {
		var ret EnumexternalServerPingOneHttpHostnameVerificationMethodProp
		return ret
	}
	return *o.HostnameVerificationMethod
}

// GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetHostnameVerificationMethodOk() (*EnumexternalServerPingOneHttpHostnameVerificationMethodProp, bool) {
	if o == nil || IsNil(o.HostnameVerificationMethod) {
		return nil, false
	}
	return o.HostnameVerificationMethod, true
}

// HasHostnameVerificationMethod returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasHostnameVerificationMethod() bool {
	if o != nil && !IsNil(o.HostnameVerificationMethod) {
		return true
	}

	return false
}

// SetHostnameVerificationMethod gets a reference to the given EnumexternalServerPingOneHttpHostnameVerificationMethodProp and assigns it to the HostnameVerificationMethod field.
func (o *AddPingOneHttpExternalServerRequest) SetHostnameVerificationMethod(v EnumexternalServerPingOneHttpHostnameVerificationMethodProp) {
	o.HostnameVerificationMethod = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *AddPingOneHttpExternalServerRequest) GetTrustManagerProvider() string {
	if o == nil || IsNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasTrustManagerProvider() bool {
	if o != nil && !IsNil(o.TrustManagerProvider) {
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
	if o == nil || IsNil(o.ConnectTimeout) {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasConnectTimeout() bool {
	if o != nil && !IsNil(o.ConnectTimeout) {
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
	if o == nil || IsNil(o.ResponseTimeout) {
		var ret string
		return ret
	}
	return *o.ResponseTimeout
}

// GetResponseTimeoutOk returns a tuple with the ResponseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetResponseTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ResponseTimeout) {
		return nil, false
	}
	return o.ResponseTimeout, true
}

// HasResponseTimeout returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasResponseTimeout() bool {
	if o != nil && !IsNil(o.ResponseTimeout) {
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
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneHttpExternalServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPingOneHttpExternalServerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPingOneHttpExternalServerRequest) SetDescription(v string) {
	o.Description = &v
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

func (o AddPingOneHttpExternalServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPingOneHttpExternalServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.HostnameVerificationMethod) {
		toSerialize["hostnameVerificationMethod"] = o.HostnameVerificationMethod
	}
	if !IsNil(o.TrustManagerProvider) {
		toSerialize["trustManagerProvider"] = o.TrustManagerProvider
	}
	if !IsNil(o.ConnectTimeout) {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if !IsNil(o.ResponseTimeout) {
		toSerialize["responseTimeout"] = o.ResponseTimeout
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["serverName"] = o.ServerName
	return toSerialize, nil
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
