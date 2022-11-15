# SaltedSha384PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsaltedSha384PasswordStorageSchemeSchemaUrn**](EnumsaltedSha384PasswordStorageSchemeSchemaUrn.md) |  | 
**SaltLengthBytes** | Pointer to **int32** | Specifies the number of bytes to use for the generated salt. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewSaltedSha384PasswordStorageSchemeResponse

`func NewSaltedSha384PasswordStorageSchemeResponse(schemas []EnumsaltedSha384PasswordStorageSchemeSchemaUrn, enabled bool, ) *SaltedSha384PasswordStorageSchemeResponse`

NewSaltedSha384PasswordStorageSchemeResponse instantiates a new SaltedSha384PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaltedSha384PasswordStorageSchemeResponseWithDefaults

`func NewSaltedSha384PasswordStorageSchemeResponseWithDefaults() *SaltedSha384PasswordStorageSchemeResponse`

NewSaltedSha384PasswordStorageSchemeResponseWithDefaults instantiates a new SaltedSha384PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SaltedSha384PasswordStorageSchemeResponse) GetSchemas() []EnumsaltedSha384PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SaltedSha384PasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumsaltedSha384PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SaltedSha384PasswordStorageSchemeResponse) SetSchemas(v []EnumsaltedSha384PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSaltLengthBytes

`func (o *SaltedSha384PasswordStorageSchemeResponse) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *SaltedSha384PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *SaltedSha384PasswordStorageSchemeResponse) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.

### HasSaltLengthBytes

`func (o *SaltedSha384PasswordStorageSchemeResponse) HasSaltLengthBytes() bool`

HasSaltLengthBytes returns a boolean if a field has been set.

### GetDescription

`func (o *SaltedSha384PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaltedSha384PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaltedSha384PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaltedSha384PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SaltedSha384PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SaltedSha384PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SaltedSha384PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


