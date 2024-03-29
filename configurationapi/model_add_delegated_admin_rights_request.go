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

// checks if the AddDelegatedAdminRightsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDelegatedAdminRightsRequest{}

// AddDelegatedAdminRightsRequest struct for AddDelegatedAdminRightsRequest
type AddDelegatedAdminRightsRequest struct {
	Schemas []EnumdelegatedAdminRightsSchemaUrn `json:"schemas,omitempty"`
	// A description for this Delegated Admin Rights
	Description *string `json:"description,omitempty"`
	// Indicates whether the Delegated Admin Rights is enabled.
	Enabled bool `json:"enabled"`
	// Specifies the DN of an administrative user who has authority to manage resources. Either admin-user-dn or admin-group-dn must be specified, but not both.
	AdminUserDN *string `json:"adminUserDN,omitempty"`
	// Specifies the DN of a group of administrative users who have authority to manage resources. Either admin-user-dn or admin-group-dn must be specified, but not both.
	AdminGroupDN *string `json:"adminGroupDN,omitempty"`
	// Name of the new Delegated Admin Rights
	RightsName string `json:"rightsName"`
}

// NewAddDelegatedAdminRightsRequest instantiates a new AddDelegatedAdminRightsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDelegatedAdminRightsRequest(enabled bool, rightsName string) *AddDelegatedAdminRightsRequest {
	this := AddDelegatedAdminRightsRequest{}
	this.Enabled = enabled
	this.RightsName = rightsName
	return &this
}

// NewAddDelegatedAdminRightsRequestWithDefaults instantiates a new AddDelegatedAdminRightsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDelegatedAdminRightsRequestWithDefaults() *AddDelegatedAdminRightsRequest {
	this := AddDelegatedAdminRightsRequest{}
	return &this
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *AddDelegatedAdminRightsRequest) GetSchemas() []EnumdelegatedAdminRightsSchemaUrn {
	if o == nil || IsNil(o.Schemas) {
		var ret []EnumdelegatedAdminRightsSchemaUrn
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDelegatedAdminRightsRequest) GetSchemasOk() ([]EnumdelegatedAdminRightsSchemaUrn, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *AddDelegatedAdminRightsRequest) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []EnumdelegatedAdminRightsSchemaUrn and assigns it to the Schemas field.
func (o *AddDelegatedAdminRightsRequest) SetSchemas(v []EnumdelegatedAdminRightsSchemaUrn) {
	o.Schemas = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDelegatedAdminRightsRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDelegatedAdminRightsRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDelegatedAdminRightsRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDelegatedAdminRightsRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddDelegatedAdminRightsRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddDelegatedAdminRightsRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddDelegatedAdminRightsRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAdminUserDN returns the AdminUserDN field value if set, zero value otherwise.
func (o *AddDelegatedAdminRightsRequest) GetAdminUserDN() string {
	if o == nil || IsNil(o.AdminUserDN) {
		var ret string
		return ret
	}
	return *o.AdminUserDN
}

// GetAdminUserDNOk returns a tuple with the AdminUserDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDelegatedAdminRightsRequest) GetAdminUserDNOk() (*string, bool) {
	if o == nil || IsNil(o.AdminUserDN) {
		return nil, false
	}
	return o.AdminUserDN, true
}

// HasAdminUserDN returns a boolean if a field has been set.
func (o *AddDelegatedAdminRightsRequest) HasAdminUserDN() bool {
	if o != nil && !IsNil(o.AdminUserDN) {
		return true
	}

	return false
}

// SetAdminUserDN gets a reference to the given string and assigns it to the AdminUserDN field.
func (o *AddDelegatedAdminRightsRequest) SetAdminUserDN(v string) {
	o.AdminUserDN = &v
}

// GetAdminGroupDN returns the AdminGroupDN field value if set, zero value otherwise.
func (o *AddDelegatedAdminRightsRequest) GetAdminGroupDN() string {
	if o == nil || IsNil(o.AdminGroupDN) {
		var ret string
		return ret
	}
	return *o.AdminGroupDN
}

// GetAdminGroupDNOk returns a tuple with the AdminGroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDelegatedAdminRightsRequest) GetAdminGroupDNOk() (*string, bool) {
	if o == nil || IsNil(o.AdminGroupDN) {
		return nil, false
	}
	return o.AdminGroupDN, true
}

// HasAdminGroupDN returns a boolean if a field has been set.
func (o *AddDelegatedAdminRightsRequest) HasAdminGroupDN() bool {
	if o != nil && !IsNil(o.AdminGroupDN) {
		return true
	}

	return false
}

// SetAdminGroupDN gets a reference to the given string and assigns it to the AdminGroupDN field.
func (o *AddDelegatedAdminRightsRequest) SetAdminGroupDN(v string) {
	o.AdminGroupDN = &v
}

// GetRightsName returns the RightsName field value
func (o *AddDelegatedAdminRightsRequest) GetRightsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RightsName
}

// GetRightsNameOk returns a tuple with the RightsName field value
// and a boolean to check if the value has been set.
func (o *AddDelegatedAdminRightsRequest) GetRightsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RightsName, true
}

// SetRightsName sets field value
func (o *AddDelegatedAdminRightsRequest) SetRightsName(v string) {
	o.RightsName = v
}

func (o AddDelegatedAdminRightsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDelegatedAdminRightsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.AdminUserDN) {
		toSerialize["adminUserDN"] = o.AdminUserDN
	}
	if !IsNil(o.AdminGroupDN) {
		toSerialize["adminGroupDN"] = o.AdminGroupDN
	}
	toSerialize["rightsName"] = o.RightsName
	return toSerialize, nil
}

type NullableAddDelegatedAdminRightsRequest struct {
	value *AddDelegatedAdminRightsRequest
	isSet bool
}

func (v NullableAddDelegatedAdminRightsRequest) Get() *AddDelegatedAdminRightsRequest {
	return v.value
}

func (v *NullableAddDelegatedAdminRightsRequest) Set(val *AddDelegatedAdminRightsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDelegatedAdminRightsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDelegatedAdminRightsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDelegatedAdminRightsRequest(val *AddDelegatedAdminRightsRequest) *NullableAddDelegatedAdminRightsRequest {
	return &NullableAddDelegatedAdminRightsRequest{value: val, isSet: true}
}

func (v NullableAddDelegatedAdminRightsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDelegatedAdminRightsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
