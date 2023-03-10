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

// checks if the AddOpendjExternalServerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOpendjExternalServerRequest{}

// AddOpendjExternalServerRequest struct for AddOpendjExternalServerRequest
type AddOpendjExternalServerRequest struct {
	// Name of the new External Server
	ServerName string                              `json:"serverName"`
	Schemas    []EnumopendjExternalServerSchemaUrn `json:"schemas"`
	// The host name or IP address of the target LDAP server.
	ServerHostName string `json:"serverHostName"`
	// The port number on which the server listens for requests.
	ServerPort *int32 `json:"serverPort,omitempty"`
	// Specifies the location for the LDAP External Server.
	Location *string `json:"location,omitempty"`
	// The DN to use to bind to the target LDAP server if simple authentication is required.
	BindDN *string `json:"bindDN,omitempty"`
	// The login password for the specified user.
	Password *string `json:"password,omitempty"`
	// The passphrase provider to use to obtain the login password for the specified user.
	PassphraseProvider      *string                                           `json:"passphraseProvider,omitempty"`
	ConnectionSecurity      *EnumexternalServerConnectionSecurityProp         `json:"connectionSecurity,omitempty"`
	AuthenticationMethod    *EnumexternalServerOpendjAuthenticationMethodProp `json:"authenticationMethod,omitempty"`
	VerifyCredentialsMethod *EnumexternalServerVerifyCredentialsMethodProp    `json:"verifyCredentialsMethod,omitempty"`
	// Specifies the maximum length of time to wait for a connection to be established for the purpose of performing a health check. If the connection cannot be established within this length of time, the server will be classified as unavailable.
	HealthCheckConnectTimeout *string `json:"healthCheckConnectTimeout,omitempty"`
	// Specifies the maximum length of time that connections to this server should be allowed to remain established before being closed and replaced with newly-established connections.
	MaxConnectionAge *string `json:"maxConnectionAge,omitempty"`
	// Specifies the minimum length of time that should pass between connection closures as a result of the connections being established for longer than the maximum connection age. This may help avoid cases in which a large number of connections are closed and re-established in a short period of time because of the maximum connection age.
	MinExpiredConnectionDisconnectInterval *string `json:"minExpiredConnectionDisconnectInterval,omitempty"`
	// Specifies the maximum length of time to wait for a connection to be established before giving up and considering the server unavailable.
	ConnectTimeout *string `json:"connectTimeout,omitempty"`
	// Specifies the maximum response size that should be supported for messages received from the LDAP external server.
	MaxResponseSize *string `json:"maxResponseSize,omitempty"`
	// The key manager provider to use if SSL or StartTLS is to be used for connection-level security. When specifying a value for this property (except when using the Null key manager provider) you must ensure that the external server trusts this server's public certificate by adding this server's public certificate to the external server's trust store.
	KeyManagerProvider *string `json:"keyManagerProvider,omitempty"`
	// The trust manager provider to use if SSL or StartTLS is to be used for connection-level security.
	TrustManagerProvider *string `json:"trustManagerProvider,omitempty"`
	// The number of connections to initially establish to the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool.
	InitialConnections *int32 `json:"initialConnections,omitempty"`
	// The maximum number of concurrent connections to maintain for the LDAP external server. A value of zero indicates that the number of connections should be dynamically based on the number of available worker threads. This will be ignored when using a thread-local connection pool.
	MaxConnections              *int32                                              `json:"maxConnections,omitempty"`
	DefunctConnectionResultCode []EnumexternalServerDefunctConnectionResultCodeProp `json:"defunctConnectionResultCode,omitempty"`
	// Indicates whether to send an abandon request for an operation for which a response timeout is encountered. A request which has timed out on one server may be retried on another server regardless of whether an abandon request is sent, but if the initial attempt is not abandoned then a long-running operation may unnecessarily continue to consume processing resources on the initial server.
	AbandonOnTimeout *bool `json:"abandonOnTimeout,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewAddOpendjExternalServerRequest instantiates a new AddOpendjExternalServerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOpendjExternalServerRequest(serverName string, schemas []EnumopendjExternalServerSchemaUrn, serverHostName string) *AddOpendjExternalServerRequest {
	this := AddOpendjExternalServerRequest{}
	this.ServerName = serverName
	this.Schemas = schemas
	this.ServerHostName = serverHostName
	return &this
}

// NewAddOpendjExternalServerRequestWithDefaults instantiates a new AddOpendjExternalServerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOpendjExternalServerRequestWithDefaults() *AddOpendjExternalServerRequest {
	this := AddOpendjExternalServerRequest{}
	return &this
}

// GetServerName returns the ServerName field value
func (o *AddOpendjExternalServerRequest) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *AddOpendjExternalServerRequest) SetServerName(v string) {
	o.ServerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddOpendjExternalServerRequest) GetSchemas() []EnumopendjExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumopendjExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetSchemasOk() ([]EnumopendjExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddOpendjExternalServerRequest) SetSchemas(v []EnumopendjExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetServerHostName returns the ServerHostName field value
func (o *AddOpendjExternalServerRequest) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetServerHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *AddOpendjExternalServerRequest) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetServerPort() int32 {
	if o == nil || IsNil(o.ServerPort) {
		var ret int32
		return ret
	}
	return *o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetServerPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerPort) {
		return nil, false
	}
	return o.ServerPort, true
}

// HasServerPort returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasServerPort() bool {
	if o != nil && !IsNil(o.ServerPort) {
		return true
	}

	return false
}

// SetServerPort gets a reference to the given int32 and assigns it to the ServerPort field.
func (o *AddOpendjExternalServerRequest) SetServerPort(v int32) {
	o.ServerPort = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AddOpendjExternalServerRequest) SetLocation(v string) {
	o.Location = &v
}

// GetBindDN returns the BindDN field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetBindDN() string {
	if o == nil || IsNil(o.BindDN) {
		var ret string
		return ret
	}
	return *o.BindDN
}

// GetBindDNOk returns a tuple with the BindDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetBindDNOk() (*string, bool) {
	if o == nil || IsNil(o.BindDN) {
		return nil, false
	}
	return o.BindDN, true
}

// HasBindDN returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasBindDN() bool {
	if o != nil && !IsNil(o.BindDN) {
		return true
	}

	return false
}

// SetBindDN gets a reference to the given string and assigns it to the BindDN field.
func (o *AddOpendjExternalServerRequest) SetBindDN(v string) {
	o.BindDN = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AddOpendjExternalServerRequest) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetPassphraseProvider() string {
	if o == nil || IsNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.PassphraseProvider) {
		return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasPassphraseProvider() bool {
	if o != nil && !IsNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *AddOpendjExternalServerRequest) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetConnectionSecurity returns the ConnectionSecurity field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp {
	if o == nil || IsNil(o.ConnectionSecurity) {
		var ret EnumexternalServerConnectionSecurityProp
		return ret
	}
	return *o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool) {
	if o == nil || IsNil(o.ConnectionSecurity) {
		return nil, false
	}
	return o.ConnectionSecurity, true
}

// HasConnectionSecurity returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasConnectionSecurity() bool {
	if o != nil && !IsNil(o.ConnectionSecurity) {
		return true
	}

	return false
}

// SetConnectionSecurity gets a reference to the given EnumexternalServerConnectionSecurityProp and assigns it to the ConnectionSecurity field.
func (o *AddOpendjExternalServerRequest) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp) {
	o.ConnectionSecurity = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetAuthenticationMethod() EnumexternalServerOpendjAuthenticationMethodProp {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret EnumexternalServerOpendjAuthenticationMethodProp
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetAuthenticationMethodOk() (*EnumexternalServerOpendjAuthenticationMethodProp, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given EnumexternalServerOpendjAuthenticationMethodProp and assigns it to the AuthenticationMethod field.
func (o *AddOpendjExternalServerRequest) SetAuthenticationMethod(v EnumexternalServerOpendjAuthenticationMethodProp) {
	o.AuthenticationMethod = &v
}

// GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp {
	if o == nil || IsNil(o.VerifyCredentialsMethod) {
		var ret EnumexternalServerVerifyCredentialsMethodProp
		return ret
	}
	return *o.VerifyCredentialsMethod
}

// GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool) {
	if o == nil || IsNil(o.VerifyCredentialsMethod) {
		return nil, false
	}
	return o.VerifyCredentialsMethod, true
}

// HasVerifyCredentialsMethod returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasVerifyCredentialsMethod() bool {
	if o != nil && !IsNil(o.VerifyCredentialsMethod) {
		return true
	}

	return false
}

// SetVerifyCredentialsMethod gets a reference to the given EnumexternalServerVerifyCredentialsMethodProp and assigns it to the VerifyCredentialsMethod field.
func (o *AddOpendjExternalServerRequest) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp) {
	o.VerifyCredentialsMethod = &v
}

// GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetHealthCheckConnectTimeout() string {
	if o == nil || IsNil(o.HealthCheckConnectTimeout) {
		var ret string
		return ret
	}
	return *o.HealthCheckConnectTimeout
}

// GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetHealthCheckConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.HealthCheckConnectTimeout) {
		return nil, false
	}
	return o.HealthCheckConnectTimeout, true
}

// HasHealthCheckConnectTimeout returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasHealthCheckConnectTimeout() bool {
	if o != nil && !IsNil(o.HealthCheckConnectTimeout) {
		return true
	}

	return false
}

// SetHealthCheckConnectTimeout gets a reference to the given string and assigns it to the HealthCheckConnectTimeout field.
func (o *AddOpendjExternalServerRequest) SetHealthCheckConnectTimeout(v string) {
	o.HealthCheckConnectTimeout = &v
}

// GetMaxConnectionAge returns the MaxConnectionAge field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetMaxConnectionAge() string {
	if o == nil || IsNil(o.MaxConnectionAge) {
		var ret string
		return ret
	}
	return *o.MaxConnectionAge
}

// GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetMaxConnectionAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxConnectionAge) {
		return nil, false
	}
	return o.MaxConnectionAge, true
}

// HasMaxConnectionAge returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasMaxConnectionAge() bool {
	if o != nil && !IsNil(o.MaxConnectionAge) {
		return true
	}

	return false
}

// SetMaxConnectionAge gets a reference to the given string and assigns it to the MaxConnectionAge field.
func (o *AddOpendjExternalServerRequest) SetMaxConnectionAge(v string) {
	o.MaxConnectionAge = &v
}

// GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetMinExpiredConnectionDisconnectInterval() string {
	if o == nil || IsNil(o.MinExpiredConnectionDisconnectInterval) {
		var ret string
		return ret
	}
	return *o.MinExpiredConnectionDisconnectInterval
}

// GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.MinExpiredConnectionDisconnectInterval) {
		return nil, false
	}
	return o.MinExpiredConnectionDisconnectInterval, true
}

// HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasMinExpiredConnectionDisconnectInterval() bool {
	if o != nil && !IsNil(o.MinExpiredConnectionDisconnectInterval) {
		return true
	}

	return false
}

// SetMinExpiredConnectionDisconnectInterval gets a reference to the given string and assigns it to the MinExpiredConnectionDisconnectInterval field.
func (o *AddOpendjExternalServerRequest) SetMinExpiredConnectionDisconnectInterval(v string) {
	o.MinExpiredConnectionDisconnectInterval = &v
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetConnectTimeout() string {
	if o == nil || IsNil(o.ConnectTimeout) {
		var ret string
		return ret
	}
	return *o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectTimeout) {
		return nil, false
	}
	return o.ConnectTimeout, true
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasConnectTimeout() bool {
	if o != nil && !IsNil(o.ConnectTimeout) {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given string and assigns it to the ConnectTimeout field.
func (o *AddOpendjExternalServerRequest) SetConnectTimeout(v string) {
	o.ConnectTimeout = &v
}

// GetMaxResponseSize returns the MaxResponseSize field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetMaxResponseSize() string {
	if o == nil || IsNil(o.MaxResponseSize) {
		var ret string
		return ret
	}
	return *o.MaxResponseSize
}

// GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetMaxResponseSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxResponseSize) {
		return nil, false
	}
	return o.MaxResponseSize, true
}

// HasMaxResponseSize returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasMaxResponseSize() bool {
	if o != nil && !IsNil(o.MaxResponseSize) {
		return true
	}

	return false
}

// SetMaxResponseSize gets a reference to the given string and assigns it to the MaxResponseSize field.
func (o *AddOpendjExternalServerRequest) SetMaxResponseSize(v string) {
	o.MaxResponseSize = &v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetKeyManagerProvider() string {
	if o == nil || IsNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyManagerProvider) {
		return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasKeyManagerProvider() bool {
	if o != nil && !IsNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *AddOpendjExternalServerRequest) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetTrustManagerProvider() string {
	if o == nil || IsNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasTrustManagerProvider() bool {
	if o != nil && !IsNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *AddOpendjExternalServerRequest) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetInitialConnections returns the InitialConnections field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetInitialConnections() int32 {
	if o == nil || IsNil(o.InitialConnections) {
		var ret int32
		return ret
	}
	return *o.InitialConnections
}

// GetInitialConnectionsOk returns a tuple with the InitialConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetInitialConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.InitialConnections) {
		return nil, false
	}
	return o.InitialConnections, true
}

// HasInitialConnections returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasInitialConnections() bool {
	if o != nil && !IsNil(o.InitialConnections) {
		return true
	}

	return false
}

// SetInitialConnections gets a reference to the given int32 and assigns it to the InitialConnections field.
func (o *AddOpendjExternalServerRequest) SetInitialConnections(v int32) {
	o.InitialConnections = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetMaxConnections() int32 {
	if o == nil || IsNil(o.MaxConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetMaxConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int32 and assigns it to the MaxConnections field.
func (o *AddOpendjExternalServerRequest) SetMaxConnections(v int32) {
	o.MaxConnections = &v
}

// GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp {
	if o == nil || IsNil(o.DefunctConnectionResultCode) {
		var ret []EnumexternalServerDefunctConnectionResultCodeProp
		return ret
	}
	return o.DefunctConnectionResultCode
}

// GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetDefunctConnectionResultCodeOk() ([]EnumexternalServerDefunctConnectionResultCodeProp, bool) {
	if o == nil || IsNil(o.DefunctConnectionResultCode) {
		return nil, false
	}
	return o.DefunctConnectionResultCode, true
}

// HasDefunctConnectionResultCode returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasDefunctConnectionResultCode() bool {
	if o != nil && !IsNil(o.DefunctConnectionResultCode) {
		return true
	}

	return false
}

// SetDefunctConnectionResultCode gets a reference to the given []EnumexternalServerDefunctConnectionResultCodeProp and assigns it to the DefunctConnectionResultCode field.
func (o *AddOpendjExternalServerRequest) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp) {
	o.DefunctConnectionResultCode = v
}

// GetAbandonOnTimeout returns the AbandonOnTimeout field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetAbandonOnTimeout() bool {
	if o == nil || IsNil(o.AbandonOnTimeout) {
		var ret bool
		return ret
	}
	return *o.AbandonOnTimeout
}

// GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetAbandonOnTimeoutOk() (*bool, bool) {
	if o == nil || IsNil(o.AbandonOnTimeout) {
		return nil, false
	}
	return o.AbandonOnTimeout, true
}

// HasAbandonOnTimeout returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasAbandonOnTimeout() bool {
	if o != nil && !IsNil(o.AbandonOnTimeout) {
		return true
	}

	return false
}

// SetAbandonOnTimeout gets a reference to the given bool and assigns it to the AbandonOnTimeout field.
func (o *AddOpendjExternalServerRequest) SetAbandonOnTimeout(v bool) {
	o.AbandonOnTimeout = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddOpendjExternalServerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOpendjExternalServerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddOpendjExternalServerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddOpendjExternalServerRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddOpendjExternalServerRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOpendjExternalServerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverName"] = o.ServerName
	toSerialize["schemas"] = o.Schemas
	toSerialize["serverHostName"] = o.ServerHostName
	if !IsNil(o.ServerPort) {
		toSerialize["serverPort"] = o.ServerPort
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.BindDN) {
		toSerialize["bindDN"] = o.BindDN
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.PassphraseProvider) {
		toSerialize["passphraseProvider"] = o.PassphraseProvider
	}
	if !IsNil(o.ConnectionSecurity) {
		toSerialize["connectionSecurity"] = o.ConnectionSecurity
	}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	if !IsNil(o.VerifyCredentialsMethod) {
		toSerialize["verifyCredentialsMethod"] = o.VerifyCredentialsMethod
	}
	if !IsNil(o.HealthCheckConnectTimeout) {
		toSerialize["healthCheckConnectTimeout"] = o.HealthCheckConnectTimeout
	}
	if !IsNil(o.MaxConnectionAge) {
		toSerialize["maxConnectionAge"] = o.MaxConnectionAge
	}
	if !IsNil(o.MinExpiredConnectionDisconnectInterval) {
		toSerialize["minExpiredConnectionDisconnectInterval"] = o.MinExpiredConnectionDisconnectInterval
	}
	if !IsNil(o.ConnectTimeout) {
		toSerialize["connectTimeout"] = o.ConnectTimeout
	}
	if !IsNil(o.MaxResponseSize) {
		toSerialize["maxResponseSize"] = o.MaxResponseSize
	}
	if !IsNil(o.KeyManagerProvider) {
		toSerialize["keyManagerProvider"] = o.KeyManagerProvider
	}
	if !IsNil(o.TrustManagerProvider) {
		toSerialize["trustManagerProvider"] = o.TrustManagerProvider
	}
	if !IsNil(o.InitialConnections) {
		toSerialize["initialConnections"] = o.InitialConnections
	}
	if !IsNil(o.MaxConnections) {
		toSerialize["maxConnections"] = o.MaxConnections
	}
	if !IsNil(o.DefunctConnectionResultCode) {
		toSerialize["defunctConnectionResultCode"] = o.DefunctConnectionResultCode
	}
	if !IsNil(o.AbandonOnTimeout) {
		toSerialize["abandonOnTimeout"] = o.AbandonOnTimeout
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAddOpendjExternalServerRequest struct {
	value *AddOpendjExternalServerRequest
	isSet bool
}

func (v NullableAddOpendjExternalServerRequest) Get() *AddOpendjExternalServerRequest {
	return v.value
}

func (v *NullableAddOpendjExternalServerRequest) Set(val *AddOpendjExternalServerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOpendjExternalServerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOpendjExternalServerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOpendjExternalServerRequest(val *AddOpendjExternalServerRequest) *NullableAddOpendjExternalServerRequest {
	return &NullableAddOpendjExternalServerRequest{value: val, isSet: true}
}

func (v NullableAddOpendjExternalServerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOpendjExternalServerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
