# PingOneHttpExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the External Server | 
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

### NewPingOneHttpExternalServerResponse

`func NewPingOneHttpExternalServerResponse(id string, schemas []EnumpingOneHttpExternalServerSchemaUrn, baseURL string, ) *PingOneHttpExternalServerResponse`

NewPingOneHttpExternalServerResponse instantiates a new PingOneHttpExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneHttpExternalServerResponseWithDefaults

`func NewPingOneHttpExternalServerResponseWithDefaults() *PingOneHttpExternalServerResponse`

NewPingOneHttpExternalServerResponseWithDefaults instantiates a new PingOneHttpExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetBaseURL

`func (o *PingOneHttpExternalServerResponse) GetBaseURL() string`

GetBaseURL returns the BaseURL field if non-nil, zero value otherwise.

### GetBaseURLOk

`func (o *PingOneHttpExternalServerResponse) GetBaseURLOk() (*string, bool)`

GetBaseURLOk returns a tuple with the BaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURL

`func (o *PingOneHttpExternalServerResponse) SetBaseURL(v string)`

SetBaseURL sets BaseURL field to given value.


### GetKeyManagerProvider

`func (o *PingOneHttpExternalServerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *PingOneHttpExternalServerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *PingOneHttpExternalServerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *PingOneHttpExternalServerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *PingOneHttpExternalServerResponse) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *PingOneHttpExternalServerResponse) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *PingOneHttpExternalServerResponse) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *PingOneHttpExternalServerResponse) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.

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


