# AddResultCodeMapRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MapName** | **string** | Name of the new Result Code Map | 
**Schemas** | Pointer to [**[]EnumresultCodeMapSchemaUrn**](EnumresultCodeMapSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Result Code Map | [optional] 
**BindAccountLockedResultCode** | Pointer to **int64** | Specifies the result code that should be returned if a bind attempt fails because the user&#39;s account is locked as a result of too many failed authentication attempts. | [optional] 
**BindMissingUserResultCode** | Pointer to **int64** | Specifies the result code that should be returned if a bind attempt fails because the target user entry does not exist in the server. | [optional] 
**BindMissingPasswordResultCode** | Pointer to **int64** | Specifies the result code that should be returned if a password-based bind attempt fails because the target user entry does not have a password. | [optional] 
**ServerErrorResultCode** | Pointer to **int64** | Specifies the result code that should be returned if a generic error occurs within the server. | [optional] 

## Methods

### NewAddResultCodeMapRequest

`func NewAddResultCodeMapRequest(mapName string, ) *AddResultCodeMapRequest`

NewAddResultCodeMapRequest instantiates a new AddResultCodeMapRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddResultCodeMapRequestWithDefaults

`func NewAddResultCodeMapRequestWithDefaults() *AddResultCodeMapRequest`

NewAddResultCodeMapRequestWithDefaults instantiates a new AddResultCodeMapRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMapName

`func (o *AddResultCodeMapRequest) GetMapName() string`

GetMapName returns the MapName field if non-nil, zero value otherwise.

### GetMapNameOk

`func (o *AddResultCodeMapRequest) GetMapNameOk() (*string, bool)`

GetMapNameOk returns a tuple with the MapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapName

`func (o *AddResultCodeMapRequest) SetMapName(v string)`

SetMapName sets MapName field to given value.


### GetSchemas

`func (o *AddResultCodeMapRequest) GetSchemas() []EnumresultCodeMapSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddResultCodeMapRequest) GetSchemasOk() (*[]EnumresultCodeMapSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddResultCodeMapRequest) SetSchemas(v []EnumresultCodeMapSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddResultCodeMapRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddResultCodeMapRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddResultCodeMapRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddResultCodeMapRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddResultCodeMapRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBindAccountLockedResultCode

`func (o *AddResultCodeMapRequest) GetBindAccountLockedResultCode() int64`

GetBindAccountLockedResultCode returns the BindAccountLockedResultCode field if non-nil, zero value otherwise.

### GetBindAccountLockedResultCodeOk

`func (o *AddResultCodeMapRequest) GetBindAccountLockedResultCodeOk() (*int64, bool)`

GetBindAccountLockedResultCodeOk returns a tuple with the BindAccountLockedResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAccountLockedResultCode

`func (o *AddResultCodeMapRequest) SetBindAccountLockedResultCode(v int64)`

SetBindAccountLockedResultCode sets BindAccountLockedResultCode field to given value.

### HasBindAccountLockedResultCode

`func (o *AddResultCodeMapRequest) HasBindAccountLockedResultCode() bool`

HasBindAccountLockedResultCode returns a boolean if a field has been set.

### GetBindMissingUserResultCode

`func (o *AddResultCodeMapRequest) GetBindMissingUserResultCode() int64`

GetBindMissingUserResultCode returns the BindMissingUserResultCode field if non-nil, zero value otherwise.

### GetBindMissingUserResultCodeOk

`func (o *AddResultCodeMapRequest) GetBindMissingUserResultCodeOk() (*int64, bool)`

GetBindMissingUserResultCodeOk returns a tuple with the BindMissingUserResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMissingUserResultCode

`func (o *AddResultCodeMapRequest) SetBindMissingUserResultCode(v int64)`

SetBindMissingUserResultCode sets BindMissingUserResultCode field to given value.

### HasBindMissingUserResultCode

`func (o *AddResultCodeMapRequest) HasBindMissingUserResultCode() bool`

HasBindMissingUserResultCode returns a boolean if a field has been set.

### GetBindMissingPasswordResultCode

`func (o *AddResultCodeMapRequest) GetBindMissingPasswordResultCode() int64`

GetBindMissingPasswordResultCode returns the BindMissingPasswordResultCode field if non-nil, zero value otherwise.

### GetBindMissingPasswordResultCodeOk

`func (o *AddResultCodeMapRequest) GetBindMissingPasswordResultCodeOk() (*int64, bool)`

GetBindMissingPasswordResultCodeOk returns a tuple with the BindMissingPasswordResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMissingPasswordResultCode

`func (o *AddResultCodeMapRequest) SetBindMissingPasswordResultCode(v int64)`

SetBindMissingPasswordResultCode sets BindMissingPasswordResultCode field to given value.

### HasBindMissingPasswordResultCode

`func (o *AddResultCodeMapRequest) HasBindMissingPasswordResultCode() bool`

HasBindMissingPasswordResultCode returns a boolean if a field has been set.

### GetServerErrorResultCode

`func (o *AddResultCodeMapRequest) GetServerErrorResultCode() int64`

GetServerErrorResultCode returns the ServerErrorResultCode field if non-nil, zero value otherwise.

### GetServerErrorResultCodeOk

`func (o *AddResultCodeMapRequest) GetServerErrorResultCodeOk() (*int64, bool)`

GetServerErrorResultCodeOk returns a tuple with the ServerErrorResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerErrorResultCode

`func (o *AddResultCodeMapRequest) SetServerErrorResultCode(v int64)`

SetServerErrorResultCode sets ServerErrorResultCode field to given value.

### HasServerErrorResultCode

`func (o *AddResultCodeMapRequest) HasServerErrorResultCode() bool`

HasServerErrorResultCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


