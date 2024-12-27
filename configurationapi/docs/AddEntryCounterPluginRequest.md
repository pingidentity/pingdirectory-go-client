# AddEntryCounterPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumentryCounterPluginSchemaUrn**](EnumentryCounterPluginSchemaUrn.md) |  | 
**TimeBetweenSearches** | Pointer to **string** | The length of time between internal searches used to identify entries that match the sets of search criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**PluginName** | **string** | Name of the new Plugin | 

## Methods

### NewAddEntryCounterPluginRequest

`func NewAddEntryCounterPluginRequest(schemas []EnumentryCounterPluginSchemaUrn, enabled bool, pluginName string, ) *AddEntryCounterPluginRequest`

NewAddEntryCounterPluginRequest instantiates a new AddEntryCounterPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEntryCounterPluginRequestWithDefaults

`func NewAddEntryCounterPluginRequestWithDefaults() *AddEntryCounterPluginRequest`

NewAddEntryCounterPluginRequestWithDefaults instantiates a new AddEntryCounterPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddEntryCounterPluginRequest) GetSchemas() []EnumentryCounterPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddEntryCounterPluginRequest) GetSchemasOk() (*[]EnumentryCounterPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddEntryCounterPluginRequest) SetSchemas(v []EnumentryCounterPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTimeBetweenSearches

`func (o *AddEntryCounterPluginRequest) GetTimeBetweenSearches() string`

GetTimeBetweenSearches returns the TimeBetweenSearches field if non-nil, zero value otherwise.

### GetTimeBetweenSearchesOk

`func (o *AddEntryCounterPluginRequest) GetTimeBetweenSearchesOk() (*string, bool)`

GetTimeBetweenSearchesOk returns a tuple with the TimeBetweenSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBetweenSearches

`func (o *AddEntryCounterPluginRequest) SetTimeBetweenSearches(v string)`

SetTimeBetweenSearches sets TimeBetweenSearches field to given value.

### HasTimeBetweenSearches

`func (o *AddEntryCounterPluginRequest) HasTimeBetweenSearches() bool`

HasTimeBetweenSearches returns a boolean if a field has been set.

### GetDescription

`func (o *AddEntryCounterPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddEntryCounterPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddEntryCounterPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddEntryCounterPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddEntryCounterPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddEntryCounterPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddEntryCounterPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddEntryCounterPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddEntryCounterPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddEntryCounterPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddEntryCounterPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetPluginName

`func (o *AddEntryCounterPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddEntryCounterPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddEntryCounterPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


