# ClearPasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumclearPasswordStorageSchemeSchemaUrn**](EnumclearPasswordStorageSchemeSchemaUrn.md) |  | 
**Enabled** | **bool** | Indicates whether the Clear Password Storage Scheme is enabled for use. | 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 

## Methods

### NewClearPasswordStorageSchemeResponse

`func NewClearPasswordStorageSchemeResponse(schemas []EnumclearPasswordStorageSchemeSchemaUrn, enabled bool, ) *ClearPasswordStorageSchemeResponse`

NewClearPasswordStorageSchemeResponse instantiates a new ClearPasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearPasswordStorageSchemeResponseWithDefaults

`func NewClearPasswordStorageSchemeResponseWithDefaults() *ClearPasswordStorageSchemeResponse`

NewClearPasswordStorageSchemeResponseWithDefaults instantiates a new ClearPasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ClearPasswordStorageSchemeResponse) GetSchemas() []EnumclearPasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ClearPasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumclearPasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ClearPasswordStorageSchemeResponse) SetSchemas(v []EnumclearPasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetEnabled

`func (o *ClearPasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClearPasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClearPasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDescription

`func (o *ClearPasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClearPasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClearPasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClearPasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


