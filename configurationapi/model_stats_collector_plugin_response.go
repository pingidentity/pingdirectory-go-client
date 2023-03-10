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

// checks if the StatsCollectorPluginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatsCollectorPluginResponse{}

// StatsCollectorPluginResponse struct for StatsCollectorPluginResponse
type StatsCollectorPluginResponse struct {
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
	Schemas                                       []EnumstatsCollectorPluginSchemaUrn                `json:"schemas"`
	// Name of the Plugin
	Id string `json:"id"`
	// The duration between statistics collections. Setting this value too small can have an impact on performance. This value should be a multiple of collection-interval.
	SampleInterval string `json:"sampleInterval"`
	// Some of the calculated statistics, such as the average and maximum queue sizes, can use multiple samples within a log interval. This value controls how often samples are gathered, and setting this value too small can have an adverse impact on performance.
	CollectionInterval      string                                               `json:"collectionInterval"`
	LdapInfo                *EnumpluginLdapInfoProp                              `json:"ldapInfo,omitempty"`
	ServerInfo              *EnumpluginServerInfoProp                            `json:"serverInfo,omitempty"`
	PerApplicationLDAPStats *EnumpluginStatsCollectorPerApplicationLDAPStatsProp `json:"perApplicationLDAPStats,omitempty"`
	LdapChangelogInfo       *EnumpluginLdapChangelogInfoProp                     `json:"ldapChangelogInfo,omitempty"`
	StatusSummaryInfo       *EnumpluginStatusSummaryInfoProp                     `json:"statusSummaryInfo,omitempty"`
	// Indicates whether this plugin should store metric samples on disk for use by the Data Metrics Server. If the Stats Collector Plugin is only being used to collect metrics for one or more StatsD Monitoring Endpoints, then this can be set to false to prevent unnecessary I/O.
	GenerateCollectorFiles *bool                             `json:"generateCollectorFiles,omitempty"`
	LocalDBBackendInfo     *EnumpluginLocalDBBackendInfoProp `json:"localDBBackendInfo,omitempty"`
	ReplicationInfo        *EnumpluginReplicationInfoProp    `json:"replicationInfo,omitempty"`
	EntryCacheInfo         *EnumpluginEntryCacheInfoProp     `json:"entryCacheInfo,omitempty"`
	HostInfo               []EnumpluginHostInfoProp          `json:"hostInfo,omitempty"`
	// If statistics should not be included for all applications, this property names the subset of applications that should be included.
	IncludedLDAPApplication []string `json:"includedLDAPApplication,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewStatsCollectorPluginResponse instantiates a new StatsCollectorPluginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsCollectorPluginResponse(schemas []EnumstatsCollectorPluginSchemaUrn, id string, sampleInterval string, collectionInterval string, enabled bool) *StatsCollectorPluginResponse {
	this := StatsCollectorPluginResponse{}
	this.Schemas = schemas
	this.Id = id
	this.SampleInterval = sampleInterval
	this.CollectionInterval = collectionInterval
	this.Enabled = enabled
	return &this
}

// NewStatsCollectorPluginResponseWithDefaults instantiates a new StatsCollectorPluginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsCollectorPluginResponseWithDefaults() *StatsCollectorPluginResponse {
	this := StatsCollectorPluginResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *StatsCollectorPluginResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *StatsCollectorPluginResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

// GetSchemas returns the Schemas field value
func (o *StatsCollectorPluginResponse) GetSchemas() []EnumstatsCollectorPluginSchemaUrn {
	if o == nil {
		var ret []EnumstatsCollectorPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetSchemasOk() ([]EnumstatsCollectorPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *StatsCollectorPluginResponse) SetSchemas(v []EnumstatsCollectorPluginSchemaUrn) {
	o.Schemas = v
}

// GetId returns the Id field value
func (o *StatsCollectorPluginResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StatsCollectorPluginResponse) SetId(v string) {
	o.Id = v
}

// GetSampleInterval returns the SampleInterval field value
func (o *StatsCollectorPluginResponse) GetSampleInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SampleInterval
}

// GetSampleIntervalOk returns a tuple with the SampleInterval field value
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetSampleIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleInterval, true
}

// SetSampleInterval sets field value
func (o *StatsCollectorPluginResponse) SetSampleInterval(v string) {
	o.SampleInterval = v
}

// GetCollectionInterval returns the CollectionInterval field value
func (o *StatsCollectorPluginResponse) GetCollectionInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionInterval
}

// GetCollectionIntervalOk returns a tuple with the CollectionInterval field value
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetCollectionIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionInterval, true
}

// SetCollectionInterval sets field value
func (o *StatsCollectorPluginResponse) SetCollectionInterval(v string) {
	o.CollectionInterval = v
}

// GetLdapInfo returns the LdapInfo field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetLdapInfo() EnumpluginLdapInfoProp {
	if o == nil || IsNil(o.LdapInfo) {
		var ret EnumpluginLdapInfoProp
		return ret
	}
	return *o.LdapInfo
}

// GetLdapInfoOk returns a tuple with the LdapInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetLdapInfoOk() (*EnumpluginLdapInfoProp, bool) {
	if o == nil || IsNil(o.LdapInfo) {
		return nil, false
	}
	return o.LdapInfo, true
}

// HasLdapInfo returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasLdapInfo() bool {
	if o != nil && !IsNil(o.LdapInfo) {
		return true
	}

	return false
}

// SetLdapInfo gets a reference to the given EnumpluginLdapInfoProp and assigns it to the LdapInfo field.
func (o *StatsCollectorPluginResponse) SetLdapInfo(v EnumpluginLdapInfoProp) {
	o.LdapInfo = &v
}

// GetServerInfo returns the ServerInfo field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetServerInfo() EnumpluginServerInfoProp {
	if o == nil || IsNil(o.ServerInfo) {
		var ret EnumpluginServerInfoProp
		return ret
	}
	return *o.ServerInfo
}

// GetServerInfoOk returns a tuple with the ServerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetServerInfoOk() (*EnumpluginServerInfoProp, bool) {
	if o == nil || IsNil(o.ServerInfo) {
		return nil, false
	}
	return o.ServerInfo, true
}

// HasServerInfo returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasServerInfo() bool {
	if o != nil && !IsNil(o.ServerInfo) {
		return true
	}

	return false
}

// SetServerInfo gets a reference to the given EnumpluginServerInfoProp and assigns it to the ServerInfo field.
func (o *StatsCollectorPluginResponse) SetServerInfo(v EnumpluginServerInfoProp) {
	o.ServerInfo = &v
}

// GetPerApplicationLDAPStats returns the PerApplicationLDAPStats field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetPerApplicationLDAPStats() EnumpluginStatsCollectorPerApplicationLDAPStatsProp {
	if o == nil || IsNil(o.PerApplicationLDAPStats) {
		var ret EnumpluginStatsCollectorPerApplicationLDAPStatsProp
		return ret
	}
	return *o.PerApplicationLDAPStats
}

// GetPerApplicationLDAPStatsOk returns a tuple with the PerApplicationLDAPStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetPerApplicationLDAPStatsOk() (*EnumpluginStatsCollectorPerApplicationLDAPStatsProp, bool) {
	if o == nil || IsNil(o.PerApplicationLDAPStats) {
		return nil, false
	}
	return o.PerApplicationLDAPStats, true
}

// HasPerApplicationLDAPStats returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasPerApplicationLDAPStats() bool {
	if o != nil && !IsNil(o.PerApplicationLDAPStats) {
		return true
	}

	return false
}

// SetPerApplicationLDAPStats gets a reference to the given EnumpluginStatsCollectorPerApplicationLDAPStatsProp and assigns it to the PerApplicationLDAPStats field.
func (o *StatsCollectorPluginResponse) SetPerApplicationLDAPStats(v EnumpluginStatsCollectorPerApplicationLDAPStatsProp) {
	o.PerApplicationLDAPStats = &v
}

// GetLdapChangelogInfo returns the LdapChangelogInfo field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetLdapChangelogInfo() EnumpluginLdapChangelogInfoProp {
	if o == nil || IsNil(o.LdapChangelogInfo) {
		var ret EnumpluginLdapChangelogInfoProp
		return ret
	}
	return *o.LdapChangelogInfo
}

// GetLdapChangelogInfoOk returns a tuple with the LdapChangelogInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetLdapChangelogInfoOk() (*EnumpluginLdapChangelogInfoProp, bool) {
	if o == nil || IsNil(o.LdapChangelogInfo) {
		return nil, false
	}
	return o.LdapChangelogInfo, true
}

// HasLdapChangelogInfo returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasLdapChangelogInfo() bool {
	if o != nil && !IsNil(o.LdapChangelogInfo) {
		return true
	}

	return false
}

// SetLdapChangelogInfo gets a reference to the given EnumpluginLdapChangelogInfoProp and assigns it to the LdapChangelogInfo field.
func (o *StatsCollectorPluginResponse) SetLdapChangelogInfo(v EnumpluginLdapChangelogInfoProp) {
	o.LdapChangelogInfo = &v
}

// GetStatusSummaryInfo returns the StatusSummaryInfo field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetStatusSummaryInfo() EnumpluginStatusSummaryInfoProp {
	if o == nil || IsNil(o.StatusSummaryInfo) {
		var ret EnumpluginStatusSummaryInfoProp
		return ret
	}
	return *o.StatusSummaryInfo
}

// GetStatusSummaryInfoOk returns a tuple with the StatusSummaryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetStatusSummaryInfoOk() (*EnumpluginStatusSummaryInfoProp, bool) {
	if o == nil || IsNil(o.StatusSummaryInfo) {
		return nil, false
	}
	return o.StatusSummaryInfo, true
}

// HasStatusSummaryInfo returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasStatusSummaryInfo() bool {
	if o != nil && !IsNil(o.StatusSummaryInfo) {
		return true
	}

	return false
}

// SetStatusSummaryInfo gets a reference to the given EnumpluginStatusSummaryInfoProp and assigns it to the StatusSummaryInfo field.
func (o *StatsCollectorPluginResponse) SetStatusSummaryInfo(v EnumpluginStatusSummaryInfoProp) {
	o.StatusSummaryInfo = &v
}

// GetGenerateCollectorFiles returns the GenerateCollectorFiles field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetGenerateCollectorFiles() bool {
	if o == nil || IsNil(o.GenerateCollectorFiles) {
		var ret bool
		return ret
	}
	return *o.GenerateCollectorFiles
}

// GetGenerateCollectorFilesOk returns a tuple with the GenerateCollectorFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetGenerateCollectorFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateCollectorFiles) {
		return nil, false
	}
	return o.GenerateCollectorFiles, true
}

// HasGenerateCollectorFiles returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasGenerateCollectorFiles() bool {
	if o != nil && !IsNil(o.GenerateCollectorFiles) {
		return true
	}

	return false
}

// SetGenerateCollectorFiles gets a reference to the given bool and assigns it to the GenerateCollectorFiles field.
func (o *StatsCollectorPluginResponse) SetGenerateCollectorFiles(v bool) {
	o.GenerateCollectorFiles = &v
}

// GetLocalDBBackendInfo returns the LocalDBBackendInfo field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetLocalDBBackendInfo() EnumpluginLocalDBBackendInfoProp {
	if o == nil || IsNil(o.LocalDBBackendInfo) {
		var ret EnumpluginLocalDBBackendInfoProp
		return ret
	}
	return *o.LocalDBBackendInfo
}

// GetLocalDBBackendInfoOk returns a tuple with the LocalDBBackendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetLocalDBBackendInfoOk() (*EnumpluginLocalDBBackendInfoProp, bool) {
	if o == nil || IsNil(o.LocalDBBackendInfo) {
		return nil, false
	}
	return o.LocalDBBackendInfo, true
}

// HasLocalDBBackendInfo returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasLocalDBBackendInfo() bool {
	if o != nil && !IsNil(o.LocalDBBackendInfo) {
		return true
	}

	return false
}

// SetLocalDBBackendInfo gets a reference to the given EnumpluginLocalDBBackendInfoProp and assigns it to the LocalDBBackendInfo field.
func (o *StatsCollectorPluginResponse) SetLocalDBBackendInfo(v EnumpluginLocalDBBackendInfoProp) {
	o.LocalDBBackendInfo = &v
}

// GetReplicationInfo returns the ReplicationInfo field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetReplicationInfo() EnumpluginReplicationInfoProp {
	if o == nil || IsNil(o.ReplicationInfo) {
		var ret EnumpluginReplicationInfoProp
		return ret
	}
	return *o.ReplicationInfo
}

// GetReplicationInfoOk returns a tuple with the ReplicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetReplicationInfoOk() (*EnumpluginReplicationInfoProp, bool) {
	if o == nil || IsNil(o.ReplicationInfo) {
		return nil, false
	}
	return o.ReplicationInfo, true
}

// HasReplicationInfo returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasReplicationInfo() bool {
	if o != nil && !IsNil(o.ReplicationInfo) {
		return true
	}

	return false
}

// SetReplicationInfo gets a reference to the given EnumpluginReplicationInfoProp and assigns it to the ReplicationInfo field.
func (o *StatsCollectorPluginResponse) SetReplicationInfo(v EnumpluginReplicationInfoProp) {
	o.ReplicationInfo = &v
}

// GetEntryCacheInfo returns the EntryCacheInfo field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetEntryCacheInfo() EnumpluginEntryCacheInfoProp {
	if o == nil || IsNil(o.EntryCacheInfo) {
		var ret EnumpluginEntryCacheInfoProp
		return ret
	}
	return *o.EntryCacheInfo
}

// GetEntryCacheInfoOk returns a tuple with the EntryCacheInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetEntryCacheInfoOk() (*EnumpluginEntryCacheInfoProp, bool) {
	if o == nil || IsNil(o.EntryCacheInfo) {
		return nil, false
	}
	return o.EntryCacheInfo, true
}

// HasEntryCacheInfo returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasEntryCacheInfo() bool {
	if o != nil && !IsNil(o.EntryCacheInfo) {
		return true
	}

	return false
}

// SetEntryCacheInfo gets a reference to the given EnumpluginEntryCacheInfoProp and assigns it to the EntryCacheInfo field.
func (o *StatsCollectorPluginResponse) SetEntryCacheInfo(v EnumpluginEntryCacheInfoProp) {
	o.EntryCacheInfo = &v
}

// GetHostInfo returns the HostInfo field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetHostInfo() []EnumpluginHostInfoProp {
	if o == nil || IsNil(o.HostInfo) {
		var ret []EnumpluginHostInfoProp
		return ret
	}
	return o.HostInfo
}

// GetHostInfoOk returns a tuple with the HostInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetHostInfoOk() ([]EnumpluginHostInfoProp, bool) {
	if o == nil || IsNil(o.HostInfo) {
		return nil, false
	}
	return o.HostInfo, true
}

// HasHostInfo returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasHostInfo() bool {
	if o != nil && !IsNil(o.HostInfo) {
		return true
	}

	return false
}

// SetHostInfo gets a reference to the given []EnumpluginHostInfoProp and assigns it to the HostInfo field.
func (o *StatsCollectorPluginResponse) SetHostInfo(v []EnumpluginHostInfoProp) {
	o.HostInfo = v
}

// GetIncludedLDAPApplication returns the IncludedLDAPApplication field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetIncludedLDAPApplication() []string {
	if o == nil || IsNil(o.IncludedLDAPApplication) {
		var ret []string
		return ret
	}
	return o.IncludedLDAPApplication
}

// GetIncludedLDAPApplicationOk returns a tuple with the IncludedLDAPApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetIncludedLDAPApplicationOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludedLDAPApplication) {
		return nil, false
	}
	return o.IncludedLDAPApplication, true
}

// HasIncludedLDAPApplication returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasIncludedLDAPApplication() bool {
	if o != nil && !IsNil(o.IncludedLDAPApplication) {
		return true
	}

	return false
}

// SetIncludedLDAPApplication gets a reference to the given []string and assigns it to the IncludedLDAPApplication field.
func (o *StatsCollectorPluginResponse) SetIncludedLDAPApplication(v []string) {
	o.IncludedLDAPApplication = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StatsCollectorPluginResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StatsCollectorPluginResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StatsCollectorPluginResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *StatsCollectorPluginResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *StatsCollectorPluginResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *StatsCollectorPluginResponse) SetEnabled(v bool) {
	o.Enabled = v
}

func (o StatsCollectorPluginResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatsCollectorPluginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	toSerialize["schemas"] = o.Schemas
	toSerialize["id"] = o.Id
	toSerialize["sampleInterval"] = o.SampleInterval
	toSerialize["collectionInterval"] = o.CollectionInterval
	if !IsNil(o.LdapInfo) {
		toSerialize["ldapInfo"] = o.LdapInfo
	}
	if !IsNil(o.ServerInfo) {
		toSerialize["serverInfo"] = o.ServerInfo
	}
	if !IsNil(o.PerApplicationLDAPStats) {
		toSerialize["perApplicationLDAPStats"] = o.PerApplicationLDAPStats
	}
	if !IsNil(o.LdapChangelogInfo) {
		toSerialize["ldapChangelogInfo"] = o.LdapChangelogInfo
	}
	if !IsNil(o.StatusSummaryInfo) {
		toSerialize["statusSummaryInfo"] = o.StatusSummaryInfo
	}
	if !IsNil(o.GenerateCollectorFiles) {
		toSerialize["generateCollectorFiles"] = o.GenerateCollectorFiles
	}
	if !IsNil(o.LocalDBBackendInfo) {
		toSerialize["localDBBackendInfo"] = o.LocalDBBackendInfo
	}
	if !IsNil(o.ReplicationInfo) {
		toSerialize["replicationInfo"] = o.ReplicationInfo
	}
	if !IsNil(o.EntryCacheInfo) {
		toSerialize["entryCacheInfo"] = o.EntryCacheInfo
	}
	if !IsNil(o.HostInfo) {
		toSerialize["hostInfo"] = o.HostInfo
	}
	if !IsNil(o.IncludedLDAPApplication) {
		toSerialize["includedLDAPApplication"] = o.IncludedLDAPApplication
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableStatsCollectorPluginResponse struct {
	value *StatsCollectorPluginResponse
	isSet bool
}

func (v NullableStatsCollectorPluginResponse) Get() *StatsCollectorPluginResponse {
	return v.value
}

func (v *NullableStatsCollectorPluginResponse) Set(val *StatsCollectorPluginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsCollectorPluginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsCollectorPluginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsCollectorPluginResponse(val *StatsCollectorPluginResponse) *NullableStatsCollectorPluginResponse {
	return &NullableStatsCollectorPluginResponse{value: val, isSet: true}
}

func (v NullableStatsCollectorPluginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsCollectorPluginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
