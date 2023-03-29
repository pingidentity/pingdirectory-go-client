# AddGaugeDataSource200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Gauge Data Source | 
**Schemas** | [**[]EnumnumericGaugeDataSourceSchemaUrn**](EnumnumericGaugeDataSourceSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Gauge Data Source | [optional] 
**AdditionalText** | Pointer to **string** | Additional information about the source of this data that is added to alerts sent as a result of gauges that use this Gauge Data Source. | [optional] 
**MonitorObjectclass** | **string** | The object class name of the monitor entries to examine for generating gauge data. | 
**MonitorAttribute** | **string** | Specifies the attribute on the monitor entries from which to derive the current gauge value. | 
**IncludeFilter** | Pointer to **string** | An optional LDAP filter that can be used restrict which monitor entries are used to compute output. | [optional] 
**ResourceAttribute** | Pointer to **string** | Specifies the attribute whose value is used to identify the specific resource being monitored (e.g. device name). | [optional] 
**ResourceType** | Pointer to **string** | A string indicating the type of resource being monitored. | [optional] 
**MinimumUpdateInterval** | Pointer to **string** | The minimum frequency with which gauges using this Gauge Data Source can be configured for update. In order to prevent undesirable side effects, some Gauge Data Sources may use this property to impose a higher bound on the update frequency of gauges. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**DataOrientation** | Pointer to [**EnumgaugeDataSourceDataOrientationProp**](EnumgaugeDataSourceDataOrientationProp.md) |  | [optional] 
**StatisticType** | [**EnumgaugeDataSourceStatisticTypeProp**](EnumgaugeDataSourceStatisticTypeProp.md) |  | 
**DivideValueBy** | Pointer to **float32** | An optional floating point value that can be used to scale the resulting value. | [optional] 
**DivideValueByAttribute** | Pointer to **string** | An optional property that can scale the resulting value by another attribute in the monitored entry. | [optional] 
**DivideValueByCounterAttribute** | Pointer to **string** | An optional property that can scale the resulting value by another attribute whose value represents a counter in the monitored entry. | [optional] 

## Methods

### NewAddGaugeDataSource200Response

`func NewAddGaugeDataSource200Response(id string, schemas []EnumnumericGaugeDataSourceSchemaUrn, monitorObjectclass string, monitorAttribute string, statisticType EnumgaugeDataSourceStatisticTypeProp, ) *AddGaugeDataSource200Response`

NewAddGaugeDataSource200Response instantiates a new AddGaugeDataSource200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGaugeDataSource200ResponseWithDefaults

`func NewAddGaugeDataSource200ResponseWithDefaults() *AddGaugeDataSource200Response`

NewAddGaugeDataSource200ResponseWithDefaults instantiates a new AddGaugeDataSource200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddGaugeDataSource200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddGaugeDataSource200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddGaugeDataSource200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddGaugeDataSource200Response) GetSchemas() []EnumnumericGaugeDataSourceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGaugeDataSource200Response) GetSchemasOk() (*[]EnumnumericGaugeDataSourceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGaugeDataSource200Response) SetSchemas(v []EnumnumericGaugeDataSourceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddGaugeDataSource200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGaugeDataSource200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGaugeDataSource200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGaugeDataSource200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalText

`func (o *AddGaugeDataSource200Response) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *AddGaugeDataSource200Response) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *AddGaugeDataSource200Response) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *AddGaugeDataSource200Response) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetMonitorObjectclass

`func (o *AddGaugeDataSource200Response) GetMonitorObjectclass() string`

GetMonitorObjectclass returns the MonitorObjectclass field if non-nil, zero value otherwise.

### GetMonitorObjectclassOk

`func (o *AddGaugeDataSource200Response) GetMonitorObjectclassOk() (*string, bool)`

GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectclass

`func (o *AddGaugeDataSource200Response) SetMonitorObjectclass(v string)`

SetMonitorObjectclass sets MonitorObjectclass field to given value.


### GetMonitorAttribute

`func (o *AddGaugeDataSource200Response) GetMonitorAttribute() string`

GetMonitorAttribute returns the MonitorAttribute field if non-nil, zero value otherwise.

### GetMonitorAttributeOk

`func (o *AddGaugeDataSource200Response) GetMonitorAttributeOk() (*string, bool)`

GetMonitorAttributeOk returns a tuple with the MonitorAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAttribute

`func (o *AddGaugeDataSource200Response) SetMonitorAttribute(v string)`

SetMonitorAttribute sets MonitorAttribute field to given value.


### GetIncludeFilter

`func (o *AddGaugeDataSource200Response) GetIncludeFilter() string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddGaugeDataSource200Response) GetIncludeFilterOk() (*string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddGaugeDataSource200Response) SetIncludeFilter(v string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddGaugeDataSource200Response) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetResourceAttribute

`func (o *AddGaugeDataSource200Response) GetResourceAttribute() string`

GetResourceAttribute returns the ResourceAttribute field if non-nil, zero value otherwise.

### GetResourceAttributeOk

`func (o *AddGaugeDataSource200Response) GetResourceAttributeOk() (*string, bool)`

GetResourceAttributeOk returns a tuple with the ResourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttribute

`func (o *AddGaugeDataSource200Response) SetResourceAttribute(v string)`

SetResourceAttribute sets ResourceAttribute field to given value.

### HasResourceAttribute

`func (o *AddGaugeDataSource200Response) HasResourceAttribute() bool`

HasResourceAttribute returns a boolean if a field has been set.

### GetResourceType

`func (o *AddGaugeDataSource200Response) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AddGaugeDataSource200Response) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AddGaugeDataSource200Response) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AddGaugeDataSource200Response) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetMinimumUpdateInterval

`func (o *AddGaugeDataSource200Response) GetMinimumUpdateInterval() string`

GetMinimumUpdateInterval returns the MinimumUpdateInterval field if non-nil, zero value otherwise.

### GetMinimumUpdateIntervalOk

`func (o *AddGaugeDataSource200Response) GetMinimumUpdateIntervalOk() (*string, bool)`

GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUpdateInterval

`func (o *AddGaugeDataSource200Response) SetMinimumUpdateInterval(v string)`

SetMinimumUpdateInterval sets MinimumUpdateInterval field to given value.

### HasMinimumUpdateInterval

`func (o *AddGaugeDataSource200Response) HasMinimumUpdateInterval() bool`

HasMinimumUpdateInterval returns a boolean if a field has been set.

### GetMeta

`func (o *AddGaugeDataSource200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddGaugeDataSource200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddGaugeDataSource200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddGaugeDataSource200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddGaugeDataSource200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddGaugeDataSource200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddGaugeDataSource200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddGaugeDataSource200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetDataOrientation

`func (o *AddGaugeDataSource200Response) GetDataOrientation() EnumgaugeDataSourceDataOrientationProp`

GetDataOrientation returns the DataOrientation field if non-nil, zero value otherwise.

### GetDataOrientationOk

`func (o *AddGaugeDataSource200Response) GetDataOrientationOk() (*EnumgaugeDataSourceDataOrientationProp, bool)`

GetDataOrientationOk returns a tuple with the DataOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOrientation

`func (o *AddGaugeDataSource200Response) SetDataOrientation(v EnumgaugeDataSourceDataOrientationProp)`

SetDataOrientation sets DataOrientation field to given value.

### HasDataOrientation

`func (o *AddGaugeDataSource200Response) HasDataOrientation() bool`

HasDataOrientation returns a boolean if a field has been set.

### GetStatisticType

`func (o *AddGaugeDataSource200Response) GetStatisticType() EnumgaugeDataSourceStatisticTypeProp`

GetStatisticType returns the StatisticType field if non-nil, zero value otherwise.

### GetStatisticTypeOk

`func (o *AddGaugeDataSource200Response) GetStatisticTypeOk() (*EnumgaugeDataSourceStatisticTypeProp, bool)`

GetStatisticTypeOk returns a tuple with the StatisticType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticType

`func (o *AddGaugeDataSource200Response) SetStatisticType(v EnumgaugeDataSourceStatisticTypeProp)`

SetStatisticType sets StatisticType field to given value.


### GetDivideValueBy

`func (o *AddGaugeDataSource200Response) GetDivideValueBy() float32`

GetDivideValueBy returns the DivideValueBy field if non-nil, zero value otherwise.

### GetDivideValueByOk

`func (o *AddGaugeDataSource200Response) GetDivideValueByOk() (*float32, bool)`

GetDivideValueByOk returns a tuple with the DivideValueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueBy

`func (o *AddGaugeDataSource200Response) SetDivideValueBy(v float32)`

SetDivideValueBy sets DivideValueBy field to given value.

### HasDivideValueBy

`func (o *AddGaugeDataSource200Response) HasDivideValueBy() bool`

HasDivideValueBy returns a boolean if a field has been set.

### GetDivideValueByAttribute

`func (o *AddGaugeDataSource200Response) GetDivideValueByAttribute() string`

GetDivideValueByAttribute returns the DivideValueByAttribute field if non-nil, zero value otherwise.

### GetDivideValueByAttributeOk

`func (o *AddGaugeDataSource200Response) GetDivideValueByAttributeOk() (*string, bool)`

GetDivideValueByAttributeOk returns a tuple with the DivideValueByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByAttribute

`func (o *AddGaugeDataSource200Response) SetDivideValueByAttribute(v string)`

SetDivideValueByAttribute sets DivideValueByAttribute field to given value.

### HasDivideValueByAttribute

`func (o *AddGaugeDataSource200Response) HasDivideValueByAttribute() bool`

HasDivideValueByAttribute returns a boolean if a field has been set.

### GetDivideValueByCounterAttribute

`func (o *AddGaugeDataSource200Response) GetDivideValueByCounterAttribute() string`

GetDivideValueByCounterAttribute returns the DivideValueByCounterAttribute field if non-nil, zero value otherwise.

### GetDivideValueByCounterAttributeOk

`func (o *AddGaugeDataSource200Response) GetDivideValueByCounterAttributeOk() (*string, bool)`

GetDivideValueByCounterAttributeOk returns a tuple with the DivideValueByCounterAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByCounterAttribute

`func (o *AddGaugeDataSource200Response) SetDivideValueByCounterAttribute(v string)`

SetDivideValueByCounterAttribute sets DivideValueByCounterAttribute field to given value.

### HasDivideValueByCounterAttribute

`func (o *AddGaugeDataSource200Response) HasDivideValueByCounterAttribute() bool`

HasDivideValueByCounterAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


