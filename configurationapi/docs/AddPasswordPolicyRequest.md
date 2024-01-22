# AddPasswordPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**ReEncodePasswordsOnSchemeConfigChange** | Pointer to **bool** | Indicates whether to re-encode passwords on authentication if the configuration for the underlying password storage scheme has changed. | [optional] 
**AllowMultiplePasswordValues** | Pointer to **bool** | Indicates whether user entries can have multiple distinct values for the password attribute. | [optional] 
**AllowPreEncodedPasswords** | Pointer to [**EnumpasswordPolicyAllowPreEncodedPasswordsProp**](EnumpasswordPolicyAllowPreEncodedPasswordsProp.md) |  | [optional] 
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
**PolicyName** | **string** | Name of the new Password Policy | 

## Methods

### NewAddPasswordPolicyRequest

`func NewAddPasswordPolicyRequest(passwordAttribute string, defaultPasswordStorageScheme []string, policyName string, ) *AddPasswordPolicyRequest`

NewAddPasswordPolicyRequest instantiates a new AddPasswordPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPasswordPolicyRequestWithDefaults

`func NewAddPasswordPolicyRequestWithDefaults() *AddPasswordPolicyRequest`

NewAddPasswordPolicyRequestWithDefaults instantiates a new AddPasswordPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddPasswordPolicyRequest) GetSchemas() []EnumpasswordPolicySchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPasswordPolicyRequest) GetSchemasOk() (*[]EnumpasswordPolicySchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPasswordPolicyRequest) SetSchemas(v []EnumpasswordPolicySchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddPasswordPolicyRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetDescription

`func (o *AddPasswordPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPasswordPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPasswordPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPasswordPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequireSecureAuthentication

`func (o *AddPasswordPolicyRequest) GetRequireSecureAuthentication() bool`

GetRequireSecureAuthentication returns the RequireSecureAuthentication field if non-nil, zero value otherwise.

### GetRequireSecureAuthenticationOk

`func (o *AddPasswordPolicyRequest) GetRequireSecureAuthenticationOk() (*bool, bool)`

GetRequireSecureAuthenticationOk returns a tuple with the RequireSecureAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSecureAuthentication

`func (o *AddPasswordPolicyRequest) SetRequireSecureAuthentication(v bool)`

SetRequireSecureAuthentication sets RequireSecureAuthentication field to given value.

### HasRequireSecureAuthentication

`func (o *AddPasswordPolicyRequest) HasRequireSecureAuthentication() bool`

HasRequireSecureAuthentication returns a boolean if a field has been set.

### GetRequireSecurePasswordChanges

`func (o *AddPasswordPolicyRequest) GetRequireSecurePasswordChanges() bool`

GetRequireSecurePasswordChanges returns the RequireSecurePasswordChanges field if non-nil, zero value otherwise.

### GetRequireSecurePasswordChangesOk

`func (o *AddPasswordPolicyRequest) GetRequireSecurePasswordChangesOk() (*bool, bool)`

GetRequireSecurePasswordChangesOk returns a tuple with the RequireSecurePasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSecurePasswordChanges

`func (o *AddPasswordPolicyRequest) SetRequireSecurePasswordChanges(v bool)`

SetRequireSecurePasswordChanges sets RequireSecurePasswordChanges field to given value.

### HasRequireSecurePasswordChanges

`func (o *AddPasswordPolicyRequest) HasRequireSecurePasswordChanges() bool`

HasRequireSecurePasswordChanges returns a boolean if a field has been set.

### GetAccountStatusNotificationHandler

`func (o *AddPasswordPolicyRequest) GetAccountStatusNotificationHandler() []string`

GetAccountStatusNotificationHandler returns the AccountStatusNotificationHandler field if non-nil, zero value otherwise.

### GetAccountStatusNotificationHandlerOk

`func (o *AddPasswordPolicyRequest) GetAccountStatusNotificationHandlerOk() (*[]string, bool)`

GetAccountStatusNotificationHandlerOk returns a tuple with the AccountStatusNotificationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatusNotificationHandler

`func (o *AddPasswordPolicyRequest) SetAccountStatusNotificationHandler(v []string)`

SetAccountStatusNotificationHandler sets AccountStatusNotificationHandler field to given value.

### HasAccountStatusNotificationHandler

`func (o *AddPasswordPolicyRequest) HasAccountStatusNotificationHandler() bool`

HasAccountStatusNotificationHandler returns a boolean if a field has been set.

### GetStateUpdateFailurePolicy

`func (o *AddPasswordPolicyRequest) GetStateUpdateFailurePolicy() EnumpasswordPolicyStateUpdateFailurePolicyProp`

GetStateUpdateFailurePolicy returns the StateUpdateFailurePolicy field if non-nil, zero value otherwise.

### GetStateUpdateFailurePolicyOk

`func (o *AddPasswordPolicyRequest) GetStateUpdateFailurePolicyOk() (*EnumpasswordPolicyStateUpdateFailurePolicyProp, bool)`

GetStateUpdateFailurePolicyOk returns a tuple with the StateUpdateFailurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateUpdateFailurePolicy

`func (o *AddPasswordPolicyRequest) SetStateUpdateFailurePolicy(v EnumpasswordPolicyStateUpdateFailurePolicyProp)`

SetStateUpdateFailurePolicy sets StateUpdateFailurePolicy field to given value.

### HasStateUpdateFailurePolicy

`func (o *AddPasswordPolicyRequest) HasStateUpdateFailurePolicy() bool`

HasStateUpdateFailurePolicy returns a boolean if a field has been set.

### GetEnableDebug

`func (o *AddPasswordPolicyRequest) GetEnableDebug() bool`

GetEnableDebug returns the EnableDebug field if non-nil, zero value otherwise.

### GetEnableDebugOk

`func (o *AddPasswordPolicyRequest) GetEnableDebugOk() (*bool, bool)`

GetEnableDebugOk returns a tuple with the EnableDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebug

`func (o *AddPasswordPolicyRequest) SetEnableDebug(v bool)`

SetEnableDebug sets EnableDebug field to given value.

### HasEnableDebug

`func (o *AddPasswordPolicyRequest) HasEnableDebug() bool`

HasEnableDebug returns a boolean if a field has been set.

### GetPasswordAttribute

`func (o *AddPasswordPolicyRequest) GetPasswordAttribute() string`

GetPasswordAttribute returns the PasswordAttribute field if non-nil, zero value otherwise.

### GetPasswordAttributeOk

`func (o *AddPasswordPolicyRequest) GetPasswordAttributeOk() (*string, bool)`

GetPasswordAttributeOk returns a tuple with the PasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAttribute

`func (o *AddPasswordPolicyRequest) SetPasswordAttribute(v string)`

SetPasswordAttribute sets PasswordAttribute field to given value.


### GetDefaultPasswordStorageScheme

`func (o *AddPasswordPolicyRequest) GetDefaultPasswordStorageScheme() []string`

GetDefaultPasswordStorageScheme returns the DefaultPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDefaultPasswordStorageSchemeOk

`func (o *AddPasswordPolicyRequest) GetDefaultPasswordStorageSchemeOk() (*[]string, bool)`

GetDefaultPasswordStorageSchemeOk returns a tuple with the DefaultPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPasswordStorageScheme

`func (o *AddPasswordPolicyRequest) SetDefaultPasswordStorageScheme(v []string)`

SetDefaultPasswordStorageScheme sets DefaultPasswordStorageScheme field to given value.


### GetDeprecatedPasswordStorageScheme

`func (o *AddPasswordPolicyRequest) GetDeprecatedPasswordStorageScheme() []string`

GetDeprecatedPasswordStorageScheme returns the DeprecatedPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDeprecatedPasswordStorageSchemeOk

`func (o *AddPasswordPolicyRequest) GetDeprecatedPasswordStorageSchemeOk() (*[]string, bool)`

GetDeprecatedPasswordStorageSchemeOk returns a tuple with the DeprecatedPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedPasswordStorageScheme

`func (o *AddPasswordPolicyRequest) SetDeprecatedPasswordStorageScheme(v []string)`

SetDeprecatedPasswordStorageScheme sets DeprecatedPasswordStorageScheme field to given value.

### HasDeprecatedPasswordStorageScheme

`func (o *AddPasswordPolicyRequest) HasDeprecatedPasswordStorageScheme() bool`

HasDeprecatedPasswordStorageScheme returns a boolean if a field has been set.

### GetReEncodePasswordsOnSchemeConfigChange

`func (o *AddPasswordPolicyRequest) GetReEncodePasswordsOnSchemeConfigChange() bool`

GetReEncodePasswordsOnSchemeConfigChange returns the ReEncodePasswordsOnSchemeConfigChange field if non-nil, zero value otherwise.

### GetReEncodePasswordsOnSchemeConfigChangeOk

`func (o *AddPasswordPolicyRequest) GetReEncodePasswordsOnSchemeConfigChangeOk() (*bool, bool)`

GetReEncodePasswordsOnSchemeConfigChangeOk returns a tuple with the ReEncodePasswordsOnSchemeConfigChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReEncodePasswordsOnSchemeConfigChange

`func (o *AddPasswordPolicyRequest) SetReEncodePasswordsOnSchemeConfigChange(v bool)`

SetReEncodePasswordsOnSchemeConfigChange sets ReEncodePasswordsOnSchemeConfigChange field to given value.

### HasReEncodePasswordsOnSchemeConfigChange

`func (o *AddPasswordPolicyRequest) HasReEncodePasswordsOnSchemeConfigChange() bool`

HasReEncodePasswordsOnSchemeConfigChange returns a boolean if a field has been set.

### GetAllowMultiplePasswordValues

`func (o *AddPasswordPolicyRequest) GetAllowMultiplePasswordValues() bool`

GetAllowMultiplePasswordValues returns the AllowMultiplePasswordValues field if non-nil, zero value otherwise.

### GetAllowMultiplePasswordValuesOk

`func (o *AddPasswordPolicyRequest) GetAllowMultiplePasswordValuesOk() (*bool, bool)`

GetAllowMultiplePasswordValuesOk returns a tuple with the AllowMultiplePasswordValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiplePasswordValues

`func (o *AddPasswordPolicyRequest) SetAllowMultiplePasswordValues(v bool)`

SetAllowMultiplePasswordValues sets AllowMultiplePasswordValues field to given value.

### HasAllowMultiplePasswordValues

`func (o *AddPasswordPolicyRequest) HasAllowMultiplePasswordValues() bool`

HasAllowMultiplePasswordValues returns a boolean if a field has been set.

### GetAllowPreEncodedPasswords

`func (o *AddPasswordPolicyRequest) GetAllowPreEncodedPasswords() EnumpasswordPolicyAllowPreEncodedPasswordsProp`

GetAllowPreEncodedPasswords returns the AllowPreEncodedPasswords field if non-nil, zero value otherwise.

### GetAllowPreEncodedPasswordsOk

`func (o *AddPasswordPolicyRequest) GetAllowPreEncodedPasswordsOk() (*EnumpasswordPolicyAllowPreEncodedPasswordsProp, bool)`

GetAllowPreEncodedPasswordsOk returns a tuple with the AllowPreEncodedPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPreEncodedPasswords

`func (o *AddPasswordPolicyRequest) SetAllowPreEncodedPasswords(v EnumpasswordPolicyAllowPreEncodedPasswordsProp)`

SetAllowPreEncodedPasswords sets AllowPreEncodedPasswords field to given value.

### HasAllowPreEncodedPasswords

`func (o *AddPasswordPolicyRequest) HasAllowPreEncodedPasswords() bool`

HasAllowPreEncodedPasswords returns a boolean if a field has been set.

### GetPasswordValidator

`func (o *AddPasswordPolicyRequest) GetPasswordValidator() []string`

GetPasswordValidator returns the PasswordValidator field if non-nil, zero value otherwise.

### GetPasswordValidatorOk

`func (o *AddPasswordPolicyRequest) GetPasswordValidatorOk() (*[]string, bool)`

GetPasswordValidatorOk returns a tuple with the PasswordValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordValidator

`func (o *AddPasswordPolicyRequest) SetPasswordValidator(v []string)`

SetPasswordValidator sets PasswordValidator field to given value.

### HasPasswordValidator

`func (o *AddPasswordPolicyRequest) HasPasswordValidator() bool`

HasPasswordValidator returns a boolean if a field has been set.

### GetBindPasswordValidator

`func (o *AddPasswordPolicyRequest) GetBindPasswordValidator() []string`

GetBindPasswordValidator returns the BindPasswordValidator field if non-nil, zero value otherwise.

### GetBindPasswordValidatorOk

`func (o *AddPasswordPolicyRequest) GetBindPasswordValidatorOk() (*[]string, bool)`

GetBindPasswordValidatorOk returns a tuple with the BindPasswordValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPasswordValidator

`func (o *AddPasswordPolicyRequest) SetBindPasswordValidator(v []string)`

SetBindPasswordValidator sets BindPasswordValidator field to given value.

### HasBindPasswordValidator

`func (o *AddPasswordPolicyRequest) HasBindPasswordValidator() bool`

HasBindPasswordValidator returns a boolean if a field has been set.

### GetMinimumBindPasswordValidationFrequency

`func (o *AddPasswordPolicyRequest) GetMinimumBindPasswordValidationFrequency() string`

GetMinimumBindPasswordValidationFrequency returns the MinimumBindPasswordValidationFrequency field if non-nil, zero value otherwise.

### GetMinimumBindPasswordValidationFrequencyOk

`func (o *AddPasswordPolicyRequest) GetMinimumBindPasswordValidationFrequencyOk() (*string, bool)`

GetMinimumBindPasswordValidationFrequencyOk returns a tuple with the MinimumBindPasswordValidationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumBindPasswordValidationFrequency

`func (o *AddPasswordPolicyRequest) SetMinimumBindPasswordValidationFrequency(v string)`

SetMinimumBindPasswordValidationFrequency sets MinimumBindPasswordValidationFrequency field to given value.

### HasMinimumBindPasswordValidationFrequency

`func (o *AddPasswordPolicyRequest) HasMinimumBindPasswordValidationFrequency() bool`

HasMinimumBindPasswordValidationFrequency returns a boolean if a field has been set.

### GetBindPasswordValidationFailureAction

`func (o *AddPasswordPolicyRequest) GetBindPasswordValidationFailureAction() EnumpasswordPolicyBindPasswordValidationFailureActionProp`

GetBindPasswordValidationFailureAction returns the BindPasswordValidationFailureAction field if non-nil, zero value otherwise.

### GetBindPasswordValidationFailureActionOk

`func (o *AddPasswordPolicyRequest) GetBindPasswordValidationFailureActionOk() (*EnumpasswordPolicyBindPasswordValidationFailureActionProp, bool)`

GetBindPasswordValidationFailureActionOk returns a tuple with the BindPasswordValidationFailureAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPasswordValidationFailureAction

`func (o *AddPasswordPolicyRequest) SetBindPasswordValidationFailureAction(v EnumpasswordPolicyBindPasswordValidationFailureActionProp)`

SetBindPasswordValidationFailureAction sets BindPasswordValidationFailureAction field to given value.

### HasBindPasswordValidationFailureAction

`func (o *AddPasswordPolicyRequest) HasBindPasswordValidationFailureAction() bool`

HasBindPasswordValidationFailureAction returns a boolean if a field has been set.

### GetPasswordGenerator

`func (o *AddPasswordPolicyRequest) GetPasswordGenerator() string`

GetPasswordGenerator returns the PasswordGenerator field if non-nil, zero value otherwise.

### GetPasswordGeneratorOk

`func (o *AddPasswordPolicyRequest) GetPasswordGeneratorOk() (*string, bool)`

GetPasswordGeneratorOk returns a tuple with the PasswordGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordGenerator

`func (o *AddPasswordPolicyRequest) SetPasswordGenerator(v string)`

SetPasswordGenerator sets PasswordGenerator field to given value.

### HasPasswordGenerator

`func (o *AddPasswordPolicyRequest) HasPasswordGenerator() bool`

HasPasswordGenerator returns a boolean if a field has been set.

### GetPasswordHistoryCount

`func (o *AddPasswordPolicyRequest) GetPasswordHistoryCount() int64`

GetPasswordHistoryCount returns the PasswordHistoryCount field if non-nil, zero value otherwise.

### GetPasswordHistoryCountOk

`func (o *AddPasswordPolicyRequest) GetPasswordHistoryCountOk() (*int64, bool)`

GetPasswordHistoryCountOk returns a tuple with the PasswordHistoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistoryCount

`func (o *AddPasswordPolicyRequest) SetPasswordHistoryCount(v int64)`

SetPasswordHistoryCount sets PasswordHistoryCount field to given value.

### HasPasswordHistoryCount

`func (o *AddPasswordPolicyRequest) HasPasswordHistoryCount() bool`

HasPasswordHistoryCount returns a boolean if a field has been set.

### GetPasswordHistoryDuration

`func (o *AddPasswordPolicyRequest) GetPasswordHistoryDuration() string`

GetPasswordHistoryDuration returns the PasswordHistoryDuration field if non-nil, zero value otherwise.

### GetPasswordHistoryDurationOk

`func (o *AddPasswordPolicyRequest) GetPasswordHistoryDurationOk() (*string, bool)`

GetPasswordHistoryDurationOk returns a tuple with the PasswordHistoryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHistoryDuration

`func (o *AddPasswordPolicyRequest) SetPasswordHistoryDuration(v string)`

SetPasswordHistoryDuration sets PasswordHistoryDuration field to given value.

### HasPasswordHistoryDuration

`func (o *AddPasswordPolicyRequest) HasPasswordHistoryDuration() bool`

HasPasswordHistoryDuration returns a boolean if a field has been set.

### GetMinPasswordAge

`func (o *AddPasswordPolicyRequest) GetMinPasswordAge() string`

GetMinPasswordAge returns the MinPasswordAge field if non-nil, zero value otherwise.

### GetMinPasswordAgeOk

`func (o *AddPasswordPolicyRequest) GetMinPasswordAgeOk() (*string, bool)`

GetMinPasswordAgeOk returns a tuple with the MinPasswordAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPasswordAge

`func (o *AddPasswordPolicyRequest) SetMinPasswordAge(v string)`

SetMinPasswordAge sets MinPasswordAge field to given value.

### HasMinPasswordAge

`func (o *AddPasswordPolicyRequest) HasMinPasswordAge() bool`

HasMinPasswordAge returns a boolean if a field has been set.

### GetMaxPasswordAge

`func (o *AddPasswordPolicyRequest) GetMaxPasswordAge() string`

GetMaxPasswordAge returns the MaxPasswordAge field if non-nil, zero value otherwise.

### GetMaxPasswordAgeOk

`func (o *AddPasswordPolicyRequest) GetMaxPasswordAgeOk() (*string, bool)`

GetMaxPasswordAgeOk returns a tuple with the MaxPasswordAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordAge

`func (o *AddPasswordPolicyRequest) SetMaxPasswordAge(v string)`

SetMaxPasswordAge sets MaxPasswordAge field to given value.

### HasMaxPasswordAge

`func (o *AddPasswordPolicyRequest) HasMaxPasswordAge() bool`

HasMaxPasswordAge returns a boolean if a field has been set.

### GetPasswordExpirationWarningInterval

`func (o *AddPasswordPolicyRequest) GetPasswordExpirationWarningInterval() string`

GetPasswordExpirationWarningInterval returns the PasswordExpirationWarningInterval field if non-nil, zero value otherwise.

### GetPasswordExpirationWarningIntervalOk

`func (o *AddPasswordPolicyRequest) GetPasswordExpirationWarningIntervalOk() (*string, bool)`

GetPasswordExpirationWarningIntervalOk returns a tuple with the PasswordExpirationWarningInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpirationWarningInterval

`func (o *AddPasswordPolicyRequest) SetPasswordExpirationWarningInterval(v string)`

SetPasswordExpirationWarningInterval sets PasswordExpirationWarningInterval field to given value.

### HasPasswordExpirationWarningInterval

`func (o *AddPasswordPolicyRequest) HasPasswordExpirationWarningInterval() bool`

HasPasswordExpirationWarningInterval returns a boolean if a field has been set.

### GetExpirePasswordsWithoutWarning

`func (o *AddPasswordPolicyRequest) GetExpirePasswordsWithoutWarning() bool`

GetExpirePasswordsWithoutWarning returns the ExpirePasswordsWithoutWarning field if non-nil, zero value otherwise.

### GetExpirePasswordsWithoutWarningOk

`func (o *AddPasswordPolicyRequest) GetExpirePasswordsWithoutWarningOk() (*bool, bool)`

GetExpirePasswordsWithoutWarningOk returns a tuple with the ExpirePasswordsWithoutWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirePasswordsWithoutWarning

`func (o *AddPasswordPolicyRequest) SetExpirePasswordsWithoutWarning(v bool)`

SetExpirePasswordsWithoutWarning sets ExpirePasswordsWithoutWarning field to given value.

### HasExpirePasswordsWithoutWarning

`func (o *AddPasswordPolicyRequest) HasExpirePasswordsWithoutWarning() bool`

HasExpirePasswordsWithoutWarning returns a boolean if a field has been set.

### GetReturnPasswordExpirationControls

`func (o *AddPasswordPolicyRequest) GetReturnPasswordExpirationControls() EnumpasswordPolicyReturnPasswordExpirationControlsProp`

GetReturnPasswordExpirationControls returns the ReturnPasswordExpirationControls field if non-nil, zero value otherwise.

### GetReturnPasswordExpirationControlsOk

`func (o *AddPasswordPolicyRequest) GetReturnPasswordExpirationControlsOk() (*EnumpasswordPolicyReturnPasswordExpirationControlsProp, bool)`

GetReturnPasswordExpirationControlsOk returns a tuple with the ReturnPasswordExpirationControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPasswordExpirationControls

`func (o *AddPasswordPolicyRequest) SetReturnPasswordExpirationControls(v EnumpasswordPolicyReturnPasswordExpirationControlsProp)`

SetReturnPasswordExpirationControls sets ReturnPasswordExpirationControls field to given value.

### HasReturnPasswordExpirationControls

`func (o *AddPasswordPolicyRequest) HasReturnPasswordExpirationControls() bool`

HasReturnPasswordExpirationControls returns a boolean if a field has been set.

### GetAllowExpiredPasswordChanges

`func (o *AddPasswordPolicyRequest) GetAllowExpiredPasswordChanges() bool`

GetAllowExpiredPasswordChanges returns the AllowExpiredPasswordChanges field if non-nil, zero value otherwise.

### GetAllowExpiredPasswordChangesOk

`func (o *AddPasswordPolicyRequest) GetAllowExpiredPasswordChangesOk() (*bool, bool)`

GetAllowExpiredPasswordChangesOk returns a tuple with the AllowExpiredPasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowExpiredPasswordChanges

`func (o *AddPasswordPolicyRequest) SetAllowExpiredPasswordChanges(v bool)`

SetAllowExpiredPasswordChanges sets AllowExpiredPasswordChanges field to given value.

### HasAllowExpiredPasswordChanges

`func (o *AddPasswordPolicyRequest) HasAllowExpiredPasswordChanges() bool`

HasAllowExpiredPasswordChanges returns a boolean if a field has been set.

### GetGraceLoginCount

`func (o *AddPasswordPolicyRequest) GetGraceLoginCount() int64`

GetGraceLoginCount returns the GraceLoginCount field if non-nil, zero value otherwise.

### GetGraceLoginCountOk

`func (o *AddPasswordPolicyRequest) GetGraceLoginCountOk() (*int64, bool)`

GetGraceLoginCountOk returns a tuple with the GraceLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceLoginCount

`func (o *AddPasswordPolicyRequest) SetGraceLoginCount(v int64)`

SetGraceLoginCount sets GraceLoginCount field to given value.

### HasGraceLoginCount

`func (o *AddPasswordPolicyRequest) HasGraceLoginCount() bool`

HasGraceLoginCount returns a boolean if a field has been set.

### GetRequireChangeByTime

`func (o *AddPasswordPolicyRequest) GetRequireChangeByTime() string`

GetRequireChangeByTime returns the RequireChangeByTime field if non-nil, zero value otherwise.

### GetRequireChangeByTimeOk

`func (o *AddPasswordPolicyRequest) GetRequireChangeByTimeOk() (*string, bool)`

GetRequireChangeByTimeOk returns a tuple with the RequireChangeByTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireChangeByTime

`func (o *AddPasswordPolicyRequest) SetRequireChangeByTime(v string)`

SetRequireChangeByTime sets RequireChangeByTime field to given value.

### HasRequireChangeByTime

`func (o *AddPasswordPolicyRequest) HasRequireChangeByTime() bool`

HasRequireChangeByTime returns a boolean if a field has been set.

### GetLockoutFailureCount

`func (o *AddPasswordPolicyRequest) GetLockoutFailureCount() int64`

GetLockoutFailureCount returns the LockoutFailureCount field if non-nil, zero value otherwise.

### GetLockoutFailureCountOk

`func (o *AddPasswordPolicyRequest) GetLockoutFailureCountOk() (*int64, bool)`

GetLockoutFailureCountOk returns a tuple with the LockoutFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutFailureCount

`func (o *AddPasswordPolicyRequest) SetLockoutFailureCount(v int64)`

SetLockoutFailureCount sets LockoutFailureCount field to given value.

### HasLockoutFailureCount

`func (o *AddPasswordPolicyRequest) HasLockoutFailureCount() bool`

HasLockoutFailureCount returns a boolean if a field has been set.

### GetLockoutDuration

`func (o *AddPasswordPolicyRequest) GetLockoutDuration() string`

GetLockoutDuration returns the LockoutDuration field if non-nil, zero value otherwise.

### GetLockoutDurationOk

`func (o *AddPasswordPolicyRequest) GetLockoutDurationOk() (*string, bool)`

GetLockoutDurationOk returns a tuple with the LockoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutDuration

`func (o *AddPasswordPolicyRequest) SetLockoutDuration(v string)`

SetLockoutDuration sets LockoutDuration field to given value.

### HasLockoutDuration

`func (o *AddPasswordPolicyRequest) HasLockoutDuration() bool`

HasLockoutDuration returns a boolean if a field has been set.

### GetLockoutFailureExpirationInterval

`func (o *AddPasswordPolicyRequest) GetLockoutFailureExpirationInterval() string`

GetLockoutFailureExpirationInterval returns the LockoutFailureExpirationInterval field if non-nil, zero value otherwise.

### GetLockoutFailureExpirationIntervalOk

`func (o *AddPasswordPolicyRequest) GetLockoutFailureExpirationIntervalOk() (*string, bool)`

GetLockoutFailureExpirationIntervalOk returns a tuple with the LockoutFailureExpirationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutFailureExpirationInterval

`func (o *AddPasswordPolicyRequest) SetLockoutFailureExpirationInterval(v string)`

SetLockoutFailureExpirationInterval sets LockoutFailureExpirationInterval field to given value.

### HasLockoutFailureExpirationInterval

`func (o *AddPasswordPolicyRequest) HasLockoutFailureExpirationInterval() bool`

HasLockoutFailureExpirationInterval returns a boolean if a field has been set.

### GetIgnoreDuplicatePasswordFailures

`func (o *AddPasswordPolicyRequest) GetIgnoreDuplicatePasswordFailures() bool`

GetIgnoreDuplicatePasswordFailures returns the IgnoreDuplicatePasswordFailures field if non-nil, zero value otherwise.

### GetIgnoreDuplicatePasswordFailuresOk

`func (o *AddPasswordPolicyRequest) GetIgnoreDuplicatePasswordFailuresOk() (*bool, bool)`

GetIgnoreDuplicatePasswordFailuresOk returns a tuple with the IgnoreDuplicatePasswordFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDuplicatePasswordFailures

`func (o *AddPasswordPolicyRequest) SetIgnoreDuplicatePasswordFailures(v bool)`

SetIgnoreDuplicatePasswordFailures sets IgnoreDuplicatePasswordFailures field to given value.

### HasIgnoreDuplicatePasswordFailures

`func (o *AddPasswordPolicyRequest) HasIgnoreDuplicatePasswordFailures() bool`

HasIgnoreDuplicatePasswordFailures returns a boolean if a field has been set.

### GetFailureLockoutAction

`func (o *AddPasswordPolicyRequest) GetFailureLockoutAction() string`

GetFailureLockoutAction returns the FailureLockoutAction field if non-nil, zero value otherwise.

### GetFailureLockoutActionOk

`func (o *AddPasswordPolicyRequest) GetFailureLockoutActionOk() (*string, bool)`

GetFailureLockoutActionOk returns a tuple with the FailureLockoutAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureLockoutAction

`func (o *AddPasswordPolicyRequest) SetFailureLockoutAction(v string)`

SetFailureLockoutAction sets FailureLockoutAction field to given value.

### HasFailureLockoutAction

`func (o *AddPasswordPolicyRequest) HasFailureLockoutAction() bool`

HasFailureLockoutAction returns a boolean if a field has been set.

### GetIdleLockoutInterval

`func (o *AddPasswordPolicyRequest) GetIdleLockoutInterval() string`

GetIdleLockoutInterval returns the IdleLockoutInterval field if non-nil, zero value otherwise.

### GetIdleLockoutIntervalOk

`func (o *AddPasswordPolicyRequest) GetIdleLockoutIntervalOk() (*string, bool)`

GetIdleLockoutIntervalOk returns a tuple with the IdleLockoutInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleLockoutInterval

`func (o *AddPasswordPolicyRequest) SetIdleLockoutInterval(v string)`

SetIdleLockoutInterval sets IdleLockoutInterval field to given value.

### HasIdleLockoutInterval

`func (o *AddPasswordPolicyRequest) HasIdleLockoutInterval() bool`

HasIdleLockoutInterval returns a boolean if a field has been set.

### GetAllowUserPasswordChanges

`func (o *AddPasswordPolicyRequest) GetAllowUserPasswordChanges() bool`

GetAllowUserPasswordChanges returns the AllowUserPasswordChanges field if non-nil, zero value otherwise.

### GetAllowUserPasswordChangesOk

`func (o *AddPasswordPolicyRequest) GetAllowUserPasswordChangesOk() (*bool, bool)`

GetAllowUserPasswordChangesOk returns a tuple with the AllowUserPasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUserPasswordChanges

`func (o *AddPasswordPolicyRequest) SetAllowUserPasswordChanges(v bool)`

SetAllowUserPasswordChanges sets AllowUserPasswordChanges field to given value.

### HasAllowUserPasswordChanges

`func (o *AddPasswordPolicyRequest) HasAllowUserPasswordChanges() bool`

HasAllowUserPasswordChanges returns a boolean if a field has been set.

### GetPasswordChangeRequiresCurrentPassword

`func (o *AddPasswordPolicyRequest) GetPasswordChangeRequiresCurrentPassword() bool`

GetPasswordChangeRequiresCurrentPassword returns the PasswordChangeRequiresCurrentPassword field if non-nil, zero value otherwise.

### GetPasswordChangeRequiresCurrentPasswordOk

`func (o *AddPasswordPolicyRequest) GetPasswordChangeRequiresCurrentPasswordOk() (*bool, bool)`

GetPasswordChangeRequiresCurrentPasswordOk returns a tuple with the PasswordChangeRequiresCurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeRequiresCurrentPassword

`func (o *AddPasswordPolicyRequest) SetPasswordChangeRequiresCurrentPassword(v bool)`

SetPasswordChangeRequiresCurrentPassword sets PasswordChangeRequiresCurrentPassword field to given value.

### HasPasswordChangeRequiresCurrentPassword

`func (o *AddPasswordPolicyRequest) HasPasswordChangeRequiresCurrentPassword() bool`

HasPasswordChangeRequiresCurrentPassword returns a boolean if a field has been set.

### GetPasswordRetirementBehavior

`func (o *AddPasswordPolicyRequest) GetPasswordRetirementBehavior() []EnumpasswordPolicyPasswordRetirementBehaviorProp`

GetPasswordRetirementBehavior returns the PasswordRetirementBehavior field if non-nil, zero value otherwise.

### GetPasswordRetirementBehaviorOk

`func (o *AddPasswordPolicyRequest) GetPasswordRetirementBehaviorOk() (*[]EnumpasswordPolicyPasswordRetirementBehaviorProp, bool)`

GetPasswordRetirementBehaviorOk returns a tuple with the PasswordRetirementBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordRetirementBehavior

`func (o *AddPasswordPolicyRequest) SetPasswordRetirementBehavior(v []EnumpasswordPolicyPasswordRetirementBehaviorProp)`

SetPasswordRetirementBehavior sets PasswordRetirementBehavior field to given value.

### HasPasswordRetirementBehavior

`func (o *AddPasswordPolicyRequest) HasPasswordRetirementBehavior() bool`

HasPasswordRetirementBehavior returns a boolean if a field has been set.

### GetMaxRetiredPasswordAge

`func (o *AddPasswordPolicyRequest) GetMaxRetiredPasswordAge() string`

GetMaxRetiredPasswordAge returns the MaxRetiredPasswordAge field if non-nil, zero value otherwise.

### GetMaxRetiredPasswordAgeOk

`func (o *AddPasswordPolicyRequest) GetMaxRetiredPasswordAgeOk() (*string, bool)`

GetMaxRetiredPasswordAgeOk returns a tuple with the MaxRetiredPasswordAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetiredPasswordAge

`func (o *AddPasswordPolicyRequest) SetMaxRetiredPasswordAge(v string)`

SetMaxRetiredPasswordAge sets MaxRetiredPasswordAge field to given value.

### HasMaxRetiredPasswordAge

`func (o *AddPasswordPolicyRequest) HasMaxRetiredPasswordAge() bool`

HasMaxRetiredPasswordAge returns a boolean if a field has been set.

### GetAllowedPasswordResetTokenUseCondition

`func (o *AddPasswordPolicyRequest) GetAllowedPasswordResetTokenUseCondition() []EnumpasswordPolicyAllowedPasswordResetTokenUseConditionProp`

GetAllowedPasswordResetTokenUseCondition returns the AllowedPasswordResetTokenUseCondition field if non-nil, zero value otherwise.

### GetAllowedPasswordResetTokenUseConditionOk

`func (o *AddPasswordPolicyRequest) GetAllowedPasswordResetTokenUseConditionOk() (*[]EnumpasswordPolicyAllowedPasswordResetTokenUseConditionProp, bool)`

GetAllowedPasswordResetTokenUseConditionOk returns a tuple with the AllowedPasswordResetTokenUseCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPasswordResetTokenUseCondition

`func (o *AddPasswordPolicyRequest) SetAllowedPasswordResetTokenUseCondition(v []EnumpasswordPolicyAllowedPasswordResetTokenUseConditionProp)`

SetAllowedPasswordResetTokenUseCondition sets AllowedPasswordResetTokenUseCondition field to given value.

### HasAllowedPasswordResetTokenUseCondition

`func (o *AddPasswordPolicyRequest) HasAllowedPasswordResetTokenUseCondition() bool`

HasAllowedPasswordResetTokenUseCondition returns a boolean if a field has been set.

### GetForceChangeOnAdd

`func (o *AddPasswordPolicyRequest) GetForceChangeOnAdd() bool`

GetForceChangeOnAdd returns the ForceChangeOnAdd field if non-nil, zero value otherwise.

### GetForceChangeOnAddOk

`func (o *AddPasswordPolicyRequest) GetForceChangeOnAddOk() (*bool, bool)`

GetForceChangeOnAddOk returns a tuple with the ForceChangeOnAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChangeOnAdd

`func (o *AddPasswordPolicyRequest) SetForceChangeOnAdd(v bool)`

SetForceChangeOnAdd sets ForceChangeOnAdd field to given value.

### HasForceChangeOnAdd

`func (o *AddPasswordPolicyRequest) HasForceChangeOnAdd() bool`

HasForceChangeOnAdd returns a boolean if a field has been set.

### GetForceChangeOnReset

`func (o *AddPasswordPolicyRequest) GetForceChangeOnReset() bool`

GetForceChangeOnReset returns the ForceChangeOnReset field if non-nil, zero value otherwise.

### GetForceChangeOnResetOk

`func (o *AddPasswordPolicyRequest) GetForceChangeOnResetOk() (*bool, bool)`

GetForceChangeOnResetOk returns a tuple with the ForceChangeOnReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChangeOnReset

`func (o *AddPasswordPolicyRequest) SetForceChangeOnReset(v bool)`

SetForceChangeOnReset sets ForceChangeOnReset field to given value.

### HasForceChangeOnReset

`func (o *AddPasswordPolicyRequest) HasForceChangeOnReset() bool`

HasForceChangeOnReset returns a boolean if a field has been set.

### GetMaxPasswordResetAge

`func (o *AddPasswordPolicyRequest) GetMaxPasswordResetAge() string`

GetMaxPasswordResetAge returns the MaxPasswordResetAge field if non-nil, zero value otherwise.

### GetMaxPasswordResetAgeOk

`func (o *AddPasswordPolicyRequest) GetMaxPasswordResetAgeOk() (*string, bool)`

GetMaxPasswordResetAgeOk returns a tuple with the MaxPasswordResetAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordResetAge

`func (o *AddPasswordPolicyRequest) SetMaxPasswordResetAge(v string)`

SetMaxPasswordResetAge sets MaxPasswordResetAge field to given value.

### HasMaxPasswordResetAge

`func (o *AddPasswordPolicyRequest) HasMaxPasswordResetAge() bool`

HasMaxPasswordResetAge returns a boolean if a field has been set.

### GetSkipValidationForAdministrators

`func (o *AddPasswordPolicyRequest) GetSkipValidationForAdministrators() bool`

GetSkipValidationForAdministrators returns the SkipValidationForAdministrators field if non-nil, zero value otherwise.

### GetSkipValidationForAdministratorsOk

`func (o *AddPasswordPolicyRequest) GetSkipValidationForAdministratorsOk() (*bool, bool)`

GetSkipValidationForAdministratorsOk returns a tuple with the SkipValidationForAdministrators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipValidationForAdministrators

`func (o *AddPasswordPolicyRequest) SetSkipValidationForAdministrators(v bool)`

SetSkipValidationForAdministrators sets SkipValidationForAdministrators field to given value.

### HasSkipValidationForAdministrators

`func (o *AddPasswordPolicyRequest) HasSkipValidationForAdministrators() bool`

HasSkipValidationForAdministrators returns a boolean if a field has been set.

### GetMaximumRecentLoginHistorySuccessfulAuthenticationCount

`func (o *AddPasswordPolicyRequest) GetMaximumRecentLoginHistorySuccessfulAuthenticationCount() int64`

GetMaximumRecentLoginHistorySuccessfulAuthenticationCount returns the MaximumRecentLoginHistorySuccessfulAuthenticationCount field if non-nil, zero value otherwise.

### GetMaximumRecentLoginHistorySuccessfulAuthenticationCountOk

`func (o *AddPasswordPolicyRequest) GetMaximumRecentLoginHistorySuccessfulAuthenticationCountOk() (*int64, bool)`

GetMaximumRecentLoginHistorySuccessfulAuthenticationCountOk returns a tuple with the MaximumRecentLoginHistorySuccessfulAuthenticationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRecentLoginHistorySuccessfulAuthenticationCount

`func (o *AddPasswordPolicyRequest) SetMaximumRecentLoginHistorySuccessfulAuthenticationCount(v int64)`

SetMaximumRecentLoginHistorySuccessfulAuthenticationCount sets MaximumRecentLoginHistorySuccessfulAuthenticationCount field to given value.

### HasMaximumRecentLoginHistorySuccessfulAuthenticationCount

`func (o *AddPasswordPolicyRequest) HasMaximumRecentLoginHistorySuccessfulAuthenticationCount() bool`

HasMaximumRecentLoginHistorySuccessfulAuthenticationCount returns a boolean if a field has been set.

### GetMaximumRecentLoginHistorySuccessfulAuthenticationDuration

`func (o *AddPasswordPolicyRequest) GetMaximumRecentLoginHistorySuccessfulAuthenticationDuration() string`

GetMaximumRecentLoginHistorySuccessfulAuthenticationDuration returns the MaximumRecentLoginHistorySuccessfulAuthenticationDuration field if non-nil, zero value otherwise.

### GetMaximumRecentLoginHistorySuccessfulAuthenticationDurationOk

`func (o *AddPasswordPolicyRequest) GetMaximumRecentLoginHistorySuccessfulAuthenticationDurationOk() (*string, bool)`

GetMaximumRecentLoginHistorySuccessfulAuthenticationDurationOk returns a tuple with the MaximumRecentLoginHistorySuccessfulAuthenticationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRecentLoginHistorySuccessfulAuthenticationDuration

`func (o *AddPasswordPolicyRequest) SetMaximumRecentLoginHistorySuccessfulAuthenticationDuration(v string)`

SetMaximumRecentLoginHistorySuccessfulAuthenticationDuration sets MaximumRecentLoginHistorySuccessfulAuthenticationDuration field to given value.

### HasMaximumRecentLoginHistorySuccessfulAuthenticationDuration

`func (o *AddPasswordPolicyRequest) HasMaximumRecentLoginHistorySuccessfulAuthenticationDuration() bool`

HasMaximumRecentLoginHistorySuccessfulAuthenticationDuration returns a boolean if a field has been set.

### GetMaximumRecentLoginHistoryFailedAuthenticationCount

`func (o *AddPasswordPolicyRequest) GetMaximumRecentLoginHistoryFailedAuthenticationCount() int64`

GetMaximumRecentLoginHistoryFailedAuthenticationCount returns the MaximumRecentLoginHistoryFailedAuthenticationCount field if non-nil, zero value otherwise.

### GetMaximumRecentLoginHistoryFailedAuthenticationCountOk

`func (o *AddPasswordPolicyRequest) GetMaximumRecentLoginHistoryFailedAuthenticationCountOk() (*int64, bool)`

GetMaximumRecentLoginHistoryFailedAuthenticationCountOk returns a tuple with the MaximumRecentLoginHistoryFailedAuthenticationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRecentLoginHistoryFailedAuthenticationCount

`func (o *AddPasswordPolicyRequest) SetMaximumRecentLoginHistoryFailedAuthenticationCount(v int64)`

SetMaximumRecentLoginHistoryFailedAuthenticationCount sets MaximumRecentLoginHistoryFailedAuthenticationCount field to given value.

### HasMaximumRecentLoginHistoryFailedAuthenticationCount

`func (o *AddPasswordPolicyRequest) HasMaximumRecentLoginHistoryFailedAuthenticationCount() bool`

HasMaximumRecentLoginHistoryFailedAuthenticationCount returns a boolean if a field has been set.

### GetMaximumRecentLoginHistoryFailedAuthenticationDuration

`func (o *AddPasswordPolicyRequest) GetMaximumRecentLoginHistoryFailedAuthenticationDuration() string`

GetMaximumRecentLoginHistoryFailedAuthenticationDuration returns the MaximumRecentLoginHistoryFailedAuthenticationDuration field if non-nil, zero value otherwise.

### GetMaximumRecentLoginHistoryFailedAuthenticationDurationOk

`func (o *AddPasswordPolicyRequest) GetMaximumRecentLoginHistoryFailedAuthenticationDurationOk() (*string, bool)`

GetMaximumRecentLoginHistoryFailedAuthenticationDurationOk returns a tuple with the MaximumRecentLoginHistoryFailedAuthenticationDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumRecentLoginHistoryFailedAuthenticationDuration

`func (o *AddPasswordPolicyRequest) SetMaximumRecentLoginHistoryFailedAuthenticationDuration(v string)`

SetMaximumRecentLoginHistoryFailedAuthenticationDuration sets MaximumRecentLoginHistoryFailedAuthenticationDuration field to given value.

### HasMaximumRecentLoginHistoryFailedAuthenticationDuration

`func (o *AddPasswordPolicyRequest) HasMaximumRecentLoginHistoryFailedAuthenticationDuration() bool`

HasMaximumRecentLoginHistoryFailedAuthenticationDuration returns a boolean if a field has been set.

### GetRecentLoginHistorySimilarAttemptBehavior

`func (o *AddPasswordPolicyRequest) GetRecentLoginHistorySimilarAttemptBehavior() EnumpasswordPolicyRecentLoginHistorySimilarAttemptBehaviorProp`

GetRecentLoginHistorySimilarAttemptBehavior returns the RecentLoginHistorySimilarAttemptBehavior field if non-nil, zero value otherwise.

### GetRecentLoginHistorySimilarAttemptBehaviorOk

`func (o *AddPasswordPolicyRequest) GetRecentLoginHistorySimilarAttemptBehaviorOk() (*EnumpasswordPolicyRecentLoginHistorySimilarAttemptBehaviorProp, bool)`

GetRecentLoginHistorySimilarAttemptBehaviorOk returns a tuple with the RecentLoginHistorySimilarAttemptBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentLoginHistorySimilarAttemptBehavior

`func (o *AddPasswordPolicyRequest) SetRecentLoginHistorySimilarAttemptBehavior(v EnumpasswordPolicyRecentLoginHistorySimilarAttemptBehaviorProp)`

SetRecentLoginHistorySimilarAttemptBehavior sets RecentLoginHistorySimilarAttemptBehavior field to given value.

### HasRecentLoginHistorySimilarAttemptBehavior

`func (o *AddPasswordPolicyRequest) HasRecentLoginHistorySimilarAttemptBehavior() bool`

HasRecentLoginHistorySimilarAttemptBehavior returns a boolean if a field has been set.

### GetLastLoginIPAddressAttribute

`func (o *AddPasswordPolicyRequest) GetLastLoginIPAddressAttribute() string`

GetLastLoginIPAddressAttribute returns the LastLoginIPAddressAttribute field if non-nil, zero value otherwise.

### GetLastLoginIPAddressAttributeOk

`func (o *AddPasswordPolicyRequest) GetLastLoginIPAddressAttributeOk() (*string, bool)`

GetLastLoginIPAddressAttributeOk returns a tuple with the LastLoginIPAddressAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginIPAddressAttribute

`func (o *AddPasswordPolicyRequest) SetLastLoginIPAddressAttribute(v string)`

SetLastLoginIPAddressAttribute sets LastLoginIPAddressAttribute field to given value.

### HasLastLoginIPAddressAttribute

`func (o *AddPasswordPolicyRequest) HasLastLoginIPAddressAttribute() bool`

HasLastLoginIPAddressAttribute returns a boolean if a field has been set.

### GetLastLoginTimeAttribute

`func (o *AddPasswordPolicyRequest) GetLastLoginTimeAttribute() string`

GetLastLoginTimeAttribute returns the LastLoginTimeAttribute field if non-nil, zero value otherwise.

### GetLastLoginTimeAttributeOk

`func (o *AddPasswordPolicyRequest) GetLastLoginTimeAttributeOk() (*string, bool)`

GetLastLoginTimeAttributeOk returns a tuple with the LastLoginTimeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimeAttribute

`func (o *AddPasswordPolicyRequest) SetLastLoginTimeAttribute(v string)`

SetLastLoginTimeAttribute sets LastLoginTimeAttribute field to given value.

### HasLastLoginTimeAttribute

`func (o *AddPasswordPolicyRequest) HasLastLoginTimeAttribute() bool`

HasLastLoginTimeAttribute returns a boolean if a field has been set.

### GetLastLoginTimeFormat

`func (o *AddPasswordPolicyRequest) GetLastLoginTimeFormat() string`

GetLastLoginTimeFormat returns the LastLoginTimeFormat field if non-nil, zero value otherwise.

### GetLastLoginTimeFormatOk

`func (o *AddPasswordPolicyRequest) GetLastLoginTimeFormatOk() (*string, bool)`

GetLastLoginTimeFormatOk returns a tuple with the LastLoginTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimeFormat

`func (o *AddPasswordPolicyRequest) SetLastLoginTimeFormat(v string)`

SetLastLoginTimeFormat sets LastLoginTimeFormat field to given value.

### HasLastLoginTimeFormat

`func (o *AddPasswordPolicyRequest) HasLastLoginTimeFormat() bool`

HasLastLoginTimeFormat returns a boolean if a field has been set.

### GetPreviousLastLoginTimeFormat

`func (o *AddPasswordPolicyRequest) GetPreviousLastLoginTimeFormat() []string`

GetPreviousLastLoginTimeFormat returns the PreviousLastLoginTimeFormat field if non-nil, zero value otherwise.

### GetPreviousLastLoginTimeFormatOk

`func (o *AddPasswordPolicyRequest) GetPreviousLastLoginTimeFormatOk() (*[]string, bool)`

GetPreviousLastLoginTimeFormatOk returns a tuple with the PreviousLastLoginTimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLastLoginTimeFormat

`func (o *AddPasswordPolicyRequest) SetPreviousLastLoginTimeFormat(v []string)`

SetPreviousLastLoginTimeFormat sets PreviousLastLoginTimeFormat field to given value.

### HasPreviousLastLoginTimeFormat

`func (o *AddPasswordPolicyRequest) HasPreviousLastLoginTimeFormat() bool`

HasPreviousLastLoginTimeFormat returns a boolean if a field has been set.

### GetPolicyName

`func (o *AddPasswordPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *AddPasswordPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *AddPasswordPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


