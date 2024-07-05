# AddVerifyPasswordExtendedOperationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumverifyPasswordExtendedOperationHandlerSchemaUrn**](EnumverifyPasswordExtendedOperationHandlerSchemaUrn.md) |  | 
**RejectInsecureRequests** | Pointer to **bool** | Indicates whether the server should reject attempts to use this extended operation over an insecure connection. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**HandlerName** | **string** | Name of the new Extended Operation Handler | 

## Methods

### NewAddVerifyPasswordExtendedOperationHandlerRequest

`func NewAddVerifyPasswordExtendedOperationHandlerRequest(schemas []EnumverifyPasswordExtendedOperationHandlerSchemaUrn, enabled bool, handlerName string, ) *AddVerifyPasswordExtendedOperationHandlerRequest`

NewAddVerifyPasswordExtendedOperationHandlerRequest instantiates a new AddVerifyPasswordExtendedOperationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVerifyPasswordExtendedOperationHandlerRequestWithDefaults

`func NewAddVerifyPasswordExtendedOperationHandlerRequestWithDefaults() *AddVerifyPasswordExtendedOperationHandlerRequest`

NewAddVerifyPasswordExtendedOperationHandlerRequestWithDefaults instantiates a new AddVerifyPasswordExtendedOperationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetSchemas() []EnumverifyPasswordExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetSchemasOk() (*[]EnumverifyPasswordExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) SetSchemas(v []EnumverifyPasswordExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRejectInsecureRequests

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetRejectInsecureRequests() bool`

GetRejectInsecureRequests returns the RejectInsecureRequests field if non-nil, zero value otherwise.

### GetRejectInsecureRequestsOk

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetRejectInsecureRequestsOk() (*bool, bool)`

GetRejectInsecureRequestsOk returns a tuple with the RejectInsecureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectInsecureRequests

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) SetRejectInsecureRequests(v bool)`

SetRejectInsecureRequests sets RejectInsecureRequests field to given value.

### HasRejectInsecureRequests

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) HasRejectInsecureRequests() bool`

HasRejectInsecureRequests returns a boolean if a field has been set.

### GetDescription

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHandlerName

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddVerifyPasswordExtendedOperationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


