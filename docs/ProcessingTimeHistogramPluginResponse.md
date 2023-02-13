# ProcessingTimeHistogramPluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumprocessingTimeHistogramPluginSchemaUrn**](EnumprocessingTimeHistogramPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**HistogramCategoryBoundary** | **[]string** | Specifies the boundary values that will be used to separate the processing times into categories. Values should be specified as durations, and all values must be greater than zero. | 
**IncludeQueueTime** | Pointer to **bool** | Indicates whether operation processing times should include the time spent waiting on the work queue. This will only be available if the work queue is configured to monitor the queue time. | [optional] 
**SeparateMonitorEntryPerTrackedApplication** | Pointer to **bool** | When enabled, separate monitor entries will be included for each application defined in the Global Configuration&#39;s tracked-application property. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewProcessingTimeHistogramPluginResponse

`func NewProcessingTimeHistogramPluginResponse(schemas []EnumprocessingTimeHistogramPluginSchemaUrn, id string, pluginType []EnumpluginPluginTypeProp, histogramCategoryBoundary []string, enabled bool, ) *ProcessingTimeHistogramPluginResponse`

NewProcessingTimeHistogramPluginResponse instantiates a new ProcessingTimeHistogramPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessingTimeHistogramPluginResponseWithDefaults

`func NewProcessingTimeHistogramPluginResponseWithDefaults() *ProcessingTimeHistogramPluginResponse`

NewProcessingTimeHistogramPluginResponseWithDefaults instantiates a new ProcessingTimeHistogramPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ProcessingTimeHistogramPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ProcessingTimeHistogramPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ProcessingTimeHistogramPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ProcessingTimeHistogramPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ProcessingTimeHistogramPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ProcessingTimeHistogramPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ProcessingTimeHistogramPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ProcessingTimeHistogramPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *ProcessingTimeHistogramPluginResponse) GetSchemas() []EnumprocessingTimeHistogramPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ProcessingTimeHistogramPluginResponse) GetSchemasOk() (*[]EnumprocessingTimeHistogramPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ProcessingTimeHistogramPluginResponse) SetSchemas(v []EnumprocessingTimeHistogramPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ProcessingTimeHistogramPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcessingTimeHistogramPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcessingTimeHistogramPluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPluginType

`func (o *ProcessingTimeHistogramPluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *ProcessingTimeHistogramPluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *ProcessingTimeHistogramPluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetHistogramCategoryBoundary

`func (o *ProcessingTimeHistogramPluginResponse) GetHistogramCategoryBoundary() []string`

GetHistogramCategoryBoundary returns the HistogramCategoryBoundary field if non-nil, zero value otherwise.

### GetHistogramCategoryBoundaryOk

`func (o *ProcessingTimeHistogramPluginResponse) GetHistogramCategoryBoundaryOk() (*[]string, bool)`

GetHistogramCategoryBoundaryOk returns a tuple with the HistogramCategoryBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramCategoryBoundary

`func (o *ProcessingTimeHistogramPluginResponse) SetHistogramCategoryBoundary(v []string)`

SetHistogramCategoryBoundary sets HistogramCategoryBoundary field to given value.


### GetIncludeQueueTime

`func (o *ProcessingTimeHistogramPluginResponse) GetIncludeQueueTime() bool`

GetIncludeQueueTime returns the IncludeQueueTime field if non-nil, zero value otherwise.

### GetIncludeQueueTimeOk

`func (o *ProcessingTimeHistogramPluginResponse) GetIncludeQueueTimeOk() (*bool, bool)`

GetIncludeQueueTimeOk returns a tuple with the IncludeQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeQueueTime

`func (o *ProcessingTimeHistogramPluginResponse) SetIncludeQueueTime(v bool)`

SetIncludeQueueTime sets IncludeQueueTime field to given value.

### HasIncludeQueueTime

`func (o *ProcessingTimeHistogramPluginResponse) HasIncludeQueueTime() bool`

HasIncludeQueueTime returns a boolean if a field has been set.

### GetSeparateMonitorEntryPerTrackedApplication

`func (o *ProcessingTimeHistogramPluginResponse) GetSeparateMonitorEntryPerTrackedApplication() bool`

GetSeparateMonitorEntryPerTrackedApplication returns the SeparateMonitorEntryPerTrackedApplication field if non-nil, zero value otherwise.

### GetSeparateMonitorEntryPerTrackedApplicationOk

`func (o *ProcessingTimeHistogramPluginResponse) GetSeparateMonitorEntryPerTrackedApplicationOk() (*bool, bool)`

GetSeparateMonitorEntryPerTrackedApplicationOk returns a tuple with the SeparateMonitorEntryPerTrackedApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateMonitorEntryPerTrackedApplication

`func (o *ProcessingTimeHistogramPluginResponse) SetSeparateMonitorEntryPerTrackedApplication(v bool)`

SetSeparateMonitorEntryPerTrackedApplication sets SeparateMonitorEntryPerTrackedApplication field to given value.

### HasSeparateMonitorEntryPerTrackedApplication

`func (o *ProcessingTimeHistogramPluginResponse) HasSeparateMonitorEntryPerTrackedApplication() bool`

HasSeparateMonitorEntryPerTrackedApplication returns a boolean if a field has been set.

### GetDescription

`func (o *ProcessingTimeHistogramPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProcessingTimeHistogramPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProcessingTimeHistogramPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProcessingTimeHistogramPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ProcessingTimeHistogramPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProcessingTimeHistogramPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProcessingTimeHistogramPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *ProcessingTimeHistogramPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *ProcessingTimeHistogramPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *ProcessingTimeHistogramPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *ProcessingTimeHistogramPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


