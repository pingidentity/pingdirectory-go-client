# LastAccessTimePluginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumlastAccessTimePluginSchemaUrn**](EnumlastAccessTimePluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin Root | 
**MaxUpdateFrequency** | Pointer to **string** | Specifies the maximum frequency with which last access time values should be written for an entry. This may help limit the rate of internal write operations processed in the server. | [optional] 
**OperationType** | Pointer to [**[]EnumpluginOperationTypeProp**](EnumpluginOperationTypeProp.md) |  | [optional] 
**InvokeForFailedBinds** | Pointer to **bool** | Indicates whether to update the last access time for an entry targeted by a bind operation if the bind is unsuccessful. | [optional] 
**MaxSearchResultEntriesToUpdate** | Pointer to **int32** | Specifies the maximum number of entries that should be updated in a search operation. Only search result entries actually returned to the client may have their last access time updated, but because a single search operation may return a very large number of entries, the plugin will only update entries if no more than a specified number of entries are updated. | [optional] 
**RequestCriteria** | Pointer to **string** | Specifies a set of request criteria that may be used to indicate whether to apply access time updates for the associated operation. | [optional] 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewLastAccessTimePluginResponse

`func NewLastAccessTimePluginResponse(schemas []EnumlastAccessTimePluginSchemaUrn, id string, enabled bool, ) *LastAccessTimePluginResponse`

NewLastAccessTimePluginResponse instantiates a new LastAccessTimePluginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastAccessTimePluginResponseWithDefaults

`func NewLastAccessTimePluginResponseWithDefaults() *LastAccessTimePluginResponse`

NewLastAccessTimePluginResponseWithDefaults instantiates a new LastAccessTimePluginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LastAccessTimePluginResponse) GetSchemas() []EnumlastAccessTimePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LastAccessTimePluginResponse) GetSchemasOk() (*[]EnumlastAccessTimePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LastAccessTimePluginResponse) SetSchemas(v []EnumlastAccessTimePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *LastAccessTimePluginResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LastAccessTimePluginResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LastAccessTimePluginResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMaxUpdateFrequency

`func (o *LastAccessTimePluginResponse) GetMaxUpdateFrequency() string`

GetMaxUpdateFrequency returns the MaxUpdateFrequency field if non-nil, zero value otherwise.

### GetMaxUpdateFrequencyOk

`func (o *LastAccessTimePluginResponse) GetMaxUpdateFrequencyOk() (*string, bool)`

GetMaxUpdateFrequencyOk returns a tuple with the MaxUpdateFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdateFrequency

`func (o *LastAccessTimePluginResponse) SetMaxUpdateFrequency(v string)`

SetMaxUpdateFrequency sets MaxUpdateFrequency field to given value.

### HasMaxUpdateFrequency

`func (o *LastAccessTimePluginResponse) HasMaxUpdateFrequency() bool`

HasMaxUpdateFrequency returns a boolean if a field has been set.

### GetOperationType

`func (o *LastAccessTimePluginResponse) GetOperationType() []EnumpluginOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *LastAccessTimePluginResponse) GetOperationTypeOk() (*[]EnumpluginOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *LastAccessTimePluginResponse) SetOperationType(v []EnumpluginOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *LastAccessTimePluginResponse) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetInvokeForFailedBinds

`func (o *LastAccessTimePluginResponse) GetInvokeForFailedBinds() bool`

GetInvokeForFailedBinds returns the InvokeForFailedBinds field if non-nil, zero value otherwise.

### GetInvokeForFailedBindsOk

`func (o *LastAccessTimePluginResponse) GetInvokeForFailedBindsOk() (*bool, bool)`

GetInvokeForFailedBindsOk returns a tuple with the InvokeForFailedBinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForFailedBinds

`func (o *LastAccessTimePluginResponse) SetInvokeForFailedBinds(v bool)`

SetInvokeForFailedBinds sets InvokeForFailedBinds field to given value.

### HasInvokeForFailedBinds

`func (o *LastAccessTimePluginResponse) HasInvokeForFailedBinds() bool`

HasInvokeForFailedBinds returns a boolean if a field has been set.

### GetMaxSearchResultEntriesToUpdate

`func (o *LastAccessTimePluginResponse) GetMaxSearchResultEntriesToUpdate() int32`

GetMaxSearchResultEntriesToUpdate returns the MaxSearchResultEntriesToUpdate field if non-nil, zero value otherwise.

### GetMaxSearchResultEntriesToUpdateOk

`func (o *LastAccessTimePluginResponse) GetMaxSearchResultEntriesToUpdateOk() (*int32, bool)`

GetMaxSearchResultEntriesToUpdateOk returns a tuple with the MaxSearchResultEntriesToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSearchResultEntriesToUpdate

`func (o *LastAccessTimePluginResponse) SetMaxSearchResultEntriesToUpdate(v int32)`

SetMaxSearchResultEntriesToUpdate sets MaxSearchResultEntriesToUpdate field to given value.

### HasMaxSearchResultEntriesToUpdate

`func (o *LastAccessTimePluginResponse) HasMaxSearchResultEntriesToUpdate() bool`

HasMaxSearchResultEntriesToUpdate returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *LastAccessTimePluginResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *LastAccessTimePluginResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *LastAccessTimePluginResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *LastAccessTimePluginResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetInvokeForInternalOperations

`func (o *LastAccessTimePluginResponse) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *LastAccessTimePluginResponse) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *LastAccessTimePluginResponse) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *LastAccessTimePluginResponse) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetDescription

`func (o *LastAccessTimePluginResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LastAccessTimePluginResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LastAccessTimePluginResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LastAccessTimePluginResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *LastAccessTimePluginResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LastAccessTimePluginResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LastAccessTimePluginResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *LastAccessTimePluginResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LastAccessTimePluginResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LastAccessTimePluginResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LastAccessTimePluginResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LastAccessTimePluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LastAccessTimePluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LastAccessTimePluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LastAccessTimePluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


