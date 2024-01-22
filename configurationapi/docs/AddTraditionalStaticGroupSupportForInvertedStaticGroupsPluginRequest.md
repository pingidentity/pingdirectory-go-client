# AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn**](EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn.md) |  | 
**TraditionalStaticGroupObjectClass** | Pointer to [**EnumpluginTraditionalStaticGroupObjectClassProp**](EnumpluginTraditionalStaticGroupObjectClassProp.md) |  | [optional] 
**MaximumMembershipUpdatesPerModify** | Pointer to **int64** | An integer property that specifies the maximum number of membership changes that will be supported in a single modify operation. A value of zero indicates that modify operations targeting the group entry should not be permitted to alter the set of members for the group. | [optional] 
**ReadOperationSupport** | Pointer to [**EnumpluginReadOperationSupportProp**](EnumpluginReadOperationSupportProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**PluginName** | **string** | Name of the new Plugin | 

## Methods

### NewAddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest

`func NewAddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest(schemas []EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn, enabled bool, pluginName string, ) *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest`

NewAddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest instantiates a new AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequestWithDefaults

`func NewAddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequestWithDefaults() *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest`

NewAddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequestWithDefaults instantiates a new AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetSchemas() []EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetSchemasOk() (*[]EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) SetSchemas(v []EnumtraditionalStaticGroupSupportForInvertedStaticGroupsPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTraditionalStaticGroupObjectClass

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetTraditionalStaticGroupObjectClass() EnumpluginTraditionalStaticGroupObjectClassProp`

GetTraditionalStaticGroupObjectClass returns the TraditionalStaticGroupObjectClass field if non-nil, zero value otherwise.

### GetTraditionalStaticGroupObjectClassOk

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetTraditionalStaticGroupObjectClassOk() (*EnumpluginTraditionalStaticGroupObjectClassProp, bool)`

GetTraditionalStaticGroupObjectClassOk returns a tuple with the TraditionalStaticGroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraditionalStaticGroupObjectClass

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) SetTraditionalStaticGroupObjectClass(v EnumpluginTraditionalStaticGroupObjectClassProp)`

SetTraditionalStaticGroupObjectClass sets TraditionalStaticGroupObjectClass field to given value.

### HasTraditionalStaticGroupObjectClass

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) HasTraditionalStaticGroupObjectClass() bool`

HasTraditionalStaticGroupObjectClass returns a boolean if a field has been set.

### GetMaximumMembershipUpdatesPerModify

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetMaximumMembershipUpdatesPerModify() int64`

GetMaximumMembershipUpdatesPerModify returns the MaximumMembershipUpdatesPerModify field if non-nil, zero value otherwise.

### GetMaximumMembershipUpdatesPerModifyOk

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetMaximumMembershipUpdatesPerModifyOk() (*int64, bool)`

GetMaximumMembershipUpdatesPerModifyOk returns a tuple with the MaximumMembershipUpdatesPerModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMembershipUpdatesPerModify

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) SetMaximumMembershipUpdatesPerModify(v int64)`

SetMaximumMembershipUpdatesPerModify sets MaximumMembershipUpdatesPerModify field to given value.

### HasMaximumMembershipUpdatesPerModify

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) HasMaximumMembershipUpdatesPerModify() bool`

HasMaximumMembershipUpdatesPerModify returns a boolean if a field has been set.

### GetReadOperationSupport

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetReadOperationSupport() EnumpluginReadOperationSupportProp`

GetReadOperationSupport returns the ReadOperationSupport field if non-nil, zero value otherwise.

### GetReadOperationSupportOk

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetReadOperationSupportOk() (*EnumpluginReadOperationSupportProp, bool)`

GetReadOperationSupportOk returns a tuple with the ReadOperationSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOperationSupport

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) SetReadOperationSupport(v EnumpluginReadOperationSupportProp)`

SetReadOperationSupport sets ReadOperationSupport field to given value.

### HasReadOperationSupport

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) HasReadOperationSupport() bool`

HasReadOperationSupport returns a boolean if a field has been set.

### GetDescription

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetPluginName

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddTraditionalStaticGroupSupportForInvertedStaticGroupsPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


