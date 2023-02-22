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

// AddRegularExpressionIdentityMapperRequest struct for AddRegularExpressionIdentityMapperRequest
type AddRegularExpressionIdentityMapperRequest struct {
	// Name of the new Identity Mapper
	MapperName string                                         `json:"mapperName"`
	Schemas    []EnumregularExpressionIdentityMapperSchemaUrn `json:"schemas"`
	// Specifies the name or OID of the attribute whose value should match the provided identifier string after it has been processed by the associated regular expression.
	MatchAttribute []string `json:"matchAttribute"`
	// Specifies the base DN(s) that should be used when performing searches to map the provided ID string to a user entry. If multiple values are given, searches are performed below all the specified base DNs.
	MatchBaseDN []string `json:"matchBaseDN,omitempty"`
	// An optional filter that mapped users must match.
	MatchFilter *string `json:"matchFilter,omitempty"`
	// Specifies the regular expression pattern that is used to identify portions of the ID string that will be replaced.
	MatchPattern string `json:"matchPattern"`
	// Specifies the replacement pattern that should be used for substrings in the ID string that match the provided regular expression pattern.
	ReplacePattern *string `json:"replacePattern,omitempty"`
	// A description for this Identity Mapper
	Description *string `json:"description,omitempty"`
	// Indicates whether the Identity Mapper is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddRegularExpressionIdentityMapperRequest instantiates a new AddRegularExpressionIdentityMapperRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRegularExpressionIdentityMapperRequest(mapperName string, schemas []EnumregularExpressionIdentityMapperSchemaUrn, matchAttribute []string, matchPattern string, enabled bool) *AddRegularExpressionIdentityMapperRequest {
	this := AddRegularExpressionIdentityMapperRequest{}
	this.MapperName = mapperName
	this.Schemas = schemas
	this.MatchAttribute = matchAttribute
	this.MatchPattern = matchPattern
	this.Enabled = enabled
	return &this
}

// NewAddRegularExpressionIdentityMapperRequestWithDefaults instantiates a new AddRegularExpressionIdentityMapperRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRegularExpressionIdentityMapperRequestWithDefaults() *AddRegularExpressionIdentityMapperRequest {
	this := AddRegularExpressionIdentityMapperRequest{}
	return &this
}

// GetMapperName returns the MapperName field value
func (o *AddRegularExpressionIdentityMapperRequest) GetMapperName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MapperName
}

// GetMapperNameOk returns a tuple with the MapperName field value
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetMapperNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MapperName, true
}

// SetMapperName sets field value
func (o *AddRegularExpressionIdentityMapperRequest) SetMapperName(v string) {
	o.MapperName = v
}

// GetSchemas returns the Schemas field value
func (o *AddRegularExpressionIdentityMapperRequest) GetSchemas() []EnumregularExpressionIdentityMapperSchemaUrn {
	if o == nil {
		var ret []EnumregularExpressionIdentityMapperSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetSchemasOk() ([]EnumregularExpressionIdentityMapperSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddRegularExpressionIdentityMapperRequest) SetSchemas(v []EnumregularExpressionIdentityMapperSchemaUrn) {
	o.Schemas = v
}

// GetMatchAttribute returns the MatchAttribute field value
func (o *AddRegularExpressionIdentityMapperRequest) GetMatchAttribute() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MatchAttribute
}

// GetMatchAttributeOk returns a tuple with the MatchAttribute field value
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetMatchAttributeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MatchAttribute, true
}

// SetMatchAttribute sets field value
func (o *AddRegularExpressionIdentityMapperRequest) SetMatchAttribute(v []string) {
	o.MatchAttribute = v
}

// GetMatchBaseDN returns the MatchBaseDN field value if set, zero value otherwise.
func (o *AddRegularExpressionIdentityMapperRequest) GetMatchBaseDN() []string {
	if o == nil || isNil(o.MatchBaseDN) {
		var ret []string
		return ret
	}
	return o.MatchBaseDN
}

// GetMatchBaseDNOk returns a tuple with the MatchBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetMatchBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.MatchBaseDN) {
		return nil, false
	}
	return o.MatchBaseDN, true
}

// HasMatchBaseDN returns a boolean if a field has been set.
func (o *AddRegularExpressionIdentityMapperRequest) HasMatchBaseDN() bool {
	if o != nil && !isNil(o.MatchBaseDN) {
		return true
	}

	return false
}

// SetMatchBaseDN gets a reference to the given []string and assigns it to the MatchBaseDN field.
func (o *AddRegularExpressionIdentityMapperRequest) SetMatchBaseDN(v []string) {
	o.MatchBaseDN = v
}

// GetMatchFilter returns the MatchFilter field value if set, zero value otherwise.
func (o *AddRegularExpressionIdentityMapperRequest) GetMatchFilter() string {
	if o == nil || isNil(o.MatchFilter) {
		var ret string
		return ret
	}
	return *o.MatchFilter
}

// GetMatchFilterOk returns a tuple with the MatchFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetMatchFilterOk() (*string, bool) {
	if o == nil || isNil(o.MatchFilter) {
		return nil, false
	}
	return o.MatchFilter, true
}

// HasMatchFilter returns a boolean if a field has been set.
func (o *AddRegularExpressionIdentityMapperRequest) HasMatchFilter() bool {
	if o != nil && !isNil(o.MatchFilter) {
		return true
	}

	return false
}

// SetMatchFilter gets a reference to the given string and assigns it to the MatchFilter field.
func (o *AddRegularExpressionIdentityMapperRequest) SetMatchFilter(v string) {
	o.MatchFilter = &v
}

// GetMatchPattern returns the MatchPattern field value
func (o *AddRegularExpressionIdentityMapperRequest) GetMatchPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MatchPattern
}

// GetMatchPatternOk returns a tuple with the MatchPattern field value
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetMatchPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchPattern, true
}

// SetMatchPattern sets field value
func (o *AddRegularExpressionIdentityMapperRequest) SetMatchPattern(v string) {
	o.MatchPattern = v
}

// GetReplacePattern returns the ReplacePattern field value if set, zero value otherwise.
func (o *AddRegularExpressionIdentityMapperRequest) GetReplacePattern() string {
	if o == nil || isNil(o.ReplacePattern) {
		var ret string
		return ret
	}
	return *o.ReplacePattern
}

// GetReplacePatternOk returns a tuple with the ReplacePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetReplacePatternOk() (*string, bool) {
	if o == nil || isNil(o.ReplacePattern) {
		return nil, false
	}
	return o.ReplacePattern, true
}

// HasReplacePattern returns a boolean if a field has been set.
func (o *AddRegularExpressionIdentityMapperRequest) HasReplacePattern() bool {
	if o != nil && !isNil(o.ReplacePattern) {
		return true
	}

	return false
}

// SetReplacePattern gets a reference to the given string and assigns it to the ReplacePattern field.
func (o *AddRegularExpressionIdentityMapperRequest) SetReplacePattern(v string) {
	o.ReplacePattern = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddRegularExpressionIdentityMapperRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddRegularExpressionIdentityMapperRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddRegularExpressionIdentityMapperRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddRegularExpressionIdentityMapperRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddRegularExpressionIdentityMapperRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddRegularExpressionIdentityMapperRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddRegularExpressionIdentityMapperRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mapperName"] = o.MapperName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["matchAttribute"] = o.MatchAttribute
	}
	if !isNil(o.MatchBaseDN) {
		toSerialize["matchBaseDN"] = o.MatchBaseDN
	}
	if !isNil(o.MatchFilter) {
		toSerialize["matchFilter"] = o.MatchFilter
	}
	if true {
		toSerialize["matchPattern"] = o.MatchPattern
	}
	if !isNil(o.ReplacePattern) {
		toSerialize["replacePattern"] = o.ReplacePattern
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddRegularExpressionIdentityMapperRequest struct {
	value *AddRegularExpressionIdentityMapperRequest
	isSet bool
}

func (v NullableAddRegularExpressionIdentityMapperRequest) Get() *AddRegularExpressionIdentityMapperRequest {
	return v.value
}

func (v *NullableAddRegularExpressionIdentityMapperRequest) Set(val *AddRegularExpressionIdentityMapperRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRegularExpressionIdentityMapperRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRegularExpressionIdentityMapperRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRegularExpressionIdentityMapperRequest(val *AddRegularExpressionIdentityMapperRequest) *NullableAddRegularExpressionIdentityMapperRequest {
	return &NullableAddRegularExpressionIdentityMapperRequest{value: val, isSet: true}
}

func (v NullableAddRegularExpressionIdentityMapperRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRegularExpressionIdentityMapperRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}