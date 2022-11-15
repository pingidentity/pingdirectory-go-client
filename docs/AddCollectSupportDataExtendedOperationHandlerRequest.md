# AddCollectSupportDataExtendedOperationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Extended Operation Handler | 
**Schemas** | [**[]EnumcollectSupportDataExtendedOperationHandlerSchemaUrn**](EnumcollectSupportDataExtendedOperationHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 

## Methods

### NewAddCollectSupportDataExtendedOperationHandlerRequest

`func NewAddCollectSupportDataExtendedOperationHandlerRequest(handlerName string, schemas []EnumcollectSupportDataExtendedOperationHandlerSchemaUrn, enabled bool, ) *AddCollectSupportDataExtendedOperationHandlerRequest`

NewAddCollectSupportDataExtendedOperationHandlerRequest instantiates a new AddCollectSupportDataExtendedOperationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCollectSupportDataExtendedOperationHandlerRequestWithDefaults

`func NewAddCollectSupportDataExtendedOperationHandlerRequestWithDefaults() *AddCollectSupportDataExtendedOperationHandlerRequest`

NewAddCollectSupportDataExtendedOperationHandlerRequestWithDefaults instantiates a new AddCollectSupportDataExtendedOperationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) GetSchemas() []EnumcollectSupportDataExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) GetSchemasOk() (*[]EnumcollectSupportDataExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) SetSchemas(v []EnumcollectSupportDataExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCollectSupportDataExtendedOperationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


