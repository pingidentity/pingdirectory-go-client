# ServerInstanceListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsyncServerInstanceSchemaUrn**](EnumsyncServerInstanceSchemaUrn.md) |  | 
**Id** | **string** | Name of the Server Instance | 
**ServerInstanceType** | Pointer to [**EnumserverInstanceServerInstanceTypeProp**](EnumserverInstanceServerInstanceTypeProp.md) |  | [optional] 
**ServerInstanceName** | **string** | The name of this Server Instance. The instance name needs to be unique if this server will be part of a topology of servers that are connected to each other. Once set, it may not be changed. | 
**ClusterName** | **string** | The name of the cluster to which this Server Instance belongs. Server instances within the same cluster will share the same cluster-wide configuration. | 
**ServerInstanceLocation** | Pointer to **string** | Specifies the location for the Server Instance. | [optional] 
**Hostname** | Pointer to **string** | The name of the host where this Server Instance is installed. | [optional] 
**ServerRoot** | Pointer to **string** | The file system path where this Server Instance is installed. | [optional] 
**ServerVersion** | **string** | The version of the server. | 
**InterServerCertificate** | Pointer to **string** | The public component of the certificate used by this instance to protect inter-server communication and to perform server-specific encryption. This will generally be managed by the server and should only be altered by administrators under explicit direction from Ping Identity support personnel. | [optional] 
**LdapPort** | Pointer to **int64** | The TCP port on which this server is listening for LDAP connections. | [optional] 
**LdapsPort** | Pointer to **int64** | The TCP port on which this server is listening for LDAP secure connections. | [optional] 
**HttpPort** | Pointer to **int64** | The TCP port on which this server is listening for HTTP connections. | [optional] 
**HttpsPort** | Pointer to **int64** | The TCP port on which this server is listening for HTTPS connections. | [optional] 
**ReplicationPort** | Pointer to **int64** | The replication TCP port. | [optional] 
**ReplicationServerID** | Pointer to **int64** | Specifies a unique identifier for the replication server on this server instance. | [optional] 
**ReplicationDomainServerID** | Pointer to **[]int64** | Specifies a unique identifier for the Directory Server within the replication domain. | [optional] 
**JmxPort** | Pointer to **int64** | The TCP port on which this server is listening for JMX connections. | [optional] 
**JmxsPort** | Pointer to **int64** | The TCP port on which this server is listening for JMX secure connections. | [optional] 
**PreferredSecurity** | Pointer to [**EnumserverInstancePreferredSecurityProp**](EnumserverInstancePreferredSecurityProp.md) |  | [optional] 
**StartTLSEnabled** | Pointer to **bool** | Indicates whether StartTLS is enabled on this server. | [optional] 
**BaseDN** | Pointer to **[]string** | The set of base DNs under the root DSE. | [optional] 
**MemberOfServerGroup** | Pointer to **[]string** | The set of groups of which this server is a member. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**ReplicationSetName** | Pointer to **string** | The name of the replication set assigned to this Directory Server. Restricted domains are only replicated within instances using the same replication set name. | [optional] 
**LoadBalancingAlgorithmName** | Pointer to **[]string** | The name of the configuration object for a load-balancing algorithm that should include this server. | [optional] 

## Methods

### NewServerInstanceListResponseResourcesInner

`func NewServerInstanceListResponseResourcesInner(schemas []EnumsyncServerInstanceSchemaUrn, id string, serverInstanceName string, clusterName string, serverVersion string, ) *ServerInstanceListResponseResourcesInner`

NewServerInstanceListResponseResourcesInner instantiates a new ServerInstanceListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInstanceListResponseResourcesInnerWithDefaults

`func NewServerInstanceListResponseResourcesInnerWithDefaults() *ServerInstanceListResponseResourcesInner`

NewServerInstanceListResponseResourcesInnerWithDefaults instantiates a new ServerInstanceListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ServerInstanceListResponseResourcesInner) GetSchemas() []EnumsyncServerInstanceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ServerInstanceListResponseResourcesInner) GetSchemasOk() (*[]EnumsyncServerInstanceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ServerInstanceListResponseResourcesInner) SetSchemas(v []EnumsyncServerInstanceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *ServerInstanceListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerInstanceListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerInstanceListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetServerInstanceType

`func (o *ServerInstanceListResponseResourcesInner) GetServerInstanceType() EnumserverInstanceServerInstanceTypeProp`

GetServerInstanceType returns the ServerInstanceType field if non-nil, zero value otherwise.

### GetServerInstanceTypeOk

`func (o *ServerInstanceListResponseResourcesInner) GetServerInstanceTypeOk() (*EnumserverInstanceServerInstanceTypeProp, bool)`

GetServerInstanceTypeOk returns a tuple with the ServerInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceType

`func (o *ServerInstanceListResponseResourcesInner) SetServerInstanceType(v EnumserverInstanceServerInstanceTypeProp)`

SetServerInstanceType sets ServerInstanceType field to given value.

### HasServerInstanceType

`func (o *ServerInstanceListResponseResourcesInner) HasServerInstanceType() bool`

HasServerInstanceType returns a boolean if a field has been set.

### GetServerInstanceName

`func (o *ServerInstanceListResponseResourcesInner) GetServerInstanceName() string`

GetServerInstanceName returns the ServerInstanceName field if non-nil, zero value otherwise.

### GetServerInstanceNameOk

`func (o *ServerInstanceListResponseResourcesInner) GetServerInstanceNameOk() (*string, bool)`

GetServerInstanceNameOk returns a tuple with the ServerInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceName

`func (o *ServerInstanceListResponseResourcesInner) SetServerInstanceName(v string)`

SetServerInstanceName sets ServerInstanceName field to given value.


### GetClusterName

`func (o *ServerInstanceListResponseResourcesInner) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ServerInstanceListResponseResourcesInner) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ServerInstanceListResponseResourcesInner) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetServerInstanceLocation

`func (o *ServerInstanceListResponseResourcesInner) GetServerInstanceLocation() string`

GetServerInstanceLocation returns the ServerInstanceLocation field if non-nil, zero value otherwise.

### GetServerInstanceLocationOk

`func (o *ServerInstanceListResponseResourcesInner) GetServerInstanceLocationOk() (*string, bool)`

GetServerInstanceLocationOk returns a tuple with the ServerInstanceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceLocation

`func (o *ServerInstanceListResponseResourcesInner) SetServerInstanceLocation(v string)`

SetServerInstanceLocation sets ServerInstanceLocation field to given value.

### HasServerInstanceLocation

`func (o *ServerInstanceListResponseResourcesInner) HasServerInstanceLocation() bool`

HasServerInstanceLocation returns a boolean if a field has been set.

### GetHostname

`func (o *ServerInstanceListResponseResourcesInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ServerInstanceListResponseResourcesInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ServerInstanceListResponseResourcesInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ServerInstanceListResponseResourcesInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServerRoot

`func (o *ServerInstanceListResponseResourcesInner) GetServerRoot() string`

GetServerRoot returns the ServerRoot field if non-nil, zero value otherwise.

### GetServerRootOk

`func (o *ServerInstanceListResponseResourcesInner) GetServerRootOk() (*string, bool)`

GetServerRootOk returns a tuple with the ServerRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRoot

`func (o *ServerInstanceListResponseResourcesInner) SetServerRoot(v string)`

SetServerRoot sets ServerRoot field to given value.

### HasServerRoot

`func (o *ServerInstanceListResponseResourcesInner) HasServerRoot() bool`

HasServerRoot returns a boolean if a field has been set.

### GetServerVersion

`func (o *ServerInstanceListResponseResourcesInner) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *ServerInstanceListResponseResourcesInner) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *ServerInstanceListResponseResourcesInner) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetInterServerCertificate

`func (o *ServerInstanceListResponseResourcesInner) GetInterServerCertificate() string`

GetInterServerCertificate returns the InterServerCertificate field if non-nil, zero value otherwise.

### GetInterServerCertificateOk

`func (o *ServerInstanceListResponseResourcesInner) GetInterServerCertificateOk() (*string, bool)`

GetInterServerCertificateOk returns a tuple with the InterServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterServerCertificate

`func (o *ServerInstanceListResponseResourcesInner) SetInterServerCertificate(v string)`

SetInterServerCertificate sets InterServerCertificate field to given value.

### HasInterServerCertificate

`func (o *ServerInstanceListResponseResourcesInner) HasInterServerCertificate() bool`

HasInterServerCertificate returns a boolean if a field has been set.

### GetLdapPort

`func (o *ServerInstanceListResponseResourcesInner) GetLdapPort() int64`

GetLdapPort returns the LdapPort field if non-nil, zero value otherwise.

### GetLdapPortOk

`func (o *ServerInstanceListResponseResourcesInner) GetLdapPortOk() (*int64, bool)`

GetLdapPortOk returns a tuple with the LdapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPort

`func (o *ServerInstanceListResponseResourcesInner) SetLdapPort(v int64)`

SetLdapPort sets LdapPort field to given value.

### HasLdapPort

`func (o *ServerInstanceListResponseResourcesInner) HasLdapPort() bool`

HasLdapPort returns a boolean if a field has been set.

### GetLdapsPort

`func (o *ServerInstanceListResponseResourcesInner) GetLdapsPort() int64`

GetLdapsPort returns the LdapsPort field if non-nil, zero value otherwise.

### GetLdapsPortOk

`func (o *ServerInstanceListResponseResourcesInner) GetLdapsPortOk() (*int64, bool)`

GetLdapsPortOk returns a tuple with the LdapsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapsPort

`func (o *ServerInstanceListResponseResourcesInner) SetLdapsPort(v int64)`

SetLdapsPort sets LdapsPort field to given value.

### HasLdapsPort

`func (o *ServerInstanceListResponseResourcesInner) HasLdapsPort() bool`

HasLdapsPort returns a boolean if a field has been set.

### GetHttpPort

`func (o *ServerInstanceListResponseResourcesInner) GetHttpPort() int64`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *ServerInstanceListResponseResourcesInner) GetHttpPortOk() (*int64, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *ServerInstanceListResponseResourcesInner) SetHttpPort(v int64)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *ServerInstanceListResponseResourcesInner) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *ServerInstanceListResponseResourcesInner) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *ServerInstanceListResponseResourcesInner) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *ServerInstanceListResponseResourcesInner) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *ServerInstanceListResponseResourcesInner) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetReplicationPort

`func (o *ServerInstanceListResponseResourcesInner) GetReplicationPort() int64`

GetReplicationPort returns the ReplicationPort field if non-nil, zero value otherwise.

### GetReplicationPortOk

`func (o *ServerInstanceListResponseResourcesInner) GetReplicationPortOk() (*int64, bool)`

GetReplicationPortOk returns a tuple with the ReplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPort

`func (o *ServerInstanceListResponseResourcesInner) SetReplicationPort(v int64)`

SetReplicationPort sets ReplicationPort field to given value.

### HasReplicationPort

`func (o *ServerInstanceListResponseResourcesInner) HasReplicationPort() bool`

HasReplicationPort returns a boolean if a field has been set.

### GetReplicationServerID

`func (o *ServerInstanceListResponseResourcesInner) GetReplicationServerID() int64`

GetReplicationServerID returns the ReplicationServerID field if non-nil, zero value otherwise.

### GetReplicationServerIDOk

`func (o *ServerInstanceListResponseResourcesInner) GetReplicationServerIDOk() (*int64, bool)`

GetReplicationServerIDOk returns a tuple with the ReplicationServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationServerID

`func (o *ServerInstanceListResponseResourcesInner) SetReplicationServerID(v int64)`

SetReplicationServerID sets ReplicationServerID field to given value.

### HasReplicationServerID

`func (o *ServerInstanceListResponseResourcesInner) HasReplicationServerID() bool`

HasReplicationServerID returns a boolean if a field has been set.

### GetReplicationDomainServerID

`func (o *ServerInstanceListResponseResourcesInner) GetReplicationDomainServerID() []int64`

GetReplicationDomainServerID returns the ReplicationDomainServerID field if non-nil, zero value otherwise.

### GetReplicationDomainServerIDOk

`func (o *ServerInstanceListResponseResourcesInner) GetReplicationDomainServerIDOk() (*[]int64, bool)`

GetReplicationDomainServerIDOk returns a tuple with the ReplicationDomainServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationDomainServerID

`func (o *ServerInstanceListResponseResourcesInner) SetReplicationDomainServerID(v []int64)`

SetReplicationDomainServerID sets ReplicationDomainServerID field to given value.

### HasReplicationDomainServerID

`func (o *ServerInstanceListResponseResourcesInner) HasReplicationDomainServerID() bool`

HasReplicationDomainServerID returns a boolean if a field has been set.

### GetJmxPort

`func (o *ServerInstanceListResponseResourcesInner) GetJmxPort() int64`

GetJmxPort returns the JmxPort field if non-nil, zero value otherwise.

### GetJmxPortOk

`func (o *ServerInstanceListResponseResourcesInner) GetJmxPortOk() (*int64, bool)`

GetJmxPortOk returns a tuple with the JmxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxPort

`func (o *ServerInstanceListResponseResourcesInner) SetJmxPort(v int64)`

SetJmxPort sets JmxPort field to given value.

### HasJmxPort

`func (o *ServerInstanceListResponseResourcesInner) HasJmxPort() bool`

HasJmxPort returns a boolean if a field has been set.

### GetJmxsPort

`func (o *ServerInstanceListResponseResourcesInner) GetJmxsPort() int64`

GetJmxsPort returns the JmxsPort field if non-nil, zero value otherwise.

### GetJmxsPortOk

`func (o *ServerInstanceListResponseResourcesInner) GetJmxsPortOk() (*int64, bool)`

GetJmxsPortOk returns a tuple with the JmxsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxsPort

`func (o *ServerInstanceListResponseResourcesInner) SetJmxsPort(v int64)`

SetJmxsPort sets JmxsPort field to given value.

### HasJmxsPort

`func (o *ServerInstanceListResponseResourcesInner) HasJmxsPort() bool`

HasJmxsPort returns a boolean if a field has been set.

### GetPreferredSecurity

`func (o *ServerInstanceListResponseResourcesInner) GetPreferredSecurity() EnumserverInstancePreferredSecurityProp`

GetPreferredSecurity returns the PreferredSecurity field if non-nil, zero value otherwise.

### GetPreferredSecurityOk

`func (o *ServerInstanceListResponseResourcesInner) GetPreferredSecurityOk() (*EnumserverInstancePreferredSecurityProp, bool)`

GetPreferredSecurityOk returns a tuple with the PreferredSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSecurity

`func (o *ServerInstanceListResponseResourcesInner) SetPreferredSecurity(v EnumserverInstancePreferredSecurityProp)`

SetPreferredSecurity sets PreferredSecurity field to given value.

### HasPreferredSecurity

`func (o *ServerInstanceListResponseResourcesInner) HasPreferredSecurity() bool`

HasPreferredSecurity returns a boolean if a field has been set.

### GetStartTLSEnabled

`func (o *ServerInstanceListResponseResourcesInner) GetStartTLSEnabled() bool`

GetStartTLSEnabled returns the StartTLSEnabled field if non-nil, zero value otherwise.

### GetStartTLSEnabledOk

`func (o *ServerInstanceListResponseResourcesInner) GetStartTLSEnabledOk() (*bool, bool)`

GetStartTLSEnabledOk returns a tuple with the StartTLSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTLSEnabled

`func (o *ServerInstanceListResponseResourcesInner) SetStartTLSEnabled(v bool)`

SetStartTLSEnabled sets StartTLSEnabled field to given value.

### HasStartTLSEnabled

`func (o *ServerInstanceListResponseResourcesInner) HasStartTLSEnabled() bool`

HasStartTLSEnabled returns a boolean if a field has been set.

### GetBaseDN

`func (o *ServerInstanceListResponseResourcesInner) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *ServerInstanceListResponseResourcesInner) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *ServerInstanceListResponseResourcesInner) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *ServerInstanceListResponseResourcesInner) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMemberOfServerGroup

`func (o *ServerInstanceListResponseResourcesInner) GetMemberOfServerGroup() []string`

GetMemberOfServerGroup returns the MemberOfServerGroup field if non-nil, zero value otherwise.

### GetMemberOfServerGroupOk

`func (o *ServerInstanceListResponseResourcesInner) GetMemberOfServerGroupOk() (*[]string, bool)`

GetMemberOfServerGroupOk returns a tuple with the MemberOfServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfServerGroup

`func (o *ServerInstanceListResponseResourcesInner) SetMemberOfServerGroup(v []string)`

SetMemberOfServerGroup sets MemberOfServerGroup field to given value.

### HasMemberOfServerGroup

`func (o *ServerInstanceListResponseResourcesInner) HasMemberOfServerGroup() bool`

HasMemberOfServerGroup returns a boolean if a field has been set.

### GetMeta

`func (o *ServerInstanceListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ServerInstanceListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ServerInstanceListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ServerInstanceListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *ServerInstanceListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *ServerInstanceListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *ServerInstanceListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *ServerInstanceListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetReplicationSetName

`func (o *ServerInstanceListResponseResourcesInner) GetReplicationSetName() string`

GetReplicationSetName returns the ReplicationSetName field if non-nil, zero value otherwise.

### GetReplicationSetNameOk

`func (o *ServerInstanceListResponseResourcesInner) GetReplicationSetNameOk() (*string, bool)`

GetReplicationSetNameOk returns a tuple with the ReplicationSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSetName

`func (o *ServerInstanceListResponseResourcesInner) SetReplicationSetName(v string)`

SetReplicationSetName sets ReplicationSetName field to given value.

### HasReplicationSetName

`func (o *ServerInstanceListResponseResourcesInner) HasReplicationSetName() bool`

HasReplicationSetName returns a boolean if a field has been set.

### GetLoadBalancingAlgorithmName

`func (o *ServerInstanceListResponseResourcesInner) GetLoadBalancingAlgorithmName() []string`

GetLoadBalancingAlgorithmName returns the LoadBalancingAlgorithmName field if non-nil, zero value otherwise.

### GetLoadBalancingAlgorithmNameOk

`func (o *ServerInstanceListResponseResourcesInner) GetLoadBalancingAlgorithmNameOk() (*[]string, bool)`

GetLoadBalancingAlgorithmNameOk returns a tuple with the LoadBalancingAlgorithmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingAlgorithmName

`func (o *ServerInstanceListResponseResourcesInner) SetLoadBalancingAlgorithmName(v []string)`

SetLoadBalancingAlgorithmName sets LoadBalancingAlgorithmName field to given value.

### HasLoadBalancingAlgorithmName

`func (o *ServerInstanceListResponseResourcesInner) HasLoadBalancingAlgorithmName() bool`

HasLoadBalancingAlgorithmName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


