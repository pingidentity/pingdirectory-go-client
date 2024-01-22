# BindAccessTokenValidatorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Schemas** | [**[]EnumbindAccessTokenValidatorSchemaUrn**](EnumbindAccessTokenValidatorSchemaUrn.md) |  | 
**Id** | **string** | Name of the Access Token Validator | 
**Enabled** | **bool** | Indicates whether this Bind Access Token Validator is enabled for use in Directory Server. | 
**PersistAccessTokens** | Pointer to **bool** | Indicates whether access tokens should be persisted in user entries. | [optional] 
**MaximumTokenLifetime** | Pointer to **string** | Specifies the maximum length of time that a generated token should be considered valid. If this is not specified, then generated access tokens will not expire. | [optional] 
**AllowedAuthenticationType** | Pointer to [**[]EnumaccessTokenValidatorAllowedAuthenticationTypeProp**](EnumaccessTokenValidatorAllowedAuthenticationTypeProp.md) |  | [optional] 
**AllowedSASLMechanism** | Pointer to **[]string** | Specifies the names of the SASL mechanisms for which access tokens may be generated, and for which generated access tokens will be accepted. | [optional] 
**GenerateTokenResultCriteria** | Pointer to **string** | A reference to a request criteria object that may be used to identify the types of bind operations for which access tokens may be generated. If no criteria is specified, then access tokens may be generated for any bind operations that satisfy the other requirements configured in this validator. | [optional] 
**IncludedScope** | Pointer to **[]string** | Specifies the names of any scopes that should be granted to a client that authenticates with a bind access token. By default, no scopes will be granted. | [optional] 
**IdentityMapper** | Pointer to **string** | Specifies the name of the Identity Mapper that should be used for associating user entries with Bearer token subject names. The claim name from which to obtain the subject (i.e. the currently logged-in user) may be configured using the subject-claim-name property. | [optional] 
**SubjectClaimName** | Pointer to **string** | The name of the token claim that contains the subject, i.e. the logged-in user in an access token. This property goes hand-in-hand with the identity-mapper property and tells the Identity Mapper which field to use to look up the user entry on the server. | [optional] 
**Description** | Pointer to **string** | A description for this Access Token Validator | [optional] 
**EvaluationOrderIndex** | **int64** | When multiple Access Token Validators are defined for a single Directory Server, this property determines the evaluation order for determining the correct validator class for an access token received by the Directory Server. Values of this property must be unique among all Access Token Validators defined within Directory Server but not necessarily contiguous. Access Token Validators with a smaller value will be evaluated first to determine if they are able to validate the access token. | 

## Methods

### NewBindAccessTokenValidatorResponse

`func NewBindAccessTokenValidatorResponse(schemas []EnumbindAccessTokenValidatorSchemaUrn, id string, enabled bool, evaluationOrderIndex int64, ) *BindAccessTokenValidatorResponse`

NewBindAccessTokenValidatorResponse instantiates a new BindAccessTokenValidatorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindAccessTokenValidatorResponseWithDefaults

`func NewBindAccessTokenValidatorResponseWithDefaults() *BindAccessTokenValidatorResponse`

NewBindAccessTokenValidatorResponseWithDefaults instantiates a new BindAccessTokenValidatorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *BindAccessTokenValidatorResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BindAccessTokenValidatorResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BindAccessTokenValidatorResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BindAccessTokenValidatorResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *BindAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *BindAccessTokenValidatorResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *BindAccessTokenValidatorResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *BindAccessTokenValidatorResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *BindAccessTokenValidatorResponse) GetSchemas() []EnumbindAccessTokenValidatorSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *BindAccessTokenValidatorResponse) GetSchemasOk() (*[]EnumbindAccessTokenValidatorSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *BindAccessTokenValidatorResponse) SetSchemas(v []EnumbindAccessTokenValidatorSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *BindAccessTokenValidatorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BindAccessTokenValidatorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BindAccessTokenValidatorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEnabled

`func (o *BindAccessTokenValidatorResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BindAccessTokenValidatorResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BindAccessTokenValidatorResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPersistAccessTokens

`func (o *BindAccessTokenValidatorResponse) GetPersistAccessTokens() bool`

GetPersistAccessTokens returns the PersistAccessTokens field if non-nil, zero value otherwise.

### GetPersistAccessTokensOk

`func (o *BindAccessTokenValidatorResponse) GetPersistAccessTokensOk() (*bool, bool)`

GetPersistAccessTokensOk returns a tuple with the PersistAccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistAccessTokens

`func (o *BindAccessTokenValidatorResponse) SetPersistAccessTokens(v bool)`

SetPersistAccessTokens sets PersistAccessTokens field to given value.

### HasPersistAccessTokens

`func (o *BindAccessTokenValidatorResponse) HasPersistAccessTokens() bool`

HasPersistAccessTokens returns a boolean if a field has been set.

### GetMaximumTokenLifetime

`func (o *BindAccessTokenValidatorResponse) GetMaximumTokenLifetime() string`

GetMaximumTokenLifetime returns the MaximumTokenLifetime field if non-nil, zero value otherwise.

### GetMaximumTokenLifetimeOk

`func (o *BindAccessTokenValidatorResponse) GetMaximumTokenLifetimeOk() (*string, bool)`

GetMaximumTokenLifetimeOk returns a tuple with the MaximumTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumTokenLifetime

`func (o *BindAccessTokenValidatorResponse) SetMaximumTokenLifetime(v string)`

SetMaximumTokenLifetime sets MaximumTokenLifetime field to given value.

### HasMaximumTokenLifetime

`func (o *BindAccessTokenValidatorResponse) HasMaximumTokenLifetime() bool`

HasMaximumTokenLifetime returns a boolean if a field has been set.

### GetAllowedAuthenticationType

`func (o *BindAccessTokenValidatorResponse) GetAllowedAuthenticationType() []EnumaccessTokenValidatorAllowedAuthenticationTypeProp`

GetAllowedAuthenticationType returns the AllowedAuthenticationType field if non-nil, zero value otherwise.

### GetAllowedAuthenticationTypeOk

`func (o *BindAccessTokenValidatorResponse) GetAllowedAuthenticationTypeOk() (*[]EnumaccessTokenValidatorAllowedAuthenticationTypeProp, bool)`

GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationType

`func (o *BindAccessTokenValidatorResponse) SetAllowedAuthenticationType(v []EnumaccessTokenValidatorAllowedAuthenticationTypeProp)`

SetAllowedAuthenticationType sets AllowedAuthenticationType field to given value.

### HasAllowedAuthenticationType

`func (o *BindAccessTokenValidatorResponse) HasAllowedAuthenticationType() bool`

HasAllowedAuthenticationType returns a boolean if a field has been set.

### GetAllowedSASLMechanism

`func (o *BindAccessTokenValidatorResponse) GetAllowedSASLMechanism() []string`

GetAllowedSASLMechanism returns the AllowedSASLMechanism field if non-nil, zero value otherwise.

### GetAllowedSASLMechanismOk

`func (o *BindAccessTokenValidatorResponse) GetAllowedSASLMechanismOk() (*[]string, bool)`

GetAllowedSASLMechanismOk returns a tuple with the AllowedSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSASLMechanism

`func (o *BindAccessTokenValidatorResponse) SetAllowedSASLMechanism(v []string)`

SetAllowedSASLMechanism sets AllowedSASLMechanism field to given value.

### HasAllowedSASLMechanism

`func (o *BindAccessTokenValidatorResponse) HasAllowedSASLMechanism() bool`

HasAllowedSASLMechanism returns a boolean if a field has been set.

### GetGenerateTokenResultCriteria

`func (o *BindAccessTokenValidatorResponse) GetGenerateTokenResultCriteria() string`

GetGenerateTokenResultCriteria returns the GenerateTokenResultCriteria field if non-nil, zero value otherwise.

### GetGenerateTokenResultCriteriaOk

`func (o *BindAccessTokenValidatorResponse) GetGenerateTokenResultCriteriaOk() (*string, bool)`

GetGenerateTokenResultCriteriaOk returns a tuple with the GenerateTokenResultCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateTokenResultCriteria

`func (o *BindAccessTokenValidatorResponse) SetGenerateTokenResultCriteria(v string)`

SetGenerateTokenResultCriteria sets GenerateTokenResultCriteria field to given value.

### HasGenerateTokenResultCriteria

`func (o *BindAccessTokenValidatorResponse) HasGenerateTokenResultCriteria() bool`

HasGenerateTokenResultCriteria returns a boolean if a field has been set.

### GetIncludedScope

`func (o *BindAccessTokenValidatorResponse) GetIncludedScope() []string`

GetIncludedScope returns the IncludedScope field if non-nil, zero value otherwise.

### GetIncludedScopeOk

`func (o *BindAccessTokenValidatorResponse) GetIncludedScopeOk() (*[]string, bool)`

GetIncludedScopeOk returns a tuple with the IncludedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedScope

`func (o *BindAccessTokenValidatorResponse) SetIncludedScope(v []string)`

SetIncludedScope sets IncludedScope field to given value.

### HasIncludedScope

`func (o *BindAccessTokenValidatorResponse) HasIncludedScope() bool`

HasIncludedScope returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *BindAccessTokenValidatorResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *BindAccessTokenValidatorResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *BindAccessTokenValidatorResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.

### HasIdentityMapper

`func (o *BindAccessTokenValidatorResponse) HasIdentityMapper() bool`

HasIdentityMapper returns a boolean if a field has been set.

### GetSubjectClaimName

`func (o *BindAccessTokenValidatorResponse) GetSubjectClaimName() string`

GetSubjectClaimName returns the SubjectClaimName field if non-nil, zero value otherwise.

### GetSubjectClaimNameOk

`func (o *BindAccessTokenValidatorResponse) GetSubjectClaimNameOk() (*string, bool)`

GetSubjectClaimNameOk returns a tuple with the SubjectClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectClaimName

`func (o *BindAccessTokenValidatorResponse) SetSubjectClaimName(v string)`

SetSubjectClaimName sets SubjectClaimName field to given value.

### HasSubjectClaimName

`func (o *BindAccessTokenValidatorResponse) HasSubjectClaimName() bool`

HasSubjectClaimName returns a boolean if a field has been set.

### GetDescription

`func (o *BindAccessTokenValidatorResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BindAccessTokenValidatorResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BindAccessTokenValidatorResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BindAccessTokenValidatorResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluationOrderIndex

`func (o *BindAccessTokenValidatorResponse) GetEvaluationOrderIndex() int64`

GetEvaluationOrderIndex returns the EvaluationOrderIndex field if non-nil, zero value otherwise.

### GetEvaluationOrderIndexOk

`func (o *BindAccessTokenValidatorResponse) GetEvaluationOrderIndexOk() (*int64, bool)`

GetEvaluationOrderIndexOk returns a tuple with the EvaluationOrderIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationOrderIndex

`func (o *BindAccessTokenValidatorResponse) SetEvaluationOrderIndex(v int64)`

SetEvaluationOrderIndex sets EvaluationOrderIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


