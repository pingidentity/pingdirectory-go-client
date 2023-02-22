# NumericGaugeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Gauge | 
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
**IncludeResource** | Pointer to **[]string** | Specifies set of resources to be monitored. | [optional] 
**ExcludeResource** | Pointer to **[]string** | Specifies resources to exclude from being monitored. | [optional] 
**ServerUnavailableSeverityLevel** | Pointer to [**EnumgaugeServerUnavailableSeverityLevelProp**](EnumgaugeServerUnavailableSeverityLevelProp.md) |  | [optional] 
**ServerDegradedSeverityLevel** | Pointer to [**EnumgaugeServerDegradedSeverityLevelProp**](EnumgaugeServerDegradedSeverityLevelProp.md) |  | [optional] 

## Methods

### NewNumericGaugeResponse

`func NewNumericGaugeResponse(id string, schemas []EnumnumericGaugeSchemaUrn, gaugeDataSource string, enabled bool, ) *NumericGaugeResponse`

NewNumericGaugeResponse instantiates a new NumericGaugeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumericGaugeResponseWithDefaults

`func NewNumericGaugeResponseWithDefaults() *NumericGaugeResponse`

NewNumericGaugeResponseWithDefaults instantiates a new NumericGaugeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *NumericGaugeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NumericGaugeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NumericGaugeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NumericGaugeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *NumericGaugeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *NumericGaugeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *NumericGaugeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *NumericGaugeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *NumericGaugeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NumericGaugeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NumericGaugeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *NumericGaugeResponse) GetSchemas() []EnumnumericGaugeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *NumericGaugeResponse) GetSchemasOk() (*[]EnumnumericGaugeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *NumericGaugeResponse) SetSchemas(v []EnumnumericGaugeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetGaugeDataSource

`func (o *NumericGaugeResponse) GetGaugeDataSource() string`

GetGaugeDataSource returns the GaugeDataSource field if non-nil, zero value otherwise.

### GetGaugeDataSourceOk

`func (o *NumericGaugeResponse) GetGaugeDataSourceOk() (*string, bool)`

GetGaugeDataSourceOk returns a tuple with the GaugeDataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeDataSource

`func (o *NumericGaugeResponse) SetGaugeDataSource(v string)`

SetGaugeDataSource sets GaugeDataSource field to given value.


### GetCriticalValue

`func (o *NumericGaugeResponse) GetCriticalValue() float32`

GetCriticalValue returns the CriticalValue field if non-nil, zero value otherwise.

### GetCriticalValueOk

`func (o *NumericGaugeResponse) GetCriticalValueOk() (*float32, bool)`

GetCriticalValueOk returns a tuple with the CriticalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalValue

`func (o *NumericGaugeResponse) SetCriticalValue(v float32)`

SetCriticalValue sets CriticalValue field to given value.

### HasCriticalValue

`func (o *NumericGaugeResponse) HasCriticalValue() bool`

HasCriticalValue returns a boolean if a field has been set.

### GetCriticalExitValue

`func (o *NumericGaugeResponse) GetCriticalExitValue() float32`

GetCriticalExitValue returns the CriticalExitValue field if non-nil, zero value otherwise.

### GetCriticalExitValueOk

`func (o *NumericGaugeResponse) GetCriticalExitValueOk() (*float32, bool)`

GetCriticalExitValueOk returns a tuple with the CriticalExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalExitValue

`func (o *NumericGaugeResponse) SetCriticalExitValue(v float32)`

SetCriticalExitValue sets CriticalExitValue field to given value.

### HasCriticalExitValue

`func (o *NumericGaugeResponse) HasCriticalExitValue() bool`

HasCriticalExitValue returns a boolean if a field has been set.

### GetMajorValue

`func (o *NumericGaugeResponse) GetMajorValue() float32`

GetMajorValue returns the MajorValue field if non-nil, zero value otherwise.

### GetMajorValueOk

`func (o *NumericGaugeResponse) GetMajorValueOk() (*float32, bool)`

GetMajorValueOk returns a tuple with the MajorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorValue

`func (o *NumericGaugeResponse) SetMajorValue(v float32)`

SetMajorValue sets MajorValue field to given value.

### HasMajorValue

`func (o *NumericGaugeResponse) HasMajorValue() bool`

HasMajorValue returns a boolean if a field has been set.

### GetMajorExitValue

`func (o *NumericGaugeResponse) GetMajorExitValue() float32`

GetMajorExitValue returns the MajorExitValue field if non-nil, zero value otherwise.

### GetMajorExitValueOk

`func (o *NumericGaugeResponse) GetMajorExitValueOk() (*float32, bool)`

GetMajorExitValueOk returns a tuple with the MajorExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorExitValue

`func (o *NumericGaugeResponse) SetMajorExitValue(v float32)`

SetMajorExitValue sets MajorExitValue field to given value.

### HasMajorExitValue

`func (o *NumericGaugeResponse) HasMajorExitValue() bool`

HasMajorExitValue returns a boolean if a field has been set.

### GetMinorValue

`func (o *NumericGaugeResponse) GetMinorValue() float32`

GetMinorValue returns the MinorValue field if non-nil, zero value otherwise.

### GetMinorValueOk

`func (o *NumericGaugeResponse) GetMinorValueOk() (*float32, bool)`

GetMinorValueOk returns a tuple with the MinorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorValue

`func (o *NumericGaugeResponse) SetMinorValue(v float32)`

SetMinorValue sets MinorValue field to given value.

### HasMinorValue

`func (o *NumericGaugeResponse) HasMinorValue() bool`

HasMinorValue returns a boolean if a field has been set.

### GetMinorExitValue

`func (o *NumericGaugeResponse) GetMinorExitValue() float32`

GetMinorExitValue returns the MinorExitValue field if non-nil, zero value otherwise.

### GetMinorExitValueOk

`func (o *NumericGaugeResponse) GetMinorExitValueOk() (*float32, bool)`

GetMinorExitValueOk returns a tuple with the MinorExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorExitValue

`func (o *NumericGaugeResponse) SetMinorExitValue(v float32)`

SetMinorExitValue sets MinorExitValue field to given value.

### HasMinorExitValue

`func (o *NumericGaugeResponse) HasMinorExitValue() bool`

HasMinorExitValue returns a boolean if a field has been set.

### GetWarningValue

`func (o *NumericGaugeResponse) GetWarningValue() float32`

GetWarningValue returns the WarningValue field if non-nil, zero value otherwise.

### GetWarningValueOk

`func (o *NumericGaugeResponse) GetWarningValueOk() (*float32, bool)`

GetWarningValueOk returns a tuple with the WarningValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningValue

`func (o *NumericGaugeResponse) SetWarningValue(v float32)`

SetWarningValue sets WarningValue field to given value.

### HasWarningValue

`func (o *NumericGaugeResponse) HasWarningValue() bool`

HasWarningValue returns a boolean if a field has been set.

### GetWarningExitValue

`func (o *NumericGaugeResponse) GetWarningExitValue() float32`

GetWarningExitValue returns the WarningExitValue field if non-nil, zero value otherwise.

### GetWarningExitValueOk

`func (o *NumericGaugeResponse) GetWarningExitValueOk() (*float32, bool)`

GetWarningExitValueOk returns a tuple with the WarningExitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningExitValue

`func (o *NumericGaugeResponse) SetWarningExitValue(v float32)`

SetWarningExitValue sets WarningExitValue field to given value.

### HasWarningExitValue

`func (o *NumericGaugeResponse) HasWarningExitValue() bool`

HasWarningExitValue returns a boolean if a field has been set.

### GetDescription

`func (o *NumericGaugeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NumericGaugeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NumericGaugeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NumericGaugeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *NumericGaugeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NumericGaugeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NumericGaugeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOverrideSeverity

`func (o *NumericGaugeResponse) GetOverrideSeverity() EnumgaugeOverrideSeverityProp`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *NumericGaugeResponse) GetOverrideSeverityOk() (*EnumgaugeOverrideSeverityProp, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *NumericGaugeResponse) SetOverrideSeverity(v EnumgaugeOverrideSeverityProp)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *NumericGaugeResponse) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.

### GetAlertLevel

`func (o *NumericGaugeResponse) GetAlertLevel() EnumgaugeAlertLevelProp`

GetAlertLevel returns the AlertLevel field if non-nil, zero value otherwise.

### GetAlertLevelOk

`func (o *NumericGaugeResponse) GetAlertLevelOk() (*EnumgaugeAlertLevelProp, bool)`

GetAlertLevelOk returns a tuple with the AlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertLevel

`func (o *NumericGaugeResponse) SetAlertLevel(v EnumgaugeAlertLevelProp)`

SetAlertLevel sets AlertLevel field to given value.

### HasAlertLevel

`func (o *NumericGaugeResponse) HasAlertLevel() bool`

HasAlertLevel returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *NumericGaugeResponse) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *NumericGaugeResponse) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *NumericGaugeResponse) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *NumericGaugeResponse) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetSamplesPerUpdateInterval

`func (o *NumericGaugeResponse) GetSamplesPerUpdateInterval() int32`

GetSamplesPerUpdateInterval returns the SamplesPerUpdateInterval field if non-nil, zero value otherwise.

### GetSamplesPerUpdateIntervalOk

`func (o *NumericGaugeResponse) GetSamplesPerUpdateIntervalOk() (*int32, bool)`

GetSamplesPerUpdateIntervalOk returns a tuple with the SamplesPerUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplesPerUpdateInterval

`func (o *NumericGaugeResponse) SetSamplesPerUpdateInterval(v int32)`

SetSamplesPerUpdateInterval sets SamplesPerUpdateInterval field to given value.

### HasSamplesPerUpdateInterval

`func (o *NumericGaugeResponse) HasSamplesPerUpdateInterval() bool`

HasSamplesPerUpdateInterval returns a boolean if a field has been set.

### GetIncludeResource

`func (o *NumericGaugeResponse) GetIncludeResource() []string`

GetIncludeResource returns the IncludeResource field if non-nil, zero value otherwise.

### GetIncludeResourceOk

`func (o *NumericGaugeResponse) GetIncludeResourceOk() (*[]string, bool)`

GetIncludeResourceOk returns a tuple with the IncludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResource

`func (o *NumericGaugeResponse) SetIncludeResource(v []string)`

SetIncludeResource sets IncludeResource field to given value.

### HasIncludeResource

`func (o *NumericGaugeResponse) HasIncludeResource() bool`

HasIncludeResource returns a boolean if a field has been set.

### GetExcludeResource

`func (o *NumericGaugeResponse) GetExcludeResource() []string`

GetExcludeResource returns the ExcludeResource field if non-nil, zero value otherwise.

### GetExcludeResourceOk

`func (o *NumericGaugeResponse) GetExcludeResourceOk() (*[]string, bool)`

GetExcludeResourceOk returns a tuple with the ExcludeResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeResource

`func (o *NumericGaugeResponse) SetExcludeResource(v []string)`

SetExcludeResource sets ExcludeResource field to given value.

### HasExcludeResource

`func (o *NumericGaugeResponse) HasExcludeResource() bool`

HasExcludeResource returns a boolean if a field has been set.

### GetServerUnavailableSeverityLevel

`func (o *NumericGaugeResponse) GetServerUnavailableSeverityLevel() EnumgaugeServerUnavailableSeverityLevelProp`

GetServerUnavailableSeverityLevel returns the ServerUnavailableSeverityLevel field if non-nil, zero value otherwise.

### GetServerUnavailableSeverityLevelOk

`func (o *NumericGaugeResponse) GetServerUnavailableSeverityLevelOk() (*EnumgaugeServerUnavailableSeverityLevelProp, bool)`

GetServerUnavailableSeverityLevelOk returns a tuple with the ServerUnavailableSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUnavailableSeverityLevel

`func (o *NumericGaugeResponse) SetServerUnavailableSeverityLevel(v EnumgaugeServerUnavailableSeverityLevelProp)`

SetServerUnavailableSeverityLevel sets ServerUnavailableSeverityLevel field to given value.

### HasServerUnavailableSeverityLevel

`func (o *NumericGaugeResponse) HasServerUnavailableSeverityLevel() bool`

HasServerUnavailableSeverityLevel returns a boolean if a field has been set.

### GetServerDegradedSeverityLevel

`func (o *NumericGaugeResponse) GetServerDegradedSeverityLevel() EnumgaugeServerDegradedSeverityLevelProp`

GetServerDegradedSeverityLevel returns the ServerDegradedSeverityLevel field if non-nil, zero value otherwise.

### GetServerDegradedSeverityLevelOk

`func (o *NumericGaugeResponse) GetServerDegradedSeverityLevelOk() (*EnumgaugeServerDegradedSeverityLevelProp, bool)`

GetServerDegradedSeverityLevelOk returns a tuple with the ServerDegradedSeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDegradedSeverityLevel

`func (o *NumericGaugeResponse) SetServerDegradedSeverityLevel(v EnumgaugeServerDegradedSeverityLevelProp)`

SetServerDegradedSeverityLevel sets ServerDegradedSeverityLevel field to given value.

### HasServerDegradedSeverityLevel

`func (o *NumericGaugeResponse) HasServerDegradedSeverityLevel() bool`

HasServerDegradedSeverityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


