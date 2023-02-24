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

// checks if the AddPwnedPasswordsPasswordValidatorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPwnedPasswordsPasswordValidatorRequest{}

// AddPwnedPasswordsPasswordValidatorRequest struct for AddPwnedPasswordsPasswordValidatorRequest
type AddPwnedPasswordsPasswordValidatorRequest struct {
	// Name of the new Password Validator
	ValidatorName string                                         `json:"validatorName"`
	Schemas       []EnumpwnedPasswordsPasswordValidatorSchemaUrn `json:"schemas"`
	// The base URL for requests used to interact with the Pwned Passwords service. The first five characters of the hexadecimal representation of the unsalted SHA-1 digest of a proposed password will be appended to this base URL to construct the HTTP GET request used to obtain information about potential matches.
	PwnedPasswordsBaseURL *string `json:"pwnedPasswordsBaseURL,omitempty"`
	// Indicates whether this password validator should be used to validate clear-text passwords provided in LDAP add requests.
	InvokeForAdd *bool `json:"invokeForAdd,omitempty"`
	// Indicates whether this password validator should be used to validate clear-text passwords provided by an end user in the course of changing their own password.
	InvokeForSelfChange *bool `json:"invokeForSelfChange,omitempty"`
	// Indicates whether this password validator should be used to validate clear-text passwords provided by administrators when changing the password for another user.
	InvokeForAdminReset *bool `json:"invokeForAdminReset,omitempty"`
	// Indicates whether to accept the proposed password if an error occurs while attempting to interact with the Pwned Passwords service.
	AcceptPasswordOnServiceError *bool `json:"acceptPasswordOnServiceError,omitempty"`
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
	ValidatorFailureMessage *string `json:"validatorFailureMessage,omitempty"`
}

// NewAddPwnedPasswordsPasswordValidatorRequest instantiates a new AddPwnedPasswordsPasswordValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPwnedPasswordsPasswordValidatorRequest(validatorName string, schemas []EnumpwnedPasswordsPasswordValidatorSchemaUrn, enabled bool) *AddPwnedPasswordsPasswordValidatorRequest {
	this := AddPwnedPasswordsPasswordValidatorRequest{}
	this.ValidatorName = validatorName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddPwnedPasswordsPasswordValidatorRequestWithDefaults instantiates a new AddPwnedPasswordsPasswordValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPwnedPasswordsPasswordValidatorRequestWithDefaults() *AddPwnedPasswordsPasswordValidatorRequest {
	this := AddPwnedPasswordsPasswordValidatorRequest{}
	return &this
}

// GetValidatorName returns the ValidatorName field value
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetSchemas() []EnumpwnedPasswordsPasswordValidatorSchemaUrn {
	if o == nil {
		var ret []EnumpwnedPasswordsPasswordValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetSchemasOk() ([]EnumpwnedPasswordsPasswordValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetSchemas(v []EnumpwnedPasswordsPasswordValidatorSchemaUrn) {
	o.Schemas = v
}

// GetPwnedPasswordsBaseURL returns the PwnedPasswordsBaseURL field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetPwnedPasswordsBaseURL() string {
	if o == nil || IsNil(o.PwnedPasswordsBaseURL) {
		var ret string
		return ret
	}
	return *o.PwnedPasswordsBaseURL
}

// GetPwnedPasswordsBaseURLOk returns a tuple with the PwnedPasswordsBaseURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetPwnedPasswordsBaseURLOk() (*string, bool) {
	if o == nil || IsNil(o.PwnedPasswordsBaseURL) {
		return nil, false
	}
	return o.PwnedPasswordsBaseURL, true
}

// HasPwnedPasswordsBaseURL returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasPwnedPasswordsBaseURL() bool {
	if o != nil && !IsNil(o.PwnedPasswordsBaseURL) {
		return true
	}

	return false
}

// SetPwnedPasswordsBaseURL gets a reference to the given string and assigns it to the PwnedPasswordsBaseURL field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetPwnedPasswordsBaseURL(v string) {
	o.PwnedPasswordsBaseURL = &v
}

// GetInvokeForAdd returns the InvokeForAdd field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetInvokeForAdd() bool {
	if o == nil || IsNil(o.InvokeForAdd) {
		var ret bool
		return ret
	}
	return *o.InvokeForAdd
}

// GetInvokeForAddOk returns a tuple with the InvokeForAdd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetInvokeForAddOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForAdd) {
		return nil, false
	}
	return o.InvokeForAdd, true
}

// HasInvokeForAdd returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasInvokeForAdd() bool {
	if o != nil && !IsNil(o.InvokeForAdd) {
		return true
	}

	return false
}

// SetInvokeForAdd gets a reference to the given bool and assigns it to the InvokeForAdd field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetInvokeForAdd(v bool) {
	o.InvokeForAdd = &v
}

// GetInvokeForSelfChange returns the InvokeForSelfChange field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetInvokeForSelfChange() bool {
	if o == nil || IsNil(o.InvokeForSelfChange) {
		var ret bool
		return ret
	}
	return *o.InvokeForSelfChange
}

// GetInvokeForSelfChangeOk returns a tuple with the InvokeForSelfChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetInvokeForSelfChangeOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForSelfChange) {
		return nil, false
	}
	return o.InvokeForSelfChange, true
}

// HasInvokeForSelfChange returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasInvokeForSelfChange() bool {
	if o != nil && !IsNil(o.InvokeForSelfChange) {
		return true
	}

	return false
}

// SetInvokeForSelfChange gets a reference to the given bool and assigns it to the InvokeForSelfChange field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetInvokeForSelfChange(v bool) {
	o.InvokeForSelfChange = &v
}

// GetInvokeForAdminReset returns the InvokeForAdminReset field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetInvokeForAdminReset() bool {
	if o == nil || IsNil(o.InvokeForAdminReset) {
		var ret bool
		return ret
	}
	return *o.InvokeForAdminReset
}

// GetInvokeForAdminResetOk returns a tuple with the InvokeForAdminReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetInvokeForAdminResetOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForAdminReset) {
		return nil, false
	}
	return o.InvokeForAdminReset, true
}

// HasInvokeForAdminReset returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasInvokeForAdminReset() bool {
	if o != nil && !IsNil(o.InvokeForAdminReset) {
		return true
	}

	return false
}

// SetInvokeForAdminReset gets a reference to the given bool and assigns it to the InvokeForAdminReset field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetInvokeForAdminReset(v bool) {
	o.InvokeForAdminReset = &v
}

// GetAcceptPasswordOnServiceError returns the AcceptPasswordOnServiceError field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetAcceptPasswordOnServiceError() bool {
	if o == nil || IsNil(o.AcceptPasswordOnServiceError) {
		var ret bool
		return ret
	}
	return *o.AcceptPasswordOnServiceError
}

// GetAcceptPasswordOnServiceErrorOk returns a tuple with the AcceptPasswordOnServiceError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetAcceptPasswordOnServiceErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptPasswordOnServiceError) {
		return nil, false
	}
	return o.AcceptPasswordOnServiceError, true
}

// HasAcceptPasswordOnServiceError returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasAcceptPasswordOnServiceError() bool {
	if o != nil && !IsNil(o.AcceptPasswordOnServiceError) {
		return true
	}

	return false
}

// SetAcceptPasswordOnServiceError gets a reference to the given bool and assigns it to the AcceptPasswordOnServiceError field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetAcceptPasswordOnServiceError(v bool) {
	o.AcceptPasswordOnServiceError = &v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetKeyManagerProvider() string {
	if o == nil || IsNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyManagerProvider) {
		return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasKeyManagerProvider() bool {
	if o != nil && !IsNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetTrustManagerProvider() string {
	if o == nil || IsNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasTrustManagerProvider() bool {
	if o != nil && !IsNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetValidatorRequirementDescription returns the ValidatorRequirementDescription field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetValidatorRequirementDescription() string {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		var ret string
		return ret
	}
	return *o.ValidatorRequirementDescription
}

// GetValidatorRequirementDescriptionOk returns a tuple with the ValidatorRequirementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetValidatorRequirementDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorRequirementDescription) {
		return nil, false
	}
	return o.ValidatorRequirementDescription, true
}

// HasValidatorRequirementDescription returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasValidatorRequirementDescription() bool {
	if o != nil && !IsNil(o.ValidatorRequirementDescription) {
		return true
	}

	return false
}

// SetValidatorRequirementDescription gets a reference to the given string and assigns it to the ValidatorRequirementDescription field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetValidatorRequirementDescription(v string) {
	o.ValidatorRequirementDescription = &v
}

// GetValidatorFailureMessage returns the ValidatorFailureMessage field value if set, zero value otherwise.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetValidatorFailureMessage() string {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		var ret string
		return ret
	}
	return *o.ValidatorFailureMessage
}

// GetValidatorFailureMessageOk returns a tuple with the ValidatorFailureMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) GetValidatorFailureMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ValidatorFailureMessage) {
		return nil, false
	}
	return o.ValidatorFailureMessage, true
}

// HasValidatorFailureMessage returns a boolean if a field has been set.
func (o *AddPwnedPasswordsPasswordValidatorRequest) HasValidatorFailureMessage() bool {
	if o != nil && !IsNil(o.ValidatorFailureMessage) {
		return true
	}

	return false
}

// SetValidatorFailureMessage gets a reference to the given string and assigns it to the ValidatorFailureMessage field.
func (o *AddPwnedPasswordsPasswordValidatorRequest) SetValidatorFailureMessage(v string) {
	o.ValidatorFailureMessage = &v
}

func (o AddPwnedPasswordsPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPwnedPasswordsPasswordValidatorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["validatorName"] = o.ValidatorName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.PwnedPasswordsBaseURL) {
		toSerialize["pwnedPasswordsBaseURL"] = o.PwnedPasswordsBaseURL
	}
	if !IsNil(o.InvokeForAdd) {
		toSerialize["invokeForAdd"] = o.InvokeForAdd
	}
	if !IsNil(o.InvokeForSelfChange) {
		toSerialize["invokeForSelfChange"] = o.InvokeForSelfChange
	}
	if !IsNil(o.InvokeForAdminReset) {
		toSerialize["invokeForAdminReset"] = o.InvokeForAdminReset
	}
	if !IsNil(o.AcceptPasswordOnServiceError) {
		toSerialize["acceptPasswordOnServiceError"] = o.AcceptPasswordOnServiceError
	}
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
	return toSerialize, nil
}

type NullableAddPwnedPasswordsPasswordValidatorRequest struct {
	value *AddPwnedPasswordsPasswordValidatorRequest
	isSet bool
}

func (v NullableAddPwnedPasswordsPasswordValidatorRequest) Get() *AddPwnedPasswordsPasswordValidatorRequest {
	return v.value
}

func (v *NullableAddPwnedPasswordsPasswordValidatorRequest) Set(val *AddPwnedPasswordsPasswordValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPwnedPasswordsPasswordValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPwnedPasswordsPasswordValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPwnedPasswordsPasswordValidatorRequest(val *AddPwnedPasswordsPasswordValidatorRequest) *NullableAddPwnedPasswordsPasswordValidatorRequest {
	return &NullableAddPwnedPasswordsPasswordValidatorRequest{value: val, isSet: true}
}

func (v NullableAddPwnedPasswordsPasswordValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPwnedPasswordsPasswordValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
