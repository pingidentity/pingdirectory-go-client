# AddExtendedOperationHandlerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandlerName** | **string** | Name of the new Extended Operation Handler | 
**Schemas** | [**[]EnumthirdPartyExtendedOperationHandlerSchemaUrn**](EnumthirdPartyExtendedOperationHandlerSchemaUrn.md) |  | 
**SharedSecretAttributeType** | Pointer to **string** | The name or OID of the attribute that will be used to hold the shared secret key used during TOTP processing. | [optional] 
**TimeIntervalDuration** | Pointer to **string** | The duration of the time interval used for TOTP processing. | [optional] 
**AdjacentIntervalsToCheck** | Pointer to **int32** | The number of adjacent time intervals (both before and after the current time) that should be checked when performing authentication. | [optional] 
**PreventTOTPReuse** | Pointer to **bool** | Indicates whether to prevent clients from re-using TOTP passwords. | [optional] 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**AllowRemotelyProvidedCertificates** | Pointer to **bool** | Indicates whether clients should be allowed to directly provide a new listener or inter-server certificate chain in the extended request. | [optional] 
**AllowedOperation** | Pointer to [**[]EnumextendedOperationHandlerAllowedOperationProp**](EnumextendedOperationHandlerAllowedOperationProp.md) | The types of replace certificate operations that clients will be allowed to request. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A set of criteria that client connections must satisfy before they will be allowed to request the associated extended operations. | [optional] 
**RequestCriteria** | Pointer to **string** | A set of criteria that the extended requests must satisfy before they will be processed by the server. | [optional] 
**PasswordGenerator** | **string** | The password generator that will be used to create the one-time password values to be delivered to the end user. | 
**DefaultOTPDeliveryMechanism** | **[]string** | The set of delivery mechanisms that may be used to deliver one-time passwords to users in requests that do not specify one or more preferred delivery mechanisms. | 
**DefaultSingleUseTokenValidityDuration** | Pointer to **string** | The default length of time that a single-use token will be considered valid by the server if the client doesn&#39;t specify a duration in the deliver single-use token request. | [optional] 
**DefaultTokenDeliveryMechanism** | **[]string** | The set of delivery mechanisms that may be used to deliver password reset tokens to users for requests that do not specify one or more preferred delivery mechanisms. | 
**PasswordResetTokenValidityDuration** | Pointer to **string** | The maximum length of time that a password reset token should be considered valid. | [optional] 
**IdentityMapper** | **string** | The identity mapper that should be used to identify the user(s) targeted by the authentication identity contained in the extended request. This will only be used for \&quot;u:\&quot;-style authentication identities. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Extended Operation Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Extended Operation Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddExtendedOperationHandlerRequest

`func NewAddExtendedOperationHandlerRequest(handlerName string, schemas []EnumthirdPartyExtendedOperationHandlerSchemaUrn, enabled bool, passwordGenerator string, defaultOTPDeliveryMechanism []string, defaultTokenDeliveryMechanism []string, identityMapper string, extensionClass string, ) *AddExtendedOperationHandlerRequest`

NewAddExtendedOperationHandlerRequest instantiates a new AddExtendedOperationHandlerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddExtendedOperationHandlerRequestWithDefaults

`func NewAddExtendedOperationHandlerRequestWithDefaults() *AddExtendedOperationHandlerRequest`

NewAddExtendedOperationHandlerRequestWithDefaults instantiates a new AddExtendedOperationHandlerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandlerName

`func (o *AddExtendedOperationHandlerRequest) GetHandlerName() string`

GetHandlerName returns the HandlerName field if non-nil, zero value otherwise.

### GetHandlerNameOk

`func (o *AddExtendedOperationHandlerRequest) GetHandlerNameOk() (*string, bool)`

GetHandlerNameOk returns a tuple with the HandlerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerName

`func (o *AddExtendedOperationHandlerRequest) SetHandlerName(v string)`

SetHandlerName sets HandlerName field to given value.


### GetSchemas

`func (o *AddExtendedOperationHandlerRequest) GetSchemas() []EnumthirdPartyExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddExtendedOperationHandlerRequest) GetSchemasOk() (*[]EnumthirdPartyExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddExtendedOperationHandlerRequest) SetSchemas(v []EnumthirdPartyExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSharedSecretAttributeType

`func (o *AddExtendedOperationHandlerRequest) GetSharedSecretAttributeType() string`

GetSharedSecretAttributeType returns the SharedSecretAttributeType field if non-nil, zero value otherwise.

### GetSharedSecretAttributeTypeOk

`func (o *AddExtendedOperationHandlerRequest) GetSharedSecretAttributeTypeOk() (*string, bool)`

GetSharedSecretAttributeTypeOk returns a tuple with the SharedSecretAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretAttributeType

`func (o *AddExtendedOperationHandlerRequest) SetSharedSecretAttributeType(v string)`

SetSharedSecretAttributeType sets SharedSecretAttributeType field to given value.

### HasSharedSecretAttributeType

`func (o *AddExtendedOperationHandlerRequest) HasSharedSecretAttributeType() bool`

HasSharedSecretAttributeType returns a boolean if a field has been set.

### GetTimeIntervalDuration

`func (o *AddExtendedOperationHandlerRequest) GetTimeIntervalDuration() string`

GetTimeIntervalDuration returns the TimeIntervalDuration field if non-nil, zero value otherwise.

### GetTimeIntervalDurationOk

`func (o *AddExtendedOperationHandlerRequest) GetTimeIntervalDurationOk() (*string, bool)`

GetTimeIntervalDurationOk returns a tuple with the TimeIntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalDuration

`func (o *AddExtendedOperationHandlerRequest) SetTimeIntervalDuration(v string)`

SetTimeIntervalDuration sets TimeIntervalDuration field to given value.

### HasTimeIntervalDuration

`func (o *AddExtendedOperationHandlerRequest) HasTimeIntervalDuration() bool`

HasTimeIntervalDuration returns a boolean if a field has been set.

### GetAdjacentIntervalsToCheck

`func (o *AddExtendedOperationHandlerRequest) GetAdjacentIntervalsToCheck() int32`

GetAdjacentIntervalsToCheck returns the AdjacentIntervalsToCheck field if non-nil, zero value otherwise.

### GetAdjacentIntervalsToCheckOk

`func (o *AddExtendedOperationHandlerRequest) GetAdjacentIntervalsToCheckOk() (*int32, bool)`

GetAdjacentIntervalsToCheckOk returns a tuple with the AdjacentIntervalsToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentIntervalsToCheck

`func (o *AddExtendedOperationHandlerRequest) SetAdjacentIntervalsToCheck(v int32)`

SetAdjacentIntervalsToCheck sets AdjacentIntervalsToCheck field to given value.

### HasAdjacentIntervalsToCheck

`func (o *AddExtendedOperationHandlerRequest) HasAdjacentIntervalsToCheck() bool`

HasAdjacentIntervalsToCheck returns a boolean if a field has been set.

### GetPreventTOTPReuse

`func (o *AddExtendedOperationHandlerRequest) GetPreventTOTPReuse() bool`

GetPreventTOTPReuse returns the PreventTOTPReuse field if non-nil, zero value otherwise.

### GetPreventTOTPReuseOk

`func (o *AddExtendedOperationHandlerRequest) GetPreventTOTPReuseOk() (*bool, bool)`

GetPreventTOTPReuseOk returns a tuple with the PreventTOTPReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventTOTPReuse

`func (o *AddExtendedOperationHandlerRequest) SetPreventTOTPReuse(v bool)`

SetPreventTOTPReuse sets PreventTOTPReuse field to given value.

### HasPreventTOTPReuse

`func (o *AddExtendedOperationHandlerRequest) HasPreventTOTPReuse() bool`

HasPreventTOTPReuse returns a boolean if a field has been set.

### GetDescription

`func (o *AddExtendedOperationHandlerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddExtendedOperationHandlerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddExtendedOperationHandlerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddExtendedOperationHandlerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddExtendedOperationHandlerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddExtendedOperationHandlerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddExtendedOperationHandlerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAllowRemotelyProvidedCertificates

`func (o *AddExtendedOperationHandlerRequest) GetAllowRemotelyProvidedCertificates() bool`

GetAllowRemotelyProvidedCertificates returns the AllowRemotelyProvidedCertificates field if non-nil, zero value otherwise.

### GetAllowRemotelyProvidedCertificatesOk

`func (o *AddExtendedOperationHandlerRequest) GetAllowRemotelyProvidedCertificatesOk() (*bool, bool)`

GetAllowRemotelyProvidedCertificatesOk returns a tuple with the AllowRemotelyProvidedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemotelyProvidedCertificates

`func (o *AddExtendedOperationHandlerRequest) SetAllowRemotelyProvidedCertificates(v bool)`

SetAllowRemotelyProvidedCertificates sets AllowRemotelyProvidedCertificates field to given value.

### HasAllowRemotelyProvidedCertificates

`func (o *AddExtendedOperationHandlerRequest) HasAllowRemotelyProvidedCertificates() bool`

HasAllowRemotelyProvidedCertificates returns a boolean if a field has been set.

### GetAllowedOperation

`func (o *AddExtendedOperationHandlerRequest) GetAllowedOperation() []EnumextendedOperationHandlerAllowedOperationProp`

GetAllowedOperation returns the AllowedOperation field if non-nil, zero value otherwise.

### GetAllowedOperationOk

`func (o *AddExtendedOperationHandlerRequest) GetAllowedOperationOk() (*[]EnumextendedOperationHandlerAllowedOperationProp, bool)`

GetAllowedOperationOk returns a tuple with the AllowedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperation

`func (o *AddExtendedOperationHandlerRequest) SetAllowedOperation(v []EnumextendedOperationHandlerAllowedOperationProp)`

SetAllowedOperation sets AllowedOperation field to given value.

### HasAllowedOperation

`func (o *AddExtendedOperationHandlerRequest) HasAllowedOperation() bool`

HasAllowedOperation returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddExtendedOperationHandlerRequest) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddExtendedOperationHandlerRequest) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddExtendedOperationHandlerRequest) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddExtendedOperationHandlerRequest) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddExtendedOperationHandlerRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddExtendedOperationHandlerRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddExtendedOperationHandlerRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddExtendedOperationHandlerRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetPasswordGenerator

`func (o *AddExtendedOperationHandlerRequest) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *AddExtendedOperationHandlerRequest) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *AddExtendedOperationHandlerRequest) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultOTPDeliveryMechanism

`func (o *AddExtendedOperationHandlerRequest) GetDefaultOTPDeliveryMechanism() []string`

GetDefaultOTPDeliveryMechanism returns the DefaultOTPDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultOTPDeliveryMechanismOk

`func (o *AddExtendedOperationHandlerRequest) GetDefaultOTPDeliveryMechanismOk() (*[]string, bool)`

GetDefaultOTPDeliveryMechanismOk returns a tuple with the DefaultOTPDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOTPDeliveryMechanism

`func (o *AddExtendedOperationHandlerRequest) SetDefaultOTPDeliveryMechanism(v []string)`

SetDefaultOTPDeliveryMechanism sets DefaultOTPDeliveryMechanism field to given value.


### GetDefaultSingleUseTokenValidityDuration

`func (o *AddExtendedOperationHandlerRequest) GetDefaultSingleUseTokenValidityDuration() string`

GetDefaultSingleUseTokenValidityDuration returns the DefaultSingleUseTokenValidityDuration field if non-nil, zero value otherwise.

### GetDefaultSingleUseTokenValidityDurationOk

`func (o *AddExtendedOperationHandlerRequest) GetDefaultSingleUseTokenValidityDurationOk() (*string, bool)`

GetDefaultSingleUseTokenValidityDurationOk returns a tuple with the DefaultSingleUseTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSingleUseTokenValidityDuration

`func (o *AddExtendedOperationHandlerRequest) SetDefaultSingleUseTokenValidityDuration(v string)`

SetDefaultSingleUseTokenValidityDuration sets DefaultSingleUseTokenValidityDuration field to given value.

### HasDefaultSingleUseTokenValidityDuration

`func (o *AddExtendedOperationHandlerRequest) HasDefaultSingleUseTokenValidityDuration() bool`

HasDefaultSingleUseTokenValidityDuration returns a boolean if a field has been set.

### GetDefaultTokenDeliveryMechanism

`func (o *AddExtendedOperationHandlerRequest) GetDefaultTokenDeliveryMechanism() []string`

GetDefaultTokenDeliveryMechanism returns the DefaultTokenDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultTokenDeliveryMechanismOk

`func (o *AddExtendedOperationHandlerRequest) GetDefaultTokenDeliveryMechanismOk() (*[]string, bool)`

GetDefaultTokenDeliveryMechanismOk returns a tuple with the DefaultTokenDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenDeliveryMechanism

`func (o *AddExtendedOperationHandlerRequest) SetDefaultTokenDeliveryMechanism(v []string)`

SetDefaultTokenDeliveryMechanism sets DefaultTokenDeliveryMechanism field to given value.


### GetPasswordResetTokenValidityDuration

`func (o *AddExtendedOperationHandlerRequest) GetPasswordResetTokenValidityDuration() string`

GetPasswordResetTokenValidityDuration returns the PasswordResetTokenValidityDuration field if non-nil, zero value otherwise.

### GetPasswordResetTokenValidityDurationOk

`func (o *AddExtendedOperationHandlerRequest) GetPasswordResetTokenValidityDurationOk() (*string, bool)`

GetPasswordResetTokenValidityDurationOk returns a tuple with the PasswordResetTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenValidityDuration

`func (o *AddExtendedOperationHandlerRequest) SetPasswordResetTokenValidityDuration(v string)`

SetPasswordResetTokenValidityDuration sets PasswordResetTokenValidityDuration field to given value.

### HasPasswordResetTokenValidityDuration

`func (o *AddExtendedOperationHandlerRequest) HasPasswordResetTokenValidityDuration() bool`

HasPasswordResetTokenValidityDuration returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *AddExtendedOperationHandlerRequest) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *AddExtendedOperationHandlerRequest) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *AddExtendedOperationHandlerRequest) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetExtensionClass

`func (o *AddExtendedOperationHandlerRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddExtendedOperationHandlerRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddExtendedOperationHandlerRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddExtendedOperationHandlerRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddExtendedOperationHandlerRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddExtendedOperationHandlerRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddExtendedOperationHandlerRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


