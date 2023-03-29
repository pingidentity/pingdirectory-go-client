# OauthBearerSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the SASL Mechanism Handler | 
**Schemas** | [**[]EnumoauthBearerSaslMechanismHandlerSchemaUrn**](EnumoauthBearerSaslMechanismHandlerSchemaUrn.md) |  | 
**AccessTokenValidator** | Pointer to **[]string** | An access token validator that will ensure that each presented OAuth access token is authentic and trustworthy. It must be configured with an identity mapper that will be used to map the access token to a local entry. | [optional] 
**IdTokenValidator** | Pointer to **[]string** | An ID token validator that will ensure that each presented OpenID Connect ID token is authentic and trustworthy, and that will map the token to a local entry. | [optional] 
**RequireBothAccessTokenAndIDToken** | Pointer to **bool** | Indicates whether bind requests will be required to have both an OAuth access token (in the \&quot;auth\&quot; element of the bind request) and an OpenID Connect ID token (in the \&quot;pingidentityidtoken\&quot; element of the bind request). | [optional] 
**ValidateAccessTokenWhenIDTokenIsAlsoProvided** | Pointer to [**EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp**](EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp.md) |  | [optional] 
**AlternateAuthorizationIdentityMapper** | Pointer to **string** | The identity mapper that will be used to map an alternate authorization identity (provided in the GS2 header of the encoded OAUTHBEARER bind request credentials) to the corresponding local entry. | [optional] 
**AllRequiredScope** | Pointer to **[]string** | The set of OAuth scopes that will all be required for any access tokens that will be allowed for authentication. | [optional] 
**AnyRequiredScope** | Pointer to **[]string** | The set of OAuth scopes that a token may have to be allowed for authentication. | [optional] 
**ServerFqdn** | Pointer to **string** | The fully-qualified name that clients are expected to use when communicating with the server. | [optional] 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewOauthBearerSaslMechanismHandlerResponse

`func NewOauthBearerSaslMechanismHandlerResponse(id string, schemas []EnumoauthBearerSaslMechanismHandlerSchemaUrn, enabled bool, ) *OauthBearerSaslMechanismHandlerResponse`

NewOauthBearerSaslMechanismHandlerResponse instantiates a new OauthBearerSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthBearerSaslMechanismHandlerResponseWithDefaults

`func NewOauthBearerSaslMechanismHandlerResponseWithDefaults() *OauthBearerSaslMechanismHandlerResponse`

NewOauthBearerSaslMechanismHandlerResponseWithDefaults instantiates a new OauthBearerSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OauthBearerSaslMechanismHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OauthBearerSaslMechanismHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *OauthBearerSaslMechanismHandlerResponse) GetSchemas() []EnumoauthBearerSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumoauthBearerSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *OauthBearerSaslMechanismHandlerResponse) SetSchemas(v []EnumoauthBearerSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccessTokenValidator

`func (o *OauthBearerSaslMechanismHandlerResponse) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *OauthBearerSaslMechanismHandlerResponse) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *OauthBearerSaslMechanismHandlerResponse) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *OauthBearerSaslMechanismHandlerResponse) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *OauthBearerSaslMechanismHandlerResponse) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *OauthBearerSaslMechanismHandlerResponse) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireBothAccessTokenAndIDToken

`func (o *OauthBearerSaslMechanismHandlerResponse) GetRequireBothAccessTokenAndIDToken() bool`

GetRequireBothAccessTokenAndIDToken returns the RequireBothAccessTokenAndIDToken field if non-nil, zero value otherwise.

### GetRequireBothAccessTokenAndIDTokenOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetRequireBothAccessTokenAndIDTokenOk() (*bool, bool)`

GetRequireBothAccessTokenAndIDTokenOk returns a tuple with the RequireBothAccessTokenAndIDToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBothAccessTokenAndIDToken

`func (o *OauthBearerSaslMechanismHandlerResponse) SetRequireBothAccessTokenAndIDToken(v bool)`

SetRequireBothAccessTokenAndIDToken sets RequireBothAccessTokenAndIDToken field to given value.

### HasRequireBothAccessTokenAndIDToken

`func (o *OauthBearerSaslMechanismHandlerResponse) HasRequireBothAccessTokenAndIDToken() bool`

HasRequireBothAccessTokenAndIDToken returns a boolean if a field has been set.

### GetValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *OauthBearerSaslMechanismHandlerResponse) GetValidateAccessTokenWhenIDTokenIsAlsoProvided() EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp`

GetValidateAccessTokenWhenIDTokenIsAlsoProvided returns the ValidateAccessTokenWhenIDTokenIsAlsoProvided field if non-nil, zero value otherwise.

### GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk() (*EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp, bool)`

GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk returns a tuple with the ValidateAccessTokenWhenIDTokenIsAlsoProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *OauthBearerSaslMechanismHandlerResponse) SetValidateAccessTokenWhenIDTokenIsAlsoProvided(v EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp)`

SetValidateAccessTokenWhenIDTokenIsAlsoProvided sets ValidateAccessTokenWhenIDTokenIsAlsoProvided field to given value.

### HasValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *OauthBearerSaslMechanismHandlerResponse) HasValidateAccessTokenWhenIDTokenIsAlsoProvided() bool`

HasValidateAccessTokenWhenIDTokenIsAlsoProvided returns a boolean if a field has been set.

### GetAlternateAuthorizationIdentityMapper

`func (o *OauthBearerSaslMechanismHandlerResponse) GetAlternateAuthorizationIdentityMapper() string`

GetAlternateAuthorizationIdentityMapper returns the AlternateAuthorizationIdentityMapper field if non-nil, zero value otherwise.

### GetAlternateAuthorizationIdentityMapperOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetAlternateAuthorizationIdentityMapperOk() (*string, bool)`

GetAlternateAuthorizationIdentityMapperOk returns a tuple with the AlternateAuthorizationIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateAuthorizationIdentityMapper

`func (o *OauthBearerSaslMechanismHandlerResponse) SetAlternateAuthorizationIdentityMapper(v string)`

SetAlternateAuthorizationIdentityMapper sets AlternateAuthorizationIdentityMapper field to given value.

### HasAlternateAuthorizationIdentityMapper

`func (o *OauthBearerSaslMechanismHandlerResponse) HasAlternateAuthorizationIdentityMapper() bool`

HasAlternateAuthorizationIdentityMapper returns a boolean if a field has been set.

### GetAllRequiredScope

`func (o *OauthBearerSaslMechanismHandlerResponse) GetAllRequiredScope() []string`

GetAllRequiredScope returns the AllRequiredScope field if non-nil, zero value otherwise.

### GetAllRequiredScopeOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetAllRequiredScopeOk() (*[]string, bool)`

GetAllRequiredScopeOk returns a tuple with the AllRequiredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequiredScope

`func (o *OauthBearerSaslMechanismHandlerResponse) SetAllRequiredScope(v []string)`

SetAllRequiredScope sets AllRequiredScope field to given value.

### HasAllRequiredScope

`func (o *OauthBearerSaslMechanismHandlerResponse) HasAllRequiredScope() bool`

HasAllRequiredScope returns a boolean if a field has been set.

### GetAnyRequiredScope

`func (o *OauthBearerSaslMechanismHandlerResponse) GetAnyRequiredScope() []string`

GetAnyRequiredScope returns the AnyRequiredScope field if non-nil, zero value otherwise.

### GetAnyRequiredScopeOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetAnyRequiredScopeOk() (*[]string, bool)`

GetAnyRequiredScopeOk returns a tuple with the AnyRequiredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredScope

`func (o *OauthBearerSaslMechanismHandlerResponse) SetAnyRequiredScope(v []string)`

SetAnyRequiredScope sets AnyRequiredScope field to given value.

### HasAnyRequiredScope

`func (o *OauthBearerSaslMechanismHandlerResponse) HasAnyRequiredScope() bool`

HasAnyRequiredScope returns a boolean if a field has been set.

### GetServerFqdn

`func (o *OauthBearerSaslMechanismHandlerResponse) GetServerFqdn() string`

GetServerFqdn returns the ServerFqdn field if non-nil, zero value otherwise.

### GetServerFqdnOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetServerFqdnOk() (*string, bool)`

GetServerFqdnOk returns a tuple with the ServerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFqdn

`func (o *OauthBearerSaslMechanismHandlerResponse) SetServerFqdn(v string)`

SetServerFqdn sets ServerFqdn field to given value.

### HasServerFqdn

`func (o *OauthBearerSaslMechanismHandlerResponse) HasServerFqdn() bool`

HasServerFqdn returns a boolean if a field has been set.

### GetDescription

`func (o *OauthBearerSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OauthBearerSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OauthBearerSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OauthBearerSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OauthBearerSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *OauthBearerSaslMechanismHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *OauthBearerSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *OauthBearerSaslMechanismHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *OauthBearerSaslMechanismHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *OauthBearerSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *OauthBearerSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *OauthBearerSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *OauthBearerSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


