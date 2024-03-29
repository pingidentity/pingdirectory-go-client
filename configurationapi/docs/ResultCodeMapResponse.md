# ResultCodeMapResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumresultCodeMapSchemaUrn**](EnumresultCodeMapSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Result Code Map | [optional] 
**BindAccountLockedResultCode** | Pointer to **int64** | Specifies the result code that should be returned if a bind attempt fails because the user&#39;s account is locked as a result of too many failed authentication attempts. | [optional] 
**BindMissingUserResultCode** | Pointer to **int64** | Specifies the result code that should be returned if a bind attempt fails because the target user entry does not exist in the server. | [optional] 
**BindMissingPasswordResultCode** | Pointer to **int64** | Specifies the result code that should be returned if a password-based bind attempt fails because the target user entry does not have a password. | [optional] 
**ServerErrorResultCode** | Pointer to **int64** | Specifies the result code that should be returned if a generic error occurs within the server. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Result Code Map | 

## Methods

### NewResultCodeMapResponse

`func NewResultCodeMapResponse(id string, ) *ResultCodeMapResponse`

NewResultCodeMapResponse instantiates a new ResultCodeMapResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultCodeMapResponseWithDefaults

`func NewResultCodeMapResponseWithDefaults() *ResultCodeMapResponse`

NewResultCodeMapResponseWithDefaults instantiates a new ResultCodeMapResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ResultCodeMapResponse) GetSchemas() []EnumresultCodeMapSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ResultCodeMapResponse) GetSchemasOk() (*[]EnumresultCodeMapSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ResultCodeMapResponse) SetSchemas(v []EnumresultCodeMapSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *ResultCodeMapResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *ResultCodeMapResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResultCodeMapResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResultCodeMapResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResultCodeMapResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBindAccountLockedResultCode

`func (o *ResultCodeMapResponse) GetBindAccountLockedResultCode() int64`

GetBindAccountLockedResultCode returns the BindAccountLockedResultCode field if non-nil, zero value otherwise.

### GetBindAccountLockedResultCodeOk

`func (o *ResultCodeMapResponse) GetBindAccountLockedResultCodeOk() (*int64, bool)`

GetBindAccountLockedResultCodeOk returns a tuple with the BindAccountLockedResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindAccountLockedResultCode

`func (o *ResultCodeMapResponse) SetBindAccountLockedResultCode(v int64)`

SetBindAccountLockedResultCode sets BindAccountLockedResultCode field to given value.

### HasBindAccountLockedResultCode

`func (o *ResultCodeMapResponse) HasBindAccountLockedResultCode() bool`

HasBindAccountLockedResultCode returns a boolean if a field has been set.

### GetBindMissingUserResultCode

`func (o *ResultCodeMapResponse) GetBindMissingUserResultCode() int64`

GetBindMissingUserResultCode returns the BindMissingUserResultCode field if non-nil, zero value otherwise.

### GetBindMissingUserResultCodeOk

`func (o *ResultCodeMapResponse) GetBindMissingUserResultCodeOk() (*int64, bool)`

GetBindMissingUserResultCodeOk returns a tuple with the BindMissingUserResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMissingUserResultCode

`func (o *ResultCodeMapResponse) SetBindMissingUserResultCode(v int64)`

SetBindMissingUserResultCode sets BindMissingUserResultCode field to given value.

### HasBindMissingUserResultCode

`func (o *ResultCodeMapResponse) HasBindMissingUserResultCode() bool`

HasBindMissingUserResultCode returns a boolean if a field has been set.

### GetBindMissingPasswordResultCode

`func (o *ResultCodeMapResponse) GetBindMissingPasswordResultCode() int64`

GetBindMissingPasswordResultCode returns the BindMissingPasswordResultCode field if non-nil, zero value otherwise.

### GetBindMissingPasswordResultCodeOk

`func (o *ResultCodeMapResponse) GetBindMissingPasswordResultCodeOk() (*int64, bool)`

GetBindMissingPasswordResultCodeOk returns a tuple with the BindMissingPasswordResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindMissingPasswordResultCode

`func (o *ResultCodeMapResponse) SetBindMissingPasswordResultCode(v int64)`

SetBindMissingPasswordResultCode sets BindMissingPasswordResultCode field to given value.

### HasBindMissingPasswordResultCode

`func (o *ResultCodeMapResponse) HasBindMissingPasswordResultCode() bool`

HasBindMissingPasswordResultCode returns a boolean if a field has been set.

### GetServerErrorResultCode

`func (o *ResultCodeMapResponse) GetServerErrorResultCode() int64`

GetServerErrorResultCode returns the ServerErrorResultCode field if non-nil, zero value otherwise.

### GetServerErrorResultCodeOk

`func (o *ResultCodeMapResponse) GetServerErrorResultCodeOk() (*int64, bool)`

GetServerErrorResultCodeOk returns a tuple with the ServerErrorResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerErrorResultCode

`func (o *ResultCodeMapResponse) SetServerErrorResultCode(v int64)`

SetServerErrorResultCode sets ServerErrorResultCode field to given value.

### HasServerErrorResultCode

`func (o *ResultCodeMapResponse) HasServerErrorResultCode() bool`

HasServerErrorResultCode returns a boolean if a field has been set.

### GetMeta

`func (o *ResultCodeMapResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ResultCodeMapResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ResultCodeMapResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ResultCodeMapResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ResultCodeMapResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ResultCodeMapResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ResultCodeMapResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ResultCodeMapResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ResultCodeMapResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResultCodeMapResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResultCodeMapResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


