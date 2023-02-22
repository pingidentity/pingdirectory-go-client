# AddUniqueAttributePluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumuniqueAttributePluginSchemaUrn**](EnumuniqueAttributePluginSchemaUrn.md) |  | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**Type** | **[]string** | Specifies the type of attributes to check for value uniqueness. | 
**MultipleAttributeBehavior** | Pointer to [**EnumpluginMultipleAttributeBehaviorProp**](EnumpluginMultipleAttributeBehaviorProp.md) |  | [optional] 
**BaseDN** | Pointer to **[]string** | Specifies a base DN within which the attribute must be unique. | [optional] 
**PreventConflictsWithSoftDeletedEntries** | Pointer to **bool** | Indicates whether this Unique Attribute Plugin should reject a change that would result in one or more conflicts, even if those conflicts only exist in soft-deleted entries. | [optional] 
**Filter** | Pointer to **string** | Specifies the search filter to apply to determine if attribute uniqueness is enforced for the matching entries. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewAddUniqueAttributePluginRequest

`func NewAddUniqueAttributePluginRequest(pluginName string, schemas []EnumuniqueAttributePluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, type_ []string, enabled bool, ) *AddUniqueAttributePluginRequest`

NewAddUniqueAttributePluginRequest instantiates a new AddUniqueAttributePluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUniqueAttributePluginRequestWithDefaults

`func NewAddUniqueAttributePluginRequestWithDefaults() *AddUniqueAttributePluginRequest`

NewAddUniqueAttributePluginRequestWithDefaults instantiates a new AddUniqueAttributePluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddUniqueAttributePluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddUniqueAttributePluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddUniqueAttributePluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddUniqueAttributePluginRequest) GetSchemas() []EnumuniqueAttributePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddUniqueAttributePluginRequest) GetSchemasOk() (*[]EnumuniqueAttributePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddUniqueAttributePluginRequest) SetSchemas(v []EnumuniqueAttributePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AddUniqueAttributePluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddUniqueAttributePluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddUniqueAttributePluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetType

`func (o *AddUniqueAttributePluginRequest) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddUniqueAttributePluginRequest) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddUniqueAttributePluginRequest) SetType(v []string)`

SetType sets Type field to given value.


### GetMultipleAttributeBehavior

`func (o *AddUniqueAttributePluginRequest) GetMultipleAttributeBehavior() EnumpluginMultipleAttributeBehaviorProp`

GetMultipleAttributeBehavior returns the MultipleAttributeBehavior field if non-nil, zero value otherwise.

### GetMultipleAttributeBehaviorOk

`func (o *AddUniqueAttributePluginRequest) GetMultipleAttributeBehaviorOk() (*EnumpluginMultipleAttributeBehaviorProp, bool)`

GetMultipleAttributeBehaviorOk returns a tuple with the MultipleAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAttributeBehavior

`func (o *AddUniqueAttributePluginRequest) SetMultipleAttributeBehavior(v EnumpluginMultipleAttributeBehaviorProp)`

SetMultipleAttributeBehavior sets MultipleAttributeBehavior field to given value.

### HasMultipleAttributeBehavior

`func (o *AddUniqueAttributePluginRequest) HasMultipleAttributeBehavior() bool`

HasMultipleAttributeBehavior returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddUniqueAttributePluginRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddUniqueAttributePluginRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddUniqueAttributePluginRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddUniqueAttributePluginRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetPreventConflictsWithSoftDeletedEntries

`func (o *AddUniqueAttributePluginRequest) GetPreventConflictsWithSoftDeletedEntries() bool`

GetPreventConflictsWithSoftDeletedEntries returns the PreventConflictsWithSoftDeletedEntries field if non-nil, zero value otherwise.

### GetPreventConflictsWithSoftDeletedEntriesOk

`func (o *AddUniqueAttributePluginRequest) GetPreventConflictsWithSoftDeletedEntriesOk() (*bool, bool)`

GetPreventConflictsWithSoftDeletedEntriesOk returns a tuple with the PreventConflictsWithSoftDeletedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventConflictsWithSoftDeletedEntries

`func (o *AddUniqueAttributePluginRequest) SetPreventConflictsWithSoftDeletedEntries(v bool)`

SetPreventConflictsWithSoftDeletedEntries sets PreventConflictsWithSoftDeletedEntries field to given value.

### HasPreventConflictsWithSoftDeletedEntries

`func (o *AddUniqueAttributePluginRequest) HasPreventConflictsWithSoftDeletedEntries() bool`

HasPreventConflictsWithSoftDeletedEntries returns a boolean if a field has been set.

### GetFilter

`func (o *AddUniqueAttributePluginRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddUniqueAttributePluginRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddUniqueAttributePluginRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddUniqueAttributePluginRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDescription

`func (o *AddUniqueAttributePluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUniqueAttributePluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUniqueAttributePluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUniqueAttributePluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddUniqueAttributePluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUniqueAttributePluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUniqueAttributePluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddUniqueAttributePluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddUniqueAttributePluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddUniqueAttributePluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddUniqueAttributePluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


