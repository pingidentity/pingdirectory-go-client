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

// OracleUnifiedDirectoryExternalServerShared struct for OracleUnifiedDirectoryExternalServerShared
type OracleUnifiedDirectoryExternalServerShared struct {
	Schemas []EnumoracleUnifiedDirectoryExternalServerSchemaUrn `json:"schemas"`
	// The host name or IP address of the target LDAP server.
	ServerHostName string `json:"serverHostName"`
	// The port number on which the server listens for requests.
	ServerPort int32 `json:"serverPort"`
	// Specifies the location for the LDAP External Server.
	Location *string `json:"location,omitempty"`
	// The DN to use to bind to the target LDAP server if simple authentication is required.
	BindDN *string `json:"bindDN,omitempty"`
	// The login password for the specified user.
	Password *string `json:"password,omitempty"`
	// The passphrase provider to use to obtain the login password for the specified user.
	PassphraseProvider *string `json:"passphraseProvider,omitempty"`
	ConnectionSecurity EnumexternalServerConnectionSecurityProp `json:"connectionSecurity"`
	AuthenticationMethod EnumexternalServerAuthenticationMethodProp `json:"authenticationMethod"`
	VerifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp `json:"verifyCredentialsMethod"`
	// Specifies the maximum length of time to wait for a connection to be established for the purpose of performing a health check. If the connection cannot be established within this length of time, the server will be classified as unavailable.
	HealthCheckConnectTimeout *string `json:"healthCheckConnectTimeout,omitempty"`
	// Specifies the maximum length of time that connections to this server should be allowed to remain established before being closed and replaced with newly-established connections.
	MaxConnectionAge string `json:"maxConnectionAge"`
	// Specifies the minimum length of time that should pass between connection closures as a result of the connections being established for longer than the maximum connection age. This may help avoid cases in which a large number of connections are closed and re-established in a short period of time because of the maximum connection age.
	MinExpiredConnectionDisconnectInterval *string `json:"minExpiredConnectionDisconnectInterval,omitempty"`
	// Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable.
	ConnectTimeout string `json:"connectTimeout"`
	// Specifies the maximum response size that should be supported for messages received from the LDAP external server.
	MaxResponseSize string `json:"maxResponseSize"`
	// The key manager provider to use if SSL or StartTLS is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server's public certificate by adding this server's public certificate to the external server's trust store.
	KeyManagerProvider *string `json:"keyManagerProvider,omitempty"`
	// The trust manager provider to use if SSL or StartTLS is to be used for connection-level security.
	TrustManagerProvider *string `json:"trustManagerProvider,omitempty"`
	// The number of connections to initially establish to the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool.
	InitialConnections *int32 `json:"initialConnections,omitempty"`
	// The maximum number of concurrent connections to maintain for the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool.
	MaxConnections *int32 `json:"maxConnections,omitempty"`
	DefunctConnectionResultCode []EnumexternalServerDefunctConnectionResultCodeProp `json:"defunctConnectionResultCode,omitempty"`
	// Indicates whether to send an abandon request for an operation for which a response timeout is encountered. A request which has timed out on one server may be retried on another server regardless of whether an abandon request is sent, but if the initial attempt is not abandoned then a long-running operation may unnecessarily continue to consume processing resources on the initial server.
	AbandonOnTimeout *bool `json:"abandonOnTimeout,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewOracleUnifiedDirectoryExternalServerShared instantiates a new OracleUnifiedDirectoryExternalServerShared object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOracleUnifiedDirectoryExternalServerShared(schemas []EnumoracleUnifiedDirectoryExternalServerSchemaUrn, serverHostName string, serverPort int32, connectionSecurity EnumexternalServerConnectionSecurityProp, authenticationMethod EnumexternalServerAuthenticationMethodProp, verifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp, maxConnectionAge string, connectTimeout string, maxResponseSize string) *OracleUnifiedDirectoryExternalServerShared {
	this := OracleUnifiedDirectoryExternalServerShared{}
	this.Schemas = schemas
	this.ServerHostName = serverHostName
	this.ServerPort = serverPort
	this.ConnectionSecurity = connectionSecurity
	this.AuthenticationMethod = authenticationMethod
	this.VerifyCredentialsMethod = verifyCredentialsMethod
	this.MaxConnectionAge = maxConnectionAge
	this.ConnectTimeout = connectTimeout
	this.MaxResponseSize = maxResponseSize
	return &this
}

// NewOracleUnifiedDirectoryExternalServerSharedWithDefaults instantiates a new OracleUnifiedDirectoryExternalServerShared object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOracleUnifiedDirectoryExternalServerSharedWithDefaults() *OracleUnifiedDirectoryExternalServerShared {
	this := OracleUnifiedDirectoryExternalServerShared{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetSchemas() []EnumoracleUnifiedDirectoryExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumoracleUnifiedDirectoryExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetSchemasOk() ([]EnumoracleUnifiedDirectoryExternalServerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetSchemas(v []EnumoracleUnifiedDirectoryExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetServerHostName returns the ServerHostName field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetServerHostNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetServerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetServerPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServerPort, true
}

// SetServerPort sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetServerPort(v int32) {
	o.ServerPort = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetLocation() string {
	if o == nil || isNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetLocationOk() (*string, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetLocation(v string) {
	o.Location = &v
}

// GetBindDN returns the BindDN field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetBindDN() string {
	if o == nil || isNil(o.BindDN) {
		var ret string
		return ret
	}
	return *o.BindDN
}

// GetBindDNOk returns a tuple with the BindDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetBindDNOk() (*string, bool) {
	if o == nil || isNil(o.BindDN) {
    return nil, false
	}
	return o.BindDN, true
}

// HasBindDN returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasBindDN() bool {
	if o != nil && !isNil(o.BindDN) {
		return true
	}

	return false
}

// SetBindDN gets a reference to the given string and assigns it to the BindDN field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetBindDN(v string) {
	o.BindDN = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetPassphraseProvider() string {
	if o == nil || isNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.PassphraseProvider) {
    return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasPassphraseProvider() bool {
	if o != nil && !isNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetConnectionSecurity returns the ConnectionSecurity field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp {
	if o == nil {
		var ret EnumexternalServerConnectionSecurityProp
		return ret
	}

	return o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionSecurity, true
}

// SetConnectionSecurity sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp) {
	o.ConnectionSecurity = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetAuthenticationMethod() EnumexternalServerAuthenticationMethodProp {
	if o == nil {
		var ret EnumexternalServerAuthenticationMethodProp
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetAuthenticationMethodOk() (*EnumexternalServerAuthenticationMethodProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetAuthenticationMethod(v EnumexternalServerAuthenticationMethodProp) {
	o.AuthenticationMethod = v
}

// GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp {
	if o == nil {
		var ret EnumexternalServerVerifyCredentialsMethodProp
		return ret
	}

	return o.VerifyCredentialsMethod
}

// GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VerifyCredentialsMethod, true
}

// SetVerifyCredentialsMethod sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp) {
	o.VerifyCredentialsMethod = v
}

// GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetHealthCheckConnectTimeout() string {
	if o == nil || isNil(o.HealthCheckConnectTimeout) {
		var ret string
		return ret
	}
	return *o.HealthCheckConnectTimeout
}

// GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetHealthCheckConnectTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.HealthCheckConnectTimeout) {
    return nil, false
	}
	return o.HealthCheckConnectTimeout, true
}

// HasHealthCheckConnectTimeout returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasHealthCheckConnectTimeout() bool {
	if o != nil && !isNil(o.HealthCheckConnectTimeout) {
		return true
	}

	return false
}

// SetHealthCheckConnectTimeout gets a reference to the given string and assigns it to the HealthCheckConnectTimeout field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetHealthCheckConnectTimeout(v string) {
	o.HealthCheckConnectTimeout = &v
}

// GetMaxConnectionAge returns the MaxConnectionAge field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetMaxConnectionAge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxConnectionAge
}

// GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetMaxConnectionAgeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxConnectionAge, true
}

// SetMaxConnectionAge sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetMaxConnectionAge(v string) {
	o.MaxConnectionAge = v
}

// GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetMinExpiredConnectionDisconnectInterval() string {
	if o == nil || isNil(o.MinExpiredConnectionDisconnectInterval) {
		var ret string
		return ret
	}
	return *o.MinExpiredConnectionDisconnectInterval
}

// GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool) {
	if o == nil || isNil(o.MinExpiredConnectionDisconnectInterval) {
    return nil, false
	}
	return o.MinExpiredConnectionDisconnectInterval, true
}

// HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasMinExpiredConnectionDisconnectInterval() bool {
	if o != nil && !isNil(o.MinExpiredConnectionDisconnectInterval) {
		return true
	}

	return false
}

// SetMinExpiredConnectionDisconnectInterval gets a reference to the given string and assigns it to the MinExpiredConnectionDisconnectInterval field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetMinExpiredConnectionDisconnectInterval(v string) {
	o.MinExpiredConnectionDisconnectInterval = &v
}

// GetConnectTimeout returns the ConnectTimeout field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetConnectTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetConnectTimeoutOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectTimeout, true
}

// SetConnectTimeout sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetConnectTimeout(v string) {
	o.ConnectTimeout = v
}

// GetMaxResponseSize returns the MaxResponseSize field value
func (o *OracleUnifiedDirectoryExternalServerShared) GetMaxResponseSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxResponseSize
}

// GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field value
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetMaxResponseSizeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxResponseSize, true
}

// SetMaxResponseSize sets field value
func (o *OracleUnifiedDirectoryExternalServerShared) SetMaxResponseSize(v string) {
	o.MaxResponseSize = v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetKeyManagerProvider() string {
	if o == nil || isNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || isNil(o.KeyManagerProvider) {
    return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasKeyManagerProvider() bool {
	if o != nil && !isNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetTrustManagerProvider() string {
	if o == nil || isNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || isNil(o.TrustManagerProvider) {
    return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasTrustManagerProvider() bool {
	if o != nil && !isNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetInitialConnections returns the InitialConnections field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetInitialConnections() int32 {
	if o == nil || isNil(o.InitialConnections) {
		var ret int32
		return ret
	}
	return *o.InitialConnections
}

// GetInitialConnectionsOk returns a tuple with the InitialConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetInitialConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.InitialConnections) {
    return nil, false
	}
	return o.InitialConnections, true
}

// HasInitialConnections returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasInitialConnections() bool {
	if o != nil && !isNil(o.InitialConnections) {
		return true
	}

	return false
}

// SetInitialConnections gets a reference to the given int32 and assigns it to the InitialConnections field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetInitialConnections(v int32) {
	o.InitialConnections = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetMaxConnections() int32 {
	if o == nil || isNil(o.MaxConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetMaxConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxConnections) {
    return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasMaxConnections() bool {
	if o != nil && !isNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int32 and assigns it to the MaxConnections field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetMaxConnections(v int32) {
	o.MaxConnections = &v
}

// GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp {
	if o == nil || isNil(o.DefunctConnectionResultCode) {
		var ret []EnumexternalServerDefunctConnectionResultCodeProp
		return ret
	}
	return o.DefunctConnectionResultCode
}

// GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetDefunctConnectionResultCodeOk() ([]EnumexternalServerDefunctConnectionResultCodeProp, bool) {
	if o == nil || isNil(o.DefunctConnectionResultCode) {
    return nil, false
	}
	return o.DefunctConnectionResultCode, true
}

// HasDefunctConnectionResultCode returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasDefunctConnectionResultCode() bool {
	if o != nil && !isNil(o.DefunctConnectionResultCode) {
		return true
	}

	return false
}

// SetDefunctConnectionResultCode gets a reference to the given []EnumexternalServerDefunctConnectionResultCodeProp and assigns it to the DefunctConnectionResultCode field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp) {
	o.DefunctConnectionResultCode = v
}

// GetAbandonOnTimeout returns the AbandonOnTimeout field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetAbandonOnTimeout() bool {
	if o == nil || isNil(o.AbandonOnTimeout) {
		var ret bool
		return ret
	}
	return *o.AbandonOnTimeout
}

// GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetAbandonOnTimeoutOk() (*bool, bool) {
	if o == nil || isNil(o.AbandonOnTimeout) {
    return nil, false
	}
	return o.AbandonOnTimeout, true
}

// HasAbandonOnTimeout returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasAbandonOnTimeout() bool {
	if o != nil && !isNil(o.AbandonOnTimeout) {
		return true
	}

	return false
}

// SetAbandonOnTimeout gets a reference to the given bool and assigns it to the AbandonOnTimeout field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetAbandonOnTimeout(v bool) {
	o.AbandonOnTimeout = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OracleUnifiedDirectoryExternalServerShared) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OracleUnifiedDirectoryExternalServerShared) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OracleUnifiedDirectoryExternalServerShared) SetDescription(v string) {
	o.Description = &v
}

func (o OracleUnifiedDirectoryExternalServerShared) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["serverHostName"] = o.ServerHostName
	}
	if true {
		toSerialize["serverPort"] = o.ServerPort
	}
	if !isNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !isNil(o.BindDN) {
		toSerialize["bindDN"] = o.BindDN
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.PassphraseProvider) {
		toSerialize["passphraseProvider"] = o.PassphraseProvider
	}
	if true {
		toSerialize["connectionSecurity"] = o.ConnectionSecurity
	}
	if true {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	if true {
		toSerialize["verifyCredentialsMethod"] = o.VerifyCredentialsMethod
	}
	if !isNil(o.HealthCheckConnectTimeout) {
		toSerialize["healthCheckConnectTimeout"] = o.HealthCheckConnectTimeout
	}
	if true {
		toSerialize["maxConnectionAge"] = o.MaxConnectionAge
	}
	if !isNil(o.MinExpiredConnectionDisconnectInterval) {
		toSerialize["minExpiredConnectionDisconnectInterval"] = o.MinExpiredConnectionDisconnectInterval
	}
	if true {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if true {
		toSerialize["maxResponseSize"] = o.MaxResponseSize
	}
	if !isNil(o.KeyManagerProvider) {
		toSerialize["keyManagerProvider"] = o.KeyManagerProvider
	}
	if !isNil(o.TrustManagerProvider) {
		toSerialize["trustManagerProvider"] = o.TrustManagerProvider
	}
	if !isNil(o.InitialConnections) {
		toSerialize["initialConnections"] = o.InitialConnections
	}
	if !isNil(o.MaxConnections) {
		toSerialize["maxConnections"] = o.MaxConnections
	}
	if !isNil(o.DefunctConnectionResultCode) {
		toSerialize["defunctConnectionResultCode"] = o.DefunctConnectionResultCode
	}
	if !isNil(o.AbandonOnTimeout) {
		toSerialize["abandonOnTimeout"] = o.AbandonOnTimeout
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableOracleUnifiedDirectoryExternalServerShared struct {
	value *OracleUnifiedDirectoryExternalServerShared
	isSet bool
}

func (v NullableOracleUnifiedDirectoryExternalServerShared) Get() *OracleUnifiedDirectoryExternalServerShared {
	return v.value
}

func (v *NullableOracleUnifiedDirectoryExternalServerShared) Set(val *OracleUnifiedDirectoryExternalServerShared) {
	v.value = val
	v.isSet = true
}

func (v NullableOracleUnifiedDirectoryExternalServerShared) IsSet() bool {
	return v.isSet
}

func (v *NullableOracleUnifiedDirectoryExternalServerShared) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOracleUnifiedDirectoryExternalServerShared(val *OracleUnifiedDirectoryExternalServerShared) *NullableOracleUnifiedDirectoryExternalServerShared {
	return &NullableOracleUnifiedDirectoryExternalServerShared{value: val, isSet: true}
}

func (v NullableOracleUnifiedDirectoryExternalServerShared) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOracleUnifiedDirectoryExternalServerShared) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


