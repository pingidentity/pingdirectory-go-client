# DiskSpaceUsageMonitorProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumdiskSpaceUsageMonitorProviderSchemaUrn**](EnumdiskSpaceUsageMonitorProviderSchemaUrn.md) |  | 
**Id** | **string** | Name of the Monitor Provider | 
**LowSpaceWarningSizeThreshold** | Pointer to **string** | Specifies the low space warning threshold value as an absolute amount of space. If the amount of usable disk space drops below this amount, then the Directory Server will begin generating warning alert notifications. | [optional] 
**LowSpaceWarningPercentThreshold** | Pointer to **int64** | Specifies the low space warning threshold value as a percentage of total space. If the amount of usable disk space drops below this amount, then the Directory Server will begin generating warning alert notifications. | [optional] 
**LowSpaceErrorSizeThreshold** | Pointer to **string** | Specifies the low space error threshold value as an absolute amount of space. If the amount of usable disk space drops below this amount, then the Directory Server will start rejecting operations requested by non-root users. | [optional] 
**LowSpaceErrorPercentThreshold** | Pointer to **int64** | Specifies the low space error threshold value as a percentage of total space. If the amount of usable disk space drops below this amount, then the Directory Server will start rejecting operations requested by non-root users. | [optional] 
**OutOfSpaceErrorSizeThreshold** | Pointer to **string** | Specifies the out of space error threshold value as an absolute amount of space. If the amount of usable disk space drops below this amount, then the Directory Server will shut itself down to avoid problems that may occur from complete exhaustion of usable space. | [optional] 
**OutOfSpaceErrorPercentThreshold** | Pointer to **int64** | Specifies the out of space error threshold value as a percentage of total space. If the amount of usable disk space drops below this amount, then the Directory Server will shut itself down to avoid problems that may occur from complete exhaustion of usable space. | [optional] 
**AlertFrequency** | **string** | Specifies the length of time between administrative alerts generated in response to lack of usable disk space. Administrative alerts will be generated whenever the amount of usable space drops below any threshold, and they will also be generated at regular intervals as long as the amount of usable space remains below the threshold value. A value of zero indicates that alerts should only be generated when the amount of usable space drops below a configured threshold. | 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Monitor Provider is enabled for use. | 

## Methods

### NewDiskSpaceUsageMonitorProviderResponse

`func NewDiskSpaceUsageMonitorProviderResponse(schemas []EnumdiskSpaceUsageMonitorProviderSchemaUrn, id string, alertFrequency string, enabled bool, ) *DiskSpaceUsageMonitorProviderResponse`

NewDiskSpaceUsageMonitorProviderResponse instantiates a new DiskSpaceUsageMonitorProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskSpaceUsageMonitorProviderResponseWithDefaults

`func NewDiskSpaceUsageMonitorProviderResponseWithDefaults() *DiskSpaceUsageMonitorProviderResponse`

NewDiskSpaceUsageMonitorProviderResponseWithDefaults instantiates a new DiskSpaceUsageMonitorProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DiskSpaceUsageMonitorProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DiskSpaceUsageMonitorProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DiskSpaceUsageMonitorProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DiskSpaceUsageMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DiskSpaceUsageMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DiskSpaceUsageMonitorProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DiskSpaceUsageMonitorProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *DiskSpaceUsageMonitorProviderResponse) GetSchemas() []EnumdiskSpaceUsageMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetSchemasOk() (*[]EnumdiskSpaceUsageMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DiskSpaceUsageMonitorProviderResponse) SetSchemas(v []EnumdiskSpaceUsageMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *DiskSpaceUsageMonitorProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiskSpaceUsageMonitorProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLowSpaceWarningSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) GetLowSpaceWarningSizeThreshold() string`

GetLowSpaceWarningSizeThreshold returns the LowSpaceWarningSizeThreshold field if non-nil, zero value otherwise.

### GetLowSpaceWarningSizeThresholdOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetLowSpaceWarningSizeThresholdOk() (*string, bool)`

GetLowSpaceWarningSizeThresholdOk returns a tuple with the LowSpaceWarningSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceWarningSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) SetLowSpaceWarningSizeThreshold(v string)`

SetLowSpaceWarningSizeThreshold sets LowSpaceWarningSizeThreshold field to given value.

### HasLowSpaceWarningSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) HasLowSpaceWarningSizeThreshold() bool`

HasLowSpaceWarningSizeThreshold returns a boolean if a field has been set.

### GetLowSpaceWarningPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) GetLowSpaceWarningPercentThreshold() int64`

GetLowSpaceWarningPercentThreshold returns the LowSpaceWarningPercentThreshold field if non-nil, zero value otherwise.

### GetLowSpaceWarningPercentThresholdOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetLowSpaceWarningPercentThresholdOk() (*int64, bool)`

GetLowSpaceWarningPercentThresholdOk returns a tuple with the LowSpaceWarningPercentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceWarningPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) SetLowSpaceWarningPercentThreshold(v int64)`

SetLowSpaceWarningPercentThreshold sets LowSpaceWarningPercentThreshold field to given value.

### HasLowSpaceWarningPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) HasLowSpaceWarningPercentThreshold() bool`

HasLowSpaceWarningPercentThreshold returns a boolean if a field has been set.

### GetLowSpaceErrorSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) GetLowSpaceErrorSizeThreshold() string`

GetLowSpaceErrorSizeThreshold returns the LowSpaceErrorSizeThreshold field if non-nil, zero value otherwise.

### GetLowSpaceErrorSizeThresholdOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetLowSpaceErrorSizeThresholdOk() (*string, bool)`

GetLowSpaceErrorSizeThresholdOk returns a tuple with the LowSpaceErrorSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceErrorSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) SetLowSpaceErrorSizeThreshold(v string)`

SetLowSpaceErrorSizeThreshold sets LowSpaceErrorSizeThreshold field to given value.

### HasLowSpaceErrorSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) HasLowSpaceErrorSizeThreshold() bool`

HasLowSpaceErrorSizeThreshold returns a boolean if a field has been set.

### GetLowSpaceErrorPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) GetLowSpaceErrorPercentThreshold() int64`

GetLowSpaceErrorPercentThreshold returns the LowSpaceErrorPercentThreshold field if non-nil, zero value otherwise.

### GetLowSpaceErrorPercentThresholdOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetLowSpaceErrorPercentThresholdOk() (*int64, bool)`

GetLowSpaceErrorPercentThresholdOk returns a tuple with the LowSpaceErrorPercentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowSpaceErrorPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) SetLowSpaceErrorPercentThreshold(v int64)`

SetLowSpaceErrorPercentThreshold sets LowSpaceErrorPercentThreshold field to given value.

### HasLowSpaceErrorPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) HasLowSpaceErrorPercentThreshold() bool`

HasLowSpaceErrorPercentThreshold returns a boolean if a field has been set.

### GetOutOfSpaceErrorSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) GetOutOfSpaceErrorSizeThreshold() string`

GetOutOfSpaceErrorSizeThreshold returns the OutOfSpaceErrorSizeThreshold field if non-nil, zero value otherwise.

### GetOutOfSpaceErrorSizeThresholdOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetOutOfSpaceErrorSizeThresholdOk() (*string, bool)`

GetOutOfSpaceErrorSizeThresholdOk returns a tuple with the OutOfSpaceErrorSizeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfSpaceErrorSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) SetOutOfSpaceErrorSizeThreshold(v string)`

SetOutOfSpaceErrorSizeThreshold sets OutOfSpaceErrorSizeThreshold field to given value.

### HasOutOfSpaceErrorSizeThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) HasOutOfSpaceErrorSizeThreshold() bool`

HasOutOfSpaceErrorSizeThreshold returns a boolean if a field has been set.

### GetOutOfSpaceErrorPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) GetOutOfSpaceErrorPercentThreshold() int64`

GetOutOfSpaceErrorPercentThreshold returns the OutOfSpaceErrorPercentThreshold field if non-nil, zero value otherwise.

### GetOutOfSpaceErrorPercentThresholdOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetOutOfSpaceErrorPercentThresholdOk() (*int64, bool)`

GetOutOfSpaceErrorPercentThresholdOk returns a tuple with the OutOfSpaceErrorPercentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfSpaceErrorPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) SetOutOfSpaceErrorPercentThreshold(v int64)`

SetOutOfSpaceErrorPercentThreshold sets OutOfSpaceErrorPercentThreshold field to given value.

### HasOutOfSpaceErrorPercentThreshold

`func (o *DiskSpaceUsageMonitorProviderResponse) HasOutOfSpaceErrorPercentThreshold() bool`

HasOutOfSpaceErrorPercentThreshold returns a boolean if a field has been set.

### GetAlertFrequency

`func (o *DiskSpaceUsageMonitorProviderResponse) GetAlertFrequency() string`

GetAlertFrequency returns the AlertFrequency field if non-nil, zero value otherwise.

### GetAlertFrequencyOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetAlertFrequencyOk() (*string, bool)`

GetAlertFrequencyOk returns a tuple with the AlertFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertFrequency

`func (o *DiskSpaceUsageMonitorProviderResponse) SetAlertFrequency(v string)`

SetAlertFrequency sets AlertFrequency field to given value.


### GetDescription

`func (o *DiskSpaceUsageMonitorProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiskSpaceUsageMonitorProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiskSpaceUsageMonitorProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DiskSpaceUsageMonitorProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DiskSpaceUsageMonitorProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DiskSpaceUsageMonitorProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


