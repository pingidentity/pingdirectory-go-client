# AddMemberVirtualAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new Virtual Attribute | 
**Schemas** | [**[]EnummemberVirtualAttributeSchemaUrn**](EnummemberVirtualAttributeSchemaUrn.md) |  | 
**ConflictBehavior** | Pointer to [**EnumvirtualAttributeConflictBehaviorProp**](EnumvirtualAttributeConflictBehaviorProp.md) |  | [optional] 
**AllowRetrievingMembership** | **bool** | Indicates whether to handle requests that request all values for the virtual attribute. | 
**Filter** | Pointer to **[]string** | Specifies the search filters to be applied against entries to determine if the virtual attribute is to be generated for those entries. | [optional] 
**AllowIndexConflicts** | Pointer to **bool** | Indicates whether the server should allow creating or altering this virtual attribute definition even if it conflicts with one or more indexes defined in the server. | [optional] 
**Description** | Pointer to **string** | A description for this Virtual Attribute | [optional] 
**Enabled** | **bool** | Indicates whether the Virtual Attribute is enabled for use. | 
**AttributeType** | **string** | Specifies the attribute type for the attribute whose values are to be dynamically assigned by the virtual attribute. | 
**BaseDN** | Pointer to **[]string** | Specifies the base DNs for the branches containing entries that are eligible to use this virtual attribute. | [optional] 
**GroupDN** | Pointer to **[]string** | Specifies the DNs of the groups whose members can be eligible to use this virtual attribute. | [optional] 
**ClientConnectionPolicy** | Pointer to **[]string** | Specifies a set of client connection policies for which this Virtual Attribute should be generated. If this is undefined, then this Virtual Attribute will always be generated. If it is associated with one or more client connection policies, then this Virtual Attribute will be generated only for operations requested by clients assigned to one of those client connection policies. | [optional] 
**RequireExplicitRequestByName** | Pointer to **bool** | Indicates whether attributes of this type must be explicitly included by name in the list of requested attributes. Note that this will only apply to virtual attributes which are associated with an attribute type that is operational. It will be ignored for virtual attributes associated with a non-operational attribute type. | [optional] 
**MultipleVirtualAttributeEvaluationOrderIndex** | Pointer to **int32** | Specifies the order in which virtual attribute definitions for the same attribute type will be evaluated when generating values for an entry. | [optional] 
**MultipleVirtualAttributeMergeBehavior** | Pointer to [**EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp**](EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp.md) |  | [optional] 

## Methods

### NewAddMemberVirtualAttributeRequest

`func NewAddMemberVirtualAttributeRequest(name string, schemas []EnummemberVirtualAttributeSchemaUrn, allowRetrievingMembership bool, enabled bool, attributeType string, ) *AddMemberVirtualAttributeRequest`

NewAddMemberVirtualAttributeRequest instantiates a new AddMemberVirtualAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMemberVirtualAttributeRequestWithDefaults

`func NewAddMemberVirtualAttributeRequestWithDefaults() *AddMemberVirtualAttributeRequest`

NewAddMemberVirtualAttributeRequestWithDefaults instantiates a new AddMemberVirtualAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddMemberVirtualAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddMemberVirtualAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddMemberVirtualAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSchemas

`func (o *AddMemberVirtualAttributeRequest) GetSchemas() []EnummemberVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddMemberVirtualAttributeRequest) GetSchemasOk() (*[]EnummemberVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddMemberVirtualAttributeRequest) SetSchemas(v []EnummemberVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetConflictBehavior

`func (o *AddMemberVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *AddMemberVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *AddMemberVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *AddMemberVirtualAttributeRequest) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetAllowRetrievingMembership

`func (o *AddMemberVirtualAttributeRequest) GetAllowRetrievingMembership() bool`

GetAllowRetrievingMembership returns the AllowRetrievingMembership field if non-nil, zero value otherwise.

### GetAllowRetrievingMembershipOk

`func (o *AddMemberVirtualAttributeRequest) GetAllowRetrievingMembershipOk() (*bool, bool)`

GetAllowRetrievingMembershipOk returns a tuple with the AllowRetrievingMembership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRetrievingMembership

`func (o *AddMemberVirtualAttributeRequest) SetAllowRetrievingMembership(v bool)`

SetAllowRetrievingMembership sets AllowRetrievingMembership field to given value.


### GetFilter

`func (o *AddMemberVirtualAttributeRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddMemberVirtualAttributeRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddMemberVirtualAttributeRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddMemberVirtualAttributeRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *AddMemberVirtualAttributeRequest) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *AddMemberVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *AddMemberVirtualAttributeRequest) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *AddMemberVirtualAttributeRequest) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.

### GetDescription

`func (o *AddMemberVirtualAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddMemberVirtualAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddMemberVirtualAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddMemberVirtualAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddMemberVirtualAttributeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddMemberVirtualAttributeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddMemberVirtualAttributeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *AddMemberVirtualAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddMemberVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddMemberVirtualAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *AddMemberVirtualAttributeRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddMemberVirtualAttributeRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddMemberVirtualAttributeRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddMemberVirtualAttributeRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *AddMemberVirtualAttributeRequest) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *AddMemberVirtualAttributeRequest) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *AddMemberVirtualAttributeRequest) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *AddMemberVirtualAttributeRequest) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *AddMemberVirtualAttributeRequest) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *AddMemberVirtualAttributeRequest) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *AddMemberVirtualAttributeRequest) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *AddMemberVirtualAttributeRequest) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *AddMemberVirtualAttributeRequest) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *AddMemberVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *AddMemberVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *AddMemberVirtualAttributeRequest) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddMemberVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *AddMemberVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddMemberVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddMemberVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *AddMemberVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *AddMemberVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *AddMemberVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *AddMemberVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


