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

// JwtAccessTokenValidatorResponse struct for JwtAccessTokenValidatorResponse
type JwtAccessTokenValidatorResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Access Token Validator
	Id                      string                                                `json:"id"`
	Schemas                 []EnumjwtAccessTokenValidatorSchemaUrn                `json:"schemas"`
	AllowedSigningAlgorithm []EnumaccessTokenValidatorAllowedSigningAlgorithmProp `json:"allowedSigningAlgorithm"`
	// Specifies the locally stored certificates that may be used to validate the signature of an incoming JWT access token. If this property is specified, the JWT Access Token Validator will not use a JWKS endpoint to retrieve public keys.
	SigningCertificate []string `json:"signingCertificate,omitempty"`
	// The relative path to JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming JWT access token. This path is relative to the base_url property defined for the validator's external authorization server. If jwks-endpoint-path is specified, the JWT Access Token Validator will not consult locally stored certificates for validating token signatures.
	JwksEndpointPath *string `json:"jwksEndpointPath,omitempty"`
	// The public-private key pair that is used to encrypt the JWT payload. If specified, the JWT Access Token Validator will use the private key to decrypt the JWT payload, and the public key must be exported to the Authorization Server that is issuing access tokens.
	EncryptionKeyPair                 *string                                                         `json:"encryptionKeyPair,omitempty"`
	AllowedKeyEncryptionAlgorithm     []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp     `json:"allowedKeyEncryptionAlgorithm"`
	AllowedContentEncryptionAlgorithm []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp `json:"allowedContentEncryptionAlgorithm"`
	// Specifies the amount of clock skew that is tolerated by the JWT Access Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token's not-before (nbf) time and added to the token's expiration (exp) time, if present, to allow for any time difference between the local server's clock and the token issuer's clock.
	ClockSkewGracePeriod *string `json:"clockSkewGracePeriod,omitempty"`
	// The name of the token claim that contains the OAuth2 client Id.
	ClientIDClaimName *string `json:"clientIDClaimName,omitempty"`
	// The name of the token claim that contains the scopes granted by the token.
	ScopeClaimName *string `json:"scopeClaimName,omitempty"`
	// When multiple JWT Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all JWT Access Token Validators defined within Directory Server but not necessarily contiguous. JWT Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token.
	EvaluationOrderIndex int32 `json:"evaluationOrderIndex"`
	// Specifies the external server that will be used to aid in validating access tokens. In most cases this will be the Authorization Server that minted the token.
	AuthorizationServer *string `json:"authorizationServer,omitempty"`
	// Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server.
	SubjectClaimName *string `json:"subjectClaimName,omitempty"`
	// A description for this Access Token Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether this Access Token Validator is enabled for use in Directory Server.
	Enabled bool `json:"enabled"`
}

// NewJwtAccessTokenValidatorResponse instantiates a new JwtAccessTokenValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJwtAccessTokenValidatorResponse(id string, schemas []EnumjwtAccessTokenValidatorSchemaUrn, allowedSigningAlgorithm []EnumaccessTokenValidatorAllowedSigningAlgorithmProp, allowedKeyEncryptionAlgorithm []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, allowedContentEncryptionAlgorithm []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, evaluationOrderIndex int32, enabled bool) *JwtAccessTokenValidatorResponse {
	this := JwtAccessTokenValidatorResponse{}
	this.Id = id
	this.Schemas = schemas
	this.AllowedSigningAlgorithm = allowedSigningAlgorithm
	this.AllowedKeyEncryptionAlgorithm = allowedKeyEncryptionAlgorithm
	this.AllowedContentEncryptionAlgorithm = allowedContentEncryptionAlgorithm
	this.EvaluationOrderIndex = evaluationOrderIndex
	this.Enabled = enabled
	return &this
}

// NewJwtAccessTokenValidatorResponseWithDefaults instantiates a new JwtAccessTokenValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJwtAccessTokenValidatorResponseWithDefaults() *JwtAccessTokenValidatorResponse {
	this := JwtAccessTokenValidatorResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JwtAccessTokenValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *JwtAccessTokenValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *JwtAccessTokenValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JwtAccessTokenValidatorResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *JwtAccessTokenValidatorResponse) GetSchemas() []EnumjwtAccessTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnumjwtAccessTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetSchemasOk() ([]EnumjwtAccessTokenValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JwtAccessTokenValidatorResponse) SetSchemas(v []EnumjwtAccessTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field value
func (o *JwtAccessTokenValidatorResponse) GetAllowedSigningAlgorithm() []EnumaccessTokenValidatorAllowedSigningAlgorithmProp {
	if o == nil {
		var ret []EnumaccessTokenValidatorAllowedSigningAlgorithmProp
		return ret
	}

	return o.AllowedSigningAlgorithm
}

// GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field value
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetAllowedSigningAlgorithmOk() ([]EnumaccessTokenValidatorAllowedSigningAlgorithmProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedSigningAlgorithm, true
}

// SetAllowedSigningAlgorithm sets field value
func (o *JwtAccessTokenValidatorResponse) SetAllowedSigningAlgorithm(v []EnumaccessTokenValidatorAllowedSigningAlgorithmProp) {
	o.AllowedSigningAlgorithm = v
}

// GetSigningCertificate returns the SigningCertificate field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetSigningCertificate() []string {
	if o == nil || isNil(o.SigningCertificate) {
		var ret []string
		return ret
	}
	return o.SigningCertificate
}

// GetSigningCertificateOk returns a tuple with the SigningCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetSigningCertificateOk() ([]string, bool) {
	if o == nil || isNil(o.SigningCertificate) {
		return nil, false
	}
	return o.SigningCertificate, true
}

// HasSigningCertificate returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasSigningCertificate() bool {
	if o != nil && !isNil(o.SigningCertificate) {
		return true
	}

	return false
}

// SetSigningCertificate gets a reference to the given []string and assigns it to the SigningCertificate field.
func (o *JwtAccessTokenValidatorResponse) SetSigningCertificate(v []string) {
	o.SigningCertificate = v
}

// GetJwksEndpointPath returns the JwksEndpointPath field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetJwksEndpointPath() string {
	if o == nil || isNil(o.JwksEndpointPath) {
		var ret string
		return ret
	}
	return *o.JwksEndpointPath
}

// GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetJwksEndpointPathOk() (*string, bool) {
	if o == nil || isNil(o.JwksEndpointPath) {
		return nil, false
	}
	return o.JwksEndpointPath, true
}

// HasJwksEndpointPath returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasJwksEndpointPath() bool {
	if o != nil && !isNil(o.JwksEndpointPath) {
		return true
	}

	return false
}

// SetJwksEndpointPath gets a reference to the given string and assigns it to the JwksEndpointPath field.
func (o *JwtAccessTokenValidatorResponse) SetJwksEndpointPath(v string) {
	o.JwksEndpointPath = &v
}

// GetEncryptionKeyPair returns the EncryptionKeyPair field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetEncryptionKeyPair() string {
	if o == nil || isNil(o.EncryptionKeyPair) {
		var ret string
		return ret
	}
	return *o.EncryptionKeyPair
}

// GetEncryptionKeyPairOk returns a tuple with the EncryptionKeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetEncryptionKeyPairOk() (*string, bool) {
	if o == nil || isNil(o.EncryptionKeyPair) {
		return nil, false
	}
	return o.EncryptionKeyPair, true
}

// HasEncryptionKeyPair returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasEncryptionKeyPair() bool {
	if o != nil && !isNil(o.EncryptionKeyPair) {
		return true
	}

	return false
}

// SetEncryptionKeyPair gets a reference to the given string and assigns it to the EncryptionKeyPair field.
func (o *JwtAccessTokenValidatorResponse) SetEncryptionKeyPair(v string) {
	o.EncryptionKeyPair = &v
}

// GetAllowedKeyEncryptionAlgorithm returns the AllowedKeyEncryptionAlgorithm field value
func (o *JwtAccessTokenValidatorResponse) GetAllowedKeyEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp {
	if o == nil {
		var ret []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp
		return ret
	}

	return o.AllowedKeyEncryptionAlgorithm
}

// GetAllowedKeyEncryptionAlgorithmOk returns a tuple with the AllowedKeyEncryptionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetAllowedKeyEncryptionAlgorithmOk() ([]EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedKeyEncryptionAlgorithm, true
}

// SetAllowedKeyEncryptionAlgorithm sets field value
func (o *JwtAccessTokenValidatorResponse) SetAllowedKeyEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) {
	o.AllowedKeyEncryptionAlgorithm = v
}

// GetAllowedContentEncryptionAlgorithm returns the AllowedContentEncryptionAlgorithm field value
func (o *JwtAccessTokenValidatorResponse) GetAllowedContentEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp {
	if o == nil {
		var ret []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp
		return ret
	}

	return o.AllowedContentEncryptionAlgorithm
}

// GetAllowedContentEncryptionAlgorithmOk returns a tuple with the AllowedContentEncryptionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetAllowedContentEncryptionAlgorithmOk() ([]EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedContentEncryptionAlgorithm, true
}

// SetAllowedContentEncryptionAlgorithm sets field value
func (o *JwtAccessTokenValidatorResponse) SetAllowedContentEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp) {
	o.AllowedContentEncryptionAlgorithm = v
}

// GetClockSkewGracePeriod returns the ClockSkewGracePeriod field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetClockSkewGracePeriod() string {
	if o == nil || isNil(o.ClockSkewGracePeriod) {
		var ret string
		return ret
	}
	return *o.ClockSkewGracePeriod
}

// GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetClockSkewGracePeriodOk() (*string, bool) {
	if o == nil || isNil(o.ClockSkewGracePeriod) {
		return nil, false
	}
	return o.ClockSkewGracePeriod, true
}

// HasClockSkewGracePeriod returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasClockSkewGracePeriod() bool {
	if o != nil && !isNil(o.ClockSkewGracePeriod) {
		return true
	}

	return false
}

// SetClockSkewGracePeriod gets a reference to the given string and assigns it to the ClockSkewGracePeriod field.
func (o *JwtAccessTokenValidatorResponse) SetClockSkewGracePeriod(v string) {
	o.ClockSkewGracePeriod = &v
}

// GetClientIDClaimName returns the ClientIDClaimName field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetClientIDClaimName() string {
	if o == nil || isNil(o.ClientIDClaimName) {
		var ret string
		return ret
	}
	return *o.ClientIDClaimName
}

// GetClientIDClaimNameOk returns a tuple with the ClientIDClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetClientIDClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.ClientIDClaimName) {
		return nil, false
	}
	return o.ClientIDClaimName, true
}

// HasClientIDClaimName returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasClientIDClaimName() bool {
	if o != nil && !isNil(o.ClientIDClaimName) {
		return true
	}

	return false
}

// SetClientIDClaimName gets a reference to the given string and assigns it to the ClientIDClaimName field.
func (o *JwtAccessTokenValidatorResponse) SetClientIDClaimName(v string) {
	o.ClientIDClaimName = &v
}

// GetScopeClaimName returns the ScopeClaimName field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetScopeClaimName() string {
	if o == nil || isNil(o.ScopeClaimName) {
		var ret string
		return ret
	}
	return *o.ScopeClaimName
}

// GetScopeClaimNameOk returns a tuple with the ScopeClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetScopeClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.ScopeClaimName) {
		return nil, false
	}
	return o.ScopeClaimName, true
}

// HasScopeClaimName returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasScopeClaimName() bool {
	if o != nil && !isNil(o.ScopeClaimName) {
		return true
	}

	return false
}

// SetScopeClaimName gets a reference to the given string and assigns it to the ScopeClaimName field.
func (o *JwtAccessTokenValidatorResponse) SetScopeClaimName(v string) {
	o.ScopeClaimName = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *JwtAccessTokenValidatorResponse) GetEvaluationOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *JwtAccessTokenValidatorResponse) SetEvaluationOrderIndex(v int32) {
	o.EvaluationOrderIndex = v
}

// GetAuthorizationServer returns the AuthorizationServer field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetAuthorizationServer() string {
	if o == nil || isNil(o.AuthorizationServer) {
		var ret string
		return ret
	}
	return *o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetAuthorizationServerOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationServer) {
		return nil, false
	}
	return o.AuthorizationServer, true
}

// HasAuthorizationServer returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasAuthorizationServer() bool {
	if o != nil && !isNil(o.AuthorizationServer) {
		return true
	}

	return false
}

// SetAuthorizationServer gets a reference to the given string and assigns it to the AuthorizationServer field.
func (o *JwtAccessTokenValidatorResponse) SetAuthorizationServer(v string) {
	o.AuthorizationServer = &v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetIdentityMapper() string {
	if o == nil || isNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil || isNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasIdentityMapper() bool {
	if o != nil && !isNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *JwtAccessTokenValidatorResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetSubjectClaimName returns the SubjectClaimName field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetSubjectClaimName() string {
	if o == nil || isNil(o.SubjectClaimName) {
		var ret string
		return ret
	}
	return *o.SubjectClaimName
}

// GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetSubjectClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectClaimName) {
		return nil, false
	}
	return o.SubjectClaimName, true
}

// HasSubjectClaimName returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasSubjectClaimName() bool {
	if o != nil && !isNil(o.SubjectClaimName) {
		return true
	}

	return false
}

// SetSubjectClaimName gets a reference to the given string and assigns it to the SubjectClaimName field.
func (o *JwtAccessTokenValidatorResponse) SetSubjectClaimName(v string) {
	o.SubjectClaimName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JwtAccessTokenValidatorResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JwtAccessTokenValidatorResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JwtAccessTokenValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *JwtAccessTokenValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *JwtAccessTokenValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *JwtAccessTokenValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o JwtAccessTokenValidatorResponse) MarshalJSON() ([]byte, error) {
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
		toSerialize["allowedSigningAlgorithm"] = o.AllowedSigningAlgorithm
	}
	if !isNil(o.SigningCertificate) {
		toSerialize["signingCertificate"] = o.SigningCertificate
	}
	if !isNil(o.JwksEndpointPath) {
		toSerialize["jwksEndpointPath"] = o.JwksEndpointPath
	}
	if !isNil(o.EncryptionKeyPair) {
		toSerialize["encryptionKeyPair"] = o.EncryptionKeyPair
	}
	if true {
		toSerialize["allowedKeyEncryptionAlgorithm"] = o.AllowedKeyEncryptionAlgorithm
	}
	if true {
		toSerialize["allowedContentEncryptionAlgorithm"] = o.AllowedContentEncryptionAlgorithm
	}
	if !isNil(o.ClockSkewGracePeriod) {
		toSerialize["clockSkewGracePeriod"] = o.ClockSkewGracePeriod
	}
	if !isNil(o.ClientIDClaimName) {
		toSerialize["clientIDClaimName"] = o.ClientIDClaimName
	}
	if !isNil(o.ScopeClaimName) {
		toSerialize["scopeClaimName"] = o.ScopeClaimName
	}
	if true {
		toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	}
	if !isNil(o.AuthorizationServer) {
		toSerialize["authorizationServer"] = o.AuthorizationServer
	}
	if !isNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !isNil(o.SubjectClaimName) {
		toSerialize["subjectClaimName"] = o.SubjectClaimName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableJwtAccessTokenValidatorResponse struct {
	value *JwtAccessTokenValidatorResponse
	isSet bool
}

func (v NullableJwtAccessTokenValidatorResponse) Get() *JwtAccessTokenValidatorResponse {
	return v.value
}

func (v *NullableJwtAccessTokenValidatorResponse) Set(val *JwtAccessTokenValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJwtAccessTokenValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJwtAccessTokenValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJwtAccessTokenValidatorResponse(val *JwtAccessTokenValidatorResponse) *NullableJwtAccessTokenValidatorResponse {
	return &NullableJwtAccessTokenValidatorResponse{value: val, isSet: true}
}

func (v NullableJwtAccessTokenValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJwtAccessTokenValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
