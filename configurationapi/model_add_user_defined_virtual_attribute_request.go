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

// checks if the AddUserDefinedVirtualAttributeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUserDefinedVirtualAttributeRequest{}

// AddUserDefinedVirtualAttributeRequest struct for AddUserDefinedVirtualAttributeRequest
type AddUserDefinedVirtualAttributeRequest struct {
	Schemas []EnumuserDefinedVirtualAttributeSchemaUrn `json:"schemas"`
	// Specifies the values to be included in the virtual attribute.
	Value []string `json:"value"`
	// A description for this Virtual Attribute
	Description *string `json:"description,omitempty"`
	// Indicates whether the Virtual Attribute is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute.
	AttributeType string `json:"attributeType"`
	// Specifies the base DNs for the branches containing entries that are eligible to use this virtual attribute.
	BaseDN []string `json:"baseDN,omitempty"`
	// Specifies the DNs of the groups whose members can be eligible to use this virtual attribute.
	GroupDN []string `json:"groupDN,omitempty"`
	// Specifies the search filters to be applied against entries to determine if the virtual attribute is to be generated for those entries.
	Filter []string `json:"filter,omitempty"`
	// Specifies a set of client connection policies for which this Virtual Attribute should be generated. If this is undefined, then this Virtual Attribute will always be generated. If it is associated with one or more client connection policies, then this Virtual Attribute will be generated only for operations requested by clients assigned to one of those client connection policies.
	ClientConnectionPolicy []string                                  `json:"clientConnectionPolicy,omitempty"`
	ConflictBehavior       *EnumvirtualAttributeConflictBehaviorProp `json:"conflictBehavior,omitempty"`
	// Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type.
	RequireExplicitRequestByName *bool `json:"requireExplicitRequestByName,omitempty"`
	// Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry.
	MultipleVirtualAttributeEvaluationOrderIndex *int64                                                         `json:"multipleVirtualAttributeEvaluationOrderIndex,omitempty"`
	MultipleVirtualAttributeMergeBehavior        *EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp `json:"multipleVirtualAttributeMergeBehavior,omitempty"`
	// Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server.
	AllowIndexConflicts *bool `json:"allowIndexConflicts,omitempty"`
	// Name of the new Virtual Attribute
	Name string `json:"name"`
}

// NewAddUserDefinedVirtualAttributeRequest instantiates a new AddUserDefinedVirtualAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserDefinedVirtualAttributeRequest(schemas []EnumuserDefinedVirtualAttributeSchemaUrn, value []string, enabled bool, attributeType string, name string) *AddUserDefinedVirtualAttributeRequest {
	this := AddUserDefinedVirtualAttributeRequest{}
	this.Schemas = schemas
	this.Value = value
	this.Enabled = enabled
	this.AttributeType = attributeType
	this.Name = name
	return &this
}

// NewAddUserDefinedVirtualAttributeRequestWithDefaults instantiates a new AddUserDefinedVirtualAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserDefinedVirtualAttributeRequestWithDefaults() *AddUserDefinedVirtualAttributeRequest {
	this := AddUserDefinedVirtualAttributeRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddUserDefinedVirtualAttributeRequest) GetSchemas() []EnumuserDefinedVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnumuserDefinedVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetSchemasOk() ([]EnumuserDefinedVirtualAttributeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddUserDefinedVirtualAttributeRequest) SetSchemas(v []EnumuserDefinedVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetValue returns the Value field value
func (o *AddUserDefinedVirtualAttributeRequest) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetValueOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *AddUserDefinedVirtualAttributeRequest) SetValue(v []string) {
	o.Value = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddUserDefinedVirtualAttributeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddUserDefinedVirtualAttributeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddUserDefinedVirtualAttributeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAttributeType returns the AttributeType field value
func (o *AddUserDefinedVirtualAttributeRequest) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *AddUserDefinedVirtualAttributeRequest) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetBaseDN() []string {
	if o == nil || IsNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *AddUserDefinedVirtualAttributeRequest) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetGroupDN() []string {
	if o == nil || IsNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetGroupDNOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupDN) {
		return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasGroupDN() bool {
	if o != nil && !IsNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *AddUserDefinedVirtualAttributeRequest) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetFilter() []string {
	if o == nil || IsNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *AddUserDefinedVirtualAttributeRequest) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetClientConnectionPolicy() []string {
	if o == nil || IsNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || IsNil(o.ClientConnectionPolicy) {
		return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasClientConnectionPolicy() bool {
	if o != nil && !IsNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *AddUserDefinedVirtualAttributeRequest) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || IsNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || IsNil(o.ConflictBehavior) {
		return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasConflictBehavior() bool {
	if o != nil && !IsNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *AddUserDefinedVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetRequireExplicitRequestByName() bool {
	if o == nil || IsNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireExplicitRequestByName) {
		return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasRequireExplicitRequestByName() bool {
	if o != nil && !IsNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *AddUserDefinedVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int64 {
	if o == nil || IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int64
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int64 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *AddUserDefinedVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int64) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *AddUserDefinedVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *AddUserDefinedVirtualAttributeRequest) GetAllowIndexConflicts() bool {
	if o == nil || IsNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowIndexConflicts) {
		return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *AddUserDefinedVirtualAttributeRequest) HasAllowIndexConflicts() bool {
	if o != nil && !IsNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *AddUserDefinedVirtualAttributeRequest) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

// GetName returns the Name field value
func (o *AddUserDefinedVirtualAttributeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddUserDefinedVirtualAttributeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddUserDefinedVirtualAttributeRequest) SetName(v string) {
	o.Name = v
}

func (o AddUserDefinedVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUserDefinedVirtualAttributeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["value"] = o.Value
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["attributeType"] = o.AttributeType
	if !IsNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !IsNil(o.GroupDN) {
		toSerialize["groupDN"] = o.GroupDN
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.ClientConnectionPolicy) {
		toSerialize["clientConnectionPolicy"] = o.ClientConnectionPolicy
	}
	if !IsNil(o.ConflictBehavior) {
		toSerialize["conflictBehavior"] = o.ConflictBehavior
	}
	if !IsNil(o.RequireExplicitRequestByName) {
		toSerialize["requireExplicitRequestByName"] = o.RequireExplicitRequestByName
	}
	if !IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		toSerialize["multipleVirtualAttributeEvaluationOrderIndex"] = o.MultipleVirtualAttributeEvaluationOrderIndex
	}
	if !IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		toSerialize["multipleVirtualAttributeMergeBehavior"] = o.MultipleVirtualAttributeMergeBehavior
	}
	if !IsNil(o.AllowIndexConflicts) {
		toSerialize["allowIndexConflicts"] = o.AllowIndexConflicts
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAddUserDefinedVirtualAttributeRequest struct {
	value *AddUserDefinedVirtualAttributeRequest
	isSet bool
}

func (v NullableAddUserDefinedVirtualAttributeRequest) Get() *AddUserDefinedVirtualAttributeRequest {
	return v.value
}

func (v *NullableAddUserDefinedVirtualAttributeRequest) Set(val *AddUserDefinedVirtualAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddUserDefinedVirtualAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddUserDefinedVirtualAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddUserDefinedVirtualAttributeRequest(val *AddUserDefinedVirtualAttributeRequest) *NullableAddUserDefinedVirtualAttributeRequest {
	return &NullableAddUserDefinedVirtualAttributeRequest{value: val, isSet: true}
}

func (v NullableAddUserDefinedVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddUserDefinedVirtualAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
