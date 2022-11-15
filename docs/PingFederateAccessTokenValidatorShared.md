# PingFederateAccessTokenValidatorShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumpingFederateAccessTokenValidatorSchemaUrn**](EnumpingFederateAccessTokenValidatorSchemaUrn.md) |  | 
**ClientID** | **string** | The client identifier to use when authenticating to the PingFederate authorization server. | 
**ClientSecret** | Pointer to **string** | The client secret to use when authenticating to the PingFederate authorization server. | [optional] 
**ClientSecretPassphraseProvider** | Pointer to **string** | The passphrase provider for obtaining the client secret to use when authenticating to the PingFederate authorization server. | [optional] 
**IncludeAudParameter** | Pointer to **bool** | Whether to include the incoming request URL as the \&quot;aud\&quot; parameter when calling the PingFederate introspection endpoint. This property is ignored if the access-token-manager-id property is set. | [optional] 
**AccessTokenManagerID** | Pointer to **string** | The Access Token Manager instance ID to specify when calling the PingFederate introspection endpoint. If this property is set the include-aud-parameter property is ignored. | [optional] 
**EndpointCacheRefresh** | Pointer to **string** | How often the Access Token Validator should refresh its stored value of the PingFederate server&#39;s token introspection endpoint. | [optional] 
**EvaluationOrderIndex** | **int32** | When multiple Ping Federate Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Ping Federate Access Token Validators defined within Directory Server but not necessarily contiguous. Ping Federate Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token. | 
**AuthorizationServer** | Pointer to **string** | Specifies the external server that will be used to aid in validating access tokens. In most cases this will be the Authorization Server that minted the token. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | [optional] 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server. | [optional] 
**Description** | Pointer to **string** | A description for this Access Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this Access Token Validator is enabled for use in Directory Server. | 

## Methods

### NewPingFederateAccessTokenValidatorShared

`func NewPingFederateAccessTokenValidatorShared(schemas []EnumpingFederateAccessTokenValidatorSchemaUrn, clientID string, evaluationOrderIndex int32, enabled bool, ) *PingFederateAccessTokenValidatorShared`

NewPingFederateAccessTokenValidatorShared instantiates a new PingFederateAccessTokenValidatorShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingFederateAccessTokenValidatorSharedWithDefaults

`func NewPingFederateAccessTokenValidatorSharedWithDefaults() *PingFederateAccessTokenValidatorShared`

NewPingFederateAccessTokenValidatorSharedWithDefaults instantiates a new PingFederateAccessTokenValidatorShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PingFederateAccessTokenValidatorShared) GetSchemas() []EnumpingFederateAccessTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PingFederateAccessTokenValidatorShared) GetSchemasOk() (*[]EnumpingFederateAccessTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PingFederateAccessTokenValidatorShared) SetSchemas(v []EnumpingFederateAccessTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetClientID

`func (o *PingFederateAccessTokenValidatorShared) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *PingFederateAccessTokenValidatorShared) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *PingFederateAccessTokenValidatorShared) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *PingFederateAccessTokenValidatorShared) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PingFederateAccessTokenValidatorShared) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PingFederateAccessTokenValidatorShared) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PingFederateAccessTokenValidatorShared) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecretPassphraseProvider

`func (o *PingFederateAccessTokenValidatorShared) GetClientSecretPassphraseProvider() string`

GetClientSecretPassphraseProvider returns the ClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetClientSecretPassphraseProviderOk

`func (o *PingFederateAccessTokenValidatorShared) GetClientSecretPassphraseProviderOk() (*string, bool)`

GetClientSecretPassphraseProviderOk returns a tuple with the ClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretPassphraseProvider

`func (o *PingFederateAccessTokenValidatorShared) SetClientSecretPassphraseProvider(v string)`

SetClientSecretPassphraseProvider sets ClientSecretPassphraseProvider field to given value.

### HasClientSecretPassphraseProvider

`func (o *PingFederateAccessTokenValidatorShared) HasClientSecretPassphraseProvider() bool`

HasClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetIncludeAudParameter

`func (o *PingFederateAccessTokenValidatorShared) GetIncludeAudParameter() bool`

GetIncludeAudParameter returns the IncludeAudParameter field if non-nil, zero value otherwise.

### GetIncludeAudParameterOk

`func (o *PingFederateAccessTokenValidatorShared) GetIncludeAudParameterOk() (*bool, bool)`

GetIncludeAudParameterOk returns a tuple with the IncludeAudParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAudParameter

`func (o *PingFederateAccessTokenValidatorShared) SetIncludeAudParameter(v bool)`

SetIncludeAudParameter sets IncludeAudParameter field to given value.

### HasIncludeAudParameter

`func (o *PingFederateAccessTokenValidatorShared) HasIncludeAudParameter() bool`

HasIncludeAudParameter returns a boolean if a field has been set.

### GetAccessTokenManagerID

`func (o *PingFederateAccessTokenValidatorShared) GetAccessTokenManagerID() string`

GetAccessTokenManagerID returns the AccessTokenManagerID field if non-nil, zero value otherwise.

### GetAccessTokenManagerIDOk

`func (o *PingFederateAccessTokenValidatorShared) GetAccessTokenManagerIDOk() (*string, bool)`

GetAccessTokenManagerIDOk returns a tuple with the AccessTokenManagerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerID

`func (o *PingFederateAccessTokenValidatorShared) SetAccessTokenManagerID(v string)`

SetAccessTokenManagerID sets AccessTokenManagerID field to given value.

### HasAccessTokenManagerID

`func (o *PingFederateAccessTokenValidatorShared) HasAccessTokenManagerID() bool`

HasAccessTokenManagerID returns a boolean if a field has been set.

### GetEndpointCacheRefresh

`func (o *PingFederateAccessTokenValidatorShared) GetEndpointCacheRefresh() string`

GetEndpointCacheRefresh returns the EndpointCacheRefresh field if non-nil, zero value otherwise.

### GetEndpointCacheRefreshOk

`func (o *PingFederateAccessTokenValidatorShared) GetEndpointCacheRefreshOk() (*string, bool)`

GetEndpointCacheRefreshOk returns a tuple with the EndpointCacheRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCacheRefresh

`func (o *PingFederateAccessTokenValidatorShared) SetEndpointCacheRefresh(v string)`

SetEndpointCacheRefresh sets EndpointCacheRefresh field to given value.

### HasEndpointCacheRefresh

`func (o *PingFederateAccessTokenValidatorShared) HasEndpointCacheRefresh() bool`

HasEndpointCacheRefresh returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *PingFederateAccessTokenValidatorShared) GetEvaluationOrderIndex() int32`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *PingFederateAccessTokenValidatorShared) GetEvaluationOrderIndexOk() (*int32, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *PingFederateAccessTokenValidatorShared) SetEvaluationOrderIndex(v int32)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetAuthorizationServer

`func (o *PingFederateAccessTokenValidatorShared) GetAuthorizationServer() string`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *PingFederateAccessTokenValidatorShared) GetAuthorizationServerOk() (*string, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *PingFederateAccessTokenValidatorShared) SetAuthorizationServer(v string)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *PingFederateAccessTokenValidatorShared) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *PingFederateAccessTokenValidatorShared) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *PingFederateAccessTokenValidatorShared) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *PingFederateAccessTokenValidatorShared) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *PingFederateAccessTokenValidatorShared) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetSubjectClaimName

`func (o *PingFederateAccessTokenValidatorShared) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *PingFederateAccessTokenValidatorShared) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *PingFederateAccessTokenValidatorShared) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *PingFederateAccessTokenValidatorShared) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetDescription

`func (o *PingFederateAccessTokenValidatorShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingFederateAccessTokenValidatorShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingFederateAccessTokenValidatorShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingFederateAccessTokenValidatorShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PingFederateAccessTokenValidatorShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PingFederateAccessTokenValidatorShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PingFederateAccessTokenValidatorShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


