# AddGaugeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GaugeName** | **string** | Name of the new Gauge | 
**Schemas** | [**[]EnumnumericGaugeSchemaUrn**](EnumnumericGaugeSchemaUrn.md) |  | 
**GaugeDataSource** | **string** | Specifies the source of data to use in determining this gauge&#39;s current severity. | 
**CriticalValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;critical&#39;. | [optional] 
**MajorValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;major&#39;. | [optional] 
**MinorValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;minor&#39;. | [optional] 
**WarningValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;warning&#39;. | [optional] 
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
**CriticalExitValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;critical&#39;. | [optional] 
**MajorExitValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;major&#39;. | [optional] 
**MinorExitValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;minor&#39;. | [optional] 
**WarningExitValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;warning&#39;. | [optional] 

## Methods

### NewAddGaugeRequest

`func NewAddGaugeRequest(gaugeName string, schemas []EnumnumericGaugeSchemaUrn, gaugeDataSource string, ) *AddGaugeRequest`

NewAddGaugeRequest instantiates a new AddGaugeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGaugeRequestWithDefaults

`func NewAddGaugeRequestWithDefaults() *AddGaugeRequest`

NewAddGaugeRequestWithDefaults instantiates a new AddGaugeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGaugeName

`func (o *AddGaugeRequest) GetGaugeName() string`

GetGaugeName returns the GaugeName field if non-nil, zero value otherwise.

### GetGaugeNameOk

`func (o *AddGaugeRequest) GetGaugeNameOk() (*string, bool)`

GetGaugeNameOk returns a tuple with the GaugeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeName

`func (o *AddGaugeRequest) SetGaugeName(v string)`

SetGaugeName sets GaugeName field to given value.


### GetSchemas

`func (o *AddGaugeRequest) GetSchemas() []EnumnumericGaugeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddGaugeRequest) GetSchemasOk() (*[]EnumnumericGaugeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddGaugeRequest) SetSchemas(v []EnumnumericGaugeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGaugeDataSource

`func (o *AddGaugeRequest) GetGaugeDataSource() string`

GetGaugeDataSource returns the GaugeDataSource field if non-nil, zero value otherwise.

### GetGaugeDataSourceOk

`func (o *AddGaugeRequest) GetGaugeDataSourceOk() (*string, bool)`

GetGaugeDataSourceOk returns a tuple with the GaugeDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeDataSource

`func (o *AddGaugeRequest) SetGaugeDataSource(v string)`

SetGaugeDataSource sets GaugeDataSource field to given value.


### GetCriticalValue

`func (o *AddGaugeRequest) GetCriticalValue() float32`

GetCriticalValue returns the CriticalValue field if non-nil, zero value otherwise.

### GetCriticalValueOk

`func (o *AddGaugeRequest) GetCriticalValueOk() (*float32, bool)`

GetCriticalValueOk returns a tuple with the CriticalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalValue

`func (o *AddGaugeRequest) SetCriticalValue(v float32)`

SetCriticalValue sets CriticalValue field to given value.

### HasCriticalValue

`func (o *AddGaugeRequest) HasCriticalValue() bool`

HasCriticalValue returns a boolean if a field has been set.

### GetMajorValue

`func (o *AddGaugeRequest) GetMajorValue() float32`

GetMajorValue returns the MajorValue field if non-nil, zero value otherwise.

### GetMajorValueOk

`func (o *AddGaugeRequest) GetMajorValueOk() (*float32, bool)`

GetMajorValueOk returns a tuple with the MajorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorValue

`func (o *AddGaugeRequest) SetMajorValue(v float32)`

SetMajorValue sets MajorValue field to given value.

### HasMajorValue

`func (o *AddGaugeRequest) HasMajorValue() bool`

HasMajorValue returns a boolean if a field has been set.

### GetMinorValue

`func (o *AddGaugeRequest) GetMinorValue() float32`

GetMinorValue returns the MinorValue field if non-nil, zero value otherwise.

### GetMinorValueOk

`func (o *AddGaugeRequest) GetMinorValueOk() (*float32, bool)`

GetMinorValueOk returns a tuple with the MinorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorValue

`func (o *AddGaugeRequest) SetMinorValue(v float32)`

SetMinorValue sets MinorValue field to given value.

### HasMinorValue

`func (o *AddGaugeRequest) HasMinorValue() bool`

HasMinorValue returns a boolean if a field has been set.

### GetWarningValue

`func (o *AddGaugeRequest) GetWarningValue() float32`

GetWarningValue returns the WarningValue field if non-nil, zero value otherwise.

### GetWarningValueOk

`func (o *AddGaugeRequest) GetWarningValueOk() (*float32, bool)`

GetWarningValueOk returns a tuple with the WarningValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningValue

`func (o *AddGaugeRequest) SetWarningValue(v float32)`

SetWarningValue sets WarningValue field to given value.

### HasWarningValue

`func (o *AddGaugeRequest) HasWarningValue() bool`

HasWarningValue returns a boolean if a field has been set.

### GetDescription

`func (o *AddGaugeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddGaugeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddGaugeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddGaugeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddGaugeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddGaugeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddGaugeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddGaugeRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *AddGaugeRequest) GetOverrideSeverity() EnumgaugeOverrideSeverityProp`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *AddGaugeRequest) GetOverrideSeverityOk() (*EnumgaugeOverrideSeverityProp, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *AddGaugeRequest) SetOverrideSeverity(v EnumgaugeOverrideSeverityProp)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *AddGaugeRequest) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetAlertLevel

`func (o *AddGaugeRequest) GetAlertLevel() EnumgaugeAlertLevelProp`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *AddGaugeRequest) GetAlertLevelOk() (*EnumgaugeAlertLevelProp, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *AddGaugeRequest) SetAlertLevel(v EnumgaugeAlertLevelProp)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *AddGaugeRequest) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *AddGaugeRequest) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *AddGaugeRequest) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *AddGaugeRequest) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *AddGaugeRequest) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetSamplesPerUpdateInterval

`func (o *AddGaugeRequest) GetSamplesPerUpdateInterval() int32`

GetSamplesPerUpdateInterval returns the SamplesPerUpdateInterval field if non-nil, zero value otherwise.

### GetSamplesPerUpdateIntervalOk

`func (o *AddGaugeRequest) GetSamplesPerUpdateIntervalOk() (*int32, bool)`

GetSamplesPerUpdateIntervalOk returns a tuple with the SamplesPerUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesPerUpdateInterval

`func (o *AddGaugeRequest) SetSamplesPerUpdateInterval(v int32)`

SetSamplesPerUpdateInterval sets SamplesPerUpdateInterval field to given value.

### HasSamplesPerUpdateInterval

`func (o *AddGaugeRequest) HasSamplesPerUpdateInterval() bool`

HasSamplesPerUpdateInterval returns a boolean if a field has been set.

### GetIncludeResource

`func (o *AddGaugeRequest) GetIncludeResource() []string`

GetIncludeResource returns the IncludeResource field if non-nil, zero value otherwise.

### GetIncludeResourceOk

`func (o *AddGaugeRequest) GetIncludeResourceOk() (*[]string, bool)`

GetIncludeResourceOk returns a tuple with the IncludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResource

`func (o *AddGaugeRequest) SetIncludeResource(v []string)`

SetIncludeResource sets IncludeResource field to given value.

### HasIncludeResource

`func (o *AddGaugeRequest) HasIncludeResource() bool`

HasIncludeResource returns a boolean if a field has been set.

### GetExcludeResource

`func (o *AddGaugeRequest) GetExcludeResource() []string`

GetExcludeResource returns the ExcludeResource field if non-nil, zero value otherwise.

### GetExcludeResourceOk

`func (o *AddGaugeRequest) GetExcludeResourceOk() (*[]string, bool)`

GetExcludeResourceOk returns a tuple with the ExcludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeResource

`func (o *AddGaugeRequest) SetExcludeResource(v []string)`

SetExcludeResource sets ExcludeResource field to given value.

### HasExcludeResource

`func (o *AddGaugeRequest) HasExcludeResource() bool`

HasExcludeResource returns a boolean if a field has been set.

### GetServerUnavailableSeverityLevel

`func (o *AddGaugeRequest) GetServerUnavailableSeverityLevel() EnumgaugeServerUnavailableSeverityLevelProp`

GetServerUnavailableSeverityLevel returns the ServerUnavailableSeverityLevel field if non-nil, zero value otherwise.

### GetServerUnavailableSeverityLevelOk

`func (o *AddGaugeRequest) GetServerUnavailableSeverityLevelOk() (*EnumgaugeServerUnavailableSeverityLevelProp, bool)`

GetServerUnavailableSeverityLevelOk returns a tuple with the ServerUnavailableSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUnavailableSeverityLevel

`func (o *AddGaugeRequest) SetServerUnavailableSeverityLevel(v EnumgaugeServerUnavailableSeverityLevelProp)`

SetServerUnavailableSeverityLevel sets ServerUnavailableSeverityLevel field to given value.

### HasServerUnavailableSeverityLevel

`func (o *AddGaugeRequest) HasServerUnavailableSeverityLevel() bool`

HasServerUnavailableSeverityLevel returns a boolean if a field has been set.

### GetServerDegradedSeverityLevel

`func (o *AddGaugeRequest) GetServerDegradedSeverityLevel() EnumgaugeServerDegradedSeverityLevelProp`

GetServerDegradedSeverityLevel returns the ServerDegradedSeverityLevel field if non-nil, zero value otherwise.

### GetServerDegradedSeverityLevelOk

`func (o *AddGaugeRequest) GetServerDegradedSeverityLevelOk() (*EnumgaugeServerDegradedSeverityLevelProp, bool)`

GetServerDegradedSeverityLevelOk returns a tuple with the ServerDegradedSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDegradedSeverityLevel

`func (o *AddGaugeRequest) SetServerDegradedSeverityLevel(v EnumgaugeServerDegradedSeverityLevelProp)`

SetServerDegradedSeverityLevel sets ServerDegradedSeverityLevel field to given value.

### HasServerDegradedSeverityLevel

`func (o *AddGaugeRequest) HasServerDegradedSeverityLevel() bool`

HasServerDegradedSeverityLevel returns a boolean if a field has been set.

### GetCriticalExitValue

`func (o *AddGaugeRequest) GetCriticalExitValue() float32`

GetCriticalExitValue returns the CriticalExitValue field if non-nil, zero value otherwise.

### GetCriticalExitValueOk

`func (o *AddGaugeRequest) GetCriticalExitValueOk() (*float32, bool)`

GetCriticalExitValueOk returns a tuple with the CriticalExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalExitValue

`func (o *AddGaugeRequest) SetCriticalExitValue(v float32)`

SetCriticalExitValue sets CriticalExitValue field to given value.

### HasCriticalExitValue

`func (o *AddGaugeRequest) HasCriticalExitValue() bool`

HasCriticalExitValue returns a boolean if a field has been set.

### GetMajorExitValue

`func (o *AddGaugeRequest) GetMajorExitValue() float32`

GetMajorExitValue returns the MajorExitValue field if non-nil, zero value otherwise.

### GetMajorExitValueOk

`func (o *AddGaugeRequest) GetMajorExitValueOk() (*float32, bool)`

GetMajorExitValueOk returns a tuple with the MajorExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorExitValue

`func (o *AddGaugeRequest) SetMajorExitValue(v float32)`

SetMajorExitValue sets MajorExitValue field to given value.

### HasMajorExitValue

`func (o *AddGaugeRequest) HasMajorExitValue() bool`

HasMajorExitValue returns a boolean if a field has been set.

### GetMinorExitValue

`func (o *AddGaugeRequest) GetMinorExitValue() float32`

GetMinorExitValue returns the MinorExitValue field if non-nil, zero value otherwise.

### GetMinorExitValueOk

`func (o *AddGaugeRequest) GetMinorExitValueOk() (*float32, bool)`

GetMinorExitValueOk returns a tuple with the MinorExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorExitValue

`func (o *AddGaugeRequest) SetMinorExitValue(v float32)`

SetMinorExitValue sets MinorExitValue field to given value.

### HasMinorExitValue

`func (o *AddGaugeRequest) HasMinorExitValue() bool`

HasMinorExitValue returns a boolean if a field has been set.

### GetWarningExitValue

`func (o *AddGaugeRequest) GetWarningExitValue() float32`

GetWarningExitValue returns the WarningExitValue field if non-nil, zero value otherwise.

### GetWarningExitValueOk

`func (o *AddGaugeRequest) GetWarningExitValueOk() (*float32, bool)`

GetWarningExitValueOk returns a tuple with the WarningExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningExitValue

`func (o *AddGaugeRequest) SetWarningExitValue(v float32)`

SetWarningExitValue sets WarningExitValue field to given value.

### HasWarningExitValue

`func (o *AddGaugeRequest) HasWarningExitValue() bool`

HasWarningExitValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


