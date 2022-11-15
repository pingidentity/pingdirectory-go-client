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

// AddIdentifyReferencesVirtualAttributeRequest struct for AddIdentifyReferencesVirtualAttributeRequest
type AddIdentifyReferencesVirtualAttributeRequest struct {
	// Name of the new Virtual Attribute
	Name string `json:"name"`
	Schemas []EnumidentifyReferencesVirtualAttributeSchemaUrn `json:"schemas"`
	ReferencedByAttribute []string `json:"referencedByAttribute"`
	ReferenceSearchBaseDN []string `json:"referenceSearchBaseDN,omitempty"`
	// A description for this Virtual Attribute
	Description *string `json:"description,omitempty"`
	// Indicates whether the Virtual Attribute is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute.
	AttributeType string `json:"attributeType"`
	BaseDN []string `json:"baseDN,omitempty"`
	GroupDN []string `json:"groupDN,omitempty"`
	Filter []string `json:"filter,omitempty"`
	ClientConnectionPolicy []string `json:"clientConnectionPolicy,omitempty"`
	ConflictBehavior *EnumvirtualAttributeConflictBehaviorProp `json:"conflictBehavior,omitempty"`
	// Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type.
	RequireExplicitRequestByName *bool `json:"requireExplicitRequestByName,omitempty"`
	// Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry.
	MultipleVirtualAttributeEvaluationOrderIndex *int32 `json:"multipleVirtualAttributeEvaluationOrderIndex,omitempty"`
	MultipleVirtualAttributeMergeBehavior *EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp `json:"multipleVirtualAttributeMergeBehavior,omitempty"`
	// Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server.
	AllowIndexConflicts *bool `json:"allowIndexConflicts,omitempty"`
}

// NewAddIdentifyReferencesVirtualAttributeRequest instantiates a new AddIdentifyReferencesVirtualAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIdentifyReferencesVirtualAttributeRequest(name string, schemas []EnumidentifyReferencesVirtualAttributeSchemaUrn, referencedByAttribute []string, enabled bool, attributeType string) *AddIdentifyReferencesVirtualAttributeRequest {
	this := AddIdentifyReferencesVirtualAttributeRequest{}
	this.Name = name
	this.Schemas = schemas
	this.ReferencedByAttribute = referencedByAttribute
	this.Enabled = enabled
	this.AttributeType = attributeType
	return &this
}

// NewAddIdentifyReferencesVirtualAttributeRequestWithDefaults instantiates a new AddIdentifyReferencesVirtualAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIdentifyReferencesVirtualAttributeRequestWithDefaults() *AddIdentifyReferencesVirtualAttributeRequest {
	this := AddIdentifyReferencesVirtualAttributeRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetName(v string) {
	o.Name = v
}

// GetSchemas returns the Schemas field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetSchemas() []EnumidentifyReferencesVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnumidentifyReferencesVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetSchemasOk() ([]EnumidentifyReferencesVirtualAttributeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetSchemas(v []EnumidentifyReferencesVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetReferencedByAttribute returns the ReferencedByAttribute field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetReferencedByAttribute() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReferencedByAttribute
}

// GetReferencedByAttributeOk returns a tuple with the ReferencedByAttribute field value
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetReferencedByAttributeOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ReferencedByAttribute, true
}

// SetReferencedByAttribute sets field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetReferencedByAttribute(v []string) {
	o.ReferencedByAttribute = v
}

// GetReferenceSearchBaseDN returns the ReferenceSearchBaseDN field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetReferenceSearchBaseDN() []string {
	if o == nil || isNil(o.ReferenceSearchBaseDN) {
		var ret []string
		return ret
	}
	return o.ReferenceSearchBaseDN
}

// GetReferenceSearchBaseDNOk returns a tuple with the ReferenceSearchBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetReferenceSearchBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.ReferenceSearchBaseDN) {
    return nil, false
	}
	return o.ReferenceSearchBaseDN, true
}

// HasReferenceSearchBaseDN returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasReferenceSearchBaseDN() bool {
	if o != nil && !isNil(o.ReferenceSearchBaseDN) {
		return true
	}

	return false
}

// SetReferenceSearchBaseDN gets a reference to the given []string and assigns it to the ReferenceSearchBaseDN field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetReferenceSearchBaseDN(v []string) {
	o.ReferenceSearchBaseDN = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAttributeType returns the AttributeType field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetBaseDN() []string {
	if o == nil || isNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.BaseDN) {
    return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetGroupDN() []string {
	if o == nil || isNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.GroupDN) {
    return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasGroupDN() bool {
	if o != nil && !isNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetFilter() []string {
	if o == nil || isNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetFilterOk() ([]string, bool) {
	if o == nil || isNil(o.Filter) {
    return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetClientConnectionPolicy() []string {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || isNil(o.ClientConnectionPolicy) {
    return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasClientConnectionPolicy() bool {
	if o != nil && !isNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || isNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || isNil(o.ConflictBehavior) {
    return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasConflictBehavior() bool {
	if o != nil && !isNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetRequireExplicitRequestByName() bool {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
    return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasRequireExplicitRequestByName() bool {
	if o != nil && !isNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int32 {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int32
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
    return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int32 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
    return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetAllowIndexConflicts() bool {
	if o == nil || isNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowIndexConflicts) {
    return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *AddIdentifyReferencesVirtualAttributeRequest) HasAllowIndexConflicts() bool {
	if o != nil && !isNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *AddIdentifyReferencesVirtualAttributeRequest) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

func (o AddIdentifyReferencesVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["referencedByAttribute"] = o.ReferencedByAttribute
	}
	if !isNil(o.ReferenceSearchBaseDN) {
		toSerialize["referenceSearchBaseDN"] = o.ReferenceSearchBaseDN
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["attributeType"] = o.AttributeType
	}
	if !isNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !isNil(o.GroupDN) {
		toSerialize["groupDN"] = o.GroupDN
	}
	if !isNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !isNil(o.ClientConnectionPolicy) {
		toSerialize["clientConnectionPolicy"] = o.ClientConnectionPolicy
	}
	if !isNil(o.ConflictBehavior) {
		toSerialize["conflictBehavior"] = o.ConflictBehavior
	}
	if !isNil(o.RequireExplicitRequestByName) {
		toSerialize["requireExplicitRequestByName"] = o.RequireExplicitRequestByName
	}
	if !isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		toSerialize["multipleVirtualAttributeEvaluationOrderIndex"] = o.MultipleVirtualAttributeEvaluationOrderIndex
	}
	if !isNil(o.MultipleVirtualAttributeMergeBehavior) {
		toSerialize["multipleVirtualAttributeMergeBehavior"] = o.MultipleVirtualAttributeMergeBehavior
	}
	if !isNil(o.AllowIndexConflicts) {
		toSerialize["allowIndexConflicts"] = o.AllowIndexConflicts
	}
	return json.Marshal(toSerialize)
}

type NullableAddIdentifyReferencesVirtualAttributeRequest struct {
	value *AddIdentifyReferencesVirtualAttributeRequest
	isSet bool
}

func (v NullableAddIdentifyReferencesVirtualAttributeRequest) Get() *AddIdentifyReferencesVirtualAttributeRequest {
	return v.value
}

func (v *NullableAddIdentifyReferencesVirtualAttributeRequest) Set(val *AddIdentifyReferencesVirtualAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIdentifyReferencesVirtualAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIdentifyReferencesVirtualAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIdentifyReferencesVirtualAttributeRequest(val *AddIdentifyReferencesVirtualAttributeRequest) *NullableAddIdentifyReferencesVirtualAttributeRequest {
	return &NullableAddIdentifyReferencesVirtualAttributeRequest{value: val, isSet: true}
}

func (v NullableAddIdentifyReferencesVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIdentifyReferencesVirtualAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


