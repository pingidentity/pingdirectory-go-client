# SoftDeletePolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Soft Delete Policy | 
**Schemas** | Pointer to [**[]EnumsoftDeletePolicySchemaUrn**](EnumsoftDeletePolicySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Soft Delete Policy | [optional] 
**AutoSoftDeleteConnectionCriteria** | Pointer to **string** | Connection criteria used to automatically identify a delete operation for processing as a soft delete request. | [optional] 
**AutoSoftDeleteRequestCriteria** | Pointer to **string** | Request criteria used to automatically identify a delete operation for processing as a soft delete request. | [optional] 
**SoftDeleteRetentionTime** | Pointer to **string** | Specifies the maximum length of time that soft delete entries are retained before they are eligible to purged automatically. | [optional] 
**SoftDeleteRetainNumberOfEntries** | Pointer to **int32** | Specifies the number of soft deleted entries to retain before the oldest entries are purged. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewSoftDeletePolicyResponse

`func NewSoftDeletePolicyResponse(id string, ) *SoftDeletePolicyResponse`

NewSoftDeletePolicyResponse instantiates a new SoftDeletePolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftDeletePolicyResponseWithDefaults

`func NewSoftDeletePolicyResponseWithDefaults() *SoftDeletePolicyResponse`

NewSoftDeletePolicyResponseWithDefaults instantiates a new SoftDeletePolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SoftDeletePolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SoftDeletePolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SoftDeletePolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SoftDeletePolicyResponse) GetSchemas() []EnumsoftDeletePolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SoftDeletePolicyResponse) GetSchemasOk() (*[]EnumsoftDeletePolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SoftDeletePolicyResponse) SetSchemas(v []EnumsoftDeletePolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *SoftDeletePolicyResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *SoftDeletePolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoftDeletePolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoftDeletePolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SoftDeletePolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutoSoftDeleteConnectionCriteria

`func (o *SoftDeletePolicyResponse) GetAutoSoftDeleteConnectionCriteria() string`

GetAutoSoftDeleteConnectionCriteria returns the AutoSoftDeleteConnectionCriteria field if non-nil, zero value otherwise.

### GetAutoSoftDeleteConnectionCriteriaOk

`func (o *SoftDeletePolicyResponse) GetAutoSoftDeleteConnectionCriteriaOk() (*string, bool)`

GetAutoSoftDeleteConnectionCriteriaOk returns a tuple with the AutoSoftDeleteConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSoftDeleteConnectionCriteria

`func (o *SoftDeletePolicyResponse) SetAutoSoftDeleteConnectionCriteria(v string)`

SetAutoSoftDeleteConnectionCriteria sets AutoSoftDeleteConnectionCriteria field to given value.

### HasAutoSoftDeleteConnectionCriteria

`func (o *SoftDeletePolicyResponse) HasAutoSoftDeleteConnectionCriteria() bool`

HasAutoSoftDeleteConnectionCriteria returns a boolean if a field has been set.

### GetAutoSoftDeleteRequestCriteria

`func (o *SoftDeletePolicyResponse) GetAutoSoftDeleteRequestCriteria() string`

GetAutoSoftDeleteRequestCriteria returns the AutoSoftDeleteRequestCriteria field if non-nil, zero value otherwise.

### GetAutoSoftDeleteRequestCriteriaOk

`func (o *SoftDeletePolicyResponse) GetAutoSoftDeleteRequestCriteriaOk() (*string, bool)`

GetAutoSoftDeleteRequestCriteriaOk returns a tuple with the AutoSoftDeleteRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSoftDeleteRequestCriteria

`func (o *SoftDeletePolicyResponse) SetAutoSoftDeleteRequestCriteria(v string)`

SetAutoSoftDeleteRequestCriteria sets AutoSoftDeleteRequestCriteria field to given value.

### HasAutoSoftDeleteRequestCriteria

`func (o *SoftDeletePolicyResponse) HasAutoSoftDeleteRequestCriteria() bool`

HasAutoSoftDeleteRequestCriteria returns a boolean if a field has been set.

### GetSoftDeleteRetentionTime

`func (o *SoftDeletePolicyResponse) GetSoftDeleteRetentionTime() string`

GetSoftDeleteRetentionTime returns the SoftDeleteRetentionTime field if non-nil, zero value otherwise.

### GetSoftDeleteRetentionTimeOk

`func (o *SoftDeletePolicyResponse) GetSoftDeleteRetentionTimeOk() (*string, bool)`

GetSoftDeleteRetentionTimeOk returns a tuple with the SoftDeleteRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteRetentionTime

`func (o *SoftDeletePolicyResponse) SetSoftDeleteRetentionTime(v string)`

SetSoftDeleteRetentionTime sets SoftDeleteRetentionTime field to given value.

### HasSoftDeleteRetentionTime

`func (o *SoftDeletePolicyResponse) HasSoftDeleteRetentionTime() bool`

HasSoftDeleteRetentionTime returns a boolean if a field has been set.

### GetSoftDeleteRetainNumberOfEntries

`func (o *SoftDeletePolicyResponse) GetSoftDeleteRetainNumberOfEntries() int32`

GetSoftDeleteRetainNumberOfEntries returns the SoftDeleteRetainNumberOfEntries field if non-nil, zero value otherwise.

### GetSoftDeleteRetainNumberOfEntriesOk

`func (o *SoftDeletePolicyResponse) GetSoftDeleteRetainNumberOfEntriesOk() (*int32, bool)`

GetSoftDeleteRetainNumberOfEntriesOk returns a tuple with the SoftDeleteRetainNumberOfEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteRetainNumberOfEntries

`func (o *SoftDeletePolicyResponse) SetSoftDeleteRetainNumberOfEntries(v int32)`

SetSoftDeleteRetainNumberOfEntries sets SoftDeleteRetainNumberOfEntries field to given value.

### HasSoftDeleteRetainNumberOfEntries

`func (o *SoftDeletePolicyResponse) HasSoftDeleteRetainNumberOfEntries() bool`

HasSoftDeleteRetainNumberOfEntries returns a boolean if a field has been set.

### GetMeta

`func (o *SoftDeletePolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SoftDeletePolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SoftDeletePolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SoftDeletePolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


