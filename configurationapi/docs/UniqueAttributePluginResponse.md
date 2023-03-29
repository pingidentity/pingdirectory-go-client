# UniqueAttributePluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Plugin Root | 
**Schemas** | [**[]EnumuniqueAttributePluginSchemaUrn**](EnumuniqueAttributePluginSchemaUrn.md) |  | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) | Specifies the set of plug-in types for the plug-in, which specifies the times at which the plug-in is invoked. | 
**Type** | **[]string** | Specifies the type of attributes to check for value uniqueness. | 
**MultipleAttributeBehavior** | Pointer to [**EnumpluginMultipleAttributeBehaviorProp**](EnumpluginMultipleAttributeBehaviorProp.md) |  | [optional] 
**BaseDN** | Pointer to **[]string** | Specifies a base DN within which the attribute must be unique. | [optional] 
**PreventConflictsWithSoftDeletedEntries** | Pointer to **bool** | Indicates whether this Unique Attribute Plugin should reject a change that would result in one or more conflicts, even if those conflicts only exist in soft-deleted entries. | [optional] 
**Filter** | Pointer to **string** | Specifies the search filter to apply to determine if attribute uniqueness is enforced for the matching entries. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewUniqueAttributePluginResponse

`func NewUniqueAttributePluginResponse(id string, schemas []EnumuniqueAttributePluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, type_ []string, enabled bool, ) *UniqueAttributePluginResponse`

NewUniqueAttributePluginResponse instantiates a new UniqueAttributePluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniqueAttributePluginResponseWithDefaults

`func NewUniqueAttributePluginResponseWithDefaults() *UniqueAttributePluginResponse`

NewUniqueAttributePluginResponseWithDefaults instantiates a new UniqueAttributePluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UniqueAttributePluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UniqueAttributePluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UniqueAttributePluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *UniqueAttributePluginResponse) GetSchemas() []EnumuniqueAttributePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UniqueAttributePluginResponse) GetSchemasOk() (*[]EnumuniqueAttributePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UniqueAttributePluginResponse) SetSchemas(v []EnumuniqueAttributePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *UniqueAttributePluginResponse) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *UniqueAttributePluginResponse) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *UniqueAttributePluginResponse) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetType

`func (o *UniqueAttributePluginResponse) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UniqueAttributePluginResponse) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UniqueAttributePluginResponse) SetType(v []string)`

SetType sets Type field to given value.


### GetMultipleAttributeBehavior

`func (o *UniqueAttributePluginResponse) GetMultipleAttributeBehavior() EnumpluginMultipleAttributeBehaviorProp`

GetMultipleAttributeBehavior returns the MultipleAttributeBehavior field if non-nil, zero value otherwise.

### GetMultipleAttributeBehaviorOk

`func (o *UniqueAttributePluginResponse) GetMultipleAttributeBehaviorOk() (*EnumpluginMultipleAttributeBehaviorProp, bool)`

GetMultipleAttributeBehaviorOk returns a tuple with the MultipleAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAttributeBehavior

`func (o *UniqueAttributePluginResponse) SetMultipleAttributeBehavior(v EnumpluginMultipleAttributeBehaviorProp)`

SetMultipleAttributeBehavior sets MultipleAttributeBehavior field to given value.

### HasMultipleAttributeBehavior

`func (o *UniqueAttributePluginResponse) HasMultipleAttributeBehavior() bool`

HasMultipleAttributeBehavior returns a boolean if a field has been set.

### GetBaseDN

`func (o *UniqueAttributePluginResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *UniqueAttributePluginResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *UniqueAttributePluginResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *UniqueAttributePluginResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetPreventConflictsWithSoftDeletedEntries

`func (o *UniqueAttributePluginResponse) GetPreventConflictsWithSoftDeletedEntries() bool`

GetPreventConflictsWithSoftDeletedEntries returns the PreventConflictsWithSoftDeletedEntries field if non-nil, zero value otherwise.

### GetPreventConflictsWithSoftDeletedEntriesOk

`func (o *UniqueAttributePluginResponse) GetPreventConflictsWithSoftDeletedEntriesOk() (*bool, bool)`

GetPreventConflictsWithSoftDeletedEntriesOk returns a tuple with the PreventConflictsWithSoftDeletedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventConflictsWithSoftDeletedEntries

`func (o *UniqueAttributePluginResponse) SetPreventConflictsWithSoftDeletedEntries(v bool)`

SetPreventConflictsWithSoftDeletedEntries sets PreventConflictsWithSoftDeletedEntries field to given value.

### HasPreventConflictsWithSoftDeletedEntries

`func (o *UniqueAttributePluginResponse) HasPreventConflictsWithSoftDeletedEntries() bool`

HasPreventConflictsWithSoftDeletedEntries returns a boolean if a field has been set.

### GetFilter

`func (o *UniqueAttributePluginResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UniqueAttributePluginResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UniqueAttributePluginResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UniqueAttributePluginResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetDescription

`func (o *UniqueAttributePluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UniqueAttributePluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UniqueAttributePluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UniqueAttributePluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UniqueAttributePluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UniqueAttributePluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UniqueAttributePluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *UniqueAttributePluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *UniqueAttributePluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *UniqueAttributePluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *UniqueAttributePluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetMeta

`func (o *UniqueAttributePluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UniqueAttributePluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UniqueAttributePluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UniqueAttributePluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *UniqueAttributePluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *UniqueAttributePluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *UniqueAttributePluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *UniqueAttributePluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


