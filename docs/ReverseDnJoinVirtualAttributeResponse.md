# ReverseDnJoinVirtualAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Virtual Attribute | 
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

### NewReverseDnJoinVirtualAttributeResponse

`func NewReverseDnJoinVirtualAttributeResponse(id string, schemas []EnumreverseDnJoinVirtualAttributeSchemaUrn, joinDNAttribute string, joinBaseDNType EnumvirtualAttributeJoinBaseDNTypeProp, enabled bool, attributeType string, ) *ReverseDnJoinVirtualAttributeResponse`

NewReverseDnJoinVirtualAttributeResponse instantiates a new ReverseDnJoinVirtualAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseDnJoinVirtualAttributeResponseWithDefaults

`func NewReverseDnJoinVirtualAttributeResponseWithDefaults() *ReverseDnJoinVirtualAttributeResponse`

NewReverseDnJoinVirtualAttributeResponseWithDefaults instantiates a new ReverseDnJoinVirtualAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReverseDnJoinVirtualAttributeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReverseDnJoinVirtualAttributeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *ReverseDnJoinVirtualAttributeResponse) GetSchemas() []EnumreverseDnJoinVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetSchemasOk() (*[]EnumreverseDnJoinVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ReverseDnJoinVirtualAttributeResponse) SetSchemas(v []EnumreverseDnJoinVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetJoinDNAttribute

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinDNAttribute() string`

GetJoinDNAttribute returns the JoinDNAttribute field if non-nil, zero value otherwise.

### GetJoinDNAttributeOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinDNAttributeOk() (*string, bool)`

GetJoinDNAttributeOk returns a tuple with the JoinDNAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDNAttribute

`func (o *ReverseDnJoinVirtualAttributeResponse) SetJoinDNAttribute(v string)`

SetJoinDNAttribute sets JoinDNAttribute field to given value.


### GetJoinBaseDNType

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinBaseDNType() EnumvirtualAttributeJoinBaseDNTypeProp`

GetJoinBaseDNType returns the JoinBaseDNType field if non-nil, zero value otherwise.

### GetJoinBaseDNTypeOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinBaseDNTypeOk() (*EnumvirtualAttributeJoinBaseDNTypeProp, bool)`

GetJoinBaseDNTypeOk returns a tuple with the JoinBaseDNType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinBaseDNType

`func (o *ReverseDnJoinVirtualAttributeResponse) SetJoinBaseDNType(v EnumvirtualAttributeJoinBaseDNTypeProp)`

SetJoinBaseDNType sets JoinBaseDNType field to given value.


### GetJoinCustomBaseDN

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinCustomBaseDN() string`

GetJoinCustomBaseDN returns the JoinCustomBaseDN field if non-nil, zero value otherwise.

### GetJoinCustomBaseDNOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinCustomBaseDNOk() (*string, bool)`

GetJoinCustomBaseDNOk returns a tuple with the JoinCustomBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinCustomBaseDN

`func (o *ReverseDnJoinVirtualAttributeResponse) SetJoinCustomBaseDN(v string)`

SetJoinCustomBaseDN sets JoinCustomBaseDN field to given value.

### HasJoinCustomBaseDN

`func (o *ReverseDnJoinVirtualAttributeResponse) HasJoinCustomBaseDN() bool`

HasJoinCustomBaseDN returns a boolean if a field has been set.

### GetJoinScope

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinScope() EnumvirtualAttributeJoinScopeProp`

GetJoinScope returns the JoinScope field if non-nil, zero value otherwise.

### GetJoinScopeOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinScopeOk() (*EnumvirtualAttributeJoinScopeProp, bool)`

GetJoinScopeOk returns a tuple with the JoinScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinScope

`func (o *ReverseDnJoinVirtualAttributeResponse) SetJoinScope(v EnumvirtualAttributeJoinScopeProp)`

SetJoinScope sets JoinScope field to given value.

### HasJoinScope

`func (o *ReverseDnJoinVirtualAttributeResponse) HasJoinScope() bool`

HasJoinScope returns a boolean if a field has been set.

### GetJoinSizeLimit

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinSizeLimit() int32`

GetJoinSizeLimit returns the JoinSizeLimit field if non-nil, zero value otherwise.

### GetJoinSizeLimitOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinSizeLimitOk() (*int32, bool)`

GetJoinSizeLimitOk returns a tuple with the JoinSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSizeLimit

`func (o *ReverseDnJoinVirtualAttributeResponse) SetJoinSizeLimit(v int32)`

SetJoinSizeLimit sets JoinSizeLimit field to given value.

### HasJoinSizeLimit

`func (o *ReverseDnJoinVirtualAttributeResponse) HasJoinSizeLimit() bool`

HasJoinSizeLimit returns a boolean if a field has been set.

### GetJoinFilter

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinFilter() string`

GetJoinFilter returns the JoinFilter field if non-nil, zero value otherwise.

### GetJoinFilterOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinFilterOk() (*string, bool)`

GetJoinFilterOk returns a tuple with the JoinFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinFilter

`func (o *ReverseDnJoinVirtualAttributeResponse) SetJoinFilter(v string)`

SetJoinFilter sets JoinFilter field to given value.

### HasJoinFilter

`func (o *ReverseDnJoinVirtualAttributeResponse) HasJoinFilter() bool`

HasJoinFilter returns a boolean if a field has been set.

### GetJoinAttribute

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinAttribute() []string`

GetJoinAttribute returns the JoinAttribute field if non-nil, zero value otherwise.

### GetJoinAttributeOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetJoinAttributeOk() (*[]string, bool)`

GetJoinAttributeOk returns a tuple with the JoinAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinAttribute

`func (o *ReverseDnJoinVirtualAttributeResponse) SetJoinAttribute(v []string)`

SetJoinAttribute sets JoinAttribute field to given value.

### HasJoinAttribute

`func (o *ReverseDnJoinVirtualAttributeResponse) HasJoinAttribute() bool`

HasJoinAttribute returns a boolean if a field has been set.

### GetDescription

`func (o *ReverseDnJoinVirtualAttributeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReverseDnJoinVirtualAttributeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReverseDnJoinVirtualAttributeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ReverseDnJoinVirtualAttributeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ReverseDnJoinVirtualAttributeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *ReverseDnJoinVirtualAttributeResponse) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *ReverseDnJoinVirtualAttributeResponse) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *ReverseDnJoinVirtualAttributeResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ReverseDnJoinVirtualAttributeResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *ReverseDnJoinVirtualAttributeResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *ReverseDnJoinVirtualAttributeResponse) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *ReverseDnJoinVirtualAttributeResponse) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *ReverseDnJoinVirtualAttributeResponse) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *ReverseDnJoinVirtualAttributeResponse) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ReverseDnJoinVirtualAttributeResponse) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ReverseDnJoinVirtualAttributeResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *ReverseDnJoinVirtualAttributeResponse) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *ReverseDnJoinVirtualAttributeResponse) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *ReverseDnJoinVirtualAttributeResponse) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetConflictBehavior

`func (o *ReverseDnJoinVirtualAttributeResponse) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *ReverseDnJoinVirtualAttributeResponse) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *ReverseDnJoinVirtualAttributeResponse) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *ReverseDnJoinVirtualAttributeResponse) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *ReverseDnJoinVirtualAttributeResponse) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *ReverseDnJoinVirtualAttributeResponse) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *ReverseDnJoinVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *ReverseDnJoinVirtualAttributeResponse) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *ReverseDnJoinVirtualAttributeResponse) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *ReverseDnJoinVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *ReverseDnJoinVirtualAttributeResponse) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *ReverseDnJoinVirtualAttributeResponse) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *ReverseDnJoinVirtualAttributeResponse) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *ReverseDnJoinVirtualAttributeResponse) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *ReverseDnJoinVirtualAttributeResponse) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *ReverseDnJoinVirtualAttributeResponse) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


