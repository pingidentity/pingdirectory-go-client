# GloballyUniqueAttributePluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumgloballyUniqueAttributePluginSchemaUrn**](EnumgloballyUniqueAttributePluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin | 
**Type** | **[]string** | The attribute type(s) for which to enforce global uniqueness. The attribute must be indexed for equality searches in all subtree views for which uniqueness should be maintained. | 
**MultipleAttributeBehavior** | Pointer to [**EnumpluginGloballyUniqueAttributeMultipleAttributeBehaviorProp**](EnumpluginGloballyUniqueAttributeMultipleAttributeBehaviorProp.md) |  | [optional] 
**SubtreeView** | **[]string** | The subtree view(s) for which to enforce uniqueness. | 
**PreventConflictsWithSoftDeletedEntries** | Pointer to **bool** | Indicates whether this Globally Unique Attribute Plugin should attempt to prevent conflicts with soft-deleted entries (i.e., entries that have been removed in a way that leaves them in the server but in a way that makes them no longer visible to or accessible by normal clients). | [optional] 
**PreCommitValidation** | Pointer to [**EnumpluginPreCommitValidationProp**](EnumpluginPreCommitValidationProp.md) |  | [optional] 
**PostCommitValidation** | Pointer to [**EnumpluginPostCommitValidationProp**](EnumpluginPostCommitValidationProp.md) |  | [optional] 
**Filter** | Pointer to **string** | Specifies the search filter to apply to determine if attribute uniqueness is enforced for the matching entries. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 

## Methods

### NewGloballyUniqueAttributePluginResponse

`func NewGloballyUniqueAttributePluginResponse(schemas []EnumgloballyUniqueAttributePluginSchemaUrn, id string, type_ []string, subtreeView []string, enabled bool, ) *GloballyUniqueAttributePluginResponse`

NewGloballyUniqueAttributePluginResponse instantiates a new GloballyUniqueAttributePluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGloballyUniqueAttributePluginResponseWithDefaults

`func NewGloballyUniqueAttributePluginResponseWithDefaults() *GloballyUniqueAttributePluginResponse`

NewGloballyUniqueAttributePluginResponseWithDefaults instantiates a new GloballyUniqueAttributePluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GloballyUniqueAttributePluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GloballyUniqueAttributePluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GloballyUniqueAttributePluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GloballyUniqueAttributePluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GloballyUniqueAttributePluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GloballyUniqueAttributePluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GloballyUniqueAttributePluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GloballyUniqueAttributePluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *GloballyUniqueAttributePluginResponse) GetSchemas() []EnumgloballyUniqueAttributePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GloballyUniqueAttributePluginResponse) GetSchemasOk() (*[]EnumgloballyUniqueAttributePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GloballyUniqueAttributePluginResponse) SetSchemas(v []EnumgloballyUniqueAttributePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GloballyUniqueAttributePluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GloballyUniqueAttributePluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GloballyUniqueAttributePluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GloballyUniqueAttributePluginResponse) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GloballyUniqueAttributePluginResponse) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GloballyUniqueAttributePluginResponse) SetType(v []string)`

SetType sets Type field to given value.


### GetMultipleAttributeBehavior

`func (o *GloballyUniqueAttributePluginResponse) GetMultipleAttributeBehavior() EnumpluginGloballyUniqueAttributeMultipleAttributeBehaviorProp`

GetMultipleAttributeBehavior returns the MultipleAttributeBehavior field if non-nil, zero value otherwise.

### GetMultipleAttributeBehaviorOk

`func (o *GloballyUniqueAttributePluginResponse) GetMultipleAttributeBehaviorOk() (*EnumpluginGloballyUniqueAttributeMultipleAttributeBehaviorProp, bool)`

GetMultipleAttributeBehaviorOk returns a tuple with the MultipleAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAttributeBehavior

`func (o *GloballyUniqueAttributePluginResponse) SetMultipleAttributeBehavior(v EnumpluginGloballyUniqueAttributeMultipleAttributeBehaviorProp)`

SetMultipleAttributeBehavior sets MultipleAttributeBehavior field to given value.

### HasMultipleAttributeBehavior

`func (o *GloballyUniqueAttributePluginResponse) HasMultipleAttributeBehavior() bool`

HasMultipleAttributeBehavior returns a boolean if a field has been set.

### GetSubtreeView

`func (o *GloballyUniqueAttributePluginResponse) GetSubtreeView() []string`

GetSubtreeView returns the SubtreeView field if non-nil, zero value otherwise.

### GetSubtreeViewOk

`func (o *GloballyUniqueAttributePluginResponse) GetSubtreeViewOk() (*[]string, bool)`

GetSubtreeViewOk returns a tuple with the SubtreeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeView

`func (o *GloballyUniqueAttributePluginResponse) SetSubtreeView(v []string)`

SetSubtreeView sets SubtreeView field to given value.


### GetPreventConflictsWithSoftDeletedEntries

`func (o *GloballyUniqueAttributePluginResponse) GetPreventConflictsWithSoftDeletedEntries() bool`

GetPreventConflictsWithSoftDeletedEntries returns the PreventConflictsWithSoftDeletedEntries field if non-nil, zero value otherwise.

### GetPreventConflictsWithSoftDeletedEntriesOk

`func (o *GloballyUniqueAttributePluginResponse) GetPreventConflictsWithSoftDeletedEntriesOk() (*bool, bool)`

GetPreventConflictsWithSoftDeletedEntriesOk returns a tuple with the PreventConflictsWithSoftDeletedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventConflictsWithSoftDeletedEntries

`func (o *GloballyUniqueAttributePluginResponse) SetPreventConflictsWithSoftDeletedEntries(v bool)`

SetPreventConflictsWithSoftDeletedEntries sets PreventConflictsWithSoftDeletedEntries field to given value.

### HasPreventConflictsWithSoftDeletedEntries

`func (o *GloballyUniqueAttributePluginResponse) HasPreventConflictsWithSoftDeletedEntries() bool`

HasPreventConflictsWithSoftDeletedEntries returns a boolean if a field has been set.

### GetPreCommitValidation

`func (o *GloballyUniqueAttributePluginResponse) GetPreCommitValidation() EnumpluginPreCommitValidationProp`

GetPreCommitValidation returns the PreCommitValidation field if non-nil, zero value otherwise.

### GetPreCommitValidationOk

`func (o *GloballyUniqueAttributePluginResponse) GetPreCommitValidationOk() (*EnumpluginPreCommitValidationProp, bool)`

GetPreCommitValidationOk returns a tuple with the PreCommitValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCommitValidation

`func (o *GloballyUniqueAttributePluginResponse) SetPreCommitValidation(v EnumpluginPreCommitValidationProp)`

SetPreCommitValidation sets PreCommitValidation field to given value.

### HasPreCommitValidation

`func (o *GloballyUniqueAttributePluginResponse) HasPreCommitValidation() bool`

HasPreCommitValidation returns a boolean if a field has been set.

### GetPostCommitValidation

`func (o *GloballyUniqueAttributePluginResponse) GetPostCommitValidation() EnumpluginPostCommitValidationProp`

GetPostCommitValidation returns the PostCommitValidation field if non-nil, zero value otherwise.

### GetPostCommitValidationOk

`func (o *GloballyUniqueAttributePluginResponse) GetPostCommitValidationOk() (*EnumpluginPostCommitValidationProp, bool)`

GetPostCommitValidationOk returns a tuple with the PostCommitValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCommitValidation

`func (o *GloballyUniqueAttributePluginResponse) SetPostCommitValidation(v EnumpluginPostCommitValidationProp)`

SetPostCommitValidation sets PostCommitValidation field to given value.

### HasPostCommitValidation

`func (o *GloballyUniqueAttributePluginResponse) HasPostCommitValidation() bool`

HasPostCommitValidation returns a boolean if a field has been set.

### GetFilter

`func (o *GloballyUniqueAttributePluginResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GloballyUniqueAttributePluginResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GloballyUniqueAttributePluginResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GloballyUniqueAttributePluginResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDescription

`func (o *GloballyUniqueAttributePluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GloballyUniqueAttributePluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GloballyUniqueAttributePluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GloballyUniqueAttributePluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GloballyUniqueAttributePluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GloballyUniqueAttributePluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GloballyUniqueAttributePluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *GloballyUniqueAttributePluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *GloballyUniqueAttributePluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *GloballyUniqueAttributePluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *GloballyUniqueAttributePluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


