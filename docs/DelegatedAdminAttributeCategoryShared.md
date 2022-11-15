# DelegatedAdminAttributeCategoryShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumdelegatedAdminAttributeCategorySchemaUrn**](EnumdelegatedAdminAttributeCategorySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute Category | [optional] 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute Category. | 
**DisplayOrderIndex** | **int32** | Delegated Admin Attribute Categories are ordered for display based on this index from least to greatest. | 

## Methods

### NewDelegatedAdminAttributeCategoryShared

`func NewDelegatedAdminAttributeCategoryShared(displayName string, displayOrderIndex int32, ) *DelegatedAdminAttributeCategoryShared`

NewDelegatedAdminAttributeCategoryShared instantiates a new DelegatedAdminAttributeCategoryShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAdminAttributeCategorySharedWithDefaults

`func NewDelegatedAdminAttributeCategorySharedWithDefaults() *DelegatedAdminAttributeCategoryShared`

NewDelegatedAdminAttributeCategorySharedWithDefaults instantiates a new DelegatedAdminAttributeCategoryShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *DelegatedAdminAttributeCategoryShared) GetSchemas() []EnumdelegatedAdminAttributeCategorySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelegatedAdminAttributeCategoryShared) GetSchemasOk() (*[]EnumdelegatedAdminAttributeCategorySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelegatedAdminAttributeCategoryShared) SetSchemas(v []EnumdelegatedAdminAttributeCategorySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DelegatedAdminAttributeCategoryShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *DelegatedAdminAttributeCategoryShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelegatedAdminAttributeCategoryShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelegatedAdminAttributeCategoryShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelegatedAdminAttributeCategoryShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DelegatedAdminAttributeCategoryShared) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DelegatedAdminAttributeCategoryShared) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DelegatedAdminAttributeCategoryShared) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDisplayOrderIndex

`func (o *DelegatedAdminAttributeCategoryShared) GetDisplayOrderIndex() int32`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *DelegatedAdminAttributeCategoryShared) GetDisplayOrderIndexOk() (*int32, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *DelegatedAdminAttributeCategoryShared) SetDisplayOrderIndex(v int32)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


