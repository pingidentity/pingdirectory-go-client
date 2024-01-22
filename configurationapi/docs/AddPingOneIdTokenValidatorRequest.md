# AddPingOneIdTokenValidatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumpingOneIdTokenValidatorSchemaUrn**](EnumpingOneIdTokenValidatorSchemaUrn.md) |  | 
**IssuerURL** | **string** | Specifies a PingOne base issuer URL. | 
**OpenIDConnectProvider** | Pointer to **string** | Specifies HTTPS connection settings for the PingOne OpenID Connect provider. | [optional] 
**OpenIDConnectMetadataCacheDuration** | Pointer to **string** | How often the PingOne ID Token Validator should refresh its stored cache of OpenID Connect-related metadata. | [optional] 
**Description** | Pointer to **string** | A description for this ID Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this ID Token Validator is enabled for use in the Directory Server. | 
**IdentityMapper** | **string** | Specifies the name of the Identity Mapper that should be used to correlate an ID token subject value to a user entry. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject; i.e., the authenticated user. | [optional] 
**ClockSkewGracePeriod** | Pointer to **string** | Specifies the amount of clock skew that is tolerated by the ID Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token&#39;s not-before (nbf) time and added to the token&#39;s expiration (exp) time, if present, to allow for any time difference between the local server&#39;s clock and the token issuer&#39;s clock. | [optional] 
**JwksCacheDuration** | Pointer to **string** | How often the ID Token Validator should refresh its cache of JWKS token signing keys. | [optional] 
**EvaluationOrderIndex** | **int64** | When multiple ID Token Validators are defined for a single Directory Server, this property determines the order in which the ID Token Validators are consulted. Values of this property must be unique among all ID Token Validators defined within Directory Server but not necessarily contiguous. ID Token Validators with lower values will be evaluated first to determine if they are able to validate the ID token. | 
**ValidatorName** | **string** | Name of the new ID Token Validator | 

## Methods

### NewAddPingOneIdTokenValidatorRequest

`func NewAddPingOneIdTokenValidatorRequest(schemas []EnumpingOneIdTokenValidatorSchemaUrn, issuerURL string, enabled bool, identityMapper string, evaluationOrderIndex int64, validatorName string, ) *AddPingOneIdTokenValidatorRequest`

NewAddPingOneIdTokenValidatorRequest instantiates a new AddPingOneIdTokenValidatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPingOneIdTokenValidatorRequestWithDefaults

`func NewAddPingOneIdTokenValidatorRequestWithDefaults() *AddPingOneIdTokenValidatorRequest`

NewAddPingOneIdTokenValidatorRequestWithDefaults instantiates a new AddPingOneIdTokenValidatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddPingOneIdTokenValidatorRequest) GetSchemas() []EnumpingOneIdTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPingOneIdTokenValidatorRequest) GetSchemasOk() (*[]EnumpingOneIdTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPingOneIdTokenValidatorRequest) SetSchemas(v []EnumpingOneIdTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIssuerURL

`func (o *AddPingOneIdTokenValidatorRequest) GetIssuerURL() string`

GetIssuerURL returns the IssuerURL field if non-nil, zero value otherwise.

### GetIssuerURLOk

`func (o *AddPingOneIdTokenValidatorRequest) GetIssuerURLOk() (*string, bool)`

GetIssuerURLOk returns a tuple with the IssuerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerURL

`func (o *AddPingOneIdTokenValidatorRequest) SetIssuerURL(v string)`

SetIssuerURL sets IssuerURL field to given value.


### GetOpenIDConnectProvider

`func (o *AddPingOneIdTokenValidatorRequest) GetOpenIDConnectProvider() string`

GetOpenIDConnectProvider returns the OpenIDConnectProvider field if non-nil, zero value otherwise.

### GetOpenIDConnectProviderOk

`func (o *AddPingOneIdTokenValidatorRequest) GetOpenIDConnectProviderOk() (*string, bool)`

GetOpenIDConnectProviderOk returns a tuple with the OpenIDConnectProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectProvider

`func (o *AddPingOneIdTokenValidatorRequest) SetOpenIDConnectProvider(v string)`

SetOpenIDConnectProvider sets OpenIDConnectProvider field to given value.

### HasOpenIDConnectProvider

`func (o *AddPingOneIdTokenValidatorRequest) HasOpenIDConnectProvider() bool`

HasOpenIDConnectProvider returns a boolean if a field has been set.

### GetOpenIDConnectMetadataCacheDuration

`func (o *AddPingOneIdTokenValidatorRequest) GetOpenIDConnectMetadataCacheDuration() string`

GetOpenIDConnectMetadataCacheDuration returns the OpenIDConnectMetadataCacheDuration field if non-nil, zero value otherwise.

### GetOpenIDConnectMetadataCacheDurationOk

`func (o *AddPingOneIdTokenValidatorRequest) GetOpenIDConnectMetadataCacheDurationOk() (*string, bool)`

GetOpenIDConnectMetadataCacheDurationOk returns a tuple with the OpenIDConnectMetadataCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectMetadataCacheDuration

`func (o *AddPingOneIdTokenValidatorRequest) SetOpenIDConnectMetadataCacheDuration(v string)`

SetOpenIDConnectMetadataCacheDuration sets OpenIDConnectMetadataCacheDuration field to given value.

### HasOpenIDConnectMetadataCacheDuration

`func (o *AddPingOneIdTokenValidatorRequest) HasOpenIDConnectMetadataCacheDuration() bool`

HasOpenIDConnectMetadataCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *AddPingOneIdTokenValidatorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPingOneIdTokenValidatorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPingOneIdTokenValidatorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPingOneIdTokenValidatorRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPingOneIdTokenValidatorRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPingOneIdTokenValidatorRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPingOneIdTokenValidatorRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdentityMapper

`func (o *AddPingOneIdTokenValidatorRequest) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddPingOneIdTokenValidatorRequest) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddPingOneIdTokenValidatorRequest) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetSubjectClaimName

`func (o *AddPingOneIdTokenValidatorRequest) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *AddPingOneIdTokenValidatorRequest) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *AddPingOneIdTokenValidatorRequest) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *AddPingOneIdTokenValidatorRequest) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetClockSkewGracePeriod

`func (o *AddPingOneIdTokenValidatorRequest) GetClockSkewGracePeriod() string`

GetClockSkewGracePeriod returns the ClockSkewGracePeriod field if non-nil, zero value otherwise.

### GetClockSkewGracePeriodOk

`func (o *AddPingOneIdTokenValidatorRequest) GetClockSkewGracePeriodOk() (*string, bool)`

GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewGracePeriod

`func (o *AddPingOneIdTokenValidatorRequest) SetClockSkewGracePeriod(v string)`

SetClockSkewGracePeriod sets ClockSkewGracePeriod field to given value.

### HasClockSkewGracePeriod

`func (o *AddPingOneIdTokenValidatorRequest) HasClockSkewGracePeriod() bool`

HasClockSkewGracePeriod returns a boolean if a field has been set.

### GetJwksCacheDuration

`func (o *AddPingOneIdTokenValidatorRequest) GetJwksCacheDuration() string`

GetJwksCacheDuration returns the JwksCacheDuration field if non-nil, zero value otherwise.

### GetJwksCacheDurationOk

`func (o *AddPingOneIdTokenValidatorRequest) GetJwksCacheDurationOk() (*string, bool)`

GetJwksCacheDurationOk returns a tuple with the JwksCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksCacheDuration

`func (o *AddPingOneIdTokenValidatorRequest) SetJwksCacheDuration(v string)`

SetJwksCacheDuration sets JwksCacheDuration field to given value.

### HasJwksCacheDuration

`func (o *AddPingOneIdTokenValidatorRequest) HasJwksCacheDuration() bool`

HasJwksCacheDuration returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *AddPingOneIdTokenValidatorRequest) GetEvaluationOrderIndex() int64`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *AddPingOneIdTokenValidatorRequest) GetEvaluationOrderIndexOk() (*int64, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *AddPingOneIdTokenValidatorRequest) SetEvaluationOrderIndex(v int64)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetValidatorName

`func (o *AddPingOneIdTokenValidatorRequest) GetValidatorName() string`

GetValidatorName returns the ValidatorName field if non-nil, zero value otherwise.

### GetValidatorNameOk

`func (o *AddPingOneIdTokenValidatorRequest) GetValidatorNameOk() (*string, bool)`

GetValidatorNameOk returns a tuple with the ValidatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorName

`func (o *AddPingOneIdTokenValidatorRequest) SetValidatorName(v string)`

SetValidatorName sets ValidatorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


