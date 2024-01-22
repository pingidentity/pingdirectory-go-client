# AddArgon2idPasswordStorageSchemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumargon2idPasswordStorageSchemeSchemaUrn**](Enumargon2idPasswordStorageSchemeSchemaUrn.md) |  | 
**IterationCount** | **int64** | The number of rounds of cryptographic processing required in the course of encoding each password. | 
**ParallelismFactor** | **int64** | The number of concurrent threads that will be used in the course of encoding each password. | 
**MemoryUsageKb** | **int64** | The number of kilobytes of memory that must be used in the course of encoding each password. | 
**SaltLengthBytes** | **int64** | The number of bytes to use for the generated salt. | 
**DerivedKeyLengthBytes** | **int64** | The number of bytes to use for the derived key. The value must be greater than or equal to 8 and less than or equal to 512. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**SchemeName** | **string** | Name of the new Password Storage Scheme | 

## Methods

### NewAddArgon2idPasswordStorageSchemeRequest

`func NewAddArgon2idPasswordStorageSchemeRequest(schemas []Enumargon2idPasswordStorageSchemeSchemaUrn, iterationCount int64, parallelismFactor int64, memoryUsageKb int64, saltLengthBytes int64, derivedKeyLengthBytes int64, enabled bool, schemeName string, ) *AddArgon2idPasswordStorageSchemeRequest`

NewAddArgon2idPasswordStorageSchemeRequest instantiates a new AddArgon2idPasswordStorageSchemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddArgon2idPasswordStorageSchemeRequestWithDefaults

`func NewAddArgon2idPasswordStorageSchemeRequestWithDefaults() *AddArgon2idPasswordStorageSchemeRequest`

NewAddArgon2idPasswordStorageSchemeRequestWithDefaults instantiates a new AddArgon2idPasswordStorageSchemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetSchemas() []Enumargon2idPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetSchemasOk() (*[]Enumargon2idPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetSchemas(v []Enumargon2idPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIterationCount

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetIterationCount() int64`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetIterationCountOk() (*int64, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetIterationCount(v int64)`

SetIterationCount sets IterationCount field to given value.


### GetParallelismFactor

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetParallelismFactor() int64`

GetParallelismFactor returns the ParallelismFactor field if non-nil, zero value otherwise.

### GetParallelismFactorOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetParallelismFactorOk() (*int64, bool)`

GetParallelismFactorOk returns a tuple with the ParallelismFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelismFactor

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetParallelismFactor(v int64)`

SetParallelismFactor sets ParallelismFactor field to given value.


### GetMemoryUsageKb

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetMemoryUsageKb() int64`

GetMemoryUsageKb returns the MemoryUsageKb field if non-nil, zero value otherwise.

### GetMemoryUsageKbOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetMemoryUsageKbOk() (*int64, bool)`

GetMemoryUsageKbOk returns a tuple with the MemoryUsageKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageKb

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetMemoryUsageKb(v int64)`

SetMemoryUsageKb sets MemoryUsageKb field to given value.


### GetSaltLengthBytes

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetSaltLengthBytes() int64`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetSaltLengthBytesOk() (*int64, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetSaltLengthBytes(v int64)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.


### GetDerivedKeyLengthBytes

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetDerivedKeyLengthBytes() int64`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetDerivedKeyLengthBytesOk() (*int64, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetDerivedKeyLengthBytes(v int64)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.


### GetDescription

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddArgon2idPasswordStorageSchemeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSchemeName

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetSchemeName() string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *AddArgon2idPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *AddArgon2idPasswordStorageSchemeRequest) SetSchemeName(v string)`

SetSchemeName sets SchemeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


