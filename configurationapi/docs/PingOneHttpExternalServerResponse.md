# PingOneHttpExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the External Server | 
**Schemas** | [**[]EnumpingOneHttpExternalServerSchemaUrn**](EnumpingOneHttpExternalServerSchemaUrn.md) |  | 
**HostnameVerificationMethod** | Pointer to [**EnumexternalServerHostnameVerificationMethodProp**](EnumexternalServerHostnameVerificationMethodProp.md) |  | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use for HTTPS connection-level security. | [optional] 
**ConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established before aborting a request to PingOne. | [optional] 
**ResponseTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for response data to be read from an established connection before aborting a request to PingOne. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewPingOneHttpExternalServerResponse

`func NewPingOneHttpExternalServerResponse(id string, schemas []EnumpingOneHttpExternalServerSchemaUrn, ) *PingOneHttpExternalServerResponse`

NewPingOneHttpExternalServerResponse instantiates a new PingOneHttpExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneHttpExternalServerResponseWithDefaults

`func NewPingOneHttpExternalServerResponseWithDefaults() *PingOneHttpExternalServerResponse`

NewPingOneHttpExternalServerResponseWithDefaults instantiates a new PingOneHttpExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *PingOneHttpExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PingOneHttpExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PingOneHttpExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PingOneHttpExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PingOneHttpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PingOneHttpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PingOneHttpExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PingOneHttpExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *PingOneHttpExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOneHttpExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOneHttpExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PingOneHttpExternalServerResponse) GetSchemas() []EnumpingOneHttpExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PingOneHttpExternalServerResponse) GetSchemasOk() (*[]EnumpingOneHttpExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PingOneHttpExternalServerResponse) SetSchemas(v []EnumpingOneHttpExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetHostnameVerificationMethod

`func (o *PingOneHttpExternalServerResponse) GetHostnameVerificationMethod() EnumexternalServerHostnameVerificationMethodProp`

GetHostnameVerificationMethod returns the HostnameVerificationMethod field if non-nil, zero value otherwise.

### GetHostnameVerificationMethodOk

`func (o *PingOneHttpExternalServerResponse) GetHostnameVerificationMethodOk() (*EnumexternalServerHostnameVerificationMethodProp, bool)`

GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationMethod

`func (o *PingOneHttpExternalServerResponse) SetHostnameVerificationMethod(v EnumexternalServerHostnameVerificationMethodProp)`

SetHostnameVerificationMethod sets HostnameVerificationMethod field to given value.

### HasHostnameVerificationMethod

`func (o *PingOneHttpExternalServerResponse) HasHostnameVerificationMethod() bool`

HasHostnameVerificationMethod returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *PingOneHttpExternalServerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *PingOneHttpExternalServerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *PingOneHttpExternalServerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *PingOneHttpExternalServerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *PingOneHttpExternalServerResponse) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PingOneHttpExternalServerResponse) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PingOneHttpExternalServerResponse) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PingOneHttpExternalServerResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *PingOneHttpExternalServerResponse) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *PingOneHttpExternalServerResponse) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *PingOneHttpExternalServerResponse) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *PingOneHttpExternalServerResponse) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *PingOneHttpExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingOneHttpExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingOneHttpExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingOneHttpExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


