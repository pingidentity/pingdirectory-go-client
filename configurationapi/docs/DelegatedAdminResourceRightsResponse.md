# DelegatedAdminResourceRightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Delegated Admin Resource Rights | 
**Schemas** | Pointer to [**[]EnumdelegatedAdminResourceRightsSchemaUrn**](EnumdelegatedAdminResourceRightsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Resource Rights | [optional] 
**Enabled** | **bool** | Indicates whether these Delegated Admin Resource Rights are enabled. | 
**RestResourceType** | **string** | Specifies the resource type applicable to these Delegated Admin Resource Rights. | 
**AdminPermission** | Pointer to [**[]EnumdelegatedAdminResourceRightsAdminPermissionProp**](EnumdelegatedAdminResourceRightsAdminPermissionProp.md) |  | [optional] 
**AdminScope** | Pointer to [**EnumdelegatedAdminResourceRightsAdminScopeProp**](EnumdelegatedAdminResourceRightsAdminScopeProp.md) |  | [optional] 
**ResourceSubtree** | Pointer to **[]string** | Specifies subtrees within the search base whose entries can be managed by the administrator(s). The admin-scope must be set to resources-in-specific-subtrees. | [optional] 
**ResourcesInGroup** | Pointer to **[]string** | Specifies groups whose members can be managed by the administrator(s). The admin-scope must be set to resources-in-specific-groups. | [optional] 

## Methods

### NewDelegatedAdminResourceRightsResponse

`func NewDelegatedAdminResourceRightsResponse(id string, enabled bool, restResourceType string, ) *DelegatedAdminResourceRightsResponse`

NewDelegatedAdminResourceRightsResponse instantiates a new DelegatedAdminResourceRightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegatedAdminResourceRightsResponseWithDefaults

`func NewDelegatedAdminResourceRightsResponseWithDefaults() *DelegatedAdminResourceRightsResponse`

NewDelegatedAdminResourceRightsResponseWithDefaults instantiates a new DelegatedAdminResourceRightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *DelegatedAdminResourceRightsResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DelegatedAdminResourceRightsResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DelegatedAdminResourceRightsResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DelegatedAdminResourceRightsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminResourceRightsResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *DelegatedAdminResourceRightsResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminResourceRightsResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *DelegatedAdminResourceRightsResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *DelegatedAdminResourceRightsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelegatedAdminResourceRightsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelegatedAdminResourceRightsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DelegatedAdminResourceRightsResponse) GetSchemas() []EnumdelegatedAdminResourceRightsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelegatedAdminResourceRightsResponse) GetSchemasOk() (*[]EnumdelegatedAdminResourceRightsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelegatedAdminResourceRightsResponse) SetSchemas(v []EnumdelegatedAdminResourceRightsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *DelegatedAdminResourceRightsResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *DelegatedAdminResourceRightsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelegatedAdminResourceRightsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelegatedAdminResourceRightsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelegatedAdminResourceRightsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DelegatedAdminResourceRightsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DelegatedAdminResourceRightsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DelegatedAdminResourceRightsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRestResourceType

`func (o *DelegatedAdminResourceRightsResponse) GetRestResourceType() string`

GetRestResourceType returns the RestResourceType field if non-nil, zero value otherwise.

### GetRestResourceTypeOk

`func (o *DelegatedAdminResourceRightsResponse) GetRestResourceTypeOk() (*string, bool)`

GetRestResourceTypeOk returns a tuple with the RestResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestResourceType

`func (o *DelegatedAdminResourceRightsResponse) SetRestResourceType(v string)`

SetRestResourceType sets RestResourceType field to given value.


### GetAdminPermission

`func (o *DelegatedAdminResourceRightsResponse) GetAdminPermission() []EnumdelegatedAdminResourceRightsAdminPermissionProp`

GetAdminPermission returns the AdminPermission field if non-nil, zero value otherwise.

### GetAdminPermissionOk

`func (o *DelegatedAdminResourceRightsResponse) GetAdminPermissionOk() (*[]EnumdelegatedAdminResourceRightsAdminPermissionProp, bool)`

GetAdminPermissionOk returns a tuple with the AdminPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPermission

`func (o *DelegatedAdminResourceRightsResponse) SetAdminPermission(v []EnumdelegatedAdminResourceRightsAdminPermissionProp)`

SetAdminPermission sets AdminPermission field to given value.

### HasAdminPermission

`func (o *DelegatedAdminResourceRightsResponse) HasAdminPermission() bool`

HasAdminPermission returns a boolean if a field has been set.

### GetAdminScope

`func (o *DelegatedAdminResourceRightsResponse) GetAdminScope() EnumdelegatedAdminResourceRightsAdminScopeProp`

GetAdminScope returns the AdminScope field if non-nil, zero value otherwise.

### GetAdminScopeOk

`func (o *DelegatedAdminResourceRightsResponse) GetAdminScopeOk() (*EnumdelegatedAdminResourceRightsAdminScopeProp, bool)`

GetAdminScopeOk returns a tuple with the AdminScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminScope

`func (o *DelegatedAdminResourceRightsResponse) SetAdminScope(v EnumdelegatedAdminResourceRightsAdminScopeProp)`

SetAdminScope sets AdminScope field to given value.

### HasAdminScope

`func (o *DelegatedAdminResourceRightsResponse) HasAdminScope() bool`

HasAdminScope returns a boolean if a field has been set.

### GetResourceSubtree

`func (o *DelegatedAdminResourceRightsResponse) GetResourceSubtree() []string`

GetResourceSubtree returns the ResourceSubtree field if non-nil, zero value otherwise.

### GetResourceSubtreeOk

`func (o *DelegatedAdminResourceRightsResponse) GetResourceSubtreeOk() (*[]string, bool)`

GetResourceSubtreeOk returns a tuple with the ResourceSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSubtree

`func (o *DelegatedAdminResourceRightsResponse) SetResourceSubtree(v []string)`

SetResourceSubtree sets ResourceSubtree field to given value.

### HasResourceSubtree

`func (o *DelegatedAdminResourceRightsResponse) HasResourceSubtree() bool`

HasResourceSubtree returns a boolean if a field has been set.

### GetResourcesInGroup

`func (o *DelegatedAdminResourceRightsResponse) GetResourcesInGroup() []string`

GetResourcesInGroup returns the ResourcesInGroup field if non-nil, zero value otherwise.

### GetResourcesInGroupOk

`func (o *DelegatedAdminResourceRightsResponse) GetResourcesInGroupOk() (*[]string, bool)`

GetResourcesInGroupOk returns a tuple with the ResourcesInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesInGroup

`func (o *DelegatedAdminResourceRightsResponse) SetResourcesInGroup(v []string)`

SetResourcesInGroup sets ResourcesInGroup field to given value.

### HasResourcesInGroup

`func (o *DelegatedAdminResourceRightsResponse) HasResourcesInGroup() bool`

HasResourcesInGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


