# AddSizeLimitLogRetentionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new Log Retention Policy | 
**Schemas** | [**[]EnumsizeLimitLogRetentionPolicySchemaUrn**](EnumsizeLimitLogRetentionPolicySchemaUrn.md) |  | 
**DiskSpaceUsed** | **string** | Specifies the maximum total disk space used by the log files. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 

## Methods

### NewAddSizeLimitLogRetentionPolicyRequest

`func NewAddSizeLimitLogRetentionPolicyRequest(policyName string, schemas []EnumsizeLimitLogRetentionPolicySchemaUrn, diskSpaceUsed string, ) *AddSizeLimitLogRetentionPolicyRequest`

NewAddSizeLimitLogRetentionPolicyRequest instantiates a new AddSizeLimitLogRetentionPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSizeLimitLogRetentionPolicyRequestWithDefaults

`func NewAddSizeLimitLogRetentionPolicyRequestWithDefaults() *AddSizeLimitLogRetentionPolicyRequest`

NewAddSizeLimitLogRetentionPolicyRequestWithDefaults instantiates a new AddSizeLimitLogRetentionPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddSizeLimitLogRetentionPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddSizeLimitLogRetentionPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddSizeLimitLogRetentionPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddSizeLimitLogRetentionPolicyRequest) GetSchemas() []EnumsizeLimitLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSizeLimitLogRetentionPolicyRequest) GetSchemasOk() (*[]EnumsizeLimitLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSizeLimitLogRetentionPolicyRequest) SetSchemas(v []EnumsizeLimitLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDiskSpaceUsed

`func (o *AddSizeLimitLogRetentionPolicyRequest) GetDiskSpaceUsed() string`

GetDiskSpaceUsed returns the DiskSpaceUsed field if non-nil, zero value otherwise.

### GetDiskSpaceUsedOk

`func (o *AddSizeLimitLogRetentionPolicyRequest) GetDiskSpaceUsedOk() (*string, bool)`

GetDiskSpaceUsedOk returns a tuple with the DiskSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsed

`func (o *AddSizeLimitLogRetentionPolicyRequest) SetDiskSpaceUsed(v string)`

SetDiskSpaceUsed sets DiskSpaceUsed field to given value.


### GetDescription

`func (o *AddSizeLimitLogRetentionPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSizeLimitLogRetentionPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSizeLimitLogRetentionPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSizeLimitLogRetentionPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


