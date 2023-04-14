# PingFederateAccessTokenValidatorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Access Token Validator | 
**Schemas** | [**[]EnumpingFederateAccessTokenValidatorSchemaUrn**](EnumpingFederateAccessTokenValidatorSchemaUrn.md) |  | 
**ClientID** | **string** | The client identifier to use when authenticating to the PingFederate authorization server. | 
**ClientSecret** | Pointer to **string** | The client secret to use when authenticating to the PingFederate authorization server. | [optional] 
**ClientSecretPassphraseProvider** | Pointer to **string** | The passphrase provider for obtaining the client secret to use when authenticating to the PingFederate authorization server. | [optional] 
**IncludeAudParameter** | Pointer to **bool** | Whether to include the incoming request URL as the \&quot;aud\&quot; parameter when calling the PingFederate introspection endpoint. This property is ignored if the access-token-manager-id property is set. | [optional] 
**AccessTokenManagerID** | Pointer to **string** | The Access Token Manager instance ID to specify when calling the PingFederate introspection endpoint. If this property is set the include-aud-parameter property is ignored. | [optional] 
**EndpointCacheRefresh** | Pointer to **string** | How often the Access Token Validator should refresh its stored value of the PingFederate server&#39;s token introspection endpoint. | [optional] 
**EvaluationOrderIndex** | **int64** | When multiple Ping Federate Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Ping Federate Access Token Validators defined within Directory Server but not necessarily contiguous. Ping Federate Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token. | 
**AuthorizationServer** | Pointer to **string** | Specifies the external server that will be used to aid in validating access tokens. In most cases this will be the Authorization Server that minted the token. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | [optional] 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server. | [optional] 
**Description** | Pointer to **string** | A description for this Access Token Validator | [optional] 
**Enabled** | **bool** | Indicates whether this Access Token Validator is enabled for use in Directory Server. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewPingFederateAccessTokenValidatorResponse

`func NewPingFederateAccessTokenValidatorResponse(id string, schemas []EnumpingFederateAccessTokenValidatorSchemaUrn, clientID string, evaluationOrderIndex int64, enabled bool, ) *PingFederateAccessTokenValidatorResponse`

NewPingFederateAccessTokenValidatorResponse instantiates a new PingFederateAccessTokenValidatorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingFederateAccessTokenValidatorResponseWithDefaults

`func NewPingFederateAccessTokenValidatorResponseWithDefaults() *PingFederateAccessTokenValidatorResponse`

NewPingFederateAccessTokenValidatorResponseWithDefaults instantiates a new PingFederateAccessTokenValidatorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PingFederateAccessTokenValidatorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingFederateAccessTokenValidatorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingFederateAccessTokenValidatorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PingFederateAccessTokenValidatorResponse) GetSchemas() []EnumpingFederateAccessTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PingFederateAccessTokenValidatorResponse) GetSchemasOk() (*[]EnumpingFederateAccessTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PingFederateAccessTokenValidatorResponse) SetSchemas(v []EnumpingFederateAccessTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetClientID

`func (o *PingFederateAccessTokenValidatorResponse) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *PingFederateAccessTokenValidatorResponse) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *PingFederateAccessTokenValidatorResponse) SetClientID(v string)`

SetClientID sets ClientID field to given value.


### GetClientSecret

`func (o *PingFederateAccessTokenValidatorResponse) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PingFederateAccessTokenValidatorResponse) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PingFederateAccessTokenValidatorResponse) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PingFederateAccessTokenValidatorResponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientSecretPassphraseProvider

`func (o *PingFederateAccessTokenValidatorResponse) GetClientSecretPassphraseProvider() string`

GetClientSecretPassphraseProvider returns the ClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetClientSecretPassphraseProviderOk

`func (o *PingFederateAccessTokenValidatorResponse) GetClientSecretPassphraseProviderOk() (*string, bool)`

GetClientSecretPassphraseProviderOk returns a tuple with the ClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretPassphraseProvider

`func (o *PingFederateAccessTokenValidatorResponse) SetClientSecretPassphraseProvider(v string)`

SetClientSecretPassphraseProvider sets ClientSecretPassphraseProvider field to given value.

### HasClientSecretPassphraseProvider

`func (o *PingFederateAccessTokenValidatorResponse) HasClientSecretPassphraseProvider() bool`

HasClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetIncludeAudParameter

`func (o *PingFederateAccessTokenValidatorResponse) GetIncludeAudParameter() bool`

GetIncludeAudParameter returns the IncludeAudParameter field if non-nil, zero value otherwise.

### GetIncludeAudParameterOk

`func (o *PingFederateAccessTokenValidatorResponse) GetIncludeAudParameterOk() (*bool, bool)`

GetIncludeAudParameterOk returns a tuple with the IncludeAudParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAudParameter

`func (o *PingFederateAccessTokenValidatorResponse) SetIncludeAudParameter(v bool)`

SetIncludeAudParameter sets IncludeAudParameter field to given value.

### HasIncludeAudParameter

`func (o *PingFederateAccessTokenValidatorResponse) HasIncludeAudParameter() bool`

HasIncludeAudParameter returns a boolean if a field has been set.

### GetAccessTokenManagerID

`func (o *PingFederateAccessTokenValidatorResponse) GetAccessTokenManagerID() string`

GetAccessTokenManagerID returns the AccessTokenManagerID field if non-nil, zero value otherwise.

### GetAccessTokenManagerIDOk

`func (o *PingFederateAccessTokenValidatorResponse) GetAccessTokenManagerIDOk() (*string, bool)`

GetAccessTokenManagerIDOk returns a tuple with the AccessTokenManagerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenManagerID

`func (o *PingFederateAccessTokenValidatorResponse) SetAccessTokenManagerID(v string)`

SetAccessTokenManagerID sets AccessTokenManagerID field to given value.

### HasAccessTokenManagerID

`func (o *PingFederateAccessTokenValidatorResponse) HasAccessTokenManagerID() bool`

HasAccessTokenManagerID returns a boolean if a field has been set.

### GetEndpointCacheRefresh

`func (o *PingFederateAccessTokenValidatorResponse) GetEndpointCacheRefresh() string`

GetEndpointCacheRefresh returns the EndpointCacheRefresh field if non-nil, zero value otherwise.

### GetEndpointCacheRefreshOk

`func (o *PingFederateAccessTokenValidatorResponse) GetEndpointCacheRefreshOk() (*string, bool)`

GetEndpointCacheRefreshOk returns a tuple with the EndpointCacheRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCacheRefresh

`func (o *PingFederateAccessTokenValidatorResponse) SetEndpointCacheRefresh(v string)`

SetEndpointCacheRefresh sets EndpointCacheRefresh field to given value.

### HasEndpointCacheRefresh

`func (o *PingFederateAccessTokenValidatorResponse) HasEndpointCacheRefresh() bool`

HasEndpointCacheRefresh returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *PingFederateAccessTokenValidatorResponse) GetEvaluationOrderIndex() int64`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *PingFederateAccessTokenValidatorResponse) GetEvaluationOrderIndexOk() (*int64, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *PingFederateAccessTokenValidatorResponse) SetEvaluationOrderIndex(v int64)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.


### GetAuthorizationServer

`func (o *PingFederateAccessTokenValidatorResponse) GetAuthorizationServer() string`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *PingFederateAccessTokenValidatorResponse) GetAuthorizationServerOk() (*string, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *PingFederateAccessTokenValidatorResponse) SetAuthorizationServer(v string)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *PingFederateAccessTokenValidatorResponse) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *PingFederateAccessTokenValidatorResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *PingFederateAccessTokenValidatorResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *PingFederateAccessTokenValidatorResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *PingFederateAccessTokenValidatorResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetSubjectClaimName

`func (o *PingFederateAccessTokenValidatorResponse) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *PingFederateAccessTokenValidatorResponse) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *PingFederateAccessTokenValidatorResponse) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *PingFederateAccessTokenValidatorResponse) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetDescription

`func (o *PingFederateAccessTokenValidatorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PingFederateAccessTokenValidatorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PingFederateAccessTokenValidatorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PingFederateAccessTokenValidatorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PingFederateAccessTokenValidatorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PingFederateAccessTokenValidatorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PingFederateAccessTokenValidatorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *PingFederateAccessTokenValidatorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PingFederateAccessTokenValidatorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PingFederateAccessTokenValidatorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PingFederateAccessTokenValidatorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PingFederateAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PingFederateAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PingFederateAccessTokenValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PingFederateAccessTokenValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


