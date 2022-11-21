# AddSaslMechanismHandler200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the SASL Mechanism Handler | 
**Schemas** | [**[]EnumthirdPartySaslMechanismHandlerSchemaUrn**](EnumthirdPartySaslMechanismHandlerSchemaUrn.md) |  | 
**IdentityMapper** | **string** | The identity mapper that may be used to map usernames to user entries. If the custom SASL mechanism involves a username or some other form of authentication and/or authorization identity, then this may be used to map that ID to an entry for that user. | 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**OtpValidityDuration** | **string** | The maximum length of time that a one-time password value should be considered valid. | 
**AccessTokenValidator** | Pointer to **[]string** | An access token validator that will ensure that each presented OAuth access token is authentic and trustworthy. It must be configured with an identity mapper that will be used to map the access token to a local entry. | [optional] 
**IdTokenValidator** | Pointer to **[]string** | An ID token validator that will ensure that each presented OpenID Connect ID token is authentic and trustworthy, and that will map the token to a local entry. | [optional] 
**RequireBothAccessTokenAndIDToken** | Pointer to **bool** | Indicates whether bind requests will be required to have both an OAuth access token (in the \&quot;auth\&quot; element of the bind request) and an OpenID Connect ID token (in the \&quot;pingidentityidtoken\&quot; element of the bind request). | [optional] 
**ValidateAccessTokenWhenIDTokenIsAlsoProvided** | Pointer to [**EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp**](EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp.md) |  | [optional] 
**AlternateAuthorizationIdentityMapper** | Pointer to **string** | The identity mapper that will be used to map an alternate authorization identity (provided in the GS2 header of the encoded OAUTHBEARER bind request credentials) to the corresponding local entry. | [optional] 
**AllRequiredScope** | Pointer to **[]string** | The set of OAuth scopes that will all be required for any access tokens that will be allowed for authentication. | [optional] 
**AnyRequiredScope** | Pointer to **[]string** | The set of OAuth scopes that a token may have to be allowed for authentication. | [optional] 
**ServerFqdn** | Pointer to **string** | The fully-qualified name that clients are expected to use when communicating with the server. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party SASL Mechanism Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party SASL Mechanism Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddSaslMechanismHandler200Response

`func NewAddSaslMechanismHandler200Response(id string, schemas []EnumthirdPartySaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool, otpValidityDuration string, extensionClass string, ) *AddSaslMechanismHandler200Response`

NewAddSaslMechanismHandler200Response instantiates a new AddSaslMechanismHandler200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSaslMechanismHandler200ResponseWithDefaults

`func NewAddSaslMechanismHandler200ResponseWithDefaults() *AddSaslMechanismHandler200Response`

NewAddSaslMechanismHandler200ResponseWithDefaults instantiates a new AddSaslMechanismHandler200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddSaslMechanismHandler200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddSaslMechanismHandler200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddSaslMechanismHandler200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddSaslMechanismHandler200Response) GetSchemas() []EnumthirdPartySaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSaslMechanismHandler200Response) GetSchemasOk() (*[]EnumthirdPartySaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSaslMechanismHandler200Response) SetSchemas(v []EnumthirdPartySaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetIdentityMapper

`func (o *AddSaslMechanismHandler200Response) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddSaslMechanismHandler200Response) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddSaslMechanismHandler200Response) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetDescription

`func (o *AddSaslMechanismHandler200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSaslMechanismHandler200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSaslMechanismHandler200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSaslMechanismHandler200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddSaslMechanismHandler200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSaslMechanismHandler200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSaslMechanismHandler200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *AddSaslMechanismHandler200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddSaslMechanismHandler200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddSaslMechanismHandler200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddSaslMechanismHandler200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetOtpValidityDuration

`func (o *AddSaslMechanismHandler200Response) GetOtpValidityDuration() string`

GetOtpValidityDuration returns the OtpValidityDuration field if non-nil, zero value otherwise.

### GetOtpValidityDurationOk

`func (o *AddSaslMechanismHandler200Response) GetOtpValidityDurationOk() (*string, bool)`

GetOtpValidityDurationOk returns a tuple with the OtpValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpValidityDuration

`func (o *AddSaslMechanismHandler200Response) SetOtpValidityDuration(v string)`

SetOtpValidityDuration sets OtpValidityDuration field to given value.


### GetAccessTokenValidator

`func (o *AddSaslMechanismHandler200Response) GetAccessTokenValidator() []string`

GetAccessTokenValidator returns the AccessTokenValidator field if non-nil, zero value otherwise.

### GetAccessTokenValidatorOk

`func (o *AddSaslMechanismHandler200Response) GetAccessTokenValidatorOk() (*[]string, bool)`

GetAccessTokenValidatorOk returns a tuple with the AccessTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenValidator

`func (o *AddSaslMechanismHandler200Response) SetAccessTokenValidator(v []string)`

SetAccessTokenValidator sets AccessTokenValidator field to given value.

### HasAccessTokenValidator

`func (o *AddSaslMechanismHandler200Response) HasAccessTokenValidator() bool`

HasAccessTokenValidator returns a boolean if a field has been set.

### GetIdTokenValidator

`func (o *AddSaslMechanismHandler200Response) GetIdTokenValidator() []string`

GetIdTokenValidator returns the IdTokenValidator field if non-nil, zero value otherwise.

### GetIdTokenValidatorOk

`func (o *AddSaslMechanismHandler200Response) GetIdTokenValidatorOk() (*[]string, bool)`

GetIdTokenValidatorOk returns a tuple with the IdTokenValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdTokenValidator

`func (o *AddSaslMechanismHandler200Response) SetIdTokenValidator(v []string)`

SetIdTokenValidator sets IdTokenValidator field to given value.

### HasIdTokenValidator

`func (o *AddSaslMechanismHandler200Response) HasIdTokenValidator() bool`

HasIdTokenValidator returns a boolean if a field has been set.

### GetRequireBothAccessTokenAndIDToken

`func (o *AddSaslMechanismHandler200Response) GetRequireBothAccessTokenAndIDToken() bool`

GetRequireBothAccessTokenAndIDToken returns the RequireBothAccessTokenAndIDToken field if non-nil, zero value otherwise.

### GetRequireBothAccessTokenAndIDTokenOk

`func (o *AddSaslMechanismHandler200Response) GetRequireBothAccessTokenAndIDTokenOk() (*bool, bool)`

GetRequireBothAccessTokenAndIDTokenOk returns a tuple with the RequireBothAccessTokenAndIDToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireBothAccessTokenAndIDToken

`func (o *AddSaslMechanismHandler200Response) SetRequireBothAccessTokenAndIDToken(v bool)`

SetRequireBothAccessTokenAndIDToken sets RequireBothAccessTokenAndIDToken field to given value.

### HasRequireBothAccessTokenAndIDToken

`func (o *AddSaslMechanismHandler200Response) HasRequireBothAccessTokenAndIDToken() bool`

HasRequireBothAccessTokenAndIDToken returns a boolean if a field has been set.

### GetValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *AddSaslMechanismHandler200Response) GetValidateAccessTokenWhenIDTokenIsAlsoProvided() EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp`

GetValidateAccessTokenWhenIDTokenIsAlsoProvided returns the ValidateAccessTokenWhenIDTokenIsAlsoProvided field if non-nil, zero value otherwise.

### GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk

`func (o *AddSaslMechanismHandler200Response) GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk() (*EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp, bool)`

GetValidateAccessTokenWhenIDTokenIsAlsoProvidedOk returns a tuple with the ValidateAccessTokenWhenIDTokenIsAlsoProvided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *AddSaslMechanismHandler200Response) SetValidateAccessTokenWhenIDTokenIsAlsoProvided(v EnumsaslMechanismHandlerValidateAccessTokenWhenIDTokenIsAlsoProvidedProp)`

SetValidateAccessTokenWhenIDTokenIsAlsoProvided sets ValidateAccessTokenWhenIDTokenIsAlsoProvided field to given value.

### HasValidateAccessTokenWhenIDTokenIsAlsoProvided

`func (o *AddSaslMechanismHandler200Response) HasValidateAccessTokenWhenIDTokenIsAlsoProvided() bool`

HasValidateAccessTokenWhenIDTokenIsAlsoProvided returns a boolean if a field has been set.

### GetAlternateAuthorizationIdentityMapper

`func (o *AddSaslMechanismHandler200Response) GetAlternateAuthorizationIdentityMapper() string`

GetAlternateAuthorizationIdentityMapper returns the AlternateAuthorizationIdentityMapper field if non-nil, zero value otherwise.

### GetAlternateAuthorizationIdentityMapperOk

`func (o *AddSaslMechanismHandler200Response) GetAlternateAuthorizationIdentityMapperOk() (*string, bool)`

GetAlternateAuthorizationIdentityMapperOk returns a tuple with the AlternateAuthorizationIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateAuthorizationIdentityMapper

`func (o *AddSaslMechanismHandler200Response) SetAlternateAuthorizationIdentityMapper(v string)`

SetAlternateAuthorizationIdentityMapper sets AlternateAuthorizationIdentityMapper field to given value.

### HasAlternateAuthorizationIdentityMapper

`func (o *AddSaslMechanismHandler200Response) HasAlternateAuthorizationIdentityMapper() bool`

HasAlternateAuthorizationIdentityMapper returns a boolean if a field has been set.

### GetAllRequiredScope

`func (o *AddSaslMechanismHandler200Response) GetAllRequiredScope() []string`

GetAllRequiredScope returns the AllRequiredScope field if non-nil, zero value otherwise.

### GetAllRequiredScopeOk

`func (o *AddSaslMechanismHandler200Response) GetAllRequiredScopeOk() (*[]string, bool)`

GetAllRequiredScopeOk returns a tuple with the AllRequiredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllRequiredScope

`func (o *AddSaslMechanismHandler200Response) SetAllRequiredScope(v []string)`

SetAllRequiredScope sets AllRequiredScope field to given value.

### HasAllRequiredScope

`func (o *AddSaslMechanismHandler200Response) HasAllRequiredScope() bool`

HasAllRequiredScope returns a boolean if a field has been set.

### GetAnyRequiredScope

`func (o *AddSaslMechanismHandler200Response) GetAnyRequiredScope() []string`

GetAnyRequiredScope returns the AnyRequiredScope field if non-nil, zero value otherwise.

### GetAnyRequiredScopeOk

`func (o *AddSaslMechanismHandler200Response) GetAnyRequiredScopeOk() (*[]string, bool)`

GetAnyRequiredScopeOk returns a tuple with the AnyRequiredScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyRequiredScope

`func (o *AddSaslMechanismHandler200Response) SetAnyRequiredScope(v []string)`

SetAnyRequiredScope sets AnyRequiredScope field to given value.

### HasAnyRequiredScope

`func (o *AddSaslMechanismHandler200Response) HasAnyRequiredScope() bool`

HasAnyRequiredScope returns a boolean if a field has been set.

### GetServerFqdn

`func (o *AddSaslMechanismHandler200Response) GetServerFqdn() string`

GetServerFqdn returns the ServerFqdn field if non-nil, zero value otherwise.

### GetServerFqdnOk

`func (o *AddSaslMechanismHandler200Response) GetServerFqdnOk() (*string, bool)`

GetServerFqdnOk returns a tuple with the ServerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFqdn

`func (o *AddSaslMechanismHandler200Response) SetServerFqdn(v string)`

SetServerFqdn sets ServerFqdn field to given value.

### HasServerFqdn

`func (o *AddSaslMechanismHandler200Response) HasServerFqdn() bool`

HasServerFqdn returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddSaslMechanismHandler200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddSaslMechanismHandler200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddSaslMechanismHandler200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddSaslMechanismHandler200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddSaslMechanismHandler200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddSaslMechanismHandler200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddSaslMechanismHandler200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


