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

// OauthBearerSaslMechanismHandlerShared struct for OauthBearerSaslMechanismHandlerShared
type OauthBearerSaslMechanismHandlerShared struct {
	Schemas []EnumoauthBearerSaslMechanismHandlerSchemaUrn `json:"schemas"`
	AccessTokenValidator []string `json:"accessTokenValidator,omitempty"`
	IdTokenValidator []string `json:"idTokenValidator,omitempty"`
	// Indicates whether bind requests will be required to have both an OAuth access token (in the \"auth\" element of the bind request) and an OpenID Connect ID token (in the \"pingidentityidtoken\" element of the bind request).
	RequireBothAccessTokenAndIDToken *bool `json:"requireBothAccessTokenAndIDToken,omitempty"`
	ValidateAccessTokenWhenIDTokenIsAlsoProvided *EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp `json:"validateAccessTokenWhenIDTokenIsAlsoProvided,omitempty"`
	// The identity mapper that will be used to map an alternate authorization identity (provided in the GS2 header of the encoded OAUTHBEARER bind request credentials) to the corresponding local entry.
	AlternateAuthorizationIdentityMapper *string `json:"alternateAuthorizationIdentityMapper,omitempty"`
	AllRequiredScope []string `json:"allRequiredScope,omitempty"`
	AnyRequiredScope []string `json:"anyRequiredScope,omitempty"`
	// The fully-qualified name that clients are expected to use when communicating with the server.
	ServerFqdn *string `json:"serverFqdn,omitempty"`
	// A description for this SASL Mechanism Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the SASL mechanism handler is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewOauthBearerSaslMechanismHandlerShared instantiates a new OauthBearerSaslMechanismHandlerShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthBearerSaslMechanismHandlerShared(schemas []EnumoauthBearerSaslMechanismHandlerSchemaUrn, enabled bool) *OauthBearerSaslMechanismHandlerShared {
	this := OauthBearerSaslMechanismHandlerShared{}
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewOauthBearerSaslMechanismHandlerSharedWithDefaults instantiates a new OauthBearerSaslMechanismHandlerShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthBearerSaslMechanismHandlerSharedWithDefaults() *OauthBearerSaslMechanismHandlerShared {
	this := OauthBearerSaslMechanismHandlerShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *OauthBearerSaslMechanismHandlerShared) GetSchemas() []EnumoauthBearerSaslMechanismHandlerSchemaUrn {
	if o == nil {
		var ret []EnumoauthBearerSaslMechanismHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetSchemasOk() ([]EnumoauthBearerSaslMechanismHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *OauthBearerSaslMechanismHandlerShared) SetSchemas(v []EnumoauthBearerSaslMechanismHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAccessTokenValidator returns the AccessTokenValidator field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetAccessTokenValidator() []string {
	if o == nil || isNil(o.AccessTokenValidator) {
		var ret []string
		return ret
	}
	return o.AccessTokenValidator
}

// GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetAccessTokenValidatorOk() ([]string, bool) {
	if o == nil || isNil(o.AccessTokenValidator) {
    return nil, false
	}
	return o.AccessTokenValidator, true
}

// HasAccessTokenValidator returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasAccessTokenValidator() bool {
	if o != nil && !isNil(o.AccessTokenValidator) {
		return true
	}

	return false
}

// SetAccessTokenValidator gets a reference to the given []string and assigns it to the AccessTokenValidator field.
func (o *OauthBearerSaslMechanismHandlerShared) SetAccessTokenValidator(v []string) {
	o.AccessTokenValidator = v
}

// GetIdTokenValidator returns the IdTokenValidator field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetIdTokenValidator() []string {
	if o == nil || isNil(o.IdTokenValidator) {
		var ret []string
		return ret
	}
	return o.IdTokenValidator
}

// GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetIdTokenValidatorOk() ([]string, bool) {
	if o == nil || isNil(o.IdTokenValidator) {
    return nil, false
	}
	return o.IdTokenValidator, true
}

// HasIdTokenValidator returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasIdTokenValidator() bool {
	if o != nil && !isNil(o.IdTokenValidator) {
		return true
	}

	return false
}

// SetIdTokenValidator gets a reference to the given []string and assigns it to the IdTokenValidator field.
func (o *OauthBearerSaslMechanismHandlerShared) SetIdTokenValidator(v []string) {
	o.IdTokenValidator = v
}

// GetRequireBothAccessTokenAndIDToken returns the RequireBothAccessTokenAndIDToken field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetRequireBothAccessTokenAndIDToken() bool {
	if o == nil || isNil(o.RequireBothAccessTokenAndIDToken) {
		var ret bool
		return ret
	}
	return *o.RequireBothAccessTokenAndIDToken
}

// GetRequireBothAccessTokenAndIDTokenOk returns a tuple with the RequireBothAccessTokenAndIDToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetRequireBothAccessTokenAndIDTokenOk() (*bool, bool) {
	if o == nil || isNil(o.RequireBothAccessTokenAndIDToken) {
    return nil, false
	}
	return o.RequireBothAccessTokenAndIDToken, true
}

// HasRequireBothAccessTokenAndIDToken returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasRequireBothAccessTokenAndIDToken() bool {
	if o != nil && !isNil(o.RequireBothAccessTokenAndIDToken) {
		return true
	}

	return false
}

// SetRequireBothAccessTokenAndIDToken gets a reference to the given bool and assigns it to the RequireBothAccessTokenAndIDToken field.
func (o *OauthBearerSaslMechanismHandlerShared) SetRequireBothAccessTokenAndIDToken(v bool) {
	o.RequireBothAccessTokenAndIDToken = &v
}

// GetValidateAccessTokenWhenIDTokenIsAlsoProvided returns the ValidateAccessTokenWhenIDTokenIsAlsoProvided field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetValidateAccessTokenWhenIDTokenIsAlsoProvided() EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp {
	if o == nil || isNil(o.ValidateAccessTokenWhenIDTokenIsAlsoProvided) {
		var ret EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp
		return ret
	}
	return *o.ValidateAccessTokenWhenIDTokenIsAlsoProvided
}

// GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk returns a tuple with the ValidateAccessTokenWhenIDTokenIsAlsoProvided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk() (*EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp, bool) {
	if o == nil || isNil(o.ValidateAccessTokenWhenIDTokenIsAlsoProvided) {
    return nil, false
	}
	return o.ValidateAccessTokenWhenIDTokenIsAlsoProvided, true
}

// HasValidateAccessTokenWhenIDTokenIsAlsoProvided returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasValidateAccessTokenWhenIDTokenIsAlsoProvided() bool {
	if o != nil && !isNil(o.ValidateAccessTokenWhenIDTokenIsAlsoProvided) {
		return true
	}

	return false
}

// SetValidateAccessTokenWhenIDTokenIsAlsoProvided gets a reference to the given EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp and assigns it to the ValidateAccessTokenWhenIDTokenIsAlsoProvided field.
func (o *OauthBearerSaslMechanismHandlerShared) SetValidateAccessTokenWhenIDTokenIsAlsoProvided(v EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp) {
	o.ValidateAccessTokenWhenIDTokenIsAlsoProvided = &v
}

// GetAlternateAuthorizationIdentityMapper returns the AlternateAuthorizationIdentityMapper field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetAlternateAuthorizationIdentityMapper() string {
	if o == nil || isNil(o.AlternateAuthorizationIdentityMapper) {
		var ret string
		return ret
	}
	return *o.AlternateAuthorizationIdentityMapper
}

// GetAlternateAuthorizationIdentityMapperOk returns a tuple with the AlternateAuthorizationIdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetAlternateAuthorizationIdentityMapperOk() (*string, bool) {
	if o == nil || isNil(o.AlternateAuthorizationIdentityMapper) {
    return nil, false
	}
	return o.AlternateAuthorizationIdentityMapper, true
}

// HasAlternateAuthorizationIdentityMapper returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasAlternateAuthorizationIdentityMapper() bool {
	if o != nil && !isNil(o.AlternateAuthorizationIdentityMapper) {
		return true
	}

	return false
}

// SetAlternateAuthorizationIdentityMapper gets a reference to the given string and assigns it to the AlternateAuthorizationIdentityMapper field.
func (o *OauthBearerSaslMechanismHandlerShared) SetAlternateAuthorizationIdentityMapper(v string) {
	o.AlternateAuthorizationIdentityMapper = &v
}

// GetAllRequiredScope returns the AllRequiredScope field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetAllRequiredScope() []string {
	if o == nil || isNil(o.AllRequiredScope) {
		var ret []string
		return ret
	}
	return o.AllRequiredScope
}

// GetAllRequiredScopeOk returns a tuple with the AllRequiredScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetAllRequiredScopeOk() ([]string, bool) {
	if o == nil || isNil(o.AllRequiredScope) {
    return nil, false
	}
	return o.AllRequiredScope, true
}

// HasAllRequiredScope returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasAllRequiredScope() bool {
	if o != nil && !isNil(o.AllRequiredScope) {
		return true
	}

	return false
}

// SetAllRequiredScope gets a reference to the given []string and assigns it to the AllRequiredScope field.
func (o *OauthBearerSaslMechanismHandlerShared) SetAllRequiredScope(v []string) {
	o.AllRequiredScope = v
}

// GetAnyRequiredScope returns the AnyRequiredScope field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetAnyRequiredScope() []string {
	if o == nil || isNil(o.AnyRequiredScope) {
		var ret []string
		return ret
	}
	return o.AnyRequiredScope
}

// GetAnyRequiredScopeOk returns a tuple with the AnyRequiredScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetAnyRequiredScopeOk() ([]string, bool) {
	if o == nil || isNil(o.AnyRequiredScope) {
    return nil, false
	}
	return o.AnyRequiredScope, true
}

// HasAnyRequiredScope returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasAnyRequiredScope() bool {
	if o != nil && !isNil(o.AnyRequiredScope) {
		return true
	}

	return false
}

// SetAnyRequiredScope gets a reference to the given []string and assigns it to the AnyRequiredScope field.
func (o *OauthBearerSaslMechanismHandlerShared) SetAnyRequiredScope(v []string) {
	o.AnyRequiredScope = v
}

// GetServerFqdn returns the ServerFqdn field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetServerFqdn() string {
	if o == nil || isNil(o.ServerFqdn) {
		var ret string
		return ret
	}
	return *o.ServerFqdn
}

// GetServerFqdnOk returns a tuple with the ServerFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetServerFqdnOk() (*string, bool) {
	if o == nil || isNil(o.ServerFqdn) {
    return nil, false
	}
	return o.ServerFqdn, true
}

// HasServerFqdn returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasServerFqdn() bool {
	if o != nil && !isNil(o.ServerFqdn) {
		return true
	}

	return false
}

// SetServerFqdn gets a reference to the given string and assigns it to the ServerFqdn field.
func (o *OauthBearerSaslMechanismHandlerShared) SetServerFqdn(v string) {
	o.ServerFqdn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OauthBearerSaslMechanismHandlerShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OauthBearerSaslMechanismHandlerShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OauthBearerSaslMechanismHandlerShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *OauthBearerSaslMechanismHandlerShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *OauthBearerSaslMechanismHandlerShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *OauthBearerSaslMechanismHandlerShared) SetEnabled(v bool) {
	o.Enabled = v
}

func (o OauthBearerSaslMechanismHandlerShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.AccessTokenValidator) {
		toSerialize["accessTokenValidator"] = o.AccessTokenValidator
	}
	if !isNil(o.IdTokenValidator) {
		toSerialize["idTokenValidator"] = o.IdTokenValidator
	}
	if !isNil(o.RequireBothAccessTokenAndIDToken) {
		toSerialize["requireBothAccessTokenAndIDToken"] = o.RequireBothAccessTokenAndIDToken
	}
	if !isNil(o.ValidateAccessTokenWhenIDTokenIsAlsoProvided) {
		toSerialize["validateAccessTokenWhenIDTokenIsAlsoProvided"] = o.ValidateAccessTokenWhenIDTokenIsAlsoProvided
	}
	if !isNil(o.AlternateAuthorizationIdentityMapper) {
		toSerialize["alternateAuthorizationIdentityMapper"] = o.AlternateAuthorizationIdentityMapper
	}
	if !isNil(o.AllRequiredScope) {
		toSerialize["allRequiredScope"] = o.AllRequiredScope
	}
	if !isNil(o.AnyRequiredScope) {
		toSerialize["anyRequiredScope"] = o.AnyRequiredScope
	}
	if !isNil(o.ServerFqdn) {
		toSerialize["serverFqdn"] = o.ServerFqdn
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableOauthBearerSaslMechanismHandlerShared struct {
	value *OauthBearerSaslMechanismHandlerShared
	isSet bool
}

func (v NullableOauthBearerSaslMechanismHandlerShared) Get() *OauthBearerSaslMechanismHandlerShared {
	return v.value
}

func (v *NullableOauthBearerSaslMechanismHandlerShared) Set(val *OauthBearerSaslMechanismHandlerShared) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthBearerSaslMechanismHandlerShared) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthBearerSaslMechanismHandlerShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthBearerSaslMechanismHandlerShared(val *OauthBearerSaslMechanismHandlerShared) *NullableOauthBearerSaslMechanismHandlerShared {
	return &NullableOauthBearerSaslMechanismHandlerShared{value: val, isSet: true}
}

func (v NullableOauthBearerSaslMechanismHandlerShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthBearerSaslMechanismHandlerShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


