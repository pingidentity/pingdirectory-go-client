# Sha1PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumsha1PasswordStorageSchemeSchemaUrn**](Enumsha1PasswordStorageSchemeSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the SHA1 Password Storage Scheme is enabled for use. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 

## Methods

### NewSha1PasswordStorageSchemeResponse

`func NewSha1PasswordStorageSchemeResponse(schemas []Enumsha1PasswordStorageSchemeSchemaUrn, enabled bool, ) *Sha1PasswordStorageSchemeResponse`

NewSha1PasswordStorageSchemeResponse instantiates a new Sha1PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSha1PasswordStorageSchemeResponseWithDefaults

`func NewSha1PasswordStorageSchemeResponseWithDefaults() *Sha1PasswordStorageSchemeResponse`

NewSha1PasswordStorageSchemeResponseWithDefaults instantiates a new Sha1PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Sha1PasswordStorageSchemeResponse) GetSchemas() []Enumsha1PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Sha1PasswordStorageSchemeResponse) GetSchemasOk() (*[]Enumsha1PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Sha1PasswordStorageSchemeResponse) SetSchemas(v []Enumsha1PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *Sha1PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Sha1PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Sha1PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDescription

`func (o *Sha1PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Sha1PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Sha1PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Sha1PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


