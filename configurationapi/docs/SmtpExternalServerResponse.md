# SmtpExternalServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the External Server | 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewSmtpExternalServerResponse

`func NewSmtpExternalServerResponse(id string, schemas []EnumsmtpExternalServerSchemaUrn, serverHostName string, ) *SmtpExternalServerResponse`

NewSmtpExternalServerResponse instantiates a new SmtpExternalServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmtpExternalServerResponseWithDefaults

`func NewSmtpExternalServerResponseWithDefaults() *SmtpExternalServerResponse`

NewSmtpExternalServerResponseWithDefaults instantiates a new SmtpExternalServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SmtpExternalServerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SmtpExternalServerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SmtpExternalServerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SmtpExternalServerResponse) GetSchemas() []EnumsmtpExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SmtpExternalServerResponse) GetSchemasOk() (*[]EnumsmtpExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SmtpExternalServerResponse) SetSchemas(v []EnumsmtpExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerHostName

`func (o *SmtpExternalServerResponse) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *SmtpExternalServerResponse) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *SmtpExternalServerResponse) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.


### GetServerPort

`func (o *SmtpExternalServerResponse) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *SmtpExternalServerResponse) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *SmtpExternalServerResponse) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *SmtpExternalServerResponse) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetSmtpSecurity

`func (o *SmtpExternalServerResponse) GetSmtpSecurity() EnumexternalServerSmtpSecurityProp`

GetSmtpSecurity returns the SmtpSecurity field if non-nil, zero value otherwise.

### GetSmtpSecurityOk

`func (o *SmtpExternalServerResponse) GetSmtpSecurityOk() (*EnumexternalServerSmtpSecurityProp, bool)`

GetSmtpSecurityOk returns a tuple with the SmtpSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpSecurity

`func (o *SmtpExternalServerResponse) SetSmtpSecurity(v EnumexternalServerSmtpSecurityProp)`

SetSmtpSecurity sets SmtpSecurity field to given value.

### HasSmtpSecurity

`func (o *SmtpExternalServerResponse) HasSmtpSecurity() bool`

HasSmtpSecurity returns a boolean if a field has been set.

### GetUserName

`func (o *SmtpExternalServerResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SmtpExternalServerResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SmtpExternalServerResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SmtpExternalServerResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *SmtpExternalServerResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SmtpExternalServerResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SmtpExternalServerResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SmtpExternalServerResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *SmtpExternalServerResponse) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *SmtpExternalServerResponse) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *SmtpExternalServerResponse) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *SmtpExternalServerResponse) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetSmtpTimeout

`func (o *SmtpExternalServerResponse) GetSmtpTimeout() string`

GetSmtpTimeout returns the SmtpTimeout field if non-nil, zero value otherwise.

### GetSmtpTimeoutOk

`func (o *SmtpExternalServerResponse) GetSmtpTimeoutOk() (*string, bool)`

GetSmtpTimeoutOk returns a tuple with the SmtpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpTimeout

`func (o *SmtpExternalServerResponse) SetSmtpTimeout(v string)`

SetSmtpTimeout sets SmtpTimeout field to given value.

### HasSmtpTimeout

`func (o *SmtpExternalServerResponse) HasSmtpTimeout() bool`

HasSmtpTimeout returns a boolean if a field has been set.

### GetSmtpConnectionProperties

`func (o *SmtpExternalServerResponse) GetSmtpConnectionProperties() []string`

GetSmtpConnectionProperties returns the SmtpConnectionProperties field if non-nil, zero value otherwise.

### GetSmtpConnectionPropertiesOk

`func (o *SmtpExternalServerResponse) GetSmtpConnectionPropertiesOk() (*[]string, bool)`

GetSmtpConnectionPropertiesOk returns a tuple with the SmtpConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpConnectionProperties

`func (o *SmtpExternalServerResponse) SetSmtpConnectionProperties(v []string)`

SetSmtpConnectionProperties sets SmtpConnectionProperties field to given value.

### HasSmtpConnectionProperties

`func (o *SmtpExternalServerResponse) HasSmtpConnectionProperties() bool`

HasSmtpConnectionProperties returns a boolean if a field has been set.

### GetDescription

`func (o *SmtpExternalServerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SmtpExternalServerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SmtpExternalServerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SmtpExternalServerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *SmtpExternalServerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SmtpExternalServerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SmtpExternalServerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SmtpExternalServerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SmtpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SmtpExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SmtpExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SmtpExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


