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

// checks if the PwnedPasswordsPasswordValidatorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PwnedPasswordsPasswordValidatorResponse{}

// PwnedPasswordsPasswordValidatorResponse struct for PwnedPasswordsPasswordValidatorResponse
type PwnedPasswordsPasswordValidatorResponse struct {
	// Name of the Password Validator
	Id      string                                         `json:"id"`
	Schemas []EnumpwnedPasswordsPasswordValidatorSchemaUrn `json:"schemas"`
	// The base URL for requests used to interact with the Pwned Passwords service. The first five characters of the hexadecimal representation of the unsalted SHA-1 digest of a proposed password will be appended to this base URL to construct the HTTP GET request used to obtain information about potential matches.
	PwnedPasswordsBaseURL string `json:"pwnedPasswordsBaseURL"`
	// A reference to an HTTP proxy server that should be used for requests sent to the Pwned Passwords service.
	HttpProxyExternalServer *string `json:"httpProxyExternalServer,omitempty"`
	// Indicates whether this password validator should be used to validate clear-text passwords provided in LDAP add requests.
	InvokeForAdd bool `json:"invokeForAdd"`
	// Indicates whether this password validator should be used to validate clear-text passwords provided by an end user in the course of changing their own password.
	InvokeForSelfChange bool `json:"invokeForSelfChange"`
	// Indicates whether this password validator should be used to validate clear-text passwords provided by administrators when changing the password for another user.
	InvokeForAdminReset bool `json:"invokeForAdminReset"`
	// Indicates whether to accept the proposed password if an error occurs while attempting to interact with the Pwned Passwords service.
	AcceptPasswordOnServiceError bool `json:"acceptPasswordOnServiceError"`
	// Specifies which key manager provider should be used to obtain a client certificate to present to the validation server when performing HTTPS communication. This may be left undefined if communication will not be secured with HTTPS, or if there is no need to present a client certificate to the validation service.
	KeyManagerProvider *string `json:"keyManagerProvider,omitempty"`
	// Specifies which trust manager provider should be used to determine whether to trust the certificate presented by the server when performing HTTPS communication. This may be left undefined if HTTPS communication is not needed, or if the validation service presents a certificate that is trusted by the default JVM configuration (which should be the case for the Pwned Password servers).
	TrustManagerProvider *string `json:"trustManagerProvider,omitempty"`
	// A description for this Password Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether the password validator is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies a message that can be used to describe the requirements imposed by this password validator to end users. If a value is provided for this property, then it will override any description that may have otherwise been generated by the validator.
	ValidatorRequirementDescription *string `json:"validatorRequirementDescription,omitempty"`
	// Specifies a message that may be provided to the end user in the event that a proposed password is rejected by this validator. If a value is provided for this property, then it will override any failure message that may have otherwise been generated by the validator.
	ValidatorFailureMessage                       *string                                            `json:"validatorFailureMessage,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewPwnedPasswordsPasswordValidatorResponse instantiates a new PwnedPasswordsPasswordValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPwnedPasswordsPasswordValidatorResponse(id string, schemas []EnumpwnedPasswordsPasswordValidatorSchemaUrn, pwnedPasswordsBaseURL string, invokeForAdd bool, invokeForSelfChange bool, invokeForAdminReset bool, acceptPasswordOnServiceError bool, enabled bool) *PwnedPasswordsPasswordValidatorResponse {
	this := PwnedPasswordsPasswordValidatorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.PwnedPasswordsBaseURL = pwnedPasswordsBaseURL
	this.InvokeForAdd = invokeForAdd
	this.InvokeForSelfChange = invokeForSelfChange
	this.InvokeForAdminReset = invokeForAdminReset
	this.AcceptPasswordOnServiceError = acceptPasswordOnServiceError
	this.Enabled = enabled
	return &this
}

// NewPwnedPasswordsPasswordValidatorResponseWithDefaults instantiates a new PwnedPasswordsPasswordValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPwnedPasswordsPasswordValidatorResponseWithDefaults() *PwnedPasswordsPasswordValidatorResponse {
	this := PwnedPasswordsPasswordValidatorResponse{}
	return &this
}

// GetId returns the Id field value
func (o *PwnedPasswordsPasswordValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PwnedPasswordsPasswordValidatorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *PwnedPasswordsPasswordValidatorResponse) GetSchemas() []EnumpwnedPasswordsPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumpwnedPasswordsPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetSchemasOk() ([]EnumpwnedPasswordsPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *PwnedPasswordsPasswordValidatorResponse) SetSchemas(v []EnumpwnedPasswordsPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetPwnedPasswordsBaseURL returns the PwnedPasswordsBaseURL field value
func (o *PwnedPasswordsPasswordValidatorResponse) GetPwnedPasswordsBaseURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PwnedPasswordsBaseURL
}

// GetPwnedPasswordsBaseURLOk returns a tuple with the PwnedPasswordsBaseURL field value
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetPwnedPasswordsBaseURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PwnedPasswordsBaseURL, true
}

// SetPwnedPasswordsBaseURL sets field value
func (o *PwnedPasswordsPasswordValidatorResponse) SetPwnedPasswordsBaseURL(v string) {
	o.PwnedPasswordsBaseURL = v
}

// GetHttpProxyExternalServer returns the HttpProxyExternalServer field value if set, zero value otherwise.
func (o *PwnedPasswordsPasswordValidatorResponse) GetHttpProxyExternalServer() string {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		var ret string
		return ret
	}
	return *o.HttpProxyExternalServer
}

// GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetHttpProxyExternalServerOk() (*string, bool) {
	if o == nil || IsNil(o.HttpProxyExternalServer) {
		return nil, false
	}
	return o.HttpProxyExternalServer, true
}

// HasHttpProxyExternalServer returns a boolean if a field has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) HasHttpProxyExternalServer() bool {
	if o != nil && !IsNil(o.HttpProxyExternalServer) {
		return true
	}

	return false
}

// SetHttpProxyExternalServer gets a reference to the given string and assigns it to the HttpProxyExternalServer field.
func (o *PwnedPasswordsPasswordValidatorResponse) SetHttpProxyExternalServer(v string) {
	o.HttpProxyExternalServer = &v
}

// GetInvokeForAdd returns the InvokeForAdd field value
func (o *PwnedPasswordsPasswordValidatorResponse) GetInvokeForAdd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InvokeForAdd
}

// GetInvokeForAddOk returns a tuple with the InvokeForAdd field value
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetInvokeForAddOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvokeForAdd, true
}

// SetInvokeForAdd sets field value
func (o *PwnedPasswordsPasswordValidatorResponse) SetInvokeForAdd(v bool) {
	o.InvokeForAdd = v
}

// GetInvokeForSelfChange returns the InvokeForSelfChange field value
func (o *PwnedPasswordsPasswordValidatorResponse) GetInvokeForSelfChange() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InvokeForSelfChange
}

// GetInvokeForSelfChangeOk returns a tuple with the InvokeForSelfChange field value
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetInvokeForSelfChangeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvokeForSelfChange, true
}

// SetInvokeForSelfChange sets field value
func (o *PwnedPasswordsPasswordValidatorResponse) SetInvokeForSelfChange(v bool) {
	o.InvokeForSelfChange = v
}

// GetInvokeForAdminReset returns the InvokeForAdminReset field value
func (o *PwnedPasswordsPasswordValidatorResponse) GetInvokeForAdminReset() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InvokeForAdminReset
}

// GetInvokeForAdminResetOk returns a tuple with the InvokeForAdminReset field value
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetInvokeForAdminResetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvokeForAdminReset, true
}

// SetInvokeForAdminReset sets field value
func (o *PwnedPasswordsPasswordValidatorResponse) SetInvokeForAdminReset(v bool) {
	o.InvokeForAdminReset = v
}

// GetAcceptPasswordOnServiceError returns the AcceptPasswordOnServiceError field value
func (o *PwnedPasswordsPasswordValidatorResponse) GetAcceptPasswordOnServiceError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptPasswordOnServiceError
}

// GetAcceptPasswordOnServiceErrorOk returns a tuple with the AcceptPasswordOnServiceError field value
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetAcceptPasswordOnServiceErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptPasswordOnServiceError, true
}

// SetAcceptPasswordOnServiceError sets field value
func (o *PwnedPasswordsPasswordValidatorResponse) SetAcceptPasswordOnServiceError(v bool) {
	o.AcceptPasswordOnServiceError = v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *PwnedPasswordsPasswordValidatorResponse) GetKeyManagerProvider() string {
	if o == nil || IsNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyManagerProvider) {
		return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) HasKeyManagerProvider() bool {
	if o != nil && !IsNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *PwnedPasswordsPasswordValidatorResponse) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *PwnedPasswordsPasswordValidatorResponse) GetTrustManagerProvider() string {
	if o == nil || IsNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) HasTrustManagerProvider() bool {
	if o != nil && !IsNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *PwnedPasswordsPasswordValidatorResponse) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PwnedPasswordsPasswordValidatorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PwnedPasswordsPasswordValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *PwnedPasswordsPasswordValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PwnedPasswordsPasswordValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *PwnedPasswordsPasswordValidatorResponse) GetValidatorRequirementDescription() string {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) HasValidatorRequirementDescription() bool {
	if o != nil && !IsNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *PwnedPasswordsPasswordValidatorResponse) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *PwnedPasswordsPasswordValidatorResponse) GetValidatorFailureMessage() string {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) HasValidatorFailureMessage() bool {
	if o != nil && !IsNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *PwnedPasswordsPasswordValidatorResponse) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PwnedPasswordsPasswordValidatorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *PwnedPasswordsPasswordValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *PwnedPasswordsPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *PwnedPasswordsPasswordValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *PwnedPasswordsPasswordValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o PwnedPasswordsPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PwnedPasswordsPasswordValidatorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["pwnedPasswordsBaseURL"] = o.PwnedPasswordsBaseURL
	if !IsNil(o.HttpProxyExternalServer) {
		toSerialize["httpProxyExternalServer"] = o.HttpProxyExternalServer
	}
	toSerialize["invokeForAdd"] = o.InvokeForAdd
	toSerialize["invokeForSelfChange"] = o.InvokeForSelfChange
	toSerialize["invokeForAdminReset"] = o.InvokeForAdminReset
	toSerialize["acceptPasswordOnServiceError"] = o.AcceptPasswordOnServiceError
	if !IsNil(o.KeyManagerProvider) {
		toSerialize["keyManagerProvider"] = o.KeyManagerProvider
	}
	if !IsNil(o.TrustManagerProvider) {
		toSerialize["trustManagerProvider"] = o.TrustManagerProvider
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.ValidatorRequirementDescription) {
		toSerialize["validatorRequirementDescription"] = o.ValidatorRequirementDescription
	}
	if !IsNil(o.ValidatorFailureMessage) {
		toSerialize["validatorFailureMessage"] = o.ValidatorFailureMessage
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullablePwnedPasswordsPasswordValidatorResponse struct {
	value *PwnedPasswordsPasswordValidatorResponse
	isSet bool
}

func (v NullablePwnedPasswordsPasswordValidatorResponse) Get() *PwnedPasswordsPasswordValidatorResponse {
	return v.value
}

func (v *NullablePwnedPasswordsPasswordValidatorResponse) Set(val *PwnedPasswordsPasswordValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePwnedPasswordsPasswordValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePwnedPasswordsPasswordValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePwnedPasswordsPasswordValidatorResponse(val *PwnedPasswordsPasswordValidatorResponse) *NullablePwnedPasswordsPasswordValidatorResponse {
	return &NullablePwnedPasswordsPasswordValidatorResponse{value: val, isSet: true}
}

func (v NullablePwnedPasswordsPasswordValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePwnedPasswordsPasswordValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
