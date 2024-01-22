# AddInvertedStaticGroupReferentialIntegrityPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn**](EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn.md) |  | 
**PreventAddingMembersToNonexistentGroups** | Pointer to **bool** | Indicates whether the server should prevent updates to user entries that attempt to add them as a member of an inverted static group that does not exist. | [optional] 
**PreventAddingGroupsAsInvertedStaticGroupMembers** | Pointer to **bool** | Indicates whether the server should prevent attempts to add a group as a regular member of an inverted static group. If the members of another group should be considered members of an inverted static group, then the other group should be added as a nested group rather than a regular member. | [optional] 
**PreventNestingNonexistentGroups** | Pointer to **bool** | Indicates whether the server should prevent updates to inverted static groups that add references to nested groups that don&#39;t exist. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**PluginName** | **string** | Name of the new Plugin | 

## Methods

### NewAddInvertedStaticGroupReferentialIntegrityPluginRequest

`func NewAddInvertedStaticGroupReferentialIntegrityPluginRequest(schemas []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn, enabled bool, pluginName string, ) *AddInvertedStaticGroupReferentialIntegrityPluginRequest`

NewAddInvertedStaticGroupReferentialIntegrityPluginRequest instantiates a new AddInvertedStaticGroupReferentialIntegrityPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInvertedStaticGroupReferentialIntegrityPluginRequestWithDefaults

`func NewAddInvertedStaticGroupReferentialIntegrityPluginRequestWithDefaults() *AddInvertedStaticGroupReferentialIntegrityPluginRequest`

NewAddInvertedStaticGroupReferentialIntegrityPluginRequestWithDefaults instantiates a new AddInvertedStaticGroupReferentialIntegrityPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetSchemas() []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetSchemasOk() (*[]EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) SetSchemas(v []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPreventAddingMembersToNonexistentGroups

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetPreventAddingMembersToNonexistentGroups() bool`

GetPreventAddingMembersToNonexistentGroups returns the PreventAddingMembersToNonexistentGroups field if non-nil, zero value otherwise.

### GetPreventAddingMembersToNonexistentGroupsOk

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetPreventAddingMembersToNonexistentGroupsOk() (*bool, bool)`

GetPreventAddingMembersToNonexistentGroupsOk returns a tuple with the PreventAddingMembersToNonexistentGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAddingMembersToNonexistentGroups

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) SetPreventAddingMembersToNonexistentGroups(v bool)`

SetPreventAddingMembersToNonexistentGroups sets PreventAddingMembersToNonexistentGroups field to given value.

### HasPreventAddingMembersToNonexistentGroups

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) HasPreventAddingMembersToNonexistentGroups() bool`

HasPreventAddingMembersToNonexistentGroups returns a boolean if a field has been set.

### GetPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetPreventAddingGroupsAsInvertedStaticGroupMembers() bool`

GetPreventAddingGroupsAsInvertedStaticGroupMembers returns the PreventAddingGroupsAsInvertedStaticGroupMembers field if non-nil, zero value otherwise.

### GetPreventAddingGroupsAsInvertedStaticGroupMembersOk

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetPreventAddingGroupsAsInvertedStaticGroupMembersOk() (*bool, bool)`

GetPreventAddingGroupsAsInvertedStaticGroupMembersOk returns a tuple with the PreventAddingGroupsAsInvertedStaticGroupMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) SetPreventAddingGroupsAsInvertedStaticGroupMembers(v bool)`

SetPreventAddingGroupsAsInvertedStaticGroupMembers sets PreventAddingGroupsAsInvertedStaticGroupMembers field to given value.

### HasPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) HasPreventAddingGroupsAsInvertedStaticGroupMembers() bool`

HasPreventAddingGroupsAsInvertedStaticGroupMembers returns a boolean if a field has been set.

### GetPreventNestingNonexistentGroups

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetPreventNestingNonexistentGroups() bool`

GetPreventNestingNonexistentGroups returns the PreventNestingNonexistentGroups field if non-nil, zero value otherwise.

### GetPreventNestingNonexistentGroupsOk

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetPreventNestingNonexistentGroupsOk() (*bool, bool)`

GetPreventNestingNonexistentGroupsOk returns a tuple with the PreventNestingNonexistentGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventNestingNonexistentGroups

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) SetPreventNestingNonexistentGroups(v bool)`

SetPreventNestingNonexistentGroups sets PreventNestingNonexistentGroups field to given value.

### HasPreventNestingNonexistentGroups

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) HasPreventNestingNonexistentGroups() bool`

HasPreventNestingNonexistentGroups returns a boolean if a field has been set.

### GetDescription

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetPluginName

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddInvertedStaticGroupReferentialIntegrityPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


