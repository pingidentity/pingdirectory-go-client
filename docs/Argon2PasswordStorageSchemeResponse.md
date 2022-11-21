# Argon2PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Storage Scheme | 
**Schemas** | [**[]Enumargon2PasswordStorageSchemeSchemaUrn**](Enumargon2PasswordStorageSchemeSchemaUrn.md) |  | 
**IterationCount** | **int32** | The number of rounds of cryptographic processing required in the course of encoding each password. | 
**ParallelismFactor** | **int32** | The number of concurrent threads that will be used in the course of encoding each password. | 
**MemoryUsageKb** | **int32** | The number of kilobytes of memory that must be used in the course of encoding each password. | 
**SaltLengthBytes** | **int32** | The number of bytes to use for the generated salt. | 
**DerivedKeyLengthBytes** | **int32** | The number of bytes to use for the derived key. The value must be greater than or equal to 8 and less than or equal to 512. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewArgon2PasswordStorageSchemeResponse

`func NewArgon2PasswordStorageSchemeResponse(id string, schemas []Enumargon2PasswordStorageSchemeSchemaUrn, iterationCount int32, parallelismFactor int32, memoryUsageKb int32, saltLengthBytes int32, derivedKeyLengthBytes int32, enabled bool, ) *Argon2PasswordStorageSchemeResponse`

NewArgon2PasswordStorageSchemeResponse instantiates a new Argon2PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArgon2PasswordStorageSchemeResponseWithDefaults

`func NewArgon2PasswordStorageSchemeResponseWithDefaults() *Argon2PasswordStorageSchemeResponse`

NewArgon2PasswordStorageSchemeResponseWithDefaults instantiates a new Argon2PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Argon2PasswordStorageSchemeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Argon2PasswordStorageSchemeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Argon2PasswordStorageSchemeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *Argon2PasswordStorageSchemeResponse) GetSchemas() []Enumargon2PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Argon2PasswordStorageSchemeResponse) GetSchemasOk() (*[]Enumargon2PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Argon2PasswordStorageSchemeResponse) SetSchemas(v []Enumargon2PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIterationCount

`func (o *Argon2PasswordStorageSchemeResponse) GetIterationCount() int32`

GetIterationCount returns the IterationCount field if non-nil, zero value otherwise.

### GetIterationCountOk

`func (o *Argon2PasswordStorageSchemeResponse) GetIterationCountOk() (*int32, bool)`

GetIterationCountOk returns a tuple with the IterationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCount

`func (o *Argon2PasswordStorageSchemeResponse) SetIterationCount(v int32)`

SetIterationCount sets IterationCount field to given value.


### GetParallelismFactor

`func (o *Argon2PasswordStorageSchemeResponse) GetParallelismFactor() int32`

GetParallelismFactor returns the ParallelismFactor field if non-nil, zero value otherwise.

### GetParallelismFactorOk

`func (o *Argon2PasswordStorageSchemeResponse) GetParallelismFactorOk() (*int32, bool)`

GetParallelismFactorOk returns a tuple with the ParallelismFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelismFactor

`func (o *Argon2PasswordStorageSchemeResponse) SetParallelismFactor(v int32)`

SetParallelismFactor sets ParallelismFactor field to given value.


### GetMemoryUsageKb

`func (o *Argon2PasswordStorageSchemeResponse) GetMemoryUsageKb() int32`

GetMemoryUsageKb returns the MemoryUsageKb field if non-nil, zero value otherwise.

### GetMemoryUsageKbOk

`func (o *Argon2PasswordStorageSchemeResponse) GetMemoryUsageKbOk() (*int32, bool)`

GetMemoryUsageKbOk returns a tuple with the MemoryUsageKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsageKb

`func (o *Argon2PasswordStorageSchemeResponse) SetMemoryUsageKb(v int32)`

SetMemoryUsageKb sets MemoryUsageKb field to given value.


### GetSaltLengthBytes

`func (o *Argon2PasswordStorageSchemeResponse) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *Argon2PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *Argon2PasswordStorageSchemeResponse) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.


### GetDerivedKeyLengthBytes

`func (o *Argon2PasswordStorageSchemeResponse) GetDerivedKeyLengthBytes() int32`

GetDerivedKeyLengthBytes returns the DerivedKeyLengthBytes field if non-nil, zero value otherwise.

### GetDerivedKeyLengthBytesOk

`func (o *Argon2PasswordStorageSchemeResponse) GetDerivedKeyLengthBytesOk() (*int32, bool)`

GetDerivedKeyLengthBytesOk returns a tuple with the DerivedKeyLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedKeyLengthBytes

`func (o *Argon2PasswordStorageSchemeResponse) SetDerivedKeyLengthBytes(v int32)`

SetDerivedKeyLengthBytes sets DerivedKeyLengthBytes field to given value.


### GetDescription

`func (o *Argon2PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Argon2PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Argon2PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Argon2PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Argon2PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Argon2PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Argon2PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *Argon2PasswordStorageSchemeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Argon2PasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Argon2PasswordStorageSchemeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Argon2PasswordStorageSchemeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


