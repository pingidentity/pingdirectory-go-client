# AddSmtpExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsmtpExternalServerSchemaUrn**](EnumsmtpExternalServerSchemaUrn.md) |  | 
**ServerHostName** | **string** | The host name of the smtp server. | 
**ServerPort** | Pointer to **int64** | The port number where the smtp server listens for requests. | [optional] 
**SmtpSecurity** | Pointer to [**EnumexternalServerSmtpSecurityProp**](EnumexternalServerSmtpSecurityProp.md) |  | [optional] 
**UserName** | Pointer to **string** | The name of the login account to use when connecting to the smtp server. Both username and password must be supplied if this attribute is set. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user name. Both username and password must be supplied if this attribute is set. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**SmtpTimeout** | Pointer to **string** | Specifies the maximum length of time that a connection or attempted connection to a SMTP server may take. | [optional] 
**SmtpConnectionProperties** | Pointer to **[]string** | Specifies the connection properties for the smtp server. | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 
**ServerName** | **string** | Name of the new External Server | 

## Methods

### NewAddSmtpExternalServerRequest

`func NewAddSmtpExternalServerRequest(schemas []EnumsmtpExternalServerSchemaUrn, serverHostName string, serverName string, ) *AddSmtpExternalServerRequest`

NewAddSmtpExternalServerRequest instantiates a new AddSmtpExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSmtpExternalServerRequestWithDefaults

`func NewAddSmtpExternalServerRequestWithDefaults() *AddSmtpExternalServerRequest`

NewAddSmtpExternalServerRequestWithDefaults instantiates a new AddSmtpExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSmtpExternalServerRequest) GetSchemas() []EnumsmtpExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSmtpExternalServerRequest) GetSchemasOk() (*[]EnumsmtpExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSmtpExternalServerRequest) SetSchemas(v []EnumsmtpExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *AddSmtpExternalServerRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddSmtpExternalServerRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddSmtpExternalServerRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *AddSmtpExternalServerRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddSmtpExternalServerRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddSmtpExternalServerRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AddSmtpExternalServerRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetSmtpSecurity

`func (o *AddSmtpExternalServerRequest) GetSmtpSecurity() EnumexternalServerSmtpSecurityProp`

GetSmtpSecurity returns the SmtpSecurity field if non-nil, zero value otherwise.

### GetSmtpSecurityOk

`func (o *AddSmtpExternalServerRequest) GetSmtpSecurityOk() (*EnumexternalServerSmtpSecurityProp, bool)`

GetSmtpSecurityOk returns a tuple with the SmtpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSecurity

`func (o *AddSmtpExternalServerRequest) SetSmtpSecurity(v EnumexternalServerSmtpSecurityProp)`

SetSmtpSecurity sets SmtpSecurity field to given value.

### HasSmtpSecurity

`func (o *AddSmtpExternalServerRequest) HasSmtpSecurity() bool`

HasSmtpSecurity returns a boolean if a field has been set.

### GetUserName

`func (o *AddSmtpExternalServerRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AddSmtpExternalServerRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AddSmtpExternalServerRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AddSmtpExternalServerRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *AddSmtpExternalServerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddSmtpExternalServerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddSmtpExternalServerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddSmtpExternalServerRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *AddSmtpExternalServerRequest) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *AddSmtpExternalServerRequest) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *AddSmtpExternalServerRequest) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *AddSmtpExternalServerRequest) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetSmtpTimeout

`func (o *AddSmtpExternalServerRequest) GetSmtpTimeout() string`

GetSmtpTimeout returns the SmtpTimeout field if non-nil, zero value otherwise.

### GetSmtpTimeoutOk

`func (o *AddSmtpExternalServerRequest) GetSmtpTimeoutOk() (*string, bool)`

GetSmtpTimeoutOk returns a tuple with the SmtpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTimeout

`func (o *AddSmtpExternalServerRequest) SetSmtpTimeout(v string)`

SetSmtpTimeout sets SmtpTimeout field to given value.

### HasSmtpTimeout

`func (o *AddSmtpExternalServerRequest) HasSmtpTimeout() bool`

HasSmtpTimeout returns a boolean if a field has been set.

### GetSmtpConnectionProperties

`func (o *AddSmtpExternalServerRequest) GetSmtpConnectionProperties() []string`

GetSmtpConnectionProperties returns the SmtpConnectionProperties field if non-nil, zero value otherwise.

### GetSmtpConnectionPropertiesOk

`func (o *AddSmtpExternalServerRequest) GetSmtpConnectionPropertiesOk() (*[]string, bool)`

GetSmtpConnectionPropertiesOk returns a tuple with the SmtpConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpConnectionProperties

`func (o *AddSmtpExternalServerRequest) SetSmtpConnectionProperties(v []string)`

SetSmtpConnectionProperties sets SmtpConnectionProperties field to given value.

### HasSmtpConnectionProperties

`func (o *AddSmtpExternalServerRequest) HasSmtpConnectionProperties() bool`

HasSmtpConnectionProperties returns a boolean if a field has been set.

### GetDescription

`func (o *AddSmtpExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSmtpExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSmtpExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSmtpExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServerName

`func (o *AddSmtpExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddSmtpExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddSmtpExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


