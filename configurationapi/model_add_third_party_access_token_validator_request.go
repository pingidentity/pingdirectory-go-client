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

// checks if the AddThirdPartyAccessTokenValidatorRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyAccessTokenValidatorRequest{}

// AddThirdPartyAccessTokenValidatorRequest struct for AddThirdPartyAccessTokenValidatorRequest
type AddThirdPartyAccessTokenValidatorRequest struct {
	// Name of the new Access Token Validator
	ValidatorName string                                        `json:"validatorName"`
	Schemas       []EnumthirdPartyAccessTokenValidatorSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Access Token Validator.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Access Token Validator. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server.
	SubjectClaimName *string `json:"subjectClaimName,omitempty"`
	// A description for this Access Token Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether this Access Token Validator is enabled for use in Directory Server.
	Enabled bool `json:"enabled"`
	// When multiple Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Access Token Validators defined within Directory Server but not necessarily contiguous. Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token.
	EvaluationOrderIndex int64 `json:"evaluationOrderIndex"`
}

// NewAddThirdPartyAccessTokenValidatorRequest instantiates a new AddThirdPartyAccessTokenValidatorRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyAccessTokenValidatorRequest(validatorName string, schemas []EnumthirdPartyAccessTokenValidatorSchemaUrn, extensionClass string, enabled bool, evaluationOrderIndex int64) *AddThirdPartyAccessTokenValidatorRequest {
	this := AddThirdPartyAccessTokenValidatorRequest{}
	this.ValidatorName = validatorName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.EvaluationOrderIndex = evaluationOrderIndex
	return &this
}

// NewAddThirdPartyAccessTokenValidatorRequestWithDefaults instantiates a new AddThirdPartyAccessTokenValidatorRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyAccessTokenValidatorRequestWithDefaults() *AddThirdPartyAccessTokenValidatorRequest {
	this := AddThirdPartyAccessTokenValidatorRequest{}
	return &this
}

// GetValidatorName returns the ValidatorName field value
func (o *AddThirdPartyAccessTokenValidatorRequest) GetValidatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ValidatorName
}

// GetValidatorNameOk returns a tuple with the ValidatorName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetValidatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorName, true
}

// SetValidatorName sets field value
func (o *AddThirdPartyAccessTokenValidatorRequest) SetValidatorName(v string) {
	o.ValidatorName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyAccessTokenValidatorRequest) GetSchemas() []EnumthirdPartyAccessTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyAccessTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetSchemasOk() ([]EnumthirdPartyAccessTokenValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyAccessTokenValidatorRequest) SetSchemas(v []EnumthirdPartyAccessTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyAccessTokenValidatorRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyAccessTokenValidatorRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyAccessTokenValidatorRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetIdentityMapper() string {
	if o == nil || IsNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetIdentityMapperOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) HasIdentityMapper() bool {
	if o != nil && !IsNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *AddThirdPartyAccessTokenValidatorRequest) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetSubjectClaimName returns the SubjectClaimName field value if set, zero value otherwise.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetSubjectClaimName() string {
	if o == nil || IsNil(o.SubjectClaimName) {
		var ret string
		return ret
	}
	return *o.SubjectClaimName
}

// GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetSubjectClaimNameOk() (*string, bool) {
	if o == nil || IsNil(o.SubjectClaimName) {
		return nil, false
	}
	return o.SubjectClaimName, true
}

// HasSubjectClaimName returns a boolean if a field has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) HasSubjectClaimName() bool {
	if o != nil && !IsNil(o.SubjectClaimName) {
		return true
	}

	return false
}

// SetSubjectClaimName gets a reference to the given string and assigns it to the SubjectClaimName field.
func (o *AddThirdPartyAccessTokenValidatorRequest) SetSubjectClaimName(v string) {
	o.SubjectClaimName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyAccessTokenValidatorRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyAccessTokenValidatorRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyAccessTokenValidatorRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *AddThirdPartyAccessTokenValidatorRequest) GetEvaluationOrderIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessTokenValidatorRequest) GetEvaluationOrderIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *AddThirdPartyAccessTokenValidatorRequest) SetEvaluationOrderIndex(v int64) {
	o.EvaluationOrderIndex = v
}

func (o AddThirdPartyAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyAccessTokenValidatorRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["validatorName"] = o.ValidatorName
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
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
	toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	return toSerialize, nil
}

type NullableAddThirdPartyAccessTokenValidatorRequest struct {
	value *AddThirdPartyAccessTokenValidatorRequest
	isSet bool
}

func (v NullableAddThirdPartyAccessTokenValidatorRequest) Get() *AddThirdPartyAccessTokenValidatorRequest {
	return v.value
}

func (v *NullableAddThirdPartyAccessTokenValidatorRequest) Set(val *AddThirdPartyAccessTokenValidatorRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyAccessTokenValidatorRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyAccessTokenValidatorRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyAccessTokenValidatorRequest(val *AddThirdPartyAccessTokenValidatorRequest) *NullableAddThirdPartyAccessTokenValidatorRequest {
	return &NullableAddThirdPartyAccessTokenValidatorRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyAccessTokenValidatorRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyAccessTokenValidatorRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
