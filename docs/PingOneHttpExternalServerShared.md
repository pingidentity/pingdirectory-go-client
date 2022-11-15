# PingOneHttpExternalServerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewPingOneHttpExternalServerShared

`func NewPingOneHttpExternalServerShared(schemas []EnumpingOneHttpExternalServerSchemaUrn, baseURL string, ) *PingOneHttpExternalServerShared`

NewPingOneHttpExternalServerShared instantiates a new PingOneHttpExternalServerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneHttpExternalServerSharedWithDefaults

`func NewPingOneHttpExternalServerSharedWithDefaults() *PingOneHttpExternalServerShared`

NewPingOneHttpExternalServerSharedWithDefaults instantiates a new PingOneHttpExternalServerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PingOneHttpExternalServerShared) GetSchemas() []EnumpingOneHttpExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PingOneHttpExternalServerShared) GetSchemasOk() (*[]EnumpingOneHttpExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PingOneHttpExternalServerShared) SetSchemas(v []EnumpingOneHttpExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetHostnameVerificationMethod

`func (o *PingOneHttpExternalServerShared) GetHostnameVerificationMethod() EnumexternalServerHostnameVerificationMethodProp`

GetHostnameVerificationMethod returns the HostnameVerificationMethod field if non-nil, zero value otherwise.

### GetHostnameVerificationMethodOk

`func (o *PingOneHttpExternalServerShared) GetHostnameVerificationMethodOk() (*EnumexternalServerHostnameVerificationMethodProp, bool)`

GetHostnameVerificationMethodOk returns a tuple with the HostnameVerificationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnameVerificationMethod

`func (o *PingOneHttpExternalServerShared) SetHostnameVerificationMethod(v EnumexternalServerHostnameVerificationMethodProp)`

SetHostnameVerificationMethod sets HostnameVerificationMethod field to given value.

### HasHostnameVerificationMethod

`func (o *PingOneHttpExternalServerShared) HasHostnameVerificationMethod() bool`

HasHostnameVerificationMethod returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *PingOneHttpExternalServerShared) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *PingOneHttpExternalServerShared) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *PingOneHttpExternalServerShared) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *PingOneHttpExternalServerShared) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetConnectTimeout

`func (o *PingOneHttpExternalServerShared) GetConnectTimeout() string`

GetConnectTimeout returns the ConnectTimeout field if non-nil, zero value otherwise.

### GetConnectTimeoutOk

`func (o *PingOneHttpExternalServerShared) GetConnectTimeoutOk() (*string, bool)`

GetConnectTimeoutOk returns a tuple with the ConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTimeout

`func (o *PingOneHttpExternalServerShared) SetConnectTimeout(v string)`

SetConnectTimeout sets ConnectTimeout field to given value.

### HasConnectTimeout

`func (o *PingOneHttpExternalServerShared) HasConnectTimeout() bool`

HasConnectTimeout returns a boolean if a field has been set.

### GetResponseTimeout

`func (o *PingOneHttpExternalServerShared) GetResponseTimeout() string`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *PingOneHttpExternalServerShared) GetResponseTimeoutOk() (*string, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *PingOneHttpExternalServerShared) SetResponseTimeout(v string)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *PingOneHttpExternalServerShared) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### GetBaseURL

`func (o *PingOneHttpExternalServerShared) GetBaseURL() string`

GetBaseURL returns the BaseURL field if non-nil, zero value otherwise.

### GetBaseURLOk

`func (o *PingOneHttpExternalServerShared) GetBaseURLOk() (*string, bool)`

GetBaseURLOk returns a tuple with the BaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURL

`func (o *PingOneHttpExternalServerShared) SetBaseURL(v string)`

SetBaseURL sets BaseURL field to given value.


### GetKeyManagerProvider

`func (o *PingOneHttpExternalServerShared) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *PingOneHttpExternalServerShared) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *PingOneHttpExternalServerShared) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *PingOneHttpExternalServerShared) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *PingOneHttpExternalServerShared) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *PingOneHttpExternalServerShared) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *PingOneHttpExternalServerShared) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *PingOneHttpExternalServerShared) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

### GetDescription

`func (o *PingOneHttpExternalServerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingOneHttpExternalServerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingOneHttpExternalServerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingOneHttpExternalServerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


