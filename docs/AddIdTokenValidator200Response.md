# AddIdTokenValidator200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the ID Token Validator | 
**Schemas** | [**[]EnumopenidConnectIdTokenValidatorSchemaUrn**](EnumopenidConnectIdTokenValidatorSchemaUrn.md) |  | 
**IssuerURL** | **string** | Specifies the OpenID Connect provider&#39;s issuer URL. | 
**OpenIDConnectProvider** | **string** | Specifies the OpenID Connect provider that issues ID tokens handled by this OpenID Connect ID Token Validator. This property is used in conjunction with the jwks-endpoint-path property. | 
**OpenIDConnectMetadataCacheDuration** | Pointer to **string** | How often the PingOne ID Token Validator should refresh its stored cache of OpenID Connect-related metadata. | [optional] 
**Description** | Pointer to **string** | A description for this ID Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this ID Token Validator is enabled for use in the Directory Server. | 
**IdentityMapper** | **string** | Specifies the name of the Identity Mapper that should be used to correlate an ID token subject value to a user entry. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject; i.e., the authenticated user. | [optional] 
**ClockSkewGracePeriod** | Pointer to **string** | Specifies the amount of clock skew that is tolerated by the ID Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token&#39;s not-before (nbf) time and added to the token&#39;s expiration (exp) time, if present, to allow for any time difference between the local server&#39;s clock and the token issuer&#39;s clock. | [optional] 
**JwksCacheDuration** | Pointer to **string** | How often the ID Token Validator should refresh its cache of JWKS token signing keys. | [optional] 
**EvaluationOrderIndex** | **int32** | When multiple ID Token Validators are defined for a single Directory Server, this property determines the order in which the ID Token Validators are consulted. Values of this property must be unique among all ID Token Validators defined within Directory Server but not necessarily contiguous. ID Token Validators with lower values will be evaluated first to determine if they are able to validate the ID token. | 
**AllowedSigningAlgorithm** | [**[]EnumidTokenValidatorAllowedSigningAlgorithmProp**](EnumidTokenValidatorAllowedSigningAlgorithmProp.md) |  | 
**SigningCertificate** | Pointer to **[]string** |  | [optional] 
**JwksEndpointPath** | Pointer to **string** | The relative path to the JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming ID token. This path is relative to the base_url property defined for the validator&#39;s OpenID Connect provider. If jwks-endpoint-path is specified, the OpenID Connect ID Token Validator will not consult locally stored certificates for validating token signatures. | [optional] 

## Methods

### NewAddIdTokenValidator200Response

`func NewAddIdTokenValidator200Response(id string, schemas []EnumopenidConnectIdTokenValidatorSchemaUrn, issuerURL string, openIDConnectProvider string, enabled bool, identityMapper string, evaluationOrderIndex int32, allowedSigningAlgorithm []EnumidTokenValidatorAllowedSigningAlgorithmProp, ) *AddIdTokenValidator200Response`

NewAddIdTokenValidator200Response instantiates a new AddIdTokenValidator200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIdTokenValidator200ResponseWithDefaults

`func NewAddIdTokenValidator200ResponseWithDefaults() *AddIdTokenValidator200Response`

NewAddIdTokenValidator200ResponseWithDefaults instantiates a new AddIdTokenValidator200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddIdTokenValidator200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddIdTokenValidator200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddIdTokenValidator200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddIdTokenValidator200Response) GetSchemas() []EnumopenidConnectIdTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddIdTokenValidator200Response) GetSchemasOk() (*[]EnumopenidConnectIdTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddIdTokenValidator200Response) SetSchemas(v []EnumopenidConnectIdTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIssuerURL

`func (o *AddIdTokenValidator200Response) GetIssuerURL() string`

GetIssuerURL returns the IssuerURL field if non-nil, zero value otherwise.

### GetIssuerURLOk

`func (o *AddIdTokenValidator200Response) GetIssuerURLOk() (*string, bool)`

GetIssuerURLOk returns a tuple with the IssuerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerURL

`func (o *AddIdTokenValidator200Response) SetIssuerURL(v string)`

SetIssuerURL sets IssuerURL field to given value.


### GetOpenIDConnectProvider

`func (o *AddIdTokenValidator200Response) GetOpenIDConnectProvider() string`

GetOpenIDConnectProvider returns the OpenIDConnectProvider field if non-nil, zero value otherwise.

### GetOpenIDConnectProviderOk

`func (o *AddIdTokenValidator200Response) GetOpenIDConnectProviderOk() (*string, bool)`

GetOpenIDConnectProviderOk returns a tuple with the OpenIDConnectProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectProvider

`func (o *AddIdTokenValidator200Response) SetOpenIDConnectProvider(v string)`

SetOpenIDConnectProvider sets OpenIDConnectProvider field to given value.


### GetOpenIDConnectMetadataCacheDuration

`func (o *AddIdTokenValidator200Response) GetOpenIDConnectMetadataCacheDuration() string`

GetOpenIDConnectMetadataCacheDuration returns the OpenIDConnectMetadataCacheDuration field if non-nil, zero value otherwise.

### GetOpenIDConnectMetadataCacheDurationOk

`func (o *AddIdTokenValidator200Response) GetOpenIDConnectMetadataCacheDurationOk() (*string, bool)`

GetOpenIDConnectMetadataCacheDurationOk returns a tuple with the OpenIDConnectMetadataCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectMetadataCacheDuration

`func (o *AddIdTokenValidator200Response) SetOpenIDConnectMetadataCacheDuration(v string)`

SetOpenIDConnectMetadataCacheDuration sets OpenIDConnectMetadataCacheDuration field to given value.

### HasOpenIDConnectMetadataCacheDuration

`func (o *AddIdTokenValidator200Response) HasOpenIDConnectMetadataCacheDuration() bool`

HasOpenIDConnectMetadataCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AddIdTokenValidator200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddIdTokenValidator200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddIdTokenValidator200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddIdTokenValidator200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIdTokenValidator200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIdTokenValidator200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIdTokenValidator200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdentityMapper

`func (o *AddIdTokenValidator200Response) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddIdTokenValidator200Response) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddIdTokenValidator200Response) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetSubjectClaimName

`func (o *AddIdTokenValidator200Response) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *AddIdTokenValidator200Response) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *AddIdTokenValidator200Response) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *AddIdTokenValidator200Response) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetClockSkewGracePeriod

`func (o *AddIdTokenValidator200Response) GetClockSkewGracePeriod() string`

GetClockSkewGracePeriod returns the ClockSkewGracePeriod field if non-nil, zero value otherwise.

### GetClockSkewGracePeriodOk

`func (o *AddIdTokenValidator200Response) GetClockSkewGracePeriodOk() (*string, bool)`

GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewGracePeriod

`func (o *AddIdTokenValidator200Response) SetClockSkewGracePeriod(v string)`

SetClockSkewGracePeriod sets ClockSkewGracePeriod field to given value.

### HasClockSkewGracePeriod

`func (o *AddIdTokenValidator200Response) HasClockSkewGracePeriod() bool`

HasClockSkewGracePeriod returns a boolean if a field has been set.

### GetJwksCacheDuration

`func (o *AddIdTokenValidator200Response) GetJwksCacheDuration() string`

GetJwksCacheDuration returns the JwksCacheDuration field if non-nil, zero value otherwise.

### GetJwksCacheDurationOk

`func (o *AddIdTokenValidator200Response) GetJwksCacheDurationOk() (*string, bool)`

GetJwksCacheDurationOk returns a tuple with the JwksCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksCacheDuration

`func (o *AddIdTokenValidator200Response) SetJwksCacheDuration(v string)`

SetJwksCacheDuration sets JwksCacheDuration field to given value.

### HasJwksCacheDuration

`func (o *AddIdTokenValidator200Response) HasJwksCacheDuration() bool`

HasJwksCacheDuration returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *AddIdTokenValidator200Response) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *AddIdTokenValidator200Response) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *AddIdTokenValidator200Response) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetAllowedSigningAlgorithm

`func (o *AddIdTokenValidator200Response) GetAllowedSigningAlgorithm() []EnumidTokenValidatorAllowedSigningAlgorithmProp`

GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field if non-nil, zero value otherwise.

### GetAllowedSigningAlgorithmOk

`func (o *AddIdTokenValidator200Response) GetAllowedSigningAlgorithmOk() (*[]EnumidTokenValidatorAllowedSigningAlgorithmProp, bool)`

GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSigningAlgorithm

`func (o *AddIdTokenValidator200Response) SetAllowedSigningAlgorithm(v []EnumidTokenValidatorAllowedSigningAlgorithmProp)`

SetAllowedSigningAlgorithm sets AllowedSigningAlgorithm field to given value.


### GetSigningCertificate

`func (o *AddIdTokenValidator200Response) GetSigningCertificate() []string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *AddIdTokenValidator200Response) GetSigningCertificateOk() (*[]string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *AddIdTokenValidator200Response) SetSigningCertificate(v []string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *AddIdTokenValidator200Response) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.

### GetJwksEndpointPath

`func (o *AddIdTokenValidator200Response) GetJwksEndpointPath() string`

GetJwksEndpointPath returns the JwksEndpointPath field if non-nil, zero value otherwise.

### GetJwksEndpointPathOk

`func (o *AddIdTokenValidator200Response) GetJwksEndpointPathOk() (*string, bool)`

GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksEndpointPath

`func (o *AddIdTokenValidator200Response) SetJwksEndpointPath(v string)`

SetJwksEndpointPath sets JwksEndpointPath field to given value.

### HasJwksEndpointPath

`func (o *AddIdTokenValidator200Response) HasJwksEndpointPath() bool`

HasJwksEndpointPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


