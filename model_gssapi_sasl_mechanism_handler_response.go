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

// GssapiSaslMechanismHandlerResponse struct for GssapiSaslMechanismHandlerResponse
type GssapiSaslMechanismHandlerResponse struct {
	Meta *MetaMeta `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas []EnumgssapiSaslMechanismHandlerSchemaUrn `json:"schemas"`
	// Name of the SASL Mechanism Handler
	Id *string `json:"id,omitempty"`
	// Specifies the realm to be used for GSSAPI authentication.
	Realm *string `json:"realm,omitempty"`
	// Specifies the address of the KDC that is to be used for Kerberos processing.
	KdcAddress *string `json:"kdcAddress,omitempty"`
	// Specifies the keytab file that should be used for Kerberos processing.
	Keytab *string `json:"keytab,omitempty"`
	// Specifies whether or not to allow a null value for the server-fqdn.
	AllowNullServerFqdn *bool `json:"allowNullServerFqdn,omitempty"`
	// Specifies the DNS-resolvable fully-qualified domain name for the system.
	ServerFqdn *string `json:"serverFqdn,omitempty"`
	AllowedQualityOfProtection []EnumsaslMechanismHandlerAllowedQualityOfProtectionProp `json:"allowedQualityOfProtection,omitempty"`
	// Specifies the name of the identity mapper that is to be used with this SASL mechanism handler to match the Kerberos principal included in the SASL bind request to the corresponding user in the directory.
	IdentityMapper string `json:"identityMapper"`
	// Specifies the name of the identity mapper that is to be used with this SASL mechanism handler to map the alternate authorization identity (if provided, and if different from the Kerberos principal used as the authentication identity) to the corresponding user in the directory. If no value is specified, then the mapper specified in the identity-mapper configuration property will be used.
	AlternateAuthorizationIdentityMapper *string `json:"alternateAuthorizationIdentityMapper,omitempty"`
	// Specifies the Kerberos service principal that the Directory Server will use to identify itself to the KDC.
	KerberosServicePrincipal *string `json:"kerberosServicePrincipal,omitempty"`
	GssapiRole *EnumsaslMechanismHandlerGssapiRoleProp `json:"gssapiRole,omitempty"`
	// Specifies the path to a JAAS (Java Authentication and Authorization Service) configuration file that provides the information that the JVM should use for Kerberos processing.
	JaasConfigFile *string `json:"jaasConfigFile,omitempty"`
	// Indicates whether to enable debugging for the Java GSSAPI provider. Debug information will be written to standard output, which should be captured in the server.out log file.
	EnableDebug *bool `json:"enableDebug,omitempty"`
	// A description for this SASL Mechanism Handler
	Description *string `json:"description,omitempty"`
	// Indicates whether the SASL mechanism handler is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewGssapiSaslMechanismHandlerResponse instantiates a new GssapiSaslMechanismHandlerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGssapiSaslMechanismHandlerResponse(schemas []EnumgssapiSaslMechanismHandlerSchemaUrn, identityMapper string, enabled bool) *GssapiSaslMechanismHandlerResponse {
	this := GssapiSaslMechanismHandlerResponse{}
	this.Schemas = schemas
	this.IdentityMapper = identityMapper
	this.Enabled = enabled
	return &this
}

// NewGssapiSaslMechanismHandlerResponseWithDefaults instantiates a new GssapiSaslMechanismHandlerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGssapiSaslMechanismHandlerResponseWithDefaults() *GssapiSaslMechanismHandlerResponse {
	this := GssapiSaslMechanismHandlerResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *GssapiSaslMechanismHandlerResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || isNil(o.Urnpingidentityschemasconfigurationmessages20) {
    return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *GssapiSaslMechanismHandlerResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *GssapiSaslMechanismHandlerResponse) GetSchemas() []EnumgssapiSaslMechanismHandlerSchemaUrn {
	if o == nil {
		var ret []EnumgssapiSaslMechanismHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetSchemasOk() ([]EnumgssapiSaslMechanismHandlerSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *GssapiSaslMechanismHandlerResponse) SetSchemas(v []EnumgssapiSaslMechanismHandlerSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GssapiSaslMechanismHandlerResponse) SetId(v string) {
	o.Id = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetRealm() string {
	if o == nil || isNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetRealmOk() (*string, bool) {
	if o == nil || isNil(o.Realm) {
    return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasRealm() bool {
	if o != nil && !isNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *GssapiSaslMechanismHandlerResponse) SetRealm(v string) {
	o.Realm = &v
}

// GetKdcAddress returns the KdcAddress field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetKdcAddress() string {
	if o == nil || isNil(o.KdcAddress) {
		var ret string
		return ret
	}
	return *o.KdcAddress
}

// GetKdcAddressOk returns a tuple with the KdcAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetKdcAddressOk() (*string, bool) {
	if o == nil || isNil(o.KdcAddress) {
    return nil, false
	}
	return o.KdcAddress, true
}

// HasKdcAddress returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasKdcAddress() bool {
	if o != nil && !isNil(o.KdcAddress) {
		return true
	}

	return false
}

// SetKdcAddress gets a reference to the given string and assigns it to the KdcAddress field.
func (o *GssapiSaslMechanismHandlerResponse) SetKdcAddress(v string) {
	o.KdcAddress = &v
}

// GetKeytab returns the Keytab field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetKeytab() string {
	if o == nil || isNil(o.Keytab) {
		var ret string
		return ret
	}
	return *o.Keytab
}

// GetKeytabOk returns a tuple with the Keytab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetKeytabOk() (*string, bool) {
	if o == nil || isNil(o.Keytab) {
    return nil, false
	}
	return o.Keytab, true
}

// HasKeytab returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasKeytab() bool {
	if o != nil && !isNil(o.Keytab) {
		return true
	}

	return false
}

// SetKeytab gets a reference to the given string and assigns it to the Keytab field.
func (o *GssapiSaslMechanismHandlerResponse) SetKeytab(v string) {
	o.Keytab = &v
}

// GetAllowNullServerFqdn returns the AllowNullServerFqdn field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetAllowNullServerFqdn() bool {
	if o == nil || isNil(o.AllowNullServerFqdn) {
		var ret bool
		return ret
	}
	return *o.AllowNullServerFqdn
}

// GetAllowNullServerFqdnOk returns a tuple with the AllowNullServerFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetAllowNullServerFqdnOk() (*bool, bool) {
	if o == nil || isNil(o.AllowNullServerFqdn) {
    return nil, false
	}
	return o.AllowNullServerFqdn, true
}

// HasAllowNullServerFqdn returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasAllowNullServerFqdn() bool {
	if o != nil && !isNil(o.AllowNullServerFqdn) {
		return true
	}

	return false
}

// SetAllowNullServerFqdn gets a reference to the given bool and assigns it to the AllowNullServerFqdn field.
func (o *GssapiSaslMechanismHandlerResponse) SetAllowNullServerFqdn(v bool) {
	o.AllowNullServerFqdn = &v
}

// GetServerFqdn returns the ServerFqdn field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetServerFqdn() string {
	if o == nil || isNil(o.ServerFqdn) {
		var ret string
		return ret
	}
	return *o.ServerFqdn
}

// GetServerFqdnOk returns a tuple with the ServerFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetServerFqdnOk() (*string, bool) {
	if o == nil || isNil(o.ServerFqdn) {
    return nil, false
	}
	return o.ServerFqdn, true
}

// HasServerFqdn returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasServerFqdn() bool {
	if o != nil && !isNil(o.ServerFqdn) {
		return true
	}

	return false
}

// SetServerFqdn gets a reference to the given string and assigns it to the ServerFqdn field.
func (o *GssapiSaslMechanismHandlerResponse) SetServerFqdn(v string) {
	o.ServerFqdn = &v
}

// GetAllowedQualityOfProtection returns the AllowedQualityOfProtection field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetAllowedQualityOfProtection() []EnumsaslMechanismHandlerAllowedQualityOfProtectionProp {
	if o == nil || isNil(o.AllowedQualityOfProtection) {
		var ret []EnumsaslMechanismHandlerAllowedQualityOfProtectionProp
		return ret
	}
	return o.AllowedQualityOfProtection
}

// GetAllowedQualityOfProtectionOk returns a tuple with the AllowedQualityOfProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetAllowedQualityOfProtectionOk() ([]EnumsaslMechanismHandlerAllowedQualityOfProtectionProp, bool) {
	if o == nil || isNil(o.AllowedQualityOfProtection) {
    return nil, false
	}
	return o.AllowedQualityOfProtection, true
}

// HasAllowedQualityOfProtection returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasAllowedQualityOfProtection() bool {
	if o != nil && !isNil(o.AllowedQualityOfProtection) {
		return true
	}

	return false
}

// SetAllowedQualityOfProtection gets a reference to the given []EnumsaslMechanismHandlerAllowedQualityOfProtectionProp and assigns it to the AllowedQualityOfProtection field.
func (o *GssapiSaslMechanismHandlerResponse) SetAllowedQualityOfProtection(v []EnumsaslMechanismHandlerAllowedQualityOfProtectionProp) {
	o.AllowedQualityOfProtection = v
}

// GetIdentityMapper returns the IdentityMapper field value
func (o *GssapiSaslMechanismHandlerResponse) GetIdentityMapper() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityMapper
}

// GetIdentityMapperOk returns a tuple with the IdentityMapper field value
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetIdentityMapperOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IdentityMapper, true
}

// SetIdentityMapper sets field value
func (o *GssapiSaslMechanismHandlerResponse) SetIdentityMapper(v string) {
	o.IdentityMapper = v
}

// GetAlternateAuthorizationIdentityMapper returns the AlternateAuthorizationIdentityMapper field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetAlternateAuthorizationIdentityMapper() string {
	if o == nil || isNil(o.AlternateAuthorizationIdentityMapper) {
		var ret string
		return ret
	}
	return *o.AlternateAuthorizationIdentityMapper
}

// GetAlternateAuthorizationIdentityMapperOk returns a tuple with the AlternateAuthorizationIdentityMapper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetAlternateAuthorizationIdentityMapperOk() (*string, bool) {
	if o == nil || isNil(o.AlternateAuthorizationIdentityMapper) {
    return nil, false
	}
	return o.AlternateAuthorizationIdentityMapper, true
}

// HasAlternateAuthorizationIdentityMapper returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasAlternateAuthorizationIdentityMapper() bool {
	if o != nil && !isNil(o.AlternateAuthorizationIdentityMapper) {
		return true
	}

	return false
}

// SetAlternateAuthorizationIdentityMapper gets a reference to the given string and assigns it to the AlternateAuthorizationIdentityMapper field.
func (o *GssapiSaslMechanismHandlerResponse) SetAlternateAuthorizationIdentityMapper(v string) {
	o.AlternateAuthorizationIdentityMapper = &v
}

// GetKerberosServicePrincipal returns the KerberosServicePrincipal field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetKerberosServicePrincipal() string {
	if o == nil || isNil(o.KerberosServicePrincipal) {
		var ret string
		return ret
	}
	return *o.KerberosServicePrincipal
}

// GetKerberosServicePrincipalOk returns a tuple with the KerberosServicePrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetKerberosServicePrincipalOk() (*string, bool) {
	if o == nil || isNil(o.KerberosServicePrincipal) {
    return nil, false
	}
	return o.KerberosServicePrincipal, true
}

// HasKerberosServicePrincipal returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasKerberosServicePrincipal() bool {
	if o != nil && !isNil(o.KerberosServicePrincipal) {
		return true
	}

	return false
}

// SetKerberosServicePrincipal gets a reference to the given string and assigns it to the KerberosServicePrincipal field.
func (o *GssapiSaslMechanismHandlerResponse) SetKerberosServicePrincipal(v string) {
	o.KerberosServicePrincipal = &v
}

// GetGssapiRole returns the GssapiRole field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetGssapiRole() EnumsaslMechanismHandlerGssapiRoleProp {
	if o == nil || isNil(o.GssapiRole) {
		var ret EnumsaslMechanismHandlerGssapiRoleProp
		return ret
	}
	return *o.GssapiRole
}

// GetGssapiRoleOk returns a tuple with the GssapiRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetGssapiRoleOk() (*EnumsaslMechanismHandlerGssapiRoleProp, bool) {
	if o == nil || isNil(o.GssapiRole) {
    return nil, false
	}
	return o.GssapiRole, true
}

// HasGssapiRole returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasGssapiRole() bool {
	if o != nil && !isNil(o.GssapiRole) {
		return true
	}

	return false
}

// SetGssapiRole gets a reference to the given EnumsaslMechanismHandlerGssapiRoleProp and assigns it to the GssapiRole field.
func (o *GssapiSaslMechanismHandlerResponse) SetGssapiRole(v EnumsaslMechanismHandlerGssapiRoleProp) {
	o.GssapiRole = &v
}

// GetJaasConfigFile returns the JaasConfigFile field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetJaasConfigFile() string {
	if o == nil || isNil(o.JaasConfigFile) {
		var ret string
		return ret
	}
	return *o.JaasConfigFile
}

// GetJaasConfigFileOk returns a tuple with the JaasConfigFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetJaasConfigFileOk() (*string, bool) {
	if o == nil || isNil(o.JaasConfigFile) {
    return nil, false
	}
	return o.JaasConfigFile, true
}

// HasJaasConfigFile returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasJaasConfigFile() bool {
	if o != nil && !isNil(o.JaasConfigFile) {
		return true
	}

	return false
}

// SetJaasConfigFile gets a reference to the given string and assigns it to the JaasConfigFile field.
func (o *GssapiSaslMechanismHandlerResponse) SetJaasConfigFile(v string) {
	o.JaasConfigFile = &v
}

// GetEnableDebug returns the EnableDebug field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetEnableDebug() bool {
	if o == nil || isNil(o.EnableDebug) {
		var ret bool
		return ret
	}
	return *o.EnableDebug
}

// GetEnableDebugOk returns a tuple with the EnableDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetEnableDebugOk() (*bool, bool) {
	if o == nil || isNil(o.EnableDebug) {
    return nil, false
	}
	return o.EnableDebug, true
}

// HasEnableDebug returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasEnableDebug() bool {
	if o != nil && !isNil(o.EnableDebug) {
		return true
	}

	return false
}

// SetEnableDebug gets a reference to the given bool and assigns it to the EnableDebug field.
func (o *GssapiSaslMechanismHandlerResponse) SetEnableDebug(v bool) {
	o.EnableDebug = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GssapiSaslMechanismHandlerResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GssapiSaslMechanismHandlerResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GssapiSaslMechanismHandlerResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *GssapiSaslMechanismHandlerResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GssapiSaslMechanismHandlerResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GssapiSaslMechanismHandlerResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o GssapiSaslMechanismHandlerResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !isNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !isNil(o.KdcAddress) {
		toSerialize["kdcAddress"] = o.KdcAddress
	}
	if !isNil(o.Keytab) {
		toSerialize["keytab"] = o.Keytab
	}
	if !isNil(o.AllowNullServerFqdn) {
		toSerialize["allowNullServerFqdn"] = o.AllowNullServerFqdn
	}
	if !isNil(o.ServerFqdn) {
		toSerialize["serverFqdn"] = o.ServerFqdn
	}
	if !isNil(o.AllowedQualityOfProtection) {
		toSerialize["allowedQualityOfProtection"] = o.AllowedQualityOfProtection
	}
	if true {
		toSerialize["identityMapper"] = o.IdentityMapper
	}
	if !isNil(o.AlternateAuthorizationIdentityMapper) {
		toSerialize["alternateAuthorizationIdentityMapper"] = o.AlternateAuthorizationIdentityMapper
	}
	if !isNil(o.KerberosServicePrincipal) {
		toSerialize["kerberosServicePrincipal"] = o.KerberosServicePrincipal
	}
	if !isNil(o.GssapiRole) {
		toSerialize["gssapiRole"] = o.GssapiRole
	}
	if !isNil(o.JaasConfigFile) {
		toSerialize["jaasConfigFile"] = o.JaasConfigFile
	}
	if !isNil(o.EnableDebug) {
		toSerialize["enableDebug"] = o.EnableDebug
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableGssapiSaslMechanismHandlerResponse struct {
	value *GssapiSaslMechanismHandlerResponse
	isSet bool
}

func (v NullableGssapiSaslMechanismHandlerResponse) Get() *GssapiSaslMechanismHandlerResponse {
	return v.value
}

func (v *NullableGssapiSaslMechanismHandlerResponse) Set(val *GssapiSaslMechanismHandlerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGssapiSaslMechanismHandlerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGssapiSaslMechanismHandlerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGssapiSaslMechanismHandlerResponse(val *GssapiSaslMechanismHandlerResponse) *NullableGssapiSaslMechanismHandlerResponse {
	return &NullableGssapiSaslMechanismHandlerResponse{value: val, isSet: true}
}

func (v NullableGssapiSaslMechanismHandlerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGssapiSaslMechanismHandlerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

