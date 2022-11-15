# IsMemberOfVirtualAttributeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumisMemberOfVirtualAttributeSchemaUrn**](EnumisMemberOfVirtualAttributeSchemaUrn.md) |  | 
**ConflictBehavior** | Pointer to [**EnumvirtualAttributeConflictBehaviorProp**](EnumvirtualAttributeConflictBehaviorProp.md) |  | [optional] 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute. | 
**DirectMembershipsOnly** | Pointer to **bool** | Specifies whether to only include groups in which the user is directly associated with and the membership maybe modified via the group entry. Groups in which the user&#39;s membership is derived dynamically or through nested groups will not be included. | [optional] 
**IncludedGroupFilter** | Pointer to **string** | A search filter that will be used to identify which groups should be included in the values of the virtual attribute. With no value defined (which is the default behavior), all groups that contain the target user will be included. | [optional] 
**RewriteSearchFilters** | Pointer to [**EnumvirtualAttributeRewriteSearchFiltersProp**](EnumvirtualAttributeRewriteSearchFiltersProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Virtual Attribute | [optional] 
**Enabled** | **bool** | Indicates whether the Virtual Attribute is enabled for use. | 
**BaseDN** | Pointer to **[]string** |  | [optional] 
**GroupDN** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to **[]string** |  | [optional] 
**ClientConnectionPolicy** | Pointer to **[]string** |  | [optional] 
**RequireExplicitRequestByName** | Pointer to **bool** | Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type. | [optional] 
**MultipleVirtualAttributeEvaluationOrderIndex** | Pointer to **int32** | Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry. | [optional] 
**MultipleVirtualAttributeMergeBehavior** | Pointer to [**EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp**](EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp.md) |  | [optional] 
**AllowIndexConflicts** | Pointer to **bool** | Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server. | [optional] 

## Methods

### NewIsMemberOfVirtualAttributeShared

`func NewIsMemberOfVirtualAttributeShared(schemas []EnumisMemberOfVirtualAttributeSchemaUrn, attributeType string, enabled bool, ) *IsMemberOfVirtualAttributeShared`

NewIsMemberOfVirtualAttributeShared instantiates a new IsMemberOfVirtualAttributeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsMemberOfVirtualAttributeSharedWithDefaults

`func NewIsMemberOfVirtualAttributeSharedWithDefaults() *IsMemberOfVirtualAttributeShared`

NewIsMemberOfVirtualAttributeSharedWithDefaults instantiates a new IsMemberOfVirtualAttributeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *IsMemberOfVirtualAttributeShared) GetSchemas() []EnumisMemberOfVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *IsMemberOfVirtualAttributeShared) GetSchemasOk() (*[]EnumisMemberOfVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *IsMemberOfVirtualAttributeShared) SetSchemas(v []EnumisMemberOfVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConflictBehavior

`func (o *IsMemberOfVirtualAttributeShared) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *IsMemberOfVirtualAttributeShared) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *IsMemberOfVirtualAttributeShared) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *IsMemberOfVirtualAttributeShared) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetAttributeType

`func (o *IsMemberOfVirtualAttributeShared) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *IsMemberOfVirtualAttributeShared) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *IsMemberOfVirtualAttributeShared) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetDirectMembershipsOnly

`func (o *IsMemberOfVirtualAttributeShared) GetDirectMembershipsOnly() bool`

GetDirectMembershipsOnly returns the DirectMembershipsOnly field if non-nil, zero value otherwise.

### GetDirectMembershipsOnlyOk

`func (o *IsMemberOfVirtualAttributeShared) GetDirectMembershipsOnlyOk() (*bool, bool)`

GetDirectMembershipsOnlyOk returns a tuple with the DirectMembershipsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMembershipsOnly

`func (o *IsMemberOfVirtualAttributeShared) SetDirectMembershipsOnly(v bool)`

SetDirectMembershipsOnly sets DirectMembershipsOnly field to given value.

### HasDirectMembershipsOnly

`func (o *IsMemberOfVirtualAttributeShared) HasDirectMembershipsOnly() bool`

HasDirectMembershipsOnly returns a boolean if a field has been set.

### GetIncludedGroupFilter

`func (o *IsMemberOfVirtualAttributeShared) GetIncludedGroupFilter() string`

GetIncludedGroupFilter returns the IncludedGroupFilter field if non-nil, zero value otherwise.

### GetIncludedGroupFilterOk

`func (o *IsMemberOfVirtualAttributeShared) GetIncludedGroupFilterOk() (*string, bool)`

GetIncludedGroupFilterOk returns a tuple with the IncludedGroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedGroupFilter

`func (o *IsMemberOfVirtualAttributeShared) SetIncludedGroupFilter(v string)`

SetIncludedGroupFilter sets IncludedGroupFilter field to given value.

### HasIncludedGroupFilter

`func (o *IsMemberOfVirtualAttributeShared) HasIncludedGroupFilter() bool`

HasIncludedGroupFilter returns a boolean if a field has been set.

### GetRewriteSearchFilters

`func (o *IsMemberOfVirtualAttributeShared) GetRewriteSearchFilters() EnumvirtualAttributeRewriteSearchFiltersProp`

GetRewriteSearchFilters returns the RewriteSearchFilters field if non-nil, zero value otherwise.

### GetRewriteSearchFiltersOk

`func (o *IsMemberOfVirtualAttributeShared) GetRewriteSearchFiltersOk() (*EnumvirtualAttributeRewriteSearchFiltersProp, bool)`

GetRewriteSearchFiltersOk returns a tuple with the RewriteSearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteSearchFilters

`func (o *IsMemberOfVirtualAttributeShared) SetRewriteSearchFilters(v EnumvirtualAttributeRewriteSearchFiltersProp)`

SetRewriteSearchFilters sets RewriteSearchFilters field to given value.

### HasRewriteSearchFilters

`func (o *IsMemberOfVirtualAttributeShared) HasRewriteSearchFilters() bool`

HasRewriteSearchFilters returns a boolean if a field has been set.

### GetDescription

`func (o *IsMemberOfVirtualAttributeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IsMemberOfVirtualAttributeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IsMemberOfVirtualAttributeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IsMemberOfVirtualAttributeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *IsMemberOfVirtualAttributeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IsMemberOfVirtualAttributeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IsMemberOfVirtualAttributeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseDN

`func (o *IsMemberOfVirtualAttributeShared) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *IsMemberOfVirtualAttributeShared) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *IsMemberOfVirtualAttributeShared) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *IsMemberOfVirtualAttributeShared) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *IsMemberOfVirtualAttributeShared) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *IsMemberOfVirtualAttributeShared) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *IsMemberOfVirtualAttributeShared) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *IsMemberOfVirtualAttributeShared) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *IsMemberOfVirtualAttributeShared) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IsMemberOfVirtualAttributeShared) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IsMemberOfVirtualAttributeShared) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IsMemberOfVirtualAttributeShared) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *IsMemberOfVirtualAttributeShared) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *IsMemberOfVirtualAttributeShared) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *IsMemberOfVirtualAttributeShared) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *IsMemberOfVirtualAttributeShared) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *IsMemberOfVirtualAttributeShared) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *IsMemberOfVirtualAttributeShared) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *IsMemberOfVirtualAttributeShared) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *IsMemberOfVirtualAttributeShared) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *IsMemberOfVirtualAttributeShared) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *IsMemberOfVirtualAttributeShared) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *IsMemberOfVirtualAttributeShared) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *IsMemberOfVirtualAttributeShared) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *IsMemberOfVirtualAttributeShared) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *IsMemberOfVirtualAttributeShared) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *IsMemberOfVirtualAttributeShared) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *IsMemberOfVirtualAttributeShared) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *IsMemberOfVirtualAttributeShared) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *IsMemberOfVirtualAttributeShared) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *IsMemberOfVirtualAttributeShared) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *IsMemberOfVirtualAttributeShared) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


