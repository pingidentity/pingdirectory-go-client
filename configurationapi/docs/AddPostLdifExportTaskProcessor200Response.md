# AddPostLdifExportTaskProcessor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Post LDIF Export Task Processor | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Post LDIF Export Task Processor. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Post LDIF Export Task Processor. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddPostLdifExportTaskProcessor200Response

`func NewAddPostLdifExportTaskProcessor200Response(id string, schemas []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn, awsExternalServer string, s3BucketName string, enabled bool, extensionClass string, ) *AddPostLdifExportTaskProcessor200Response`

NewAddPostLdifExportTaskProcessor200Response instantiates a new AddPostLdifExportTaskProcessor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPostLdifExportTaskProcessor200ResponseWithDefaults

`func NewAddPostLdifExportTaskProcessor200ResponseWithDefaults() *AddPostLdifExportTaskProcessor200Response`

NewAddPostLdifExportTaskProcessor200ResponseWithDefaults instantiates a new AddPostLdifExportTaskProcessor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddPostLdifExportTaskProcessor200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPostLdifExportTaskProcessor200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddPostLdifExportTaskProcessor200Response) GetSchemas() []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetSchemasOk() (*[]EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPostLdifExportTaskProcessor200Response) SetSchemas(v []EnumthirdPartyPostLdifExportTaskProcessorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAwsExternalServer

`func (o *AddPostLdifExportTaskProcessor200Response) GetAwsExternalServer() string`

GetAwsExternalServer returns the AwsExternalServer field if non-nil, zero value otherwise.

### GetAwsExternalServerOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetAwsExternalServerOk() (*string, bool)`

GetAwsExternalServerOk returns a tuple with the AwsExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsExternalServer

`func (o *AddPostLdifExportTaskProcessor200Response) SetAwsExternalServer(v string)`

SetAwsExternalServer sets AwsExternalServer field to given value.


### GetS3BucketName

`func (o *AddPostLdifExportTaskProcessor200Response) GetS3BucketName() string`

GetS3BucketName returns the S3BucketName field if non-nil, zero value otherwise.

### GetS3BucketNameOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetS3BucketNameOk() (*string, bool)`

GetS3BucketNameOk returns a tuple with the S3BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketName

`func (o *AddPostLdifExportTaskProcessor200Response) SetS3BucketName(v string)`

SetS3BucketName sets S3BucketName field to given value.


### GetTargetThroughputInMegabitsPerSecond

`func (o *AddPostLdifExportTaskProcessor200Response) GetTargetThroughputInMegabitsPerSecond() int64`

GetTargetThroughputInMegabitsPerSecond returns the TargetThroughputInMegabitsPerSecond field if non-nil, zero value otherwise.

### GetTargetThroughputInMegabitsPerSecondOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetTargetThroughputInMegabitsPerSecondOk() (*int64, bool)`

GetTargetThroughputInMegabitsPerSecondOk returns a tuple with the TargetThroughputInMegabitsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetThroughputInMegabitsPerSecond

`func (o *AddPostLdifExportTaskProcessor200Response) SetTargetThroughputInMegabitsPerSecond(v int64)`

SetTargetThroughputInMegabitsPerSecond sets TargetThroughputInMegabitsPerSecond field to given value.

### HasTargetThroughputInMegabitsPerSecond

`func (o *AddPostLdifExportTaskProcessor200Response) HasTargetThroughputInMegabitsPerSecond() bool`

HasTargetThroughputInMegabitsPerSecond returns a boolean if a field has been set.

### GetMaximumConcurrentTransferConnections

`func (o *AddPostLdifExportTaskProcessor200Response) GetMaximumConcurrentTransferConnections() int64`

GetMaximumConcurrentTransferConnections returns the MaximumConcurrentTransferConnections field if non-nil, zero value otherwise.

### GetMaximumConcurrentTransferConnectionsOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetMaximumConcurrentTransferConnectionsOk() (*int64, bool)`

GetMaximumConcurrentTransferConnectionsOk returns a tuple with the MaximumConcurrentTransferConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumConcurrentTransferConnections

`func (o *AddPostLdifExportTaskProcessor200Response) SetMaximumConcurrentTransferConnections(v int64)`

SetMaximumConcurrentTransferConnections sets MaximumConcurrentTransferConnections field to given value.

### HasMaximumConcurrentTransferConnections

`func (o *AddPostLdifExportTaskProcessor200Response) HasMaximumConcurrentTransferConnections() bool`

HasMaximumConcurrentTransferConnections returns a boolean if a field has been set.

### GetMaximumFileCountToRetain

`func (o *AddPostLdifExportTaskProcessor200Response) GetMaximumFileCountToRetain() int64`

GetMaximumFileCountToRetain returns the MaximumFileCountToRetain field if non-nil, zero value otherwise.

### GetMaximumFileCountToRetainOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetMaximumFileCountToRetainOk() (*int64, bool)`

GetMaximumFileCountToRetainOk returns a tuple with the MaximumFileCountToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileCountToRetain

`func (o *AddPostLdifExportTaskProcessor200Response) SetMaximumFileCountToRetain(v int64)`

SetMaximumFileCountToRetain sets MaximumFileCountToRetain field to given value.

### HasMaximumFileCountToRetain

`func (o *AddPostLdifExportTaskProcessor200Response) HasMaximumFileCountToRetain() bool`

HasMaximumFileCountToRetain returns a boolean if a field has been set.

### GetMaximumFileAgeToRetain

`func (o *AddPostLdifExportTaskProcessor200Response) GetMaximumFileAgeToRetain() string`

GetMaximumFileAgeToRetain returns the MaximumFileAgeToRetain field if non-nil, zero value otherwise.

### GetMaximumFileAgeToRetainOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetMaximumFileAgeToRetainOk() (*string, bool)`

GetMaximumFileAgeToRetainOk returns a tuple with the MaximumFileAgeToRetain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumFileAgeToRetain

`func (o *AddPostLdifExportTaskProcessor200Response) SetMaximumFileAgeToRetain(v string)`

SetMaximumFileAgeToRetain sets MaximumFileAgeToRetain field to given value.

### HasMaximumFileAgeToRetain

`func (o *AddPostLdifExportTaskProcessor200Response) HasMaximumFileAgeToRetain() bool`

HasMaximumFileAgeToRetain returns a boolean if a field has been set.

### GetFileRetentionPattern

`func (o *AddPostLdifExportTaskProcessor200Response) GetFileRetentionPattern() string`

GetFileRetentionPattern returns the FileRetentionPattern field if non-nil, zero value otherwise.

### GetFileRetentionPatternOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetFileRetentionPatternOk() (*string, bool)`

GetFileRetentionPatternOk returns a tuple with the FileRetentionPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRetentionPattern

`func (o *AddPostLdifExportTaskProcessor200Response) SetFileRetentionPattern(v string)`

SetFileRetentionPattern sets FileRetentionPattern field to given value.

### HasFileRetentionPattern

`func (o *AddPostLdifExportTaskProcessor200Response) HasFileRetentionPattern() bool`

HasFileRetentionPattern returns a boolean if a field has been set.

### GetDescription

`func (o *AddPostLdifExportTaskProcessor200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPostLdifExportTaskProcessor200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPostLdifExportTaskProcessor200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPostLdifExportTaskProcessor200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPostLdifExportTaskProcessor200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AddPostLdifExportTaskProcessor200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddPostLdifExportTaskProcessor200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddPostLdifExportTaskProcessor200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddPostLdifExportTaskProcessor200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddPostLdifExportTaskProcessor200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddPostLdifExportTaskProcessor200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddPostLdifExportTaskProcessor200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddPostLdifExportTaskProcessor200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPostLdifExportTaskProcessor200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPostLdifExportTaskProcessor200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPostLdifExportTaskProcessor200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPostLdifExportTaskProcessor200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPostLdifExportTaskProcessor200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


