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

// checks if the AggregateSearchReferenceCriteriaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AggregateSearchReferenceCriteriaResponse{}

// AggregateSearchReferenceCriteriaResponse struct for AggregateSearchReferenceCriteriaResponse
type AggregateSearchReferenceCriteriaResponse struct {
	Schemas []EnumaggregateSearchReferenceCriteriaSchemaUrn `json:"schemas"`
	// Specifies a search reference criteria object that must match the associated search result reference in order to match the aggregate search reference criteria. If one or more all-included search reference criteria objects are provided, then a search result reference must match all of them in order to match the aggregate search reference criteria.
	AllIncludedSearchReferenceCriteria []string `json:"allIncludedSearchReferenceCriteria,omitempty"`
	// Specifies a search reference criteria object that may match the associated search result reference in order to match the aggregate search reference criteria. If one or more any-included search reference criteria objects are provided, then a search result reference must match at least one of them in order to match the aggregate search reference criteria.
	AnyIncludedSearchReferenceCriteria []string `json:"anyIncludedSearchReferenceCriteria,omitempty"`
	// Specifies a search reference criteria object that should not match the associated search result reference in order to match the aggregate search reference criteria. If one or more not-all-included search reference criteria objects are provided, then a search result reference must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate search reference criteria.
	NotAllIncludedSearchReferenceCriteria []string `json:"notAllIncludedSearchReferenceCriteria,omitempty"`
	// Specifies a search reference criteria object that must not match the associated search result reference in order to match the aggregate search reference criteria. If one or more none-included search reference criteria objects are provided, then a search result reference must not match any of them in order to match the aggregate search reference criteria.
	NoneIncludedSearchReferenceCriteria []string `json:"noneIncludedSearchReferenceCriteria,omitempty"`
	// A description for this Search Reference Criteria
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Search Reference Criteria
	Id string `json:"id"`
}

// NewAggregateSearchReferenceCriteriaResponse instantiates a new AggregateSearchReferenceCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateSearchReferenceCriteriaResponse(schemas []EnumaggregateSearchReferenceCriteriaSchemaUrn, id string) *AggregateSearchReferenceCriteriaResponse {
	this := AggregateSearchReferenceCriteriaResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewAggregateSearchReferenceCriteriaResponseWithDefaults instantiates a new AggregateSearchReferenceCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateSearchReferenceCriteriaResponseWithDefaults() *AggregateSearchReferenceCriteriaResponse {
	this := AggregateSearchReferenceCriteriaResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AggregateSearchReferenceCriteriaResponse) GetSchemas() []EnumaggregateSearchReferenceCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumaggregateSearchReferenceCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetSchemasOk() ([]EnumaggregateSearchReferenceCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AggregateSearchReferenceCriteriaResponse) SetSchemas(v []EnumaggregateSearchReferenceCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetAllIncludedSearchReferenceCriteria returns the AllIncludedSearchReferenceCriteria field value if set, zero value otherwise.
func (o *AggregateSearchReferenceCriteriaResponse) GetAllIncludedSearchReferenceCriteria() []string {
	if o == nil || IsNil(o.AllIncludedSearchReferenceCriteria) {
		var ret []string
		return ret
	}
	return o.AllIncludedSearchReferenceCriteria
}

// GetAllIncludedSearchReferenceCriteriaOk returns a tuple with the AllIncludedSearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetAllIncludedSearchReferenceCriteriaOk() ([]string, bool) {
	if o == nil || IsNil(o.AllIncludedSearchReferenceCriteria) {
		return nil, false
	}
	return o.AllIncludedSearchReferenceCriteria, true
}

// HasAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.
func (o *AggregateSearchReferenceCriteriaResponse) HasAllIncludedSearchReferenceCriteria() bool {
	if o != nil && !IsNil(o.AllIncludedSearchReferenceCriteria) {
		return true
	}

	return false
}

// SetAllIncludedSearchReferenceCriteria gets a reference to the given []string and assigns it to the AllIncludedSearchReferenceCriteria field.
func (o *AggregateSearchReferenceCriteriaResponse) SetAllIncludedSearchReferenceCriteria(v []string) {
	o.AllIncludedSearchReferenceCriteria = v
}

// GetAnyIncludedSearchReferenceCriteria returns the AnyIncludedSearchReferenceCriteria field value if set, zero value otherwise.
func (o *AggregateSearchReferenceCriteriaResponse) GetAnyIncludedSearchReferenceCriteria() []string {
	if o == nil || IsNil(o.AnyIncludedSearchReferenceCriteria) {
		var ret []string
		return ret
	}
	return o.AnyIncludedSearchReferenceCriteria
}

// GetAnyIncludedSearchReferenceCriteriaOk returns a tuple with the AnyIncludedSearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetAnyIncludedSearchReferenceCriteriaOk() ([]string, bool) {
	if o == nil || IsNil(o.AnyIncludedSearchReferenceCriteria) {
		return nil, false
	}
	return o.AnyIncludedSearchReferenceCriteria, true
}

// HasAnyIncludedSearchReferenceCriteria returns a boolean if a field has been set.
func (o *AggregateSearchReferenceCriteriaResponse) HasAnyIncludedSearchReferenceCriteria() bool {
	if o != nil && !IsNil(o.AnyIncludedSearchReferenceCriteria) {
		return true
	}

	return false
}

// SetAnyIncludedSearchReferenceCriteria gets a reference to the given []string and assigns it to the AnyIncludedSearchReferenceCriteria field.
func (o *AggregateSearchReferenceCriteriaResponse) SetAnyIncludedSearchReferenceCriteria(v []string) {
	o.AnyIncludedSearchReferenceCriteria = v
}

// GetNotAllIncludedSearchReferenceCriteria returns the NotAllIncludedSearchReferenceCriteria field value if set, zero value otherwise.
func (o *AggregateSearchReferenceCriteriaResponse) GetNotAllIncludedSearchReferenceCriteria() []string {
	if o == nil || IsNil(o.NotAllIncludedSearchReferenceCriteria) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedSearchReferenceCriteria
}

// GetNotAllIncludedSearchReferenceCriteriaOk returns a tuple with the NotAllIncludedSearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetNotAllIncludedSearchReferenceCriteriaOk() ([]string, bool) {
	if o == nil || IsNil(o.NotAllIncludedSearchReferenceCriteria) {
		return nil, false
	}
	return o.NotAllIncludedSearchReferenceCriteria, true
}

// HasNotAllIncludedSearchReferenceCriteria returns a boolean if a field has been set.
func (o *AggregateSearchReferenceCriteriaResponse) HasNotAllIncludedSearchReferenceCriteria() bool {
	if o != nil && !IsNil(o.NotAllIncludedSearchReferenceCriteria) {
		return true
	}

	return false
}

// SetNotAllIncludedSearchReferenceCriteria gets a reference to the given []string and assigns it to the NotAllIncludedSearchReferenceCriteria field.
func (o *AggregateSearchReferenceCriteriaResponse) SetNotAllIncludedSearchReferenceCriteria(v []string) {
	o.NotAllIncludedSearchReferenceCriteria = v
}

// GetNoneIncludedSearchReferenceCriteria returns the NoneIncludedSearchReferenceCriteria field value if set, zero value otherwise.
func (o *AggregateSearchReferenceCriteriaResponse) GetNoneIncludedSearchReferenceCriteria() []string {
	if o == nil || IsNil(o.NoneIncludedSearchReferenceCriteria) {
		var ret []string
		return ret
	}
	return o.NoneIncludedSearchReferenceCriteria
}

// GetNoneIncludedSearchReferenceCriteriaOk returns a tuple with the NoneIncludedSearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetNoneIncludedSearchReferenceCriteriaOk() ([]string, bool) {
	if o == nil || IsNil(o.NoneIncludedSearchReferenceCriteria) {
		return nil, false
	}
	return o.NoneIncludedSearchReferenceCriteria, true
}

// HasNoneIncludedSearchReferenceCriteria returns a boolean if a field has been set.
func (o *AggregateSearchReferenceCriteriaResponse) HasNoneIncludedSearchReferenceCriteria() bool {
	if o != nil && !IsNil(o.NoneIncludedSearchReferenceCriteria) {
		return true
	}

	return false
}

// SetNoneIncludedSearchReferenceCriteria gets a reference to the given []string and assigns it to the NoneIncludedSearchReferenceCriteria field.
func (o *AggregateSearchReferenceCriteriaResponse) SetNoneIncludedSearchReferenceCriteria(v []string) {
	o.NoneIncludedSearchReferenceCriteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AggregateSearchReferenceCriteriaResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AggregateSearchReferenceCriteriaResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AggregateSearchReferenceCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AggregateSearchReferenceCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AggregateSearchReferenceCriteriaResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AggregateSearchReferenceCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AggregateSearchReferenceCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AggregateSearchReferenceCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AggregateSearchReferenceCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AggregateSearchReferenceCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AggregateSearchReferenceCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AggregateSearchReferenceCriteriaResponse) SetId(v string) {
	o.Id = v
}

func (o AggregateSearchReferenceCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AggregateSearchReferenceCriteriaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.AllIncludedSearchReferenceCriteria) {
		toSerialize["allIncludedSearchReferenceCriteria"] = o.AllIncludedSearchReferenceCriteria
	}
	if !IsNil(o.AnyIncludedSearchReferenceCriteria) {
		toSerialize["anyIncludedSearchReferenceCriteria"] = o.AnyIncludedSearchReferenceCriteria
	}
	if !IsNil(o.NotAllIncludedSearchReferenceCriteria) {
		toSerialize["notAllIncludedSearchReferenceCriteria"] = o.NotAllIncludedSearchReferenceCriteria
	}
	if !IsNil(o.NoneIncludedSearchReferenceCriteria) {
		toSerialize["noneIncludedSearchReferenceCriteria"] = o.NoneIncludedSearchReferenceCriteria
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAggregateSearchReferenceCriteriaResponse struct {
	value *AggregateSearchReferenceCriteriaResponse
	isSet bool
}

func (v NullableAggregateSearchReferenceCriteriaResponse) Get() *AggregateSearchReferenceCriteriaResponse {
	return v.value
}

func (v *NullableAggregateSearchReferenceCriteriaResponse) Set(val *AggregateSearchReferenceCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateSearchReferenceCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateSearchReferenceCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateSearchReferenceCriteriaResponse(val *AggregateSearchReferenceCriteriaResponse) *NullableAggregateSearchReferenceCriteriaResponse {
	return &NullableAggregateSearchReferenceCriteriaResponse{value: val, isSet: true}
}

func (v NullableAggregateSearchReferenceCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateSearchReferenceCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
