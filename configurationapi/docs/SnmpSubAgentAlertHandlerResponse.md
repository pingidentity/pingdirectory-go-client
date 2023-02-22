# SnmpSubAgentAlertHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Alert Handler | 
**Schemas** | [**[]EnumsnmpSubAgentAlertHandlerSchemaUrn**](EnumsnmpSubAgentAlertHandlerSchemaUrn.md) |  | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this SNMP Sub Agent Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 

## Methods

### NewSnmpSubAgentAlertHandlerResponse

`func NewSnmpSubAgentAlertHandlerResponse(id string, schemas []EnumsnmpSubAgentAlertHandlerSchemaUrn, enabled bool, ) *SnmpSubAgentAlertHandlerResponse`

NewSnmpSubAgentAlertHandlerResponse instantiates a new SnmpSubAgentAlertHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnmpSubAgentAlertHandlerResponseWithDefaults

`func NewSnmpSubAgentAlertHandlerResponseWithDefaults() *SnmpSubAgentAlertHandlerResponse`

NewSnmpSubAgentAlertHandlerResponseWithDefaults instantiates a new SnmpSubAgentAlertHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SnmpSubAgentAlertHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SnmpSubAgentAlertHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SnmpSubAgentAlertHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SnmpSubAgentAlertHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SnmpSubAgentAlertHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SnmpSubAgentAlertHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SnmpSubAgentAlertHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *SnmpSubAgentAlertHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnmpSubAgentAlertHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SnmpSubAgentAlertHandlerResponse) GetSchemas() []EnumsnmpSubAgentAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetSchemasOk() (*[]EnumsnmpSubAgentAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SnmpSubAgentAlertHandlerResponse) SetSchemas(v []EnumsnmpSubAgentAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAsynchronous

`func (o *SnmpSubAgentAlertHandlerResponse) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *SnmpSubAgentAlertHandlerResponse) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *SnmpSubAgentAlertHandlerResponse) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetDescription

`func (o *SnmpSubAgentAlertHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SnmpSubAgentAlertHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SnmpSubAgentAlertHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SnmpSubAgentAlertHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SnmpSubAgentAlertHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledAlertSeverity

`func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *SnmpSubAgentAlertHandlerResponse) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *SnmpSubAgentAlertHandlerResponse) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *SnmpSubAgentAlertHandlerResponse) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *SnmpSubAgentAlertHandlerResponse) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *SnmpSubAgentAlertHandlerResponse) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *SnmpSubAgentAlertHandlerResponse) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *SnmpSubAgentAlertHandlerResponse) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *SnmpSubAgentAlertHandlerResponse) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


