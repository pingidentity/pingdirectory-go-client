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

// AggregateResultCriteriaResponse struct for AggregateResultCriteriaResponse
type AggregateResultCriteriaResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Result Criteria
	Id      string                                 `json:"id"`
	Schemas []EnumaggregateResultCriteriaSchemaUrn `json:"schemas"`
	// Specifies a result criteria object that must match the associated operation result in order to match the aggregate result criteria. If one or more all-included result criteria objects are provided, then an operation result must match all of them in order to match the aggregate result criteria.
	AllIncludedResultCriteria []string `json:"allIncludedResultCriteria,omitempty"`
	// Specifies a result criteria object that may match the associated operation result in order to match the aggregate result criteria. If one or more any-included result criteria objects are provided, then an operation result must match at least one of them in order to match the aggregate result criteria.
	AnyIncludedResultCriteria []string `json:"anyIncludedResultCriteria,omitempty"`
	// Specifies a result criteria object that should not match the associated operation result in order to match the aggregate result criteria. If one or more not-all-included result criteria objects are provided, then an operation result must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate result criteria.
	NotAllIncludedResultCriteria []string `json:"notAllIncludedResultCriteria,omitempty"`
	// Specifies a result criteria object that must not match the associated operation result in order to match the aggregate result criteria. If one or more none-included result criteria objects are provided, then an operation result must not match any of them in order to match the aggregate result criteria.
	NoneIncludedResultCriteria []string `json:"noneIncludedResultCriteria,omitempty"`
	// A description for this Result Criteria
	Description *string `json:"description,omitempty"`
}

// NewAggregateResultCriteriaResponse instantiates a new AggregateResultCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregateResultCriteriaResponse(id string, schemas []EnumaggregateResultCriteriaSchemaUrn) *AggregateResultCriteriaResponse {
	this := AggregateResultCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	return &this
}

// NewAggregateResultCriteriaResponseWithDefaults instantiates a new AggregateResultCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregateResultCriteriaResponseWithDefaults() *AggregateResultCriteriaResponse {
	this := AggregateResultCriteriaResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AggregateResultCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AggregateResultCriteriaResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *AggregateResultCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *AggregateResultCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *AggregateResultCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *AggregateResultCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *AggregateResultCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AggregateResultCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *AggregateResultCriteriaResponse) GetSchemas() []EnumaggregateResultCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumaggregateResultCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetSchemasOk() ([]EnumaggregateResultCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AggregateResultCriteriaResponse) SetSchemas(v []EnumaggregateResultCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetAllIncludedResultCriteria returns the AllIncludedResultCriteria field value if set, zero value otherwise.
func (o *AggregateResultCriteriaResponse) GetAllIncludedResultCriteria() []string {
	if o == nil || isNil(o.AllIncludedResultCriteria) {
		var ret []string
		return ret
	}
	return o.AllIncludedResultCriteria
}

// GetAllIncludedResultCriteriaOk returns a tuple with the AllIncludedResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetAllIncludedResultCriteriaOk() ([]string, bool) {
	if o == nil || isNil(o.AllIncludedResultCriteria) {
		return nil, false
	}
	return o.AllIncludedResultCriteria, true
}

// HasAllIncludedResultCriteria returns a boolean if a field has been set.
func (o *AggregateResultCriteriaResponse) HasAllIncludedResultCriteria() bool {
	if o != nil && !isNil(o.AllIncludedResultCriteria) {
		return true
	}

	return false
}

// SetAllIncludedResultCriteria gets a reference to the given []string and assigns it to the AllIncludedResultCriteria field.
func (o *AggregateResultCriteriaResponse) SetAllIncludedResultCriteria(v []string) {
	o.AllIncludedResultCriteria = v
}

// GetAnyIncludedResultCriteria returns the AnyIncludedResultCriteria field value if set, zero value otherwise.
func (o *AggregateResultCriteriaResponse) GetAnyIncludedResultCriteria() []string {
	if o == nil || isNil(o.AnyIncludedResultCriteria) {
		var ret []string
		return ret
	}
	return o.AnyIncludedResultCriteria
}

// GetAnyIncludedResultCriteriaOk returns a tuple with the AnyIncludedResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetAnyIncludedResultCriteriaOk() ([]string, bool) {
	if o == nil || isNil(o.AnyIncludedResultCriteria) {
		return nil, false
	}
	return o.AnyIncludedResultCriteria, true
}

// HasAnyIncludedResultCriteria returns a boolean if a field has been set.
func (o *AggregateResultCriteriaResponse) HasAnyIncludedResultCriteria() bool {
	if o != nil && !isNil(o.AnyIncludedResultCriteria) {
		return true
	}

	return false
}

// SetAnyIncludedResultCriteria gets a reference to the given []string and assigns it to the AnyIncludedResultCriteria field.
func (o *AggregateResultCriteriaResponse) SetAnyIncludedResultCriteria(v []string) {
	o.AnyIncludedResultCriteria = v
}

// GetNotAllIncludedResultCriteria returns the NotAllIncludedResultCriteria field value if set, zero value otherwise.
func (o *AggregateResultCriteriaResponse) GetNotAllIncludedResultCriteria() []string {
	if o == nil || isNil(o.NotAllIncludedResultCriteria) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedResultCriteria
}

// GetNotAllIncludedResultCriteriaOk returns a tuple with the NotAllIncludedResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetNotAllIncludedResultCriteriaOk() ([]string, bool) {
	if o == nil || isNil(o.NotAllIncludedResultCriteria) {
		return nil, false
	}
	return o.NotAllIncludedResultCriteria, true
}

// HasNotAllIncludedResultCriteria returns a boolean if a field has been set.
func (o *AggregateResultCriteriaResponse) HasNotAllIncludedResultCriteria() bool {
	if o != nil && !isNil(o.NotAllIncludedResultCriteria) {
		return true
	}

	return false
}

// SetNotAllIncludedResultCriteria gets a reference to the given []string and assigns it to the NotAllIncludedResultCriteria field.
func (o *AggregateResultCriteriaResponse) SetNotAllIncludedResultCriteria(v []string) {
	o.NotAllIncludedResultCriteria = v
}

// GetNoneIncludedResultCriteria returns the NoneIncludedResultCriteria field value if set, zero value otherwise.
func (o *AggregateResultCriteriaResponse) GetNoneIncludedResultCriteria() []string {
	if o == nil || isNil(o.NoneIncludedResultCriteria) {
		var ret []string
		return ret
	}
	return o.NoneIncludedResultCriteria
}

// GetNoneIncludedResultCriteriaOk returns a tuple with the NoneIncludedResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetNoneIncludedResultCriteriaOk() ([]string, bool) {
	if o == nil || isNil(o.NoneIncludedResultCriteria) {
		return nil, false
	}
	return o.NoneIncludedResultCriteria, true
}

// HasNoneIncludedResultCriteria returns a boolean if a field has been set.
func (o *AggregateResultCriteriaResponse) HasNoneIncludedResultCriteria() bool {
	if o != nil && !isNil(o.NoneIncludedResultCriteria) {
		return true
	}

	return false
}

// SetNoneIncludedResultCriteria gets a reference to the given []string and assigns it to the NoneIncludedResultCriteria field.
func (o *AggregateResultCriteriaResponse) SetNoneIncludedResultCriteria(v []string) {
	o.NoneIncludedResultCriteria = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AggregateResultCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AggregateResultCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AggregateResultCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AggregateResultCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

func (o AggregateResultCriteriaResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.AllIncludedResultCriteria) {
		toSerialize["allIncludedResultCriteria"] = o.AllIncludedResultCriteria
	}
	if !isNil(o.AnyIncludedResultCriteria) {
		toSerialize["anyIncludedResultCriteria"] = o.AnyIncludedResultCriteria
	}
	if !isNil(o.NotAllIncludedResultCriteria) {
		toSerialize["notAllIncludedResultCriteria"] = o.NotAllIncludedResultCriteria
	}
	if !isNil(o.NoneIncludedResultCriteria) {
		toSerialize["noneIncludedResultCriteria"] = o.NoneIncludedResultCriteria
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAggregateResultCriteriaResponse struct {
	value *AggregateResultCriteriaResponse
	isSet bool
}

func (v NullableAggregateResultCriteriaResponse) Get() *AggregateResultCriteriaResponse {
	return v.value
}

func (v *NullableAggregateResultCriteriaResponse) Set(val *AggregateResultCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregateResultCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregateResultCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregateResultCriteriaResponse(val *AggregateResultCriteriaResponse) *NullableAggregateResultCriteriaResponse {
	return &NullableAggregateResultCriteriaResponse{value: val, isSet: true}
}

func (v NullableAggregateResultCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregateResultCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
