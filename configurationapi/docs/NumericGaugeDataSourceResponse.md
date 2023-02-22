# NumericGaugeDataSourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Gauge Data Source | 
**Schemas** | [**[]EnumnumericGaugeDataSourceSchemaUrn**](EnumnumericGaugeDataSourceSchemaUrn.md) |  | 
**DataOrientation** | Pointer to [**EnumgaugeDataSourceDataOrientationProp**](EnumgaugeDataSourceDataOrientationProp.md) |  | [optional] 
**StatisticType** | [**EnumgaugeDataSourceStatisticTypeProp**](EnumgaugeDataSourceStatisticTypeProp.md) |  | 
**DivideValueBy** | Pointer to **float32** | An optional floating point value that can be used to scale the resulting value. | [optional] 
**DivideValueByAttribute** | Pointer to **string** | An optional property that can scale the resulting value by another attribute in the monitored entry. | [optional] 
**DivideValueByCounterAttribute** | Pointer to **string** | An optional property that can scale the resulting value by another attribute whose value represents a counter in the monitored entry. | [optional] 
**Description** | Pointer to **string** | A description for this Gauge Data Source | [optional] 
**AdditionalText** | Pointer to **string** | Additional information about the source of this data that is added to alerts sent as a result of gauges that use this Gauge Data Source. | [optional] 
**MonitorObjectclass** | **string** | The object class name of the monitor entries to examine for generating gauge data. | 
**MonitorAttribute** | **string** | Specifies the attribute on the monitor entries from which to derive the current gauge value. | 
**IncludeFilter** | Pointer to **string** | An optional LDAP filter that can be used restrict which monitor entries are used to compute output. | [optional] 
**ResourceAttribute** | Pointer to **string** | Specifies the attribute whose value is used to identify the specific resource being monitored (e.g. device name). | [optional] 
**ResourceType** | Pointer to **string** | A string indicating the type of resource being monitored. | [optional] 
**MinimumUpdateInterval** | Pointer to **string** | The minimum frequency with which gauges using this Gauge Data Source can be configured for update. In order to prevent undesirable side effects, some Gauge Data Sources may use this property to impose a higher bound on the update frequency of gauges. | [optional] 

## Methods

### NewNumericGaugeDataSourceResponse

`func NewNumericGaugeDataSourceResponse(id string, schemas []EnumnumericGaugeDataSourceSchemaUrn, statisticType EnumgaugeDataSourceStatisticTypeProp, monitorObjectclass string, monitorAttribute string, ) *NumericGaugeDataSourceResponse`

NewNumericGaugeDataSourceResponse instantiates a new NumericGaugeDataSourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumericGaugeDataSourceResponseWithDefaults

`func NewNumericGaugeDataSourceResponseWithDefaults() *NumericGaugeDataSourceResponse`

NewNumericGaugeDataSourceResponseWithDefaults instantiates a new NumericGaugeDataSourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *NumericGaugeDataSourceResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NumericGaugeDataSourceResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NumericGaugeDataSourceResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NumericGaugeDataSourceResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *NumericGaugeDataSourceResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *NumericGaugeDataSourceResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *NumericGaugeDataSourceResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *NumericGaugeDataSourceResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *NumericGaugeDataSourceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumericGaugeDataSourceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumericGaugeDataSourceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *NumericGaugeDataSourceResponse) GetSchemas() []EnumnumericGaugeDataSourceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NumericGaugeDataSourceResponse) GetSchemasOk() (*[]EnumnumericGaugeDataSourceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NumericGaugeDataSourceResponse) SetSchemas(v []EnumnumericGaugeDataSourceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDataOrientation

`func (o *NumericGaugeDataSourceResponse) GetDataOrientation() EnumgaugeDataSourceDataOrientationProp`

GetDataOrientation returns the DataOrientation field if non-nil, zero value otherwise.

### GetDataOrientationOk

`func (o *NumericGaugeDataSourceResponse) GetDataOrientationOk() (*EnumgaugeDataSourceDataOrientationProp, bool)`

GetDataOrientationOk returns a tuple with the DataOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOrientation

`func (o *NumericGaugeDataSourceResponse) SetDataOrientation(v EnumgaugeDataSourceDataOrientationProp)`

SetDataOrientation sets DataOrientation field to given value.

### HasDataOrientation

`func (o *NumericGaugeDataSourceResponse) HasDataOrientation() bool`

HasDataOrientation returns a boolean if a field has been set.

### GetStatisticType

`func (o *NumericGaugeDataSourceResponse) GetStatisticType() EnumgaugeDataSourceStatisticTypeProp`

GetStatisticType returns the StatisticType field if non-nil, zero value otherwise.

### GetStatisticTypeOk

`func (o *NumericGaugeDataSourceResponse) GetStatisticTypeOk() (*EnumgaugeDataSourceStatisticTypeProp, bool)`

GetStatisticTypeOk returns a tuple with the StatisticType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticType

`func (o *NumericGaugeDataSourceResponse) SetStatisticType(v EnumgaugeDataSourceStatisticTypeProp)`

SetStatisticType sets StatisticType field to given value.


### GetDivideValueBy

`func (o *NumericGaugeDataSourceResponse) GetDivideValueBy() float32`

GetDivideValueBy returns the DivideValueBy field if non-nil, zero value otherwise.

### GetDivideValueByOk

`func (o *NumericGaugeDataSourceResponse) GetDivideValueByOk() (*float32, bool)`

GetDivideValueByOk returns a tuple with the DivideValueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueBy

`func (o *NumericGaugeDataSourceResponse) SetDivideValueBy(v float32)`

SetDivideValueBy sets DivideValueBy field to given value.

### HasDivideValueBy

`func (o *NumericGaugeDataSourceResponse) HasDivideValueBy() bool`

HasDivideValueBy returns a boolean if a field has been set.

### GetDivideValueByAttribute

`func (o *NumericGaugeDataSourceResponse) GetDivideValueByAttribute() string`

GetDivideValueByAttribute returns the DivideValueByAttribute field if non-nil, zero value otherwise.

### GetDivideValueByAttributeOk

`func (o *NumericGaugeDataSourceResponse) GetDivideValueByAttributeOk() (*string, bool)`

GetDivideValueByAttributeOk returns a tuple with the DivideValueByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByAttribute

`func (o *NumericGaugeDataSourceResponse) SetDivideValueByAttribute(v string)`

SetDivideValueByAttribute sets DivideValueByAttribute field to given value.

### HasDivideValueByAttribute

`func (o *NumericGaugeDataSourceResponse) HasDivideValueByAttribute() bool`

HasDivideValueByAttribute returns a boolean if a field has been set.

### GetDivideValueByCounterAttribute

`func (o *NumericGaugeDataSourceResponse) GetDivideValueByCounterAttribute() string`

GetDivideValueByCounterAttribute returns the DivideValueByCounterAttribute field if non-nil, zero value otherwise.

### GetDivideValueByCounterAttributeOk

`func (o *NumericGaugeDataSourceResponse) GetDivideValueByCounterAttributeOk() (*string, bool)`

GetDivideValueByCounterAttributeOk returns a tuple with the DivideValueByCounterAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByCounterAttribute

`func (o *NumericGaugeDataSourceResponse) SetDivideValueByCounterAttribute(v string)`

SetDivideValueByCounterAttribute sets DivideValueByCounterAttribute field to given value.

### HasDivideValueByCounterAttribute

`func (o *NumericGaugeDataSourceResponse) HasDivideValueByCounterAttribute() bool`

HasDivideValueByCounterAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *NumericGaugeDataSourceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NumericGaugeDataSourceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NumericGaugeDataSourceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NumericGaugeDataSourceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NumericGaugeDataSourceResponse) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NumericGaugeDataSourceResponse) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NumericGaugeDataSourceResponse) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NumericGaugeDataSourceResponse) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetMonitorObjectclass

`func (o *NumericGaugeDataSourceResponse) GetMonitorObjectclass() string`

GetMonitorObjectclass returns the MonitorObjectclass field if non-nil, zero value otherwise.

### GetMonitorObjectclassOk

`func (o *NumericGaugeDataSourceResponse) GetMonitorObjectclassOk() (*string, bool)`

GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectclass

`func (o *NumericGaugeDataSourceResponse) SetMonitorObjectclass(v string)`

SetMonitorObjectclass sets MonitorObjectclass field to given value.


### GetMonitorAttribute

`func (o *NumericGaugeDataSourceResponse) GetMonitorAttribute() string`

GetMonitorAttribute returns the MonitorAttribute field if non-nil, zero value otherwise.

### GetMonitorAttributeOk

`func (o *NumericGaugeDataSourceResponse) GetMonitorAttributeOk() (*string, bool)`

GetMonitorAttributeOk returns a tuple with the MonitorAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAttribute

`func (o *NumericGaugeDataSourceResponse) SetMonitorAttribute(v string)`

SetMonitorAttribute sets MonitorAttribute field to given value.


### GetIncludeFilter

`func (o *NumericGaugeDataSourceResponse) GetIncludeFilter() string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *NumericGaugeDataSourceResponse) GetIncludeFilterOk() (*string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *NumericGaugeDataSourceResponse) SetIncludeFilter(v string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *NumericGaugeDataSourceResponse) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetResourceAttribute

`func (o *NumericGaugeDataSourceResponse) GetResourceAttribute() string`

GetResourceAttribute returns the ResourceAttribute field if non-nil, zero value otherwise.

### GetResourceAttributeOk

`func (o *NumericGaugeDataSourceResponse) GetResourceAttributeOk() (*string, bool)`

GetResourceAttributeOk returns a tuple with the ResourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttribute

`func (o *NumericGaugeDataSourceResponse) SetResourceAttribute(v string)`

SetResourceAttribute sets ResourceAttribute field to given value.

### HasResourceAttribute

`func (o *NumericGaugeDataSourceResponse) HasResourceAttribute() bool`

HasResourceAttribute returns a boolean if a field has been set.

### GetResourceType

`func (o *NumericGaugeDataSourceResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *NumericGaugeDataSourceResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *NumericGaugeDataSourceResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *NumericGaugeDataSourceResponse) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetMinimumUpdateInterval

`func (o *NumericGaugeDataSourceResponse) GetMinimumUpdateInterval() string`

GetMinimumUpdateInterval returns the MinimumUpdateInterval field if non-nil, zero value otherwise.

### GetMinimumUpdateIntervalOk

`func (o *NumericGaugeDataSourceResponse) GetMinimumUpdateIntervalOk() (*string, bool)`

GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUpdateInterval

`func (o *NumericGaugeDataSourceResponse) SetMinimumUpdateInterval(v string)`

SetMinimumUpdateInterval sets MinimumUpdateInterval field to given value.

### HasMinimumUpdateInterval

`func (o *NumericGaugeDataSourceResponse) HasMinimumUpdateInterval() bool`

HasMinimumUpdateInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


