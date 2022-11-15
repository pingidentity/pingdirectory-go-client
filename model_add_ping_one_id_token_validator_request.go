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

// AddPingOneIdTokenValidatorRequest struct for AddPingOneIdTokenValidatorRequest
type AddPingOneIdTokenValidatorRequest struct {
	// Name of the new ID Token Validator
	ValidatorName string `json:"validatorName"`
	Schemas []EnumpingOneIdTokenValidatorSchemaUrn `json:"schemas"`
	// Specifies a PingOne base issuer URL.
	IssuerURL string `json:"issuerURL"`
	// Specifies HTTPS connection settings for the PingOne OpenID Connect provider.
	OpenIDConnectProvider string `json:"OpenIDConnectProvider"`
	// How often the PingOne ID Token Validator should refresh its stored cache of OpenID Connect-related metadata.
	OpenIDConnectMetadataCacheDuration *string `json:"OpenIDConnectMetadataCacheDuration,omitempty"`
	// A description for this ID Token Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether this ID Token Validator is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
	// Specifies the name of the Identity Mapper that should be used to correlate an ID token subject value to a user entry. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property.
	IdentityMapper string `json:"identityMapper"`
	// The name of the token claim that contains the subject; i.e., the authenticated user.
	SubjectClaimName *string `json:"subjectClaimName,omitempty"`
	// Specifies the amount of clock skew that is tolerated by the ID Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token's not-before (nbf) time and added to the token's expiration (exp) time, if present, to allow for any time difference between the local server's clock and the token issuer's clock.
	ClockSkewGracePeriod *string `json:"clockSkewGracePeriod,omitempty"`
	// How often the ID Token Validator should refresh its cache of JWKS token signing keys.
	JwksCacheDuration *string `json:"jwksCacheDuration,omitempty"`
	// When multiple ID Token Validators are defined for a single Directory Server, this property determines the order in which the ID Token Validators are consulted. Values of this property must be unique among all ID Token Validators defined within Directory Server but not necessarily contiguous. ID Token Validators with lower values will be evaluated first to determine if they are able to validate the ID token.
	EvaluationOrderIndex int32 `json:"evaluationOrderIndex"`
}

// NewAddPingOneIdTokenValidatorRequest instantiates a new AddPingOneIdTokenValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPingOneIdTokenValidatorRequest(validatorName string, schemas []EnumpingOneIdTokenValidatorSchemaUrn, issuerURL string, openIDConnectProvider string, enabled bool, identityMapper string, evaluationOrderIndex int32) *AddPingOneIdTokenValidatorRequest {
	this := AddPingOneIdTokenValidatorRequest{}
	this.ValidatorName = validatorName
	this.Schemas = schemas
	this.IssuerURL = issuerURL
	this.OpenIDConnectProvider = openIDConnectProvider
	this.Enabled = enabled
	this.IdentityMapper = identityMapper
	this.EvaluationOrderIndex = evaluationOrderIndex
	return &this
}

// NewAddPingOneIdTokenValidatorRequestWithDefaults instantiates a new AddPingOneIdTokenValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPingOneIdTokenValidatorRequestWithDefaults() *AddPingOneIdTokenValidatorRequest {
	this := AddPingOneIdTokenValidatorRequest{}
	return &this
}

// GetValidatorName returns the ValidatorName field value
func (o *AddPingOneIdTokenValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddPingOneIdTokenValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPingOneIdTokenValidatorRequest) GetSchemas() []EnumpingOneIdTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnumpingOneIdTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetSchemasOk() ([]EnumpingOneIdTokenValidatorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPingOneIdTokenValidatorRequest) SetSchemas(v []EnumpingOneIdTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetIssuerURL returns the IssuerURL field value
func (o *AddPingOneIdTokenValidatorRequest) GetIssuerURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerURL
}

// GetIssuerURLOk returns a tuple with the IssuerURL field value
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetIssuerURLOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IssuerURL, true
}

// SetIssuerURL sets field value
func (o *AddPingOneIdTokenValidatorRequest) SetIssuerURL(v string) {
	o.IssuerURL = v
}

// GetOpenIDConnectProvider returns the OpenIDConnectProvider field value
func (o *AddPingOneIdTokenValidatorRequest) GetOpenIDConnectProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenIDConnectProvider
}

// GetOpenIDConnectProviderOk returns a tuple with the OpenIDConnectProvider field value
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetOpenIDConnectProviderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OpenIDConnectProvider, true
}

// SetOpenIDConnectProvider sets field value
func (o *AddPingOneIdTokenValidatorRequest) SetOpenIDConnectProvider(v string) {
	o.OpenIDConnectProvider = v
}

// GetOpenIDConnectMetadataCacheDuration returns the OpenIDConnectMetadataCacheDuration field value if set, zero value otherwise.
func (o *AddPingOneIdTokenValidatorRequest) GetOpenIDConnectMetadataCacheDuration() string {
	if o == nil || isNil(o.OpenIDConnectMetadataCacheDuration) {
		var ret string
		return ret
	}
	return *o.OpenIDConnectMetadataCacheDuration
}

// GetOpenIDConnectMetadataCacheDurationOk returns a tuple with the OpenIDConnectMetadataCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetOpenIDConnectMetadataCacheDurationOk() (*string, bool) {
	if o == nil || isNil(o.OpenIDConnectMetadataCacheDuration) {
    return nil, false
	}
	return o.OpenIDConnectMetadataCacheDuration, true
}

// HasOpenIDConnectMetadataCacheDuration returns a boolean if a field has been set.
func (o *AddPingOneIdTokenValidatorRequest) HasOpenIDConnectMetadataCacheDuration() bool {
	if o != nil && !isNil(o.OpenIDConnectMetadataCacheDuration) {
		return true
	}

	return false
}

// SetOpenIDConnectMetadataCacheDuration gets a reference to the given string and assigns it to the OpenIDConnectMetadataCacheDuration field.
func (o *AddPingOneIdTokenValidatorRequest) SetOpenIDConnectMetadataCacheDuration(v string) {
	o.OpenIDConnectMetadataCacheDuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPingOneIdTokenValidatorRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPingOneIdTokenValidatorRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPingOneIdTokenValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPingOneIdTokenValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPingOneIdTokenValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *AddPingOneIdTokenValidatorRequest) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *AddPingOneIdTokenValidatorRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetSubjectClaimName returns the SubjectClaimName field value if set, zero value otherwise.
func (o *AddPingOneIdTokenValidatorRequest) GetSubjectClaimName() string {
	if o == nil || isNil(o.SubjectClaimName) {
		var ret string
		return ret
	}
	return *o.SubjectClaimName
}

// GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetSubjectClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectClaimName) {
    return nil, false
	}
	return o.SubjectClaimName, true
}

// HasSubjectClaimName returns a boolean if a field has been set.
func (o *AddPingOneIdTokenValidatorRequest) HasSubjectClaimName() bool {
	if o != nil && !isNil(o.SubjectClaimName) {
		return true
	}

	return false
}

// SetSubjectClaimName gets a reference to the given string and assigns it to the SubjectClaimName field.
func (o *AddPingOneIdTokenValidatorRequest) SetSubjectClaimName(v string) {
	o.SubjectClaimName = &v
}

// GetClockSkewGracePeriod returns the ClockSkewGracePeriod field value if set, zero value otherwise.
func (o *AddPingOneIdTokenValidatorRequest) GetClockSkewGracePeriod() string {
	if o == nil || isNil(o.ClockSkewGracePeriod) {
		var ret string
		return ret
	}
	return *o.ClockSkewGracePeriod
}

// GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetClockSkewGracePeriodOk() (*string, bool) {
	if o == nil || isNil(o.ClockSkewGracePeriod) {
    return nil, false
	}
	return o.ClockSkewGracePeriod, true
}

// HasClockSkewGracePeriod returns a boolean if a field has been set.
func (o *AddPingOneIdTokenValidatorRequest) HasClockSkewGracePeriod() bool {
	if o != nil && !isNil(o.ClockSkewGracePeriod) {
		return true
	}

	return false
}

// SetClockSkewGracePeriod gets a reference to the given string and assigns it to the ClockSkewGracePeriod field.
func (o *AddPingOneIdTokenValidatorRequest) SetClockSkewGracePeriod(v string) {
	o.ClockSkewGracePeriod = &v
}

// GetJwksCacheDuration returns the JwksCacheDuration field value if set, zero value otherwise.
func (o *AddPingOneIdTokenValidatorRequest) GetJwksCacheDuration() string {
	if o == nil || isNil(o.JwksCacheDuration) {
		var ret string
		return ret
	}
	return *o.JwksCacheDuration
}

// GetJwksCacheDurationOk returns a tuple with the JwksCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetJwksCacheDurationOk() (*string, bool) {
	if o == nil || isNil(o.JwksCacheDuration) {
    return nil, false
	}
	return o.JwksCacheDuration, true
}

// HasJwksCacheDuration returns a boolean if a field has been set.
func (o *AddPingOneIdTokenValidatorRequest) HasJwksCacheDuration() bool {
	if o != nil && !isNil(o.JwksCacheDuration) {
		return true
	}

	return false
}

// SetJwksCacheDuration gets a reference to the given string and assigns it to the JwksCacheDuration field.
func (o *AddPingOneIdTokenValidatorRequest) SetJwksCacheDuration(v string) {
	o.JwksCacheDuration = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *AddPingOneIdTokenValidatorRequest) GetEvaluationOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *AddPingOneIdTokenValidatorRequest) GetEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *AddPingOneIdTokenValidatorRequest) SetEvaluationOrderIndex(v int32) {
	o.EvaluationOrderIndex = v
}

func (o AddPingOneIdTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["validatorName"] = o.ValidatorName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["issuerURL"] = o.IssuerURL
	}
	if true {
		toSerialize["OpenIDConnectProvider"] = o.OpenIDConnectProvider
	}
	if !isNil(o.OpenIDConnectMetadataCacheDuration) {
		toSerialize["OpenIDConnectMetadataCacheDuration"] = o.OpenIDConnectMetadataCacheDuration
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !isNil(o.SubjectClaimName) {
		toSerialize["subjectClaimName"] = o.SubjectClaimName
	}
	if !isNil(o.ClockSkewGracePeriod) {
		toSerialize["clockSkewGracePeriod"] = o.ClockSkewGracePeriod
	}
	if !isNil(o.JwksCacheDuration) {
		toSerialize["jwksCacheDuration"] = o.JwksCacheDuration
	}
	if true {
		toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	}
	return json.Marshal(toSerialize)
}

type NullableAddPingOneIdTokenValidatorRequest struct {
	value *AddPingOneIdTokenValidatorRequest
	isSet bool
}

func (v NullableAddPingOneIdTokenValidatorRequest) Get() *AddPingOneIdTokenValidatorRequest {
	return v.value
}

func (v *NullableAddPingOneIdTokenValidatorRequest) Set(val *AddPingOneIdTokenValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPingOneIdTokenValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPingOneIdTokenValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPingOneIdTokenValidatorRequest(val *AddPingOneIdTokenValidatorRequest) *NullableAddPingOneIdTokenValidatorRequest {
	return &NullableAddPingOneIdTokenValidatorRequest{value: val, isSet: true}
}

func (v NullableAddPingOneIdTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPingOneIdTokenValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


