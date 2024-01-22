# AddSoftDeletePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumsoftDeletePolicySchemaUrn**](EnumsoftDeletePolicySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Soft Delete Policy | [optional] 
**AutoSoftDeleteConnectionCriteria** | Pointer to **string** | Connection criteria used to automatically identify a delete operation for processing as a soft delete request. | [optional] 
**AutoSoftDeleteRequestCriteria** | Pointer to **string** | Request criteria used to automatically identify a delete operation for processing as a soft delete request. | [optional] 
**SoftDeleteRetentionTime** | Pointer to **string** | Specifies the maximum length of time that soft delete entries are retained before they are eligible to purged automatically. | [optional] 
**SoftDeleteRetainNumberOfEntries** | Pointer to **int64** | Specifies the number of soft deleted entries to retain before the oldest entries are purged. | [optional] 
**PolicyName** | **string** | Name of the new Soft Delete Policy | 

## Methods

### NewAddSoftDeletePolicyRequest

`func NewAddSoftDeletePolicyRequest(policyName string, ) *AddSoftDeletePolicyRequest`

NewAddSoftDeletePolicyRequest instantiates a new AddSoftDeletePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSoftDeletePolicyRequestWithDefaults

`func NewAddSoftDeletePolicyRequestWithDefaults() *AddSoftDeletePolicyRequest`

NewAddSoftDeletePolicyRequestWithDefaults instantiates a new AddSoftDeletePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSoftDeletePolicyRequest) GetSchemas() []EnumsoftDeletePolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSoftDeletePolicyRequest) GetSchemasOk() (*[]EnumsoftDeletePolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSoftDeletePolicyRequest) SetSchemas(v []EnumsoftDeletePolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddSoftDeletePolicyRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddSoftDeletePolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSoftDeletePolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSoftDeletePolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSoftDeletePolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutoSoftDeleteConnectionCriteria

`func (o *AddSoftDeletePolicyRequest) GetAutoSoftDeleteConnectionCriteria() string`

GetAutoSoftDeleteConnectionCriteria returns the AutoSoftDeleteConnectionCriteria field if non-nil, zero value otherwise.

### GetAutoSoftDeleteConnectionCriteriaOk

`func (o *AddSoftDeletePolicyRequest) GetAutoSoftDeleteConnectionCriteriaOk() (*string, bool)`

GetAutoSoftDeleteConnectionCriteriaOk returns a tuple with the AutoSoftDeleteConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSoftDeleteConnectionCriteria

`func (o *AddSoftDeletePolicyRequest) SetAutoSoftDeleteConnectionCriteria(v string)`

SetAutoSoftDeleteConnectionCriteria sets AutoSoftDeleteConnectionCriteria field to given value.

### HasAutoSoftDeleteConnectionCriteria

`func (o *AddSoftDeletePolicyRequest) HasAutoSoftDeleteConnectionCriteria() bool`

HasAutoSoftDeleteConnectionCriteria returns a boolean if a field has been set.

### GetAutoSoftDeleteRequestCriteria

`func (o *AddSoftDeletePolicyRequest) GetAutoSoftDeleteRequestCriteria() string`

GetAutoSoftDeleteRequestCriteria returns the AutoSoftDeleteRequestCriteria field if non-nil, zero value otherwise.

### GetAutoSoftDeleteRequestCriteriaOk

`func (o *AddSoftDeletePolicyRequest) GetAutoSoftDeleteRequestCriteriaOk() (*string, bool)`

GetAutoSoftDeleteRequestCriteriaOk returns a tuple with the AutoSoftDeleteRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSoftDeleteRequestCriteria

`func (o *AddSoftDeletePolicyRequest) SetAutoSoftDeleteRequestCriteria(v string)`

SetAutoSoftDeleteRequestCriteria sets AutoSoftDeleteRequestCriteria field to given value.

### HasAutoSoftDeleteRequestCriteria

`func (o *AddSoftDeletePolicyRequest) HasAutoSoftDeleteRequestCriteria() bool`

HasAutoSoftDeleteRequestCriteria returns a boolean if a field has been set.

### GetSoftDeleteRetentionTime

`func (o *AddSoftDeletePolicyRequest) GetSoftDeleteRetentionTime() string`

GetSoftDeleteRetentionTime returns the SoftDeleteRetentionTime field if non-nil, zero value otherwise.

### GetSoftDeleteRetentionTimeOk

`func (o *AddSoftDeletePolicyRequest) GetSoftDeleteRetentionTimeOk() (*string, bool)`

GetSoftDeleteRetentionTimeOk returns a tuple with the SoftDeleteRetentionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteRetentionTime

`func (o *AddSoftDeletePolicyRequest) SetSoftDeleteRetentionTime(v string)`

SetSoftDeleteRetentionTime sets SoftDeleteRetentionTime field to given value.

### HasSoftDeleteRetentionTime

`func (o *AddSoftDeletePolicyRequest) HasSoftDeleteRetentionTime() bool`

HasSoftDeleteRetentionTime returns a boolean if a field has been set.

### GetSoftDeleteRetainNumberOfEntries

`func (o *AddSoftDeletePolicyRequest) GetSoftDeleteRetainNumberOfEntries() int64`

GetSoftDeleteRetainNumberOfEntries returns the SoftDeleteRetainNumberOfEntries field if non-nil, zero value otherwise.

### GetSoftDeleteRetainNumberOfEntriesOk

`func (o *AddSoftDeletePolicyRequest) GetSoftDeleteRetainNumberOfEntriesOk() (*int64, bool)`

GetSoftDeleteRetainNumberOfEntriesOk returns a tuple with the SoftDeleteRetainNumberOfEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeleteRetainNumberOfEntries

`func (o *AddSoftDeletePolicyRequest) SetSoftDeleteRetainNumberOfEntries(v int64)`

SetSoftDeleteRetainNumberOfEntries sets SoftDeleteRetainNumberOfEntries field to given value.

### HasSoftDeleteRetainNumberOfEntries

`func (o *AddSoftDeletePolicyRequest) HasSoftDeleteRetainNumberOfEntries() bool`

HasSoftDeleteRetainNumberOfEntries returns a boolean if a field has been set.

### GetPolicyName

`func (o *AddSoftDeletePolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddSoftDeletePolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddSoftDeletePolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


