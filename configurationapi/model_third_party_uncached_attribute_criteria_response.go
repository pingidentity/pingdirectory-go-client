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

// ThirdPartyUncachedAttributeCriteriaResponse struct for ThirdPartyUncachedAttributeCriteriaResponse
type ThirdPartyUncachedAttributeCriteriaResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Uncached Attribute Criteria
	Id      string                                             `json:"id"`
	Schemas []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Uncached Attribute Criteria.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Uncached Attribute Criteria. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Uncached Attribute Criteria
	Description *string `json:"description,omitempty"`
	// Indicates whether this Uncached Attribute Criteria is enabled for use in the server.
	Enabled bool `json:"enabled"`
}

// NewThirdPartyUncachedAttributeCriteriaResponse instantiates a new ThirdPartyUncachedAttributeCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyUncachedAttributeCriteriaResponse(id string, schemas []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn, extensionClass string, enabled bool) *ThirdPartyUncachedAttributeCriteriaResponse {
	this := ThirdPartyUncachedAttributeCriteriaResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyUncachedAttributeCriteriaResponseWithDefaults instantiates a new ThirdPartyUncachedAttributeCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyUncachedAttributeCriteriaResponseWithDefaults() *ThirdPartyUncachedAttributeCriteriaResponse {
	this := ThirdPartyUncachedAttributeCriteriaResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyUncachedAttributeCriteriaResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetSchemas() []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetSchemasOk() ([]EnumthirdPartyUncachedAttributeCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyUncachedAttributeCriteriaResponse) SetSchemas(v []EnumthirdPartyUncachedAttributeCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyUncachedAttributeCriteriaResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyUncachedAttributeCriteriaResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyUncachedAttributeCriteriaResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o ThirdPartyUncachedAttributeCriteriaResponse) MarshalJSON() ([]byte, error) {
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
		toSerialize["extensionClass"] = o.ExtensionClass
	}
	if !isNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableThirdPartyUncachedAttributeCriteriaResponse struct {
	value *ThirdPartyUncachedAttributeCriteriaResponse
	isSet bool
}

func (v NullableThirdPartyUncachedAttributeCriteriaResponse) Get() *ThirdPartyUncachedAttributeCriteriaResponse {
	return v.value
}

func (v *NullableThirdPartyUncachedAttributeCriteriaResponse) Set(val *ThirdPartyUncachedAttributeCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyUncachedAttributeCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyUncachedAttributeCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyUncachedAttributeCriteriaResponse(val *ThirdPartyUncachedAttributeCriteriaResponse) *NullableThirdPartyUncachedAttributeCriteriaResponse {
	return &NullableThirdPartyUncachedAttributeCriteriaResponse{value: val, isSet: true}
}

func (v NullableThirdPartyUncachedAttributeCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyUncachedAttributeCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}