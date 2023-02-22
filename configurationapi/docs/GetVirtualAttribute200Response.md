# GetVirtualAttribute200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Virtual Attribute | 
**Schemas** | [**[]EnumthirdPartyVirtualAttributeSchemaUrn**](EnumthirdPartyVirtualAttributeSchemaUrn.md) |  | 
**ConflictBehavior** | Pointer to [**EnumvirtualAttributeConflictBehaviorProp**](EnumvirtualAttributeConflictBehaviorProp.md) |  | [optional] 
**SourceAttribute** | **string** | Specifies the source attribute containing the values to use for this virtual attribute. | 
**SourceEntryDNAttribute** | Pointer to **string** | Specifies the attribute containing the DN of another entry from which to obtain the source attribute providing the values for this virtual attribute. | [optional] 
**SourceEntryDNMap** | Pointer to **string** | Specifies a DN map that will be used to identify the entry from which to obtain the source attribute providing the values for this virtual attribute. | [optional] 
**BypassAccessControlForSearches** | Pointer to **bool** | Indicates whether searches performed by this virtual attribute provider should be exempted from access control restrictions. | [optional] 
**Description** | Pointer to **string** | A description for this Virtual Attribute | [optional] 
**Enabled** | **bool** | Indicates whether the Virtual Attribute is enabled for use. | 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute. | 
**BaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches containing entries that are eligible to use this virtual attribute. | [optional] 
**GroupDN** | Pointer to **[]string** | Specifies the DNs of the groups whose members can be eligible to use this virtual attribute. | [optional] 
**Filter** | Pointer to **[]string** | Specifies the search filters to be applied against entries to determine if the virtual attribute is to be generated for those entries. | [optional] 
**ClientConnectionPolicy** | Pointer to **[]string** | Specifies a set of client connection policies for which this Virtual Attribute should be generated. If this is undefined, then this Virtual Attribute will always be generated. If it is associated with one or more client connection policies, then this Virtual Attribute will be generated only for operations requested by clients assigned to one of those client connection policies. | [optional] 
**RequireExplicitRequestByName** | Pointer to **bool** | Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type. | [optional] 
**MultipleVirtualAttributeEvaluationOrderIndex** | Pointer to **int32** | Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry. | [optional] 
**MultipleVirtualAttributeMergeBehavior** | Pointer to [**EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp**](EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp.md) |  | [optional] 
**AllowIndexConflicts** | Pointer to **bool** | Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server. | [optional] 
**ExcludeOperationalAttributes** | Pointer to **bool** | Indicates whether all operational attributes should be excluded from the generated checksum. | [optional] 
**ExcludedAttribute** | Pointer to **[]string** | Specifies the attributes that should be excluded from the checksum calculation. | [optional] 
**ValuePattern** | **[]string** | Specifies a pattern for constructing the virtual attribute value using fixed text and attribute values from the entry. | 
**DirectMembershipsOnly** | Pointer to **bool** | Specifies whether to only include groups in which the user is directly associated with and the membership maybe modified via the group entry. Groups in which the user&#39;s membership is derived dynamically or through nested groups will not be included. | [optional] 
**IncludedGroupFilter** | Pointer to **string** | A search filter that will be used to identify which groups should be included in the values of the virtual attribute. With no value defined (which is the default behavior), all groups that contain the target user will be included. | [optional] 
**RewriteSearchFilters** | Pointer to [**EnumvirtualAttributeRewriteSearchFiltersProp**](EnumvirtualAttributeRewriteSearchFiltersProp.md) |  | [optional] 
**JoinDNAttribute** | **string** | The attribute whose values are the DNs of the entries to be joined with the search result entry. | 
**JoinBaseDNType** | [**EnumvirtualAttributeJoinBaseDNTypeProp**](EnumvirtualAttributeJoinBaseDNTypeProp.md) |  | 
**JoinCustomBaseDN** | Pointer to **string** | The fixed, administrator-specified base DN for the internal searches used to identify joined entries. | [optional] 
**JoinScope** | Pointer to [**EnumvirtualAttributeJoinScopeProp**](EnumvirtualAttributeJoinScopeProp.md) |  | [optional] 
**JoinSizeLimit** | Pointer to **int32** | The maximum number of entries that may be joined with the source entry, which also corresponds to the maximum number of values that the virtual attribute provider will generate for an entry. | [optional] 
**JoinFilter** | Pointer to **string** | An optional filter that specifies additional criteria for identifying joined entries. If a join-filter value is specified, then only entries matching that filter (in addition to satisfying the other join criteria) will be joined with the search result entry. | [optional] 
**JoinAttribute** | Pointer to **[]string** | An optional set of the names of the attributes to include with joined entries. | [optional] 
**ReferencedByAttribute** | **[]string** | The name or OID of an attribute type whose values will be searched for references to the target entry. The attribute type must be defined in the server schema, must have a syntax of either \&quot;distinguished name\&quot; or \&quot;name and optional UID\&quot;, and must be indexed for equality. | 
**ReferenceSearchBaseDN** | Pointer to **[]string** | The base DN that will be used when searching for references to the target entry. If no reference search base DN is specified, the default behavior will be to search below all public naming contexts configured in the server. | [optional] 
**Value** | **[]string** | Specifies the values to be included in the virtual attribute. | 
**ReturnUtcTime** | Pointer to **bool** | Indicates whether to return current time in UTC. | [optional] 
**IncludeMilliseconds** | Pointer to **bool** | Indicates whether the current time includes millisecond precision. | [optional] 
**SequenceNumberAttribute** | **string** | Specifies the name or OID of the attribute which contains the sequence number from which unique identifiers are generated. The attribute should have Integer syntax or a String syntax permitting integer values. If this property is modified then the filter property should be updated accordingly so that only entries containing the sequence number attribute are eligible to have a value generated for this virtual attribute. | 
**JoinSourceAttribute** | **string** | The attribute containing the value(s) in the source entry to use to identify related entries. | 
**JoinTargetAttribute** | **string** | The attribute in target entries whose value(s) match values of the source attribute in the source entry. | 
**JoinMatchAll** | Pointer to **bool** | Indicates whether joined entries will be required to have all values for the source attribute, or only at least one of its values. | [optional] 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Virtual Attribute. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Virtual Attribute. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**AllowRetrievingMembership** | **bool** | Indicates whether to handle requests that request all values for the virtual attribute. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Virtual Attribute. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Virtual Attribute. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewGetVirtualAttribute200Response

`func NewGetVirtualAttribute200Response(id string, schemas []EnumthirdPartyVirtualAttributeSchemaUrn, sourceAttribute string, enabled bool, attributeType string, valuePattern []string, joinDNAttribute string, joinBaseDNType EnumvirtualAttributeJoinBaseDNTypeProp, referencedByAttribute []string, value []string, sequenceNumberAttribute string, joinSourceAttribute string, joinTargetAttribute string, scriptClass string, allowRetrievingMembership bool, extensionClass string, ) *GetVirtualAttribute200Response`

NewGetVirtualAttribute200Response instantiates a new GetVirtualAttribute200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVirtualAttribute200ResponseWithDefaults

`func NewGetVirtualAttribute200ResponseWithDefaults() *GetVirtualAttribute200Response`

NewGetVirtualAttribute200ResponseWithDefaults instantiates a new GetVirtualAttribute200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetVirtualAttribute200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetVirtualAttribute200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetVirtualAttribute200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetVirtualAttribute200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetVirtualAttribute200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetVirtualAttribute200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetVirtualAttribute200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetVirtualAttribute200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *GetVirtualAttribute200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVirtualAttribute200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVirtualAttribute200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *GetVirtualAttribute200Response) GetSchemas() []EnumthirdPartyVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetVirtualAttribute200Response) GetSchemasOk() (*[]EnumthirdPartyVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetVirtualAttribute200Response) SetSchemas(v []EnumthirdPartyVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConflictBehavior

`func (o *GetVirtualAttribute200Response) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *GetVirtualAttribute200Response) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *GetVirtualAttribute200Response) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *GetVirtualAttribute200Response) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetSourceAttribute

`func (o *GetVirtualAttribute200Response) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *GetVirtualAttribute200Response) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *GetVirtualAttribute200Response) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetSourceEntryDNAttribute

`func (o *GetVirtualAttribute200Response) GetSourceEntryDNAttribute() string`

GetSourceEntryDNAttribute returns the SourceEntryDNAttribute field if non-nil, zero value otherwise.

### GetSourceEntryDNAttributeOk

`func (o *GetVirtualAttribute200Response) GetSourceEntryDNAttributeOk() (*string, bool)`

GetSourceEntryDNAttributeOk returns a tuple with the SourceEntryDNAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntryDNAttribute

`func (o *GetVirtualAttribute200Response) SetSourceEntryDNAttribute(v string)`

SetSourceEntryDNAttribute sets SourceEntryDNAttribute field to given value.

### HasSourceEntryDNAttribute

`func (o *GetVirtualAttribute200Response) HasSourceEntryDNAttribute() bool`

HasSourceEntryDNAttribute returns a boolean if a field has been set.

### GetSourceEntryDNMap

`func (o *GetVirtualAttribute200Response) GetSourceEntryDNMap() string`

GetSourceEntryDNMap returns the SourceEntryDNMap field if non-nil, zero value otherwise.

### GetSourceEntryDNMapOk

`func (o *GetVirtualAttribute200Response) GetSourceEntryDNMapOk() (*string, bool)`

GetSourceEntryDNMapOk returns a tuple with the SourceEntryDNMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntryDNMap

`func (o *GetVirtualAttribute200Response) SetSourceEntryDNMap(v string)`

SetSourceEntryDNMap sets SourceEntryDNMap field to given value.

### HasSourceEntryDNMap

`func (o *GetVirtualAttribute200Response) HasSourceEntryDNMap() bool`

HasSourceEntryDNMap returns a boolean if a field has been set.

### GetBypassAccessControlForSearches

`func (o *GetVirtualAttribute200Response) GetBypassAccessControlForSearches() bool`

GetBypassAccessControlForSearches returns the BypassAccessControlForSearches field if non-nil, zero value otherwise.

### GetBypassAccessControlForSearchesOk

`func (o *GetVirtualAttribute200Response) GetBypassAccessControlForSearchesOk() (*bool, bool)`

GetBypassAccessControlForSearchesOk returns a tuple with the BypassAccessControlForSearches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassAccessControlForSearches

`func (o *GetVirtualAttribute200Response) SetBypassAccessControlForSearches(v bool)`

SetBypassAccessControlForSearches sets BypassAccessControlForSearches field to given value.

### HasBypassAccessControlForSearches

`func (o *GetVirtualAttribute200Response) HasBypassAccessControlForSearches() bool`

HasBypassAccessControlForSearches returns a boolean if a field has been set.

### GetDescription

`func (o *GetVirtualAttribute200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVirtualAttribute200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVirtualAttribute200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVirtualAttribute200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetVirtualAttribute200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetVirtualAttribute200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetVirtualAttribute200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *GetVirtualAttribute200Response) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *GetVirtualAttribute200Response) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *GetVirtualAttribute200Response) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *GetVirtualAttribute200Response) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *GetVirtualAttribute200Response) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *GetVirtualAttribute200Response) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *GetVirtualAttribute200Response) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *GetVirtualAttribute200Response) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *GetVirtualAttribute200Response) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *GetVirtualAttribute200Response) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *GetVirtualAttribute200Response) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *GetVirtualAttribute200Response) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetVirtualAttribute200Response) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetVirtualAttribute200Response) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetVirtualAttribute200Response) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *GetVirtualAttribute200Response) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *GetVirtualAttribute200Response) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *GetVirtualAttribute200Response) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *GetVirtualAttribute200Response) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *GetVirtualAttribute200Response) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *GetVirtualAttribute200Response) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *GetVirtualAttribute200Response) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *GetVirtualAttribute200Response) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *GetVirtualAttribute200Response) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *GetVirtualAttribute200Response) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *GetVirtualAttribute200Response) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *GetVirtualAttribute200Response) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *GetVirtualAttribute200Response) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *GetVirtualAttribute200Response) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *GetVirtualAttribute200Response) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *GetVirtualAttribute200Response) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *GetVirtualAttribute200Response) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *GetVirtualAttribute200Response) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *GetVirtualAttribute200Response) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *GetVirtualAttribute200Response) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.

### GetExcludeOperationalAttributes

`func (o *GetVirtualAttribute200Response) GetExcludeOperationalAttributes() bool`

GetExcludeOperationalAttributes returns the ExcludeOperationalAttributes field if non-nil, zero value otherwise.

### GetExcludeOperationalAttributesOk

`func (o *GetVirtualAttribute200Response) GetExcludeOperationalAttributesOk() (*bool, bool)`

GetExcludeOperationalAttributesOk returns a tuple with the ExcludeOperationalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeOperationalAttributes

`func (o *GetVirtualAttribute200Response) SetExcludeOperationalAttributes(v bool)`

SetExcludeOperationalAttributes sets ExcludeOperationalAttributes field to given value.

### HasExcludeOperationalAttributes

`func (o *GetVirtualAttribute200Response) HasExcludeOperationalAttributes() bool`

HasExcludeOperationalAttributes returns a boolean if a field has been set.

### GetExcludedAttribute

`func (o *GetVirtualAttribute200Response) GetExcludedAttribute() []string`

GetExcludedAttribute returns the ExcludedAttribute field if non-nil, zero value otherwise.

### GetExcludedAttributeOk

`func (o *GetVirtualAttribute200Response) GetExcludedAttributeOk() (*[]string, bool)`

GetExcludedAttributeOk returns a tuple with the ExcludedAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAttribute

`func (o *GetVirtualAttribute200Response) SetExcludedAttribute(v []string)`

SetExcludedAttribute sets ExcludedAttribute field to given value.

### HasExcludedAttribute

`func (o *GetVirtualAttribute200Response) HasExcludedAttribute() bool`

HasExcludedAttribute returns a boolean if a field has been set.

### GetValuePattern

`func (o *GetVirtualAttribute200Response) GetValuePattern() []string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *GetVirtualAttribute200Response) GetValuePatternOk() (*[]string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *GetVirtualAttribute200Response) SetValuePattern(v []string)`

SetValuePattern sets ValuePattern field to given value.


### GetDirectMembershipsOnly

`func (o *GetVirtualAttribute200Response) GetDirectMembershipsOnly() bool`

GetDirectMembershipsOnly returns the DirectMembershipsOnly field if non-nil, zero value otherwise.

### GetDirectMembershipsOnlyOk

`func (o *GetVirtualAttribute200Response) GetDirectMembershipsOnlyOk() (*bool, bool)`

GetDirectMembershipsOnlyOk returns a tuple with the DirectMembershipsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectMembershipsOnly

`func (o *GetVirtualAttribute200Response) SetDirectMembershipsOnly(v bool)`

SetDirectMembershipsOnly sets DirectMembershipsOnly field to given value.

### HasDirectMembershipsOnly

`func (o *GetVirtualAttribute200Response) HasDirectMembershipsOnly() bool`

HasDirectMembershipsOnly returns a boolean if a field has been set.

### GetIncludedGroupFilter

`func (o *GetVirtualAttribute200Response) GetIncludedGroupFilter() string`

GetIncludedGroupFilter returns the IncludedGroupFilter field if non-nil, zero value otherwise.

### GetIncludedGroupFilterOk

`func (o *GetVirtualAttribute200Response) GetIncludedGroupFilterOk() (*string, bool)`

GetIncludedGroupFilterOk returns a tuple with the IncludedGroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedGroupFilter

`func (o *GetVirtualAttribute200Response) SetIncludedGroupFilter(v string)`

SetIncludedGroupFilter sets IncludedGroupFilter field to given value.

### HasIncludedGroupFilter

`func (o *GetVirtualAttribute200Response) HasIncludedGroupFilter() bool`

HasIncludedGroupFilter returns a boolean if a field has been set.

### GetRewriteSearchFilters

`func (o *GetVirtualAttribute200Response) GetRewriteSearchFilters() EnumvirtualAttributeRewriteSearchFiltersProp`

GetRewriteSearchFilters returns the RewriteSearchFilters field if non-nil, zero value otherwise.

### GetRewriteSearchFiltersOk

`func (o *GetVirtualAttribute200Response) GetRewriteSearchFiltersOk() (*EnumvirtualAttributeRewriteSearchFiltersProp, bool)`

GetRewriteSearchFiltersOk returns a tuple with the RewriteSearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewriteSearchFilters

`func (o *GetVirtualAttribute200Response) SetRewriteSearchFilters(v EnumvirtualAttributeRewriteSearchFiltersProp)`

SetRewriteSearchFilters sets RewriteSearchFilters field to given value.

### HasRewriteSearchFilters

`func (o *GetVirtualAttribute200Response) HasRewriteSearchFilters() bool`

HasRewriteSearchFilters returns a boolean if a field has been set.

### GetJoinDNAttribute

`func (o *GetVirtualAttribute200Response) GetJoinDNAttribute() string`

GetJoinDNAttribute returns the JoinDNAttribute field if non-nil, zero value otherwise.

### GetJoinDNAttributeOk

`func (o *GetVirtualAttribute200Response) GetJoinDNAttributeOk() (*string, bool)`

GetJoinDNAttributeOk returns a tuple with the JoinDNAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDNAttribute

`func (o *GetVirtualAttribute200Response) SetJoinDNAttribute(v string)`

SetJoinDNAttribute sets JoinDNAttribute field to given value.


### GetJoinBaseDNType

`func (o *GetVirtualAttribute200Response) GetJoinBaseDNType() EnumvirtualAttributeJoinBaseDNTypeProp`

GetJoinBaseDNType returns the JoinBaseDNType field if non-nil, zero value otherwise.

### GetJoinBaseDNTypeOk

`func (o *GetVirtualAttribute200Response) GetJoinBaseDNTypeOk() (*EnumvirtualAttributeJoinBaseDNTypeProp, bool)`

GetJoinBaseDNTypeOk returns a tuple with the JoinBaseDNType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinBaseDNType

`func (o *GetVirtualAttribute200Response) SetJoinBaseDNType(v EnumvirtualAttributeJoinBaseDNTypeProp)`

SetJoinBaseDNType sets JoinBaseDNType field to given value.


### GetJoinCustomBaseDN

`func (o *GetVirtualAttribute200Response) GetJoinCustomBaseDN() string`

GetJoinCustomBaseDN returns the JoinCustomBaseDN field if non-nil, zero value otherwise.

### GetJoinCustomBaseDNOk

`func (o *GetVirtualAttribute200Response) GetJoinCustomBaseDNOk() (*string, bool)`

GetJoinCustomBaseDNOk returns a tuple with the JoinCustomBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinCustomBaseDN

`func (o *GetVirtualAttribute200Response) SetJoinCustomBaseDN(v string)`

SetJoinCustomBaseDN sets JoinCustomBaseDN field to given value.

### HasJoinCustomBaseDN

`func (o *GetVirtualAttribute200Response) HasJoinCustomBaseDN() bool`

HasJoinCustomBaseDN returns a boolean if a field has been set.

### GetJoinScope

`func (o *GetVirtualAttribute200Response) GetJoinScope() EnumvirtualAttributeJoinScopeProp`

GetJoinScope returns the JoinScope field if non-nil, zero value otherwise.

### GetJoinScopeOk

`func (o *GetVirtualAttribute200Response) GetJoinScopeOk() (*EnumvirtualAttributeJoinScopeProp, bool)`

GetJoinScopeOk returns a tuple with the JoinScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinScope

`func (o *GetVirtualAttribute200Response) SetJoinScope(v EnumvirtualAttributeJoinScopeProp)`

SetJoinScope sets JoinScope field to given value.

### HasJoinScope

`func (o *GetVirtualAttribute200Response) HasJoinScope() bool`

HasJoinScope returns a boolean if a field has been set.

### GetJoinSizeLimit

`func (o *GetVirtualAttribute200Response) GetJoinSizeLimit() int32`

GetJoinSizeLimit returns the JoinSizeLimit field if non-nil, zero value otherwise.

### GetJoinSizeLimitOk

`func (o *GetVirtualAttribute200Response) GetJoinSizeLimitOk() (*int32, bool)`

GetJoinSizeLimitOk returns a tuple with the JoinSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSizeLimit

`func (o *GetVirtualAttribute200Response) SetJoinSizeLimit(v int32)`

SetJoinSizeLimit sets JoinSizeLimit field to given value.

### HasJoinSizeLimit

`func (o *GetVirtualAttribute200Response) HasJoinSizeLimit() bool`

HasJoinSizeLimit returns a boolean if a field has been set.

### GetJoinFilter

`func (o *GetVirtualAttribute200Response) GetJoinFilter() string`

GetJoinFilter returns the JoinFilter field if non-nil, zero value otherwise.

### GetJoinFilterOk

`func (o *GetVirtualAttribute200Response) GetJoinFilterOk() (*string, bool)`

GetJoinFilterOk returns a tuple with the JoinFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinFilter

`func (o *GetVirtualAttribute200Response) SetJoinFilter(v string)`

SetJoinFilter sets JoinFilter field to given value.

### HasJoinFilter

`func (o *GetVirtualAttribute200Response) HasJoinFilter() bool`

HasJoinFilter returns a boolean if a field has been set.

### GetJoinAttribute

`func (o *GetVirtualAttribute200Response) GetJoinAttribute() []string`

GetJoinAttribute returns the JoinAttribute field if non-nil, zero value otherwise.

### GetJoinAttributeOk

`func (o *GetVirtualAttribute200Response) GetJoinAttributeOk() (*[]string, bool)`

GetJoinAttributeOk returns a tuple with the JoinAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinAttribute

`func (o *GetVirtualAttribute200Response) SetJoinAttribute(v []string)`

SetJoinAttribute sets JoinAttribute field to given value.

### HasJoinAttribute

`func (o *GetVirtualAttribute200Response) HasJoinAttribute() bool`

HasJoinAttribute returns a boolean if a field has been set.

### GetReferencedByAttribute

`func (o *GetVirtualAttribute200Response) GetReferencedByAttribute() []string`

GetReferencedByAttribute returns the ReferencedByAttribute field if non-nil, zero value otherwise.

### GetReferencedByAttributeOk

`func (o *GetVirtualAttribute200Response) GetReferencedByAttributeOk() (*[]string, bool)`

GetReferencedByAttributeOk returns a tuple with the ReferencedByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByAttribute

`func (o *GetVirtualAttribute200Response) SetReferencedByAttribute(v []string)`

SetReferencedByAttribute sets ReferencedByAttribute field to given value.


### GetReferenceSearchBaseDN

`func (o *GetVirtualAttribute200Response) GetReferenceSearchBaseDN() []string`

GetReferenceSearchBaseDN returns the ReferenceSearchBaseDN field if non-nil, zero value otherwise.

### GetReferenceSearchBaseDNOk

`func (o *GetVirtualAttribute200Response) GetReferenceSearchBaseDNOk() (*[]string, bool)`

GetReferenceSearchBaseDNOk returns a tuple with the ReferenceSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSearchBaseDN

`func (o *GetVirtualAttribute200Response) SetReferenceSearchBaseDN(v []string)`

SetReferenceSearchBaseDN sets ReferenceSearchBaseDN field to given value.

### HasReferenceSearchBaseDN

`func (o *GetVirtualAttribute200Response) HasReferenceSearchBaseDN() bool`

HasReferenceSearchBaseDN returns a boolean if a field has been set.

### GetValue

`func (o *GetVirtualAttribute200Response) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetVirtualAttribute200Response) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetVirtualAttribute200Response) SetValue(v []string)`

SetValue sets Value field to given value.


### GetReturnUtcTime

`func (o *GetVirtualAttribute200Response) GetReturnUtcTime() bool`

GetReturnUtcTime returns the ReturnUtcTime field if non-nil, zero value otherwise.

### GetReturnUtcTimeOk

`func (o *GetVirtualAttribute200Response) GetReturnUtcTimeOk() (*bool, bool)`

GetReturnUtcTimeOk returns a tuple with the ReturnUtcTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUtcTime

`func (o *GetVirtualAttribute200Response) SetReturnUtcTime(v bool)`

SetReturnUtcTime sets ReturnUtcTime field to given value.

### HasReturnUtcTime

`func (o *GetVirtualAttribute200Response) HasReturnUtcTime() bool`

HasReturnUtcTime returns a boolean if a field has been set.

### GetIncludeMilliseconds

`func (o *GetVirtualAttribute200Response) GetIncludeMilliseconds() bool`

GetIncludeMilliseconds returns the IncludeMilliseconds field if non-nil, zero value otherwise.

### GetIncludeMillisecondsOk

`func (o *GetVirtualAttribute200Response) GetIncludeMillisecondsOk() (*bool, bool)`

GetIncludeMillisecondsOk returns a tuple with the IncludeMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMilliseconds

`func (o *GetVirtualAttribute200Response) SetIncludeMilliseconds(v bool)`

SetIncludeMilliseconds sets IncludeMilliseconds field to given value.

### HasIncludeMilliseconds

`func (o *GetVirtualAttribute200Response) HasIncludeMilliseconds() bool`

HasIncludeMilliseconds returns a boolean if a field has been set.

### GetSequenceNumberAttribute

`func (o *GetVirtualAttribute200Response) GetSequenceNumberAttribute() string`

GetSequenceNumberAttribute returns the SequenceNumberAttribute field if non-nil, zero value otherwise.

### GetSequenceNumberAttributeOk

`func (o *GetVirtualAttribute200Response) GetSequenceNumberAttributeOk() (*string, bool)`

GetSequenceNumberAttributeOk returns a tuple with the SequenceNumberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumberAttribute

`func (o *GetVirtualAttribute200Response) SetSequenceNumberAttribute(v string)`

SetSequenceNumberAttribute sets SequenceNumberAttribute field to given value.


### GetJoinSourceAttribute

`func (o *GetVirtualAttribute200Response) GetJoinSourceAttribute() string`

GetJoinSourceAttribute returns the JoinSourceAttribute field if non-nil, zero value otherwise.

### GetJoinSourceAttributeOk

`func (o *GetVirtualAttribute200Response) GetJoinSourceAttributeOk() (*string, bool)`

GetJoinSourceAttributeOk returns a tuple with the JoinSourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinSourceAttribute

`func (o *GetVirtualAttribute200Response) SetJoinSourceAttribute(v string)`

SetJoinSourceAttribute sets JoinSourceAttribute field to given value.


### GetJoinTargetAttribute

`func (o *GetVirtualAttribute200Response) GetJoinTargetAttribute() string`

GetJoinTargetAttribute returns the JoinTargetAttribute field if non-nil, zero value otherwise.

### GetJoinTargetAttributeOk

`func (o *GetVirtualAttribute200Response) GetJoinTargetAttributeOk() (*string, bool)`

GetJoinTargetAttributeOk returns a tuple with the JoinTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTargetAttribute

`func (o *GetVirtualAttribute200Response) SetJoinTargetAttribute(v string)`

SetJoinTargetAttribute sets JoinTargetAttribute field to given value.


### GetJoinMatchAll

`func (o *GetVirtualAttribute200Response) GetJoinMatchAll() bool`

GetJoinMatchAll returns the JoinMatchAll field if non-nil, zero value otherwise.

### GetJoinMatchAllOk

`func (o *GetVirtualAttribute200Response) GetJoinMatchAllOk() (*bool, bool)`

GetJoinMatchAllOk returns a tuple with the JoinMatchAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinMatchAll

`func (o *GetVirtualAttribute200Response) SetJoinMatchAll(v bool)`

SetJoinMatchAll sets JoinMatchAll field to given value.

### HasJoinMatchAll

`func (o *GetVirtualAttribute200Response) HasJoinMatchAll() bool`

HasJoinMatchAll returns a boolean if a field has been set.

### GetScriptClass

`func (o *GetVirtualAttribute200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GetVirtualAttribute200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GetVirtualAttribute200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GetVirtualAttribute200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GetVirtualAttribute200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GetVirtualAttribute200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GetVirtualAttribute200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetAllowRetrievingMembership

`func (o *GetVirtualAttribute200Response) GetAllowRetrievingMembership() bool`

GetAllowRetrievingMembership returns the AllowRetrievingMembership field if non-nil, zero value otherwise.

### GetAllowRetrievingMembershipOk

`func (o *GetVirtualAttribute200Response) GetAllowRetrievingMembershipOk() (*bool, bool)`

GetAllowRetrievingMembershipOk returns a tuple with the AllowRetrievingMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRetrievingMembership

`func (o *GetVirtualAttribute200Response) SetAllowRetrievingMembership(v bool)`

SetAllowRetrievingMembership sets AllowRetrievingMembership field to given value.


### GetExtensionClass

`func (o *GetVirtualAttribute200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *GetVirtualAttribute200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *GetVirtualAttribute200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *GetVirtualAttribute200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *GetVirtualAttribute200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *GetVirtualAttribute200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *GetVirtualAttribute200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


