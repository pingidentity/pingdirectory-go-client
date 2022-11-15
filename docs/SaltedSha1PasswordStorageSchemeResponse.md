# SaltedSha1PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsaltedSha1PasswordStorageSchemeSchemaUrn**](EnumsaltedSha1PasswordStorageSchemeSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Salted SHA1 Password Storage Scheme is enabled for use. | 
**SaltLengthBytes** | Pointer to **int32** | Specifies the number of bytes to use for the generated salt. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 

## Methods

### NewSaltedSha1PasswordStorageSchemeResponse

`func NewSaltedSha1PasswordStorageSchemeResponse(schemas []EnumsaltedSha1PasswordStorageSchemeSchemaUrn, enabled bool, ) *SaltedSha1PasswordStorageSchemeResponse`

NewSaltedSha1PasswordStorageSchemeResponse instantiates a new SaltedSha1PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaltedSha1PasswordStorageSchemeResponseWithDefaults

`func NewSaltedSha1PasswordStorageSchemeResponseWithDefaults() *SaltedSha1PasswordStorageSchemeResponse`

NewSaltedSha1PasswordStorageSchemeResponseWithDefaults instantiates a new SaltedSha1PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SaltedSha1PasswordStorageSchemeResponse) GetSchemas() []EnumsaltedSha1PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SaltedSha1PasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumsaltedSha1PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SaltedSha1PasswordStorageSchemeResponse) SetSchemas(v []EnumsaltedSha1PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *SaltedSha1PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SaltedSha1PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SaltedSha1PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSaltLengthBytes

`func (o *SaltedSha1PasswordStorageSchemeResponse) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *SaltedSha1PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *SaltedSha1PasswordStorageSchemeResponse) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.

### HasSaltLengthBytes

`func (o *SaltedSha1PasswordStorageSchemeResponse) HasSaltLengthBytes() bool`

HasSaltLengthBytes returns a boolean if a field has been set.

### GetDescription

`func (o *SaltedSha1PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaltedSha1PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaltedSha1PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaltedSha1PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


