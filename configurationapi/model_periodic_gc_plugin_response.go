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

// PeriodicGcPluginResponse struct for PeriodicGcPluginResponse
type PeriodicGcPluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the Plugin
	Id                string                            `json:"id"`
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

// NewPeriodicGcPluginResponse instantiates a new PeriodicGcPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriodicGcPluginResponse(id string, schemas []EnumperiodicGcPluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, invokeGCTimeUtc []string, enabled bool) *PeriodicGcPluginResponse {
	this := PeriodicGcPluginResponse{}
	this.Id = id
	this.Schemas = schemas
	this.PluginType = pluginType
	this.InvokeGCTimeUtc = invokeGCTimeUtc
	this.Enabled = enabled
	return &this
}

// NewPeriodicGcPluginResponseWithDefaults instantiates a new PeriodicGcPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodicGcPluginResponseWithDefaults() *PeriodicGcPluginResponse {
	this := PeriodicGcPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PeriodicGcPluginResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PeriodicGcPluginResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *PeriodicGcPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *PeriodicGcPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *PeriodicGcPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *PeriodicGcPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *PeriodicGcPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PeriodicGcPluginResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *PeriodicGcPluginResponse) GetSchemas() []EnumperiodicGcPluginSchemaUrn {
	if o == nil {
		var ret []EnumperiodicGcPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetSchemasOk() ([]EnumperiodicGcPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *PeriodicGcPluginResponse) SetSchemas(v []EnumperiodicGcPluginSchemaUrn) {
	o.Schemas = v
}

// GetPluginType returns the PluginType field value
func (o *PeriodicGcPluginResponse) GetPluginType() []EnumpluginPluginTypeProp {
	if o == nil {
		var ret []EnumpluginPluginTypeProp
		return ret
	}

	return o.PluginType
}

// GetPluginTypeOk returns a tuple with the PluginType field value
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetPluginTypeOk() ([]EnumpluginPluginTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginType, true
}

// SetPluginType sets field value
func (o *PeriodicGcPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp) {
	o.PluginType = v
}

// GetInvokeGCDayOfWeek returns the InvokeGCDayOfWeek field value if set, zero value otherwise.
func (o *PeriodicGcPluginResponse) GetInvokeGCDayOfWeek() []EnumpluginInvokeGCDayOfWeekProp {
	if o == nil || isNil(o.InvokeGCDayOfWeek) {
		var ret []EnumpluginInvokeGCDayOfWeekProp
		return ret
	}
	return o.InvokeGCDayOfWeek
}

// GetInvokeGCDayOfWeekOk returns a tuple with the InvokeGCDayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetInvokeGCDayOfWeekOk() ([]EnumpluginInvokeGCDayOfWeekProp, bool) {
	if o == nil || isNil(o.InvokeGCDayOfWeek) {
		return nil, false
	}
	return o.InvokeGCDayOfWeek, true
}

// HasInvokeGCDayOfWeek returns a boolean if a field has been set.
func (o *PeriodicGcPluginResponse) HasInvokeGCDayOfWeek() bool {
	if o != nil && !isNil(o.InvokeGCDayOfWeek) {
		return true
	}

	return false
}

// SetInvokeGCDayOfWeek gets a reference to the given []EnumpluginInvokeGCDayOfWeekProp and assigns it to the InvokeGCDayOfWeek field.
func (o *PeriodicGcPluginResponse) SetInvokeGCDayOfWeek(v []EnumpluginInvokeGCDayOfWeekProp) {
	o.InvokeGCDayOfWeek = v
}

// GetInvokeGCTimeUtc returns the InvokeGCTimeUtc field value
func (o *PeriodicGcPluginResponse) GetInvokeGCTimeUtc() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InvokeGCTimeUtc
}

// GetInvokeGCTimeUtcOk returns a tuple with the InvokeGCTimeUtc field value
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetInvokeGCTimeUtcOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvokeGCTimeUtc, true
}

// SetInvokeGCTimeUtc sets field value
func (o *PeriodicGcPluginResponse) SetInvokeGCTimeUtc(v []string) {
	o.InvokeGCTimeUtc = v
}

// GetDelayAfterAlert returns the DelayAfterAlert field value if set, zero value otherwise.
func (o *PeriodicGcPluginResponse) GetDelayAfterAlert() string {
	if o == nil || isNil(o.DelayAfterAlert) {
		var ret string
		return ret
	}
	return *o.DelayAfterAlert
}

// GetDelayAfterAlertOk returns a tuple with the DelayAfterAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetDelayAfterAlertOk() (*string, bool) {
	if o == nil || isNil(o.DelayAfterAlert) {
		return nil, false
	}
	return o.DelayAfterAlert, true
}

// HasDelayAfterAlert returns a boolean if a field has been set.
func (o *PeriodicGcPluginResponse) HasDelayAfterAlert() bool {
	if o != nil && !isNil(o.DelayAfterAlert) {
		return true
	}

	return false
}

// SetDelayAfterAlert gets a reference to the given string and assigns it to the DelayAfterAlert field.
func (o *PeriodicGcPluginResponse) SetDelayAfterAlert(v string) {
	o.DelayAfterAlert = &v
}

// GetDelayPostGC returns the DelayPostGC field value if set, zero value otherwise.
func (o *PeriodicGcPluginResponse) GetDelayPostGC() string {
	if o == nil || isNil(o.DelayPostGC) {
		var ret string
		return ret
	}
	return *o.DelayPostGC
}

// GetDelayPostGCOk returns a tuple with the DelayPostGC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetDelayPostGCOk() (*string, bool) {
	if o == nil || isNil(o.DelayPostGC) {
		return nil, false
	}
	return o.DelayPostGC, true
}

// HasDelayPostGC returns a boolean if a field has been set.
func (o *PeriodicGcPluginResponse) HasDelayPostGC() bool {
	if o != nil && !isNil(o.DelayPostGC) {
		return true
	}

	return false
}

// SetDelayPostGC gets a reference to the given string and assigns it to the DelayPostGC field.
func (o *PeriodicGcPluginResponse) SetDelayPostGC(v string) {
	o.DelayPostGC = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PeriodicGcPluginResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PeriodicGcPluginResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PeriodicGcPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *PeriodicGcPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *PeriodicGcPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetInvokeForInternalOperations returns the InvokeForInternalOperations field value if set, zero value otherwise.
func (o *PeriodicGcPluginResponse) GetInvokeForInternalOperations() bool {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		var ret bool
		return ret
	}
	return *o.InvokeForInternalOperations
}

// GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeriodicGcPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeForInternalOperations) {
		return nil, false
	}
	return o.InvokeForInternalOperations, true
}

// HasInvokeForInternalOperations returns a boolean if a field has been set.
func (o *PeriodicGcPluginResponse) HasInvokeForInternalOperations() bool {
	if o != nil && !isNil(o.InvokeForInternalOperations) {
		return true
	}

	return false
}

// SetInvokeForInternalOperations gets a reference to the given bool and assigns it to the InvokeForInternalOperations field.
func (o *PeriodicGcPluginResponse) SetInvokeForInternalOperations(v bool) {
	o.InvokeForInternalOperations = &v
}

func (o PeriodicGcPluginResponse) MarshalJSON() ([]byte, error) {
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

type NullablePeriodicGcPluginResponse struct {
	value *PeriodicGcPluginResponse
	isSet bool
}

func (v NullablePeriodicGcPluginResponse) Get() *PeriodicGcPluginResponse {
	return v.value
}

func (v *NullablePeriodicGcPluginResponse) Set(val *PeriodicGcPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodicGcPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodicGcPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodicGcPluginResponse(val *PeriodicGcPluginResponse) *NullablePeriodicGcPluginResponse {
	return &NullablePeriodicGcPluginResponse{value: val, isSet: true}
}

func (v NullablePeriodicGcPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodicGcPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}