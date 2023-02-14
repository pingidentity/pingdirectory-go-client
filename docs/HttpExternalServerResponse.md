# HttpExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the External Server | 
**Schemas** | [**[]EnumhttpExternalServerSchemaUrn**](EnumhttpExternalServerSchemaUrn.md) |  | 
**BaseURL** | **string** | The base URL of the external server, optionally including port number, for example \&quot;https://externalService:9031\&quot;. | 
**HostnameVerificationMethod** | Pointer to [**EnumexternalServerHostnameVerificationMethodProp**](EnumexternalServerHostnameVerificationMethodProp.md) |  | [optional] 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if SSL (HTTPS) is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use if SSL (HTTPS) is to be used for connection-level security. | [optional] 
**SslCertNickname** | Pointer to **string** | The certificate alias within the keystore to use if SSL (HTTPS) is to be used for connection-level security. When specifying a value for this property you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**ConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established before aborting a request to the server. | [optional] 
**ResponseTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for response data to be read from an established connection before aborting a request to the server. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewHttpExternalServerResponse

`func NewHttpExternalServerResponse(id string, schemas []EnumhttpExternalServerSchemaUrn, baseURL string, ) *HttpExternalServerResponse`

NewHttpExternalServerResponse instantiates a new HttpExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpExternalServerResponseWithDefaults

`func NewHttpExternalServerResponseWithDefaults() *HttpExternalServerResponse`

NewHttpExternalServerResponseWithDefaults instantiates a new HttpExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *HttpExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HttpExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HttpExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HttpExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *HttpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *HttpExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *HttpExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *HttpExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HttpExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HttpExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *HttpExternalServerResponse) GetSchemas() []EnumhttpExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HttpExternalServerResponse) GetSchemasOk() (*[]EnumhttpExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HttpExternalServerResponse) SetSchemas(v []EnumhttpExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetBaseURL

`func (o *HttpExternalServerResponse) GetBaseURL() string`

GetBaseURL returns the BaseURL field if non-nil, zero value otherwise.

### GetBaseURLOk

`func (o *HttpExternalServerResponse) GetBaseURLOk() (*string, bool)`

GetBaseURLOk returns a tuple with the BaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURL

`func (o *HttpExternalServerResponse) SetBaseURL(v string)`

SetBaseURL sets BaseURL field to given value.


### GetHostnameVerificationMethod

`func (o *HttpExternalServerResponse) GetHostnameVerificationMethod() EnumexternalServerHostnameVerificationMethodProp`

GetHostnameVerificationMethod returns the HostnameVerificationMethod field if non-nil, zero value otherwise.

### GetHostnameVerificationMethodOk

`func (o *HttpExternalServerResponse) GetHostnameVerificationMethodOk() (*EnumexternalServerHostnameVerificationMethodProp, bool)`

GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationMethod

`func (o *HttpExternalServerResponse) SetHostnameVerificationMethod(v EnumexternalServerHostnameVerificationMethodProp)`

SetHostnameVerificationMethod sets HostnameVerificationMethod field to given value.

### HasHostnameVerificationMethod

`func (o *HttpExternalServerResponse) HasHostnameVerificationMethod() bool`

HasHostnameVerificationMethod returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *HttpExternalServerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *HttpExternalServerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *HttpExternalServerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *HttpExternalServerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *HttpExternalServerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *HttpExternalServerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *HttpExternalServerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *HttpExternalServerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *HttpExternalServerResponse) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *HttpExternalServerResponse) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *HttpExternalServerResponse) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *HttpExternalServerResponse) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *HttpExternalServerResponse) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *HttpExternalServerResponse) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *HttpExternalServerResponse) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *HttpExternalServerResponse) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *HttpExternalServerResponse) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *HttpExternalServerResponse) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *HttpExternalServerResponse) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *HttpExternalServerResponse) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetDescription

`func (o *HttpExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HttpExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HttpExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HttpExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


