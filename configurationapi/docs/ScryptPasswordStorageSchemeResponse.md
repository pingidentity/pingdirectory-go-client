# ScryptPasswordStorageSchemeResponse

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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Password Storage Scheme | 

## Methods

### NewScryptPasswordStorageSchemeResponse

`func NewScryptPasswordStorageSchemeResponse(schemas []EnumscryptPasswordStorageSchemeSchemaUrn, enabled bool, id string, ) *ScryptPasswordStorageSchemeResponse`

NewScryptPasswordStorageSchemeResponse instantiates a new ScryptPasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScryptPasswordStorageSchemeResponseWithDefaults

`func NewScryptPasswordStorageSchemeResponseWithDefaults() *ScryptPasswordStorageSchemeResponse`

NewScryptPasswordStorageSchemeResponseWithDefaults instantiates a new ScryptPasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScryptPasswordStorageSchemeResponse) GetSchemas() []EnumscryptPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScryptPasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumscryptPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScryptPasswordStorageSchemeResponse) SetSchemas(v []EnumscryptPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetScryptCpuMemoryCostFactorExponent

`func (o *ScryptPasswordStorageSchemeResponse) GetScryptCpuMemoryCostFactorExponent() int64`

GetScryptCpuMemoryCostFactorExponent returns the ScryptCpuMemoryCostFactorExponent field if non-nil, zero value otherwise.

### GetScryptCpuMemoryCostFactorExponentOk

`func (o *ScryptPasswordStorageSchemeResponse) GetScryptCpuMemoryCostFactorExponentOk() (*int64, bool)`

GetScryptCpuMemoryCostFactorExponentOk returns a tuple with the ScryptCpuMemoryCostFactorExponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptCpuMemoryCostFactorExponent

`func (o *ScryptPasswordStorageSchemeResponse) SetScryptCpuMemoryCostFactorExponent(v int64)`

SetScryptCpuMemoryCostFactorExponent sets ScryptCpuMemoryCostFactorExponent field to given value.

### HasScryptCpuMemoryCostFactorExponent

`func (o *ScryptPasswordStorageSchemeResponse) HasScryptCpuMemoryCostFactorExponent() bool`

HasScryptCpuMemoryCostFactorExponent returns a boolean if a field has been set.

### GetScryptBlockSize

`func (o *ScryptPasswordStorageSchemeResponse) GetScryptBlockSize() int64`

GetScryptBlockSize returns the ScryptBlockSize field if non-nil, zero value otherwise.

### GetScryptBlockSizeOk

`func (o *ScryptPasswordStorageSchemeResponse) GetScryptBlockSizeOk() (*int64, bool)`

GetScryptBlockSizeOk returns a tuple with the ScryptBlockSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptBlockSize

`func (o *ScryptPasswordStorageSchemeResponse) SetScryptBlockSize(v int64)`

SetScryptBlockSize sets ScryptBlockSize field to given value.

### HasScryptBlockSize

`func (o *ScryptPasswordStorageSchemeResponse) HasScryptBlockSize() bool`

HasScryptBlockSize returns a boolean if a field has been set.

### GetScryptParallelizationParameter

`func (o *ScryptPasswordStorageSchemeResponse) GetScryptParallelizationParameter() int64`

GetScryptParallelizationParameter returns the ScryptParallelizationParameter field if non-nil, zero value otherwise.

### GetScryptParallelizationParameterOk

`func (o *ScryptPasswordStorageSchemeResponse) GetScryptParallelizationParameterOk() (*int64, bool)`

GetScryptParallelizationParameterOk returns a tuple with the ScryptParallelizationParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScryptParallelizationParameter

`func (o *ScryptPasswordStorageSchemeResponse) SetScryptParallelizationParameter(v int64)`

SetScryptParallelizationParameter sets ScryptParallelizationParameter field to given value.

### HasScryptParallelizationParameter

`func (o *ScryptPasswordStorageSchemeResponse) HasScryptParallelizationParameter() bool`

HasScryptParallelizationParameter returns a boolean if a field has been set.

### GetMaxPasswordLength

`func (o *ScryptPasswordStorageSchemeResponse) GetMaxPasswordLength() int64`

GetMaxPasswordLength returns the MaxPasswordLength field if non-nil, zero value otherwise.

### GetMaxPasswordLengthOk

`func (o *ScryptPasswordStorageSchemeResponse) GetMaxPasswordLengthOk() (*int64, bool)`

GetMaxPasswordLengthOk returns a tuple with the MaxPasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordLength

`func (o *ScryptPasswordStorageSchemeResponse) SetMaxPasswordLength(v int64)`

SetMaxPasswordLength sets MaxPasswordLength field to given value.

### HasMaxPasswordLength

`func (o *ScryptPasswordStorageSchemeResponse) HasMaxPasswordLength() bool`

HasMaxPasswordLength returns a boolean if a field has been set.

### GetDescription

`func (o *ScryptPasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScryptPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScryptPasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScryptPasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ScryptPasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ScryptPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ScryptPasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ScryptPasswordStorageSchemeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScryptPasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScryptPasswordStorageSchemeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScryptPasswordStorageSchemeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ScryptPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ScryptPasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ScryptPasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ScryptPasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *ScryptPasswordStorageSchemeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScryptPasswordStorageSchemeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScryptPasswordStorageSchemeResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


