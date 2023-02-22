# SaltedSha512PasswordStorageSchemeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumsaltedSha512PasswordStorageSchemeSchemaUrn**](EnumsaltedSha512PasswordStorageSchemeSchemaUrn.md) |  | 
**Id** | **string** | Name of the Password Storage Scheme | 
**SaltLengthBytes** | Pointer to **int32** | Specifies the number of bytes to use for the generated salt. | [optional] 
**Description** | Pointer to **string** | A description for this Password Storage Scheme | [optional] 
**Enabled** | **bool** | Indicates whether the Password Storage Scheme is enabled for use. | 

## Methods

### NewSaltedSha512PasswordStorageSchemeResponse

`func NewSaltedSha512PasswordStorageSchemeResponse(schemas []EnumsaltedSha512PasswordStorageSchemeSchemaUrn, id string, enabled bool, ) *SaltedSha512PasswordStorageSchemeResponse`

NewSaltedSha512PasswordStorageSchemeResponse instantiates a new SaltedSha512PasswordStorageSchemeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaltedSha512PasswordStorageSchemeResponseWithDefaults

`func NewSaltedSha512PasswordStorageSchemeResponseWithDefaults() *SaltedSha512PasswordStorageSchemeResponse`

NewSaltedSha512PasswordStorageSchemeResponseWithDefaults instantiates a new SaltedSha512PasswordStorageSchemeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SaltedSha512PasswordStorageSchemeResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SaltedSha512PasswordStorageSchemeResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SaltedSha512PasswordStorageSchemeResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SaltedSha512PasswordStorageSchemeResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetSchemas() []EnumsaltedSha512PasswordStorageSchemeSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetSchemasOk() (*[]EnumsaltedSha512PasswordStorageSchemeSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SaltedSha512PasswordStorageSchemeResponse) SetSchemas(v []EnumsaltedSha512PasswordStorageSchemeSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SaltedSha512PasswordStorageSchemeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSaltLengthBytes

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetSaltLengthBytes() int32`

GetSaltLengthBytes returns the SaltLengthBytes field if non-nil, zero value otherwise.

### GetSaltLengthBytesOk

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetSaltLengthBytesOk() (*int32, bool)`

GetSaltLengthBytesOk returns a tuple with the SaltLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaltLengthBytes

`func (o *SaltedSha512PasswordStorageSchemeResponse) SetSaltLengthBytes(v int32)`

SetSaltLengthBytes sets SaltLengthBytes field to given value.

### HasSaltLengthBytes

`func (o *SaltedSha512PasswordStorageSchemeResponse) HasSaltLengthBytes() bool`

HasSaltLengthBytes returns a boolean if a field has been set.

### GetDescription

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaltedSha512PasswordStorageSchemeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SaltedSha512PasswordStorageSchemeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SaltedSha512PasswordStorageSchemeResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SaltedSha512PasswordStorageSchemeResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


