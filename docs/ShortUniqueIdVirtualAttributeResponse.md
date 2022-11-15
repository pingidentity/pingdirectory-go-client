# ShortUniqueIdVirtualAttributeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumshortUniqueIdVirtualAttributeSchemaUrn**](EnumshortUniqueIdVirtualAttributeSchemaUrn.md) |  | 
**SequenceNumberAttribute** | **string** | Specifies the name or OID of the attribute which contains the sequence number from which unique identifiers are generated. The attribute should have Integer syntax or a String syntax permitting integer values. If this property is modified then the filter property should be updated accordingly so that only entries containing the sequence number attribute are eligible to have a value generated for this virtual attribute. | 
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

### NewShortUniqueIdVirtualAttributeResponse

`func NewShortUniqueIdVirtualAttributeResponse(schemas []EnumshortUniqueIdVirtualAttributeSchemaUrn, sequenceNumberAttribute string, enabled bool, attributeType string, ) *ShortUniqueIdVirtualAttributeResponse`

NewShortUniqueIdVirtualAttributeResponse instantiates a new ShortUniqueIdVirtualAttributeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortUniqueIdVirtualAttributeResponseWithDefaults

`func NewShortUniqueIdVirtualAttributeResponseWithDefaults() *ShortUniqueIdVirtualAttributeResponse`

NewShortUniqueIdVirtualAttributeResponseWithDefaults instantiates a new ShortUniqueIdVirtualAttributeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ShortUniqueIdVirtualAttributeResponse) GetSchemas() []EnumshortUniqueIdVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetSchemasOk() (*[]EnumshortUniqueIdVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ShortUniqueIdVirtualAttributeResponse) SetSchemas(v []EnumshortUniqueIdVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSequenceNumberAttribute

`func (o *ShortUniqueIdVirtualAttributeResponse) GetSequenceNumberAttribute() string`

GetSequenceNumberAttribute returns the SequenceNumberAttribute field if non-nil, zero value otherwise.

### GetSequenceNumberAttributeOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetSequenceNumberAttributeOk() (*string, bool)`

GetSequenceNumberAttributeOk returns a tuple with the SequenceNumberAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumberAttribute

`func (o *ShortUniqueIdVirtualAttributeResponse) SetSequenceNumberAttribute(v string)`

SetSequenceNumberAttribute sets SequenceNumberAttribute field to given value.


### GetDescription

`func (o *ShortUniqueIdVirtualAttributeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ShortUniqueIdVirtualAttributeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ShortUniqueIdVirtualAttributeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ShortUniqueIdVirtualAttributeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ShortUniqueIdVirtualAttributeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *ShortUniqueIdVirtualAttributeResponse) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *ShortUniqueIdVirtualAttributeResponse) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *ShortUniqueIdVirtualAttributeResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ShortUniqueIdVirtualAttributeResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *ShortUniqueIdVirtualAttributeResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *ShortUniqueIdVirtualAttributeResponse) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *ShortUniqueIdVirtualAttributeResponse) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *ShortUniqueIdVirtualAttributeResponse) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *ShortUniqueIdVirtualAttributeResponse) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ShortUniqueIdVirtualAttributeResponse) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ShortUniqueIdVirtualAttributeResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *ShortUniqueIdVirtualAttributeResponse) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *ShortUniqueIdVirtualAttributeResponse) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *ShortUniqueIdVirtualAttributeResponse) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetConflictBehavior

`func (o *ShortUniqueIdVirtualAttributeResponse) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *ShortUniqueIdVirtualAttributeResponse) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *ShortUniqueIdVirtualAttributeResponse) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *ShortUniqueIdVirtualAttributeResponse) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *ShortUniqueIdVirtualAttributeResponse) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *ShortUniqueIdVirtualAttributeResponse) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *ShortUniqueIdVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *ShortUniqueIdVirtualAttributeResponse) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *ShortUniqueIdVirtualAttributeResponse) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *ShortUniqueIdVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *ShortUniqueIdVirtualAttributeResponse) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *ShortUniqueIdVirtualAttributeResponse) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *ShortUniqueIdVirtualAttributeResponse) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *ShortUniqueIdVirtualAttributeResponse) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *ShortUniqueIdVirtualAttributeResponse) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *ShortUniqueIdVirtualAttributeResponse) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


