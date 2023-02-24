# AddGaugeDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceName** | **string** | Name of the new Gauge Data Source | 
**Schemas** | [**[]EnumnumericGaugeDataSourceSchemaUrn**](EnumnumericGaugeDataSourceSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Gauge Data Source | [optional] 
**AdditionalText** | Pointer to **string** | Additional information about the source of this data that is added to alerts sent as a result of gauges that use this Gauge Data Source. | [optional] 
**MonitorObjectclass** | **string** | The object class name of the monitor entries to examine for generating gauge data. | 
**MonitorAttribute** | **string** | Specifies the attribute on the monitor entries from which to derive the current gauge value. | 
**IncludeFilter** | Pointer to **string** | An optional LDAP filter that can be used restrict which monitor entries are used to compute output. | [optional] 
**ResourceAttribute** | Pointer to **string** | Specifies the attribute whose value is used to identify the specific resource being monitored (e.g. device name). | [optional] 
**ResourceType** | Pointer to **string** | A string indicating the type of resource being monitored. | [optional] 
**MinimumUpdateInterval** | Pointer to **string** | The minimum frequency with which gauges using this Gauge Data Source can be configured for update. In order to prevent undesirable side effects, some Gauge Data Sources may use this property to impose a higher bound on the update frequency of gauges. | [optional] 
**DataOrientation** | Pointer to [**EnumgaugeDataSourceDataOrientationProp**](EnumgaugeDataSourceDataOrientationProp.md) |  | [optional] 
**StatisticType** | Pointer to [**EnumgaugeDataSourceStatisticTypeProp**](EnumgaugeDataSourceStatisticTypeProp.md) |  | [optional] 
**DivideValueBy** | Pointer to **float32** | An optional floating point value that can be used to scale the resulting value. | [optional] 
**DivideValueByAttribute** | Pointer to **string** | An optional property that can scale the resulting value by another attribute in the monitored entry. | [optional] 
**DivideValueByCounterAttribute** | Pointer to **string** | An optional property that can scale the resulting value by another attribute whose value represents a counter in the monitored entry. | [optional] 

## Methods

### NewAddGaugeDataSourceRequest

`func NewAddGaugeDataSourceRequest(sourceName string, schemas []EnumnumericGaugeDataSourceSchemaUrn, monitorObjectclass string, monitorAttribute string, ) *AddGaugeDataSourceRequest`

NewAddGaugeDataSourceRequest instantiates a new AddGaugeDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGaugeDataSourceRequestWithDefaults

`func NewAddGaugeDataSourceRequestWithDefaults() *AddGaugeDataSourceRequest`

NewAddGaugeDataSourceRequestWithDefaults instantiates a new AddGaugeDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceName

`func (o *AddGaugeDataSourceRequest) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AddGaugeDataSourceRequest) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AddGaugeDataSourceRequest) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetSchemas

`func (o *AddGaugeDataSourceRequest) GetSchemas() []EnumnumericGaugeDataSourceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGaugeDataSourceRequest) GetSchemasOk() (*[]EnumnumericGaugeDataSourceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGaugeDataSourceRequest) SetSchemas(v []EnumnumericGaugeDataSourceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddGaugeDataSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGaugeDataSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGaugeDataSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGaugeDataSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalText

`func (o *AddGaugeDataSourceRequest) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *AddGaugeDataSourceRequest) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *AddGaugeDataSourceRequest) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *AddGaugeDataSourceRequest) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetMonitorObjectclass

`func (o *AddGaugeDataSourceRequest) GetMonitorObjectclass() string`

GetMonitorObjectclass returns the MonitorObjectclass field if non-nil, zero value otherwise.

### GetMonitorObjectclassOk

`func (o *AddGaugeDataSourceRequest) GetMonitorObjectclassOk() (*string, bool)`

GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectclass

`func (o *AddGaugeDataSourceRequest) SetMonitorObjectclass(v string)`

SetMonitorObjectclass sets MonitorObjectclass field to given value.


### GetMonitorAttribute

`func (o *AddGaugeDataSourceRequest) GetMonitorAttribute() string`

GetMonitorAttribute returns the MonitorAttribute field if non-nil, zero value otherwise.

### GetMonitorAttributeOk

`func (o *AddGaugeDataSourceRequest) GetMonitorAttributeOk() (*string, bool)`

GetMonitorAttributeOk returns a tuple with the MonitorAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAttribute

`func (o *AddGaugeDataSourceRequest) SetMonitorAttribute(v string)`

SetMonitorAttribute sets MonitorAttribute field to given value.


### GetIncludeFilter

`func (o *AddGaugeDataSourceRequest) GetIncludeFilter() string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddGaugeDataSourceRequest) GetIncludeFilterOk() (*string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddGaugeDataSourceRequest) SetIncludeFilter(v string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddGaugeDataSourceRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetResourceAttribute

`func (o *AddGaugeDataSourceRequest) GetResourceAttribute() string`

GetResourceAttribute returns the ResourceAttribute field if non-nil, zero value otherwise.

### GetResourceAttributeOk

`func (o *AddGaugeDataSourceRequest) GetResourceAttributeOk() (*string, bool)`

GetResourceAttributeOk returns a tuple with the ResourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttribute

`func (o *AddGaugeDataSourceRequest) SetResourceAttribute(v string)`

SetResourceAttribute sets ResourceAttribute field to given value.

### HasResourceAttribute

`func (o *AddGaugeDataSourceRequest) HasResourceAttribute() bool`

HasResourceAttribute returns a boolean if a field has been set.

### GetResourceType

`func (o *AddGaugeDataSourceRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AddGaugeDataSourceRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AddGaugeDataSourceRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AddGaugeDataSourceRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetMinimumUpdateInterval

`func (o *AddGaugeDataSourceRequest) GetMinimumUpdateInterval() string`

GetMinimumUpdateInterval returns the MinimumUpdateInterval field if non-nil, zero value otherwise.

### GetMinimumUpdateIntervalOk

`func (o *AddGaugeDataSourceRequest) GetMinimumUpdateIntervalOk() (*string, bool)`

GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUpdateInterval

`func (o *AddGaugeDataSourceRequest) SetMinimumUpdateInterval(v string)`

SetMinimumUpdateInterval sets MinimumUpdateInterval field to given value.

### HasMinimumUpdateInterval

`func (o *AddGaugeDataSourceRequest) HasMinimumUpdateInterval() bool`

HasMinimumUpdateInterval returns a boolean if a field has been set.

### GetDataOrientation

`func (o *AddGaugeDataSourceRequest) GetDataOrientation() EnumgaugeDataSourceDataOrientationProp`

GetDataOrientation returns the DataOrientation field if non-nil, zero value otherwise.

### GetDataOrientationOk

`func (o *AddGaugeDataSourceRequest) GetDataOrientationOk() (*EnumgaugeDataSourceDataOrientationProp, bool)`

GetDataOrientationOk returns a tuple with the DataOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOrientation

`func (o *AddGaugeDataSourceRequest) SetDataOrientation(v EnumgaugeDataSourceDataOrientationProp)`

SetDataOrientation sets DataOrientation field to given value.

### HasDataOrientation

`func (o *AddGaugeDataSourceRequest) HasDataOrientation() bool`

HasDataOrientation returns a boolean if a field has been set.

### GetStatisticType

`func (o *AddGaugeDataSourceRequest) GetStatisticType() EnumgaugeDataSourceStatisticTypeProp`

GetStatisticType returns the StatisticType field if non-nil, zero value otherwise.

### GetStatisticTypeOk

`func (o *AddGaugeDataSourceRequest) GetStatisticTypeOk() (*EnumgaugeDataSourceStatisticTypeProp, bool)`

GetStatisticTypeOk returns a tuple with the StatisticType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticType

`func (o *AddGaugeDataSourceRequest) SetStatisticType(v EnumgaugeDataSourceStatisticTypeProp)`

SetStatisticType sets StatisticType field to given value.

### HasStatisticType

`func (o *AddGaugeDataSourceRequest) HasStatisticType() bool`

HasStatisticType returns a boolean if a field has been set.

### GetDivideValueBy

`func (o *AddGaugeDataSourceRequest) GetDivideValueBy() float32`

GetDivideValueBy returns the DivideValueBy field if non-nil, zero value otherwise.

### GetDivideValueByOk

`func (o *AddGaugeDataSourceRequest) GetDivideValueByOk() (*float32, bool)`

GetDivideValueByOk returns a tuple with the DivideValueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueBy

`func (o *AddGaugeDataSourceRequest) SetDivideValueBy(v float32)`

SetDivideValueBy sets DivideValueBy field to given value.

### HasDivideValueBy

`func (o *AddGaugeDataSourceRequest) HasDivideValueBy() bool`

HasDivideValueBy returns a boolean if a field has been set.

### GetDivideValueByAttribute

`func (o *AddGaugeDataSourceRequest) GetDivideValueByAttribute() string`

GetDivideValueByAttribute returns the DivideValueByAttribute field if non-nil, zero value otherwise.

### GetDivideValueByAttributeOk

`func (o *AddGaugeDataSourceRequest) GetDivideValueByAttributeOk() (*string, bool)`

GetDivideValueByAttributeOk returns a tuple with the DivideValueByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByAttribute

`func (o *AddGaugeDataSourceRequest) SetDivideValueByAttribute(v string)`

SetDivideValueByAttribute sets DivideValueByAttribute field to given value.

### HasDivideValueByAttribute

`func (o *AddGaugeDataSourceRequest) HasDivideValueByAttribute() bool`

HasDivideValueByAttribute returns a boolean if a field has been set.

### GetDivideValueByCounterAttribute

`func (o *AddGaugeDataSourceRequest) GetDivideValueByCounterAttribute() string`

GetDivideValueByCounterAttribute returns the DivideValueByCounterAttribute field if non-nil, zero value otherwise.

### GetDivideValueByCounterAttributeOk

`func (o *AddGaugeDataSourceRequest) GetDivideValueByCounterAttributeOk() (*string, bool)`

GetDivideValueByCounterAttributeOk returns a tuple with the DivideValueByCounterAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByCounterAttribute

`func (o *AddGaugeDataSourceRequest) SetDivideValueByCounterAttribute(v string)`

SetDivideValueByCounterAttribute sets DivideValueByCounterAttribute field to given value.

### HasDivideValueByCounterAttribute

`func (o *AddGaugeDataSourceRequest) HasDivideValueByCounterAttribute() bool`

HasDivideValueByCounterAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


