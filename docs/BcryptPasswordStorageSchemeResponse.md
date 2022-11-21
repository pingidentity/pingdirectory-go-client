# BcryptPasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Storage Scheme | 
**Schemas** | [**[]EnumbcryptPasswordStorageSchemeSchemaUrn**](EnumbcryptPasswordStorageSchemeSchemaUrn.md) |  | 
**BcryptCostFactor** | Pointer to **int32** | Specifies the cost factor to use when encoding passwords with Bcrypt. A higher cost factor requires more processing to generate a password, which makes attacks against the password more expensive. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewBcryptPasswordStorageSchemeResponse

`func NewBcryptPasswordStorageSchemeResponse(id string, schemas []EnumbcryptPasswordStorageSchemeSchemaUrn, enabled bool, ) *BcryptPasswordStorageSchemeResponse`

NewBcryptPasswordStorageSchemeResponse instantiates a new BcryptPasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcryptPasswordStorageSchemeResponseWithDefaults

`func NewBcryptPasswordStorageSchemeResponseWithDefaults() *BcryptPasswordStorageSchemeResponse`

NewBcryptPasswordStorageSchemeResponseWithDefaults instantiates a new BcryptPasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcryptPasswordStorageSchemeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcryptPasswordStorageSchemeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcryptPasswordStorageSchemeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *BcryptPasswordStorageSchemeResponse) GetSchemas() []EnumbcryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BcryptPasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumbcryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BcryptPasswordStorageSchemeResponse) SetSchemas(v []EnumbcryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBcryptCostFactor

`func (o *BcryptPasswordStorageSchemeResponse) GetBcryptCostFactor() int32`

GetBcryptCostFactor returns the BcryptCostFactor field if non-nil, zero value otherwise.

### GetBcryptCostFactorOk

`func (o *BcryptPasswordStorageSchemeResponse) GetBcryptCostFactorOk() (*int32, bool)`

GetBcryptCostFactorOk returns a tuple with the BcryptCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcryptCostFactor

`func (o *BcryptPasswordStorageSchemeResponse) SetBcryptCostFactor(v int32)`

SetBcryptCostFactor sets BcryptCostFactor field to given value.

### HasBcryptCostFactor

`func (o *BcryptPasswordStorageSchemeResponse) HasBcryptCostFactor() bool`

HasBcryptCostFactor returns a boolean if a field has been set.

### GetDescription

`func (o *BcryptPasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcryptPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcryptPasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcryptPasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BcryptPasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BcryptPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BcryptPasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *BcryptPasswordStorageSchemeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BcryptPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BcryptPasswordStorageSchemeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BcryptPasswordStorageSchemeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


