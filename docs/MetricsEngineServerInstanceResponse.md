# MetricsEngineServerInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnummetricsEngineServerInstanceSchemaUrn**](EnummetricsEngineServerInstanceSchemaUrn.md) |  | 
**ServerInstanceType** | Pointer to [**EnumserverInstanceServerInstanceTypeProp**](EnumserverInstanceServerInstanceTypeProp.md) |  | [optional] 
**ServerInstanceName** | **string** | The name of this Server Instance. The instance name needs to be unique if this server will be part of a topology of servers that are connected to each other. Once set, it may not be changed. | 
**ClusterName** | **string** | The name of the cluster to which this Server Instance belongs. Server instances within the same cluster will share the same cluster-wide configuration. | 
**ServerInstanceLocation** | Pointer to **string** | Specifies the location for the Server Instance. | [optional] 
**Hostname** | Pointer to **string** | The name of the host where this Server Instance is installed. | [optional] 
**ServerRoot** | Pointer to **string** | The file system path where this Server Instance is installed. | [optional] 
**ServerVersion** | **string** | The version of the server. | 
**InterServerCertificate** | Pointer to **string** | The public component of the certificate used by this instance to protect inter-server communication and to perform server-specific encryption. This will generally be managed by the server and should only be altered by administrators under explicit direction from Ping Identity support personnel. | [optional] 
**LdapPort** | Pointer to **int32** | The TCP port on which this server is listening for LDAP connections. | [optional] 
**LdapsPort** | Pointer to **int32** | The TCP port on which this server is listening for LDAP secure connections. | [optional] 
**HttpPort** | Pointer to **int32** | The TCP port on which this server is listening for HTTP connections. | [optional] 
**HttpsPort** | Pointer to **int32** | The TCP port on which this server is listening for HTTPS connections. | [optional] 
**ReplicationPort** | Pointer to **int32** | The replication TCP port. | [optional] 
**ReplicationServerID** | Pointer to **int32** | Specifies a unique identifier for the replication server on this server instance. | [optional] 
**ReplicationDomainServerID** | Pointer to **[]int32** | Specifies a unique identifier for the Directory Server within the replication domain. | [optional] 
**JmxPort** | Pointer to **int32** | The TCP port on which this server is listening for JMX connections. | [optional] 
**JmxsPort** | Pointer to **int32** | The TCP port on which this server is listening for JMX secure connections. | [optional] 
**PreferredSecurity** | Pointer to [**EnumserverInstancePreferredSecurityProp**](EnumserverInstancePreferredSecurityProp.md) |  | [optional] 
**StartTLSEnabled** | Pointer to **bool** | Indicates whether StartTLS is enabled on this server. | [optional] 
**BaseDN** | Pointer to **[]string** | The set of base DNs under the root DSE. | [optional] 
**MemberOfServerGroup** | Pointer to **[]string** | The set of groups of which this server is a member. | [optional] 

## Methods

### NewMetricsEngineServerInstanceResponse

`func NewMetricsEngineServerInstanceResponse(schemas []EnummetricsEngineServerInstanceSchemaUrn, serverInstanceName string, clusterName string, serverVersion string, ) *MetricsEngineServerInstanceResponse`

NewMetricsEngineServerInstanceResponse instantiates a new MetricsEngineServerInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsEngineServerInstanceResponseWithDefaults

`func NewMetricsEngineServerInstanceResponseWithDefaults() *MetricsEngineServerInstanceResponse`

NewMetricsEngineServerInstanceResponseWithDefaults instantiates a new MetricsEngineServerInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *MetricsEngineServerInstanceResponse) GetSchemas() []EnummetricsEngineServerInstanceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *MetricsEngineServerInstanceResponse) GetSchemasOk() (*[]EnummetricsEngineServerInstanceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *MetricsEngineServerInstanceResponse) SetSchemas(v []EnummetricsEngineServerInstanceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerInstanceType

`func (o *MetricsEngineServerInstanceResponse) GetServerInstanceType() EnumserverInstanceServerInstanceTypeProp`

GetServerInstanceType returns the ServerInstanceType field if non-nil, zero value otherwise.

### GetServerInstanceTypeOk

`func (o *MetricsEngineServerInstanceResponse) GetServerInstanceTypeOk() (*EnumserverInstanceServerInstanceTypeProp, bool)`

GetServerInstanceTypeOk returns a tuple with the ServerInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceType

`func (o *MetricsEngineServerInstanceResponse) SetServerInstanceType(v EnumserverInstanceServerInstanceTypeProp)`

SetServerInstanceType sets ServerInstanceType field to given value.

### HasServerInstanceType

`func (o *MetricsEngineServerInstanceResponse) HasServerInstanceType() bool`

HasServerInstanceType returns a boolean if a field has been set.

### GetServerInstanceName

`func (o *MetricsEngineServerInstanceResponse) GetServerInstanceName() string`

GetServerInstanceName returns the ServerInstanceName field if non-nil, zero value otherwise.

### GetServerInstanceNameOk

`func (o *MetricsEngineServerInstanceResponse) GetServerInstanceNameOk() (*string, bool)`

GetServerInstanceNameOk returns a tuple with the ServerInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceName

`func (o *MetricsEngineServerInstanceResponse) SetServerInstanceName(v string)`

SetServerInstanceName sets ServerInstanceName field to given value.


### GetClusterName

`func (o *MetricsEngineServerInstanceResponse) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *MetricsEngineServerInstanceResponse) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *MetricsEngineServerInstanceResponse) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetServerInstanceLocation

`func (o *MetricsEngineServerInstanceResponse) GetServerInstanceLocation() string`

GetServerInstanceLocation returns the ServerInstanceLocation field if non-nil, zero value otherwise.

### GetServerInstanceLocationOk

`func (o *MetricsEngineServerInstanceResponse) GetServerInstanceLocationOk() (*string, bool)`

GetServerInstanceLocationOk returns a tuple with the ServerInstanceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceLocation

`func (o *MetricsEngineServerInstanceResponse) SetServerInstanceLocation(v string)`

SetServerInstanceLocation sets ServerInstanceLocation field to given value.

### HasServerInstanceLocation

`func (o *MetricsEngineServerInstanceResponse) HasServerInstanceLocation() bool`

HasServerInstanceLocation returns a boolean if a field has been set.

### GetHostname

`func (o *MetricsEngineServerInstanceResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *MetricsEngineServerInstanceResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *MetricsEngineServerInstanceResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *MetricsEngineServerInstanceResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServerRoot

`func (o *MetricsEngineServerInstanceResponse) GetServerRoot() string`

GetServerRoot returns the ServerRoot field if non-nil, zero value otherwise.

### GetServerRootOk

`func (o *MetricsEngineServerInstanceResponse) GetServerRootOk() (*string, bool)`

GetServerRootOk returns a tuple with the ServerRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRoot

`func (o *MetricsEngineServerInstanceResponse) SetServerRoot(v string)`

SetServerRoot sets ServerRoot field to given value.

### HasServerRoot

`func (o *MetricsEngineServerInstanceResponse) HasServerRoot() bool`

HasServerRoot returns a boolean if a field has been set.

### GetServerVersion

`func (o *MetricsEngineServerInstanceResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *MetricsEngineServerInstanceResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *MetricsEngineServerInstanceResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetInterServerCertificate

`func (o *MetricsEngineServerInstanceResponse) GetInterServerCertificate() string`

GetInterServerCertificate returns the InterServerCertificate field if non-nil, zero value otherwise.

### GetInterServerCertificateOk

`func (o *MetricsEngineServerInstanceResponse) GetInterServerCertificateOk() (*string, bool)`

GetInterServerCertificateOk returns a tuple with the InterServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterServerCertificate

`func (o *MetricsEngineServerInstanceResponse) SetInterServerCertificate(v string)`

SetInterServerCertificate sets InterServerCertificate field to given value.

### HasInterServerCertificate

`func (o *MetricsEngineServerInstanceResponse) HasInterServerCertificate() bool`

HasInterServerCertificate returns a boolean if a field has been set.

### GetLdapPort

`func (o *MetricsEngineServerInstanceResponse) GetLdapPort() int32`

GetLdapPort returns the LdapPort field if non-nil, zero value otherwise.

### GetLdapPortOk

`func (o *MetricsEngineServerInstanceResponse) GetLdapPortOk() (*int32, bool)`

GetLdapPortOk returns a tuple with the LdapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPort

`func (o *MetricsEngineServerInstanceResponse) SetLdapPort(v int32)`

SetLdapPort sets LdapPort field to given value.

### HasLdapPort

`func (o *MetricsEngineServerInstanceResponse) HasLdapPort() bool`

HasLdapPort returns a boolean if a field has been set.

### GetLdapsPort

`func (o *MetricsEngineServerInstanceResponse) GetLdapsPort() int32`

GetLdapsPort returns the LdapsPort field if non-nil, zero value otherwise.

### GetLdapsPortOk

`func (o *MetricsEngineServerInstanceResponse) GetLdapsPortOk() (*int32, bool)`

GetLdapsPortOk returns a tuple with the LdapsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapsPort

`func (o *MetricsEngineServerInstanceResponse) SetLdapsPort(v int32)`

SetLdapsPort sets LdapsPort field to given value.

### HasLdapsPort

`func (o *MetricsEngineServerInstanceResponse) HasLdapsPort() bool`

HasLdapsPort returns a boolean if a field has been set.

### GetHttpPort

`func (o *MetricsEngineServerInstanceResponse) GetHttpPort() int32`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *MetricsEngineServerInstanceResponse) GetHttpPortOk() (*int32, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *MetricsEngineServerInstanceResponse) SetHttpPort(v int32)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *MetricsEngineServerInstanceResponse) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *MetricsEngineServerInstanceResponse) GetHttpsPort() int32`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *MetricsEngineServerInstanceResponse) GetHttpsPortOk() (*int32, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *MetricsEngineServerInstanceResponse) SetHttpsPort(v int32)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *MetricsEngineServerInstanceResponse) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetReplicationPort

`func (o *MetricsEngineServerInstanceResponse) GetReplicationPort() int32`

GetReplicationPort returns the ReplicationPort field if non-nil, zero value otherwise.

### GetReplicationPortOk

`func (o *MetricsEngineServerInstanceResponse) GetReplicationPortOk() (*int32, bool)`

GetReplicationPortOk returns a tuple with the ReplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPort

`func (o *MetricsEngineServerInstanceResponse) SetReplicationPort(v int32)`

SetReplicationPort sets ReplicationPort field to given value.

### HasReplicationPort

`func (o *MetricsEngineServerInstanceResponse) HasReplicationPort() bool`

HasReplicationPort returns a boolean if a field has been set.

### GetReplicationServerID

`func (o *MetricsEngineServerInstanceResponse) GetReplicationServerID() int32`

GetReplicationServerID returns the ReplicationServerID field if non-nil, zero value otherwise.

### GetReplicationServerIDOk

`func (o *MetricsEngineServerInstanceResponse) GetReplicationServerIDOk() (*int32, bool)`

GetReplicationServerIDOk returns a tuple with the ReplicationServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationServerID

`func (o *MetricsEngineServerInstanceResponse) SetReplicationServerID(v int32)`

SetReplicationServerID sets ReplicationServerID field to given value.

### HasReplicationServerID

`func (o *MetricsEngineServerInstanceResponse) HasReplicationServerID() bool`

HasReplicationServerID returns a boolean if a field has been set.

### GetReplicationDomainServerID

`func (o *MetricsEngineServerInstanceResponse) GetReplicationDomainServerID() []int32`

GetReplicationDomainServerID returns the ReplicationDomainServerID field if non-nil, zero value otherwise.

### GetReplicationDomainServerIDOk

`func (o *MetricsEngineServerInstanceResponse) GetReplicationDomainServerIDOk() (*[]int32, bool)`

GetReplicationDomainServerIDOk returns a tuple with the ReplicationDomainServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationDomainServerID

`func (o *MetricsEngineServerInstanceResponse) SetReplicationDomainServerID(v []int32)`

SetReplicationDomainServerID sets ReplicationDomainServerID field to given value.

### HasReplicationDomainServerID

`func (o *MetricsEngineServerInstanceResponse) HasReplicationDomainServerID() bool`

HasReplicationDomainServerID returns a boolean if a field has been set.

### GetJmxPort

`func (o *MetricsEngineServerInstanceResponse) GetJmxPort() int32`

GetJmxPort returns the JmxPort field if non-nil, zero value otherwise.

### GetJmxPortOk

`func (o *MetricsEngineServerInstanceResponse) GetJmxPortOk() (*int32, bool)`

GetJmxPortOk returns a tuple with the JmxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxPort

`func (o *MetricsEngineServerInstanceResponse) SetJmxPort(v int32)`

SetJmxPort sets JmxPort field to given value.

### HasJmxPort

`func (o *MetricsEngineServerInstanceResponse) HasJmxPort() bool`

HasJmxPort returns a boolean if a field has been set.

### GetJmxsPort

`func (o *MetricsEngineServerInstanceResponse) GetJmxsPort() int32`

GetJmxsPort returns the JmxsPort field if non-nil, zero value otherwise.

### GetJmxsPortOk

`func (o *MetricsEngineServerInstanceResponse) GetJmxsPortOk() (*int32, bool)`

GetJmxsPortOk returns a tuple with the JmxsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxsPort

`func (o *MetricsEngineServerInstanceResponse) SetJmxsPort(v int32)`

SetJmxsPort sets JmxsPort field to given value.

### HasJmxsPort

`func (o *MetricsEngineServerInstanceResponse) HasJmxsPort() bool`

HasJmxsPort returns a boolean if a field has been set.

### GetPreferredSecurity

`func (o *MetricsEngineServerInstanceResponse) GetPreferredSecurity() EnumserverInstancePreferredSecurityProp`

GetPreferredSecurity returns the PreferredSecurity field if non-nil, zero value otherwise.

### GetPreferredSecurityOk

`func (o *MetricsEngineServerInstanceResponse) GetPreferredSecurityOk() (*EnumserverInstancePreferredSecurityProp, bool)`

GetPreferredSecurityOk returns a tuple with the PreferredSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSecurity

`func (o *MetricsEngineServerInstanceResponse) SetPreferredSecurity(v EnumserverInstancePreferredSecurityProp)`

SetPreferredSecurity sets PreferredSecurity field to given value.

### HasPreferredSecurity

`func (o *MetricsEngineServerInstanceResponse) HasPreferredSecurity() bool`

HasPreferredSecurity returns a boolean if a field has been set.

### GetStartTLSEnabled

`func (o *MetricsEngineServerInstanceResponse) GetStartTLSEnabled() bool`

GetStartTLSEnabled returns the StartTLSEnabled field if non-nil, zero value otherwise.

### GetStartTLSEnabledOk

`func (o *MetricsEngineServerInstanceResponse) GetStartTLSEnabledOk() (*bool, bool)`

GetStartTLSEnabledOk returns a tuple with the StartTLSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTLSEnabled

`func (o *MetricsEngineServerInstanceResponse) SetStartTLSEnabled(v bool)`

SetStartTLSEnabled sets StartTLSEnabled field to given value.

### HasStartTLSEnabled

`func (o *MetricsEngineServerInstanceResponse) HasStartTLSEnabled() bool`

HasStartTLSEnabled returns a boolean if a field has been set.

### GetBaseDN

`func (o *MetricsEngineServerInstanceResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *MetricsEngineServerInstanceResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *MetricsEngineServerInstanceResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *MetricsEngineServerInstanceResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMemberOfServerGroup

`func (o *MetricsEngineServerInstanceResponse) GetMemberOfServerGroup() []string`

GetMemberOfServerGroup returns the MemberOfServerGroup field if non-nil, zero value otherwise.

### GetMemberOfServerGroupOk

`func (o *MetricsEngineServerInstanceResponse) GetMemberOfServerGroupOk() (*[]string, bool)`

GetMemberOfServerGroupOk returns a tuple with the MemberOfServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfServerGroup

`func (o *MetricsEngineServerInstanceResponse) SetMemberOfServerGroup(v []string)`

SetMemberOfServerGroup sets MemberOfServerGroup field to given value.

### HasMemberOfServerGroup

`func (o *MetricsEngineServerInstanceResponse) HasMemberOfServerGroup() bool`

HasMemberOfServerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


