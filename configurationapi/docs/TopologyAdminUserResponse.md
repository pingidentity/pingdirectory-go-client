# TopologyAdminUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Topology Admin User | 
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
**InheritDefaultRootPrivileges** | **bool** | Indicates whether this User should be automatically granted the set of privileges defined in the default-root-privilege-name property of the Root DN configuration object. | 
**Privilege** | Pointer to [**[]EnumtopologyAdminUserPrivilegeProp**](EnumtopologyAdminUserPrivilegeProp.md) | Privileges that are either explicitly granted or revoked from the root user. Privileges can be revoked by including a minus sign (-) before the privilege name. This is stored in the ds-privilege-name LDAP attribute. | [optional] 
**SearchResultEntryLimit** | **int32** | Specifies the maximum number of entries that the server may return to the user in response to any single search request. A value of 0 indicates no limit should be enforced. This is stored in the ds-rlim-size-limit LDAP attribute. | 
**TimeLimitSeconds** | **int32** | Specifies the maximum length of time (in seconds) that the server may spend processing any single search request. A value of 0 indicates no limit should be enforced. This is stored in the ds-rlim-time-limit LDAP attribute. | 
**LookThroughEntryLimit** | **int32** | Specifies the maximum number of candidate entries that the server may examine in the course of processing any single search request. A value of 0 indicates no limit should be enforced. This is stored in the ds-rlim-lookthrough-limit LDAP attribute. | 
**IdleTimeLimitSeconds** | **int32** | Specifies the maximum length of time (in seconds) that a connection authenticated as this user may remain established without issuing any requests. A value of 0 indicates no limit should be enforced. This is stored in the ds-rlim-idle-time-limit LDAP attribute. | 
**PasswordPolicy** | **string** | Specifies the password policy for the user. This is stored in the ds-pwp-password-policy-dn LDAP attribute. | 
**Disabled** | Pointer to **bool** | Specifies whether the root user account should be disabled. A disabled account is not permitted to authenticate, nor can it be used as an authorization identity. This is stored in the ds-pwp-account-disabled LDAP attribute. | [optional] 
**AccountActivationTime** | Pointer to **string** | Specifies the time, in generalized time format (e.g., &#39;20160101070000Z&#39;), that the root user account should become active. If an activation time is specified, the user will not be permitted to authenticate, nor can the account be used as an authorization identity, until the activation time has arrived. This is stored in the ds-pwp-account-activation-time LDAP attribute. | [optional] 
**AccountExpirationTime** | Pointer to **string** | Specifies the time, in generalized time format (e.g., &#39;20240101070000Z&#39;), that the root user account should expire. If an expiration time is specified, the user will not be permitted to authenticate, nor can the account be used as an authorization identity, after this time has passed. This is stored in the ds-pwp-account-expiration-time LDAP attribute. | [optional] 
**RequireSecureAuthentication** | **bool** | Indicates whether this User must authenticate in a secure manner. When set to \&quot;true\&quot;, the User will only be allowed to authenticate over a secure connection or using a mechanism that does not expose user credentials (e.g., the CRAM-MD5, DIGEST-MD5, and GSSAPI SASL mechanisms). | 
**RequireSecureConnections** | **bool** | Indicates whether this User must be required to communicate with the server over a secure connection. When set to \&quot;true\&quot;, the User will only be allowed to communicate with the server over a secure connection (i.e., using TLS or the StartTLS extended operation). | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewTopologyAdminUserResponse

`func NewTopologyAdminUserResponse(id string, inheritDefaultRootPrivileges bool, searchResultEntryLimit int32, timeLimitSeconds int32, lookThroughEntryLimit int32, idleTimeLimitSeconds int32, passwordPolicy string, requireSecureAuthentication bool, requireSecureConnections bool, ) *TopologyAdminUserResponse`

NewTopologyAdminUserResponse instantiates a new TopologyAdminUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyAdminUserResponseWithDefaults

`func NewTopologyAdminUserResponseWithDefaults() *TopologyAdminUserResponse`

NewTopologyAdminUserResponseWithDefaults instantiates a new TopologyAdminUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopologyAdminUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopologyAdminUserResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopologyAdminUserResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *TopologyAdminUserResponse) GetSchemas() []EnumtopologyAdminUserSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *TopologyAdminUserResponse) GetSchemasOk() (*[]EnumtopologyAdminUserSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *TopologyAdminUserResponse) SetSchemas(v []EnumtopologyAdminUserSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *TopologyAdminUserResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetAlternateBindDN

`func (o *TopologyAdminUserResponse) GetAlternateBindDN() []string`

GetAlternateBindDN returns the AlternateBindDN field if non-nil, zero value otherwise.

### GetAlternateBindDNOk

`func (o *TopologyAdminUserResponse) GetAlternateBindDNOk() (*[]string, bool)`

GetAlternateBindDNOk returns a tuple with the AlternateBindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateBindDN

`func (o *TopologyAdminUserResponse) SetAlternateBindDN(v []string)`

SetAlternateBindDN sets AlternateBindDN field to given value.

### HasAlternateBindDN

`func (o *TopologyAdminUserResponse) HasAlternateBindDN() bool`

HasAlternateBindDN returns a boolean if a field has been set.

### GetDescription

`func (o *TopologyAdminUserResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TopologyAdminUserResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TopologyAdminUserResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TopologyAdminUserResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPassword

`func (o *TopologyAdminUserResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TopologyAdminUserResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TopologyAdminUserResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TopologyAdminUserResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFirstName

`func (o *TopologyAdminUserResponse) GetFirstName() []string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TopologyAdminUserResponse) GetFirstNameOk() (*[]string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TopologyAdminUserResponse) SetFirstName(v []string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *TopologyAdminUserResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *TopologyAdminUserResponse) GetLastName() []string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TopologyAdminUserResponse) GetLastNameOk() (*[]string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TopologyAdminUserResponse) SetLastName(v []string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *TopologyAdminUserResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUserID

`func (o *TopologyAdminUserResponse) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *TopologyAdminUserResponse) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *TopologyAdminUserResponse) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *TopologyAdminUserResponse) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetEmailAddress

`func (o *TopologyAdminUserResponse) GetEmailAddress() []string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *TopologyAdminUserResponse) GetEmailAddressOk() (*[]string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *TopologyAdminUserResponse) SetEmailAddress(v []string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *TopologyAdminUserResponse) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetWorkTelephoneNumber

`func (o *TopologyAdminUserResponse) GetWorkTelephoneNumber() []string`

GetWorkTelephoneNumber returns the WorkTelephoneNumber field if non-nil, zero value otherwise.

### GetWorkTelephoneNumberOk

`func (o *TopologyAdminUserResponse) GetWorkTelephoneNumberOk() (*[]string, bool)`

GetWorkTelephoneNumberOk returns a tuple with the WorkTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTelephoneNumber

`func (o *TopologyAdminUserResponse) SetWorkTelephoneNumber(v []string)`

SetWorkTelephoneNumber sets WorkTelephoneNumber field to given value.

### HasWorkTelephoneNumber

`func (o *TopologyAdminUserResponse) HasWorkTelephoneNumber() bool`

HasWorkTelephoneNumber returns a boolean if a field has been set.

### GetHomeTelephoneNumber

`func (o *TopologyAdminUserResponse) GetHomeTelephoneNumber() []string`

GetHomeTelephoneNumber returns the HomeTelephoneNumber field if non-nil, zero value otherwise.

### GetHomeTelephoneNumberOk

`func (o *TopologyAdminUserResponse) GetHomeTelephoneNumberOk() (*[]string, bool)`

GetHomeTelephoneNumberOk returns a tuple with the HomeTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeTelephoneNumber

`func (o *TopologyAdminUserResponse) SetHomeTelephoneNumber(v []string)`

SetHomeTelephoneNumber sets HomeTelephoneNumber field to given value.

### HasHomeTelephoneNumber

`func (o *TopologyAdminUserResponse) HasHomeTelephoneNumber() bool`

HasHomeTelephoneNumber returns a boolean if a field has been set.

### GetMobileTelephoneNumber

`func (o *TopologyAdminUserResponse) GetMobileTelephoneNumber() []string`

GetMobileTelephoneNumber returns the MobileTelephoneNumber field if non-nil, zero value otherwise.

### GetMobileTelephoneNumberOk

`func (o *TopologyAdminUserResponse) GetMobileTelephoneNumberOk() (*[]string, bool)`

GetMobileTelephoneNumberOk returns a tuple with the MobileTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileTelephoneNumber

`func (o *TopologyAdminUserResponse) SetMobileTelephoneNumber(v []string)`

SetMobileTelephoneNumber sets MobileTelephoneNumber field to given value.

### HasMobileTelephoneNumber

`func (o *TopologyAdminUserResponse) HasMobileTelephoneNumber() bool`

HasMobileTelephoneNumber returns a boolean if a field has been set.

### GetPagerTelephoneNumber

`func (o *TopologyAdminUserResponse) GetPagerTelephoneNumber() []string`

GetPagerTelephoneNumber returns the PagerTelephoneNumber field if non-nil, zero value otherwise.

### GetPagerTelephoneNumberOk

`func (o *TopologyAdminUserResponse) GetPagerTelephoneNumberOk() (*[]string, bool)`

GetPagerTelephoneNumberOk returns a tuple with the PagerTelephoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerTelephoneNumber

`func (o *TopologyAdminUserResponse) SetPagerTelephoneNumber(v []string)`

SetPagerTelephoneNumber sets PagerTelephoneNumber field to given value.

### HasPagerTelephoneNumber

`func (o *TopologyAdminUserResponse) HasPagerTelephoneNumber() bool`

HasPagerTelephoneNumber returns a boolean if a field has been set.

### GetInheritDefaultRootPrivileges

`func (o *TopologyAdminUserResponse) GetInheritDefaultRootPrivileges() bool`

GetInheritDefaultRootPrivileges returns the InheritDefaultRootPrivileges field if non-nil, zero value otherwise.

### GetInheritDefaultRootPrivilegesOk

`func (o *TopologyAdminUserResponse) GetInheritDefaultRootPrivilegesOk() (*bool, bool)`

GetInheritDefaultRootPrivilegesOk returns a tuple with the InheritDefaultRootPrivileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritDefaultRootPrivileges

`func (o *TopologyAdminUserResponse) SetInheritDefaultRootPrivileges(v bool)`

SetInheritDefaultRootPrivileges sets InheritDefaultRootPrivileges field to given value.


### GetPrivilege

`func (o *TopologyAdminUserResponse) GetPrivilege() []EnumtopologyAdminUserPrivilegeProp`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *TopologyAdminUserResponse) GetPrivilegeOk() (*[]EnumtopologyAdminUserPrivilegeProp, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *TopologyAdminUserResponse) SetPrivilege(v []EnumtopologyAdminUserPrivilegeProp)`

SetPrivilege sets Privilege field to given value.

### HasPrivilege

`func (o *TopologyAdminUserResponse) HasPrivilege() bool`

HasPrivilege returns a boolean if a field has been set.

### GetSearchResultEntryLimit

`func (o *TopologyAdminUserResponse) GetSearchResultEntryLimit() int32`

GetSearchResultEntryLimit returns the SearchResultEntryLimit field if non-nil, zero value otherwise.

### GetSearchResultEntryLimitOk

`func (o *TopologyAdminUserResponse) GetSearchResultEntryLimitOk() (*int32, bool)`

GetSearchResultEntryLimitOk returns a tuple with the SearchResultEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResultEntryLimit

`func (o *TopologyAdminUserResponse) SetSearchResultEntryLimit(v int32)`

SetSearchResultEntryLimit sets SearchResultEntryLimit field to given value.


### GetTimeLimitSeconds

`func (o *TopologyAdminUserResponse) GetTimeLimitSeconds() int32`

GetTimeLimitSeconds returns the TimeLimitSeconds field if non-nil, zero value otherwise.

### GetTimeLimitSecondsOk

`func (o *TopologyAdminUserResponse) GetTimeLimitSecondsOk() (*int32, bool)`

GetTimeLimitSecondsOk returns a tuple with the TimeLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimitSeconds

`func (o *TopologyAdminUserResponse) SetTimeLimitSeconds(v int32)`

SetTimeLimitSeconds sets TimeLimitSeconds field to given value.


### GetLookThroughEntryLimit

`func (o *TopologyAdminUserResponse) GetLookThroughEntryLimit() int32`

GetLookThroughEntryLimit returns the LookThroughEntryLimit field if non-nil, zero value otherwise.

### GetLookThroughEntryLimitOk

`func (o *TopologyAdminUserResponse) GetLookThroughEntryLimitOk() (*int32, bool)`

GetLookThroughEntryLimitOk returns a tuple with the LookThroughEntryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookThroughEntryLimit

`func (o *TopologyAdminUserResponse) SetLookThroughEntryLimit(v int32)`

SetLookThroughEntryLimit sets LookThroughEntryLimit field to given value.


### GetIdleTimeLimitSeconds

`func (o *TopologyAdminUserResponse) GetIdleTimeLimitSeconds() int32`

GetIdleTimeLimitSeconds returns the IdleTimeLimitSeconds field if non-nil, zero value otherwise.

### GetIdleTimeLimitSecondsOk

`func (o *TopologyAdminUserResponse) GetIdleTimeLimitSecondsOk() (*int32, bool)`

GetIdleTimeLimitSecondsOk returns a tuple with the IdleTimeLimitSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleTimeLimitSeconds

`func (o *TopologyAdminUserResponse) SetIdleTimeLimitSeconds(v int32)`

SetIdleTimeLimitSeconds sets IdleTimeLimitSeconds field to given value.


### GetPasswordPolicy

`func (o *TopologyAdminUserResponse) GetPasswordPolicy() string`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *TopologyAdminUserResponse) GetPasswordPolicyOk() (*string, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *TopologyAdminUserResponse) SetPasswordPolicy(v string)`

SetPasswordPolicy sets PasswordPolicy field to given value.


### GetDisabled

`func (o *TopologyAdminUserResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *TopologyAdminUserResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *TopologyAdminUserResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *TopologyAdminUserResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetAccountActivationTime

`func (o *TopologyAdminUserResponse) GetAccountActivationTime() string`

GetAccountActivationTime returns the AccountActivationTime field if non-nil, zero value otherwise.

### GetAccountActivationTimeOk

`func (o *TopologyAdminUserResponse) GetAccountActivationTimeOk() (*string, bool)`

GetAccountActivationTimeOk returns a tuple with the AccountActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountActivationTime

`func (o *TopologyAdminUserResponse) SetAccountActivationTime(v string)`

SetAccountActivationTime sets AccountActivationTime field to given value.

### HasAccountActivationTime

`func (o *TopologyAdminUserResponse) HasAccountActivationTime() bool`

HasAccountActivationTime returns a boolean if a field has been set.

### GetAccountExpirationTime

`func (o *TopologyAdminUserResponse) GetAccountExpirationTime() string`

GetAccountExpirationTime returns the AccountExpirationTime field if non-nil, zero value otherwise.

### GetAccountExpirationTimeOk

`func (o *TopologyAdminUserResponse) GetAccountExpirationTimeOk() (*string, bool)`

GetAccountExpirationTimeOk returns a tuple with the AccountExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountExpirationTime

`func (o *TopologyAdminUserResponse) SetAccountExpirationTime(v string)`

SetAccountExpirationTime sets AccountExpirationTime field to given value.

### HasAccountExpirationTime

`func (o *TopologyAdminUserResponse) HasAccountExpirationTime() bool`

HasAccountExpirationTime returns a boolean if a field has been set.

### GetRequireSecureAuthentication

`func (o *TopologyAdminUserResponse) GetRequireSecureAuthentication() bool`

GetRequireSecureAuthentication returns the RequireSecureAuthentication field if non-nil, zero value otherwise.

### GetRequireSecureAuthenticationOk

`func (o *TopologyAdminUserResponse) GetRequireSecureAuthenticationOk() (*bool, bool)`

GetRequireSecureAuthenticationOk returns a tuple with the RequireSecureAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSecureAuthentication

`func (o *TopologyAdminUserResponse) SetRequireSecureAuthentication(v bool)`

SetRequireSecureAuthentication sets RequireSecureAuthentication field to given value.


### GetRequireSecureConnections

`func (o *TopologyAdminUserResponse) GetRequireSecureConnections() bool`

GetRequireSecureConnections returns the RequireSecureConnections field if non-nil, zero value otherwise.

### GetRequireSecureConnectionsOk

`func (o *TopologyAdminUserResponse) GetRequireSecureConnectionsOk() (*bool, bool)`

GetRequireSecureConnectionsOk returns a tuple with the RequireSecureConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSecureConnections

`func (o *TopologyAdminUserResponse) SetRequireSecureConnections(v bool)`

SetRequireSecureConnections sets RequireSecureConnections field to given value.


### GetAllowedAuthenticationType

`func (o *TopologyAdminUserResponse) GetAllowedAuthenticationType() []string`

GetAllowedAuthenticationType returns the AllowedAuthenticationType field if non-nil, zero value otherwise.

### GetAllowedAuthenticationTypeOk

`func (o *TopologyAdminUserResponse) GetAllowedAuthenticationTypeOk() (*[]string, bool)`

GetAllowedAuthenticationTypeOk returns a tuple with the AllowedAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationType

`func (o *TopologyAdminUserResponse) SetAllowedAuthenticationType(v []string)`

SetAllowedAuthenticationType sets AllowedAuthenticationType field to given value.

### HasAllowedAuthenticationType

`func (o *TopologyAdminUserResponse) HasAllowedAuthenticationType() bool`

HasAllowedAuthenticationType returns a boolean if a field has been set.

### GetAllowedAuthenticationIPAddress

`func (o *TopologyAdminUserResponse) GetAllowedAuthenticationIPAddress() []string`

GetAllowedAuthenticationIPAddress returns the AllowedAuthenticationIPAddress field if non-nil, zero value otherwise.

### GetAllowedAuthenticationIPAddressOk

`func (o *TopologyAdminUserResponse) GetAllowedAuthenticationIPAddressOk() (*[]string, bool)`

GetAllowedAuthenticationIPAddressOk returns a tuple with the AllowedAuthenticationIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticationIPAddress

`func (o *TopologyAdminUserResponse) SetAllowedAuthenticationIPAddress(v []string)`

SetAllowedAuthenticationIPAddress sets AllowedAuthenticationIPAddress field to given value.

### HasAllowedAuthenticationIPAddress

`func (o *TopologyAdminUserResponse) HasAllowedAuthenticationIPAddress() bool`

HasAllowedAuthenticationIPAddress returns a boolean if a field has been set.

### GetPreferredOTPDeliveryMechanism

`func (o *TopologyAdminUserResponse) GetPreferredOTPDeliveryMechanism() []string`

GetPreferredOTPDeliveryMechanism returns the PreferredOTPDeliveryMechanism field if non-nil, zero value otherwise.

### GetPreferredOTPDeliveryMechanismOk

`func (o *TopologyAdminUserResponse) GetPreferredOTPDeliveryMechanismOk() (*[]string, bool)`

GetPreferredOTPDeliveryMechanismOk returns a tuple with the PreferredOTPDeliveryMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredOTPDeliveryMechanism

`func (o *TopologyAdminUserResponse) SetPreferredOTPDeliveryMechanism(v []string)`

SetPreferredOTPDeliveryMechanism sets PreferredOTPDeliveryMechanism field to given value.

### HasPreferredOTPDeliveryMechanism

`func (o *TopologyAdminUserResponse) HasPreferredOTPDeliveryMechanism() bool`

HasPreferredOTPDeliveryMechanism returns a boolean if a field has been set.

### GetIsProxyable

`func (o *TopologyAdminUserResponse) GetIsProxyable() EnumtopologyAdminUserIsProxyableProp`

GetIsProxyable returns the IsProxyable field if non-nil, zero value otherwise.

### GetIsProxyableOk

`func (o *TopologyAdminUserResponse) GetIsProxyableOk() (*EnumtopologyAdminUserIsProxyableProp, bool)`

GetIsProxyableOk returns a tuple with the IsProxyable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyable

`func (o *TopologyAdminUserResponse) SetIsProxyable(v EnumtopologyAdminUserIsProxyableProp)`

SetIsProxyable sets IsProxyable field to given value.

### HasIsProxyable

`func (o *TopologyAdminUserResponse) HasIsProxyable() bool`

HasIsProxyable returns a boolean if a field has been set.

### GetIsProxyableByDN

`func (o *TopologyAdminUserResponse) GetIsProxyableByDN() []string`

GetIsProxyableByDN returns the IsProxyableByDN field if non-nil, zero value otherwise.

### GetIsProxyableByDNOk

`func (o *TopologyAdminUserResponse) GetIsProxyableByDNOk() (*[]string, bool)`

GetIsProxyableByDNOk returns a tuple with the IsProxyableByDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyableByDN

`func (o *TopologyAdminUserResponse) SetIsProxyableByDN(v []string)`

SetIsProxyableByDN sets IsProxyableByDN field to given value.

### HasIsProxyableByDN

`func (o *TopologyAdminUserResponse) HasIsProxyableByDN() bool`

HasIsProxyableByDN returns a boolean if a field has been set.

### GetIsProxyableByGroup

`func (o *TopologyAdminUserResponse) GetIsProxyableByGroup() []string`

GetIsProxyableByGroup returns the IsProxyableByGroup field if non-nil, zero value otherwise.

### GetIsProxyableByGroupOk

`func (o *TopologyAdminUserResponse) GetIsProxyableByGroupOk() (*[]string, bool)`

GetIsProxyableByGroupOk returns a tuple with the IsProxyableByGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyableByGroup

`func (o *TopologyAdminUserResponse) SetIsProxyableByGroup(v []string)`

SetIsProxyableByGroup sets IsProxyableByGroup field to given value.

### HasIsProxyableByGroup

`func (o *TopologyAdminUserResponse) HasIsProxyableByGroup() bool`

HasIsProxyableByGroup returns a boolean if a field has been set.

### GetIsProxyableByURL

`func (o *TopologyAdminUserResponse) GetIsProxyableByURL() []string`

GetIsProxyableByURL returns the IsProxyableByURL field if non-nil, zero value otherwise.

### GetIsProxyableByURLOk

`func (o *TopologyAdminUserResponse) GetIsProxyableByURLOk() (*[]string, bool)`

GetIsProxyableByURLOk returns a tuple with the IsProxyableByURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxyableByURL

`func (o *TopologyAdminUserResponse) SetIsProxyableByURL(v []string)`

SetIsProxyableByURL sets IsProxyableByURL field to given value.

### HasIsProxyableByURL

`func (o *TopologyAdminUserResponse) HasIsProxyableByURL() bool`

HasIsProxyableByURL returns a boolean if a field has been set.

### GetMayProxyAsDN

`func (o *TopologyAdminUserResponse) GetMayProxyAsDN() []string`

GetMayProxyAsDN returns the MayProxyAsDN field if non-nil, zero value otherwise.

### GetMayProxyAsDNOk

`func (o *TopologyAdminUserResponse) GetMayProxyAsDNOk() (*[]string, bool)`

GetMayProxyAsDNOk returns a tuple with the MayProxyAsDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayProxyAsDN

`func (o *TopologyAdminUserResponse) SetMayProxyAsDN(v []string)`

SetMayProxyAsDN sets MayProxyAsDN field to given value.

### HasMayProxyAsDN

`func (o *TopologyAdminUserResponse) HasMayProxyAsDN() bool`

HasMayProxyAsDN returns a boolean if a field has been set.

### GetMayProxyAsGroup

`func (o *TopologyAdminUserResponse) GetMayProxyAsGroup() []string`

GetMayProxyAsGroup returns the MayProxyAsGroup field if non-nil, zero value otherwise.

### GetMayProxyAsGroupOk

`func (o *TopologyAdminUserResponse) GetMayProxyAsGroupOk() (*[]string, bool)`

GetMayProxyAsGroupOk returns a tuple with the MayProxyAsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayProxyAsGroup

`func (o *TopologyAdminUserResponse) SetMayProxyAsGroup(v []string)`

SetMayProxyAsGroup sets MayProxyAsGroup field to given value.

### HasMayProxyAsGroup

`func (o *TopologyAdminUserResponse) HasMayProxyAsGroup() bool`

HasMayProxyAsGroup returns a boolean if a field has been set.

### GetMayProxyAsURL

`func (o *TopologyAdminUserResponse) GetMayProxyAsURL() []string`

GetMayProxyAsURL returns the MayProxyAsURL field if non-nil, zero value otherwise.

### GetMayProxyAsURLOk

`func (o *TopologyAdminUserResponse) GetMayProxyAsURLOk() (*[]string, bool)`

GetMayProxyAsURLOk returns a tuple with the MayProxyAsURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayProxyAsURL

`func (o *TopologyAdminUserResponse) SetMayProxyAsURL(v []string)`

SetMayProxyAsURL sets MayProxyAsURL field to given value.

### HasMayProxyAsURL

`func (o *TopologyAdminUserResponse) HasMayProxyAsURL() bool`

HasMayProxyAsURL returns a boolean if a field has been set.

### GetMeta

`func (o *TopologyAdminUserResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TopologyAdminUserResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TopologyAdminUserResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TopologyAdminUserResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *TopologyAdminUserResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *TopologyAdminUserResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *TopologyAdminUserResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *TopologyAdminUserResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


