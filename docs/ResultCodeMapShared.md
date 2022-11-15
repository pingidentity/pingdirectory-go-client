# ResultCodeMapShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumresultCodeMapSchemaUrn**](EnumresultCodeMapSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Result Code Map | [optional] 
**BindAccountLockedResultCode** | Pointer to **int32** | Specifies the result code that should be returned if a bind attempt fails because the user&#39;s account is locked as a result of too many failed authentication attempts. | [optional] 
**BindMissingUserResultCode** | Pointer to **int32** | Specifies the result code that should be returned if a bind attempt fails because the target user entry does not exist in the server. | [optional] 
**BindMissingPasswordResultCode** | Pointer to **int32** | Specifies the result code that should be returned if a password-based bind attempt fails because the target user entry does not have a password. | [optional] 
**ServerErrorResultCode** | Pointer to **int32** | Specifies the result code that should be returned if a generic error occurs within the server. | [optional] 

## Methods

### NewResultCodeMapShared

`func NewResultCodeMapShared() *ResultCodeMapShared`

NewResultCodeMapShared instantiates a new ResultCodeMapShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCodeMapSharedWithDefaults

`func NewResultCodeMapSharedWithDefaults() *ResultCodeMapShared`

NewResultCodeMapSharedWithDefaults instantiates a new ResultCodeMapShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ResultCodeMapShared) GetSchemas() []EnumresultCodeMapSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ResultCodeMapShared) GetSchemasOk() (*[]EnumresultCodeMapSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ResultCodeMapShared) SetSchemas(v []EnumresultCodeMapSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ResultCodeMapShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ResultCodeMapShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResultCodeMapShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResultCodeMapShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResultCodeMapShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBindAccountLockedResultCode

`func (o *ResultCodeMapShared) GetBindAccountLockedResultCode() int32`

GetBindAccountLockedResultCode returns the BindAccountLockedResultCode field if non-nil, zero value otherwise.

### GetBindAccountLockedResultCodeOk

`func (o *ResultCodeMapShared) GetBindAccountLockedResultCodeOk() (*int32, bool)`

GetBindAccountLockedResultCodeOk returns a tuple with the BindAccountLockedResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAccountLockedResultCode

`func (o *ResultCodeMapShared) SetBindAccountLockedResultCode(v int32)`

SetBindAccountLockedResultCode sets BindAccountLockedResultCode field to given value.

### HasBindAccountLockedResultCode

`func (o *ResultCodeMapShared) HasBindAccountLockedResultCode() bool`

HasBindAccountLockedResultCode returns a boolean if a field has been set.

### GetBindMissingUserResultCode

`func (o *ResultCodeMapShared) GetBindMissingUserResultCode() int32`

GetBindMissingUserResultCode returns the BindMissingUserResultCode field if non-nil, zero value otherwise.

### GetBindMissingUserResultCodeOk

`func (o *ResultCodeMapShared) GetBindMissingUserResultCodeOk() (*int32, bool)`

GetBindMissingUserResultCodeOk returns a tuple with the BindMissingUserResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMissingUserResultCode

`func (o *ResultCodeMapShared) SetBindMissingUserResultCode(v int32)`

SetBindMissingUserResultCode sets BindMissingUserResultCode field to given value.

### HasBindMissingUserResultCode

`func (o *ResultCodeMapShared) HasBindMissingUserResultCode() bool`

HasBindMissingUserResultCode returns a boolean if a field has been set.

### GetBindMissingPasswordResultCode

`func (o *ResultCodeMapShared) GetBindMissingPasswordResultCode() int32`

GetBindMissingPasswordResultCode returns the BindMissingPasswordResultCode field if non-nil, zero value otherwise.

### GetBindMissingPasswordResultCodeOk

`func (o *ResultCodeMapShared) GetBindMissingPasswordResultCodeOk() (*int32, bool)`

GetBindMissingPasswordResultCodeOk returns a tuple with the BindMissingPasswordResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMissingPasswordResultCode

`func (o *ResultCodeMapShared) SetBindMissingPasswordResultCode(v int32)`

SetBindMissingPasswordResultCode sets BindMissingPasswordResultCode field to given value.

### HasBindMissingPasswordResultCode

`func (o *ResultCodeMapShared) HasBindMissingPasswordResultCode() bool`

HasBindMissingPasswordResultCode returns a boolean if a field has been set.

### GetServerErrorResultCode

`func (o *ResultCodeMapShared) GetServerErrorResultCode() int32`

GetServerErrorResultCode returns the ServerErrorResultCode field if non-nil, zero value otherwise.

### GetServerErrorResultCodeOk

`func (o *ResultCodeMapShared) GetServerErrorResultCodeOk() (*int32, bool)`

GetServerErrorResultCodeOk returns a tuple with the ServerErrorResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerErrorResultCode

`func (o *ResultCodeMapShared) SetServerErrorResultCode(v int32)`

SetServerErrorResultCode sets ServerErrorResultCode field to given value.

### HasServerErrorResultCode

`func (o *ResultCodeMapShared) HasServerErrorResultCode() bool`

HasServerErrorResultCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


