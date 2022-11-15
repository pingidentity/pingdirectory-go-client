# ExecAlertHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumexecAlertHandlerSchemaUrn**](EnumexecAlertHandlerSchemaUrn.md) |  | 
**Command** | **string** | Specifies the path of the command to execute, without any arguments. It must be an absolute path for reasons of security and reliability. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Exec Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 

## Methods

### NewExecAlertHandlerShared

`func NewExecAlertHandlerShared(schemas []EnumexecAlertHandlerSchemaUrn, command string, enabled bool, ) *ExecAlertHandlerShared`

NewExecAlertHandlerShared instantiates a new ExecAlertHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecAlertHandlerSharedWithDefaults

`func NewExecAlertHandlerSharedWithDefaults() *ExecAlertHandlerShared`

NewExecAlertHandlerSharedWithDefaults instantiates a new ExecAlertHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ExecAlertHandlerShared) GetSchemas() []EnumexecAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExecAlertHandlerShared) GetSchemasOk() (*[]EnumexecAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExecAlertHandlerShared) SetSchemas(v []EnumexecAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetCommand

`func (o *ExecAlertHandlerShared) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *ExecAlertHandlerShared) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *ExecAlertHandlerShared) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetAsynchronous

`func (o *ExecAlertHandlerShared) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *ExecAlertHandlerShared) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *ExecAlertHandlerShared) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *ExecAlertHandlerShared) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetDescription

`func (o *ExecAlertHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExecAlertHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExecAlertHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExecAlertHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ExecAlertHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExecAlertHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExecAlertHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnabledAlertSeverity

`func (o *ExecAlertHandlerShared) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *ExecAlertHandlerShared) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *ExecAlertHandlerShared) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *ExecAlertHandlerShared) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *ExecAlertHandlerShared) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *ExecAlertHandlerShared) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *ExecAlertHandlerShared) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *ExecAlertHandlerShared) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *ExecAlertHandlerShared) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *ExecAlertHandlerShared) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *ExecAlertHandlerShared) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *ExecAlertHandlerShared) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


