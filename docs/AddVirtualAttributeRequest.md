# AddVirtualAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new Virtual Attribute | 
**Schemas** | [**[]EnumthirdPartyVirtualAttributeSchemaUrn**](EnumthirdPartyVirtualAttributeSchemaUrn.md) |  | 
**ConflictBehavior** | Pointer to [**EnumvirtualAttributeConflictBehaviorProp**](EnumvirtualAttributeConflictBehaviorProp.md) |  | [optional] 
**SourceAttribute** | **string** | Specifies the source attribute containing the values to use for this virtual attribute. | 
**SourceEntryDNAttribute** | Pointer to **string** | Specifies the attribute containing the DN of another entry from which to obtain the source attribute providing the values for this virtual attribute. | [optional] 
**SourceEntryDNMap** | Pointer to **string** | Specifies a DN map that will be used to identify the entry from which to obtain the source attribute providing the values for this virtual attribute. | [optional] 
**BypassAccessControlForSearches** | Pointer to **bool** | Indicates whether searches performed by this virtual attribute provider should be exempted from access control restrictions. | [optional] 
**Description** | Pointer to **string** | A description for this Virtual Attribute | [optional] 
**Enabled** | **bool** | Indicates whether the Virtual Attribute is enabled for use. | 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute. | 
**BaseDN** | Pointer to **[]string** |  | [optional] 
**GroupDN** | Pointer to **[]string** |  | [optional] 
**Filter** | Pointer to **[]string** |  | [optional] 
**ClientConnectionPolicy** | Pointer to **[]string** |  | [optional] 
**RequireExplicitRequestByName** | Pointer to **bool** | Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type. | [optional] 
**MultipleVirtualAttributeEvaluationOrderIndex** | Pointer to **int32** | Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry. | [optional] 
**MultipleVirtualAttributeMergeBehavior** | Pointer to [**EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp**](EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp.md) |  | [optional] 
**AllowIndexConflicts** | Pointer to **bool** | Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server. | [optional] 
**ValuePattern** | **[]string** |  | 
**DirectMembershipsOnly** | Pointer to **bool** | Specifies whether to only include groups in which the user is directly associated with and the membership maybe modified via the group entry. Groups in which the user&#39;s membership is derived dynamically or through nested groups will not be included. | [optional] 
**IncludedGroupFilter** | Pointer to **string** | A search filter that will be used to identify which groups should be included in the values of the virtual attribute. With no value defined (which is the default behavior), all groups that contain the target user will be included. | [optional] 
**RewriteSearchFilters** | Pointer to [**EnumvirtualAttributeRewriteSearchFiltersProp**](EnumvirtualAttributeRewriteSearchFiltersProp.md) |  | [optional] 
**JoinDNAttribute** | **string** | The attribute whose values are the DNs of the entries to be joined with the search result entry. | 
**JoinBaseDNType** | [**EnumvirtualAttributeJoinBaseDNTypeProp**](EnumvirtualAttributeJoinBaseDNTypeProp.md) |  | 
**JoinCustomBaseDN** | Pointer to **string** | The fixed, administrator-specified base DN for the internal searches used to identify joined entries. | [optional] 
**JoinScope** | Pointer to [**EnumvirtualAttributeJoinScopeProp**](EnumvirtualAttributeJoinScopeProp.md) |  | [optional] 
**JoinSizeLimit** | Pointer to **int32** | The maximum number of entries that may be joined with the source entry, which also corresponds to the maximum number of values that the virtual attribute provider will generate for an entry. | [optional] 
**JoinFilter** | Pointer to **string** | An optional filter that specifies additional criteria for identifying joined entries. If a join-filter value is specified, then only entries matching that filter (in addition to satisfying the other join criteria) will be joined with the search result entry. | [optional] 
**JoinAttribute** | Pointer to **[]string** |  | [optional] 
**ReferencedByAttribute** | **[]string** |  | 
**ReferenceSearchBaseDN** | Pointer to **[]string** |  | [optional] 
**Value** | **[]string** |  | 
**JoinSourceAttribute** | **string** | The attribute containing the value(s) in the source entry to use to identify related entries. | 
**JoinTargetAttribute** | **string** | The attribute in target entries whose value(s) match values of the source attribute in the source entry. | 
**JoinMatchAll** | Pointer to **bool** | Indicates whether joined entries will be required to have all values for the source attribute, or only at least one of its values. | [optional] 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Virtual Attribute. | 
**ScriptArgument** | Pointer to **[]string** |  | [optional] 
**AllowRetrievingMembership** | **bool** | Indicates whether to handle requests that request all values for the virtual attribute. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Virtual Attribute. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddVirtualAttributeRequest

`func NewAddVirtualAttributeRequest(name string, schemas []EnumthirdPartyVirtualAttributeSchemaUrn, sourceAttribute string, enabled bool, attributeType string, valuePattern []string, joinDNAttribute string, joinBaseDNType EnumvirtualAttributeJoinBaseDNTypeProp, referencedByAttribute []string, value []string, joinSourceAttribute string, joinTargetAttribute string, scriptClass string, allowRetrievingMembership bool, extensionClass string, ) *AddVirtualAttributeRequest`

NewAddVirtualAttributeRequest instantiates a new AddVirtualAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVirtualAttributeRequestWithDefaults

`func NewAddVirtualAttributeRequestWithDefaults() *AddVirtualAttributeRequest`

NewAddVirtualAttributeRequestWithDefaults instantiates a new AddVirtualAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddVirtualAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVirtualAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVirtualAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSchemas

`func (o *AddVirtualAttributeRequest) GetSchemas() []EnumthirdPartyVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddVirtualAttributeRequest) GetSchemasOk() (*[]EnumthirdPartyVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddVirtualAttributeRequest) SetSchemas(v []EnumthirdPartyVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConflictBehavior

`func (o *AddVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *AddVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *AddVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *AddVirtualAttributeRequest) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetSourceAttribute

`func (o *AddVirtualAttributeRequest) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *AddVirtualAttributeRequest) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *AddVirtualAttributeRequest) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetSourceEntryDNAttribute

`func (o *AddVirtualAttributeRequest) GetSourceEntryDNAttribute() string`

GetSourceEntryDNAttribute returns the SourceEntryDNAttribute field if non-nil, zero value otherwise.

### GetSourceEntryDNAttributeOk

`func (o *AddVirtualAttributeRequest) GetSourceEntryDNAttributeOk() (*string, bool)`

GetSourceEntryDNAttributeOk returns a tuple with the SourceEntryDNAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntryDNAttribute

`func (o *AddVirtualAttributeRequest) SetSourceEntryDNAttribute(v string)`

SetSourceEntryDNAttribute sets SourceEntryDNAttribute field to given value.

### HasSourceEntryDNAttribute

`func (o *AddVirtualAttributeRequest) HasSourceEntryDNAttribute() bool`

HasSourceEntryDNAttribute returns a boolean if a field has been set.

### GetSourceEntryDNMap

`func (o *AddVirtualAttributeRequest) GetSourceEntryDNMap() string`

GetSourceEntryDNMap returns the SourceEntryDNMap field if non-nil, zero value otherwise.

### GetSourceEntryDNMapOk

`func (o *AddVirtualAttributeRequest) GetSourceEntryDNMapOk() (*string, bool)`

GetSourceEntryDNMapOk returns a tuple with the SourceEntryDNMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntryDNMap

`func (o *AddVirtualAttributeRequest) SetSourceEntryDNMap(v string)`

SetSourceEntryDNMap sets SourceEntryDNMap field to given value.

### HasSourceEntryDNMap

`func (o *AddVirtualAttributeRequest) HasSourceEntryDNMap() bool`

HasSourceEntryDNMap returns a boolean if a field has been set.

### GetBypassAccessControlForSearches

`func (o *AddVirtualAttributeRequest) GetBypassAccessControlForSearches() bool`

GetBypassAccessControlForSearches returns the BypassAccessControlForSearches field if non-nil, zero value otherwise.

### GetBypassAccessControlForSearchesOk

`func (o *AddVirtualAttributeRequest) GetBypassAccessControlForSearchesOk() (*bool, bool)`

GetBypassAccessControlForSearchesOk returns a tuple with the BypassAccessControlForSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAccessControlForSearches

`func (o *AddVirtualAttributeRequest) SetBypassAccessControlForSearches(v bool)`

SetBypassAccessControlForSearches sets BypassAccessControlForSearches field to given value.

### HasBypassAccessControlForSearches

`func (o *AddVirtualAttributeRequest) HasBypassAccessControlForSearches() bool`

HasBypassAccessControlForSearches returns a boolean if a field has been set.

### GetDescription

`func (o *AddVirtualAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVirtualAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVirtualAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVirtualAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddVirtualAttributeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddVirtualAttributeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddVirtualAttributeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *AddVirtualAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddVirtualAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *AddVirtualAttributeRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddVirtualAttributeRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddVirtualAttributeRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddVirtualAttributeRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *AddVirtualAttributeRequest) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *AddVirtualAttributeRequest) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *AddVirtualAttributeRequest) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *AddVirtualAttributeRequest) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *AddVirtualAttributeRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddVirtualAttributeRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddVirtualAttributeRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddVirtualAttributeRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *AddVirtualAttributeRequest) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *AddVirtualAttributeRequest) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *AddVirtualAttributeRequest) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *AddVirtualAttributeRequest) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *AddVirtualAttributeRequest) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *AddVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *AddVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *AddVirtualAttributeRequest) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *AddVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *AddVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *AddVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *AddVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *AddVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *AddVirtualAttributeRequest) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *AddVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *AddVirtualAttributeRequest) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *AddVirtualAttributeRequest) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.

### GetValuePattern

`func (o *AddVirtualAttributeRequest) GetValuePattern() []string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *AddVirtualAttributeRequest) GetValuePatternOk() (*[]string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *AddVirtualAttributeRequest) SetValuePattern(v []string)`

SetValuePattern sets ValuePattern field to given value.


### GetDirectMembershipsOnly

`func (o *AddVirtualAttributeRequest) GetDirectMembershipsOnly() bool`

GetDirectMembershipsOnly returns the DirectMembershipsOnly field if non-nil, zero value otherwise.

### GetDirectMembershipsOnlyOk

`func (o *AddVirtualAttributeRequest) GetDirectMembershipsOnlyOk() (*bool, bool)`

GetDirectMembershipsOnlyOk returns a tuple with the DirectMembershipsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMembershipsOnly

`func (o *AddVirtualAttributeRequest) SetDirectMembershipsOnly(v bool)`

SetDirectMembershipsOnly sets DirectMembershipsOnly field to given value.

### HasDirectMembershipsOnly

`func (o *AddVirtualAttributeRequest) HasDirectMembershipsOnly() bool`

HasDirectMembershipsOnly returns a boolean if a field has been set.

### GetIncludedGroupFilter

`func (o *AddVirtualAttributeRequest) GetIncludedGroupFilter() string`

GetIncludedGroupFilter returns the IncludedGroupFilter field if non-nil, zero value otherwise.

### GetIncludedGroupFilterOk

`func (o *AddVirtualAttributeRequest) GetIncludedGroupFilterOk() (*string, bool)`

GetIncludedGroupFilterOk returns a tuple with the IncludedGroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedGroupFilter

`func (o *AddVirtualAttributeRequest) SetIncludedGroupFilter(v string)`

SetIncludedGroupFilter sets IncludedGroupFilter field to given value.

### HasIncludedGroupFilter

`func (o *AddVirtualAttributeRequest) HasIncludedGroupFilter() bool`

HasIncludedGroupFilter returns a boolean if a field has been set.

### GetRewriteSearchFilters

`func (o *AddVirtualAttributeRequest) GetRewriteSearchFilters() EnumvirtualAttributeRewriteSearchFiltersProp`

GetRewriteSearchFilters returns the RewriteSearchFilters field if non-nil, zero value otherwise.

### GetRewriteSearchFiltersOk

`func (o *AddVirtualAttributeRequest) GetRewriteSearchFiltersOk() (*EnumvirtualAttributeRewriteSearchFiltersProp, bool)`

GetRewriteSearchFiltersOk returns a tuple with the RewriteSearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteSearchFilters

`func (o *AddVirtualAttributeRequest) SetRewriteSearchFilters(v EnumvirtualAttributeRewriteSearchFiltersProp)`

SetRewriteSearchFilters sets RewriteSearchFilters field to given value.

### HasRewriteSearchFilters

`func (o *AddVirtualAttributeRequest) HasRewriteSearchFilters() bool`

HasRewriteSearchFilters returns a boolean if a field has been set.

### GetJoinDNAttribute

`func (o *AddVirtualAttributeRequest) GetJoinDNAttribute() string`

GetJoinDNAttribute returns the JoinDNAttribute field if non-nil, zero value otherwise.

### GetJoinDNAttributeOk

`func (o *AddVirtualAttributeRequest) GetJoinDNAttributeOk() (*string, bool)`

GetJoinDNAttributeOk returns a tuple with the JoinDNAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDNAttribute

`func (o *AddVirtualAttributeRequest) SetJoinDNAttribute(v string)`

SetJoinDNAttribute sets JoinDNAttribute field to given value.


### GetJoinBaseDNType

`func (o *AddVirtualAttributeRequest) GetJoinBaseDNType() EnumvirtualAttributeJoinBaseDNTypeProp`

GetJoinBaseDNType returns the JoinBaseDNType field if non-nil, zero value otherwise.

### GetJoinBaseDNTypeOk

`func (o *AddVirtualAttributeRequest) GetJoinBaseDNTypeOk() (*EnumvirtualAttributeJoinBaseDNTypeProp, bool)`

GetJoinBaseDNTypeOk returns a tuple with the JoinBaseDNType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinBaseDNType

`func (o *AddVirtualAttributeRequest) SetJoinBaseDNType(v EnumvirtualAttributeJoinBaseDNTypeProp)`

SetJoinBaseDNType sets JoinBaseDNType field to given value.


### GetJoinCustomBaseDN

`func (o *AddVirtualAttributeRequest) GetJoinCustomBaseDN() string`

GetJoinCustomBaseDN returns the JoinCustomBaseDN field if non-nil, zero value otherwise.

### GetJoinCustomBaseDNOk

`func (o *AddVirtualAttributeRequest) GetJoinCustomBaseDNOk() (*string, bool)`

GetJoinCustomBaseDNOk returns a tuple with the JoinCustomBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinCustomBaseDN

`func (o *AddVirtualAttributeRequest) SetJoinCustomBaseDN(v string)`

SetJoinCustomBaseDN sets JoinCustomBaseDN field to given value.

### HasJoinCustomBaseDN

`func (o *AddVirtualAttributeRequest) HasJoinCustomBaseDN() bool`

HasJoinCustomBaseDN returns a boolean if a field has been set.

### GetJoinScope

`func (o *AddVirtualAttributeRequest) GetJoinScope() EnumvirtualAttributeJoinScopeProp`

GetJoinScope returns the JoinScope field if non-nil, zero value otherwise.

### GetJoinScopeOk

`func (o *AddVirtualAttributeRequest) GetJoinScopeOk() (*EnumvirtualAttributeJoinScopeProp, bool)`

GetJoinScopeOk returns a tuple with the JoinScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinScope

`func (o *AddVirtualAttributeRequest) SetJoinScope(v EnumvirtualAttributeJoinScopeProp)`

SetJoinScope sets JoinScope field to given value.

### HasJoinScope

`func (o *AddVirtualAttributeRequest) HasJoinScope() bool`

HasJoinScope returns a boolean if a field has been set.

### GetJoinSizeLimit

`func (o *AddVirtualAttributeRequest) GetJoinSizeLimit() int32`

GetJoinSizeLimit returns the JoinSizeLimit field if non-nil, zero value otherwise.

### GetJoinSizeLimitOk

`func (o *AddVirtualAttributeRequest) GetJoinSizeLimitOk() (*int32, bool)`

GetJoinSizeLimitOk returns a tuple with the JoinSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSizeLimit

`func (o *AddVirtualAttributeRequest) SetJoinSizeLimit(v int32)`

SetJoinSizeLimit sets JoinSizeLimit field to given value.

### HasJoinSizeLimit

`func (o *AddVirtualAttributeRequest) HasJoinSizeLimit() bool`

HasJoinSizeLimit returns a boolean if a field has been set.

### GetJoinFilter

`func (o *AddVirtualAttributeRequest) GetJoinFilter() string`

GetJoinFilter returns the JoinFilter field if non-nil, zero value otherwise.

### GetJoinFilterOk

`func (o *AddVirtualAttributeRequest) GetJoinFilterOk() (*string, bool)`

GetJoinFilterOk returns a tuple with the JoinFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinFilter

`func (o *AddVirtualAttributeRequest) SetJoinFilter(v string)`

SetJoinFilter sets JoinFilter field to given value.

### HasJoinFilter

`func (o *AddVirtualAttributeRequest) HasJoinFilter() bool`

HasJoinFilter returns a boolean if a field has been set.

### GetJoinAttribute

`func (o *AddVirtualAttributeRequest) GetJoinAttribute() []string`

GetJoinAttribute returns the JoinAttribute field if non-nil, zero value otherwise.

### GetJoinAttributeOk

`func (o *AddVirtualAttributeRequest) GetJoinAttributeOk() (*[]string, bool)`

GetJoinAttributeOk returns a tuple with the JoinAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinAttribute

`func (o *AddVirtualAttributeRequest) SetJoinAttribute(v []string)`

SetJoinAttribute sets JoinAttribute field to given value.

### HasJoinAttribute

`func (o *AddVirtualAttributeRequest) HasJoinAttribute() bool`

HasJoinAttribute returns a boolean if a field has been set.

### GetReferencedByAttribute

`func (o *AddVirtualAttributeRequest) GetReferencedByAttribute() []string`

GetReferencedByAttribute returns the ReferencedByAttribute field if non-nil, zero value otherwise.

### GetReferencedByAttributeOk

`func (o *AddVirtualAttributeRequest) GetReferencedByAttributeOk() (*[]string, bool)`

GetReferencedByAttributeOk returns a tuple with the ReferencedByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByAttribute

`func (o *AddVirtualAttributeRequest) SetReferencedByAttribute(v []string)`

SetReferencedByAttribute sets ReferencedByAttribute field to given value.


### GetReferenceSearchBaseDN

`func (o *AddVirtualAttributeRequest) GetReferenceSearchBaseDN() []string`

GetReferenceSearchBaseDN returns the ReferenceSearchBaseDN field if non-nil, zero value otherwise.

### GetReferenceSearchBaseDNOk

`func (o *AddVirtualAttributeRequest) GetReferenceSearchBaseDNOk() (*[]string, bool)`

GetReferenceSearchBaseDNOk returns a tuple with the ReferenceSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSearchBaseDN

`func (o *AddVirtualAttributeRequest) SetReferenceSearchBaseDN(v []string)`

SetReferenceSearchBaseDN sets ReferenceSearchBaseDN field to given value.

### HasReferenceSearchBaseDN

`func (o *AddVirtualAttributeRequest) HasReferenceSearchBaseDN() bool`

HasReferenceSearchBaseDN returns a boolean if a field has been set.

### GetValue

`func (o *AddVirtualAttributeRequest) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddVirtualAttributeRequest) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddVirtualAttributeRequest) SetValue(v []string)`

SetValue sets Value field to given value.


### GetJoinSourceAttribute

`func (o *AddVirtualAttributeRequest) GetJoinSourceAttribute() string`

GetJoinSourceAttribute returns the JoinSourceAttribute field if non-nil, zero value otherwise.

### GetJoinSourceAttributeOk

`func (o *AddVirtualAttributeRequest) GetJoinSourceAttributeOk() (*string, bool)`

GetJoinSourceAttributeOk returns a tuple with the JoinSourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSourceAttribute

`func (o *AddVirtualAttributeRequest) SetJoinSourceAttribute(v string)`

SetJoinSourceAttribute sets JoinSourceAttribute field to given value.


### GetJoinTargetAttribute

`func (o *AddVirtualAttributeRequest) GetJoinTargetAttribute() string`

GetJoinTargetAttribute returns the JoinTargetAttribute field if non-nil, zero value otherwise.

### GetJoinTargetAttributeOk

`func (o *AddVirtualAttributeRequest) GetJoinTargetAttributeOk() (*string, bool)`

GetJoinTargetAttributeOk returns a tuple with the JoinTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTargetAttribute

`func (o *AddVirtualAttributeRequest) SetJoinTargetAttribute(v string)`

SetJoinTargetAttribute sets JoinTargetAttribute field to given value.


### GetJoinMatchAll

`func (o *AddVirtualAttributeRequest) GetJoinMatchAll() bool`

GetJoinMatchAll returns the JoinMatchAll field if non-nil, zero value otherwise.

### GetJoinMatchAllOk

`func (o *AddVirtualAttributeRequest) GetJoinMatchAllOk() (*bool, bool)`

GetJoinMatchAllOk returns a tuple with the JoinMatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinMatchAll

`func (o *AddVirtualAttributeRequest) SetJoinMatchAll(v bool)`

SetJoinMatchAll sets JoinMatchAll field to given value.

### HasJoinMatchAll

`func (o *AddVirtualAttributeRequest) HasJoinMatchAll() bool`

HasJoinMatchAll returns a boolean if a field has been set.

### GetScriptClass

`func (o *AddVirtualAttributeRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddVirtualAttributeRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddVirtualAttributeRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddVirtualAttributeRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddVirtualAttributeRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddVirtualAttributeRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddVirtualAttributeRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetAllowRetrievingMembership

`func (o *AddVirtualAttributeRequest) GetAllowRetrievingMembership() bool`

GetAllowRetrievingMembership returns the AllowRetrievingMembership field if non-nil, zero value otherwise.

### GetAllowRetrievingMembershipOk

`func (o *AddVirtualAttributeRequest) GetAllowRetrievingMembershipOk() (*bool, bool)`

GetAllowRetrievingMembershipOk returns a tuple with the AllowRetrievingMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRetrievingMembership

`func (o *AddVirtualAttributeRequest) SetAllowRetrievingMembership(v bool)`

SetAllowRetrievingMembership sets AllowRetrievingMembership field to given value.


### GetExtensionClass

`func (o *AddVirtualAttributeRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddVirtualAttributeRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddVirtualAttributeRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddVirtualAttributeRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddVirtualAttributeRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddVirtualAttributeRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddVirtualAttributeRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


