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

// AddPingFederateAccessTokenValidatorRequest struct for AddPingFederateAccessTokenValidatorRequest
type AddPingFederateAccessTokenValidatorRequest struct {
	// Name of the new Access Token Validator
	ValidatorName string                                          `json:"validatorName"`
	Schemas       []EnumpingFederateAccessTokenValidatorSchemaUrn `json:"schemas"`
	// The client identifier to use when authenticating to the PingFederate authorization server.
	ClientID string `json:"clientID"`
	// The client secret to use when authenticating to the PingFederate authorization server.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The passphrase provider for obtaining the client secret to use when authenticating to the PingFederate authorization server.
	ClientSecretPassphraseProvider *string `json:"clientSecretPassphraseProvider,omitempty"`
	// Whether to include the incoming request URL as the \"aud\" parameter when calling the PingFederate introspection endpoint. This property is ignored if the access-token-manager-id property is set.
	IncludeAudParameter *bool `json:"includeAudParameter,omitempty"`
	// The Access Token Manager instance ID to specify when calling the PingFederate introspection endpoint. If this property is set the include-aud-parameter property is ignored.
	AccessTokenManagerID *string `json:"accessTokenManagerID,omitempty"`
	// How often the Access Token Validator should refresh its stored value of the PingFederate server's token introspection endpoint.
	EndpointCacheRefresh *string `json:"endpointCacheRefresh,omitempty"`
	// When multiple Ping Federate Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Ping Federate Access Token Validators defined within Directory Server but not necessarily contiguous. Ping Federate Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token.
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

// NewAddPingFederateAccessTokenValidatorRequest instantiates a new AddPingFederateAccessTokenValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPingFederateAccessTokenValidatorRequest(validatorName string, schemas []EnumpingFederateAccessTokenValidatorSchemaUrn, clientID string, evaluationOrderIndex int32, enabled bool) *AddPingFederateAccessTokenValidatorRequest {
	this := AddPingFederateAccessTokenValidatorRequest{}
	this.ValidatorName = validatorName
	this.Schemas = schemas
	this.ClientID = clientID
	this.EvaluationOrderIndex = evaluationOrderIndex
	this.Enabled = enabled
	return &this
}

// NewAddPingFederateAccessTokenValidatorRequestWithDefaults instantiates a new AddPingFederateAccessTokenValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPingFederateAccessTokenValidatorRequestWithDefaults() *AddPingFederateAccessTokenValidatorRequest {
	this := AddPingFederateAccessTokenValidatorRequest{}
	return &this
}

// GetValidatorName returns the ValidatorName field value
func (o *AddPingFederateAccessTokenValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddPingFederateAccessTokenValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPingFederateAccessTokenValidatorRequest) GetSchemas() []EnumpingFederateAccessTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnumpingFederateAccessTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetSchemasOk() ([]EnumpingFederateAccessTokenValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPingFederateAccessTokenValidatorRequest) SetSchemas(v []EnumpingFederateAccessTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetClientID returns the ClientID field value
func (o *AddPingFederateAccessTokenValidatorRequest) GetClientID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetClientIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientID, true
}

// SetClientID sets field value
func (o *AddPingFederateAccessTokenValidatorRequest) SetClientID(v string) {
	o.ClientID = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetClientSecret() string {
	if o == nil || isNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetClientSecretOk() (*string, bool) {
	if o == nil || isNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasClientSecret() bool {
	if o != nil && !isNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientSecretPassphraseProvider returns the ClientSecretPassphraseProvider field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetClientSecretPassphraseProvider() string {
	if o == nil || isNil(o.ClientSecretPassphraseProvider) {
		var ret string
		return ret
	}
	return *o.ClientSecretPassphraseProvider
}

// GetClientSecretPassphraseProviderOk returns a tuple with the ClientSecretPassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetClientSecretPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.ClientSecretPassphraseProvider) {
		return nil, false
	}
	return o.ClientSecretPassphraseProvider, true
}

// HasClientSecretPassphraseProvider returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasClientSecretPassphraseProvider() bool {
	if o != nil && !isNil(o.ClientSecretPassphraseProvider) {
		return true
	}

	return false
}

// SetClientSecretPassphraseProvider gets a reference to the given string and assigns it to the ClientSecretPassphraseProvider field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetClientSecretPassphraseProvider(v string) {
	o.ClientSecretPassphraseProvider = &v
}

// GetIncludeAudParameter returns the IncludeAudParameter field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetIncludeAudParameter() bool {
	if o == nil || isNil(o.IncludeAudParameter) {
		var ret bool
		return ret
	}
	return *o.IncludeAudParameter
}

// GetIncludeAudParameterOk returns a tuple with the IncludeAudParameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetIncludeAudParameterOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeAudParameter) {
		return nil, false
	}
	return o.IncludeAudParameter, true
}

// HasIncludeAudParameter returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasIncludeAudParameter() bool {
	if o != nil && !isNil(o.IncludeAudParameter) {
		return true
	}

	return false
}

// SetIncludeAudParameter gets a reference to the given bool and assigns it to the IncludeAudParameter field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetIncludeAudParameter(v bool) {
	o.IncludeAudParameter = &v
}

// GetAccessTokenManagerID returns the AccessTokenManagerID field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetAccessTokenManagerID() string {
	if o == nil || isNil(o.AccessTokenManagerID) {
		var ret string
		return ret
	}
	return *o.AccessTokenManagerID
}

// GetAccessTokenManagerIDOk returns a tuple with the AccessTokenManagerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetAccessTokenManagerIDOk() (*string, bool) {
	if o == nil || isNil(o.AccessTokenManagerID) {
		return nil, false
	}
	return o.AccessTokenManagerID, true
}

// HasAccessTokenManagerID returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasAccessTokenManagerID() bool {
	if o != nil && !isNil(o.AccessTokenManagerID) {
		return true
	}

	return false
}

// SetAccessTokenManagerID gets a reference to the given string and assigns it to the AccessTokenManagerID field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetAccessTokenManagerID(v string) {
	o.AccessTokenManagerID = &v
}

// GetEndpointCacheRefresh returns the EndpointCacheRefresh field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetEndpointCacheRefresh() string {
	if o == nil || isNil(o.EndpointCacheRefresh) {
		var ret string
		return ret
	}
	return *o.EndpointCacheRefresh
}

// GetEndpointCacheRefreshOk returns a tuple with the EndpointCacheRefresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetEndpointCacheRefreshOk() (*string, bool) {
	if o == nil || isNil(o.EndpointCacheRefresh) {
		return nil, false
	}
	return o.EndpointCacheRefresh, true
}

// HasEndpointCacheRefresh returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasEndpointCacheRefresh() bool {
	if o != nil && !isNil(o.EndpointCacheRefresh) {
		return true
	}

	return false
}

// SetEndpointCacheRefresh gets a reference to the given string and assigns it to the EndpointCacheRefresh field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetEndpointCacheRefresh(v string) {
	o.EndpointCacheRefresh = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *AddPingFederateAccessTokenValidatorRequest) GetEvaluationOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *AddPingFederateAccessTokenValidatorRequest) SetEvaluationOrderIndex(v int32) {
	o.EvaluationOrderIndex = v
}

// GetAuthorizationServer returns the AuthorizationServer field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetAuthorizationServer() string {
	if o == nil || isNil(o.AuthorizationServer) {
		var ret string
		return ret
	}
	return *o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetAuthorizationServerOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizationServer) {
		return nil, false
	}
	return o.AuthorizationServer, true
}

// HasAuthorizationServer returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasAuthorizationServer() bool {
	if o != nil && !isNil(o.AuthorizationServer) {
		return true
	}

	return false
}

// SetAuthorizationServer gets a reference to the given string and assigns it to the AuthorizationServer field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetAuthorizationServer(v string) {
	o.AuthorizationServer = &v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetIdentityMapper() string {
	if o == nil || isNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil || isNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasIdentityMapper() bool {
	if o != nil && !isNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetSubjectClaimName returns the SubjectClaimName field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetSubjectClaimName() string {
	if o == nil || isNil(o.SubjectClaimName) {
		var ret string
		return ret
	}
	return *o.SubjectClaimName
}

// GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetSubjectClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectClaimName) {
		return nil, false
	}
	return o.SubjectClaimName, true
}

// HasSubjectClaimName returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasSubjectClaimName() bool {
	if o != nil && !isNil(o.SubjectClaimName) {
		return true
	}

	return false
}

// SetSubjectClaimName gets a reference to the given string and assigns it to the SubjectClaimName field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetSubjectClaimName(v string) {
	o.SubjectClaimName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPingFederateAccessTokenValidatorRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPingFederateAccessTokenValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPingFederateAccessTokenValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPingFederateAccessTokenValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPingFederateAccessTokenValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddPingFederateAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["validatorName"] = o.ValidatorName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["clientID"] = o.ClientID
	}
	if !isNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !isNil(o.ClientSecretPassphraseProvider) {
		toSerialize["clientSecretPassphraseProvider"] = o.ClientSecretPassphraseProvider
	}
	if !isNil(o.IncludeAudParameter) {
		toSerialize["includeAudParameter"] = o.IncludeAudParameter
	}
	if !isNil(o.AccessTokenManagerID) {
		toSerialize["accessTokenManagerID"] = o.AccessTokenManagerID
	}
	if !isNil(o.EndpointCacheRefresh) {
		toSerialize["endpointCacheRefresh"] = o.EndpointCacheRefresh
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

type NullableAddPingFederateAccessTokenValidatorRequest struct {
	value *AddPingFederateAccessTokenValidatorRequest
	isSet bool
}

func (v NullableAddPingFederateAccessTokenValidatorRequest) Get() *AddPingFederateAccessTokenValidatorRequest {
	return v.value
}

func (v *NullableAddPingFederateAccessTokenValidatorRequest) Set(val *AddPingFederateAccessTokenValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPingFederateAccessTokenValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPingFederateAccessTokenValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPingFederateAccessTokenValidatorRequest(val *AddPingFederateAccessTokenValidatorRequest) *NullableAddPingFederateAccessTokenValidatorRequest {
	return &NullableAddPingFederateAccessTokenValidatorRequest{value: val, isSet: true}
}

func (v NullableAddPingFederateAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPingFederateAccessTokenValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
