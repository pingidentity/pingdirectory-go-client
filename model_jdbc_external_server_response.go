/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// JdbcExternalServerResponse struct for JdbcExternalServerResponse
type JdbcExternalServerResponse struct {
	// Name of the External Server
	Id string `json:"id"`
	Schemas []EnumjdbcExternalServerSchemaUrn `json:"schemas"`
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
	JdbcConnectionProperties []string `json:"jdbcConnectionProperties,omitempty"`
	TransactionIsolationLevel *EnumexternalServerTransactionIsolationLevelProp `json:"transactionIsolationLevel,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewJdbcExternalServerResponse instantiates a new JdbcExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJdbcExternalServerResponse(id string, schemas []EnumjdbcExternalServerSchemaUrn, jdbcDriverType EnumexternalServerJdbcDriverTypeProp) *JdbcExternalServerResponse {
	this := JdbcExternalServerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.JdbcDriverType = jdbcDriverType
	return &this
}

// NewJdbcExternalServerResponseWithDefaults instantiates a new JdbcExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJdbcExternalServerResponseWithDefaults() *JdbcExternalServerResponse {
	this := JdbcExternalServerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *JdbcExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JdbcExternalServerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *JdbcExternalServerResponse) GetSchemas() []EnumjdbcExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumjdbcExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetSchemasOk() ([]EnumjdbcExternalServerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JdbcExternalServerResponse) SetSchemas(v []EnumjdbcExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetJdbcDriverType returns the JdbcDriverType field value
func (o *JdbcExternalServerResponse) GetJdbcDriverType() EnumexternalServerJdbcDriverTypeProp {
	if o == nil {
		var ret EnumexternalServerJdbcDriverTypeProp
		return ret
	}

	return o.JdbcDriverType
}

// GetJdbcDriverTypeOk returns a tuple with the JdbcDriverType field value
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetJdbcDriverTypeOk() (*EnumexternalServerJdbcDriverTypeProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.JdbcDriverType, true
}

// SetJdbcDriverType sets field value
func (o *JdbcExternalServerResponse) SetJdbcDriverType(v EnumexternalServerJdbcDriverTypeProp) {
	o.JdbcDriverType = v
}

// GetJdbcDriverURL returns the JdbcDriverURL field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetJdbcDriverURL() string {
	if o == nil || isNil(o.JdbcDriverURL) {
		var ret string
		return ret
	}
	return *o.JdbcDriverURL
}

// GetJdbcDriverURLOk returns a tuple with the JdbcDriverURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetJdbcDriverURLOk() (*string, bool) {
	if o == nil || isNil(o.JdbcDriverURL) {
    return nil, false
	}
	return o.JdbcDriverURL, true
}

// HasJdbcDriverURL returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasJdbcDriverURL() bool {
	if o != nil && !isNil(o.JdbcDriverURL) {
		return true
	}

	return false
}

// SetJdbcDriverURL gets a reference to the given string and assigns it to the JdbcDriverURL field.
func (o *JdbcExternalServerResponse) SetJdbcDriverURL(v string) {
	o.JdbcDriverURL = &v
}

// GetDatabaseName returns the DatabaseName field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetDatabaseName() string {
	if o == nil || isNil(o.DatabaseName) {
		var ret string
		return ret
	}
	return *o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetDatabaseNameOk() (*string, bool) {
	if o == nil || isNil(o.DatabaseName) {
    return nil, false
	}
	return o.DatabaseName, true
}

// HasDatabaseName returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasDatabaseName() bool {
	if o != nil && !isNil(o.DatabaseName) {
		return true
	}

	return false
}

// SetDatabaseName gets a reference to the given string and assigns it to the DatabaseName field.
func (o *JdbcExternalServerResponse) SetDatabaseName(v string) {
	o.DatabaseName = &v
}

// GetServerHostName returns the ServerHostName field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetServerHostName() string {
	if o == nil || isNil(o.ServerHostName) {
		var ret string
		return ret
	}
	return *o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetServerHostNameOk() (*string, bool) {
	if o == nil || isNil(o.ServerHostName) {
    return nil, false
	}
	return o.ServerHostName, true
}

// HasServerHostName returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasServerHostName() bool {
	if o != nil && !isNil(o.ServerHostName) {
		return true
	}

	return false
}

// SetServerHostName gets a reference to the given string and assigns it to the ServerHostName field.
func (o *JdbcExternalServerResponse) SetServerHostName(v string) {
	o.ServerHostName = &v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetServerPort() int32 {
	if o == nil || isNil(o.ServerPort) {
		var ret int32
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetServerPortOk() (*int32, bool) {
	if o == nil || isNil(o.ServerPort) {
    return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasServerPort() bool {
	if o != nil && !isNil(o.ServerPort) {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given int32 and assigns it to the ServerPort field.
func (o *JdbcExternalServerResponse) SetServerPort(v int32) {
	o.ServerPort = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetUserName() string {
	if o == nil || isNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetUserNameOk() (*string, bool) {
	if o == nil || isNil(o.UserName) {
    return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasUserName() bool {
	if o != nil && !isNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *JdbcExternalServerResponse) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *JdbcExternalServerResponse) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetPassphraseProvider() string {
	if o == nil || isNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.PassphraseProvider) {
    return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasPassphraseProvider() bool {
	if o != nil && !isNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *JdbcExternalServerResponse) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetValidationQuery returns the ValidationQuery field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetValidationQuery() string {
	if o == nil || isNil(o.ValidationQuery) {
		var ret string
		return ret
	}
	return *o.ValidationQuery
}

// GetValidationQueryOk returns a tuple with the ValidationQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetValidationQueryOk() (*string, bool) {
	if o == nil || isNil(o.ValidationQuery) {
    return nil, false
	}
	return o.ValidationQuery, true
}

// HasValidationQuery returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasValidationQuery() bool {
	if o != nil && !isNil(o.ValidationQuery) {
		return true
	}

	return false
}

// SetValidationQuery gets a reference to the given string and assigns it to the ValidationQuery field.
func (o *JdbcExternalServerResponse) SetValidationQuery(v string) {
	o.ValidationQuery = &v
}

// GetValidationQueryTimeout returns the ValidationQueryTimeout field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetValidationQueryTimeout() string {
	if o == nil || isNil(o.ValidationQueryTimeout) {
		var ret string
		return ret
	}
	return *o.ValidationQueryTimeout
}

// GetValidationQueryTimeoutOk returns a tuple with the ValidationQueryTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetValidationQueryTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.ValidationQueryTimeout) {
    return nil, false
	}
	return o.ValidationQueryTimeout, true
}

// HasValidationQueryTimeout returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasValidationQueryTimeout() bool {
	if o != nil && !isNil(o.ValidationQueryTimeout) {
		return true
	}

	return false
}

// SetValidationQueryTimeout gets a reference to the given string and assigns it to the ValidationQueryTimeout field.
func (o *JdbcExternalServerResponse) SetValidationQueryTimeout(v string) {
	o.ValidationQueryTimeout = &v
}

// GetJdbcConnectionProperties returns the JdbcConnectionProperties field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetJdbcConnectionProperties() []string {
	if o == nil || isNil(o.JdbcConnectionProperties) {
		var ret []string
		return ret
	}
	return o.JdbcConnectionProperties
}

// GetJdbcConnectionPropertiesOk returns a tuple with the JdbcConnectionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetJdbcConnectionPropertiesOk() ([]string, bool) {
	if o == nil || isNil(o.JdbcConnectionProperties) {
    return nil, false
	}
	return o.JdbcConnectionProperties, true
}

// HasJdbcConnectionProperties returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasJdbcConnectionProperties() bool {
	if o != nil && !isNil(o.JdbcConnectionProperties) {
		return true
	}

	return false
}

// SetJdbcConnectionProperties gets a reference to the given []string and assigns it to the JdbcConnectionProperties field.
func (o *JdbcExternalServerResponse) SetJdbcConnectionProperties(v []string) {
	o.JdbcConnectionProperties = v
}

// GetTransactionIsolationLevel returns the TransactionIsolationLevel field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetTransactionIsolationLevel() EnumexternalServerTransactionIsolationLevelProp {
	if o == nil || isNil(o.TransactionIsolationLevel) {
		var ret EnumexternalServerTransactionIsolationLevelProp
		return ret
	}
	return *o.TransactionIsolationLevel
}

// GetTransactionIsolationLevelOk returns a tuple with the TransactionIsolationLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetTransactionIsolationLevelOk() (*EnumexternalServerTransactionIsolationLevelProp, bool) {
	if o == nil || isNil(o.TransactionIsolationLevel) {
    return nil, false
	}
	return o.TransactionIsolationLevel, true
}

// HasTransactionIsolationLevel returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasTransactionIsolationLevel() bool {
	if o != nil && !isNil(o.TransactionIsolationLevel) {
		return true
	}

	return false
}

// SetTransactionIsolationLevel gets a reference to the given EnumexternalServerTransactionIsolationLevelProp and assigns it to the TransactionIsolationLevel field.
func (o *JdbcExternalServerResponse) SetTransactionIsolationLevel(v EnumexternalServerTransactionIsolationLevelProp) {
	o.TransactionIsolationLevel = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JdbcExternalServerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JdbcExternalServerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JdbcExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

func (o JdbcExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["jdbcDriverType"] = o.JdbcDriverType
	}
	if !isNil(o.JdbcDriverURL) {
		toSerialize["jdbcDriverURL"] = o.JdbcDriverURL
	}
	if !isNil(o.DatabaseName) {
		toSerialize["databaseName"] = o.DatabaseName
	}
	if !isNil(o.ServerHostName) {
		toSerialize["serverHostName"] = o.ServerHostName
	}
	if !isNil(o.ServerPort) {
		toSerialize["serverPort"] = o.ServerPort
	}
	if !isNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.PassphraseProvider) {
		toSerialize["passphraseProvider"] = o.PassphraseProvider
	}
	if !isNil(o.ValidationQuery) {
		toSerialize["validationQuery"] = o.ValidationQuery
	}
	if !isNil(o.ValidationQueryTimeout) {
		toSerialize["validationQueryTimeout"] = o.ValidationQueryTimeout
	}
	if !isNil(o.JdbcConnectionProperties) {
		toSerialize["jdbcConnectionProperties"] = o.JdbcConnectionProperties
	}
	if !isNil(o.TransactionIsolationLevel) {
		toSerialize["transactionIsolationLevel"] = o.TransactionIsolationLevel
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableJdbcExternalServerResponse struct {
	value *JdbcExternalServerResponse
	isSet bool
}

func (v NullableJdbcExternalServerResponse) Get() *JdbcExternalServerResponse {
	return v.value
}

func (v *NullableJdbcExternalServerResponse) Set(val *JdbcExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJdbcExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJdbcExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJdbcExternalServerResponse(val *JdbcExternalServerResponse) *NullableJdbcExternalServerResponse {
	return &NullableJdbcExternalServerResponse{value: val, isSet: true}
}

func (v NullableJdbcExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJdbcExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


