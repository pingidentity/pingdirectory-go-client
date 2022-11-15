# LockAccountFailureLockoutActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Failure Lockout Action | 
**Schemas** | [**[]EnumlockAccountFailureLockoutActionSchemaUrn**](EnumlockAccountFailureLockoutActionSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Failure Lockout Action | [optional] 

## Methods

### NewLockAccountFailureLockoutActionResponse

`func NewLockAccountFailureLockoutActionResponse(id string, schemas []EnumlockAccountFailureLockoutActionSchemaUrn, ) *LockAccountFailureLockoutActionResponse`

NewLockAccountFailureLockoutActionResponse instantiates a new LockAccountFailureLockoutActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockAccountFailureLockoutActionResponseWithDefaults

`func NewLockAccountFailureLockoutActionResponseWithDefaults() *LockAccountFailureLockoutActionResponse`

NewLockAccountFailureLockoutActionResponseWithDefaults instantiates a new LockAccountFailureLockoutActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LockAccountFailureLockoutActionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LockAccountFailureLockoutActionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LockAccountFailureLockoutActionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LockAccountFailureLockoutActionResponse) GetSchemas() []EnumlockAccountFailureLockoutActionSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LockAccountFailureLockoutActionResponse) GetSchemasOk() (*[]EnumlockAccountFailureLockoutActionSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LockAccountFailureLockoutActionResponse) SetSchemas(v []EnumlockAccountFailureLockoutActionSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *LockAccountFailureLockoutActionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LockAccountFailureLockoutActionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LockAccountFailureLockoutActionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LockAccountFailureLockoutActionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


