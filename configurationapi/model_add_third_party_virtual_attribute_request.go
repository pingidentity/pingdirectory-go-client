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

// checks if the AddThirdPartyVirtualAttributeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddThirdPartyVirtualAttributeRequest{}

// AddThirdPartyVirtualAttributeRequest struct for AddThirdPartyVirtualAttributeRequest
type AddThirdPartyVirtualAttributeRequest struct {
	Schemas []EnumthirdPartyVirtualAttributeSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Virtual Attribute.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Virtual Attribute. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
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

// NewAddThirdPartyVirtualAttributeRequest instantiates a new AddThirdPartyVirtualAttributeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyVirtualAttributeRequest(schemas []EnumthirdPartyVirtualAttributeSchemaUrn, extensionClass string, enabled bool, attributeType string, name string) *AddThirdPartyVirtualAttributeRequest {
	this := AddThirdPartyVirtualAttributeRequest{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	this.AttributeType = attributeType
	this.Name = name
	return &this
}

// NewAddThirdPartyVirtualAttributeRequestWithDefaults instantiates a new AddThirdPartyVirtualAttributeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyVirtualAttributeRequestWithDefaults() *AddThirdPartyVirtualAttributeRequest {
	this := AddThirdPartyVirtualAttributeRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyVirtualAttributeRequest) GetSchemas() []EnumthirdPartyVirtualAttributeSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyVirtualAttributeSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetSchemasOk() ([]EnumthirdPartyVirtualAttributeSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyVirtualAttributeRequest) SetSchemas(v []EnumthirdPartyVirtualAttributeSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyVirtualAttributeRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyVirtualAttributeRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyVirtualAttributeRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyVirtualAttributeRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyVirtualAttributeRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyVirtualAttributeRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAttributeType returns the AttributeType field value
func (o *AddThirdPartyVirtualAttributeRequest) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *AddThirdPartyVirtualAttributeRequest) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetBaseDN() []string {
	if o == nil || IsNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *AddThirdPartyVirtualAttributeRequest) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetGroupDN returns the GroupDN field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetGroupDN() []string {
	if o == nil || IsNil(o.GroupDN) {
		var ret []string
		return ret
	}
	return o.GroupDN
}

// GetGroupDNOk returns a tuple with the GroupDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetGroupDNOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupDN) {
		return nil, false
	}
	return o.GroupDN, true
}

// HasGroupDN returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasGroupDN() bool {
	if o != nil && !IsNil(o.GroupDN) {
		return true
	}

	return false
}

// SetGroupDN gets a reference to the given []string and assigns it to the GroupDN field.
func (o *AddThirdPartyVirtualAttributeRequest) SetGroupDN(v []string) {
	o.GroupDN = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetFilter() []string {
	if o == nil || IsNil(o.Filter) {
		var ret []string
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []string and assigns it to the Filter field.
func (o *AddThirdPartyVirtualAttributeRequest) SetFilter(v []string) {
	o.Filter = v
}

// GetClientConnectionPolicy returns the ClientConnectionPolicy field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetClientConnectionPolicy() []string {
	if o == nil || IsNil(o.ClientConnectionPolicy) {
		var ret []string
		return ret
	}
	return o.ClientConnectionPolicy
}

// GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetClientConnectionPolicyOk() ([]string, bool) {
	if o == nil || IsNil(o.ClientConnectionPolicy) {
		return nil, false
	}
	return o.ClientConnectionPolicy, true
}

// HasClientConnectionPolicy returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasClientConnectionPolicy() bool {
	if o != nil && !IsNil(o.ClientConnectionPolicy) {
		return true
	}

	return false
}

// SetClientConnectionPolicy gets a reference to the given []string and assigns it to the ClientConnectionPolicy field.
func (o *AddThirdPartyVirtualAttributeRequest) SetClientConnectionPolicy(v []string) {
	o.ClientConnectionPolicy = v
}

// GetConflictBehavior returns the ConflictBehavior field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp {
	if o == nil || IsNil(o.ConflictBehavior) {
		var ret EnumvirtualAttributeConflictBehaviorProp
		return ret
	}
	return *o.ConflictBehavior
}

// GetConflictBehaviorOk returns a tuple with the ConflictBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool) {
	if o == nil || IsNil(o.ConflictBehavior) {
		return nil, false
	}
	return o.ConflictBehavior, true
}

// HasConflictBehavior returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasConflictBehavior() bool {
	if o != nil && !IsNil(o.ConflictBehavior) {
		return true
	}

	return false
}

// SetConflictBehavior gets a reference to the given EnumvirtualAttributeConflictBehaviorProp and assigns it to the ConflictBehavior field.
func (o *AddThirdPartyVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp) {
	o.ConflictBehavior = &v
}

// GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetRequireExplicitRequestByName() bool {
	if o == nil || IsNil(o.RequireExplicitRequestByName) {
		var ret bool
		return ret
	}
	return *o.RequireExplicitRequestByName
}

// GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireExplicitRequestByName) {
		return nil, false
	}
	return o.RequireExplicitRequestByName, true
}

// HasRequireExplicitRequestByName returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasRequireExplicitRequestByName() bool {
	if o != nil && !IsNil(o.RequireExplicitRequestByName) {
		return true
	}

	return false
}

// SetRequireExplicitRequestByName gets a reference to the given bool and assigns it to the RequireExplicitRequestByName field.
func (o *AddThirdPartyVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool) {
	o.RequireExplicitRequestByName = &v
}

// GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int64 {
	if o == nil || IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		var ret int64
		return ret
	}
	return *o.MultipleVirtualAttributeEvaluationOrderIndex
}

// GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return nil, false
	}
	return o.MultipleVirtualAttributeEvaluationOrderIndex, true
}

// HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool {
	if o != nil && !IsNil(o.MultipleVirtualAttributeEvaluationOrderIndex) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeEvaluationOrderIndex gets a reference to the given int64 and assigns it to the MultipleVirtualAttributeEvaluationOrderIndex field.
func (o *AddThirdPartyVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int64) {
	o.MultipleVirtualAttributeEvaluationOrderIndex = &v
}

// GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp {
	if o == nil || IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		var ret EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp
		return ret
	}
	return *o.MultipleVirtualAttributeMergeBehavior
}

// GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool) {
	if o == nil || IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		return nil, false
	}
	return o.MultipleVirtualAttributeMergeBehavior, true
}

// HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool {
	if o != nil && !IsNil(o.MultipleVirtualAttributeMergeBehavior) {
		return true
	}

	return false
}

// SetMultipleVirtualAttributeMergeBehavior gets a reference to the given EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp and assigns it to the MultipleVirtualAttributeMergeBehavior field.
func (o *AddThirdPartyVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp) {
	o.MultipleVirtualAttributeMergeBehavior = &v
}

// GetAllowIndexConflicts returns the AllowIndexConflicts field value if set, zero value otherwise.
func (o *AddThirdPartyVirtualAttributeRequest) GetAllowIndexConflicts() bool {
	if o == nil || IsNil(o.AllowIndexConflicts) {
		var ret bool
		return ret
	}
	return *o.AllowIndexConflicts
}

// GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowIndexConflicts) {
		return nil, false
	}
	return o.AllowIndexConflicts, true
}

// HasAllowIndexConflicts returns a boolean if a field has been set.
func (o *AddThirdPartyVirtualAttributeRequest) HasAllowIndexConflicts() bool {
	if o != nil && !IsNil(o.AllowIndexConflicts) {
		return true
	}

	return false
}

// SetAllowIndexConflicts gets a reference to the given bool and assigns it to the AllowIndexConflicts field.
func (o *AddThirdPartyVirtualAttributeRequest) SetAllowIndexConflicts(v bool) {
	o.AllowIndexConflicts = &v
}

// GetName returns the Name field value
func (o *AddThirdPartyVirtualAttributeRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyVirtualAttributeRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddThirdPartyVirtualAttributeRequest) SetName(v string) {
	o.Name = v
}

func (o AddThirdPartyVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddThirdPartyVirtualAttributeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
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

type NullableAddThirdPartyVirtualAttributeRequest struct {
	value *AddThirdPartyVirtualAttributeRequest
	isSet bool
}

func (v NullableAddThirdPartyVirtualAttributeRequest) Get() *AddThirdPartyVirtualAttributeRequest {
	return v.value
}

func (v *NullableAddThirdPartyVirtualAttributeRequest) Set(val *AddThirdPartyVirtualAttributeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyVirtualAttributeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyVirtualAttributeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyVirtualAttributeRequest(val *AddThirdPartyVirtualAttributeRequest) *NullableAddThirdPartyVirtualAttributeRequest {
	return &NullableAddThirdPartyVirtualAttributeRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyVirtualAttributeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyVirtualAttributeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
