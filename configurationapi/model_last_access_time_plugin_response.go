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

// checks if the LastAccessTimePluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LastAccessTimePluginResponse{}

// LastAccessTimePluginResponse struct for LastAccessTimePluginResponse
type LastAccessTimePluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumlastAccessTimePluginSchemaUrn                `json:"schemas"`
	// Name of the Plugin
	Id string `json:"id"`
	// Specifies the maximum frequency with which last access time values should be written for an entry. This may help limit the rate of internal write operations processed in the server.
	MaxUpdateFrequency *string                       `json:"maxUpdateFrequency,omitempty"`
	OperationType      []EnumpluginOperationTypeProp `json:"operationType,omitempty"`
	// Indicates whether to update the last access time for an entry targeted by a bind operation if the bind is unsuccessful.
	InvokeForFailedBinds *bool `json:"invokeForFailedBinds,omitempty"`
	// Specifies the maximum number of entries that should be updated in a search operation. Only search result entries actually returned to the client may have their last access time updated, but because a single search operation may return a very large number of entries, the plugin will only update entries if no more than a specified number of entries are updated.
	MaxSearchResultEntriesToUpdate *int32 `json:"maxSearchResultEntriesToUpdate,omitempty"`
	// Specifies a set of request criteria that may be used to indicate whether to apply access time updates for the associated operation.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewLastAccessTimePluginResponse instantiates a new LastAccessTimePluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLastAccessTimePluginResponse(schemas []EnumlastAccessTimePluginSchemaUrn, id string, enabled bool) *LastAccessTimePluginResponse {
	this := LastAccessTimePluginResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	return &this
}

// NewLastAccessTimePluginResponseWithDefaults instantiates a new LastAccessTimePluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLastAccessTimePluginResponseWithDefaults() *LastAccessTimePluginResponse {
	this := LastAccessTimePluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *LastAccessTimePluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *LastAccessTimePluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *LastAccessTimePluginResponse) GetSchemas() []EnumlastAccessTimePluginSchemaUrn {
	if o == nil {
		var ret []EnumlastAccessTimePluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetSchemasOk() ([]EnumlastAccessTimePluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *LastAccessTimePluginResponse) SetSchemas(v []EnumlastAccessTimePluginSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *LastAccessTimePluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LastAccessTimePluginResponse) SetId(v string) {
	o.Id = v
}

// GetMaxUpdateFrequency returns the MaxUpdateFrequency field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetMaxUpdateFrequency() string {
	if o == nil || IsNil(o.MaxUpdateFrequency) {
		var ret string
		return ret
	}
	return *o.MaxUpdateFrequency
}

// GetMaxUpdateFrequencyOk returns a tuple with the MaxUpdateFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetMaxUpdateFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.MaxUpdateFrequency) {
		return nil, false
	}
	return o.MaxUpdateFrequency, true
}

// HasMaxUpdateFrequency returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasMaxUpdateFrequency() bool {
	if o != nil && !IsNil(o.MaxUpdateFrequency) {
		return true
	}

	return false
}

// SetMaxUpdateFrequency gets a reference to the given string and assigns it to the MaxUpdateFrequency field.
func (o *LastAccessTimePluginResponse) SetMaxUpdateFrequency(v string) {
	o.MaxUpdateFrequency = &v
}

// GetOperationType returns the OperationType field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetOperationType() []EnumpluginOperationTypeProp {
	if o == nil || IsNil(o.OperationType) {
		var ret []EnumpluginOperationTypeProp
		return ret
	}
	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetOperationTypeOk() ([]EnumpluginOperationTypeProp, bool) {
	if o == nil || IsNil(o.OperationType) {
		return nil, false
	}
	return o.OperationType, true
}

// HasOperationType returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasOperationType() bool {
	if o != nil && !IsNil(o.OperationType) {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given []EnumpluginOperationTypeProp and assigns it to the OperationType field.
func (o *LastAccessTimePluginResponse) SetOperationType(v []EnumpluginOperationTypeProp) {
	o.OperationType = v
}

// GetInvokeForFailedBinds returns the InvokeForFailedBinds field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetInvokeForFailedBinds() bool {
	if o == nil || IsNil(o.InvokeForFailedBinds) {
		var ret bool
		return ret
	}
	return *o.InvokeForFailedBinds
}

// GetInvokeForFailedBindsOk returns a tuple with the InvokeForFailedBinds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetInvokeForFailedBindsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForFailedBinds) {
		return nil, false
	}
	return o.InvokeForFailedBinds, true
}

// HasInvokeForFailedBinds returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasInvokeForFailedBinds() bool {
	if o != nil && !IsNil(o.InvokeForFailedBinds) {
		return true
	}

	return false
}

// SetInvokeForFailedBinds gets a reference to the given bool and assigns it to the InvokeForFailedBinds field.
func (o *LastAccessTimePluginResponse) SetInvokeForFailedBinds(v bool) {
	o.InvokeForFailedBinds = &v
}

// GetMaxSearchResultEntriesToUpdate returns the MaxSearchResultEntriesToUpdate field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetMaxSearchResultEntriesToUpdate() int32 {
	if o == nil || IsNil(o.MaxSearchResultEntriesToUpdate) {
		var ret int32
		return ret
	}
	return *o.MaxSearchResultEntriesToUpdate
}

// GetMaxSearchResultEntriesToUpdateOk returns a tuple with the MaxSearchResultEntriesToUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetMaxSearchResultEntriesToUpdateOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxSearchResultEntriesToUpdate) {
		return nil, false
	}
	return o.MaxSearchResultEntriesToUpdate, true
}

// HasMaxSearchResultEntriesToUpdate returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasMaxSearchResultEntriesToUpdate() bool {
	if o != nil && !IsNil(o.MaxSearchResultEntriesToUpdate) {
		return true
	}

	return false
}

// SetMaxSearchResultEntriesToUpdate gets a reference to the given int32 and assigns it to the MaxSearchResultEntriesToUpdate field.
func (o *LastAccessTimePluginResponse) SetMaxSearchResultEntriesToUpdate(v int32) {
	o.MaxSearchResultEntriesToUpdate = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *LastAccessTimePluginResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *LastAccessTimePluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LastAccessTimePluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LastAccessTimePluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LastAccessTimePluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *LastAccessTimePluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *LastAccessTimePluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *LastAccessTimePluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o LastAccessTimePluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LastAccessTimePluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.MaxUpdateFrequency) {
		toSerialize["maxUpdateFrequency"] = o.MaxUpdateFrequency
	}
	if !IsNil(o.OperationType) {
		toSerialize["operationType"] = o.OperationType
	}
	if !IsNil(o.InvokeForFailedBinds) {
		toSerialize["invokeForFailedBinds"] = o.InvokeForFailedBinds
	}
	if !IsNil(o.MaxSearchResultEntriesToUpdate) {
		toSerialize["maxSearchResultEntriesToUpdate"] = o.MaxSearchResultEntriesToUpdate
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableLastAccessTimePluginResponse struct {
	value *LastAccessTimePluginResponse
	isSet bool
}

func (v NullableLastAccessTimePluginResponse) Get() *LastAccessTimePluginResponse {
	return v.value
}

func (v *NullableLastAccessTimePluginResponse) Set(val *LastAccessTimePluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLastAccessTimePluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLastAccessTimePluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLastAccessTimePluginResponse(val *LastAccessTimePluginResponse) *NullableLastAccessTimePluginResponse {
	return &NullableLastAccessTimePluginResponse{value: val, isSet: true}
}

func (v NullableLastAccessTimePluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLastAccessTimePluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
