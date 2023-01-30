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

// DnJoinVirtualAttributeResponse struct for DnJoinVirtualAttributeResponse
type DnJoinVirtualAttributeResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Virtual Attribute
	Id      string                                `json:"id"`
	Schemas []EnumdnJoinVirtualAttributeSchemaUrn `json:"schemas"`
	// The attribute whose values are the DNs of the entries to be joined with the search result entry.
	JoinDNAttribute string                                 `json:"joinDNAttribute"`
	JoinBaseDNType  EnumvirtualAttributeJoinBaseDNTypeProp `json:"joinBaseDNType"`
	// The fixed, administrator-specified base DN for the internal searches used to identify joined entries.
	JoinCustomBaseDN *string                            `json:"joinCustomBaseDN,omitempty"`
	JoinScope        *EnumvirtualAttributeJoinScopeProp `json:"joinScope,omitempty"`
	// The maximum number of entries that may be joined with the source entry, which also corresponds to the maximum number of values that the virtual attribute provider will generate for an entry.
	JoinSizeLimit *int32 `json:"joinSizeLimit,omitempty"`
	// An optional filter that specifies additional criteria for identifying joined entries. If a join-filter value is specified, then only entries matching that filter (in addition to satisfying the other join criteria) will be joined with the search result entry.
	JoinFilter *string `json:"joinFilter,omitempty"`
	// An optional set of the names of the attributes to include with joined entries.
	JoinAttribute []string `json:"joinAttribute,omitempty"`
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
	MultipleVirtualAttributeEvaluationOrderIndex *int32                                                         `json:"multipleVirtualAttributeEvaluationOrderIndex,omitempty"`
	MultipleVirtualAttributeMergeBehavior        *EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp `json:"multipleVirtualAttributeMergeBehavior,omitempty"`
	// Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server.
	AllowIndexConflicts *bool `json:"allowIndexConflicts,omitempty"`
}

// NewDnJoinVirtualAttributeResponse instantiates a new DnJoinVirtualAttributeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnJoinVirtualAttributeResponse(id string, schemas []EnumdnJoinVirtualAttributeSchemaUrn, joinDNAttribute string, joinBaseDNType EnumvirtualAttributeJoinBaseDNTypeProp, enabled bool, attributeType string) *DnJoinVirtualAttributeResponse {
	this := DnJoinVirtualAttributeResponse{}
	this.Id = id
	this.Schemas = schemas
	this.JoinDNAttribute = joinDNAttribute
	this.JoinBaseDNType = joinBaseDNType
	this.Enabled = enabled
	this.AttributeType = attributeType
	return &this
}

// NewDnJoinVirtualAttributeResponseWithDefaults instantiates a new DnJoinVirtualAttributeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnJoinVirtualAttributeResponseWithDefaults() *DnJoinVirtualAttributeResponse {
	this := DnJoinVirtualAttributeResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *DnJoinVirtualAttributeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *DnJoinVirtualAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *DnJoinVirtualAttributeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DnJoinVirtualAttributeResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *DnJoinVirtualAttributeResponse) GetSchemas() []EnumdnJoinVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnumdnJoinVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetSchemasOk() ([]EnumdnJoinVirtualAttributeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *DnJoinVirtualAttributeResponse) SetSchemas(v []EnumdnJoinVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetJoinDNAttribute returns the JoinDNAttribute field value
func (o *DnJoinVirtualAttributeResponse) GetJoinDNAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinDNAttribute
}

// GetJoinDNAttributeOk returns a tuple with the JoinDNAttribute field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetJoinDNAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinDNAttribute, true
}

// SetJoinDNAttribute sets field value
func (o *DnJoinVirtualAttributeResponse) SetJoinDNAttribute(v string) {
	o.JoinDNAttribute = v
}

// GetJoinBaseDNType returns the JoinBaseDNType field value
func (o *DnJoinVirtualAttributeResponse) GetJoinBaseDNType() EnumvirtualAttributeJoinBaseDNTypeProp {
	if o == nil {
		var ret EnumvirtualAttributeJoinBaseDNTypeProp
		return ret
	}

	return o.JoinBaseDNType
}

// GetJoinBaseDNTypeOk returns a tuple with the JoinBaseDNType field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetJoinBaseDNTypeOk() (*EnumvirtualAttributeJoinBaseDNTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinBaseDNType, true
}

// SetJoinBaseDNType sets field value
func (o *DnJoinVirtualAttributeResponse) SetJoinBaseDNType(v EnumvirtualAttributeJoinBaseDNTypeProp) {
	o.JoinBaseDNType = v
}

// GetJoinCustomBaseDN returns the JoinCustomBaseDN field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetJoinCustomBaseDN() string {
	if o == nil || isNil(o.JoinCustomBaseDN) {
		var ret string
		return ret
	}
	return *o.JoinCustomBaseDN
}

// GetJoinCustomBaseDNOk returns a tuple with the JoinCustomBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetJoinCustomBaseDNOk() (*string, bool) {
	if o == nil || isNil(o.JoinCustomBaseDN) {
		return nil, false
	}
	return o.JoinCustomBaseDN, true
}

// HasJoinCustomBaseDN returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasJoinCustomBaseDN() bool {
	if o != nil && !isNil(o.JoinCustomBaseDN) {
		return true
	}

	return false
}

// SetJoinCustomBaseDN gets a reference to the given string and assigns it to the JoinCustomBaseDN field.
func (o *DnJoinVirtualAttributeResponse) SetJoinCustomBaseDN(v string) {
	o.JoinCustomBaseDN = &v
}

// GetJoinScope returns the JoinScope field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetJoinScope() EnumvirtualAttributeJoinScopeProp {
	if o == nil || isNil(o.JoinScope) {
		var ret EnumvirtualAttributeJoinScopeProp
		return ret
	}
	return *o.JoinScope
}

// GetJoinScopeOk returns a tuple with the JoinScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetJoinScopeOk() (*EnumvirtualAttributeJoinScopeProp, bool) {
	if o == nil || isNil(o.JoinScope) {
		return nil, false
	}
	return o.JoinScope, true
}

// HasJoinScope returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasJoinScope() bool {
	if o != nil && !isNil(o.JoinScope) {
		return true
	}

	return false
}

// SetJoinScope gets a reference to the given EnumvirtualAttributeJoinScopeProp and assigns it to the JoinScope field.
func (o *DnJoinVirtualAttributeResponse) SetJoinScope(v EnumvirtualAttributeJoinScopeProp) {
	o.JoinScope = &v
}

// GetJoinSizeLimit returns the JoinSizeLimit field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetJoinSizeLimit() int32 {
	if o == nil || isNil(o.JoinSizeLimit) {
		var ret int32
		return ret
	}
	return *o.JoinSizeLimit
}

// GetJoinSizeLimitOk returns a tuple with the JoinSizeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetJoinSizeLimitOk() (*int32, bool) {
	if o == nil || isNil(o.JoinSizeLimit) {
		return nil, false
	}
	return o.JoinSizeLimit, true
}

// HasJoinSizeLimit returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasJoinSizeLimit() bool {
	if o != nil && !isNil(o.JoinSizeLimit) {
		return true
	}

	return false
}

// SetJoinSizeLimit gets a reference to the given int32 and assigns it to the JoinSizeLimit field.
func (o *DnJoinVirtualAttributeResponse) SetJoinSizeLimit(v int32) {
	o.JoinSizeLimit = &v
}

// GetJoinFilter returns the JoinFilter field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetJoinFilter() string {
	if o == nil || isNil(o.JoinFilter) {
		var ret string
		return ret
	}
	return *o.JoinFilter
}

// GetJoinFilterOk returns a tuple with the JoinFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetJoinFilterOk() (*string, bool) {
	if o == nil || isNil(o.JoinFilter) {
		return nil, false
	}
	return o.JoinFilter, true
}

// HasJoinFilter returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasJoinFilter() bool {
	if o != nil && !isNil(o.JoinFilter) {
		return true
	}

	return false
}

// SetJoinFilter gets a reference to the given string and assigns it to the JoinFilter field.
func (o *DnJoinVirtualAttributeResponse) SetJoinFilter(v string) {
	o.JoinFilter = &v
}

// GetJoinAttribute returns the JoinAttribute field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetJoinAttribute() []string {
	if o == nil || isNil(o.JoinAttribute) {
		var ret []string
		return ret
	}
	return o.JoinAttribute
}

// GetJoinAttributeOk returns a tuple with the JoinAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetJoinAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.JoinAttribute) {
		return nil, false
	}
	return o.JoinAttribute, true
}

// HasJoinAttribute returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasJoinAttribute() bool {
	if o != nil && !isNil(o.JoinAttribute) {
		return true
	}

	return false
}

// SetJoinAttribute gets a reference to the given []string and assigns it to the JoinAttribute field.
func (o *DnJoinVirtualAttributeResponse) SetJoinAttribute(v []string) {
	o.JoinAttribute = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DnJoinVirtualAttributeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *DnJoinVirtualAttributeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DnJoinVirtualAttributeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAttributeType returns the AttributeType field value
func (o *DnJoinVirtualAttributeResponse) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *DnJoinVirtualAttributeResponse) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetBaseDN() []string {
	if o == nil || isNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *DnJoinVirtualAttributeResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetGroupDN() []string {
	if o == nil || isNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetGroupDNOk() ([]string, bool) {
	if o == nil || isNil(o.GroupDN) {
		return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasGroupDN() bool {
	if o != nil && !isNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *DnJoinVirtualAttributeResponse) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetFilter() []string {
	if o == nil || isNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetFilterOk() ([]string, bool) {
	if o == nil || isNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasFilter() bool {
	if o != nil && !isNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *DnJoinVirtualAttributeResponse) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetClientConnectionPolicy() []string {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || isNil(o.ClientConnectionPolicy) {
		return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasClientConnectionPolicy() bool {
	if o != nil && !isNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *DnJoinVirtualAttributeResponse) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || isNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || isNil(o.ConflictBehavior) {
		return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasConflictBehavior() bool {
	if o != nil && !isNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *DnJoinVirtualAttributeResponse) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetRequireExplicitRequestByName() bool {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || isNil(o.RequireExplicitRequestByName) {
		return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasRequireExplicitRequestByName() bool {
	if o != nil && !isNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *DnJoinVirtualAttributeResponse) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndex() int32 {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int32
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int32 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *DnJoinVirtualAttributeResponse) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !isNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *DnJoinVirtualAttributeResponse) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *DnJoinVirtualAttributeResponse) GetAllowIndexConflicts() bool {
	if o == nil || isNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnJoinVirtualAttributeResponse) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || isNil(o.AllowIndexConflicts) {
		return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *DnJoinVirtualAttributeResponse) HasAllowIndexConflicts() bool {
	if o != nil && !isNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *DnJoinVirtualAttributeResponse) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

func (o DnJoinVirtualAttributeResponse) MarshalJSON() ([]byte, error) {
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

type NullableDnJoinVirtualAttributeResponse struct {
	value *DnJoinVirtualAttributeResponse
	isSet bool
}

func (v NullableDnJoinVirtualAttributeResponse) Get() *DnJoinVirtualAttributeResponse {
	return v.value
}

func (v *NullableDnJoinVirtualAttributeResponse) Set(val *DnJoinVirtualAttributeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDnJoinVirtualAttributeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDnJoinVirtualAttributeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnJoinVirtualAttributeResponse(val *DnJoinVirtualAttributeResponse) *NullableDnJoinVirtualAttributeResponse {
	return &NullableDnJoinVirtualAttributeResponse{value: val, isSet: true}
}

func (v NullableDnJoinVirtualAttributeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnJoinVirtualAttributeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
