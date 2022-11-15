# SizeLimitLogRetentionPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Retention Policy | 
**Schemas** | [**[]EnumsizeLimitLogRetentionPolicySchemaUrn**](EnumsizeLimitLogRetentionPolicySchemaUrn.md) |  | 
**DiskSpaceUsed** | **string** | Specifies the maximum total disk space used by the log files. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 

## Methods

### NewSizeLimitLogRetentionPolicyResponse

`func NewSizeLimitLogRetentionPolicyResponse(id string, schemas []EnumsizeLimitLogRetentionPolicySchemaUrn, diskSpaceUsed string, ) *SizeLimitLogRetentionPolicyResponse`

NewSizeLimitLogRetentionPolicyResponse instantiates a new SizeLimitLogRetentionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeLimitLogRetentionPolicyResponseWithDefaults

`func NewSizeLimitLogRetentionPolicyResponseWithDefaults() *SizeLimitLogRetentionPolicyResponse`

NewSizeLimitLogRetentionPolicyResponseWithDefaults instantiates a new SizeLimitLogRetentionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SizeLimitLogRetentionPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SizeLimitLogRetentionPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SizeLimitLogRetentionPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SizeLimitLogRetentionPolicyResponse) GetSchemas() []EnumsizeLimitLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SizeLimitLogRetentionPolicyResponse) GetSchemasOk() (*[]EnumsizeLimitLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SizeLimitLogRetentionPolicyResponse) SetSchemas(v []EnumsizeLimitLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDiskSpaceUsed

`func (o *SizeLimitLogRetentionPolicyResponse) GetDiskSpaceUsed() string`

GetDiskSpaceUsed returns the DiskSpaceUsed field if non-nil, zero value otherwise.

### GetDiskSpaceUsedOk

`func (o *SizeLimitLogRetentionPolicyResponse) GetDiskSpaceUsedOk() (*string, bool)`

GetDiskSpaceUsedOk returns a tuple with the DiskSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsed

`func (o *SizeLimitLogRetentionPolicyResponse) SetDiskSpaceUsed(v string)`

SetDiskSpaceUsed sets DiskSpaceUsed field to given value.


### GetDescription

`func (o *SizeLimitLogRetentionPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SizeLimitLogRetentionPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SizeLimitLogRetentionPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SizeLimitLogRetentionPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


