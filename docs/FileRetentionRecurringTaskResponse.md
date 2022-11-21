# FileRetentionRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Recurring Task | 
**Schemas** | [**[]EnumfileRetentionRecurringTaskSchemaUrn**](EnumfileRetentionRecurringTaskSchemaUrn.md) |  | 
**TargetDirectory** | **string** | The path to the directory containing the files to examine. The directory must exist. | 
**FilenamePattern** | **string** | A pattern that specifies the names of the files to examine. The pattern may contain zero or more asterisks as wildcards, where each wildcard matches zero or more characters. It may also contain at most one occurrence of the special string \&quot;${timestamp}\&quot;, which will match a timestamp with the format specified using the timestamp-format property. All other characters in the pattern will be treated literally. | 
**TimestampFormat** | [**EnumrecurringTaskTimestampFormatProp**](EnumrecurringTaskTimestampFormatProp.md) |  | 
**RetainFileCount** | Pointer to **int32** | The minimum number of files matching the pattern that will be retained. | [optional] 
**RetainFileAge** | Pointer to **string** | The minimum age of files matching the pattern that will be retained. | [optional] 
**RetainAggregateFileSize** | Pointer to **string** | The minimum aggregate size of files that will be retained. The size should be specified as an integer followed by a unit that is one of \&quot;b\&quot; or \&quot;bytes\&quot;, \&quot;kb\&quot; or \&quot;kilobytes\&quot;, \&quot;mb\&quot; or \&quot;megabytes\&quot;, \&quot;gb\&quot; or \&quot;gigabytes\&quot;, or \&quot;tb\&quot; or \&quot;terabytes\&quot;. For example, a value of \&quot;1 gb\&quot; indicates that at least one gigabyte of files should be retained. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Methods

### NewFileRetentionRecurringTaskResponse

`func NewFileRetentionRecurringTaskResponse(id string, schemas []EnumfileRetentionRecurringTaskSchemaUrn, targetDirectory string, filenamePattern string, timestampFormat EnumrecurringTaskTimestampFormatProp, ) *FileRetentionRecurringTaskResponse`

NewFileRetentionRecurringTaskResponse instantiates a new FileRetentionRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileRetentionRecurringTaskResponseWithDefaults

`func NewFileRetentionRecurringTaskResponseWithDefaults() *FileRetentionRecurringTaskResponse`

NewFileRetentionRecurringTaskResponseWithDefaults instantiates a new FileRetentionRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FileRetentionRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileRetentionRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileRetentionRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FileRetentionRecurringTaskResponse) GetSchemas() []EnumfileRetentionRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileRetentionRecurringTaskResponse) GetSchemasOk() (*[]EnumfileRetentionRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileRetentionRecurringTaskResponse) SetSchemas(v []EnumfileRetentionRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTargetDirectory

`func (o *FileRetentionRecurringTaskResponse) GetTargetDirectory() string`

GetTargetDirectory returns the TargetDirectory field if non-nil, zero value otherwise.

### GetTargetDirectoryOk

`func (o *FileRetentionRecurringTaskResponse) GetTargetDirectoryOk() (*string, bool)`

GetTargetDirectoryOk returns a tuple with the TargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectory

`func (o *FileRetentionRecurringTaskResponse) SetTargetDirectory(v string)`

SetTargetDirectory sets TargetDirectory field to given value.


### GetFilenamePattern

`func (o *FileRetentionRecurringTaskResponse) GetFilenamePattern() string`

GetFilenamePattern returns the FilenamePattern field if non-nil, zero value otherwise.

### GetFilenamePatternOk

`func (o *FileRetentionRecurringTaskResponse) GetFilenamePatternOk() (*string, bool)`

GetFilenamePatternOk returns a tuple with the FilenamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenamePattern

`func (o *FileRetentionRecurringTaskResponse) SetFilenamePattern(v string)`

SetFilenamePattern sets FilenamePattern field to given value.


### GetTimestampFormat

`func (o *FileRetentionRecurringTaskResponse) GetTimestampFormat() EnumrecurringTaskTimestampFormatProp`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *FileRetentionRecurringTaskResponse) GetTimestampFormatOk() (*EnumrecurringTaskTimestampFormatProp, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *FileRetentionRecurringTaskResponse) SetTimestampFormat(v EnumrecurringTaskTimestampFormatProp)`

SetTimestampFormat sets TimestampFormat field to given value.


### GetRetainFileCount

`func (o *FileRetentionRecurringTaskResponse) GetRetainFileCount() int32`

GetRetainFileCount returns the RetainFileCount field if non-nil, zero value otherwise.

### GetRetainFileCountOk

`func (o *FileRetentionRecurringTaskResponse) GetRetainFileCountOk() (*int32, bool)`

GetRetainFileCountOk returns a tuple with the RetainFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileCount

`func (o *FileRetentionRecurringTaskResponse) SetRetainFileCount(v int32)`

SetRetainFileCount sets RetainFileCount field to given value.

### HasRetainFileCount

`func (o *FileRetentionRecurringTaskResponse) HasRetainFileCount() bool`

HasRetainFileCount returns a boolean if a field has been set.

### GetRetainFileAge

`func (o *FileRetentionRecurringTaskResponse) GetRetainFileAge() string`

GetRetainFileAge returns the RetainFileAge field if non-nil, zero value otherwise.

### GetRetainFileAgeOk

`func (o *FileRetentionRecurringTaskResponse) GetRetainFileAgeOk() (*string, bool)`

GetRetainFileAgeOk returns a tuple with the RetainFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileAge

`func (o *FileRetentionRecurringTaskResponse) SetRetainFileAge(v string)`

SetRetainFileAge sets RetainFileAge field to given value.

### HasRetainFileAge

`func (o *FileRetentionRecurringTaskResponse) HasRetainFileAge() bool`

HasRetainFileAge returns a boolean if a field has been set.

### GetRetainAggregateFileSize

`func (o *FileRetentionRecurringTaskResponse) GetRetainAggregateFileSize() string`

GetRetainAggregateFileSize returns the RetainAggregateFileSize field if non-nil, zero value otherwise.

### GetRetainAggregateFileSizeOk

`func (o *FileRetentionRecurringTaskResponse) GetRetainAggregateFileSizeOk() (*string, bool)`

GetRetainAggregateFileSizeOk returns a tuple with the RetainAggregateFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAggregateFileSize

`func (o *FileRetentionRecurringTaskResponse) SetRetainAggregateFileSize(v string)`

SetRetainAggregateFileSize sets RetainAggregateFileSize field to given value.

### HasRetainAggregateFileSize

`func (o *FileRetentionRecurringTaskResponse) HasRetainAggregateFileSize() bool`

HasRetainAggregateFileSize returns a boolean if a field has been set.

### GetDescription

`func (o *FileRetentionRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileRetentionRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileRetentionRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileRetentionRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *FileRetentionRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *FileRetentionRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *FileRetentionRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *FileRetentionRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *FileRetentionRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *FileRetentionRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *FileRetentionRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *FileRetentionRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *FileRetentionRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *FileRetentionRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *FileRetentionRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *FileRetentionRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *FileRetentionRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *FileRetentionRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *FileRetentionRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *FileRetentionRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *FileRetentionRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *FileRetentionRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *FileRetentionRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *FileRetentionRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *FileRetentionRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *FileRetentionRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *FileRetentionRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *FileRetentionRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *FileRetentionRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *FileRetentionRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *FileRetentionRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *FileRetentionRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.

### GetMeta

`func (o *FileRetentionRecurringTaskResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileRetentionRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileRetentionRecurringTaskResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FileRetentionRecurringTaskResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


