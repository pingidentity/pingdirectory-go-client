# HostSystemMonitorProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumhostSystemMonitorProviderSchemaUrn**](EnumhostSystemMonitorProviderSchemaUrn.md) |  | 
**Id** | **string** | Name of the Monitor Provider | 
**Enabled** | **bool** | Indicates whether the Host System Monitor Provider is enabled for use. | 
**DiskDevices** | Pointer to **[]string** | Specifies which disk devices to monitor for I/O activity. Should be the device name as displayed by iostat -d. | [optional] 
**NetworkDevices** | Pointer to **[]string** | Specifies which network interfaces to monitor for I/O activity. Should be the device name as displayed by netstat -i. | [optional] 
**SystemUtilizationMonitorLogDirectory** | **string** | Specifies a relative or absolute path to the directory on the local filesystem containing the log files used by the system utilization monitor. The path must exist, and it must be a writable directory by the server process. | 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 

## Methods

### NewHostSystemMonitorProviderResponse

`func NewHostSystemMonitorProviderResponse(schemas []EnumhostSystemMonitorProviderSchemaUrn, id string, enabled bool, systemUtilizationMonitorLogDirectory string, ) *HostSystemMonitorProviderResponse`

NewHostSystemMonitorProviderResponse instantiates a new HostSystemMonitorProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostSystemMonitorProviderResponseWithDefaults

`func NewHostSystemMonitorProviderResponseWithDefaults() *HostSystemMonitorProviderResponse`

NewHostSystemMonitorProviderResponseWithDefaults instantiates a new HostSystemMonitorProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *HostSystemMonitorProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HostSystemMonitorProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HostSystemMonitorProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HostSystemMonitorProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *HostSystemMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *HostSystemMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *HostSystemMonitorProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *HostSystemMonitorProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

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


### GetId

`func (o *HostSystemMonitorProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostSystemMonitorProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostSystemMonitorProviderResponse) SetId(v string)`

SetId sets Id field to given value.


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


