# AddAccessTokenValidator200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Access Token Validator | 
**Schemas** | [**[]EnumthirdPartyAccessTokenValidatorSchemaUrn**](EnumthirdPartyAccessTokenValidatorSchemaUrn.md) |  | 
**ClientID** | **string** | The client identifier to use when authenticating to the PingFederate authorization server. | 
**ClientSecret** | Pointer to **string** | The client secret to use when authenticating to the PingFederate authorization server. | [optional] 
**ClientSecretPassphraseProvider** | Pointer to **string** | The passphrase provider for obtaining the client secret to use when authenticating to the PingFederate authorization server. | [optional] 
**IncludeAudParameter** | Pointer to **bool** | Whether to include the incoming request URL as the \&quot;aud\&quot; parameter when calling the PingFederate introspection endpoint. This property is ignored if the access-token-manager-id property is set. | [optional] 
**AccessTokenManagerID** | Pointer to **string** | The Access Token Manager instance ID to specify when calling the PingFederate introspection endpoint. If this property is set the include-aud-parameter property is ignored. | [optional] 
**EndpointCacheRefresh** | Pointer to **string** | How often the Access Token Validator should refresh its stored value of the PingFederate server&#39;s token introspection endpoint. | [optional] 
**EvaluationOrderIndex** | **int64** | When multiple Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Access Token Validators defined within Directory Server but not necessarily contiguous. Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token. | 
**AuthorizationServer** | Pointer to **string** | Specifies the external server that will be used to aid in validating access tokens. In most cases this will be the Authorization Server that minted the token. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | [optional] 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server. | [optional] 
**Description** | Pointer to **string** | A description for this Access Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this Access Token Validator is enabled for use in Directory Server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**AllowedSigningAlgorithm** | [**[]EnumaccessTokenValidatorAllowedSigningAlgorithmProp**](EnumaccessTokenValidatorAllowedSigningAlgorithmProp.md) |  | 
**SigningCertificate** | Pointer to **[]string** | Specifies the locally stored certificates that may be used to validate the signature of an incoming JWT access token. If this property is specified, the JWT Access Token Validator will not use a JWKS endpoint to retrieve public keys. | [optional] 
**JwksEndpointPath** | Pointer to **string** | The relative path to JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming JWT access token. This path is relative to the base_url property defined for the validator&#39;s external authorization server. If jwks-endpoint-path is specified, the JWT Access Token Validator will not consult locally stored certificates for validating token signatures. | [optional] 
**EncryptionKeyPair** | Pointer to **string** | The public-private key pair that is used to encrypt the JWT payload. If specified, the JWT Access Token Validator will use the private key to decrypt the JWT payload, and the public key must be exported to the Authorization Server that is issuing access tokens. | [optional] 
**AllowedKeyEncryptionAlgorithm** | [**[]EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp**](EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp.md) |  | 
**AllowedContentEncryptionAlgorithm** | [**[]EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp**](EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp.md) |  | 
**ClockSkewGracePeriod** | Pointer to **string** | Specifies the amount of clock skew that is tolerated by the JWT Access Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token&#39;s not-before (nbf) time and added to the token&#39;s expiration (exp) time, if present, to allow for any time difference between the local server&#39;s clock and the token issuer&#39;s clock. | [optional] 
**ClientIDClaimName** | Pointer to **string** | The name of the token claim that contains the OAuth2 client ID. | [optional] 
**ScopeClaimName** | Pointer to **string** | The name of the token claim that contains the scopes granted by the token. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Access Token Validator. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Access Token Validator. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddAccessTokenValidator200Response

`func NewAddAccessTokenValidator200Response(id string, schemas []EnumthirdPartyAccessTokenValidatorSchemaUrn, clientID string, evaluationOrderIndex int64, enabled bool, allowedSigningAlgorithm []EnumaccessTokenValidatorAllowedSigningAlgorithmProp, allowedKeyEncryptionAlgorithm []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, allowedContentEncryptionAlgorithm []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, extensionClass string, ) *AddAccessTokenValidator200Response`

NewAddAccessTokenValidator200Response instantiates a new AddAccessTokenValidator200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccessTokenValidator200ResponseWithDefaults

`func NewAddAccessTokenValidator200ResponseWithDefaults() *AddAccessTokenValidator200Response`

NewAddAccessTokenValidator200ResponseWithDefaults instantiates a new AddAccessTokenValidator200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddAccessTokenValidator200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddAccessTokenValidator200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddAccessTokenValidator200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddAccessTokenValidator200Response) GetSchemas() []EnumthirdPartyAccessTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAccessTokenValidator200Response) GetSchemasOk() (*[]EnumthirdPartyAccessTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAccessTokenValidator200Response) SetSchemas(v []EnumthirdPartyAccessTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetClientID

`func (o *AddAccessTokenValidator200Response) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AddAccessTokenValidator200Response) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AddAccessTokenValidator200Response) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *AddAccessTokenValidator200Response) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AddAccessTokenValidator200Response) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AddAccessTokenValidator200Response) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AddAccessTokenValidator200Response) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecretPassphraseProvider

`func (o *AddAccessTokenValidator200Response) GetClientSecretPassphraseProvider() string`

GetClientSecretPassphraseProvider returns the ClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetClientSecretPassphraseProviderOk

`func (o *AddAccessTokenValidator200Response) GetClientSecretPassphraseProviderOk() (*string, bool)`

GetClientSecretPassphraseProviderOk returns a tuple with the ClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretPassphraseProvider

`func (o *AddAccessTokenValidator200Response) SetClientSecretPassphraseProvider(v string)`

SetClientSecretPassphraseProvider sets ClientSecretPassphraseProvider field to given value.

### HasClientSecretPassphraseProvider

`func (o *AddAccessTokenValidator200Response) HasClientSecretPassphraseProvider() bool`

HasClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetIncludeAudParameter

`func (o *AddAccessTokenValidator200Response) GetIncludeAudParameter() bool`

GetIncludeAudParameter returns the IncludeAudParameter field if non-nil, zero value otherwise.

### GetIncludeAudParameterOk

`func (o *AddAccessTokenValidator200Response) GetIncludeAudParameterOk() (*bool, bool)`

GetIncludeAudParameterOk returns a tuple with the IncludeAudParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAudParameter

`func (o *AddAccessTokenValidator200Response) SetIncludeAudParameter(v bool)`

SetIncludeAudParameter sets IncludeAudParameter field to given value.

### HasIncludeAudParameter

`func (o *AddAccessTokenValidator200Response) HasIncludeAudParameter() bool`

HasIncludeAudParameter returns a boolean if a field has been set.

### GetAccessTokenManagerID

`func (o *AddAccessTokenValidator200Response) GetAccessTokenManagerID() string`

GetAccessTokenManagerID returns the AccessTokenManagerID field if non-nil, zero value otherwise.

### GetAccessTokenManagerIDOk

`func (o *AddAccessTokenValidator200Response) GetAccessTokenManagerIDOk() (*string, bool)`

GetAccessTokenManagerIDOk returns a tuple with the AccessTokenManagerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerID

`func (o *AddAccessTokenValidator200Response) SetAccessTokenManagerID(v string)`

SetAccessTokenManagerID sets AccessTokenManagerID field to given value.

### HasAccessTokenManagerID

`func (o *AddAccessTokenValidator200Response) HasAccessTokenManagerID() bool`

HasAccessTokenManagerID returns a boolean if a field has been set.

### GetEndpointCacheRefresh

`func (o *AddAccessTokenValidator200Response) GetEndpointCacheRefresh() string`

GetEndpointCacheRefresh returns the EndpointCacheRefresh field if non-nil, zero value otherwise.

### GetEndpointCacheRefreshOk

`func (o *AddAccessTokenValidator200Response) GetEndpointCacheRefreshOk() (*string, bool)`

GetEndpointCacheRefreshOk returns a tuple with the EndpointCacheRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCacheRefresh

`func (o *AddAccessTokenValidator200Response) SetEndpointCacheRefresh(v string)`

SetEndpointCacheRefresh sets EndpointCacheRefresh field to given value.

### HasEndpointCacheRefresh

`func (o *AddAccessTokenValidator200Response) HasEndpointCacheRefresh() bool`

HasEndpointCacheRefresh returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *AddAccessTokenValidator200Response) GetEvaluationOrderIndex() int64`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *AddAccessTokenValidator200Response) GetEvaluationOrderIndexOk() (*int64, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *AddAccessTokenValidator200Response) SetEvaluationOrderIndex(v int64)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetAuthorizationServer

`func (o *AddAccessTokenValidator200Response) GetAuthorizationServer() string`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *AddAccessTokenValidator200Response) GetAuthorizationServerOk() (*string, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *AddAccessTokenValidator200Response) SetAuthorizationServer(v string)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *AddAccessTokenValidator200Response) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *AddAccessTokenValidator200Response) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddAccessTokenValidator200Response) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddAccessTokenValidator200Response) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *AddAccessTokenValidator200Response) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetSubjectClaimName

`func (o *AddAccessTokenValidator200Response) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *AddAccessTokenValidator200Response) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *AddAccessTokenValidator200Response) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *AddAccessTokenValidator200Response) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetDescription

`func (o *AddAccessTokenValidator200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAccessTokenValidator200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAccessTokenValidator200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAccessTokenValidator200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAccessTokenValidator200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAccessTokenValidator200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAccessTokenValidator200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AddAccessTokenValidator200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddAccessTokenValidator200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddAccessTokenValidator200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddAccessTokenValidator200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddAccessTokenValidator200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddAccessTokenValidator200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddAccessTokenValidator200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddAccessTokenValidator200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetAllowedSigningAlgorithm

`func (o *AddAccessTokenValidator200Response) GetAllowedSigningAlgorithm() []EnumaccessTokenValidatorAllowedSigningAlgorithmProp`

GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field if non-nil, zero value otherwise.

### GetAllowedSigningAlgorithmOk

`func (o *AddAccessTokenValidator200Response) GetAllowedSigningAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedSigningAlgorithmProp, bool)`

GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSigningAlgorithm

`func (o *AddAccessTokenValidator200Response) SetAllowedSigningAlgorithm(v []EnumaccessTokenValidatorAllowedSigningAlgorithmProp)`

SetAllowedSigningAlgorithm sets AllowedSigningAlgorithm field to given value.


### GetSigningCertificate

`func (o *AddAccessTokenValidator200Response) GetSigningCertificate() []string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *AddAccessTokenValidator200Response) GetSigningCertificateOk() (*[]string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *AddAccessTokenValidator200Response) SetSigningCertificate(v []string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *AddAccessTokenValidator200Response) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.

### GetJwksEndpointPath

`func (o *AddAccessTokenValidator200Response) GetJwksEndpointPath() string`

GetJwksEndpointPath returns the JwksEndpointPath field if non-nil, zero value otherwise.

### GetJwksEndpointPathOk

`func (o *AddAccessTokenValidator200Response) GetJwksEndpointPathOk() (*string, bool)`

GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksEndpointPath

`func (o *AddAccessTokenValidator200Response) SetJwksEndpointPath(v string)`

SetJwksEndpointPath sets JwksEndpointPath field to given value.

### HasJwksEndpointPath

`func (o *AddAccessTokenValidator200Response) HasJwksEndpointPath() bool`

HasJwksEndpointPath returns a boolean if a field has been set.

### GetEncryptionKeyPair

`func (o *AddAccessTokenValidator200Response) GetEncryptionKeyPair() string`

GetEncryptionKeyPair returns the EncryptionKeyPair field if non-nil, zero value otherwise.

### GetEncryptionKeyPairOk

`func (o *AddAccessTokenValidator200Response) GetEncryptionKeyPairOk() (*string, bool)`

GetEncryptionKeyPairOk returns a tuple with the EncryptionKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyPair

`func (o *AddAccessTokenValidator200Response) SetEncryptionKeyPair(v string)`

SetEncryptionKeyPair sets EncryptionKeyPair field to given value.

### HasEncryptionKeyPair

`func (o *AddAccessTokenValidator200Response) HasEncryptionKeyPair() bool`

HasEncryptionKeyPair returns a boolean if a field has been set.

### GetAllowedKeyEncryptionAlgorithm

`func (o *AddAccessTokenValidator200Response) GetAllowedKeyEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp`

GetAllowedKeyEncryptionAlgorithm returns the AllowedKeyEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetAllowedKeyEncryptionAlgorithmOk

`func (o *AddAccessTokenValidator200Response) GetAllowedKeyEncryptionAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, bool)`

GetAllowedKeyEncryptionAlgorithmOk returns a tuple with the AllowedKeyEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyEncryptionAlgorithm

`func (o *AddAccessTokenValidator200Response) SetAllowedKeyEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp)`

SetAllowedKeyEncryptionAlgorithm sets AllowedKeyEncryptionAlgorithm field to given value.


### GetAllowedContentEncryptionAlgorithm

`func (o *AddAccessTokenValidator200Response) GetAllowedContentEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp`

GetAllowedContentEncryptionAlgorithm returns the AllowedContentEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetAllowedContentEncryptionAlgorithmOk

`func (o *AddAccessTokenValidator200Response) GetAllowedContentEncryptionAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, bool)`

GetAllowedContentEncryptionAlgorithmOk returns a tuple with the AllowedContentEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedContentEncryptionAlgorithm

`func (o *AddAccessTokenValidator200Response) SetAllowedContentEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp)`

SetAllowedContentEncryptionAlgorithm sets AllowedContentEncryptionAlgorithm field to given value.


### GetClockSkewGracePeriod

`func (o *AddAccessTokenValidator200Response) GetClockSkewGracePeriod() string`

GetClockSkewGracePeriod returns the ClockSkewGracePeriod field if non-nil, zero value otherwise.

### GetClockSkewGracePeriodOk

`func (o *AddAccessTokenValidator200Response) GetClockSkewGracePeriodOk() (*string, bool)`

GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewGracePeriod

`func (o *AddAccessTokenValidator200Response) SetClockSkewGracePeriod(v string)`

SetClockSkewGracePeriod sets ClockSkewGracePeriod field to given value.

### HasClockSkewGracePeriod

`func (o *AddAccessTokenValidator200Response) HasClockSkewGracePeriod() bool`

HasClockSkewGracePeriod returns a boolean if a field has been set.

### GetClientIDClaimName

`func (o *AddAccessTokenValidator200Response) GetClientIDClaimName() string`

GetClientIDClaimName returns the ClientIDClaimName field if non-nil, zero value otherwise.

### GetClientIDClaimNameOk

`func (o *AddAccessTokenValidator200Response) GetClientIDClaimNameOk() (*string, bool)`

GetClientIDClaimNameOk returns a tuple with the ClientIDClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIDClaimName

`func (o *AddAccessTokenValidator200Response) SetClientIDClaimName(v string)`

SetClientIDClaimName sets ClientIDClaimName field to given value.

### HasClientIDClaimName

`func (o *AddAccessTokenValidator200Response) HasClientIDClaimName() bool`

HasClientIDClaimName returns a boolean if a field has been set.

### GetScopeClaimName

`func (o *AddAccessTokenValidator200Response) GetScopeClaimName() string`

GetScopeClaimName returns the ScopeClaimName field if non-nil, zero value otherwise.

### GetScopeClaimNameOk

`func (o *AddAccessTokenValidator200Response) GetScopeClaimNameOk() (*string, bool)`

GetScopeClaimNameOk returns a tuple with the ScopeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClaimName

`func (o *AddAccessTokenValidator200Response) SetScopeClaimName(v string)`

SetScopeClaimName sets ScopeClaimName field to given value.

### HasScopeClaimName

`func (o *AddAccessTokenValidator200Response) HasScopeClaimName() bool`

HasScopeClaimName returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddAccessTokenValidator200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddAccessTokenValidator200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddAccessTokenValidator200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddAccessTokenValidator200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddAccessTokenValidator200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddAccessTokenValidator200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddAccessTokenValidator200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


