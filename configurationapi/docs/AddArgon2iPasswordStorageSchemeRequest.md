# AddArgon2iPasswordStorageSchemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SchemeName** | **string** | Name of the new Password Storage Scheme | 
**Schemas** | [**[]Enumargon2iPasswordStorageSchemeSchemaUrn**](Enumargon2iPasswordStorageSchemeSchemaUrn.md) |  | 
**IterationCount** | **int64** | The number of rounds of cryptographic processing required in the course of encoding each password. | 
**ParallelismFactor** | **int64** | The number of concurrent threads that will be used in the course of encoding each password. | 
**MemoryUsageKb** | **int64** | The number of kilobytes of memory that must be used in the course of encoding each password. | 
**SaltLengthBytes** | **int64** | The number of bytes to use for the generated salt. | 
**DerivedKeyLengthBytes** | **int64** | The number of bytes to use for the derived key. The value must be greater than or equal to 8 and less than or equal to 512. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewAddArgon2iPasswordStorageSchemeRequest

`func NewAddArgon2iPasswordStorageSchemeRequest(schemeName string, schemas []Enumargon2iPasswordStorageSchemeSchemaUrn, iterationCount int64, parallelismFactor int64, memoryUsageKb int64, saltLengthBytes int64, derivedKeyLengthBytes int64, enabled bool, ) *AddArgon2iPasswordStorageSchemeRequest`

NewAddArgon2iPasswordStorageSchemeRequest instantiates a new AddArgon2iPasswordStorageSchemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddArgon2iPasswordStorageSchemeRequestWithDefaults

`func NewAddArgon2iPasswordStorageSchemeRequestWithDefaults() *AddArgon2iPasswordStorageSchemeRequest`

NewAddArgon2iPasswordStorageSchemeRequestWithDefaults instantiates a new AddArgon2iPasswordStorageSchemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemeName

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetSchemeName() string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetSchemeName(v string)`

SetSchemeName sets SchemeName field to given value.


### GetSchemas

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetSchemas() []Enumargon2iPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetSchemasOk() (*[]Enumargon2iPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetSchemas(v []Enumargon2iPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIterationCount

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.


### GetParallelismFactor

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetParallelismFactor() int64`

GetParallelismFactor returns the ParallelismFactor field if non-nil, zero value otherwise.

### GetParallelismFactorOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetParallelismFactorOk() (*int64, bool)`

GetParallelismFactorOk returns a tuple with the ParallelismFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelismFactor

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetParallelismFactor(v int64)`

SetParallelismFactor sets ParallelismFactor field to given value.


### GetMemoryUsageKb

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetMemoryUsageKb() int64`

GetMemoryUsageKb returns the MemoryUsageKb field if non-nil, zero value otherwise.

### GetMemoryUsageKbOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetMemoryUsageKbOk() (*int64, bool)`

GetMemoryUsageKbOk returns a tuple with the MemoryUsageKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageKb

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetMemoryUsageKb(v int64)`

SetMemoryUsageKb sets MemoryUsageKb field to given value.


### GetSaltLengthBytes

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetSaltLengthBytes() int64`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetSaltLengthBytesOk() (*int64, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetSaltLengthBytes(v int64)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.


### GetDerivedKeyLengthBytes

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetDerivedKeyLengthBytes() int64`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetDerivedKeyLengthBytesOk() (*int64, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetDerivedKeyLengthBytes(v int64)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.


### GetDescription

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddArgon2iPasswordStorageSchemeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddArgon2iPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddArgon2iPasswordStorageSchemeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


