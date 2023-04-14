# AddDelegatedAdminAttributeCategoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute Category. | 
**Schemas** | Pointer to [**[]EnumdelegatedAdminAttributeCategorySchemaUrn**](EnumdelegatedAdminAttributeCategorySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute Category | [optional] 
**DisplayOrderIndex** | **int64** | Delegated Admin Attribute Categories are ordered for display based on this index from least to greatest. | 

## Methods

### NewAddDelegatedAdminAttributeCategoryRequest

`func NewAddDelegatedAdminAttributeCategoryRequest(displayName string, displayOrderIndex int64, ) *AddDelegatedAdminAttributeCategoryRequest`

NewAddDelegatedAdminAttributeCategoryRequest instantiates a new AddDelegatedAdminAttributeCategoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDelegatedAdminAttributeCategoryRequestWithDefaults

`func NewAddDelegatedAdminAttributeCategoryRequestWithDefaults() *AddDelegatedAdminAttributeCategoryRequest`

NewAddDelegatedAdminAttributeCategoryRequestWithDefaults instantiates a new AddDelegatedAdminAttributeCategoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AddDelegatedAdminAttributeCategoryRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AddDelegatedAdminAttributeCategoryRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AddDelegatedAdminAttributeCategoryRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSchemas

`func (o *AddDelegatedAdminAttributeCategoryRequest) GetSchemas() []EnumdelegatedAdminAttributeCategorySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDelegatedAdminAttributeCategoryRequest) GetSchemasOk() (*[]EnumdelegatedAdminAttributeCategorySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDelegatedAdminAttributeCategoryRequest) SetSchemas(v []EnumdelegatedAdminAttributeCategorySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddDelegatedAdminAttributeCategoryRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddDelegatedAdminAttributeCategoryRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDelegatedAdminAttributeCategoryRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDelegatedAdminAttributeCategoryRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDelegatedAdminAttributeCategoryRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrderIndex

`func (o *AddDelegatedAdminAttributeCategoryRequest) GetDisplayOrderIndex() int64`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *AddDelegatedAdminAttributeCategoryRequest) GetDisplayOrderIndexOk() (*int64, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *AddDelegatedAdminAttributeCategoryRequest) SetDisplayOrderIndex(v int64)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


