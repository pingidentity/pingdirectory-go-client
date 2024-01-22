# AddDelegatedAdminResourceRightsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumdelegatedAdminResourceRightsSchemaUrn**](EnumdelegatedAdminResourceRightsSchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Delegated Admin Resource Rights | [optional] 
**Enabled** | **bool** | Indicates whether these Delegated Admin Resource Rights are enabled. | 
**RestResourceType** | **string** | Name of the new Delegated Admin Resource Rights | 
**AdminPermission** | Pointer to [**[]EnumdelegatedAdminResourceRightsAdminPermissionProp**](EnumdelegatedAdminResourceRightsAdminPermissionProp.md) |  | [optional] 
**AdminScope** | Pointer to [**EnumdelegatedAdminResourceRightsAdminScopeProp**](EnumdelegatedAdminResourceRightsAdminScopeProp.md) |  | [optional] 
**ResourceSubtree** | Pointer to **[]string** | Specifies subtrees within the search base whose entries can be managed by the administrator(s). The admin-scope must be set to resources-in-specific-subtrees. | [optional] 
**ResourcesInGroup** | Pointer to **[]string** | Specifies groups whose members can be managed by the administrator(s). The admin-scope must be set to resources-in-specific-groups. | [optional] 

## Methods

### NewAddDelegatedAdminResourceRightsRequest

`func NewAddDelegatedAdminResourceRightsRequest(enabled bool, restResourceType string, ) *AddDelegatedAdminResourceRightsRequest`

NewAddDelegatedAdminResourceRightsRequest instantiates a new AddDelegatedAdminResourceRightsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDelegatedAdminResourceRightsRequestWithDefaults

`func NewAddDelegatedAdminResourceRightsRequestWithDefaults() *AddDelegatedAdminResourceRightsRequest`

NewAddDelegatedAdminResourceRightsRequestWithDefaults instantiates a new AddDelegatedAdminResourceRightsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddDelegatedAdminResourceRightsRequest) GetSchemas() []EnumdelegatedAdminResourceRightsSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddDelegatedAdminResourceRightsRequest) GetSchemasOk() (*[]EnumdelegatedAdminResourceRightsSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddDelegatedAdminResourceRightsRequest) SetSchemas(v []EnumdelegatedAdminResourceRightsSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddDelegatedAdminResourceRightsRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddDelegatedAdminResourceRightsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDelegatedAdminResourceRightsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDelegatedAdminResourceRightsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDelegatedAdminResourceRightsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddDelegatedAdminResourceRightsRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddDelegatedAdminResourceRightsRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddDelegatedAdminResourceRightsRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRestResourceType

`func (o *AddDelegatedAdminResourceRightsRequest) GetRestResourceType() string`

GetRestResourceType returns the RestResourceType field if non-nil, zero value otherwise.

### GetRestResourceTypeOk

`func (o *AddDelegatedAdminResourceRightsRequest) GetRestResourceTypeOk() (*string, bool)`

GetRestResourceTypeOk returns a tuple with the RestResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestResourceType

`func (o *AddDelegatedAdminResourceRightsRequest) SetRestResourceType(v string)`

SetRestResourceType sets RestResourceType field to given value.


### GetAdminPermission

`func (o *AddDelegatedAdminResourceRightsRequest) GetAdminPermission() []EnumdelegatedAdminResourceRightsAdminPermissionProp`

GetAdminPermission returns the AdminPermission field if non-nil, zero value otherwise.

### GetAdminPermissionOk

`func (o *AddDelegatedAdminResourceRightsRequest) GetAdminPermissionOk() (*[]EnumdelegatedAdminResourceRightsAdminPermissionProp, bool)`

GetAdminPermissionOk returns a tuple with the AdminPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPermission

`func (o *AddDelegatedAdminResourceRightsRequest) SetAdminPermission(v []EnumdelegatedAdminResourceRightsAdminPermissionProp)`

SetAdminPermission sets AdminPermission field to given value.

### HasAdminPermission

`func (o *AddDelegatedAdminResourceRightsRequest) HasAdminPermission() bool`

HasAdminPermission returns a boolean if a field has been set.

### GetAdminScope

`func (o *AddDelegatedAdminResourceRightsRequest) GetAdminScope() EnumdelegatedAdminResourceRightsAdminScopeProp`

GetAdminScope returns the AdminScope field if non-nil, zero value otherwise.

### GetAdminScopeOk

`func (o *AddDelegatedAdminResourceRightsRequest) GetAdminScopeOk() (*EnumdelegatedAdminResourceRightsAdminScopeProp, bool)`

GetAdminScopeOk returns a tuple with the AdminScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminScope

`func (o *AddDelegatedAdminResourceRightsRequest) SetAdminScope(v EnumdelegatedAdminResourceRightsAdminScopeProp)`

SetAdminScope sets AdminScope field to given value.

### HasAdminScope

`func (o *AddDelegatedAdminResourceRightsRequest) HasAdminScope() bool`

HasAdminScope returns a boolean if a field has been set.

### GetResourceSubtree

`func (o *AddDelegatedAdminResourceRightsRequest) GetResourceSubtree() []string`

GetResourceSubtree returns the ResourceSubtree field if non-nil, zero value otherwise.

### GetResourceSubtreeOk

`func (o *AddDelegatedAdminResourceRightsRequest) GetResourceSubtreeOk() (*[]string, bool)`

GetResourceSubtreeOk returns a tuple with the ResourceSubtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceSubtree

`func (o *AddDelegatedAdminResourceRightsRequest) SetResourceSubtree(v []string)`

SetResourceSubtree sets ResourceSubtree field to given value.

### HasResourceSubtree

`func (o *AddDelegatedAdminResourceRightsRequest) HasResourceSubtree() bool`

HasResourceSubtree returns a boolean if a field has been set.

### GetResourcesInGroup

`func (o *AddDelegatedAdminResourceRightsRequest) GetResourcesInGroup() []string`

GetResourcesInGroup returns the ResourcesInGroup field if non-nil, zero value otherwise.

### GetResourcesInGroupOk

`func (o *AddDelegatedAdminResourceRightsRequest) GetResourcesInGroupOk() (*[]string, bool)`

GetResourcesInGroupOk returns a tuple with the ResourcesInGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesInGroup

`func (o *AddDelegatedAdminResourceRightsRequest) SetResourcesInGroup(v []string)`

SetResourcesInGroup sets ResourcesInGroup field to given value.

### HasResourcesInGroup

`func (o *AddDelegatedAdminResourceRightsRequest) HasResourcesInGroup() bool`

HasResourcesInGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


