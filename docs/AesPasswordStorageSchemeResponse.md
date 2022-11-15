# AesPasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaesPasswordStorageSchemeSchemaUrn**](EnumaesPasswordStorageSchemeSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewAesPasswordStorageSchemeResponse

`func NewAesPasswordStorageSchemeResponse(schemas []EnumaesPasswordStorageSchemeSchemaUrn, enabled bool, ) *AesPasswordStorageSchemeResponse`

NewAesPasswordStorageSchemeResponse instantiates a new AesPasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAesPasswordStorageSchemeResponseWithDefaults

`func NewAesPasswordStorageSchemeResponseWithDefaults() *AesPasswordStorageSchemeResponse`

NewAesPasswordStorageSchemeResponseWithDefaults instantiates a new AesPasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AesPasswordStorageSchemeResponse) GetSchemas() []EnumaesPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AesPasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumaesPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AesPasswordStorageSchemeResponse) SetSchemas(v []EnumaesPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AesPasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AesPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AesPasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AesPasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AesPasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AesPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AesPasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


