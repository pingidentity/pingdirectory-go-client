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

// AuthorizeServerInstanceResponse struct for AuthorizeServerInstanceResponse
type AuthorizeServerInstanceResponse struct {
	Schemas []EnumauthorizeServerInstanceSchemaUrn `json:"schemas"`
	ServerInstanceType *EnumserverInstanceServerInstanceTypeProp `json:"serverInstanceType,omitempty"`
	// The name of this Server Instance. The instance name needs to be unique if this server will be part of a topology of servers that are connected to each other. Once set, it may not be changed.
	ServerInstanceName string `json:"serverInstanceName"`
	// The name of the cluster to which this Server Instance belongs. Server instances within the same cluster will share the same cluster-wide configuration.
	ClusterName string `json:"clusterName"`
	// Specifies the location for the Server Instance.
	ServerInstanceLocation *string `json:"serverInstanceLocation,omitempty"`
	// The name of the host where this Server Instance is installed.
	Hostname *string `json:"hostname,omitempty"`
	// The file system path where this Server Instance is installed.
	ServerRoot *string `json:"serverRoot,omitempty"`
	// The version of the server.
	ServerVersion string `json:"serverVersion"`
	// The public component of the certificate used by this instance to protect inter-server communication and to perform server-specific encryption. This will generally be managed by the server and should only be altered by administrators under explicit direction from Ping Identity support personnel.
	InterServerCertificate *string `json:"interServerCertificate,omitempty"`
	// The TCP port on which this server is listening for LDAP connections.
	LdapPort *int32 `json:"ldapPort,omitempty"`
	// The TCP port on which this server is listening for LDAP secure connections.
	LdapsPort *int32 `json:"ldapsPort,omitempty"`
	// The TCP port on which this server is listening for HTTP connections.
	HttpPort *int32 `json:"httpPort,omitempty"`
	// The TCP port on which this server is listening for HTTPS connections.
	HttpsPort *int32 `json:"httpsPort,omitempty"`
	// The replication TCP port.
	ReplicationPort *int32 `json:"replicationPort,omitempty"`
	// Specifies a unique identifier for the replication server on this server instance.
	ReplicationServerID *int32 `json:"replicationServerID,omitempty"`
	ReplicationDomainServerID []int32 `json:"replicationDomainServerID,omitempty"`
	// The TCP port on which this server is listening for JMX connections.
	JmxPort *int32 `json:"jmxPort,omitempty"`
	// The TCP port on which this server is listening for JMX secure connections.
	JmxsPort *int32 `json:"jmxsPort,omitempty"`
	PreferredSecurity *EnumserverInstancePreferredSecurityProp `json:"preferredSecurity,omitempty"`
	// Indicates whether StartTLS is enabled on this server.
	StartTLSEnabled *bool `json:"startTLSEnabled,omitempty"`
	BaseDN []string `json:"baseDN,omitempty"`
	MemberOfServerGroup []string `json:"memberOfServerGroup,omitempty"`
}

// NewAuthorizeServerInstanceResponse instantiates a new AuthorizeServerInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeServerInstanceResponse(schemas []EnumauthorizeServerInstanceSchemaUrn, serverInstanceName string, clusterName string, serverVersion string) *AuthorizeServerInstanceResponse {
	this := AuthorizeServerInstanceResponse{}
	this.Schemas = schemas
	this.ServerInstanceName = serverInstanceName
	this.ClusterName = clusterName
	this.ServerVersion = serverVersion
	return &this
}

// NewAuthorizeServerInstanceResponseWithDefaults instantiates a new AuthorizeServerInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeServerInstanceResponseWithDefaults() *AuthorizeServerInstanceResponse {
	this := AuthorizeServerInstanceResponse{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AuthorizeServerInstanceResponse) GetSchemas() []EnumauthorizeServerInstanceSchemaUrn {
	if o == nil {
		var ret []EnumauthorizeServerInstanceSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetSchemasOk() ([]EnumauthorizeServerInstanceSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AuthorizeServerInstanceResponse) SetSchemas(v []EnumauthorizeServerInstanceSchemaUrn) {
	o.Schemas = v
}

// GetServerInstanceType returns the ServerInstanceType field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetServerInstanceType() EnumserverInstanceServerInstanceTypeProp {
	if o == nil || isNil(o.ServerInstanceType) {
		var ret EnumserverInstanceServerInstanceTypeProp
		return ret
	}
	return *o.ServerInstanceType
}

// GetServerInstanceTypeOk returns a tuple with the ServerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetServerInstanceTypeOk() (*EnumserverInstanceServerInstanceTypeProp, bool) {
	if o == nil || isNil(o.ServerInstanceType) {
    return nil, false
	}
	return o.ServerInstanceType, true
}

// HasServerInstanceType returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasServerInstanceType() bool {
	if o != nil && !isNil(o.ServerInstanceType) {
		return true
	}

	return false
}

// SetServerInstanceType gets a reference to the given EnumserverInstanceServerInstanceTypeProp and assigns it to the ServerInstanceType field.
func (o *AuthorizeServerInstanceResponse) SetServerInstanceType(v EnumserverInstanceServerInstanceTypeProp) {
	o.ServerInstanceType = &v
}

// GetServerInstanceName returns the ServerInstanceName field value
func (o *AuthorizeServerInstanceResponse) GetServerInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerInstanceName
}

// GetServerInstanceNameOk returns a tuple with the ServerInstanceName field value
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetServerInstanceNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServerInstanceName, true
}

// SetServerInstanceName sets field value
func (o *AuthorizeServerInstanceResponse) SetServerInstanceName(v string) {
	o.ServerInstanceName = v
}

// GetClusterName returns the ClusterName field value
func (o *AuthorizeServerInstanceResponse) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetClusterNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *AuthorizeServerInstanceResponse) SetClusterName(v string) {
	o.ClusterName = v
}

// GetServerInstanceLocation returns the ServerInstanceLocation field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetServerInstanceLocation() string {
	if o == nil || isNil(o.ServerInstanceLocation) {
		var ret string
		return ret
	}
	return *o.ServerInstanceLocation
}

// GetServerInstanceLocationOk returns a tuple with the ServerInstanceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetServerInstanceLocationOk() (*string, bool) {
	if o == nil || isNil(o.ServerInstanceLocation) {
    return nil, false
	}
	return o.ServerInstanceLocation, true
}

// HasServerInstanceLocation returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasServerInstanceLocation() bool {
	if o != nil && !isNil(o.ServerInstanceLocation) {
		return true
	}

	return false
}

// SetServerInstanceLocation gets a reference to the given string and assigns it to the ServerInstanceLocation field.
func (o *AuthorizeServerInstanceResponse) SetServerInstanceLocation(v string) {
	o.ServerInstanceLocation = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetHostname() string {
	if o == nil || isNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetHostnameOk() (*string, bool) {
	if o == nil || isNil(o.Hostname) {
    return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasHostname() bool {
	if o != nil && !isNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *AuthorizeServerInstanceResponse) SetHostname(v string) {
	o.Hostname = &v
}

// GetServerRoot returns the ServerRoot field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetServerRoot() string {
	if o == nil || isNil(o.ServerRoot) {
		var ret string
		return ret
	}
	return *o.ServerRoot
}

// GetServerRootOk returns a tuple with the ServerRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetServerRootOk() (*string, bool) {
	if o == nil || isNil(o.ServerRoot) {
    return nil, false
	}
	return o.ServerRoot, true
}

// HasServerRoot returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasServerRoot() bool {
	if o != nil && !isNil(o.ServerRoot) {
		return true
	}

	return false
}

// SetServerRoot gets a reference to the given string and assigns it to the ServerRoot field.
func (o *AuthorizeServerInstanceResponse) SetServerRoot(v string) {
	o.ServerRoot = &v
}

// GetServerVersion returns the ServerVersion field value
func (o *AuthorizeServerInstanceResponse) GetServerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetServerVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServerVersion, true
}

// SetServerVersion sets field value
func (o *AuthorizeServerInstanceResponse) SetServerVersion(v string) {
	o.ServerVersion = v
}

// GetInterServerCertificate returns the InterServerCertificate field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetInterServerCertificate() string {
	if o == nil || isNil(o.InterServerCertificate) {
		var ret string
		return ret
	}
	return *o.InterServerCertificate
}

// GetInterServerCertificateOk returns a tuple with the InterServerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetInterServerCertificateOk() (*string, bool) {
	if o == nil || isNil(o.InterServerCertificate) {
    return nil, false
	}
	return o.InterServerCertificate, true
}

// HasInterServerCertificate returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasInterServerCertificate() bool {
	if o != nil && !isNil(o.InterServerCertificate) {
		return true
	}

	return false
}

// SetInterServerCertificate gets a reference to the given string and assigns it to the InterServerCertificate field.
func (o *AuthorizeServerInstanceResponse) SetInterServerCertificate(v string) {
	o.InterServerCertificate = &v
}

// GetLdapPort returns the LdapPort field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetLdapPort() int32 {
	if o == nil || isNil(o.LdapPort) {
		var ret int32
		return ret
	}
	return *o.LdapPort
}

// GetLdapPortOk returns a tuple with the LdapPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetLdapPortOk() (*int32, bool) {
	if o == nil || isNil(o.LdapPort) {
    return nil, false
	}
	return o.LdapPort, true
}

// HasLdapPort returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasLdapPort() bool {
	if o != nil && !isNil(o.LdapPort) {
		return true
	}

	return false
}

// SetLdapPort gets a reference to the given int32 and assigns it to the LdapPort field.
func (o *AuthorizeServerInstanceResponse) SetLdapPort(v int32) {
	o.LdapPort = &v
}

// GetLdapsPort returns the LdapsPort field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetLdapsPort() int32 {
	if o == nil || isNil(o.LdapsPort) {
		var ret int32
		return ret
	}
	return *o.LdapsPort
}

// GetLdapsPortOk returns a tuple with the LdapsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetLdapsPortOk() (*int32, bool) {
	if o == nil || isNil(o.LdapsPort) {
    return nil, false
	}
	return o.LdapsPort, true
}

// HasLdapsPort returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasLdapsPort() bool {
	if o != nil && !isNil(o.LdapsPort) {
		return true
	}

	return false
}

// SetLdapsPort gets a reference to the given int32 and assigns it to the LdapsPort field.
func (o *AuthorizeServerInstanceResponse) SetLdapsPort(v int32) {
	o.LdapsPort = &v
}

// GetHttpPort returns the HttpPort field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetHttpPort() int32 {
	if o == nil || isNil(o.HttpPort) {
		var ret int32
		return ret
	}
	return *o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetHttpPortOk() (*int32, bool) {
	if o == nil || isNil(o.HttpPort) {
    return nil, false
	}
	return o.HttpPort, true
}

// HasHttpPort returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasHttpPort() bool {
	if o != nil && !isNil(o.HttpPort) {
		return true
	}

	return false
}

// SetHttpPort gets a reference to the given int32 and assigns it to the HttpPort field.
func (o *AuthorizeServerInstanceResponse) SetHttpPort(v int32) {
	o.HttpPort = &v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetHttpsPort() int32 {
	if o == nil || isNil(o.HttpsPort) {
		var ret int32
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetHttpsPortOk() (*int32, bool) {
	if o == nil || isNil(o.HttpsPort) {
    return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasHttpsPort() bool {
	if o != nil && !isNil(o.HttpsPort) {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int32 and assigns it to the HttpsPort field.
func (o *AuthorizeServerInstanceResponse) SetHttpsPort(v int32) {
	o.HttpsPort = &v
}

// GetReplicationPort returns the ReplicationPort field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetReplicationPort() int32 {
	if o == nil || isNil(o.ReplicationPort) {
		var ret int32
		return ret
	}
	return *o.ReplicationPort
}

// GetReplicationPortOk returns a tuple with the ReplicationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetReplicationPortOk() (*int32, bool) {
	if o == nil || isNil(o.ReplicationPort) {
    return nil, false
	}
	return o.ReplicationPort, true
}

// HasReplicationPort returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasReplicationPort() bool {
	if o != nil && !isNil(o.ReplicationPort) {
		return true
	}

	return false
}

// SetReplicationPort gets a reference to the given int32 and assigns it to the ReplicationPort field.
func (o *AuthorizeServerInstanceResponse) SetReplicationPort(v int32) {
	o.ReplicationPort = &v
}

// GetReplicationServerID returns the ReplicationServerID field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetReplicationServerID() int32 {
	if o == nil || isNil(o.ReplicationServerID) {
		var ret int32
		return ret
	}
	return *o.ReplicationServerID
}

// GetReplicationServerIDOk returns a tuple with the ReplicationServerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetReplicationServerIDOk() (*int32, bool) {
	if o == nil || isNil(o.ReplicationServerID) {
    return nil, false
	}
	return o.ReplicationServerID, true
}

// HasReplicationServerID returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasReplicationServerID() bool {
	if o != nil && !isNil(o.ReplicationServerID) {
		return true
	}

	return false
}

// SetReplicationServerID gets a reference to the given int32 and assigns it to the ReplicationServerID field.
func (o *AuthorizeServerInstanceResponse) SetReplicationServerID(v int32) {
	o.ReplicationServerID = &v
}

// GetReplicationDomainServerID returns the ReplicationDomainServerID field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetReplicationDomainServerID() []int32 {
	if o == nil || isNil(o.ReplicationDomainServerID) {
		var ret []int32
		return ret
	}
	return o.ReplicationDomainServerID
}

// GetReplicationDomainServerIDOk returns a tuple with the ReplicationDomainServerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetReplicationDomainServerIDOk() ([]int32, bool) {
	if o == nil || isNil(o.ReplicationDomainServerID) {
    return nil, false
	}
	return o.ReplicationDomainServerID, true
}

// HasReplicationDomainServerID returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasReplicationDomainServerID() bool {
	if o != nil && !isNil(o.ReplicationDomainServerID) {
		return true
	}

	return false
}

// SetReplicationDomainServerID gets a reference to the given []int32 and assigns it to the ReplicationDomainServerID field.
func (o *AuthorizeServerInstanceResponse) SetReplicationDomainServerID(v []int32) {
	o.ReplicationDomainServerID = v
}

// GetJmxPort returns the JmxPort field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetJmxPort() int32 {
	if o == nil || isNil(o.JmxPort) {
		var ret int32
		return ret
	}
	return *o.JmxPort
}

// GetJmxPortOk returns a tuple with the JmxPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetJmxPortOk() (*int32, bool) {
	if o == nil || isNil(o.JmxPort) {
    return nil, false
	}
	return o.JmxPort, true
}

// HasJmxPort returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasJmxPort() bool {
	if o != nil && !isNil(o.JmxPort) {
		return true
	}

	return false
}

// SetJmxPort gets a reference to the given int32 and assigns it to the JmxPort field.
func (o *AuthorizeServerInstanceResponse) SetJmxPort(v int32) {
	o.JmxPort = &v
}

// GetJmxsPort returns the JmxsPort field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetJmxsPort() int32 {
	if o == nil || isNil(o.JmxsPort) {
		var ret int32
		return ret
	}
	return *o.JmxsPort
}

// GetJmxsPortOk returns a tuple with the JmxsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetJmxsPortOk() (*int32, bool) {
	if o == nil || isNil(o.JmxsPort) {
    return nil, false
	}
	return o.JmxsPort, true
}

// HasJmxsPort returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasJmxsPort() bool {
	if o != nil && !isNil(o.JmxsPort) {
		return true
	}

	return false
}

// SetJmxsPort gets a reference to the given int32 and assigns it to the JmxsPort field.
func (o *AuthorizeServerInstanceResponse) SetJmxsPort(v int32) {
	o.JmxsPort = &v
}

// GetPreferredSecurity returns the PreferredSecurity field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetPreferredSecurity() EnumserverInstancePreferredSecurityProp {
	if o == nil || isNil(o.PreferredSecurity) {
		var ret EnumserverInstancePreferredSecurityProp
		return ret
	}
	return *o.PreferredSecurity
}

// GetPreferredSecurityOk returns a tuple with the PreferredSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetPreferredSecurityOk() (*EnumserverInstancePreferredSecurityProp, bool) {
	if o == nil || isNil(o.PreferredSecurity) {
    return nil, false
	}
	return o.PreferredSecurity, true
}

// HasPreferredSecurity returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasPreferredSecurity() bool {
	if o != nil && !isNil(o.PreferredSecurity) {
		return true
	}

	return false
}

// SetPreferredSecurity gets a reference to the given EnumserverInstancePreferredSecurityProp and assigns it to the PreferredSecurity field.
func (o *AuthorizeServerInstanceResponse) SetPreferredSecurity(v EnumserverInstancePreferredSecurityProp) {
	o.PreferredSecurity = &v
}

// GetStartTLSEnabled returns the StartTLSEnabled field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetStartTLSEnabled() bool {
	if o == nil || isNil(o.StartTLSEnabled) {
		var ret bool
		return ret
	}
	return *o.StartTLSEnabled
}

// GetStartTLSEnabledOk returns a tuple with the StartTLSEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetStartTLSEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.StartTLSEnabled) {
    return nil, false
	}
	return o.StartTLSEnabled, true
}

// HasStartTLSEnabled returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasStartTLSEnabled() bool {
	if o != nil && !isNil(o.StartTLSEnabled) {
		return true
	}

	return false
}

// SetStartTLSEnabled gets a reference to the given bool and assigns it to the StartTLSEnabled field.
func (o *AuthorizeServerInstanceResponse) SetStartTLSEnabled(v bool) {
	o.StartTLSEnabled = &v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetBaseDN() []string {
	if o == nil || isNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil || isNil(o.BaseDN) {
    return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasBaseDN() bool {
	if o != nil && !isNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *AuthorizeServerInstanceResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetMemberOfServerGroup returns the MemberOfServerGroup field value if set, zero value otherwise.
func (o *AuthorizeServerInstanceResponse) GetMemberOfServerGroup() []string {
	if o == nil || isNil(o.MemberOfServerGroup) {
		var ret []string
		return ret
	}
	return o.MemberOfServerGroup
}

// GetMemberOfServerGroupOk returns a tuple with the MemberOfServerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeServerInstanceResponse) GetMemberOfServerGroupOk() ([]string, bool) {
	if o == nil || isNil(o.MemberOfServerGroup) {
    return nil, false
	}
	return o.MemberOfServerGroup, true
}

// HasMemberOfServerGroup returns a boolean if a field has been set.
func (o *AuthorizeServerInstanceResponse) HasMemberOfServerGroup() bool {
	if o != nil && !isNil(o.MemberOfServerGroup) {
		return true
	}

	return false
}

// SetMemberOfServerGroup gets a reference to the given []string and assigns it to the MemberOfServerGroup field.
func (o *AuthorizeServerInstanceResponse) SetMemberOfServerGroup(v []string) {
	o.MemberOfServerGroup = v
}

func (o AuthorizeServerInstanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.ServerInstanceType) {
		toSerialize["serverInstanceType"] = o.ServerInstanceType
	}
	if true {
		toSerialize["serverInstanceName"] = o.ServerInstanceName
	}
	if true {
		toSerialize["clusterName"] = o.ClusterName
	}
	if !isNil(o.ServerInstanceLocation) {
		toSerialize["serverInstanceLocation"] = o.ServerInstanceLocation
	}
	if !isNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !isNil(o.ServerRoot) {
		toSerialize["serverRoot"] = o.ServerRoot
	}
	if true {
		toSerialize["serverVersion"] = o.ServerVersion
	}
	if !isNil(o.InterServerCertificate) {
		toSerialize["interServerCertificate"] = o.InterServerCertificate
	}
	if !isNil(o.LdapPort) {
		toSerialize["ldapPort"] = o.LdapPort
	}
	if !isNil(o.LdapsPort) {
		toSerialize["ldapsPort"] = o.LdapsPort
	}
	if !isNil(o.HttpPort) {
		toSerialize["httpPort"] = o.HttpPort
	}
	if !isNil(o.HttpsPort) {
		toSerialize["httpsPort"] = o.HttpsPort
	}
	if !isNil(o.ReplicationPort) {
		toSerialize["replicationPort"] = o.ReplicationPort
	}
	if !isNil(o.ReplicationServerID) {
		toSerialize["replicationServerID"] = o.ReplicationServerID
	}
	if !isNil(o.ReplicationDomainServerID) {
		toSerialize["replicationDomainServerID"] = o.ReplicationDomainServerID
	}
	if !isNil(o.JmxPort) {
		toSerialize["jmxPort"] = o.JmxPort
	}
	if !isNil(o.JmxsPort) {
		toSerialize["jmxsPort"] = o.JmxsPort
	}
	if !isNil(o.PreferredSecurity) {
		toSerialize["preferredSecurity"] = o.PreferredSecurity
	}
	if !isNil(o.StartTLSEnabled) {
		toSerialize["startTLSEnabled"] = o.StartTLSEnabled
	}
	if !isNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !isNil(o.MemberOfServerGroup) {
		toSerialize["memberOfServerGroup"] = o.MemberOfServerGroup
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizeServerInstanceResponse struct {
	value *AuthorizeServerInstanceResponse
	isSet bool
}

func (v NullableAuthorizeServerInstanceResponse) Get() *AuthorizeServerInstanceResponse {
	return v.value
}

func (v *NullableAuthorizeServerInstanceResponse) Set(val *AuthorizeServerInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeServerInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeServerInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeServerInstanceResponse(val *AuthorizeServerInstanceResponse) *NullableAuthorizeServerInstanceResponse {
	return &NullableAuthorizeServerInstanceResponse{value: val, isSet: true}
}

func (v NullableAuthorizeServerInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeServerInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


