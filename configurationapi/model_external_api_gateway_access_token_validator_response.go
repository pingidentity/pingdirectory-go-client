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

// checks if the ExternalApiGatewayAccessTokenValidatorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalApiGatewayAccessTokenValidatorResponse{}

// ExternalApiGatewayAccessTokenValidatorResponse struct for ExternalApiGatewayAccessTokenValidatorResponse
type ExternalApiGatewayAccessTokenValidatorResponse struct {
	Meta                                          *MetaMeta                                             `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20    `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumexternalApiGatewayAccessTokenValidatorSchemaUrn `json:"schemas"`
	// Name of the Access Token Validator
	Id string `json:"id"`
	// When multiple External API Gateway Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all External API Gateway Access Token Validators defined within Directory Server but not necessarily contiguous. External API Gateway Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token.
	EvaluationOrderIndex int64 `json:"evaluationOrderIndex"`
	// Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property.
	IdentityMapper *string `json:"identityMapper,omitempty"`
	// A description for this Access Token Validator
	Description *string `json:"description,omitempty"`
	// Indicates whether this Access Token Validator is enabled for use in Directory Server.
	Enabled bool `json:"enabled"`
}

// NewExternalApiGatewayAccessTokenValidatorResponse instantiates a new ExternalApiGatewayAccessTokenValidatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalApiGatewayAccessTokenValidatorResponse(schemas []EnumexternalApiGatewayAccessTokenValidatorSchemaUrn, id string, evaluationOrderIndex int64, enabled bool) *ExternalApiGatewayAccessTokenValidatorResponse {
	this := ExternalApiGatewayAccessTokenValidatorResponse{}
	this.Schemas = schemas
	this.Id = id
	this.EvaluationOrderIndex = evaluationOrderIndex
	this.Enabled = enabled
	return &this
}

// NewExternalApiGatewayAccessTokenValidatorResponseWithDefaults instantiates a new ExternalApiGatewayAccessTokenValidatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalApiGatewayAccessTokenValidatorResponseWithDefaults() *ExternalApiGatewayAccessTokenValidatorResponse {
	this := ExternalApiGatewayAccessTokenValidatorResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetSchemas() []EnumexternalApiGatewayAccessTokenValidatorSchemaUrn {
	if o == nil {
		var ret []EnumexternalApiGatewayAccessTokenValidatorSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetSchemasOk() ([]EnumexternalApiGatewayAccessTokenValidatorSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ExternalApiGatewayAccessTokenValidatorResponse) SetSchemas(v []EnumexternalApiGatewayAccessTokenValidatorSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalApiGatewayAccessTokenValidatorResponse) SetId(v string) {
	o.Id = v
}

// GetEvaluationOrderIndex returns the EvaluationOrderIndex field value
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetEvaluationOrderIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EvaluationOrderIndex
}

// GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field value
// and a boolean to check if the value has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetEvaluationOrderIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationOrderIndex, true
}

// SetEvaluationOrderIndex sets field value
func (o *ExternalApiGatewayAccessTokenValidatorResponse) SetEvaluationOrderIndex(v int64) {
	o.EvaluationOrderIndex = v
}

// GetIdentityMapper returns the IdentityMapper field value if set, zero value otherwise.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetIdentityMapper() string {
	if o == nil || IsNil(o.IdentityMapper) {
		var ret string
		return ret
	}
	return *o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityMapper) {
		return nil, false
	}
	return o.IdentityMapper, true
}

// HasIdentityMapper returns a boolean if a field has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) HasIdentityMapper() bool {
	if o != nil && !IsNil(o.IdentityMapper) {
		return true
	}

	return false
}

// SetIdentityMapper gets a reference to the given string and assigns it to the IdentityMapper field.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ExternalApiGatewayAccessTokenValidatorResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ExternalApiGatewayAccessTokenValidatorResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ExternalApiGatewayAccessTokenValidatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalApiGatewayAccessTokenValidatorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["evaluationOrderIndex"] = o.EvaluationOrderIndex
	if !IsNil(o.IdentityMapper) {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableExternalApiGatewayAccessTokenValidatorResponse struct {
	value *ExternalApiGatewayAccessTokenValidatorResponse
	isSet bool
}

func (v NullableExternalApiGatewayAccessTokenValidatorResponse) Get() *ExternalApiGatewayAccessTokenValidatorResponse {
	return v.value
}

func (v *NullableExternalApiGatewayAccessTokenValidatorResponse) Set(val *ExternalApiGatewayAccessTokenValidatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalApiGatewayAccessTokenValidatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalApiGatewayAccessTokenValidatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalApiGatewayAccessTokenValidatorResponse(val *ExternalApiGatewayAccessTokenValidatorResponse) *NullableExternalApiGatewayAccessTokenValidatorResponse {
	return &NullableExternalApiGatewayAccessTokenValidatorResponse{value: val, isSet: true}
}

func (v NullableExternalApiGatewayAccessTokenValidatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalApiGatewayAccessTokenValidatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
