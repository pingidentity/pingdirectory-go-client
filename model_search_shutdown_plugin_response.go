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

// SearchShutdownPluginResponse struct for SearchShutdownPluginResponse
type SearchShutdownPluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id      string                              `json:"id"`
	Schemas []EnumsearchShutdownPluginSchemaUrn `json:"schemas"`
	// The base DN to use for the search.
	BaseDN *string             `json:"baseDN,omitempty"`
	Scope  EnumpluginScopeProp `json:"scope"`
	// The filter to use for the search.
	Filter string `json:"filter"`
	// The name of an attribute that should be included in the results. This may include any token which is allowed as a requested attribute in search requests, including the name of an attribute, an asterisk (to indicate all user attributes), a plus sign (to indicate all operational attributes), an object class name preceded with an at symbol (to indicate all attributes associated with that object class), an attribute name preceded by a caret (to indicate that attribute should be excluded), or an object class name preceded by a caret and an at symbol (to indicate that all attributes associated with that object class should be excluded).
	IncludeAttribute []string `json:"includeAttribute,omitempty"`
	// The path of an LDIF file that should be created with the results of the search.
	OutputFile string `json:"outputFile"`
	// An extension that should be appended to the name of an existing output file rather than deleting it. If a file already exists with the full previous file name, then it will be deleted before the current file is renamed to become the previous file.
	PreviousFileExtension *string `json:"previousFileExtension,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewSearchShutdownPluginResponse instantiates a new SearchShutdownPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchShutdownPluginResponse(id string, schemas []EnumsearchShutdownPluginSchemaUrn, scope EnumpluginScopeProp, filter string, outputFile string, enabled bool) *SearchShutdownPluginResponse {
	this := SearchShutdownPluginResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Scope = scope
	this.Filter = filter
	this.OutputFile = outputFile
	this.Enabled = enabled
	return &this
}

// NewSearchShutdownPluginResponseWithDefaults instantiates a new SearchShutdownPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchShutdownPluginResponseWithDefaults() *SearchShutdownPluginResponse {
	this := SearchShutdownPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SearchShutdownPluginResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SearchShutdownPluginResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SearchShutdownPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SearchShutdownPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SearchShutdownPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SearchShutdownPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *SearchShutdownPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SearchShutdownPluginResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *SearchShutdownPluginResponse) GetSchemas() []EnumsearchShutdownPluginSchemaUrn {
	if o == nil {
		var ret []EnumsearchShutdownPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetSchemasOk() ([]EnumsearchShutdownPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SearchShutdownPluginResponse) SetSchemas(v []EnumsearchShutdownPluginSchemaUrn) {
	o.Schemas = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *SearchShutdownPluginResponse) GetBaseDN() string {
	if o == nil || isNil(o.BaseDN) {
		var ret string
		return ret
	}
	return *o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetBaseDNOk() (*string, bool) {
	if o == nil || isNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *SearchShutdownPluginResponse) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given string and assigns it to the BaseDN field.
func (o *SearchShutdownPluginResponse) SetBaseDN(v string) {
	o.BaseDN = &v
}

// GetScope returns the Scope field value
func (o *SearchShutdownPluginResponse) GetScope() EnumpluginScopeProp {
	if o == nil {
		var ret EnumpluginScopeProp
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetScopeOk() (*EnumpluginScopeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *SearchShutdownPluginResponse) SetScope(v EnumpluginScopeProp) {
	o.Scope = v
}

// GetFilter returns the Filter field value
func (o *SearchShutdownPluginResponse) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *SearchShutdownPluginResponse) SetFilter(v string) {
	o.Filter = v
}

// GetIncludeAttribute returns the IncludeAttribute field value if set, zero value otherwise.
func (o *SearchShutdownPluginResponse) GetIncludeAttribute() []string {
	if o == nil || isNil(o.IncludeAttribute) {
		var ret []string
		return ret
	}
	return o.IncludeAttribute
}

// GetIncludeAttributeOk returns a tuple with the IncludeAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetIncludeAttributeOk() ([]string, bool) {
	if o == nil || isNil(o.IncludeAttribute) {
		return nil, false
	}
	return o.IncludeAttribute, true
}

// HasIncludeAttribute returns a boolean if a field has been set.
func (o *SearchShutdownPluginResponse) HasIncludeAttribute() bool {
	if o != nil && !isNil(o.IncludeAttribute) {
		return true
	}

	return false
}

// SetIncludeAttribute gets a reference to the given []string and assigns it to the IncludeAttribute field.
func (o *SearchShutdownPluginResponse) SetIncludeAttribute(v []string) {
	o.IncludeAttribute = v
}

// GetOutputFile returns the OutputFile field value
func (o *SearchShutdownPluginResponse) GetOutputFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OutputFile
}

// GetOutputFileOk returns a tuple with the OutputFile field value
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetOutputFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputFile, true
}

// SetOutputFile sets field value
func (o *SearchShutdownPluginResponse) SetOutputFile(v string) {
	o.OutputFile = v
}

// GetPreviousFileExtension returns the PreviousFileExtension field value if set, zero value otherwise.
func (o *SearchShutdownPluginResponse) GetPreviousFileExtension() string {
	if o == nil || isNil(o.PreviousFileExtension) {
		var ret string
		return ret
	}
	return *o.PreviousFileExtension
}

// GetPreviousFileExtensionOk returns a tuple with the PreviousFileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetPreviousFileExtensionOk() (*string, bool) {
	if o == nil || isNil(o.PreviousFileExtension) {
		return nil, false
	}
	return o.PreviousFileExtension, true
}

// HasPreviousFileExtension returns a boolean if a field has been set.
func (o *SearchShutdownPluginResponse) HasPreviousFileExtension() bool {
	if o != nil && !isNil(o.PreviousFileExtension) {
		return true
	}

	return false
}

// SetPreviousFileExtension gets a reference to the given string and assigns it to the PreviousFileExtension field.
func (o *SearchShutdownPluginResponse) SetPreviousFileExtension(v string) {
	o.PreviousFileExtension = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SearchShutdownPluginResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SearchShutdownPluginResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SearchShutdownPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *SearchShutdownPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SearchShutdownPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SearchShutdownPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o SearchShutdownPluginResponse) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["filter"] = o.Filter
	}
	if !isNil(o.IncludeAttribute) {
		toSerialize["includeAttribute"] = o.IncludeAttribute
	}
	if true {
		toSerialize["outputFile"] = o.OutputFile
	}
	if !isNil(o.PreviousFileExtension) {
		toSerialize["previousFileExtension"] = o.PreviousFileExtension
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableSearchShutdownPluginResponse struct {
	value *SearchShutdownPluginResponse
	isSet bool
}

func (v NullableSearchShutdownPluginResponse) Get() *SearchShutdownPluginResponse {
	return v.value
}

func (v *NullableSearchShutdownPluginResponse) Set(val *SearchShutdownPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchShutdownPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchShutdownPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchShutdownPluginResponse(val *SearchShutdownPluginResponse) *NullableSearchShutdownPluginResponse {
	return &NullableSearchShutdownPluginResponse{value: val, isSet: true}
}

func (v NullableSearchShutdownPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchShutdownPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
