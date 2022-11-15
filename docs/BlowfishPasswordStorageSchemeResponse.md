# BlowfishPasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumblowfishPasswordStorageSchemeSchemaUrn**](EnumblowfishPasswordStorageSchemeSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewBlowfishPasswordStorageSchemeResponse

`func NewBlowfishPasswordStorageSchemeResponse(schemas []EnumblowfishPasswordStorageSchemeSchemaUrn, enabled bool, ) *BlowfishPasswordStorageSchemeResponse`

NewBlowfishPasswordStorageSchemeResponse instantiates a new BlowfishPasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlowfishPasswordStorageSchemeResponseWithDefaults

`func NewBlowfishPasswordStorageSchemeResponseWithDefaults() *BlowfishPasswordStorageSchemeResponse`

NewBlowfishPasswordStorageSchemeResponseWithDefaults instantiates a new BlowfishPasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BlowfishPasswordStorageSchemeResponse) GetSchemas() []EnumblowfishPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BlowfishPasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumblowfishPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BlowfishPasswordStorageSchemeResponse) SetSchemas(v []EnumblowfishPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *BlowfishPasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlowfishPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlowfishPasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlowfishPasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BlowfishPasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BlowfishPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BlowfishPasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


