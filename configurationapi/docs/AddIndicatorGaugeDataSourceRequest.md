# AddIndicatorGaugeDataSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceName** | **string** | Name of the new Gauge Data Source | 
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

### NewAddIndicatorGaugeDataSourceRequest

`func NewAddIndicatorGaugeDataSourceRequest(sourceName string, schemas []EnumindicatorGaugeDataSourceSchemaUrn, monitorObjectclass string, monitorAttribute string, ) *AddIndicatorGaugeDataSourceRequest`

NewAddIndicatorGaugeDataSourceRequest instantiates a new AddIndicatorGaugeDataSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIndicatorGaugeDataSourceRequestWithDefaults

`func NewAddIndicatorGaugeDataSourceRequestWithDefaults() *AddIndicatorGaugeDataSourceRequest`

NewAddIndicatorGaugeDataSourceRequestWithDefaults instantiates a new AddIndicatorGaugeDataSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceName

`func (o *AddIndicatorGaugeDataSourceRequest) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *AddIndicatorGaugeDataSourceRequest) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetSchemas

`func (o *AddIndicatorGaugeDataSourceRequest) GetSchemas() []EnumindicatorGaugeDataSourceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetSchemasOk() (*[]EnumindicatorGaugeDataSourceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddIndicatorGaugeDataSourceRequest) SetSchemas(v []EnumindicatorGaugeDataSourceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddIndicatorGaugeDataSourceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddIndicatorGaugeDataSourceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddIndicatorGaugeDataSourceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAdditionalText

`func (o *AddIndicatorGaugeDataSourceRequest) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *AddIndicatorGaugeDataSourceRequest) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *AddIndicatorGaugeDataSourceRequest) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetMonitorObjectclass

`func (o *AddIndicatorGaugeDataSourceRequest) GetMonitorObjectclass() string`

GetMonitorObjectclass returns the MonitorObjectclass field if non-nil, zero value otherwise.

### GetMonitorObjectclassOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetMonitorObjectclassOk() (*string, bool)`

GetMonitorObjectclassOk returns a tuple with the MonitorObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorObjectclass

`func (o *AddIndicatorGaugeDataSourceRequest) SetMonitorObjectclass(v string)`

SetMonitorObjectclass sets MonitorObjectclass field to given value.


### GetMonitorAttribute

`func (o *AddIndicatorGaugeDataSourceRequest) GetMonitorAttribute() string`

GetMonitorAttribute returns the MonitorAttribute field if non-nil, zero value otherwise.

### GetMonitorAttributeOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetMonitorAttributeOk() (*string, bool)`

GetMonitorAttributeOk returns a tuple with the MonitorAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAttribute

`func (o *AddIndicatorGaugeDataSourceRequest) SetMonitorAttribute(v string)`

SetMonitorAttribute sets MonitorAttribute field to given value.


### GetIncludeFilter

`func (o *AddIndicatorGaugeDataSourceRequest) GetIncludeFilter() string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetIncludeFilterOk() (*string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddIndicatorGaugeDataSourceRequest) SetIncludeFilter(v string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddIndicatorGaugeDataSourceRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetResourceAttribute

`func (o *AddIndicatorGaugeDataSourceRequest) GetResourceAttribute() string`

GetResourceAttribute returns the ResourceAttribute field if non-nil, zero value otherwise.

### GetResourceAttributeOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetResourceAttributeOk() (*string, bool)`

GetResourceAttributeOk returns a tuple with the ResourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttribute

`func (o *AddIndicatorGaugeDataSourceRequest) SetResourceAttribute(v string)`

SetResourceAttribute sets ResourceAttribute field to given value.

### HasResourceAttribute

`func (o *AddIndicatorGaugeDataSourceRequest) HasResourceAttribute() bool`

HasResourceAttribute returns a boolean if a field has been set.

### GetResourceType

`func (o *AddIndicatorGaugeDataSourceRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *AddIndicatorGaugeDataSourceRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *AddIndicatorGaugeDataSourceRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetMinimumUpdateInterval

`func (o *AddIndicatorGaugeDataSourceRequest) GetMinimumUpdateInterval() string`

GetMinimumUpdateInterval returns the MinimumUpdateInterval field if non-nil, zero value otherwise.

### GetMinimumUpdateIntervalOk

`func (o *AddIndicatorGaugeDataSourceRequest) GetMinimumUpdateIntervalOk() (*string, bool)`

GetMinimumUpdateIntervalOk returns a tuple with the MinimumUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumUpdateInterval

`func (o *AddIndicatorGaugeDataSourceRequest) SetMinimumUpdateInterval(v string)`

SetMinimumUpdateInterval sets MinimumUpdateInterval field to given value.

### HasMinimumUpdateInterval

`func (o *AddIndicatorGaugeDataSourceRequest) HasMinimumUpdateInterval() bool`

HasMinimumUpdateInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


