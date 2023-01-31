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

// AddLdapPassThroughAuthenticationHandlerRequest struct for AddLdapPassThroughAuthenticationHandlerRequest
type AddLdapPassThroughAuthenticationHandlerRequest struct {
	// Name of the new Pass Through Authentication Handler
	HandlerName string                                              `json:"handlerName"`
	Schemas     []EnumldapPassThroughAuthenticationHandlerSchemaUrn `json:"schemas"`
	// Specifies the LDAP external server(s) to which authentication attempts should be forwarded.
	Server           []string                                                 `json:"server"`
	ServerAccessMode EnumpassThroughAuthenticationHandlerServerAccessModeProp `json:"serverAccessMode"`
	// Specifies one or more DN mappings that may be used to transform bind DNs before attempting to bind to the external servers.
	DnMap []string `json:"dnMap,omitempty"`
	// A pattern to use to construct the bind DN for the simple bind request to send to the remote server. This may consist of a combination of static text and attribute values and other directives enclosed in curly braces.  For example, the value \"cn={cn},ou=People,dc=example,dc=com\" indicates that the remote bind DN should be constructed from the text \"cn=\" followed by the value of the local entry's cn attribute followed by the text \"ou=People,dc=example,dc=com\". If an attribute contains the value to use as the bind DN for pass-through authentication, then the pattern may simply be the name of that attribute in curly braces (e.g., if the seeAlso attribute contains the bind DN for the target user, then a bind DN pattern of \"{seeAlso}\" would be appropriate).  Note that a bind DN pattern can be used to construct a bind DN that is not actually a valid LDAP distinguished name. For example, if authentication is being passed through to a Microsoft Active Directory server, then a bind DN pattern could be used to construct a user principal name (UPN) as an alternative to a distinguished name.
	BindDNPattern *string `json:"bindDNPattern,omitempty"`
	// The base DN to use when searching for the user entry using a filter constructed from the pattern defined in the search-filter-pattern property. If no base DN is specified, the null DN will be used as the search base DN.
	SearchBaseDN *string `json:"searchBaseDN,omitempty"`
	// A pattern to use to construct a filter to use when searching an external server for the entry of the user as whom to bind. For example, \"(mail={uid:ldapFilterEscape}@example.com)\" would construct a search filter to search for a user whose entry in the local server contains a uid attribute whose value appears before \"@example.com\" in the mail attribute in the external server. Note that the \"ldapFilterEscape\" modifier should almost always be used with attributes specified in the pattern.
	SearchFilterPattern *string `json:"searchFilterPattern,omitempty"`
	// Specifies the initial number of connections to establish to each external server against which authentication may be attempted.
	InitialConnections int32 `json:"initialConnections"`
	// Specifies the maximum number of connections to maintain to each external server against which authentication may be attempted. This value must be greater than or equal to the value for the initial-connections property.
	MaxConnections int32 `json:"maxConnections"`
	// Indicates whether to take server locations into account when prioritizing the servers to use for pass-through authentication attempts.
	UseLocation *bool `json:"useLocation,omitempty"`
	// The maximum length of time to wait for a response from an external server in the same location as this Directory Server before considering it unavailable.
	MaximumAllowedLocalResponseTime *string `json:"maximumAllowedLocalResponseTime,omitempty"`
	// The maximum length of time to wait for a response from an external server in a different location from this Directory Server before considering it unavailable.
	MaximumAllowedNonlocalResponseTime *string `json:"maximumAllowedNonlocalResponseTime,omitempty"`
	// Indicates whether to include the password policy request control (as defined in draft-behera-ldap-password-policy-10) in bind requests sent to the external server.
	UsePasswordPolicyControl *bool `json:"usePasswordPolicyControl,omitempty"`
	// A description for this Pass Through Authentication Handler
	Description *string `json:"description,omitempty"`
}

// NewAddLdapPassThroughAuthenticationHandlerRequest instantiates a new AddLdapPassThroughAuthenticationHandlerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLdapPassThroughAuthenticationHandlerRequest(handlerName string, schemas []EnumldapPassThroughAuthenticationHandlerSchemaUrn, server []string, serverAccessMode EnumpassThroughAuthenticationHandlerServerAccessModeProp, initialConnections int32, maxConnections int32) *AddLdapPassThroughAuthenticationHandlerRequest {
	this := AddLdapPassThroughAuthenticationHandlerRequest{}
	this.HandlerName = handlerName
	this.Schemas = schemas
	this.Server = server
	this.ServerAccessMode = serverAccessMode
	this.InitialConnections = initialConnections
	this.MaxConnections = maxConnections
	return &this
}

// NewAddLdapPassThroughAuthenticationHandlerRequestWithDefaults instantiates a new AddLdapPassThroughAuthenticationHandlerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLdapPassThroughAuthenticationHandlerRequestWithDefaults() *AddLdapPassThroughAuthenticationHandlerRequest {
	this := AddLdapPassThroughAuthenticationHandlerRequest{}
	return &this
}

// GetHandlerName returns the HandlerName field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetHandlerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HandlerName
}

// GetHandlerNameOk returns a tuple with the HandlerName field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetHandlerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandlerName, true
}

// SetHandlerName sets field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetHandlerName(v string) {
	o.HandlerName = v
}

// GetSchemas returns the Schemas field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSchemas() []EnumldapPassThroughAuthenticationHandlerSchemaUrn {
	if o == nil {
		var ret []EnumldapPassThroughAuthenticationHandlerSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSchemasOk() ([]EnumldapPassThroughAuthenticationHandlerSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetSchemas(v []EnumldapPassThroughAuthenticationHandlerSchemaUrn) {
	o.Schemas = v
}

// GetServer returns the Server field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetServer() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetServerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server, true
}

// SetServer sets field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetServer(v []string) {
	o.Server = v
}

// GetServerAccessMode returns the ServerAccessMode field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetServerAccessMode() EnumpassThroughAuthenticationHandlerServerAccessModeProp {
	if o == nil {
		var ret EnumpassThroughAuthenticationHandlerServerAccessModeProp
		return ret
	}

	return o.ServerAccessMode
}

// GetServerAccessModeOk returns a tuple with the ServerAccessMode field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetServerAccessModeOk() (*EnumpassThroughAuthenticationHandlerServerAccessModeProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerAccessMode, true
}

// SetServerAccessMode sets field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetServerAccessMode(v EnumpassThroughAuthenticationHandlerServerAccessModeProp) {
	o.ServerAccessMode = v
}

// GetDnMap returns the DnMap field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetDnMap() []string {
	if o == nil || isNil(o.DnMap) {
		var ret []string
		return ret
	}
	return o.DnMap
}

// GetDnMapOk returns a tuple with the DnMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetDnMapOk() ([]string, bool) {
	if o == nil || isNil(o.DnMap) {
		return nil, false
	}
	return o.DnMap, true
}

// HasDnMap returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasDnMap() bool {
	if o != nil && !isNil(o.DnMap) {
		return true
	}

	return false
}

// SetDnMap gets a reference to the given []string and assigns it to the DnMap field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetDnMap(v []string) {
	o.DnMap = v
}

// GetBindDNPattern returns the BindDNPattern field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetBindDNPattern() string {
	if o == nil || isNil(o.BindDNPattern) {
		var ret string
		return ret
	}
	return *o.BindDNPattern
}

// GetBindDNPatternOk returns a tuple with the BindDNPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetBindDNPatternOk() (*string, bool) {
	if o == nil || isNil(o.BindDNPattern) {
		return nil, false
	}
	return o.BindDNPattern, true
}

// HasBindDNPattern returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasBindDNPattern() bool {
	if o != nil && !isNil(o.BindDNPattern) {
		return true
	}

	return false
}

// SetBindDNPattern gets a reference to the given string and assigns it to the BindDNPattern field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetBindDNPattern(v string) {
	o.BindDNPattern = &v
}

// GetSearchBaseDN returns the SearchBaseDN field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSearchBaseDN() string {
	if o == nil || isNil(o.SearchBaseDN) {
		var ret string
		return ret
	}
	return *o.SearchBaseDN
}

// GetSearchBaseDNOk returns a tuple with the SearchBaseDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSearchBaseDNOk() (*string, bool) {
	if o == nil || isNil(o.SearchBaseDN) {
		return nil, false
	}
	return o.SearchBaseDN, true
}

// HasSearchBaseDN returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasSearchBaseDN() bool {
	if o != nil && !isNil(o.SearchBaseDN) {
		return true
	}

	return false
}

// SetSearchBaseDN gets a reference to the given string and assigns it to the SearchBaseDN field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetSearchBaseDN(v string) {
	o.SearchBaseDN = &v
}

// GetSearchFilterPattern returns the SearchFilterPattern field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSearchFilterPattern() string {
	if o == nil || isNil(o.SearchFilterPattern) {
		var ret string
		return ret
	}
	return *o.SearchFilterPattern
}

// GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetSearchFilterPatternOk() (*string, bool) {
	if o == nil || isNil(o.SearchFilterPattern) {
		return nil, false
	}
	return o.SearchFilterPattern, true
}

// HasSearchFilterPattern returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasSearchFilterPattern() bool {
	if o != nil && !isNil(o.SearchFilterPattern) {
		return true
	}

	return false
}

// SetSearchFilterPattern gets a reference to the given string and assigns it to the SearchFilterPattern field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetSearchFilterPattern(v string) {
	o.SearchFilterPattern = &v
}

// GetInitialConnections returns the InitialConnections field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetInitialConnections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InitialConnections
}

// GetInitialConnectionsOk returns a tuple with the InitialConnections field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetInitialConnectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialConnections, true
}

// SetInitialConnections sets field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetInitialConnections(v int32) {
	o.InitialConnections = v
}

// GetMaxConnections returns the MaxConnections field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaxConnections() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxConnections
}

// GetMaxConnectionsOk returns a tuple with the MaxConnections field value
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaxConnectionsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxConnections, true
}

// SetMaxConnections sets field value
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetMaxConnections(v int32) {
	o.MaxConnections = v
}

// GetUseLocation returns the UseLocation field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetUseLocation() bool {
	if o == nil || isNil(o.UseLocation) {
		var ret bool
		return ret
	}
	return *o.UseLocation
}

// GetUseLocationOk returns a tuple with the UseLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetUseLocationOk() (*bool, bool) {
	if o == nil || isNil(o.UseLocation) {
		return nil, false
	}
	return o.UseLocation, true
}

// HasUseLocation returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasUseLocation() bool {
	if o != nil && !isNil(o.UseLocation) {
		return true
	}

	return false
}

// SetUseLocation gets a reference to the given bool and assigns it to the UseLocation field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetUseLocation(v bool) {
	o.UseLocation = &v
}

// GetMaximumAllowedLocalResponseTime returns the MaximumAllowedLocalResponseTime field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaximumAllowedLocalResponseTime() string {
	if o == nil || isNil(o.MaximumAllowedLocalResponseTime) {
		var ret string
		return ret
	}
	return *o.MaximumAllowedLocalResponseTime
}

// GetMaximumAllowedLocalResponseTimeOk returns a tuple with the MaximumAllowedLocalResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaximumAllowedLocalResponseTimeOk() (*string, bool) {
	if o == nil || isNil(o.MaximumAllowedLocalResponseTime) {
		return nil, false
	}
	return o.MaximumAllowedLocalResponseTime, true
}

// HasMaximumAllowedLocalResponseTime returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasMaximumAllowedLocalResponseTime() bool {
	if o != nil && !isNil(o.MaximumAllowedLocalResponseTime) {
		return true
	}

	return false
}

// SetMaximumAllowedLocalResponseTime gets a reference to the given string and assigns it to the MaximumAllowedLocalResponseTime field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetMaximumAllowedLocalResponseTime(v string) {
	o.MaximumAllowedLocalResponseTime = &v
}

// GetMaximumAllowedNonlocalResponseTime returns the MaximumAllowedNonlocalResponseTime field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaximumAllowedNonlocalResponseTime() string {
	if o == nil || isNil(o.MaximumAllowedNonlocalResponseTime) {
		var ret string
		return ret
	}
	return *o.MaximumAllowedNonlocalResponseTime
}

// GetMaximumAllowedNonlocalResponseTimeOk returns a tuple with the MaximumAllowedNonlocalResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetMaximumAllowedNonlocalResponseTimeOk() (*string, bool) {
	if o == nil || isNil(o.MaximumAllowedNonlocalResponseTime) {
		return nil, false
	}
	return o.MaximumAllowedNonlocalResponseTime, true
}

// HasMaximumAllowedNonlocalResponseTime returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasMaximumAllowedNonlocalResponseTime() bool {
	if o != nil && !isNil(o.MaximumAllowedNonlocalResponseTime) {
		return true
	}

	return false
}

// SetMaximumAllowedNonlocalResponseTime gets a reference to the given string and assigns it to the MaximumAllowedNonlocalResponseTime field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetMaximumAllowedNonlocalResponseTime(v string) {
	o.MaximumAllowedNonlocalResponseTime = &v
}

// GetUsePasswordPolicyControl returns the UsePasswordPolicyControl field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetUsePasswordPolicyControl() bool {
	if o == nil || isNil(o.UsePasswordPolicyControl) {
		var ret bool
		return ret
	}
	return *o.UsePasswordPolicyControl
}

// GetUsePasswordPolicyControlOk returns a tuple with the UsePasswordPolicyControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetUsePasswordPolicyControlOk() (*bool, bool) {
	if o == nil || isNil(o.UsePasswordPolicyControl) {
		return nil, false
	}
	return o.UsePasswordPolicyControl, true
}

// HasUsePasswordPolicyControl returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasUsePasswordPolicyControl() bool {
	if o != nil && !isNil(o.UsePasswordPolicyControl) {
		return true
	}

	return false
}

// SetUsePasswordPolicyControl gets a reference to the given bool and assigns it to the UsePasswordPolicyControl field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetUsePasswordPolicyControl(v bool) {
	o.UsePasswordPolicyControl = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddLdapPassThroughAuthenticationHandlerRequest) SetDescription(v string) {
	o.Description = &v
}

func (o AddLdapPassThroughAuthenticationHandlerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["handlerName"] = o.HandlerName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["server"] = o.Server
	}
	if true {
		toSerialize["serverAccessMode"] = o.ServerAccessMode
	}
	if !isNil(o.DnMap) {
		toSerialize["dnMap"] = o.DnMap
	}
	if !isNil(o.BindDNPattern) {
		toSerialize["bindDNPattern"] = o.BindDNPattern
	}
	if !isNil(o.SearchBaseDN) {
		toSerialize["searchBaseDN"] = o.SearchBaseDN
	}
	if !isNil(o.SearchFilterPattern) {
		toSerialize["searchFilterPattern"] = o.SearchFilterPattern
	}
	if true {
		toSerialize["initialConnections"] = o.InitialConnections
	}
	if true {
		toSerialize["maxConnections"] = o.MaxConnections
	}
	if !isNil(o.UseLocation) {
		toSerialize["useLocation"] = o.UseLocation
	}
	if !isNil(o.MaximumAllowedLocalResponseTime) {
		toSerialize["maximumAllowedLocalResponseTime"] = o.MaximumAllowedLocalResponseTime
	}
	if !isNil(o.MaximumAllowedNonlocalResponseTime) {
		toSerialize["maximumAllowedNonlocalResponseTime"] = o.MaximumAllowedNonlocalResponseTime
	}
	if !isNil(o.UsePasswordPolicyControl) {
		toSerialize["usePasswordPolicyControl"] = o.UsePasswordPolicyControl
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableAddLdapPassThroughAuthenticationHandlerRequest struct {
	value *AddLdapPassThroughAuthenticationHandlerRequest
	isSet bool
}

func (v NullableAddLdapPassThroughAuthenticationHandlerRequest) Get() *AddLdapPassThroughAuthenticationHandlerRequest {
	return v.value
}

func (v *NullableAddLdapPassThroughAuthenticationHandlerRequest) Set(val *AddLdapPassThroughAuthenticationHandlerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLdapPassThroughAuthenticationHandlerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLdapPassThroughAuthenticationHandlerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLdapPassThroughAuthenticationHandlerRequest(val *AddLdapPassThroughAuthenticationHandlerRequest) *NullableAddLdapPassThroughAuthenticationHandlerRequest {
	return &NullableAddLdapPassThroughAuthenticationHandlerRequest{value: val, isSet: true}
}

func (v NullableAddLdapPassThroughAuthenticationHandlerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLdapPassThroughAuthenticationHandlerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
