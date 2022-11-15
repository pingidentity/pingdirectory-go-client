# HostSystemMonitorProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumhostSystemMonitorProviderSchemaUrn**](EnumhostSystemMonitorProviderSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Host System Monitor Provider is enabled for use. | 
**DiskDevices** | Pointer to **[]string** |  | [optional] 
**NetworkDevices** | Pointer to **[]string** |  | [optional] 
**SystemUtilizationMonitorLogDirectory** | **string** | Specifies a relative or absolute path to the directory on the local filesystem containing the log files used by the system utilization monitor. The path must exist, and it must be a writable directory by the server process. | 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 

## Methods

### NewHostSystemMonitorProviderResponse

`func NewHostSystemMonitorProviderResponse(schemas []EnumhostSystemMonitorProviderSchemaUrn, enabled bool, systemUtilizationMonitorLogDirectory string, ) *HostSystemMonitorProviderResponse`

NewHostSystemMonitorProviderResponse instantiates a new HostSystemMonitorProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostSystemMonitorProviderResponseWithDefaults

`func NewHostSystemMonitorProviderResponseWithDefaults() *HostSystemMonitorProviderResponse`

NewHostSystemMonitorProviderResponseWithDefaults instantiates a new HostSystemMonitorProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *HostSystemMonitorProviderResponse) GetSchemas() []EnumhostSystemMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HostSystemMonitorProviderResponse) GetSchemasOk() (*[]EnumhostSystemMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HostSystemMonitorProviderResponse) SetSchemas(v []EnumhostSystemMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *HostSystemMonitorProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *HostSystemMonitorProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *HostSystemMonitorProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDiskDevices

`func (o *HostSystemMonitorProviderResponse) GetDiskDevices() []string`

GetDiskDevices returns the DiskDevices field if non-nil, zero value otherwise.

### GetDiskDevicesOk

`func (o *HostSystemMonitorProviderResponse) GetDiskDevicesOk() (*[]string, bool)`

GetDiskDevicesOk returns a tuple with the DiskDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDevices

`func (o *HostSystemMonitorProviderResponse) SetDiskDevices(v []string)`

SetDiskDevices sets DiskDevices field to given value.

### HasDiskDevices

`func (o *HostSystemMonitorProviderResponse) HasDiskDevices() bool`

HasDiskDevices returns a boolean if a field has been set.

### GetNetworkDevices

`func (o *HostSystemMonitorProviderResponse) GetNetworkDevices() []string`

GetNetworkDevices returns the NetworkDevices field if non-nil, zero value otherwise.

### GetNetworkDevicesOk

`func (o *HostSystemMonitorProviderResponse) GetNetworkDevicesOk() (*[]string, bool)`

GetNetworkDevicesOk returns a tuple with the NetworkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevices

`func (o *HostSystemMonitorProviderResponse) SetNetworkDevices(v []string)`

SetNetworkDevices sets NetworkDevices field to given value.

### HasNetworkDevices

`func (o *HostSystemMonitorProviderResponse) HasNetworkDevices() bool`

HasNetworkDevices returns a boolean if a field has been set.

### GetSystemUtilizationMonitorLogDirectory

`func (o *HostSystemMonitorProviderResponse) GetSystemUtilizationMonitorLogDirectory() string`

GetSystemUtilizationMonitorLogDirectory returns the SystemUtilizationMonitorLogDirectory field if non-nil, zero value otherwise.

### GetSystemUtilizationMonitorLogDirectoryOk

`func (o *HostSystemMonitorProviderResponse) GetSystemUtilizationMonitorLogDirectoryOk() (*string, bool)`

GetSystemUtilizationMonitorLogDirectoryOk returns a tuple with the SystemUtilizationMonitorLogDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUtilizationMonitorLogDirectory

`func (o *HostSystemMonitorProviderResponse) SetSystemUtilizationMonitorLogDirectory(v string)`

SetSystemUtilizationMonitorLogDirectory sets SystemUtilizationMonitorLogDirectory field to given value.


### GetDescription

`func (o *HostSystemMonitorProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HostSystemMonitorProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HostSystemMonitorProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HostSystemMonitorProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


