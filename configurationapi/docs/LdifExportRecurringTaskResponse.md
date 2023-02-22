# LdifExportRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Recurring Task | 
**Schemas** | [**[]EnumldifExportRecurringTaskSchemaUrn**](EnumldifExportRecurringTaskSchemaUrn.md) |  | 
**LdifDirectory** | **string** | The directory in which LDIF export files will be placed. The directory must already exist. | 
**BackendID** | Pointer to **[]string** | The backend ID for a backend to be exported. | [optional] 
**ExcludeBackendID** | Pointer to **[]string** | The backend ID for a backend to be excluded from the export. | [optional] 
**Compress** | Pointer to **bool** | Indicates whether to compress the LDIF data as it is exported. | [optional] 
**Encrypt** | Pointer to **bool** | Indicates whether to encrypt the LDIF data as it exported. | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | The ID of an encryption settings definition to use to obtain the LDIF export encryption key. | [optional] 
**Sign** | Pointer to **bool** | Indicates whether to cryptographically sign the exported data, which will make it possible to detect whether the LDIF data has been altered since it was exported. | [optional] 
**RetainPreviousLDIFExportCount** | Pointer to **int32** | The minimum number of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**RetainPreviousLDIFExportAge** | Pointer to **string** | The minimum age of previous LDIF exports that should be preserved after a new export completes successfully. | [optional] 
**MaxMegabytesPerSecond** | Pointer to **int32** | The maximum rate, in megabytes per second, at which LDIF exports should be written. | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task starts running. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnSuccess** | Pointer to **[]string** | The email addresses to which a message should be sent whenever an instance of this Recurring Task completes successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**EmailOnFailure** | Pointer to **[]string** | The email addresses to which a message should be sent if an instance of this Recurring Task fails to complete successfully. If this option is used, then at least one smtp-server must be configured in the global configuration. | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewLdifExportRecurringTaskResponse

`func NewLdifExportRecurringTaskResponse(id string, schemas []EnumldifExportRecurringTaskSchemaUrn, ldifDirectory string, ) *LdifExportRecurringTaskResponse`

NewLdifExportRecurringTaskResponse instantiates a new LdifExportRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdifExportRecurringTaskResponseWithDefaults

`func NewLdifExportRecurringTaskResponseWithDefaults() *LdifExportRecurringTaskResponse`

NewLdifExportRecurringTaskResponseWithDefaults instantiates a new LdifExportRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *LdifExportRecurringTaskResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *LdifExportRecurringTaskResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *LdifExportRecurringTaskResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *LdifExportRecurringTaskResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *LdifExportRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *LdifExportRecurringTaskResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *LdifExportRecurringTaskResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *LdifExportRecurringTaskResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *LdifExportRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdifExportRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdifExportRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *LdifExportRecurringTaskResponse) GetSchemas() []EnumldifExportRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *LdifExportRecurringTaskResponse) GetSchemasOk() (*[]EnumldifExportRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *LdifExportRecurringTaskResponse) SetSchemas(v []EnumldifExportRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLdifDirectory

`func (o *LdifExportRecurringTaskResponse) GetLdifDirectory() string`

GetLdifDirectory returns the LdifDirectory field if non-nil, zero value otherwise.

### GetLdifDirectoryOk

`func (o *LdifExportRecurringTaskResponse) GetLdifDirectoryOk() (*string, bool)`

GetLdifDirectoryOk returns a tuple with the LdifDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdifDirectory

`func (o *LdifExportRecurringTaskResponse) SetLdifDirectory(v string)`

SetLdifDirectory sets LdifDirectory field to given value.


### GetBackendID

`func (o *LdifExportRecurringTaskResponse) GetBackendID() []string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *LdifExportRecurringTaskResponse) GetBackendIDOk() (*[]string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *LdifExportRecurringTaskResponse) SetBackendID(v []string)`

SetBackendID sets BackendID field to given value.

### HasBackendID

`func (o *LdifExportRecurringTaskResponse) HasBackendID() bool`

HasBackendID returns a boolean if a field has been set.

### GetExcludeBackendID

`func (o *LdifExportRecurringTaskResponse) GetExcludeBackendID() []string`

GetExcludeBackendID returns the ExcludeBackendID field if non-nil, zero value otherwise.

### GetExcludeBackendIDOk

`func (o *LdifExportRecurringTaskResponse) GetExcludeBackendIDOk() (*[]string, bool)`

GetExcludeBackendIDOk returns a tuple with the ExcludeBackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBackendID

`func (o *LdifExportRecurringTaskResponse) SetExcludeBackendID(v []string)`

SetExcludeBackendID sets ExcludeBackendID field to given value.

### HasExcludeBackendID

`func (o *LdifExportRecurringTaskResponse) HasExcludeBackendID() bool`

HasExcludeBackendID returns a boolean if a field has been set.

### GetCompress

`func (o *LdifExportRecurringTaskResponse) GetCompress() bool`

GetCompress returns the Compress field if non-nil, zero value otherwise.

### GetCompressOk

`func (o *LdifExportRecurringTaskResponse) GetCompressOk() (*bool, bool)`

GetCompressOk returns a tuple with the Compress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompress

`func (o *LdifExportRecurringTaskResponse) SetCompress(v bool)`

SetCompress sets Compress field to given value.

### HasCompress

`func (o *LdifExportRecurringTaskResponse) HasCompress() bool`

HasCompress returns a boolean if a field has been set.

### GetEncrypt

`func (o *LdifExportRecurringTaskResponse) GetEncrypt() bool`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *LdifExportRecurringTaskResponse) GetEncryptOk() (*bool, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *LdifExportRecurringTaskResponse) SetEncrypt(v bool)`

SetEncrypt sets Encrypt field to given value.

### HasEncrypt

`func (o *LdifExportRecurringTaskResponse) HasEncrypt() bool`

HasEncrypt returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *LdifExportRecurringTaskResponse) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *LdifExportRecurringTaskResponse) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *LdifExportRecurringTaskResponse) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *LdifExportRecurringTaskResponse) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetSign

`func (o *LdifExportRecurringTaskResponse) GetSign() bool`

GetSign returns the Sign field if non-nil, zero value otherwise.

### GetSignOk

`func (o *LdifExportRecurringTaskResponse) GetSignOk() (*bool, bool)`

GetSignOk returns a tuple with the Sign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSign

`func (o *LdifExportRecurringTaskResponse) SetSign(v bool)`

SetSign sets Sign field to given value.

### HasSign

`func (o *LdifExportRecurringTaskResponse) HasSign() bool`

HasSign returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportCount

`func (o *LdifExportRecurringTaskResponse) GetRetainPreviousLDIFExportCount() int32`

GetRetainPreviousLDIFExportCount returns the RetainPreviousLDIFExportCount field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportCountOk

`func (o *LdifExportRecurringTaskResponse) GetRetainPreviousLDIFExportCountOk() (*int32, bool)`

GetRetainPreviousLDIFExportCountOk returns a tuple with the RetainPreviousLDIFExportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportCount

`func (o *LdifExportRecurringTaskResponse) SetRetainPreviousLDIFExportCount(v int32)`

SetRetainPreviousLDIFExportCount sets RetainPreviousLDIFExportCount field to given value.

### HasRetainPreviousLDIFExportCount

`func (o *LdifExportRecurringTaskResponse) HasRetainPreviousLDIFExportCount() bool`

HasRetainPreviousLDIFExportCount returns a boolean if a field has been set.

### GetRetainPreviousLDIFExportAge

`func (o *LdifExportRecurringTaskResponse) GetRetainPreviousLDIFExportAge() string`

GetRetainPreviousLDIFExportAge returns the RetainPreviousLDIFExportAge field if non-nil, zero value otherwise.

### GetRetainPreviousLDIFExportAgeOk

`func (o *LdifExportRecurringTaskResponse) GetRetainPreviousLDIFExportAgeOk() (*string, bool)`

GetRetainPreviousLDIFExportAgeOk returns a tuple with the RetainPreviousLDIFExportAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainPreviousLDIFExportAge

`func (o *LdifExportRecurringTaskResponse) SetRetainPreviousLDIFExportAge(v string)`

SetRetainPreviousLDIFExportAge sets RetainPreviousLDIFExportAge field to given value.

### HasRetainPreviousLDIFExportAge

`func (o *LdifExportRecurringTaskResponse) HasRetainPreviousLDIFExportAge() bool`

HasRetainPreviousLDIFExportAge returns a boolean if a field has been set.

### GetMaxMegabytesPerSecond

`func (o *LdifExportRecurringTaskResponse) GetMaxMegabytesPerSecond() int32`

GetMaxMegabytesPerSecond returns the MaxMegabytesPerSecond field if non-nil, zero value otherwise.

### GetMaxMegabytesPerSecondOk

`func (o *LdifExportRecurringTaskResponse) GetMaxMegabytesPerSecondOk() (*int32, bool)`

GetMaxMegabytesPerSecondOk returns a tuple with the MaxMegabytesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMegabytesPerSecond

`func (o *LdifExportRecurringTaskResponse) SetMaxMegabytesPerSecond(v int32)`

SetMaxMegabytesPerSecond sets MaxMegabytesPerSecond field to given value.

### HasMaxMegabytesPerSecond

`func (o *LdifExportRecurringTaskResponse) HasMaxMegabytesPerSecond() bool`

HasMaxMegabytesPerSecond returns a boolean if a field has been set.

### GetDescription

`func (o *LdifExportRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LdifExportRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LdifExportRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LdifExportRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *LdifExportRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *LdifExportRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *LdifExportRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *LdifExportRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *LdifExportRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *LdifExportRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *LdifExportRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *LdifExportRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *LdifExportRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *LdifExportRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *LdifExportRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *LdifExportRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *LdifExportRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *LdifExportRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *LdifExportRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *LdifExportRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *LdifExportRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *LdifExportRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *LdifExportRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *LdifExportRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *LdifExportRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *LdifExportRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *LdifExportRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *LdifExportRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *LdifExportRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *LdifExportRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *LdifExportRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *LdifExportRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


