# AddOauthBearerSaslMechanismHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new SASL Mechanism Handler | 
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

### NewAddOauthBearerSaslMechanismHandlerRequest

`func NewAddOauthBearerSaslMechanismHandlerRequest(handlerName string, schemas []EnumoauthBearerSaslMechanismHandlerSchemaUrn, enabled bool, ) *AddOauthBearerSaslMechanismHandlerRequest`

NewAddOauthBearerSaslMechanismHandlerRequest instantiates a new AddOauthBearerSaslMechanismHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOauthBearerSaslMechanismHandlerRequestWithDefaults

`func NewAddOauthBearerSaslMechanismHandlerRequestWithDefaults() *AddOauthBearerSaslMechanismHandlerRequest`

NewAddOauthBearerSaslMechanismHandlerRequestWithDefaults instantiates a new AddOauthBearerSaslMechanismHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetSchemas() []EnumoauthBearerSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetSchemasOk() (*[]EnumoauthBearerSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetSchemas(v []EnumoauthBearerSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetAccessTokenValidator

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireBothAccessTokenAndIDToken

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetRequireBothAccessTokenAndIDToken() bool`

GetRequireBothAccessTokenAndIDToken returns the RequireBothAccessTokenAndIDToken field if non-nil, zero value otherwise.

### GetRequireBothAccessTokenAndIDTokenOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetRequireBothAccessTokenAndIDTokenOk() (*bool, bool)`

GetRequireBothAccessTokenAndIDTokenOk returns a tuple with the RequireBothAccessTokenAndIDToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBothAccessTokenAndIDToken

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetRequireBothAccessTokenAndIDToken(v bool)`

SetRequireBothAccessTokenAndIDToken sets RequireBothAccessTokenAndIDToken field to given value.

### HasRequireBothAccessTokenAndIDToken

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasRequireBothAccessTokenAndIDToken() bool`

HasRequireBothAccessTokenAndIDToken returns a boolean if a field has been set.

### GetValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetValidateAccessTokenWhenIDTokenIsAlsoProvided() EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp`

GetValidateAccessTokenWhenIDTokenIsAlsoProvided returns the ValidateAccessTokenWhenIDTokenIsAlsoProvided field if non-nil, zero value otherwise.

### GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk() (*EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp, bool)`

GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk returns a tuple with the ValidateAccessTokenWhenIDTokenIsAlsoProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetValidateAccessTokenWhenIDTokenIsAlsoProvided(v EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp)`

SetValidateAccessTokenWhenIDTokenIsAlsoProvided sets ValidateAccessTokenWhenIDTokenIsAlsoProvided field to given value.

### HasValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasValidateAccessTokenWhenIDTokenIsAlsoProvided() bool`

HasValidateAccessTokenWhenIDTokenIsAlsoProvided returns a boolean if a field has been set.

### GetAlternateAuthorizationIdentityMapper

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAlternateAuthorizationIdentityMapper() string`

GetAlternateAuthorizationIdentityMapper returns the AlternateAuthorizationIdentityMapper field if non-nil, zero value otherwise.

### GetAlternateAuthorizationIdentityMapperOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAlternateAuthorizationIdentityMapperOk() (*string, bool)`

GetAlternateAuthorizationIdentityMapperOk returns a tuple with the AlternateAuthorizationIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateAuthorizationIdentityMapper

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetAlternateAuthorizationIdentityMapper(v string)`

SetAlternateAuthorizationIdentityMapper sets AlternateAuthorizationIdentityMapper field to given value.

### HasAlternateAuthorizationIdentityMapper

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasAlternateAuthorizationIdentityMapper() bool`

HasAlternateAuthorizationIdentityMapper returns a boolean if a field has been set.

### GetAllRequiredScope

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAllRequiredScope() []string`

GetAllRequiredScope returns the AllRequiredScope field if non-nil, zero value otherwise.

### GetAllRequiredScopeOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAllRequiredScopeOk() (*[]string, bool)`

GetAllRequiredScopeOk returns a tuple with the AllRequiredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequiredScope

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetAllRequiredScope(v []string)`

SetAllRequiredScope sets AllRequiredScope field to given value.

### HasAllRequiredScope

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasAllRequiredScope() bool`

HasAllRequiredScope returns a boolean if a field has been set.

### GetAnyRequiredScope

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAnyRequiredScope() []string`

GetAnyRequiredScope returns the AnyRequiredScope field if non-nil, zero value otherwise.

### GetAnyRequiredScopeOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetAnyRequiredScopeOk() (*[]string, bool)`

GetAnyRequiredScopeOk returns a tuple with the AnyRequiredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredScope

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetAnyRequiredScope(v []string)`

SetAnyRequiredScope sets AnyRequiredScope field to given value.

### HasAnyRequiredScope

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasAnyRequiredScope() bool`

HasAnyRequiredScope returns a boolean if a field has been set.

### GetServerFqdn

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetServerFqdn() string`

GetServerFqdn returns the ServerFqdn field if non-nil, zero value otherwise.

### GetServerFqdnOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetServerFqdnOk() (*string, bool)`

GetServerFqdnOk returns a tuple with the ServerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFqdn

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetServerFqdn(v string)`

SetServerFqdn sets ServerFqdn field to given value.

### HasServerFqdn

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasServerFqdn() bool`

HasServerFqdn returns a boolean if a field has been set.

### GetDescription

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOauthBearerSaslMechanismHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddOauthBearerSaslMechanismHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddOauthBearerSaslMechanismHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


