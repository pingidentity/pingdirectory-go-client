# CurrentTimeVirtualAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumcurrentTimeVirtualAttributeSchemaUrn**](EnumcurrentTimeVirtualAttributeSchemaUrn.md) |  | 
**Id** | **string** | Name of the Virtual Attribute | 
**ConflictBehavior** | Pointer to [**EnumvirtualAttributeConflictBehaviorProp**](EnumvirtualAttributeConflictBehaviorProp.md) |  | [optional] 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute. | 
**ReturnUtcTime** | Pointer to **bool** | Indicates whether to return current time in UTC. | [optional] 
**IncludeMilliseconds** | Pointer to **bool** | Indicates whether the current time includes millisecond precision. | [optional] 
**Description** | Pointer to **string** | A description for this Virtual Attribute | [optional] 
**Enabled** | **bool** | Indicates whether the Virtual Attribute is enabled for use. | 
**BaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches containing entries that are eligible to use this virtual attribute. | [optional] 
**GroupDN** | Pointer to **[]string** | Specifies the DNs of the groups whose members can be eligible to use this virtual attribute. | [optional] 
**Filter** | Pointer to **[]string** | Specifies the search filters to be applied against entries to determine if the virtual attribute is to be generated for those entries. | [optional] 
**ClientConnectionPolicy** | Pointer to **[]string** | Specifies a set of client connection policies for which this Virtual Attribute should be generated. If this is undefined, then this Virtual Attribute will always be generated. If it is associated with one or more client connection policies, then this Virtual Attribute will be generated only for operations requested by clients assigned to one of those client connection policies. | [optional] 
**RequireExplicitRequestByName** | Pointer to **bool** | Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type. | [optional] 
**MultipleVirtualAttributeEvaluationOrderIndex** | Pointer to **int32** | Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry. | [optional] 
**MultipleVirtualAttributeMergeBehavior** | Pointer to [**EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp**](EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp.md) |  | [optional] 
**AllowIndexConflicts** | Pointer to **bool** | Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server. | [optional] 

## Methods

### NewCurrentTimeVirtualAttributeResponse

`func NewCurrentTimeVirtualAttributeResponse(schemas []EnumcurrentTimeVirtualAttributeSchemaUrn, id string, attributeType string, enabled bool, ) *CurrentTimeVirtualAttributeResponse`

NewCurrentTimeVirtualAttributeResponse instantiates a new CurrentTimeVirtualAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentTimeVirtualAttributeResponseWithDefaults

`func NewCurrentTimeVirtualAttributeResponseWithDefaults() *CurrentTimeVirtualAttributeResponse`

NewCurrentTimeVirtualAttributeResponseWithDefaults instantiates a new CurrentTimeVirtualAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CurrentTimeVirtualAttributeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CurrentTimeVirtualAttributeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CurrentTimeVirtualAttributeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CurrentTimeVirtualAttributeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CurrentTimeVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CurrentTimeVirtualAttributeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CurrentTimeVirtualAttributeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CurrentTimeVirtualAttributeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *CurrentTimeVirtualAttributeResponse) GetSchemas() []EnumcurrentTimeVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CurrentTimeVirtualAttributeResponse) GetSchemasOk() (*[]EnumcurrentTimeVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CurrentTimeVirtualAttributeResponse) SetSchemas(v []EnumcurrentTimeVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *CurrentTimeVirtualAttributeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentTimeVirtualAttributeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentTimeVirtualAttributeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetConflictBehavior

`func (o *CurrentTimeVirtualAttributeResponse) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *CurrentTimeVirtualAttributeResponse) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *CurrentTimeVirtualAttributeResponse) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *CurrentTimeVirtualAttributeResponse) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetAttributeType

`func (o *CurrentTimeVirtualAttributeResponse) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *CurrentTimeVirtualAttributeResponse) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *CurrentTimeVirtualAttributeResponse) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetReturnUtcTime

`func (o *CurrentTimeVirtualAttributeResponse) GetReturnUtcTime() bool`

GetReturnUtcTime returns the ReturnUtcTime field if non-nil, zero value otherwise.

### GetReturnUtcTimeOk

`func (o *CurrentTimeVirtualAttributeResponse) GetReturnUtcTimeOk() (*bool, bool)`

GetReturnUtcTimeOk returns a tuple with the ReturnUtcTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUtcTime

`func (o *CurrentTimeVirtualAttributeResponse) SetReturnUtcTime(v bool)`

SetReturnUtcTime sets ReturnUtcTime field to given value.

### HasReturnUtcTime

`func (o *CurrentTimeVirtualAttributeResponse) HasReturnUtcTime() bool`

HasReturnUtcTime returns a boolean if a field has been set.

### GetIncludeMilliseconds

`func (o *CurrentTimeVirtualAttributeResponse) GetIncludeMilliseconds() bool`

GetIncludeMilliseconds returns the IncludeMilliseconds field if non-nil, zero value otherwise.

### GetIncludeMillisecondsOk

`func (o *CurrentTimeVirtualAttributeResponse) GetIncludeMillisecondsOk() (*bool, bool)`

GetIncludeMillisecondsOk returns a tuple with the IncludeMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMilliseconds

`func (o *CurrentTimeVirtualAttributeResponse) SetIncludeMilliseconds(v bool)`

SetIncludeMilliseconds sets IncludeMilliseconds field to given value.

### HasIncludeMilliseconds

`func (o *CurrentTimeVirtualAttributeResponse) HasIncludeMilliseconds() bool`

HasIncludeMilliseconds returns a boolean if a field has been set.

### GetDescription

`func (o *CurrentTimeVirtualAttributeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CurrentTimeVirtualAttributeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CurrentTimeVirtualAttributeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CurrentTimeVirtualAttributeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CurrentTimeVirtualAttributeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CurrentTimeVirtualAttributeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CurrentTimeVirtualAttributeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseDN

`func (o *CurrentTimeVirtualAttributeResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *CurrentTimeVirtualAttributeResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *CurrentTimeVirtualAttributeResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *CurrentTimeVirtualAttributeResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *CurrentTimeVirtualAttributeResponse) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *CurrentTimeVirtualAttributeResponse) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *CurrentTimeVirtualAttributeResponse) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *CurrentTimeVirtualAttributeResponse) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *CurrentTimeVirtualAttributeResponse) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CurrentTimeVirtualAttributeResponse) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CurrentTimeVirtualAttributeResponse) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CurrentTimeVirtualAttributeResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *CurrentTimeVirtualAttributeResponse) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *CurrentTimeVirtualAttributeResponse) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *CurrentTimeVirtualAttributeResponse) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *CurrentTimeVirtualAttributeResponse) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *CurrentTimeVirtualAttributeResponse) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *CurrentTimeVirtualAttributeResponse) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *CurrentTimeVirtualAttributeResponse) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *CurrentTimeVirtualAttributeResponse) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *CurrentTimeVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *CurrentTimeVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *CurrentTimeVirtualAttributeResponse) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *CurrentTimeVirtualAttributeResponse) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *CurrentTimeVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *CurrentTimeVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *CurrentTimeVirtualAttributeResponse) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *CurrentTimeVirtualAttributeResponse) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *CurrentTimeVirtualAttributeResponse) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *CurrentTimeVirtualAttributeResponse) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *CurrentTimeVirtualAttributeResponse) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *CurrentTimeVirtualAttributeResponse) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


