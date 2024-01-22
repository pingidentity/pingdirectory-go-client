# SyncServerInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
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

## Methods

### NewSyncServerInstanceResponse

`func NewSyncServerInstanceResponse(schemas []EnumsyncServerInstanceSchemaUrn, id string, serverInstanceName string, clusterName string, serverVersion string, ) *SyncServerInstanceResponse`

NewSyncServerInstanceResponse instantiates a new SyncServerInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncServerInstanceResponseWithDefaults

`func NewSyncServerInstanceResponseWithDefaults() *SyncServerInstanceResponse`

NewSyncServerInstanceResponseWithDefaults instantiates a new SyncServerInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SyncServerInstanceResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SyncServerInstanceResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SyncServerInstanceResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SyncServerInstanceResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SyncServerInstanceResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SyncServerInstanceResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SyncServerInstanceResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SyncServerInstanceResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSchemas

`func (o *SyncServerInstanceResponse) GetSchemas() []EnumsyncServerInstanceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SyncServerInstanceResponse) GetSchemasOk() (*[]EnumsyncServerInstanceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SyncServerInstanceResponse) SetSchemas(v []EnumsyncServerInstanceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *SyncServerInstanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SyncServerInstanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SyncServerInstanceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetServerInstanceType

`func (o *SyncServerInstanceResponse) GetServerInstanceType() EnumserverInstanceServerInstanceTypeProp`

GetServerInstanceType returns the ServerInstanceType field if non-nil, zero value otherwise.

### GetServerInstanceTypeOk

`func (o *SyncServerInstanceResponse) GetServerInstanceTypeOk() (*EnumserverInstanceServerInstanceTypeProp, bool)`

GetServerInstanceTypeOk returns a tuple with the ServerInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceType

`func (o *SyncServerInstanceResponse) SetServerInstanceType(v EnumserverInstanceServerInstanceTypeProp)`

SetServerInstanceType sets ServerInstanceType field to given value.

### HasServerInstanceType

`func (o *SyncServerInstanceResponse) HasServerInstanceType() bool`

HasServerInstanceType returns a boolean if a field has been set.

### GetServerInstanceName

`func (o *SyncServerInstanceResponse) GetServerInstanceName() string`

GetServerInstanceName returns the ServerInstanceName field if non-nil, zero value otherwise.

### GetServerInstanceNameOk

`func (o *SyncServerInstanceResponse) GetServerInstanceNameOk() (*string, bool)`

GetServerInstanceNameOk returns a tuple with the ServerInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceName

`func (o *SyncServerInstanceResponse) SetServerInstanceName(v string)`

SetServerInstanceName sets ServerInstanceName field to given value.


### GetClusterName

`func (o *SyncServerInstanceResponse) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *SyncServerInstanceResponse) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *SyncServerInstanceResponse) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetServerInstanceLocation

`func (o *SyncServerInstanceResponse) GetServerInstanceLocation() string`

GetServerInstanceLocation returns the ServerInstanceLocation field if non-nil, zero value otherwise.

### GetServerInstanceLocationOk

`func (o *SyncServerInstanceResponse) GetServerInstanceLocationOk() (*string, bool)`

GetServerInstanceLocationOk returns a tuple with the ServerInstanceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceLocation

`func (o *SyncServerInstanceResponse) SetServerInstanceLocation(v string)`

SetServerInstanceLocation sets ServerInstanceLocation field to given value.

### HasServerInstanceLocation

`func (o *SyncServerInstanceResponse) HasServerInstanceLocation() bool`

HasServerInstanceLocation returns a boolean if a field has been set.

### GetHostname

`func (o *SyncServerInstanceResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *SyncServerInstanceResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *SyncServerInstanceResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *SyncServerInstanceResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServerRoot

`func (o *SyncServerInstanceResponse) GetServerRoot() string`

GetServerRoot returns the ServerRoot field if non-nil, zero value otherwise.

### GetServerRootOk

`func (o *SyncServerInstanceResponse) GetServerRootOk() (*string, bool)`

GetServerRootOk returns a tuple with the ServerRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRoot

`func (o *SyncServerInstanceResponse) SetServerRoot(v string)`

SetServerRoot sets ServerRoot field to given value.

### HasServerRoot

`func (o *SyncServerInstanceResponse) HasServerRoot() bool`

HasServerRoot returns a boolean if a field has been set.

### GetServerVersion

`func (o *SyncServerInstanceResponse) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *SyncServerInstanceResponse) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *SyncServerInstanceResponse) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetInterServerCertificate

`func (o *SyncServerInstanceResponse) GetInterServerCertificate() string`

GetInterServerCertificate returns the InterServerCertificate field if non-nil, zero value otherwise.

### GetInterServerCertificateOk

`func (o *SyncServerInstanceResponse) GetInterServerCertificateOk() (*string, bool)`

GetInterServerCertificateOk returns a tuple with the InterServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterServerCertificate

`func (o *SyncServerInstanceResponse) SetInterServerCertificate(v string)`

SetInterServerCertificate sets InterServerCertificate field to given value.

### HasInterServerCertificate

`func (o *SyncServerInstanceResponse) HasInterServerCertificate() bool`

HasInterServerCertificate returns a boolean if a field has been set.

### GetLdapPort

`func (o *SyncServerInstanceResponse) GetLdapPort() int64`

GetLdapPort returns the LdapPort field if non-nil, zero value otherwise.

### GetLdapPortOk

`func (o *SyncServerInstanceResponse) GetLdapPortOk() (*int64, bool)`

GetLdapPortOk returns a tuple with the LdapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPort

`func (o *SyncServerInstanceResponse) SetLdapPort(v int64)`

SetLdapPort sets LdapPort field to given value.

### HasLdapPort

`func (o *SyncServerInstanceResponse) HasLdapPort() bool`

HasLdapPort returns a boolean if a field has been set.

### GetLdapsPort

`func (o *SyncServerInstanceResponse) GetLdapsPort() int64`

GetLdapsPort returns the LdapsPort field if non-nil, zero value otherwise.

### GetLdapsPortOk

`func (o *SyncServerInstanceResponse) GetLdapsPortOk() (*int64, bool)`

GetLdapsPortOk returns a tuple with the LdapsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapsPort

`func (o *SyncServerInstanceResponse) SetLdapsPort(v int64)`

SetLdapsPort sets LdapsPort field to given value.

### HasLdapsPort

`func (o *SyncServerInstanceResponse) HasLdapsPort() bool`

HasLdapsPort returns a boolean if a field has been set.

### GetHttpPort

`func (o *SyncServerInstanceResponse) GetHttpPort() int64`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *SyncServerInstanceResponse) GetHttpPortOk() (*int64, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *SyncServerInstanceResponse) SetHttpPort(v int64)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *SyncServerInstanceResponse) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *SyncServerInstanceResponse) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *SyncServerInstanceResponse) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *SyncServerInstanceResponse) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *SyncServerInstanceResponse) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetReplicationPort

`func (o *SyncServerInstanceResponse) GetReplicationPort() int64`

GetReplicationPort returns the ReplicationPort field if non-nil, zero value otherwise.

### GetReplicationPortOk

`func (o *SyncServerInstanceResponse) GetReplicationPortOk() (*int64, bool)`

GetReplicationPortOk returns a tuple with the ReplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPort

`func (o *SyncServerInstanceResponse) SetReplicationPort(v int64)`

SetReplicationPort sets ReplicationPort field to given value.

### HasReplicationPort

`func (o *SyncServerInstanceResponse) HasReplicationPort() bool`

HasReplicationPort returns a boolean if a field has been set.

### GetReplicationServerID

`func (o *SyncServerInstanceResponse) GetReplicationServerID() int64`

GetReplicationServerID returns the ReplicationServerID field if non-nil, zero value otherwise.

### GetReplicationServerIDOk

`func (o *SyncServerInstanceResponse) GetReplicationServerIDOk() (*int64, bool)`

GetReplicationServerIDOk returns a tuple with the ReplicationServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationServerID

`func (o *SyncServerInstanceResponse) SetReplicationServerID(v int64)`

SetReplicationServerID sets ReplicationServerID field to given value.

### HasReplicationServerID

`func (o *SyncServerInstanceResponse) HasReplicationServerID() bool`

HasReplicationServerID returns a boolean if a field has been set.

### GetReplicationDomainServerID

`func (o *SyncServerInstanceResponse) GetReplicationDomainServerID() []int64`

GetReplicationDomainServerID returns the ReplicationDomainServerID field if non-nil, zero value otherwise.

### GetReplicationDomainServerIDOk

`func (o *SyncServerInstanceResponse) GetReplicationDomainServerIDOk() (*[]int64, bool)`

GetReplicationDomainServerIDOk returns a tuple with the ReplicationDomainServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationDomainServerID

`func (o *SyncServerInstanceResponse) SetReplicationDomainServerID(v []int64)`

SetReplicationDomainServerID sets ReplicationDomainServerID field to given value.

### HasReplicationDomainServerID

`func (o *SyncServerInstanceResponse) HasReplicationDomainServerID() bool`

HasReplicationDomainServerID returns a boolean if a field has been set.

### GetJmxPort

`func (o *SyncServerInstanceResponse) GetJmxPort() int64`

GetJmxPort returns the JmxPort field if non-nil, zero value otherwise.

### GetJmxPortOk

`func (o *SyncServerInstanceResponse) GetJmxPortOk() (*int64, bool)`

GetJmxPortOk returns a tuple with the JmxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxPort

`func (o *SyncServerInstanceResponse) SetJmxPort(v int64)`

SetJmxPort sets JmxPort field to given value.

### HasJmxPort

`func (o *SyncServerInstanceResponse) HasJmxPort() bool`

HasJmxPort returns a boolean if a field has been set.

### GetJmxsPort

`func (o *SyncServerInstanceResponse) GetJmxsPort() int64`

GetJmxsPort returns the JmxsPort field if non-nil, zero value otherwise.

### GetJmxsPortOk

`func (o *SyncServerInstanceResponse) GetJmxsPortOk() (*int64, bool)`

GetJmxsPortOk returns a tuple with the JmxsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxsPort

`func (o *SyncServerInstanceResponse) SetJmxsPort(v int64)`

SetJmxsPort sets JmxsPort field to given value.

### HasJmxsPort

`func (o *SyncServerInstanceResponse) HasJmxsPort() bool`

HasJmxsPort returns a boolean if a field has been set.

### GetPreferredSecurity

`func (o *SyncServerInstanceResponse) GetPreferredSecurity() EnumserverInstancePreferredSecurityProp`

GetPreferredSecurity returns the PreferredSecurity field if non-nil, zero value otherwise.

### GetPreferredSecurityOk

`func (o *SyncServerInstanceResponse) GetPreferredSecurityOk() (*EnumserverInstancePreferredSecurityProp, bool)`

GetPreferredSecurityOk returns a tuple with the PreferredSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSecurity

`func (o *SyncServerInstanceResponse) SetPreferredSecurity(v EnumserverInstancePreferredSecurityProp)`

SetPreferredSecurity sets PreferredSecurity field to given value.

### HasPreferredSecurity

`func (o *SyncServerInstanceResponse) HasPreferredSecurity() bool`

HasPreferredSecurity returns a boolean if a field has been set.

### GetStartTLSEnabled

`func (o *SyncServerInstanceResponse) GetStartTLSEnabled() bool`

GetStartTLSEnabled returns the StartTLSEnabled field if non-nil, zero value otherwise.

### GetStartTLSEnabledOk

`func (o *SyncServerInstanceResponse) GetStartTLSEnabledOk() (*bool, bool)`

GetStartTLSEnabledOk returns a tuple with the StartTLSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTLSEnabled

`func (o *SyncServerInstanceResponse) SetStartTLSEnabled(v bool)`

SetStartTLSEnabled sets StartTLSEnabled field to given value.

### HasStartTLSEnabled

`func (o *SyncServerInstanceResponse) HasStartTLSEnabled() bool`

HasStartTLSEnabled returns a boolean if a field has been set.

### GetBaseDN

`func (o *SyncServerInstanceResponse) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *SyncServerInstanceResponse) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *SyncServerInstanceResponse) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *SyncServerInstanceResponse) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMemberOfServerGroup

`func (o *SyncServerInstanceResponse) GetMemberOfServerGroup() []string`

GetMemberOfServerGroup returns the MemberOfServerGroup field if non-nil, zero value otherwise.

### GetMemberOfServerGroupOk

`func (o *SyncServerInstanceResponse) GetMemberOfServerGroupOk() (*[]string, bool)`

GetMemberOfServerGroupOk returns a tuple with the MemberOfServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfServerGroup

`func (o *SyncServerInstanceResponse) SetMemberOfServerGroup(v []string)`

SetMemberOfServerGroup sets MemberOfServerGroup field to given value.

### HasMemberOfServerGroup

`func (o *SyncServerInstanceResponse) HasMemberOfServerGroup() bool`

HasMemberOfServerGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


