# AddLogRetentionPolicy200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Retention Policy | 
**Schemas** | [**[]EnumsizeLimitLogRetentionPolicySchemaUrn**](EnumsizeLimitLogRetentionPolicySchemaUrn.md) |  | 
**RetainDuration** | **string** | Specifies the desired minimum length of time that each log file should be retained. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**NumberOfFiles** | **int64** | Specifies the number of archived log files to retain before the oldest ones are cleaned. | 
**FreeDiskSpace** | **string** | Specifies the minimum amount of free disk space that should be available on the file system on which the archived log files are stored. | 
**DiskSpaceUsed** | **string** | Specifies the maximum total disk space used by the log files. | 

## Methods

### NewAddLogRetentionPolicy200Response

`func NewAddLogRetentionPolicy200Response(id string, schemas []EnumsizeLimitLogRetentionPolicySchemaUrn, retainDuration string, numberOfFiles int64, freeDiskSpace string, diskSpaceUsed string, ) *AddLogRetentionPolicy200Response`

NewAddLogRetentionPolicy200Response instantiates a new AddLogRetentionPolicy200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogRetentionPolicy200ResponseWithDefaults

`func NewAddLogRetentionPolicy200ResponseWithDefaults() *AddLogRetentionPolicy200Response`

NewAddLogRetentionPolicy200ResponseWithDefaults instantiates a new AddLogRetentionPolicy200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddLogRetentionPolicy200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddLogRetentionPolicy200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddLogRetentionPolicy200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddLogRetentionPolicy200Response) GetSchemas() []EnumsizeLimitLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogRetentionPolicy200Response) GetSchemasOk() (*[]EnumsizeLimitLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogRetentionPolicy200Response) SetSchemas(v []EnumsizeLimitLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRetainDuration

`func (o *AddLogRetentionPolicy200Response) GetRetainDuration() string`

GetRetainDuration returns the RetainDuration field if non-nil, zero value otherwise.

### GetRetainDurationOk

`func (o *AddLogRetentionPolicy200Response) GetRetainDurationOk() (*string, bool)`

GetRetainDurationOk returns a tuple with the RetainDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainDuration

`func (o *AddLogRetentionPolicy200Response) SetRetainDuration(v string)`

SetRetainDuration sets RetainDuration field to given value.


### GetDescription

`func (o *AddLogRetentionPolicy200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogRetentionPolicy200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogRetentionPolicy200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogRetentionPolicy200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *AddLogRetentionPolicy200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddLogRetentionPolicy200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddLogRetentionPolicy200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddLogRetentionPolicy200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogRetentionPolicy200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddLogRetentionPolicy200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogRetentionPolicy200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddLogRetentionPolicy200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetNumberOfFiles

`func (o *AddLogRetentionPolicy200Response) GetNumberOfFiles() int64`

GetNumberOfFiles returns the NumberOfFiles field if non-nil, zero value otherwise.

### GetNumberOfFilesOk

`func (o *AddLogRetentionPolicy200Response) GetNumberOfFilesOk() (*int64, bool)`

GetNumberOfFilesOk returns a tuple with the NumberOfFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFiles

`func (o *AddLogRetentionPolicy200Response) SetNumberOfFiles(v int64)`

SetNumberOfFiles sets NumberOfFiles field to given value.


### GetFreeDiskSpace

`func (o *AddLogRetentionPolicy200Response) GetFreeDiskSpace() string`

GetFreeDiskSpace returns the FreeDiskSpace field if non-nil, zero value otherwise.

### GetFreeDiskSpaceOk

`func (o *AddLogRetentionPolicy200Response) GetFreeDiskSpaceOk() (*string, bool)`

GetFreeDiskSpaceOk returns a tuple with the FreeDiskSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeDiskSpace

`func (o *AddLogRetentionPolicy200Response) SetFreeDiskSpace(v string)`

SetFreeDiskSpace sets FreeDiskSpace field to given value.


### GetDiskSpaceUsed

`func (o *AddLogRetentionPolicy200Response) GetDiskSpaceUsed() string`

GetDiskSpaceUsed returns the DiskSpaceUsed field if non-nil, zero value otherwise.

### GetDiskSpaceUsedOk

`func (o *AddLogRetentionPolicy200Response) GetDiskSpaceUsedOk() (*string, bool)`

GetDiskSpaceUsedOk returns a tuple with the DiskSpaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceUsed

`func (o *AddLogRetentionPolicy200Response) SetDiskSpaceUsed(v string)`

SetDiskSpaceUsed sets DiskSpaceUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


