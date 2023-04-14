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

// checks if the AddJwtAccessTokenValidatorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddJwtAccessTokenValidatorRequest{}

// AddJwtAccessTokenValidatorRequest struct for AddJwtAccessTokenValidatorRequest
type AddJwtAccessTokenValidatorRequest struct {
	// Name of the new Access Token Validator
	ValidatorName           string                                                `json:"validatorName"`
	Schemas                 []EnumjwtAccessTokenValidatorSchemaUrn                `json:"schemas"`
	AllowedSigningAlgorithm []EnumaccessTokenValidatorAllowedSigningAlgorithmProp `json:"allowedSigningAlgorithm,omitempty"`
	// Specifies the locally stored certificates that may be used to validate the signature of an incoming JWT access token. If this property is specified, the JWT Access Token Validator will not use a JWKS endpoint to retrieve public keys.
	SigningCertificate []string `json:"signingCertificate,omitempty"`
	// The relative path to JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming JWT access token. This path is relative to the base_url property defined for the validator's external authorization server. If jwks-endpoint-path is specified, the JWT Access Token Validator will not consult locally stored certificates for validating token signatures.
	JwksEndpointPath *string `json:"jwksEndpointPath,omitempty"`
	// The public-private key pair that is used to encrypt the JWT payload. If specified, the JWT Access Token Validator will use the private key to decrypt the JWT payload, and the public key must be exported to the Authorization Server that is issuing access tokens.
	EncryptionKeyPair                 *string                                                         `json:"encryptionKeyPair,omitempty"`
	AllowedKeyEncryptionAlgorithm     []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp     `json:"allowedKeyEncryptionAlgorithm,omitempty"`
	AllowedContentEncryptionAlgorithm []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp `json:"allowedContentEncryptionAlgorithm,omitempty"`
	// Specifies the amount of clock skew that is tolerated by the JWT Access Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token's not-before (nbf) time and added to the token's expiration (exp) time, if present, to allow for any time difference between the local server's clock and the token issuer's clock.
	ClockSkewGracePeriod *string `json:"clockSkewGracePeriod,omitempty"`
	// The name of the token claim that contains the OAuth2 client Id.
	ClientIDClaimName *string `json:"clientIDClaimName,omitempty"`
	// The name of the token claim that contains the scopes granted by the token.
	ScopeClaimName *string `json:"scopeClaimName,omitempty"`
	// When multiple JWT Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all JWT Access Token Validators defined within Directory Server but not necessarily contiguous. JWT Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token.
	EvaluationOrderIndex *int64 `json:"evaluationOrderIndex,omitempty"`
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

// NewAddJwtAccessTokenValidatorRequest instantiates a new AddJwtAccessTokenValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddJwtAccessTokenValidatorRequest(validatorName string, schemas []EnumjwtAccessTokenValidatorSchemaUrn, enabled bool) *AddJwtAccessTokenValidatorRequest {
	this := AddJwtAccessTokenValidatorRequest{}
	this.ValidatorName = validatorName
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddJwtAccessTokenValidatorRequestWithDefaults instantiates a new AddJwtAccessTokenValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddJwtAccessTokenValidatorRequestWithDefaults() *AddJwtAccessTokenValidatorRequest {
	this := AddJwtAccessTokenValidatorRequest{}
	return &this
}

// GetValidatorName returns the ValidatorName field value
func (o *AddJwtAccessTokenValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddJwtAccessTokenValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddJwtAccessTokenValidatorRequest) GetSchemas() []EnumjwtAccessTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnumjwtAccessTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetSchemasOk() ([]EnumjwtAccessTokenValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddJwtAccessTokenValidatorRequest) SetSchemas(v []EnumjwtAccessTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetAllowedSigningAlgorithm() []EnumaccessTokenValidatorAllowedSigningAlgorithmProp {
	if o == nil || IsNil(o.AllowedSigningAlgorithm) {
		var ret []EnumaccessTokenValidatorAllowedSigningAlgorithmProp
		return ret
	}
	return o.AllowedSigningAlgorithm
}

// GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetAllowedSigningAlgorithmOk() ([]EnumaccessTokenValidatorAllowedSigningAlgorithmProp, bool) {
	if o == nil || IsNil(o.AllowedSigningAlgorithm) {
		return nil, false
	}
	return o.AllowedSigningAlgorithm, true
}

// HasAllowedSigningAlgorithm returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasAllowedSigningAlgorithm() bool {
	if o != nil && !IsNil(o.AllowedSigningAlgorithm) {
		return true
	}

	return false
}

// SetAllowedSigningAlgorithm gets a reference to the given []EnumaccessTokenValidatorAllowedSigningAlgorithmProp and assigns it to the AllowedSigningAlgorithm field.
func (o *AddJwtAccessTokenValidatorRequest) SetAllowedSigningAlgorithm(v []EnumaccessTokenValidatorAllowedSigningAlgorithmProp) {
	o.AllowedSigningAlgorithm = v
}

// GetSigningCertificate returns the SigningCertificate field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetSigningCertificate() []string {
	if o == nil || IsNil(o.SigningCertificate) {
		var ret []string
		return ret
	}
	return o.SigningCertificate
}

// GetSigningCertificateOk returns a tuple with the SigningCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetSigningCertificateOk() ([]string, bool) {
	if o == nil || IsNil(o.SigningCertificate) {
		return nil, false
	}
	return o.SigningCertificate, true
}

// HasSigningCertificate returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasSigningCertificate() bool {
	if o != nil && !IsNil(o.SigningCertificate) {
		return true
	}

	return false
}

// SetSigningCertificate gets a reference to the given []string and assigns it to the SigningCertificate field.
func (o *AddJwtAccessTokenValidatorRequest) SetSigningCertificate(v []string) {
	o.SigningCertificate = v
}

// GetJwksEndpointPath returns the JwksEndpointPath field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetJwksEndpointPath() string {
	if o == nil || IsNil(o.JwksEndpointPath) {
		var ret string
		return ret
	}
	return *o.JwksEndpointPath
}

// GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetJwksEndpointPathOk() (*string, bool) {
	if o == nil || IsNil(o.JwksEndpointPath) {
		return nil, false
	}
	return o.JwksEndpointPath, true
}

// HasJwksEndpointPath returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasJwksEndpointPath() bool {
	if o != nil && !IsNil(o.JwksEndpointPath) {
		return true
	}

	return false
}

// SetJwksEndpointPath gets a reference to the given string and assigns it to the JwksEndpointPath field.
func (o *AddJwtAccessTokenValidatorRequest) SetJwksEndpointPath(v string) {
	o.JwksEndpointPath = &v
}

// GetEncryptionKeyPair returns the EncryptionKeyPair field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetEncryptionKeyPair() string {
	if o == nil || IsNil(o.EncryptionKeyPair) {
		var ret string
		return ret
	}
	return *o.EncryptionKeyPair
}

// GetEncryptionKeyPairOk returns a tuple with the EncryptionKeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetEncryptionKeyPairOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptionKeyPair) {
		return nil, false
	}
	return o.EncryptionKeyPair, true
}

// HasEncryptionKeyPair returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasEncryptionKeyPair() bool {
	if o != nil && !IsNil(o.EncryptionKeyPair) {
		return true
	}

	return false
}

// SetEncryptionKeyPair gets a reference to the given string and assigns it to the EncryptionKeyPair field.
func (o *AddJwtAccessTokenValidatorRequest) SetEncryptionKeyPair(v string) {
	o.EncryptionKeyPair = &v
}

// GetAllowedKeyEncryptionAlgorithm returns the AllowedKeyEncryptionAlgorithm field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetAllowedKeyEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp {
	if o == nil || IsNil(o.AllowedKeyEncryptionAlgorithm) {
		var ret []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp
		return ret
	}
	return o.AllowedKeyEncryptionAlgorithm
}

// GetAllowedKeyEncryptionAlgorithmOk returns a tuple with the AllowedKeyEncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetAllowedKeyEncryptionAlgorithmOk() ([]EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, bool) {
	if o == nil || IsNil(o.AllowedKeyEncryptionAlgorithm) {
		return nil, false
	}
	return o.AllowedKeyEncryptionAlgorithm, true
}

// HasAllowedKeyEncryptionAlgorithm returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasAllowedKeyEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.AllowedKeyEncryptionAlgorithm) {
		return true
	}

	return false
}

// SetAllowedKeyEncryptionAlgorithm gets a reference to the given []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp and assigns it to the AllowedKeyEncryptionAlgorithm field.
func (o *AddJwtAccessTokenValidatorRequest) SetAllowedKeyEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp) {
	o.AllowedKeyEncryptionAlgorithm = v
}

// GetAllowedContentEncryptionAlgorithm returns the AllowedContentEncryptionAlgorithm field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetAllowedContentEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp {
	if o == nil || IsNil(o.AllowedContentEncryptionAlgorithm) {
		var ret []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp
		return ret
	}
	return o.AllowedContentEncryptionAlgorithm
}

// GetAllowedContentEncryptionAlgorithmOk returns a tuple with the AllowedContentEncryptionAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetAllowedContentEncryptionAlgorithmOk() ([]EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, bool) {
	if o == nil || IsNil(o.AllowedContentEncryptionAlgorithm) {
		return nil, false
	}
	return o.AllowedContentEncryptionAlgorithm, true
}

// HasAllowedContentEncryptionAlgorithm returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasAllowedContentEncryptionAlgorithm() bool {
	if o != nil && !IsNil(o.AllowedContentEncryptionAlgorithm) {
		return true
	}

	return false
}

// SetAllowedContentEncryptionAlgorithm gets a reference to the given []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp and assigns it to the AllowedContentEncryptionAlgorithm field.
func (o *AddJwtAccessTokenValidatorRequest) SetAllowedContentEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp) {
	o.AllowedContentEncryptionAlgorithm = v
}

// GetClockSkewGracePeriod returns the ClockSkewGracePeriod field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetClockSkewGracePeriod() string {
	if o == nil || IsNil(o.ClockSkewGracePeriod) {
		var ret string
		return ret
	}
	return *o.ClockSkewGracePeriod
}

// GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetClockSkewGracePeriodOk() (*string, bool) {
	if o == nil || IsNil(o.ClockSkewGracePeriod) {
		return nil, false
	}
	return o.ClockSkewGracePeriod, true
}

// HasClockSkewGracePeriod returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasClockSkewGracePeriod() bool {
	if o != nil && !IsNil(o.ClockSkewGracePeriod) {
		return true
	}

	return false
}

// SetClockSkewGracePeriod gets a reference to the given string and assigns it to the ClockSkewGracePeriod field.
func (o *AddJwtAccessTokenValidatorRequest) SetClockSkewGracePeriod(v string) {
	o.ClockSkewGracePeriod = &v
}

// GetClientIDClaimName returns the ClientIDClaimName field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetClientIDClaimName() string {
	if o == nil || IsNil(o.ClientIDClaimName) {
		var ret string
		return ret
	}
	return *o.ClientIDClaimName
}

// GetClientIDClaimNameOk returns a tuple with the ClientIDClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetClientIDClaimNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIDClaimName) {
		return nil, false
	}
	return o.ClientIDClaimName, true
}

// HasClientIDClaimName returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasClientIDClaimName() bool {
	if o != nil && !IsNil(o.ClientIDClaimName) {
		return true
	}

	return false
}

// SetClientIDClaimName gets a reference to the given string and assigns it to the ClientIDClaimName field.
func (o *AddJwtAccessTokenValidatorRequest) SetClientIDClaimName(v string) {
	o.ClientIDClaimName = &v
}

// GetScopeClaimName returns the ScopeClaimName field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetScopeClaimName() string {
	if o == nil || IsNil(o.ScopeClaimName) {
		var ret string
		return ret
	}
	return *o.ScopeClaimName
}

// GetScopeClaimNameOk returns a tuple with the ScopeClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetScopeClaimNameOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeClaimName) {
		return nil, false
	}
	return o.ScopeClaimName, true
}

// HasScopeClaimName returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasScopeClaimName() bool {
	if o != nil && !IsNil(o.ScopeClaimName) {
		return true
	}

	return false
}

// SetScopeClaimName gets a reference to the given string and assigns it to the ScopeClaimName field.
func (o *AddJwtAccessTokenValidatorRequest) SetScopeClaimName(v string) {
	o.ScopeClaimName = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetEvaluationOrderIndex() int64 {
	if o == nil || IsNil(o.EvaluationOrderIndex) {
		var ret int64
		return ret
	}
	return *o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetEvaluationOrderIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.EvaluationOrderIndex) {
		return nil, false
	}
	return o.EvaluationOrderIndex, true
}

// HasEvaluationOrderIndex returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasEvaluationOrderIndex() bool {
	if o != nil && !IsNil(o.EvaluationOrderIndex) {
		return true
	}

	return false
}

// SetEvaluationOrderIndex gets a reference to the given int64 and assigns it to the EvaluationOrderIndex field.
func (o *AddJwtAccessTokenValidatorRequest) SetEvaluationOrderIndex(v int64) {
	o.EvaluationOrderIndex = &v
}

// GetAuthorizationServer returns the AuthorizationServer field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetAuthorizationServer() string {
	if o == nil || IsNil(o.AuthorizationServer) {
		var ret string
		return ret
	}
	return *o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetAuthorizationServerOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationServer) {
		return nil, false
	}
	return o.AuthorizationServer, true
}

// HasAuthorizationServer returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasAuthorizationServer() bool {
	if o != nil && !IsNil(o.AuthorizationServer) {
		return true
	}

	return false
}

// SetAuthorizationServer gets a reference to the given string and assigns it to the AuthorizationServer field.
func (o *AddJwtAccessTokenValidatorRequest) SetAuthorizationServer(v string) {
	o.AuthorizationServer = &v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetIdentityMapper() string {
	if o == nil || IsNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasIdentityMapper() bool {
	if o != nil && !IsNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *AddJwtAccessTokenValidatorRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetSubjectClaimName returns the SubjectClaimName field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetSubjectClaimName() string {
	if o == nil || IsNil(o.SubjectClaimName) {
		var ret string
		return ret
	}
	return *o.SubjectClaimName
}

// GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetSubjectClaimNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectClaimName) {
		return nil, false
	}
	return o.SubjectClaimName, true
}

// HasSubjectClaimName returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasSubjectClaimName() bool {
	if o != nil && !IsNil(o.SubjectClaimName) {
		return true
	}

	return false
}

// SetSubjectClaimName gets a reference to the given string and assigns it to the SubjectClaimName field.
func (o *AddJwtAccessTokenValidatorRequest) SetSubjectClaimName(v string) {
	o.SubjectClaimName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddJwtAccessTokenValidatorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddJwtAccessTokenValidatorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddJwtAccessTokenValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddJwtAccessTokenValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddJwtAccessTokenValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddJwtAccessTokenValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddJwtAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddJwtAccessTokenValidatorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["validatorName"] = o.ValidatorName
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.AllowedSigningAlgorithm) {
		toSerialize["allowedSigningAlgorithm"] = o.AllowedSigningAlgorithm
	}
	if !IsNil(o.SigningCertificate) {
		toSerialize["signingCertificate"] = o.SigningCertificate
	}
	if !IsNil(o.JwksEndpointPath) {
		toSerialize["jwksEndpointPath"] = o.JwksEndpointPath
	}
	if !IsNil(o.EncryptionKeyPair) {
		toSerialize["encryptionKeyPair"] = o.EncryptionKeyPair
	}
	if !IsNil(o.AllowedKeyEncryptionAlgorithm) {
		toSerialize["allowedKeyEncryptionAlgorithm"] = o.AllowedKeyEncryptionAlgorithm
	}
	if !IsNil(o.AllowedContentEncryptionAlgorithm) {
		toSerialize["allowedContentEncryptionAlgorithm"] = o.AllowedContentEncryptionAlgorithm
	}
	if !IsNil(o.ClockSkewGracePeriod) {
		toSerialize["clockSkewGracePeriod"] = o.ClockSkewGracePeriod
	}
	if !IsNil(o.ClientIDClaimName) {
		toSerialize["clientIDClaimName"] = o.ClientIDClaimName
	}
	if !IsNil(o.ScopeClaimName) {
		toSerialize["scopeClaimName"] = o.ScopeClaimName
	}
	if !IsNil(o.EvaluationOrderIndex) {
		toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	}
	if !IsNil(o.AuthorizationServer) {
		toSerialize["authorizationServer"] = o.AuthorizationServer
	}
	if !IsNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !IsNil(o.SubjectClaimName) {
		toSerialize["subjectClaimName"] = o.SubjectClaimName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableAddJwtAccessTokenValidatorRequest struct {
	value *AddJwtAccessTokenValidatorRequest
	isSet bool
}

func (v NullableAddJwtAccessTokenValidatorRequest) Get() *AddJwtAccessTokenValidatorRequest {
	return v.value
}

func (v *NullableAddJwtAccessTokenValidatorRequest) Set(val *AddJwtAccessTokenValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddJwtAccessTokenValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddJwtAccessTokenValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddJwtAccessTokenValidatorRequest(val *AddJwtAccessTokenValidatorRequest) *NullableAddJwtAccessTokenValidatorRequest {
	return &NullableAddJwtAccessTokenValidatorRequest{value: val, isSet: true}
}

func (v NullableAddJwtAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddJwtAccessTokenValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
