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

// AddPeriodicGcPluginRequest struct for AddPeriodicGcPluginRequest
type AddPeriodicGcPluginRequest struct {
	// Name of the new Plugin
	PluginName        string                            `json:"pluginName"`
	Schemas           []EnumperiodicGcPluginSchemaUrn   `json:"schemas"`
	PluginType        []EnumpluginPluginTypeProp        `json:"pluginType"`
	InvokeGCDayOfWeek []EnumpluginInvokeGCDayOfWeekProp `json:"invokeGCDayOfWeek,omitempty"`
	// Specifies the times of the day at which garbage collection may be explicitly invoked. The times should be specified in \"HH:MM\" format, with \"HH\" as a two-digit numeric value between 00 and 23 representing the hour of the day, and MM as a two-digit numeric value between 00 and 59 representing the minute of the hour. All times will be interpreted in the UTC time zone.
	InvokeGCTimeUtc []string `json:"invokeGCTimeUtc"`
	// Specifies the length of time that the Directory Server should wait after sending the \"force-gc-starting\" administrative alert before actually invoking the garbage collection processing.
	DelayAfterAlert *string `json:"delayAfterAlert,omitempty"`
	// Specifies the length of time that the Directory Server should wait after successfully completing the garbage collection processing, before removing the \"force-gc-starting\" administrative alert, which marks the server as unavailable.
	DelayPostGC *string `json:"delayPostGC,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
	// Indicates whether the plug-in should be invoked for internal operations.
	InvokeForInternalOperations *bool `json:"invokeForInternalOperations,omitempty"`
}

// NewAddPeriodicGcPluginRequest instantiates a new AddPeriodicGcPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPeriodicGcPluginRequest(pluginName string, schemas []EnumperiodicGcPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, invokeGCTimeUtc []string, enabled bool) *AddPeriodicGcPluginRequest {
	this := AddPeriodicGcPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.PluginType = pluginType
	this.InvokeGCTimeUtc = invokeGCTimeUtc
	this.Enabled = enabled
	return &this
}

// NewAddPeriodicGcPluginRequestWithDefaults instantiates a new AddPeriodicGcPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPeriodicGcPluginRequestWithDefaults() *AddPeriodicGcPluginRequest {
	this := AddPeriodicGcPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddPeriodicGcPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddPeriodicGcPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPeriodicGcPluginRequest) GetSchemas() []EnumperiodicGcPluginSchemaUrn {
	if o == nil {
		var ret []EnumperiodicGcPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetSchemasOk() ([]EnumperiodicGcPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPeriodicGcPluginRequest) SetSchemas(v []EnumperiodicGcPluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value
func (o *AddPeriodicGcPluginRequest) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *AddPeriodicGcPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetInvokeGCDayOfWeek returns the InvokeGCDayOfWeek field value if set, zero value otherwise.
func (o *AddPeriodicGcPluginRequest) GetInvokeGCDayOfWeek() []EnumpluginInvokeGCDayOfWeekProp {
	if o == nil || isNil(o.InvokeGCDayOfWeek) {
		var ret []EnumpluginInvokeGCDayOfWeekProp
		return ret
	}
	return o.InvokeGCDayOfWeek
}

// GetInvokeGCDayOfWeekOk returns a tuple with the InvokeGCDayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetInvokeGCDayOfWeekOk() ([]EnumpluginInvokeGCDayOfWeekProp, bool) {
	if o == nil || isNil(o.InvokeGCDayOfWeek) {
		return nil, false
	}
	return o.InvokeGCDayOfWeek, true
}

// HasInvokeGCDayOfWeek returns a boolean if a field has been set.
func (o *AddPeriodicGcPluginRequest) HasInvokeGCDayOfWeek() bool {
	if o != nil && !isNil(o.InvokeGCDayOfWeek) {
		return true
	}

	return false
}

// SetInvokeGCDayOfWeek gets a reference to the given []EnumpluginInvokeGCDayOfWeekProp and assigns it to the InvokeGCDayOfWeek field.
func (o *AddPeriodicGcPluginRequest) SetInvokeGCDayOfWeek(v []EnumpluginInvokeGCDayOfWeekProp) {
	o.InvokeGCDayOfWeek = v
}

// GetInvokeGCTimeUtc returns the InvokeGCTimeUtc field value
func (o *AddPeriodicGcPluginRequest) GetInvokeGCTimeUtc() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InvokeGCTimeUtc
}

// GetInvokeGCTimeUtcOk returns a tuple with the InvokeGCTimeUtc field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetInvokeGCTimeUtcOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvokeGCTimeUtc, true
}

// SetInvokeGCTimeUtc sets field value
func (o *AddPeriodicGcPluginRequest) SetInvokeGCTimeUtc(v []string) {
	o.InvokeGCTimeUtc = v
}

// GetDelayAfterAlert returns the DelayAfterAlert field value if set, zero value otherwise.
func (o *AddPeriodicGcPluginRequest) GetDelayAfterAlert() string {
	if o == nil || isNil(o.DelayAfterAlert) {
		var ret string
		return ret
	}
	return *o.DelayAfterAlert
}

// GetDelayAfterAlertOk returns a tuple with the DelayAfterAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetDelayAfterAlertOk() (*string, bool) {
	if o == nil || isNil(o.DelayAfterAlert) {
		return nil, false
	}
	return o.DelayAfterAlert, true
}

// HasDelayAfterAlert returns a boolean if a field has been set.
func (o *AddPeriodicGcPluginRequest) HasDelayAfterAlert() bool {
	if o != nil && !isNil(o.DelayAfterAlert) {
		return true
	}

	return false
}

// SetDelayAfterAlert gets a reference to the given string and assigns it to the DelayAfterAlert field.
func (o *AddPeriodicGcPluginRequest) SetDelayAfterAlert(v string) {
	o.DelayAfterAlert = &v
}

// GetDelayPostGC returns the DelayPostGC field value if set, zero value otherwise.
func (o *AddPeriodicGcPluginRequest) GetDelayPostGC() string {
	if o == nil || isNil(o.DelayPostGC) {
		var ret string
		return ret
	}
	return *o.DelayPostGC
}

// GetDelayPostGCOk returns a tuple with the DelayPostGC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetDelayPostGCOk() (*string, bool) {
	if o == nil || isNil(o.DelayPostGC) {
		return nil, false
	}
	return o.DelayPostGC, true
}

// HasDelayPostGC returns a boolean if a field has been set.
func (o *AddPeriodicGcPluginRequest) HasDelayPostGC() bool {
	if o != nil && !isNil(o.DelayPostGC) {
		return true
	}

	return false
}

// SetDelayPostGC gets a reference to the given string and assigns it to the DelayPostGC field.
func (o *AddPeriodicGcPluginRequest) SetDelayPostGC(v string) {
	o.DelayPostGC = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPeriodicGcPluginRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPeriodicGcPluginRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPeriodicGcPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPeriodicGcPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPeriodicGcPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *AddPeriodicGcPluginRequest) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicGcPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *AddPeriodicGcPluginRequest) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *AddPeriodicGcPluginRequest) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o AddPeriodicGcPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pluginName"] = o.PluginName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["pluginType"] = o.PluginType
	}
	if !isNil(o.InvokeGCDayOfWeek) {
		toSerialize["invokeGCDayOfWeek"] = o.InvokeGCDayOfWeek
	}
	if true {
		toSerialize["invokeGCTimeUtc"] = o.InvokeGCTimeUtc
	}
	if !isNil(o.DelayAfterAlert) {
		toSerialize["delayAfterAlert"] = o.DelayAfterAlert
	}
	if !isNil(o.DelayPostGC) {
		toSerialize["delayPostGC"] = o.DelayPostGC
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.InvokeForInternalOperations) {
		toSerialize["invokeForInternalOperations"] = o.InvokeForInternalOperations
	}
	return json.Marshal(toSerialize)
}

type NullableAddPeriodicGcPluginRequest struct {
	value *AddPeriodicGcPluginRequest
	isSet bool
}

func (v NullableAddPeriodicGcPluginRequest) Get() *AddPeriodicGcPluginRequest {
	return v.value
}

func (v *NullableAddPeriodicGcPluginRequest) Set(val *AddPeriodicGcPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPeriodicGcPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPeriodicGcPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPeriodicGcPluginRequest(val *AddPeriodicGcPluginRequest) *NullableAddPeriodicGcPluginRequest {
	return &NullableAddPeriodicGcPluginRequest{value: val, isSet: true}
}

func (v NullableAddPeriodicGcPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPeriodicGcPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}