# BatchedTransactionsExtendedOperationHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn**](EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewBatchedTransactionsExtendedOperationHandlerResponse

`func NewBatchedTransactionsExtendedOperationHandlerResponse(schemas []EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn, enabled bool, ) *BatchedTransactionsExtendedOperationHandlerResponse`

NewBatchedTransactionsExtendedOperationHandlerResponse instantiates a new BatchedTransactionsExtendedOperationHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchedTransactionsExtendedOperationHandlerResponseWithDefaults

`func NewBatchedTransactionsExtendedOperationHandlerResponseWithDefaults() *BatchedTransactionsExtendedOperationHandlerResponse`

NewBatchedTransactionsExtendedOperationHandlerResponseWithDefaults instantiates a new BatchedTransactionsExtendedOperationHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetSchemas() []EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetSchemasOk() (*[]EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) SetSchemas(v []EnumbatchedTransactionsExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BatchedTransactionsExtendedOperationHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


