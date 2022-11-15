# LdifExportRecurringTaskShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumldifExportRecurringTaskSchemaUrn**](EnumldifExportRecurringTaskSchemaUrn.md) |  | 
**LdifDirectory** | **string** | The directory in which LDIF export files will be placed. The directory must already exist. | 
**BackendID** | Pointer to **[]string** |  | [optional] 
**ExcludeBackendID** | Pointer to **[]string** |  | [optional] 
**Compress** | Pointer to **bool** | Indicates whether to compress the LDIF data as it is exported. | [optional] 
**Encrypt** | Pointer to **bool** | Indicates whether to encrypt the LDIF data as it exported. | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | The ID of an encryption settings definition to use to obtain the LDIF export encryption key. | [optional] 
**Sign** | Pointer to **bool** | Indicates whether to cryptographically sign the exported data, which will make it possible to detect whether the LDIF data has been altered since it was exported. | [optional] 
**RetainPreviousLDIFExportCount** | Pointer to **int32** | The minimum number of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**RetainPreviousLDIFExportAge** | Pointer to **string** | The minimum age of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**MaxMegabytesPerSecond** | Pointer to **int32** | The maximum rate, in megabytes per second, at which LDIF exports should be written. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewLdifExportRecurringTaskShared

`func NewLdifExportRecurringTaskShared(schemas []EnumldifExportRecurringTaskSchemaUrn, ldifDirectory string, ) *LdifExportRecurringTaskShared`

NewLdifExportRecurringTaskShared instantiates a new LdifExportRecurringTaskShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdifExportRecurringTaskSharedWithDefaults

`func NewLdifExportRecurringTaskSharedWithDefaults() *LdifExportRecurringTaskShared`

NewLdifExportRecurringTaskSharedWithDefaults instantiates a new LdifExportRecurringTaskShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *LdifExportRecurringTaskShared) GetSchemas() []EnumldifExportRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdifExportRecurringTaskShared) GetSchemasOk() (*[]EnumldifExportRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdifExportRecurringTaskShared) SetSchemas(v []EnumldifExportRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLdifDirectory

`func (o *LdifExportRecurringTaskShared) GetLdifDirectory() string`

GetLdifDirectory returns the LdifDirectory field if non-nil, zero value otherwise.

### GetLdifDirectoryOk

`func (o *LdifExportRecurringTaskShared) GetLdifDirectoryOk() (*string, bool)`

GetLdifDirectoryOk returns a tuple with the LdifDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifDirectory

`func (o *LdifExportRecurringTaskShared) SetLdifDirectory(v string)`

SetLdifDirectory sets LdifDirectory field to given value.


### GetBackendID

`func (o *LdifExportRecurringTaskShared) GetBackendID() []string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *LdifExportRecurringTaskShared) GetBackendIDOk() (*[]string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *LdifExportRecurringTaskShared) SetBackendID(v []string)`

SetBackendID sets BackendID field to given value.

### HasBackendID

`func (o *LdifExportRecurringTaskShared) HasBackendID() bool`

HasBackendID returns a boolean if a field has been set.

### GetExcludeBackendID

`func (o *LdifExportRecurringTaskShared) GetExcludeBackendID() []string`

GetExcludeBackendID returns the ExcludeBackendID field if non-nil, zero value otherwise.

### GetExcludeBackendIDOk

`func (o *LdifExportRecurringTaskShared) GetExcludeBackendIDOk() (*[]string, bool)`

GetExcludeBackendIDOk returns a tuple with the ExcludeBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBackendID

`func (o *LdifExportRecurringTaskShared) SetExcludeBackendID(v []string)`

SetExcludeBackendID sets ExcludeBackendID field to given value.

### HasExcludeBackendID

`func (o *LdifExportRecurringTaskShared) HasExcludeBackendID() bool`

HasExcludeBackendID returns a boolean if a field has been set.

### GetCompress

`func (o *LdifExportRecurringTaskShared) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *LdifExportRecurringTaskShared) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *LdifExportRecurringTaskShared) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *LdifExportRecurringTaskShared) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncrypt

`func (o *LdifExportRecurringTaskShared) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *LdifExportRecurringTaskShared) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *LdifExportRecurringTaskShared) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *LdifExportRecurringTaskShared) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *LdifExportRecurringTaskShared) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *LdifExportRecurringTaskShared) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *LdifExportRecurringTaskShared) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *LdifExportRecurringTaskShared) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetSign

`func (o *LdifExportRecurringTaskShared) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *LdifExportRecurringTaskShared) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *LdifExportRecurringTaskShared) SetSign(v bool)`

SetSign sets Sign field to given value.

### HasSign

`func (o *LdifExportRecurringTaskShared) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportCount

`func (o *LdifExportRecurringTaskShared) GetRetainPreviousLDIFExportCount() int32`

GetRetainPreviousLDIFExportCount returns the RetainPreviousLDIFExportCount field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportCountOk

`func (o *LdifExportRecurringTaskShared) GetRetainPreviousLDIFExportCountOk() (*int32, bool)`

GetRetainPreviousLDIFExportCountOk returns a tuple with the RetainPreviousLDIFExportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportCount

`func (o *LdifExportRecurringTaskShared) SetRetainPreviousLDIFExportCount(v int32)`

SetRetainPreviousLDIFExportCount sets RetainPreviousLDIFExportCount field to given value.

### HasRetainPreviousLDIFExportCount

`func (o *LdifExportRecurringTaskShared) HasRetainPreviousLDIFExportCount() bool`

HasRetainPreviousLDIFExportCount returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportAge

`func (o *LdifExportRecurringTaskShared) GetRetainPreviousLDIFExportAge() string`

GetRetainPreviousLDIFExportAge returns the RetainPreviousLDIFExportAge field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportAgeOk

`func (o *LdifExportRecurringTaskShared) GetRetainPreviousLDIFExportAgeOk() (*string, bool)`

GetRetainPreviousLDIFExportAgeOk returns a tuple with the RetainPreviousLDIFExportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportAge

`func (o *LdifExportRecurringTaskShared) SetRetainPreviousLDIFExportAge(v string)`

SetRetainPreviousLDIFExportAge sets RetainPreviousLDIFExportAge field to given value.

### HasRetainPreviousLDIFExportAge

`func (o *LdifExportRecurringTaskShared) HasRetainPreviousLDIFExportAge() bool`

HasRetainPreviousLDIFExportAge returns a boolean if a field has been set.

### GetMaxMegabytesPerSecond

`func (o *LdifExportRecurringTaskShared) GetMaxMegabytesPerSecond() int32`

GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field if non-nil, zero value otherwise.

### GetMaxMegabytesPerSecondOk

`func (o *LdifExportRecurringTaskShared) GetMaxMegabytesPerSecondOk() (*int32, bool)`

GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMegabytesPerSecond

`func (o *LdifExportRecurringTaskShared) SetMaxMegabytesPerSecond(v int32)`

SetMaxMegabytesPerSecond sets MaxMegabytesPerSecond field to given value.

### HasMaxMegabytesPerSecond

`func (o *LdifExportRecurringTaskShared) HasMaxMegabytesPerSecond() bool`

HasMaxMegabytesPerSecond returns a boolean if a field has been set.

### GetDescription

`func (o *LdifExportRecurringTaskShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdifExportRecurringTaskShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdifExportRecurringTaskShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdifExportRecurringTaskShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *LdifExportRecurringTaskShared) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *LdifExportRecurringTaskShared) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *LdifExportRecurringTaskShared) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *LdifExportRecurringTaskShared) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *LdifExportRecurringTaskShared) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *LdifExportRecurringTaskShared) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *LdifExportRecurringTaskShared) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *LdifExportRecurringTaskShared) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *LdifExportRecurringTaskShared) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *LdifExportRecurringTaskShared) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *LdifExportRecurringTaskShared) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *LdifExportRecurringTaskShared) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *LdifExportRecurringTaskShared) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *LdifExportRecurringTaskShared) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *LdifExportRecurringTaskShared) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *LdifExportRecurringTaskShared) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *LdifExportRecurringTaskShared) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *LdifExportRecurringTaskShared) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *LdifExportRecurringTaskShared) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *LdifExportRecurringTaskShared) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *LdifExportRecurringTaskShared) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *LdifExportRecurringTaskShared) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *LdifExportRecurringTaskShared) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *LdifExportRecurringTaskShared) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *LdifExportRecurringTaskShared) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *LdifExportRecurringTaskShared) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *LdifExportRecurringTaskShared) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *LdifExportRecurringTaskShared) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


