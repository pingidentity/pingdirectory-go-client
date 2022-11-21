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

// HostSystemMonitorProviderResponse struct for HostSystemMonitorProviderResponse
type HostSystemMonitorProviderResponse struct {
	Schemas []EnumhostSystemMonitorProviderSchemaUrn `json:"schemas"`
	// Indicates whether the Host System Monitor Provider is enabled for use.
	Enabled bool `json:"enabled"`
	// Specifies which disk devices to monitor for I/O activity. Should be the device name as displayed by iostat -d.
	DiskDevices []string `json:"diskDevices,omitempty"`
	// Specifies which network interfaces to monitor for I/O activity. Should be the device name as displayed by netstat -i.
	NetworkDevices []string `json:"networkDevices,omitempty"`
	// Specifies a relative or absolute path to the directory on the local filesystem containing the log files used by the system utilization monitor. The path must exist, and it must be a writable directory by the server process.
	SystemUtilizationMonitorLogDirectory string `json:"systemUtilizationMonitorLogDirectory"`
	// A description for this Monitor Provider
	Description *string `json:"description,omitempty"`
}

// NewHostSystemMonitorProviderResponse instantiates a new HostSystemMonitorProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostSystemMonitorProviderResponse(schemas []EnumhostSystemMonitorProviderSchemaUrn, enabled bool, systemUtilizationMonitorLogDirectory string) *HostSystemMonitorProviderResponse {
	this := HostSystemMonitorProviderResponse{}
	this.Schemas = schemas
	this.Enabled = enabled
	this.SystemUtilizationMonitorLogDirectory = systemUtilizationMonitorLogDirectory
	return &this
}

// NewHostSystemMonitorProviderResponseWithDefaults instantiates a new HostSystemMonitorProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostSystemMonitorProviderResponseWithDefaults() *HostSystemMonitorProviderResponse {
	this := HostSystemMonitorProviderResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *HostSystemMonitorProviderResponse) GetSchemas() []EnumhostSystemMonitorProviderSchemaUrn {
	if o == nil {
		var ret []EnumhostSystemMonitorProviderSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *HostSystemMonitorProviderResponse) GetSchemasOk() ([]EnumhostSystemMonitorProviderSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *HostSystemMonitorProviderResponse) SetSchemas(v []EnumhostSystemMonitorProviderSchemaUrn) {
	o.Schemas = v
}

// GetEnabled returns the Enabled field value
func (o *HostSystemMonitorProviderResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *HostSystemMonitorProviderResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *HostSystemMonitorProviderResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetDiskDevices returns the DiskDevices field value if set, zero value otherwise.
func (o *HostSystemMonitorProviderResponse) GetDiskDevices() []string {
	if o == nil || isNil(o.DiskDevices) {
		var ret []string
		return ret
	}
	return o.DiskDevices
}

// GetDiskDevicesOk returns a tuple with the DiskDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostSystemMonitorProviderResponse) GetDiskDevicesOk() ([]string, bool) {
	if o == nil || isNil(o.DiskDevices) {
    return nil, false
	}
	return o.DiskDevices, true
}

// HasDiskDevices returns a boolean if a field has been set.
func (o *HostSystemMonitorProviderResponse) HasDiskDevices() bool {
	if o != nil && !isNil(o.DiskDevices) {
		return true
	}

	return false
}

// SetDiskDevices gets a reference to the given []string and assigns it to the DiskDevices field.
func (o *HostSystemMonitorProviderResponse) SetDiskDevices(v []string) {
	o.DiskDevices = v
}

// GetNetworkDevices returns the NetworkDevices field value if set, zero value otherwise.
func (o *HostSystemMonitorProviderResponse) GetNetworkDevices() []string {
	if o == nil || isNil(o.NetworkDevices) {
		var ret []string
		return ret
	}
	return o.NetworkDevices
}

// GetNetworkDevicesOk returns a tuple with the NetworkDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostSystemMonitorProviderResponse) GetNetworkDevicesOk() ([]string, bool) {
	if o == nil || isNil(o.NetworkDevices) {
    return nil, false
	}
	return o.NetworkDevices, true
}

// HasNetworkDevices returns a boolean if a field has been set.
func (o *HostSystemMonitorProviderResponse) HasNetworkDevices() bool {
	if o != nil && !isNil(o.NetworkDevices) {
		return true
	}

	return false
}

// SetNetworkDevices gets a reference to the given []string and assigns it to the NetworkDevices field.
func (o *HostSystemMonitorProviderResponse) SetNetworkDevices(v []string) {
	o.NetworkDevices = v
}

// GetSystemUtilizationMonitorLogDirectory returns the SystemUtilizationMonitorLogDirectory field value
func (o *HostSystemMonitorProviderResponse) GetSystemUtilizationMonitorLogDirectory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemUtilizationMonitorLogDirectory
}

// GetSystemUtilizationMonitorLogDirectoryOk returns a tuple with the SystemUtilizationMonitorLogDirectory field value
// and a boolean to check if the value has been set.
func (o *HostSystemMonitorProviderResponse) GetSystemUtilizationMonitorLogDirectoryOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SystemUtilizationMonitorLogDirectory, true
}

// SetSystemUtilizationMonitorLogDirectory sets field value
func (o *HostSystemMonitorProviderResponse) SetSystemUtilizationMonitorLogDirectory(v string) {
	o.SystemUtilizationMonitorLogDirectory = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *HostSystemMonitorProviderResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostSystemMonitorProviderResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *HostSystemMonitorProviderResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *HostSystemMonitorProviderResponse) SetDescription(v string) {
	o.Description = &v
}

func (o HostSystemMonitorProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DiskDevices) {
		toSerialize["diskDevices"] = o.DiskDevices
	}
	if !isNil(o.NetworkDevices) {
		toSerialize["networkDevices"] = o.NetworkDevices
	}
	if true {
		toSerialize["systemUtilizationMonitorLogDirectory"] = o.SystemUtilizationMonitorLogDirectory
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableHostSystemMonitorProviderResponse struct {
	value *HostSystemMonitorProviderResponse
	isSet bool
}

func (v NullableHostSystemMonitorProviderResponse) Get() *HostSystemMonitorProviderResponse {
	return v.value
}

func (v *NullableHostSystemMonitorProviderResponse) Set(val *HostSystemMonitorProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableHostSystemMonitorProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableHostSystemMonitorProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostSystemMonitorProviderResponse(val *HostSystemMonitorProviderResponse) *NullableHostSystemMonitorProviderResponse {
	return &NullableHostSystemMonitorProviderResponse{value: val, isSet: true}
}

func (v NullableHostSystemMonitorProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostSystemMonitorProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


