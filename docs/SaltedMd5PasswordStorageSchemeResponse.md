# SaltedMd5PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsaltedMd5PasswordStorageSchemeSchemaUrn**](EnumsaltedMd5PasswordStorageSchemeSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Salted MD5 Password Storage Scheme is enabled for use. | 
**SaltLengthBytes** | Pointer to **int32** | Specifies the number of bytes to use for the generated salt. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 

## Methods

### NewSaltedMd5PasswordStorageSchemeResponse

`func NewSaltedMd5PasswordStorageSchemeResponse(schemas []EnumsaltedMd5PasswordStorageSchemeSchemaUrn, enabled bool, ) *SaltedMd5PasswordStorageSchemeResponse`

NewSaltedMd5PasswordStorageSchemeResponse instantiates a new SaltedMd5PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaltedMd5PasswordStorageSchemeResponseWithDefaults

`func NewSaltedMd5PasswordStorageSchemeResponseWithDefaults() *SaltedMd5PasswordStorageSchemeResponse`

NewSaltedMd5PasswordStorageSchemeResponseWithDefaults instantiates a new SaltedMd5PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SaltedMd5PasswordStorageSchemeResponse) GetSchemas() []EnumsaltedMd5PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SaltedMd5PasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumsaltedMd5PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SaltedMd5PasswordStorageSchemeResponse) SetSchemas(v []EnumsaltedMd5PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *SaltedMd5PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SaltedMd5PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SaltedMd5PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSaltLengthBytes

`func (o *SaltedMd5PasswordStorageSchemeResponse) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *SaltedMd5PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *SaltedMd5PasswordStorageSchemeResponse) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.

### HasSaltLengthBytes

`func (o *SaltedMd5PasswordStorageSchemeResponse) HasSaltLengthBytes() bool`

HasSaltLengthBytes returns a boolean if a field has been set.

### GetDescription

`func (o *SaltedMd5PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaltedMd5PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaltedMd5PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaltedMd5PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


