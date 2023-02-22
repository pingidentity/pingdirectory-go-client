# AddEqualityJoinVirtualAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new Virtual Attribute | 
**Schemas** | [**[]EnumequalityJoinVirtualAttributeSchemaUrn**](EnumequalityJoinVirtualAttributeSchemaUrn.md) |  | 
**JoinSourceAttribute** | **string** | The attribute containing the value(s) in the source entry to use to identify related entries. | 
**JoinTargetAttribute** | **string** | The attribute in target entries whose value(s) match values of the source attribute in the source entry. | 
**JoinMatchAll** | Pointer to **bool** | Indicates whether joined entries will be required to have all values for the source attribute, or only at least one of its values. | [optional] 
**JoinBaseDNType** | [**EnumvirtualAttributeJoinBaseDNTypeProp**](EnumvirtualAttributeJoinBaseDNTypeProp.md) |  | 
**JoinCustomBaseDN** | Pointer to **string** | The fixed, administrator-specified base DN for the internal searches used to identify joined entries. | [optional] 
**JoinScope** | Pointer to [**EnumvirtualAttributeJoinScopeProp**](EnumvirtualAttributeJoinScopeProp.md) |  | [optional] 
**JoinSizeLimit** | Pointer to **int32** | The maximum number of entries that may be joined with the source entry, which also corresponds to the maximum number of values that the virtual attribute provider will generate for an entry. | [optional] 
**JoinFilter** | Pointer to **string** | An optional filter that specifies additional criteria for identifying joined entries. If a join-filter value is specified, then only entries matching that filter (in addition to satisfying the other join criteria) will be joined with the search result entry. | [optional] 
**JoinAttribute** | Pointer to **[]string** | An optional set of the names of the attributes to include with joined entries. | [optional] 
**Description** | Pointer to **string** | A description for this Virtual Attribute | [optional] 
**Enabled** | **bool** | Indicates whether the Virtual Attribute is enabled for use. | 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute. | 
**BaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches containing entries that are eligible to use this virtual attribute. | [optional] 
**GroupDN** | Pointer to **[]string** | Specifies the DNs of the groups whose members can be eligible to use this virtual attribute. | [optional] 
**Filter** | Pointer to **[]string** | Specifies the search filters to be applied against entries to determine if the virtual attribute is to be generated for those entries. | [optional] 
**ClientConnectionPolicy** | Pointer to **[]string** | Specifies a set of client connection policies for which this Virtual Attribute should be generated. If this is undefined, then this Virtual Attribute will always be generated. If it is associated with one or more client connection policies, then this Virtual Attribute will be generated only for operations requested by clients assigned to one of those client connection policies. | [optional] 
**ConflictBehavior** | Pointer to [**EnumvirtualAttributeConflictBehaviorProp**](EnumvirtualAttributeConflictBehaviorProp.md) |  | [optional] 
**RequireExplicitRequestByName** | Pointer to **bool** | Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type. | [optional] 
**MultipleVirtualAttributeEvaluationOrderIndex** | Pointer to **int32** | Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry. | [optional] 
**MultipleVirtualAttributeMergeBehavior** | Pointer to [**EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp**](EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp.md) |  | [optional] 
**AllowIndexConflicts** | Pointer to **bool** | Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server. | [optional] 

## Methods

### NewAddEqualityJoinVirtualAttributeRequest

`func NewAddEqualityJoinVirtualAttributeRequest(name string, schemas []EnumequalityJoinVirtualAttributeSchemaUrn, joinSourceAttribute string, joinTargetAttribute string, joinBaseDNType EnumvirtualAttributeJoinBaseDNTypeProp, enabled bool, attributeType string, ) *AddEqualityJoinVirtualAttributeRequest`

NewAddEqualityJoinVirtualAttributeRequest instantiates a new AddEqualityJoinVirtualAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEqualityJoinVirtualAttributeRequestWithDefaults

`func NewAddEqualityJoinVirtualAttributeRequestWithDefaults() *AddEqualityJoinVirtualAttributeRequest`

NewAddEqualityJoinVirtualAttributeRequestWithDefaults instantiates a new AddEqualityJoinVirtualAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddEqualityJoinVirtualAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddEqualityJoinVirtualAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSchemas

`func (o *AddEqualityJoinVirtualAttributeRequest) GetSchemas() []EnumequalityJoinVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetSchemasOk() (*[]EnumequalityJoinVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddEqualityJoinVirtualAttributeRequest) SetSchemas(v []EnumequalityJoinVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetJoinSourceAttribute

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinSourceAttribute() string`

GetJoinSourceAttribute returns the JoinSourceAttribute field if non-nil, zero value otherwise.

### GetJoinSourceAttributeOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinSourceAttributeOk() (*string, bool)`

GetJoinSourceAttributeOk returns a tuple with the JoinSourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSourceAttribute

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinSourceAttribute(v string)`

SetJoinSourceAttribute sets JoinSourceAttribute field to given value.


### GetJoinTargetAttribute

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinTargetAttribute() string`

GetJoinTargetAttribute returns the JoinTargetAttribute field if non-nil, zero value otherwise.

### GetJoinTargetAttributeOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinTargetAttributeOk() (*string, bool)`

GetJoinTargetAttributeOk returns a tuple with the JoinTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTargetAttribute

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinTargetAttribute(v string)`

SetJoinTargetAttribute sets JoinTargetAttribute field to given value.


### GetJoinMatchAll

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinMatchAll() bool`

GetJoinMatchAll returns the JoinMatchAll field if non-nil, zero value otherwise.

### GetJoinMatchAllOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinMatchAllOk() (*bool, bool)`

GetJoinMatchAllOk returns a tuple with the JoinMatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinMatchAll

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinMatchAll(v bool)`

SetJoinMatchAll sets JoinMatchAll field to given value.

### HasJoinMatchAll

`func (o *AddEqualityJoinVirtualAttributeRequest) HasJoinMatchAll() bool`

HasJoinMatchAll returns a boolean if a field has been set.

### GetJoinBaseDNType

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinBaseDNType() EnumvirtualAttributeJoinBaseDNTypeProp`

GetJoinBaseDNType returns the JoinBaseDNType field if non-nil, zero value otherwise.

### GetJoinBaseDNTypeOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinBaseDNTypeOk() (*EnumvirtualAttributeJoinBaseDNTypeProp, bool)`

GetJoinBaseDNTypeOk returns a tuple with the JoinBaseDNType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinBaseDNType

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinBaseDNType(v EnumvirtualAttributeJoinBaseDNTypeProp)`

SetJoinBaseDNType sets JoinBaseDNType field to given value.


### GetJoinCustomBaseDN

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinCustomBaseDN() string`

GetJoinCustomBaseDN returns the JoinCustomBaseDN field if non-nil, zero value otherwise.

### GetJoinCustomBaseDNOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinCustomBaseDNOk() (*string, bool)`

GetJoinCustomBaseDNOk returns a tuple with the JoinCustomBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinCustomBaseDN

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinCustomBaseDN(v string)`

SetJoinCustomBaseDN sets JoinCustomBaseDN field to given value.

### HasJoinCustomBaseDN

`func (o *AddEqualityJoinVirtualAttributeRequest) HasJoinCustomBaseDN() bool`

HasJoinCustomBaseDN returns a boolean if a field has been set.

### GetJoinScope

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinScope() EnumvirtualAttributeJoinScopeProp`

GetJoinScope returns the JoinScope field if non-nil, zero value otherwise.

### GetJoinScopeOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinScopeOk() (*EnumvirtualAttributeJoinScopeProp, bool)`

GetJoinScopeOk returns a tuple with the JoinScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinScope

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinScope(v EnumvirtualAttributeJoinScopeProp)`

SetJoinScope sets JoinScope field to given value.

### HasJoinScope

`func (o *AddEqualityJoinVirtualAttributeRequest) HasJoinScope() bool`

HasJoinScope returns a boolean if a field has been set.

### GetJoinSizeLimit

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinSizeLimit() int32`

GetJoinSizeLimit returns the JoinSizeLimit field if non-nil, zero value otherwise.

### GetJoinSizeLimitOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinSizeLimitOk() (*int32, bool)`

GetJoinSizeLimitOk returns a tuple with the JoinSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSizeLimit

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinSizeLimit(v int32)`

SetJoinSizeLimit sets JoinSizeLimit field to given value.

### HasJoinSizeLimit

`func (o *AddEqualityJoinVirtualAttributeRequest) HasJoinSizeLimit() bool`

HasJoinSizeLimit returns a boolean if a field has been set.

### GetJoinFilter

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinFilter() string`

GetJoinFilter returns the JoinFilter field if non-nil, zero value otherwise.

### GetJoinFilterOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinFilterOk() (*string, bool)`

GetJoinFilterOk returns a tuple with the JoinFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinFilter

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinFilter(v string)`

SetJoinFilter sets JoinFilter field to given value.

### HasJoinFilter

`func (o *AddEqualityJoinVirtualAttributeRequest) HasJoinFilter() bool`

HasJoinFilter returns a boolean if a field has been set.

### GetJoinAttribute

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinAttribute() []string`

GetJoinAttribute returns the JoinAttribute field if non-nil, zero value otherwise.

### GetJoinAttributeOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetJoinAttributeOk() (*[]string, bool)`

GetJoinAttributeOk returns a tuple with the JoinAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinAttribute

`func (o *AddEqualityJoinVirtualAttributeRequest) SetJoinAttribute(v []string)`

SetJoinAttribute sets JoinAttribute field to given value.

### HasJoinAttribute

`func (o *AddEqualityJoinVirtualAttributeRequest) HasJoinAttribute() bool`

HasJoinAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *AddEqualityJoinVirtualAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddEqualityJoinVirtualAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddEqualityJoinVirtualAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddEqualityJoinVirtualAttributeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddEqualityJoinVirtualAttributeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *AddEqualityJoinVirtualAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddEqualityJoinVirtualAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *AddEqualityJoinVirtualAttributeRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddEqualityJoinVirtualAttributeRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddEqualityJoinVirtualAttributeRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *AddEqualityJoinVirtualAttributeRequest) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *AddEqualityJoinVirtualAttributeRequest) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *AddEqualityJoinVirtualAttributeRequest) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *AddEqualityJoinVirtualAttributeRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddEqualityJoinVirtualAttributeRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddEqualityJoinVirtualAttributeRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *AddEqualityJoinVirtualAttributeRequest) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *AddEqualityJoinVirtualAttributeRequest) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *AddEqualityJoinVirtualAttributeRequest) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetConflictBehavior

`func (o *AddEqualityJoinVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *AddEqualityJoinVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *AddEqualityJoinVirtualAttributeRequest) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *AddEqualityJoinVirtualAttributeRequest) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *AddEqualityJoinVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *AddEqualityJoinVirtualAttributeRequest) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddEqualityJoinVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddEqualityJoinVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddEqualityJoinVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *AddEqualityJoinVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *AddEqualityJoinVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *AddEqualityJoinVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *AddEqualityJoinVirtualAttributeRequest) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *AddEqualityJoinVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *AddEqualityJoinVirtualAttributeRequest) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *AddEqualityJoinVirtualAttributeRequest) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


