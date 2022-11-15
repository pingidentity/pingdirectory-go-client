# CancelExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumcancelExtendedOperationHandlerSchemaUrn**](EnumcancelExtendedOperationHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewCancelExtendedOperationHandlerResponse

`func NewCancelExtendedOperationHandlerResponse(schemas []EnumcancelExtendedOperationHandlerSchemaUrn, enabled bool, ) *CancelExtendedOperationHandlerResponse`

NewCancelExtendedOperationHandlerResponse instantiates a new CancelExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelExtendedOperationHandlerResponseWithDefaults

`func NewCancelExtendedOperationHandlerResponseWithDefaults() *CancelExtendedOperationHandlerResponse`

NewCancelExtendedOperationHandlerResponseWithDefaults instantiates a new CancelExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *CancelExtendedOperationHandlerResponse) GetSchemas() []EnumcancelExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CancelExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumcancelExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CancelExtendedOperationHandlerResponse) SetSchemas(v []EnumcancelExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *CancelExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CancelExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CancelExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CancelExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CancelExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CancelExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CancelExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


