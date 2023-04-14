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

// checks if the HasSubordinatesVirtualAttributeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasSubordinatesVirtualAttributeResponse{}

// HasSubordinatesVirtualAttributeResponse struct for HasSubordinatesVirtualAttributeResponse
type HasSubordinatesVirtualAttributeResponse struct {
	Schemas []EnumhasSubordinatesVirtualAttributeSchemaUrn `json:"schemas"`
	// Name of the Virtual Attribute
	Id               string                                    `json:"id"`
	ConflictBehavior *EnumvirtualAttributeConflictBehaviorProp `json:"conflictBehavior,omitempty"`
	// Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute.
	AttributeType string `json:"attributeType"`
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
	MultipleVirtualAttributeEvaluationOrderIndex *int64                                                         `json:"multipleVirtualAttributeEvaluationOrderIndex,omitempty"`
	MultipleVirtualAttributeMergeBehavior        *EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp `json:"multipleVirtualAttributeMergeBehavior,omitempty"`
	// Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server.
	AllowIndexConflicts                           *bool                                              `json:"allowIndexConflicts,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewHasSubordinatesVirtualAttributeResponse instantiates a new HasSubordinatesVirtualAttributeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasSubordinatesVirtualAttributeResponse(schemas []EnumhasSubordinatesVirtualAttributeSchemaUrn, id string, attributeType string, enabled bool) *HasSubordinatesVirtualAttributeResponse {
	this := HasSubordinatesVirtualAttributeResponse{}
	this.Schemas = schemas
	this.Id = id
	this.AttributeType = attributeType
	this.Enabled = enabled
	return &this
}

// NewHasSubordinatesVirtualAttributeResponseWithDefaults instantiates a new HasSubordinatesVirtualAttributeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasSubordinatesVirtualAttributeResponseWithDefaults() *HasSubordinatesVirtualAttributeResponse {
	this := HasSubordinatesVirtualAttributeResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *HasSubordinatesVirtualAttributeResponse) GetSchemas() []EnumhasSubordinatesVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnumhasSubordinatesVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetSchemasOk() ([]EnumhasSubordinatesVirtualAttributeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *HasSubordinatesVirtualAttributeResponse) SetSchemas(v []EnumhasSubordinatesVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *HasSubordinatesVirtualAttributeResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HasSubordinatesVirtualAttributeResponse) SetId(v string) {
	o.Id = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || IsNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || IsNil(o.ConflictBehavior) {
		return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasConflictBehavior() bool {
	if o != nil && !IsNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *HasSubordinatesVirtualAttributeResponse) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetAttributeType returns the AttributeType field value
func (o *HasSubordinatesVirtualAttributeResponse) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *HasSubordinatesVirtualAttributeResponse) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HasSubordinatesVirtualAttributeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *HasSubordinatesVirtualAttributeResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *HasSubordinatesVirtualAttributeResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetBaseDN() []string {
	if o == nil || IsNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *HasSubordinatesVirtualAttributeResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetGroupDN() []string {
	if o == nil || IsNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetGroupDNOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupDN) {
		return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasGroupDN() bool {
	if o != nil && !IsNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *HasSubordinatesVirtualAttributeResponse) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetFilter() []string {
	if o == nil || IsNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *HasSubordinatesVirtualAttributeResponse) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetClientConnectionPolicy() []string {
	if o == nil || IsNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || IsNil(o.ClientConnectionPolicy) {
		return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasClientConnectionPolicy() bool {
	if o != nil && !IsNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *HasSubordinatesVirtualAttributeResponse) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetRequireExplicitRequestByName() bool {
	if o == nil || IsNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireExplicitRequestByName) {
		return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasRequireExplicitRequestByName() bool {
	if o != nil && !IsNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *HasSubordinatesVirtualAttributeResponse) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndex() int64 {
	if o == nil || IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int64
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int64 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *HasSubordinatesVirtualAttributeResponse) SetMultipleVirtualAttributeEvaluationOrderIndex(v int64) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *HasSubordinatesVirtualAttributeResponse) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetAllowIndexConflicts() bool {
	if o == nil || IsNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowIndexConflicts) {
		return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasAllowIndexConflicts() bool {
	if o != nil && !IsNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *HasSubordinatesVirtualAttributeResponse) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *HasSubordinatesVirtualAttributeResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *HasSubordinatesVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasSubordinatesVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *HasSubordinatesVirtualAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *HasSubordinatesVirtualAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o HasSubordinatesVirtualAttributeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasSubordinatesVirtualAttributeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.ConflictBehavior) {
		toSerialize["conflictBehavior"] = o.ConflictBehavior
	}
	toSerialize["attributeType"] = o.AttributeType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
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
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableHasSubordinatesVirtualAttributeResponse struct {
	value *HasSubordinatesVirtualAttributeResponse
	isSet bool
}

func (v NullableHasSubordinatesVirtualAttributeResponse) Get() *HasSubordinatesVirtualAttributeResponse {
	return v.value
}

func (v *NullableHasSubordinatesVirtualAttributeResponse) Set(val *HasSubordinatesVirtualAttributeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHasSubordinatesVirtualAttributeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHasSubordinatesVirtualAttributeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasSubordinatesVirtualAttributeResponse(val *HasSubordinatesVirtualAttributeResponse) *NullableHasSubordinatesVirtualAttributeResponse {
	return &NullableHasSubordinatesVirtualAttributeResponse{value: val, isSet: true}
}

func (v NullableHasSubordinatesVirtualAttributeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasSubordinatesVirtualAttributeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
