# CryptoManagerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | Pointer to [**[]EnumcryptoManagerSchemaUrn**](EnumcryptoManagerSchemaUrn.md) |  | [optional] 
**DigestAlgorithm** | Pointer to **string** | Specifies the preferred message digest algorithm for the Directory Server. | [optional] 
**MacAlgorithm** | Pointer to **string** | Specifies the preferred MAC algorithm for the Directory Server. | [optional] 
**MacKeyLength** | Pointer to **int32** | Specifies the key length in bits for the preferred MAC algorithm. | [optional] 
**CipherTransformation** | Pointer to **string** | Specifies the cipher for the Directory Server using the syntax algorithm/mode/padding. | [optional] 
**CipherKeyLength** | Pointer to **int32** | Specifies the key length in bits for the preferred cipher. | [optional] 
**KeyWrappingTransformation** | Pointer to **string** | The preferred key wrapping transformation for the Directory Server. This value must be the same for all server instances in a replication topology. | [optional] 
**SslProtocol** | Pointer to **[]string** | Specifies the names of TLS protocols that are allowed for use in secure communication. | [optional] 
**SslCipherSuite** | Pointer to **[]string** | Specifies the names of the TLS cipher suites that are allowed for use in secure communication. | [optional] 
**OutboundSSLProtocol** | Pointer to **[]string** | Specifies the names of the TLS protocols that will be enabled for outbound connections initiated by the Directory Server. | [optional] 
**OutboundSSLCipherSuite** | Pointer to **[]string** | Specifies the names of the TLS cipher suites that will be enabled for outbound connections initiated by the Directory Server. | [optional] 
**EnableSha1CipherSuites** | Pointer to **bool** | Indicates whether to enable support for TLS cipher suites that use the SHA-1 digest algorithm. The SHA-1 digest algorithm is no longer considered secure and is not recommended for use. | [optional] 
**EnableRsaKeyExchangeCipherSuites** | Pointer to **bool** | Indicates whether to enable support for TLS cipher suites that use the RSA key exchange algorithm. Cipher suites that rely on RSA key exchange are not recommended because they do not support forward secrecy, which means that if the private key is compromised, then any communication negotiated using that private key should also be considered compromised. | [optional] 
**SslCertNickname** | Pointer to **string** | Specifies the nickname (also called the alias) of the certificate that the Crypto Manager should use when performing SSL communication. | [optional] 

## Methods

### NewCryptoManagerResponse

`func NewCryptoManagerResponse() *CryptoManagerResponse`

NewCryptoManagerResponse instantiates a new CryptoManagerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptoManagerResponseWithDefaults

`func NewCryptoManagerResponseWithDefaults() *CryptoManagerResponse`

NewCryptoManagerResponseWithDefaults instantiates a new CryptoManagerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CryptoManagerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CryptoManagerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CryptoManagerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CryptoManagerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *CryptoManagerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *CryptoManagerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *CryptoManagerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *CryptoManagerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *CryptoManagerResponse) GetSchemas() []EnumcryptoManagerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *CryptoManagerResponse) GetSchemasOk() (*[]EnumcryptoManagerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *CryptoManagerResponse) SetSchemas(v []EnumcryptoManagerSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *CryptoManagerResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDigestAlgorithm

`func (o *CryptoManagerResponse) GetDigestAlgorithm() string`

GetDigestAlgorithm returns the DigestAlgorithm field if non-nil, zero value otherwise.

### GetDigestAlgorithmOk

`func (o *CryptoManagerResponse) GetDigestAlgorithmOk() (*string, bool)`

GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestAlgorithm

`func (o *CryptoManagerResponse) SetDigestAlgorithm(v string)`

SetDigestAlgorithm sets DigestAlgorithm field to given value.

### HasDigestAlgorithm

`func (o *CryptoManagerResponse) HasDigestAlgorithm() bool`

HasDigestAlgorithm returns a boolean if a field has been set.

### GetMacAlgorithm

`func (o *CryptoManagerResponse) GetMacAlgorithm() string`

GetMacAlgorithm returns the MacAlgorithm field if non-nil, zero value otherwise.

### GetMacAlgorithmOk

`func (o *CryptoManagerResponse) GetMacAlgorithmOk() (*string, bool)`

GetMacAlgorithmOk returns a tuple with the MacAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAlgorithm

`func (o *CryptoManagerResponse) SetMacAlgorithm(v string)`

SetMacAlgorithm sets MacAlgorithm field to given value.

### HasMacAlgorithm

`func (o *CryptoManagerResponse) HasMacAlgorithm() bool`

HasMacAlgorithm returns a boolean if a field has been set.

### GetMacKeyLength

`func (o *CryptoManagerResponse) GetMacKeyLength() int32`

GetMacKeyLength returns the MacKeyLength field if non-nil, zero value otherwise.

### GetMacKeyLengthOk

`func (o *CryptoManagerResponse) GetMacKeyLengthOk() (*int32, bool)`

GetMacKeyLengthOk returns a tuple with the MacKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacKeyLength

`func (o *CryptoManagerResponse) SetMacKeyLength(v int32)`

SetMacKeyLength sets MacKeyLength field to given value.

### HasMacKeyLength

`func (o *CryptoManagerResponse) HasMacKeyLength() bool`

HasMacKeyLength returns a boolean if a field has been set.

### GetCipherTransformation

`func (o *CryptoManagerResponse) GetCipherTransformation() string`

GetCipherTransformation returns the CipherTransformation field if non-nil, zero value otherwise.

### GetCipherTransformationOk

`func (o *CryptoManagerResponse) GetCipherTransformationOk() (*string, bool)`

GetCipherTransformationOk returns a tuple with the CipherTransformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherTransformation

`func (o *CryptoManagerResponse) SetCipherTransformation(v string)`

SetCipherTransformation sets CipherTransformation field to given value.

### HasCipherTransformation

`func (o *CryptoManagerResponse) HasCipherTransformation() bool`

HasCipherTransformation returns a boolean if a field has been set.

### GetCipherKeyLength

`func (o *CryptoManagerResponse) GetCipherKeyLength() int32`

GetCipherKeyLength returns the CipherKeyLength field if non-nil, zero value otherwise.

### GetCipherKeyLengthOk

`func (o *CryptoManagerResponse) GetCipherKeyLengthOk() (*int32, bool)`

GetCipherKeyLengthOk returns a tuple with the CipherKeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherKeyLength

`func (o *CryptoManagerResponse) SetCipherKeyLength(v int32)`

SetCipherKeyLength sets CipherKeyLength field to given value.

### HasCipherKeyLength

`func (o *CryptoManagerResponse) HasCipherKeyLength() bool`

HasCipherKeyLength returns a boolean if a field has been set.

### GetKeyWrappingTransformation

`func (o *CryptoManagerResponse) GetKeyWrappingTransformation() string`

GetKeyWrappingTransformation returns the KeyWrappingTransformation field if non-nil, zero value otherwise.

### GetKeyWrappingTransformationOk

`func (o *CryptoManagerResponse) GetKeyWrappingTransformationOk() (*string, bool)`

GetKeyWrappingTransformationOk returns a tuple with the KeyWrappingTransformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWrappingTransformation

`func (o *CryptoManagerResponse) SetKeyWrappingTransformation(v string)`

SetKeyWrappingTransformation sets KeyWrappingTransformation field to given value.

### HasKeyWrappingTransformation

`func (o *CryptoManagerResponse) HasKeyWrappingTransformation() bool`

HasKeyWrappingTransformation returns a boolean if a field has been set.

### GetSslProtocol

`func (o *CryptoManagerResponse) GetSslProtocol() []string`

GetSslProtocol returns the SslProtocol field if non-nil, zero value otherwise.

### GetSslProtocolOk

`func (o *CryptoManagerResponse) GetSslProtocolOk() (*[]string, bool)`

GetSslProtocolOk returns a tuple with the SslProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslProtocol

`func (o *CryptoManagerResponse) SetSslProtocol(v []string)`

SetSslProtocol sets SslProtocol field to given value.

### HasSslProtocol

`func (o *CryptoManagerResponse) HasSslProtocol() bool`

HasSslProtocol returns a boolean if a field has been set.

### GetSslCipherSuite

`func (o *CryptoManagerResponse) GetSslCipherSuite() []string`

GetSslCipherSuite returns the SslCipherSuite field if non-nil, zero value otherwise.

### GetSslCipherSuiteOk

`func (o *CryptoManagerResponse) GetSslCipherSuiteOk() (*[]string, bool)`

GetSslCipherSuiteOk returns a tuple with the SslCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCipherSuite

`func (o *CryptoManagerResponse) SetSslCipherSuite(v []string)`

SetSslCipherSuite sets SslCipherSuite field to given value.

### HasSslCipherSuite

`func (o *CryptoManagerResponse) HasSslCipherSuite() bool`

HasSslCipherSuite returns a boolean if a field has been set.

### GetOutboundSSLProtocol

`func (o *CryptoManagerResponse) GetOutboundSSLProtocol() []string`

GetOutboundSSLProtocol returns the OutboundSSLProtocol field if non-nil, zero value otherwise.

### GetOutboundSSLProtocolOk

`func (o *CryptoManagerResponse) GetOutboundSSLProtocolOk() (*[]string, bool)`

GetOutboundSSLProtocolOk returns a tuple with the OutboundSSLProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSSLProtocol

`func (o *CryptoManagerResponse) SetOutboundSSLProtocol(v []string)`

SetOutboundSSLProtocol sets OutboundSSLProtocol field to given value.

### HasOutboundSSLProtocol

`func (o *CryptoManagerResponse) HasOutboundSSLProtocol() bool`

HasOutboundSSLProtocol returns a boolean if a field has been set.

### GetOutboundSSLCipherSuite

`func (o *CryptoManagerResponse) GetOutboundSSLCipherSuite() []string`

GetOutboundSSLCipherSuite returns the OutboundSSLCipherSuite field if non-nil, zero value otherwise.

### GetOutboundSSLCipherSuiteOk

`func (o *CryptoManagerResponse) GetOutboundSSLCipherSuiteOk() (*[]string, bool)`

GetOutboundSSLCipherSuiteOk returns a tuple with the OutboundSSLCipherSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundSSLCipherSuite

`func (o *CryptoManagerResponse) SetOutboundSSLCipherSuite(v []string)`

SetOutboundSSLCipherSuite sets OutboundSSLCipherSuite field to given value.

### HasOutboundSSLCipherSuite

`func (o *CryptoManagerResponse) HasOutboundSSLCipherSuite() bool`

HasOutboundSSLCipherSuite returns a boolean if a field has been set.

### GetEnableSha1CipherSuites

`func (o *CryptoManagerResponse) GetEnableSha1CipherSuites() bool`

GetEnableSha1CipherSuites returns the EnableSha1CipherSuites field if non-nil, zero value otherwise.

### GetEnableSha1CipherSuitesOk

`func (o *CryptoManagerResponse) GetEnableSha1CipherSuitesOk() (*bool, bool)`

GetEnableSha1CipherSuitesOk returns a tuple with the EnableSha1CipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSha1CipherSuites

`func (o *CryptoManagerResponse) SetEnableSha1CipherSuites(v bool)`

SetEnableSha1CipherSuites sets EnableSha1CipherSuites field to given value.

### HasEnableSha1CipherSuites

`func (o *CryptoManagerResponse) HasEnableSha1CipherSuites() bool`

HasEnableSha1CipherSuites returns a boolean if a field has been set.

### GetEnableRsaKeyExchangeCipherSuites

`func (o *CryptoManagerResponse) GetEnableRsaKeyExchangeCipherSuites() bool`

GetEnableRsaKeyExchangeCipherSuites returns the EnableRsaKeyExchangeCipherSuites field if non-nil, zero value otherwise.

### GetEnableRsaKeyExchangeCipherSuitesOk

`func (o *CryptoManagerResponse) GetEnableRsaKeyExchangeCipherSuitesOk() (*bool, bool)`

GetEnableRsaKeyExchangeCipherSuitesOk returns a tuple with the EnableRsaKeyExchangeCipherSuites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRsaKeyExchangeCipherSuites

`func (o *CryptoManagerResponse) SetEnableRsaKeyExchangeCipherSuites(v bool)`

SetEnableRsaKeyExchangeCipherSuites sets EnableRsaKeyExchangeCipherSuites field to given value.

### HasEnableRsaKeyExchangeCipherSuites

`func (o *CryptoManagerResponse) HasEnableRsaKeyExchangeCipherSuites() bool`

HasEnableRsaKeyExchangeCipherSuites returns a boolean if a field has been set.

### GetSslCertNickname

`func (o *CryptoManagerResponse) GetSslCertNickname() string`

GetSslCertNickname returns the SslCertNickname field if non-nil, zero value otherwise.

### GetSslCertNicknameOk

`func (o *CryptoManagerResponse) GetSslCertNicknameOk() (*string, bool)`

GetSslCertNicknameOk returns a tuple with the SslCertNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertNickname

`func (o *CryptoManagerResponse) SetSslCertNickname(v string)`

SetSslCertNickname sets SslCertNickname field to given value.

### HasSslCertNickname

`func (o *CryptoManagerResponse) HasSslCertNickname() bool`

HasSslCertNickname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


