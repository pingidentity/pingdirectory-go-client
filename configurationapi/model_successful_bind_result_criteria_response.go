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

// checks if the SuccessfulBindResultCriteriaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SuccessfulBindResultCriteriaResponse{}

// SuccessfulBindResultCriteriaResponse struct for SuccessfulBindResultCriteriaResponse
type SuccessfulBindResultCriteriaResponse struct {
	Schemas []EnumsuccessfulBindResultCriteriaSchemaUrn `json:"schemas"`
	// Specifies a request criteria object that must match the associated request for operations included in this Successful Bind Result Criteria.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Indicates whether this criteria will be permitted to match bind operations that resulted in anonymous authentication.
	IncludeAnonymousBinds *bool `json:"includeAnonymousBinds,omitempty"`
	// A set of base DNs for authenticated users that will be permitted to match this criteria.
	IncludedUserBaseDN []string `json:"includedUserBaseDN,omitempty"`
	// A set of base DNs for authenticated users that will not be permitted to match this criteria.
	ExcludedUserBaseDN []string `json:"excludedUserBaseDN,omitempty"`
	// A set of filters that may be used to identify entries for authenticated users that will be permitted to match this criteria.
	IncludedUserFilter []string `json:"includedUserFilter,omitempty"`
	// A set of filters that may be used to identify entries for authenticated users that will not be permitted to match this criteria.
	ExcludedUserFilter []string `json:"excludedUserFilter,omitempty"`
	// The DNs of the groups whose members will be permitted to match this criteria.
	IncludedUserGroupDN []string `json:"includedUserGroupDN,omitempty"`
	// The DNs of the groups whose members will not be permitted to match this criteria.
	ExcludedUserGroupDN []string `json:"excludedUserGroupDN,omitempty"`
	// A description for this Result Criteria
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Result Criteria
	Id string `json:"id"`
}

// NewSuccessfulBindResultCriteriaResponse instantiates a new SuccessfulBindResultCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessfulBindResultCriteriaResponse(schemas []EnumsuccessfulBindResultCriteriaSchemaUrn, id string) *SuccessfulBindResultCriteriaResponse {
	this := SuccessfulBindResultCriteriaResponse{}
	this.Schemas = schemas
	this.Id = id
	return &this
}

// NewSuccessfulBindResultCriteriaResponseWithDefaults instantiates a new SuccessfulBindResultCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessfulBindResultCriteriaResponseWithDefaults() *SuccessfulBindResultCriteriaResponse {
	this := SuccessfulBindResultCriteriaResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *SuccessfulBindResultCriteriaResponse) GetSchemas() []EnumsuccessfulBindResultCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumsuccessfulBindResultCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetSchemasOk() ([]EnumsuccessfulBindResultCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SuccessfulBindResultCriteriaResponse) SetSchemas(v []EnumsuccessfulBindResultCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *SuccessfulBindResultCriteriaResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetIncludeAnonymousBinds returns the IncludeAnonymousBinds field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetIncludeAnonymousBinds() bool {
	if o == nil || IsNil(o.IncludeAnonymousBinds) {
		var ret bool
		return ret
	}
	return *o.IncludeAnonymousBinds
}

// GetIncludeAnonymousBindsOk returns a tuple with the IncludeAnonymousBinds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetIncludeAnonymousBindsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeAnonymousBinds) {
		return nil, false
	}
	return o.IncludeAnonymousBinds, true
}

// HasIncludeAnonymousBinds returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasIncludeAnonymousBinds() bool {
	if o != nil && !IsNil(o.IncludeAnonymousBinds) {
		return true
	}

	return false
}

// SetIncludeAnonymousBinds gets a reference to the given bool and assigns it to the IncludeAnonymousBinds field.
func (o *SuccessfulBindResultCriteriaResponse) SetIncludeAnonymousBinds(v bool) {
	o.IncludeAnonymousBinds = &v
}

// GetIncludedUserBaseDN returns the IncludedUserBaseDN field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetIncludedUserBaseDN() []string {
	if o == nil || IsNil(o.IncludedUserBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludedUserBaseDN
}

// GetIncludedUserBaseDNOk returns a tuple with the IncludedUserBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetIncludedUserBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedUserBaseDN) {
		return nil, false
	}
	return o.IncludedUserBaseDN, true
}

// HasIncludedUserBaseDN returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasIncludedUserBaseDN() bool {
	if o != nil && !IsNil(o.IncludedUserBaseDN) {
		return true
	}

	return false
}

// SetIncludedUserBaseDN gets a reference to the given []string and assigns it to the IncludedUserBaseDN field.
func (o *SuccessfulBindResultCriteriaResponse) SetIncludedUserBaseDN(v []string) {
	o.IncludedUserBaseDN = v
}

// GetExcludedUserBaseDN returns the ExcludedUserBaseDN field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetExcludedUserBaseDN() []string {
	if o == nil || IsNil(o.ExcludedUserBaseDN) {
		var ret []string
		return ret
	}
	return o.ExcludedUserBaseDN
}

// GetExcludedUserBaseDNOk returns a tuple with the ExcludedUserBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetExcludedUserBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedUserBaseDN) {
		return nil, false
	}
	return o.ExcludedUserBaseDN, true
}

// HasExcludedUserBaseDN returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasExcludedUserBaseDN() bool {
	if o != nil && !IsNil(o.ExcludedUserBaseDN) {
		return true
	}

	return false
}

// SetExcludedUserBaseDN gets a reference to the given []string and assigns it to the ExcludedUserBaseDN field.
func (o *SuccessfulBindResultCriteriaResponse) SetExcludedUserBaseDN(v []string) {
	o.ExcludedUserBaseDN = v
}

// GetIncludedUserFilter returns the IncludedUserFilter field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetIncludedUserFilter() []string {
	if o == nil || IsNil(o.IncludedUserFilter) {
		var ret []string
		return ret
	}
	return o.IncludedUserFilter
}

// GetIncludedUserFilterOk returns a tuple with the IncludedUserFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetIncludedUserFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedUserFilter) {
		return nil, false
	}
	return o.IncludedUserFilter, true
}

// HasIncludedUserFilter returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasIncludedUserFilter() bool {
	if o != nil && !IsNil(o.IncludedUserFilter) {
		return true
	}

	return false
}

// SetIncludedUserFilter gets a reference to the given []string and assigns it to the IncludedUserFilter field.
func (o *SuccessfulBindResultCriteriaResponse) SetIncludedUserFilter(v []string) {
	o.IncludedUserFilter = v
}

// GetExcludedUserFilter returns the ExcludedUserFilter field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetExcludedUserFilter() []string {
	if o == nil || IsNil(o.ExcludedUserFilter) {
		var ret []string
		return ret
	}
	return o.ExcludedUserFilter
}

// GetExcludedUserFilterOk returns a tuple with the ExcludedUserFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetExcludedUserFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedUserFilter) {
		return nil, false
	}
	return o.ExcludedUserFilter, true
}

// HasExcludedUserFilter returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasExcludedUserFilter() bool {
	if o != nil && !IsNil(o.ExcludedUserFilter) {
		return true
	}

	return false
}

// SetExcludedUserFilter gets a reference to the given []string and assigns it to the ExcludedUserFilter field.
func (o *SuccessfulBindResultCriteriaResponse) SetExcludedUserFilter(v []string) {
	o.ExcludedUserFilter = v
}

// GetIncludedUserGroupDN returns the IncludedUserGroupDN field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetIncludedUserGroupDN() []string {
	if o == nil || IsNil(o.IncludedUserGroupDN) {
		var ret []string
		return ret
	}
	return o.IncludedUserGroupDN
}

// GetIncludedUserGroupDNOk returns a tuple with the IncludedUserGroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetIncludedUserGroupDNOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedUserGroupDN) {
		return nil, false
	}
	return o.IncludedUserGroupDN, true
}

// HasIncludedUserGroupDN returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasIncludedUserGroupDN() bool {
	if o != nil && !IsNil(o.IncludedUserGroupDN) {
		return true
	}

	return false
}

// SetIncludedUserGroupDN gets a reference to the given []string and assigns it to the IncludedUserGroupDN field.
func (o *SuccessfulBindResultCriteriaResponse) SetIncludedUserGroupDN(v []string) {
	o.IncludedUserGroupDN = v
}

// GetExcludedUserGroupDN returns the ExcludedUserGroupDN field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetExcludedUserGroupDN() []string {
	if o == nil || IsNil(o.ExcludedUserGroupDN) {
		var ret []string
		return ret
	}
	return o.ExcludedUserGroupDN
}

// GetExcludedUserGroupDNOk returns a tuple with the ExcludedUserGroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetExcludedUserGroupDNOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludedUserGroupDN) {
		return nil, false
	}
	return o.ExcludedUserGroupDN, true
}

// HasExcludedUserGroupDN returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasExcludedUserGroupDN() bool {
	if o != nil && !IsNil(o.ExcludedUserGroupDN) {
		return true
	}

	return false
}

// SetExcludedUserGroupDN gets a reference to the given []string and assigns it to the ExcludedUserGroupDN field.
func (o *SuccessfulBindResultCriteriaResponse) SetExcludedUserGroupDN(v []string) {
	o.ExcludedUserGroupDN = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SuccessfulBindResultCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SuccessfulBindResultCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SuccessfulBindResultCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SuccessfulBindResultCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SuccessfulBindResultCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SuccessfulBindResultCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SuccessfulBindResultCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SuccessfulBindResultCriteriaResponse) SetId(v string) {
	o.Id = v
}

func (o SuccessfulBindResultCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuccessfulBindResultCriteriaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.IncludeAnonymousBinds) {
		toSerialize["includeAnonymousBinds"] = o.IncludeAnonymousBinds
	}
	if !IsNil(o.IncludedUserBaseDN) {
		toSerialize["includedUserBaseDN"] = o.IncludedUserBaseDN
	}
	if !IsNil(o.ExcludedUserBaseDN) {
		toSerialize["excludedUserBaseDN"] = o.ExcludedUserBaseDN
	}
	if !IsNil(o.IncludedUserFilter) {
		toSerialize["includedUserFilter"] = o.IncludedUserFilter
	}
	if !IsNil(o.ExcludedUserFilter) {
		toSerialize["excludedUserFilter"] = o.ExcludedUserFilter
	}
	if !IsNil(o.IncludedUserGroupDN) {
		toSerialize["includedUserGroupDN"] = o.IncludedUserGroupDN
	}
	if !IsNil(o.ExcludedUserGroupDN) {
		toSerialize["excludedUserGroupDN"] = o.ExcludedUserGroupDN
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

type NullableSuccessfulBindResultCriteriaResponse struct {
	value *SuccessfulBindResultCriteriaResponse
	isSet bool
}

func (v NullableSuccessfulBindResultCriteriaResponse) Get() *SuccessfulBindResultCriteriaResponse {
	return v.value
}

func (v *NullableSuccessfulBindResultCriteriaResponse) Set(val *SuccessfulBindResultCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessfulBindResultCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessfulBindResultCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessfulBindResultCriteriaResponse(val *SuccessfulBindResultCriteriaResponse) *NullableSuccessfulBindResultCriteriaResponse {
	return &NullableSuccessfulBindResultCriteriaResponse{value: val, isSet: true}
}

func (v NullableSuccessfulBindResultCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessfulBindResultCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
