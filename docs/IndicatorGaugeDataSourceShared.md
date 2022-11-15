# IndicatorGaugeDataSourceShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumindicatorGaugeDataSourceSchemaUrn**](EnumindicatorGaugeDataSourceSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Gauge Data Source | [optional] 
**AdditionalText** | Pointer to **string** | Additional information about the source of this data that is added to alerts sent as a result of gauges that use this Gauge Data Source. | [optional] 
**MonitorObjectclass** | **string** | The object class name of the monitor entries to examine for generating gauge data. | 
**MonitorAttribute** | **string** | Specifies the attribute on the monitor entries from which to derive the current gauge value. | 
**IncludeFilter** | Pointer to **string** | An optional LDAP filter that can be used restrict which monitor entries are used to compute output. | [optional] 
**ResourceAttribute** | Pointer to **string** | Specifies the attribute whose value is used to identify the specific resource being monitored (e.g. device name). | [optional] 
**ResourceType** | Pointer to **string** | A string indicating the type of resource being monitored. | [optional] 
**MinimumUpdateInterval** | Pointer to **string** | The minimum frequency with which gauges using this Gauge Data Source can be configured for update. In order to prevent undesirable side effects, some Gauge Data Sources may use this property to impose a higher bound on the update frequency of gauges. | [optional] 

## Methods

### NewIndicatorGaugeDataSourceShared

`func NewIndicatorGaugeDataSourceShared(schemas []EnumindicatorGaugeDataSourceSchemaUrn, monitorObjectclass string, monitorAttribute string, ) *IndicatorGaugeDataSourceShared`

NewIndicatorGaugeDataSourceShared instantiates a new IndicatorGaugeDataSourceShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorGaugeDataSourceSharedWithDefaults

`func NewIndicatorGaugeDataSourceSharedWithDefaults() *IndicatorGaugeDataSourceShared`

NewIndicatorGaugeDataSourceSharedWithDefaults instantiates a new IndicatorGaugeDataSourceShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *IndicatorGaugeDataSourceShared) GetSchemas() []EnumindicatorGaugeDataSourceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *IndicatorGaugeDataSourceShared) GetSchemasOk() (*[]EnumindicatorGaugeDataSourceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *IndicatorGaugeDataSourceShared) SetSchemas(v []EnumindicatorGaugeDataSourceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *IndicatorGaugeDataSourceShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IndicatorGaugeDataSourceShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IndicatorGaugeDataSourceShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IndicatorGaugeDataSourceShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalText

`func (o *IndicatorGaugeDataSourceShared) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *IndicatorGaugeDataSourceShared) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *IndicatorGaugeDataSourceShared) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *IndicatorGaugeDataSourceShared) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetMonitorObjectclass

`func (o *IndicatorGaugeDataSourceShared) GetMonitorObjectclass() string`

GetMonitorObjectclass returns the MonitorObjectclass field if non-nil, zero value otherwise.

### GetMonitorObjectclassOk

`func (o *IndicatorGaugeDataSourceShared) GetMonitorObjectclassOk() (*string, bool)`

GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectclass

`func (o *IndicatorGaugeDataSourceShared) SetMonitorObjectclass(v string)`

SetMonitorObjectclass sets MonitorObjectclass field to given value.


### GetMonitorAttribute

`func (o *IndicatorGaugeDataSourceShared) GetMonitorAttribute() string`

GetMonitorAttribute returns the MonitorAttribute field if non-nil, zero value otherwise.

### GetMonitorAttributeOk

`func (o *IndicatorGaugeDataSourceShared) GetMonitorAttributeOk() (*string, bool)`

GetMonitorAttributeOk returns a tuple with the MonitorAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAttribute

`func (o *IndicatorGaugeDataSourceShared) SetMonitorAttribute(v string)`

SetMonitorAttribute sets MonitorAttribute field to given value.


### GetIncludeFilter

`func (o *IndicatorGaugeDataSourceShared) GetIncludeFilter() string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *IndicatorGaugeDataSourceShared) GetIncludeFilterOk() (*string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *IndicatorGaugeDataSourceShared) SetIncludeFilter(v string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *IndicatorGaugeDataSourceShared) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetResourceAttribute

`func (o *IndicatorGaugeDataSourceShared) GetResourceAttribute() string`

GetResourceAttribute returns the ResourceAttribute field if non-nil, zero value otherwise.

### GetResourceAttributeOk

`func (o *IndicatorGaugeDataSourceShared) GetResourceAttributeOk() (*string, bool)`

GetResourceAttributeOk returns a tuple with the ResourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttribute

`func (o *IndicatorGaugeDataSourceShared) SetResourceAttribute(v string)`

SetResourceAttribute sets ResourceAttribute field to given value.

### HasResourceAttribute

`func (o *IndicatorGaugeDataSourceShared) HasResourceAttribute() bool`

HasResourceAttribute returns a boolean if a field has been set.

### GetResourceType

`func (o *IndicatorGaugeDataSourceShared) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *IndicatorGaugeDataSourceShared) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *IndicatorGaugeDataSourceShared) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *IndicatorGaugeDataSourceShared) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetMinimumUpdateInterval

`func (o *IndicatorGaugeDataSourceShared) GetMinimumUpdateInterval() string`

GetMinimumUpdateInterval returns the MinimumUpdateInterval field if non-nil, zero value otherwise.

### GetMinimumUpdateIntervalOk

`func (o *IndicatorGaugeDataSourceShared) GetMinimumUpdateIntervalOk() (*string, bool)`

GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUpdateInterval

`func (o *IndicatorGaugeDataSourceShared) SetMinimumUpdateInterval(v string)`

SetMinimumUpdateInterval sets MinimumUpdateInterval field to given value.

### HasMinimumUpdateInterval

`func (o *IndicatorGaugeDataSourceShared) HasMinimumUpdateInterval() bool`

HasMinimumUpdateInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


