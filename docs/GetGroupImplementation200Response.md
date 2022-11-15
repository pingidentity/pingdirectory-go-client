# GetGroupImplementation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumdynamicGroupImplementationSchemaUrn**](EnumdynamicGroupImplementationSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Group Implementation | [optional] 
**Enabled** | **bool** | Indicates whether the Group Implementation is enabled. | 

## Methods

### NewGetGroupImplementation200Response

`func NewGetGroupImplementation200Response(schemas []EnumdynamicGroupImplementationSchemaUrn, enabled bool, ) *GetGroupImplementation200Response`

NewGetGroupImplementation200Response instantiates a new GetGroupImplementation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupImplementation200ResponseWithDefaults

`func NewGetGroupImplementation200ResponseWithDefaults() *GetGroupImplementation200Response`

NewGetGroupImplementation200ResponseWithDefaults instantiates a new GetGroupImplementation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetGroupImplementation200Response) GetSchemas() []EnumdynamicGroupImplementationSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetGroupImplementation200Response) GetSchemasOk() (*[]EnumdynamicGroupImplementationSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetGroupImplementation200Response) SetSchemas(v []EnumdynamicGroupImplementationSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *GetGroupImplementation200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGroupImplementation200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGroupImplementation200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGroupImplementation200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetGroupImplementation200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetGroupImplementation200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetGroupImplementation200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


