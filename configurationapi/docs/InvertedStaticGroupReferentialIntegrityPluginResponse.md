# InvertedStaticGroupReferentialIntegrityPluginResponse

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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Plugin | 

## Methods

### NewInvertedStaticGroupReferentialIntegrityPluginResponse

`func NewInvertedStaticGroupReferentialIntegrityPluginResponse(schemas []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn, enabled bool, id string, ) *InvertedStaticGroupReferentialIntegrityPluginResponse`

NewInvertedStaticGroupReferentialIntegrityPluginResponse instantiates a new InvertedStaticGroupReferentialIntegrityPluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvertedStaticGroupReferentialIntegrityPluginResponseWithDefaults

`func NewInvertedStaticGroupReferentialIntegrityPluginResponseWithDefaults() *InvertedStaticGroupReferentialIntegrityPluginResponse`

NewInvertedStaticGroupReferentialIntegrityPluginResponseWithDefaults instantiates a new InvertedStaticGroupReferentialIntegrityPluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetSchemas() []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetSchemasOk() (*[]EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetSchemas(v []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPreventAddingMembersToNonexistentGroups

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetPreventAddingMembersToNonexistentGroups() bool`

GetPreventAddingMembersToNonexistentGroups returns the PreventAddingMembersToNonexistentGroups field if non-nil, zero value otherwise.

### GetPreventAddingMembersToNonexistentGroupsOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetPreventAddingMembersToNonexistentGroupsOk() (*bool, bool)`

GetPreventAddingMembersToNonexistentGroupsOk returns a tuple with the PreventAddingMembersToNonexistentGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAddingMembersToNonexistentGroups

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetPreventAddingMembersToNonexistentGroups(v bool)`

SetPreventAddingMembersToNonexistentGroups sets PreventAddingMembersToNonexistentGroups field to given value.

### HasPreventAddingMembersToNonexistentGroups

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) HasPreventAddingMembersToNonexistentGroups() bool`

HasPreventAddingMembersToNonexistentGroups returns a boolean if a field has been set.

### GetPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetPreventAddingGroupsAsInvertedStaticGroupMembers() bool`

GetPreventAddingGroupsAsInvertedStaticGroupMembers returns the PreventAddingGroupsAsInvertedStaticGroupMembers field if non-nil, zero value otherwise.

### GetPreventAddingGroupsAsInvertedStaticGroupMembersOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetPreventAddingGroupsAsInvertedStaticGroupMembersOk() (*bool, bool)`

GetPreventAddingGroupsAsInvertedStaticGroupMembersOk returns a tuple with the PreventAddingGroupsAsInvertedStaticGroupMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetPreventAddingGroupsAsInvertedStaticGroupMembers(v bool)`

SetPreventAddingGroupsAsInvertedStaticGroupMembers sets PreventAddingGroupsAsInvertedStaticGroupMembers field to given value.

### HasPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) HasPreventAddingGroupsAsInvertedStaticGroupMembers() bool`

HasPreventAddingGroupsAsInvertedStaticGroupMembers returns a boolean if a field has been set.

### GetPreventNestingNonexistentGroups

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetPreventNestingNonexistentGroups() bool`

GetPreventNestingNonexistentGroups returns the PreventNestingNonexistentGroups field if non-nil, zero value otherwise.

### GetPreventNestingNonexistentGroupsOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetPreventNestingNonexistentGroupsOk() (*bool, bool)`

GetPreventNestingNonexistentGroupsOk returns a tuple with the PreventNestingNonexistentGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventNestingNonexistentGroups

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetPreventNestingNonexistentGroups(v bool)`

SetPreventNestingNonexistentGroups sets PreventNestingNonexistentGroups field to given value.

### HasPreventNestingNonexistentGroups

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) HasPreventNestingNonexistentGroups() bool`

HasPreventNestingNonexistentGroups returns a boolean if a field has been set.

### GetDescription

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvertedStaticGroupReferentialIntegrityPluginResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


