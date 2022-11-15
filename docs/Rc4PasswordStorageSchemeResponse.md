# Rc4PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]Enumrc4PasswordStorageSchemeSchemaUrn**](Enumrc4PasswordStorageSchemeSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the RC4 Password Storage Scheme is enabled for use. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 

## Methods

### NewRc4PasswordStorageSchemeResponse

`func NewRc4PasswordStorageSchemeResponse(schemas []Enumrc4PasswordStorageSchemeSchemaUrn, enabled bool, ) *Rc4PasswordStorageSchemeResponse`

NewRc4PasswordStorageSchemeResponse instantiates a new Rc4PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRc4PasswordStorageSchemeResponseWithDefaults

`func NewRc4PasswordStorageSchemeResponseWithDefaults() *Rc4PasswordStorageSchemeResponse`

NewRc4PasswordStorageSchemeResponseWithDefaults instantiates a new Rc4PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *Rc4PasswordStorageSchemeResponse) GetSchemas() []Enumrc4PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Rc4PasswordStorageSchemeResponse) GetSchemasOk() (*[]Enumrc4PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Rc4PasswordStorageSchemeResponse) SetSchemas(v []Enumrc4PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *Rc4PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Rc4PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Rc4PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDescription

`func (o *Rc4PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Rc4PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Rc4PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Rc4PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


