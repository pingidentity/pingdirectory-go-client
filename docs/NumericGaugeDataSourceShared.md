# NumericGaugeDataSourceShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewNumericGaugeDataSourceShared

`func NewNumericGaugeDataSourceShared(schemas []EnumnumericGaugeDataSourceSchemaUrn, statisticType EnumgaugeDataSourceStatisticTypeProp, monitorObjectclass string, monitorAttribute string, ) *NumericGaugeDataSourceShared`

NewNumericGaugeDataSourceShared instantiates a new NumericGaugeDataSourceShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumericGaugeDataSourceSharedWithDefaults

`func NewNumericGaugeDataSourceSharedWithDefaults() *NumericGaugeDataSourceShared`

NewNumericGaugeDataSourceSharedWithDefaults instantiates a new NumericGaugeDataSourceShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *NumericGaugeDataSourceShared) GetSchemas() []EnumnumericGaugeDataSourceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NumericGaugeDataSourceShared) GetSchemasOk() (*[]EnumnumericGaugeDataSourceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NumericGaugeDataSourceShared) SetSchemas(v []EnumnumericGaugeDataSourceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDataOrientation

`func (o *NumericGaugeDataSourceShared) GetDataOrientation() EnumgaugeDataSourceDataOrientationProp`

GetDataOrientation returns the DataOrientation field if non-nil, zero value otherwise.

### GetDataOrientationOk

`func (o *NumericGaugeDataSourceShared) GetDataOrientationOk() (*EnumgaugeDataSourceDataOrientationProp, bool)`

GetDataOrientationOk returns a tuple with the DataOrientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataOrientation

`func (o *NumericGaugeDataSourceShared) SetDataOrientation(v EnumgaugeDataSourceDataOrientationProp)`

SetDataOrientation sets DataOrientation field to given value.

### HasDataOrientation

`func (o *NumericGaugeDataSourceShared) HasDataOrientation() bool`

HasDataOrientation returns a boolean if a field has been set.

### GetStatisticType

`func (o *NumericGaugeDataSourceShared) GetStatisticType() EnumgaugeDataSourceStatisticTypeProp`

GetStatisticType returns the StatisticType field if non-nil, zero value otherwise.

### GetStatisticTypeOk

`func (o *NumericGaugeDataSourceShared) GetStatisticTypeOk() (*EnumgaugeDataSourceStatisticTypeProp, bool)`

GetStatisticTypeOk returns a tuple with the StatisticType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticType

`func (o *NumericGaugeDataSourceShared) SetStatisticType(v EnumgaugeDataSourceStatisticTypeProp)`

SetStatisticType sets StatisticType field to given value.


### GetDivideValueBy

`func (o *NumericGaugeDataSourceShared) GetDivideValueBy() float32`

GetDivideValueBy returns the DivideValueBy field if non-nil, zero value otherwise.

### GetDivideValueByOk

`func (o *NumericGaugeDataSourceShared) GetDivideValueByOk() (*float32, bool)`

GetDivideValueByOk returns a tuple with the DivideValueBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueBy

`func (o *NumericGaugeDataSourceShared) SetDivideValueBy(v float32)`

SetDivideValueBy sets DivideValueBy field to given value.

### HasDivideValueBy

`func (o *NumericGaugeDataSourceShared) HasDivideValueBy() bool`

HasDivideValueBy returns a boolean if a field has been set.

### GetDivideValueByAttribute

`func (o *NumericGaugeDataSourceShared) GetDivideValueByAttribute() string`

GetDivideValueByAttribute returns the DivideValueByAttribute field if non-nil, zero value otherwise.

### GetDivideValueByAttributeOk

`func (o *NumericGaugeDataSourceShared) GetDivideValueByAttributeOk() (*string, bool)`

GetDivideValueByAttributeOk returns a tuple with the DivideValueByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByAttribute

`func (o *NumericGaugeDataSourceShared) SetDivideValueByAttribute(v string)`

SetDivideValueByAttribute sets DivideValueByAttribute field to given value.

### HasDivideValueByAttribute

`func (o *NumericGaugeDataSourceShared) HasDivideValueByAttribute() bool`

HasDivideValueByAttribute returns a boolean if a field has been set.

### GetDivideValueByCounterAttribute

`func (o *NumericGaugeDataSourceShared) GetDivideValueByCounterAttribute() string`

GetDivideValueByCounterAttribute returns the DivideValueByCounterAttribute field if non-nil, zero value otherwise.

### GetDivideValueByCounterAttributeOk

`func (o *NumericGaugeDataSourceShared) GetDivideValueByCounterAttributeOk() (*string, bool)`

GetDivideValueByCounterAttributeOk returns a tuple with the DivideValueByCounterAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivideValueByCounterAttribute

`func (o *NumericGaugeDataSourceShared) SetDivideValueByCounterAttribute(v string)`

SetDivideValueByCounterAttribute sets DivideValueByCounterAttribute field to given value.

### HasDivideValueByCounterAttribute

`func (o *NumericGaugeDataSourceShared) HasDivideValueByCounterAttribute() bool`

HasDivideValueByCounterAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *NumericGaugeDataSourceShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NumericGaugeDataSourceShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NumericGaugeDataSourceShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NumericGaugeDataSourceShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NumericGaugeDataSourceShared) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NumericGaugeDataSourceShared) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NumericGaugeDataSourceShared) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NumericGaugeDataSourceShared) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetMonitorObjectclass

`func (o *NumericGaugeDataSourceShared) GetMonitorObjectclass() string`

GetMonitorObjectclass returns the MonitorObjectclass field if non-nil, zero value otherwise.

### GetMonitorObjectclassOk

`func (o *NumericGaugeDataSourceShared) GetMonitorObjectclassOk() (*string, bool)`

GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectclass

`func (o *NumericGaugeDataSourceShared) SetMonitorObjectclass(v string)`

SetMonitorObjectclass sets MonitorObjectclass field to given value.


### GetMonitorAttribute

`func (o *NumericGaugeDataSourceShared) GetMonitorAttribute() string`

GetMonitorAttribute returns the MonitorAttribute field if non-nil, zero value otherwise.

### GetMonitorAttributeOk

`func (o *NumericGaugeDataSourceShared) GetMonitorAttributeOk() (*string, bool)`

GetMonitorAttributeOk returns a tuple with the MonitorAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAttribute

`func (o *NumericGaugeDataSourceShared) SetMonitorAttribute(v string)`

SetMonitorAttribute sets MonitorAttribute field to given value.


### GetIncludeFilter

`func (o *NumericGaugeDataSourceShared) GetIncludeFilter() string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *NumericGaugeDataSourceShared) GetIncludeFilterOk() (*string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *NumericGaugeDataSourceShared) SetIncludeFilter(v string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *NumericGaugeDataSourceShared) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetResourceAttribute

`func (o *NumericGaugeDataSourceShared) GetResourceAttribute() string`

GetResourceAttribute returns the ResourceAttribute field if non-nil, zero value otherwise.

### GetResourceAttributeOk

`func (o *NumericGaugeDataSourceShared) GetResourceAttributeOk() (*string, bool)`

GetResourceAttributeOk returns a tuple with the ResourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttribute

`func (o *NumericGaugeDataSourceShared) SetResourceAttribute(v string)`

SetResourceAttribute sets ResourceAttribute field to given value.

### HasResourceAttribute

`func (o *NumericGaugeDataSourceShared) HasResourceAttribute() bool`

HasResourceAttribute returns a boolean if a field has been set.

### GetResourceType

`func (o *NumericGaugeDataSourceShared) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *NumericGaugeDataSourceShared) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *NumericGaugeDataSourceShared) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *NumericGaugeDataSourceShared) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetMinimumUpdateInterval

`func (o *NumericGaugeDataSourceShared) GetMinimumUpdateInterval() string`

GetMinimumUpdateInterval returns the MinimumUpdateInterval field if non-nil, zero value otherwise.

### GetMinimumUpdateIntervalOk

`func (o *NumericGaugeDataSourceShared) GetMinimumUpdateIntervalOk() (*string, bool)`

GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUpdateInterval

`func (o *NumericGaugeDataSourceShared) SetMinimumUpdateInterval(v string)`

SetMinimumUpdateInterval sets MinimumUpdateInterval field to given value.

### HasMinimumUpdateInterval

`func (o *NumericGaugeDataSourceShared) HasMinimumUpdateInterval() bool`

HasMinimumUpdateInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


