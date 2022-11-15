# AddReverseDnJoinVirtualAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new Virtual Attribute | 
**Schemas** | [**[]EnumreverseDnJoinVirtualAttributeSchemaUrn**](EnumreverseDnJoinVirtualAttributeSchemaUrn.md) |  | 
**JoinDNAttribute** | **string** | The attribute in related entries whose set of values must contain the DN of the search result entry to be joined with that entry. | 
**JoinBaseDNType** | [**EnumvirtualAttributeJoinBaseDNTypeProp**](EnumvirtualAttributeJoinBaseDNTypeProp.md) |  | 
**JoinCustomBaseDN** | Pointer to **string** | The fixed, administrator-specified base DN for the internal searches used to identify joined entries. | [optional] 
**JoinScope** | Pointer to [**EnumvirtualAttributeJoinScopeProp**](EnumvirtualAttributeJoinScopeProp.md) |  | [optional] 
**JoinSizeLimit** | Pointer to **int32** | The maximum number of entries that may be joined with the source entry, which also corresponds to the maximum number of values that the virtual attribute provider will generate for an entry. | [optional] 
**JoinFilter** | Pointer to **string** | An optional filter that specifies additional criteria for identifying joined entries. If a join-filter value is specified, then only entries matching that filter (in addition to satisfying the other join criteria) will be joined with the search result entry. | [optional] 
**JoinAttribute** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Virtual Attribute | [optional] 
**Enabled** | **bool** | Indicates whether the Virtual Attribute is enabled for use. | 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute. | 
**BaseDN** | Pointer to **[]string** |  | [optional] 
**GroupDN** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to **[]string** |  | [optional] 
**ClientConnectionPolicy** | Pointer to **[]string** |  | [optional] 
**ConflictBehavior** | Pointer to [**EnumvirtualAttributeConflictBehaviorProp**](EnumvirtualAttributeConflictBehaviorProp.md) |  | [optional] 
**RequireExplicitRequestByName** | Pointer to **bool** | Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type. | [optional] 
**MultipleVirtualAttributeEvaluationOrderIndex** | Pointer to **int32** | Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry. | [optional] 
**MultipleVirtualAttributeMergeBehavior** | Pointer to [**EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp**](EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp.md) |  | [optional] 
**AllowIndexConflicts** | Pointer to **bool** | Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server. | [optional] 

## Methods

### NewAddReverseDnJoinVirtualAttributeRequest

`func NewAddReverseDnJoinVirtualAttributeRequest(name string, schemas []EnumreverseDnJoinVirtualAttributeSchemaUrn, joinDNAttribute string, joinBaseDNType EnumvirtualAttributeJoinBaseDNTypeProp, enabled bool, attributeType string, ) *AddReverseDnJoinVirtualAttributeRequest`

NewAddReverseDnJoinVirtualAttributeRequest instantiates a new AddReverseDnJoinVirtualAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddReverseDnJoinVirtualAttributeRequestWithDefaults

`func NewAddReverseDnJoinVirtualAttributeRequestWithDefaults() *AddReverseDnJoinVirtualAttributeRequest`

NewAddReverseDnJoinVirtualAttributeRequestWithDefaults instantiates a new AddReverseDnJoinVirtualAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSchemas

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetSchemas() []EnumreverseDnJoinVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetSchemasOk() (*[]EnumreverseDnJoinVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetSchemas(v []EnumreverseDnJoinVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetJoinDNAttribute

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinDNAttribute() string`

GetJoinDNAttribute returns the JoinDNAttribute field if non-nil, zero value otherwise.

### GetJoinDNAttributeOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinDNAttributeOk() (*string, bool)`

GetJoinDNAttributeOk returns a tuple with the JoinDNAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDNAttribute

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetJoinDNAttribute(v string)`

SetJoinDNAttribute sets JoinDNAttribute field to given value.


### GetJoinBaseDNType

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinBaseDNType() EnumvirtualAttributeJoinBaseDNTypeProp`

GetJoinBaseDNType returns the JoinBaseDNType field if non-nil, zero value otherwise.

### GetJoinBaseDNTypeOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinBaseDNTypeOk() (*EnumvirtualAttributeJoinBaseDNTypeProp, bool)`

GetJoinBaseDNTypeOk returns a tuple with the JoinBaseDNType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinBaseDNType

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetJoinBaseDNType(v EnumvirtualAttributeJoinBaseDNTypeProp)`

SetJoinBaseDNType sets JoinBaseDNType field to given value.


### GetJoinCustomBaseDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinCustomBaseDN() string`

GetJoinCustomBaseDN returns the JoinCustomBaseDN field if non-nil, zero value otherwise.

### GetJoinCustomBaseDNOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinCustomBaseDNOk() (*string, bool)`

GetJoinCustomBaseDNOk returns a tuple with the JoinCustomBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinCustomBaseDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetJoinCustomBaseDN(v string)`

SetJoinCustomBaseDN sets JoinCustomBaseDN field to given value.

### HasJoinCustomBaseDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasJoinCustomBaseDN() bool`

HasJoinCustomBaseDN returns a boolean if a field has been set.

### GetJoinScope

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinScope() EnumvirtualAttributeJoinScopeProp`

GetJoinScope returns the JoinScope field if non-nil, zero value otherwise.

### GetJoinScopeOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinScopeOk() (*EnumvirtualAttributeJoinScopeProp, bool)`

GetJoinScopeOk returns a tuple with the JoinScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinScope

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetJoinScope(v EnumvirtualAttributeJoinScopeProp)`

SetJoinScope sets JoinScope field to given value.

### HasJoinScope

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasJoinScope() bool`

HasJoinScope returns a boolean if a field has been set.

### GetJoinSizeLimit

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinSizeLimit() int32`

GetJoinSizeLimit returns the JoinSizeLimit field if non-nil, zero value otherwise.

### GetJoinSizeLimitOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinSizeLimitOk() (*int32, bool)`

GetJoinSizeLimitOk returns a tuple with the JoinSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSizeLimit

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetJoinSizeLimit(v int32)`

SetJoinSizeLimit sets JoinSizeLimit field to given value.

### HasJoinSizeLimit

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasJoinSizeLimit() bool`

HasJoinSizeLimit returns a boolean if a field has been set.

### GetJoinFilter

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinFilter() string`

GetJoinFilter returns the JoinFilter field if non-nil, zero value otherwise.

### GetJoinFilterOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinFilterOk() (*string, bool)`

GetJoinFilterOk returns a tuple with the JoinFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinFilter

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetJoinFilter(v string)`

SetJoinFilter sets JoinFilter field to given value.

### HasJoinFilter

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasJoinFilter() bool`

HasJoinFilter returns a boolean if a field has been set.

### GetJoinAttribute

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinAttribute() []string`

GetJoinAttribute returns the JoinAttribute field if non-nil, zero value otherwise.

### GetJoinAttributeOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetJoinAttributeOk() (*[]string, bool)`

GetJoinAttributeOk returns a tuple with the JoinAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinAttribute

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetJoinAttribute(v []string)`

SetJoinAttribute sets JoinAttribute field to given value.

### HasJoinAttribute

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasJoinAttribute() bool`

HasJoinAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetConflictBehavior

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *AddReverseDnJoinVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *AddReverseDnJoinVirtualAttributeRequest) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *AddReverseDnJoinVirtualAttributeRequest) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


