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

// DnJoinVirtualAttributeShared struct for DnJoinVirtualAttributeShared
type DnJoinVirtualAttributeShared struct {
	Schemas []EnumdnJoinVirtualAttributeSchemaUrn `json:"schemas"`
	// The attribute whose values are the DNs of the entries to be joined with the search result entry.
	JoinDNAttribute string `json:"joinDNAttribute"`
	JoinBaseDNType EnumvirtualAttributeJoinBaseDNTypeProp `json:"joinBaseDNType"`
	// The fixed, administrator-specified base DN for the internal searches used to identify joined entries.
	JoinCustomBaseDN *string `json:"joinCustomBaseDN,omitempty"`
	JoinScope *EnumvirtualAttributeJoinScopeProp `json:"joinScope,omitempty"`
	// The maximum number of entries that may be joined with the source entry, which also corresponds to the maximum number of values that the virtual attribute provider will generate for an entry.
	JoinSizeLimit *int32 `json:"joinSizeLimit,omitempty"`
	// An optional filter that specifies additional criteria for identifying joined entries. If a join-filter value is specified, then only entries matching that filter (in addition to satisfying the other join criteria) will be joined with the search result entry.
	JoinFilter *string `json:"joinFilter,omitempty"`
	JoinAttribute []string `json:"joinAttribute,omitempty"`
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

// NewDnJoinVirtualAttributeShared instantiates a new DnJoinVirtualAttributeShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnJoinVirtualAttributeShared(schemas []EnumdnJoinVirtualAttributeSchemaUrn, joinDNAttribute string, joinBaseDNType EnumvirtualAttributeJoinBaseDNTypeProp, enabled bool, attributeType string) *DnJoinVirtualAttributeShared {
	this := DnJoinVirtualAttributeShared{}
	this.Schemas = schemas
	this.JoinDNAttribute = joinDNAttribute
	this.JoinBaseDNType = joinBaseDNType
	this.Enabled = enabled
	this.AttributeType = attributeType
	return &this
}

// NewDnJoinVirtualAttributeSharedWithDefaults instantiates a new DnJoinVirtualAttributeShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnJoinVirtualAttributeSharedWithDefaults() *DnJoinVirtualAttributeShared {
	this := DnJoinVirtualAttributeShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *DnJoinVirtualAttributeShared) GetSchemas() []EnumdnJoinVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnumdnJoinVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetSchemasOk() ([]EnumdnJoinVirtualAttributeSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DnJoinVirtualAttributeShared) SetSchemas(v []EnumdnJoinVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetJoinDNAttribute returns the JoinDNAttribute field value
func (o *DnJoinVirtualAttributeShared) GetJoinDNAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinDNAttribute
}

// GetJoinDNAttributeOk returns a tuple with the JoinDNAttribute field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetJoinDNAttributeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.JoinDNAttribute, true
}

// SetJoinDNAttribute sets field value
func (o *DnJoinVirtualAttributeShared) SetJoinDNAttribute(v string) {
	o.JoinDNAttribute = v
}

// GetJoinBaseDNType returns the JoinBaseDNType field value
func (o *DnJoinVirtualAttributeShared) GetJoinBaseDNType() EnumvirtualAttributeJoinBaseDNTypeProp {
	if o == nil {
		var ret EnumvirtualAttributeJoinBaseDNTypeProp
		return ret
	}

	return o.JoinBaseDNType
}

// GetJoinBaseDNTypeOk returns a tuple with the JoinBaseDNType field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetJoinBaseDNTypeOk() (*EnumvirtualAttributeJoinBaseDNTypeProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.JoinBaseDNType, true
}

// SetJoinBaseDNType sets field value
func (o *DnJoinVirtualAttributeShared) SetJoinBaseDNType(v EnumvirtualAttributeJoinBaseDNTypeProp) {
	o.JoinBaseDNType = v
}

// GetJoinCustomBaseDN returns the JoinCustomBaseDN field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetJoinCustomBaseDN() string {
	if o == nil || isNil(o.JoinCustomBaseDN) {
		var ret string
		return ret
	}
	return *o.JoinCustomBaseDN
}

// GetJoinCustomBaseDNOk returns a tuple with the JoinCustomBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetJoinCustomBaseDNOk() (*string, bool) {
	if o == nil || isNil(o.JoinCustomBaseDN) {
    return nil, false
	}
	return o.JoinCustomBaseDN, true
}

// HasJoinCustomBaseDN returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasJoinCustomBaseDN() bool {
	if o != nil && !isNil(o.JoinCustomBaseDN) {
		return true
	}

	return false
}

// SetJoinCustomBaseDN gets a reference to the given string and assigns it to the JoinCustomBaseDN field.
func (o *DnJoinVirtualAttributeShared) SetJoinCustomBaseDN(v string) {
	o.JoinCustomBaseDN = &v
}

// GetJoinScope returns the JoinScope field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetJoinScope() EnumvirtualAttributeJoinScopeProp {
	if o == nil || isNil(o.JoinScope) {
		var ret EnumvirtualAttributeJoinScopeProp
		return ret
	}
	return *o.JoinScope
}

// GetJoinScopeOk returns a tuple with the JoinScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetJoinScopeOk() (*EnumvirtualAttributeJoinScopeProp, bool) {
	if o == nil || isNil(o.JoinScope) {
    return nil, false
	}
	return o.JoinScope, true
}

// HasJoinScope returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasJoinScope() bool {
	if o != nil && !isNil(o.JoinScope) {
		return true
	}

	return false
}

// SetJoinScope gets a reference to the given EnumvirtualAttributeJoinScopeProp and assigns it to the JoinScope field.
func (o *DnJoinVirtualAttributeShared) SetJoinScope(v EnumvirtualAttributeJoinScopeProp) {
	o.JoinScope = &v
}

// GetJoinSizeLimit returns the JoinSizeLimit field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetJoinSizeLimit() int32 {
	if o == nil || isNil(o.JoinSizeLimit) {
		var ret int32
		return ret
	}
	return *o.JoinSizeLimit
}

// GetJoinSizeLimitOk returns a tuple with the JoinSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetJoinSizeLimitOk() (*int32, bool) {
	if o == nil || isNil(o.JoinSizeLimit) {
    return nil, false
	}
	return o.JoinSizeLimit, true
}

// HasJoinSizeLimit returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasJoinSizeLimit() bool {
	if o != nil && !isNil(o.JoinSizeLimit) {
		return true
	}

	return false
}

// SetJoinSizeLimit gets a reference to the given int32 and assigns it to the JoinSizeLimit field.
func (o *DnJoinVirtualAttributeShared) SetJoinSizeLimit(v int32) {
	o.JoinSizeLimit = &v
}

// GetJoinFilter returns the JoinFilter field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetJoinFilter() string {
	if o == nil || isNil(o.JoinFilter) {
		var ret string
		return ret
	}
	return *o.JoinFilter
}

// GetJoinFilterOk returns a tuple with the JoinFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetJoinFilterOk() (*string, bool) {
	if o == nil || isNil(o.JoinFilter) {
    return nil, false
	}
	return o.JoinFilter, true
}

// HasJoinFilter returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasJoinFilter() bool {
	if o != nil && !isNil(o.JoinFilter) {
		return true
	}

	return false
}

// SetJoinFilter gets a reference to the given string and assigns it to the JoinFilter field.
func (o *DnJoinVirtualAttributeShared) SetJoinFilter(v string) {
	o.JoinFilter = &v
}

// GetJoinAttribute returns the JoinAttribute field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetJoinAttribute() []string {
	if o == nil || isNil(o.JoinAttribute) {
		var ret []string
		return ret
	}
	return o.JoinAttribute
}

// GetJoinAttributeOk returns a tuple with the JoinAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetJoinAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.JoinAttribute) {
    return nil, false
	}
	return o.JoinAttribute, true
}

// HasJoinAttribute returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasJoinAttribute() bool {
	if o != nil && !isNil(o.JoinAttribute) {
		return true
	}

	return false
}

// SetJoinAttribute gets a reference to the given []string and assigns it to the JoinAttribute field.
func (o *DnJoinVirtualAttributeShared) SetJoinAttribute(v []string) {
	o.JoinAttribute = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DnJoinVirtualAttributeShared) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *DnJoinVirtualAttributeShared) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DnJoinVirtualAttributeShared) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAttributeType returns the AttributeType field value
func (o *DnJoinVirtualAttributeShared) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *DnJoinVirtualAttributeShared) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetBaseDN() []string {
	if o == nil || isNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.BaseDN) {
    return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *DnJoinVirtualAttributeShared) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetGroupDN() []string {
	if o == nil || isNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.GroupDN) {
    return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasGroupDN() bool {
	if o != nil && !isNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *DnJoinVirtualAttributeShared) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetFilter() []string {
	if o == nil || isNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetFilterOk() ([]string, bool) {
	if o == nil || isNil(o.Filter) {
    return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *DnJoinVirtualAttributeShared) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetClientConnectionPolicy() []string {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || isNil(o.ClientConnectionPolicy) {
    return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasClientConnectionPolicy() bool {
	if o != nil && !isNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *DnJoinVirtualAttributeShared) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || isNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || isNil(o.ConflictBehavior) {
    return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasConflictBehavior() bool {
	if o != nil && !isNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *DnJoinVirtualAttributeShared) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetRequireExplicitRequestByName() bool {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
    return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasRequireExplicitRequestByName() bool {
	if o != nil && !isNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *DnJoinVirtualAttributeShared) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetMultipleVirtualAttributeEvaluationOrderIndex() int32 {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int32
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
    return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int32 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *DnJoinVirtualAttributeShared) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
    return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *DnJoinVirtualAttributeShared) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeShared) GetAllowIndexConflicts() bool {
	if o == nil || isNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeShared) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowIndexConflicts) {
    return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeShared) HasAllowIndexConflicts() bool {
	if o != nil && !isNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *DnJoinVirtualAttributeShared) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

func (o DnJoinVirtualAttributeShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["joinDNAttribute"] = o.JoinDNAttribute
	}
	if true {
		toSerialize["joinBaseDNType"] = o.JoinBaseDNType
	}
	if !isNil(o.JoinCustomBaseDN) {
		toSerialize["joinCustomBaseDN"] = o.JoinCustomBaseDN
	}
	if !isNil(o.JoinScope) {
		toSerialize["joinScope"] = o.JoinScope
	}
	if !isNil(o.JoinSizeLimit) {
		toSerialize["joinSizeLimit"] = o.JoinSizeLimit
	}
	if !isNil(o.JoinFilter) {
		toSerialize["joinFilter"] = o.JoinFilter
	}
	if !isNil(o.JoinAttribute) {
		toSerialize["joinAttribute"] = o.JoinAttribute
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

type NullableDnJoinVirtualAttributeShared struct {
	value *DnJoinVirtualAttributeShared
	isSet bool
}

func (v NullableDnJoinVirtualAttributeShared) Get() *DnJoinVirtualAttributeShared {
	return v.value
}

func (v *NullableDnJoinVirtualAttributeShared) Set(val *DnJoinVirtualAttributeShared) {
	v.value = val
	v.isSet = true
}

func (v NullableDnJoinVirtualAttributeShared) IsSet() bool {
	return v.isSet
}

func (v *NullableDnJoinVirtualAttributeShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnJoinVirtualAttributeShared(val *DnJoinVirtualAttributeShared) *NullableDnJoinVirtualAttributeShared {
	return &NullableDnJoinVirtualAttributeShared{value: val, isSet: true}
}

func (v NullableDnJoinVirtualAttributeShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnJoinVirtualAttributeShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


