# PingOneIdTokenValidatorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the ID Token Validator | 
**Schemas** | [**[]EnumpingOneIdTokenValidatorSchemaUrn**](EnumpingOneIdTokenValidatorSchemaUrn.md) |  | 
**IssuerURL** | **string** | Specifies a PingOne base issuer URL. | 
**OpenIDConnectProvider** | **string** | Specifies HTTPS connection settings for the PingOne OpenID Connect provider. | 
**OpenIDConnectMetadataCacheDuration** | Pointer to **string** | How often the PingOne ID Token Validator should refresh its stored cache of OpenID Connect-related metadata. | [optional] 
**Description** | Pointer to **string** | A description for this ID Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this ID Token Validator is enabled for use in the Directory Server. | 
**IdentityMapper** | **string** | Specifies the name of the Identity Mapper that should be used to correlate an ID token subject value to a user entry. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject; i.e., the authenticated user. | [optional] 
**ClockSkewGracePeriod** | Pointer to **string** | Specifies the amount of clock skew that is tolerated by the ID Token Validator when evaluating whether a token is within its valid time interval. The duration specified by this parameter will be subtracted from the token&#39;s not-before (nbf) time and added to the token&#39;s expiration (exp) time, if present, to allow for any time difference between the local server&#39;s clock and the token issuer&#39;s clock. | [optional] 
**JwksCacheDuration** | Pointer to **string** | How often the ID Token Validator should refresh its cache of JWKS token signing keys. | [optional] 
**EvaluationOrderIndex** | **int32** | When multiple ID Token Validators are defined for a single Directory Server, this property determines the order in which the ID Token Validators are consulted. Values of this property must be unique among all ID Token Validators defined within Directory Server but not necessarily contiguous. ID Token Validators with lower values will be evaluated first to determine if they are able to validate the ID token. | 

## Methods

### NewPingOneIdTokenValidatorResponse

`func NewPingOneIdTokenValidatorResponse(id string, schemas []EnumpingOneIdTokenValidatorSchemaUrn, issuerURL string, openIDConnectProvider string, enabled bool, identityMapper string, evaluationOrderIndex int32, ) *PingOneIdTokenValidatorResponse`

NewPingOneIdTokenValidatorResponse instantiates a new PingOneIdTokenValidatorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneIdTokenValidatorResponseWithDefaults

`func NewPingOneIdTokenValidatorResponseWithDefaults() *PingOneIdTokenValidatorResponse`

NewPingOneIdTokenValidatorResponseWithDefaults instantiates a new PingOneIdTokenValidatorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PingOneIdTokenValidatorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOneIdTokenValidatorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOneIdTokenValidatorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PingOneIdTokenValidatorResponse) GetSchemas() []EnumpingOneIdTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PingOneIdTokenValidatorResponse) GetSchemasOk() (*[]EnumpingOneIdTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PingOneIdTokenValidatorResponse) SetSchemas(v []EnumpingOneIdTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIssuerURL

`func (o *PingOneIdTokenValidatorResponse) GetIssuerURL() string`

GetIssuerURL returns the IssuerURL field if non-nil, zero value otherwise.

### GetIssuerURLOk

`func (o *PingOneIdTokenValidatorResponse) GetIssuerURLOk() (*string, bool)`

GetIssuerURLOk returns a tuple with the IssuerURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerURL

`func (o *PingOneIdTokenValidatorResponse) SetIssuerURL(v string)`

SetIssuerURL sets IssuerURL field to given value.


### GetOpenIDConnectProvider

`func (o *PingOneIdTokenValidatorResponse) GetOpenIDConnectProvider() string`

GetOpenIDConnectProvider returns the OpenIDConnectProvider field if non-nil, zero value otherwise.

### GetOpenIDConnectProviderOk

`func (o *PingOneIdTokenValidatorResponse) GetOpenIDConnectProviderOk() (*string, bool)`

GetOpenIDConnectProviderOk returns a tuple with the OpenIDConnectProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectProvider

`func (o *PingOneIdTokenValidatorResponse) SetOpenIDConnectProvider(v string)`

SetOpenIDConnectProvider sets OpenIDConnectProvider field to given value.


### GetOpenIDConnectMetadataCacheDuration

`func (o *PingOneIdTokenValidatorResponse) GetOpenIDConnectMetadataCacheDuration() string`

GetOpenIDConnectMetadataCacheDuration returns the OpenIDConnectMetadataCacheDuration field if non-nil, zero value otherwise.

### GetOpenIDConnectMetadataCacheDurationOk

`func (o *PingOneIdTokenValidatorResponse) GetOpenIDConnectMetadataCacheDurationOk() (*string, bool)`

GetOpenIDConnectMetadataCacheDurationOk returns a tuple with the OpenIDConnectMetadataCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenIDConnectMetadataCacheDuration

`func (o *PingOneIdTokenValidatorResponse) SetOpenIDConnectMetadataCacheDuration(v string)`

SetOpenIDConnectMetadataCacheDuration sets OpenIDConnectMetadataCacheDuration field to given value.

### HasOpenIDConnectMetadataCacheDuration

`func (o *PingOneIdTokenValidatorResponse) HasOpenIDConnectMetadataCacheDuration() bool`

HasOpenIDConnectMetadataCacheDuration returns a boolean if a field has been set.

### GetDescription

`func (o *PingOneIdTokenValidatorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingOneIdTokenValidatorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingOneIdTokenValidatorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingOneIdTokenValidatorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PingOneIdTokenValidatorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PingOneIdTokenValidatorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PingOneIdTokenValidatorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIdentityMapper

`func (o *PingOneIdTokenValidatorResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *PingOneIdTokenValidatorResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *PingOneIdTokenValidatorResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetSubjectClaimName

`func (o *PingOneIdTokenValidatorResponse) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *PingOneIdTokenValidatorResponse) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *PingOneIdTokenValidatorResponse) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *PingOneIdTokenValidatorResponse) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetClockSkewGracePeriod

`func (o *PingOneIdTokenValidatorResponse) GetClockSkewGracePeriod() string`

GetClockSkewGracePeriod returns the ClockSkewGracePeriod field if non-nil, zero value otherwise.

### GetClockSkewGracePeriodOk

`func (o *PingOneIdTokenValidatorResponse) GetClockSkewGracePeriodOk() (*string, bool)`

GetClockSkewGracePeriodOk returns a tuple with the ClockSkewGracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClockSkewGracePeriod

`func (o *PingOneIdTokenValidatorResponse) SetClockSkewGracePeriod(v string)`

SetClockSkewGracePeriod sets ClockSkewGracePeriod field to given value.

### HasClockSkewGracePeriod

`func (o *PingOneIdTokenValidatorResponse) HasClockSkewGracePeriod() bool`

HasClockSkewGracePeriod returns a boolean if a field has been set.

### GetJwksCacheDuration

`func (o *PingOneIdTokenValidatorResponse) GetJwksCacheDuration() string`

GetJwksCacheDuration returns the JwksCacheDuration field if non-nil, zero value otherwise.

### GetJwksCacheDurationOk

`func (o *PingOneIdTokenValidatorResponse) GetJwksCacheDurationOk() (*string, bool)`

GetJwksCacheDurationOk returns a tuple with the JwksCacheDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksCacheDuration

`func (o *PingOneIdTokenValidatorResponse) SetJwksCacheDuration(v string)`

SetJwksCacheDuration sets JwksCacheDuration field to given value.

### HasJwksCacheDuration

`func (o *PingOneIdTokenValidatorResponse) HasJwksCacheDuration() bool`

HasJwksCacheDuration returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *PingOneIdTokenValidatorResponse) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *PingOneIdTokenValidatorResponse) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *PingOneIdTokenValidatorResponse) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


