# DelegatedAdminAttributeCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Delegated Admin Attribute Category | 
**Schemas** | Pointer to [**[]EnumdelegatedAdminAttributeCategorySchemaUrn**](EnumdelegatedAdminAttributeCategorySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Attribute Category | [optional] 
**DisplayName** | **string** | A human readable display name for this Delegated Admin Attribute Category. | 
**DisplayOrderIndex** | **int32** | Delegated Admin Attribute Categories are ordered for display based on this index from least to greatest. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewDelegatedAdminAttributeCategoryResponse

`func NewDelegatedAdminAttributeCategoryResponse(id string, displayName string, displayOrderIndex int32, ) *DelegatedAdminAttributeCategoryResponse`

NewDelegatedAdminAttributeCategoryResponse instantiates a new DelegatedAdminAttributeCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAdminAttributeCategoryResponseWithDefaults

`func NewDelegatedAdminAttributeCategoryResponseWithDefaults() *DelegatedAdminAttributeCategoryResponse`

NewDelegatedAdminAttributeCategoryResponseWithDefaults instantiates a new DelegatedAdminAttributeCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DelegatedAdminAttributeCategoryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegatedAdminAttributeCategoryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegatedAdminAttributeCategoryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DelegatedAdminAttributeCategoryResponse) GetSchemas() []EnumdelegatedAdminAttributeCategorySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelegatedAdminAttributeCategoryResponse) GetSchemasOk() (*[]EnumdelegatedAdminAttributeCategorySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelegatedAdminAttributeCategoryResponse) SetSchemas(v []EnumdelegatedAdminAttributeCategorySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DelegatedAdminAttributeCategoryResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *DelegatedAdminAttributeCategoryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelegatedAdminAttributeCategoryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelegatedAdminAttributeCategoryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelegatedAdminAttributeCategoryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *DelegatedAdminAttributeCategoryResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DelegatedAdminAttributeCategoryResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DelegatedAdminAttributeCategoryResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDisplayOrderIndex

`func (o *DelegatedAdminAttributeCategoryResponse) GetDisplayOrderIndex() int32`

GetDisplayOrderIndex returns the DisplayOrderIndex field if non-nil, zero value otherwise.

### GetDisplayOrderIndexOk

`func (o *DelegatedAdminAttributeCategoryResponse) GetDisplayOrderIndexOk() (*int32, bool)`

GetDisplayOrderIndexOk returns a tuple with the DisplayOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrderIndex

`func (o *DelegatedAdminAttributeCategoryResponse) SetDisplayOrderIndex(v int32)`

SetDisplayOrderIndex sets DisplayOrderIndex field to given value.


### GetMeta

`func (o *DelegatedAdminAttributeCategoryResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DelegatedAdminAttributeCategoryResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DelegatedAdminAttributeCategoryResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DelegatedAdminAttributeCategoryResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminAttributeCategoryResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DelegatedAdminAttributeCategoryResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminAttributeCategoryResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminAttributeCategoryResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


