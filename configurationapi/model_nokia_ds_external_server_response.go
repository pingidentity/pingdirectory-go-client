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

// checks if the NokiaDsExternalServerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NokiaDsExternalServerResponse{}

// NokiaDsExternalServerResponse struct for NokiaDsExternalServerResponse
type NokiaDsExternalServerResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	// Name of the External Server
	Id                      string                                        `json:"id"`
	Schemas                 []EnumnokiaDsExternalServerSchemaUrn          `json:"schemas"`
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
	PassphraseProvider   *string                                    `json:"passphraseProvider,omitempty"`
	ConnectionSecurity   EnumexternalServerConnectionSecurityProp   `json:"connectionSecurity"`
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
	MaxConnections              *int32                                              `json:"maxConnections,omitempty"`
	DefunctConnectionResultCode []EnumexternalServerDefunctConnectionResultCodeProp `json:"defunctConnectionResultCode,omitempty"`
	// Indicates whether to send an abandon request for an operation for which a response timeout is encountered. A request which has timed out on one server may be retried on another server regardless of whether an abandon request is sent, but if the initial attempt is not abandoned then a long-running operation may unnecessarily continue to consume processing resources on the initial server.
	AbandonOnTimeout *bool `json:"abandonOnTimeout,omitempty"`
	// A description for this External Server
	Description *string `json:"description,omitempty"`
}

// NewNokiaDsExternalServerResponse instantiates a new NokiaDsExternalServerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNokiaDsExternalServerResponse(id string, schemas []EnumnokiaDsExternalServerSchemaUrn, verifyCredentialsMethod EnumexternalServerVerifyCredentialsMethodProp, serverHostName string, serverPort int32, connectionSecurity EnumexternalServerConnectionSecurityProp, authenticationMethod EnumexternalServerAuthenticationMethodProp, maxConnectionAge string, connectTimeout string, maxResponseSize string) *NokiaDsExternalServerResponse {
	this := NokiaDsExternalServerResponse{}
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

// NewNokiaDsExternalServerResponseWithDefaults instantiates a new NokiaDsExternalServerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNokiaDsExternalServerResponseWithDefaults() *NokiaDsExternalServerResponse {
	this := NokiaDsExternalServerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *NokiaDsExternalServerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *NokiaDsExternalServerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetId returns the Id field value
func (o *NokiaDsExternalServerResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NokiaDsExternalServerResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *NokiaDsExternalServerResponse) GetSchemas() []EnumnokiaDsExternalServerSchemaUrn {
	if o == nil {
		var ret []EnumnokiaDsExternalServerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetSchemasOk() ([]EnumnokiaDsExternalServerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *NokiaDsExternalServerResponse) SetSchemas(v []EnumnokiaDsExternalServerSchemaUrn) {
	o.Schemas = v
}

// GetVerifyCredentialsMethod returns the VerifyCredentialsMethod field value
func (o *NokiaDsExternalServerResponse) GetVerifyCredentialsMethod() EnumexternalServerVerifyCredentialsMethodProp {
	if o == nil {
		var ret EnumexternalServerVerifyCredentialsMethodProp
		return ret
	}

	return o.VerifyCredentialsMethod
}

// GetVerifyCredentialsMethodOk returns a tuple with the VerifyCredentialsMethod field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetVerifyCredentialsMethodOk() (*EnumexternalServerVerifyCredentialsMethodProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyCredentialsMethod, true
}

// SetVerifyCredentialsMethod sets field value
func (o *NokiaDsExternalServerResponse) SetVerifyCredentialsMethod(v EnumexternalServerVerifyCredentialsMethodProp) {
	o.VerifyCredentialsMethod = v
}

// GetUseAdministrativeOperationControl returns the UseAdministrativeOperationControl field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetUseAdministrativeOperationControl() bool {
	if o == nil || IsNil(o.UseAdministrativeOperationControl) {
		var ret bool
		return ret
	}
	return *o.UseAdministrativeOperationControl
}

// GetUseAdministrativeOperationControlOk returns a tuple with the UseAdministrativeOperationControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetUseAdministrativeOperationControlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseAdministrativeOperationControl) {
		return nil, false
	}
	return o.UseAdministrativeOperationControl, true
}

// HasUseAdministrativeOperationControl returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasUseAdministrativeOperationControl() bool {
	if o != nil && !IsNil(o.UseAdministrativeOperationControl) {
		return true
	}

	return false
}

// SetUseAdministrativeOperationControl gets a reference to the given bool and assigns it to the UseAdministrativeOperationControl field.
func (o *NokiaDsExternalServerResponse) SetUseAdministrativeOperationControl(v bool) {
	o.UseAdministrativeOperationControl = &v
}

// GetServerHostName returns the ServerHostName field value
func (o *NokiaDsExternalServerResponse) GetServerHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerHostName
}

// GetServerHostNameOk returns a tuple with the ServerHostName field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetServerHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerHostName, true
}

// SetServerHostName sets field value
func (o *NokiaDsExternalServerResponse) SetServerHostName(v string) {
	o.ServerHostName = v
}

// GetServerPort returns the ServerPort field value
func (o *NokiaDsExternalServerResponse) GetServerPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ServerPort
}

// GetServerPortOk returns a tuple with the ServerPort field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetServerPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerPort, true
}

// SetServerPort sets field value
func (o *NokiaDsExternalServerResponse) SetServerPort(v int32) {
	o.ServerPort = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *NokiaDsExternalServerResponse) SetLocation(v string) {
	o.Location = &v
}

// GetBindDN returns the BindDN field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetBindDN() string {
	if o == nil || IsNil(o.BindDN) {
		var ret string
		return ret
	}
	return *o.BindDN
}

// GetBindDNOk returns a tuple with the BindDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetBindDNOk() (*string, bool) {
	if o == nil || IsNil(o.BindDN) {
		return nil, false
	}
	return o.BindDN, true
}

// HasBindDN returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasBindDN() bool {
	if o != nil && !IsNil(o.BindDN) {
		return true
	}

	return false
}

// SetBindDN gets a reference to the given string and assigns it to the BindDN field.
func (o *NokiaDsExternalServerResponse) SetBindDN(v string) {
	o.BindDN = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *NokiaDsExternalServerResponse) SetPassword(v string) {
	o.Password = &v
}

// GetPassphraseProvider returns the PassphraseProvider field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetPassphraseProvider() string {
	if o == nil || IsNil(o.PassphraseProvider) {
		var ret string
		return ret
	}
	return *o.PassphraseProvider
}

// GetPassphraseProviderOk returns a tuple with the PassphraseProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetPassphraseProviderOk() (*string, bool) {
	if o == nil || IsNil(o.PassphraseProvider) {
		return nil, false
	}
	return o.PassphraseProvider, true
}

// HasPassphraseProvider returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasPassphraseProvider() bool {
	if o != nil && !IsNil(o.PassphraseProvider) {
		return true
	}

	return false
}

// SetPassphraseProvider gets a reference to the given string and assigns it to the PassphraseProvider field.
func (o *NokiaDsExternalServerResponse) SetPassphraseProvider(v string) {
	o.PassphraseProvider = &v
}

// GetConnectionSecurity returns the ConnectionSecurity field value
func (o *NokiaDsExternalServerResponse) GetConnectionSecurity() EnumexternalServerConnectionSecurityProp {
	if o == nil {
		var ret EnumexternalServerConnectionSecurityProp
		return ret
	}

	return o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetConnectionSecurityOk() (*EnumexternalServerConnectionSecurityProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionSecurity, true
}

// SetConnectionSecurity sets field value
func (o *NokiaDsExternalServerResponse) SetConnectionSecurity(v EnumexternalServerConnectionSecurityProp) {
	o.ConnectionSecurity = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *NokiaDsExternalServerResponse) GetAuthenticationMethod() EnumexternalServerAuthenticationMethodProp {
	if o == nil {
		var ret EnumexternalServerAuthenticationMethodProp
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetAuthenticationMethodOk() (*EnumexternalServerAuthenticationMethodProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *NokiaDsExternalServerResponse) SetAuthenticationMethod(v EnumexternalServerAuthenticationMethodProp) {
	o.AuthenticationMethod = v
}

// GetHealthCheckConnectTimeout returns the HealthCheckConnectTimeout field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetHealthCheckConnectTimeout() string {
	if o == nil || IsNil(o.HealthCheckConnectTimeout) {
		var ret string
		return ret
	}
	return *o.HealthCheckConnectTimeout
}

// GetHealthCheckConnectTimeoutOk returns a tuple with the HealthCheckConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetHealthCheckConnectTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.HealthCheckConnectTimeout) {
		return nil, false
	}
	return o.HealthCheckConnectTimeout, true
}

// HasHealthCheckConnectTimeout returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasHealthCheckConnectTimeout() bool {
	if o != nil && !IsNil(o.HealthCheckConnectTimeout) {
		return true
	}

	return false
}

// SetHealthCheckConnectTimeout gets a reference to the given string and assigns it to the HealthCheckConnectTimeout field.
func (o *NokiaDsExternalServerResponse) SetHealthCheckConnectTimeout(v string) {
	o.HealthCheckConnectTimeout = &v
}

// GetMaxConnectionAge returns the MaxConnectionAge field value
func (o *NokiaDsExternalServerResponse) GetMaxConnectionAge() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxConnectionAge
}

// GetMaxConnectionAgeOk returns a tuple with the MaxConnectionAge field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetMaxConnectionAgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxConnectionAge, true
}

// SetMaxConnectionAge sets field value
func (o *NokiaDsExternalServerResponse) SetMaxConnectionAge(v string) {
	o.MaxConnectionAge = v
}

// GetMinExpiredConnectionDisconnectInterval returns the MinExpiredConnectionDisconnectInterval field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetMinExpiredConnectionDisconnectInterval() string {
	if o == nil || IsNil(o.MinExpiredConnectionDisconnectInterval) {
		var ret string
		return ret
	}
	return *o.MinExpiredConnectionDisconnectInterval
}

// GetMinExpiredConnectionDisconnectIntervalOk returns a tuple with the MinExpiredConnectionDisconnectInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetMinExpiredConnectionDisconnectIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.MinExpiredConnectionDisconnectInterval) {
		return nil, false
	}
	return o.MinExpiredConnectionDisconnectInterval, true
}

// HasMinExpiredConnectionDisconnectInterval returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasMinExpiredConnectionDisconnectInterval() bool {
	if o != nil && !IsNil(o.MinExpiredConnectionDisconnectInterval) {
		return true
	}

	return false
}

// SetMinExpiredConnectionDisconnectInterval gets a reference to the given string and assigns it to the MinExpiredConnectionDisconnectInterval field.
func (o *NokiaDsExternalServerResponse) SetMinExpiredConnectionDisconnectInterval(v string) {
	o.MinExpiredConnectionDisconnectInterval = &v
}

// GetConnectTimeout returns the ConnectTimeout field value
func (o *NokiaDsExternalServerResponse) GetConnectTimeout() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectTimeout
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetConnectTimeoutOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectTimeout, true
}

// SetConnectTimeout sets field value
func (o *NokiaDsExternalServerResponse) SetConnectTimeout(v string) {
	o.ConnectTimeout = v
}

// GetMaxResponseSize returns the MaxResponseSize field value
func (o *NokiaDsExternalServerResponse) GetMaxResponseSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxResponseSize
}

// GetMaxResponseSizeOk returns a tuple with the MaxResponseSize field value
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetMaxResponseSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxResponseSize, true
}

// SetMaxResponseSize sets field value
func (o *NokiaDsExternalServerResponse) SetMaxResponseSize(v string) {
	o.MaxResponseSize = v
}

// GetKeyManagerProvider returns the KeyManagerProvider field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetKeyManagerProvider() string {
	if o == nil || IsNil(o.KeyManagerProvider) {
		var ret string
		return ret
	}
	return *o.KeyManagerProvider
}

// GetKeyManagerProviderOk returns a tuple with the KeyManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetKeyManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.KeyManagerProvider) {
		return nil, false
	}
	return o.KeyManagerProvider, true
}

// HasKeyManagerProvider returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasKeyManagerProvider() bool {
	if o != nil && !IsNil(o.KeyManagerProvider) {
		return true
	}

	return false
}

// SetKeyManagerProvider gets a reference to the given string and assigns it to the KeyManagerProvider field.
func (o *NokiaDsExternalServerResponse) SetKeyManagerProvider(v string) {
	o.KeyManagerProvider = &v
}

// GetTrustManagerProvider returns the TrustManagerProvider field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetTrustManagerProvider() string {
	if o == nil || IsNil(o.TrustManagerProvider) {
		var ret string
		return ret
	}
	return *o.TrustManagerProvider
}

// GetTrustManagerProviderOk returns a tuple with the TrustManagerProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetTrustManagerProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TrustManagerProvider) {
		return nil, false
	}
	return o.TrustManagerProvider, true
}

// HasTrustManagerProvider returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasTrustManagerProvider() bool {
	if o != nil && !IsNil(o.TrustManagerProvider) {
		return true
	}

	return false
}

// SetTrustManagerProvider gets a reference to the given string and assigns it to the TrustManagerProvider field.
func (o *NokiaDsExternalServerResponse) SetTrustManagerProvider(v string) {
	o.TrustManagerProvider = &v
}

// GetInitialConnections returns the InitialConnections field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetInitialConnections() int32 {
	if o == nil || IsNil(o.InitialConnections) {
		var ret int32
		return ret
	}
	return *o.InitialConnections
}

// GetInitialConnectionsOk returns a tuple with the InitialConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetInitialConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.InitialConnections) {
		return nil, false
	}
	return o.InitialConnections, true
}

// HasInitialConnections returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasInitialConnections() bool {
	if o != nil && !IsNil(o.InitialConnections) {
		return true
	}

	return false
}

// SetInitialConnections gets a reference to the given int32 and assigns it to the InitialConnections field.
func (o *NokiaDsExternalServerResponse) SetInitialConnections(v int32) {
	o.InitialConnections = &v
}

// GetMaxConnections returns the MaxConnections field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetMaxConnections() int32 {
	if o == nil || IsNil(o.MaxConnections) {
		var ret int32
		return ret
	}
	return *o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetMaxConnectionsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxConnections) {
		return nil, false
	}
	return o.MaxConnections, true
}

// HasMaxConnections returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasMaxConnections() bool {
	if o != nil && !IsNil(o.MaxConnections) {
		return true
	}

	return false
}

// SetMaxConnections gets a reference to the given int32 and assigns it to the MaxConnections field.
func (o *NokiaDsExternalServerResponse) SetMaxConnections(v int32) {
	o.MaxConnections = &v
}

// GetDefunctConnectionResultCode returns the DefunctConnectionResultCode field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetDefunctConnectionResultCode() []EnumexternalServerDefunctConnectionResultCodeProp {
	if o == nil || IsNil(o.DefunctConnectionResultCode) {
		var ret []EnumexternalServerDefunctConnectionResultCodeProp
		return ret
	}
	return o.DefunctConnectionResultCode
}

// GetDefunctConnectionResultCodeOk returns a tuple with the DefunctConnectionResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetDefunctConnectionResultCodeOk() ([]EnumexternalServerDefunctConnectionResultCodeProp, bool) {
	if o == nil || IsNil(o.DefunctConnectionResultCode) {
		return nil, false
	}
	return o.DefunctConnectionResultCode, true
}

// HasDefunctConnectionResultCode returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasDefunctConnectionResultCode() bool {
	if o != nil && !IsNil(o.DefunctConnectionResultCode) {
		return true
	}

	return false
}

// SetDefunctConnectionResultCode gets a reference to the given []EnumexternalServerDefunctConnectionResultCodeProp and assigns it to the DefunctConnectionResultCode field.
func (o *NokiaDsExternalServerResponse) SetDefunctConnectionResultCode(v []EnumexternalServerDefunctConnectionResultCodeProp) {
	o.DefunctConnectionResultCode = v
}

// GetAbandonOnTimeout returns the AbandonOnTimeout field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetAbandonOnTimeout() bool {
	if o == nil || IsNil(o.AbandonOnTimeout) {
		var ret bool
		return ret
	}
	return *o.AbandonOnTimeout
}

// GetAbandonOnTimeoutOk returns a tuple with the AbandonOnTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetAbandonOnTimeoutOk() (*bool, bool) {
	if o == nil || IsNil(o.AbandonOnTimeout) {
		return nil, false
	}
	return o.AbandonOnTimeout, true
}

// HasAbandonOnTimeout returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasAbandonOnTimeout() bool {
	if o != nil && !IsNil(o.AbandonOnTimeout) {
		return true
	}

	return false
}

// SetAbandonOnTimeout gets a reference to the given bool and assigns it to the AbandonOnTimeout field.
func (o *NokiaDsExternalServerResponse) SetAbandonOnTimeout(v bool) {
	o.AbandonOnTimeout = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NokiaDsExternalServerResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NokiaDsExternalServerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NokiaDsExternalServerResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NokiaDsExternalServerResponse) SetDescription(v string) {
	o.Description = &v
}

func (o NokiaDsExternalServerResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NokiaDsExternalServerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["verifyCredentialsMethod"] = o.VerifyCredentialsMethod
	if !IsNil(o.UseAdministrativeOperationControl) {
		toSerialize["useAdministrativeOperationControl"] = o.UseAdministrativeOperationControl
	}
	toSerialize["serverHostName"] = o.ServerHostName
	toSerialize["serverPort"] = o.ServerPort
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
	toSerialize["connectionSecurity"] = o.ConnectionSecurity
	toSerialize["authenticationMethod"] = o.AuthenticationMethod
	if !IsNil(o.HealthCheckConnectTimeout) {
		toSerialize["healthCheckConnectTimeout"] = o.HealthCheckConnectTimeout
	}
	toSerialize["maxConnectionAge"] = o.MaxConnectionAge
	if !IsNil(o.MinExpiredConnectionDisconnectInterval) {
		toSerialize["minExpiredConnectionDisconnectInterval"] = o.MinExpiredConnectionDisconnectInterval
	}
	toSerialize["connectTimeout"] = o.ConnectTimeout
	toSerialize["maxResponseSize"] = o.MaxResponseSize
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

type NullableNokiaDsExternalServerResponse struct {
	value *NokiaDsExternalServerResponse
	isSet bool
}

func (v NullableNokiaDsExternalServerResponse) Get() *NokiaDsExternalServerResponse {
	return v.value
}

func (v *NullableNokiaDsExternalServerResponse) Set(val *NokiaDsExternalServerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNokiaDsExternalServerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNokiaDsExternalServerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNokiaDsExternalServerResponse(val *NokiaDsExternalServerResponse) *NullableNokiaDsExternalServerResponse {
	return &NullableNokiaDsExternalServerResponse{value: val, isSet: true}
}

func (v NullableNokiaDsExternalServerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNokiaDsExternalServerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
