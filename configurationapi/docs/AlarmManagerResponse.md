# AlarmManagerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | Pointer to [**[]EnumalarmManagerSchemaUrn**](EnumalarmManagerSchemaUrn.md) |  | [optional] 
**DefaultGaugeAlertLevel** | [**EnumalarmManagerDefaultGaugeAlertLevelProp**](EnumalarmManagerDefaultGaugeAlertLevelProp.md) |  | 
**GeneratedAlertTypes** | [**[]EnumalarmManagerGeneratedAlertTypesProp**](EnumalarmManagerGeneratedAlertTypesProp.md) |  | 
**SuppressedAlarm** | Pointer to [**[]EnumalarmManagerSuppressedAlarmProp**](EnumalarmManagerSuppressedAlarmProp.md) |  | [optional] 

## Methods

### NewAlarmManagerResponse

`func NewAlarmManagerResponse(defaultGaugeAlertLevel EnumalarmManagerDefaultGaugeAlertLevelProp, generatedAlertTypes []EnumalarmManagerGeneratedAlertTypesProp, ) *AlarmManagerResponse`

NewAlarmManagerResponse instantiates a new AlarmManagerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmManagerResponseWithDefaults

`func NewAlarmManagerResponseWithDefaults() *AlarmManagerResponse`

NewAlarmManagerResponseWithDefaults instantiates a new AlarmManagerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AlarmManagerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AlarmManagerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AlarmManagerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AlarmManagerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AlarmManagerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AlarmManagerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AlarmManagerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AlarmManagerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *AlarmManagerResponse) GetSchemas() []EnumalarmManagerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AlarmManagerResponse) GetSchemasOk() (*[]EnumalarmManagerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AlarmManagerResponse) SetSchemas(v []EnumalarmManagerSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AlarmManagerResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDefaultGaugeAlertLevel

`func (o *AlarmManagerResponse) GetDefaultGaugeAlertLevel() EnumalarmManagerDefaultGaugeAlertLevelProp`

GetDefaultGaugeAlertLevel returns the DefaultGaugeAlertLevel field if non-nil, zero value otherwise.

### GetDefaultGaugeAlertLevelOk

`func (o *AlarmManagerResponse) GetDefaultGaugeAlertLevelOk() (*EnumalarmManagerDefaultGaugeAlertLevelProp, bool)`

GetDefaultGaugeAlertLevelOk returns a tuple with the DefaultGaugeAlertLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGaugeAlertLevel

`func (o *AlarmManagerResponse) SetDefaultGaugeAlertLevel(v EnumalarmManagerDefaultGaugeAlertLevelProp)`

SetDefaultGaugeAlertLevel sets DefaultGaugeAlertLevel field to given value.


### GetGeneratedAlertTypes

`func (o *AlarmManagerResponse) GetGeneratedAlertTypes() []EnumalarmManagerGeneratedAlertTypesProp`

GetGeneratedAlertTypes returns the GeneratedAlertTypes field if non-nil, zero value otherwise.

### GetGeneratedAlertTypesOk

`func (o *AlarmManagerResponse) GetGeneratedAlertTypesOk() (*[]EnumalarmManagerGeneratedAlertTypesProp, bool)`

GetGeneratedAlertTypesOk returns a tuple with the GeneratedAlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAlertTypes

`func (o *AlarmManagerResponse) SetGeneratedAlertTypes(v []EnumalarmManagerGeneratedAlertTypesProp)`

SetGeneratedAlertTypes sets GeneratedAlertTypes field to given value.


### GetSuppressedAlarm

`func (o *AlarmManagerResponse) GetSuppressedAlarm() []EnumalarmManagerSuppressedAlarmProp`

GetSuppressedAlarm returns the SuppressedAlarm field if non-nil, zero value otherwise.

### GetSuppressedAlarmOk

`func (o *AlarmManagerResponse) GetSuppressedAlarmOk() (*[]EnumalarmManagerSuppressedAlarmProp, bool)`

GetSuppressedAlarmOk returns a tuple with the SuppressedAlarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressedAlarm

`func (o *AlarmManagerResponse) SetSuppressedAlarm(v []EnumalarmManagerSuppressedAlarmProp)`

SetSuppressedAlarm sets SuppressedAlarm field to given value.

### HasSuppressedAlarm

`func (o *AlarmManagerResponse) HasSuppressedAlarm() bool`

HasSuppressedAlarm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


