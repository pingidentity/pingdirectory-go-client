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

// checks if the ThirdPartyConnectionCriteriaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartyConnectionCriteriaResponse{}

// ThirdPartyConnectionCriteriaResponse struct for ThirdPartyConnectionCriteriaResponse
type ThirdPartyConnectionCriteriaResponse struct {
	Schemas []EnumthirdPartyConnectionCriteriaSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Connection Criteria.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Connection Criteria. Each configuration property should be given in the form 'name=value'.
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
	// A description for this Connection Criteria
	Description                                   *string                                            `json:"description,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Connection Criteria
	Id string `json:"id"`
}

// NewThirdPartyConnectionCriteriaResponse instantiates a new ThirdPartyConnectionCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyConnectionCriteriaResponse(schemas []EnumthirdPartyConnectionCriteriaSchemaUrn, extensionClass string, id string) *ThirdPartyConnectionCriteriaResponse {
	this := ThirdPartyConnectionCriteriaResponse{}
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Id = id
	return &this
}

// NewThirdPartyConnectionCriteriaResponseWithDefaults instantiates a new ThirdPartyConnectionCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyConnectionCriteriaResponseWithDefaults() *ThirdPartyConnectionCriteriaResponse {
	this := ThirdPartyConnectionCriteriaResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyConnectionCriteriaResponse) GetSchemas() []EnumthirdPartyConnectionCriteriaSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyConnectionCriteriaSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyConnectionCriteriaResponse) GetSchemasOk() ([]EnumthirdPartyConnectionCriteriaSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyConnectionCriteriaResponse) SetSchemas(v []EnumthirdPartyConnectionCriteriaSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyConnectionCriteriaResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyConnectionCriteriaResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyConnectionCriteriaResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyConnectionCriteriaResponse) GetExtensionArgument() []string {
	if o == nil || IsNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyConnectionCriteriaResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtensionArgument) {
		return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyConnectionCriteriaResponse) HasExtensionArgument() bool {
	if o != nil && !IsNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyConnectionCriteriaResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyConnectionCriteriaResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyConnectionCriteriaResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyConnectionCriteriaResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyConnectionCriteriaResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyConnectionCriteriaResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyConnectionCriteriaResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyConnectionCriteriaResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyConnectionCriteriaResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *ThirdPartyConnectionCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyConnectionCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *ThirdPartyConnectionCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *ThirdPartyConnectionCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *ThirdPartyConnectionCriteriaResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyConnectionCriteriaResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyConnectionCriteriaResponse) SetId(v string) {
	o.Id = v
}

func (o ThirdPartyConnectionCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartyConnectionCriteriaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["extensionClass"] = o.ExtensionClass
	if !IsNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableThirdPartyConnectionCriteriaResponse struct {
	value *ThirdPartyConnectionCriteriaResponse
	isSet bool
}

func (v NullableThirdPartyConnectionCriteriaResponse) Get() *ThirdPartyConnectionCriteriaResponse {
	return v.value
}

func (v *NullableThirdPartyConnectionCriteriaResponse) Set(val *ThirdPartyConnectionCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyConnectionCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyConnectionCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyConnectionCriteriaResponse(val *ThirdPartyConnectionCriteriaResponse) *NullableThirdPartyConnectionCriteriaResponse {
	return &NullableThirdPartyConnectionCriteriaResponse{value: val, isSet: true}
}

func (v NullableThirdPartyConnectionCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyConnectionCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
