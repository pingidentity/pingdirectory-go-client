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

// checks if the SyncServerInstanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncServerInstanceResponse{}

// SyncServerInstanceResponse struct for SyncServerInstanceResponse
type SyncServerInstanceResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumsyncServerInstanceSchemaUrn                  `json:"schemas"`
	// Name of the Server Instance
	Id                 string                                    `json:"id"`
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
	LdapPort *int64 `json:"ldapPort,omitempty"`
	// The TCP port on which this server is listening for LDAP secure connections.
	LdapsPort *int64 `json:"ldapsPort,omitempty"`
	// The TCP port on which this server is listening for HTTP connections.
	HttpPort *int64 `json:"httpPort,omitempty"`
	// The TCP port on which this server is listening for HTTPS connections.
	HttpsPort *int64 `json:"httpsPort,omitempty"`
	// The replication TCP port.
	ReplicationPort *int64 `json:"replicationPort,omitempty"`
	// Specifies a unique identifier for the replication server on this server instance.
	ReplicationServerID *int64 `json:"replicationServerID,omitempty"`
	// Specifies a unique identifier for the Directory Server within the replication domain.
	ReplicationDomainServerID []int64 `json:"replicationDomainServerID,omitempty"`
	// The TCP port on which this server is listening for JMX connections.
	JmxPort *int64 `json:"jmxPort,omitempty"`
	// The TCP port on which this server is listening for JMX secure connections.
	JmxsPort          *int64                                   `json:"jmxsPort,omitempty"`
	PreferredSecurity *EnumserverInstancePreferredSecurityProp `json:"preferredSecurity,omitempty"`
	// Indicates whether StartTLS is enabled on this server.
	StartTLSEnabled *bool `json:"startTLSEnabled,omitempty"`
	// The set of base DNs under the root DSE.
	BaseDN []string `json:"baseDN,omitempty"`
	// The set of groups of which this server is a member.
	MemberOfServerGroup []string `json:"memberOfServerGroup,omitempty"`
}

// NewSyncServerInstanceResponse instantiates a new SyncServerInstanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncServerInstanceResponse(schemas []EnumsyncServerInstanceSchemaUrn, id string, serverInstanceName string, clusterName string, serverVersion string) *SyncServerInstanceResponse {
	this := SyncServerInstanceResponse{}
	this.Schemas = schemas
	this.Id = id
	this.ServerInstanceName = serverInstanceName
	this.ClusterName = clusterName
	this.ServerVersion = serverVersion
	return &this
}

// NewSyncServerInstanceResponseWithDefaults instantiates a new SyncServerInstanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncServerInstanceResponseWithDefaults() *SyncServerInstanceResponse {
	this := SyncServerInstanceResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *SyncServerInstanceResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *SyncServerInstanceResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *SyncServerInstanceResponse) GetSchemas() []EnumsyncServerInstanceSchemaUrn {
	if o == nil {
		var ret []EnumsyncServerInstanceSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetSchemasOk() ([]EnumsyncServerInstanceSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *SyncServerInstanceResponse) SetSchemas(v []EnumsyncServerInstanceSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *SyncServerInstanceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SyncServerInstanceResponse) SetId(v string) {
	o.Id = v
}

// GetServerInstanceType returns the ServerInstanceType field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetServerInstanceType() EnumserverInstanceServerInstanceTypeProp {
	if o == nil || IsNil(o.ServerInstanceType) {
		var ret EnumserverInstanceServerInstanceTypeProp
		return ret
	}
	return *o.ServerInstanceType
}

// GetServerInstanceTypeOk returns a tuple with the ServerInstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetServerInstanceTypeOk() (*EnumserverInstanceServerInstanceTypeProp, bool) {
	if o == nil || IsNil(o.ServerInstanceType) {
		return nil, false
	}
	return o.ServerInstanceType, true
}

// HasServerInstanceType returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasServerInstanceType() bool {
	if o != nil && !IsNil(o.ServerInstanceType) {
		return true
	}

	return false
}

// SetServerInstanceType gets a reference to the given EnumserverInstanceServerInstanceTypeProp and assigns it to the ServerInstanceType field.
func (o *SyncServerInstanceResponse) SetServerInstanceType(v EnumserverInstanceServerInstanceTypeProp) {
	o.ServerInstanceType = &v
}

// GetServerInstanceName returns the ServerInstanceName field value
func (o *SyncServerInstanceResponse) GetServerInstanceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerInstanceName
}

// GetServerInstanceNameOk returns a tuple with the ServerInstanceName field value
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetServerInstanceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerInstanceName, true
}

// SetServerInstanceName sets field value
func (o *SyncServerInstanceResponse) SetServerInstanceName(v string) {
	o.ServerInstanceName = v
}

// GetClusterName returns the ClusterName field value
func (o *SyncServerInstanceResponse) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *SyncServerInstanceResponse) SetClusterName(v string) {
	o.ClusterName = v
}

// GetServerInstanceLocation returns the ServerInstanceLocation field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetServerInstanceLocation() string {
	if o == nil || IsNil(o.ServerInstanceLocation) {
		var ret string
		return ret
	}
	return *o.ServerInstanceLocation
}

// GetServerInstanceLocationOk returns a tuple with the ServerInstanceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetServerInstanceLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ServerInstanceLocation) {
		return nil, false
	}
	return o.ServerInstanceLocation, true
}

// HasServerInstanceLocation returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasServerInstanceLocation() bool {
	if o != nil && !IsNil(o.ServerInstanceLocation) {
		return true
	}

	return false
}

// SetServerInstanceLocation gets a reference to the given string and assigns it to the ServerInstanceLocation field.
func (o *SyncServerInstanceResponse) SetServerInstanceLocation(v string) {
	o.ServerInstanceLocation = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *SyncServerInstanceResponse) SetHostname(v string) {
	o.Hostname = &v
}

// GetServerRoot returns the ServerRoot field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetServerRoot() string {
	if o == nil || IsNil(o.ServerRoot) {
		var ret string
		return ret
	}
	return *o.ServerRoot
}

// GetServerRootOk returns a tuple with the ServerRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetServerRootOk() (*string, bool) {
	if o == nil || IsNil(o.ServerRoot) {
		return nil, false
	}
	return o.ServerRoot, true
}

// HasServerRoot returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasServerRoot() bool {
	if o != nil && !IsNil(o.ServerRoot) {
		return true
	}

	return false
}

// SetServerRoot gets a reference to the given string and assigns it to the ServerRoot field.
func (o *SyncServerInstanceResponse) SetServerRoot(v string) {
	o.ServerRoot = &v
}

// GetServerVersion returns the ServerVersion field value
func (o *SyncServerInstanceResponse) GetServerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerVersion
}

// GetServerVersionOk returns a tuple with the ServerVersion field value
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetServerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerVersion, true
}

// SetServerVersion sets field value
func (o *SyncServerInstanceResponse) SetServerVersion(v string) {
	o.ServerVersion = v
}

// GetInterServerCertificate returns the InterServerCertificate field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetInterServerCertificate() string {
	if o == nil || IsNil(o.InterServerCertificate) {
		var ret string
		return ret
	}
	return *o.InterServerCertificate
}

// GetInterServerCertificateOk returns a tuple with the InterServerCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetInterServerCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.InterServerCertificate) {
		return nil, false
	}
	return o.InterServerCertificate, true
}

// HasInterServerCertificate returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasInterServerCertificate() bool {
	if o != nil && !IsNil(o.InterServerCertificate) {
		return true
	}

	return false
}

// SetInterServerCertificate gets a reference to the given string and assigns it to the InterServerCertificate field.
func (o *SyncServerInstanceResponse) SetInterServerCertificate(v string) {
	o.InterServerCertificate = &v
}

// GetLdapPort returns the LdapPort field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetLdapPort() int64 {
	if o == nil || IsNil(o.LdapPort) {
		var ret int64
		return ret
	}
	return *o.LdapPort
}

// GetLdapPortOk returns a tuple with the LdapPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetLdapPortOk() (*int64, bool) {
	if o == nil || IsNil(o.LdapPort) {
		return nil, false
	}
	return o.LdapPort, true
}

// HasLdapPort returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasLdapPort() bool {
	if o != nil && !IsNil(o.LdapPort) {
		return true
	}

	return false
}

// SetLdapPort gets a reference to the given int64 and assigns it to the LdapPort field.
func (o *SyncServerInstanceResponse) SetLdapPort(v int64) {
	o.LdapPort = &v
}

// GetLdapsPort returns the LdapsPort field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetLdapsPort() int64 {
	if o == nil || IsNil(o.LdapsPort) {
		var ret int64
		return ret
	}
	return *o.LdapsPort
}

// GetLdapsPortOk returns a tuple with the LdapsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetLdapsPortOk() (*int64, bool) {
	if o == nil || IsNil(o.LdapsPort) {
		return nil, false
	}
	return o.LdapsPort, true
}

// HasLdapsPort returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasLdapsPort() bool {
	if o != nil && !IsNil(o.LdapsPort) {
		return true
	}

	return false
}

// SetLdapsPort gets a reference to the given int64 and assigns it to the LdapsPort field.
func (o *SyncServerInstanceResponse) SetLdapsPort(v int64) {
	o.LdapsPort = &v
}

// GetHttpPort returns the HttpPort field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetHttpPort() int64 {
	if o == nil || IsNil(o.HttpPort) {
		var ret int64
		return ret
	}
	return *o.HttpPort
}

// GetHttpPortOk returns a tuple with the HttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetHttpPortOk() (*int64, bool) {
	if o == nil || IsNil(o.HttpPort) {
		return nil, false
	}
	return o.HttpPort, true
}

// HasHttpPort returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasHttpPort() bool {
	if o != nil && !IsNil(o.HttpPort) {
		return true
	}

	return false
}

// SetHttpPort gets a reference to the given int64 and assigns it to the HttpPort field.
func (o *SyncServerInstanceResponse) SetHttpPort(v int64) {
	o.HttpPort = &v
}

// GetHttpsPort returns the HttpsPort field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetHttpsPort() int64 {
	if o == nil || IsNil(o.HttpsPort) {
		var ret int64
		return ret
	}
	return *o.HttpsPort
}

// GetHttpsPortOk returns a tuple with the HttpsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetHttpsPortOk() (*int64, bool) {
	if o == nil || IsNil(o.HttpsPort) {
		return nil, false
	}
	return o.HttpsPort, true
}

// HasHttpsPort returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasHttpsPort() bool {
	if o != nil && !IsNil(o.HttpsPort) {
		return true
	}

	return false
}

// SetHttpsPort gets a reference to the given int64 and assigns it to the HttpsPort field.
func (o *SyncServerInstanceResponse) SetHttpsPort(v int64) {
	o.HttpsPort = &v
}

// GetReplicationPort returns the ReplicationPort field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetReplicationPort() int64 {
	if o == nil || IsNil(o.ReplicationPort) {
		var ret int64
		return ret
	}
	return *o.ReplicationPort
}

// GetReplicationPortOk returns a tuple with the ReplicationPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetReplicationPortOk() (*int64, bool) {
	if o == nil || IsNil(o.ReplicationPort) {
		return nil, false
	}
	return o.ReplicationPort, true
}

// HasReplicationPort returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasReplicationPort() bool {
	if o != nil && !IsNil(o.ReplicationPort) {
		return true
	}

	return false
}

// SetReplicationPort gets a reference to the given int64 and assigns it to the ReplicationPort field.
func (o *SyncServerInstanceResponse) SetReplicationPort(v int64) {
	o.ReplicationPort = &v
}

// GetReplicationServerID returns the ReplicationServerID field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetReplicationServerID() int64 {
	if o == nil || IsNil(o.ReplicationServerID) {
		var ret int64
		return ret
	}
	return *o.ReplicationServerID
}

// GetReplicationServerIDOk returns a tuple with the ReplicationServerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetReplicationServerIDOk() (*int64, bool) {
	if o == nil || IsNil(o.ReplicationServerID) {
		return nil, false
	}
	return o.ReplicationServerID, true
}

// HasReplicationServerID returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasReplicationServerID() bool {
	if o != nil && !IsNil(o.ReplicationServerID) {
		return true
	}

	return false
}

// SetReplicationServerID gets a reference to the given int64 and assigns it to the ReplicationServerID field.
func (o *SyncServerInstanceResponse) SetReplicationServerID(v int64) {
	o.ReplicationServerID = &v
}

// GetReplicationDomainServerID returns the ReplicationDomainServerID field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetReplicationDomainServerID() []int64 {
	if o == nil || IsNil(o.ReplicationDomainServerID) {
		var ret []int64
		return ret
	}
	return o.ReplicationDomainServerID
}

// GetReplicationDomainServerIDOk returns a tuple with the ReplicationDomainServerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetReplicationDomainServerIDOk() ([]int64, bool) {
	if o == nil || IsNil(o.ReplicationDomainServerID) {
		return nil, false
	}
	return o.ReplicationDomainServerID, true
}

// HasReplicationDomainServerID returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasReplicationDomainServerID() bool {
	if o != nil && !IsNil(o.ReplicationDomainServerID) {
		return true
	}

	return false
}

// SetReplicationDomainServerID gets a reference to the given []int64 and assigns it to the ReplicationDomainServerID field.
func (o *SyncServerInstanceResponse) SetReplicationDomainServerID(v []int64) {
	o.ReplicationDomainServerID = v
}

// GetJmxPort returns the JmxPort field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetJmxPort() int64 {
	if o == nil || IsNil(o.JmxPort) {
		var ret int64
		return ret
	}
	return *o.JmxPort
}

// GetJmxPortOk returns a tuple with the JmxPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetJmxPortOk() (*int64, bool) {
	if o == nil || IsNil(o.JmxPort) {
		return nil, false
	}
	return o.JmxPort, true
}

// HasJmxPort returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasJmxPort() bool {
	if o != nil && !IsNil(o.JmxPort) {
		return true
	}

	return false
}

// SetJmxPort gets a reference to the given int64 and assigns it to the JmxPort field.
func (o *SyncServerInstanceResponse) SetJmxPort(v int64) {
	o.JmxPort = &v
}

// GetJmxsPort returns the JmxsPort field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetJmxsPort() int64 {
	if o == nil || IsNil(o.JmxsPort) {
		var ret int64
		return ret
	}
	return *o.JmxsPort
}

// GetJmxsPortOk returns a tuple with the JmxsPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetJmxsPortOk() (*int64, bool) {
	if o == nil || IsNil(o.JmxsPort) {
		return nil, false
	}
	return o.JmxsPort, true
}

// HasJmxsPort returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasJmxsPort() bool {
	if o != nil && !IsNil(o.JmxsPort) {
		return true
	}

	return false
}

// SetJmxsPort gets a reference to the given int64 and assigns it to the JmxsPort field.
func (o *SyncServerInstanceResponse) SetJmxsPort(v int64) {
	o.JmxsPort = &v
}

// GetPreferredSecurity returns the PreferredSecurity field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetPreferredSecurity() EnumserverInstancePreferredSecurityProp {
	if o == nil || IsNil(o.PreferredSecurity) {
		var ret EnumserverInstancePreferredSecurityProp
		return ret
	}
	return *o.PreferredSecurity
}

// GetPreferredSecurityOk returns a tuple with the PreferredSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetPreferredSecurityOk() (*EnumserverInstancePreferredSecurityProp, bool) {
	if o == nil || IsNil(o.PreferredSecurity) {
		return nil, false
	}
	return o.PreferredSecurity, true
}

// HasPreferredSecurity returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasPreferredSecurity() bool {
	if o != nil && !IsNil(o.PreferredSecurity) {
		return true
	}

	return false
}

// SetPreferredSecurity gets a reference to the given EnumserverInstancePreferredSecurityProp and assigns it to the PreferredSecurity field.
func (o *SyncServerInstanceResponse) SetPreferredSecurity(v EnumserverInstancePreferredSecurityProp) {
	o.PreferredSecurity = &v
}

// GetStartTLSEnabled returns the StartTLSEnabled field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetStartTLSEnabled() bool {
	if o == nil || IsNil(o.StartTLSEnabled) {
		var ret bool
		return ret
	}
	return *o.StartTLSEnabled
}

// GetStartTLSEnabledOk returns a tuple with the StartTLSEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetStartTLSEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StartTLSEnabled) {
		return nil, false
	}
	return o.StartTLSEnabled, true
}

// HasStartTLSEnabled returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasStartTLSEnabled() bool {
	if o != nil && !IsNil(o.StartTLSEnabled) {
		return true
	}

	return false
}

// SetStartTLSEnabled gets a reference to the given bool and assigns it to the StartTLSEnabled field.
func (o *SyncServerInstanceResponse) SetStartTLSEnabled(v bool) {
	o.StartTLSEnabled = &v
}

// GetBaseDN returns the BaseDN field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetBaseDN() []string {
	if o == nil || IsNil(o.BaseDN) {
		var ret []string
		return ret
	}
	return o.BaseDN
}

// GetBaseDNOk returns a tuple with the BaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetBaseDNOk() ([]string, bool) {
	if o == nil || IsNil(o.BaseDN) {
		return nil, false
	}
	return o.BaseDN, true
}

// HasBaseDN returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasBaseDN() bool {
	if o != nil && !IsNil(o.BaseDN) {
		return true
	}

	return false
}

// SetBaseDN gets a reference to the given []string and assigns it to the BaseDN field.
func (o *SyncServerInstanceResponse) SetBaseDN(v []string) {
	o.BaseDN = v
}

// GetMemberOfServerGroup returns the MemberOfServerGroup field value if set, zero value otherwise.
func (o *SyncServerInstanceResponse) GetMemberOfServerGroup() []string {
	if o == nil || IsNil(o.MemberOfServerGroup) {
		var ret []string
		return ret
	}
	return o.MemberOfServerGroup
}

// GetMemberOfServerGroupOk returns a tuple with the MemberOfServerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncServerInstanceResponse) GetMemberOfServerGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.MemberOfServerGroup) {
		return nil, false
	}
	return o.MemberOfServerGroup, true
}

// HasMemberOfServerGroup returns a boolean if a field has been set.
func (o *SyncServerInstanceResponse) HasMemberOfServerGroup() bool {
	if o != nil && !IsNil(o.MemberOfServerGroup) {
		return true
	}

	return false
}

// SetMemberOfServerGroup gets a reference to the given []string and assigns it to the MemberOfServerGroup field.
func (o *SyncServerInstanceResponse) SetMemberOfServerGroup(v []string) {
	o.MemberOfServerGroup = v
}

func (o SyncServerInstanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncServerInstanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	if !IsNil(o.ServerInstanceType) {
		toSerialize["serverInstanceType"] = o.ServerInstanceType
	}
	toSerialize["serverInstanceName"] = o.ServerInstanceName
	toSerialize["clusterName"] = o.ClusterName
	if !IsNil(o.ServerInstanceLocation) {
		toSerialize["serverInstanceLocation"] = o.ServerInstanceLocation
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.ServerRoot) {
		toSerialize["serverRoot"] = o.ServerRoot
	}
	toSerialize["serverVersion"] = o.ServerVersion
	if !IsNil(o.InterServerCertificate) {
		toSerialize["interServerCertificate"] = o.InterServerCertificate
	}
	if !IsNil(o.LdapPort) {
		toSerialize["ldapPort"] = o.LdapPort
	}
	if !IsNil(o.LdapsPort) {
		toSerialize["ldapsPort"] = o.LdapsPort
	}
	if !IsNil(o.HttpPort) {
		toSerialize["httpPort"] = o.HttpPort
	}
	if !IsNil(o.HttpsPort) {
		toSerialize["httpsPort"] = o.HttpsPort
	}
	if !IsNil(o.ReplicationPort) {
		toSerialize["replicationPort"] = o.ReplicationPort
	}
	if !IsNil(o.ReplicationServerID) {
		toSerialize["replicationServerID"] = o.ReplicationServerID
	}
	if !IsNil(o.ReplicationDomainServerID) {
		toSerialize["replicationDomainServerID"] = o.ReplicationDomainServerID
	}
	if !IsNil(o.JmxPort) {
		toSerialize["jmxPort"] = o.JmxPort
	}
	if !IsNil(o.JmxsPort) {
		toSerialize["jmxsPort"] = o.JmxsPort
	}
	if !IsNil(o.PreferredSecurity) {
		toSerialize["preferredSecurity"] = o.PreferredSecurity
	}
	if !IsNil(o.StartTLSEnabled) {
		toSerialize["startTLSEnabled"] = o.StartTLSEnabled
	}
	if !IsNil(o.BaseDN) {
		toSerialize["baseDN"] = o.BaseDN
	}
	if !IsNil(o.MemberOfServerGroup) {
		toSerialize["memberOfServerGroup"] = o.MemberOfServerGroup
	}
	return toSerialize, nil
}

type NullableSyncServerInstanceResponse struct {
	value *SyncServerInstanceResponse
	isSet bool
}

func (v NullableSyncServerInstanceResponse) Get() *SyncServerInstanceResponse {
	return v.value
}

func (v *NullableSyncServerInstanceResponse) Set(val *SyncServerInstanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncServerInstanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncServerInstanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncServerInstanceResponse(val *SyncServerInstanceResponse) *NullableSyncServerInstanceResponse {
	return &NullableSyncServerInstanceResponse{value: val, isSet: true}
}

func (v NullableSyncServerInstanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncServerInstanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
