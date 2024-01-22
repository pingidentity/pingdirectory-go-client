# Scim2ExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]Enumscim2ExternalServerSchemaUrn**](Enumscim2ExternalServerSchemaUrn.md) |  | 
**Id** | **string** | Name of the External Server | 
**ScimServiceURL** | **string** | The base URL for the SCIMv2 service. It must include the hostname, port, and base path to use to make SCIMv2 calls. | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the SCIMv2 service. | [optional] 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if it is necessary to present a client certificate to the SCIMv2 server. | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use to determine whether to trust the certificate presented by the SCIMv2 server during TLS negotiation. | [optional] 
**SslCertNickname** | Pointer to **string** | The nickname (alias) of the entry in the associated key store that holds the client certificate chain to present to the SCIMv2 server during TLS negotiation. | [optional] 
**HostnameVerificationMethod** | Pointer to [**EnumexternalServerScim2HostnameVerificationMethodProp**](EnumexternalServerScim2HostnameVerificationMethodProp.md) |  | [optional] 
**HttpAuthorizationMethod** | **string** | The method to use to authorize requests sent to the SCIMv2 server. | 
**ResponseTimeout** | Pointer to **string** | The maximum length of time to wait for a response from the SCIMv2 server when processing operations. | [optional] 
**ClientReconnectInterval** | Pointer to **string** | The maximum length of time that a client instance should remain active before being recreated. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewScim2ExternalServerResponse

`func NewScim2ExternalServerResponse(schemas []Enumscim2ExternalServerSchemaUrn, id string, scimServiceURL string, httpAuthorizationMethod string, ) *Scim2ExternalServerResponse`

NewScim2ExternalServerResponse instantiates a new Scim2ExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScim2ExternalServerResponseWithDefaults

`func NewScim2ExternalServerResponseWithDefaults() *Scim2ExternalServerResponse`

NewScim2ExternalServerResponseWithDefaults instantiates a new Scim2ExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Scim2ExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Scim2ExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Scim2ExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Scim2ExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *Scim2ExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *Scim2ExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *Scim2ExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *Scim2ExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *Scim2ExternalServerResponse) GetSchemas() []Enumscim2ExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *Scim2ExternalServerResponse) GetSchemasOk() (*[]Enumscim2ExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *Scim2ExternalServerResponse) SetSchemas(v []Enumscim2ExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *Scim2ExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Scim2ExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Scim2ExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetScimServiceURL

`func (o *Scim2ExternalServerResponse) GetScimServiceURL() string`

GetScimServiceURL returns the ScimServiceURL field if non-nil, zero value otherwise.

### GetScimServiceURLOk

`func (o *Scim2ExternalServerResponse) GetScimServiceURLOk() (*string, bool)`

GetScimServiceURLOk returns a tuple with the ScimServiceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScimServiceURL

`func (o *Scim2ExternalServerResponse) SetScimServiceURL(v string)`

SetScimServiceURL sets ScimServiceURL field to given value.


### GetHttpProxyExternalServer

`func (o *Scim2ExternalServerResponse) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *Scim2ExternalServerResponse) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *Scim2ExternalServerResponse) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *Scim2ExternalServerResponse) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *Scim2ExternalServerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *Scim2ExternalServerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *Scim2ExternalServerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *Scim2ExternalServerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *Scim2ExternalServerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *Scim2ExternalServerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *Scim2ExternalServerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *Scim2ExternalServerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *Scim2ExternalServerResponse) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *Scim2ExternalServerResponse) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *Scim2ExternalServerResponse) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *Scim2ExternalServerResponse) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetHostnameVerificationMethod

`func (o *Scim2ExternalServerResponse) GetHostnameVerificationMethod() EnumexternalServerScim2HostnameVerificationMethodProp`

GetHostnameVerificationMethod returns the HostnameVerificationMethod field if non-nil, zero value otherwise.

### GetHostnameVerificationMethodOk

`func (o *Scim2ExternalServerResponse) GetHostnameVerificationMethodOk() (*EnumexternalServerScim2HostnameVerificationMethodProp, bool)`

GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationMethod

`func (o *Scim2ExternalServerResponse) SetHostnameVerificationMethod(v EnumexternalServerScim2HostnameVerificationMethodProp)`

SetHostnameVerificationMethod sets HostnameVerificationMethod field to given value.

### HasHostnameVerificationMethod

`func (o *Scim2ExternalServerResponse) HasHostnameVerificationMethod() bool`

HasHostnameVerificationMethod returns a boolean if a field has been set.

### GetHttpAuthorizationMethod

`func (o *Scim2ExternalServerResponse) GetHttpAuthorizationMethod() string`

GetHttpAuthorizationMethod returns the HttpAuthorizationMethod field if non-nil, zero value otherwise.

### GetHttpAuthorizationMethodOk

`func (o *Scim2ExternalServerResponse) GetHttpAuthorizationMethodOk() (*string, bool)`

GetHttpAuthorizationMethodOk returns a tuple with the HttpAuthorizationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAuthorizationMethod

`func (o *Scim2ExternalServerResponse) SetHttpAuthorizationMethod(v string)`

SetHttpAuthorizationMethod sets HttpAuthorizationMethod field to given value.


### GetResponseTimeout

`func (o *Scim2ExternalServerResponse) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *Scim2ExternalServerResponse) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *Scim2ExternalServerResponse) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *Scim2ExternalServerResponse) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetClientReconnectInterval

`func (o *Scim2ExternalServerResponse) GetClientReconnectInterval() string`

GetClientReconnectInterval returns the ClientReconnectInterval field if non-nil, zero value otherwise.

### GetClientReconnectIntervalOk

`func (o *Scim2ExternalServerResponse) GetClientReconnectIntervalOk() (*string, bool)`

GetClientReconnectIntervalOk returns a tuple with the ClientReconnectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReconnectInterval

`func (o *Scim2ExternalServerResponse) SetClientReconnectInterval(v string)`

SetClientReconnectInterval sets ClientReconnectInterval field to given value.

### HasClientReconnectInterval

`func (o *Scim2ExternalServerResponse) HasClientReconnectInterval() bool`

HasClientReconnectInterval returns a boolean if a field has been set.

### GetDescription

`func (o *Scim2ExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Scim2ExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Scim2ExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Scim2ExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


