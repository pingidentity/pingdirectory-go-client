# Md5PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enummd5PasswordStorageSchemeSchemaUrn**](Enummd5PasswordStorageSchemeSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the MD5 Password Storage Scheme is enabled for use. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 

## Methods

### NewMd5PasswordStorageSchemeResponse

`func NewMd5PasswordStorageSchemeResponse(schemas []Enummd5PasswordStorageSchemeSchemaUrn, enabled bool, ) *Md5PasswordStorageSchemeResponse`

NewMd5PasswordStorageSchemeResponse instantiates a new Md5PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMd5PasswordStorageSchemeResponseWithDefaults

`func NewMd5PasswordStorageSchemeResponseWithDefaults() *Md5PasswordStorageSchemeResponse`

NewMd5PasswordStorageSchemeResponseWithDefaults instantiates a new Md5PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Md5PasswordStorageSchemeResponse) GetSchemas() []Enummd5PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Md5PasswordStorageSchemeResponse) GetSchemasOk() (*[]Enummd5PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Md5PasswordStorageSchemeResponse) SetSchemas(v []Enummd5PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *Md5PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Md5PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Md5PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDescription

`func (o *Md5PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Md5PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Md5PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Md5PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


