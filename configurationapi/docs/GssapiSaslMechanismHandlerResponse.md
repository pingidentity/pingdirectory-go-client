# GssapiSaslMechanismHandlerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumgssapiSaslMechanismHandlerSchemaUrn**](EnumgssapiSaslMechanismHandlerSchemaUrn.md) |  | 
**Id** | **string** | Name of the SASL Mechanism Handler | 
**Realm** | Pointer to **string** | Specifies the realm to be used for GSSAPI authentication. | [optional] 
**KdcAddress** | Pointer to **string** | Specifies the address of the KDC that is to be used for Kerberos processing. | [optional] 
**Keytab** | Pointer to **string** | Specifies the keytab file that should be used for Kerberos processing. | [optional] 
**AllowNullServerFqdn** | Pointer to **bool** | Specifies whether or not to allow a null value for the server-fqdn. | [optional] 
**ServerFqdn** | Pointer to **string** | Specifies the DNS-resolvable fully-qualified domain name for the system. | [optional] 
**AllowedQualityOfProtection** | Pointer to [**[]EnumsaslMechanismHandlerAllowedQualityOfProtectionProp**](EnumsaslMechanismHandlerAllowedQualityOfProtectionProp.md) | Specifies the supported quality of protection (QoP) levels that clients will be permitted to request when performing GSSAPI authentication. | [optional] 
**IdentityMapper** | **string** | Specifies the name of the identity mapper that is to be used with this SASL mechanism handler to match the Kerberos principal included in the SASL bind request to the corresponding user in the directory. | 
**AlternateAuthorizationIdentityMapper** | Pointer to **string** | Specifies the name of the identity mapper that is to be used with this SASL mechanism handler to map the alternate authorization identity (if provided, and if different from the Kerberos principal used as the authentication identity) to the corresponding user in the directory. If no value is specified, then the mapper specified in the identity-mapper configuration property will be used. | [optional] 
**KerberosServicePrincipal** | Pointer to **string** | Specifies the Kerberos service principal that the Directory Server will use to identify itself to the KDC. | [optional] 
**GssapiRole** | Pointer to [**EnumsaslMechanismHandlerGssapiRoleProp**](EnumsaslMechanismHandlerGssapiRoleProp.md) |  | [optional] 
**JaasConfigFile** | Pointer to **string** | Specifies the path to a JAAS (Java Authentication and Authorization Service) configuration file that provides the information that the JVM should use for Kerberos processing. | [optional] 
**EnableDebug** | Pointer to **bool** | Indicates whether to enable debugging for the Java GSSAPI provider. Debug information will be written to standard output, which should be captured in the server.out log file. | [optional] 
**Description** | Pointer to **string** | A description for this SASL Mechanism Handler | [optional] 
**Enabled** | **bool** | Indicates whether the SASL mechanism handler is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewGssapiSaslMechanismHandlerResponse

`func NewGssapiSaslMechanismHandlerResponse(schemas []EnumgssapiSaslMechanismHandlerSchemaUrn, id string, identityMapper string, enabled bool, ) *GssapiSaslMechanismHandlerResponse`

NewGssapiSaslMechanismHandlerResponse instantiates a new GssapiSaslMechanismHandlerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGssapiSaslMechanismHandlerResponseWithDefaults

`func NewGssapiSaslMechanismHandlerResponseWithDefaults() *GssapiSaslMechanismHandlerResponse`

NewGssapiSaslMechanismHandlerResponseWithDefaults instantiates a new GssapiSaslMechanismHandlerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GssapiSaslMechanismHandlerResponse) GetSchemas() []EnumgssapiSaslMechanismHandlerSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GssapiSaslMechanismHandlerResponse) GetSchemasOk() (*[]EnumgssapiSaslMechanismHandlerSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GssapiSaslMechanismHandlerResponse) SetSchemas(v []EnumgssapiSaslMechanismHandlerSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GssapiSaslMechanismHandlerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GssapiSaslMechanismHandlerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GssapiSaslMechanismHandlerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetRealm

`func (o *GssapiSaslMechanismHandlerResponse) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *GssapiSaslMechanismHandlerResponse) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *GssapiSaslMechanismHandlerResponse) SetRealm(v string)`

SetRealm sets Realm field to given value.

### HasRealm

`func (o *GssapiSaslMechanismHandlerResponse) HasRealm() bool`

HasRealm returns a boolean if a field has been set.

### GetKdcAddress

`func (o *GssapiSaslMechanismHandlerResponse) GetKdcAddress() string`

GetKdcAddress returns the KdcAddress field if non-nil, zero value otherwise.

### GetKdcAddressOk

`func (o *GssapiSaslMechanismHandlerResponse) GetKdcAddressOk() (*string, bool)`

GetKdcAddressOk returns a tuple with the KdcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKdcAddress

`func (o *GssapiSaslMechanismHandlerResponse) SetKdcAddress(v string)`

SetKdcAddress sets KdcAddress field to given value.

### HasKdcAddress

`func (o *GssapiSaslMechanismHandlerResponse) HasKdcAddress() bool`

HasKdcAddress returns a boolean if a field has been set.

### GetKeytab

`func (o *GssapiSaslMechanismHandlerResponse) GetKeytab() string`

GetKeytab returns the Keytab field if non-nil, zero value otherwise.

### GetKeytabOk

`func (o *GssapiSaslMechanismHandlerResponse) GetKeytabOk() (*string, bool)`

GetKeytabOk returns a tuple with the Keytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytab

`func (o *GssapiSaslMechanismHandlerResponse) SetKeytab(v string)`

SetKeytab sets Keytab field to given value.

### HasKeytab

`func (o *GssapiSaslMechanismHandlerResponse) HasKeytab() bool`

HasKeytab returns a boolean if a field has been set.

### GetAllowNullServerFqdn

`func (o *GssapiSaslMechanismHandlerResponse) GetAllowNullServerFqdn() bool`

GetAllowNullServerFqdn returns the AllowNullServerFqdn field if non-nil, zero value otherwise.

### GetAllowNullServerFqdnOk

`func (o *GssapiSaslMechanismHandlerResponse) GetAllowNullServerFqdnOk() (*bool, bool)`

GetAllowNullServerFqdnOk returns a tuple with the AllowNullServerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNullServerFqdn

`func (o *GssapiSaslMechanismHandlerResponse) SetAllowNullServerFqdn(v bool)`

SetAllowNullServerFqdn sets AllowNullServerFqdn field to given value.

### HasAllowNullServerFqdn

`func (o *GssapiSaslMechanismHandlerResponse) HasAllowNullServerFqdn() bool`

HasAllowNullServerFqdn returns a boolean if a field has been set.

### GetServerFqdn

`func (o *GssapiSaslMechanismHandlerResponse) GetServerFqdn() string`

GetServerFqdn returns the ServerFqdn field if non-nil, zero value otherwise.

### GetServerFqdnOk

`func (o *GssapiSaslMechanismHandlerResponse) GetServerFqdnOk() (*string, bool)`

GetServerFqdnOk returns a tuple with the ServerFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFqdn

`func (o *GssapiSaslMechanismHandlerResponse) SetServerFqdn(v string)`

SetServerFqdn sets ServerFqdn field to given value.

### HasServerFqdn

`func (o *GssapiSaslMechanismHandlerResponse) HasServerFqdn() bool`

HasServerFqdn returns a boolean if a field has been set.

### GetAllowedQualityOfProtection

`func (o *GssapiSaslMechanismHandlerResponse) GetAllowedQualityOfProtection() []EnumsaslMechanismHandlerAllowedQualityOfProtectionProp`

GetAllowedQualityOfProtection returns the AllowedQualityOfProtection field if non-nil, zero value otherwise.

### GetAllowedQualityOfProtectionOk

`func (o *GssapiSaslMechanismHandlerResponse) GetAllowedQualityOfProtectionOk() (*[]EnumsaslMechanismHandlerAllowedQualityOfProtectionProp, bool)`

GetAllowedQualityOfProtectionOk returns a tuple with the AllowedQualityOfProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedQualityOfProtection

`func (o *GssapiSaslMechanismHandlerResponse) SetAllowedQualityOfProtection(v []EnumsaslMechanismHandlerAllowedQualityOfProtectionProp)`

SetAllowedQualityOfProtection sets AllowedQualityOfProtection field to given value.

### HasAllowedQualityOfProtection

`func (o *GssapiSaslMechanismHandlerResponse) HasAllowedQualityOfProtection() bool`

HasAllowedQualityOfProtection returns a boolean if a field has been set.

### GetIdentityMapper

`func (o *GssapiSaslMechanismHandlerResponse) GetIdentityMapper() string`

GetIdentityMapper returns the IdentityMapper field if non-nil, zero value otherwise.

### GetIdentityMapperOk

`func (o *GssapiSaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool)`

GetIdentityMapperOk returns a tuple with the IdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityMapper

`func (o *GssapiSaslMechanismHandlerResponse) SetIdentityMapper(v string)`

SetIdentityMapper sets IdentityMapper field to given value.


### GetAlternateAuthorizationIdentityMapper

`func (o *GssapiSaslMechanismHandlerResponse) GetAlternateAuthorizationIdentityMapper() string`

GetAlternateAuthorizationIdentityMapper returns the AlternateAuthorizationIdentityMapper field if non-nil, zero value otherwise.

### GetAlternateAuthorizationIdentityMapperOk

`func (o *GssapiSaslMechanismHandlerResponse) GetAlternateAuthorizationIdentityMapperOk() (*string, bool)`

GetAlternateAuthorizationIdentityMapperOk returns a tuple with the AlternateAuthorizationIdentityMapper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateAuthorizationIdentityMapper

`func (o *GssapiSaslMechanismHandlerResponse) SetAlternateAuthorizationIdentityMapper(v string)`

SetAlternateAuthorizationIdentityMapper sets AlternateAuthorizationIdentityMapper field to given value.

### HasAlternateAuthorizationIdentityMapper

`func (o *GssapiSaslMechanismHandlerResponse) HasAlternateAuthorizationIdentityMapper() bool`

HasAlternateAuthorizationIdentityMapper returns a boolean if a field has been set.

### GetKerberosServicePrincipal

`func (o *GssapiSaslMechanismHandlerResponse) GetKerberosServicePrincipal() string`

GetKerberosServicePrincipal returns the KerberosServicePrincipal field if non-nil, zero value otherwise.

### GetKerberosServicePrincipalOk

`func (o *GssapiSaslMechanismHandlerResponse) GetKerberosServicePrincipalOk() (*string, bool)`

GetKerberosServicePrincipalOk returns a tuple with the KerberosServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosServicePrincipal

`func (o *GssapiSaslMechanismHandlerResponse) SetKerberosServicePrincipal(v string)`

SetKerberosServicePrincipal sets KerberosServicePrincipal field to given value.

### HasKerberosServicePrincipal

`func (o *GssapiSaslMechanismHandlerResponse) HasKerberosServicePrincipal() bool`

HasKerberosServicePrincipal returns a boolean if a field has been set.

### GetGssapiRole

`func (o *GssapiSaslMechanismHandlerResponse) GetGssapiRole() EnumsaslMechanismHandlerGssapiRoleProp`

GetGssapiRole returns the GssapiRole field if non-nil, zero value otherwise.

### GetGssapiRoleOk

`func (o *GssapiSaslMechanismHandlerResponse) GetGssapiRoleOk() (*EnumsaslMechanismHandlerGssapiRoleProp, bool)`

GetGssapiRoleOk returns a tuple with the GssapiRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssapiRole

`func (o *GssapiSaslMechanismHandlerResponse) SetGssapiRole(v EnumsaslMechanismHandlerGssapiRoleProp)`

SetGssapiRole sets GssapiRole field to given value.

### HasGssapiRole

`func (o *GssapiSaslMechanismHandlerResponse) HasGssapiRole() bool`

HasGssapiRole returns a boolean if a field has been set.

### GetJaasConfigFile

`func (o *GssapiSaslMechanismHandlerResponse) GetJaasConfigFile() string`

GetJaasConfigFile returns the JaasConfigFile field if non-nil, zero value otherwise.

### GetJaasConfigFileOk

`func (o *GssapiSaslMechanismHandlerResponse) GetJaasConfigFileOk() (*string, bool)`

GetJaasConfigFileOk returns a tuple with the JaasConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJaasConfigFile

`func (o *GssapiSaslMechanismHandlerResponse) SetJaasConfigFile(v string)`

SetJaasConfigFile sets JaasConfigFile field to given value.

### HasJaasConfigFile

`func (o *GssapiSaslMechanismHandlerResponse) HasJaasConfigFile() bool`

HasJaasConfigFile returns a boolean if a field has been set.

### GetEnableDebug

`func (o *GssapiSaslMechanismHandlerResponse) GetEnableDebug() bool`

GetEnableDebug returns the EnableDebug field if non-nil, zero value otherwise.

### GetEnableDebugOk

`func (o *GssapiSaslMechanismHandlerResponse) GetEnableDebugOk() (*bool, bool)`

GetEnableDebugOk returns a tuple with the EnableDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebug

`func (o *GssapiSaslMechanismHandlerResponse) SetEnableDebug(v bool)`

SetEnableDebug sets EnableDebug field to given value.

### HasEnableDebug

`func (o *GssapiSaslMechanismHandlerResponse) HasEnableDebug() bool`

HasEnableDebug returns a boolean if a field has been set.

### GetDescription

`func (o *GssapiSaslMechanismHandlerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GssapiSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GssapiSaslMechanismHandlerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GssapiSaslMechanismHandlerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GssapiSaslMechanismHandlerResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GssapiSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GssapiSaslMechanismHandlerResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *GssapiSaslMechanismHandlerResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GssapiSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GssapiSaslMechanismHandlerResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GssapiSaslMechanismHandlerResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GssapiSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GssapiSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GssapiSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GssapiSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


