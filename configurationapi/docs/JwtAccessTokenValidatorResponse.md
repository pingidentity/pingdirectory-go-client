# JwtAccessTokenValidatorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Access Token Validator | 
**Schemas** | [**[]EnumjwtAccessTokenValidatorSchemaUrn**](EnumjwtAccessTokenValidatorSchemaUrn.md) |  | 
**AllowedSigningAlgorithm** | [**[]EnumaccessTokenValidatorAllowedSigningAlgorithmProp**](EnumaccessTokenValidatorAllowedSigningAlgorithmProp.md) |  | 
**SigningCertificate** | Pointer to **[]string** | Specifies the locally stored certificates that may be used to validate the signature of an incoming JWT access token. If this property is specified, the JWT Access Token Validator will not use a JWKS endpoint to retrieve public keys. | [optional] 
**JwksEndpointPath** | Pointer to **string** | The relative path to JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming JWT access token. This path is relative to the base_url property defined for the validator&#39;s external authorization server. If jwks-endpoint-path is specified, the JWT Access Token Validator will not consult locally stored certificates for validating token signatures. | [optional] 
**EncryptionKeyPair** | Pointer to **string** | The public-private key pair that is used to encrypt the JWT payload. If specified, the JWT Access Token Validator will use the private key to decrypt the JWT payload, and the public key must be exported to the Authorization Server that is issuing access tokens. | [optional] 
**AllowedKeyEncryptionAlgorithm** | [**[]EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp**](EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp.md) |  | 
**AllowedContentEncryptionAlgorithm** | [**[]EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp**](EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp.md) |  | 
**ClockSkewGracePeriod** | Pointer to **string** | Specifies the amount of clock skew that is tolerated by the JWT Access Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token&#39;s not-before (nbf) time and added to the token&#39;s expiration (exp) time, if present, to allow for any time difference between the local server&#39;s clock and the token issuer&#39;s clock. | [optional] 
**ClientIDClaimName** | Pointer to **string** | The name of the token claim that contains the OAuth2 client Id. | [optional] 
**ScopeClaimName** | Pointer to **string** | The name of the token claim that contains the scopes granted by the token. | [optional] 
**EvaluationOrderIndex** | **int32** | When multiple JWT Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all JWT Access Token Validators defined within Directory Server but not necessarily contiguous. JWT Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token. | 
**AuthorizationServer** | Pointer to **string** | Specifies the external server that will be used to aid in validating access tokens. In most cases this will be the Authorization Server that minted the token. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | [optional] 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server. | [optional] 
**Description** | Pointer to **string** | A description for this Access Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this Access Token Validator is enabled for use in Directory Server. | 

## Methods

### NewJwtAccessTokenValidatorResponse

`func NewJwtAccessTokenValidatorResponse(id string, schemas []EnumjwtAccessTokenValidatorSchemaUrn, allowedSigningAlgorithm []EnumaccessTokenValidatorAllowedSigningAlgorithmProp, allowedKeyEncryptionAlgorithm []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, allowedContentEncryptionAlgorithm []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, evaluationOrderIndex int32, enabled bool, ) *JwtAccessTokenValidatorResponse`

NewJwtAccessTokenValidatorResponse instantiates a new JwtAccessTokenValidatorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJwtAccessTokenValidatorResponseWithDefaults

`func NewJwtAccessTokenValidatorResponseWithDefaults() *JwtAccessTokenValidatorResponse`

NewJwtAccessTokenValidatorResponseWithDefaults instantiates a new JwtAccessTokenValidatorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *JwtAccessTokenValidatorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JwtAccessTokenValidatorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JwtAccessTokenValidatorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JwtAccessTokenValidatorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *JwtAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *JwtAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *JwtAccessTokenValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *JwtAccessTokenValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *JwtAccessTokenValidatorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JwtAccessTokenValidatorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JwtAccessTokenValidatorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *JwtAccessTokenValidatorResponse) GetSchemas() []EnumjwtAccessTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JwtAccessTokenValidatorResponse) GetSchemasOk() (*[]EnumjwtAccessTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JwtAccessTokenValidatorResponse) SetSchemas(v []EnumjwtAccessTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowedSigningAlgorithm

`func (o *JwtAccessTokenValidatorResponse) GetAllowedSigningAlgorithm() []EnumaccessTokenValidatorAllowedSigningAlgorithmProp`

GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field if non-nil, zero value otherwise.

### GetAllowedSigningAlgorithmOk

`func (o *JwtAccessTokenValidatorResponse) GetAllowedSigningAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedSigningAlgorithmProp, bool)`

GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSigningAlgorithm

`func (o *JwtAccessTokenValidatorResponse) SetAllowedSigningAlgorithm(v []EnumaccessTokenValidatorAllowedSigningAlgorithmProp)`

SetAllowedSigningAlgorithm sets AllowedSigningAlgorithm field to given value.


### GetSigningCertificate

`func (o *JwtAccessTokenValidatorResponse) GetSigningCertificate() []string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *JwtAccessTokenValidatorResponse) GetSigningCertificateOk() (*[]string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *JwtAccessTokenValidatorResponse) SetSigningCertificate(v []string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *JwtAccessTokenValidatorResponse) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.

### GetJwksEndpointPath

`func (o *JwtAccessTokenValidatorResponse) GetJwksEndpointPath() string`

GetJwksEndpointPath returns the JwksEndpointPath field if non-nil, zero value otherwise.

### GetJwksEndpointPathOk

`func (o *JwtAccessTokenValidatorResponse) GetJwksEndpointPathOk() (*string, bool)`

GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksEndpointPath

`func (o *JwtAccessTokenValidatorResponse) SetJwksEndpointPath(v string)`

SetJwksEndpointPath sets JwksEndpointPath field to given value.

### HasJwksEndpointPath

`func (o *JwtAccessTokenValidatorResponse) HasJwksEndpointPath() bool`

HasJwksEndpointPath returns a boolean if a field has been set.

### GetEncryptionKeyPair

`func (o *JwtAccessTokenValidatorResponse) GetEncryptionKeyPair() string`

GetEncryptionKeyPair returns the EncryptionKeyPair field if non-nil, zero value otherwise.

### GetEncryptionKeyPairOk

`func (o *JwtAccessTokenValidatorResponse) GetEncryptionKeyPairOk() (*string, bool)`

GetEncryptionKeyPairOk returns a tuple with the EncryptionKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKeyPair

`func (o *JwtAccessTokenValidatorResponse) SetEncryptionKeyPair(v string)`

SetEncryptionKeyPair sets EncryptionKeyPair field to given value.

### HasEncryptionKeyPair

`func (o *JwtAccessTokenValidatorResponse) HasEncryptionKeyPair() bool`

HasEncryptionKeyPair returns a boolean if a field has been set.

### GetAllowedKeyEncryptionAlgorithm

`func (o *JwtAccessTokenValidatorResponse) GetAllowedKeyEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp`

GetAllowedKeyEncryptionAlgorithm returns the AllowedKeyEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetAllowedKeyEncryptionAlgorithmOk

`func (o *JwtAccessTokenValidatorResponse) GetAllowedKeyEncryptionAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp, bool)`

GetAllowedKeyEncryptionAlgorithmOk returns a tuple with the AllowedKeyEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedKeyEncryptionAlgorithm

`func (o *JwtAccessTokenValidatorResponse) SetAllowedKeyEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedKeyEncryptionAlgorithmProp)`

SetAllowedKeyEncryptionAlgorithm sets AllowedKeyEncryptionAlgorithm field to given value.


### GetAllowedContentEncryptionAlgorithm

`func (o *JwtAccessTokenValidatorResponse) GetAllowedContentEncryptionAlgorithm() []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp`

GetAllowedContentEncryptionAlgorithm returns the AllowedContentEncryptionAlgorithm field if non-nil, zero value otherwise.

### GetAllowedContentEncryptionAlgorithmOk

`func (o *JwtAccessTokenValidatorResponse) GetAllowedContentEncryptionAlgorithmOk() (*[]EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp, bool)`

GetAllowedContentEncryptionAlgorithmOk returns a tuple with the AllowedContentEncryptionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedContentEncryptionAlgorithm

`func (o *JwtAccessTokenValidatorResponse) SetAllowedContentEncryptionAlgorithm(v []EnumaccessTokenValidatorAllowedContentEncryptionAlgorithmProp)`

SetAllowedContentEncryptionAlgorithm sets AllowedContentEncryptionAlgorithm field to given value.


### GetClockSkewGracePeriod

`func (o *JwtAccessTokenValidatorResponse) GetClockSkewGracePeriod() string`

GetClockSkewGracePeriod returns the ClockSkewGracePeriod field if non-nil, zero value otherwise.

### GetClockSkewGracePeriodOk

`func (o *JwtAccessTokenValidatorResponse) GetClockSkewGracePeriodOk() (*string, bool)`

GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewGracePeriod

`func (o *JwtAccessTokenValidatorResponse) SetClockSkewGracePeriod(v string)`

SetClockSkewGracePeriod sets ClockSkewGracePeriod field to given value.

### HasClockSkewGracePeriod

`func (o *JwtAccessTokenValidatorResponse) HasClockSkewGracePeriod() bool`

HasClockSkewGracePeriod returns a boolean if a field has been set.

### GetClientIDClaimName

`func (o *JwtAccessTokenValidatorResponse) GetClientIDClaimName() string`

GetClientIDClaimName returns the ClientIDClaimName field if non-nil, zero value otherwise.

### GetClientIDClaimNameOk

`func (o *JwtAccessTokenValidatorResponse) GetClientIDClaimNameOk() (*string, bool)`

GetClientIDClaimNameOk returns a tuple with the ClientIDClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIDClaimName

`func (o *JwtAccessTokenValidatorResponse) SetClientIDClaimName(v string)`

SetClientIDClaimName sets ClientIDClaimName field to given value.

### HasClientIDClaimName

`func (o *JwtAccessTokenValidatorResponse) HasClientIDClaimName() bool`

HasClientIDClaimName returns a boolean if a field has been set.

### GetScopeClaimName

`func (o *JwtAccessTokenValidatorResponse) GetScopeClaimName() string`

GetScopeClaimName returns the ScopeClaimName field if non-nil, zero value otherwise.

### GetScopeClaimNameOk

`func (o *JwtAccessTokenValidatorResponse) GetScopeClaimNameOk() (*string, bool)`

GetScopeClaimNameOk returns a tuple with the ScopeClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClaimName

`func (o *JwtAccessTokenValidatorResponse) SetScopeClaimName(v string)`

SetScopeClaimName sets ScopeClaimName field to given value.

### HasScopeClaimName

`func (o *JwtAccessTokenValidatorResponse) HasScopeClaimName() bool`

HasScopeClaimName returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *JwtAccessTokenValidatorResponse) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *JwtAccessTokenValidatorResponse) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *JwtAccessTokenValidatorResponse) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetAuthorizationServer

`func (o *JwtAccessTokenValidatorResponse) GetAuthorizationServer() string`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *JwtAccessTokenValidatorResponse) GetAuthorizationServerOk() (*string, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *JwtAccessTokenValidatorResponse) SetAuthorizationServer(v string)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *JwtAccessTokenValidatorResponse) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *JwtAccessTokenValidatorResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *JwtAccessTokenValidatorResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *JwtAccessTokenValidatorResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *JwtAccessTokenValidatorResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetSubjectClaimName

`func (o *JwtAccessTokenValidatorResponse) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *JwtAccessTokenValidatorResponse) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *JwtAccessTokenValidatorResponse) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *JwtAccessTokenValidatorResponse) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetDescription

`func (o *JwtAccessTokenValidatorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JwtAccessTokenValidatorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JwtAccessTokenValidatorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JwtAccessTokenValidatorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *JwtAccessTokenValidatorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *JwtAccessTokenValidatorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *JwtAccessTokenValidatorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


