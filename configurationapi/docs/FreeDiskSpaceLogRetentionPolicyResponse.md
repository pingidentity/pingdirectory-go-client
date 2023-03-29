# FreeDiskSpaceLogRetentionPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Log Retention Policy | 
**Schemas** | [**[]EnumfreeDiskSpaceLogRetentionPolicySchemaUrn**](EnumfreeDiskSpaceLogRetentionPolicySchemaUrn.md) |  | 
**FreeDiskSpace** | **string** | Specifies the minimum amount of free disk space that should be available on the file system on which the archived log files are stored. | 
**Description** | Pointer to **string** | A description for this Log Retention Policy | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewFreeDiskSpaceLogRetentionPolicyResponse

`func NewFreeDiskSpaceLogRetentionPolicyResponse(id string, schemas []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn, freeDiskSpace string, ) *FreeDiskSpaceLogRetentionPolicyResponse`

NewFreeDiskSpaceLogRetentionPolicyResponse instantiates a new FreeDiskSpaceLogRetentionPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeDiskSpaceLogRetentionPolicyResponseWithDefaults

`func NewFreeDiskSpaceLogRetentionPolicyResponseWithDefaults() *FreeDiskSpaceLogRetentionPolicyResponse`

NewFreeDiskSpaceLogRetentionPolicyResponseWithDefaults instantiates a new FreeDiskSpaceLogRetentionPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetSchemas() []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetSchemasOk() (*[]EnumfreeDiskSpaceLogRetentionPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) SetSchemas(v []EnumfreeDiskSpaceLogRetentionPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFreeDiskSpace

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetFreeDiskSpace() string`

GetFreeDiskSpace returns the FreeDiskSpace field if non-nil, zero value otherwise.

### GetFreeDiskSpaceOk

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetFreeDiskSpaceOk() (*string, bool)`

GetFreeDiskSpaceOk returns a tuple with the FreeDiskSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeDiskSpace

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) SetFreeDiskSpace(v string)`

SetFreeDiskSpace sets FreeDiskSpace field to given value.


### GetDescription

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *FreeDiskSpaceLogRetentionPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


