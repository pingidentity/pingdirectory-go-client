/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configurationapi

import (
	"encoding/json"
)

// checks if the AddJdbcExternalServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddJdbcExternalServerRequest{}

// AddJdbcExternalServerRequest struct for AddJdbcExternalServerRequest
type AddJdbcExternalServerRequest struct {
	// Name of the new External Server
	ServerName     string                               `json:"serverName"`
	Schemas        []EnumjdbcExternalServerSchemaUrn    `json:"schemas"`
	JdbcDriverType EnumexternalServerJdbcDriverTypeProp `json:"jdbcDriverType"`
	// Specify the complete JDBC URL which will be used instead of the automatic URL format. You must select type 'other' for the jdbc-driver-type.
	JdbcDriverURL *string `json:"jdbcDriverURL,omitempty"`
	// Specifies which database to connect to. This is ignored if jdbc-driver-url is specified.
	DatabaseName *string `json:"databaseName,omitempty"`
	// The host name of the database server. This is ignored if jdbc-driver-url is specified.
	ServerHostName *string `json:"serverHostName,omitempty"`
	// The port number where the database server listens for requests. This is ignored if jdbc-driver-url is specified
	ServerPort *int32 `json:"serverPort,omitempty"`
	// The name of the login account to use when connecting to the database server.
	UserName *string `json:"userName,omitempty"`
	// The login password for the specified user name.
	Password *string `json:"password,omitempty"`
	// The passphrase provider to use to obtain the login password for the specified user.
	PassphraseProvider *string `json:"passphraseProvider,omitempty"`
	// The SQL query that will be used to validate connections to the database before making them available to the Directory Server.
	ValidationQuery *string `json:"validationQuery,omitempty"`
	// Specifies the amount of time to wait for a response from the database when executing the validation query, if one is set. If the timeout is exceeded, the Directory Server will drop the connection and obtain a new one. A value of zero indicates no timeout.
	ValidationQueryTimeout *string `json:"validationQueryTimeout,omitempty"`
	// Specifies the connection properties for the JDBC datasource.
	JdbcConnectionProperties  []string                                         `json:"jdbcConnectionProperties,omitempty"`
	TransactionIsolationLevel *EnumexternalServerTransactionIsolationLevelProp `json:"transactionIsolationLevel,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewAddJdbcExternalServerRequest instantiates a new AddJdbcExternalServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddJdbcExternalServerRequest(serverName string, schemas []EnumjdbcExternalServerSchemaUrn, jdbcDriverType EnumexternalServerJdbcDriverTypeProp) *AddJdbcExternalServerRequest {
	this := AddJdbcExternalServerRequest{}
	this.ServerName = serverName
	this.Schemas = schemas
	this.JdbcDriverType = jdbcDriverType
	return &this
}

// NewAddJdbcExternalServerRequestWithDefaults instantiates a new AddJdbcExternalServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddJdbcExternalServerRequestWithDefaults() *AddJdbcExternalServerRequest {
	this := AddJdbcExternalServerRequest{}
	return &this
}

// GetServerName returns the ServerName field value
func (o *AddJdbcExternalServerRequest) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *AddJdbcExternalServerRequest) SetServerName(v string) {
	o.ServerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddJdbcExternalServerRequest) GetSchemas() []EnumjdbcExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumjdbcExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetSchemasOk() ([]EnumjdbcExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddJdbcExternalServerRequest) SetSchemas(v []EnumjdbcExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetJdbcDriverType returns the JdbcDriverType field value
func (o *AddJdbcExternalServerRequest) GetJdbcDriverType() EnumexternalServerJdbcDriverTypeProp {
	if o == nil {
		var ret EnumexternalServerJdbcDriverTypeProp
		return ret
	}

	return o.JdbcDriverType
}

// GetJdbcDriverTypeOk returns a tuple with the JdbcDriverType field value
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetJdbcDriverTypeOk() (*EnumexternalServerJdbcDriverTypeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JdbcDriverType, true
}

// SetJdbcDriverType sets field value
func (o *AddJdbcExternalServerRequest) SetJdbcDriverType(v EnumexternalServerJdbcDriverTypeProp) {
	o.JdbcDriverType = v
}

// GetJdbcDriverURL returns the JdbcDriverURL field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetJdbcDriverURL() string {
	if o == nil || IsNil(o.JdbcDriverURL) {
		var ret string
		return ret
	}
	return *o.JdbcDriverURL
}

// GetJdbcDriverURLOk returns a tuple with the JdbcDriverURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetJdbcDriverURLOk() (*string, bool) {
	if o == nil || IsNil(o.JdbcDriverURL) {
		return nil, false
	}
	return o.JdbcDriverURL, true
}

// HasJdbcDriverURL returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasJdbcDriverURL() bool {
	if o != nil && !IsNil(o.JdbcDriverURL) {
		return true
	}

	return false
}

// SetJdbcDriverURL gets a reference to the given string and assigns it to the JdbcDriverURL field.
func (o *AddJdbcExternalServerRequest) SetJdbcDriverURL(v string) {
	o.JdbcDriverURL = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetDatabaseName() string {
	if o == nil || IsNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetDatabaseNameOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseName) {
		return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasDatabaseName() bool {
	if o != nil && !IsNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *AddJdbcExternalServerRequest) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetServerHostName returns the ServerHostName field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetServerHostName() string {
	if o == nil || IsNil(o.ServerHostName) {
		var ret string
		return ret
	}
	return *o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetServerHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerHostName) {
		return nil, false
	}
	return o.ServerHostName, true
}

// HasServerHostName returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasServerHostName() bool {
	if o != nil && !IsNil(o.ServerHostName) {
		return true
	}

	return false
}

// SetServerHostName gets a reference to the given string and assigns it to the ServerHostName field.
func (o *AddJdbcExternalServerRequest) SetServerHostName(v string) {
	o.ServerHostName = &v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetServerPort() int32 {
	if o == nil || IsNil(o.ServerPort) {
		var ret int32
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetServerPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerPort) {
		return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasServerPort() bool {
	if o != nil && !IsNil(o.ServerPort) {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given int32 and assigns it to the ServerPort field.
func (o *AddJdbcExternalServerRequest) SetServerPort(v int32) {
	o.ServerPort = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *AddJdbcExternalServerRequest) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AddJdbcExternalServerRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetPassphraseProvider() string {
	if o == nil || IsNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.PassphraseProvider) {
		return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasPassphraseProvider() bool {
	if o != nil && !IsNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *AddJdbcExternalServerRequest) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetValidationQuery returns the ValidationQuery field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetValidationQuery() string {
	if o == nil || IsNil(o.ValidationQuery) {
		var ret string
		return ret
	}
	return *o.ValidationQuery
}

// GetValidationQueryOk returns a tuple with the ValidationQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetValidationQueryOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationQuery) {
		return nil, false
	}
	return o.ValidationQuery, true
}

// HasValidationQuery returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasValidationQuery() bool {
	if o != nil && !IsNil(o.ValidationQuery) {
		return true
	}

	return false
}

// SetValidationQuery gets a reference to the given string and assigns it to the ValidationQuery field.
func (o *AddJdbcExternalServerRequest) SetValidationQuery(v string) {
	o.ValidationQuery = &v
}

// GetValidationQueryTimeout returns the ValidationQueryTimeout field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetValidationQueryTimeout() string {
	if o == nil || IsNil(o.ValidationQueryTimeout) {
		var ret string
		return ret
	}
	return *o.ValidationQueryTimeout
}

// GetValidationQueryTimeoutOk returns a tuple with the ValidationQueryTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetValidationQueryTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationQueryTimeout) {
		return nil, false
	}
	return o.ValidationQueryTimeout, true
}

// HasValidationQueryTimeout returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasValidationQueryTimeout() bool {
	if o != nil && !IsNil(o.ValidationQueryTimeout) {
		return true
	}

	return false
}

// SetValidationQueryTimeout gets a reference to the given string and assigns it to the ValidationQueryTimeout field.
func (o *AddJdbcExternalServerRequest) SetValidationQueryTimeout(v string) {
	o.ValidationQueryTimeout = &v
}

// GetJdbcConnectionProperties returns the JdbcConnectionProperties field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetJdbcConnectionProperties() []string {
	if o == nil || IsNil(o.JdbcConnectionProperties) {
		var ret []string
		return ret
	}
	return o.JdbcConnectionProperties
}

// GetJdbcConnectionPropertiesOk returns a tuple with the JdbcConnectionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetJdbcConnectionPropertiesOk() ([]string, bool) {
	if o == nil || IsNil(o.JdbcConnectionProperties) {
		return nil, false
	}
	return o.JdbcConnectionProperties, true
}

// HasJdbcConnectionProperties returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasJdbcConnectionProperties() bool {
	if o != nil && !IsNil(o.JdbcConnectionProperties) {
		return true
	}

	return false
}

// SetJdbcConnectionProperties gets a reference to the given []string and assigns it to the JdbcConnectionProperties field.
func (o *AddJdbcExternalServerRequest) SetJdbcConnectionProperties(v []string) {
	o.JdbcConnectionProperties = v
}

// GetTransactionIsolationLevel returns the TransactionIsolationLevel field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetTransactionIsolationLevel() EnumexternalServerTransactionIsolationLevelProp {
	if o == nil || IsNil(o.TransactionIsolationLevel) {
		var ret EnumexternalServerTransactionIsolationLevelProp
		return ret
	}
	return *o.TransactionIsolationLevel
}

// GetTransactionIsolationLevelOk returns a tuple with the TransactionIsolationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetTransactionIsolationLevelOk() (*EnumexternalServerTransactionIsolationLevelProp, bool) {
	if o == nil || IsNil(o.TransactionIsolationLevel) {
		return nil, false
	}
	return o.TransactionIsolationLevel, true
}

// HasTransactionIsolationLevel returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasTransactionIsolationLevel() bool {
	if o != nil && !IsNil(o.TransactionIsolationLevel) {
		return true
	}

	return false
}

// SetTransactionIsolationLevel gets a reference to the given EnumexternalServerTransactionIsolationLevelProp and assigns it to the TransactionIsolationLevel field.
func (o *AddJdbcExternalServerRequest) SetTransactionIsolationLevel(v EnumexternalServerTransactionIsolationLevelProp) {
	o.TransactionIsolationLevel = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddJdbcExternalServerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcExternalServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddJdbcExternalServerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddJdbcExternalServerRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddJdbcExternalServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddJdbcExternalServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverName"] = o.ServerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["jdbcDriverType"] = o.JdbcDriverType
	if !IsNil(o.JdbcDriverURL) {
		toSerialize["jdbcDriverURL"] = o.JdbcDriverURL
	}
	if !IsNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !IsNil(o.ServerHostName) {
		toSerialize["serverHostName"] = o.ServerHostName
	}
	if !IsNil(o.ServerPort) {
		toSerialize["serverPort"] = o.ServerPort
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PassphraseProvider) {
		toSerialize["passphraseProvider"] = o.PassphraseProvider
	}
	if !IsNil(o.ValidationQuery) {
		toSerialize["validationQuery"] = o.ValidationQuery
	}
	if !IsNil(o.ValidationQueryTimeout) {
		toSerialize["validationQueryTimeout"] = o.ValidationQueryTimeout
	}
	if !IsNil(o.JdbcConnectionProperties) {
		toSerialize["jdbcConnectionProperties"] = o.JdbcConnectionProperties
	}
	if !IsNil(o.TransactionIsolationLevel) {
		toSerialize["transactionIsolationLevel"] = o.TransactionIsolationLevel
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAddJdbcExternalServerRequest struct {
	value *AddJdbcExternalServerRequest
	isSet bool
}

func (v NullableAddJdbcExternalServerRequest) Get() *AddJdbcExternalServerRequest {
	return v.value
}

func (v *NullableAddJdbcExternalServerRequest) Set(val *AddJdbcExternalServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddJdbcExternalServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddJdbcExternalServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddJdbcExternalServerRequest(val *AddJdbcExternalServerRequest) *NullableAddJdbcExternalServerRequest {
	return &NullableAddJdbcExternalServerRequest{value: val, isSet: true}
}

func (v NullableAddJdbcExternalServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddJdbcExternalServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
