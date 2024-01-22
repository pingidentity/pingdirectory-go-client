# AddLogFileRotationListenerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListenerName** | **string** | Name of the new Log File Rotation Listener | 
**Schemas** | [**[]EnumthirdPartyLogFileRotationListenerSchemaUrn**](EnumthirdPartyLogFileRotationListenerSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS S3 service. | 
**S3BucketName** | **string** | The name of the S3 bucket into which rotated log files should be copied. | 
**TargetThroughputInMegabitsPerSecond** | Pointer to **int64** | The target throughput to attempt to achieve for data transfers to or from S3, in megabits per second. | [optional] 
**MaximumConcurrentTransferConnections** | Pointer to **int64** | The maximum number of concurrent connections that may be used when transferring data to or from S3. | [optional] 
**MaximumFileCountToRetain** | Pointer to **int64** | The maximum number of existing files matching the file retention pattern that should be retained in the S3 bucket after successfully uploading a newly rotated file. | [optional] 
**MaximumFileAgeToRetain** | Pointer to **string** | The maximum length of time to retain files matching the file retention pattern that should be retained in the S3 bucket after successfully uploading a newly rotated file. | [optional] 
**FileRetentionPattern** | Pointer to **string** | A regular expression pattern that will be used to identify which files are candidates for automatic removal based on the maximum-file-count-to-retain and maximum-file-age-to-retain properties. By default, all files in the bucket will be eligible for removal by retention processing. | [optional] 
**Description** | Pointer to **string** | A description for this Log File Rotation Listener | [optional] 
**Enabled** | **bool** | Indicates whether the Log File Rotation Listener is enabled for use. | 
**OutputDirectory** | Pointer to **string** | The path to the directory in which the summarize-access-log output should be written. If no value is provided, the output file will be written into the same directory as the rotated log file. | [optional] 
**CopyToDirectory** | **string** | The path to the directory to which log files should be copied. It must be different from the directory to which the log file is originally written, and administrators should ensure that the filesystem has sufficient space to hold files as they are copied. | 
**CompressOnCopy** | Pointer to **bool** | Indicates whether the file should be gzip-compressed as it is copied into the destination directory. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Log File Rotation Listener. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Log File Rotation Listener. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddLogFileRotationListenerRequest

`func NewAddLogFileRotationListenerRequest(listenerName string, schemas []EnumthirdPartyLogFileRotationListenerSchemaUrn, awsExternalServer string, s3BucketName string, enabled bool, copyToDirectory string, extensionClass string, ) *AddLogFileRotationListenerRequest`

NewAddLogFileRotationListenerRequest instantiates a new AddLogFileRotationListenerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogFileRotationListenerRequestWithDefaults

`func NewAddLogFileRotationListenerRequestWithDefaults() *AddLogFileRotationListenerRequest`

NewAddLogFileRotationListenerRequestWithDefaults instantiates a new AddLogFileRotationListenerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListenerName

`func (o *AddLogFileRotationListenerRequest) GetListenerName() string`

GetListenerName returns the ListenerName field if non-nil, zero value otherwise.

### GetListenerNameOk

`func (o *AddLogFileRotationListenerRequest) GetListenerNameOk() (*string, bool)`

GetListenerNameOk returns a tuple with the ListenerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerName

`func (o *AddLogFileRotationListenerRequest) SetListenerName(v string)`

SetListenerName sets ListenerName field to given value.


### GetSchemas

`func (o *AddLogFileRotationListenerRequest) GetSchemas() []EnumthirdPartyLogFileRotationListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogFileRotationListenerRequest) GetSchemasOk() (*[]EnumthirdPartyLogFileRotationListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogFileRotationListenerRequest) SetSchemas(v []EnumthirdPartyLogFileRotationListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AddLogFileRotationListenerRequest) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddLogFileRotationListenerRequest) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddLogFileRotationListenerRequest) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetS3BucketName

`func (o *AddLogFileRotationListenerRequest) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *AddLogFileRotationListenerRequest) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *AddLogFileRotationListenerRequest) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.


### GetTargetThroughputInMegabitsPerSecond

`func (o *AddLogFileRotationListenerRequest) GetTargetThroughputInMegabitsPerSecond() int64`

GetTargetThroughputInMegabitsPerSecond returns the TargetThroughputInMegabitsPerSecond field if non-nil, zero value otherwise.

### GetTargetThroughputInMegabitsPerSecondOk

`func (o *AddLogFileRotationListenerRequest) GetTargetThroughputInMegabitsPerSecondOk() (*int64, bool)`

GetTargetThroughputInMegabitsPerSecondOk returns a tuple with the TargetThroughputInMegabitsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetThroughputInMegabitsPerSecond

`func (o *AddLogFileRotationListenerRequest) SetTargetThroughputInMegabitsPerSecond(v int64)`

SetTargetThroughputInMegabitsPerSecond sets TargetThroughputInMegabitsPerSecond field to given value.

### HasTargetThroughputInMegabitsPerSecond

`func (o *AddLogFileRotationListenerRequest) HasTargetThroughputInMegabitsPerSecond() bool`

HasTargetThroughputInMegabitsPerSecond returns a boolean if a field has been set.

### GetMaximumConcurrentTransferConnections

`func (o *AddLogFileRotationListenerRequest) GetMaximumConcurrentTransferConnections() int64`

GetMaximumConcurrentTransferConnections returns the MaximumConcurrentTransferConnections field if non-nil, zero value otherwise.

### GetMaximumConcurrentTransferConnectionsOk

`func (o *AddLogFileRotationListenerRequest) GetMaximumConcurrentTransferConnectionsOk() (*int64, bool)`

GetMaximumConcurrentTransferConnectionsOk returns a tuple with the MaximumConcurrentTransferConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentTransferConnections

`func (o *AddLogFileRotationListenerRequest) SetMaximumConcurrentTransferConnections(v int64)`

SetMaximumConcurrentTransferConnections sets MaximumConcurrentTransferConnections field to given value.

### HasMaximumConcurrentTransferConnections

`func (o *AddLogFileRotationListenerRequest) HasMaximumConcurrentTransferConnections() bool`

HasMaximumConcurrentTransferConnections returns a boolean if a field has been set.

### GetMaximumFileCountToRetain

`func (o *AddLogFileRotationListenerRequest) GetMaximumFileCountToRetain() int64`

GetMaximumFileCountToRetain returns the MaximumFileCountToRetain field if non-nil, zero value otherwise.

### GetMaximumFileCountToRetainOk

`func (o *AddLogFileRotationListenerRequest) GetMaximumFileCountToRetainOk() (*int64, bool)`

GetMaximumFileCountToRetainOk returns a tuple with the MaximumFileCountToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileCountToRetain

`func (o *AddLogFileRotationListenerRequest) SetMaximumFileCountToRetain(v int64)`

SetMaximumFileCountToRetain sets MaximumFileCountToRetain field to given value.

### HasMaximumFileCountToRetain

`func (o *AddLogFileRotationListenerRequest) HasMaximumFileCountToRetain() bool`

HasMaximumFileCountToRetain returns a boolean if a field has been set.

### GetMaximumFileAgeToRetain

`func (o *AddLogFileRotationListenerRequest) GetMaximumFileAgeToRetain() string`

GetMaximumFileAgeToRetain returns the MaximumFileAgeToRetain field if non-nil, zero value otherwise.

### GetMaximumFileAgeToRetainOk

`func (o *AddLogFileRotationListenerRequest) GetMaximumFileAgeToRetainOk() (*string, bool)`

GetMaximumFileAgeToRetainOk returns a tuple with the MaximumFileAgeToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileAgeToRetain

`func (o *AddLogFileRotationListenerRequest) SetMaximumFileAgeToRetain(v string)`

SetMaximumFileAgeToRetain sets MaximumFileAgeToRetain field to given value.

### HasMaximumFileAgeToRetain

`func (o *AddLogFileRotationListenerRequest) HasMaximumFileAgeToRetain() bool`

HasMaximumFileAgeToRetain returns a boolean if a field has been set.

### GetFileRetentionPattern

`func (o *AddLogFileRotationListenerRequest) GetFileRetentionPattern() string`

GetFileRetentionPattern returns the FileRetentionPattern field if non-nil, zero value otherwise.

### GetFileRetentionPatternOk

`func (o *AddLogFileRotationListenerRequest) GetFileRetentionPatternOk() (*string, bool)`

GetFileRetentionPatternOk returns a tuple with the FileRetentionPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRetentionPattern

`func (o *AddLogFileRotationListenerRequest) SetFileRetentionPattern(v string)`

SetFileRetentionPattern sets FileRetentionPattern field to given value.

### HasFileRetentionPattern

`func (o *AddLogFileRotationListenerRequest) HasFileRetentionPattern() bool`

HasFileRetentionPattern returns a boolean if a field has been set.

### GetDescription

`func (o *AddLogFileRotationListenerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogFileRotationListenerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogFileRotationListenerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogFileRotationListenerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddLogFileRotationListenerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddLogFileRotationListenerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddLogFileRotationListenerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOutputDirectory

`func (o *AddLogFileRotationListenerRequest) GetOutputDirectory() string`

GetOutputDirectory returns the OutputDirectory field if non-nil, zero value otherwise.

### GetOutputDirectoryOk

`func (o *AddLogFileRotationListenerRequest) GetOutputDirectoryOk() (*string, bool)`

GetOutputDirectoryOk returns a tuple with the OutputDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDirectory

`func (o *AddLogFileRotationListenerRequest) SetOutputDirectory(v string)`

SetOutputDirectory sets OutputDirectory field to given value.

### HasOutputDirectory

`func (o *AddLogFileRotationListenerRequest) HasOutputDirectory() bool`

HasOutputDirectory returns a boolean if a field has been set.

### GetCopyToDirectory

`func (o *AddLogFileRotationListenerRequest) GetCopyToDirectory() string`

GetCopyToDirectory returns the CopyToDirectory field if non-nil, zero value otherwise.

### GetCopyToDirectoryOk

`func (o *AddLogFileRotationListenerRequest) GetCopyToDirectoryOk() (*string, bool)`

GetCopyToDirectoryOk returns a tuple with the CopyToDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToDirectory

`func (o *AddLogFileRotationListenerRequest) SetCopyToDirectory(v string)`

SetCopyToDirectory sets CopyToDirectory field to given value.


### GetCompressOnCopy

`func (o *AddLogFileRotationListenerRequest) GetCompressOnCopy() bool`

GetCompressOnCopy returns the CompressOnCopy field if non-nil, zero value otherwise.

### GetCompressOnCopyOk

`func (o *AddLogFileRotationListenerRequest) GetCompressOnCopyOk() (*bool, bool)`

GetCompressOnCopyOk returns a tuple with the CompressOnCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressOnCopy

`func (o *AddLogFileRotationListenerRequest) SetCompressOnCopy(v bool)`

SetCompressOnCopy sets CompressOnCopy field to given value.

### HasCompressOnCopy

`func (o *AddLogFileRotationListenerRequest) HasCompressOnCopy() bool`

HasCompressOnCopy returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddLogFileRotationListenerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddLogFileRotationListenerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddLogFileRotationListenerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddLogFileRotationListenerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddLogFileRotationListenerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddLogFileRotationListenerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddLogFileRotationListenerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


