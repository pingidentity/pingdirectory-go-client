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

// checks if the ComposedAttributePluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComposedAttributePluginResponse{}

// ComposedAttributePluginResponse struct for ComposedAttributePluginResponse
type ComposedAttributePluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id         string                                 `json:"id"`
	Schemas    []EnumcomposedAttributePluginSchemaUrn `json:"schemas"`
	PluginType []EnumpluginPluginTypeProp             `json:"pluginType"`
	// The name or OID of the attribute type for which values are to be generated.
	AttributeType string `json:"attributeType"`
	// Specifies a pattern for constructing the values to use for the target attribute type.
	ValuePattern                                         []string                                                            `json:"valuePattern"`
	MultipleValuePatternBehavior                         *EnumpluginMultipleValuePatternBehaviorProp                         `json:"multipleValuePatternBehavior,omitempty"`
	MultiValuedAttributeBehavior                         *EnumpluginMultiValuedAttributeBehaviorProp                         `json:"multiValuedAttributeBehavior,omitempty"`
	TargetAttributeExistsDuringInitialPopulationBehavior *EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp `json:"targetAttributeExistsDuringInitialPopulationBehavior,omitempty"`
	UpdateSourceAttributeBehavior                        *EnumpluginUpdateSourceAttributeBehaviorProp                        `json:"updateSourceAttributeBehavior,omitempty"`
	SourceAttributeRemovalBehavior                       *EnumpluginSourceAttributeRemovalBehaviorProp                       `json:"sourceAttributeRemovalBehavior,omitempty"`
	UpdateTargetAttributeBehavior                        *EnumpluginUpdateTargetAttributeBehaviorProp                        `json:"updateTargetAttributeBehavior,omitempty"`
	// The set of base DNs below which composed values may be generated.
	IncludeBaseDN []string `json:"includeBaseDN,omitempty"`
	// The set of base DNs below which composed values will not be generated.
	ExcludeBaseDN []string `json:"excludeBaseDN,omitempty"`
	// The set of search filters that identify entries for which composed values may be generated.
	IncludeFilter []string `json:"includeFilter,omitempty"`
	// The set of search filters that identify entries for which composed values will not be generated.
	ExcludeFilter                               []string                                                   `json:"excludeFilter,omitempty"`
	UpdatedEntryNewlyMatchesCriteriaBehavior    *EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp    `json:"updatedEntryNewlyMatchesCriteriaBehavior,omitempty"`
	UpdatedEntryNoLongerMatchesCriteriaBehavior *EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp `json:"updatedEntryNoLongerMatchesCriteriaBehavior,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
}

// NewComposedAttributePluginResponse instantiates a new ComposedAttributePluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComposedAttributePluginResponse(id string, schemas []EnumcomposedAttributePluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, attributeType string, valuePattern []string, enabled bool) *ComposedAttributePluginResponse {
	this := ComposedAttributePluginResponse{}
	this.Id = id
	this.Schemas = schemas
	this.PluginType = pluginType
	this.AttributeType = attributeType
	this.ValuePattern = valuePattern
	this.Enabled = enabled
	return &this
}

// NewComposedAttributePluginResponseWithDefaults instantiates a new ComposedAttributePluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComposedAttributePluginResponseWithDefaults() *ComposedAttributePluginResponse {
	this := ComposedAttributePluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ComposedAttributePluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ComposedAttributePluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ComposedAttributePluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ComposedAttributePluginResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ComposedAttributePluginResponse) GetSchemas() []EnumcomposedAttributePluginSchemaUrn {
	if o == nil {
		var ret []EnumcomposedAttributePluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetSchemasOk() ([]EnumcomposedAttributePluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ComposedAttributePluginResponse) SetSchemas(v []EnumcomposedAttributePluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value
func (o *ComposedAttributePluginResponse) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *ComposedAttributePluginResponse) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetAttributeType returns the AttributeType field value
func (o *ComposedAttributePluginResponse) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *ComposedAttributePluginResponse) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetValuePattern returns the ValuePattern field value
func (o *ComposedAttributePluginResponse) GetValuePattern() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ValuePattern
}

// GetValuePatternOk returns a tuple with the ValuePattern field value
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetValuePatternOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValuePattern, true
}

// SetValuePattern sets field value
func (o *ComposedAttributePluginResponse) SetValuePattern(v []string) {
	o.ValuePattern = v
}

// GetMultipleValuePatternBehavior returns the MultipleValuePatternBehavior field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetMultipleValuePatternBehavior() EnumpluginMultipleValuePatternBehaviorProp {
	if o == nil || IsNil(o.MultipleValuePatternBehavior) {
		var ret EnumpluginMultipleValuePatternBehaviorProp
		return ret
	}
	return *o.MultipleValuePatternBehavior
}

// GetMultipleValuePatternBehaviorOk returns a tuple with the MultipleValuePatternBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetMultipleValuePatternBehaviorOk() (*EnumpluginMultipleValuePatternBehaviorProp, bool) {
	if o == nil || IsNil(o.MultipleValuePatternBehavior) {
		return nil, false
	}
	return o.MultipleValuePatternBehavior, true
}

// HasMultipleValuePatternBehavior returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasMultipleValuePatternBehavior() bool {
	if o != nil && !IsNil(o.MultipleValuePatternBehavior) {
		return true
	}

	return false
}

// SetMultipleValuePatternBehavior gets a reference to the given EnumpluginMultipleValuePatternBehaviorProp and assigns it to the MultipleValuePatternBehavior field.
func (o *ComposedAttributePluginResponse) SetMultipleValuePatternBehavior(v EnumpluginMultipleValuePatternBehaviorProp) {
	o.MultipleValuePatternBehavior = &v
}

// GetMultiValuedAttributeBehavior returns the MultiValuedAttributeBehavior field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetMultiValuedAttributeBehavior() EnumpluginMultiValuedAttributeBehaviorProp {
	if o == nil || IsNil(o.MultiValuedAttributeBehavior) {
		var ret EnumpluginMultiValuedAttributeBehaviorProp
		return ret
	}
	return *o.MultiValuedAttributeBehavior
}

// GetMultiValuedAttributeBehaviorOk returns a tuple with the MultiValuedAttributeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetMultiValuedAttributeBehaviorOk() (*EnumpluginMultiValuedAttributeBehaviorProp, bool) {
	if o == nil || IsNil(o.MultiValuedAttributeBehavior) {
		return nil, false
	}
	return o.MultiValuedAttributeBehavior, true
}

// HasMultiValuedAttributeBehavior returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasMultiValuedAttributeBehavior() bool {
	if o != nil && !IsNil(o.MultiValuedAttributeBehavior) {
		return true
	}

	return false
}

// SetMultiValuedAttributeBehavior gets a reference to the given EnumpluginMultiValuedAttributeBehaviorProp and assigns it to the MultiValuedAttributeBehavior field.
func (o *ComposedAttributePluginResponse) SetMultiValuedAttributeBehavior(v EnumpluginMultiValuedAttributeBehaviorProp) {
	o.MultiValuedAttributeBehavior = &v
}

// GetTargetAttributeExistsDuringInitialPopulationBehavior returns the TargetAttributeExistsDuringInitialPopulationBehavior field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetTargetAttributeExistsDuringInitialPopulationBehavior() EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp {
	if o == nil || IsNil(o.TargetAttributeExistsDuringInitialPopulationBehavior) {
		var ret EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp
		return ret
	}
	return *o.TargetAttributeExistsDuringInitialPopulationBehavior
}

// GetTargetAttributeExistsDuringInitialPopulationBehaviorOk returns a tuple with the TargetAttributeExistsDuringInitialPopulationBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetTargetAttributeExistsDuringInitialPopulationBehaviorOk() (*EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp, bool) {
	if o == nil || IsNil(o.TargetAttributeExistsDuringInitialPopulationBehavior) {
		return nil, false
	}
	return o.TargetAttributeExistsDuringInitialPopulationBehavior, true
}

// HasTargetAttributeExistsDuringInitialPopulationBehavior returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasTargetAttributeExistsDuringInitialPopulationBehavior() bool {
	if o != nil && !IsNil(o.TargetAttributeExistsDuringInitialPopulationBehavior) {
		return true
	}

	return false
}

// SetTargetAttributeExistsDuringInitialPopulationBehavior gets a reference to the given EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp and assigns it to the TargetAttributeExistsDuringInitialPopulationBehavior field.
func (o *ComposedAttributePluginResponse) SetTargetAttributeExistsDuringInitialPopulationBehavior(v EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp) {
	o.TargetAttributeExistsDuringInitialPopulationBehavior = &v
}

// GetUpdateSourceAttributeBehavior returns the UpdateSourceAttributeBehavior field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetUpdateSourceAttributeBehavior() EnumpluginUpdateSourceAttributeBehaviorProp {
	if o == nil || IsNil(o.UpdateSourceAttributeBehavior) {
		var ret EnumpluginUpdateSourceAttributeBehaviorProp
		return ret
	}
	return *o.UpdateSourceAttributeBehavior
}

// GetUpdateSourceAttributeBehaviorOk returns a tuple with the UpdateSourceAttributeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetUpdateSourceAttributeBehaviorOk() (*EnumpluginUpdateSourceAttributeBehaviorProp, bool) {
	if o == nil || IsNil(o.UpdateSourceAttributeBehavior) {
		return nil, false
	}
	return o.UpdateSourceAttributeBehavior, true
}

// HasUpdateSourceAttributeBehavior returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasUpdateSourceAttributeBehavior() bool {
	if o != nil && !IsNil(o.UpdateSourceAttributeBehavior) {
		return true
	}

	return false
}

// SetUpdateSourceAttributeBehavior gets a reference to the given EnumpluginUpdateSourceAttributeBehaviorProp and assigns it to the UpdateSourceAttributeBehavior field.
func (o *ComposedAttributePluginResponse) SetUpdateSourceAttributeBehavior(v EnumpluginUpdateSourceAttributeBehaviorProp) {
	o.UpdateSourceAttributeBehavior = &v
}

// GetSourceAttributeRemovalBehavior returns the SourceAttributeRemovalBehavior field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetSourceAttributeRemovalBehavior() EnumpluginSourceAttributeRemovalBehaviorProp {
	if o == nil || IsNil(o.SourceAttributeRemovalBehavior) {
		var ret EnumpluginSourceAttributeRemovalBehaviorProp
		return ret
	}
	return *o.SourceAttributeRemovalBehavior
}

// GetSourceAttributeRemovalBehaviorOk returns a tuple with the SourceAttributeRemovalBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetSourceAttributeRemovalBehaviorOk() (*EnumpluginSourceAttributeRemovalBehaviorProp, bool) {
	if o == nil || IsNil(o.SourceAttributeRemovalBehavior) {
		return nil, false
	}
	return o.SourceAttributeRemovalBehavior, true
}

// HasSourceAttributeRemovalBehavior returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasSourceAttributeRemovalBehavior() bool {
	if o != nil && !IsNil(o.SourceAttributeRemovalBehavior) {
		return true
	}

	return false
}

// SetSourceAttributeRemovalBehavior gets a reference to the given EnumpluginSourceAttributeRemovalBehaviorProp and assigns it to the SourceAttributeRemovalBehavior field.
func (o *ComposedAttributePluginResponse) SetSourceAttributeRemovalBehavior(v EnumpluginSourceAttributeRemovalBehaviorProp) {
	o.SourceAttributeRemovalBehavior = &v
}

// GetUpdateTargetAttributeBehavior returns the UpdateTargetAttributeBehavior field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetUpdateTargetAttributeBehavior() EnumpluginUpdateTargetAttributeBehaviorProp {
	if o == nil || IsNil(o.UpdateTargetAttributeBehavior) {
		var ret EnumpluginUpdateTargetAttributeBehaviorProp
		return ret
	}
	return *o.UpdateTargetAttributeBehavior
}

// GetUpdateTargetAttributeBehaviorOk returns a tuple with the UpdateTargetAttributeBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetUpdateTargetAttributeBehaviorOk() (*EnumpluginUpdateTargetAttributeBehaviorProp, bool) {
	if o == nil || IsNil(o.UpdateTargetAttributeBehavior) {
		return nil, false
	}
	return o.UpdateTargetAttributeBehavior, true
}

// HasUpdateTargetAttributeBehavior returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasUpdateTargetAttributeBehavior() bool {
	if o != nil && !IsNil(o.UpdateTargetAttributeBehavior) {
		return true
	}

	return false
}

// SetUpdateTargetAttributeBehavior gets a reference to the given EnumpluginUpdateTargetAttributeBehaviorProp and assigns it to the UpdateTargetAttributeBehavior field.
func (o *ComposedAttributePluginResponse) SetUpdateTargetAttributeBehavior(v EnumpluginUpdateTargetAttributeBehaviorProp) {
	o.UpdateTargetAttributeBehavior = &v
}

// GetIncludeBaseDN returns the IncludeBaseDN field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetIncludeBaseDN() []string {
	if o == nil || IsNil(o.IncludeBaseDN) {
		var ret []string
		return ret
	}
	return o.IncludeBaseDN
}

// GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetIncludeBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeBaseDN) {
		return nil, false
	}
	return o.IncludeBaseDN, true
}

// HasIncludeBaseDN returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasIncludeBaseDN() bool {
	if o != nil && !IsNil(o.IncludeBaseDN) {
		return true
	}

	return false
}

// SetIncludeBaseDN gets a reference to the given []string and assigns it to the IncludeBaseDN field.
func (o *ComposedAttributePluginResponse) SetIncludeBaseDN(v []string) {
	o.IncludeBaseDN = v
}

// GetExcludeBaseDN returns the ExcludeBaseDN field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetExcludeBaseDN() []string {
	if o == nil || IsNil(o.ExcludeBaseDN) {
		var ret []string
		return ret
	}
	return o.ExcludeBaseDN
}

// GetExcludeBaseDNOk returns a tuple with the ExcludeBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetExcludeBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeBaseDN) {
		return nil, false
	}
	return o.ExcludeBaseDN, true
}

// HasExcludeBaseDN returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasExcludeBaseDN() bool {
	if o != nil && !IsNil(o.ExcludeBaseDN) {
		return true
	}

	return false
}

// SetExcludeBaseDN gets a reference to the given []string and assigns it to the ExcludeBaseDN field.
func (o *ComposedAttributePluginResponse) SetExcludeBaseDN(v []string) {
	o.ExcludeBaseDN = v
}

// GetIncludeFilter returns the IncludeFilter field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetIncludeFilter() []string {
	if o == nil || IsNil(o.IncludeFilter) {
		var ret []string
		return ret
	}
	return o.IncludeFilter
}

// GetIncludeFilterOk returns a tuple with the IncludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetIncludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeFilter) {
		return nil, false
	}
	return o.IncludeFilter, true
}

// HasIncludeFilter returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasIncludeFilter() bool {
	if o != nil && !IsNil(o.IncludeFilter) {
		return true
	}

	return false
}

// SetIncludeFilter gets a reference to the given []string and assigns it to the IncludeFilter field.
func (o *ComposedAttributePluginResponse) SetIncludeFilter(v []string) {
	o.IncludeFilter = v
}

// GetExcludeFilter returns the ExcludeFilter field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetExcludeFilter() []string {
	if o == nil || IsNil(o.ExcludeFilter) {
		var ret []string
		return ret
	}
	return o.ExcludeFilter
}

// GetExcludeFilterOk returns a tuple with the ExcludeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetExcludeFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeFilter) {
		return nil, false
	}
	return o.ExcludeFilter, true
}

// HasExcludeFilter returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasExcludeFilter() bool {
	if o != nil && !IsNil(o.ExcludeFilter) {
		return true
	}

	return false
}

// SetExcludeFilter gets a reference to the given []string and assigns it to the ExcludeFilter field.
func (o *ComposedAttributePluginResponse) SetExcludeFilter(v []string) {
	o.ExcludeFilter = v
}

// GetUpdatedEntryNewlyMatchesCriteriaBehavior returns the UpdatedEntryNewlyMatchesCriteriaBehavior field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetUpdatedEntryNewlyMatchesCriteriaBehavior() EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp {
	if o == nil || IsNil(o.UpdatedEntryNewlyMatchesCriteriaBehavior) {
		var ret EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp
		return ret
	}
	return *o.UpdatedEntryNewlyMatchesCriteriaBehavior
}

// GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNewlyMatchesCriteriaBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp, bool) {
	if o == nil || IsNil(o.UpdatedEntryNewlyMatchesCriteriaBehavior) {
		return nil, false
	}
	return o.UpdatedEntryNewlyMatchesCriteriaBehavior, true
}

// HasUpdatedEntryNewlyMatchesCriteriaBehavior returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasUpdatedEntryNewlyMatchesCriteriaBehavior() bool {
	if o != nil && !IsNil(o.UpdatedEntryNewlyMatchesCriteriaBehavior) {
		return true
	}

	return false
}

// SetUpdatedEntryNewlyMatchesCriteriaBehavior gets a reference to the given EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp and assigns it to the UpdatedEntryNewlyMatchesCriteriaBehavior field.
func (o *ComposedAttributePluginResponse) SetUpdatedEntryNewlyMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp) {
	o.UpdatedEntryNewlyMatchesCriteriaBehavior = &v
}

// GetUpdatedEntryNoLongerMatchesCriteriaBehavior returns the UpdatedEntryNoLongerMatchesCriteriaBehavior field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetUpdatedEntryNoLongerMatchesCriteriaBehavior() EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp {
	if o == nil || IsNil(o.UpdatedEntryNoLongerMatchesCriteriaBehavior) {
		var ret EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp
		return ret
	}
	return *o.UpdatedEntryNoLongerMatchesCriteriaBehavior
}

// GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNoLongerMatchesCriteriaBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp, bool) {
	if o == nil || IsNil(o.UpdatedEntryNoLongerMatchesCriteriaBehavior) {
		return nil, false
	}
	return o.UpdatedEntryNoLongerMatchesCriteriaBehavior, true
}

// HasUpdatedEntryNoLongerMatchesCriteriaBehavior returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasUpdatedEntryNoLongerMatchesCriteriaBehavior() bool {
	if o != nil && !IsNil(o.UpdatedEntryNoLongerMatchesCriteriaBehavior) {
		return true
	}

	return false
}

// SetUpdatedEntryNoLongerMatchesCriteriaBehavior gets a reference to the given EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp and assigns it to the UpdatedEntryNoLongerMatchesCriteriaBehavior field.
func (o *ComposedAttributePluginResponse) SetUpdatedEntryNoLongerMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp) {
	o.UpdatedEntryNoLongerMatchesCriteriaBehavior = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ComposedAttributePluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ComposedAttributePluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ComposedAttributePluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *ComposedAttributePluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComposedAttributePluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *ComposedAttributePluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *ComposedAttributePluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o ComposedAttributePluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComposedAttributePluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["pluginType"] = o.PluginType
	toSerialize["attributeType"] = o.AttributeType
	toSerialize["valuePattern"] = o.ValuePattern
	if !IsNil(o.MultipleValuePatternBehavior) {
		toSerialize["multipleValuePatternBehavior"] = o.MultipleValuePatternBehavior
	}
	if !IsNil(o.MultiValuedAttributeBehavior) {
		toSerialize["multiValuedAttributeBehavior"] = o.MultiValuedAttributeBehavior
	}
	if !IsNil(o.TargetAttributeExistsDuringInitialPopulationBehavior) {
		toSerialize["targetAttributeExistsDuringInitialPopulationBehavior"] = o.TargetAttributeExistsDuringInitialPopulationBehavior
	}
	if !IsNil(o.UpdateSourceAttributeBehavior) {
		toSerialize["updateSourceAttributeBehavior"] = o.UpdateSourceAttributeBehavior
	}
	if !IsNil(o.SourceAttributeRemovalBehavior) {
		toSerialize["sourceAttributeRemovalBehavior"] = o.SourceAttributeRemovalBehavior
	}
	if !IsNil(o.UpdateTargetAttributeBehavior) {
		toSerialize["updateTargetAttributeBehavior"] = o.UpdateTargetAttributeBehavior
	}
	if !IsNil(o.IncludeBaseDN) {
		toSerialize["includeBaseDN"] = o.IncludeBaseDN
	}
	if !IsNil(o.ExcludeBaseDN) {
		toSerialize["excludeBaseDN"] = o.ExcludeBaseDN
	}
	if !IsNil(o.IncludeFilter) {
		toSerialize["includeFilter"] = o.IncludeFilter
	}
	if !IsNil(o.ExcludeFilter) {
		toSerialize["excludeFilter"] = o.ExcludeFilter
	}
	if !IsNil(o.UpdatedEntryNewlyMatchesCriteriaBehavior) {
		toSerialize["updatedEntryNewlyMatchesCriteriaBehavior"] = o.UpdatedEntryNewlyMatchesCriteriaBehavior
	}
	if !IsNil(o.UpdatedEntryNoLongerMatchesCriteriaBehavior) {
		toSerialize["updatedEntryNoLongerMatchesCriteriaBehavior"] = o.UpdatedEntryNoLongerMatchesCriteriaBehavior
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	return toSerialize, nil
}

type NullableComposedAttributePluginResponse struct {
	value *ComposedAttributePluginResponse
	isSet bool
}

func (v NullableComposedAttributePluginResponse) Get() *ComposedAttributePluginResponse {
	return v.value
}

func (v *NullableComposedAttributePluginResponse) Set(val *ComposedAttributePluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableComposedAttributePluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableComposedAttributePluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComposedAttributePluginResponse(val *ComposedAttributePluginResponse) *NullableComposedAttributePluginResponse {
	return &NullableComposedAttributePluginResponse{value: val, isSet: true}
}

func (v NullableComposedAttributePluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComposedAttributePluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
