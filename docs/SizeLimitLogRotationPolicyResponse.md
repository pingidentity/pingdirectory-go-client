# SizeLimitLogRotationPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Log Rotation Policy | 
**Schemas** | [**[]EnumsizeLimitLogRotationPolicySchemaUrn**](EnumsizeLimitLogRotationPolicySchemaUrn.md) |  | 
**FileSizeLimit** | **string** | Specifies the maximum size that a log file can reach before it is rotated. | 
**Description** | Pointer to **string** | A description for this Log Rotation Policy | [optional] 

## Methods

### NewSizeLimitLogRotationPolicyResponse

`func NewSizeLimitLogRotationPolicyResponse(id string, schemas []EnumsizeLimitLogRotationPolicySchemaUrn, fileSizeLimit string, ) *SizeLimitLogRotationPolicyResponse`

NewSizeLimitLogRotationPolicyResponse instantiates a new SizeLimitLogRotationPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeLimitLogRotationPolicyResponseWithDefaults

`func NewSizeLimitLogRotationPolicyResponseWithDefaults() *SizeLimitLogRotationPolicyResponse`

NewSizeLimitLogRotationPolicyResponseWithDefaults instantiates a new SizeLimitLogRotationPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SizeLimitLogRotationPolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SizeLimitLogRotationPolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SizeLimitLogRotationPolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SizeLimitLogRotationPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SizeLimitLogRotationPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SizeLimitLogRotationPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SizeLimitLogRotationPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SizeLimitLogRotationPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *SizeLimitLogRotationPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SizeLimitLogRotationPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SizeLimitLogRotationPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SizeLimitLogRotationPolicyResponse) GetSchemas() []EnumsizeLimitLogRotationPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SizeLimitLogRotationPolicyResponse) GetSchemasOk() (*[]EnumsizeLimitLogRotationPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SizeLimitLogRotationPolicyResponse) SetSchemas(v []EnumsizeLimitLogRotationPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetFileSizeLimit

`func (o *SizeLimitLogRotationPolicyResponse) GetFileSizeLimit() string`

GetFileSizeLimit returns the FileSizeLimit field if non-nil, zero value otherwise.

### GetFileSizeLimitOk

`func (o *SizeLimitLogRotationPolicyResponse) GetFileSizeLimitOk() (*string, bool)`

GetFileSizeLimitOk returns a tuple with the FileSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeLimit

`func (o *SizeLimitLogRotationPolicyResponse) SetFileSizeLimit(v string)`

SetFileSizeLimit sets FileSizeLimit field to given value.


### GetDescription

`func (o *SizeLimitLogRotationPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SizeLimitLogRotationPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SizeLimitLogRotationPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SizeLimitLogRotationPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


