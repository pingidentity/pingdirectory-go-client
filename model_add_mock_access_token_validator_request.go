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

// AddMockAccessTokenValidatorRequest struct for AddMockAccessTokenValidatorRequest
type AddMockAccessTokenValidatorRequest struct {
	// Name of the new Access Token Validator
	ValidatorName string `json:"validatorName"`
	Schemas []EnummockAccessTokenValidatorSchemaUrn `json:"schemas"`
	// The name of the token claim that contains the OAuth2 client ID.
	ClientIDClaimName *string `json:"clientIDClaimName,omitempty"`
	// The name of the token claim that contains the scopes granted by the token.
	ScopeClaimName *string `json:"scopeClaimName,omitempty"`
	// When multiple Mock Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Mock Access Token Validators defined within Directory Server but not necessarily contiguous. Mock Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token.
	EvaluationOrderIndex int32 `json:"evaluationOrderIndex"`
	// Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server.
	SubjectClaimName *string `json:"subjectClaimName,omitempty"`
	// A description for this Access Token Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether this Access Token Validator is enabled for use in Directory Server.
	Enabled bool `json:"enabled"`
}

// NewAddMockAccessTokenValidatorRequest instantiates a new AddMockAccessTokenValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMockAccessTokenValidatorRequest(validatorName string, schemas []EnummockAccessTokenValidatorSchemaUrn, evaluationOrderIndex int32, enabled bool) *AddMockAccessTokenValidatorRequest {
	this := AddMockAccessTokenValidatorRequest{}
	this.ValidatorName = validatorName
	this.Schemas = schemas
	this.EvaluationOrderIndex = evaluationOrderIndex
	this.Enabled = enabled
	return &this
}

// NewAddMockAccessTokenValidatorRequestWithDefaults instantiates a new AddMockAccessTokenValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMockAccessTokenValidatorRequestWithDefaults() *AddMockAccessTokenValidatorRequest {
	this := AddMockAccessTokenValidatorRequest{}
	return &this
}

// GetValidatorName returns the ValidatorName field value
func (o *AddMockAccessTokenValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddMockAccessTokenValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddMockAccessTokenValidatorRequest) GetSchemas() []EnummockAccessTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnummockAccessTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetSchemasOk() ([]EnummockAccessTokenValidatorSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddMockAccessTokenValidatorRequest) SetSchemas(v []EnummockAccessTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetClientIDClaimName returns the ClientIDClaimName field value if set, zero value otherwise.
func (o *AddMockAccessTokenValidatorRequest) GetClientIDClaimName() string {
	if o == nil || isNil(o.ClientIDClaimName) {
		var ret string
		return ret
	}
	return *o.ClientIDClaimName
}

// GetClientIDClaimNameOk returns a tuple with the ClientIDClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetClientIDClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.ClientIDClaimName) {
    return nil, false
	}
	return o.ClientIDClaimName, true
}

// HasClientIDClaimName returns a boolean if a field has been set.
func (o *AddMockAccessTokenValidatorRequest) HasClientIDClaimName() bool {
	if o != nil && !isNil(o.ClientIDClaimName) {
		return true
	}

	return false
}

// SetClientIDClaimName gets a reference to the given string and assigns it to the ClientIDClaimName field.
func (o *AddMockAccessTokenValidatorRequest) SetClientIDClaimName(v string) {
	o.ClientIDClaimName = &v
}

// GetScopeClaimName returns the ScopeClaimName field value if set, zero value otherwise.
func (o *AddMockAccessTokenValidatorRequest) GetScopeClaimName() string {
	if o == nil || isNil(o.ScopeClaimName) {
		var ret string
		return ret
	}
	return *o.ScopeClaimName
}

// GetScopeClaimNameOk returns a tuple with the ScopeClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetScopeClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.ScopeClaimName) {
    return nil, false
	}
	return o.ScopeClaimName, true
}

// HasScopeClaimName returns a boolean if a field has been set.
func (o *AddMockAccessTokenValidatorRequest) HasScopeClaimName() bool {
	if o != nil && !isNil(o.ScopeClaimName) {
		return true
	}

	return false
}

// SetScopeClaimName gets a reference to the given string and assigns it to the ScopeClaimName field.
func (o *AddMockAccessTokenValidatorRequest) SetScopeClaimName(v string) {
	o.ScopeClaimName = &v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *AddMockAccessTokenValidatorRequest) GetEvaluationOrderIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *AddMockAccessTokenValidatorRequest) SetEvaluationOrderIndex(v int32) {
	o.EvaluationOrderIndex = v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *AddMockAccessTokenValidatorRequest) GetIdentityMapper() string {
	if o == nil || isNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil || isNil(o.IdentityMapper) {
    return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *AddMockAccessTokenValidatorRequest) HasIdentityMapper() bool {
	if o != nil && !isNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *AddMockAccessTokenValidatorRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetSubjectClaimName returns the SubjectClaimName field value if set, zero value otherwise.
func (o *AddMockAccessTokenValidatorRequest) GetSubjectClaimName() string {
	if o == nil || isNil(o.SubjectClaimName) {
		var ret string
		return ret
	}
	return *o.SubjectClaimName
}

// GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetSubjectClaimNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectClaimName) {
    return nil, false
	}
	return o.SubjectClaimName, true
}

// HasSubjectClaimName returns a boolean if a field has been set.
func (o *AddMockAccessTokenValidatorRequest) HasSubjectClaimName() bool {
	if o != nil && !isNil(o.SubjectClaimName) {
		return true
	}

	return false
}

// SetSubjectClaimName gets a reference to the given string and assigns it to the SubjectClaimName field.
func (o *AddMockAccessTokenValidatorRequest) SetSubjectClaimName(v string) {
	o.SubjectClaimName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddMockAccessTokenValidatorRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddMockAccessTokenValidatorRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddMockAccessTokenValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddMockAccessTokenValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddMockAccessTokenValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddMockAccessTokenValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddMockAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["validatorName"] = o.ValidatorName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
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

type NullableAddMockAccessTokenValidatorRequest struct {
	value *AddMockAccessTokenValidatorRequest
	isSet bool
}

func (v NullableAddMockAccessTokenValidatorRequest) Get() *AddMockAccessTokenValidatorRequest {
	return v.value
}

func (v *NullableAddMockAccessTokenValidatorRequest) Set(val *AddMockAccessTokenValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMockAccessTokenValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMockAccessTokenValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMockAccessTokenValidatorRequest(val *AddMockAccessTokenValidatorRequest) *NullableAddMockAccessTokenValidatorRequest {
	return &NullableAddMockAccessTokenValidatorRequest{value: val, isSet: true}
}

func (v NullableAddMockAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMockAccessTokenValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

