# SizeLimitLogRetentionPolicyShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsizeLimitLogRetentionPolicySchemaUrn**](EnumsizeLimitLogRetentionPolicySchemaUrn.md) |  | 
**DiskSpaceUsed** | **string** | Specifies the maximum total disk space used by the log files. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 

## Methods

### NewSizeLimitLogRetentionPolicyShared

`func NewSizeLimitLogRetentionPolicyShared(schemas []EnumsizeLimitLogRetentionPolicySchemaUrn, diskSpaceUsed string, ) *SizeLimitLogRetentionPolicyShared`

NewSizeLimitLogRetentionPolicyShared instantiates a new SizeLimitLogRetentionPolicyShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeLimitLogRetentionPolicySharedWithDefaults

`func NewSizeLimitLogRetentionPolicySharedWithDefaults() *SizeLimitLogRetentionPolicyShared`

NewSizeLimitLogRetentionPolicySharedWithDefaults instantiates a new SizeLimitLogRetentionPolicyShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *SizeLimitLogRetentionPolicyShared) GetSchemas() []EnumsizeLimitLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SizeLimitLogRetentionPolicyShared) GetSchemasOk() (*[]EnumsizeLimitLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SizeLimitLogRetentionPolicyShared) SetSchemas(v []EnumsizeLimitLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDiskSpaceUsed

`func (o *SizeLimitLogRetentionPolicyShared) GetDiskSpaceUsed() string`

GetDiskSpaceUsed returns the DiskSpaceUsed field if non-nil, zero value otherwise.

### GetDiskSpaceUsedOk

`func (o *SizeLimitLogRetentionPolicyShared) GetDiskSpaceUsedOk() (*string, bool)`

GetDiskSpaceUsedOk returns a tuple with the DiskSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsed

`func (o *SizeLimitLogRetentionPolicyShared) SetDiskSpaceUsed(v string)`

SetDiskSpaceUsed sets DiskSpaceUsed field to given value.


### GetDescription

`func (o *SizeLimitLogRetentionPolicyShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SizeLimitLogRetentionPolicyShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SizeLimitLogRetentionPolicyShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SizeLimitLogRetentionPolicyShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


