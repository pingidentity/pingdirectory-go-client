# AddAccessTokenValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidatorName** | **string** | Name of the new Access Token Validator | 
**Schemas** | [**[]EnumthirdPartyAccessTokenValidatorSchemaUrn**](EnumthirdPartyAccessTokenValidatorSchemaUrn.md) |  | 
**ClientID** | **string** | The client identifier to use when authenticating to the PingFederate authorization server. | 
**ClientSecret** | Pointer to **string** | The client secret to use when authenticating to the PingFederate authorization server. | [optional] 
**ClientSecretPassphraseProvider** | Pointer to **string** | The passphrase provider for obtaining the client secret to use when authenticating to the PingFederate authorization server. | [optional] 
**IncludeAudParameter** | Pointer to **bool** | Whether to include the incoming request URL as the \&quot;aud\&quot; parameter when calling the PingFederate introspection endpoint. This property is ignored if the access-token-manager-id property is set. | [optional] 
**AccessTokenManagerID** | Pointer to **string** | The Access Token Manager instance ID to specify when calling the PingFederate introspection endpoint. If this property is set the include-aud-parameter property is ignored. | [optional] 
**EndpointCacheRefresh** | Pointer to **string** | How often the Access Token Validator should refresh its stored value of the PingFederate server&#39;s token introspection endpoint. | [optional] 
**EvaluationOrderIndex** | **int32** | When multiple Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Access Token Validators defined within Directory Server but not necessarily contiguous. Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token. | 
**AuthorizationServer** | Pointer to **string** | Specifies the external server that will be used to aid in validating access tokens. In most cases this will be the Authorization Server that minted the token. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | [optional] 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server. | [optional] 
**Description** | Pointer to **string** | A description for this Access Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this Access Token Validator is enabled for use in Directory Server. | 
**AllowedSigningAlgorithm** | [**[]EnumaccessTokenValidatorAllowedSigningAlgorithmProp**](EnumaccessTokenValidatorAllowedSigningAlgorithmProp.md) |  | 
**SigningCertificate** | Pointer to **[]string** |  | [optional] 
**JwksEndpointPath** | Pointer to **string** | The relative path to JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming JWT access token. This path is relative to the base_url property defined for the validator&#39;s external authorization server. If jwks-endpoint-path is specified, the JWT Access Token Validator will not consult locally stored certificates for validating token signatures. | [optional] 
**EncryptionKeyPair** | Pointer to **string** | The public-private key pair that is used to encrypt the JWT payload. If specified, the JWT Access Token Validator will use the private key to decrypt the JWT payload, and the public key must be exported to the Authorization Server that is issuing access tokens. | [optional] 
**AllowedKeyEncryptionAlgorithm** | [**[]EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp**](EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp.md) |  | 
**AllowedContentEncryptionAlgorithm** | [**[]EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp**](EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp.md) |  | 
**ClockSkewGracePeriod** | Pointer to **string** | Specifies the amount of clock skew that is tolerated by the JWT Access Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token&#39;s not-before (nbf) time and added to the token&#39;s expiration (exp) time, if present, to allow for any time difference between the local server&#39;s clock and the token issuer&#39;s clock. | [optional] 
**ClientIDClaimName** | Pointer to **string** | The name of the token claim that contains the OAuth2 client ID. | [optional] 
**ScopeClaimName** | Pointer to **string** | The name of the token claim that contains the scopes granted by the token. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Access Token Validator. | 
**ExtensionArgument** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAddAccessTokenValidatorRequest

`func NewAddAccessTokenValidatorRequest(validatorName string, schemas []EnumthirdPartyAccessTokenValidatorSchemaUrn, clientID string, evaluationOrderIndex int32, enabled bool, allowedSigningAlgorithm []EnumaccessTokenValidatorAllowedSigningAlgorithmProp, allowedKeyEncryptionAlgorithm []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, allowedContentEncryptionAlgorithm []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, extensionClass string, ) *AddAccessTokenValidatorRequest`

NewAddAccessTokenValidatorRequest instantiates a new AddAccessTokenValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAccessTokenValidatorRequestWithDefaults

`func NewAddAccessTokenValidatorRequestWithDefaults() *AddAccessTokenValidatorRequest`

NewAddAccessTokenValidatorRequestWithDefaults instantiates a new AddAccessTokenValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidatorName

`func (o *AddAccessTokenValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddAccessTokenValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddAccessTokenValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.


### GetSchemas

`func (o *AddAccessTokenValidatorRequest) GetSchemas() []EnumthirdPartyAccessTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddAccessTokenValidatorRequest) GetSchemasOk() (*[]EnumthirdPartyAccessTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddAccessTokenValidatorRequest) SetSchemas(v []EnumthirdPartyAccessTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetClientID

`func (o *AddAccessTokenValidatorRequest) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *AddAccessTokenValidatorRequest) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *AddAccessTokenValidatorRequest) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *AddAccessTokenValidatorRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AddAccessTokenValidatorRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AddAccessTokenValidatorRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AddAccessTokenValidatorRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecretPassphraseProvider

`func (o *AddAccessTokenValidatorRequest) GetClientSecretPassphraseProvider() string`

GetClientSecretPassphraseProvider returns the ClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetClientSecretPassphraseProviderOk

`func (o *AddAccessTokenValidatorRequest) GetClientSecretPassphraseProviderOk() (*string, bool)`

GetClientSecretPassphraseProviderOk returns a tuple with the ClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretPassphraseProvider

`func (o *AddAccessTokenValidatorRequest) SetClientSecretPassphraseProvider(v string)`

SetClientSecretPassphraseProvider sets ClientSecretPassphraseProvider field to given value.

### HasClientSecretPassphraseProvider

`func (o *AddAccessTokenValidatorRequest) HasClientSecretPassphraseProvider() bool`

HasClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetIncludeAudParameter

`func (o *AddAccessTokenValidatorRequest) GetIncludeAudParameter() bool`

GetIncludeAudParameter returns the IncludeAudParameter field if non-nil, zero value otherwise.

### GetIncludeAudParameterOk

`func (o *AddAccessTokenValidatorRequest) GetIncludeAudParameterOk() (*bool, bool)`

GetIncludeAudParameterOk returns a tuple with the IncludeAudParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAudParameter

`func (o *AddAccessTokenValidatorRequest) SetIncludeAudParameter(v bool)`

SetIncludeAudParameter sets IncludeAudParameter field to given value.

### HasIncludeAudParameter

`func (o *AddAccessTokenValidatorRequest) HasIncludeAudParameter() bool`

HasIncludeAudParameter returns a boolean if a field has been set.

### GetAccessTokenManagerID

`func (o *AddAccessTokenValidatorRequest) GetAccessTokenManagerID() string`

GetAccessTokenManagerID returns the AccessTokenManagerID field if non-nil, zero value otherwise.

### GetAccessTokenManagerIDOk

`func (o *AddAccessTokenValidatorRequest) GetAccessTokenManagerIDOk() (*string, bool)`

GetAccessTokenManagerIDOk returns a tuple with the AccessTokenManagerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerID

`func (o *AddAccessTokenValidatorRequest) SetAccessTokenManagerID(v string)`

SetAccessTokenManagerID sets AccessTokenManagerID field to given value.

### HasAccessTokenManagerID

`func (o *AddAccessTokenValidatorRequest) HasAccessTokenManagerID() bool`

HasAccessTokenManagerID returns a boolean if a field has been set.

### GetEndpointCacheRefresh

`func (o *AddAccessTokenValidatorRequest) GetEndpointCacheRefresh() string`

GetEndpointCacheRefresh returns the EndpointCacheRefresh field if non-nil, zero value otherwise.

### GetEndpointCacheRefreshOk

`func (o *AddAccessTokenValidatorRequest) GetEndpointCacheRefreshOk() (*string, bool)`

GetEndpointCacheRefreshOk returns a tuple with the EndpointCacheRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCacheRefresh

`func (o *AddAccessTokenValidatorRequest) SetEndpointCacheRefresh(v string)`

SetEndpointCacheRefresh sets EndpointCacheRefresh field to given value.

### HasEndpointCacheRefresh

`func (o *AddAccessTokenValidatorRequest) HasEndpointCacheRefresh() bool`

HasEndpointCacheRefresh returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *AddAccessTokenValidatorRequest) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *AddAccessTokenValidatorRequest) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *AddAccessTokenValidatorRequest) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetAuthorizationServer

`func (o *AddAccessTokenValidatorRequest) GetAuthorizationServer() string`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *AddAccessTokenValidatorRequest) GetAuthorizationServerOk() (*string, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *AddAccessTokenValidatorRequest) SetAuthorizationServer(v string)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *AddAccessTokenValidatorRequest) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *AddAccessTokenValidatorRequest) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddAccessTokenValidatorRequest) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddAccessTokenValidatorRequest) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *AddAccessTokenValidatorRequest) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetSubjectClaimName

`func (o *AddAccessTokenValidatorRequest) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *AddAccessTokenValidatorRequest) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *AddAccessTokenValidatorRequest) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *AddAccessTokenValidatorRequest) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetDescription

`func (o *AddAccessTokenValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddAccessTokenValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddAccessTokenValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddAccessTokenValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddAccessTokenValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddAccessTokenValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddAccessTokenValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowedSigningAlgorithm

`func (o *AddAccessTokenValidatorRequest) GetAllowedSigningAlgorithm() []EnumaccessTokenValidatorAllowedSigningAlgorithmProp`

GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field if non-nil, zero value otherwise.

### GetAllowedSigningAlgorithmOk

`func (o *AddAccessTokenValidatorRequest) GetAllowedSigningAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedSigningAlgorithmProp, bool)`

GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSigningAlgorithm

`func (o *AddAccessTokenValidatorRequest) SetAllowedSigningAlgorithm(v []EnumaccessTokenValidatorAllowedSigningAlgorithmProp)`

SetAllowedSigningAlgorithm sets AllowedSigningAlgorithm field to given value.


### GetSigningCertificate

`func (o *AddAccessTokenValidatorRequest) GetSigningCertificate() []string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *AddAccessTokenValidatorRequest) GetSigningCertificateOk() (*[]string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *AddAccessTokenValidatorRequest) SetSigningCertificate(v []string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *AddAccessTokenValidatorRequest) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.

### GetJwksEndpointPath

`func (o *AddAccessTokenValidatorRequest) GetJwksEndpointPath() string`

GetJwksEndpointPath returns the JwksEndpointPath field if non-nil, zero value otherwise.

### GetJwksEndpointPathOk

`func (o *AddAccessTokenValidatorRequest) GetJwksEndpointPathOk() (*string, bool)`

GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksEndpointPath

`func (o *AddAccessTokenValidatorRequest) SetJwksEndpointPath(v string)`

SetJwksEndpointPath sets JwksEndpointPath field to given value.

### HasJwksEndpointPath

`func (o *AddAccessTokenValidatorRequest) HasJwksEndpointPath() bool`

HasJwksEndpointPath returns a boolean if a field has been set.

### GetEncryptionKeyPair

`func (o *AddAccessTokenValidatorRequest) GetEncryptionKeyPair() string`

GetEncryptionKeyPair returns the EncryptionKeyPair field if non-nil, zero value otherwise.

### GetEncryptionKeyPairOk

`func (o *AddAccessTokenValidatorRequest) GetEncryptionKeyPairOk() (*string, bool)`

GetEncryptionKeyPairOk returns a tuple with the EncryptionKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyPair

`func (o *AddAccessTokenValidatorRequest) SetEncryptionKeyPair(v string)`

SetEncryptionKeyPair sets EncryptionKeyPair field to given value.

### HasEncryptionKeyPair

`func (o *AddAccessTokenValidatorRequest) HasEncryptionKeyPair() bool`

HasEncryptionKeyPair returns a boolean if a field has been set.

### GetAllowedKeyEncryptionAlgorithm

`func (o *AddAccessTokenValidatorRequest) GetAllowedKeyEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp`

GetAllowedKeyEncryptionAlgorithm returns the AllowedKeyEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetAllowedKeyEncryptionAlgorithmOk

`func (o *AddAccessTokenValidatorRequest) GetAllowedKeyEncryptionAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, bool)`

GetAllowedKeyEncryptionAlgorithmOk returns a tuple with the AllowedKeyEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyEncryptionAlgorithm

`func (o *AddAccessTokenValidatorRequest) SetAllowedKeyEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp)`

SetAllowedKeyEncryptionAlgorithm sets AllowedKeyEncryptionAlgorithm field to given value.


### GetAllowedContentEncryptionAlgorithm

`func (o *AddAccessTokenValidatorRequest) GetAllowedContentEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp`

GetAllowedContentEncryptionAlgorithm returns the AllowedContentEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetAllowedContentEncryptionAlgorithmOk

`func (o *AddAccessTokenValidatorRequest) GetAllowedContentEncryptionAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, bool)`

GetAllowedContentEncryptionAlgorithmOk returns a tuple with the AllowedContentEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedContentEncryptionAlgorithm

`func (o *AddAccessTokenValidatorRequest) SetAllowedContentEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp)`

SetAllowedContentEncryptionAlgorithm sets AllowedContentEncryptionAlgorithm field to given value.


### GetClockSkewGracePeriod

`func (o *AddAccessTokenValidatorRequest) GetClockSkewGracePeriod() string`

GetClockSkewGracePeriod returns the ClockSkewGracePeriod field if non-nil, zero value otherwise.

### GetClockSkewGracePeriodOk

`func (o *AddAccessTokenValidatorRequest) GetClockSkewGracePeriodOk() (*string, bool)`

GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewGracePeriod

`func (o *AddAccessTokenValidatorRequest) SetClockSkewGracePeriod(v string)`

SetClockSkewGracePeriod sets ClockSkewGracePeriod field to given value.

### HasClockSkewGracePeriod

`func (o *AddAccessTokenValidatorRequest) HasClockSkewGracePeriod() bool`

HasClockSkewGracePeriod returns a boolean if a field has been set.

### GetClientIDClaimName

`func (o *AddAccessTokenValidatorRequest) GetClientIDClaimName() string`

GetClientIDClaimName returns the ClientIDClaimName field if non-nil, zero value otherwise.

### GetClientIDClaimNameOk

`func (o *AddAccessTokenValidatorRequest) GetClientIDClaimNameOk() (*string, bool)`

GetClientIDClaimNameOk returns a tuple with the ClientIDClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIDClaimName

`func (o *AddAccessTokenValidatorRequest) SetClientIDClaimName(v string)`

SetClientIDClaimName sets ClientIDClaimName field to given value.

### HasClientIDClaimName

`func (o *AddAccessTokenValidatorRequest) HasClientIDClaimName() bool`

HasClientIDClaimName returns a boolean if a field has been set.

### GetScopeClaimName

`func (o *AddAccessTokenValidatorRequest) GetScopeClaimName() string`

GetScopeClaimName returns the ScopeClaimName field if non-nil, zero value otherwise.

### GetScopeClaimNameOk

`func (o *AddAccessTokenValidatorRequest) GetScopeClaimNameOk() (*string, bool)`

GetScopeClaimNameOk returns a tuple with the ScopeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClaimName

`func (o *AddAccessTokenValidatorRequest) SetScopeClaimName(v string)`

SetScopeClaimName sets ScopeClaimName field to given value.

### HasScopeClaimName

`func (o *AddAccessTokenValidatorRequest) HasScopeClaimName() bool`

HasScopeClaimName returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddAccessTokenValidatorRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddAccessTokenValidatorRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddAccessTokenValidatorRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddAccessTokenValidatorRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddAccessTokenValidatorRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddAccessTokenValidatorRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddAccessTokenValidatorRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


