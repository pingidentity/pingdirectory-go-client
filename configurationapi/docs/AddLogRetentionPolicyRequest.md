# AddLogRetentionPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** | Name of the new Log Retention Policy | 
**Schemas** | [**[]EnumsizeLimitLogRetentionPolicySchemaUrn**](EnumsizeLimitLogRetentionPolicySchemaUrn.md) |  | 
**RetainDuration** | **string** | Specifies the desired minimum length of time that each log file should be retained. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 
**NumberOfFiles** | **int32** | Specifies the number of archived log files to retain before the oldest ones are cleaned. | 
**FreeDiskSpace** | **string** | Specifies the minimum amount of free disk space that should be available on the file system on which the archived log files are stored. | 
**DiskSpaceUsed** | **string** | Specifies the maximum total disk space used by the log files. | 

## Methods

### NewAddLogRetentionPolicyRequest

`func NewAddLogRetentionPolicyRequest(policyName string, schemas []EnumsizeLimitLogRetentionPolicySchemaUrn, retainDuration string, numberOfFiles int32, freeDiskSpace string, diskSpaceUsed string, ) *AddLogRetentionPolicyRequest`

NewAddLogRetentionPolicyRequest instantiates a new AddLogRetentionPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogRetentionPolicyRequestWithDefaults

`func NewAddLogRetentionPolicyRequestWithDefaults() *AddLogRetentionPolicyRequest`

NewAddLogRetentionPolicyRequestWithDefaults instantiates a new AddLogRetentionPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *AddLogRetentionPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddLogRetentionPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddLogRetentionPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetSchemas

`func (o *AddLogRetentionPolicyRequest) GetSchemas() []EnumsizeLimitLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogRetentionPolicyRequest) GetSchemasOk() (*[]EnumsizeLimitLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogRetentionPolicyRequest) SetSchemas(v []EnumsizeLimitLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRetainDuration

`func (o *AddLogRetentionPolicyRequest) GetRetainDuration() string`

GetRetainDuration returns the RetainDuration field if non-nil, zero value otherwise.

### GetRetainDurationOk

`func (o *AddLogRetentionPolicyRequest) GetRetainDurationOk() (*string, bool)`

GetRetainDurationOk returns a tuple with the RetainDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainDuration

`func (o *AddLogRetentionPolicyRequest) SetRetainDuration(v string)`

SetRetainDuration sets RetainDuration field to given value.


### GetDescription

`func (o *AddLogRetentionPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogRetentionPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogRetentionPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogRetentionPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNumberOfFiles

`func (o *AddLogRetentionPolicyRequest) GetNumberOfFiles() int32`

GetNumberOfFiles returns the NumberOfFiles field if non-nil, zero value otherwise.

### GetNumberOfFilesOk

`func (o *AddLogRetentionPolicyRequest) GetNumberOfFilesOk() (*int32, bool)`

GetNumberOfFilesOk returns a tuple with the NumberOfFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFiles

`func (o *AddLogRetentionPolicyRequest) SetNumberOfFiles(v int32)`

SetNumberOfFiles sets NumberOfFiles field to given value.


### GetFreeDiskSpace

`func (o *AddLogRetentionPolicyRequest) GetFreeDiskSpace() string`

GetFreeDiskSpace returns the FreeDiskSpace field if non-nil, zero value otherwise.

### GetFreeDiskSpaceOk

`func (o *AddLogRetentionPolicyRequest) GetFreeDiskSpaceOk() (*string, bool)`

GetFreeDiskSpaceOk returns a tuple with the FreeDiskSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeDiskSpace

`func (o *AddLogRetentionPolicyRequest) SetFreeDiskSpace(v string)`

SetFreeDiskSpace sets FreeDiskSpace field to given value.


### GetDiskSpaceUsed

`func (o *AddLogRetentionPolicyRequest) GetDiskSpaceUsed() string`

GetDiskSpaceUsed returns the DiskSpaceUsed field if non-nil, zero value otherwise.

### GetDiskSpaceUsedOk

`func (o *AddLogRetentionPolicyRequest) GetDiskSpaceUsedOk() (*string, bool)`

GetDiskSpaceUsedOk returns a tuple with the DiskSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsed

`func (o *AddLogRetentionPolicyRequest) SetDiskSpaceUsed(v string)`

SetDiskSpaceUsed sets DiskSpaceUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


