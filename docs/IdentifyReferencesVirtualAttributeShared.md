# IdentifyReferencesVirtualAttributeShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumidentifyReferencesVirtualAttributeSchemaUrn**](EnumidentifyReferencesVirtualAttributeSchemaUrn.md) |  | 
**ReferencedByAttribute** | **[]string** |  | 
**ReferenceSearchBaseDN** | Pointer to **[]string** |  | [optional] 
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

### NewIdentifyReferencesVirtualAttributeShared

`func NewIdentifyReferencesVirtualAttributeShared(schemas []EnumidentifyReferencesVirtualAttributeSchemaUrn, referencedByAttribute []string, enabled bool, attributeType string, ) *IdentifyReferencesVirtualAttributeShared`

NewIdentifyReferencesVirtualAttributeShared instantiates a new IdentifyReferencesVirtualAttributeShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifyReferencesVirtualAttributeSharedWithDefaults

`func NewIdentifyReferencesVirtualAttributeSharedWithDefaults() *IdentifyReferencesVirtualAttributeShared`

NewIdentifyReferencesVirtualAttributeSharedWithDefaults instantiates a new IdentifyReferencesVirtualAttributeShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *IdentifyReferencesVirtualAttributeShared) GetSchemas() []EnumidentifyReferencesVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetSchemasOk() (*[]EnumidentifyReferencesVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *IdentifyReferencesVirtualAttributeShared) SetSchemas(v []EnumidentifyReferencesVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetReferencedByAttribute

`func (o *IdentifyReferencesVirtualAttributeShared) GetReferencedByAttribute() []string`

GetReferencedByAttribute returns the ReferencedByAttribute field if non-nil, zero value otherwise.

### GetReferencedByAttributeOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetReferencedByAttributeOk() (*[]string, bool)`

GetReferencedByAttributeOk returns a tuple with the ReferencedByAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByAttribute

`func (o *IdentifyReferencesVirtualAttributeShared) SetReferencedByAttribute(v []string)`

SetReferencedByAttribute sets ReferencedByAttribute field to given value.


### GetReferenceSearchBaseDN

`func (o *IdentifyReferencesVirtualAttributeShared) GetReferenceSearchBaseDN() []string`

GetReferenceSearchBaseDN returns the ReferenceSearchBaseDN field if non-nil, zero value otherwise.

### GetReferenceSearchBaseDNOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetReferenceSearchBaseDNOk() (*[]string, bool)`

GetReferenceSearchBaseDNOk returns a tuple with the ReferenceSearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSearchBaseDN

`func (o *IdentifyReferencesVirtualAttributeShared) SetReferenceSearchBaseDN(v []string)`

SetReferenceSearchBaseDN sets ReferenceSearchBaseDN field to given value.

### HasReferenceSearchBaseDN

`func (o *IdentifyReferencesVirtualAttributeShared) HasReferenceSearchBaseDN() bool`

HasReferenceSearchBaseDN returns a boolean if a field has been set.

### GetDescription

`func (o *IdentifyReferencesVirtualAttributeShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentifyReferencesVirtualAttributeShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentifyReferencesVirtualAttributeShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentifyReferencesVirtualAttributeShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentifyReferencesVirtualAttributeShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *IdentifyReferencesVirtualAttributeShared) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *IdentifyReferencesVirtualAttributeShared) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *IdentifyReferencesVirtualAttributeShared) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *IdentifyReferencesVirtualAttributeShared) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *IdentifyReferencesVirtualAttributeShared) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *IdentifyReferencesVirtualAttributeShared) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *IdentifyReferencesVirtualAttributeShared) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *IdentifyReferencesVirtualAttributeShared) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *IdentifyReferencesVirtualAttributeShared) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IdentifyReferencesVirtualAttributeShared) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IdentifyReferencesVirtualAttributeShared) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *IdentifyReferencesVirtualAttributeShared) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *IdentifyReferencesVirtualAttributeShared) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *IdentifyReferencesVirtualAttributeShared) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetConflictBehavior

`func (o *IdentifyReferencesVirtualAttributeShared) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *IdentifyReferencesVirtualAttributeShared) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *IdentifyReferencesVirtualAttributeShared) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *IdentifyReferencesVirtualAttributeShared) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *IdentifyReferencesVirtualAttributeShared) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *IdentifyReferencesVirtualAttributeShared) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *IdentifyReferencesVirtualAttributeShared) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *IdentifyReferencesVirtualAttributeShared) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *IdentifyReferencesVirtualAttributeShared) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *IdentifyReferencesVirtualAttributeShared) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *IdentifyReferencesVirtualAttributeShared) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *IdentifyReferencesVirtualAttributeShared) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *IdentifyReferencesVirtualAttributeShared) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *IdentifyReferencesVirtualAttributeShared) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *IdentifyReferencesVirtualAttributeShared) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *IdentifyReferencesVirtualAttributeShared) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


