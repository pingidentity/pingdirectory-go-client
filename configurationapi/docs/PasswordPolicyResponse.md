# PasswordPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Password Policy | 
**Schemas** | Pointer to [**[]EnumpasswordPolicySchemaUrn**](EnumpasswordPolicySchemaUrn.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Password Policy | [optional] 
**RequireSecureAuthentication** | Pointer to **bool** | Indicates whether users with the associated password policy are required to authenticate in a secure manner. | [optional] 
**RequireSecurePasswordChanges** | Pointer to **bool** | Indicates whether users with the associated password policy are required to change their password in a secure manner that does not expose the credentials. | [optional] 
**AccountStatusNotificationHandler** | Pointer to **[]string** | Specifies the names of the account status notification handlers that are used with the associated password storage scheme. | [optional] 
**StateUpdateFailurePolicy** | Pointer to [**EnumpasswordPolicyStateUpdateFailurePolicyProp**](EnumpasswordPolicyStateUpdateFailurePolicyProp.md) |  | [optional] 
**EnableDebug** | Pointer to **bool** | Indicates whether to enable debugging for the password policy state. | [optional] 
**PasswordAttribute** | **string** | Specifies the attribute type used to hold user passwords. | 
**DefaultPasswordStorageScheme** | **[]string** | Specifies the names of the password storage schemes that are used to encode clear-text passwords for this password policy. | 
**DeprecatedPasswordStorageScheme** | Pointer to **[]string** | Specifies the names of the password storage schemes that are considered deprecated for this password policy. | [optional] 
**AllowMultiplePasswordValues** | Pointer to **bool** | Indicates whether user entries can have multiple distinct values for the password attribute. | [optional] 
**AllowPreEncodedPasswords** | Pointer to **bool** | Indicates whether users can change their passwords by providing a pre-encoded value. | [optional] 
**PasswordValidator** | Pointer to **[]string** | Specifies the names of the password validators that are used with the associated password storage scheme. | [optional] 
**BindPasswordValidator** | Pointer to **[]string** | Specifies the names of the password validators that should be invoked for bind operations. | [optional] 
**MinimumBindPasswordValidationFrequency** | Pointer to **string** | Indicates how frequently password validation should be performed during bind operations for each user to whom this password policy is assigned. | [optional] 
**BindPasswordValidationFailureAction** | Pointer to [**EnumpasswordPolicyBindPasswordValidationFailureActionProp**](EnumpasswordPolicyBindPasswordValidationFailureActionProp.md) |  | [optional] 
**PasswordGenerator** | Pointer to **string** | Specifies the name of the password generator that is used with the associated password policy. | [optional] 
**PasswordHistoryCount** | Pointer to **int64** | Specifies the maximum number of former passwords to maintain in the password history. | [optional] 
**PasswordHistoryDuration** | Pointer to **string** | Specifies the maximum length of time that passwords remain in the password history. | [optional] 
**MinPasswordAge** | Pointer to **string** | Specifies the minimum length of time after a password change before the user is allowed to change the password again. | [optional] 
**MaxPasswordAge** | Pointer to **string** | Specifies the maximum length of time that a user can continue using the same password before it must be changed (that is, the password expiration interval). | [optional] 
**PasswordExpirationWarningInterval** | Pointer to **string** | Specifies the maximum length of time before a user&#39;s password actually expires that the server begins to include warning notifications in bind responses for that user. | [optional] 
**ExpirePasswordsWithoutWarning** | Pointer to **bool** | Indicates whether the Directory Server allows a user&#39;s password to expire even if that user has never seen an expiration warning notification. | [optional] 
**ReturnPasswordExpirationControls** | Pointer to [**EnumpasswordPolicyReturnPasswordExpirationControlsProp**](EnumpasswordPolicyReturnPasswordExpirationControlsProp.md) |  | [optional] 
**AllowExpiredPasswordChanges** | Pointer to **bool** | Indicates whether a user whose password is expired is still allowed to change that password using the password modify extended operation. | [optional] 
**GraceLoginCount** | Pointer to **int64** | Specifies the number of grace logins that a user is allowed after the account has expired to allow that user to choose a new password. | [optional] 
**RequireChangeByTime** | Pointer to **string** | Specifies the time by which all users with the associated password policy must change their passwords. | [optional] 
**LockoutFailureCount** | Pointer to **int64** | Specifies the maximum number of authentication failures that a user is allowed before the account is locked out. | [optional] 
**LockoutDuration** | Pointer to **string** | Specifies the length of time that an account is locked after too many authentication failures. | [optional] 
**LockoutFailureExpirationInterval** | Pointer to **string** | Specifies the length of time before an authentication failure is no longer counted against a user for the purposes of account lockout. | [optional] 
**IgnoreDuplicatePasswordFailures** | Pointer to **bool** | Indicates whether to ignore subsequent authentication failures using the same password as an earlier failed authentication attempt (within the time frame defined by the lockout failure expiration interval). If this option is \&quot;true\&quot;, then multiple failed attempts using the same password will be considered only a single failure. If this option is \&quot;false\&quot;, then any failure will be tracked regardless of whether it used the same password as an earlier attempt. | [optional] 
**FailureLockoutAction** | Pointer to **string** | The action that the server should take for authentication attempts that target a user with more than the configured number of outstanding authentication failures. | [optional] 
**IdleLockoutInterval** | Pointer to **string** | Specifies the maximum length of time that an account may remain idle (that is, the associated user does not authenticate to the server) before that user is locked out. | [optional] 
**AllowUserPasswordChanges** | Pointer to **bool** | Indicates whether users can change their own passwords. | [optional] 
**PasswordChangeRequiresCurrentPassword** | Pointer to **bool** | Indicates whether user password changes must use the password modify extended operation and must include the user&#39;s current password before the change is allowed. | [optional] 
**PasswordRetirementBehavior** | Pointer to [**[]EnumpasswordPolicyPasswordRetirementBehaviorProp**](EnumpasswordPolicyPasswordRetirementBehaviorProp.md) |  | [optional] 
**MaxRetiredPasswordAge** | Pointer to **string** | Specifies the maximum length of time that a retired password should be considered valid and may be used to authenticate to the server. | [optional] 
**AllowedPasswordResetTokenUseCondition** | Pointer to [**[]EnumpasswordPolicyAllowedPasswordResetTokenUseConditionProp**](EnumpasswordPolicyAllowedPasswordResetTokenUseConditionProp.md) |  | [optional] 
**ForceChangeOnAdd** | Pointer to **bool** | Indicates whether users are forced to change their passwords upon first authenticating to the Directory Server after their account has been created. | [optional] 
**ForceChangeOnReset** | Pointer to **bool** | Indicates whether users are forced to change their passwords if they are reset by an administrator. If a user&#39;s password is changed by any other user, that is considered an administrative password reset. | [optional] 
**MaxPasswordResetAge** | Pointer to **string** | Specifies the maximum length of time that users have to change passwords after they have been reset by an administrator before they become locked. | [optional] 
**SkipValidationForAdministrators** | Pointer to **bool** | Indicates whether passwords set by administrators are allowed to bypass the password validation process that is required for user password changes. | [optional] 
**MaximumRecentLoginHistorySuccessfulAuthenticationCount** | Pointer to **int64** | The maximum number of successful authentication attempts to include in the recent login history for each account. | [optional] 
**MaximumRecentLoginHistorySuccessfulAuthenticationDuration** | Pointer to **string** | The maximum age of successful authentication attempts to include in the recent login history for each account. | [optional] 
**MaximumRecentLoginHistoryFailedAuthenticationCount** | Pointer to **int64** | The maximum number of failed authentication attempts to include in the recent login history for each account. | [optional] 
**MaximumRecentLoginHistoryFailedAuthenticationDuration** | Pointer to **string** | The maximum age of failed authentication attempts to include in the recent login history for each account. | [optional] 
**RecentLoginHistorySimilarAttemptBehavior** | Pointer to [**EnumpasswordPolicyRecentLoginHistorySimilarAttemptBehaviorProp**](EnumpasswordPolicyRecentLoginHistorySimilarAttemptBehaviorProp.md) |  | [optional] 
**LastLoginIPAddressAttribute** | Pointer to **string** | Specifies the name or OID of the attribute type that is used to hold the IP address of the client from which the user last authenticated. | [optional] 
**LastLoginTimeAttribute** | Pointer to **string** | Specifies the name or OID of the attribute type that is used to hold the last login time for users with the associated password policy. | [optional] 
**LastLoginTimeFormat** | Pointer to **string** | Specifies the format string that is used to generate the last login time value for users with the associated password policy. Last login time values will be written using the UTC (also known as GMT, or Greenwich Mean Time) time zone. | [optional] 
**PreviousLastLoginTimeFormat** | Pointer to **[]string** | Specifies the format string(s) that might have been used with the last login time at any point in the past for users associated with the password policy. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewPasswordPolicyResponse

`func NewPasswordPolicyResponse(id string, passwordAttribute string, defaultPasswordStorageScheme []string, ) *PasswordPolicyResponse`

NewPasswordPolicyResponse instantiates a new PasswordPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyResponseWithDefaults

`func NewPasswordPolicyResponseWithDefaults() *PasswordPolicyResponse`

NewPasswordPolicyResponseWithDefaults instantiates a new PasswordPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PasswordPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PasswordPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PasswordPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *PasswordPolicyResponse) GetSchemas() []EnumpasswordPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PasswordPolicyResponse) GetSchemasOk() (*[]EnumpasswordPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PasswordPolicyResponse) SetSchemas(v []EnumpasswordPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *PasswordPolicyResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *PasswordPolicyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PasswordPolicyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PasswordPolicyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PasswordPolicyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequireSecureAuthentication

`func (o *PasswordPolicyResponse) GetRequireSecureAuthentication() bool`

GetRequireSecureAuthentication returns the RequireSecureAuthentication field if non-nil, zero value otherwise.

### GetRequireSecureAuthenticationOk

`func (o *PasswordPolicyResponse) GetRequireSecureAuthenticationOk() (*bool, bool)`

GetRequireSecureAuthenticationOk returns a tuple with the RequireSecureAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSecureAuthentication

`func (o *PasswordPolicyResponse) SetRequireSecureAuthentication(v bool)`

SetRequireSecureAuthentication sets RequireSecureAuthentication field to given value.

### HasRequireSecureAuthentication

`func (o *PasswordPolicyResponse) HasRequireSecureAuthentication() bool`

HasRequireSecureAuthentication returns a boolean if a field has been set.

### GetRequireSecurePasswordChanges

`func (o *PasswordPolicyResponse) GetRequireSecurePasswordChanges() bool`

GetRequireSecurePasswordChanges returns the RequireSecurePasswordChanges field if non-nil, zero value otherwise.

### GetRequireSecurePasswordChangesOk

`func (o *PasswordPolicyResponse) GetRequireSecurePasswordChangesOk() (*bool, bool)`

GetRequireSecurePasswordChangesOk returns a tuple with the RequireSecurePasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSecurePasswordChanges

`func (o *PasswordPolicyResponse) SetRequireSecurePasswordChanges(v bool)`

SetRequireSecurePasswordChanges sets RequireSecurePasswordChanges field to given value.

### HasRequireSecurePasswordChanges

`func (o *PasswordPolicyResponse) HasRequireSecurePasswordChanges() bool`

HasRequireSecurePasswordChanges returns a boolean if a field has been set.

### GetAccountStatusNotificationHandler

`func (o *PasswordPolicyResponse) GetAccountStatusNotificationHandler() []string`

GetAccountStatusNotificationHandler returns the AccountStatusNotificationHandler field if non-nil, zero value otherwise.

### GetAccountStatusNotificationHandlerOk

`func (o *PasswordPolicyResponse) GetAccountStatusNotificationHandlerOk() (*[]string, bool)`

GetAccountStatusNotificationHandlerOk returns a tuple with the AccountStatusNotificationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatusNotificationHandler

`func (o *PasswordPolicyResponse) SetAccountStatusNotificationHandler(v []string)`

SetAccountStatusNotificationHandler sets AccountStatusNotificationHandler field to given value.

### HasAccountStatusNotificationHandler

`func (o *PasswordPolicyResponse) HasAccountStatusNotificationHandler() bool`

HasAccountStatusNotificationHandler returns a boolean if a field has been set.

### GetStateUpdateFailurePolicy

`func (o *PasswordPolicyResponse) GetStateUpdateFailurePolicy() EnumpasswordPolicyStateUpdateFailurePolicyProp`

GetStateUpdateFailurePolicy returns the StateUpdateFailurePolicy field if non-nil, zero value otherwise.

### GetStateUpdateFailurePolicyOk

`func (o *PasswordPolicyResponse) GetStateUpdateFailurePolicyOk() (*EnumpasswordPolicyStateUpdateFailurePolicyProp, bool)`

GetStateUpdateFailurePolicyOk returns a tuple with the StateUpdateFailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateUpdateFailurePolicy

`func (o *PasswordPolicyResponse) SetStateUpdateFailurePolicy(v EnumpasswordPolicyStateUpdateFailurePolicyProp)`

SetStateUpdateFailurePolicy sets StateUpdateFailurePolicy field to given value.

### HasStateUpdateFailurePolicy

`func (o *PasswordPolicyResponse) HasStateUpdateFailurePolicy() bool`

HasStateUpdateFailurePolicy returns a boolean if a field has been set.

### GetEnableDebug

`func (o *PasswordPolicyResponse) GetEnableDebug() bool`

GetEnableDebug returns the EnableDebug field if non-nil, zero value otherwise.

### GetEnableDebugOk

`func (o *PasswordPolicyResponse) GetEnableDebugOk() (*bool, bool)`

GetEnableDebugOk returns a tuple with the EnableDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebug

`func (o *PasswordPolicyResponse) SetEnableDebug(v bool)`

SetEnableDebug sets EnableDebug field to given value.

### HasEnableDebug

`func (o *PasswordPolicyResponse) HasEnableDebug() bool`

HasEnableDebug returns a boolean if a field has been set.

### GetPasswordAttribute

`func (o *PasswordPolicyResponse) GetPasswordAttribute() string`

GetPasswordAttribute returns the PasswordAttribute field if non-nil, zero value otherwise.

### GetPasswordAttributeOk

`func (o *PasswordPolicyResponse) GetPasswordAttributeOk() (*string, bool)`

GetPasswordAttributeOk returns a tuple with the PasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAttribute

`func (o *PasswordPolicyResponse) SetPasswordAttribute(v string)`

SetPasswordAttribute sets PasswordAttribute field to given value.


### GetDefaultPasswordStorageScheme

`func (o *PasswordPolicyResponse) GetDefaultPasswordStorageScheme() []string`

GetDefaultPasswordStorageScheme returns the DefaultPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDefaultPasswordStorageSchemeOk

`func (o *PasswordPolicyResponse) GetDefaultPasswordStorageSchemeOk() (*[]string, bool)`

GetDefaultPasswordStorageSchemeOk returns a tuple with the DefaultPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordStorageScheme

`func (o *PasswordPolicyResponse) SetDefaultPasswordStorageScheme(v []string)`

SetDefaultPasswordStorageScheme sets DefaultPasswordStorageScheme field to given value.


### GetDeprecatedPasswordStorageScheme

`func (o *PasswordPolicyResponse) GetDeprecatedPasswordStorageScheme() []string`

GetDeprecatedPasswordStorageScheme returns the DeprecatedPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDeprecatedPasswordStorageSchemeOk

`func (o *PasswordPolicyResponse) GetDeprecatedPasswordStorageSchemeOk() (*[]string, bool)`

GetDeprecatedPasswordStorageSchemeOk returns a tuple with the DeprecatedPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedPasswordStorageScheme

`func (o *PasswordPolicyResponse) SetDeprecatedPasswordStorageScheme(v []string)`

SetDeprecatedPasswordStorageScheme sets DeprecatedPasswordStorageScheme field to given value.

### HasDeprecatedPasswordStorageScheme

`func (o *PasswordPolicyResponse) HasDeprecatedPasswordStorageScheme() bool`

HasDeprecatedPasswordStorageScheme returns a boolean if a field has been set.

### GetAllowMultiplePasswordValues

`func (o *PasswordPolicyResponse) GetAllowMultiplePasswordValues() bool`

GetAllowMultiplePasswordValues returns the AllowMultiplePasswordValues field if non-nil, zero value otherwise.

### GetAllowMultiplePasswordValuesOk

`func (o *PasswordPolicyResponse) GetAllowMultiplePasswordValuesOk() (*bool, bool)`

GetAllowMultiplePasswordValuesOk returns a tuple with the AllowMultiplePasswordValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiplePasswordValues

`func (o *PasswordPolicyResponse) SetAllowMultiplePasswordValues(v bool)`

SetAllowMultiplePasswordValues sets AllowMultiplePasswordValues field to given value.

### HasAllowMultiplePasswordValues

`func (o *PasswordPolicyResponse) HasAllowMultiplePasswordValues() bool`

HasAllowMultiplePasswordValues returns a boolean if a field has been set.

### GetAllowPreEncodedPasswords

`func (o *PasswordPolicyResponse) GetAllowPreEncodedPasswords() bool`

GetAllowPreEncodedPasswords returns the AllowPreEncodedPasswords field if non-nil, zero value otherwise.

### GetAllowPreEncodedPasswordsOk

`func (o *PasswordPolicyResponse) GetAllowPreEncodedPasswordsOk() (*bool, bool)`

GetAllowPreEncodedPasswordsOk returns a tuple with the AllowPreEncodedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPreEncodedPasswords

`func (o *PasswordPolicyResponse) SetAllowPreEncodedPasswords(v bool)`

SetAllowPreEncodedPasswords sets AllowPreEncodedPasswords field to given value.

### HasAllowPreEncodedPasswords

`func (o *PasswordPolicyResponse) HasAllowPreEncodedPasswords() bool`

HasAllowPreEncodedPasswords returns a boolean if a field has been set.

### GetPasswordValidator

`func (o *PasswordPolicyResponse) GetPasswordValidator() []string`

GetPasswordValidator returns the PasswordValidator field if non-nil, zero value otherwise.

### GetPasswordValidatorOk

`func (o *PasswordPolicyResponse) GetPasswordValidatorOk() (*[]string, bool)`

GetPasswordValidatorOk returns a tuple with the PasswordValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordValidator

`func (o *PasswordPolicyResponse) SetPasswordValidator(v []string)`

SetPasswordValidator sets PasswordValidator field to given value.

### HasPasswordValidator

`func (o *PasswordPolicyResponse) HasPasswordValidator() bool`

HasPasswordValidator returns a boolean if a field has been set.

### GetBindPasswordValidator

`func (o *PasswordPolicyResponse) GetBindPasswordValidator() []string`

GetBindPasswordValidator returns the BindPasswordValidator field if non-nil, zero value otherwise.

### GetBindPasswordValidatorOk

`func (o *PasswordPolicyResponse) GetBindPasswordValidatorOk() (*[]string, bool)`

GetBindPasswordValidatorOk returns a tuple with the BindPasswordValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPasswordValidator

`func (o *PasswordPolicyResponse) SetBindPasswordValidator(v []string)`

SetBindPasswordValidator sets BindPasswordValidator field to given value.

### HasBindPasswordValidator

`func (o *PasswordPolicyResponse) HasBindPasswordValidator() bool`

HasBindPasswordValidator returns a boolean if a field has been set.

### GetMinimumBindPasswordValidationFrequency

`func (o *PasswordPolicyResponse) GetMinimumBindPasswordValidationFrequency() string`

GetMinimumBindPasswordValidationFrequency returns the MinimumBindPasswordValidationFrequency field if non-nil, zero value otherwise.

### GetMinimumBindPasswordValidationFrequencyOk

`func (o *PasswordPolicyResponse) GetMinimumBindPasswordValidationFrequencyOk() (*string, bool)`

GetMinimumBindPasswordValidationFrequencyOk returns a tuple with the MinimumBindPasswordValidationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBindPasswordValidationFrequency

`func (o *PasswordPolicyResponse) SetMinimumBindPasswordValidationFrequency(v string)`

SetMinimumBindPasswordValidationFrequency sets MinimumBindPasswordValidationFrequency field to given value.

### HasMinimumBindPasswordValidationFrequency

`func (o *PasswordPolicyResponse) HasMinimumBindPasswordValidationFrequency() bool`

HasMinimumBindPasswordValidationFrequency returns a boolean if a field has been set.

### GetBindPasswordValidationFailureAction

`func (o *PasswordPolicyResponse) GetBindPasswordValidationFailureAction() EnumpasswordPolicyBindPasswordValidationFailureActionProp`

GetBindPasswordValidationFailureAction returns the BindPasswordValidationFailureAction field if non-nil, zero value otherwise.

### GetBindPasswordValidationFailureActionOk

`func (o *PasswordPolicyResponse) GetBindPasswordValidationFailureActionOk() (*EnumpasswordPolicyBindPasswordValidationFailureActionProp, bool)`

GetBindPasswordValidationFailureActionOk returns a tuple with the BindPasswordValidationFailureAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPasswordValidationFailureAction

`func (o *PasswordPolicyResponse) SetBindPasswordValidationFailureAction(v EnumpasswordPolicyBindPasswordValidationFailureActionProp)`

SetBindPasswordValidationFailureAction sets BindPasswordValidationFailureAction field to given value.

### HasBindPasswordValidationFailureAction

`func (o *PasswordPolicyResponse) HasBindPasswordValidationFailureAction() bool`

HasBindPasswordValidationFailureAction returns a boolean if a field has been set.

### GetPasswordGenerator

`func (o *PasswordPolicyResponse) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *PasswordPolicyResponse) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *PasswordPolicyResponse) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.

### HasPasswordGenerator

`func (o *PasswordPolicyResponse) HasPasswordGenerator() bool`

HasPasswordGenerator returns a boolean if a field has been set.

### GetPasswordHistoryCount

`func (o *PasswordPolicyResponse) GetPasswordHistoryCount() int64`

GetPasswordHistoryCount returns the PasswordHistoryCount field if non-nil, zero value otherwise.

### GetPasswordHistoryCountOk

`func (o *PasswordPolicyResponse) GetPasswordHistoryCountOk() (*int64, bool)`

GetPasswordHistoryCountOk returns a tuple with the PasswordHistoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistoryCount

`func (o *PasswordPolicyResponse) SetPasswordHistoryCount(v int64)`

SetPasswordHistoryCount sets PasswordHistoryCount field to given value.

### HasPasswordHistoryCount

`func (o *PasswordPolicyResponse) HasPasswordHistoryCount() bool`

HasPasswordHistoryCount returns a boolean if a field has been set.

### GetPasswordHistoryDuration

`func (o *PasswordPolicyResponse) GetPasswordHistoryDuration() string`

GetPasswordHistoryDuration returns the PasswordHistoryDuration field if non-nil, zero value otherwise.

### GetPasswordHistoryDurationOk

`func (o *PasswordPolicyResponse) GetPasswordHistoryDurationOk() (*string, bool)`

GetPasswordHistoryDurationOk returns a tuple with the PasswordHistoryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistoryDuration

`func (o *PasswordPolicyResponse) SetPasswordHistoryDuration(v string)`

SetPasswordHistoryDuration sets PasswordHistoryDuration field to given value.

### HasPasswordHistoryDuration

`func (o *PasswordPolicyResponse) HasPasswordHistoryDuration() bool`

HasPasswordHistoryDuration returns a boolean if a field has been set.

### GetMinPasswordAge

`func (o *PasswordPolicyResponse) GetMinPasswordAge() string`

GetMinPasswordAge returns the MinPasswordAge field if non-nil, zero value otherwise.

### GetMinPasswordAgeOk

`func (o *PasswordPolicyResponse) GetMinPasswordAgeOk() (*string, bool)`

GetMinPasswordAgeOk returns a tuple with the MinPasswordAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPasswordAge

`func (o *PasswordPolicyResponse) SetMinPasswordAge(v string)`

SetMinPasswordAge sets MinPasswordAge field to given value.

### HasMinPasswordAge

`func (o *PasswordPolicyResponse) HasMinPasswordAge() bool`

HasMinPasswordAge returns a boolean if a field has been set.

### GetMaxPasswordAge

`func (o *PasswordPolicyResponse) GetMaxPasswordAge() string`

GetMaxPasswordAge returns the MaxPasswordAge field if non-nil, zero value otherwise.

### GetMaxPasswordAgeOk

`func (o *PasswordPolicyResponse) GetMaxPasswordAgeOk() (*string, bool)`

GetMaxPasswordAgeOk returns a tuple with the MaxPasswordAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordAge

`func (o *PasswordPolicyResponse) SetMaxPasswordAge(v string)`

SetMaxPasswordAge sets MaxPasswordAge field to given value.

### HasMaxPasswordAge

`func (o *PasswordPolicyResponse) HasMaxPasswordAge() bool`

HasMaxPasswordAge returns a boolean if a field has been set.

### GetPasswordExpirationWarningInterval

`func (o *PasswordPolicyResponse) GetPasswordExpirationWarningInterval() string`

GetPasswordExpirationWarningInterval returns the PasswordExpirationWarningInterval field if non-nil, zero value otherwise.

### GetPasswordExpirationWarningIntervalOk

`func (o *PasswordPolicyResponse) GetPasswordExpirationWarningIntervalOk() (*string, bool)`

GetPasswordExpirationWarningIntervalOk returns a tuple with the PasswordExpirationWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationWarningInterval

`func (o *PasswordPolicyResponse) SetPasswordExpirationWarningInterval(v string)`

SetPasswordExpirationWarningInterval sets PasswordExpirationWarningInterval field to given value.

### HasPasswordExpirationWarningInterval

`func (o *PasswordPolicyResponse) HasPasswordExpirationWarningInterval() bool`

HasPasswordExpirationWarningInterval returns a boolean if a field has been set.

### GetExpirePasswordsWithoutWarning

`func (o *PasswordPolicyResponse) GetExpirePasswordsWithoutWarning() bool`

GetExpirePasswordsWithoutWarning returns the ExpirePasswordsWithoutWarning field if non-nil, zero value otherwise.

### GetExpirePasswordsWithoutWarningOk

`func (o *PasswordPolicyResponse) GetExpirePasswordsWithoutWarningOk() (*bool, bool)`

GetExpirePasswordsWithoutWarningOk returns a tuple with the ExpirePasswordsWithoutWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePasswordsWithoutWarning

`func (o *PasswordPolicyResponse) SetExpirePasswordsWithoutWarning(v bool)`

SetExpirePasswordsWithoutWarning sets ExpirePasswordsWithoutWarning field to given value.

### HasExpirePasswordsWithoutWarning

`func (o *PasswordPolicyResponse) HasExpirePasswordsWithoutWarning() bool`

HasExpirePasswordsWithoutWarning returns a boolean if a field has been set.

### GetReturnPasswordExpirationControls

`func (o *PasswordPolicyResponse) GetReturnPasswordExpirationControls() EnumpasswordPolicyReturnPasswordExpirationControlsProp`

GetReturnPasswordExpirationControls returns the ReturnPasswordExpirationControls field if non-nil, zero value otherwise.

### GetReturnPasswordExpirationControlsOk

`func (o *PasswordPolicyResponse) GetReturnPasswordExpirationControlsOk() (*EnumpasswordPolicyReturnPasswordExpirationControlsProp, bool)`

GetReturnPasswordExpirationControlsOk returns a tuple with the ReturnPasswordExpirationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPasswordExpirationControls

`func (o *PasswordPolicyResponse) SetReturnPasswordExpirationControls(v EnumpasswordPolicyReturnPasswordExpirationControlsProp)`

SetReturnPasswordExpirationControls sets ReturnPasswordExpirationControls field to given value.

### HasReturnPasswordExpirationControls

`func (o *PasswordPolicyResponse) HasReturnPasswordExpirationControls() bool`

HasReturnPasswordExpirationControls returns a boolean if a field has been set.

### GetAllowExpiredPasswordChanges

`func (o *PasswordPolicyResponse) GetAllowExpiredPasswordChanges() bool`

GetAllowExpiredPasswordChanges returns the AllowExpiredPasswordChanges field if non-nil, zero value otherwise.

### GetAllowExpiredPasswordChangesOk

`func (o *PasswordPolicyResponse) GetAllowExpiredPasswordChangesOk() (*bool, bool)`

GetAllowExpiredPasswordChangesOk returns a tuple with the AllowExpiredPasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExpiredPasswordChanges

`func (o *PasswordPolicyResponse) SetAllowExpiredPasswordChanges(v bool)`

SetAllowExpiredPasswordChanges sets AllowExpiredPasswordChanges field to given value.

### HasAllowExpiredPasswordChanges

`func (o *PasswordPolicyResponse) HasAllowExpiredPasswordChanges() bool`

HasAllowExpiredPasswordChanges returns a boolean if a field has been set.

### GetGraceLoginCount

`func (o *PasswordPolicyResponse) GetGraceLoginCount() int64`

GetGraceLoginCount returns the GraceLoginCount field if non-nil, zero value otherwise.

### GetGraceLoginCountOk

`func (o *PasswordPolicyResponse) GetGraceLoginCountOk() (*int64, bool)`

GetGraceLoginCountOk returns a tuple with the GraceLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceLoginCount

`func (o *PasswordPolicyResponse) SetGraceLoginCount(v int64)`

SetGraceLoginCount sets GraceLoginCount field to given value.

### HasGraceLoginCount

`func (o *PasswordPolicyResponse) HasGraceLoginCount() bool`

HasGraceLoginCount returns a boolean if a field has been set.

### GetRequireChangeByTime

`func (o *PasswordPolicyResponse) GetRequireChangeByTime() string`

GetRequireChangeByTime returns the RequireChangeByTime field if non-nil, zero value otherwise.

### GetRequireChangeByTimeOk

`func (o *PasswordPolicyResponse) GetRequireChangeByTimeOk() (*string, bool)`

GetRequireChangeByTimeOk returns a tuple with the RequireChangeByTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireChangeByTime

`func (o *PasswordPolicyResponse) SetRequireChangeByTime(v string)`

SetRequireChangeByTime sets RequireChangeByTime field to given value.

### HasRequireChangeByTime

`func (o *PasswordPolicyResponse) HasRequireChangeByTime() bool`

HasRequireChangeByTime returns a boolean if a field has been set.

### GetLockoutFailureCount

`func (o *PasswordPolicyResponse) GetLockoutFailureCount() int64`

GetLockoutFailureCount returns the LockoutFailureCount field if non-nil, zero value otherwise.

### GetLockoutFailureCountOk

`func (o *PasswordPolicyResponse) GetLockoutFailureCountOk() (*int64, bool)`

GetLockoutFailureCountOk returns a tuple with the LockoutFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutFailureCount

`func (o *PasswordPolicyResponse) SetLockoutFailureCount(v int64)`

SetLockoutFailureCount sets LockoutFailureCount field to given value.

### HasLockoutFailureCount

`func (o *PasswordPolicyResponse) HasLockoutFailureCount() bool`

HasLockoutFailureCount returns a boolean if a field has been set.

### GetLockoutDuration

`func (o *PasswordPolicyResponse) GetLockoutDuration() string`

GetLockoutDuration returns the LockoutDuration field if non-nil, zero value otherwise.

### GetLockoutDurationOk

`func (o *PasswordPolicyResponse) GetLockoutDurationOk() (*string, bool)`

GetLockoutDurationOk returns a tuple with the LockoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutDuration

`func (o *PasswordPolicyResponse) SetLockoutDuration(v string)`

SetLockoutDuration sets LockoutDuration field to given value.

### HasLockoutDuration

`func (o *PasswordPolicyResponse) HasLockoutDuration() bool`

HasLockoutDuration returns a boolean if a field has been set.

### GetLockoutFailureExpirationInterval

`func (o *PasswordPolicyResponse) GetLockoutFailureExpirationInterval() string`

GetLockoutFailureExpirationInterval returns the LockoutFailureExpirationInterval field if non-nil, zero value otherwise.

### GetLockoutFailureExpirationIntervalOk

`func (o *PasswordPolicyResponse) GetLockoutFailureExpirationIntervalOk() (*string, bool)`

GetLockoutFailureExpirationIntervalOk returns a tuple with the LockoutFailureExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutFailureExpirationInterval

`func (o *PasswordPolicyResponse) SetLockoutFailureExpirationInterval(v string)`

SetLockoutFailureExpirationInterval sets LockoutFailureExpirationInterval field to given value.

### HasLockoutFailureExpirationInterval

`func (o *PasswordPolicyResponse) HasLockoutFailureExpirationInterval() bool`

HasLockoutFailureExpirationInterval returns a boolean if a field has been set.

### GetIgnoreDuplicatePasswordFailures

`func (o *PasswordPolicyResponse) GetIgnoreDuplicatePasswordFailures() bool`

GetIgnoreDuplicatePasswordFailures returns the IgnoreDuplicatePasswordFailures field if non-nil, zero value otherwise.

### GetIgnoreDuplicatePasswordFailuresOk

`func (o *PasswordPolicyResponse) GetIgnoreDuplicatePasswordFailuresOk() (*bool, bool)`

GetIgnoreDuplicatePasswordFailuresOk returns a tuple with the IgnoreDuplicatePasswordFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDuplicatePasswordFailures

`func (o *PasswordPolicyResponse) SetIgnoreDuplicatePasswordFailures(v bool)`

SetIgnoreDuplicatePasswordFailures sets IgnoreDuplicatePasswordFailures field to given value.

### HasIgnoreDuplicatePasswordFailures

`func (o *PasswordPolicyResponse) HasIgnoreDuplicatePasswordFailures() bool`

HasIgnoreDuplicatePasswordFailures returns a boolean if a field has been set.

### GetFailureLockoutAction

`func (o *PasswordPolicyResponse) GetFailureLockoutAction() string`

GetFailureLockoutAction returns the FailureLockoutAction field if non-nil, zero value otherwise.

### GetFailureLockoutActionOk

`func (o *PasswordPolicyResponse) GetFailureLockoutActionOk() (*string, bool)`

GetFailureLockoutActionOk returns a tuple with the FailureLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLockoutAction

`func (o *PasswordPolicyResponse) SetFailureLockoutAction(v string)`

SetFailureLockoutAction sets FailureLockoutAction field to given value.

### HasFailureLockoutAction

`func (o *PasswordPolicyResponse) HasFailureLockoutAction() bool`

HasFailureLockoutAction returns a boolean if a field has been set.

### GetIdleLockoutInterval

`func (o *PasswordPolicyResponse) GetIdleLockoutInterval() string`

GetIdleLockoutInterval returns the IdleLockoutInterval field if non-nil, zero value otherwise.

### GetIdleLockoutIntervalOk

`func (o *PasswordPolicyResponse) GetIdleLockoutIntervalOk() (*string, bool)`

GetIdleLockoutIntervalOk returns a tuple with the IdleLockoutInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleLockoutInterval

`func (o *PasswordPolicyResponse) SetIdleLockoutInterval(v string)`

SetIdleLockoutInterval sets IdleLockoutInterval field to given value.

### HasIdleLockoutInterval

`func (o *PasswordPolicyResponse) HasIdleLockoutInterval() bool`

HasIdleLockoutInterval returns a boolean if a field has been set.

### GetAllowUserPasswordChanges

`func (o *PasswordPolicyResponse) GetAllowUserPasswordChanges() bool`

GetAllowUserPasswordChanges returns the AllowUserPasswordChanges field if non-nil, zero value otherwise.

### GetAllowUserPasswordChangesOk

`func (o *PasswordPolicyResponse) GetAllowUserPasswordChangesOk() (*bool, bool)`

GetAllowUserPasswordChangesOk returns a tuple with the AllowUserPasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserPasswordChanges

`func (o *PasswordPolicyResponse) SetAllowUserPasswordChanges(v bool)`

SetAllowUserPasswordChanges sets AllowUserPasswordChanges field to given value.

### HasAllowUserPasswordChanges

`func (o *PasswordPolicyResponse) HasAllowUserPasswordChanges() bool`

HasAllowUserPasswordChanges returns a boolean if a field has been set.

### GetPasswordChangeRequiresCurrentPassword

`func (o *PasswordPolicyResponse) GetPasswordChangeRequiresCurrentPassword() bool`

GetPasswordChangeRequiresCurrentPassword returns the PasswordChangeRequiresCurrentPassword field if non-nil, zero value otherwise.

### GetPasswordChangeRequiresCurrentPasswordOk

`func (o *PasswordPolicyResponse) GetPasswordChangeRequiresCurrentPasswordOk() (*bool, bool)`

GetPasswordChangeRequiresCurrentPasswordOk returns a tuple with the PasswordChangeRequiresCurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeRequiresCurrentPassword

`func (o *PasswordPolicyResponse) SetPasswordChangeRequiresCurrentPassword(v bool)`

SetPasswordChangeRequiresCurrentPassword sets PasswordChangeRequiresCurrentPassword field to given value.

### HasPasswordChangeRequiresCurrentPassword

`func (o *PasswordPolicyResponse) HasPasswordChangeRequiresCurrentPassword() bool`

HasPasswordChangeRequiresCurrentPassword returns a boolean if a field has been set.

### GetPasswordRetirementBehavior

`func (o *PasswordPolicyResponse) GetPasswordRetirementBehavior() []EnumpasswordPolicyPasswordRetirementBehaviorProp`

GetPasswordRetirementBehavior returns the PasswordRetirementBehavior field if non-nil, zero value otherwise.

### GetPasswordRetirementBehaviorOk

`func (o *PasswordPolicyResponse) GetPasswordRetirementBehaviorOk() (*[]EnumpasswordPolicyPasswordRetirementBehaviorProp, bool)`

GetPasswordRetirementBehaviorOk returns a tuple with the PasswordRetirementBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRetirementBehavior

`func (o *PasswordPolicyResponse) SetPasswordRetirementBehavior(v []EnumpasswordPolicyPasswordRetirementBehaviorProp)`

SetPasswordRetirementBehavior sets PasswordRetirementBehavior field to given value.

### HasPasswordRetirementBehavior

`func (o *PasswordPolicyResponse) HasPasswordRetirementBehavior() bool`

HasPasswordRetirementBehavior returns a boolean if a field has been set.

### GetMaxRetiredPasswordAge

`func (o *PasswordPolicyResponse) GetMaxRetiredPasswordAge() string`

GetMaxRetiredPasswordAge returns the MaxRetiredPasswordAge field if non-nil, zero value otherwise.

### GetMaxRetiredPasswordAgeOk

`func (o *PasswordPolicyResponse) GetMaxRetiredPasswordAgeOk() (*string, bool)`

GetMaxRetiredPasswordAgeOk returns a tuple with the MaxRetiredPasswordAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetiredPasswordAge

`func (o *PasswordPolicyResponse) SetMaxRetiredPasswordAge(v string)`

SetMaxRetiredPasswordAge sets MaxRetiredPasswordAge field to given value.

### HasMaxRetiredPasswordAge

`func (o *PasswordPolicyResponse) HasMaxRetiredPasswordAge() bool`

HasMaxRetiredPasswordAge returns a boolean if a field has been set.

### GetAllowedPasswordResetTokenUseCondition

`func (o *PasswordPolicyResponse) GetAllowedPasswordResetTokenUseCondition() []EnumpasswordPolicyAllowedPasswordResetTokenUseConditionProp`

GetAllowedPasswordResetTokenUseCondition returns the AllowedPasswordResetTokenUseCondition field if non-nil, zero value otherwise.

### GetAllowedPasswordResetTokenUseConditionOk

`func (o *PasswordPolicyResponse) GetAllowedPasswordResetTokenUseConditionOk() (*[]EnumpasswordPolicyAllowedPasswordResetTokenUseConditionProp, bool)`

GetAllowedPasswordResetTokenUseConditionOk returns a tuple with the AllowedPasswordResetTokenUseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPasswordResetTokenUseCondition

`func (o *PasswordPolicyResponse) SetAllowedPasswordResetTokenUseCondition(v []EnumpasswordPolicyAllowedPasswordResetTokenUseConditionProp)`

SetAllowedPasswordResetTokenUseCondition sets AllowedPasswordResetTokenUseCondition field to given value.

### HasAllowedPasswordResetTokenUseCondition

`func (o *PasswordPolicyResponse) HasAllowedPasswordResetTokenUseCondition() bool`

HasAllowedPasswordResetTokenUseCondition returns a boolean if a field has been set.

### GetForceChangeOnAdd

`func (o *PasswordPolicyResponse) GetForceChangeOnAdd() bool`

GetForceChangeOnAdd returns the ForceChangeOnAdd field if non-nil, zero value otherwise.

### GetForceChangeOnAddOk

`func (o *PasswordPolicyResponse) GetForceChangeOnAddOk() (*bool, bool)`

GetForceChangeOnAddOk returns a tuple with the ForceChangeOnAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChangeOnAdd

`func (o *PasswordPolicyResponse) SetForceChangeOnAdd(v bool)`

SetForceChangeOnAdd sets ForceChangeOnAdd field to given value.

### HasForceChangeOnAdd

`func (o *PasswordPolicyResponse) HasForceChangeOnAdd() bool`

HasForceChangeOnAdd returns a boolean if a field has been set.

### GetForceChangeOnReset

`func (o *PasswordPolicyResponse) GetForceChangeOnReset() bool`

GetForceChangeOnReset returns the ForceChangeOnReset field if non-nil, zero value otherwise.

### GetForceChangeOnResetOk

`func (o *PasswordPolicyResponse) GetForceChangeOnResetOk() (*bool, bool)`

GetForceChangeOnResetOk returns a tuple with the ForceChangeOnReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChangeOnReset

`func (o *PasswordPolicyResponse) SetForceChangeOnReset(v bool)`

SetForceChangeOnReset sets ForceChangeOnReset field to given value.

### HasForceChangeOnReset

`func (o *PasswordPolicyResponse) HasForceChangeOnReset() bool`

HasForceChangeOnReset returns a boolean if a field has been set.

### GetMaxPasswordResetAge

`func (o *PasswordPolicyResponse) GetMaxPasswordResetAge() string`

GetMaxPasswordResetAge returns the MaxPasswordResetAge field if non-nil, zero value otherwise.

### GetMaxPasswordResetAgeOk

`func (o *PasswordPolicyResponse) GetMaxPasswordResetAgeOk() (*string, bool)`

GetMaxPasswordResetAgeOk returns a tuple with the MaxPasswordResetAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordResetAge

`func (o *PasswordPolicyResponse) SetMaxPasswordResetAge(v string)`

SetMaxPasswordResetAge sets MaxPasswordResetAge field to given value.

### HasMaxPasswordResetAge

`func (o *PasswordPolicyResponse) HasMaxPasswordResetAge() bool`

HasMaxPasswordResetAge returns a boolean if a field has been set.

### GetSkipValidationForAdministrators

`func (o *PasswordPolicyResponse) GetSkipValidationForAdministrators() bool`

GetSkipValidationForAdministrators returns the SkipValidationForAdministrators field if non-nil, zero value otherwise.

### GetSkipValidationForAdministratorsOk

`func (o *PasswordPolicyResponse) GetSkipValidationForAdministratorsOk() (*bool, bool)`

GetSkipValidationForAdministratorsOk returns a tuple with the SkipValidationForAdministrators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipValidationForAdministrators

`func (o *PasswordPolicyResponse) SetSkipValidationForAdministrators(v bool)`

SetSkipValidationForAdministrators sets SkipValidationForAdministrators field to given value.

### HasSkipValidationForAdministrators

`func (o *PasswordPolicyResponse) HasSkipValidationForAdministrators() bool`

HasSkipValidationForAdministrators returns a boolean if a field has been set.

### GetMaximumRecentLoginHistorySuccessfulAuthenticationCount

`func (o *PasswordPolicyResponse) GetMaximumRecentLoginHistorySuccessfulAuthenticationCount() int64`

GetMaximumRecentLoginHistorySuccessfulAuthenticationCount returns the MaximumRecentLoginHistorySuccessfulAuthenticationCount field if non-nil, zero value otherwise.

### GetMaximumRecentLoginHistorySuccessfulAuthenticationCountOk

`func (o *PasswordPolicyResponse) GetMaximumRecentLoginHistorySuccessfulAuthenticationCountOk() (*int64, bool)`

GetMaximumRecentLoginHistorySuccessfulAuthenticationCountOk returns a tuple with the MaximumRecentLoginHistorySuccessfulAuthenticationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRecentLoginHistorySuccessfulAuthenticationCount

`func (o *PasswordPolicyResponse) SetMaximumRecentLoginHistorySuccessfulAuthenticationCount(v int64)`

SetMaximumRecentLoginHistorySuccessfulAuthenticationCount sets MaximumRecentLoginHistorySuccessfulAuthenticationCount field to given value.

### HasMaximumRecentLoginHistorySuccessfulAuthenticationCount

`func (o *PasswordPolicyResponse) HasMaximumRecentLoginHistorySuccessfulAuthenticationCount() bool`

HasMaximumRecentLoginHistorySuccessfulAuthenticationCount returns a boolean if a field has been set.

### GetMaximumRecentLoginHistorySuccessfulAuthenticationDuration

`func (o *PasswordPolicyResponse) GetMaximumRecentLoginHistorySuccessfulAuthenticationDuration() string`

GetMaximumRecentLoginHistorySuccessfulAuthenticationDuration returns the MaximumRecentLoginHistorySuccessfulAuthenticationDuration field if non-nil, zero value otherwise.

### GetMaximumRecentLoginHistorySuccessfulAuthenticationDurationOk

`func (o *PasswordPolicyResponse) GetMaximumRecentLoginHistorySuccessfulAuthenticationDurationOk() (*string, bool)`

GetMaximumRecentLoginHistorySuccessfulAuthenticationDurationOk returns a tuple with the MaximumRecentLoginHistorySuccessfulAuthenticationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRecentLoginHistorySuccessfulAuthenticationDuration

`func (o *PasswordPolicyResponse) SetMaximumRecentLoginHistorySuccessfulAuthenticationDuration(v string)`

SetMaximumRecentLoginHistorySuccessfulAuthenticationDuration sets MaximumRecentLoginHistorySuccessfulAuthenticationDuration field to given value.

### HasMaximumRecentLoginHistorySuccessfulAuthenticationDuration

`func (o *PasswordPolicyResponse) HasMaximumRecentLoginHistorySuccessfulAuthenticationDuration() bool`

HasMaximumRecentLoginHistorySuccessfulAuthenticationDuration returns a boolean if a field has been set.

### GetMaximumRecentLoginHistoryFailedAuthenticationCount

`func (o *PasswordPolicyResponse) GetMaximumRecentLoginHistoryFailedAuthenticationCount() int64`

GetMaximumRecentLoginHistoryFailedAuthenticationCount returns the MaximumRecentLoginHistoryFailedAuthenticationCount field if non-nil, zero value otherwise.

### GetMaximumRecentLoginHistoryFailedAuthenticationCountOk

`func (o *PasswordPolicyResponse) GetMaximumRecentLoginHistoryFailedAuthenticationCountOk() (*int64, bool)`

GetMaximumRecentLoginHistoryFailedAuthenticationCountOk returns a tuple with the MaximumRecentLoginHistoryFailedAuthenticationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRecentLoginHistoryFailedAuthenticationCount

`func (o *PasswordPolicyResponse) SetMaximumRecentLoginHistoryFailedAuthenticationCount(v int64)`

SetMaximumRecentLoginHistoryFailedAuthenticationCount sets MaximumRecentLoginHistoryFailedAuthenticationCount field to given value.

### HasMaximumRecentLoginHistoryFailedAuthenticationCount

`func (o *PasswordPolicyResponse) HasMaximumRecentLoginHistoryFailedAuthenticationCount() bool`

HasMaximumRecentLoginHistoryFailedAuthenticationCount returns a boolean if a field has been set.

### GetMaximumRecentLoginHistoryFailedAuthenticationDuration

`func (o *PasswordPolicyResponse) GetMaximumRecentLoginHistoryFailedAuthenticationDuration() string`

GetMaximumRecentLoginHistoryFailedAuthenticationDuration returns the MaximumRecentLoginHistoryFailedAuthenticationDuration field if non-nil, zero value otherwise.

### GetMaximumRecentLoginHistoryFailedAuthenticationDurationOk

`func (o *PasswordPolicyResponse) GetMaximumRecentLoginHistoryFailedAuthenticationDurationOk() (*string, bool)`

GetMaximumRecentLoginHistoryFailedAuthenticationDurationOk returns a tuple with the MaximumRecentLoginHistoryFailedAuthenticationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRecentLoginHistoryFailedAuthenticationDuration

`func (o *PasswordPolicyResponse) SetMaximumRecentLoginHistoryFailedAuthenticationDuration(v string)`

SetMaximumRecentLoginHistoryFailedAuthenticationDuration sets MaximumRecentLoginHistoryFailedAuthenticationDuration field to given value.

### HasMaximumRecentLoginHistoryFailedAuthenticationDuration

`func (o *PasswordPolicyResponse) HasMaximumRecentLoginHistoryFailedAuthenticationDuration() bool`

HasMaximumRecentLoginHistoryFailedAuthenticationDuration returns a boolean if a field has been set.

### GetRecentLoginHistorySimilarAttemptBehavior

`func (o *PasswordPolicyResponse) GetRecentLoginHistorySimilarAttemptBehavior() EnumpasswordPolicyRecentLoginHistorySimilarAttemptBehaviorProp`

GetRecentLoginHistorySimilarAttemptBehavior returns the RecentLoginHistorySimilarAttemptBehavior field if non-nil, zero value otherwise.

### GetRecentLoginHistorySimilarAttemptBehaviorOk

`func (o *PasswordPolicyResponse) GetRecentLoginHistorySimilarAttemptBehaviorOk() (*EnumpasswordPolicyRecentLoginHistorySimilarAttemptBehaviorProp, bool)`

GetRecentLoginHistorySimilarAttemptBehaviorOk returns a tuple with the RecentLoginHistorySimilarAttemptBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentLoginHistorySimilarAttemptBehavior

`func (o *PasswordPolicyResponse) SetRecentLoginHistorySimilarAttemptBehavior(v EnumpasswordPolicyRecentLoginHistorySimilarAttemptBehaviorProp)`

SetRecentLoginHistorySimilarAttemptBehavior sets RecentLoginHistorySimilarAttemptBehavior field to given value.

### HasRecentLoginHistorySimilarAttemptBehavior

`func (o *PasswordPolicyResponse) HasRecentLoginHistorySimilarAttemptBehavior() bool`

HasRecentLoginHistorySimilarAttemptBehavior returns a boolean if a field has been set.

### GetLastLoginIPAddressAttribute

`func (o *PasswordPolicyResponse) GetLastLoginIPAddressAttribute() string`

GetLastLoginIPAddressAttribute returns the LastLoginIPAddressAttribute field if non-nil, zero value otherwise.

### GetLastLoginIPAddressAttributeOk

`func (o *PasswordPolicyResponse) GetLastLoginIPAddressAttributeOk() (*string, bool)`

GetLastLoginIPAddressAttributeOk returns a tuple with the LastLoginIPAddressAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginIPAddressAttribute

`func (o *PasswordPolicyResponse) SetLastLoginIPAddressAttribute(v string)`

SetLastLoginIPAddressAttribute sets LastLoginIPAddressAttribute field to given value.

### HasLastLoginIPAddressAttribute

`func (o *PasswordPolicyResponse) HasLastLoginIPAddressAttribute() bool`

HasLastLoginIPAddressAttribute returns a boolean if a field has been set.

### GetLastLoginTimeAttribute

`func (o *PasswordPolicyResponse) GetLastLoginTimeAttribute() string`

GetLastLoginTimeAttribute returns the LastLoginTimeAttribute field if non-nil, zero value otherwise.

### GetLastLoginTimeAttributeOk

`func (o *PasswordPolicyResponse) GetLastLoginTimeAttributeOk() (*string, bool)`

GetLastLoginTimeAttributeOk returns a tuple with the LastLoginTimeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimeAttribute

`func (o *PasswordPolicyResponse) SetLastLoginTimeAttribute(v string)`

SetLastLoginTimeAttribute sets LastLoginTimeAttribute field to given value.

### HasLastLoginTimeAttribute

`func (o *PasswordPolicyResponse) HasLastLoginTimeAttribute() bool`

HasLastLoginTimeAttribute returns a boolean if a field has been set.

### GetLastLoginTimeFormat

`func (o *PasswordPolicyResponse) GetLastLoginTimeFormat() string`

GetLastLoginTimeFormat returns the LastLoginTimeFormat field if non-nil, zero value otherwise.

### GetLastLoginTimeFormatOk

`func (o *PasswordPolicyResponse) GetLastLoginTimeFormatOk() (*string, bool)`

GetLastLoginTimeFormatOk returns a tuple with the LastLoginTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimeFormat

`func (o *PasswordPolicyResponse) SetLastLoginTimeFormat(v string)`

SetLastLoginTimeFormat sets LastLoginTimeFormat field to given value.

### HasLastLoginTimeFormat

`func (o *PasswordPolicyResponse) HasLastLoginTimeFormat() bool`

HasLastLoginTimeFormat returns a boolean if a field has been set.

### GetPreviousLastLoginTimeFormat

`func (o *PasswordPolicyResponse) GetPreviousLastLoginTimeFormat() []string`

GetPreviousLastLoginTimeFormat returns the PreviousLastLoginTimeFormat field if non-nil, zero value otherwise.

### GetPreviousLastLoginTimeFormatOk

`func (o *PasswordPolicyResponse) GetPreviousLastLoginTimeFormatOk() (*[]string, bool)`

GetPreviousLastLoginTimeFormatOk returns a tuple with the PreviousLastLoginTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLastLoginTimeFormat

`func (o *PasswordPolicyResponse) SetPreviousLastLoginTimeFormat(v []string)`

SetPreviousLastLoginTimeFormat sets PreviousLastLoginTimeFormat field to given value.

### HasPreviousLastLoginTimeFormat

`func (o *PasswordPolicyResponse) HasPreviousLastLoginTimeFormat() bool`

HasPreviousLastLoginTimeFormat returns a boolean if a field has been set.

### GetMeta

`func (o *PasswordPolicyResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PasswordPolicyResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PasswordPolicyResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PasswordPolicyResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PasswordPolicyResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordPolicyResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PasswordPolicyResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


