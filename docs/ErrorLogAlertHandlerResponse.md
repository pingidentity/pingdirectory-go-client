# ErrorLogAlertHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Alert Handler | 
**Schemas** | [**[]EnumerrorLogAlertHandlerSchemaUrn**](EnumerrorLogAlertHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Alert Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Alert Handler is enabled. | 
**Asynchronous** | Pointer to **bool** | Indicates whether the server should attempt to invoke this Alert Handler in a background thread so that any potentially-expensive processing (e.g., performing network communication to deliver the alert notification) will not delay whatever processing the server was performing when the alert was generated. | [optional] 
**EnabledAlertSeverity** | Pointer to [**[]EnumalertHandlerEnabledAlertSeverityProp**](EnumalertHandlerEnabledAlertSeverityProp.md) |  | [optional] 
**EnabledAlertType** | Pointer to [**[]EnumalertHandlerEnabledAlertTypeProp**](EnumalertHandlerEnabledAlertTypeProp.md) |  | [optional] 
**DisabledAlertType** | Pointer to [**[]EnumalertHandlerDisabledAlertTypeProp**](EnumalertHandlerDisabledAlertTypeProp.md) |  | [optional] 

## Methods

### NewErrorLogAlertHandlerResponse

`func NewErrorLogAlertHandlerResponse(id string, schemas []EnumerrorLogAlertHandlerSchemaUrn, enabled bool, ) *ErrorLogAlertHandlerResponse`

NewErrorLogAlertHandlerResponse instantiates a new ErrorLogAlertHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLogAlertHandlerResponseWithDefaults

`func NewErrorLogAlertHandlerResponseWithDefaults() *ErrorLogAlertHandlerResponse`

NewErrorLogAlertHandlerResponseWithDefaults instantiates a new ErrorLogAlertHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ErrorLogAlertHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorLogAlertHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorLogAlertHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ErrorLogAlertHandlerResponse) GetSchemas() []EnumerrorLogAlertHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ErrorLogAlertHandlerResponse) GetSchemasOk() (*[]EnumerrorLogAlertHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ErrorLogAlertHandlerResponse) SetSchemas(v []EnumerrorLogAlertHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *ErrorLogAlertHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorLogAlertHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorLogAlertHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorLogAlertHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ErrorLogAlertHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ErrorLogAlertHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ErrorLogAlertHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAsynchronous

`func (o *ErrorLogAlertHandlerResponse) GetAsynchronous() bool`

GetAsynchronous returns the Asynchronous field if non-nil, zero value otherwise.

### GetAsynchronousOk

`func (o *ErrorLogAlertHandlerResponse) GetAsynchronousOk() (*bool, bool)`

GetAsynchronousOk returns a tuple with the Asynchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsynchronous

`func (o *ErrorLogAlertHandlerResponse) SetAsynchronous(v bool)`

SetAsynchronous sets Asynchronous field to given value.

### HasAsynchronous

`func (o *ErrorLogAlertHandlerResponse) HasAsynchronous() bool`

HasAsynchronous returns a boolean if a field has been set.

### GetEnabledAlertSeverity

`func (o *ErrorLogAlertHandlerResponse) GetEnabledAlertSeverity() []EnumalertHandlerEnabledAlertSeverityProp`

GetEnabledAlertSeverity returns the EnabledAlertSeverity field if non-nil, zero value otherwise.

### GetEnabledAlertSeverityOk

`func (o *ErrorLogAlertHandlerResponse) GetEnabledAlertSeverityOk() (*[]EnumalertHandlerEnabledAlertSeverityProp, bool)`

GetEnabledAlertSeverityOk returns a tuple with the EnabledAlertSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertSeverity

`func (o *ErrorLogAlertHandlerResponse) SetEnabledAlertSeverity(v []EnumalertHandlerEnabledAlertSeverityProp)`

SetEnabledAlertSeverity sets EnabledAlertSeverity field to given value.

### HasEnabledAlertSeverity

`func (o *ErrorLogAlertHandlerResponse) HasEnabledAlertSeverity() bool`

HasEnabledAlertSeverity returns a boolean if a field has been set.

### GetEnabledAlertType

`func (o *ErrorLogAlertHandlerResponse) GetEnabledAlertType() []EnumalertHandlerEnabledAlertTypeProp`

GetEnabledAlertType returns the EnabledAlertType field if non-nil, zero value otherwise.

### GetEnabledAlertTypeOk

`func (o *ErrorLogAlertHandlerResponse) GetEnabledAlertTypeOk() (*[]EnumalertHandlerEnabledAlertTypeProp, bool)`

GetEnabledAlertTypeOk returns a tuple with the EnabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAlertType

`func (o *ErrorLogAlertHandlerResponse) SetEnabledAlertType(v []EnumalertHandlerEnabledAlertTypeProp)`

SetEnabledAlertType sets EnabledAlertType field to given value.

### HasEnabledAlertType

`func (o *ErrorLogAlertHandlerResponse) HasEnabledAlertType() bool`

HasEnabledAlertType returns a boolean if a field has been set.

### GetDisabledAlertType

`func (o *ErrorLogAlertHandlerResponse) GetDisabledAlertType() []EnumalertHandlerDisabledAlertTypeProp`

GetDisabledAlertType returns the DisabledAlertType field if non-nil, zero value otherwise.

### GetDisabledAlertTypeOk

`func (o *ErrorLogAlertHandlerResponse) GetDisabledAlertTypeOk() (*[]EnumalertHandlerDisabledAlertTypeProp, bool)`

GetDisabledAlertTypeOk returns a tuple with the DisabledAlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAlertType

`func (o *ErrorLogAlertHandlerResponse) SetDisabledAlertType(v []EnumalertHandlerDisabledAlertTypeProp)`

SetDisabledAlertType sets DisabledAlertType field to given value.

### HasDisabledAlertType

`func (o *ErrorLogAlertHandlerResponse) HasDisabledAlertType() bool`

HasDisabledAlertType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


