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

// checks if the SoftReferenceEntryCacheResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoftReferenceEntryCacheResponse{}

// SoftReferenceEntryCacheResponse struct for SoftReferenceEntryCacheResponse
type SoftReferenceEntryCacheResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumsoftReferenceEntryCacheSchemaUrn             `json:"schemas"`
	// Name of the Entry Cache
	Id string `json:"id"`
	// The set of filters that define the entries that should be included in the cache.
	IncludeFilter []string `json:"includeFilter,omitempty"`
	// The set of filters that define the entries that should be excluded from the cache.
	ExcludeFilter []string `json:"excludeFilter,omitempty"`
	// A description for this Entry Cache
	Description *string `json:"description,omitempty"`
	// Indicates whether the Entry Cache is enabled.
	Enabled bool `json:"enabled"`
	// Specifies the cache level in the cache order if more than one instance of the cache is configured.
	CacheLevel int64 `json:"cacheLevel"`
	// Indicates whether the entry cache should be updated with entries that have been returned to the client during the course of processing an unindexed search.
	CacheUnindexedSearchResults *bool `json:"cacheUnindexedSearchResults,omitempty"`
}

// NewSoftReferenceEntryCacheResponse instantiates a new SoftReferenceEntryCacheResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftReferenceEntryCacheResponse(schemas []EnumsoftReferenceEntryCacheSchemaUrn, id string, enabled bool, cacheLevel int64) *SoftReferenceEntryCacheResponse {
	this := SoftReferenceEntryCacheResponse{}
	this.Schemas = schemas
	this.Id = id
	this.Enabled = enabled
	this.CacheLevel = cacheLevel
	return &this
}

// NewSoftReferenceEntryCacheResponseWithDefaults instantiates a new SoftReferenceEntryCacheResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftReferenceEntryCacheResponseWithDefaults() *SoftReferenceEntryCacheResponse {
	this := SoftReferenceEntryCacheResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SoftReferenceEntryCacheResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SoftReferenceEntryCacheResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SoftReferenceEntryCacheResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SoftReferenceEntryCacheResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SoftReferenceEntryCacheResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SoftReferenceEntryCacheResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *SoftReferenceEntryCacheResponse) GetSchemas() []EnumsoftReferenceEntryCacheSchemaUrn {
	if o == nil {
		var ret []EnumsoftReferenceEntryCacheSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetSchemasOk() ([]EnumsoftReferenceEntryCacheSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SoftReferenceEntryCacheResponse) SetSchemas(v []EnumsoftReferenceEntryCacheSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *SoftReferenceEntryCacheResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SoftReferenceEntryCacheResponse) SetId(v string) {
	o.Id = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *SoftReferenceEntryCacheResponse) GetIncludeFilter() []string {
	if o == nil || IsNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *SoftReferenceEntryCacheResponse) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *SoftReferenceEntryCacheResponse) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetExcludeFilter returns the ExcludeFilter field value if set, zero value otherwise.
func (o *SoftReferenceEntryCacheResponse) GetExcludeFilter() []string {
	if o == nil || IsNil(o.ExcludeFilter) {
		var ret []string
		return ret
	}
	return o.ExcludeFilter
}

// GetExcludeFilterOk returns a tuple with the ExcludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetExcludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeFilter) {
		return nil, false
	}
	return o.ExcludeFilter, true
}

// HasExcludeFilter returns a boolean if a field has been set.
func (o *SoftReferenceEntryCacheResponse) HasExcludeFilter() bool {
	if o != nil && !IsNil(o.ExcludeFilter) {
		return true
	}

	return false
}

// SetExcludeFilter gets a reference to the given []string and assigns it to the ExcludeFilter field.
func (o *SoftReferenceEntryCacheResponse) SetExcludeFilter(v []string) {
	o.ExcludeFilter = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SoftReferenceEntryCacheResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SoftReferenceEntryCacheResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SoftReferenceEntryCacheResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SoftReferenceEntryCacheResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SoftReferenceEntryCacheResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetCacheLevel returns the CacheLevel field value
func (o *SoftReferenceEntryCacheResponse) GetCacheLevel() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CacheLevel
}

// GetCacheLevelOk returns a tuple with the CacheLevel field value
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetCacheLevelOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheLevel, true
}

// SetCacheLevel sets field value
func (o *SoftReferenceEntryCacheResponse) SetCacheLevel(v int64) {
	o.CacheLevel = v
}

// GetCacheUnindexedSearchResults returns the CacheUnindexedSearchResults field value if set, zero value otherwise.
func (o *SoftReferenceEntryCacheResponse) GetCacheUnindexedSearchResults() bool {
	if o == nil || IsNil(o.CacheUnindexedSearchResults) {
		var ret bool
		return ret
	}
	return *o.CacheUnindexedSearchResults
}

// GetCacheUnindexedSearchResultsOk returns a tuple with the CacheUnindexedSearchResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftReferenceEntryCacheResponse) GetCacheUnindexedSearchResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.CacheUnindexedSearchResults) {
		return nil, false
	}
	return o.CacheUnindexedSearchResults, true
}

// HasCacheUnindexedSearchResults returns a boolean if a field has been set.
func (o *SoftReferenceEntryCacheResponse) HasCacheUnindexedSearchResults() bool {
	if o != nil && !IsNil(o.CacheUnindexedSearchResults) {
		return true
	}

	return false
}

// SetCacheUnindexedSearchResults gets a reference to the given bool and assigns it to the CacheUnindexedSearchResults field.
func (o *SoftReferenceEntryCacheResponse) SetCacheUnindexedSearchResults(v bool) {
	o.CacheUnindexedSearchResults = &v
}

func (o SoftReferenceEntryCacheResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoftReferenceEntryCacheResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !IsNil(o.ExcludeFilter) {
		toSerialize["excludeFilter"] = o.ExcludeFilter
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["cacheLevel"] = o.CacheLevel
	if !IsNil(o.CacheUnindexedSearchResults) {
		toSerialize["cacheUnindexedSearchResults"] = o.CacheUnindexedSearchResults
	}
	return toSerialize, nil
}

type NullableSoftReferenceEntryCacheResponse struct {
	value *SoftReferenceEntryCacheResponse
	isSet bool
}

func (v NullableSoftReferenceEntryCacheResponse) Get() *SoftReferenceEntryCacheResponse {
	return v.value
}

func (v *NullableSoftReferenceEntryCacheResponse) Set(val *SoftReferenceEntryCacheResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftReferenceEntryCacheResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftReferenceEntryCacheResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftReferenceEntryCacheResponse(val *SoftReferenceEntryCacheResponse) *NullableSoftReferenceEntryCacheResponse {
	return &NullableSoftReferenceEntryCacheResponse{value: val, isSet: true}
}

func (v NullableSoftReferenceEntryCacheResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftReferenceEntryCacheResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
