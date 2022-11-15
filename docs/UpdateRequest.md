# UpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | [**[]Operation**](Operation.md) |  | 

## Methods

### NewUpdateRequest

`func NewUpdateRequest(operations []Operation, ) *UpdateRequest`

NewUpdateRequest instantiates a new UpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRequestWithDefaults

`func NewUpdateRequestWithDefaults() *UpdateRequest`

NewUpdateRequestWithDefaults instantiates a new UpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperations

`func (o *UpdateRequest) GetOperations() []Operation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *UpdateRequest) GetOperationsOk() (*[]Operation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *UpdateRequest) SetOperations(v []Operation)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


