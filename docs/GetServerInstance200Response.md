# GetServerInstance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsyncServerInstanceSchemaUrn**](EnumsyncServerInstanceSchemaUrn.md) |  | 
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
**ReplicationSetName** | Pointer to **string** | The name of the replication set assigned to this Directory Server. Restricted domains are only replicated within instances using the same replication set name. | [optional] 
**LoadBalancingAlgorithmName** | Pointer to **[]string** | The name of the configuration object for a load-balancing algorithm that should include this server. | [optional] 

## Methods

### NewGetServerInstance200Response

`func NewGetServerInstance200Response(schemas []EnumsyncServerInstanceSchemaUrn, serverInstanceName string, clusterName string, serverVersion string, ) *GetServerInstance200Response`

NewGetServerInstance200Response instantiates a new GetServerInstance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServerInstance200ResponseWithDefaults

`func NewGetServerInstance200ResponseWithDefaults() *GetServerInstance200Response`

NewGetServerInstance200ResponseWithDefaults instantiates a new GetServerInstance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetServerInstance200Response) GetSchemas() []EnumsyncServerInstanceSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetServerInstance200Response) GetSchemasOk() (*[]EnumsyncServerInstanceSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetServerInstance200Response) SetSchemas(v []EnumsyncServerInstanceSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetServerInstanceType

`func (o *GetServerInstance200Response) GetServerInstanceType() EnumserverInstanceServerInstanceTypeProp`

GetServerInstanceType returns the ServerInstanceType field if non-nil, zero value otherwise.

### GetServerInstanceTypeOk

`func (o *GetServerInstance200Response) GetServerInstanceTypeOk() (*EnumserverInstanceServerInstanceTypeProp, bool)`

GetServerInstanceTypeOk returns a tuple with the ServerInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceType

`func (o *GetServerInstance200Response) SetServerInstanceType(v EnumserverInstanceServerInstanceTypeProp)`

SetServerInstanceType sets ServerInstanceType field to given value.

### HasServerInstanceType

`func (o *GetServerInstance200Response) HasServerInstanceType() bool`

HasServerInstanceType returns a boolean if a field has been set.

### GetServerInstanceName

`func (o *GetServerInstance200Response) GetServerInstanceName() string`

GetServerInstanceName returns the ServerInstanceName field if non-nil, zero value otherwise.

### GetServerInstanceNameOk

`func (o *GetServerInstance200Response) GetServerInstanceNameOk() (*string, bool)`

GetServerInstanceNameOk returns a tuple with the ServerInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceName

`func (o *GetServerInstance200Response) SetServerInstanceName(v string)`

SetServerInstanceName sets ServerInstanceName field to given value.


### GetClusterName

`func (o *GetServerInstance200Response) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *GetServerInstance200Response) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *GetServerInstance200Response) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetServerInstanceLocation

`func (o *GetServerInstance200Response) GetServerInstanceLocation() string`

GetServerInstanceLocation returns the ServerInstanceLocation field if non-nil, zero value otherwise.

### GetServerInstanceLocationOk

`func (o *GetServerInstance200Response) GetServerInstanceLocationOk() (*string, bool)`

GetServerInstanceLocationOk returns a tuple with the ServerInstanceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInstanceLocation

`func (o *GetServerInstance200Response) SetServerInstanceLocation(v string)`

SetServerInstanceLocation sets ServerInstanceLocation field to given value.

### HasServerInstanceLocation

`func (o *GetServerInstance200Response) HasServerInstanceLocation() bool`

HasServerInstanceLocation returns a boolean if a field has been set.

### GetHostname

`func (o *GetServerInstance200Response) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetServerInstance200Response) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetServerInstance200Response) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetServerInstance200Response) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetServerRoot

`func (o *GetServerInstance200Response) GetServerRoot() string`

GetServerRoot returns the ServerRoot field if non-nil, zero value otherwise.

### GetServerRootOk

`func (o *GetServerInstance200Response) GetServerRootOk() (*string, bool)`

GetServerRootOk returns a tuple with the ServerRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRoot

`func (o *GetServerInstance200Response) SetServerRoot(v string)`

SetServerRoot sets ServerRoot field to given value.

### HasServerRoot

`func (o *GetServerInstance200Response) HasServerRoot() bool`

HasServerRoot returns a boolean if a field has been set.

### GetServerVersion

`func (o *GetServerInstance200Response) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *GetServerInstance200Response) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *GetServerInstance200Response) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetInterServerCertificate

`func (o *GetServerInstance200Response) GetInterServerCertificate() string`

GetInterServerCertificate returns the InterServerCertificate field if non-nil, zero value otherwise.

### GetInterServerCertificateOk

`func (o *GetServerInstance200Response) GetInterServerCertificateOk() (*string, bool)`

GetInterServerCertificateOk returns a tuple with the InterServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterServerCertificate

`func (o *GetServerInstance200Response) SetInterServerCertificate(v string)`

SetInterServerCertificate sets InterServerCertificate field to given value.

### HasInterServerCertificate

`func (o *GetServerInstance200Response) HasInterServerCertificate() bool`

HasInterServerCertificate returns a boolean if a field has been set.

### GetLdapPort

`func (o *GetServerInstance200Response) GetLdapPort() int32`

GetLdapPort returns the LdapPort field if non-nil, zero value otherwise.

### GetLdapPortOk

`func (o *GetServerInstance200Response) GetLdapPortOk() (*int32, bool)`

GetLdapPortOk returns a tuple with the LdapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPort

`func (o *GetServerInstance200Response) SetLdapPort(v int32)`

SetLdapPort sets LdapPort field to given value.

### HasLdapPort

`func (o *GetServerInstance200Response) HasLdapPort() bool`

HasLdapPort returns a boolean if a field has been set.

### GetLdapsPort

`func (o *GetServerInstance200Response) GetLdapsPort() int32`

GetLdapsPort returns the LdapsPort field if non-nil, zero value otherwise.

### GetLdapsPortOk

`func (o *GetServerInstance200Response) GetLdapsPortOk() (*int32, bool)`

GetLdapsPortOk returns a tuple with the LdapsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapsPort

`func (o *GetServerInstance200Response) SetLdapsPort(v int32)`

SetLdapsPort sets LdapsPort field to given value.

### HasLdapsPort

`func (o *GetServerInstance200Response) HasLdapsPort() bool`

HasLdapsPort returns a boolean if a field has been set.

### GetHttpPort

`func (o *GetServerInstance200Response) GetHttpPort() int32`

GetHttpPort returns the HttpPort field if non-nil, zero value otherwise.

### GetHttpPortOk

`func (o *GetServerInstance200Response) GetHttpPortOk() (*int32, bool)`

GetHttpPortOk returns a tuple with the HttpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPort

`func (o *GetServerInstance200Response) SetHttpPort(v int32)`

SetHttpPort sets HttpPort field to given value.

### HasHttpPort

`func (o *GetServerInstance200Response) HasHttpPort() bool`

HasHttpPort returns a boolean if a field has been set.

### GetHttpsPort

`func (o *GetServerInstance200Response) GetHttpsPort() int32`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *GetServerInstance200Response) GetHttpsPortOk() (*int32, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *GetServerInstance200Response) SetHttpsPort(v int32)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *GetServerInstance200Response) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetReplicationPort

`func (o *GetServerInstance200Response) GetReplicationPort() int32`

GetReplicationPort returns the ReplicationPort field if non-nil, zero value otherwise.

### GetReplicationPortOk

`func (o *GetServerInstance200Response) GetReplicationPortOk() (*int32, bool)`

GetReplicationPortOk returns a tuple with the ReplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPort

`func (o *GetServerInstance200Response) SetReplicationPort(v int32)`

SetReplicationPort sets ReplicationPort field to given value.

### HasReplicationPort

`func (o *GetServerInstance200Response) HasReplicationPort() bool`

HasReplicationPort returns a boolean if a field has been set.

### GetReplicationServerID

`func (o *GetServerInstance200Response) GetReplicationServerID() int32`

GetReplicationServerID returns the ReplicationServerID field if non-nil, zero value otherwise.

### GetReplicationServerIDOk

`func (o *GetServerInstance200Response) GetReplicationServerIDOk() (*int32, bool)`

GetReplicationServerIDOk returns a tuple with the ReplicationServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationServerID

`func (o *GetServerInstance200Response) SetReplicationServerID(v int32)`

SetReplicationServerID sets ReplicationServerID field to given value.

### HasReplicationServerID

`func (o *GetServerInstance200Response) HasReplicationServerID() bool`

HasReplicationServerID returns a boolean if a field has been set.

### GetReplicationDomainServerID

`func (o *GetServerInstance200Response) GetReplicationDomainServerID() []int32`

GetReplicationDomainServerID returns the ReplicationDomainServerID field if non-nil, zero value otherwise.

### GetReplicationDomainServerIDOk

`func (o *GetServerInstance200Response) GetReplicationDomainServerIDOk() (*[]int32, bool)`

GetReplicationDomainServerIDOk returns a tuple with the ReplicationDomainServerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationDomainServerID

`func (o *GetServerInstance200Response) SetReplicationDomainServerID(v []int32)`

SetReplicationDomainServerID sets ReplicationDomainServerID field to given value.

### HasReplicationDomainServerID

`func (o *GetServerInstance200Response) HasReplicationDomainServerID() bool`

HasReplicationDomainServerID returns a boolean if a field has been set.

### GetJmxPort

`func (o *GetServerInstance200Response) GetJmxPort() int32`

GetJmxPort returns the JmxPort field if non-nil, zero value otherwise.

### GetJmxPortOk

`func (o *GetServerInstance200Response) GetJmxPortOk() (*int32, bool)`

GetJmxPortOk returns a tuple with the JmxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxPort

`func (o *GetServerInstance200Response) SetJmxPort(v int32)`

SetJmxPort sets JmxPort field to given value.

### HasJmxPort

`func (o *GetServerInstance200Response) HasJmxPort() bool`

HasJmxPort returns a boolean if a field has been set.

### GetJmxsPort

`func (o *GetServerInstance200Response) GetJmxsPort() int32`

GetJmxsPort returns the JmxsPort field if non-nil, zero value otherwise.

### GetJmxsPortOk

`func (o *GetServerInstance200Response) GetJmxsPortOk() (*int32, bool)`

GetJmxsPortOk returns a tuple with the JmxsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJmxsPort

`func (o *GetServerInstance200Response) SetJmxsPort(v int32)`

SetJmxsPort sets JmxsPort field to given value.

### HasJmxsPort

`func (o *GetServerInstance200Response) HasJmxsPort() bool`

HasJmxsPort returns a boolean if a field has been set.

### GetPreferredSecurity

`func (o *GetServerInstance200Response) GetPreferredSecurity() EnumserverInstancePreferredSecurityProp`

GetPreferredSecurity returns the PreferredSecurity field if non-nil, zero value otherwise.

### GetPreferredSecurityOk

`func (o *GetServerInstance200Response) GetPreferredSecurityOk() (*EnumserverInstancePreferredSecurityProp, bool)`

GetPreferredSecurityOk returns a tuple with the PreferredSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredSecurity

`func (o *GetServerInstance200Response) SetPreferredSecurity(v EnumserverInstancePreferredSecurityProp)`

SetPreferredSecurity sets PreferredSecurity field to given value.

### HasPreferredSecurity

`func (o *GetServerInstance200Response) HasPreferredSecurity() bool`

HasPreferredSecurity returns a boolean if a field has been set.

### GetStartTLSEnabled

`func (o *GetServerInstance200Response) GetStartTLSEnabled() bool`

GetStartTLSEnabled returns the StartTLSEnabled field if non-nil, zero value otherwise.

### GetStartTLSEnabledOk

`func (o *GetServerInstance200Response) GetStartTLSEnabledOk() (*bool, bool)`

GetStartTLSEnabledOk returns a tuple with the StartTLSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTLSEnabled

`func (o *GetServerInstance200Response) SetStartTLSEnabled(v bool)`

SetStartTLSEnabled sets StartTLSEnabled field to given value.

### HasStartTLSEnabled

`func (o *GetServerInstance200Response) HasStartTLSEnabled() bool`

HasStartTLSEnabled returns a boolean if a field has been set.

### GetBaseDN

`func (o *GetServerInstance200Response) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *GetServerInstance200Response) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *GetServerInstance200Response) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.

### HasBaseDN

`func (o *GetServerInstance200Response) HasBaseDN() bool`

HasBaseDN returns a boolean if a field has been set.

### GetMemberOfServerGroup

`func (o *GetServerInstance200Response) GetMemberOfServerGroup() []string`

GetMemberOfServerGroup returns the MemberOfServerGroup field if non-nil, zero value otherwise.

### GetMemberOfServerGroupOk

`func (o *GetServerInstance200Response) GetMemberOfServerGroupOk() (*[]string, bool)`

GetMemberOfServerGroupOk returns a tuple with the MemberOfServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOfServerGroup

`func (o *GetServerInstance200Response) SetMemberOfServerGroup(v []string)`

SetMemberOfServerGroup sets MemberOfServerGroup field to given value.

### HasMemberOfServerGroup

`func (o *GetServerInstance200Response) HasMemberOfServerGroup() bool`

HasMemberOfServerGroup returns a boolean if a field has been set.

### GetReplicationSetName

`func (o *GetServerInstance200Response) GetReplicationSetName() string`

GetReplicationSetName returns the ReplicationSetName field if non-nil, zero value otherwise.

### GetReplicationSetNameOk

`func (o *GetServerInstance200Response) GetReplicationSetNameOk() (*string, bool)`

GetReplicationSetNameOk returns a tuple with the ReplicationSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationSetName

`func (o *GetServerInstance200Response) SetReplicationSetName(v string)`

SetReplicationSetName sets ReplicationSetName field to given value.

### HasReplicationSetName

`func (o *GetServerInstance200Response) HasReplicationSetName() bool`

HasReplicationSetName returns a boolean if a field has been set.

### GetLoadBalancingAlgorithmName

`func (o *GetServerInstance200Response) GetLoadBalancingAlgorithmName() []string`

GetLoadBalancingAlgorithmName returns the LoadBalancingAlgorithmName field if non-nil, zero value otherwise.

### GetLoadBalancingAlgorithmNameOk

`func (o *GetServerInstance200Response) GetLoadBalancingAlgorithmNameOk() (*[]string, bool)`

GetLoadBalancingAlgorithmNameOk returns a tuple with the LoadBalancingAlgorithmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancingAlgorithmName

`func (o *GetServerInstance200Response) SetLoadBalancingAlgorithmName(v []string)`

SetLoadBalancingAlgorithmName sets LoadBalancingAlgorithmName field to given value.

### HasLoadBalancingAlgorithmName

`func (o *GetServerInstance200Response) HasLoadBalancingAlgorithmName() bool`

HasLoadBalancingAlgorithmName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


