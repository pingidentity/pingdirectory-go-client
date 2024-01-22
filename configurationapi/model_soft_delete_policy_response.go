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

// checks if the SoftDeletePolicyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftDeletePolicyResponse{}

// SoftDeletePolicyResponse struct for SoftDeletePolicyResponse
type SoftDeletePolicyResponse struct {
	Schemas []EnumsoftDeletePolicySchemaUrn `json:"schemas,omitempty"`
	// A description for this Soft Delete Policy
	Description *string `json:"description,omitempty"`
	// Connection criteria used to automatically identify a delete operation for processing as a soft delete request.
	AutoSoftDeleteConnectionCriteria *string `json:"autoSoftDeleteConnectionCriteria,omitempty"`
	// Request criteria used to automatically identify a delete operation for processing as a soft delete request.
	AutoSoftDeleteRequestCriteria *string `json:"autoSoftDeleteRequestCriteria,omitempty"`
	// Specifies the maximum length of time that soft delete entries are retained before they are eligible to purged automatically.
	SoftDeleteRetentionTime *string `json:"softDeleteRetentionTime,omitempty"`
	// Specifies the number of soft deleted entries to retain before the oldest entries are purged.
	SoftDeleteRetainNumberOfEntries               *int64                                             `json:"softDeleteRetainNumberOfEntries,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Soft Delete Policy
	Id string `json:"id"`
}

// NewSoftDeletePolicyResponse instantiates a new SoftDeletePolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftDeletePolicyResponse(id string) *SoftDeletePolicyResponse {
	this := SoftDeletePolicyResponse{}
	this.Id = id
	return &this
}

// NewSoftDeletePolicyResponseWithDefaults instantiates a new SoftDeletePolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftDeletePolicyResponseWithDefaults() *SoftDeletePolicyResponse {
	this := SoftDeletePolicyResponse{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *SoftDeletePolicyResponse) GetSchemas() []EnumsoftDeletePolicySchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumsoftDeletePolicySchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetSchemasOk() ([]EnumsoftDeletePolicySchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *SoftDeletePolicyResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumsoftDeletePolicySchemaUrn and assigns it to the Schemas field.
func (o *SoftDeletePolicyResponse) SetSchemas(v []EnumsoftDeletePolicySchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SoftDeletePolicyResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SoftDeletePolicyResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SoftDeletePolicyResponse) SetDescription(v string) {
	o.Description = &v
}

// GetAutoSoftDeleteConnectionCriteria returns the AutoSoftDeleteConnectionCriteria field value if set, zero value otherwise.
func (o *SoftDeletePolicyResponse) GetAutoSoftDeleteConnectionCriteria() string {
	if o == nil || IsNil(o.AutoSoftDeleteConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.AutoSoftDeleteConnectionCriteria
}

// GetAutoSoftDeleteConnectionCriteriaOk returns a tuple with the AutoSoftDeleteConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetAutoSoftDeleteConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AutoSoftDeleteConnectionCriteria) {
		return nil, false
	}
	return o.AutoSoftDeleteConnectionCriteria, true
}

// HasAutoSoftDeleteConnectionCriteria returns a boolean if a field has been set.
func (o *SoftDeletePolicyResponse) HasAutoSoftDeleteConnectionCriteria() bool {
	if o != nil && !IsNil(o.AutoSoftDeleteConnectionCriteria) {
		return true
	}

	return false
}

// SetAutoSoftDeleteConnectionCriteria gets a reference to the given string and assigns it to the AutoSoftDeleteConnectionCriteria field.
func (o *SoftDeletePolicyResponse) SetAutoSoftDeleteConnectionCriteria(v string) {
	o.AutoSoftDeleteConnectionCriteria = &v
}

// GetAutoSoftDeleteRequestCriteria returns the AutoSoftDeleteRequestCriteria field value if set, zero value otherwise.
func (o *SoftDeletePolicyResponse) GetAutoSoftDeleteRequestCriteria() string {
	if o == nil || IsNil(o.AutoSoftDeleteRequestCriteria) {
		var ret string
		return ret
	}
	return *o.AutoSoftDeleteRequestCriteria
}

// GetAutoSoftDeleteRequestCriteriaOk returns a tuple with the AutoSoftDeleteRequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetAutoSoftDeleteRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.AutoSoftDeleteRequestCriteria) {
		return nil, false
	}
	return o.AutoSoftDeleteRequestCriteria, true
}

// HasAutoSoftDeleteRequestCriteria returns a boolean if a field has been set.
func (o *SoftDeletePolicyResponse) HasAutoSoftDeleteRequestCriteria() bool {
	if o != nil && !IsNil(o.AutoSoftDeleteRequestCriteria) {
		return true
	}

	return false
}

// SetAutoSoftDeleteRequestCriteria gets a reference to the given string and assigns it to the AutoSoftDeleteRequestCriteria field.
func (o *SoftDeletePolicyResponse) SetAutoSoftDeleteRequestCriteria(v string) {
	o.AutoSoftDeleteRequestCriteria = &v
}

// GetSoftDeleteRetentionTime returns the SoftDeleteRetentionTime field value if set, zero value otherwise.
func (o *SoftDeletePolicyResponse) GetSoftDeleteRetentionTime() string {
	if o == nil || IsNil(o.SoftDeleteRetentionTime) {
		var ret string
		return ret
	}
	return *o.SoftDeleteRetentionTime
}

// GetSoftDeleteRetentionTimeOk returns a tuple with the SoftDeleteRetentionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetSoftDeleteRetentionTimeOk() (*string, bool) {
	if o == nil || IsNil(o.SoftDeleteRetentionTime) {
		return nil, false
	}
	return o.SoftDeleteRetentionTime, true
}

// HasSoftDeleteRetentionTime returns a boolean if a field has been set.
func (o *SoftDeletePolicyResponse) HasSoftDeleteRetentionTime() bool {
	if o != nil && !IsNil(o.SoftDeleteRetentionTime) {
		return true
	}

	return false
}

// SetSoftDeleteRetentionTime gets a reference to the given string and assigns it to the SoftDeleteRetentionTime field.
func (o *SoftDeletePolicyResponse) SetSoftDeleteRetentionTime(v string) {
	o.SoftDeleteRetentionTime = &v
}

// GetSoftDeleteRetainNumberOfEntries returns the SoftDeleteRetainNumberOfEntries field value if set, zero value otherwise.
func (o *SoftDeletePolicyResponse) GetSoftDeleteRetainNumberOfEntries() int64 {
	if o == nil || IsNil(o.SoftDeleteRetainNumberOfEntries) {
		var ret int64
		return ret
	}
	return *o.SoftDeleteRetainNumberOfEntries
}

// GetSoftDeleteRetainNumberOfEntriesOk returns a tuple with the SoftDeleteRetainNumberOfEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetSoftDeleteRetainNumberOfEntriesOk() (*int64, bool) {
	if o == nil || IsNil(o.SoftDeleteRetainNumberOfEntries) {
		return nil, false
	}
	return o.SoftDeleteRetainNumberOfEntries, true
}

// HasSoftDeleteRetainNumberOfEntries returns a boolean if a field has been set.
func (o *SoftDeletePolicyResponse) HasSoftDeleteRetainNumberOfEntries() bool {
	if o != nil && !IsNil(o.SoftDeleteRetainNumberOfEntries) {
		return true
	}

	return false
}

// SetSoftDeleteRetainNumberOfEntries gets a reference to the given int64 and assigns it to the SoftDeleteRetainNumberOfEntries field.
func (o *SoftDeletePolicyResponse) SetSoftDeleteRetainNumberOfEntries(v int64) {
	o.SoftDeleteRetainNumberOfEntries = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SoftDeletePolicyResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SoftDeletePolicyResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SoftDeletePolicyResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SoftDeletePolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SoftDeletePolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SoftDeletePolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SoftDeletePolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SoftDeletePolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SoftDeletePolicyResponse) SetId(v string) {
	o.Id = v
}

func (o SoftDeletePolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftDeletePolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AutoSoftDeleteConnectionCriteria) {
		toSerialize["autoSoftDeleteConnectionCriteria"] = o.AutoSoftDeleteConnectionCriteria
	}
	if !IsNil(o.AutoSoftDeleteRequestCriteria) {
		toSerialize["autoSoftDeleteRequestCriteria"] = o.AutoSoftDeleteRequestCriteria
	}
	if !IsNil(o.SoftDeleteRetentionTime) {
		toSerialize["softDeleteRetentionTime"] = o.SoftDeleteRetentionTime
	}
	if !IsNil(o.SoftDeleteRetainNumberOfEntries) {
		toSerialize["softDeleteRetainNumberOfEntries"] = o.SoftDeleteRetainNumberOfEntries
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

type NullableSoftDeletePolicyResponse struct {
	value *SoftDeletePolicyResponse
	isSet bool
}

func (v NullableSoftDeletePolicyResponse) Get() *SoftDeletePolicyResponse {
	return v.value
}

func (v *NullableSoftDeletePolicyResponse) Set(val *SoftDeletePolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftDeletePolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftDeletePolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftDeletePolicyResponse(val *SoftDeletePolicyResponse) *NullableSoftDeletePolicyResponse {
	return &NullableSoftDeletePolicyResponse{value: val, isSet: true}
}

func (v NullableSoftDeletePolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftDeletePolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
