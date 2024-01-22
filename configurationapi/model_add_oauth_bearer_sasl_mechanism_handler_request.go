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

// checks if the AddOauthBearerSaslMechanismHandlerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOauthBearerSaslMechanismHandlerRequest{}

// AddOauthBearerSaslMechanismHandlerRequest struct for AddOauthBearerSaslMechanismHandlerRequest
type AddOauthBearerSaslMechanismHandlerRequest struct {
	Schemas []EnumoauthBearerSaslMechanismHandlerSchemaUrn `json:"schemas"`
	// An access token validator that will ensure that each presented OAuth access token is authentic and trustworthy. It must be configured with an identity mapper that will be used to map the access token to a local entry.
	AccessTokenValidator []string `json:"accessTokenValidator,omitempty"`
	// An ID token validator that will ensure that each presented OpenID Connect ID token is authentic and trustworthy, and that will map the token to a local entry.
	IdTokenValidator []string `json:"idTokenValidator,omitempty"`
	// Indicates whether bind requests will be required to have both an OAuth access token (in the \"auth\" element of the bind request) and an OpenID Connect ID token (in the \"pingidentityidtoken\" element of the bind request).
	RequireBothAccessTokenAndIDToken             *bool                                                                     `json:"requireBothAccessTokenAndIDToken,omitempty"`
	ValidateAccessTokenWhenIDTokenIsAlsoProvided *EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp `json:"validateAccessTokenWhenIDTokenIsAlsoProvided,omitempty"`
	// The identity mapper that will be used to map an alternate authorization identity (provided in the GS2 header of the encoded OAUTHBEARER bind request credentials) to the corresponding local entry.
	AlternateAuthorizationIdentityMapper *string `json:"alternateAuthorizationIdentityMapper,omitempty"`
	// The set of OAuth scopes that will all be required for any access tokens that will be allowed for authentication.
	AllRequiredScope []string `json:"allRequiredScope,omitempty"`
	// The set of OAuth scopes that a token may have to be allowed for authentication.
	AnyRequiredScope []string `json:"anyRequiredScope,omitempty"`
	// The fully-qualified name that clients are expected to use when communicating with the server.
	ServerFqdn *string `json:"serverFqdn,omitempty"`
	// A description for this SASL Mechanism Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the SASL mechanism handler is enabled for use.
	Enabled bool `json:"enabled"`
	// Name of the new SASL Mechanism Handler
	HandlerName string `json:"handlerName"`
}

// NewAddOauthBearerSaslMechanismHandlerRequest instantiates a new AddOauthBearerSaslMechanismHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOauthBearerSaslMechanismHandlerRequest(schemas []EnumoauthBearerSaslMechanismHandlerSchemaUrn, enabled bool, handlerName string) *AddOauthBearerSaslMechanismHandlerRequest {
	this := AddOauthBearerSaslMechanismHandlerRequest{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.HandlerName = handlerName
	return &this
}

// NewAddOauthBearerSaslMechanismHandlerRequestWithDefaults instantiates a new AddOauthBearerSaslMechanismHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOauthBearerSaslMechanismHandlerRequestWithDefaults() *AddOauthBearerSaslMechanismHandlerRequest {
	this := AddOauthBearerSaslMechanismHandlerRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetSchemas() []EnumoauthBearerSaslMechanismHandlerSchemaUrn {
	if o == nil {
		var ret []EnumoauthBearerSaslMechanismHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetSchemasOk() ([]EnumoauthBearerSaslMechanismHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetSchemas(v []EnumoauthBearerSaslMechanismHandlerSchemaUrn) {
	o.Schemas = v
}

// GetAccessTokenValidator returns the AccessTokenValidator field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAccessTokenValidator() []string {
	if o == nil || IsNil(o.AccessTokenValidator) {
		var ret []string
		return ret
	}
	return o.AccessTokenValidator
}

// GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAccessTokenValidatorOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessTokenValidator) {
		return nil, false
	}
	return o.AccessTokenValidator, true
}

// HasAccessTokenValidator returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasAccessTokenValidator() bool {
	if o != nil && !IsNil(o.AccessTokenValidator) {
		return true
	}

	return false
}

// SetAccessTokenValidator gets a reference to the given []string and assigns it to the AccessTokenValidator field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetAccessTokenValidator(v []string) {
	o.AccessTokenValidator = v
}

// GetIdTokenValidator returns the IdTokenValidator field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetIdTokenValidator() []string {
	if o == nil || IsNil(o.IdTokenValidator) {
		var ret []string
		return ret
	}
	return o.IdTokenValidator
}

// GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetIdTokenValidatorOk() ([]string, bool) {
	if o == nil || IsNil(o.IdTokenValidator) {
		return nil, false
	}
	return o.IdTokenValidator, true
}

// HasIdTokenValidator returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasIdTokenValidator() bool {
	if o != nil && !IsNil(o.IdTokenValidator) {
		return true
	}

	return false
}

// SetIdTokenValidator gets a reference to the given []string and assigns it to the IdTokenValidator field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetIdTokenValidator(v []string) {
	o.IdTokenValidator = v
}

// GetRequireBothAccessTokenAndIDToken returns the RequireBothAccessTokenAndIDToken field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetRequireBothAccessTokenAndIDToken() bool {
	if o == nil || IsNil(o.RequireBothAccessTokenAndIDToken) {
		var ret bool
		return ret
	}
	return *o.RequireBothAccessTokenAndIDToken
}

// GetRequireBothAccessTokenAndIDTokenOk returns a tuple with the RequireBothAccessTokenAndIDToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetRequireBothAccessTokenAndIDTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireBothAccessTokenAndIDToken) {
		return nil, false
	}
	return o.RequireBothAccessTokenAndIDToken, true
}

// HasRequireBothAccessTokenAndIDToken returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasRequireBothAccessTokenAndIDToken() bool {
	if o != nil && !IsNil(o.RequireBothAccessTokenAndIDToken) {
		return true
	}

	return false
}

// SetRequireBothAccessTokenAndIDToken gets a reference to the given bool and assigns it to the RequireBothAccessTokenAndIDToken field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetRequireBothAccessTokenAndIDToken(v bool) {
	o.RequireBothAccessTokenAndIDToken = &v
}

// GetValidateAccessTokenWhenIDTokenIsAlsoProvided returns the ValidateAccessTokenWhenIDTokenIsAlsoProvided field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetValidateAccessTokenWhenIDTokenIsAlsoProvided() EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp {
	if o == nil || IsNil(o.ValidateAccessTokenWhenIDTokenIsAlsoProvided) {
		var ret EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp
		return ret
	}
	return *o.ValidateAccessTokenWhenIDTokenIsAlsoProvided
}

// GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk returns a tuple with the ValidateAccessTokenWhenIDTokenIsAlsoProvided field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk() (*EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp, bool) {
	if o == nil || IsNil(o.ValidateAccessTokenWhenIDTokenIsAlsoProvided) {
		return nil, false
	}
	return o.ValidateAccessTokenWhenIDTokenIsAlsoProvided, true
}

// HasValidateAccessTokenWhenIDTokenIsAlsoProvided returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasValidateAccessTokenWhenIDTokenIsAlsoProvided() bool {
	if o != nil && !IsNil(o.ValidateAccessTokenWhenIDTokenIsAlsoProvided) {
		return true
	}

	return false
}

// SetValidateAccessTokenWhenIDTokenIsAlsoProvided gets a reference to the given EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp and assigns it to the ValidateAccessTokenWhenIDTokenIsAlsoProvided field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetValidateAccessTokenWhenIDTokenIsAlsoProvided(v EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp) {
	o.ValidateAccessTokenWhenIDTokenIsAlsoProvided = &v
}

// GetAlternateAuthorizationIdentityMapper returns the AlternateAuthorizationIdentityMapper field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAlternateAuthorizationIdentityMapper() string {
	if o == nil || IsNil(o.AlternateAuthorizationIdentityMapper) {
		var ret string
		return ret
	}
	return *o.AlternateAuthorizationIdentityMapper
}

// GetAlternateAuthorizationIdentityMapperOk returns a tuple with the AlternateAuthorizationIdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAlternateAuthorizationIdentityMapperOk() (*string, bool) {
	if o == nil || IsNil(o.AlternateAuthorizationIdentityMapper) {
		return nil, false
	}
	return o.AlternateAuthorizationIdentityMapper, true
}

// HasAlternateAuthorizationIdentityMapper returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasAlternateAuthorizationIdentityMapper() bool {
	if o != nil && !IsNil(o.AlternateAuthorizationIdentityMapper) {
		return true
	}

	return false
}

// SetAlternateAuthorizationIdentityMapper gets a reference to the given string and assigns it to the AlternateAuthorizationIdentityMapper field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetAlternateAuthorizationIdentityMapper(v string) {
	o.AlternateAuthorizationIdentityMapper = &v
}

// GetAllRequiredScope returns the AllRequiredScope field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAllRequiredScope() []string {
	if o == nil || IsNil(o.AllRequiredScope) {
		var ret []string
		return ret
	}
	return o.AllRequiredScope
}

// GetAllRequiredScopeOk returns a tuple with the AllRequiredScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAllRequiredScopeOk() ([]string, bool) {
	if o == nil || IsNil(o.AllRequiredScope) {
		return nil, false
	}
	return o.AllRequiredScope, true
}

// HasAllRequiredScope returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasAllRequiredScope() bool {
	if o != nil && !IsNil(o.AllRequiredScope) {
		return true
	}

	return false
}

// SetAllRequiredScope gets a reference to the given []string and assigns it to the AllRequiredScope field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetAllRequiredScope(v []string) {
	o.AllRequiredScope = v
}

// GetAnyRequiredScope returns the AnyRequiredScope field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAnyRequiredScope() []string {
	if o == nil || IsNil(o.AnyRequiredScope) {
		var ret []string
		return ret
	}
	return o.AnyRequiredScope
}

// GetAnyRequiredScopeOk returns a tuple with the AnyRequiredScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAnyRequiredScopeOk() ([]string, bool) {
	if o == nil || IsNil(o.AnyRequiredScope) {
		return nil, false
	}
	return o.AnyRequiredScope, true
}

// HasAnyRequiredScope returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasAnyRequiredScope() bool {
	if o != nil && !IsNil(o.AnyRequiredScope) {
		return true
	}

	return false
}

// SetAnyRequiredScope gets a reference to the given []string and assigns it to the AnyRequiredScope field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetAnyRequiredScope(v []string) {
	o.AnyRequiredScope = v
}

// GetServerFqdn returns the ServerFqdn field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetServerFqdn() string {
	if o == nil || IsNil(o.ServerFqdn) {
		var ret string
		return ret
	}
	return *o.ServerFqdn
}

// GetServerFqdnOk returns a tuple with the ServerFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetServerFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ServerFqdn) {
		return nil, false
	}
	return o.ServerFqdn, true
}

// HasServerFqdn returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasServerFqdn() bool {
	if o != nil && !IsNil(o.ServerFqdn) {
		return true
	}

	return false
}

// SetServerFqdn gets a reference to the given string and assigns it to the ServerFqdn field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetServerFqdn(v string) {
	o.ServerFqdn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHandlerName returns the HandlerName field value
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddOauthBearerSaslMechanismHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddOauthBearerSaslMechanismHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

func (o AddOauthBearerSaslMechanismHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOauthBearerSaslMechanismHandlerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.AccessTokenValidator) {
		toSerialize["accessTokenValidator"] = o.AccessTokenValidator
	}
	if !IsNil(o.IdTokenValidator) {
		toSerialize["idTokenValidator"] = o.IdTokenValidator
	}
	if !IsNil(o.RequireBothAccessTokenAndIDToken) {
		toSerialize["requireBothAccessTokenAndIDToken"] = o.RequireBothAccessTokenAndIDToken
	}
	if !IsNil(o.ValidateAccessTokenWhenIDTokenIsAlsoProvided) {
		toSerialize["validateAccessTokenWhenIDTokenIsAlsoProvided"] = o.ValidateAccessTokenWhenIDTokenIsAlsoProvided
	}
	if !IsNil(o.AlternateAuthorizationIdentityMapper) {
		toSerialize["alternateAuthorizationIdentityMapper"] = o.AlternateAuthorizationIdentityMapper
	}
	if !IsNil(o.AllRequiredScope) {
		toSerialize["allRequiredScope"] = o.AllRequiredScope
	}
	if !IsNil(o.AnyRequiredScope) {
		toSerialize["anyRequiredScope"] = o.AnyRequiredScope
	}
	if !IsNil(o.ServerFqdn) {
		toSerialize["serverFqdn"] = o.ServerFqdn
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["handlerName"] = o.HandlerName
	return toSerialize, nil
}

type NullableAddOauthBearerSaslMechanismHandlerRequest struct {
	value *AddOauthBearerSaslMechanismHandlerRequest
	isSet bool
}

func (v NullableAddOauthBearerSaslMechanismHandlerRequest) Get() *AddOauthBearerSaslMechanismHandlerRequest {
	return v.value
}

func (v *NullableAddOauthBearerSaslMechanismHandlerRequest) Set(val *AddOauthBearerSaslMechanismHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOauthBearerSaslMechanismHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOauthBearerSaslMechanismHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOauthBearerSaslMechanismHandlerRequest(val *AddOauthBearerSaslMechanismHandlerRequest) *NullableAddOauthBearerSaslMechanismHandlerRequest {
	return &NullableAddOauthBearerSaslMechanismHandlerRequest{value: val, isSet: true}
}

func (v NullableAddOauthBearerSaslMechanismHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOauthBearerSaslMechanismHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
