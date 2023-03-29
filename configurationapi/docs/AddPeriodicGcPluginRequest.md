# AddPeriodicGcPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumperiodicGcPluginSchemaUrn**](EnumperiodicGcPluginSchemaUrn.md) |  | 
**PluginType** | Pointer to [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) | Specifies the set of plug-in types for the plug-in, which specifies the times at which the plug-in is invoked. | [optional] 
**InvokeGCDayOfWeek** | Pointer to [**[]EnumpluginInvokeGCDayOfWeekProp**](EnumpluginInvokeGCDayOfWeekProp.md) | Specifies the days of the week which the Periodic GC Plugin should run. If no values are provided, then the plugin will run every day at the specified time. | [optional] 
**InvokeGCTimeUtc** | **[]string** | Specifies the times of the day at which garbage collection may be explicitly invoked. The times should be specified in \&quot;HH:MM\&quot; format, with \&quot;HH\&quot; as a two-digit numeric value between 00 and 23 representing the hour of the day, and MM as a two-digit numeric value between 00 and 59 representing the minute of the hour. All times will be interpreted in the UTC time zone. | 
**DelayAfterAlert** | Pointer to **string** | Specifies the length of time that the Directory Server should wait after sending the \&quot;force-gc-starting\&quot; administrative alert before actually invoking the garbage collection processing. | [optional] 
**DelayPostGC** | Pointer to **string** | Specifies the length of time that the Directory Server should wait after successfully completing the garbage collection processing, before removing the \&quot;force-gc-starting\&quot; administrative alert, which marks the server as unavailable. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewAddPeriodicGcPluginRequest

`func NewAddPeriodicGcPluginRequest(pluginName string, schemas []EnumperiodicGcPluginSchemaUrn, invokeGCTimeUtc []string, enabled bool, ) *AddPeriodicGcPluginRequest`

NewAddPeriodicGcPluginRequest instantiates a new AddPeriodicGcPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPeriodicGcPluginRequestWithDefaults

`func NewAddPeriodicGcPluginRequestWithDefaults() *AddPeriodicGcPluginRequest`

NewAddPeriodicGcPluginRequestWithDefaults instantiates a new AddPeriodicGcPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddPeriodicGcPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddPeriodicGcPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddPeriodicGcPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddPeriodicGcPluginRequest) GetSchemas() []EnumperiodicGcPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPeriodicGcPluginRequest) GetSchemasOk() (*[]EnumperiodicGcPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPeriodicGcPluginRequest) SetSchemas(v []EnumperiodicGcPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AddPeriodicGcPluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddPeriodicGcPluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddPeriodicGcPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.

### HasPluginType

`func (o *AddPeriodicGcPluginRequest) HasPluginType() bool`

HasPluginType returns a boolean if a field has been set.

### GetInvokeGCDayOfWeek

`func (o *AddPeriodicGcPluginRequest) GetInvokeGCDayOfWeek() []EnumpluginInvokeGCDayOfWeekProp`

GetInvokeGCDayOfWeek returns the InvokeGCDayOfWeek field if non-nil, zero value otherwise.

### GetInvokeGCDayOfWeekOk

`func (o *AddPeriodicGcPluginRequest) GetInvokeGCDayOfWeekOk() (*[]EnumpluginInvokeGCDayOfWeekProp, bool)`

GetInvokeGCDayOfWeekOk returns a tuple with the InvokeGCDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeGCDayOfWeek

`func (o *AddPeriodicGcPluginRequest) SetInvokeGCDayOfWeek(v []EnumpluginInvokeGCDayOfWeekProp)`

SetInvokeGCDayOfWeek sets InvokeGCDayOfWeek field to given value.

### HasInvokeGCDayOfWeek

`func (o *AddPeriodicGcPluginRequest) HasInvokeGCDayOfWeek() bool`

HasInvokeGCDayOfWeek returns a boolean if a field has been set.

### GetInvokeGCTimeUtc

`func (o *AddPeriodicGcPluginRequest) GetInvokeGCTimeUtc() []string`

GetInvokeGCTimeUtc returns the InvokeGCTimeUtc field if non-nil, zero value otherwise.

### GetInvokeGCTimeUtcOk

`func (o *AddPeriodicGcPluginRequest) GetInvokeGCTimeUtcOk() (*[]string, bool)`

GetInvokeGCTimeUtcOk returns a tuple with the InvokeGCTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeGCTimeUtc

`func (o *AddPeriodicGcPluginRequest) SetInvokeGCTimeUtc(v []string)`

SetInvokeGCTimeUtc sets InvokeGCTimeUtc field to given value.


### GetDelayAfterAlert

`func (o *AddPeriodicGcPluginRequest) GetDelayAfterAlert() string`

GetDelayAfterAlert returns the DelayAfterAlert field if non-nil, zero value otherwise.

### GetDelayAfterAlertOk

`func (o *AddPeriodicGcPluginRequest) GetDelayAfterAlertOk() (*string, bool)`

GetDelayAfterAlertOk returns a tuple with the DelayAfterAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayAfterAlert

`func (o *AddPeriodicGcPluginRequest) SetDelayAfterAlert(v string)`

SetDelayAfterAlert sets DelayAfterAlert field to given value.

### HasDelayAfterAlert

`func (o *AddPeriodicGcPluginRequest) HasDelayAfterAlert() bool`

HasDelayAfterAlert returns a boolean if a field has been set.

### GetDelayPostGC

`func (o *AddPeriodicGcPluginRequest) GetDelayPostGC() string`

GetDelayPostGC returns the DelayPostGC field if non-nil, zero value otherwise.

### GetDelayPostGCOk

`func (o *AddPeriodicGcPluginRequest) GetDelayPostGCOk() (*string, bool)`

GetDelayPostGCOk returns a tuple with the DelayPostGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayPostGC

`func (o *AddPeriodicGcPluginRequest) SetDelayPostGC(v string)`

SetDelayPostGC sets DelayPostGC field to given value.

### HasDelayPostGC

`func (o *AddPeriodicGcPluginRequest) HasDelayPostGC() bool`

HasDelayPostGC returns a boolean if a field has been set.

### GetDescription

`func (o *AddPeriodicGcPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPeriodicGcPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPeriodicGcPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPeriodicGcPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPeriodicGcPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPeriodicGcPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPeriodicGcPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddPeriodicGcPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddPeriodicGcPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddPeriodicGcPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddPeriodicGcPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


