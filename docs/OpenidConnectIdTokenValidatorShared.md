# OpenidConnectIdTokenValidatorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumopenidConnectIdTokenValidatorSchemaUrn**](EnumopenidConnectIdTokenValidatorSchemaUrn.md) |  | 
**AllowedSigningAlgorithm** | [**[]EnumidTokenValidatorAllowedSigningAlgorithmProp**](EnumidTokenValidatorAllowedSigningAlgorithmProp.md) |  | 
**SigningCertificate** | Pointer to **[]string** |  | [optional] 
**OpenIDConnectProvider** | Pointer to **string** | Specifies the OpenID Connect provider that issues ID tokens handled by this OpenID Connect ID Token Validator. This property is used in conjunction with the jwks-endpoint-path property. | [optional] 
**JwksEndpointPath** | Pointer to **string** | The relative path to the JWKS endpoint from which to retrieve one or more public signing keys that may be used to validate the signature of an incoming ID token. This path is relative to the base_url property defined for the validator&#39;s OpenID Connect provider. If jwks-endpoint-path is specified, the OpenID Connect ID Token Validator will not consult locally stored certificates for validating token signatures. | [optional] 
**Description** | Pointer to **string** | A description for this ID Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this ID Token Validator is enabled for use in the Directory Server. | 
**IdentityMapper** | **string** | Specifies the name of the Identity Mapper that should be used to correlate an ID token subject value to a user entry. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject; i.e., the authenticated user. | [optional] 
**IssuerURL** | **string** | Specifies the OpenID Connect provider&#39;s issuer URL. | 
**ClockSkewGracePeriod** | Pointer to **string** | Specifies the amount of clock skew that is tolerated by the ID Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token&#39;s not-before (nbf) time and added to the token&#39;s expiration (exp) time, if present, to allow for any time difference between the local server&#39;s clock and the token issuer&#39;s clock. | [optional] 
**JwksCacheDuration** | Pointer to **string** | How often the ID Token Validator should refresh its cache of JWKS token signing keys. | [optional] 
**EvaluationOrderIndex** | **int32** | When multiple ID Token Validators are defined for a single Directory Server, this property determines the order in which the ID Token Validators are consulted. Values of this property must be unique among all ID Token Validators defined within Directory Server but not necessarily contiguous. ID Token Validators with lower values will be evaluated first to determine if they are able to validate the ID token. | 

## Methods

### NewOpenidConnectIdTokenValidatorShared

`func NewOpenidConnectIdTokenValidatorShared(schemas []EnumopenidConnectIdTokenValidatorSchemaUrn, allowedSigningAlgorithm []EnumidTokenValidatorAllowedSigningAlgorithmProp, enabled bool, identityMapper string, issuerURL string, evaluationOrderIndex int32, ) *OpenidConnectIdTokenValidatorShared`

NewOpenidConnectIdTokenValidatorShared instantiates a new OpenidConnectIdTokenValidatorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenidConnectIdTokenValidatorSharedWithDefaults

`func NewOpenidConnectIdTokenValidatorSharedWithDefaults() *OpenidConnectIdTokenValidatorShared`

NewOpenidConnectIdTokenValidatorSharedWithDefaults instantiates a new OpenidConnectIdTokenValidatorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *OpenidConnectIdTokenValidatorShared) GetSchemas() []EnumopenidConnectIdTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *OpenidConnectIdTokenValidatorShared) GetSchemasOk() (*[]EnumopenidConnectIdTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *OpenidConnectIdTokenValidatorShared) SetSchemas(v []EnumopenidConnectIdTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAllowedSigningAlgorithm

`func (o *OpenidConnectIdTokenValidatorShared) GetAllowedSigningAlgorithm() []EnumidTokenValidatorAllowedSigningAlgorithmProp`

GetAllowedSigningAlgorithm returns the AllowedSigningAlgorithm field if non-nil, zero value otherwise.

### GetAllowedSigningAlgorithmOk

`func (o *OpenidConnectIdTokenValidatorShared) GetAllowedSigningAlgorithmOk() (*[]EnumidTokenValidatorAllowedSigningAlgorithmProp, bool)`

GetAllowedSigningAlgorithmOk returns a tuple with the AllowedSigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSigningAlgorithm

`func (o *OpenidConnectIdTokenValidatorShared) SetAllowedSigningAlgorithm(v []EnumidTokenValidatorAllowedSigningAlgorithmProp)`

SetAllowedSigningAlgorithm sets AllowedSigningAlgorithm field to given value.


### GetSigningCertificate

`func (o *OpenidConnectIdTokenValidatorShared) GetSigningCertificate() []string`

GetSigningCertificate returns the SigningCertificate field if non-nil, zero value otherwise.

### GetSigningCertificateOk

`func (o *OpenidConnectIdTokenValidatorShared) GetSigningCertificateOk() (*[]string, bool)`

GetSigningCertificateOk returns a tuple with the SigningCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertificate

`func (o *OpenidConnectIdTokenValidatorShared) SetSigningCertificate(v []string)`

SetSigningCertificate sets SigningCertificate field to given value.

### HasSigningCertificate

`func (o *OpenidConnectIdTokenValidatorShared) HasSigningCertificate() bool`

HasSigningCertificate returns a boolean if a field has been set.

### GetOpenIDConnectProvider

`func (o *OpenidConnectIdTokenValidatorShared) GetOpenIDConnectProvider() string`

GetOpenIDConnectProvider returns the OpenIDConnectProvider field if non-nil, zero value otherwise.

### GetOpenIDConnectProviderOk

`func (o *OpenidConnectIdTokenValidatorShared) GetOpenIDConnectProviderOk() (*string, bool)`

GetOpenIDConnectProviderOk returns a tuple with the OpenIDConnectProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectProvider

`func (o *OpenidConnectIdTokenValidatorShared) SetOpenIDConnectProvider(v string)`

SetOpenIDConnectProvider sets OpenIDConnectProvider field to given value.

### HasOpenIDConnectProvider

`func (o *OpenidConnectIdTokenValidatorShared) HasOpenIDConnectProvider() bool`

HasOpenIDConnectProvider returns a boolean if a field has been set.

### GetJwksEndpointPath

`func (o *OpenidConnectIdTokenValidatorShared) GetJwksEndpointPath() string`

GetJwksEndpointPath returns the JwksEndpointPath field if non-nil, zero value otherwise.

### GetJwksEndpointPathOk

`func (o *OpenidConnectIdTokenValidatorShared) GetJwksEndpointPathOk() (*string, bool)`

GetJwksEndpointPathOk returns a tuple with the JwksEndpointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksEndpointPath

`func (o *OpenidConnectIdTokenValidatorShared) SetJwksEndpointPath(v string)`

SetJwksEndpointPath sets JwksEndpointPath field to given value.

### HasJwksEndpointPath

`func (o *OpenidConnectIdTokenValidatorShared) HasJwksEndpointPath() bool`

HasJwksEndpointPath returns a boolean if a field has been set.

### GetDescription

`func (o *OpenidConnectIdTokenValidatorShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OpenidConnectIdTokenValidatorShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OpenidConnectIdTokenValidatorShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OpenidConnectIdTokenValidatorShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OpenidConnectIdTokenValidatorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OpenidConnectIdTokenValidatorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OpenidConnectIdTokenValidatorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdentityMapper

`func (o *OpenidConnectIdTokenValidatorShared) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *OpenidConnectIdTokenValidatorShared) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *OpenidConnectIdTokenValidatorShared) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetSubjectClaimName

`func (o *OpenidConnectIdTokenValidatorShared) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *OpenidConnectIdTokenValidatorShared) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *OpenidConnectIdTokenValidatorShared) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *OpenidConnectIdTokenValidatorShared) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetIssuerURL

`func (o *OpenidConnectIdTokenValidatorShared) GetIssuerURL() string`

GetIssuerURL returns the IssuerURL field if non-nil, zero value otherwise.

### GetIssuerURLOk

`func (o *OpenidConnectIdTokenValidatorShared) GetIssuerURLOk() (*string, bool)`

GetIssuerURLOk returns a tuple with the IssuerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerURL

`func (o *OpenidConnectIdTokenValidatorShared) SetIssuerURL(v string)`

SetIssuerURL sets IssuerURL field to given value.


### GetClockSkewGracePeriod

`func (o *OpenidConnectIdTokenValidatorShared) GetClockSkewGracePeriod() string`

GetClockSkewGracePeriod returns the ClockSkewGracePeriod field if non-nil, zero value otherwise.

### GetClockSkewGracePeriodOk

`func (o *OpenidConnectIdTokenValidatorShared) GetClockSkewGracePeriodOk() (*string, bool)`

GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewGracePeriod

`func (o *OpenidConnectIdTokenValidatorShared) SetClockSkewGracePeriod(v string)`

SetClockSkewGracePeriod sets ClockSkewGracePeriod field to given value.

### HasClockSkewGracePeriod

`func (o *OpenidConnectIdTokenValidatorShared) HasClockSkewGracePeriod() bool`

HasClockSkewGracePeriod returns a boolean if a field has been set.

### GetJwksCacheDuration

`func (o *OpenidConnectIdTokenValidatorShared) GetJwksCacheDuration() string`

GetJwksCacheDuration returns the JwksCacheDuration field if non-nil, zero value otherwise.

### GetJwksCacheDurationOk

`func (o *OpenidConnectIdTokenValidatorShared) GetJwksCacheDurationOk() (*string, bool)`

GetJwksCacheDurationOk returns a tuple with the JwksCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksCacheDuration

`func (o *OpenidConnectIdTokenValidatorShared) SetJwksCacheDuration(v string)`

SetJwksCacheDuration sets JwksCacheDuration field to given value.

### HasJwksCacheDuration

`func (o *OpenidConnectIdTokenValidatorShared) HasJwksCacheDuration() bool`

HasJwksCacheDuration returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *OpenidConnectIdTokenValidatorShared) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *OpenidConnectIdTokenValidatorShared) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *OpenidConnectIdTokenValidatorShared) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


