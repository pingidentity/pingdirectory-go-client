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

// PingOneIdTokenValidatorResponse struct for PingOneIdTokenValidatorResponse
type PingOneIdTokenValidatorResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the ID Token Validator
	Id string `json:"id"`
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

// NewPingOneIdTokenValidatorResponse instantiates a new PingOneIdTokenValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingOneIdTokenValidatorResponse(id string, schemas []EnumpingOneIdTokenValidatorSchemaUrn, issuerURL string, openIDConnectProvider string, enabled bool, identityMapper string, evaluationOrderIndex int32) *PingOneIdTokenValidatorResponse {
	this := PingOneIdTokenValidatorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.IssuerURL = issuerURL
	this.OpenIDConnectProvider = openIDConnectProvider
	this.Enabled = enabled
	this.IdentityMapper = identityMapper
	this.EvaluationOrderIndex = evaluationOrderIndex
	return &this
}

// NewPingOneIdTokenValidatorResponseWithDefaults instantiates a new PingOneIdTokenValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingOneIdTokenValidatorResponseWithDefaults() *PingOneIdTokenValidatorResponse {
	this := PingOneIdTokenValidatorResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PingOneIdTokenValidatorResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PingOneIdTokenValidatorResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *PingOneIdTokenValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *PingOneIdTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *PingOneIdTokenValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *PingOneIdTokenValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *PingOneIdTokenValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PingOneIdTokenValidatorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *PingOneIdTokenValidatorResponse) GetSchemas() []EnumpingOneIdTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnumpingOneIdTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetSchemasOk() ([]EnumpingOneIdTokenValidatorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *PingOneIdTokenValidatorResponse) SetSchemas(v []EnumpingOneIdTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetIssuerURL returns the IssuerURL field value
func (o *PingOneIdTokenValidatorResponse) GetIssuerURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerURL
}

// GetIssuerURLOk returns a tuple with the IssuerURL field value
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetIssuerURLOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IssuerURL, true
}

// SetIssuerURL sets field value
func (o *PingOneIdTokenValidatorResponse) SetIssuerURL(v string) {
	o.IssuerURL = v
}

// GetOpenIDConnectProvider returns the OpenIDConnectProvider field value
func (o *PingOneIdTokenValidatorResponse) GetOpenIDConnectProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenIDConnectProvider
}

// GetOpenIDConnectProviderOk returns a tuple with the OpenIDConnectProvider field value
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetOpenIDConnectProviderOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.OpenIDConnectProvider, true
}

// SetOpenIDConnectProvider sets field value
func (o *PingOneIdTokenValidatorResponse) SetOpenIDConnectProvider(v string) {
	o.OpenIDConnectProvider = v
}

// GetOpenIDConnectMetadataCacheDuration returns the OpenIDConnectMetadataCacheDuration field value if set, zero value otherwise.
func (o *PingOneIdTokenValidatorResponse) GetOpenIDConnectMetadataCacheDuration() string {
	if o == nil || isNil(o.OpenIDConnectMetadataCacheDuration) {
		var ret string
		return ret
	}
	return *o.OpenIDConnectMetadataCacheDuration
}

// GetOpenIDConnectMetadataCacheDurationOk returns a tuple with the OpenIDConnectMetadataCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetOpenIDConnectMetadataCacheDurationOk() (*string, bool) {
	if o == nil || isNil(o.OpenIDConnectMetadataCacheDuration) {
    return nil, false
	}
	return o.OpenIDConnectMetadataCacheDuration, true
}

// HasOpenIDConnectMetadataCacheDuration returns a boolean if a field has been set.
func (o *PingOneIdTokenValidatorResponse) HasOpenIDConnectMetadataCacheDuration() bool {
	if o != nil && !isNil(o.OpenIDConnectMetadataCacheDuration) {
		return true
	}

	return false
}

// SetOpenIDConnectMetadataCacheDuration gets a reference to the given string and assigns it to the OpenIDConnectMetadataCacheDuration field.
func (o *PingOneIdTokenValidatorResponse) SetOpenIDConnectMetadataCacheDuration(v string) {
	o.OpenIDConnectMetadataCacheDuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PingOneIdTokenValidatorResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PingOneIdTokenValidatorResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PingOneIdTokenValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *PingOneIdTokenValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PingOneIdTokenValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *PingOneIdTokenValidatorResponse) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *PingOneIdTokenValidatorResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetSubjectClaimName returns the SubjectClaimName field value if set, zero value otherwise.
func (o *PingOneIdTokenValidatorResponse) GetSubjectClaimName() string {
	if o == nil || isNil(o.SubjectClaimName) {
		var ret string
		return ret
	}
	return *o.SubjectClaimName
}

// GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetSubjectClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectClaimName) {
    return nil, false
	}
	return o.SubjectClaimName, true
}

// HasSubjectClaimName returns a boolean if a field has been set.
func (o *PingOneIdTokenValidatorResponse) HasSubjectClaimName() bool {
	if o != nil && !isNil(o.SubjectClaimName) {
		return true
	}

	return false
}

// SetSubjectClaimName gets a reference to the given string and assigns it to the SubjectClaimName field.
func (o *PingOneIdTokenValidatorResponse) SetSubjectClaimName(v string) {
	o.SubjectClaimName = &v
}

// GetClockSkewGracePeriod returns the ClockSkewGracePeriod field value if set, zero value otherwise.
func (o *PingOneIdTokenValidatorResponse) GetClockSkewGracePeriod() string {
	if o == nil || isNil(o.ClockSkewGracePeriod) {
		var ret string
		return ret
	}
	return *o.ClockSkewGracePeriod
}

// GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetClockSkewGracePeriodOk() (*string, bool) {
	if o == nil || isNil(o.ClockSkewGracePeriod) {
    return nil, false
	}
	return o.ClockSkewGracePeriod, true
}

// HasClockSkewGracePeriod returns a boolean if a field has been set.
func (o *PingOneIdTokenValidatorResponse) HasClockSkewGracePeriod() bool {
	if o != nil && !isNil(o.ClockSkewGracePeriod) {
		return true
	}

	return false
}

// SetClockSkewGracePeriod gets a reference to the given string and assigns it to the ClockSkewGracePeriod field.
func (o *PingOneIdTokenValidatorResponse) SetClockSkewGracePeriod(v string) {
	o.ClockSkewGracePeriod = &v
}

// GetJwksCacheDuration returns the JwksCacheDuration field value if set, zero value otherwise.
func (o *PingOneIdTokenValidatorResponse) GetJwksCacheDuration() string {
	if o == nil || isNil(o.JwksCacheDuration) {
		var ret string
		return ret
	}
	return *o.JwksCacheDuration
}

// GetJwksCacheDurationOk returns a tuple with the JwksCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetJwksCacheDurationOk() (*string, bool) {
	if o == nil || isNil(o.JwksCacheDuration) {
    return nil, false
	}
	return o.JwksCacheDuration, true
}

// HasJwksCacheDuration returns a boolean if a field has been set.
func (o *PingOneIdTokenValidatorResponse) HasJwksCacheDuration() bool {
	if o != nil && !isNil(o.JwksCacheDuration) {
		return true
	}

	return false
}

// SetJwksCacheDuration gets a reference to the given string and assigns it to the JwksCacheDuration field.
func (o *PingOneIdTokenValidatorResponse) SetJwksCacheDuration(v string) {
	o.JwksCacheDuration = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *PingOneIdTokenValidatorResponse) GetEvaluationOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *PingOneIdTokenValidatorResponse) GetEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *PingOneIdTokenValidatorResponse) SetEvaluationOrderIndex(v int32) {
	o.EvaluationOrderIndex = v
}

func (o PingOneIdTokenValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
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

type NullablePingOneIdTokenValidatorResponse struct {
	value *PingOneIdTokenValidatorResponse
	isSet bool
}

func (v NullablePingOneIdTokenValidatorResponse) Get() *PingOneIdTokenValidatorResponse {
	return v.value
}

func (v *NullablePingOneIdTokenValidatorResponse) Set(val *PingOneIdTokenValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePingOneIdTokenValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePingOneIdTokenValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingOneIdTokenValidatorResponse(val *PingOneIdTokenValidatorResponse) *NullablePingOneIdTokenValidatorResponse {
	return &NullablePingOneIdTokenValidatorResponse{value: val, isSet: true}
}

func (v NullablePingOneIdTokenValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingOneIdTokenValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


