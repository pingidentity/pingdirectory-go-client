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

// checks if the DelegatedAdminResourceRightsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelegatedAdminResourceRightsResponse{}

// DelegatedAdminResourceRightsResponse struct for DelegatedAdminResourceRightsResponse
type DelegatedAdminResourceRightsResponse struct {
	// Name of the Delegated Admin Rights
	Id      string                                      `json:"id"`
	Schemas []EnumdelegatedAdminResourceRightsSchemaUrn `json:"schemas,omitempty"`
	// A description for this Delegated Admin Resource Rights
	Description *string `json:"description,omitempty"`
	// Indicates whether these Delegated Admin Resource Rights are enabled.
	Enabled bool `json:"enabled"`
	// Specifies the resource type applicable to these Delegated Admin Resource Rights.
	RestResourceType string `json:"restResourceType"`
	// Specifies administrator(s) permissions.
	AdminPermission []EnumdelegatedAdminResourceRightsAdminPermissionProp `json:"adminPermission,omitempty"`
	AdminScope      *EnumdelegatedAdminResourceRightsAdminScopeProp       `json:"adminScope,omitempty"`
	// Specifies subtrees within the search base whose entries can be managed by the administrator(s). The admin-scope must be set to resources-in-specific-subtrees.
	ResourceSubtree []string `json:"resourceSubtree,omitempty"`
	// Specifies groups whose members can be managed by the administrator(s). The admin-scope must be set to resources-in-specific-groups.
	ResourcesInGroup                              []string                                           `json:"resourcesInGroup,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewDelegatedAdminResourceRightsResponse instantiates a new DelegatedAdminResourceRightsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelegatedAdminResourceRightsResponse(id string, enabled bool, restResourceType string) *DelegatedAdminResourceRightsResponse {
	this := DelegatedAdminResourceRightsResponse{}
	this.Id = id
	this.Enabled = enabled
	this.RestResourceType = restResourceType
	return &this
}

// NewDelegatedAdminResourceRightsResponseWithDefaults instantiates a new DelegatedAdminResourceRightsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelegatedAdminResourceRightsResponseWithDefaults() *DelegatedAdminResourceRightsResponse {
	this := DelegatedAdminResourceRightsResponse{}
	return &this
}

// GetId returns the Id field value
func (o *DelegatedAdminResourceRightsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DelegatedAdminResourceRightsResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *DelegatedAdminResourceRightsResponse) GetSchemas() []EnumdelegatedAdminResourceRightsSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumdelegatedAdminResourceRightsSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetSchemasOk() ([]EnumdelegatedAdminResourceRightsSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *DelegatedAdminResourceRightsResponse) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumdelegatedAdminResourceRightsSchemaUrn and assigns it to the Schemas field.
func (o *DelegatedAdminResourceRightsResponse) SetSchemas(v []EnumdelegatedAdminResourceRightsSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DelegatedAdminResourceRightsResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DelegatedAdminResourceRightsResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DelegatedAdminResourceRightsResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *DelegatedAdminResourceRightsResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DelegatedAdminResourceRightsResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetRestResourceType returns the RestResourceType field value
func (o *DelegatedAdminResourceRightsResponse) GetRestResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RestResourceType
}

// GetRestResourceTypeOk returns a tuple with the RestResourceType field value
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetRestResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestResourceType, true
}

// SetRestResourceType sets field value
func (o *DelegatedAdminResourceRightsResponse) SetRestResourceType(v string) {
	o.RestResourceType = v
}

// GetAdminPermission returns the AdminPermission field value if set, zero value otherwise.
func (o *DelegatedAdminResourceRightsResponse) GetAdminPermission() []EnumdelegatedAdminResourceRightsAdminPermissionProp {
	if o == nil || IsNil(o.AdminPermission) {
		var ret []EnumdelegatedAdminResourceRightsAdminPermissionProp
		return ret
	}
	return o.AdminPermission
}

// GetAdminPermissionOk returns a tuple with the AdminPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetAdminPermissionOk() ([]EnumdelegatedAdminResourceRightsAdminPermissionProp, bool) {
	if o == nil || IsNil(o.AdminPermission) {
		return nil, false
	}
	return o.AdminPermission, true
}

// HasAdminPermission returns a boolean if a field has been set.
func (o *DelegatedAdminResourceRightsResponse) HasAdminPermission() bool {
	if o != nil && !IsNil(o.AdminPermission) {
		return true
	}

	return false
}

// SetAdminPermission gets a reference to the given []EnumdelegatedAdminResourceRightsAdminPermissionProp and assigns it to the AdminPermission field.
func (o *DelegatedAdminResourceRightsResponse) SetAdminPermission(v []EnumdelegatedAdminResourceRightsAdminPermissionProp) {
	o.AdminPermission = v
}

// GetAdminScope returns the AdminScope field value if set, zero value otherwise.
func (o *DelegatedAdminResourceRightsResponse) GetAdminScope() EnumdelegatedAdminResourceRightsAdminScopeProp {
	if o == nil || IsNil(o.AdminScope) {
		var ret EnumdelegatedAdminResourceRightsAdminScopeProp
		return ret
	}
	return *o.AdminScope
}

// GetAdminScopeOk returns a tuple with the AdminScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetAdminScopeOk() (*EnumdelegatedAdminResourceRightsAdminScopeProp, bool) {
	if o == nil || IsNil(o.AdminScope) {
		return nil, false
	}
	return o.AdminScope, true
}

// HasAdminScope returns a boolean if a field has been set.
func (o *DelegatedAdminResourceRightsResponse) HasAdminScope() bool {
	if o != nil && !IsNil(o.AdminScope) {
		return true
	}

	return false
}

// SetAdminScope gets a reference to the given EnumdelegatedAdminResourceRightsAdminScopeProp and assigns it to the AdminScope field.
func (o *DelegatedAdminResourceRightsResponse) SetAdminScope(v EnumdelegatedAdminResourceRightsAdminScopeProp) {
	o.AdminScope = &v
}

// GetResourceSubtree returns the ResourceSubtree field value if set, zero value otherwise.
func (o *DelegatedAdminResourceRightsResponse) GetResourceSubtree() []string {
	if o == nil || IsNil(o.ResourceSubtree) {
		var ret []string
		return ret
	}
	return o.ResourceSubtree
}

// GetResourceSubtreeOk returns a tuple with the ResourceSubtree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetResourceSubtreeOk() ([]string, bool) {
	if o == nil || IsNil(o.ResourceSubtree) {
		return nil, false
	}
	return o.ResourceSubtree, true
}

// HasResourceSubtree returns a boolean if a field has been set.
func (o *DelegatedAdminResourceRightsResponse) HasResourceSubtree() bool {
	if o != nil && !IsNil(o.ResourceSubtree) {
		return true
	}

	return false
}

// SetResourceSubtree gets a reference to the given []string and assigns it to the ResourceSubtree field.
func (o *DelegatedAdminResourceRightsResponse) SetResourceSubtree(v []string) {
	o.ResourceSubtree = v
}

// GetResourcesInGroup returns the ResourcesInGroup field value if set, zero value otherwise.
func (o *DelegatedAdminResourceRightsResponse) GetResourcesInGroup() []string {
	if o == nil || IsNil(o.ResourcesInGroup) {
		var ret []string
		return ret
	}
	return o.ResourcesInGroup
}

// GetResourcesInGroupOk returns a tuple with the ResourcesInGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetResourcesInGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.ResourcesInGroup) {
		return nil, false
	}
	return o.ResourcesInGroup, true
}

// HasResourcesInGroup returns a boolean if a field has been set.
func (o *DelegatedAdminResourceRightsResponse) HasResourcesInGroup() bool {
	if o != nil && !IsNil(o.ResourcesInGroup) {
		return true
	}

	return false
}

// SetResourcesInGroup gets a reference to the given []string and assigns it to the ResourcesInGroup field.
func (o *DelegatedAdminResourceRightsResponse) SetResourcesInGroup(v []string) {
	o.ResourcesInGroup = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DelegatedAdminResourceRightsResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DelegatedAdminResourceRightsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DelegatedAdminResourceRightsResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DelegatedAdminResourceRightsResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelegatedAdminResourceRightsResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DelegatedAdminResourceRightsResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DelegatedAdminResourceRightsResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o DelegatedAdminResourceRightsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelegatedAdminResourceRightsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["restResourceType"] = o.RestResourceType
	if !IsNil(o.AdminPermission) {
		toSerialize["adminPermission"] = o.AdminPermission
	}
	if !IsNil(o.AdminScope) {
		toSerialize["adminScope"] = o.AdminScope
	}
	if !IsNil(o.ResourceSubtree) {
		toSerialize["resourceSubtree"] = o.ResourceSubtree
	}
	if !IsNil(o.ResourcesInGroup) {
		toSerialize["resourcesInGroup"] = o.ResourcesInGroup
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableDelegatedAdminResourceRightsResponse struct {
	value *DelegatedAdminResourceRightsResponse
	isSet bool
}

func (v NullableDelegatedAdminResourceRightsResponse) Get() *DelegatedAdminResourceRightsResponse {
	return v.value
}

func (v *NullableDelegatedAdminResourceRightsResponse) Set(val *DelegatedAdminResourceRightsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDelegatedAdminResourceRightsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDelegatedAdminResourceRightsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelegatedAdminResourceRightsResponse(val *DelegatedAdminResourceRightsResponse) *NullableDelegatedAdminResourceRightsResponse {
	return &NullableDelegatedAdminResourceRightsResponse{value: val, isSet: true}
}

func (v NullableDelegatedAdminResourceRightsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelegatedAdminResourceRightsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
