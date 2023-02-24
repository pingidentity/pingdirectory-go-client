# AddIndicatorGaugeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GaugeName** | **string** | Name of the new Gauge | 
**Schemas** | [**[]EnumindicatorGaugeSchemaUrn**](EnumindicatorGaugeSchemaUrn.md) |  | 
**GaugeDataSource** | **string** | Specifies the source of data to use in determining this Indicator Gauge&#39;s severity and status. | 
**CriticalValue** | Pointer to **string** | A regular expression pattern that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be critical. | [optional] 
**MajorValue** | Pointer to **string** | A regular expression pattern that is used to determine whether the current monitored value indicates this gauge&#39;s severity will be &#39;major&#39;. | [optional] 
**MinorValue** | Pointer to **string** | A regular expression pattern that is used to determine whether the current monitored value indicates this gauge&#39;s severity will be &#39;minor&#39;. | [optional] 
**WarningValue** | Pointer to **string** | A regular expression pattern that is used to determine whether the current monitored value indicates this gauge&#39;s severity will be &#39;warning&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Gauge | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Gauge is enabled. | [optional] 
**OverrideSeverity** | Pointer to [**EnumgaugeOverrideSeverityProp**](EnumgaugeOverrideSeverityProp.md) |  | [optional] 
**AlertLevel** | Pointer to [**EnumgaugeAlertLevelProp**](EnumgaugeAlertLevelProp.md) |  | [optional] 
**UpdateInterval** | Pointer to **string** | The frequency with which this Gauge is updated. | [optional] 
**SamplesPerUpdateInterval** | Pointer to **int32** | Indicates the number of times the monitor data source value will be collected during the update interval. | [optional] 
**IncludeResource** | Pointer to **[]string** | Specifies set of resources to be monitored. | [optional] 
**ExcludeResource** | Pointer to **[]string** | Specifies resources to exclude from being monitored. | [optional] 
**ServerUnavailableSeverityLevel** | Pointer to [**EnumgaugeServerUnavailableSeverityLevelProp**](EnumgaugeServerUnavailableSeverityLevelProp.md) |  | [optional] 
**ServerDegradedSeverityLevel** | Pointer to [**EnumgaugeServerDegradedSeverityLevelProp**](EnumgaugeServerDegradedSeverityLevelProp.md) |  | [optional] 

## Methods

### NewAddIndicatorGaugeRequest

`func NewAddIndicatorGaugeRequest(gaugeName string, schemas []EnumindicatorGaugeSchemaUrn, gaugeDataSource string, ) *AddIndicatorGaugeRequest`

NewAddIndicatorGaugeRequest instantiates a new AddIndicatorGaugeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIndicatorGaugeRequestWithDefaults

`func NewAddIndicatorGaugeRequestWithDefaults() *AddIndicatorGaugeRequest`

NewAddIndicatorGaugeRequestWithDefaults instantiates a new AddIndicatorGaugeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGaugeName

`func (o *AddIndicatorGaugeRequest) GetGaugeName() string`

GetGaugeName returns the GaugeName field if non-nil, zero value otherwise.

### GetGaugeNameOk

`func (o *AddIndicatorGaugeRequest) GetGaugeNameOk() (*string, bool)`

GetGaugeNameOk returns a tuple with the GaugeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeName

`func (o *AddIndicatorGaugeRequest) SetGaugeName(v string)`

SetGaugeName sets GaugeName field to given value.


### GetSchemas

`func (o *AddIndicatorGaugeRequest) GetSchemas() []EnumindicatorGaugeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddIndicatorGaugeRequest) GetSchemasOk() (*[]EnumindicatorGaugeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddIndicatorGaugeRequest) SetSchemas(v []EnumindicatorGaugeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGaugeDataSource

`func (o *AddIndicatorGaugeRequest) GetGaugeDataSource() string`

GetGaugeDataSource returns the GaugeDataSource field if non-nil, zero value otherwise.

### GetGaugeDataSourceOk

`func (o *AddIndicatorGaugeRequest) GetGaugeDataSourceOk() (*string, bool)`

GetGaugeDataSourceOk returns a tuple with the GaugeDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeDataSource

`func (o *AddIndicatorGaugeRequest) SetGaugeDataSource(v string)`

SetGaugeDataSource sets GaugeDataSource field to given value.


### GetCriticalValue

`func (o *AddIndicatorGaugeRequest) GetCriticalValue() string`

GetCriticalValue returns the CriticalValue field if non-nil, zero value otherwise.

### GetCriticalValueOk

`func (o *AddIndicatorGaugeRequest) GetCriticalValueOk() (*string, bool)`

GetCriticalValueOk returns a tuple with the CriticalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalValue

`func (o *AddIndicatorGaugeRequest) SetCriticalValue(v string)`

SetCriticalValue sets CriticalValue field to given value.

### HasCriticalValue

`func (o *AddIndicatorGaugeRequest) HasCriticalValue() bool`

HasCriticalValue returns a boolean if a field has been set.

### GetMajorValue

`func (o *AddIndicatorGaugeRequest) GetMajorValue() string`

GetMajorValue returns the MajorValue field if non-nil, zero value otherwise.

### GetMajorValueOk

`func (o *AddIndicatorGaugeRequest) GetMajorValueOk() (*string, bool)`

GetMajorValueOk returns a tuple with the MajorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorValue

`func (o *AddIndicatorGaugeRequest) SetMajorValue(v string)`

SetMajorValue sets MajorValue field to given value.

### HasMajorValue

`func (o *AddIndicatorGaugeRequest) HasMajorValue() bool`

HasMajorValue returns a boolean if a field has been set.

### GetMinorValue

`func (o *AddIndicatorGaugeRequest) GetMinorValue() string`

GetMinorValue returns the MinorValue field if non-nil, zero value otherwise.

### GetMinorValueOk

`func (o *AddIndicatorGaugeRequest) GetMinorValueOk() (*string, bool)`

GetMinorValueOk returns a tuple with the MinorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorValue

`func (o *AddIndicatorGaugeRequest) SetMinorValue(v string)`

SetMinorValue sets MinorValue field to given value.

### HasMinorValue

`func (o *AddIndicatorGaugeRequest) HasMinorValue() bool`

HasMinorValue returns a boolean if a field has been set.

### GetWarningValue

`func (o *AddIndicatorGaugeRequest) GetWarningValue() string`

GetWarningValue returns the WarningValue field if non-nil, zero value otherwise.

### GetWarningValueOk

`func (o *AddIndicatorGaugeRequest) GetWarningValueOk() (*string, bool)`

GetWarningValueOk returns a tuple with the WarningValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningValue

`func (o *AddIndicatorGaugeRequest) SetWarningValue(v string)`

SetWarningValue sets WarningValue field to given value.

### HasWarningValue

`func (o *AddIndicatorGaugeRequest) HasWarningValue() bool`

HasWarningValue returns a boolean if a field has been set.

### GetDescription

`func (o *AddIndicatorGaugeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddIndicatorGaugeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddIndicatorGaugeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddIndicatorGaugeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIndicatorGaugeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIndicatorGaugeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIndicatorGaugeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddIndicatorGaugeRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *AddIndicatorGaugeRequest) GetOverrideSeverity() EnumgaugeOverrideSeverityProp`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *AddIndicatorGaugeRequest) GetOverrideSeverityOk() (*EnumgaugeOverrideSeverityProp, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *AddIndicatorGaugeRequest) SetOverrideSeverity(v EnumgaugeOverrideSeverityProp)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *AddIndicatorGaugeRequest) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetAlertLevel

`func (o *AddIndicatorGaugeRequest) GetAlertLevel() EnumgaugeAlertLevelProp`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *AddIndicatorGaugeRequest) GetAlertLevelOk() (*EnumgaugeAlertLevelProp, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *AddIndicatorGaugeRequest) SetAlertLevel(v EnumgaugeAlertLevelProp)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *AddIndicatorGaugeRequest) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *AddIndicatorGaugeRequest) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *AddIndicatorGaugeRequest) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *AddIndicatorGaugeRequest) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *AddIndicatorGaugeRequest) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetSamplesPerUpdateInterval

`func (o *AddIndicatorGaugeRequest) GetSamplesPerUpdateInterval() int32`

GetSamplesPerUpdateInterval returns the SamplesPerUpdateInterval field if non-nil, zero value otherwise.

### GetSamplesPerUpdateIntervalOk

`func (o *AddIndicatorGaugeRequest) GetSamplesPerUpdateIntervalOk() (*int32, bool)`

GetSamplesPerUpdateIntervalOk returns a tuple with the SamplesPerUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesPerUpdateInterval

`func (o *AddIndicatorGaugeRequest) SetSamplesPerUpdateInterval(v int32)`

SetSamplesPerUpdateInterval sets SamplesPerUpdateInterval field to given value.

### HasSamplesPerUpdateInterval

`func (o *AddIndicatorGaugeRequest) HasSamplesPerUpdateInterval() bool`

HasSamplesPerUpdateInterval returns a boolean if a field has been set.

### GetIncludeResource

`func (o *AddIndicatorGaugeRequest) GetIncludeResource() []string`

GetIncludeResource returns the IncludeResource field if non-nil, zero value otherwise.

### GetIncludeResourceOk

`func (o *AddIndicatorGaugeRequest) GetIncludeResourceOk() (*[]string, bool)`

GetIncludeResourceOk returns a tuple with the IncludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResource

`func (o *AddIndicatorGaugeRequest) SetIncludeResource(v []string)`

SetIncludeResource sets IncludeResource field to given value.

### HasIncludeResource

`func (o *AddIndicatorGaugeRequest) HasIncludeResource() bool`

HasIncludeResource returns a boolean if a field has been set.

### GetExcludeResource

`func (o *AddIndicatorGaugeRequest) GetExcludeResource() []string`

GetExcludeResource returns the ExcludeResource field if non-nil, zero value otherwise.

### GetExcludeResourceOk

`func (o *AddIndicatorGaugeRequest) GetExcludeResourceOk() (*[]string, bool)`

GetExcludeResourceOk returns a tuple with the ExcludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeResource

`func (o *AddIndicatorGaugeRequest) SetExcludeResource(v []string)`

SetExcludeResource sets ExcludeResource field to given value.

### HasExcludeResource

`func (o *AddIndicatorGaugeRequest) HasExcludeResource() bool`

HasExcludeResource returns a boolean if a field has been set.

### GetServerUnavailableSeverityLevel

`func (o *AddIndicatorGaugeRequest) GetServerUnavailableSeverityLevel() EnumgaugeServerUnavailableSeverityLevelProp`

GetServerUnavailableSeverityLevel returns the ServerUnavailableSeverityLevel field if non-nil, zero value otherwise.

### GetServerUnavailableSeverityLevelOk

`func (o *AddIndicatorGaugeRequest) GetServerUnavailableSeverityLevelOk() (*EnumgaugeServerUnavailableSeverityLevelProp, bool)`

GetServerUnavailableSeverityLevelOk returns a tuple with the ServerUnavailableSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUnavailableSeverityLevel

`func (o *AddIndicatorGaugeRequest) SetServerUnavailableSeverityLevel(v EnumgaugeServerUnavailableSeverityLevelProp)`

SetServerUnavailableSeverityLevel sets ServerUnavailableSeverityLevel field to given value.

### HasServerUnavailableSeverityLevel

`func (o *AddIndicatorGaugeRequest) HasServerUnavailableSeverityLevel() bool`

HasServerUnavailableSeverityLevel returns a boolean if a field has been set.

### GetServerDegradedSeverityLevel

`func (o *AddIndicatorGaugeRequest) GetServerDegradedSeverityLevel() EnumgaugeServerDegradedSeverityLevelProp`

GetServerDegradedSeverityLevel returns the ServerDegradedSeverityLevel field if non-nil, zero value otherwise.

### GetServerDegradedSeverityLevelOk

`func (o *AddIndicatorGaugeRequest) GetServerDegradedSeverityLevelOk() (*EnumgaugeServerDegradedSeverityLevelProp, bool)`

GetServerDegradedSeverityLevelOk returns a tuple with the ServerDegradedSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDegradedSeverityLevel

`func (o *AddIndicatorGaugeRequest) SetServerDegradedSeverityLevel(v EnumgaugeServerDegradedSeverityLevelProp)`

SetServerDegradedSeverityLevel sets ServerDegradedSeverityLevel field to given value.

### HasServerDegradedSeverityLevel

`func (o *AddIndicatorGaugeRequest) HasServerDegradedSeverityLevel() bool`

HasServerDegradedSeverityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


