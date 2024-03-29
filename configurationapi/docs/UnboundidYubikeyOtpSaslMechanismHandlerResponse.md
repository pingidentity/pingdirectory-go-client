# UnboundidYubikeyOtpSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn**](EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the SASL Mechanism Handler | 
**YubikeyClientID** | Pointer to **string** | The client ID to include in requests to the YubiKey validation server. A client ID and API key may be obtained for free from https://upgrade.yubico.com/getapikey/. | [optional] 
**YubikeyAPIKey** | Pointer to **string** | The API key needed to verify signatures generated by the YubiKey validation server. A client ID and API key may be obtained for free from https://upgrade.yubico.com/getapikey/. | [optional] 
**YubikeyAPIKeyPassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the API key needed to verify signatures generated by the YubiKey validation server. A client ID and API key may be obtained for free from https://upgrade.yubico.com/getapikey/. | [optional] 
**YubikeyValidationServerBaseURL** | **[]string** | The base URL of the validation server to use to verify one-time passwords. You should only need to change the value if you wish to use your own validation server instead of using one of the Yubico servers. The server must use the YubiKey Validation Protocol version 2.0. | 
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the YubiKey validation service. | [optional] 
**HttpConnectTimeout** | Pointer to **string** | The maximum length of time to wait to obtain an HTTP connection. | [optional] 
**HttpResponseTimeout** | Pointer to **string** | The maximum length of time to wait for a response to an HTTP request. | [optional] 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user(s) targeted in the authentication and/or authorization identities contained in the bind request. This will only be used for \&quot;u:\&quot;-style identities. | 
**RequireStaticPassword** | Pointer to **bool** | Indicates whether a user will be required to provide a static password when authenticating via the UNBOUNDID-YUBIKEY-OTP SASL mechanism. | [optional] 
**KeyManagerProvider** | Pointer to **string** | Specifies which key manager provider should be used to obtain a client certificate to present to the validation server when performing HTTPS communication. This may be left undefined if communication will not be secured with HTTPS, or if there is no need to present a client certificate to the validation service. | [optional] 
**TrustManagerProvider** | Pointer to **string** | Specifies which trust manager provider should be used to determine whether to trust the certificate presented by the server when performing HTTPS communication. This may be left undefined if HTTPS communication is not needed, or if the validation service presents a certificate that is trusted by the default JVM configuration (which should be the case for the validation servers that Yubico provides, but may not be the case if an alternate validation server is configured). | [optional] 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewUnboundidYubikeyOtpSaslMechanismHandlerResponse

`func NewUnboundidYubikeyOtpSaslMechanismHandlerResponse(schemas []EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn, id string, yubikeyValidationServerBaseURL []string, identityMapper string, enabled bool, ) *UnboundidYubikeyOtpSaslMechanismHandlerResponse`

NewUnboundidYubikeyOtpSaslMechanismHandlerResponse instantiates a new UnboundidYubikeyOtpSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnboundidYubikeyOtpSaslMechanismHandlerResponseWithDefaults

`func NewUnboundidYubikeyOtpSaslMechanismHandlerResponseWithDefaults() *UnboundidYubikeyOtpSaslMechanismHandlerResponse`

NewUnboundidYubikeyOtpSaslMechanismHandlerResponseWithDefaults instantiates a new UnboundidYubikeyOtpSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetSchemas() []EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetSchemas(v []EnumunboundidYubikeyOtpSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetYubikeyClientID

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetYubikeyClientID() string`

GetYubikeyClientID returns the YubikeyClientID field if non-nil, zero value otherwise.

### GetYubikeyClientIDOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetYubikeyClientIDOk() (*string, bool)`

GetYubikeyClientIDOk returns a tuple with the YubikeyClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYubikeyClientID

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetYubikeyClientID(v string)`

SetYubikeyClientID sets YubikeyClientID field to given value.

### HasYubikeyClientID

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasYubikeyClientID() bool`

HasYubikeyClientID returns a boolean if a field has been set.

### GetYubikeyAPIKey

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetYubikeyAPIKey() string`

GetYubikeyAPIKey returns the YubikeyAPIKey field if non-nil, zero value otherwise.

### GetYubikeyAPIKeyOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetYubikeyAPIKeyOk() (*string, bool)`

GetYubikeyAPIKeyOk returns a tuple with the YubikeyAPIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYubikeyAPIKey

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetYubikeyAPIKey(v string)`

SetYubikeyAPIKey sets YubikeyAPIKey field to given value.

### HasYubikeyAPIKey

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasYubikeyAPIKey() bool`

HasYubikeyAPIKey returns a boolean if a field has been set.

### GetYubikeyAPIKeyPassphraseProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetYubikeyAPIKeyPassphraseProvider() string`

GetYubikeyAPIKeyPassphraseProvider returns the YubikeyAPIKeyPassphraseProvider field if non-nil, zero value otherwise.

### GetYubikeyAPIKeyPassphraseProviderOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetYubikeyAPIKeyPassphraseProviderOk() (*string, bool)`

GetYubikeyAPIKeyPassphraseProviderOk returns a tuple with the YubikeyAPIKeyPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYubikeyAPIKeyPassphraseProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetYubikeyAPIKeyPassphraseProvider(v string)`

SetYubikeyAPIKeyPassphraseProvider sets YubikeyAPIKeyPassphraseProvider field to given value.

### HasYubikeyAPIKeyPassphraseProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasYubikeyAPIKeyPassphraseProvider() bool`

HasYubikeyAPIKeyPassphraseProvider returns a boolean if a field has been set.

### GetYubikeyValidationServerBaseURL

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetYubikeyValidationServerBaseURL() []string`

GetYubikeyValidationServerBaseURL returns the YubikeyValidationServerBaseURL field if non-nil, zero value otherwise.

### GetYubikeyValidationServerBaseURLOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetYubikeyValidationServerBaseURLOk() (*[]string, bool)`

GetYubikeyValidationServerBaseURLOk returns a tuple with the YubikeyValidationServerBaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYubikeyValidationServerBaseURL

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetYubikeyValidationServerBaseURL(v []string)`

SetYubikeyValidationServerBaseURL sets YubikeyValidationServerBaseURL field to given value.


### GetHttpProxyExternalServer

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetHttpConnectTimeout

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetHttpConnectTimeout() string`

GetHttpConnectTimeout returns the HttpConnectTimeout field if non-nil, zero value otherwise.

### GetHttpConnectTimeoutOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetHttpConnectTimeoutOk() (*string, bool)`

GetHttpConnectTimeoutOk returns a tuple with the HttpConnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpConnectTimeout

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetHttpConnectTimeout(v string)`

SetHttpConnectTimeout sets HttpConnectTimeout field to given value.

### HasHttpConnectTimeout

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasHttpConnectTimeout() bool`

HasHttpConnectTimeout returns a boolean if a field has been set.

### GetHttpResponseTimeout

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetHttpResponseTimeout() string`

GetHttpResponseTimeout returns the HttpResponseTimeout field if non-nil, zero value otherwise.

### GetHttpResponseTimeoutOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetHttpResponseTimeoutOk() (*string, bool)`

GetHttpResponseTimeoutOk returns a tuple with the HttpResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseTimeout

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetHttpResponseTimeout(v string)`

SetHttpResponseTimeout sets HttpResponseTimeout field to given value.

### HasHttpResponseTimeout

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasHttpResponseTimeout() bool`

HasHttpResponseTimeout returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetRequireStaticPassword

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetRequireStaticPassword() bool`

GetRequireStaticPassword returns the RequireStaticPassword field if non-nil, zero value otherwise.

### GetRequireStaticPasswordOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetRequireStaticPasswordOk() (*bool, bool)`

GetRequireStaticPasswordOk returns a tuple with the RequireStaticPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireStaticPassword

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetRequireStaticPassword(v bool)`

SetRequireStaticPassword sets RequireStaticPassword field to given value.

### HasRequireStaticPassword

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasRequireStaticPassword() bool`

HasRequireStaticPassword returns a boolean if a field has been set.

### GetKeyManagerProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetKeyManagerProvider() string`

GetKeyManagerProvider returns the KeyManagerProvider field if non-nil, zero value otherwise.

### GetKeyManagerProviderOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetKeyManagerProviderOk() (*string, bool)`

GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyManagerProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetKeyManagerProvider(v string)`

SetKeyManagerProvider sets KeyManagerProvider field to given value.

### HasKeyManagerProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasKeyManagerProvider() bool`

HasKeyManagerProvider returns a boolean if a field has been set.

### GetTrustManagerProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetTrustManagerProvider() string`

GetTrustManagerProvider returns the TrustManagerProvider field if non-nil, zero value otherwise.

### GetTrustManagerProviderOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetTrustManagerProviderOk() (*string, bool)`

GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustManagerProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetTrustManagerProvider(v string)`

SetTrustManagerProvider sets TrustManagerProvider field to given value.

### HasTrustManagerProvider

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasTrustManagerProvider() bool`

HasTrustManagerProvider returns a boolean if a field has been set.

### GetDescription

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UnboundidYubikeyOtpSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


