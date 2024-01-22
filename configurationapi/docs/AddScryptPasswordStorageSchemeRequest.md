# AddScryptPasswordStorageSchemeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumscryptPasswordStorageSchemeSchemaUrn**](EnumscryptPasswordStorageSchemeSchemaUrn.md) |  | 
**ScryptCpuMemoryCostFactorExponent** | Pointer to **int64** | Specifies the exponent that should be used for the CPU/memory cost factor. The cost factor must be a power of two, so the value of this property represents the power to which two is raised. The CPU/memory cost factor specifies the number of iterations required for encoding the password, and also affects the amount of memory required during processing. A higher cost factor requires more processing and more memory to generate a password, which makes attacks against the password more expensive. | [optional] 
**ScryptBlockSize** | Pointer to **int64** | Specifies the block size for the digest that will be used in the course of encoding passwords. Increasing the block size while keeping the CPU/memory cost factor constant will increase the amount of memory required to encode a password, but it also increases the ratio of sequential memory access to random memory access (and sequential memory access is generally faster than random memory access). | [optional] 
**ScryptParallelizationParameter** | Pointer to **int64** | Specifies the number of times that scrypt has to perform the entire encoding process to produce the final result. | [optional] 
**MaxPasswordLength** | Pointer to **int64** | Specifies the maximum allowed length, in bytes, for passwords encoded with this scheme, which can help mitigate denial of service attacks from clients that attempt to bind with very long passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 
**SchemeName** | **string** | Name of the new Password Storage Scheme | 

## Methods

### NewAddScryptPasswordStorageSchemeRequest

`func NewAddScryptPasswordStorageSchemeRequest(schemas []EnumscryptPasswordStorageSchemeSchemaUrn, enabled bool, schemeName string, ) *AddScryptPasswordStorageSchemeRequest`

NewAddScryptPasswordStorageSchemeRequest instantiates a new AddScryptPasswordStorageSchemeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddScryptPasswordStorageSchemeRequestWithDefaults

`func NewAddScryptPasswordStorageSchemeRequestWithDefaults() *AddScryptPasswordStorageSchemeRequest`

NewAddScryptPasswordStorageSchemeRequestWithDefaults instantiates a new AddScryptPasswordStorageSchemeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddScryptPasswordStorageSchemeRequest) GetSchemas() []EnumscryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddScryptPasswordStorageSchemeRequest) GetSchemasOk() (*[]EnumscryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddScryptPasswordStorageSchemeRequest) SetSchemas(v []EnumscryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScryptCpuMemoryCostFactorExponent

`func (o *AddScryptPasswordStorageSchemeRequest) GetScryptCpuMemoryCostFactorExponent() int64`

GetScryptCpuMemoryCostFactorExponent returns the ScryptCpuMemoryCostFactorExponent field if non-nil, zero value otherwise.

### GetScryptCpuMemoryCostFactorExponentOk

`func (o *AddScryptPasswordStorageSchemeRequest) GetScryptCpuMemoryCostFactorExponentOk() (*int64, bool)`

GetScryptCpuMemoryCostFactorExponentOk returns a tuple with the ScryptCpuMemoryCostFactorExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptCpuMemoryCostFactorExponent

`func (o *AddScryptPasswordStorageSchemeRequest) SetScryptCpuMemoryCostFactorExponent(v int64)`

SetScryptCpuMemoryCostFactorExponent sets ScryptCpuMemoryCostFactorExponent field to given value.

### HasScryptCpuMemoryCostFactorExponent

`func (o *AddScryptPasswordStorageSchemeRequest) HasScryptCpuMemoryCostFactorExponent() bool`

HasScryptCpuMemoryCostFactorExponent returns a boolean if a field has been set.

### GetScryptBlockSize

`func (o *AddScryptPasswordStorageSchemeRequest) GetScryptBlockSize() int64`

GetScryptBlockSize returns the ScryptBlockSize field if non-nil, zero value otherwise.

### GetScryptBlockSizeOk

`func (o *AddScryptPasswordStorageSchemeRequest) GetScryptBlockSizeOk() (*int64, bool)`

GetScryptBlockSizeOk returns a tuple with the ScryptBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptBlockSize

`func (o *AddScryptPasswordStorageSchemeRequest) SetScryptBlockSize(v int64)`

SetScryptBlockSize sets ScryptBlockSize field to given value.

### HasScryptBlockSize

`func (o *AddScryptPasswordStorageSchemeRequest) HasScryptBlockSize() bool`

HasScryptBlockSize returns a boolean if a field has been set.

### GetScryptParallelizationParameter

`func (o *AddScryptPasswordStorageSchemeRequest) GetScryptParallelizationParameter() int64`

GetScryptParallelizationParameter returns the ScryptParallelizationParameter field if non-nil, zero value otherwise.

### GetScryptParallelizationParameterOk

`func (o *AddScryptPasswordStorageSchemeRequest) GetScryptParallelizationParameterOk() (*int64, bool)`

GetScryptParallelizationParameterOk returns a tuple with the ScryptParallelizationParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptParallelizationParameter

`func (o *AddScryptPasswordStorageSchemeRequest) SetScryptParallelizationParameter(v int64)`

SetScryptParallelizationParameter sets ScryptParallelizationParameter field to given value.

### HasScryptParallelizationParameter

`func (o *AddScryptPasswordStorageSchemeRequest) HasScryptParallelizationParameter() bool`

HasScryptParallelizationParameter returns a boolean if a field has been set.

### GetMaxPasswordLength

`func (o *AddScryptPasswordStorageSchemeRequest) GetMaxPasswordLength() int64`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *AddScryptPasswordStorageSchemeRequest) GetMaxPasswordLengthOk() (*int64, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *AddScryptPasswordStorageSchemeRequest) SetMaxPasswordLength(v int64)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *AddScryptPasswordStorageSchemeRequest) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetDescription

`func (o *AddScryptPasswordStorageSchemeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddScryptPasswordStorageSchemeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddScryptPasswordStorageSchemeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddScryptPasswordStorageSchemeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddScryptPasswordStorageSchemeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddScryptPasswordStorageSchemeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddScryptPasswordStorageSchemeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSchemeName

`func (o *AddScryptPasswordStorageSchemeRequest) GetSchemeName() string`

GetSchemeName returns the SchemeName field if non-nil, zero value otherwise.

### GetSchemeNameOk

`func (o *AddScryptPasswordStorageSchemeRequest) GetSchemeNameOk() (*string, bool)`

GetSchemeNameOk returns a tuple with the SchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemeName

`func (o *AddScryptPasswordStorageSchemeRequest) SetSchemeName(v string)`

SetSchemeName sets SchemeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


