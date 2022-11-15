# Argon2PasswordStorageSchemeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumargon2PasswordStorageSchemeSchemaUrn**](Enumargon2PasswordStorageSchemeSchemaUrn.md) |  | 
**IterationCount** | **int32** | The number of rounds of cryptographic processing required in the course of encoding each password. | 
**ParallelismFactor** | **int32** | The number of concurrent threads that will be used in the course of encoding each password. | 
**MemoryUsageKb** | **int32** | The number of kilobytes of memory that must be used in the course of encoding each password. | 
**SaltLengthBytes** | **int32** | The number of bytes to use for the generated salt. | 
**DerivedKeyLengthBytes** | **int32** | The number of bytes to use for the derived key. The value must be greater than or equal to 8 and less than or equal to 512. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewArgon2PasswordStorageSchemeShared

`func NewArgon2PasswordStorageSchemeShared(schemas []Enumargon2PasswordStorageSchemeSchemaUrn, iterationCount int32, parallelismFactor int32, memoryUsageKb int32, saltLengthBytes int32, derivedKeyLengthBytes int32, enabled bool, ) *Argon2PasswordStorageSchemeShared`

NewArgon2PasswordStorageSchemeShared instantiates a new Argon2PasswordStorageSchemeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgon2PasswordStorageSchemeSharedWithDefaults

`func NewArgon2PasswordStorageSchemeSharedWithDefaults() *Argon2PasswordStorageSchemeShared`

NewArgon2PasswordStorageSchemeSharedWithDefaults instantiates a new Argon2PasswordStorageSchemeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Argon2PasswordStorageSchemeShared) GetSchemas() []Enumargon2PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Argon2PasswordStorageSchemeShared) GetSchemasOk() (*[]Enumargon2PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Argon2PasswordStorageSchemeShared) SetSchemas(v []Enumargon2PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIterationCount

`func (o *Argon2PasswordStorageSchemeShared) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *Argon2PasswordStorageSchemeShared) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *Argon2PasswordStorageSchemeShared) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.


### GetParallelismFactor

`func (o *Argon2PasswordStorageSchemeShared) GetParallelismFactor() int32`

GetParallelismFactor returns the ParallelismFactor field if non-nil, zero value otherwise.

### GetParallelismFactorOk

`func (o *Argon2PasswordStorageSchemeShared) GetParallelismFactorOk() (*int32, bool)`

GetParallelismFactorOk returns a tuple with the ParallelismFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelismFactor

`func (o *Argon2PasswordStorageSchemeShared) SetParallelismFactor(v int32)`

SetParallelismFactor sets ParallelismFactor field to given value.


### GetMemoryUsageKb

`func (o *Argon2PasswordStorageSchemeShared) GetMemoryUsageKb() int32`

GetMemoryUsageKb returns the MemoryUsageKb field if non-nil, zero value otherwise.

### GetMemoryUsageKbOk

`func (o *Argon2PasswordStorageSchemeShared) GetMemoryUsageKbOk() (*int32, bool)`

GetMemoryUsageKbOk returns a tuple with the MemoryUsageKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageKb

`func (o *Argon2PasswordStorageSchemeShared) SetMemoryUsageKb(v int32)`

SetMemoryUsageKb sets MemoryUsageKb field to given value.


### GetSaltLengthBytes

`func (o *Argon2PasswordStorageSchemeShared) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *Argon2PasswordStorageSchemeShared) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *Argon2PasswordStorageSchemeShared) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.


### GetDerivedKeyLengthBytes

`func (o *Argon2PasswordStorageSchemeShared) GetDerivedKeyLengthBytes() int32`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *Argon2PasswordStorageSchemeShared) GetDerivedKeyLengthBytesOk() (*int32, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *Argon2PasswordStorageSchemeShared) SetDerivedKeyLengthBytes(v int32)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.


### GetDescription

`func (o *Argon2PasswordStorageSchemeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Argon2PasswordStorageSchemeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Argon2PasswordStorageSchemeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Argon2PasswordStorageSchemeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Argon2PasswordStorageSchemeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Argon2PasswordStorageSchemeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Argon2PasswordStorageSchemeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


