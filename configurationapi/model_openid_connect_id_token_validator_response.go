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

// checks if the OpenidConnectIdTokenValidatorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenidConnectIdTokenValidatorResponse{}

// OpenidConnectIdTokenValidatorResponse struct for OpenidConnectIdTokenValidatorResponse
type OpenidConnectIdTokenValidatorResponse struct {
	Schemas                 []EnumopenidConnectIdTokenValidatorSchemaUrn      `json:"schemas"`
	AllowedSigningAlgorithm []EnumidTokenValidatorAllowedSigningAlgorithmProp `json:"allowedSigningAlgorithm"`
	// Specifies the locally stored certificates that may be used to validate the signature of an incoming ID token. This property may be specified if a JWKS endpoint should not be used to retrieve public signing keys.
	SigningCertificate []string `json:"signingCertificate,omitempty"`
	// Specifies the OpenID Connect provider that issues ID tokens handled by this OpenID Connect ID Token Validator. This property is used in conjunction with the jwks-endpoint-path property.
	OpenIDConnectProvider *string `json:"OpenIDConnectProvider,omitempty"`
	// The relative path to the JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming ID token. This path is relative to the base_url property defined for the validator's OpenID Connect provider. If jwks-endpoint-path is specified, the OpenID Connect ID Token Validator will not consult locally stored certificates for validating token signatures.
	JwksEndpointPath *string `json:"jwksEndpointPath,omitempty"`
	// A description for this ID Token Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether this ID Token Validator is enabled for use in the Directory Server.
	Enabled bool `json:"enabled"`
	// Specifies the name of the Identity Mapper that should be used to correlate an ID token subject value to a user entry. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property.
	IdentityMapper string `json:"identityMapper"`
	// The name of the token claim that contains the subject; i.e., the authenticated user.
	SubjectClaimName *string `json:"subjectClaimName,omitempty"`
	// Specifies the OpenID Connect provider's issuer URL.
	IssuerURL string `json:"issuerURL"`
	// Specifies the amount of clock skew that is tolerated by the ID Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token's not-before (nbf) time and added to the token's expiration (exp) time, if present, to allow for any time difference between the local server's clock and the token issuer's clock.
	ClockSkewGracePeriod *string `json:"clockSkewGracePeriod,omitempty"`
	// How often the ID Token Validator should refresh its cache of JWKS token signing keys.
	JwksCacheDuration *string `json:"jwksCacheDuration,omitempty"`
	// When multiple ID Token Validators are defined for a single Directory Server, this property determines the order in which the ID Token Validators are consulted. Values of this property must be unique among all ID Token Validators defined within Directory Server but not necessarily contiguous. ID Token Validators with lower values will be evaluated first to determine if they are able to validate the ID token.
	EvaluationOrderIndex                          int64                                              `json:"evaluationOrderIndex"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the ID Token Validator
	Id string `json:"id"`
}

// NewOpenidConnectIdTokenValidatorResponse instantiates a new OpenidConnectIdTokenValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenidConnectIdTokenValidatorResponse(schemas []EnumopenidConnectIdTokenValidatorSchemaUrn, allowedSigningAlgorithm []EnumidTokenValidatorAllowedSigningAlgorithmProp, enabled bool, identityMapper string, issuerURL string, evaluationOrderIndex int64, id string) *OpenidConnectIdTokenValidatorResponse {
	this := OpenidConnectIdTokenValidatorResponse{}
	this.Schemas = schemas
	this.AllowedSigningAlgorithm = allowedSigningAlgorithm
	this.Enabled = enabled
	this.IdentityMapper = identityMapper
	this.IssuerURL = issuerURL
	this.EvaluationOrderIndex = evaluationOrderIndex
	this.Id = id
	return &this
}

// NewOpenidConnectIdTokenValidatorResponseWithDefaults instantiates a new OpenidConnectIdTokenValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenidConnectIdTokenValidatorResponseWithDefaults() *OpenidConnectIdTokenValidatorResponse {
	this := OpenidConnectIdTokenValidatorResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *OpenidConnectIdTokenValidatorResponse) GetSchemas() []EnumopenidConnectIdTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnumopenidConnectIdTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetSchemasOk() ([]EnumopenidConnectIdTokenValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *OpenidConnectIdTokenValidatorResponse) SetSchemas(v []EnumopenidConnectIdTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field value
func (o *OpenidConnectIdTokenValidatorResponse) GetAllowedSigningAlgorithm() []EnumidTokenValidatorAllowedSigningAlgorithmProp {
	if o == nil {
		var ret []EnumidTokenValidatorAllowedSigningAlgorithmProp
		return ret
	}

	return o.AllowedSigningAlgorithm
}

// GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field value
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetAllowedSigningAlgorithmOk() ([]EnumidTokenValidatorAllowedSigningAlgorithmProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedSigningAlgorithm, true
}

// SetAllowedSigningAlgorithm sets field value
func (o *OpenidConnectIdTokenValidatorResponse) SetAllowedSigningAlgorithm(v []EnumidTokenValidatorAllowedSigningAlgorithmProp) {
	o.AllowedSigningAlgorithm = v
}

// GetSigningCertificate returns the SigningCertificate field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetSigningCertificate() []string {
	if o == nil || IsNil(o.SigningCertificate) {
		var ret []string
		return ret
	}
	return o.SigningCertificate
}

// GetSigningCertificateOk returns a tuple with the SigningCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetSigningCertificateOk() ([]string, bool) {
	if o == nil || IsNil(o.SigningCertificate) {
		return nil, false
	}
	return o.SigningCertificate, true
}

// HasSigningCertificate returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasSigningCertificate() bool {
	if o != nil && !IsNil(o.SigningCertificate) {
		return true
	}

	return false
}

// SetSigningCertificate gets a reference to the given []string and assigns it to the SigningCertificate field.
func (o *OpenidConnectIdTokenValidatorResponse) SetSigningCertificate(v []string) {
	o.SigningCertificate = v
}

// GetOpenIDConnectProvider returns the OpenIDConnectProvider field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetOpenIDConnectProvider() string {
	if o == nil || IsNil(o.OpenIDConnectProvider) {
		var ret string
		return ret
	}
	return *o.OpenIDConnectProvider
}

// GetOpenIDConnectProviderOk returns a tuple with the OpenIDConnectProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetOpenIDConnectProviderOk() (*string, bool) {
	if o == nil || IsNil(o.OpenIDConnectProvider) {
		return nil, false
	}
	return o.OpenIDConnectProvider, true
}

// HasOpenIDConnectProvider returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasOpenIDConnectProvider() bool {
	if o != nil && !IsNil(o.OpenIDConnectProvider) {
		return true
	}

	return false
}

// SetOpenIDConnectProvider gets a reference to the given string and assigns it to the OpenIDConnectProvider field.
func (o *OpenidConnectIdTokenValidatorResponse) SetOpenIDConnectProvider(v string) {
	o.OpenIDConnectProvider = &v
}

// GetJwksEndpointPath returns the JwksEndpointPath field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetJwksEndpointPath() string {
	if o == nil || IsNil(o.JwksEndpointPath) {
		var ret string
		return ret
	}
	return *o.JwksEndpointPath
}

// GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetJwksEndpointPathOk() (*string, bool) {
	if o == nil || IsNil(o.JwksEndpointPath) {
		return nil, false
	}
	return o.JwksEndpointPath, true
}

// HasJwksEndpointPath returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasJwksEndpointPath() bool {
	if o != nil && !IsNil(o.JwksEndpointPath) {
		return true
	}

	return false
}

// SetJwksEndpointPath gets a reference to the given string and assigns it to the JwksEndpointPath field.
func (o *OpenidConnectIdTokenValidatorResponse) SetJwksEndpointPath(v string) {
	o.JwksEndpointPath = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OpenidConnectIdTokenValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *OpenidConnectIdTokenValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *OpenidConnectIdTokenValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *OpenidConnectIdTokenValidatorResponse) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *OpenidConnectIdTokenValidatorResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetSubjectClaimName returns the SubjectClaimName field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetSubjectClaimName() string {
	if o == nil || IsNil(o.SubjectClaimName) {
		var ret string
		return ret
	}
	return *o.SubjectClaimName
}

// GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetSubjectClaimNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectClaimName) {
		return nil, false
	}
	return o.SubjectClaimName, true
}

// HasSubjectClaimName returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasSubjectClaimName() bool {
	if o != nil && !IsNil(o.SubjectClaimName) {
		return true
	}

	return false
}

// SetSubjectClaimName gets a reference to the given string and assigns it to the SubjectClaimName field.
func (o *OpenidConnectIdTokenValidatorResponse) SetSubjectClaimName(v string) {
	o.SubjectClaimName = &v
}

// GetIssuerURL returns the IssuerURL field value
func (o *OpenidConnectIdTokenValidatorResponse) GetIssuerURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerURL
}

// GetIssuerURLOk returns a tuple with the IssuerURL field value
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetIssuerURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerURL, true
}

// SetIssuerURL sets field value
func (o *OpenidConnectIdTokenValidatorResponse) SetIssuerURL(v string) {
	o.IssuerURL = v
}

// GetClockSkewGracePeriod returns the ClockSkewGracePeriod field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetClockSkewGracePeriod() string {
	if o == nil || IsNil(o.ClockSkewGracePeriod) {
		var ret string
		return ret
	}
	return *o.ClockSkewGracePeriod
}

// GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetClockSkewGracePeriodOk() (*string, bool) {
	if o == nil || IsNil(o.ClockSkewGracePeriod) {
		return nil, false
	}
	return o.ClockSkewGracePeriod, true
}

// HasClockSkewGracePeriod returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasClockSkewGracePeriod() bool {
	if o != nil && !IsNil(o.ClockSkewGracePeriod) {
		return true
	}

	return false
}

// SetClockSkewGracePeriod gets a reference to the given string and assigns it to the ClockSkewGracePeriod field.
func (o *OpenidConnectIdTokenValidatorResponse) SetClockSkewGracePeriod(v string) {
	o.ClockSkewGracePeriod = &v
}

// GetJwksCacheDuration returns the JwksCacheDuration field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetJwksCacheDuration() string {
	if o == nil || IsNil(o.JwksCacheDuration) {
		var ret string
		return ret
	}
	return *o.JwksCacheDuration
}

// GetJwksCacheDurationOk returns a tuple with the JwksCacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetJwksCacheDurationOk() (*string, bool) {
	if o == nil || IsNil(o.JwksCacheDuration) {
		return nil, false
	}
	return o.JwksCacheDuration, true
}

// HasJwksCacheDuration returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasJwksCacheDuration() bool {
	if o != nil && !IsNil(o.JwksCacheDuration) {
		return true
	}

	return false
}

// SetJwksCacheDuration gets a reference to the given string and assigns it to the JwksCacheDuration field.
func (o *OpenidConnectIdTokenValidatorResponse) SetJwksCacheDuration(v string) {
	o.JwksCacheDuration = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *OpenidConnectIdTokenValidatorResponse) GetEvaluationOrderIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetEvaluationOrderIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *OpenidConnectIdTokenValidatorResponse) SetEvaluationOrderIndex(v int64) {
	o.EvaluationOrderIndex = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *OpenidConnectIdTokenValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *OpenidConnectIdTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *OpenidConnectIdTokenValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *OpenidConnectIdTokenValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *OpenidConnectIdTokenValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpenidConnectIdTokenValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpenidConnectIdTokenValidatorResponse) SetId(v string) {
	o.Id = v
}

func (o OpenidConnectIdTokenValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenidConnectIdTokenValidatorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["allowedSigningAlgorithm"] = o.AllowedSigningAlgorithm
	if !IsNil(o.SigningCertificate) {
		toSerialize["signingCertificate"] = o.SigningCertificate
	}
	if !IsNil(o.OpenIDConnectProvider) {
		toSerialize["OpenIDConnectProvider"] = o.OpenIDConnectProvider
	}
	if !IsNil(o.JwksEndpointPath) {
		toSerialize["jwksEndpointPath"] = o.JwksEndpointPath
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["identityMapper"] = o.IdentityMapper
	if !IsNil(o.SubjectClaimName) {
		toSerialize["subjectClaimName"] = o.SubjectClaimName
	}
	toSerialize["issuerURL"] = o.IssuerURL
	if !IsNil(o.ClockSkewGracePeriod) {
		toSerialize["clockSkewGracePeriod"] = o.ClockSkewGracePeriod
	}
	if !IsNil(o.JwksCacheDuration) {
		toSerialize["jwksCacheDuration"] = o.JwksCacheDuration
	}
	toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableOpenidConnectIdTokenValidatorResponse struct {
	value *OpenidConnectIdTokenValidatorResponse
	isSet bool
}

func (v NullableOpenidConnectIdTokenValidatorResponse) Get() *OpenidConnectIdTokenValidatorResponse {
	return v.value
}

func (v *NullableOpenidConnectIdTokenValidatorResponse) Set(val *OpenidConnectIdTokenValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenidConnectIdTokenValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenidConnectIdTokenValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenidConnectIdTokenValidatorResponse(val *OpenidConnectIdTokenValidatorResponse) *NullableOpenidConnectIdTokenValidatorResponse {
	return &NullableOpenidConnectIdTokenValidatorResponse{value: val, isSet: true}
}

func (v NullableOpenidConnectIdTokenValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenidConnectIdTokenValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
