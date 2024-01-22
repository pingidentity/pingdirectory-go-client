# AddJdbcExternalServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumjdbcExternalServerSchemaUrn**](EnumjdbcExternalServerSchemaUrn.md) |  | 
**JdbcDriverType** | [**EnumexternalServerJdbcDriverTypeProp**](EnumexternalServerJdbcDriverTypeProp.md) |  | 
**JdbcDriverURL** | Pointer to **string** | Specify the complete JDBC URL which will be used instead of the automatic URL format. You must select type &#39;other&#39; for the jdbc-driver-type. | [optional] 
**DatabaseName** | Pointer to **string** | Specifies which database to connect to. This is ignored if jdbc-driver-url is specified. | [optional] 
**ServerHostName** | Pointer to **string** | The host name of the database server. This is ignored if jdbc-driver-url is specified. | [optional] 
**ServerPort** | Pointer to **int64** | The port number where the database server listens for requests. This is ignored if jdbc-driver-url is specified | [optional] 
**UserName** | Pointer to **string** | The name of the login account to use when connecting to the database server. | [optional] 
**Password** | Pointer to **string** | The login password for the specified user name. | [optional] 
**PassphraseProvider** | Pointer to **string** | The passphrase provider to use to obtain the login password for the specified user. | [optional] 
**ValidationQuery** | Pointer to **string** | The SQL query that will be used to validate connections to the database before making them available to the Directory Server. | [optional] 
**ValidationQueryTimeout** | Pointer to **string** | Specifies the amount of time to wait for a response from the database when executing the validation query, if one is set. If the timeout is exceeded, the Directory Server will drop the connection and obtain a new one. A value of zero indicates no timeout. | [optional] 
**JdbcConnectionProperties** | Pointer to **[]string** | Specifies the connection properties for the JDBC datasource. | [optional] 
**TransactionIsolationLevel** | Pointer to [**EnumexternalServerTransactionIsolationLevelProp**](EnumexternalServerTransactionIsolationLevelProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this External Server | [optional] 
**ServerName** | **string** | Name of the new External Server | 

## Methods

### NewAddJdbcExternalServerRequest

`func NewAddJdbcExternalServerRequest(schemas []EnumjdbcExternalServerSchemaUrn, jdbcDriverType EnumexternalServerJdbcDriverTypeProp, serverName string, ) *AddJdbcExternalServerRequest`

NewAddJdbcExternalServerRequest instantiates a new AddJdbcExternalServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJdbcExternalServerRequestWithDefaults

`func NewAddJdbcExternalServerRequestWithDefaults() *AddJdbcExternalServerRequest`

NewAddJdbcExternalServerRequestWithDefaults instantiates a new AddJdbcExternalServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddJdbcExternalServerRequest) GetSchemas() []EnumjdbcExternalServerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddJdbcExternalServerRequest) GetSchemasOk() (*[]EnumjdbcExternalServerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddJdbcExternalServerRequest) SetSchemas(v []EnumjdbcExternalServerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetJdbcDriverType

`func (o *AddJdbcExternalServerRequest) GetJdbcDriverType() EnumexternalServerJdbcDriverTypeProp`

GetJdbcDriverType returns the JdbcDriverType field if non-nil, zero value otherwise.

### GetJdbcDriverTypeOk

`func (o *AddJdbcExternalServerRequest) GetJdbcDriverTypeOk() (*EnumexternalServerJdbcDriverTypeProp, bool)`

GetJdbcDriverTypeOk returns a tuple with the JdbcDriverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcDriverType

`func (o *AddJdbcExternalServerRequest) SetJdbcDriverType(v EnumexternalServerJdbcDriverTypeProp)`

SetJdbcDriverType sets JdbcDriverType field to given value.


### GetJdbcDriverURL

`func (o *AddJdbcExternalServerRequest) GetJdbcDriverURL() string`

GetJdbcDriverURL returns the JdbcDriverURL field if non-nil, zero value otherwise.

### GetJdbcDriverURLOk

`func (o *AddJdbcExternalServerRequest) GetJdbcDriverURLOk() (*string, bool)`

GetJdbcDriverURLOk returns a tuple with the JdbcDriverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcDriverURL

`func (o *AddJdbcExternalServerRequest) SetJdbcDriverURL(v string)`

SetJdbcDriverURL sets JdbcDriverURL field to given value.

### HasJdbcDriverURL

`func (o *AddJdbcExternalServerRequest) HasJdbcDriverURL() bool`

HasJdbcDriverURL returns a boolean if a field has been set.

### GetDatabaseName

`func (o *AddJdbcExternalServerRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *AddJdbcExternalServerRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *AddJdbcExternalServerRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *AddJdbcExternalServerRequest) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetServerHostName

`func (o *AddJdbcExternalServerRequest) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *AddJdbcExternalServerRequest) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *AddJdbcExternalServerRequest) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.

### HasServerHostName

`func (o *AddJdbcExternalServerRequest) HasServerHostName() bool`

HasServerHostName returns a boolean if a field has been set.

### GetServerPort

`func (o *AddJdbcExternalServerRequest) GetServerPort() int64`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *AddJdbcExternalServerRequest) GetServerPortOk() (*int64, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *AddJdbcExternalServerRequest) SetServerPort(v int64)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *AddJdbcExternalServerRequest) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetUserName

`func (o *AddJdbcExternalServerRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AddJdbcExternalServerRequest) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AddJdbcExternalServerRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AddJdbcExternalServerRequest) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *AddJdbcExternalServerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddJdbcExternalServerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddJdbcExternalServerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddJdbcExternalServerRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPassphraseProvider

`func (o *AddJdbcExternalServerRequest) GetPassphraseProvider() string`

GetPassphraseProvider returns the PassphraseProvider field if non-nil, zero value otherwise.

### GetPassphraseProviderOk

`func (o *AddJdbcExternalServerRequest) GetPassphraseProviderOk() (*string, bool)`

GetPassphraseProviderOk returns a tuple with the PassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphraseProvider

`func (o *AddJdbcExternalServerRequest) SetPassphraseProvider(v string)`

SetPassphraseProvider sets PassphraseProvider field to given value.

### HasPassphraseProvider

`func (o *AddJdbcExternalServerRequest) HasPassphraseProvider() bool`

HasPassphraseProvider returns a boolean if a field has been set.

### GetValidationQuery

`func (o *AddJdbcExternalServerRequest) GetValidationQuery() string`

GetValidationQuery returns the ValidationQuery field if non-nil, zero value otherwise.

### GetValidationQueryOk

`func (o *AddJdbcExternalServerRequest) GetValidationQueryOk() (*string, bool)`

GetValidationQueryOk returns a tuple with the ValidationQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationQuery

`func (o *AddJdbcExternalServerRequest) SetValidationQuery(v string)`

SetValidationQuery sets ValidationQuery field to given value.

### HasValidationQuery

`func (o *AddJdbcExternalServerRequest) HasValidationQuery() bool`

HasValidationQuery returns a boolean if a field has been set.

### GetValidationQueryTimeout

`func (o *AddJdbcExternalServerRequest) GetValidationQueryTimeout() string`

GetValidationQueryTimeout returns the ValidationQueryTimeout field if non-nil, zero value otherwise.

### GetValidationQueryTimeoutOk

`func (o *AddJdbcExternalServerRequest) GetValidationQueryTimeoutOk() (*string, bool)`

GetValidationQueryTimeoutOk returns a tuple with the ValidationQueryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationQueryTimeout

`func (o *AddJdbcExternalServerRequest) SetValidationQueryTimeout(v string)`

SetValidationQueryTimeout sets ValidationQueryTimeout field to given value.

### HasValidationQueryTimeout

`func (o *AddJdbcExternalServerRequest) HasValidationQueryTimeout() bool`

HasValidationQueryTimeout returns a boolean if a field has been set.

### GetJdbcConnectionProperties

`func (o *AddJdbcExternalServerRequest) GetJdbcConnectionProperties() []string`

GetJdbcConnectionProperties returns the JdbcConnectionProperties field if non-nil, zero value otherwise.

### GetJdbcConnectionPropertiesOk

`func (o *AddJdbcExternalServerRequest) GetJdbcConnectionPropertiesOk() (*[]string, bool)`

GetJdbcConnectionPropertiesOk returns a tuple with the JdbcConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJdbcConnectionProperties

`func (o *AddJdbcExternalServerRequest) SetJdbcConnectionProperties(v []string)`

SetJdbcConnectionProperties sets JdbcConnectionProperties field to given value.

### HasJdbcConnectionProperties

`func (o *AddJdbcExternalServerRequest) HasJdbcConnectionProperties() bool`

HasJdbcConnectionProperties returns a boolean if a field has been set.

### GetTransactionIsolationLevel

`func (o *AddJdbcExternalServerRequest) GetTransactionIsolationLevel() EnumexternalServerTransactionIsolationLevelProp`

GetTransactionIsolationLevel returns the TransactionIsolationLevel field if non-nil, zero value otherwise.

### GetTransactionIsolationLevelOk

`func (o *AddJdbcExternalServerRequest) GetTransactionIsolationLevelOk() (*EnumexternalServerTransactionIsolationLevelProp, bool)`

GetTransactionIsolationLevelOk returns a tuple with the TransactionIsolationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionIsolationLevel

`func (o *AddJdbcExternalServerRequest) SetTransactionIsolationLevel(v EnumexternalServerTransactionIsolationLevelProp)`

SetTransactionIsolationLevel sets TransactionIsolationLevel field to given value.

### HasTransactionIsolationLevel

`func (o *AddJdbcExternalServerRequest) HasTransactionIsolationLevel() bool`

HasTransactionIsolationLevel returns a boolean if a field has been set.

### GetDescription

`func (o *AddJdbcExternalServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddJdbcExternalServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddJdbcExternalServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddJdbcExternalServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServerName

`func (o *AddJdbcExternalServerRequest) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AddJdbcExternalServerRequest) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AddJdbcExternalServerRequest) SetServerName(v string)`

SetServerName sets ServerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


