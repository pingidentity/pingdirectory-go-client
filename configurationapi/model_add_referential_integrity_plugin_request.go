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

// checks if the AddReferentialIntegrityPluginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddReferentialIntegrityPluginRequest{}

// AddReferentialIntegrityPluginRequest struct for AddReferentialIntegrityPluginRequest
type AddReferentialIntegrityPluginRequest struct {
	Schemas    []EnumreferentialIntegrityPluginSchemaUrn `json:"schemas"`
	PluginType []EnumpluginPluginTypeProp                `json:"pluginType,omitempty"`
	// Specifies the attribute types for which referential integrity is to be maintained.
	AttributeType []string `json:"attributeType"`
	// Specifies the base DN that limits the scope within which referential integrity is maintained.
	BaseDN []string `json:"baseDN,omitempty"`
	// Specifies the log file location where the update records are written when the plug-in is in background-mode processing.
	LogFile *string `json:"logFile,omitempty"`
	// Specifies the interval in seconds when referential integrity updates are made.
	UpdateInterval *string `json:"updateInterval,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
	// Name of the new Plugin
	PluginName string `json:"pluginName"`
}

// NewAddReferentialIntegrityPluginRequest instantiates a new AddReferentialIntegrityPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddReferentialIntegrityPluginRequest(schemas []EnumreferentialIntegrityPluginSchemaUrn, attributeType []string, enabled bool, pluginName string) *AddReferentialIntegrityPluginRequest {
	this := AddReferentialIntegrityPluginRequest{}
	this.Schemas = schemas
	this.AttributeType = attributeType
	this.Enabled = enabled
	this.PluginName = pluginName
	return &this
}

// NewAddReferentialIntegrityPluginRequestWithDefaults instantiates a new AddReferentialIntegrityPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddReferentialIntegrityPluginRequestWithDefaults() *AddReferentialIntegrityPluginRequest {
	this := AddReferentialIntegrityPluginRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddReferentialIntegrityPluginRequest) GetSchemas() []EnumreferentialIntegrityPluginSchemaUrn {
	if o == nil {
		var ret []EnumreferentialIntegrityPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetSchemasOk() ([]EnumreferentialIntegrityPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddReferentialIntegrityPluginRequest) SetSchemas(v []EnumreferentialIntegrityPluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value if set, zero value otherwise.
func (o *AddReferentialIntegrityPluginRequest) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil || IsNil(o.PluginType) {
		var ret []EnumpluginPluginTypeProp
		return ret
	}
	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil || IsNil(o.PluginType) {
		return nil, false
	}
	return o.PluginType, true
}

// HasPluginType returns a boolean if a field has been set.
func (o *AddReferentialIntegrityPluginRequest) HasPluginType() bool {
	if o != nil && !IsNil(o.PluginType) {
		return true
	}

	return false
}

// SetPluginType gets a reference to the given []EnumpluginPluginTypeProp and assigns it to the PluginType field.
func (o *AddReferentialIntegrityPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetAttributeType returns the AttributeType field value
func (o *AddReferentialIntegrityPluginRequest) GetAttributeType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetAttributeTypeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeType, true
}

// SetAttributeType sets field value
func (o *AddReferentialIntegrityPluginRequest) SetAttributeType(v []string) {
	o.AttributeType = v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AddReferentialIntegrityPluginRequest) GetBaseDN() []string {
	if o == nil || IsNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AddReferentialIntegrityPluginRequest) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *AddReferentialIntegrityPluginRequest) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetLogFile returns the LogFile field value if set, zero value otherwise.
func (o *AddReferentialIntegrityPluginRequest) GetLogFile() string {
	if o == nil || IsNil(o.LogFile) {
		var ret string
		return ret
	}
	return *o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetLogFileOk() (*string, bool) {
	if o == nil || IsNil(o.LogFile) {
		return nil, false
	}
	return o.LogFile, true
}

// HasLogFile returns a boolean if a field has been set.
func (o *AddReferentialIntegrityPluginRequest) HasLogFile() bool {
	if o != nil && !IsNil(o.LogFile) {
		return true
	}

	return false
}

// SetLogFile gets a reference to the given string and assigns it to the LogFile field.
func (o *AddReferentialIntegrityPluginRequest) SetLogFile(v string) {
	o.LogFile = &v
}

// GetUpdateInterval returns the UpdateInterval field value if set, zero value otherwise.
func (o *AddReferentialIntegrityPluginRequest) GetUpdateInterval() string {
	if o == nil || IsNil(o.UpdateInterval) {
		var ret string
		return ret
	}
	return *o.UpdateInterval
}

// GetUpdateIntervalOk returns a tuple with the UpdateInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetUpdateIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateInterval) {
		return nil, false
	}
	return o.UpdateInterval, true
}

// HasUpdateInterval returns a boolean if a field has been set.
func (o *AddReferentialIntegrityPluginRequest) HasUpdateInterval() bool {
	if o != nil && !IsNil(o.UpdateInterval) {
		return true
	}

	return false
}

// SetUpdateInterval gets a reference to the given string and assigns it to the UpdateInterval field.
func (o *AddReferentialIntegrityPluginRequest) SetUpdateInterval(v string) {
	o.UpdateInterval = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddReferentialIntegrityPluginRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddReferentialIntegrityPluginRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddReferentialIntegrityPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddReferentialIntegrityPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddReferentialIntegrityPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddReferentialIntegrityPluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddReferentialIntegrityPluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !IsNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddReferentialIntegrityPluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

// GetPluginName returns the PluginName field value
func (o *AddReferentialIntegrityPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddReferentialIntegrityPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddReferentialIntegrityPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

func (o AddReferentialIntegrityPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddReferentialIntegrityPluginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	if !IsNil(o.PluginType) {
		toSerialize["pluginType"] = o.PluginType
	}
	toSerialize["attributeType"] = o.AttributeType
	if !IsNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !IsNil(o.LogFile) {
		toSerialize["logFile"] = o.LogFile
	}
	if !IsNil(o.UpdateInterval) {
		toSerialize["updateInterval"] = o.UpdateInterval
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	toSerialize["pluginName"] = o.PluginName
	return toSerialize, nil
}

type NullableAddReferentialIntegrityPluginRequest struct {
	value *AddReferentialIntegrityPluginRequest
	isSet bool
}

func (v NullableAddReferentialIntegrityPluginRequest) Get() *AddReferentialIntegrityPluginRequest {
	return v.value
}

func (v *NullableAddReferentialIntegrityPluginRequest) Set(val *AddReferentialIntegrityPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddReferentialIntegrityPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddReferentialIntegrityPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddReferentialIntegrityPluginRequest(val *AddReferentialIntegrityPluginRequest) *NullableAddReferentialIntegrityPluginRequest {
	return &NullableAddReferentialIntegrityPluginRequest{value: val, isSet: true}
}

func (v NullableAddReferentialIntegrityPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddReferentialIntegrityPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
