# Base64PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumbase64PasswordStorageSchemeSchemaUrn**](Enumbase64PasswordStorageSchemeSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Base64 Password Storage Scheme is enabled for use. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 

## Methods

### NewBase64PasswordStorageSchemeResponse

`func NewBase64PasswordStorageSchemeResponse(schemas []Enumbase64PasswordStorageSchemeSchemaUrn, enabled bool, ) *Base64PasswordStorageSchemeResponse`

NewBase64PasswordStorageSchemeResponse instantiates a new Base64PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBase64PasswordStorageSchemeResponseWithDefaults

`func NewBase64PasswordStorageSchemeResponseWithDefaults() *Base64PasswordStorageSchemeResponse`

NewBase64PasswordStorageSchemeResponseWithDefaults instantiates a new Base64PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Base64PasswordStorageSchemeResponse) GetSchemas() []Enumbase64PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Base64PasswordStorageSchemeResponse) GetSchemasOk() (*[]Enumbase64PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Base64PasswordStorageSchemeResponse) SetSchemas(v []Enumbase64PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *Base64PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Base64PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Base64PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDescription

`func (o *Base64PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Base64PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Base64PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Base64PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


