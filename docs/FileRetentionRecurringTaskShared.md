# FileRetentionRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileRetentionRecurringTaskSchemaUrn**](EnumfileRetentionRecurringTaskSchemaUrn.md) |  | 
**TargetDirectory** | **string** | The path to the directory containing the files to examine. The directory must exist. | 
**FilenamePattern** | **string** | A pattern that specifies the names of the files to examine. The pattern may contain zero or more asterisks as wildcards, where each wildcard matches zero or more characters. It may also contain at most one occurrence of the special string \&quot;${timestamp}\&quot;, which will match a timestamp with the format specified using the timestamp-format property. All other characters in the pattern will be treated literally. | 
**TimestampFormat** | [**EnumrecurringTaskTimestampFormatProp**](EnumrecurringTaskTimestampFormatProp.md) |  | 
**RetainFileCount** | Pointer to **int32** | The minimum number of files matching the pattern that will be retained. | [optional] 
**RetainFileAge** | Pointer to **string** | The minimum age of files matching the pattern that will be retained. | [optional] 
**RetainAggregateFileSize** | Pointer to **string** | The minimum aggregate size of files that will be retained. The size should be specified as an integer followed by a unit that is one of \&quot;b\&quot; or \&quot;bytes\&quot;, \&quot;kb\&quot; or \&quot;kilobytes\&quot;, \&quot;mb\&quot; or \&quot;megabytes\&quot;, \&quot;gb\&quot; or \&quot;gigabytes\&quot;, or \&quot;tb\&quot; or \&quot;terabytes\&quot;. For example, a value of \&quot;1 gb\&quot; indicates that at least one gigabyte of files should be retained. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewFileRetentionRecurringTaskShared

`func NewFileRetentionRecurringTaskShared(schemas []EnumfileRetentionRecurringTaskSchemaUrn, targetDirectory string, filenamePattern string, timestampFormat EnumrecurringTaskTimestampFormatProp, ) *FileRetentionRecurringTaskShared`

NewFileRetentionRecurringTaskShared instantiates a new FileRetentionRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileRetentionRecurringTaskSharedWithDefaults

`func NewFileRetentionRecurringTaskSharedWithDefaults() *FileRetentionRecurringTaskShared`

NewFileRetentionRecurringTaskSharedWithDefaults instantiates a new FileRetentionRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FileRetentionRecurringTaskShared) GetSchemas() []EnumfileRetentionRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileRetentionRecurringTaskShared) GetSchemasOk() (*[]EnumfileRetentionRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileRetentionRecurringTaskShared) SetSchemas(v []EnumfileRetentionRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetTargetDirectory

`func (o *FileRetentionRecurringTaskShared) GetTargetDirectory() string`

GetTargetDirectory returns the TargetDirectory field if non-nil, zero value otherwise.

### GetTargetDirectoryOk

`func (o *FileRetentionRecurringTaskShared) GetTargetDirectoryOk() (*string, bool)`

GetTargetDirectoryOk returns a tuple with the TargetDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDirectory

`func (o *FileRetentionRecurringTaskShared) SetTargetDirectory(v string)`

SetTargetDirectory sets TargetDirectory field to given value.


### GetFilenamePattern

`func (o *FileRetentionRecurringTaskShared) GetFilenamePattern() string`

GetFilenamePattern returns the FilenamePattern field if non-nil, zero value otherwise.

### GetFilenamePatternOk

`func (o *FileRetentionRecurringTaskShared) GetFilenamePatternOk() (*string, bool)`

GetFilenamePatternOk returns a tuple with the FilenamePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenamePattern

`func (o *FileRetentionRecurringTaskShared) SetFilenamePattern(v string)`

SetFilenamePattern sets FilenamePattern field to given value.


### GetTimestampFormat

`func (o *FileRetentionRecurringTaskShared) GetTimestampFormat() EnumrecurringTaskTimestampFormatProp`

GetTimestampFormat returns the TimestampFormat field if non-nil, zero value otherwise.

### GetTimestampFormatOk

`func (o *FileRetentionRecurringTaskShared) GetTimestampFormatOk() (*EnumrecurringTaskTimestampFormatProp, bool)`

GetTimestampFormatOk returns a tuple with the TimestampFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampFormat

`func (o *FileRetentionRecurringTaskShared) SetTimestampFormat(v EnumrecurringTaskTimestampFormatProp)`

SetTimestampFormat sets TimestampFormat field to given value.


### GetRetainFileCount

`func (o *FileRetentionRecurringTaskShared) GetRetainFileCount() int32`

GetRetainFileCount returns the RetainFileCount field if non-nil, zero value otherwise.

### GetRetainFileCountOk

`func (o *FileRetentionRecurringTaskShared) GetRetainFileCountOk() (*int32, bool)`

GetRetainFileCountOk returns a tuple with the RetainFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileCount

`func (o *FileRetentionRecurringTaskShared) SetRetainFileCount(v int32)`

SetRetainFileCount sets RetainFileCount field to given value.

### HasRetainFileCount

`func (o *FileRetentionRecurringTaskShared) HasRetainFileCount() bool`

HasRetainFileCount returns a boolean if a field has been set.

### GetRetainFileAge

`func (o *FileRetentionRecurringTaskShared) GetRetainFileAge() string`

GetRetainFileAge returns the RetainFileAge field if non-nil, zero value otherwise.

### GetRetainFileAgeOk

`func (o *FileRetentionRecurringTaskShared) GetRetainFileAgeOk() (*string, bool)`

GetRetainFileAgeOk returns a tuple with the RetainFileAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFileAge

`func (o *FileRetentionRecurringTaskShared) SetRetainFileAge(v string)`

SetRetainFileAge sets RetainFileAge field to given value.

### HasRetainFileAge

`func (o *FileRetentionRecurringTaskShared) HasRetainFileAge() bool`

HasRetainFileAge returns a boolean if a field has been set.

### GetRetainAggregateFileSize

`func (o *FileRetentionRecurringTaskShared) GetRetainAggregateFileSize() string`

GetRetainAggregateFileSize returns the RetainAggregateFileSize field if non-nil, zero value otherwise.

### GetRetainAggregateFileSizeOk

`func (o *FileRetentionRecurringTaskShared) GetRetainAggregateFileSizeOk() (*string, bool)`

GetRetainAggregateFileSizeOk returns a tuple with the RetainAggregateFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainAggregateFileSize

`func (o *FileRetentionRecurringTaskShared) SetRetainAggregateFileSize(v string)`

SetRetainAggregateFileSize sets RetainAggregateFileSize field to given value.

### HasRetainAggregateFileSize

`func (o *FileRetentionRecurringTaskShared) HasRetainAggregateFileSize() bool`

HasRetainAggregateFileSize returns a boolean if a field has been set.

### GetDescription

`func (o *FileRetentionRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileRetentionRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileRetentionRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileRetentionRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *FileRetentionRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *FileRetentionRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *FileRetentionRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *FileRetentionRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *FileRetentionRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *FileRetentionRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *FileRetentionRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *FileRetentionRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *FileRetentionRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *FileRetentionRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *FileRetentionRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *FileRetentionRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *FileRetentionRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *FileRetentionRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *FileRetentionRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *FileRetentionRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *FileRetentionRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *FileRetentionRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *FileRetentionRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *FileRetentionRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *FileRetentionRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *FileRetentionRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *FileRetentionRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *FileRetentionRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *FileRetentionRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *FileRetentionRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *FileRetentionRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *FileRetentionRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


