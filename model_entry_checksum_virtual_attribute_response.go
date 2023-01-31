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

// EntryChecksumVirtualAttributeResponse struct for EntryChecksumVirtualAttributeResponse
type EntryChecksumVirtualAttributeResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumentryChecksumVirtualAttributeSchemaUrn       `json:"schemas"`
	// Name of the Virtual Attribute
	Id               string                                    `json:"id"`
	ConflictBehavior *EnumvirtualAttributeConflictBehaviorProp `json:"conflictBehavior,omitempty"`
	// Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute.
	AttributeType string `json:"attributeType"`
	// Indicates whether all operational attributes should be excluded from the generated checksum.
	ExcludeOperationalAttributes *bool `json:"excludeOperationalAttributes,omitempty"`
	// Specifies the attributes that should be excluded from the checksum calculation.
	ExcludedAttribute []string `json:"excludedAttribute,omitempty"`
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

// NewEntryChecksumVirtualAttributeResponse instantiates a new EntryChecksumVirtualAttributeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryChecksumVirtualAttributeResponse(schemas []EnumentryChecksumVirtualAttributeSchemaUrn, id string, attributeType string, enabled bool) *EntryChecksumVirtualAttributeResponse {
	this := EntryChecksumVirtualAttributeResponse{}
	this.Schemas = schemas
	this.Id = id
	this.AttributeType = attributeType
	this.Enabled = enabled
	return &this
}

// NewEntryChecksumVirtualAttributeResponseWithDefaults instantiates a new EntryChecksumVirtualAttributeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryChecksumVirtualAttributeResponseWithDefaults() *EntryChecksumVirtualAttributeResponse {
	this := EntryChecksumVirtualAttributeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *EntryChecksumVirtualAttributeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *EntryChecksumVirtualAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *EntryChecksumVirtualAttributeResponse) GetSchemas() []EnumentryChecksumVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnumentryChecksumVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetSchemasOk() ([]EnumentryChecksumVirtualAttributeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *EntryChecksumVirtualAttributeResponse) SetSchemas(v []EnumentryChecksumVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *EntryChecksumVirtualAttributeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntryChecksumVirtualAttributeResponse) SetId(v string) {
	o.Id = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || isNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || isNil(o.ConflictBehavior) {
		return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasConflictBehavior() bool {
	if o != nil && !isNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *EntryChecksumVirtualAttributeResponse) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetAttributeType returns the AttributeType field value
func (o *EntryChecksumVirtualAttributeResponse) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *EntryChecksumVirtualAttributeResponse) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetExcludeOperationalAttributes returns the ExcludeOperationalAttributes field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetExcludeOperationalAttributes() bool {
	if o == nil || isNil(o.ExcludeOperationalAttributes) {
		var ret bool
		return ret
	}
	return *o.ExcludeOperationalAttributes
}

// GetExcludeOperationalAttributesOk returns a tuple with the ExcludeOperationalAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetExcludeOperationalAttributesOk() (*bool, bool) {
	if o == nil || isNil(o.ExcludeOperationalAttributes) {
		return nil, false
	}
	return o.ExcludeOperationalAttributes, true
}

// HasExcludeOperationalAttributes returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasExcludeOperationalAttributes() bool {
	if o != nil && !isNil(o.ExcludeOperationalAttributes) {
		return true
	}

	return false
}

// SetExcludeOperationalAttributes gets a reference to the given bool and assigns it to the ExcludeOperationalAttributes field.
func (o *EntryChecksumVirtualAttributeResponse) SetExcludeOperationalAttributes(v bool) {
	o.ExcludeOperationalAttributes = &v
}

// GetExcludedAttribute returns the ExcludedAttribute field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetExcludedAttribute() []string {
	if o == nil || isNil(o.ExcludedAttribute) {
		var ret []string
		return ret
	}
	return o.ExcludedAttribute
}

// GetExcludedAttributeOk returns a tuple with the ExcludedAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetExcludedAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.ExcludedAttribute) {
		return nil, false
	}
	return o.ExcludedAttribute, true
}

// HasExcludedAttribute returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasExcludedAttribute() bool {
	if o != nil && !isNil(o.ExcludedAttribute) {
		return true
	}

	return false
}

// SetExcludedAttribute gets a reference to the given []string and assigns it to the ExcludedAttribute field.
func (o *EntryChecksumVirtualAttributeResponse) SetExcludedAttribute(v []string) {
	o.ExcludedAttribute = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntryChecksumVirtualAttributeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *EntryChecksumVirtualAttributeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *EntryChecksumVirtualAttributeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetBaseDN() []string {
	if o == nil || isNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *EntryChecksumVirtualAttributeResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetGroupDN() []string {
	if o == nil || isNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.GroupDN) {
		return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasGroupDN() bool {
	if o != nil && !isNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *EntryChecksumVirtualAttributeResponse) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetFilter() []string {
	if o == nil || isNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetFilterOk() ([]string, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *EntryChecksumVirtualAttributeResponse) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetClientConnectionPolicy() []string {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasClientConnectionPolicy() bool {
	if o != nil && !isNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *EntryChecksumVirtualAttributeResponse) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetRequireExplicitRequestByName() bool {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasRequireExplicitRequestByName() bool {
	if o != nil && !isNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *EntryChecksumVirtualAttributeResponse) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndex() int32 {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int32
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int32 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *EntryChecksumVirtualAttributeResponse) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *EntryChecksumVirtualAttributeResponse) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *EntryChecksumVirtualAttributeResponse) GetAllowIndexConflicts() bool {
	if o == nil || isNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryChecksumVirtualAttributeResponse) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowIndexConflicts) {
		return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *EntryChecksumVirtualAttributeResponse) HasAllowIndexConflicts() bool {
	if o != nil && !isNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *EntryChecksumVirtualAttributeResponse) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

func (o EntryChecksumVirtualAttributeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ConflictBehavior) {
		toSerialize["conflictBehavior"] = o.ConflictBehavior
	}
	if true {
		toSerialize["attributeType"] = o.AttributeType
	}
	if !isNil(o.ExcludeOperationalAttributes) {
		toSerialize["excludeOperationalAttributes"] = o.ExcludeOperationalAttributes
	}
	if !isNil(o.ExcludedAttribute) {
		toSerialize["excludedAttribute"] = o.ExcludedAttribute
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

type NullableEntryChecksumVirtualAttributeResponse struct {
	value *EntryChecksumVirtualAttributeResponse
	isSet bool
}

func (v NullableEntryChecksumVirtualAttributeResponse) Get() *EntryChecksumVirtualAttributeResponse {
	return v.value
}

func (v *NullableEntryChecksumVirtualAttributeResponse) Set(val *EntryChecksumVirtualAttributeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryChecksumVirtualAttributeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryChecksumVirtualAttributeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryChecksumVirtualAttributeResponse(val *EntryChecksumVirtualAttributeResponse) *NullableEntryChecksumVirtualAttributeResponse {
	return &NullableEntryChecksumVirtualAttributeResponse{value: val, isSet: true}
}

func (v NullableEntryChecksumVirtualAttributeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryChecksumVirtualAttributeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
