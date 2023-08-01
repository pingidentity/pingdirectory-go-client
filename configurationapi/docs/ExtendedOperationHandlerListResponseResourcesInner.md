# ExtendedOperationHandlerListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgetConfigurationExtendedOperationHandlerSchemaUrn**](EnumgetConfigurationExtendedOperationHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the Extended Operation Handler | 
**Description** | Pointer to **string** | A description for this Extended Operation Handler | [optional] 
**Enabled** | **bool** | Indicates whether the Extended Operation Handler is enabled (that is, whether the types of extended operations are allowed in the server). | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**SharedSecretAttributeType** | Pointer to **string** | The name or OID of the attribute that will be used to hold the shared secret key used during TOTP processing. | [optional] 
**TimeIntervalDuration** | Pointer to **string** | The duration of the time interval used for TOTP processing. | [optional] 
**AdjacentIntervalsToCheck** | Pointer to **int64** | The number of adjacent time intervals (both before and after the current time) that should be checked when performing authentication. | [optional] 
**PreventTOTPReuse** | Pointer to **bool** | Indicates whether to prevent clients from re-using TOTP passwords. | [optional] 
**BackendID** | **string** | Specifies the backend ID for the backend in which the transactions will be allowed. Interactive transactions will only be supported for operations in a single backend. | 
**AllowRemotelyProvidedCertificates** | Pointer to **bool** | Indicates whether clients should be allowed to directly provide a new listener or inter-server certificate chain in the extended request. | [optional] 
**AllowedOperation** | Pointer to [**[]EnumextendedOperationHandlerAllowedOperationProp**](EnumextendedOperationHandlerAllowedOperationProp.md) |  | [optional] 
**ConnectionCriteria** | Pointer to **string** | A set of criteria that client connections must satisfy before they will be allowed to request the associated extended operations. | [optional] 
**RequestCriteria** | Pointer to **string** | A set of criteria that the extended requests must satisfy before they will be processed by the server. | [optional] 
**PasswordGenerator** | **string** | The password generator that will be used to create the password reset token values to be delivered to the end user. | 
**DefaultOTPDeliveryMechanism** | **[]string** | The set of delivery mechanisms that may be used to deliver one-time passwords to users in requests that do not specify one or more preferred delivery mechanisms. | 
**DefaultSingleUseTokenValidityDuration** | Pointer to **string** | The default length of time that a single-use token will be considered valid by the server if the client doesn&#39;t specify a duration in the deliver single-use token request. | [optional] 
**ValuesPerStreamResponse** | Pointer to **int64** | The maximum number of values to include per response when responding to a stream values extended request, when the client does not specify a value. | [optional] 
**NonAccessibleSubtreeUnauthorizedBindResultCode** | Pointer to **int64** | The integer value for the result code that the server should return to clients that attempt to perform an unauthorized bind operation in a non-accessible subtree. | [optional] 
**NonAccessibleSubtreeUnauthorizedReadResultCode** | Pointer to **int64** | The integer value for the result code that the server should return to clients that attempt to perform an unauthorized read (e.g., search or compare) operation in a non-accessible subtree. | [optional] 
**NonAccessibleSubtreeUnauthorizedWriteResultCode** | Pointer to **int64** | The integer value for the result code that the server should return to clients that attempt to perform an unauthorized write (e.g., add, delete, modify or modify DN) operation in a non-accessible subtree. | [optional] 
**IdentityMapper** | **string** | Specifies the name of the identity mapper that should be used in conjunction with the password modify extended operation. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Proxied Extended Operation Handler. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Proxied Extended Operation Handler. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**RouteToBackendSetBehavior** | Pointer to [**EnumextendedOperationHandlerRouteToBackendSetBehaviorProp**](EnumextendedOperationHandlerRouteToBackendSetBehaviorProp.md) |  | [optional] 
**DefaultPasswordPolicy** | Pointer to **string** | The default password policy that should be used when generating and validating passwords if the request does not specify an alternate policy. If this is not provided, then this Generate Password Extended Operation Handler will use the default password policy defined in the global configuration. | [optional] 
**DefaultPasswordGenerator** | **string** | The default password generator that will be used if the selected password policy is not configured with a password generator. | 
**MaximumPasswordsPerRequest** | Pointer to **int64** | The maximum number of passwords that may be generated and returned to the client for a single request. | [optional] 
**MaximumValidationAttemptsPerPassword** | Pointer to **int64** | The maximum number of attempts that the server may use to generate a password that passes validation. | [optional] 
**DefaultTokenDeliveryMechanism** | **[]string** | The set of delivery mechanisms that may be used to deliver password reset tokens to users for requests that do not specify one or more preferred delivery mechanisms. | 
**PasswordResetTokenValidityDuration** | **string** | The maximum length of time that a password reset token should be considered valid. | 

## Methods

### NewExtendedOperationHandlerListResponseResourcesInner

`func NewExtendedOperationHandlerListResponseResourcesInner(schemas []EnumgetConfigurationExtendedOperationHandlerSchemaUrn, id string, enabled bool, backendID string, passwordGenerator string, defaultOTPDeliveryMechanism []string, identityMapper string, extensionClass string, defaultPasswordGenerator string, defaultTokenDeliveryMechanism []string, passwordResetTokenValidityDuration string, ) *ExtendedOperationHandlerListResponseResourcesInner`

NewExtendedOperationHandlerListResponseResourcesInner instantiates a new ExtendedOperationHandlerListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedOperationHandlerListResponseResourcesInnerWithDefaults

`func NewExtendedOperationHandlerListResponseResourcesInnerWithDefaults() *ExtendedOperationHandlerListResponseResourcesInner`

NewExtendedOperationHandlerListResponseResourcesInnerWithDefaults instantiates a new ExtendedOperationHandlerListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetSchemas() []EnumgetConfigurationExtendedOperationHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetSchemasOk() (*[]EnumgetConfigurationExtendedOperationHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetSchemas(v []EnumgetConfigurationExtendedOperationHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSharedSecretAttributeType

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetSharedSecretAttributeType() string`

GetSharedSecretAttributeType returns the SharedSecretAttributeType field if non-nil, zero value otherwise.

### GetSharedSecretAttributeTypeOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetSharedSecretAttributeTypeOk() (*string, bool)`

GetSharedSecretAttributeTypeOk returns a tuple with the SharedSecretAttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecretAttributeType

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetSharedSecretAttributeType(v string)`

SetSharedSecretAttributeType sets SharedSecretAttributeType field to given value.

### HasSharedSecretAttributeType

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasSharedSecretAttributeType() bool`

HasSharedSecretAttributeType returns a boolean if a field has been set.

### GetTimeIntervalDuration

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetTimeIntervalDuration() string`

GetTimeIntervalDuration returns the TimeIntervalDuration field if non-nil, zero value otherwise.

### GetTimeIntervalDurationOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetTimeIntervalDurationOk() (*string, bool)`

GetTimeIntervalDurationOk returns a tuple with the TimeIntervalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeIntervalDuration

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetTimeIntervalDuration(v string)`

SetTimeIntervalDuration sets TimeIntervalDuration field to given value.

### HasTimeIntervalDuration

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasTimeIntervalDuration() bool`

HasTimeIntervalDuration returns a boolean if a field has been set.

### GetAdjacentIntervalsToCheck

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetAdjacentIntervalsToCheck() int64`

GetAdjacentIntervalsToCheck returns the AdjacentIntervalsToCheck field if non-nil, zero value otherwise.

### GetAdjacentIntervalsToCheckOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetAdjacentIntervalsToCheckOk() (*int64, bool)`

GetAdjacentIntervalsToCheckOk returns a tuple with the AdjacentIntervalsToCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacentIntervalsToCheck

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetAdjacentIntervalsToCheck(v int64)`

SetAdjacentIntervalsToCheck sets AdjacentIntervalsToCheck field to given value.

### HasAdjacentIntervalsToCheck

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasAdjacentIntervalsToCheck() bool`

HasAdjacentIntervalsToCheck returns a boolean if a field has been set.

### GetPreventTOTPReuse

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetPreventTOTPReuse() bool`

GetPreventTOTPReuse returns the PreventTOTPReuse field if non-nil, zero value otherwise.

### GetPreventTOTPReuseOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetPreventTOTPReuseOk() (*bool, bool)`

GetPreventTOTPReuseOk returns a tuple with the PreventTOTPReuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventTOTPReuse

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetPreventTOTPReuse(v bool)`

SetPreventTOTPReuse sets PreventTOTPReuse field to given value.

### HasPreventTOTPReuse

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasPreventTOTPReuse() bool`

HasPreventTOTPReuse returns a boolean if a field has been set.

### GetBackendID

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetBackendID() string`

GetBackendID returns the BackendID field if non-nil, zero value otherwise.

### GetBackendIDOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetBackendIDOk() (*string, bool)`

GetBackendIDOk returns a tuple with the BackendID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendID

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetBackendID(v string)`

SetBackendID sets BackendID field to given value.


### GetAllowRemotelyProvidedCertificates

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetAllowRemotelyProvidedCertificates() bool`

GetAllowRemotelyProvidedCertificates returns the AllowRemotelyProvidedCertificates field if non-nil, zero value otherwise.

### GetAllowRemotelyProvidedCertificatesOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetAllowRemotelyProvidedCertificatesOk() (*bool, bool)`

GetAllowRemotelyProvidedCertificatesOk returns a tuple with the AllowRemotelyProvidedCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemotelyProvidedCertificates

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetAllowRemotelyProvidedCertificates(v bool)`

SetAllowRemotelyProvidedCertificates sets AllowRemotelyProvidedCertificates field to given value.

### HasAllowRemotelyProvidedCertificates

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasAllowRemotelyProvidedCertificates() bool`

HasAllowRemotelyProvidedCertificates returns a boolean if a field has been set.

### GetAllowedOperation

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetAllowedOperation() []EnumextendedOperationHandlerAllowedOperationProp`

GetAllowedOperation returns the AllowedOperation field if non-nil, zero value otherwise.

### GetAllowedOperationOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetAllowedOperationOk() (*[]EnumextendedOperationHandlerAllowedOperationProp, bool)`

GetAllowedOperationOk returns a tuple with the AllowedOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperation

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetAllowedOperation(v []EnumextendedOperationHandlerAllowedOperationProp)`

SetAllowedOperation sets AllowedOperation field to given value.

### HasAllowedOperation

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasAllowedOperation() bool`

HasAllowedOperation returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetPasswordGenerator

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.


### GetDefaultOTPDeliveryMechanism

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultOTPDeliveryMechanism() []string`

GetDefaultOTPDeliveryMechanism returns the DefaultOTPDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultOTPDeliveryMechanismOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultOTPDeliveryMechanismOk() (*[]string, bool)`

GetDefaultOTPDeliveryMechanismOk returns a tuple with the DefaultOTPDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOTPDeliveryMechanism

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetDefaultOTPDeliveryMechanism(v []string)`

SetDefaultOTPDeliveryMechanism sets DefaultOTPDeliveryMechanism field to given value.


### GetDefaultSingleUseTokenValidityDuration

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultSingleUseTokenValidityDuration() string`

GetDefaultSingleUseTokenValidityDuration returns the DefaultSingleUseTokenValidityDuration field if non-nil, zero value otherwise.

### GetDefaultSingleUseTokenValidityDurationOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultSingleUseTokenValidityDurationOk() (*string, bool)`

GetDefaultSingleUseTokenValidityDurationOk returns a tuple with the DefaultSingleUseTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSingleUseTokenValidityDuration

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetDefaultSingleUseTokenValidityDuration(v string)`

SetDefaultSingleUseTokenValidityDuration sets DefaultSingleUseTokenValidityDuration field to given value.

### HasDefaultSingleUseTokenValidityDuration

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasDefaultSingleUseTokenValidityDuration() bool`

HasDefaultSingleUseTokenValidityDuration returns a boolean if a field has been set.

### GetValuesPerStreamResponse

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetValuesPerStreamResponse() int64`

GetValuesPerStreamResponse returns the ValuesPerStreamResponse field if non-nil, zero value otherwise.

### GetValuesPerStreamResponseOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetValuesPerStreamResponseOk() (*int64, bool)`

GetValuesPerStreamResponseOk returns a tuple with the ValuesPerStreamResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPerStreamResponse

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetValuesPerStreamResponse(v int64)`

SetValuesPerStreamResponse sets ValuesPerStreamResponse field to given value.

### HasValuesPerStreamResponse

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasValuesPerStreamResponse() bool`

HasValuesPerStreamResponse returns a boolean if a field has been set.

### GetNonAccessibleSubtreeUnauthorizedBindResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetNonAccessibleSubtreeUnauthorizedBindResultCode() int64`

GetNonAccessibleSubtreeUnauthorizedBindResultCode returns the NonAccessibleSubtreeUnauthorizedBindResultCode field if non-nil, zero value otherwise.

### GetNonAccessibleSubtreeUnauthorizedBindResultCodeOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetNonAccessibleSubtreeUnauthorizedBindResultCodeOk() (*int64, bool)`

GetNonAccessibleSubtreeUnauthorizedBindResultCodeOk returns a tuple with the NonAccessibleSubtreeUnauthorizedBindResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonAccessibleSubtreeUnauthorizedBindResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetNonAccessibleSubtreeUnauthorizedBindResultCode(v int64)`

SetNonAccessibleSubtreeUnauthorizedBindResultCode sets NonAccessibleSubtreeUnauthorizedBindResultCode field to given value.

### HasNonAccessibleSubtreeUnauthorizedBindResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasNonAccessibleSubtreeUnauthorizedBindResultCode() bool`

HasNonAccessibleSubtreeUnauthorizedBindResultCode returns a boolean if a field has been set.

### GetNonAccessibleSubtreeUnauthorizedReadResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetNonAccessibleSubtreeUnauthorizedReadResultCode() int64`

GetNonAccessibleSubtreeUnauthorizedReadResultCode returns the NonAccessibleSubtreeUnauthorizedReadResultCode field if non-nil, zero value otherwise.

### GetNonAccessibleSubtreeUnauthorizedReadResultCodeOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetNonAccessibleSubtreeUnauthorizedReadResultCodeOk() (*int64, bool)`

GetNonAccessibleSubtreeUnauthorizedReadResultCodeOk returns a tuple with the NonAccessibleSubtreeUnauthorizedReadResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonAccessibleSubtreeUnauthorizedReadResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetNonAccessibleSubtreeUnauthorizedReadResultCode(v int64)`

SetNonAccessibleSubtreeUnauthorizedReadResultCode sets NonAccessibleSubtreeUnauthorizedReadResultCode field to given value.

### HasNonAccessibleSubtreeUnauthorizedReadResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasNonAccessibleSubtreeUnauthorizedReadResultCode() bool`

HasNonAccessibleSubtreeUnauthorizedReadResultCode returns a boolean if a field has been set.

### GetNonAccessibleSubtreeUnauthorizedWriteResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetNonAccessibleSubtreeUnauthorizedWriteResultCode() int64`

GetNonAccessibleSubtreeUnauthorizedWriteResultCode returns the NonAccessibleSubtreeUnauthorizedWriteResultCode field if non-nil, zero value otherwise.

### GetNonAccessibleSubtreeUnauthorizedWriteResultCodeOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetNonAccessibleSubtreeUnauthorizedWriteResultCodeOk() (*int64, bool)`

GetNonAccessibleSubtreeUnauthorizedWriteResultCodeOk returns a tuple with the NonAccessibleSubtreeUnauthorizedWriteResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonAccessibleSubtreeUnauthorizedWriteResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetNonAccessibleSubtreeUnauthorizedWriteResultCode(v int64)`

SetNonAccessibleSubtreeUnauthorizedWriteResultCode sets NonAccessibleSubtreeUnauthorizedWriteResultCode field to given value.

### HasNonAccessibleSubtreeUnauthorizedWriteResultCode

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasNonAccessibleSubtreeUnauthorizedWriteResultCode() bool`

HasNonAccessibleSubtreeUnauthorizedWriteResultCode returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetExtensionClass

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetRouteToBackendSetBehavior

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetRouteToBackendSetBehavior() EnumextendedOperationHandlerRouteToBackendSetBehaviorProp`

GetRouteToBackendSetBehavior returns the RouteToBackendSetBehavior field if non-nil, zero value otherwise.

### GetRouteToBackendSetBehaviorOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetRouteToBackendSetBehaviorOk() (*EnumextendedOperationHandlerRouteToBackendSetBehaviorProp, bool)`

GetRouteToBackendSetBehaviorOk returns a tuple with the RouteToBackendSetBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteToBackendSetBehavior

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetRouteToBackendSetBehavior(v EnumextendedOperationHandlerRouteToBackendSetBehaviorProp)`

SetRouteToBackendSetBehavior sets RouteToBackendSetBehavior field to given value.

### HasRouteToBackendSetBehavior

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasRouteToBackendSetBehavior() bool`

HasRouteToBackendSetBehavior returns a boolean if a field has been set.

### GetDefaultPasswordPolicy

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultPasswordPolicy() string`

GetDefaultPasswordPolicy returns the DefaultPasswordPolicy field if non-nil, zero value otherwise.

### GetDefaultPasswordPolicyOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultPasswordPolicyOk() (*string, bool)`

GetDefaultPasswordPolicyOk returns a tuple with the DefaultPasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordPolicy

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetDefaultPasswordPolicy(v string)`

SetDefaultPasswordPolicy sets DefaultPasswordPolicy field to given value.

### HasDefaultPasswordPolicy

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasDefaultPasswordPolicy() bool`

HasDefaultPasswordPolicy returns a boolean if a field has been set.

### GetDefaultPasswordGenerator

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultPasswordGenerator() string`

GetDefaultPasswordGenerator returns the DefaultPasswordGenerator field if non-nil, zero value otherwise.

### GetDefaultPasswordGeneratorOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultPasswordGeneratorOk() (*string, bool)`

GetDefaultPasswordGeneratorOk returns a tuple with the DefaultPasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordGenerator

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetDefaultPasswordGenerator(v string)`

SetDefaultPasswordGenerator sets DefaultPasswordGenerator field to given value.


### GetMaximumPasswordsPerRequest

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetMaximumPasswordsPerRequest() int64`

GetMaximumPasswordsPerRequest returns the MaximumPasswordsPerRequest field if non-nil, zero value otherwise.

### GetMaximumPasswordsPerRequestOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetMaximumPasswordsPerRequestOk() (*int64, bool)`

GetMaximumPasswordsPerRequestOk returns a tuple with the MaximumPasswordsPerRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPasswordsPerRequest

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetMaximumPasswordsPerRequest(v int64)`

SetMaximumPasswordsPerRequest sets MaximumPasswordsPerRequest field to given value.

### HasMaximumPasswordsPerRequest

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasMaximumPasswordsPerRequest() bool`

HasMaximumPasswordsPerRequest returns a boolean if a field has been set.

### GetMaximumValidationAttemptsPerPassword

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetMaximumValidationAttemptsPerPassword() int64`

GetMaximumValidationAttemptsPerPassword returns the MaximumValidationAttemptsPerPassword field if non-nil, zero value otherwise.

### GetMaximumValidationAttemptsPerPasswordOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetMaximumValidationAttemptsPerPasswordOk() (*int64, bool)`

GetMaximumValidationAttemptsPerPasswordOk returns a tuple with the MaximumValidationAttemptsPerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValidationAttemptsPerPassword

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetMaximumValidationAttemptsPerPassword(v int64)`

SetMaximumValidationAttemptsPerPassword sets MaximumValidationAttemptsPerPassword field to given value.

### HasMaximumValidationAttemptsPerPassword

`func (o *ExtendedOperationHandlerListResponseResourcesInner) HasMaximumValidationAttemptsPerPassword() bool`

HasMaximumValidationAttemptsPerPassword returns a boolean if a field has been set.

### GetDefaultTokenDeliveryMechanism

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultTokenDeliveryMechanism() []string`

GetDefaultTokenDeliveryMechanism returns the DefaultTokenDeliveryMechanism field if non-nil, zero value otherwise.

### GetDefaultTokenDeliveryMechanismOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetDefaultTokenDeliveryMechanismOk() (*[]string, bool)`

GetDefaultTokenDeliveryMechanismOk returns a tuple with the DefaultTokenDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTokenDeliveryMechanism

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetDefaultTokenDeliveryMechanism(v []string)`

SetDefaultTokenDeliveryMechanism sets DefaultTokenDeliveryMechanism field to given value.


### GetPasswordResetTokenValidityDuration

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetPasswordResetTokenValidityDuration() string`

GetPasswordResetTokenValidityDuration returns the PasswordResetTokenValidityDuration field if non-nil, zero value otherwise.

### GetPasswordResetTokenValidityDurationOk

`func (o *ExtendedOperationHandlerListResponseResourcesInner) GetPasswordResetTokenValidityDurationOk() (*string, bool)`

GetPasswordResetTokenValidityDurationOk returns a tuple with the PasswordResetTokenValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetTokenValidityDuration

`func (o *ExtendedOperationHandlerListResponseResourcesInner) SetPasswordResetTokenValidityDuration(v string)`

SetPasswordResetTokenValidityDuration sets PasswordResetTokenValidityDuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


