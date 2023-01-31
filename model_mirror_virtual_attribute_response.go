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

// MirrorVirtualAttributeResponse struct for MirrorVirtualAttributeResponse
type MirrorVirtualAttributeResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Virtual Attribute
	Id               string                                    `json:"id"`
	Schemas          []EnummirrorVirtualAttributeSchemaUrn     `json:"schemas"`
	ConflictBehavior *EnumvirtualAttributeConflictBehaviorProp `json:"conflictBehavior,omitempty"`
	// Specifies the source attribute containing the values to use for this virtual attribute.
	SourceAttribute string `json:"sourceAttribute"`
	// Specifies the attribute containing the DN of another entry from which to obtain the source attribute providing the values for this virtual attribute.
	SourceEntryDNAttribute *string `json:"sourceEntryDNAttribute,omitempty"`
	// Specifies a DN map that will be used to identify the entry from which to obtain the source attribute providing the values for this virtual attribute.
	SourceEntryDNMap *string `json:"sourceEntryDNMap,omitempty"`
	// Indicates whether searches performed by this virtual attribute provider should be exempted from access control restrictions.
	BypassAccessControlForSearches *bool `json:"bypassAccessControlForSearches,omitempty"`
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
	ClientConnectionPolicy []string `json:"clientConnectionPolicy,omitempty"`
	// Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type.
	RequireExplicitRequestByName *bool `json:"requireExplicitRequestByName,omitempty"`
	// Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry.
	MultipleVirtualAttributeEvaluationOrderIndex *int32                                                         `json:"multipleVirtualAttributeEvaluationOrderIndex,omitempty"`
	MultipleVirtualAttributeMergeBehavior        *EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp `json:"multipleVirtualAttributeMergeBehavior,omitempty"`
	// Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server.
	AllowIndexConflicts *bool `json:"allowIndexConflicts,omitempty"`
}

// NewMirrorVirtualAttributeResponse instantiates a new MirrorVirtualAttributeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMirrorVirtualAttributeResponse(id string, schemas []EnummirrorVirtualAttributeSchemaUrn, sourceAttribute string, enabled bool, attributeType string) *MirrorVirtualAttributeResponse {
	this := MirrorVirtualAttributeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.SourceAttribute = sourceAttribute
	this.Enabled = enabled
	this.AttributeType = attributeType
	return &this
}

// NewMirrorVirtualAttributeResponseWithDefaults instantiates a new MirrorVirtualAttributeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMirrorVirtualAttributeResponseWithDefaults() *MirrorVirtualAttributeResponse {
	this := MirrorVirtualAttributeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *MirrorVirtualAttributeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *MirrorVirtualAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *MirrorVirtualAttributeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MirrorVirtualAttributeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *MirrorVirtualAttributeResponse) GetSchemas() []EnummirrorVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnummirrorVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetSchemasOk() ([]EnummirrorVirtualAttributeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *MirrorVirtualAttributeResponse) SetSchemas(v []EnummirrorVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || isNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || isNil(o.ConflictBehavior) {
		return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasConflictBehavior() bool {
	if o != nil && !isNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *MirrorVirtualAttributeResponse) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetSourceAttribute returns the SourceAttribute field value
func (o *MirrorVirtualAttributeResponse) GetSourceAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceAttribute
}

// GetSourceAttributeOk returns a tuple with the SourceAttribute field value
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetSourceAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceAttribute, true
}

// SetSourceAttribute sets field value
func (o *MirrorVirtualAttributeResponse) SetSourceAttribute(v string) {
	o.SourceAttribute = v
}

// GetSourceEntryDNAttribute returns the SourceEntryDNAttribute field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetSourceEntryDNAttribute() string {
	if o == nil || isNil(o.SourceEntryDNAttribute) {
		var ret string
		return ret
	}
	return *o.SourceEntryDNAttribute
}

// GetSourceEntryDNAttributeOk returns a tuple with the SourceEntryDNAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetSourceEntryDNAttributeOk() (*string, bool) {
	if o == nil || isNil(o.SourceEntryDNAttribute) {
		return nil, false
	}
	return o.SourceEntryDNAttribute, true
}

// HasSourceEntryDNAttribute returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasSourceEntryDNAttribute() bool {
	if o != nil && !isNil(o.SourceEntryDNAttribute) {
		return true
	}

	return false
}

// SetSourceEntryDNAttribute gets a reference to the given string and assigns it to the SourceEntryDNAttribute field.
func (o *MirrorVirtualAttributeResponse) SetSourceEntryDNAttribute(v string) {
	o.SourceEntryDNAttribute = &v
}

// GetSourceEntryDNMap returns the SourceEntryDNMap field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetSourceEntryDNMap() string {
	if o == nil || isNil(o.SourceEntryDNMap) {
		var ret string
		return ret
	}
	return *o.SourceEntryDNMap
}

// GetSourceEntryDNMapOk returns a tuple with the SourceEntryDNMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetSourceEntryDNMapOk() (*string, bool) {
	if o == nil || isNil(o.SourceEntryDNMap) {
		return nil, false
	}
	return o.SourceEntryDNMap, true
}

// HasSourceEntryDNMap returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasSourceEntryDNMap() bool {
	if o != nil && !isNil(o.SourceEntryDNMap) {
		return true
	}

	return false
}

// SetSourceEntryDNMap gets a reference to the given string and assigns it to the SourceEntryDNMap field.
func (o *MirrorVirtualAttributeResponse) SetSourceEntryDNMap(v string) {
	o.SourceEntryDNMap = &v
}

// GetBypassAccessControlForSearches returns the BypassAccessControlForSearches field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetBypassAccessControlForSearches() bool {
	if o == nil || isNil(o.BypassAccessControlForSearches) {
		var ret bool
		return ret
	}
	return *o.BypassAccessControlForSearches
}

// GetBypassAccessControlForSearchesOk returns a tuple with the BypassAccessControlForSearches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetBypassAccessControlForSearchesOk() (*bool, bool) {
	if o == nil || isNil(o.BypassAccessControlForSearches) {
		return nil, false
	}
	return o.BypassAccessControlForSearches, true
}

// HasBypassAccessControlForSearches returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasBypassAccessControlForSearches() bool {
	if o != nil && !isNil(o.BypassAccessControlForSearches) {
		return true
	}

	return false
}

// SetBypassAccessControlForSearches gets a reference to the given bool and assigns it to the BypassAccessControlForSearches field.
func (o *MirrorVirtualAttributeResponse) SetBypassAccessControlForSearches(v bool) {
	o.BypassAccessControlForSearches = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MirrorVirtualAttributeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *MirrorVirtualAttributeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MirrorVirtualAttributeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAttributeType returns the AttributeType field value
func (o *MirrorVirtualAttributeResponse) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *MirrorVirtualAttributeResponse) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetBaseDN() []string {
	if o == nil || isNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *MirrorVirtualAttributeResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetGroupDN() []string {
	if o == nil || isNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.GroupDN) {
		return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasGroupDN() bool {
	if o != nil && !isNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *MirrorVirtualAttributeResponse) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetFilter() []string {
	if o == nil || isNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetFilterOk() ([]string, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *MirrorVirtualAttributeResponse) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetClientConnectionPolicy() []string {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasClientConnectionPolicy() bool {
	if o != nil && !isNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *MirrorVirtualAttributeResponse) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetRequireExplicitRequestByName() bool {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasRequireExplicitRequestByName() bool {
	if o != nil && !isNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *MirrorVirtualAttributeResponse) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndex() int32 {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int32
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int32 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *MirrorVirtualAttributeResponse) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *MirrorVirtualAttributeResponse) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *MirrorVirtualAttributeResponse) GetAllowIndexConflicts() bool {
	if o == nil || isNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MirrorVirtualAttributeResponse) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowIndexConflicts) {
		return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *MirrorVirtualAttributeResponse) HasAllowIndexConflicts() bool {
	if o != nil && !isNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *MirrorVirtualAttributeResponse) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

func (o MirrorVirtualAttributeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.ConflictBehavior) {
		toSerialize["conflictBehavior"] = o.ConflictBehavior
	}
	if true {
		toSerialize["sourceAttribute"] = o.SourceAttribute
	}
	if !isNil(o.SourceEntryDNAttribute) {
		toSerialize["sourceEntryDNAttribute"] = o.SourceEntryDNAttribute
	}
	if !isNil(o.SourceEntryDNMap) {
		toSerialize["sourceEntryDNMap"] = o.SourceEntryDNMap
	}
	if !isNil(o.BypassAccessControlForSearches) {
		toSerialize["bypassAccessControlForSearches"] = o.BypassAccessControlForSearches
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

type NullableMirrorVirtualAttributeResponse struct {
	value *MirrorVirtualAttributeResponse
	isSet bool
}

func (v NullableMirrorVirtualAttributeResponse) Get() *MirrorVirtualAttributeResponse {
	return v.value
}

func (v *NullableMirrorVirtualAttributeResponse) Set(val *MirrorVirtualAttributeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMirrorVirtualAttributeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMirrorVirtualAttributeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMirrorVirtualAttributeResponse(val *MirrorVirtualAttributeResponse) *NullableMirrorVirtualAttributeResponse {
	return &NullableMirrorVirtualAttributeResponse{value: val, isSet: true}
}

func (v NullableMirrorVirtualAttributeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMirrorVirtualAttributeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
