# GetMonitorProvider200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyMonitorProviderSchemaUrn**](EnumthirdPartyMonitorProviderSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Monitor Provider is enabled for use. | 
**LowSpaceWarningSizeThreshold** | Pointer to **string** | Specifies the low space warning threshold value as an absolute amount of space. If the amount of usable disk space drops below this amount, then the Directory Server will begin generating warning alert notifications. | [optional] 
**LowSpaceWarningPercentThreshold** | Pointer to **int32** | Specifies the low space warning threshold value as a percentage of total space. If the amount of usable disk space drops below this amount, then the Directory Server will begin generating warning alert notifications. | [optional] 
**LowSpaceErrorSizeThreshold** | Pointer to **string** | Specifies the low space error threshold value as an absolute amount of space. If the amount of usable disk space drops below this amount, then the Directory Server will start rejecting operations requested by non-root users. | [optional] 
**LowSpaceErrorPercentThreshold** | Pointer to **int32** | Specifies the low space error threshold value as a percentage of total space. If the amount of usable disk space drops below this amount, then the Directory Server will start rejecting operations requested by non-root users. | [optional] 
**OutOfSpaceErrorSizeThreshold** | Pointer to **string** | Specifies the out of space error threshold value as an absolute amount of space. If the amount of usable disk space drops below this amount, then the Directory Server will shut itself down to avoid problems that may occur from complete exhaustion of usable space. | [optional] 
**OutOfSpaceErrorPercentThreshold** | Pointer to **int32** | Specifies the out of space error threshold value as a percentage of total space. If the amount of usable disk space drops below this amount, then the Directory Server will shut itself down to avoid problems that may occur from complete exhaustion of usable space. | [optional] 
**AlertFrequency** | **string** | Specifies the length of time between administrative alerts generated in response to lack of usable disk space. Administrative alerts will be generated whenever the amount of usable space drops below any threshold, and they will also be generated at regular intervals as long as the amount of usable space remains below the threshold value. A value of zero indicates that alerts should only be generated when the amount of usable space drops below a configured threshold. | 
**DiskDevices** | Pointer to **[]string** | Specifies which disk devices to monitor for I/O activity. Should be the device name as displayed by iostat -d. | [optional] 
**NetworkDevices** | Pointer to **[]string** | Specifies which network interfaces to monitor for I/O activity. Should be the device name as displayed by netstat -i. | [optional] 
**SystemUtilizationMonitorLogDirectory** | **string** | Specifies a relative or absolute path to the directory on the local filesystem containing the log files used by the system utilization monitor. The path must exist, and it must be a writable directory by the server process. | 
**Id** | **string** | Name of the Monitor Provider | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Monitor Provider. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Monitor Provider. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewGetMonitorProvider200Response

`func NewGetMonitorProvider200Response(schemas []EnumthirdPartyMonitorProviderSchemaUrn, enabled bool, alertFrequency string, systemUtilizationMonitorLogDirectory string, id string, extensionClass string, ) *GetMonitorProvider200Response`

NewGetMonitorProvider200Response instantiates a new GetMonitorProvider200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMonitorProvider200ResponseWithDefaults

`func NewGetMonitorProvider200ResponseWithDefaults() *GetMonitorProvider200Response`

NewGetMonitorProvider200ResponseWithDefaults instantiates a new GetMonitorProvider200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetMonitorProvider200Response) GetSchemas() []EnumthirdPartyMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetMonitorProvider200Response) GetSchemasOk() (*[]EnumthirdPartyMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetMonitorProvider200Response) SetSchemas(v []EnumthirdPartyMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *GetMonitorProvider200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMonitorProvider200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMonitorProvider200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMonitorProvider200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetMonitorProvider200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetMonitorProvider200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetMonitorProvider200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLowSpaceWarningSizeThreshold

`func (o *GetMonitorProvider200Response) GetLowSpaceWarningSizeThreshold() string`

GetLowSpaceWarningSizeThreshold returns the LowSpaceWarningSizeThreshold field if non-nil, zero value otherwise.

### GetLowSpaceWarningSizeThresholdOk

`func (o *GetMonitorProvider200Response) GetLowSpaceWarningSizeThresholdOk() (*string, bool)`

GetLowSpaceWarningSizeThresholdOk returns a tuple with the LowSpaceWarningSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceWarningSizeThreshold

`func (o *GetMonitorProvider200Response) SetLowSpaceWarningSizeThreshold(v string)`

SetLowSpaceWarningSizeThreshold sets LowSpaceWarningSizeThreshold field to given value.

### HasLowSpaceWarningSizeThreshold

`func (o *GetMonitorProvider200Response) HasLowSpaceWarningSizeThreshold() bool`

HasLowSpaceWarningSizeThreshold returns a boolean if a field has been set.

### GetLowSpaceWarningPercentThreshold

`func (o *GetMonitorProvider200Response) GetLowSpaceWarningPercentThreshold() int32`

GetLowSpaceWarningPercentThreshold returns the LowSpaceWarningPercentThreshold field if non-nil, zero value otherwise.

### GetLowSpaceWarningPercentThresholdOk

`func (o *GetMonitorProvider200Response) GetLowSpaceWarningPercentThresholdOk() (*int32, bool)`

GetLowSpaceWarningPercentThresholdOk returns a tuple with the LowSpaceWarningPercentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceWarningPercentThreshold

`func (o *GetMonitorProvider200Response) SetLowSpaceWarningPercentThreshold(v int32)`

SetLowSpaceWarningPercentThreshold sets LowSpaceWarningPercentThreshold field to given value.

### HasLowSpaceWarningPercentThreshold

`func (o *GetMonitorProvider200Response) HasLowSpaceWarningPercentThreshold() bool`

HasLowSpaceWarningPercentThreshold returns a boolean if a field has been set.

### GetLowSpaceErrorSizeThreshold

`func (o *GetMonitorProvider200Response) GetLowSpaceErrorSizeThreshold() string`

GetLowSpaceErrorSizeThreshold returns the LowSpaceErrorSizeThreshold field if non-nil, zero value otherwise.

### GetLowSpaceErrorSizeThresholdOk

`func (o *GetMonitorProvider200Response) GetLowSpaceErrorSizeThresholdOk() (*string, bool)`

GetLowSpaceErrorSizeThresholdOk returns a tuple with the LowSpaceErrorSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceErrorSizeThreshold

`func (o *GetMonitorProvider200Response) SetLowSpaceErrorSizeThreshold(v string)`

SetLowSpaceErrorSizeThreshold sets LowSpaceErrorSizeThreshold field to given value.

### HasLowSpaceErrorSizeThreshold

`func (o *GetMonitorProvider200Response) HasLowSpaceErrorSizeThreshold() bool`

HasLowSpaceErrorSizeThreshold returns a boolean if a field has been set.

### GetLowSpaceErrorPercentThreshold

`func (o *GetMonitorProvider200Response) GetLowSpaceErrorPercentThreshold() int32`

GetLowSpaceErrorPercentThreshold returns the LowSpaceErrorPercentThreshold field if non-nil, zero value otherwise.

### GetLowSpaceErrorPercentThresholdOk

`func (o *GetMonitorProvider200Response) GetLowSpaceErrorPercentThresholdOk() (*int32, bool)`

GetLowSpaceErrorPercentThresholdOk returns a tuple with the LowSpaceErrorPercentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceErrorPercentThreshold

`func (o *GetMonitorProvider200Response) SetLowSpaceErrorPercentThreshold(v int32)`

SetLowSpaceErrorPercentThreshold sets LowSpaceErrorPercentThreshold field to given value.

### HasLowSpaceErrorPercentThreshold

`func (o *GetMonitorProvider200Response) HasLowSpaceErrorPercentThreshold() bool`

HasLowSpaceErrorPercentThreshold returns a boolean if a field has been set.

### GetOutOfSpaceErrorSizeThreshold

`func (o *GetMonitorProvider200Response) GetOutOfSpaceErrorSizeThreshold() string`

GetOutOfSpaceErrorSizeThreshold returns the OutOfSpaceErrorSizeThreshold field if non-nil, zero value otherwise.

### GetOutOfSpaceErrorSizeThresholdOk

`func (o *GetMonitorProvider200Response) GetOutOfSpaceErrorSizeThresholdOk() (*string, bool)`

GetOutOfSpaceErrorSizeThresholdOk returns a tuple with the OutOfSpaceErrorSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfSpaceErrorSizeThreshold

`func (o *GetMonitorProvider200Response) SetOutOfSpaceErrorSizeThreshold(v string)`

SetOutOfSpaceErrorSizeThreshold sets OutOfSpaceErrorSizeThreshold field to given value.

### HasOutOfSpaceErrorSizeThreshold

`func (o *GetMonitorProvider200Response) HasOutOfSpaceErrorSizeThreshold() bool`

HasOutOfSpaceErrorSizeThreshold returns a boolean if a field has been set.

### GetOutOfSpaceErrorPercentThreshold

`func (o *GetMonitorProvider200Response) GetOutOfSpaceErrorPercentThreshold() int32`

GetOutOfSpaceErrorPercentThreshold returns the OutOfSpaceErrorPercentThreshold field if non-nil, zero value otherwise.

### GetOutOfSpaceErrorPercentThresholdOk

`func (o *GetMonitorProvider200Response) GetOutOfSpaceErrorPercentThresholdOk() (*int32, bool)`

GetOutOfSpaceErrorPercentThresholdOk returns a tuple with the OutOfSpaceErrorPercentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfSpaceErrorPercentThreshold

`func (o *GetMonitorProvider200Response) SetOutOfSpaceErrorPercentThreshold(v int32)`

SetOutOfSpaceErrorPercentThreshold sets OutOfSpaceErrorPercentThreshold field to given value.

### HasOutOfSpaceErrorPercentThreshold

`func (o *GetMonitorProvider200Response) HasOutOfSpaceErrorPercentThreshold() bool`

HasOutOfSpaceErrorPercentThreshold returns a boolean if a field has been set.

### GetAlertFrequency

`func (o *GetMonitorProvider200Response) GetAlertFrequency() string`

GetAlertFrequency returns the AlertFrequency field if non-nil, zero value otherwise.

### GetAlertFrequencyOk

`func (o *GetMonitorProvider200Response) GetAlertFrequencyOk() (*string, bool)`

GetAlertFrequencyOk returns a tuple with the AlertFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertFrequency

`func (o *GetMonitorProvider200Response) SetAlertFrequency(v string)`

SetAlertFrequency sets AlertFrequency field to given value.


### GetDiskDevices

`func (o *GetMonitorProvider200Response) GetDiskDevices() []string`

GetDiskDevices returns the DiskDevices field if non-nil, zero value otherwise.

### GetDiskDevicesOk

`func (o *GetMonitorProvider200Response) GetDiskDevicesOk() (*[]string, bool)`

GetDiskDevicesOk returns a tuple with the DiskDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDevices

`func (o *GetMonitorProvider200Response) SetDiskDevices(v []string)`

SetDiskDevices sets DiskDevices field to given value.

### HasDiskDevices

`func (o *GetMonitorProvider200Response) HasDiskDevices() bool`

HasDiskDevices returns a boolean if a field has been set.

### GetNetworkDevices

`func (o *GetMonitorProvider200Response) GetNetworkDevices() []string`

GetNetworkDevices returns the NetworkDevices field if non-nil, zero value otherwise.

### GetNetworkDevicesOk

`func (o *GetMonitorProvider200Response) GetNetworkDevicesOk() (*[]string, bool)`

GetNetworkDevicesOk returns a tuple with the NetworkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDevices

`func (o *GetMonitorProvider200Response) SetNetworkDevices(v []string)`

SetNetworkDevices sets NetworkDevices field to given value.

### HasNetworkDevices

`func (o *GetMonitorProvider200Response) HasNetworkDevices() bool`

HasNetworkDevices returns a boolean if a field has been set.

### GetSystemUtilizationMonitorLogDirectory

`func (o *GetMonitorProvider200Response) GetSystemUtilizationMonitorLogDirectory() string`

GetSystemUtilizationMonitorLogDirectory returns the SystemUtilizationMonitorLogDirectory field if non-nil, zero value otherwise.

### GetSystemUtilizationMonitorLogDirectoryOk

`func (o *GetMonitorProvider200Response) GetSystemUtilizationMonitorLogDirectoryOk() (*string, bool)`

GetSystemUtilizationMonitorLogDirectoryOk returns a tuple with the SystemUtilizationMonitorLogDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUtilizationMonitorLogDirectory

`func (o *GetMonitorProvider200Response) SetSystemUtilizationMonitorLogDirectory(v string)`

SetSystemUtilizationMonitorLogDirectory sets SystemUtilizationMonitorLogDirectory field to given value.


### GetId

`func (o *GetMonitorProvider200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMonitorProvider200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMonitorProvider200Response) SetId(v string)`

SetId sets Id field to given value.


### GetExtensionClass

`func (o *GetMonitorProvider200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *GetMonitorProvider200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *GetMonitorProvider200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *GetMonitorProvider200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *GetMonitorProvider200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *GetMonitorProvider200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *GetMonitorProvider200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetMeta

`func (o *GetMonitorProvider200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetMonitorProvider200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetMonitorProvider200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetMonitorProvider200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


