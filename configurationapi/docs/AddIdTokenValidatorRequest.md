# AddIdTokenValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidatorName** | **string** | Name of the new ID Token Validator | 
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
**SigningCertificate** | Pointer to **[]string** | Specifies the locally stored certificates that may be used to validate the signature of an incoming ID token. This property may be specified if a JWKS endpoint should not be used to retrieve public signing keys. | [optional] 
**JwksEndpointPath** | Pointer to **string** | The relative path to the JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming ID token. This path is relative to the base_url property defined for the validator&#39;s OpenID Connect provider. If jwks-endpoint-path is specified, the OpenID Connect ID Token Validator will not consult locally stored certificates for validating token signatures. | [optional] 

## Methods

### NewAddIdTokenValidatorRequest

`func NewAddIdTokenValidatorRequest(validatorName string, schemas []EnumopenidConnectIdTokenValidatorSchemaUrn, issuerURL string, openIDConnectProvider string, enabled bool, identityMapper string, evaluationOrderIndex int32, allowedSigningAlgorithm []EnumidTokenValidatorAllowedSigningAlgorithmProp, ) *AddIdTokenValidatorRequest`

NewAddIdTokenValidatorRequest instantiates a new AddIdTokenValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIdTokenValidatorRequestWithDefaults

`func NewAddIdTokenValidatorRequestWithDefaults() *AddIdTokenValidatorRequest`

NewAddIdTokenValidatorRequestWithDefaults instantiates a new AddIdTokenValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidatorName

`func (o *AddIdTokenValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddIdTokenValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddIdTokenValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.


### GetSchemas

`func (o *AddIdTokenValidatorRequest) GetSchemas() []EnumopenidConnectIdTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddIdTokenValidatorRequest) GetSchemasOk() (*[]EnumopenidConnectIdTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddIdTokenValidatorRequest) SetSchemas(v []EnumopenidConnectIdTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIssuerURL

`func (o *AddIdTokenValidatorRequest) GetIssuerURL() string`

GetIssuerURL returns the IssuerURL field if non-nil, zero value otherwise.

### GetIssuerURLOk

`func (o *AddIdTokenValidatorRequest) GetIssuerURLOk() (*string, bool)`

GetIssuerURLOk returns a tuple with the IssuerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerURL

`func (o *AddIdTokenValidatorRequest) SetIssuerURL(v string)`

SetIssuerURL sets IssuerURL field to given value.


### GetOpenIDConnectProvider

`func (o *AddIdTokenValidatorRequest) GetOpenIDConnectProvider() string`

GetOpenIDConnectProvider returns the OpenIDConnectProvider field if non-nil, zero value otherwise.

### GetOpenIDConnectProviderOk

`func (o *AddIdTokenValidatorRequest) GetOpenIDConnectProviderOk() (*string, bool)`

GetOpenIDConnectProviderOk returns a tuple with the OpenIDConnectProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectProvider

`func (o *AddIdTokenValidatorRequest) SetOpenIDConnectProvider(v string)`

SetOpenIDConnectProvider sets OpenIDConnectProvider field to given value.


### GetOpenIDConnectMetadataCacheDuration

`func (o *AddIdTokenValidatorRequest) GetOpenIDConnectMetadataCacheDuration() string`

GetOpenIDConnectMetadataCacheDuration returns the OpenIDConnectMetadataCacheDuration field if non-nil, zero value otherwise.

### GetOpenIDConnectMetadataCacheDurationOk

`func (o *AddIdTokenValidatorRequest) GetOpenIDConnectMetadataCacheDurationOk() (*string, bool)`

GetOpenIDConnectMetadataCacheDurationOk returns a tuple with the OpenIDConnectMetadataCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectMetadataCacheDuration

`func (o *AddIdTokenValidatorRequest) SetOpenIDConnectMetadataCacheDuration(v string)`

SetOpenIDConnectMetadataCacheDuration sets OpenIDConnectMetadataCacheDuration field to given value.

### HasOpenIDConnectMetadataCacheDuration

`func (o *AddIdTokenValidatorRequest) HasOpenIDConnectMetadataCacheDuration() bool`

HasOpenIDConnectMetadataCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AddIdTokenValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddIdTokenValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddIdTokenValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddIdTokenValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIdTokenValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIdTokenValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIdTokenValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdentityMapper

`func (o *AddIdTokenValidatorRequest) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddIdTokenValidatorRequest) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddIdTokenValidatorRequest) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetSubjectClaimName

`func (o *AddIdTokenValidatorRequest) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *AddIdTokenValidatorRequest) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *AddIdTokenValidatorRequest) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *AddIdTokenValidatorRequest) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetClockSkewGracePeriod

`func (o *AddIdTokenValidatorRequest) GetClockSkewGracePeriod() string`

GetClockSkewGracePeriod returns the ClockSkewGracePeriod field if non-nil, zero value otherwise.

### GetClockSkewGracePeriodOk

`func (o *AddIdTokenValidatorRequest) GetClockSkewGracePeriodOk() (*string, bool)`

GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewGracePeriod

`func (o *AddIdTokenValidatorRequest) SetClockSkewGracePeriod(v string)`

SetClockSkewGracePeriod sets ClockSkewGracePeriod field to given value.

### HasClockSkewGracePeriod

`func (o *AddIdTokenValidatorRequest) HasClockSkewGracePeriod() bool`

HasClockSkewGracePeriod returns a boolean if a field has been set.

### GetJwksCacheDuration

`func (o *AddIdTokenValidatorRequest) GetJwksCacheDuration() string`

GetJwksCacheDuration returns the JwksCacheDuration field if non-nil, zero value otherwise.

### GetJwksCacheDurationOk

`func (o *AddIdTokenValidatorRequest) GetJwksCacheDurationOk() (*string, bool)`

GetJwksCacheDurationOk returns a tuple with the JwksCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksCacheDuration

`func (o *AddIdTokenValidatorRequest) SetJwksCacheDuration(v string)`

SetJwksCacheDuration sets JwksCacheDuration field to given value.

### HasJwksCacheDuration

`func (o *AddIdTokenValidatorRequest) HasJwksCacheDuration() bool`

HasJwksCacheDuration returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *AddIdTokenValidatorRequest) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *AddIdTokenValidatorRequest) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *AddIdTokenValidatorRequest) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetAllowedSigningAlgorithm

`func (o *AddIdTokenValidatorRequest) GetAllowedSigningAlgorithm() []EnumidTokenValidatorAllowedSigningAlgorithmProp`

GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field if non-nil, zero value otherwise.

### GetAllowedSigningAlgorithmOk

`func (o *AddIdTokenValidatorRequest) GetAllowedSigningAlgorithmOk() (*[]EnumidTokenValidatorAllowedSigningAlgorithmProp, bool)`

GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSigningAlgorithm

`func (o *AddIdTokenValidatorRequest) SetAllowedSigningAlgorithm(v []EnumidTokenValidatorAllowedSigningAlgorithmProp)`

SetAllowedSigningAlgorithm sets AllowedSigningAlgorithm field to given value.


### GetSigningCertificate

`func (o *AddIdTokenValidatorRequest) GetSigningCertificate() []string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *AddIdTokenValidatorRequest) GetSigningCertificateOk() (*[]string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *AddIdTokenValidatorRequest) SetSigningCertificate(v []string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *AddIdTokenValidatorRequest) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.

### GetJwksEndpointPath

`func (o *AddIdTokenValidatorRequest) GetJwksEndpointPath() string`

GetJwksEndpointPath returns the JwksEndpointPath field if non-nil, zero value otherwise.

### GetJwksEndpointPathOk

`func (o *AddIdTokenValidatorRequest) GetJwksEndpointPathOk() (*string, bool)`

GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksEndpointPath

`func (o *AddIdTokenValidatorRequest) SetJwksEndpointPath(v string)`

SetJwksEndpointPath sets JwksEndpointPath field to given value.

### HasJwksEndpointPath

`func (o *AddIdTokenValidatorRequest) HasJwksEndpointPath() bool`

HasJwksEndpointPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


