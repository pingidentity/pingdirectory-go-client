# AddSubOperationTimingPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumsubOperationTimingPluginSchemaUrn**](EnumsubOperationTimingPluginSchemaUrn.md) |  | 
**PluginType** | Pointer to [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) | Specifies the set of plug-in types for the plug-in, which specifies the times at which the plug-in is invoked. | [optional] 
**RequestCriteria** | Pointer to **string** | Specifies a set of request criteria used to indicate that only operations for requests matching this criteria should be counted when aggregating timing data. | [optional] 
**NumMostExpensivePhasesShown** | Pointer to **int32** | This controls how many of the most expensive phases are included per operation type in the monitor entry. | [optional] 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewAddSubOperationTimingPluginRequest

`func NewAddSubOperationTimingPluginRequest(pluginName string, schemas []EnumsubOperationTimingPluginSchemaUrn, enabled bool, ) *AddSubOperationTimingPluginRequest`

NewAddSubOperationTimingPluginRequest instantiates a new AddSubOperationTimingPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSubOperationTimingPluginRequestWithDefaults

`func NewAddSubOperationTimingPluginRequestWithDefaults() *AddSubOperationTimingPluginRequest`

NewAddSubOperationTimingPluginRequestWithDefaults instantiates a new AddSubOperationTimingPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddSubOperationTimingPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddSubOperationTimingPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddSubOperationTimingPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddSubOperationTimingPluginRequest) GetSchemas() []EnumsubOperationTimingPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSubOperationTimingPluginRequest) GetSchemasOk() (*[]EnumsubOperationTimingPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSubOperationTimingPluginRequest) SetSchemas(v []EnumsubOperationTimingPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AddSubOperationTimingPluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddSubOperationTimingPluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddSubOperationTimingPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.

### HasPluginType

`func (o *AddSubOperationTimingPluginRequest) HasPluginType() bool`

HasPluginType returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddSubOperationTimingPluginRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddSubOperationTimingPluginRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddSubOperationTimingPluginRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddSubOperationTimingPluginRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetNumMostExpensivePhasesShown

`func (o *AddSubOperationTimingPluginRequest) GetNumMostExpensivePhasesShown() int32`

GetNumMostExpensivePhasesShown returns the NumMostExpensivePhasesShown field if non-nil, zero value otherwise.

### GetNumMostExpensivePhasesShownOk

`func (o *AddSubOperationTimingPluginRequest) GetNumMostExpensivePhasesShownOk() (*int32, bool)`

GetNumMostExpensivePhasesShownOk returns a tuple with the NumMostExpensivePhasesShown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMostExpensivePhasesShown

`func (o *AddSubOperationTimingPluginRequest) SetNumMostExpensivePhasesShown(v int32)`

SetNumMostExpensivePhasesShown sets NumMostExpensivePhasesShown field to given value.

### HasNumMostExpensivePhasesShown

`func (o *AddSubOperationTimingPluginRequest) HasNumMostExpensivePhasesShown() bool`

HasNumMostExpensivePhasesShown returns a boolean if a field has been set.

### GetInvokeForInternalOperations

`func (o *AddSubOperationTimingPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddSubOperationTimingPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddSubOperationTimingPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddSubOperationTimingPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetDescription

`func (o *AddSubOperationTimingPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSubOperationTimingPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSubOperationTimingPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSubOperationTimingPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSubOperationTimingPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSubOperationTimingPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSubOperationTimingPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


