# BcryptPasswordStorageSchemeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumbcryptPasswordStorageSchemeSchemaUrn**](EnumbcryptPasswordStorageSchemeSchemaUrn.md) |  | 
**BcryptCostFactor** | Pointer to **int32** | Specifies the cost factor to use when encoding passwords with Bcrypt. A higher cost factor requires more processing to generate a password, which makes attacks against the password more expensive. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewBcryptPasswordStorageSchemeShared

`func NewBcryptPasswordStorageSchemeShared(schemas []EnumbcryptPasswordStorageSchemeSchemaUrn, enabled bool, ) *BcryptPasswordStorageSchemeShared`

NewBcryptPasswordStorageSchemeShared instantiates a new BcryptPasswordStorageSchemeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcryptPasswordStorageSchemeSharedWithDefaults

`func NewBcryptPasswordStorageSchemeSharedWithDefaults() *BcryptPasswordStorageSchemeShared`

NewBcryptPasswordStorageSchemeSharedWithDefaults instantiates a new BcryptPasswordStorageSchemeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BcryptPasswordStorageSchemeShared) GetSchemas() []EnumbcryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BcryptPasswordStorageSchemeShared) GetSchemasOk() (*[]EnumbcryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BcryptPasswordStorageSchemeShared) SetSchemas(v []EnumbcryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBcryptCostFactor

`func (o *BcryptPasswordStorageSchemeShared) GetBcryptCostFactor() int32`

GetBcryptCostFactor returns the BcryptCostFactor field if non-nil, zero value otherwise.

### GetBcryptCostFactorOk

`func (o *BcryptPasswordStorageSchemeShared) GetBcryptCostFactorOk() (*int32, bool)`

GetBcryptCostFactorOk returns a tuple with the BcryptCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcryptCostFactor

`func (o *BcryptPasswordStorageSchemeShared) SetBcryptCostFactor(v int32)`

SetBcryptCostFactor sets BcryptCostFactor field to given value.

### HasBcryptCostFactor

`func (o *BcryptPasswordStorageSchemeShared) HasBcryptCostFactor() bool`

HasBcryptCostFactor returns a boolean if a field has been set.

### GetDescription

`func (o *BcryptPasswordStorageSchemeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcryptPasswordStorageSchemeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcryptPasswordStorageSchemeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcryptPasswordStorageSchemeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BcryptPasswordStorageSchemeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BcryptPasswordStorageSchemeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BcryptPasswordStorageSchemeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


