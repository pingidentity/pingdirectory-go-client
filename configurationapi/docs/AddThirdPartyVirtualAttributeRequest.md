# AddThirdPartyVirtualAttributeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new Virtual Attribute | 
**Schemas** | [**[]EnumthirdPartyVirtualAttributeSchemaUrn**](EnumthirdPartyVirtualAttributeSchemaUrn.md) |  | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Virtual Attribute. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Virtual Attribute. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
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

### NewAddThirdPartyVirtualAttributeRequest

`func NewAddThirdPartyVirtualAttributeRequest(name string, schemas []EnumthirdPartyVirtualAttributeSchemaUrn, extensionClass string, enabled bool, attributeType string, ) *AddThirdPartyVirtualAttributeRequest`

NewAddThirdPartyVirtualAttributeRequest instantiates a new AddThirdPartyVirtualAttributeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddThirdPartyVirtualAttributeRequestWithDefaults

`func NewAddThirdPartyVirtualAttributeRequestWithDefaults() *AddThirdPartyVirtualAttributeRequest`

NewAddThirdPartyVirtualAttributeRequestWithDefaults instantiates a new AddThirdPartyVirtualAttributeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddThirdPartyVirtualAttributeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddThirdPartyVirtualAttributeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSchemas

`func (o *AddThirdPartyVirtualAttributeRequest) GetSchemas() []EnumthirdPartyVirtualAttributeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetSchemasOk() (*[]EnumthirdPartyVirtualAttributeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddThirdPartyVirtualAttributeRequest) SetSchemas(v []EnumthirdPartyVirtualAttributeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetExtensionClass

`func (o *AddThirdPartyVirtualAttributeRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddThirdPartyVirtualAttributeRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddThirdPartyVirtualAttributeRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddThirdPartyVirtualAttributeRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddThirdPartyVirtualAttributeRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetDescription

`func (o *AddThirdPartyVirtualAttributeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddThirdPartyVirtualAttributeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddThirdPartyVirtualAttributeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddThirdPartyVirtualAttributeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddThirdPartyVirtualAttributeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAttributeType

`func (o *AddThirdPartyVirtualAttributeRequest) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddThirdPartyVirtualAttributeRequest) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.


### GetBaseDN

`func (o *AddThirdPartyVirtualAttributeRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddThirdPartyVirtualAttributeRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *AddThirdPartyVirtualAttributeRequest) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetGroupDN

`func (o *AddThirdPartyVirtualAttributeRequest) GetGroupDN() []string`

GetGroupDN returns the GroupDN field if non-nil, zero value otherwise.

### GetGroupDNOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetGroupDNOk() (*[]string, bool)`

GetGroupDNOk returns a tuple with the GroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDN

`func (o *AddThirdPartyVirtualAttributeRequest) SetGroupDN(v []string)`

SetGroupDN sets GroupDN field to given value.

### HasGroupDN

`func (o *AddThirdPartyVirtualAttributeRequest) HasGroupDN() bool`

HasGroupDN returns a boolean if a field has been set.

### GetFilter

`func (o *AddThirdPartyVirtualAttributeRequest) GetFilter() []string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetFilterOk() (*[]string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddThirdPartyVirtualAttributeRequest) SetFilter(v []string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AddThirdPartyVirtualAttributeRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetClientConnectionPolicy

`func (o *AddThirdPartyVirtualAttributeRequest) GetClientConnectionPolicy() []string`

GetClientConnectionPolicy returns the ClientConnectionPolicy field if non-nil, zero value otherwise.

### GetClientConnectionPolicyOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetClientConnectionPolicyOk() (*[]string, bool)`

GetClientConnectionPolicyOk returns a tuple with the ClientConnectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionPolicy

`func (o *AddThirdPartyVirtualAttributeRequest) SetClientConnectionPolicy(v []string)`

SetClientConnectionPolicy sets ClientConnectionPolicy field to given value.

### HasClientConnectionPolicy

`func (o *AddThirdPartyVirtualAttributeRequest) HasClientConnectionPolicy() bool`

HasClientConnectionPolicy returns a boolean if a field has been set.

### GetConflictBehavior

`func (o *AddThirdPartyVirtualAttributeRequest) GetConflictBehavior() EnumvirtualAttributeConflictBehaviorProp`

GetConflictBehavior returns the ConflictBehavior field if non-nil, zero value otherwise.

### GetConflictBehaviorOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetConflictBehaviorOk() (*EnumvirtualAttributeConflictBehaviorProp, bool)`

GetConflictBehaviorOk returns a tuple with the ConflictBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictBehavior

`func (o *AddThirdPartyVirtualAttributeRequest) SetConflictBehavior(v EnumvirtualAttributeConflictBehaviorProp)`

SetConflictBehavior sets ConflictBehavior field to given value.

### HasConflictBehavior

`func (o *AddThirdPartyVirtualAttributeRequest) HasConflictBehavior() bool`

HasConflictBehavior returns a boolean if a field has been set.

### GetRequireExplicitRequestByName

`func (o *AddThirdPartyVirtualAttributeRequest) GetRequireExplicitRequestByName() bool`

GetRequireExplicitRequestByName returns the RequireExplicitRequestByName field if non-nil, zero value otherwise.

### GetRequireExplicitRequestByNameOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetRequireExplicitRequestByNameOk() (*bool, bool)`

GetRequireExplicitRequestByNameOk returns a tuple with the RequireExplicitRequestByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExplicitRequestByName

`func (o *AddThirdPartyVirtualAttributeRequest) SetRequireExplicitRequestByName(v bool)`

SetRequireExplicitRequestByName sets RequireExplicitRequestByName field to given value.

### HasRequireExplicitRequestByName

`func (o *AddThirdPartyVirtualAttributeRequest) HasRequireExplicitRequestByName() bool`

HasRequireExplicitRequestByName returns a boolean if a field has been set.

### GetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddThirdPartyVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndex() int32`

GetMultipleVirtualAttributeEvaluationOrderIndex returns the MultipleVirtualAttributeEvaluationOrderIndex field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeEvaluationOrderIndexOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetMultipleVirtualAttributeEvaluationOrderIndexOk() (*int32, bool)`

GetMultipleVirtualAttributeEvaluationOrderIndexOk returns a tuple with the MultipleVirtualAttributeEvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddThirdPartyVirtualAttributeRequest) SetMultipleVirtualAttributeEvaluationOrderIndex(v int32)`

SetMultipleVirtualAttributeEvaluationOrderIndex sets MultipleVirtualAttributeEvaluationOrderIndex field to given value.

### HasMultipleVirtualAttributeEvaluationOrderIndex

`func (o *AddThirdPartyVirtualAttributeRequest) HasMultipleVirtualAttributeEvaluationOrderIndex() bool`

HasMultipleVirtualAttributeEvaluationOrderIndex returns a boolean if a field has been set.

### GetMultipleVirtualAttributeMergeBehavior

`func (o *AddThirdPartyVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehavior() EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp`

GetMultipleVirtualAttributeMergeBehavior returns the MultipleVirtualAttributeMergeBehavior field if non-nil, zero value otherwise.

### GetMultipleVirtualAttributeMergeBehaviorOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetMultipleVirtualAttributeMergeBehaviorOk() (*EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp, bool)`

GetMultipleVirtualAttributeMergeBehaviorOk returns a tuple with the MultipleVirtualAttributeMergeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleVirtualAttributeMergeBehavior

`func (o *AddThirdPartyVirtualAttributeRequest) SetMultipleVirtualAttributeMergeBehavior(v EnumvirtualAttributeMultipleVirtualAttributeMergeBehaviorProp)`

SetMultipleVirtualAttributeMergeBehavior sets MultipleVirtualAttributeMergeBehavior field to given value.

### HasMultipleVirtualAttributeMergeBehavior

`func (o *AddThirdPartyVirtualAttributeRequest) HasMultipleVirtualAttributeMergeBehavior() bool`

HasMultipleVirtualAttributeMergeBehavior returns a boolean if a field has been set.

### GetAllowIndexConflicts

`func (o *AddThirdPartyVirtualAttributeRequest) GetAllowIndexConflicts() bool`

GetAllowIndexConflicts returns the AllowIndexConflicts field if non-nil, zero value otherwise.

### GetAllowIndexConflictsOk

`func (o *AddThirdPartyVirtualAttributeRequest) GetAllowIndexConflictsOk() (*bool, bool)`

GetAllowIndexConflictsOk returns a tuple with the AllowIndexConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIndexConflicts

`func (o *AddThirdPartyVirtualAttributeRequest) SetAllowIndexConflicts(v bool)`

SetAllowIndexConflicts sets AllowIndexConflicts field to given value.

### HasAllowIndexConflicts

`func (o *AddThirdPartyVirtualAttributeRequest) HasAllowIndexConflicts() bool`

HasAllowIndexConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


