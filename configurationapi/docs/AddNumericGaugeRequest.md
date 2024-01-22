# AddNumericGaugeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumnumericGaugeSchemaUrn**](EnumnumericGaugeSchemaUrn.md) |  | 
**GaugeDataSource** | **string** | Specifies the source of data to use in determining this gauge&#39;s current severity. | 
**CriticalValue** | Pointer to **float64** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;critical&#39;. | [optional] 
**CriticalExitValue** | Pointer to **float64** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;critical&#39;. | [optional] 
**MajorValue** | Pointer to **float64** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;major&#39;. | [optional] 
**MajorExitValue** | Pointer to **float64** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;major&#39;. | [optional] 
**MinorValue** | Pointer to **float64** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;minor&#39;. | [optional] 
**MinorExitValue** | Pointer to **float64** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;minor&#39;. | [optional] 
**WarningValue** | Pointer to **float64** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;warning&#39;. | [optional] 
**WarningExitValue** | Pointer to **float64** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;warning&#39;. | [optional] 
**Description** | Pointer to **string** | A description for this Gauge | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this Gauge is enabled. | [optional] 
**OverrideSeverity** | Pointer to [**EnumgaugeOverrideSeverityProp**](EnumgaugeOverrideSeverityProp.md) |  | [optional] 
**AlertLevel** | Pointer to [**EnumgaugeAlertLevelProp**](EnumgaugeAlertLevelProp.md) |  | [optional] 
**UpdateInterval** | Pointer to **string** | The frequency with which this Gauge is updated. | [optional] 
**SamplesPerUpdateInterval** | Pointer to **int64** | Indicates the number of times the monitor data source value will be collected during the update interval. | [optional] 
**IncludeResource** | Pointer to **[]string** | Specifies set of resources to be monitored. | [optional] 
**ExcludeResource** | Pointer to **[]string** | Specifies resources to exclude from being monitored. | [optional] 
**ServerUnavailableSeverityLevel** | Pointer to [**EnumgaugeServerUnavailableSeverityLevelProp**](EnumgaugeServerUnavailableSeverityLevelProp.md) |  | [optional] 
**ServerDegradedSeverityLevel** | Pointer to [**EnumgaugeServerDegradedSeverityLevelProp**](EnumgaugeServerDegradedSeverityLevelProp.md) |  | [optional] 
**GaugeName** | **string** | Name of the new Gauge | 

## Methods

### NewAddNumericGaugeRequest

`func NewAddNumericGaugeRequest(schemas []EnumnumericGaugeSchemaUrn, gaugeDataSource string, gaugeName string, ) *AddNumericGaugeRequest`

NewAddNumericGaugeRequest instantiates a new AddNumericGaugeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddNumericGaugeRequestWithDefaults

`func NewAddNumericGaugeRequestWithDefaults() *AddNumericGaugeRequest`

NewAddNumericGaugeRequestWithDefaults instantiates a new AddNumericGaugeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddNumericGaugeRequest) GetSchemas() []EnumnumericGaugeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddNumericGaugeRequest) GetSchemasOk() (*[]EnumnumericGaugeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddNumericGaugeRequest) SetSchemas(v []EnumnumericGaugeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGaugeDataSource

`func (o *AddNumericGaugeRequest) GetGaugeDataSource() string`

GetGaugeDataSource returns the GaugeDataSource field if non-nil, zero value otherwise.

### GetGaugeDataSourceOk

`func (o *AddNumericGaugeRequest) GetGaugeDataSourceOk() (*string, bool)`

GetGaugeDataSourceOk returns a tuple with the GaugeDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeDataSource

`func (o *AddNumericGaugeRequest) SetGaugeDataSource(v string)`

SetGaugeDataSource sets GaugeDataSource field to given value.


### GetCriticalValue

`func (o *AddNumericGaugeRequest) GetCriticalValue() float64`

GetCriticalValue returns the CriticalValue field if non-nil, zero value otherwise.

### GetCriticalValueOk

`func (o *AddNumericGaugeRequest) GetCriticalValueOk() (*float64, bool)`

GetCriticalValueOk returns a tuple with the CriticalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalValue

`func (o *AddNumericGaugeRequest) SetCriticalValue(v float64)`

SetCriticalValue sets CriticalValue field to given value.

### HasCriticalValue

`func (o *AddNumericGaugeRequest) HasCriticalValue() bool`

HasCriticalValue returns a boolean if a field has been set.

### GetCriticalExitValue

`func (o *AddNumericGaugeRequest) GetCriticalExitValue() float64`

GetCriticalExitValue returns the CriticalExitValue field if non-nil, zero value otherwise.

### GetCriticalExitValueOk

`func (o *AddNumericGaugeRequest) GetCriticalExitValueOk() (*float64, bool)`

GetCriticalExitValueOk returns a tuple with the CriticalExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalExitValue

`func (o *AddNumericGaugeRequest) SetCriticalExitValue(v float64)`

SetCriticalExitValue sets CriticalExitValue field to given value.

### HasCriticalExitValue

`func (o *AddNumericGaugeRequest) HasCriticalExitValue() bool`

HasCriticalExitValue returns a boolean if a field has been set.

### GetMajorValue

`func (o *AddNumericGaugeRequest) GetMajorValue() float64`

GetMajorValue returns the MajorValue field if non-nil, zero value otherwise.

### GetMajorValueOk

`func (o *AddNumericGaugeRequest) GetMajorValueOk() (*float64, bool)`

GetMajorValueOk returns a tuple with the MajorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorValue

`func (o *AddNumericGaugeRequest) SetMajorValue(v float64)`

SetMajorValue sets MajorValue field to given value.

### HasMajorValue

`func (o *AddNumericGaugeRequest) HasMajorValue() bool`

HasMajorValue returns a boolean if a field has been set.

### GetMajorExitValue

`func (o *AddNumericGaugeRequest) GetMajorExitValue() float64`

GetMajorExitValue returns the MajorExitValue field if non-nil, zero value otherwise.

### GetMajorExitValueOk

`func (o *AddNumericGaugeRequest) GetMajorExitValueOk() (*float64, bool)`

GetMajorExitValueOk returns a tuple with the MajorExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorExitValue

`func (o *AddNumericGaugeRequest) SetMajorExitValue(v float64)`

SetMajorExitValue sets MajorExitValue field to given value.

### HasMajorExitValue

`func (o *AddNumericGaugeRequest) HasMajorExitValue() bool`

HasMajorExitValue returns a boolean if a field has been set.

### GetMinorValue

`func (o *AddNumericGaugeRequest) GetMinorValue() float64`

GetMinorValue returns the MinorValue field if non-nil, zero value otherwise.

### GetMinorValueOk

`func (o *AddNumericGaugeRequest) GetMinorValueOk() (*float64, bool)`

GetMinorValueOk returns a tuple with the MinorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorValue

`func (o *AddNumericGaugeRequest) SetMinorValue(v float64)`

SetMinorValue sets MinorValue field to given value.

### HasMinorValue

`func (o *AddNumericGaugeRequest) HasMinorValue() bool`

HasMinorValue returns a boolean if a field has been set.

### GetMinorExitValue

`func (o *AddNumericGaugeRequest) GetMinorExitValue() float64`

GetMinorExitValue returns the MinorExitValue field if non-nil, zero value otherwise.

### GetMinorExitValueOk

`func (o *AddNumericGaugeRequest) GetMinorExitValueOk() (*float64, bool)`

GetMinorExitValueOk returns a tuple with the MinorExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorExitValue

`func (o *AddNumericGaugeRequest) SetMinorExitValue(v float64)`

SetMinorExitValue sets MinorExitValue field to given value.

### HasMinorExitValue

`func (o *AddNumericGaugeRequest) HasMinorExitValue() bool`

HasMinorExitValue returns a boolean if a field has been set.

### GetWarningValue

`func (o *AddNumericGaugeRequest) GetWarningValue() float64`

GetWarningValue returns the WarningValue field if non-nil, zero value otherwise.

### GetWarningValueOk

`func (o *AddNumericGaugeRequest) GetWarningValueOk() (*float64, bool)`

GetWarningValueOk returns a tuple with the WarningValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningValue

`func (o *AddNumericGaugeRequest) SetWarningValue(v float64)`

SetWarningValue sets WarningValue field to given value.

### HasWarningValue

`func (o *AddNumericGaugeRequest) HasWarningValue() bool`

HasWarningValue returns a boolean if a field has been set.

### GetWarningExitValue

`func (o *AddNumericGaugeRequest) GetWarningExitValue() float64`

GetWarningExitValue returns the WarningExitValue field if non-nil, zero value otherwise.

### GetWarningExitValueOk

`func (o *AddNumericGaugeRequest) GetWarningExitValueOk() (*float64, bool)`

GetWarningExitValueOk returns a tuple with the WarningExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningExitValue

`func (o *AddNumericGaugeRequest) SetWarningExitValue(v float64)`

SetWarningExitValue sets WarningExitValue field to given value.

### HasWarningExitValue

`func (o *AddNumericGaugeRequest) HasWarningExitValue() bool`

HasWarningExitValue returns a boolean if a field has been set.

### GetDescription

`func (o *AddNumericGaugeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddNumericGaugeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddNumericGaugeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddNumericGaugeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddNumericGaugeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddNumericGaugeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddNumericGaugeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddNumericGaugeRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *AddNumericGaugeRequest) GetOverrideSeverity() EnumgaugeOverrideSeverityProp`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *AddNumericGaugeRequest) GetOverrideSeverityOk() (*EnumgaugeOverrideSeverityProp, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *AddNumericGaugeRequest) SetOverrideSeverity(v EnumgaugeOverrideSeverityProp)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *AddNumericGaugeRequest) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetAlertLevel

`func (o *AddNumericGaugeRequest) GetAlertLevel() EnumgaugeAlertLevelProp`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *AddNumericGaugeRequest) GetAlertLevelOk() (*EnumgaugeAlertLevelProp, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *AddNumericGaugeRequest) SetAlertLevel(v EnumgaugeAlertLevelProp)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *AddNumericGaugeRequest) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *AddNumericGaugeRequest) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *AddNumericGaugeRequest) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *AddNumericGaugeRequest) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *AddNumericGaugeRequest) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetSamplesPerUpdateInterval

`func (o *AddNumericGaugeRequest) GetSamplesPerUpdateInterval() int64`

GetSamplesPerUpdateInterval returns the SamplesPerUpdateInterval field if non-nil, zero value otherwise.

### GetSamplesPerUpdateIntervalOk

`func (o *AddNumericGaugeRequest) GetSamplesPerUpdateIntervalOk() (*int64, bool)`

GetSamplesPerUpdateIntervalOk returns a tuple with the SamplesPerUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesPerUpdateInterval

`func (o *AddNumericGaugeRequest) SetSamplesPerUpdateInterval(v int64)`

SetSamplesPerUpdateInterval sets SamplesPerUpdateInterval field to given value.

### HasSamplesPerUpdateInterval

`func (o *AddNumericGaugeRequest) HasSamplesPerUpdateInterval() bool`

HasSamplesPerUpdateInterval returns a boolean if a field has been set.

### GetIncludeResource

`func (o *AddNumericGaugeRequest) GetIncludeResource() []string`

GetIncludeResource returns the IncludeResource field if non-nil, zero value otherwise.

### GetIncludeResourceOk

`func (o *AddNumericGaugeRequest) GetIncludeResourceOk() (*[]string, bool)`

GetIncludeResourceOk returns a tuple with the IncludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResource

`func (o *AddNumericGaugeRequest) SetIncludeResource(v []string)`

SetIncludeResource sets IncludeResource field to given value.

### HasIncludeResource

`func (o *AddNumericGaugeRequest) HasIncludeResource() bool`

HasIncludeResource returns a boolean if a field has been set.

### GetExcludeResource

`func (o *AddNumericGaugeRequest) GetExcludeResource() []string`

GetExcludeResource returns the ExcludeResource field if non-nil, zero value otherwise.

### GetExcludeResourceOk

`func (o *AddNumericGaugeRequest) GetExcludeResourceOk() (*[]string, bool)`

GetExcludeResourceOk returns a tuple with the ExcludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeResource

`func (o *AddNumericGaugeRequest) SetExcludeResource(v []string)`

SetExcludeResource sets ExcludeResource field to given value.

### HasExcludeResource

`func (o *AddNumericGaugeRequest) HasExcludeResource() bool`

HasExcludeResource returns a boolean if a field has been set.

### GetServerUnavailableSeverityLevel

`func (o *AddNumericGaugeRequest) GetServerUnavailableSeverityLevel() EnumgaugeServerUnavailableSeverityLevelProp`

GetServerUnavailableSeverityLevel returns the ServerUnavailableSeverityLevel field if non-nil, zero value otherwise.

### GetServerUnavailableSeverityLevelOk

`func (o *AddNumericGaugeRequest) GetServerUnavailableSeverityLevelOk() (*EnumgaugeServerUnavailableSeverityLevelProp, bool)`

GetServerUnavailableSeverityLevelOk returns a tuple with the ServerUnavailableSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUnavailableSeverityLevel

`func (o *AddNumericGaugeRequest) SetServerUnavailableSeverityLevel(v EnumgaugeServerUnavailableSeverityLevelProp)`

SetServerUnavailableSeverityLevel sets ServerUnavailableSeverityLevel field to given value.

### HasServerUnavailableSeverityLevel

`func (o *AddNumericGaugeRequest) HasServerUnavailableSeverityLevel() bool`

HasServerUnavailableSeverityLevel returns a boolean if a field has been set.

### GetServerDegradedSeverityLevel

`func (o *AddNumericGaugeRequest) GetServerDegradedSeverityLevel() EnumgaugeServerDegradedSeverityLevelProp`

GetServerDegradedSeverityLevel returns the ServerDegradedSeverityLevel field if non-nil, zero value otherwise.

### GetServerDegradedSeverityLevelOk

`func (o *AddNumericGaugeRequest) GetServerDegradedSeverityLevelOk() (*EnumgaugeServerDegradedSeverityLevelProp, bool)`

GetServerDegradedSeverityLevelOk returns a tuple with the ServerDegradedSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDegradedSeverityLevel

`func (o *AddNumericGaugeRequest) SetServerDegradedSeverityLevel(v EnumgaugeServerDegradedSeverityLevelProp)`

SetServerDegradedSeverityLevel sets ServerDegradedSeverityLevel field to given value.

### HasServerDegradedSeverityLevel

`func (o *AddNumericGaugeRequest) HasServerDegradedSeverityLevel() bool`

HasServerDegradedSeverityLevel returns a boolean if a field has been set.

### GetGaugeName

`func (o *AddNumericGaugeRequest) GetGaugeName() string`

GetGaugeName returns the GaugeName field if non-nil, zero value otherwise.

### GetGaugeNameOk

`func (o *AddNumericGaugeRequest) GetGaugeNameOk() (*string, bool)`

GetGaugeNameOk returns a tuple with the GaugeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeName

`func (o *AddNumericGaugeRequest) SetGaugeName(v string)`

SetGaugeName sets GaugeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


