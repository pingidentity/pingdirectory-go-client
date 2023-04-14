# AddBcryptPasswordStorageSchemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemeName** | **string** | Name of the new Password Storage Scheme | 
**Schemas** | [**[]EnumbcryptPasswordStorageSchemeSchemaUrn**](EnumbcryptPasswordStorageSchemeSchemaUrn.md) |  | 
**BcryptCostFactor** | Pointer to **int64** | Specifies the cost factor to use when encoding passwords with Bcrypt. A higher cost factor requires more processing to generate a password, which makes attacks against the password more expensive. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewAddBcryptPasswordStorageSchemeRequest

`func NewAddBcryptPasswordStorageSchemeRequest(schemeName string, schemas []EnumbcryptPasswordStorageSchemeSchemaUrn, enabled bool, ) *AddBcryptPasswordStorageSchemeRequest`

NewAddBcryptPasswordStorageSchemeRequest instantiates a new AddBcryptPasswordStorageSchemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBcryptPasswordStorageSchemeRequestWithDefaults

`func NewAddBcryptPasswordStorageSchemeRequestWithDefaults() *AddBcryptPasswordStorageSchemeRequest`

NewAddBcryptPasswordStorageSchemeRequestWithDefaults instantiates a new AddBcryptPasswordStorageSchemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemeName

`func (o *AddBcryptPasswordStorageSchemeRequest) GetSchemeName() string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *AddBcryptPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *AddBcryptPasswordStorageSchemeRequest) SetSchemeName(v string)`

SetSchemeName sets SchemeName field to given value.


### GetSchemas

`func (o *AddBcryptPasswordStorageSchemeRequest) GetSchemas() []EnumbcryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddBcryptPasswordStorageSchemeRequest) GetSchemasOk() (*[]EnumbcryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddBcryptPasswordStorageSchemeRequest) SetSchemas(v []EnumbcryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBcryptCostFactor

`func (o *AddBcryptPasswordStorageSchemeRequest) GetBcryptCostFactor() int64`

GetBcryptCostFactor returns the BcryptCostFactor field if non-nil, zero value otherwise.

### GetBcryptCostFactorOk

`func (o *AddBcryptPasswordStorageSchemeRequest) GetBcryptCostFactorOk() (*int64, bool)`

GetBcryptCostFactorOk returns a tuple with the BcryptCostFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcryptCostFactor

`func (o *AddBcryptPasswordStorageSchemeRequest) SetBcryptCostFactor(v int64)`

SetBcryptCostFactor sets BcryptCostFactor field to given value.

### HasBcryptCostFactor

`func (o *AddBcryptPasswordStorageSchemeRequest) HasBcryptCostFactor() bool`

HasBcryptCostFactor returns a boolean if a field has been set.

### GetDescription

`func (o *AddBcryptPasswordStorageSchemeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddBcryptPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddBcryptPasswordStorageSchemeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddBcryptPasswordStorageSchemeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddBcryptPasswordStorageSchemeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddBcryptPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddBcryptPasswordStorageSchemeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


