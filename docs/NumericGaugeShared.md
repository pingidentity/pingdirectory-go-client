# NumericGaugeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumnumericGaugeSchemaUrn**](EnumnumericGaugeSchemaUrn.md) |  | 
**GaugeDataSource** | **string** | Specifies the source of data to use in determining this gauge&#39;s current severity. | 
**CriticalValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;critical&#39;. | [optional] 
**CriticalExitValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;critical&#39;. | [optional] 
**MajorValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;major&#39;. | [optional] 
**MajorExitValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;major&#39;. | [optional] 
**MinorValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;minor&#39;. | [optional] 
**MinorExitValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;minor&#39;. | [optional] 
**WarningValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should be &#39;warning&#39;. | [optional] 
**WarningExitValue** | Pointer to **float32** | A value that is used to determine whether the current monitored value indicates this gauge&#39;s severity should no longer be &#39;warning&#39;. | [optional] 
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

### NewNumericGaugeShared

`func NewNumericGaugeShared(schemas []EnumnumericGaugeSchemaUrn, gaugeDataSource string, enabled bool, ) *NumericGaugeShared`

NewNumericGaugeShared instantiates a new NumericGaugeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumericGaugeSharedWithDefaults

`func NewNumericGaugeSharedWithDefaults() *NumericGaugeShared`

NewNumericGaugeSharedWithDefaults instantiates a new NumericGaugeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *NumericGaugeShared) GetSchemas() []EnumnumericGaugeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NumericGaugeShared) GetSchemasOk() (*[]EnumnumericGaugeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NumericGaugeShared) SetSchemas(v []EnumnumericGaugeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGaugeDataSource

`func (o *NumericGaugeShared) GetGaugeDataSource() string`

GetGaugeDataSource returns the GaugeDataSource field if non-nil, zero value otherwise.

### GetGaugeDataSourceOk

`func (o *NumericGaugeShared) GetGaugeDataSourceOk() (*string, bool)`

GetGaugeDataSourceOk returns a tuple with the GaugeDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeDataSource

`func (o *NumericGaugeShared) SetGaugeDataSource(v string)`

SetGaugeDataSource sets GaugeDataSource field to given value.


### GetCriticalValue

`func (o *NumericGaugeShared) GetCriticalValue() float32`

GetCriticalValue returns the CriticalValue field if non-nil, zero value otherwise.

### GetCriticalValueOk

`func (o *NumericGaugeShared) GetCriticalValueOk() (*float32, bool)`

GetCriticalValueOk returns a tuple with the CriticalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalValue

`func (o *NumericGaugeShared) SetCriticalValue(v float32)`

SetCriticalValue sets CriticalValue field to given value.

### HasCriticalValue

`func (o *NumericGaugeShared) HasCriticalValue() bool`

HasCriticalValue returns a boolean if a field has been set.

### GetCriticalExitValue

`func (o *NumericGaugeShared) GetCriticalExitValue() float32`

GetCriticalExitValue returns the CriticalExitValue field if non-nil, zero value otherwise.

### GetCriticalExitValueOk

`func (o *NumericGaugeShared) GetCriticalExitValueOk() (*float32, bool)`

GetCriticalExitValueOk returns a tuple with the CriticalExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalExitValue

`func (o *NumericGaugeShared) SetCriticalExitValue(v float32)`

SetCriticalExitValue sets CriticalExitValue field to given value.

### HasCriticalExitValue

`func (o *NumericGaugeShared) HasCriticalExitValue() bool`

HasCriticalExitValue returns a boolean if a field has been set.

### GetMajorValue

`func (o *NumericGaugeShared) GetMajorValue() float32`

GetMajorValue returns the MajorValue field if non-nil, zero value otherwise.

### GetMajorValueOk

`func (o *NumericGaugeShared) GetMajorValueOk() (*float32, bool)`

GetMajorValueOk returns a tuple with the MajorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorValue

`func (o *NumericGaugeShared) SetMajorValue(v float32)`

SetMajorValue sets MajorValue field to given value.

### HasMajorValue

`func (o *NumericGaugeShared) HasMajorValue() bool`

HasMajorValue returns a boolean if a field has been set.

### GetMajorExitValue

`func (o *NumericGaugeShared) GetMajorExitValue() float32`

GetMajorExitValue returns the MajorExitValue field if non-nil, zero value otherwise.

### GetMajorExitValueOk

`func (o *NumericGaugeShared) GetMajorExitValueOk() (*float32, bool)`

GetMajorExitValueOk returns a tuple with the MajorExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorExitValue

`func (o *NumericGaugeShared) SetMajorExitValue(v float32)`

SetMajorExitValue sets MajorExitValue field to given value.

### HasMajorExitValue

`func (o *NumericGaugeShared) HasMajorExitValue() bool`

HasMajorExitValue returns a boolean if a field has been set.

### GetMinorValue

`func (o *NumericGaugeShared) GetMinorValue() float32`

GetMinorValue returns the MinorValue field if non-nil, zero value otherwise.

### GetMinorValueOk

`func (o *NumericGaugeShared) GetMinorValueOk() (*float32, bool)`

GetMinorValueOk returns a tuple with the MinorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorValue

`func (o *NumericGaugeShared) SetMinorValue(v float32)`

SetMinorValue sets MinorValue field to given value.

### HasMinorValue

`func (o *NumericGaugeShared) HasMinorValue() bool`

HasMinorValue returns a boolean if a field has been set.

### GetMinorExitValue

`func (o *NumericGaugeShared) GetMinorExitValue() float32`

GetMinorExitValue returns the MinorExitValue field if non-nil, zero value otherwise.

### GetMinorExitValueOk

`func (o *NumericGaugeShared) GetMinorExitValueOk() (*float32, bool)`

GetMinorExitValueOk returns a tuple with the MinorExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorExitValue

`func (o *NumericGaugeShared) SetMinorExitValue(v float32)`

SetMinorExitValue sets MinorExitValue field to given value.

### HasMinorExitValue

`func (o *NumericGaugeShared) HasMinorExitValue() bool`

HasMinorExitValue returns a boolean if a field has been set.

### GetWarningValue

`func (o *NumericGaugeShared) GetWarningValue() float32`

GetWarningValue returns the WarningValue field if non-nil, zero value otherwise.

### GetWarningValueOk

`func (o *NumericGaugeShared) GetWarningValueOk() (*float32, bool)`

GetWarningValueOk returns a tuple with the WarningValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningValue

`func (o *NumericGaugeShared) SetWarningValue(v float32)`

SetWarningValue sets WarningValue field to given value.

### HasWarningValue

`func (o *NumericGaugeShared) HasWarningValue() bool`

HasWarningValue returns a boolean if a field has been set.

### GetWarningExitValue

`func (o *NumericGaugeShared) GetWarningExitValue() float32`

GetWarningExitValue returns the WarningExitValue field if non-nil, zero value otherwise.

### GetWarningExitValueOk

`func (o *NumericGaugeShared) GetWarningExitValueOk() (*float32, bool)`

GetWarningExitValueOk returns a tuple with the WarningExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningExitValue

`func (o *NumericGaugeShared) SetWarningExitValue(v float32)`

SetWarningExitValue sets WarningExitValue field to given value.

### HasWarningExitValue

`func (o *NumericGaugeShared) HasWarningExitValue() bool`

HasWarningExitValue returns a boolean if a field has been set.

### GetDescription

`func (o *NumericGaugeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NumericGaugeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NumericGaugeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NumericGaugeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *NumericGaugeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NumericGaugeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NumericGaugeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOverrideSeverity

`func (o *NumericGaugeShared) GetOverrideSeverity() EnumgaugeOverrideSeverityProp`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *NumericGaugeShared) GetOverrideSeverityOk() (*EnumgaugeOverrideSeverityProp, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *NumericGaugeShared) SetOverrideSeverity(v EnumgaugeOverrideSeverityProp)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *NumericGaugeShared) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetAlertLevel

`func (o *NumericGaugeShared) GetAlertLevel() EnumgaugeAlertLevelProp`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *NumericGaugeShared) GetAlertLevelOk() (*EnumgaugeAlertLevelProp, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *NumericGaugeShared) SetAlertLevel(v EnumgaugeAlertLevelProp)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *NumericGaugeShared) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *NumericGaugeShared) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *NumericGaugeShared) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *NumericGaugeShared) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *NumericGaugeShared) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetSamplesPerUpdateInterval

`func (o *NumericGaugeShared) GetSamplesPerUpdateInterval() int32`

GetSamplesPerUpdateInterval returns the SamplesPerUpdateInterval field if non-nil, zero value otherwise.

### GetSamplesPerUpdateIntervalOk

`func (o *NumericGaugeShared) GetSamplesPerUpdateIntervalOk() (*int32, bool)`

GetSamplesPerUpdateIntervalOk returns a tuple with the SamplesPerUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesPerUpdateInterval

`func (o *NumericGaugeShared) SetSamplesPerUpdateInterval(v int32)`

SetSamplesPerUpdateInterval sets SamplesPerUpdateInterval field to given value.

### HasSamplesPerUpdateInterval

`func (o *NumericGaugeShared) HasSamplesPerUpdateInterval() bool`

HasSamplesPerUpdateInterval returns a boolean if a field has been set.

### GetIncludeResource

`func (o *NumericGaugeShared) GetIncludeResource() []string`

GetIncludeResource returns the IncludeResource field if non-nil, zero value otherwise.

### GetIncludeResourceOk

`func (o *NumericGaugeShared) GetIncludeResourceOk() (*[]string, bool)`

GetIncludeResourceOk returns a tuple with the IncludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResource

`func (o *NumericGaugeShared) SetIncludeResource(v []string)`

SetIncludeResource sets IncludeResource field to given value.

### HasIncludeResource

`func (o *NumericGaugeShared) HasIncludeResource() bool`

HasIncludeResource returns a boolean if a field has been set.

### GetExcludeResource

`func (o *NumericGaugeShared) GetExcludeResource() []string`

GetExcludeResource returns the ExcludeResource field if non-nil, zero value otherwise.

### GetExcludeResourceOk

`func (o *NumericGaugeShared) GetExcludeResourceOk() (*[]string, bool)`

GetExcludeResourceOk returns a tuple with the ExcludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeResource

`func (o *NumericGaugeShared) SetExcludeResource(v []string)`

SetExcludeResource sets ExcludeResource field to given value.

### HasExcludeResource

`func (o *NumericGaugeShared) HasExcludeResource() bool`

HasExcludeResource returns a boolean if a field has been set.

### GetServerUnavailableSeverityLevel

`func (o *NumericGaugeShared) GetServerUnavailableSeverityLevel() EnumgaugeServerUnavailableSeverityLevelProp`

GetServerUnavailableSeverityLevel returns the ServerUnavailableSeverityLevel field if non-nil, zero value otherwise.

### GetServerUnavailableSeverityLevelOk

`func (o *NumericGaugeShared) GetServerUnavailableSeverityLevelOk() (*EnumgaugeServerUnavailableSeverityLevelProp, bool)`

GetServerUnavailableSeverityLevelOk returns a tuple with the ServerUnavailableSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUnavailableSeverityLevel

`func (o *NumericGaugeShared) SetServerUnavailableSeverityLevel(v EnumgaugeServerUnavailableSeverityLevelProp)`

SetServerUnavailableSeverityLevel sets ServerUnavailableSeverityLevel field to given value.

### HasServerUnavailableSeverityLevel

`func (o *NumericGaugeShared) HasServerUnavailableSeverityLevel() bool`

HasServerUnavailableSeverityLevel returns a boolean if a field has been set.

### GetServerDegradedSeverityLevel

`func (o *NumericGaugeShared) GetServerDegradedSeverityLevel() EnumgaugeServerDegradedSeverityLevelProp`

GetServerDegradedSeverityLevel returns the ServerDegradedSeverityLevel field if non-nil, zero value otherwise.

### GetServerDegradedSeverityLevelOk

`func (o *NumericGaugeShared) GetServerDegradedSeverityLevelOk() (*EnumgaugeServerDegradedSeverityLevelProp, bool)`

GetServerDegradedSeverityLevelOk returns a tuple with the ServerDegradedSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDegradedSeverityLevel

`func (o *NumericGaugeShared) SetServerDegradedSeverityLevel(v EnumgaugeServerDegradedSeverityLevelProp)`

SetServerDegradedSeverityLevel sets ServerDegradedSeverityLevel field to given value.

### HasServerDegradedSeverityLevel

`func (o *NumericGaugeShared) HasServerDegradedSeverityLevel() bool`

HasServerDegradedSeverityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


