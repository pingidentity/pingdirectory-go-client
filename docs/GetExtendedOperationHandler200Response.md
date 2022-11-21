# GetExtendedOperationHandler200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumthirdPartyExtendedOperationHandlerSchemaUrn**](EnumthirdPartyExtendedOperationHandlerSchemaUrn.md) |  | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**Id** | **string** | Name of the Extended Operation Handler | 
**SharedSecretAttributeType** | Pointer to **string** | The name or OID of the attribute that will be used to hold the shared secret key used during TOTP processing. | [optional] 
**TimeIntervalDuration** | Pointer to **string** | The duration of the time interval used for TOTP processing. | [optional] 
**AdjacentIntervalsToCheck** | Pointer to **int32** | The number of adjacent time intervals (both before and after the current time) that should be checked when performing authentication. | [optional] 
**PreventTOTPReuse** | Pointer to **bool** | Indicates whether to prevent clients from re-using TOTP passwords. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**AllowRemotelyProvidedCertificates** | Pointer to **bool** | Indicates whether clients should be allowed to directly provide a new listener or inter-server certificate chain in the extended request. | [optional] 
**AllowedOperation** | Pointer to [**[]EnumextendedOperationHandlerAllowedOperationProp**](EnumextendedOperationHandlerAllowedOperationProp.md) |  | [optional] 
**ConnectionCriteria** | Pointer to **string** | A set of criteria that client connections must satisfy before they will be allowed to request the associated extended operations. | [optional] 
**RequestCriteria** | Pointer to **string** | A set of criteria that the extended requests must satisfy before they will be processed by the server. | [optional] 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user(s) targeted by the authentication identity contained in the extended request. This will only be used for \&quot;u:\&quot;-style authentication identities. | 
**PasswordGenerator** | **string** | The password generator that will be used to create the one-time password values to be delivered to the end user. | 
**DefaultOTPDeliveryMechanism** | **[]string** | The set of delivery mechanisms that may be used to deliver one-time passwords to users in requests that do not specify one or more preferred delivery mechanisms. | 
**DefaultSingleUseTokenValidityDuration** | Pointer to **string** | The default length of time that a single-use token will be considered valid by the server if the client doesn&#39;t specify a duration in the deliver single-use token request. | [optional] 
**DefaultPasswordPolicy** | Pointer to **string** | The default password policy that should be used when generating and validating passwords if the request does not specify an alternate policy. If this is not provided, then this Generate Password Extended Operation Handler will use the default password policy defined in the global configuration. | [optional] 
**DefaultPasswordGenerator** | **string** | The default password generator that will be used if the selected password policy is not configured with a password generator. | 
**MaximumPasswordsPerRequest** | Pointer to **int32** | The maximum number of passwords that may be generated and returned to the client for a single request. | [optional] 
**MaximumValidationAttemptsPerPassword** | Pointer to **int32** | The maximum number of attempts that the server may use to generate a password that passes validation. | [optional] 
**DefaultTokenDeliveryMechanism** | **[]string** | The set of delivery mechanisms that may be used to deliver password reset tokens to users for requests that do not specify one or more preferred delivery mechanisms. | 
**PasswordResetTokenValidityDuration** | **string** | The maximum length of time that a password reset token should be considered valid. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Extended Operation Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Extended Operation Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewGetExtendedOperationHandler200Response

`func NewGetExtendedOperationHandler200Response(schemas []EnumthirdPartyExtendedOperationHandlerSchemaUrn, enabled bool, id string, identityMapper string, passwordGenerator string, defaultOTPDeliveryMechanism []string, defaultPasswordGenerator string, defaultTokenDeliveryMechanism []string, passwordResetTokenValidityDuration string, extensionClass string, ) *GetExtendedOperationHandler200Response`

NewGetExtendedOperationHandler200Response instantiates a new GetExtendedOperationHandler200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExtendedOperationHandler200ResponseWithDefaults

`func NewGetExtendedOperationHandler200ResponseWithDefaults() *GetExtendedOperationHandler200Response`

NewGetExtendedOperationHandler200ResponseWithDefaults instantiates a new GetExtendedOperationHandler200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetExtendedOperationHandler200Response) GetSchemas() []EnumthirdPartyExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetExtendedOperationHandler200Response) GetSchemasOk() (*[]EnumthirdPartyExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetExtendedOperationHandler200Response) SetSchemas(v []EnumthirdPartyExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetDescription

`func (o *GetExtendedOperationHandler200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetExtendedOperationHandler200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetExtendedOperationHandler200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetExtendedOperationHandler200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetExtendedOperationHandler200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetExtendedOperationHandler200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetExtendedOperationHandler200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *GetExtendedOperationHandler200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetExtendedOperationHandler200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetExtendedOperationHandler200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSharedSecretAttributeType

`func (o *GetExtendedOperationHandler200Response) GetSharedSecretAttributeType() string`

GetSharedSecretAttributeType returns the SharedSecretAttributeType field if non-nil, zero value otherwise.

### GetSharedSecretAttributeTypeOk

`func (o *GetExtendedOperationHandler200Response) GetSharedSecretAttributeTypeOk() (*string, bool)`

GetSharedSecretAttributeTypeOk returns a tuple with the SharedSecretAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretAttributeType

`func (o *GetExtendedOperationHandler200Response) SetSharedSecretAttributeType(v string)`

SetSharedSecretAttributeType sets SharedSecretAttributeType field to given value.

### HasSharedSecretAttributeType

`func (o *GetExtendedOperationHandler200Response) HasSharedSecretAttributeType() bool`

HasSharedSecretAttributeType returns a boolean if a field has been set.

### GetTimeIntervalDuration

`func (o *GetExtendedOperationHandler200Response) GetTimeIntervalDuration() string`

GetTimeIntervalDuration returns the TimeIntervalDuration field if non-nil, zero value otherwise.

### GetTimeIntervalDurationOk

`func (o *GetExtendedOperationHandler200Response) GetTimeIntervalDurationOk() (*string, bool)`

GetTimeIntervalDurationOk returns a tuple with the TimeIntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalDuration

`func (o *GetExtendedOperationHandler200Response) SetTimeIntervalDuration(v string)`

SetTimeIntervalDuration sets TimeIntervalDuration field to given value.

### HasTimeIntervalDuration

`func (o *GetExtendedOperationHandler200Response) HasTimeIntervalDuration() bool`

HasTimeIntervalDuration returns a boolean if a field has been set.

### GetAdjacentIntervalsToCheck

`func (o *GetExtendedOperationHandler200Response) GetAdjacentIntervalsToCheck() int32`

GetAdjacentIntervalsToCheck returns the AdjacentIntervalsToCheck field if non-nil, zero value otherwise.

### GetAdjacentIntervalsToCheckOk

`func (o *GetExtendedOperationHandler200Response) GetAdjacentIntervalsToCheckOk() (*int32, bool)`

GetAdjacentIntervalsToCheckOk returns a tuple with the AdjacentIntervalsToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentIntervalsToCheck

`func (o *GetExtendedOperationHandler200Response) SetAdjacentIntervalsToCheck(v int32)`

SetAdjacentIntervalsToCheck sets AdjacentIntervalsToCheck field to given value.

### HasAdjacentIntervalsToCheck

`func (o *GetExtendedOperationHandler200Response) HasAdjacentIntervalsToCheck() bool`

HasAdjacentIntervalsToCheck returns a boolean if a field has been set.

### GetPreventTOTPReuse

`func (o *GetExtendedOperationHandler200Response) GetPreventTOTPReuse() bool`

GetPreventTOTPReuse returns the PreventTOTPReuse field if non-nil, zero value otherwise.

### GetPreventTOTPReuseOk

`func (o *GetExtendedOperationHandler200Response) GetPreventTOTPReuseOk() (*bool, bool)`

GetPreventTOTPReuseOk returns a tuple with the PreventTOTPReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventTOTPReuse

`func (o *GetExtendedOperationHandler200Response) SetPreventTOTPReuse(v bool)`

SetPreventTOTPReuse sets PreventTOTPReuse field to given value.

### HasPreventTOTPReuse

`func (o *GetExtendedOperationHandler200Response) HasPreventTOTPReuse() bool`

HasPreventTOTPReuse returns a boolean if a field has been set.

### GetMeta

`func (o *GetExtendedOperationHandler200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetExtendedOperationHandler200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetExtendedOperationHandler200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetExtendedOperationHandler200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetAllowRemotelyProvidedCertificates

`func (o *GetExtendedOperationHandler200Response) GetAllowRemotelyProvidedCertificates() bool`

GetAllowRemotelyProvidedCertificates returns the AllowRemotelyProvidedCertificates field if non-nil, zero value otherwise.

### GetAllowRemotelyProvidedCertificatesOk

`func (o *GetExtendedOperationHandler200Response) GetAllowRemotelyProvidedCertificatesOk() (*bool, bool)`

GetAllowRemotelyProvidedCertificatesOk returns a tuple with the AllowRemotelyProvidedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemotelyProvidedCertificates

`func (o *GetExtendedOperationHandler200Response) SetAllowRemotelyProvidedCertificates(v bool)`

SetAllowRemotelyProvidedCertificates sets AllowRemotelyProvidedCertificates field to given value.

### HasAllowRemotelyProvidedCertificates

`func (o *GetExtendedOperationHandler200Response) HasAllowRemotelyProvidedCertificates() bool`

HasAllowRemotelyProvidedCertificates returns a boolean if a field has been set.

### GetAllowedOperation

`func (o *GetExtendedOperationHandler200Response) GetAllowedOperation() []EnumextendedOperationHandlerAllowedOperationProp`

GetAllowedOperation returns the AllowedOperation field if non-nil, zero value otherwise.

### GetAllowedOperationOk

`func (o *GetExtendedOperationHandler200Response) GetAllowedOperationOk() (*[]EnumextendedOperationHandlerAllowedOperationProp, bool)`

GetAllowedOperationOk returns a tuple with the AllowedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperation

`func (o *GetExtendedOperationHandler200Response) SetAllowedOperation(v []EnumextendedOperationHandlerAllowedOperationProp)`

SetAllowedOperation sets AllowedOperation field to given value.

### HasAllowedOperation

`func (o *GetExtendedOperationHandler200Response) HasAllowedOperation() bool`

HasAllowedOperation returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *GetExtendedOperationHandler200Response) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *GetExtendedOperationHandler200Response) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *GetExtendedOperationHandler200Response) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *GetExtendedOperationHandler200Response) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *GetExtendedOperationHandler200Response) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *GetExtendedOperationHandler200Response) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *GetExtendedOperationHandler200Response) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *GetExtendedOperationHandler200Response) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *GetExtendedOperationHandler200Response) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *GetExtendedOperationHandler200Response) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *GetExtendedOperationHandler200Response) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetPasswordGenerator

`func (o *GetExtendedOperationHandler200Response) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *GetExtendedOperationHandler200Response) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *GetExtendedOperationHandler200Response) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultOTPDeliveryMechanism

`func (o *GetExtendedOperationHandler200Response) GetDefaultOTPDeliveryMechanism() []string`

GetDefaultOTPDeliveryMechanism returns the DefaultOTPDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultOTPDeliveryMechanismOk

`func (o *GetExtendedOperationHandler200Response) GetDefaultOTPDeliveryMechanismOk() (*[]string, bool)`

GetDefaultOTPDeliveryMechanismOk returns a tuple with the DefaultOTPDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOTPDeliveryMechanism

`func (o *GetExtendedOperationHandler200Response) SetDefaultOTPDeliveryMechanism(v []string)`

SetDefaultOTPDeliveryMechanism sets DefaultOTPDeliveryMechanism field to given value.


### GetDefaultSingleUseTokenValidityDuration

`func (o *GetExtendedOperationHandler200Response) GetDefaultSingleUseTokenValidityDuration() string`

GetDefaultSingleUseTokenValidityDuration returns the DefaultSingleUseTokenValidityDuration field if non-nil, zero value otherwise.

### GetDefaultSingleUseTokenValidityDurationOk

`func (o *GetExtendedOperationHandler200Response) GetDefaultSingleUseTokenValidityDurationOk() (*string, bool)`

GetDefaultSingleUseTokenValidityDurationOk returns a tuple with the DefaultSingleUseTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSingleUseTokenValidityDuration

`func (o *GetExtendedOperationHandler200Response) SetDefaultSingleUseTokenValidityDuration(v string)`

SetDefaultSingleUseTokenValidityDuration sets DefaultSingleUseTokenValidityDuration field to given value.

### HasDefaultSingleUseTokenValidityDuration

`func (o *GetExtendedOperationHandler200Response) HasDefaultSingleUseTokenValidityDuration() bool`

HasDefaultSingleUseTokenValidityDuration returns a boolean if a field has been set.

### GetDefaultPasswordPolicy

`func (o *GetExtendedOperationHandler200Response) GetDefaultPasswordPolicy() string`

GetDefaultPasswordPolicy returns the DefaultPasswordPolicy field if non-nil, zero value otherwise.

### GetDefaultPasswordPolicyOk

`func (o *GetExtendedOperationHandler200Response) GetDefaultPasswordPolicyOk() (*string, bool)`

GetDefaultPasswordPolicyOk returns a tuple with the DefaultPasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordPolicy

`func (o *GetExtendedOperationHandler200Response) SetDefaultPasswordPolicy(v string)`

SetDefaultPasswordPolicy sets DefaultPasswordPolicy field to given value.

### HasDefaultPasswordPolicy

`func (o *GetExtendedOperationHandler200Response) HasDefaultPasswordPolicy() bool`

HasDefaultPasswordPolicy returns a boolean if a field has been set.

### GetDefaultPasswordGenerator

`func (o *GetExtendedOperationHandler200Response) GetDefaultPasswordGenerator() string`

GetDefaultPasswordGenerator returns the DefaultPasswordGenerator field if non-nil, zero value otherwise.

### GetDefaultPasswordGeneratorOk

`func (o *GetExtendedOperationHandler200Response) GetDefaultPasswordGeneratorOk() (*string, bool)`

GetDefaultPasswordGeneratorOk returns a tuple with the DefaultPasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordGenerator

`func (o *GetExtendedOperationHandler200Response) SetDefaultPasswordGenerator(v string)`

SetDefaultPasswordGenerator sets DefaultPasswordGenerator field to given value.


### GetMaximumPasswordsPerRequest

`func (o *GetExtendedOperationHandler200Response) GetMaximumPasswordsPerRequest() int32`

GetMaximumPasswordsPerRequest returns the MaximumPasswordsPerRequest field if non-nil, zero value otherwise.

### GetMaximumPasswordsPerRequestOk

`func (o *GetExtendedOperationHandler200Response) GetMaximumPasswordsPerRequestOk() (*int32, bool)`

GetMaximumPasswordsPerRequestOk returns a tuple with the MaximumPasswordsPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPasswordsPerRequest

`func (o *GetExtendedOperationHandler200Response) SetMaximumPasswordsPerRequest(v int32)`

SetMaximumPasswordsPerRequest sets MaximumPasswordsPerRequest field to given value.

### HasMaximumPasswordsPerRequest

`func (o *GetExtendedOperationHandler200Response) HasMaximumPasswordsPerRequest() bool`

HasMaximumPasswordsPerRequest returns a boolean if a field has been set.

### GetMaximumValidationAttemptsPerPassword

`func (o *GetExtendedOperationHandler200Response) GetMaximumValidationAttemptsPerPassword() int32`

GetMaximumValidationAttemptsPerPassword returns the MaximumValidationAttemptsPerPassword field if non-nil, zero value otherwise.

### GetMaximumValidationAttemptsPerPasswordOk

`func (o *GetExtendedOperationHandler200Response) GetMaximumValidationAttemptsPerPasswordOk() (*int32, bool)`

GetMaximumValidationAttemptsPerPasswordOk returns a tuple with the MaximumValidationAttemptsPerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValidationAttemptsPerPassword

`func (o *GetExtendedOperationHandler200Response) SetMaximumValidationAttemptsPerPassword(v int32)`

SetMaximumValidationAttemptsPerPassword sets MaximumValidationAttemptsPerPassword field to given value.

### HasMaximumValidationAttemptsPerPassword

`func (o *GetExtendedOperationHandler200Response) HasMaximumValidationAttemptsPerPassword() bool`

HasMaximumValidationAttemptsPerPassword returns a boolean if a field has been set.

### GetDefaultTokenDeliveryMechanism

`func (o *GetExtendedOperationHandler200Response) GetDefaultTokenDeliveryMechanism() []string`

GetDefaultTokenDeliveryMechanism returns the DefaultTokenDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultTokenDeliveryMechanismOk

`func (o *GetExtendedOperationHandler200Response) GetDefaultTokenDeliveryMechanismOk() (*[]string, bool)`

GetDefaultTokenDeliveryMechanismOk returns a tuple with the DefaultTokenDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenDeliveryMechanism

`func (o *GetExtendedOperationHandler200Response) SetDefaultTokenDeliveryMechanism(v []string)`

SetDefaultTokenDeliveryMechanism sets DefaultTokenDeliveryMechanism field to given value.


### GetPasswordResetTokenValidityDuration

`func (o *GetExtendedOperationHandler200Response) GetPasswordResetTokenValidityDuration() string`

GetPasswordResetTokenValidityDuration returns the PasswordResetTokenValidityDuration field if non-nil, zero value otherwise.

### GetPasswordResetTokenValidityDurationOk

`func (o *GetExtendedOperationHandler200Response) GetPasswordResetTokenValidityDurationOk() (*string, bool)`

GetPasswordResetTokenValidityDurationOk returns a tuple with the PasswordResetTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenValidityDuration

`func (o *GetExtendedOperationHandler200Response) SetPasswordResetTokenValidityDuration(v string)`

SetPasswordResetTokenValidityDuration sets PasswordResetTokenValidityDuration field to given value.


### GetExtensionClass

`func (o *GetExtendedOperationHandler200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *GetExtendedOperationHandler200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *GetExtendedOperationHandler200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *GetExtendedOperationHandler200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *GetExtendedOperationHandler200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *GetExtendedOperationHandler200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *GetExtendedOperationHandler200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


