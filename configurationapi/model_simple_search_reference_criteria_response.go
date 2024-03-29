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

// checks if the SimpleSearchReferenceCriteriaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimpleSearchReferenceCriteriaResponse{}

// SimpleSearchReferenceCriteriaResponse struct for SimpleSearchReferenceCriteriaResponse
type SimpleSearchReferenceCriteriaResponse struct {
	Schemas []EnumsimpleSearchReferenceCriteriaSchemaUrn `json:"schemas"`
	// Specifies a request criteria object that must match the associated request for references included in this Simple Search Reference Criteria.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Specifies the OID of a control that must be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain all of those controls.
	AllIncludedReferenceControl []string `json:"allIncludedReferenceControl,omitempty"`
	// Specifies the OID of a control that may be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must contain at least one of those controls.
	AnyIncludedReferenceControl []string `json:"anyIncludedReferenceControl,omitempty"`
	// Specifies the OID of a control that should not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain at least one of those controls (that is, it may contain zero or more of those controls, but not all of them).
	NotAllIncludedReferenceControl []string `json:"notAllIncludedReferenceControl,omitempty"`
	// Specifies the OID of a control that must not be present in search result references included in this Simple Search Reference Criteria. If any control OIDs are provided, then the reference must not contain any of those controls.
	NoneIncludedReferenceControl []string `json:"noneIncludedReferenceControl,omitempty"`
	// A description for this Search Reference Criteria
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Search Reference Criteria
	Id string `json:"id"`
}

// NewSimpleSearchReferenceCriteriaResponse instantiates a new SimpleSearchReferenceCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleSearchReferenceCriteriaResponse(schemas []EnumsimpleSearchReferenceCriteriaSchemaUrn, id string) *SimpleSearchReferenceCriteriaResponse {
	this := SimpleSearchReferenceCriteriaResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewSimpleSearchReferenceCriteriaResponseWithDefaults instantiates a new SimpleSearchReferenceCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleSearchReferenceCriteriaResponseWithDefaults() *SimpleSearchReferenceCriteriaResponse {
	this := SimpleSearchReferenceCriteriaResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *SimpleSearchReferenceCriteriaResponse) GetSchemas() []EnumsimpleSearchReferenceCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumsimpleSearchReferenceCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetSchemasOk() ([]EnumsimpleSearchReferenceCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SimpleSearchReferenceCriteriaResponse) SetSchemas(v []EnumsimpleSearchReferenceCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *SimpleSearchReferenceCriteriaResponse) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *SimpleSearchReferenceCriteriaResponse) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *SimpleSearchReferenceCriteriaResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetAllIncludedReferenceControl returns the AllIncludedReferenceControl field value if set, zero value otherwise.
func (o *SimpleSearchReferenceCriteriaResponse) GetAllIncludedReferenceControl() []string {
	if o == nil || IsNil(o.AllIncludedReferenceControl) {
		var ret []string
		return ret
	}
	return o.AllIncludedReferenceControl
}

// GetAllIncludedReferenceControlOk returns a tuple with the AllIncludedReferenceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetAllIncludedReferenceControlOk() ([]string, bool) {
	if o == nil || IsNil(o.AllIncludedReferenceControl) {
		return nil, false
	}
	return o.AllIncludedReferenceControl, true
}

// HasAllIncludedReferenceControl returns a boolean if a field has been set.
func (o *SimpleSearchReferenceCriteriaResponse) HasAllIncludedReferenceControl() bool {
	if o != nil && !IsNil(o.AllIncludedReferenceControl) {
		return true
	}

	return false
}

// SetAllIncludedReferenceControl gets a reference to the given []string and assigns it to the AllIncludedReferenceControl field.
func (o *SimpleSearchReferenceCriteriaResponse) SetAllIncludedReferenceControl(v []string) {
	o.AllIncludedReferenceControl = v
}

// GetAnyIncludedReferenceControl returns the AnyIncludedReferenceControl field value if set, zero value otherwise.
func (o *SimpleSearchReferenceCriteriaResponse) GetAnyIncludedReferenceControl() []string {
	if o == nil || IsNil(o.AnyIncludedReferenceControl) {
		var ret []string
		return ret
	}
	return o.AnyIncludedReferenceControl
}

// GetAnyIncludedReferenceControlOk returns a tuple with the AnyIncludedReferenceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetAnyIncludedReferenceControlOk() ([]string, bool) {
	if o == nil || IsNil(o.AnyIncludedReferenceControl) {
		return nil, false
	}
	return o.AnyIncludedReferenceControl, true
}

// HasAnyIncludedReferenceControl returns a boolean if a field has been set.
func (o *SimpleSearchReferenceCriteriaResponse) HasAnyIncludedReferenceControl() bool {
	if o != nil && !IsNil(o.AnyIncludedReferenceControl) {
		return true
	}

	return false
}

// SetAnyIncludedReferenceControl gets a reference to the given []string and assigns it to the AnyIncludedReferenceControl field.
func (o *SimpleSearchReferenceCriteriaResponse) SetAnyIncludedReferenceControl(v []string) {
	o.AnyIncludedReferenceControl = v
}

// GetNotAllIncludedReferenceControl returns the NotAllIncludedReferenceControl field value if set, zero value otherwise.
func (o *SimpleSearchReferenceCriteriaResponse) GetNotAllIncludedReferenceControl() []string {
	if o == nil || IsNil(o.NotAllIncludedReferenceControl) {
		var ret []string
		return ret
	}
	return o.NotAllIncludedReferenceControl
}

// GetNotAllIncludedReferenceControlOk returns a tuple with the NotAllIncludedReferenceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetNotAllIncludedReferenceControlOk() ([]string, bool) {
	if o == nil || IsNil(o.NotAllIncludedReferenceControl) {
		return nil, false
	}
	return o.NotAllIncludedReferenceControl, true
}

// HasNotAllIncludedReferenceControl returns a boolean if a field has been set.
func (o *SimpleSearchReferenceCriteriaResponse) HasNotAllIncludedReferenceControl() bool {
	if o != nil && !IsNil(o.NotAllIncludedReferenceControl) {
		return true
	}

	return false
}

// SetNotAllIncludedReferenceControl gets a reference to the given []string and assigns it to the NotAllIncludedReferenceControl field.
func (o *SimpleSearchReferenceCriteriaResponse) SetNotAllIncludedReferenceControl(v []string) {
	o.NotAllIncludedReferenceControl = v
}

// GetNoneIncludedReferenceControl returns the NoneIncludedReferenceControl field value if set, zero value otherwise.
func (o *SimpleSearchReferenceCriteriaResponse) GetNoneIncludedReferenceControl() []string {
	if o == nil || IsNil(o.NoneIncludedReferenceControl) {
		var ret []string
		return ret
	}
	return o.NoneIncludedReferenceControl
}

// GetNoneIncludedReferenceControlOk returns a tuple with the NoneIncludedReferenceControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetNoneIncludedReferenceControlOk() ([]string, bool) {
	if o == nil || IsNil(o.NoneIncludedReferenceControl) {
		return nil, false
	}
	return o.NoneIncludedReferenceControl, true
}

// HasNoneIncludedReferenceControl returns a boolean if a field has been set.
func (o *SimpleSearchReferenceCriteriaResponse) HasNoneIncludedReferenceControl() bool {
	if o != nil && !IsNil(o.NoneIncludedReferenceControl) {
		return true
	}

	return false
}

// SetNoneIncludedReferenceControl gets a reference to the given []string and assigns it to the NoneIncludedReferenceControl field.
func (o *SimpleSearchReferenceCriteriaResponse) SetNoneIncludedReferenceControl(v []string) {
	o.NoneIncludedReferenceControl = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SimpleSearchReferenceCriteriaResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SimpleSearchReferenceCriteriaResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SimpleSearchReferenceCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SimpleSearchReferenceCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SimpleSearchReferenceCriteriaResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SimpleSearchReferenceCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SimpleSearchReferenceCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SimpleSearchReferenceCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SimpleSearchReferenceCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SimpleSearchReferenceCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimpleSearchReferenceCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimpleSearchReferenceCriteriaResponse) SetId(v string) {
	o.Id = v
}

func (o SimpleSearchReferenceCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimpleSearchReferenceCriteriaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.AllIncludedReferenceControl) {
		toSerialize["allIncludedReferenceControl"] = o.AllIncludedReferenceControl
	}
	if !IsNil(o.AnyIncludedReferenceControl) {
		toSerialize["anyIncludedReferenceControl"] = o.AnyIncludedReferenceControl
	}
	if !IsNil(o.NotAllIncludedReferenceControl) {
		toSerialize["notAllIncludedReferenceControl"] = o.NotAllIncludedReferenceControl
	}
	if !IsNil(o.NoneIncludedReferenceControl) {
		toSerialize["noneIncludedReferenceControl"] = o.NoneIncludedReferenceControl
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

type NullableSimpleSearchReferenceCriteriaResponse struct {
	value *SimpleSearchReferenceCriteriaResponse
	isSet bool
}

func (v NullableSimpleSearchReferenceCriteriaResponse) Get() *SimpleSearchReferenceCriteriaResponse {
	return v.value
}

func (v *NullableSimpleSearchReferenceCriteriaResponse) Set(val *SimpleSearchReferenceCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleSearchReferenceCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleSearchReferenceCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleSearchReferenceCriteriaResponse(val *SimpleSearchReferenceCriteriaResponse) *NullableSimpleSearchReferenceCriteriaResponse {
	return &NullableSimpleSearchReferenceCriteriaResponse{value: val, isSet: true}
}

func (v NullableSimpleSearchReferenceCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleSearchReferenceCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
