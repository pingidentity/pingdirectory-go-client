# JdbcExternalServerShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjdbcExternalServerSchemaUrn**](EnumjdbcExternalServerSchemaUrn.md) |  | 
**JdbcDriverType** | [**EnumexternalServerJdbcDriverTypeProp**](EnumexternalServerJdbcDriverTypeProp.md) |  | 
**JdbcDriverURL** | Pointer to **string** | Specify the complete JDBC URL which will be used instead of the automatic URL format. You must select type &#39;other&#39; for the jdbc-driver-type. | [optional] 
**DatabaseName** | Pointer to **string** | Specifies which database to connect to. This is ignored if jdbc-driver-url is specified. | [optional] 
**ServerHostName** | Pointer to **string** | The host name of the database server. This is ignored if jdbc-driver-url is specified. | [optional] 
**ServerPort** | Pointer to **int32** | The port number where the database server listens for requests. This is ignored if jdbc-driver-url is specified | [optional] 
**UserName** | Pointer to **string** | The name of the login account to use when connecting to the database server. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user name. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**ValidationQuery** | Pointer to **string** | The SQL query that will be used to validate connections to the database before making them available to the Directory Server. | [optional] 
**ValidationQueryTimeout** | Pointer to **string** | Specifies the amount of time to wait for a response from the database when executing the validation query, if one is set. If the timeout is exceeded, the Directory Server will drop the connection and obtain a new one. A value of zero indicates no timeout. | [optional] 
**JdbcConnectionProperties** | Pointer to **[]string** |  | [optional] 
**TransactionIsolationLevel** | Pointer to [**EnumexternalServerTransactionIsolationLevelProp**](EnumexternalServerTransactionIsolationLevelProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 

## Methods

### NewJdbcExternalServerShared

`func NewJdbcExternalServerShared(schemas []EnumjdbcExternalServerSchemaUrn, jdbcDriverType EnumexternalServerJdbcDriverTypeProp, ) *JdbcExternalServerShared`

NewJdbcExternalServerShared instantiates a new JdbcExternalServerShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJdbcExternalServerSharedWithDefaults

`func NewJdbcExternalServerSharedWithDefaults() *JdbcExternalServerShared`

NewJdbcExternalServerSharedWithDefaults instantiates a new JdbcExternalServerShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *JdbcExternalServerShared) GetSchemas() []EnumjdbcExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *JdbcExternalServerShared) GetSchemasOk() (*[]EnumjdbcExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *JdbcExternalServerShared) SetSchemas(v []EnumjdbcExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetJdbcDriverType

`func (o *JdbcExternalServerShared) GetJdbcDriverType() EnumexternalServerJdbcDriverTypeProp`

GetJdbcDriverType returns the JdbcDriverType field if non-nil, zero value otherwise.

### GetJdbcDriverTypeOk

`func (o *JdbcExternalServerShared) GetJdbcDriverTypeOk() (*EnumexternalServerJdbcDriverTypeProp, bool)`

GetJdbcDriverTypeOk returns a tuple with the JdbcDriverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcDriverType

`func (o *JdbcExternalServerShared) SetJdbcDriverType(v EnumexternalServerJdbcDriverTypeProp)`

SetJdbcDriverType sets JdbcDriverType field to given value.


### GetJdbcDriverURL

`func (o *JdbcExternalServerShared) GetJdbcDriverURL() string`

GetJdbcDriverURL returns the JdbcDriverURL field if non-nil, zero value otherwise.

### GetJdbcDriverURLOk

`func (o *JdbcExternalServerShared) GetJdbcDriverURLOk() (*string, bool)`

GetJdbcDriverURLOk returns a tuple with the JdbcDriverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcDriverURL

`func (o *JdbcExternalServerShared) SetJdbcDriverURL(v string)`

SetJdbcDriverURL sets JdbcDriverURL field to given value.

### HasJdbcDriverURL

`func (o *JdbcExternalServerShared) HasJdbcDriverURL() bool`

HasJdbcDriverURL returns a boolean if a field has been set.

### GetDatabaseName

`func (o *JdbcExternalServerShared) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *JdbcExternalServerShared) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *JdbcExternalServerShared) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *JdbcExternalServerShared) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetServerHostName

`func (o *JdbcExternalServerShared) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *JdbcExternalServerShared) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *JdbcExternalServerShared) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.

### HasServerHostName

`func (o *JdbcExternalServerShared) HasServerHostName() bool`

HasServerHostName returns a boolean if a field has been set.

### GetServerPort

`func (o *JdbcExternalServerShared) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *JdbcExternalServerShared) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *JdbcExternalServerShared) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *JdbcExternalServerShared) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetUserName

`func (o *JdbcExternalServerShared) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *JdbcExternalServerShared) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *JdbcExternalServerShared) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *JdbcExternalServerShared) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *JdbcExternalServerShared) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *JdbcExternalServerShared) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *JdbcExternalServerShared) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *JdbcExternalServerShared) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *JdbcExternalServerShared) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *JdbcExternalServerShared) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *JdbcExternalServerShared) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *JdbcExternalServerShared) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetValidationQuery

`func (o *JdbcExternalServerShared) GetValidationQuery() string`

GetValidationQuery returns the ValidationQuery field if non-nil, zero value otherwise.

### GetValidationQueryOk

`func (o *JdbcExternalServerShared) GetValidationQueryOk() (*string, bool)`

GetValidationQueryOk returns a tuple with the ValidationQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationQuery

`func (o *JdbcExternalServerShared) SetValidationQuery(v string)`

SetValidationQuery sets ValidationQuery field to given value.

### HasValidationQuery

`func (o *JdbcExternalServerShared) HasValidationQuery() bool`

HasValidationQuery returns a boolean if a field has been set.

### GetValidationQueryTimeout

`func (o *JdbcExternalServerShared) GetValidationQueryTimeout() string`

GetValidationQueryTimeout returns the ValidationQueryTimeout field if non-nil, zero value otherwise.

### GetValidationQueryTimeoutOk

`func (o *JdbcExternalServerShared) GetValidationQueryTimeoutOk() (*string, bool)`

GetValidationQueryTimeoutOk returns a tuple with the ValidationQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationQueryTimeout

`func (o *JdbcExternalServerShared) SetValidationQueryTimeout(v string)`

SetValidationQueryTimeout sets ValidationQueryTimeout field to given value.

### HasValidationQueryTimeout

`func (o *JdbcExternalServerShared) HasValidationQueryTimeout() bool`

HasValidationQueryTimeout returns a boolean if a field has been set.

### GetJdbcConnectionProperties

`func (o *JdbcExternalServerShared) GetJdbcConnectionProperties() []string`

GetJdbcConnectionProperties returns the JdbcConnectionProperties field if non-nil, zero value otherwise.

### GetJdbcConnectionPropertiesOk

`func (o *JdbcExternalServerShared) GetJdbcConnectionPropertiesOk() (*[]string, bool)`

GetJdbcConnectionPropertiesOk returns a tuple with the JdbcConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcConnectionProperties

`func (o *JdbcExternalServerShared) SetJdbcConnectionProperties(v []string)`

SetJdbcConnectionProperties sets JdbcConnectionProperties field to given value.

### HasJdbcConnectionProperties

`func (o *JdbcExternalServerShared) HasJdbcConnectionProperties() bool`

HasJdbcConnectionProperties returns a boolean if a field has been set.

### GetTransactionIsolationLevel

`func (o *JdbcExternalServerShared) GetTransactionIsolationLevel() EnumexternalServerTransactionIsolationLevelProp`

GetTransactionIsolationLevel returns the TransactionIsolationLevel field if non-nil, zero value otherwise.

### GetTransactionIsolationLevelOk

`func (o *JdbcExternalServerShared) GetTransactionIsolationLevelOk() (*EnumexternalServerTransactionIsolationLevelProp, bool)`

GetTransactionIsolationLevelOk returns a tuple with the TransactionIsolationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIsolationLevel

`func (o *JdbcExternalServerShared) SetTransactionIsolationLevel(v EnumexternalServerTransactionIsolationLevelProp)`

SetTransactionIsolationLevel sets TransactionIsolationLevel field to given value.

### HasTransactionIsolationLevel

`func (o *JdbcExternalServerShared) HasTransactionIsolationLevel() bool`

HasTransactionIsolationLevel returns a boolean if a field has been set.

### GetDescription

`func (o *JdbcExternalServerShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JdbcExternalServerShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JdbcExternalServerShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JdbcExternalServerShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


