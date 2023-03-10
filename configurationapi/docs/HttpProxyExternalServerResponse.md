# HttpProxyExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the External Server | 
**Schemas** | [**[]EnumhttpProxyExternalServerSchemaUrn**](EnumhttpProxyExternalServerSchemaUrn.md) |  | 
**ServerHostName** | **string** | The host name or IP address of the HTTP Proxy External Server. | 
**ServerPort** | **int32** | The port on which the HTTP Proxy External Server is listening for connections. | 
**BasicAuthenticationUsername** | Pointer to **string** | The username to use to authenticate to the HTTP Proxy External Server. | [optional] 
**BasicAuthenticationPassphraseProvider** | Pointer to **string** | A passphrase provider that provides access to the password to use to authenticate to the HTTP Proxy External Server. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewHttpProxyExternalServerResponse

`func NewHttpProxyExternalServerResponse(id string, schemas []EnumhttpProxyExternalServerSchemaUrn, serverHostName string, serverPort int32, ) *HttpProxyExternalServerResponse`

NewHttpProxyExternalServerResponse instantiates a new HttpProxyExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpProxyExternalServerResponseWithDefaults

`func NewHttpProxyExternalServerResponseWithDefaults() *HttpProxyExternalServerResponse`

NewHttpProxyExternalServerResponseWithDefaults instantiates a new HttpProxyExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *HttpProxyExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HttpProxyExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HttpProxyExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HttpProxyExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpProxyExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *HttpProxyExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpProxyExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *HttpProxyExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *HttpProxyExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpProxyExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpProxyExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *HttpProxyExternalServerResponse) GetSchemas() []EnumhttpProxyExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HttpProxyExternalServerResponse) GetSchemasOk() (*[]EnumhttpProxyExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HttpProxyExternalServerResponse) SetSchemas(v []EnumhttpProxyExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *HttpProxyExternalServerResponse) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *HttpProxyExternalServerResponse) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *HttpProxyExternalServerResponse) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *HttpProxyExternalServerResponse) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *HttpProxyExternalServerResponse) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *HttpProxyExternalServerResponse) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.


### GetBasicAuthenticationUsername

`func (o *HttpProxyExternalServerResponse) GetBasicAuthenticationUsername() string`

GetBasicAuthenticationUsername returns the BasicAuthenticationUsername field if non-nil, zero value otherwise.

### GetBasicAuthenticationUsernameOk

`func (o *HttpProxyExternalServerResponse) GetBasicAuthenticationUsernameOk() (*string, bool)`

GetBasicAuthenticationUsernameOk returns a tuple with the BasicAuthenticationUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthenticationUsername

`func (o *HttpProxyExternalServerResponse) SetBasicAuthenticationUsername(v string)`

SetBasicAuthenticationUsername sets BasicAuthenticationUsername field to given value.

### HasBasicAuthenticationUsername

`func (o *HttpProxyExternalServerResponse) HasBasicAuthenticationUsername() bool`

HasBasicAuthenticationUsername returns a boolean if a field has been set.

### GetBasicAuthenticationPassphraseProvider

`func (o *HttpProxyExternalServerResponse) GetBasicAuthenticationPassphraseProvider() string`

GetBasicAuthenticationPassphraseProvider returns the BasicAuthenticationPassphraseProvider field if non-nil, zero value otherwise.

### GetBasicAuthenticationPassphraseProviderOk

`func (o *HttpProxyExternalServerResponse) GetBasicAuthenticationPassphraseProviderOk() (*string, bool)`

GetBasicAuthenticationPassphraseProviderOk returns a tuple with the BasicAuthenticationPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthenticationPassphraseProvider

`func (o *HttpProxyExternalServerResponse) SetBasicAuthenticationPassphraseProvider(v string)`

SetBasicAuthenticationPassphraseProvider sets BasicAuthenticationPassphraseProvider field to given value.

### HasBasicAuthenticationPassphraseProvider

`func (o *HttpProxyExternalServerResponse) HasBasicAuthenticationPassphraseProvider() bool`

HasBasicAuthenticationPassphraseProvider returns a boolean if a field has been set.

### GetDescription

`func (o *HttpProxyExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpProxyExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpProxyExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpProxyExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


