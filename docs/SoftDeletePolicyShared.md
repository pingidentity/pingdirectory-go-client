# SoftDeletePolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumsoftDeletePolicySchemaUrn**](EnumsoftDeletePolicySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Soft Delete Policy | [optional] 
**AutoSoftDeleteConnectionCriteria** | Pointer to **string** | Connection criteria used to automatically identify a delete operation for processing as a soft delete request. | [optional] 
**AutoSoftDeleteRequestCriteria** | Pointer to **string** | Request criteria used to automatically identify a delete operation for processing as a soft delete request. | [optional] 
**SoftDeleteRetentionTime** | Pointer to **string** | Specifies the maximum length of time that soft delete entries are retained before they are eligible to purged automatically. | [optional] 
**SoftDeleteRetainNumberOfEntries** | Pointer to **int32** | Specifies the number of soft deleted entries to retain before the oldest entries are purged. | [optional] 

## Methods

### NewSoftDeletePolicyShared

`func NewSoftDeletePolicyShared() *SoftDeletePolicyShared`

NewSoftDeletePolicyShared instantiates a new SoftDeletePolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftDeletePolicySharedWithDefaults

`func NewSoftDeletePolicySharedWithDefaults() *SoftDeletePolicyShared`

NewSoftDeletePolicySharedWithDefaults instantiates a new SoftDeletePolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SoftDeletePolicyShared) GetSchemas() []EnumsoftDeletePolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SoftDeletePolicyShared) GetSchemasOk() (*[]EnumsoftDeletePolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SoftDeletePolicyShared) SetSchemas(v []EnumsoftDeletePolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *SoftDeletePolicyShared) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *SoftDeletePolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SoftDeletePolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SoftDeletePolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SoftDeletePolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutoSoftDeleteConnectionCriteria

`func (o *SoftDeletePolicyShared) GetAutoSoftDeleteConnectionCriteria() string`

GetAutoSoftDeleteConnectionCriteria returns the AutoSoftDeleteConnectionCriteria field if non-nil, zero value otherwise.

### GetAutoSoftDeleteConnectionCriteriaOk

`func (o *SoftDeletePolicyShared) GetAutoSoftDeleteConnectionCriteriaOk() (*string, bool)`

GetAutoSoftDeleteConnectionCriteriaOk returns a tuple with the AutoSoftDeleteConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSoftDeleteConnectionCriteria

`func (o *SoftDeletePolicyShared) SetAutoSoftDeleteConnectionCriteria(v string)`

SetAutoSoftDeleteConnectionCriteria sets AutoSoftDeleteConnectionCriteria field to given value.

### HasAutoSoftDeleteConnectionCriteria

`func (o *SoftDeletePolicyShared) HasAutoSoftDeleteConnectionCriteria() bool`

HasAutoSoftDeleteConnectionCriteria returns a boolean if a field has been set.

### GetAutoSoftDeleteRequestCriteria

`func (o *SoftDeletePolicyShared) GetAutoSoftDeleteRequestCriteria() string`

GetAutoSoftDeleteRequestCriteria returns the AutoSoftDeleteRequestCriteria field if non-nil, zero value otherwise.

### GetAutoSoftDeleteRequestCriteriaOk

`func (o *SoftDeletePolicyShared) GetAutoSoftDeleteRequestCriteriaOk() (*string, bool)`

GetAutoSoftDeleteRequestCriteriaOk returns a tuple with the AutoSoftDeleteRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSoftDeleteRequestCriteria

`func (o *SoftDeletePolicyShared) SetAutoSoftDeleteRequestCriteria(v string)`

SetAutoSoftDeleteRequestCriteria sets AutoSoftDeleteRequestCriteria field to given value.

### HasAutoSoftDeleteRequestCriteria

`func (o *SoftDeletePolicyShared) HasAutoSoftDeleteRequestCriteria() bool`

HasAutoSoftDeleteRequestCriteria returns a boolean if a field has been set.

### GetSoftDeleteRetentionTime

`func (o *SoftDeletePolicyShared) GetSoftDeleteRetentionTime() string`

GetSoftDeleteRetentionTime returns the SoftDeleteRetentionTime field if non-nil, zero value otherwise.

### GetSoftDeleteRetentionTimeOk

`func (o *SoftDeletePolicyShared) GetSoftDeleteRetentionTimeOk() (*string, bool)`

GetSoftDeleteRetentionTimeOk returns a tuple with the SoftDeleteRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteRetentionTime

`func (o *SoftDeletePolicyShared) SetSoftDeleteRetentionTime(v string)`

SetSoftDeleteRetentionTime sets SoftDeleteRetentionTime field to given value.

### HasSoftDeleteRetentionTime

`func (o *SoftDeletePolicyShared) HasSoftDeleteRetentionTime() bool`

HasSoftDeleteRetentionTime returns a boolean if a field has been set.

### GetSoftDeleteRetainNumberOfEntries

`func (o *SoftDeletePolicyShared) GetSoftDeleteRetainNumberOfEntries() int32`

GetSoftDeleteRetainNumberOfEntries returns the SoftDeleteRetainNumberOfEntries field if non-nil, zero value otherwise.

### GetSoftDeleteRetainNumberOfEntriesOk

`func (o *SoftDeletePolicyShared) GetSoftDeleteRetainNumberOfEntriesOk() (*int32, bool)`

GetSoftDeleteRetainNumberOfEntriesOk returns a tuple with the SoftDeleteRetainNumberOfEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteRetainNumberOfEntries

`func (o *SoftDeletePolicyShared) SetSoftDeleteRetainNumberOfEntries(v int32)`

SetSoftDeleteRetainNumberOfEntries sets SoftDeleteRetainNumberOfEntries field to given value.

### HasSoftDeleteRetainNumberOfEntries

`func (o *SoftDeletePolicyShared) HasSoftDeleteRetainNumberOfEntries() bool`

HasSoftDeleteRetainNumberOfEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


