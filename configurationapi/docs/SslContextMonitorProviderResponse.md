# SslContextMonitorProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumsslContextMonitorProviderSchemaUrn**](EnumsslContextMonitorProviderSchemaUrn.md) |  | 
**Id** | **string** | Name of the Monitor Provider | 
**Description** | Pointer to **string** | A description for this Monitor Provider | [optional] 
**Enabled** | **bool** | Indicates whether the Monitor Provider is enabled for use. | 

## Methods

### NewSslContextMonitorProviderResponse

`func NewSslContextMonitorProviderResponse(schemas []EnumsslContextMonitorProviderSchemaUrn, id string, enabled bool, ) *SslContextMonitorProviderResponse`

NewSslContextMonitorProviderResponse instantiates a new SslContextMonitorProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslContextMonitorProviderResponseWithDefaults

`func NewSslContextMonitorProviderResponseWithDefaults() *SslContextMonitorProviderResponse`

NewSslContextMonitorProviderResponseWithDefaults instantiates a new SslContextMonitorProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SslContextMonitorProviderResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SslContextMonitorProviderResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SslContextMonitorProviderResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SslContextMonitorProviderResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SslContextMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SslContextMonitorProviderResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SslContextMonitorProviderResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SslContextMonitorProviderResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *SslContextMonitorProviderResponse) GetSchemas() []EnumsslContextMonitorProviderSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SslContextMonitorProviderResponse) GetSchemasOk() (*[]EnumsslContextMonitorProviderSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SslContextMonitorProviderResponse) SetSchemas(v []EnumsslContextMonitorProviderSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SslContextMonitorProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SslContextMonitorProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SslContextMonitorProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *SslContextMonitorProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SslContextMonitorProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SslContextMonitorProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SslContextMonitorProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SslContextMonitorProviderResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SslContextMonitorProviderResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SslContextMonitorProviderResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


