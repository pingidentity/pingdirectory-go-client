# AddPingOneHttpExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Name of the new External Server | 
**Schemas** | [**[]EnumpingOneHttpExternalServerSchemaUrn**](EnumpingOneHttpExternalServerSchemaUrn.md) |  | 
**HostnameVerificationMethod** | Pointer to [**EnumexternalServerHostnameVerificationMethodProp**](EnumexternalServerHostnameVerificationMethodProp.md) |  | [optional] 
**TrustManagerProvider** | Pointer to **string** | The trust manager provider to use for HTTPS connection-level security. | [optional] 
**ConnectTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for a connection to be established before aborting a request to PingOne. | [optional] 
**ResponseTimeout** | Pointer to **string** | Specifies the maximum length of time to wait for response data to be read from an established connection before aborting a request to PingOne. | [optional] 
**BaseURL** | **string** | The base URL of the external server, optionally including port number, for example \&quot;https://externalService:9031\&quot;. | 
**KeyManagerProvider** | Pointer to **string** | The key manager provider to use if SSL (HTTPS) is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**SslCertNickname** | Pointer to **string** | The certificate alias within the keystore to use if SSL (HTTPS) is to be used for connection-level security. When specifying a value for this property you must ensure that the external server trusts this server&#39;s public certificate by adding this server&#39;s public certificate to the external server&#39;s trust store. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewAddPingOneHttpExternalServerRequest

`func NewAddPingOneHttpExternalServerRequest(serverName string, schemas []EnumpingOneHttpExternalServerSchemaUrn, baseURL string, ) *AddPingOneHttpExternalServerRequest`

NewAddPingOneHttpExternalServerRequest instantiates a new AddPingOneHttpExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPingOneHttpExternalServerRequestWithDefaults

`func NewAddPingOneHttpExternalServerRequestWithDefaults() *AddPingOneHttpExternalServerRequest`

NewAddPingOneHttpExternalServerRequestWithDefaults instantiates a new AddPingOneHttpExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AddPingOneHttpExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddPingOneHttpExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddPingOneHttpExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSchemas

`func (o *AddPingOneHttpExternalServerRequest) GetSchemas() []EnumpingOneHttpExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPingOneHttpExternalServerRequest) GetSchemasOk() (*[]EnumpingOneHttpExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPingOneHttpExternalServerRequest) SetSchemas(v []EnumpingOneHttpExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetHostnameVerificationMethod

`func (o *AddPingOneHttpExternalServerRequest) GetHostnameVerificationMethod() EnumexternalServerHostnameVerificationMethodProp`

GetHostnameVerificationMethod returns the HostnameVerificationMethod field if non-nil, zero value otherwise.

### GetHostnameVerificationMethodOk

`func (o *AddPingOneHttpExternalServerRequest) GetHostnameVerificationMethodOk() (*EnumexternalServerHostnameVerificationMethodProp, bool)`

GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationMethod

`func (o *AddPingOneHttpExternalServerRequest) SetHostnameVerificationMethod(v EnumexternalServerHostnameVerificationMethodProp)`

SetHostnameVerificationMethod sets HostnameVerificationMethod field to given value.

### HasHostnameVerificationMethod

`func (o *AddPingOneHttpExternalServerRequest) HasHostnameVerificationMethod() bool`

HasHostnameVerificationMethod returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *AddPingOneHttpExternalServerRequest) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *AddPingOneHttpExternalServerRequest) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *AddPingOneHttpExternalServerRequest) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *AddPingOneHttpExternalServerRequest) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *AddPingOneHttpExternalServerRequest) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *AddPingOneHttpExternalServerRequest) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *AddPingOneHttpExternalServerRequest) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *AddPingOneHttpExternalServerRequest) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *AddPingOneHttpExternalServerRequest) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *AddPingOneHttpExternalServerRequest) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *AddPingOneHttpExternalServerRequest) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *AddPingOneHttpExternalServerRequest) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetBaseURL

`func (o *AddPingOneHttpExternalServerRequest) GetBaseURL() string`

GetBaseURL returns the BaseURL field if non-nil, zero value otherwise.

### GetBaseURLOk

`func (o *AddPingOneHttpExternalServerRequest) GetBaseURLOk() (*string, bool)`

GetBaseURLOk returns a tuple with the BaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURL

`func (o *AddPingOneHttpExternalServerRequest) SetBaseURL(v string)`

SetBaseURL sets BaseURL field to given value.


### GetKeyManagerProvider

`func (o *AddPingOneHttpExternalServerRequest) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *AddPingOneHttpExternalServerRequest) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *AddPingOneHttpExternalServerRequest) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *AddPingOneHttpExternalServerRequest) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *AddPingOneHttpExternalServerRequest) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *AddPingOneHttpExternalServerRequest) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *AddPingOneHttpExternalServerRequest) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *AddPingOneHttpExternalServerRequest) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetDescription

`func (o *AddPingOneHttpExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPingOneHttpExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPingOneHttpExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPingOneHttpExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


