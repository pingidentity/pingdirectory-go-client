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

// AddIsMemberOfVirtualAttributeRequest struct for AddIsMemberOfVirtualAttributeRequest
type AddIsMemberOfVirtualAttributeRequest struct {
	// Name of the new Virtual Attribute
	Name             string                                    `json:"name"`
	Schemas          []EnumisMemberOfVirtualAttributeSchemaUrn `json:"schemas"`
	ConflictBehavior *EnumvirtualAttributeConflictBehaviorProp `json:"conflictBehavior,omitempty"`
	// Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute.
	AttributeType *string `json:"attributeType,omitempty"`
	// Specifies whether to only include groups in which the user is directly associated with and the membership maybe modified via the group entry. Groups in which the user's membership is derived dynamically or through nested groups will not be included.
	DirectMembershipsOnly *bool `json:"directMembershipsOnly,omitempty"`
	// A search filter that will be used to identify which groups should be included in the values of the virtual attribute. With no value defined (which is the default behavior), all groups that contain the target user will be included.
	IncludedGroupFilter  *string                                       `json:"includedGroupFilter,omitempty"`
	RewriteSearchFilters *EnumvirtualAttributeRewriteSearchFiltersProp `json:"rewriteSearchFilters,omitempty"`
	// A description for this Virtual Attribute
	Description *string `json:"description,omitempty"`
	// Indicates whether the Virtual Attribute is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies the base DNs for the branches containing entries that are eligible to use this virtual attribute.
	BaseDN []string `json:"baseDN,omitempty"`
	// Specifies the DNs of the groups whose members can be eligible to use this virtual attribute.
	GroupDN []string `json:"groupDN,omitempty"`
	// Specifies the search filters to be applied against entries to determine if the virtual attribute is to be generated for those entries.
	Filter []string `json:"filter,omitempty"`
	// Specifies a set of client connection policies for which this Virtual Attribute should be generated. If this is undefined, then this Virtual Attribute will always be generated. If it is associated with one or more client connection policies, then this Virtual Attribute will be generated only for operations requested by clients assigned to one of those client connection policies.
	ClientConnectionPolicy []string `json:"clientConnectionPolicy,omitempty"`
	// Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type.
	RequireExplicitRequestByName *bool `json:"requireExplicitRequestByName,omitempty"`
	// Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry.
	MultipleVirtualAttributeEvaluationOrderIndex *int32                                                         `json:"multipleVirtualAttributeEvaluationOrderIndex,omitempty"`
	MultipleVirtualAttributeMergeBehavior        *EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp `json:"multipleVirtualAttributeMergeBehavior,omitempty"`
	// Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server.
	AllowIndexConflicts *bool `json:"allowIndexConflicts,omitempty"`
}

// NewAddIsMemberOfVirtualAttributeRequest instantiates a new AddIsMemberOfVirtualAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIsMemberOfVirtualAttributeRequest(name string, schemas []EnumisMemberOfVirtualAttributeSchemaUrn, enabled bool) *AddIsMemberOfVirtualAttributeRequest {
	this := AddIsMemberOfVirtualAttributeRequest{}
	this.Name = name
	this.Schemas = schemas
	this.Enabled = enabled
	return &this
}

// NewAddIsMemberOfVirtualAttributeRequestWithDefaults instantiates a new AddIsMemberOfVirtualAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIsMemberOfVirtualAttributeRequestWithDefaults() *AddIsMemberOfVirtualAttributeRequest {
	this := AddIsMemberOfVirtualAttributeRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AddIsMemberOfVirtualAttributeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddIsMemberOfVirtualAttributeRequest) SetName(v string) {
	o.Name = v
}

// GetSchemas returns the Schemas field value
func (o *AddIsMemberOfVirtualAttributeRequest) GetSchemas() []EnumisMemberOfVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnumisMemberOfVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetSchemasOk() ([]EnumisMemberOfVirtualAttributeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddIsMemberOfVirtualAttributeRequest) SetSchemas(v []EnumisMemberOfVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || isNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || isNil(o.ConflictBehavior) {
		return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasConflictBehavior() bool {
	if o != nil && !isNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetAttributeType returns the AttributeType field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetAttributeType() string {
	if o == nil || isNil(o.AttributeType) {
		var ret string
		return ret
	}
	return *o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool) {
	if o == nil || isNil(o.AttributeType) {
		return nil, false
	}
	return o.AttributeType, true
}

// HasAttributeType returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasAttributeType() bool {
	if o != nil && !isNil(o.AttributeType) {
		return true
	}

	return false
}

// SetAttributeType gets a reference to the given string and assigns it to the AttributeType field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetAttributeType(v string) {
	o.AttributeType = &v
}

// GetDirectMembershipsOnly returns the DirectMembershipsOnly field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetDirectMembershipsOnly() bool {
	if o == nil || isNil(o.DirectMembershipsOnly) {
		var ret bool
		return ret
	}
	return *o.DirectMembershipsOnly
}

// GetDirectMembershipsOnlyOk returns a tuple with the DirectMembershipsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetDirectMembershipsOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.DirectMembershipsOnly) {
		return nil, false
	}
	return o.DirectMembershipsOnly, true
}

// HasDirectMembershipsOnly returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasDirectMembershipsOnly() bool {
	if o != nil && !isNil(o.DirectMembershipsOnly) {
		return true
	}

	return false
}

// SetDirectMembershipsOnly gets a reference to the given bool and assigns it to the DirectMembershipsOnly field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetDirectMembershipsOnly(v bool) {
	o.DirectMembershipsOnly = &v
}

// GetIncludedGroupFilter returns the IncludedGroupFilter field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetIncludedGroupFilter() string {
	if o == nil || isNil(o.IncludedGroupFilter) {
		var ret string
		return ret
	}
	return *o.IncludedGroupFilter
}

// GetIncludedGroupFilterOk returns a tuple with the IncludedGroupFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetIncludedGroupFilterOk() (*string, bool) {
	if o == nil || isNil(o.IncludedGroupFilter) {
		return nil, false
	}
	return o.IncludedGroupFilter, true
}

// HasIncludedGroupFilter returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasIncludedGroupFilter() bool {
	if o != nil && !isNil(o.IncludedGroupFilter) {
		return true
	}

	return false
}

// SetIncludedGroupFilter gets a reference to the given string and assigns it to the IncludedGroupFilter field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetIncludedGroupFilter(v string) {
	o.IncludedGroupFilter = &v
}

// GetRewriteSearchFilters returns the RewriteSearchFilters field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetRewriteSearchFilters() EnumvirtualAttributeRewriteSearchFiltersProp {
	if o == nil || isNil(o.RewriteSearchFilters) {
		var ret EnumvirtualAttributeRewriteSearchFiltersProp
		return ret
	}
	return *o.RewriteSearchFilters
}

// GetRewriteSearchFiltersOk returns a tuple with the RewriteSearchFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetRewriteSearchFiltersOk() (*EnumvirtualAttributeRewriteSearchFiltersProp, bool) {
	if o == nil || isNil(o.RewriteSearchFilters) {
		return nil, false
	}
	return o.RewriteSearchFilters, true
}

// HasRewriteSearchFilters returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasRewriteSearchFilters() bool {
	if o != nil && !isNil(o.RewriteSearchFilters) {
		return true
	}

	return false
}

// SetRewriteSearchFilters gets a reference to the given EnumvirtualAttributeRewriteSearchFiltersProp and assigns it to the RewriteSearchFilters field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetRewriteSearchFilters(v EnumvirtualAttributeRewriteSearchFiltersProp) {
	o.RewriteSearchFilters = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddIsMemberOfVirtualAttributeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddIsMemberOfVirtualAttributeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetBaseDN() []string {
	if o == nil || isNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetGroupDN() []string {
	if o == nil || isNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.GroupDN) {
		return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasGroupDN() bool {
	if o != nil && !isNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetFilter() []string {
	if o == nil || isNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetFilterOk() ([]string, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetClientConnectionPolicy() []string {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasClientConnectionPolicy() bool {
	if o != nil && !isNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetRequireExplicitRequestByName() bool {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasRequireExplicitRequestByName() bool {
	if o != nil && !isNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int32 {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int32
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int32 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *AddIsMemberOfVirtualAttributeRequest) GetAllowIndexConflicts() bool {
	if o == nil || isNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowIndexConflicts) {
		return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *AddIsMemberOfVirtualAttributeRequest) HasAllowIndexConflicts() bool {
	if o != nil && !isNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *AddIsMemberOfVirtualAttributeRequest) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

func (o AddIsMemberOfVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.ConflictBehavior) {
		toSerialize["conflictBehavior"] = o.ConflictBehavior
	}
	if !isNil(o.AttributeType) {
		toSerialize["attributeType"] = o.AttributeType
	}
	if !isNil(o.DirectMembershipsOnly) {
		toSerialize["directMembershipsOnly"] = o.DirectMembershipsOnly
	}
	if !isNil(o.IncludedGroupFilter) {
		toSerialize["includedGroupFilter"] = o.IncludedGroupFilter
	}
	if !isNil(o.RewriteSearchFilters) {
		toSerialize["rewriteSearchFilters"] = o.RewriteSearchFilters
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
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

type NullableAddIsMemberOfVirtualAttributeRequest struct {
	value *AddIsMemberOfVirtualAttributeRequest
	isSet bool
}

func (v NullableAddIsMemberOfVirtualAttributeRequest) Get() *AddIsMemberOfVirtualAttributeRequest {
	return v.value
}

func (v *NullableAddIsMemberOfVirtualAttributeRequest) Set(val *AddIsMemberOfVirtualAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIsMemberOfVirtualAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIsMemberOfVirtualAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIsMemberOfVirtualAttributeRequest(val *AddIsMemberOfVirtualAttributeRequest) *NullableAddIsMemberOfVirtualAttributeRequest {
	return &NullableAddIsMemberOfVirtualAttributeRequest{value: val, isSet: true}
}

func (v NullableAddIsMemberOfVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIsMemberOfVirtualAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
