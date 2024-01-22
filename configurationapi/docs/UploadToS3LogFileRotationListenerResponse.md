# UploadToS3LogFileRotationListenerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumuploadToS3LogFileRotationListenerSchemaUrn**](EnumuploadToS3LogFileRotationListenerSchemaUrn.md) |  | 
**AwsExternalServer** | **string** | The external server with information to use when interacting with the AWS S3 service. | 
**S3BucketName** | **string** | The name of the S3 bucket into which rotated log files should be copied. | 
**TargetThroughputInMegabitsPerSecond** | Pointer to **int64** | The target throughput to attempt to achieve for data transfers to or from S3, in megabits per second. | [optional] 
**MaximumConcurrentTransferConnections** | Pointer to **int64** | The maximum number of concurrent connections that may be used when transferring data to or from S3. | [optional] 
**MaximumFileCountToRetain** | Pointer to **int64** | The maximum number of existing files matching the file retention pattern that should be retained in the S3 bucket after successfully uploading a newly rotated file. | [optional] 
**MaximumFileAgeToRetain** | Pointer to **string** | The maximum length of time to retain files matching the file retention pattern that should be retained in the S3 bucket after successfully uploading a newly rotated file. | [optional] 
**FileRetentionPattern** | Pointer to **string** | A regular expression pattern that will be used to identify which files are candidates for automatic removal based on the maximum-file-count-to-retain and maximum-file-age-to-retain properties. By default, all files in the bucket will be eligible for removal by retention processing. | [optional] 
**Description** | Pointer to **string** | A description for this Log File Rotation Listener | [optional] 
**Enabled** | **bool** | Indicates whether the Log File Rotation Listener is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Log File Rotation Listener | 

## Methods

### NewUploadToS3LogFileRotationListenerResponse

`func NewUploadToS3LogFileRotationListenerResponse(schemas []EnumuploadToS3LogFileRotationListenerSchemaUrn, awsExternalServer string, s3BucketName string, enabled bool, id string, ) *UploadToS3LogFileRotationListenerResponse`

NewUploadToS3LogFileRotationListenerResponse instantiates a new UploadToS3LogFileRotationListenerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadToS3LogFileRotationListenerResponseWithDefaults

`func NewUploadToS3LogFileRotationListenerResponseWithDefaults() *UploadToS3LogFileRotationListenerResponse`

NewUploadToS3LogFileRotationListenerResponseWithDefaults instantiates a new UploadToS3LogFileRotationListenerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *UploadToS3LogFileRotationListenerResponse) GetSchemas() []EnumuploadToS3LogFileRotationListenerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetSchemasOk() (*[]EnumuploadToS3LogFileRotationListenerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UploadToS3LogFileRotationListenerResponse) SetSchemas(v []EnumuploadToS3LogFileRotationListenerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *UploadToS3LogFileRotationListenerResponse) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *UploadToS3LogFileRotationListenerResponse) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetS3BucketName

`func (o *UploadToS3LogFileRotationListenerResponse) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *UploadToS3LogFileRotationListenerResponse) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.


### GetTargetThroughputInMegabitsPerSecond

`func (o *UploadToS3LogFileRotationListenerResponse) GetTargetThroughputInMegabitsPerSecond() int64`

GetTargetThroughputInMegabitsPerSecond returns the TargetThroughputInMegabitsPerSecond field if non-nil, zero value otherwise.

### GetTargetThroughputInMegabitsPerSecondOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetTargetThroughputInMegabitsPerSecondOk() (*int64, bool)`

GetTargetThroughputInMegabitsPerSecondOk returns a tuple with the TargetThroughputInMegabitsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetThroughputInMegabitsPerSecond

`func (o *UploadToS3LogFileRotationListenerResponse) SetTargetThroughputInMegabitsPerSecond(v int64)`

SetTargetThroughputInMegabitsPerSecond sets TargetThroughputInMegabitsPerSecond field to given value.

### HasTargetThroughputInMegabitsPerSecond

`func (o *UploadToS3LogFileRotationListenerResponse) HasTargetThroughputInMegabitsPerSecond() bool`

HasTargetThroughputInMegabitsPerSecond returns a boolean if a field has been set.

### GetMaximumConcurrentTransferConnections

`func (o *UploadToS3LogFileRotationListenerResponse) GetMaximumConcurrentTransferConnections() int64`

GetMaximumConcurrentTransferConnections returns the MaximumConcurrentTransferConnections field if non-nil, zero value otherwise.

### GetMaximumConcurrentTransferConnectionsOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetMaximumConcurrentTransferConnectionsOk() (*int64, bool)`

GetMaximumConcurrentTransferConnectionsOk returns a tuple with the MaximumConcurrentTransferConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentTransferConnections

`func (o *UploadToS3LogFileRotationListenerResponse) SetMaximumConcurrentTransferConnections(v int64)`

SetMaximumConcurrentTransferConnections sets MaximumConcurrentTransferConnections field to given value.

### HasMaximumConcurrentTransferConnections

`func (o *UploadToS3LogFileRotationListenerResponse) HasMaximumConcurrentTransferConnections() bool`

HasMaximumConcurrentTransferConnections returns a boolean if a field has been set.

### GetMaximumFileCountToRetain

`func (o *UploadToS3LogFileRotationListenerResponse) GetMaximumFileCountToRetain() int64`

GetMaximumFileCountToRetain returns the MaximumFileCountToRetain field if non-nil, zero value otherwise.

### GetMaximumFileCountToRetainOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetMaximumFileCountToRetainOk() (*int64, bool)`

GetMaximumFileCountToRetainOk returns a tuple with the MaximumFileCountToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileCountToRetain

`func (o *UploadToS3LogFileRotationListenerResponse) SetMaximumFileCountToRetain(v int64)`

SetMaximumFileCountToRetain sets MaximumFileCountToRetain field to given value.

### HasMaximumFileCountToRetain

`func (o *UploadToS3LogFileRotationListenerResponse) HasMaximumFileCountToRetain() bool`

HasMaximumFileCountToRetain returns a boolean if a field has been set.

### GetMaximumFileAgeToRetain

`func (o *UploadToS3LogFileRotationListenerResponse) GetMaximumFileAgeToRetain() string`

GetMaximumFileAgeToRetain returns the MaximumFileAgeToRetain field if non-nil, zero value otherwise.

### GetMaximumFileAgeToRetainOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetMaximumFileAgeToRetainOk() (*string, bool)`

GetMaximumFileAgeToRetainOk returns a tuple with the MaximumFileAgeToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileAgeToRetain

`func (o *UploadToS3LogFileRotationListenerResponse) SetMaximumFileAgeToRetain(v string)`

SetMaximumFileAgeToRetain sets MaximumFileAgeToRetain field to given value.

### HasMaximumFileAgeToRetain

`func (o *UploadToS3LogFileRotationListenerResponse) HasMaximumFileAgeToRetain() bool`

HasMaximumFileAgeToRetain returns a boolean if a field has been set.

### GetFileRetentionPattern

`func (o *UploadToS3LogFileRotationListenerResponse) GetFileRetentionPattern() string`

GetFileRetentionPattern returns the FileRetentionPattern field if non-nil, zero value otherwise.

### GetFileRetentionPatternOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetFileRetentionPatternOk() (*string, bool)`

GetFileRetentionPatternOk returns a tuple with the FileRetentionPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRetentionPattern

`func (o *UploadToS3LogFileRotationListenerResponse) SetFileRetentionPattern(v string)`

SetFileRetentionPattern sets FileRetentionPattern field to given value.

### HasFileRetentionPattern

`func (o *UploadToS3LogFileRotationListenerResponse) HasFileRetentionPattern() bool`

HasFileRetentionPattern returns a boolean if a field has been set.

### GetDescription

`func (o *UploadToS3LogFileRotationListenerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UploadToS3LogFileRotationListenerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UploadToS3LogFileRotationListenerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UploadToS3LogFileRotationListenerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UploadToS3LogFileRotationListenerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *UploadToS3LogFileRotationListenerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UploadToS3LogFileRotationListenerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UploadToS3LogFileRotationListenerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *UploadToS3LogFileRotationListenerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *UploadToS3LogFileRotationListenerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *UploadToS3LogFileRotationListenerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *UploadToS3LogFileRotationListenerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *UploadToS3LogFileRotationListenerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadToS3LogFileRotationListenerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadToS3LogFileRotationListenerResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


