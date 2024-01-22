# FileCountLogRetentionPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumfileCountLogRetentionPolicySchemaUrn**](EnumfileCountLogRetentionPolicySchemaUrn.md) |  | 
**NumberOfFiles** | **int64** | Specifies the number of archived log files to retain before the oldest ones are cleaned. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Log Retention Policy | 

## Methods

### NewFileCountLogRetentionPolicyResponse

`func NewFileCountLogRetentionPolicyResponse(schemas []EnumfileCountLogRetentionPolicySchemaUrn, numberOfFiles int64, id string, ) *FileCountLogRetentionPolicyResponse`

NewFileCountLogRetentionPolicyResponse instantiates a new FileCountLogRetentionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileCountLogRetentionPolicyResponseWithDefaults

`func NewFileCountLogRetentionPolicyResponseWithDefaults() *FileCountLogRetentionPolicyResponse`

NewFileCountLogRetentionPolicyResponseWithDefaults instantiates a new FileCountLogRetentionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *FileCountLogRetentionPolicyResponse) GetSchemas() []EnumfileCountLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FileCountLogRetentionPolicyResponse) GetSchemasOk() (*[]EnumfileCountLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FileCountLogRetentionPolicyResponse) SetSchemas(v []EnumfileCountLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetNumberOfFiles

`func (o *FileCountLogRetentionPolicyResponse) GetNumberOfFiles() int64`

GetNumberOfFiles returns the NumberOfFiles field if non-nil, zero value otherwise.

### GetNumberOfFilesOk

`func (o *FileCountLogRetentionPolicyResponse) GetNumberOfFilesOk() (*int64, bool)`

GetNumberOfFilesOk returns a tuple with the NumberOfFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFiles

`func (o *FileCountLogRetentionPolicyResponse) SetNumberOfFiles(v int64)`

SetNumberOfFiles sets NumberOfFiles field to given value.


### GetDescription

`func (o *FileCountLogRetentionPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FileCountLogRetentionPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FileCountLogRetentionPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FileCountLogRetentionPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *FileCountLogRetentionPolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FileCountLogRetentionPolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FileCountLogRetentionPolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FileCountLogRetentionPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FileCountLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FileCountLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FileCountLogRetentionPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FileCountLogRetentionPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *FileCountLogRetentionPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileCountLogRetentionPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileCountLogRetentionPolicyResponse) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


