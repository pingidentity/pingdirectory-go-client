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

// NokiaProxyServerExternalServerResponse struct for NokiaProxyServerExternalServerResponse
type NokiaProxyServerExternalServerResponse struct {
	// Name of the External Server
	Id string `json:"id"`
	Schemas []EnumnokiaProxyServerExternalServerSchemaUrn `json:"schemas"`
	VerifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp `json:"verifyCredentialsMethod"`
	// Indicates whether to include the administrative operation request control in requests sent to this server which are intended for administrative operations (e.g., health checking) rather than requests directly from clients.
	UseAdministrativeOperationControl *bool `json:"useAdministrativeOperationControl,omitempty"`
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
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewNokiaProxyServerExternalServerResponse instantiates a new NokiaProxyServerExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNokiaProxyServerExternalServerResponse(id string, schemas []EnumnokiaProxyServerExternalServerSchemaUrn, verifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp, serverHostName string, serverPort int32, connectionSecurity EnumexternalServerConnectionSecurityProp, authenticationMethod EnumexternalServerAuthenticationMethodProp, maxConnectionAge string, connectTimeout string, maxResponseSize string) *NokiaProxyServerExternalServerResponse {
	this := NokiaProxyServerExternalServerResponse{}
	this.Id = id
	this.Schemas = schemas
	this.VerifyCredentialsMethod = verifyCredentialsMethod
	this.ServerHostName = serverHostName
	this.ServerPort = serverPort
	this.ConnectionSecurity = connectionSecurity
	this.AuthenticationMethod = authenticationMethod
	this.MaxConnectionAge = maxConnectionAge
	this.ConnectTimeout = connectTimeout
	this.MaxResponseSize = maxResponseSize
	return &this
}

// NewNokiaProxyServerExternalServerResponseWithDefaults instantiates a new NokiaProxyServerExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNokiaProxyServerExternalServerResponseWithDefaults() *NokiaProxyServerExternalServerResponse {
	this := NokiaProxyServerExternalServerResponse{}
	return &this
}

// GetId returns the Id field value
func (o *NokiaProxyServerExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NokiaProxyServerExternalServerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *NokiaProxyServerExternalServerResponse) GetSchemas() []EnumnokiaProxyServerExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumnokiaProxyServerExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetSchemasOk() ([]EnumnokiaProxyServerExternalServerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *NokiaProxyServerExternalServerResponse) SetSchemas(v []EnumnokiaProxyServerExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field value
func (o *NokiaProxyServerExternalServerResponse) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp {
	if o == nil {
		var ret EnumexternalServerVerifyCredentialsMethodProp
		return ret
	}

	return o.VerifyCredentialsMethod
}

// GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VerifyCredentialsMethod, true
}

// SetVerifyCredentialsMethod sets field value
func (o *NokiaProxyServerExternalServerResponse) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp) {
	o.VerifyCredentialsMethod = v
}

// GetUseAdministrativeOperationControl returns the UseAdministrativeOperationControl field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetUseAdministrativeOperationControl() bool {
	if o == nil || isNil(o.UseAdministrativeOperationControl) {
		var ret bool
		return ret
	}
	return *o.UseAdministrativeOperationControl
}

// GetUseAdministrativeOperationControlOk returns a tuple with the UseAdministrativeOperationControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetUseAdministrativeOperationControlOk() (*bool, bool) {
	if o == nil || isNil(o.UseAdministrativeOperationControl) {
    return nil, false
	}
	return o.UseAdministrativeOperationControl, true
}

// HasUseAdministrativeOperationControl returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasUseAdministrativeOperationControl() bool {
	if o != nil && !isNil(o.UseAdministrativeOperationControl) {
		return true
	}

	return false
}

// SetUseAdministrativeOperationControl gets a reference to the given bool and assigns it to the UseAdministrativeOperationControl field.
func (o *NokiaProxyServerExternalServerResponse) SetUseAdministrativeOperationControl(v bool) {
	o.UseAdministrativeOperationControl = &v
}

// GetServerHostName returns the ServerHostName field value
func (o *NokiaProxyServerExternalServerResponse) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetServerHostNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *NokiaProxyServerExternalServerResponse) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value
func (o *NokiaProxyServerExternalServerResponse) GetServerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetServerPortOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServerPort, true
}

// SetServerPort sets field value
func (o *NokiaProxyServerExternalServerResponse) SetServerPort(v int32) {
	o.ServerPort = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetLocation() string {
	if o == nil || isNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetLocationOk() (*string, bool) {
	if o == nil || isNil(o.Location) {
    return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasLocation() bool {
	if o != nil && !isNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *NokiaProxyServerExternalServerResponse) SetLocation(v string) {
	o.Location = &v
}

// GetBindDN returns the BindDN field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetBindDN() string {
	if o == nil || isNil(o.BindDN) {
		var ret string
		return ret
	}
	return *o.BindDN
}

// GetBindDNOk returns a tuple with the BindDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetBindDNOk() (*string, bool) {
	if o == nil || isNil(o.BindDN) {
    return nil, false
	}
	return o.BindDN, true
}

// HasBindDN returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasBindDN() bool {
	if o != nil && !isNil(o.BindDN) {
		return true
	}

	return false
}

// SetBindDN gets a reference to the given string and assigns it to the BindDN field.
func (o *NokiaProxyServerExternalServerResponse) SetBindDN(v string) {
	o.BindDN = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NokiaProxyServerExternalServerResponse) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetPassphraseProvider() string {
	if o == nil || isNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || isNil(o.PassphraseProvider) {
    return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasPassphraseProvider() bool {
	if o != nil && !isNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *NokiaProxyServerExternalServerResponse) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetConnectionSecurity returns the ConnectionSecurity field value
func (o *NokiaProxyServerExternalServerResponse) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp {
	if o == nil {
		var ret EnumexternalServerConnectionSecurityProp
		return ret
	}

	return o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectionSecurity, true
}

// SetConnectionSecurity sets field value
func (o *NokiaProxyServerExternalServerResponse) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp) {
	o.ConnectionSecurity = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *NokiaProxyServerExternalServerResponse) GetAuthenticationMethod() EnumexternalServerAuthenticationMethodProp {
	if o == nil {
		var ret EnumexternalServerAuthenticationMethodProp
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetAuthenticationMethodOk() (*EnumexternalServerAuthenticationMethodProp, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *NokiaProxyServerExternalServerResponse) SetAuthenticationMethod(v EnumexternalServerAuthenticationMethodProp) {
	o.AuthenticationMethod = v
}

// GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetHealthCheckConnectTimeout() string {
	if o == nil || isNil(o.HealthCheckConnectTimeout) {
		var ret string
		return ret
	}
	return *o.HealthCheckConnectTimeout
}

// GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetHealthCheckConnectTimeoutOk() (*string, bool) {
	if o == nil || isNil(o.HealthCheckConnectTimeout) {
    return nil, false
	}
	return o.HealthCheckConnectTimeout, true
}

// HasHealthCheckConnectTimeout returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasHealthCheckConnectTimeout() bool {
	if o != nil && !isNil(o.HealthCheckConnectTimeout) {
		return true
	}

	return false
}

// SetHealthCheckConnectTimeout gets a reference to the given string and assigns it to the HealthCheckConnectTimeout field.
func (o *NokiaProxyServerExternalServerResponse) SetHealthCheckConnectTimeout(v string) {
	o.HealthCheckConnectTimeout = &v
}

// GetMaxConnectionAge returns the MaxConnectionAge field value
func (o *NokiaProxyServerExternalServerResponse) GetMaxConnectionAge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxConnectionAge
}

// GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetMaxConnectionAgeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxConnectionAge, true
}

// SetMaxConnectionAge sets field value
func (o *NokiaProxyServerExternalServerResponse) SetMaxConnectionAge(v string) {
	o.MaxConnectionAge = v
}

// GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetMinExpiredConnectionDisconnectInterval() string {
	if o == nil || isNil(o.MinExpiredConnectionDisconnectInterval) {
		var ret string
		return ret
	}
	return *o.MinExpiredConnectionDisconnectInterval
}

// GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool) {
	if o == nil || isNil(o.MinExpiredConnectionDisconnectInterval) {
    return nil, false
	}
	return o.MinExpiredConnectionDisconnectInterval, true
}

// HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasMinExpiredConnectionDisconnectInterval() bool {
	if o != nil && !isNil(o.MinExpiredConnectionDisconnectInterval) {
		return true
	}

	return false
}

// SetMinExpiredConnectionDisconnectInterval gets a reference to the given string and assigns it to the MinExpiredConnectionDisconnectInterval field.
func (o *NokiaProxyServerExternalServerResponse) SetMinExpiredConnectionDisconnectInterval(v string) {
	o.MinExpiredConnectionDisconnectInterval = &v
}

// GetConnectTimeout returns the ConnectTimeout field value
func (o *NokiaProxyServerExternalServerResponse) GetConnectTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetConnectTimeoutOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ConnectTimeout, true
}

// SetConnectTimeout sets field value
func (o *NokiaProxyServerExternalServerResponse) SetConnectTimeout(v string) {
	o.ConnectTimeout = v
}

// GetMaxResponseSize returns the MaxResponseSize field value
func (o *NokiaProxyServerExternalServerResponse) GetMaxResponseSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxResponseSize
}

// GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field value
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetMaxResponseSizeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MaxResponseSize, true
}

// SetMaxResponseSize sets field value
func (o *NokiaProxyServerExternalServerResponse) SetMaxResponseSize(v string) {
	o.MaxResponseSize = v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetKeyManagerProvider() string {
	if o == nil || isNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || isNil(o.KeyManagerProvider) {
    return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasKeyManagerProvider() bool {
	if o != nil && !isNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *NokiaProxyServerExternalServerResponse) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetTrustManagerProvider() string {
	if o == nil || isNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || isNil(o.TrustManagerProvider) {
    return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasTrustManagerProvider() bool {
	if o != nil && !isNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *NokiaProxyServerExternalServerResponse) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetInitialConnections returns the InitialConnections field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetInitialConnections() int32 {
	if o == nil || isNil(o.InitialConnections) {
		var ret int32
		return ret
	}
	return *o.InitialConnections
}

// GetInitialConnectionsOk returns a tuple with the InitialConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetInitialConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.InitialConnections) {
    return nil, false
	}
	return o.InitialConnections, true
}

// HasInitialConnections returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasInitialConnections() bool {
	if o != nil && !isNil(o.InitialConnections) {
		return true
	}

	return false
}

// SetInitialConnections gets a reference to the given int32 and assigns it to the InitialConnections field.
func (o *NokiaProxyServerExternalServerResponse) SetInitialConnections(v int32) {
	o.InitialConnections = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetMaxConnections() int32 {
	if o == nil || isNil(o.MaxConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetMaxConnectionsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxConnections) {
    return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasMaxConnections() bool {
	if o != nil && !isNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int32 and assigns it to the MaxConnections field.
func (o *NokiaProxyServerExternalServerResponse) SetMaxConnections(v int32) {
	o.MaxConnections = &v
}

// GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp {
	if o == nil || isNil(o.DefunctConnectionResultCode) {
		var ret []EnumexternalServerDefunctConnectionResultCodeProp
		return ret
	}
	return o.DefunctConnectionResultCode
}

// GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetDefunctConnectionResultCodeOk() ([]EnumexternalServerDefunctConnectionResultCodeProp, bool) {
	if o == nil || isNil(o.DefunctConnectionResultCode) {
    return nil, false
	}
	return o.DefunctConnectionResultCode, true
}

// HasDefunctConnectionResultCode returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasDefunctConnectionResultCode() bool {
	if o != nil && !isNil(o.DefunctConnectionResultCode) {
		return true
	}

	return false
}

// SetDefunctConnectionResultCode gets a reference to the given []EnumexternalServerDefunctConnectionResultCodeProp and assigns it to the DefunctConnectionResultCode field.
func (o *NokiaProxyServerExternalServerResponse) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp) {
	o.DefunctConnectionResultCode = v
}

// GetAbandonOnTimeout returns the AbandonOnTimeout field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetAbandonOnTimeout() bool {
	if o == nil || isNil(o.AbandonOnTimeout) {
		var ret bool
		return ret
	}
	return *o.AbandonOnTimeout
}

// GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetAbandonOnTimeoutOk() (*bool, bool) {
	if o == nil || isNil(o.AbandonOnTimeout) {
    return nil, false
	}
	return o.AbandonOnTimeout, true
}

// HasAbandonOnTimeout returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasAbandonOnTimeout() bool {
	if o != nil && !isNil(o.AbandonOnTimeout) {
		return true
	}

	return false
}

// SetAbandonOnTimeout gets a reference to the given bool and assigns it to the AbandonOnTimeout field.
func (o *NokiaProxyServerExternalServerResponse) SetAbandonOnTimeout(v bool) {
	o.AbandonOnTimeout = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NokiaProxyServerExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NokiaProxyServerExternalServerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaProxyServerExternalServerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NokiaProxyServerExternalServerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *NokiaProxyServerExternalServerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o NokiaProxyServerExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["verifyCredentialsMethod"] = o.VerifyCredentialsMethod
	}
	if !isNil(o.UseAdministrativeOperationControl) {
		toSerialize["useAdministrativeOperationControl"] = o.UseAdministrativeOperationControl
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
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableNokiaProxyServerExternalServerResponse struct {
	value *NokiaProxyServerExternalServerResponse
	isSet bool
}

func (v NullableNokiaProxyServerExternalServerResponse) Get() *NokiaProxyServerExternalServerResponse {
	return v.value
}

func (v *NullableNokiaProxyServerExternalServerResponse) Set(val *NokiaProxyServerExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNokiaProxyServerExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNokiaProxyServerExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNokiaProxyServerExternalServerResponse(val *NokiaProxyServerExternalServerResponse) *NullableNokiaProxyServerExternalServerResponse {
	return &NullableNokiaProxyServerExternalServerResponse{value: val, isSet: true}
}

func (v NullableNokiaProxyServerExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNokiaProxyServerExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


