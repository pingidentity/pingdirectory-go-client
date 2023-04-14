# AddTopologyAdminUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | Name of the new Topology Admin User | 
**Schemas** | Pointer to [**[]EnumtopologyAdminUserSchemaUrn**](EnumtopologyAdminUserSchemaUrn.md) |  | [optional] 
**AlternateBindDN** | Pointer to **[]string** | Specifies one or more alternate DNs that can be used to bind to the server as this User. | [optional] 
**Description** | Pointer to **string** | A description for this User. | [optional] 
**Password** | Pointer to **string** | Specifies the user&#39;s password. This is stored in the userPassword LDAP attribute. To set a pre-hashed value, the account making the change must have the bypass-pw-policy privilege. | [optional] 
**FirstName** | Pointer to **[]string** | Specifies the user&#39;s first name. This is stored in the givenName LDAP attribute. | [optional] 
**LastName** | Pointer to **[]string** | Specifies the user&#39;s last name. This is stored in the sn LDAP attribute. | [optional] 
**UserID** | Pointer to **string** | Specifies the user&#39;s user ID. This is stored in the uid LDAP attribute. | [optional] 
**EmailAddress** | Pointer to **[]string** | Specifies the user&#39;s email address. This is stored in the mail LDAP attribute. | [optional] 
**WorkTelephoneNumber** | Pointer to **[]string** | Specifies the user&#39;s work telephone number. This is stored in the telephoneNumber LDAP attribute. | [optional] 
**HomeTelephoneNumber** | Pointer to **[]string** | Specifies the user&#39;s home telephone number. This is stored in the homePhone LDAP attribute. | [optional] 
**MobileTelephoneNumber** | Pointer to **[]string** | Specifies the user&#39;s mobile telephone number. This is stored in the mobile LDAP attribute. | [optional] 
**PagerTelephoneNumber** | Pointer to **[]string** | Specifies the user&#39;s pager telephone number. This is stored in the pager LDAP attribute. | [optional] 
**InheritDefaultRootPrivileges** | Pointer to **bool** | Indicates whether this User should be automatically granted the set of privileges defined in the default-root-privilege-name property of the Root DN configuration object. | [optional] 
**Privilege** | Pointer to [**[]EnumtopologyAdminUserPrivilegeProp**](EnumtopologyAdminUserPrivilegeProp.md) |  | [optional] 
**SearchResultEntryLimit** | Pointer to **int64** | Specifies the maximum number of entries that the server may return to the user in response to any single search request. A value of 0 indicates no limit should be enforced. This is stored in the ds-rlim-size-limit LDAP attribute. | [optional] 
**TimeLimitSeconds** | Pointer to **int64** | Specifies the maximum length of time (in seconds) that the server may spend processing any single search request. A value of 0 indicates no limit should be enforced. This is stored in the ds-rlim-time-limit LDAP attribute. | [optional] 
**LookThroughEntryLimit** | Pointer to **int64** | Specifies the maximum number of candidate entries that the server may examine in the course of processing any single search request. A value of 0 indicates no limit should be enforced. This is stored in the ds-rlim-lookthrough-limit LDAP attribute. | [optional] 
**IdleTimeLimitSeconds** | Pointer to **int64** | Specifies the maximum length of time (in seconds) that a connection authenticated as this user may remain established without issuing any requests. A value of 0 indicates no limit should be enforced. This is stored in the ds-rlim-idle-time-limit LDAP attribute. | [optional] 
**PasswordPolicy** | Pointer to **string** | Specifies the password policy for the user. This is stored in the ds-pwp-password-policy-dn LDAP attribute. | [optional] 
**Disabled** | Pointer to **bool** | Specifies whether the root user account should be disabled. A disabled account is not permitted to authenticate, nor can it be used as an authorization identity. This is stored in the ds-pwp-account-disabled LDAP attribute. | [optional] 
**AccountActivationTime** | Pointer to **string** | Specifies the time, in generalized time format (e.g., &#39;20160101070000Z&#39;), that the root user account should become active. If an activation time is specified, the user will not be permitted to authenticate, nor can the account be used as an authorization identity, until the activation time has arrived. This is stored in the ds-pwp-account-activation-time LDAP attribute. | [optional] 
**AccountExpirationTime** | Pointer to **string** | Specifies the time, in generalized time format (e.g., &#39;20240101070000Z&#39;), that the root user account should expire. If an expiration time is specified, the user will not be permitted to authenticate, nor can the account be used as an authorization identity, after this time has passed. This is stored in the ds-pwp-account-expiration-time LDAP attribute. | [optional] 
**RequireSecureAuthentication** | Pointer to **bool** | Indicates whether this User must authenticate in a secure manner. When set to \&quot;true\&quot;, the User will only be allowed to authenticate over a secure connection or using a mechanism that does not expose user credentials (e.g., the CRAM-MD5, DIGEST-MD5, and GSSAPI SASL mechanisms). | [optional] 
**RequireSecureConnections** | Pointer to **bool** | Indicates whether this User must be required to communicate with the server over a secure connection. When set to \&quot;true\&quot;, the User will only be allowed to communicate with the server over a secure connection (i.e., using TLS or the StartTLS extended operation). | [optional] 
**AllowedAuthenticationType** | Pointer to **[]string** | Indicates that User should only be allowed to authenticate in certain ways. Allowed values include \&quot;simple\&quot; (to indicate that the user should be allowed to bind using simple authentication) or \&quot;sasl {mech}\&quot; (to indicate that the user should be allowed to bind using the specified SASL mechanism, like \&quot;sasl PLAIN\&quot;). The list of available SASL mechanisms can be retrieved by running \&quot;dsconfig --advanced list-sasl-mechanism-handlers\&quot;. | [optional] 
**AllowedAuthenticationIPAddress** | Pointer to **[]string** | An IPv4 or IPv6 address mask that controls the set of IP addresses from which this User can authenticate to the server. For instance a value of 127.0.0.1 (or ::1 in IPv6) would restricted access only to localhost connections, whereas 10.6.1.* would restrict access to servers on the 10.6.1.* subnet. | [optional] 
**PreferredOTPDeliveryMechanism** | Pointer to **[]string** | Overrides the default settings for the mechanisms (e.g., email or SMS) that are used to deliver one time passwords to Users. | [optional] 
**IsProxyable** | Pointer to [**EnumtopologyAdminUserIsProxyableProp**](EnumtopologyAdminUserIsProxyableProp.md) |  | [optional] 
**IsProxyableByDN** | Pointer to **[]string** | Specifies the DNs of accounts that can proxy as this User using the proxied authorization v1 or v2 control, the intermediate client control, or a SASL mechanism that allows specifying an alternate authorization identity. This property is only applicable if is-proxyable is set to \&quot;allowed\&quot; or \&quot;required\&quot;. | [optional] 
**IsProxyableByGroup** | Pointer to **[]string** | Specifies the DNs of groups whose members can proxy as this User using the proxied authorization v1 or v2 control, the intermediate client control, or a SASL mechanism that allows specifying an alternate authorization identity. This property is only applicable if is-proxyable is set to \&quot;allowed\&quot; or \&quot;required\&quot;. | [optional] 
**IsProxyableByURL** | Pointer to **[]string** | Specifies LDAP URLs of accounts that can proxy as this User using the proxied authorization v1 or v2 control, the intermediate client control, or a SASL mechanism that allows specifying an alternate authorization identity. This property is only applicable if is-proxyable is set to \&quot;allowed\&quot; or \&quot;required\&quot;. | [optional] 
**MayProxyAsDN** | Pointer to **[]string** | This restricts the set of accounts that this User can proxy as to entries with the specified DNs. | [optional] 
**MayProxyAsGroup** | Pointer to **[]string** | This restricts the set of accounts that this User can proxy as to entries that are in the group with the specified DN. | [optional] 
**MayProxyAsURL** | Pointer to **[]string** | This restricts the set of accounts that this User can proxy as to entries that are matched by the specified LDAP URL. | [optional] 

## Methods

### NewAddTopologyAdminUserRequest

`func NewAddTopologyAdminUserRequest(userName string, ) *AddTopologyAdminUserRequest`

NewAddTopologyAdminUserRequest instantiates a new AddTopologyAdminUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTopologyAdminUserRequestWithDefaults

`func NewAddTopologyAdminUserRequestWithDefaults() *AddTopologyAdminUserRequest`

NewAddTopologyAdminUserRequestWithDefaults instantiates a new AddTopologyAdminUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *AddTopologyAdminUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AddTopologyAdminUserRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AddTopologyAdminUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetSchemas

`func (o *AddTopologyAdminUserRequest) GetSchemas() []EnumtopologyAdminUserSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddTopologyAdminUserRequest) GetSchemasOk() (*[]EnumtopologyAdminUserSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddTopologyAdminUserRequest) SetSchemas(v []EnumtopologyAdminUserSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *AddTopologyAdminUserRequest) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetAlternateBindDN

`func (o *AddTopologyAdminUserRequest) GetAlternateBindDN() []string`

GetAlternateBindDN returns the AlternateBindDN field if non-nil, zero value otherwise.

### GetAlternateBindDNOk

`func (o *AddTopologyAdminUserRequest) GetAlternateBindDNOk() (*[]string, bool)`

GetAlternateBindDNOk returns a tuple with the AlternateBindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateBindDN

`func (o *AddTopologyAdminUserRequest) SetAlternateBindDN(v []string)`

SetAlternateBindDN sets AlternateBindDN field to given value.

### HasAlternateBindDN

`func (o *AddTopologyAdminUserRequest) HasAlternateBindDN() bool`

HasAlternateBindDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddTopologyAdminUserRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddTopologyAdminUserRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddTopologyAdminUserRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddTopologyAdminUserRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPassword

`func (o *AddTopologyAdminUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddTopologyAdminUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddTopologyAdminUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddTopologyAdminUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFirstName

`func (o *AddTopologyAdminUserRequest) GetFirstName() []string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddTopologyAdminUserRequest) GetFirstNameOk() (*[]string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddTopologyAdminUserRequest) SetFirstName(v []string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AddTopologyAdminUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AddTopologyAdminUserRequest) GetLastName() []string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddTopologyAdminUserRequest) GetLastNameOk() (*[]string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddTopologyAdminUserRequest) SetLastName(v []string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AddTopologyAdminUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUserID

`func (o *AddTopologyAdminUserRequest) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AddTopologyAdminUserRequest) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AddTopologyAdminUserRequest) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *AddTopologyAdminUserRequest) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetEmailAddress

`func (o *AddTopologyAdminUserRequest) GetEmailAddress() []string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *AddTopologyAdminUserRequest) GetEmailAddressOk() (*[]string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *AddTopologyAdminUserRequest) SetEmailAddress(v []string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *AddTopologyAdminUserRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetWorkTelephoneNumber

`func (o *AddTopologyAdminUserRequest) GetWorkTelephoneNumber() []string`

GetWorkTelephoneNumber returns the WorkTelephoneNumber field if non-nil, zero value otherwise.

### GetWorkTelephoneNumberOk

`func (o *AddTopologyAdminUserRequest) GetWorkTelephoneNumberOk() (*[]string, bool)`

GetWorkTelephoneNumberOk returns a tuple with the WorkTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTelephoneNumber

`func (o *AddTopologyAdminUserRequest) SetWorkTelephoneNumber(v []string)`

SetWorkTelephoneNumber sets WorkTelephoneNumber field to given value.

### HasWorkTelephoneNumber

`func (o *AddTopologyAdminUserRequest) HasWorkTelephoneNumber() bool`

HasWorkTelephoneNumber returns a boolean if a field has been set.

### GetHomeTelephoneNumber

`func (o *AddTopologyAdminUserRequest) GetHomeTelephoneNumber() []string`

GetHomeTelephoneNumber returns the HomeTelephoneNumber field if non-nil, zero value otherwise.

### GetHomeTelephoneNumberOk

`func (o *AddTopologyAdminUserRequest) GetHomeTelephoneNumberOk() (*[]string, bool)`

GetHomeTelephoneNumberOk returns a tuple with the HomeTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeTelephoneNumber

`func (o *AddTopologyAdminUserRequest) SetHomeTelephoneNumber(v []string)`

SetHomeTelephoneNumber sets HomeTelephoneNumber field to given value.

### HasHomeTelephoneNumber

`func (o *AddTopologyAdminUserRequest) HasHomeTelephoneNumber() bool`

HasHomeTelephoneNumber returns a boolean if a field has been set.

### GetMobileTelephoneNumber

`func (o *AddTopologyAdminUserRequest) GetMobileTelephoneNumber() []string`

GetMobileTelephoneNumber returns the MobileTelephoneNumber field if non-nil, zero value otherwise.

### GetMobileTelephoneNumberOk

`func (o *AddTopologyAdminUserRequest) GetMobileTelephoneNumberOk() (*[]string, bool)`

GetMobileTelephoneNumberOk returns a tuple with the MobileTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileTelephoneNumber

`func (o *AddTopologyAdminUserRequest) SetMobileTelephoneNumber(v []string)`

SetMobileTelephoneNumber sets MobileTelephoneNumber field to given value.

### HasMobileTelephoneNumber

`func (o *AddTopologyAdminUserRequest) HasMobileTelephoneNumber() bool`

HasMobileTelephoneNumber returns a boolean if a field has been set.

### GetPagerTelephoneNumber

`func (o *AddTopologyAdminUserRequest) GetPagerTelephoneNumber() []string`

GetPagerTelephoneNumber returns the PagerTelephoneNumber field if non-nil, zero value otherwise.

### GetPagerTelephoneNumberOk

`func (o *AddTopologyAdminUserRequest) GetPagerTelephoneNumberOk() (*[]string, bool)`

GetPagerTelephoneNumberOk returns a tuple with the PagerTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerTelephoneNumber

`func (o *AddTopologyAdminUserRequest) SetPagerTelephoneNumber(v []string)`

SetPagerTelephoneNumber sets PagerTelephoneNumber field to given value.

### HasPagerTelephoneNumber

`func (o *AddTopologyAdminUserRequest) HasPagerTelephoneNumber() bool`

HasPagerTelephoneNumber returns a boolean if a field has been set.

### GetInheritDefaultRootPrivileges

`func (o *AddTopologyAdminUserRequest) GetInheritDefaultRootPrivileges() bool`

GetInheritDefaultRootPrivileges returns the InheritDefaultRootPrivileges field if non-nil, zero value otherwise.

### GetInheritDefaultRootPrivilegesOk

`func (o *AddTopologyAdminUserRequest) GetInheritDefaultRootPrivilegesOk() (*bool, bool)`

GetInheritDefaultRootPrivilegesOk returns a tuple with the InheritDefaultRootPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritDefaultRootPrivileges

`func (o *AddTopologyAdminUserRequest) SetInheritDefaultRootPrivileges(v bool)`

SetInheritDefaultRootPrivileges sets InheritDefaultRootPrivileges field to given value.

### HasInheritDefaultRootPrivileges

`func (o *AddTopologyAdminUserRequest) HasInheritDefaultRootPrivileges() bool`

HasInheritDefaultRootPrivileges returns a boolean if a field has been set.

### GetPrivilege

`func (o *AddTopologyAdminUserRequest) GetPrivilege() []EnumtopologyAdminUserPrivilegeProp`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *AddTopologyAdminUserRequest) GetPrivilegeOk() (*[]EnumtopologyAdminUserPrivilegeProp, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *AddTopologyAdminUserRequest) SetPrivilege(v []EnumtopologyAdminUserPrivilegeProp)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *AddTopologyAdminUserRequest) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetSearchResultEntryLimit

`func (o *AddTopologyAdminUserRequest) GetSearchResultEntryLimit() int64`

GetSearchResultEntryLimit returns the SearchResultEntryLimit field if non-nil, zero value otherwise.

### GetSearchResultEntryLimitOk

`func (o *AddTopologyAdminUserRequest) GetSearchResultEntryLimitOk() (*int64, bool)`

GetSearchResultEntryLimitOk returns a tuple with the SearchResultEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResultEntryLimit

`func (o *AddTopologyAdminUserRequest) SetSearchResultEntryLimit(v int64)`

SetSearchResultEntryLimit sets SearchResultEntryLimit field to given value.

### HasSearchResultEntryLimit

`func (o *AddTopologyAdminUserRequest) HasSearchResultEntryLimit() bool`

HasSearchResultEntryLimit returns a boolean if a field has been set.

### GetTimeLimitSeconds

`func (o *AddTopologyAdminUserRequest) GetTimeLimitSeconds() int64`

GetTimeLimitSeconds returns the TimeLimitSeconds field if non-nil, zero value otherwise.

### GetTimeLimitSecondsOk

`func (o *AddTopologyAdminUserRequest) GetTimeLimitSecondsOk() (*int64, bool)`

GetTimeLimitSecondsOk returns a tuple with the TimeLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimitSeconds

`func (o *AddTopologyAdminUserRequest) SetTimeLimitSeconds(v int64)`

SetTimeLimitSeconds sets TimeLimitSeconds field to given value.

### HasTimeLimitSeconds

`func (o *AddTopologyAdminUserRequest) HasTimeLimitSeconds() bool`

HasTimeLimitSeconds returns a boolean if a field has been set.

### GetLookThroughEntryLimit

`func (o *AddTopologyAdminUserRequest) GetLookThroughEntryLimit() int64`

GetLookThroughEntryLimit returns the LookThroughEntryLimit field if non-nil, zero value otherwise.

### GetLookThroughEntryLimitOk

`func (o *AddTopologyAdminUserRequest) GetLookThroughEntryLimitOk() (*int64, bool)`

GetLookThroughEntryLimitOk returns a tuple with the LookThroughEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookThroughEntryLimit

`func (o *AddTopologyAdminUserRequest) SetLookThroughEntryLimit(v int64)`

SetLookThroughEntryLimit sets LookThroughEntryLimit field to given value.

### HasLookThroughEntryLimit

`func (o *AddTopologyAdminUserRequest) HasLookThroughEntryLimit() bool`

HasLookThroughEntryLimit returns a boolean if a field has been set.

### GetIdleTimeLimitSeconds

`func (o *AddTopologyAdminUserRequest) GetIdleTimeLimitSeconds() int64`

GetIdleTimeLimitSeconds returns the IdleTimeLimitSeconds field if non-nil, zero value otherwise.

### GetIdleTimeLimitSecondsOk

`func (o *AddTopologyAdminUserRequest) GetIdleTimeLimitSecondsOk() (*int64, bool)`

GetIdleTimeLimitSecondsOk returns a tuple with the IdleTimeLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeLimitSeconds

`func (o *AddTopologyAdminUserRequest) SetIdleTimeLimitSeconds(v int64)`

SetIdleTimeLimitSeconds sets IdleTimeLimitSeconds field to given value.

### HasIdleTimeLimitSeconds

`func (o *AddTopologyAdminUserRequest) HasIdleTimeLimitSeconds() bool`

HasIdleTimeLimitSeconds returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *AddTopologyAdminUserRequest) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *AddTopologyAdminUserRequest) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *AddTopologyAdminUserRequest) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *AddTopologyAdminUserRequest) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetDisabled

`func (o *AddTopologyAdminUserRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AddTopologyAdminUserRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AddTopologyAdminUserRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AddTopologyAdminUserRequest) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetAccountActivationTime

`func (o *AddTopologyAdminUserRequest) GetAccountActivationTime() string`

GetAccountActivationTime returns the AccountActivationTime field if non-nil, zero value otherwise.

### GetAccountActivationTimeOk

`func (o *AddTopologyAdminUserRequest) GetAccountActivationTimeOk() (*string, bool)`

GetAccountActivationTimeOk returns a tuple with the AccountActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountActivationTime

`func (o *AddTopologyAdminUserRequest) SetAccountActivationTime(v string)`

SetAccountActivationTime sets AccountActivationTime field to given value.

### HasAccountActivationTime

`func (o *AddTopologyAdminUserRequest) HasAccountActivationTime() bool`

HasAccountActivationTime returns a boolean if a field has been set.

### GetAccountExpirationTime

`func (o *AddTopologyAdminUserRequest) GetAccountExpirationTime() string`

GetAccountExpirationTime returns the AccountExpirationTime field if non-nil, zero value otherwise.

### GetAccountExpirationTimeOk

`func (o *AddTopologyAdminUserRequest) GetAccountExpirationTimeOk() (*string, bool)`

GetAccountExpirationTimeOk returns a tuple with the AccountExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationTime

`func (o *AddTopologyAdminUserRequest) SetAccountExpirationTime(v string)`

SetAccountExpirationTime sets AccountExpirationTime field to given value.

### HasAccountExpirationTime

`func (o *AddTopologyAdminUserRequest) HasAccountExpirationTime() bool`

HasAccountExpirationTime returns a boolean if a field has been set.

### GetRequireSecureAuthentication

`func (o *AddTopologyAdminUserRequest) GetRequireSecureAuthentication() bool`

GetRequireSecureAuthentication returns the RequireSecureAuthentication field if non-nil, zero value otherwise.

### GetRequireSecureAuthenticationOk

`func (o *AddTopologyAdminUserRequest) GetRequireSecureAuthenticationOk() (*bool, bool)`

GetRequireSecureAuthenticationOk returns a tuple with the RequireSecureAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSecureAuthentication

`func (o *AddTopologyAdminUserRequest) SetRequireSecureAuthentication(v bool)`

SetRequireSecureAuthentication sets RequireSecureAuthentication field to given value.

### HasRequireSecureAuthentication

`func (o *AddTopologyAdminUserRequest) HasRequireSecureAuthentication() bool`

HasRequireSecureAuthentication returns a boolean if a field has been set.

### GetRequireSecureConnections

`func (o *AddTopologyAdminUserRequest) GetRequireSecureConnections() bool`

GetRequireSecureConnections returns the RequireSecureConnections field if non-nil, zero value otherwise.

### GetRequireSecureConnectionsOk

`func (o *AddTopologyAdminUserRequest) GetRequireSecureConnectionsOk() (*bool, bool)`

GetRequireSecureConnectionsOk returns a tuple with the RequireSecureConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSecureConnections

`func (o *AddTopologyAdminUserRequest) SetRequireSecureConnections(v bool)`

SetRequireSecureConnections sets RequireSecureConnections field to given value.

### HasRequireSecureConnections

`func (o *AddTopologyAdminUserRequest) HasRequireSecureConnections() bool`

HasRequireSecureConnections returns a boolean if a field has been set.

### GetAllowedAuthenticationType

`func (o *AddTopologyAdminUserRequest) GetAllowedAuthenticationType() []string`

GetAllowedAuthenticationType returns the AllowedAuthenticationType field if non-nil, zero value otherwise.

### GetAllowedAuthenticationTypeOk

`func (o *AddTopologyAdminUserRequest) GetAllowedAuthenticationTypeOk() (*[]string, bool)`

GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationType

`func (o *AddTopologyAdminUserRequest) SetAllowedAuthenticationType(v []string)`

SetAllowedAuthenticationType sets AllowedAuthenticationType field to given value.

### HasAllowedAuthenticationType

`func (o *AddTopologyAdminUserRequest) HasAllowedAuthenticationType() bool`

HasAllowedAuthenticationType returns a boolean if a field has been set.

### GetAllowedAuthenticationIPAddress

`func (o *AddTopologyAdminUserRequest) GetAllowedAuthenticationIPAddress() []string`

GetAllowedAuthenticationIPAddress returns the AllowedAuthenticationIPAddress field if non-nil, zero value otherwise.

### GetAllowedAuthenticationIPAddressOk

`func (o *AddTopologyAdminUserRequest) GetAllowedAuthenticationIPAddressOk() (*[]string, bool)`

GetAllowedAuthenticationIPAddressOk returns a tuple with the AllowedAuthenticationIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationIPAddress

`func (o *AddTopologyAdminUserRequest) SetAllowedAuthenticationIPAddress(v []string)`

SetAllowedAuthenticationIPAddress sets AllowedAuthenticationIPAddress field to given value.

### HasAllowedAuthenticationIPAddress

`func (o *AddTopologyAdminUserRequest) HasAllowedAuthenticationIPAddress() bool`

HasAllowedAuthenticationIPAddress returns a boolean if a field has been set.

### GetPreferredOTPDeliveryMechanism

`func (o *AddTopologyAdminUserRequest) GetPreferredOTPDeliveryMechanism() []string`

GetPreferredOTPDeliveryMechanism returns the PreferredOTPDeliveryMechanism field if non-nil, zero value otherwise.

### GetPreferredOTPDeliveryMechanismOk

`func (o *AddTopologyAdminUserRequest) GetPreferredOTPDeliveryMechanismOk() (*[]string, bool)`

GetPreferredOTPDeliveryMechanismOk returns a tuple with the PreferredOTPDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOTPDeliveryMechanism

`func (o *AddTopologyAdminUserRequest) SetPreferredOTPDeliveryMechanism(v []string)`

SetPreferredOTPDeliveryMechanism sets PreferredOTPDeliveryMechanism field to given value.

### HasPreferredOTPDeliveryMechanism

`func (o *AddTopologyAdminUserRequest) HasPreferredOTPDeliveryMechanism() bool`

HasPreferredOTPDeliveryMechanism returns a boolean if a field has been set.

### GetIsProxyable

`func (o *AddTopologyAdminUserRequest) GetIsProxyable() EnumtopologyAdminUserIsProxyableProp`

GetIsProxyable returns the IsProxyable field if non-nil, zero value otherwise.

### GetIsProxyableOk

`func (o *AddTopologyAdminUserRequest) GetIsProxyableOk() (*EnumtopologyAdminUserIsProxyableProp, bool)`

GetIsProxyableOk returns a tuple with the IsProxyable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyable

`func (o *AddTopologyAdminUserRequest) SetIsProxyable(v EnumtopologyAdminUserIsProxyableProp)`

SetIsProxyable sets IsProxyable field to given value.

### HasIsProxyable

`func (o *AddTopologyAdminUserRequest) HasIsProxyable() bool`

HasIsProxyable returns a boolean if a field has been set.

### GetIsProxyableByDN

`func (o *AddTopologyAdminUserRequest) GetIsProxyableByDN() []string`

GetIsProxyableByDN returns the IsProxyableByDN field if non-nil, zero value otherwise.

### GetIsProxyableByDNOk

`func (o *AddTopologyAdminUserRequest) GetIsProxyableByDNOk() (*[]string, bool)`

GetIsProxyableByDNOk returns a tuple with the IsProxyableByDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyableByDN

`func (o *AddTopologyAdminUserRequest) SetIsProxyableByDN(v []string)`

SetIsProxyableByDN sets IsProxyableByDN field to given value.

### HasIsProxyableByDN

`func (o *AddTopologyAdminUserRequest) HasIsProxyableByDN() bool`

HasIsProxyableByDN returns a boolean if a field has been set.

### GetIsProxyableByGroup

`func (o *AddTopologyAdminUserRequest) GetIsProxyableByGroup() []string`

GetIsProxyableByGroup returns the IsProxyableByGroup field if non-nil, zero value otherwise.

### GetIsProxyableByGroupOk

`func (o *AddTopologyAdminUserRequest) GetIsProxyableByGroupOk() (*[]string, bool)`

GetIsProxyableByGroupOk returns a tuple with the IsProxyableByGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyableByGroup

`func (o *AddTopologyAdminUserRequest) SetIsProxyableByGroup(v []string)`

SetIsProxyableByGroup sets IsProxyableByGroup field to given value.

### HasIsProxyableByGroup

`func (o *AddTopologyAdminUserRequest) HasIsProxyableByGroup() bool`

HasIsProxyableByGroup returns a boolean if a field has been set.

### GetIsProxyableByURL

`func (o *AddTopologyAdminUserRequest) GetIsProxyableByURL() []string`

GetIsProxyableByURL returns the IsProxyableByURL field if non-nil, zero value otherwise.

### GetIsProxyableByURLOk

`func (o *AddTopologyAdminUserRequest) GetIsProxyableByURLOk() (*[]string, bool)`

GetIsProxyableByURLOk returns a tuple with the IsProxyableByURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyableByURL

`func (o *AddTopologyAdminUserRequest) SetIsProxyableByURL(v []string)`

SetIsProxyableByURL sets IsProxyableByURL field to given value.

### HasIsProxyableByURL

`func (o *AddTopologyAdminUserRequest) HasIsProxyableByURL() bool`

HasIsProxyableByURL returns a boolean if a field has been set.

### GetMayProxyAsDN

`func (o *AddTopologyAdminUserRequest) GetMayProxyAsDN() []string`

GetMayProxyAsDN returns the MayProxyAsDN field if non-nil, zero value otherwise.

### GetMayProxyAsDNOk

`func (o *AddTopologyAdminUserRequest) GetMayProxyAsDNOk() (*[]string, bool)`

GetMayProxyAsDNOk returns a tuple with the MayProxyAsDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayProxyAsDN

`func (o *AddTopologyAdminUserRequest) SetMayProxyAsDN(v []string)`

SetMayProxyAsDN sets MayProxyAsDN field to given value.

### HasMayProxyAsDN

`func (o *AddTopologyAdminUserRequest) HasMayProxyAsDN() bool`

HasMayProxyAsDN returns a boolean if a field has been set.

### GetMayProxyAsGroup

`func (o *AddTopologyAdminUserRequest) GetMayProxyAsGroup() []string`

GetMayProxyAsGroup returns the MayProxyAsGroup field if non-nil, zero value otherwise.

### GetMayProxyAsGroupOk

`func (o *AddTopologyAdminUserRequest) GetMayProxyAsGroupOk() (*[]string, bool)`

GetMayProxyAsGroupOk returns a tuple with the MayProxyAsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayProxyAsGroup

`func (o *AddTopologyAdminUserRequest) SetMayProxyAsGroup(v []string)`

SetMayProxyAsGroup sets MayProxyAsGroup field to given value.

### HasMayProxyAsGroup

`func (o *AddTopologyAdminUserRequest) HasMayProxyAsGroup() bool`

HasMayProxyAsGroup returns a boolean if a field has been set.

### GetMayProxyAsURL

`func (o *AddTopologyAdminUserRequest) GetMayProxyAsURL() []string`

GetMayProxyAsURL returns the MayProxyAsURL field if non-nil, zero value otherwise.

### GetMayProxyAsURLOk

`func (o *AddTopologyAdminUserRequest) GetMayProxyAsURLOk() (*[]string, bool)`

GetMayProxyAsURLOk returns a tuple with the MayProxyAsURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayProxyAsURL

`func (o *AddTopologyAdminUserRequest) SetMayProxyAsURL(v []string)`

SetMayProxyAsURL sets MayProxyAsURL field to given value.

### HasMayProxyAsURL

`func (o *AddTopologyAdminUserRequest) HasMayProxyAsURL() bool`

HasMayProxyAsURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


