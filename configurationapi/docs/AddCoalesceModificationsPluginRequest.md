# AddCoalesceModificationsPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumcoalesceModificationsPluginSchemaUrn**](EnumcoalesceModificationsPluginSchemaUrn.md) |  | 
**RequestCriteria** | **string** | A reference to request criteria that indicates which modify requests should be coalesced. | 
**AllowedRequestControl** | Pointer to **[]string** | Specifies the OIDs of the controls that are allowed to be present in operations to coalesce. These controls are passed through when the request is validated, but they will not be included when the background thread applies the coalesced modify requests. | [optional] 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 

## Methods

### NewAddCoalesceModificationsPluginRequest

`func NewAddCoalesceModificationsPluginRequest(pluginName string, schemas []EnumcoalesceModificationsPluginSchemaUrn, requestCriteria string, enabled bool, ) *AddCoalesceModificationsPluginRequest`

NewAddCoalesceModificationsPluginRequest instantiates a new AddCoalesceModificationsPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCoalesceModificationsPluginRequestWithDefaults

`func NewAddCoalesceModificationsPluginRequestWithDefaults() *AddCoalesceModificationsPluginRequest`

NewAddCoalesceModificationsPluginRequestWithDefaults instantiates a new AddCoalesceModificationsPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddCoalesceModificationsPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddCoalesceModificationsPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddCoalesceModificationsPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddCoalesceModificationsPluginRequest) GetSchemas() []EnumcoalesceModificationsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddCoalesceModificationsPluginRequest) GetSchemasOk() (*[]EnumcoalesceModificationsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddCoalesceModificationsPluginRequest) SetSchemas(v []EnumcoalesceModificationsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *AddCoalesceModificationsPluginRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddCoalesceModificationsPluginRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddCoalesceModificationsPluginRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.


### GetAllowedRequestControl

`func (o *AddCoalesceModificationsPluginRequest) GetAllowedRequestControl() []string`

GetAllowedRequestControl returns the AllowedRequestControl field if non-nil, zero value otherwise.

### GetAllowedRequestControlOk

`func (o *AddCoalesceModificationsPluginRequest) GetAllowedRequestControlOk() (*[]string, bool)`

GetAllowedRequestControlOk returns a tuple with the AllowedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequestControl

`func (o *AddCoalesceModificationsPluginRequest) SetAllowedRequestControl(v []string)`

SetAllowedRequestControl sets AllowedRequestControl field to given value.

### HasAllowedRequestControl

`func (o *AddCoalesceModificationsPluginRequest) HasAllowedRequestControl() bool`

HasAllowedRequestControl returns a boolean if a field has been set.

### GetInvokeForInternalOperations

`func (o *AddCoalesceModificationsPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddCoalesceModificationsPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddCoalesceModificationsPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddCoalesceModificationsPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetDescription

`func (o *AddCoalesceModificationsPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCoalesceModificationsPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCoalesceModificationsPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCoalesceModificationsPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCoalesceModificationsPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCoalesceModificationsPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCoalesceModificationsPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


