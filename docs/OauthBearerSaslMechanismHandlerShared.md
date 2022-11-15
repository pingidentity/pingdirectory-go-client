# OauthBearerSaslMechanismHandlerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumoauthBearerSaslMechanismHandlerSchemaUrn**](EnumoauthBearerSaslMechanismHandlerSchemaUrn.md) |  | 
**AccessTokenValidator** | Pointer to **[]string** |  | [optional] 
**IdTokenValidator** | Pointer to **[]string** |  | [optional] 
**RequireBothAccessTokenAndIDToken** | Pointer to **bool** | Indicates whether bind requests will be required to have both an OAuth access token (in the \&quot;auth\&quot; element of the bind request) and an OpenID Connect ID token (in the \&quot;pingidentityidtoken\&quot; element of the bind request). | [optional] 
**ValidateAccessTokenWhenIDTokenIsAlsoProvided** | Pointer to [**EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp**](EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp.md) |  | [optional] 
**AlternateAuthorizationIdentityMapper** | Pointer to **string** | The identity mapper that will be used to map an alternate authorization identity (provided in the GS2 header of the encoded OAUTHBEARER bind request credentials) to the corresponding local entry. | [optional] 
**AllRequiredScope** | Pointer to **[]string** |  | [optional] 
**AnyRequiredScope** | Pointer to **[]string** |  | [optional] 
**ServerFqdn** | Pointer to **string** | The fully-qualified name that clients are expected to use when communicating with the server. | [optional] 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 

## Methods

### NewOauthBearerSaslMechanismHandlerShared

`func NewOauthBearerSaslMechanismHandlerShared(schemas []EnumoauthBearerSaslMechanismHandlerSchemaUrn, enabled bool, ) *OauthBearerSaslMechanismHandlerShared`

NewOauthBearerSaslMechanismHandlerShared instantiates a new OauthBearerSaslMechanismHandlerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthBearerSaslMechanismHandlerSharedWithDefaults

`func NewOauthBearerSaslMechanismHandlerSharedWithDefaults() *OauthBearerSaslMechanismHandlerShared`

NewOauthBearerSaslMechanismHandlerSharedWithDefaults instantiates a new OauthBearerSaslMechanismHandlerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *OauthBearerSaslMechanismHandlerShared) GetSchemas() []EnumoauthBearerSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetSchemasOk() (*[]EnumoauthBearerSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *OauthBearerSaslMechanismHandlerShared) SetSchemas(v []EnumoauthBearerSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccessTokenValidator

`func (o *OauthBearerSaslMechanismHandlerShared) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *OauthBearerSaslMechanismHandlerShared) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *OauthBearerSaslMechanismHandlerShared) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *OauthBearerSaslMechanismHandlerShared) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *OauthBearerSaslMechanismHandlerShared) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *OauthBearerSaslMechanismHandlerShared) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireBothAccessTokenAndIDToken

`func (o *OauthBearerSaslMechanismHandlerShared) GetRequireBothAccessTokenAndIDToken() bool`

GetRequireBothAccessTokenAndIDToken returns the RequireBothAccessTokenAndIDToken field if non-nil, zero value otherwise.

### GetRequireBothAccessTokenAndIDTokenOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetRequireBothAccessTokenAndIDTokenOk() (*bool, bool)`

GetRequireBothAccessTokenAndIDTokenOk returns a tuple with the RequireBothAccessTokenAndIDToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBothAccessTokenAndIDToken

`func (o *OauthBearerSaslMechanismHandlerShared) SetRequireBothAccessTokenAndIDToken(v bool)`

SetRequireBothAccessTokenAndIDToken sets RequireBothAccessTokenAndIDToken field to given value.

### HasRequireBothAccessTokenAndIDToken

`func (o *OauthBearerSaslMechanismHandlerShared) HasRequireBothAccessTokenAndIDToken() bool`

HasRequireBothAccessTokenAndIDToken returns a boolean if a field has been set.

### GetValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *OauthBearerSaslMechanismHandlerShared) GetValidateAccessTokenWhenIDTokenIsAlsoProvided() EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp`

GetValidateAccessTokenWhenIDTokenIsAlsoProvided returns the ValidateAccessTokenWhenIDTokenIsAlsoProvided field if non-nil, zero value otherwise.

### GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk() (*EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp, bool)`

GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk returns a tuple with the ValidateAccessTokenWhenIDTokenIsAlsoProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *OauthBearerSaslMechanismHandlerShared) SetValidateAccessTokenWhenIDTokenIsAlsoProvided(v EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp)`

SetValidateAccessTokenWhenIDTokenIsAlsoProvided sets ValidateAccessTokenWhenIDTokenIsAlsoProvided field to given value.

### HasValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *OauthBearerSaslMechanismHandlerShared) HasValidateAccessTokenWhenIDTokenIsAlsoProvided() bool`

HasValidateAccessTokenWhenIDTokenIsAlsoProvided returns a boolean if a field has been set.

### GetAlternateAuthorizationIdentityMapper

`func (o *OauthBearerSaslMechanismHandlerShared) GetAlternateAuthorizationIdentityMapper() string`

GetAlternateAuthorizationIdentityMapper returns the AlternateAuthorizationIdentityMapper field if non-nil, zero value otherwise.

### GetAlternateAuthorizationIdentityMapperOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetAlternateAuthorizationIdentityMapperOk() (*string, bool)`

GetAlternateAuthorizationIdentityMapperOk returns a tuple with the AlternateAuthorizationIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateAuthorizationIdentityMapper

`func (o *OauthBearerSaslMechanismHandlerShared) SetAlternateAuthorizationIdentityMapper(v string)`

SetAlternateAuthorizationIdentityMapper sets AlternateAuthorizationIdentityMapper field to given value.

### HasAlternateAuthorizationIdentityMapper

`func (o *OauthBearerSaslMechanismHandlerShared) HasAlternateAuthorizationIdentityMapper() bool`

HasAlternateAuthorizationIdentityMapper returns a boolean if a field has been set.

### GetAllRequiredScope

`func (o *OauthBearerSaslMechanismHandlerShared) GetAllRequiredScope() []string`

GetAllRequiredScope returns the AllRequiredScope field if non-nil, zero value otherwise.

### GetAllRequiredScopeOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetAllRequiredScopeOk() (*[]string, bool)`

GetAllRequiredScopeOk returns a tuple with the AllRequiredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequiredScope

`func (o *OauthBearerSaslMechanismHandlerShared) SetAllRequiredScope(v []string)`

SetAllRequiredScope sets AllRequiredScope field to given value.

### HasAllRequiredScope

`func (o *OauthBearerSaslMechanismHandlerShared) HasAllRequiredScope() bool`

HasAllRequiredScope returns a boolean if a field has been set.

### GetAnyRequiredScope

`func (o *OauthBearerSaslMechanismHandlerShared) GetAnyRequiredScope() []string`

GetAnyRequiredScope returns the AnyRequiredScope field if non-nil, zero value otherwise.

### GetAnyRequiredScopeOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetAnyRequiredScopeOk() (*[]string, bool)`

GetAnyRequiredScopeOk returns a tuple with the AnyRequiredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredScope

`func (o *OauthBearerSaslMechanismHandlerShared) SetAnyRequiredScope(v []string)`

SetAnyRequiredScope sets AnyRequiredScope field to given value.

### HasAnyRequiredScope

`func (o *OauthBearerSaslMechanismHandlerShared) HasAnyRequiredScope() bool`

HasAnyRequiredScope returns a boolean if a field has been set.

### GetServerFqdn

`func (o *OauthBearerSaslMechanismHandlerShared) GetServerFqdn() string`

GetServerFqdn returns the ServerFqdn field if non-nil, zero value otherwise.

### GetServerFqdnOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetServerFqdnOk() (*string, bool)`

GetServerFqdnOk returns a tuple with the ServerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFqdn

`func (o *OauthBearerSaslMechanismHandlerShared) SetServerFqdn(v string)`

SetServerFqdn sets ServerFqdn field to given value.

### HasServerFqdn

`func (o *OauthBearerSaslMechanismHandlerShared) HasServerFqdn() bool`

HasServerFqdn returns a boolean if a field has been set.

### GetDescription

`func (o *OauthBearerSaslMechanismHandlerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OauthBearerSaslMechanismHandlerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OauthBearerSaslMechanismHandlerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *OauthBearerSaslMechanismHandlerShared) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OauthBearerSaslMechanismHandlerShared) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OauthBearerSaslMechanismHandlerShared) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


