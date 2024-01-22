# AddPostLdifExportTaskProcessorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessorName** | **string** | Name of the new Post LDIF Export Task Processor | 
**Schemas** | [**[]EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn**](EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS S3 service. | 
**S3BucketName** | **string** | The name of the S3 bucket into which LDIF files should be copied. | 
**TargetThroughputInMegabitsPerSecond** | Pointer to **int64** | The target throughput to attempt to achieve for data transfers to or from S3, in megabits per second. | [optional] 
**MaximumConcurrentTransferConnections** | Pointer to **int64** | The maximum number of concurrent connections that may be used when transferring data to or from S3. | [optional] 
**MaximumFileCountToRetain** | Pointer to **int64** | The maximum number of existing files matching the file retention pattern that should be retained in the S3 bucket after successfully uploading a newly exported file. | [optional] 
**MaximumFileAgeToRetain** | Pointer to **string** | The maximum length of time to retain files matching the file retention pattern that should be retained in the S3 bucket after successfully uploading a newly exported file. | [optional] 
**FileRetentionPattern** | Pointer to **string** | A regular expression pattern that will be used to identify which files are candidates for automatic removal based on the maximum-file-count-to-retain and maximum-file-age-to-retain properties. By default, all files in the bucket will be eligible for removal by retention processing. | [optional] 
**Description** | Pointer to **string** | A description for this Post LDIF Export Task Processor | [optional] 
**Enabled** | **bool** | Indicates whether the Post LDIF Export Task Processor is enabled for use. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Post LDIF Export Task Processor. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Post LDIF Export Task Processor. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddPostLdifExportTaskProcessorRequest

`func NewAddPostLdifExportTaskProcessorRequest(processorName string, schemas []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn, awsExternalServer string, s3BucketName string, enabled bool, extensionClass string, ) *AddPostLdifExportTaskProcessorRequest`

NewAddPostLdifExportTaskProcessorRequest instantiates a new AddPostLdifExportTaskProcessorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPostLdifExportTaskProcessorRequestWithDefaults

`func NewAddPostLdifExportTaskProcessorRequestWithDefaults() *AddPostLdifExportTaskProcessorRequest`

NewAddPostLdifExportTaskProcessorRequestWithDefaults instantiates a new AddPostLdifExportTaskProcessorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessorName

`func (o *AddPostLdifExportTaskProcessorRequest) GetProcessorName() string`

GetProcessorName returns the ProcessorName field if non-nil, zero value otherwise.

### GetProcessorNameOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetProcessorNameOk() (*string, bool)`

GetProcessorNameOk returns a tuple with the ProcessorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorName

`func (o *AddPostLdifExportTaskProcessorRequest) SetProcessorName(v string)`

SetProcessorName sets ProcessorName field to given value.


### GetSchemas

`func (o *AddPostLdifExportTaskProcessorRequest) GetSchemas() []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetSchemasOk() (*[]EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPostLdifExportTaskProcessorRequest) SetSchemas(v []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AddPostLdifExportTaskProcessorRequest) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddPostLdifExportTaskProcessorRequest) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetS3BucketName

`func (o *AddPostLdifExportTaskProcessorRequest) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *AddPostLdifExportTaskProcessorRequest) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.


### GetTargetThroughputInMegabitsPerSecond

`func (o *AddPostLdifExportTaskProcessorRequest) GetTargetThroughputInMegabitsPerSecond() int64`

GetTargetThroughputInMegabitsPerSecond returns the TargetThroughputInMegabitsPerSecond field if non-nil, zero value otherwise.

### GetTargetThroughputInMegabitsPerSecondOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetTargetThroughputInMegabitsPerSecondOk() (*int64, bool)`

GetTargetThroughputInMegabitsPerSecondOk returns a tuple with the TargetThroughputInMegabitsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetThroughputInMegabitsPerSecond

`func (o *AddPostLdifExportTaskProcessorRequest) SetTargetThroughputInMegabitsPerSecond(v int64)`

SetTargetThroughputInMegabitsPerSecond sets TargetThroughputInMegabitsPerSecond field to given value.

### HasTargetThroughputInMegabitsPerSecond

`func (o *AddPostLdifExportTaskProcessorRequest) HasTargetThroughputInMegabitsPerSecond() bool`

HasTargetThroughputInMegabitsPerSecond returns a boolean if a field has been set.

### GetMaximumConcurrentTransferConnections

`func (o *AddPostLdifExportTaskProcessorRequest) GetMaximumConcurrentTransferConnections() int64`

GetMaximumConcurrentTransferConnections returns the MaximumConcurrentTransferConnections field if non-nil, zero value otherwise.

### GetMaximumConcurrentTransferConnectionsOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetMaximumConcurrentTransferConnectionsOk() (*int64, bool)`

GetMaximumConcurrentTransferConnectionsOk returns a tuple with the MaximumConcurrentTransferConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentTransferConnections

`func (o *AddPostLdifExportTaskProcessorRequest) SetMaximumConcurrentTransferConnections(v int64)`

SetMaximumConcurrentTransferConnections sets MaximumConcurrentTransferConnections field to given value.

### HasMaximumConcurrentTransferConnections

`func (o *AddPostLdifExportTaskProcessorRequest) HasMaximumConcurrentTransferConnections() bool`

HasMaximumConcurrentTransferConnections returns a boolean if a field has been set.

### GetMaximumFileCountToRetain

`func (o *AddPostLdifExportTaskProcessorRequest) GetMaximumFileCountToRetain() int64`

GetMaximumFileCountToRetain returns the MaximumFileCountToRetain field if non-nil, zero value otherwise.

### GetMaximumFileCountToRetainOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetMaximumFileCountToRetainOk() (*int64, bool)`

GetMaximumFileCountToRetainOk returns a tuple with the MaximumFileCountToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileCountToRetain

`func (o *AddPostLdifExportTaskProcessorRequest) SetMaximumFileCountToRetain(v int64)`

SetMaximumFileCountToRetain sets MaximumFileCountToRetain field to given value.

### HasMaximumFileCountToRetain

`func (o *AddPostLdifExportTaskProcessorRequest) HasMaximumFileCountToRetain() bool`

HasMaximumFileCountToRetain returns a boolean if a field has been set.

### GetMaximumFileAgeToRetain

`func (o *AddPostLdifExportTaskProcessorRequest) GetMaximumFileAgeToRetain() string`

GetMaximumFileAgeToRetain returns the MaximumFileAgeToRetain field if non-nil, zero value otherwise.

### GetMaximumFileAgeToRetainOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetMaximumFileAgeToRetainOk() (*string, bool)`

GetMaximumFileAgeToRetainOk returns a tuple with the MaximumFileAgeToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileAgeToRetain

`func (o *AddPostLdifExportTaskProcessorRequest) SetMaximumFileAgeToRetain(v string)`

SetMaximumFileAgeToRetain sets MaximumFileAgeToRetain field to given value.

### HasMaximumFileAgeToRetain

`func (o *AddPostLdifExportTaskProcessorRequest) HasMaximumFileAgeToRetain() bool`

HasMaximumFileAgeToRetain returns a boolean if a field has been set.

### GetFileRetentionPattern

`func (o *AddPostLdifExportTaskProcessorRequest) GetFileRetentionPattern() string`

GetFileRetentionPattern returns the FileRetentionPattern field if non-nil, zero value otherwise.

### GetFileRetentionPatternOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetFileRetentionPatternOk() (*string, bool)`

GetFileRetentionPatternOk returns a tuple with the FileRetentionPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRetentionPattern

`func (o *AddPostLdifExportTaskProcessorRequest) SetFileRetentionPattern(v string)`

SetFileRetentionPattern sets FileRetentionPattern field to given value.

### HasFileRetentionPattern

`func (o *AddPostLdifExportTaskProcessorRequest) HasFileRetentionPattern() bool`

HasFileRetentionPattern returns a boolean if a field has been set.

### GetDescription

`func (o *AddPostLdifExportTaskProcessorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPostLdifExportTaskProcessorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPostLdifExportTaskProcessorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPostLdifExportTaskProcessorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPostLdifExportTaskProcessorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExtensionClass

`func (o *AddPostLdifExportTaskProcessorRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPostLdifExportTaskProcessorRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPostLdifExportTaskProcessorRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPostLdifExportTaskProcessorRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPostLdifExportTaskProcessorRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPostLdifExportTaskProcessorRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


