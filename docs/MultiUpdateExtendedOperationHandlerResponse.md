# MultiUpdateExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummultiUpdateExtendedOperationHandlerSchemaUrn**](EnummultiUpdateExtendedOperationHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewMultiUpdateExtendedOperationHandlerResponse

`func NewMultiUpdateExtendedOperationHandlerResponse(schemas []EnummultiUpdateExtendedOperationHandlerSchemaUrn, enabled bool, ) *MultiUpdateExtendedOperationHandlerResponse`

NewMultiUpdateExtendedOperationHandlerResponse instantiates a new MultiUpdateExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiUpdateExtendedOperationHandlerResponseWithDefaults

`func NewMultiUpdateExtendedOperationHandlerResponseWithDefaults() *MultiUpdateExtendedOperationHandlerResponse`

NewMultiUpdateExtendedOperationHandlerResponseWithDefaults instantiates a new MultiUpdateExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetSchemas() []EnummultiUpdateExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnummultiUpdateExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetSchemas(v []EnummultiUpdateExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MultiUpdateExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MultiUpdateExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MultiUpdateExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


