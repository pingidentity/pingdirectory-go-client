# AddEntryDnVirtualAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new Virtual Attribute | 
**Schemas** | [**[]EnumentryDnVirtualAttributeSchemaUrn**](EnumentryDnVirtualAttributeSchemaUrn.md) |  | 
**ConflictBehavior** | Pointer to [**EnumvirtualAttributeConflictBehaviorProp**](EnumvirtualAttributeConflictBehaviorProp.md) |  | [optional] 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute. | 
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

### NewAddEntryDnVirtualAttributeRequest

`func NewAddEntryDnVirtualAttributeRequest(name string, schemas []EnumentryDnVirtualAttributeSchemaUrn, attributeType string, enabled bool, ) *AddEntryDnVirtualAttributeRequest`

NewAddEntryDnVirtualAttributeRequest instantiates a new AddEntryDnVirtualAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEntryDnVirtualAttributeRequestWithDefaults

`func NewAddEntryDnVirtualAttributeRequestWithDefaults() *AddEntryDnVirtualAttributeRequest`

NewAddEntryDnVirtualAttributeRequestWithDefaults instantiates a new AddEntryDnVirtualAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddEntryDnVirtualAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddEntryDnVirtualAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddEntryDnVirtualAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSchemas

`func (o *AddEntryDnVirtualAttributeRequest) GetSchemas() []EnumentryDnVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddEntryDnVirtualAttributeRequest) GetSchemasOk() (*[]EnumentryDnVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddEntryDnVirtualAttributeRequest) SetSchemas(v []EnumentryDnVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConflictBehavior

`func (o *AddEntryDnVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *AddEntryDnVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *AddEntryDnVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *AddEntryDnVirtualAttributeRequest) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetAttributeType

`func (o *AddEntryDnVirtualAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddEntryDnVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddEntryDnVirtualAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetDescription

`func (o *AddEntryDnVirtualAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddEntryDnVirtualAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddEntryDnVirtualAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddEntryDnVirtualAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddEntryDnVirtualAttributeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddEntryDnVirtualAttributeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddEntryDnVirtualAttributeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetBaseDN

`func (o *AddEntryDnVirtualAttributeRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddEntryDnVirtualAttributeRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddEntryDnVirtualAttributeRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddEntryDnVirtualAttributeRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *AddEntryDnVirtualAttributeRequest) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *AddEntryDnVirtualAttributeRequest) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *AddEntryDnVirtualAttributeRequest) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *AddEntryDnVirtualAttributeRequest) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *AddEntryDnVirtualAttributeRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddEntryDnVirtualAttributeRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddEntryDnVirtualAttributeRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddEntryDnVirtualAttributeRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *AddEntryDnVirtualAttributeRequest) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *AddEntryDnVirtualAttributeRequest) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *AddEntryDnVirtualAttributeRequest) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *AddEntryDnVirtualAttributeRequest) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *AddEntryDnVirtualAttributeRequest) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *AddEntryDnVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *AddEntryDnVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *AddEntryDnVirtualAttributeRequest) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddEntryDnVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *AddEntryDnVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddEntryDnVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddEntryDnVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *AddEntryDnVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *AddEntryDnVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *AddEntryDnVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *AddEntryDnVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *AddEntryDnVirtualAttributeRequest) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *AddEntryDnVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *AddEntryDnVirtualAttributeRequest) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *AddEntryDnVirtualAttributeRequest) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


