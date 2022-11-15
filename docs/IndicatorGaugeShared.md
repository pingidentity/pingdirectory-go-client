# IndicatorGaugeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumindicatorGaugeSchemaUrn**](EnumindicatorGaugeSchemaUrn.md) |  | 
**GaugeDataSource** | **string** | Specifies the source of data to use in determining this Indicator Gauge&#39;s severity and status. | 
**CriticalValue** | Pointer to **string** | A regular expression pattern that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be critical. | [optional] 
**MajorValue** | Pointer to **string** | A regular expression pattern that is used to determine whether the current monitored value indicates this gauge&#39;s severity will be &#39;major&#39;. | [optional] 
**MinorValue** | Pointer to **string** | A regular expression pattern that is used to determine whether the current monitored value indicates this gauge&#39;s severity will be &#39;minor&#39;. | [optional] 
**WarningValue** | Pointer to **string** | A regular expression pattern that is used to determine whether the current monitored value indicates this gauge&#39;s severity will be &#39;warning&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Gauge | [optional] 
**Enabled** | **bool** | Indicates whether this Gauge is enabled. | 
**OverrideSeverity** | Pointer to [**EnumgaugeOverrideSeverityProp**](EnumgaugeOverrideSeverityProp.md) |  | [optional] 
**AlertLevel** | Pointer to [**EnumgaugeAlertLevelProp**](EnumgaugeAlertLevelProp.md) |  | [optional] 
**UpdateInterval** | Pointer to **string** | The frequency with which this Gauge is updated. | [optional] 
**SamplesPerUpdateInterval** | Pointer to **int32** | Indicates the number of times the monitor data source value will be collected during the update interval. | [optional] 
**IncludeResource** | Pointer to **[]string** |  | [optional] 
**ExcludeResource** | Pointer to **[]string** |  | [optional] 
**ServerUnavailableSeverityLevel** | Pointer to [**EnumgaugeServerUnavailableSeverityLevelProp**](EnumgaugeServerUnavailableSeverityLevelProp.md) |  | [optional] 
**ServerDegradedSeverityLevel** | Pointer to [**EnumgaugeServerDegradedSeverityLevelProp**](EnumgaugeServerDegradedSeverityLevelProp.md) |  | [optional] 

## Methods

### NewIndicatorGaugeShared

`func NewIndicatorGaugeShared(schemas []EnumindicatorGaugeSchemaUrn, gaugeDataSource string, enabled bool, ) *IndicatorGaugeShared`

NewIndicatorGaugeShared instantiates a new IndicatorGaugeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorGaugeSharedWithDefaults

`func NewIndicatorGaugeSharedWithDefaults() *IndicatorGaugeShared`

NewIndicatorGaugeSharedWithDefaults instantiates a new IndicatorGaugeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *IndicatorGaugeShared) GetSchemas() []EnumindicatorGaugeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *IndicatorGaugeShared) GetSchemasOk() (*[]EnumindicatorGaugeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *IndicatorGaugeShared) SetSchemas(v []EnumindicatorGaugeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGaugeDataSource

`func (o *IndicatorGaugeShared) GetGaugeDataSource() string`

GetGaugeDataSource returns the GaugeDataSource field if non-nil, zero value otherwise.

### GetGaugeDataSourceOk

`func (o *IndicatorGaugeShared) GetGaugeDataSourceOk() (*string, bool)`

GetGaugeDataSourceOk returns a tuple with the GaugeDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeDataSource

`func (o *IndicatorGaugeShared) SetGaugeDataSource(v string)`

SetGaugeDataSource sets GaugeDataSource field to given value.


### GetCriticalValue

`func (o *IndicatorGaugeShared) GetCriticalValue() string`

GetCriticalValue returns the CriticalValue field if non-nil, zero value otherwise.

### GetCriticalValueOk

`func (o *IndicatorGaugeShared) GetCriticalValueOk() (*string, bool)`

GetCriticalValueOk returns a tuple with the CriticalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalValue

`func (o *IndicatorGaugeShared) SetCriticalValue(v string)`

SetCriticalValue sets CriticalValue field to given value.

### HasCriticalValue

`func (o *IndicatorGaugeShared) HasCriticalValue() bool`

HasCriticalValue returns a boolean if a field has been set.

### GetMajorValue

`func (o *IndicatorGaugeShared) GetMajorValue() string`

GetMajorValue returns the MajorValue field if non-nil, zero value otherwise.

### GetMajorValueOk

`func (o *IndicatorGaugeShared) GetMajorValueOk() (*string, bool)`

GetMajorValueOk returns a tuple with the MajorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorValue

`func (o *IndicatorGaugeShared) SetMajorValue(v string)`

SetMajorValue sets MajorValue field to given value.

### HasMajorValue

`func (o *IndicatorGaugeShared) HasMajorValue() bool`

HasMajorValue returns a boolean if a field has been set.

### GetMinorValue

`func (o *IndicatorGaugeShared) GetMinorValue() string`

GetMinorValue returns the MinorValue field if non-nil, zero value otherwise.

### GetMinorValueOk

`func (o *IndicatorGaugeShared) GetMinorValueOk() (*string, bool)`

GetMinorValueOk returns a tuple with the MinorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorValue

`func (o *IndicatorGaugeShared) SetMinorValue(v string)`

SetMinorValue sets MinorValue field to given value.

### HasMinorValue

`func (o *IndicatorGaugeShared) HasMinorValue() bool`

HasMinorValue returns a boolean if a field has been set.

### GetWarningValue

`func (o *IndicatorGaugeShared) GetWarningValue() string`

GetWarningValue returns the WarningValue field if non-nil, zero value otherwise.

### GetWarningValueOk

`func (o *IndicatorGaugeShared) GetWarningValueOk() (*string, bool)`

GetWarningValueOk returns a tuple with the WarningValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningValue

`func (o *IndicatorGaugeShared) SetWarningValue(v string)`

SetWarningValue sets WarningValue field to given value.

### HasWarningValue

`func (o *IndicatorGaugeShared) HasWarningValue() bool`

HasWarningValue returns a boolean if a field has been set.

### GetDescription

`func (o *IndicatorGaugeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IndicatorGaugeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IndicatorGaugeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IndicatorGaugeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *IndicatorGaugeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IndicatorGaugeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IndicatorGaugeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOverrideSeverity

`func (o *IndicatorGaugeShared) GetOverrideSeverity() EnumgaugeOverrideSeverityProp`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *IndicatorGaugeShared) GetOverrideSeverityOk() (*EnumgaugeOverrideSeverityProp, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *IndicatorGaugeShared) SetOverrideSeverity(v EnumgaugeOverrideSeverityProp)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *IndicatorGaugeShared) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetAlertLevel

`func (o *IndicatorGaugeShared) GetAlertLevel() EnumgaugeAlertLevelProp`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *IndicatorGaugeShared) GetAlertLevelOk() (*EnumgaugeAlertLevelProp, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *IndicatorGaugeShared) SetAlertLevel(v EnumgaugeAlertLevelProp)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *IndicatorGaugeShared) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *IndicatorGaugeShared) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *IndicatorGaugeShared) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *IndicatorGaugeShared) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *IndicatorGaugeShared) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetSamplesPerUpdateInterval

`func (o *IndicatorGaugeShared) GetSamplesPerUpdateInterval() int32`

GetSamplesPerUpdateInterval returns the SamplesPerUpdateInterval field if non-nil, zero value otherwise.

### GetSamplesPerUpdateIntervalOk

`func (o *IndicatorGaugeShared) GetSamplesPerUpdateIntervalOk() (*int32, bool)`

GetSamplesPerUpdateIntervalOk returns a tuple with the SamplesPerUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesPerUpdateInterval

`func (o *IndicatorGaugeShared) SetSamplesPerUpdateInterval(v int32)`

SetSamplesPerUpdateInterval sets SamplesPerUpdateInterval field to given value.

### HasSamplesPerUpdateInterval

`func (o *IndicatorGaugeShared) HasSamplesPerUpdateInterval() bool`

HasSamplesPerUpdateInterval returns a boolean if a field has been set.

### GetIncludeResource

`func (o *IndicatorGaugeShared) GetIncludeResource() []string`

GetIncludeResource returns the IncludeResource field if non-nil, zero value otherwise.

### GetIncludeResourceOk

`func (o *IndicatorGaugeShared) GetIncludeResourceOk() (*[]string, bool)`

GetIncludeResourceOk returns a tuple with the IncludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResource

`func (o *IndicatorGaugeShared) SetIncludeResource(v []string)`

SetIncludeResource sets IncludeResource field to given value.

### HasIncludeResource

`func (o *IndicatorGaugeShared) HasIncludeResource() bool`

HasIncludeResource returns a boolean if a field has been set.

### GetExcludeResource

`func (o *IndicatorGaugeShared) GetExcludeResource() []string`

GetExcludeResource returns the ExcludeResource field if non-nil, zero value otherwise.

### GetExcludeResourceOk

`func (o *IndicatorGaugeShared) GetExcludeResourceOk() (*[]string, bool)`

GetExcludeResourceOk returns a tuple with the ExcludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeResource

`func (o *IndicatorGaugeShared) SetExcludeResource(v []string)`

SetExcludeResource sets ExcludeResource field to given value.

### HasExcludeResource

`func (o *IndicatorGaugeShared) HasExcludeResource() bool`

HasExcludeResource returns a boolean if a field has been set.

### GetServerUnavailableSeverityLevel

`func (o *IndicatorGaugeShared) GetServerUnavailableSeverityLevel() EnumgaugeServerUnavailableSeverityLevelProp`

GetServerUnavailableSeverityLevel returns the ServerUnavailableSeverityLevel field if non-nil, zero value otherwise.

### GetServerUnavailableSeverityLevelOk

`func (o *IndicatorGaugeShared) GetServerUnavailableSeverityLevelOk() (*EnumgaugeServerUnavailableSeverityLevelProp, bool)`

GetServerUnavailableSeverityLevelOk returns a tuple with the ServerUnavailableSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUnavailableSeverityLevel

`func (o *IndicatorGaugeShared) SetServerUnavailableSeverityLevel(v EnumgaugeServerUnavailableSeverityLevelProp)`

SetServerUnavailableSeverityLevel sets ServerUnavailableSeverityLevel field to given value.

### HasServerUnavailableSeverityLevel

`func (o *IndicatorGaugeShared) HasServerUnavailableSeverityLevel() bool`

HasServerUnavailableSeverityLevel returns a boolean if a field has been set.

### GetServerDegradedSeverityLevel

`func (o *IndicatorGaugeShared) GetServerDegradedSeverityLevel() EnumgaugeServerDegradedSeverityLevelProp`

GetServerDegradedSeverityLevel returns the ServerDegradedSeverityLevel field if non-nil, zero value otherwise.

### GetServerDegradedSeverityLevelOk

`func (o *IndicatorGaugeShared) GetServerDegradedSeverityLevelOk() (*EnumgaugeServerDegradedSeverityLevelProp, bool)`

GetServerDegradedSeverityLevelOk returns a tuple with the ServerDegradedSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDegradedSeverityLevel

`func (o *IndicatorGaugeShared) SetServerDegradedSeverityLevel(v EnumgaugeServerDegradedSeverityLevelProp)`

SetServerDegradedSeverityLevel sets ServerDegradedSeverityLevel field to given value.

### HasServerDegradedSeverityLevel

`func (o *IndicatorGaugeShared) HasServerDegradedSeverityLevel() bool`

HasServerDegradedSeverityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


