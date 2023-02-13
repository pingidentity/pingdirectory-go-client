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

// AddPeriodicStatsLoggerPluginRequest struct for AddPeriodicStatsLoggerPluginRequest
type AddPeriodicStatsLoggerPluginRequest struct {
	// Name of the new Plugin
	PluginName string                                   `json:"pluginName"`
	Schemas    []EnumperiodicStatsLoggerPluginSchemaUrn `json:"schemas"`
	// The duration between statistics collection and logging. A new line is logged to the output for each interval. Setting this value too small can have an impact on performance.
	LogInterval string `json:"logInterval"`
	// Some of the calculated statistics, such as the average and maximum queue sizes, can use multiple samples within a log interval. This value controls how often samples are gathered. It should be a multiple of the log-interval.
	CollectionInterval string `json:"collectionInterval"`
	// If the server is idle during the specified interval, then do not log any output if this property is set to true. The server is idle if during the interval, no new connections were established, no operations were processed, and no operations are pending.
	SuppressIfIdle bool `json:"suppressIfIdle"`
	// This property controls whether the header prefix, which applies to a group of columns, appears at the start of each column header or only the first column in a group.
	HeaderPrefixPerColumn *bool `json:"headerPrefixPerColumn,omitempty"`
	// This property controls whether a value in the output is shown as empty if the value is zero.
	EmptyInsteadOfZero *bool `json:"emptyInsteadOfZero,omitempty"`
	// The number of lines to log between logging the header line that summarizes the columns in the table.
	LinesBetweenHeader      int32                                  `json:"linesBetweenHeader"`
	IncludedLDAPStat        []EnumpluginIncludedLDAPStatProp       `json:"includedLDAPStat,omitempty"`
	IncludedResourceStat    []EnumpluginIncludedResourceStatProp   `json:"includedResourceStat,omitempty"`
	HistogramFormat         EnumpluginHistogramFormatProp          `json:"histogramFormat"`
	HistogramOpType         []EnumpluginHistogramOpTypeProp        `json:"histogramOpType,omitempty"`
	PerApplicationLDAPStats *EnumpluginPerApplicationLDAPStatsProp `json:"perApplicationLDAPStats,omitempty"`
	StatusSummaryInfo       *EnumpluginStatusSummaryInfoProp       `json:"statusSummaryInfo,omitempty"`
	LdapChangelogInfo       *EnumpluginLdapChangelogInfoProp       `json:"ldapChangelogInfo,omitempty"`
	GaugeInfo               *EnumpluginGaugeInfoProp               `json:"gaugeInfo,omitempty"`
	LogFileFormat           *EnumpluginLogFileFormatProp           `json:"logFileFormat,omitempty"`
	// The file name to use for the log files generated by the Periodic Stats Logger Plugin. The path to the file can be specified either as relative to the server root or as an absolute path.
	LogFile string `json:"logFile"`
	// The UNIX permissions of the log files created by this Periodic Stats Logger Plugin.
	LogFilePermissions string `json:"logFilePermissions"`
	// Specifies whether to append to existing log files.
	Append *bool `json:"append,omitempty"`
	// The rotation policy to use for the Periodic Stats Logger Plugin .
	RotationPolicy []string `json:"rotationPolicy"`
	// A listener that should be notified whenever a log file is rotated out of service.
	RotationListener []string `json:"rotationListener,omitempty"`
	// The retention policy to use for the Periodic Stats Logger Plugin .
	RetentionPolicy      []string                            `json:"retentionPolicy"`
	LoggingErrorBehavior *EnumpluginLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
	LocalDBBackendInfo   *EnumpluginLocalDBBackendInfoProp   `json:"localDBBackendInfo,omitempty"`
	ReplicationInfo      *EnumpluginReplicationInfoProp      `json:"replicationInfo,omitempty"`
	EntryCacheInfo       *EnumpluginEntryCacheInfoProp       `json:"entryCacheInfo,omitempty"`
	HostInfo             []EnumpluginHostInfoProp            `json:"hostInfo,omitempty"`
	// If statistics should not be included for all applications, this property names the subset of applications that should be included.
	IncludedLDAPApplication []string `json:"includedLDAPApplication,omitempty"`
	// A description for this Plugin
	Description *string `json:"description,omitempty"`
	// Indicates whether the plug-in is enabled for use.
	Enabled bool `json:"enabled"`
}

// NewAddPeriodicStatsLoggerPluginRequest instantiates a new AddPeriodicStatsLoggerPluginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPeriodicStatsLoggerPluginRequest(pluginName string, schemas []EnumperiodicStatsLoggerPluginSchemaUrn, logInterval string, collectionInterval string, suppressIfIdle bool, linesBetweenHeader int32, histogramFormat EnumpluginHistogramFormatProp, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, enabled bool) *AddPeriodicStatsLoggerPluginRequest {
	this := AddPeriodicStatsLoggerPluginRequest{}
	this.PluginName = pluginName
	this.Schemas = schemas
	this.LogInterval = logInterval
	this.CollectionInterval = collectionInterval
	this.SuppressIfIdle = suppressIfIdle
	this.LinesBetweenHeader = linesBetweenHeader
	this.HistogramFormat = histogramFormat
	this.LogFile = logFile
	this.LogFilePermissions = logFilePermissions
	this.RotationPolicy = rotationPolicy
	this.RetentionPolicy = retentionPolicy
	this.Enabled = enabled
	return &this
}

// NewAddPeriodicStatsLoggerPluginRequestWithDefaults instantiates a new AddPeriodicStatsLoggerPluginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPeriodicStatsLoggerPluginRequestWithDefaults() *AddPeriodicStatsLoggerPluginRequest {
	this := AddPeriodicStatsLoggerPluginRequest{}
	return &this
}

// GetPluginName returns the PluginName field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetPluginName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PluginName
}

// GetPluginNameOk returns a tuple with the PluginName field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetPluginNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PluginName, true
}

// SetPluginName sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetPluginName(v string) {
	o.PluginName = v
}

// GetSchemas returns the Schemas field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetSchemas() []EnumperiodicStatsLoggerPluginSchemaUrn {
	if o == nil {
		var ret []EnumperiodicStatsLoggerPluginSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetSchemasOk() ([]EnumperiodicStatsLoggerPluginSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetSchemas(v []EnumperiodicStatsLoggerPluginSchemaUrn) {
	o.Schemas = v
}

// GetLogInterval returns the LogInterval field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetLogInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogInterval
}

// GetLogIntervalOk returns a tuple with the LogInterval field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLogIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogInterval, true
}

// SetLogInterval sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetLogInterval(v string) {
	o.LogInterval = v
}

// GetCollectionInterval returns the CollectionInterval field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetCollectionInterval() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionInterval
}

// GetCollectionIntervalOk returns a tuple with the CollectionInterval field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetCollectionIntervalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionInterval, true
}

// SetCollectionInterval sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetCollectionInterval(v string) {
	o.CollectionInterval = v
}

// GetSuppressIfIdle returns the SuppressIfIdle field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetSuppressIfIdle() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SuppressIfIdle
}

// GetSuppressIfIdleOk returns a tuple with the SuppressIfIdle field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetSuppressIfIdleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppressIfIdle, true
}

// SetSuppressIfIdle sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetSuppressIfIdle(v bool) {
	o.SuppressIfIdle = v
}

// GetHeaderPrefixPerColumn returns the HeaderPrefixPerColumn field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetHeaderPrefixPerColumn() bool {
	if o == nil || isNil(o.HeaderPrefixPerColumn) {
		var ret bool
		return ret
	}
	return *o.HeaderPrefixPerColumn
}

// GetHeaderPrefixPerColumnOk returns a tuple with the HeaderPrefixPerColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetHeaderPrefixPerColumnOk() (*bool, bool) {
	if o == nil || isNil(o.HeaderPrefixPerColumn) {
		return nil, false
	}
	return o.HeaderPrefixPerColumn, true
}

// HasHeaderPrefixPerColumn returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasHeaderPrefixPerColumn() bool {
	if o != nil && !isNil(o.HeaderPrefixPerColumn) {
		return true
	}

	return false
}

// SetHeaderPrefixPerColumn gets a reference to the given bool and assigns it to the HeaderPrefixPerColumn field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetHeaderPrefixPerColumn(v bool) {
	o.HeaderPrefixPerColumn = &v
}

// GetEmptyInsteadOfZero returns the EmptyInsteadOfZero field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetEmptyInsteadOfZero() bool {
	if o == nil || isNil(o.EmptyInsteadOfZero) {
		var ret bool
		return ret
	}
	return *o.EmptyInsteadOfZero
}

// GetEmptyInsteadOfZeroOk returns a tuple with the EmptyInsteadOfZero field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetEmptyInsteadOfZeroOk() (*bool, bool) {
	if o == nil || isNil(o.EmptyInsteadOfZero) {
		return nil, false
	}
	return o.EmptyInsteadOfZero, true
}

// HasEmptyInsteadOfZero returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasEmptyInsteadOfZero() bool {
	if o != nil && !isNil(o.EmptyInsteadOfZero) {
		return true
	}

	return false
}

// SetEmptyInsteadOfZero gets a reference to the given bool and assigns it to the EmptyInsteadOfZero field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetEmptyInsteadOfZero(v bool) {
	o.EmptyInsteadOfZero = &v
}

// GetLinesBetweenHeader returns the LinesBetweenHeader field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetLinesBetweenHeader() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LinesBetweenHeader
}

// GetLinesBetweenHeaderOk returns a tuple with the LinesBetweenHeader field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLinesBetweenHeaderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinesBetweenHeader, true
}

// SetLinesBetweenHeader sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetLinesBetweenHeader(v int32) {
	o.LinesBetweenHeader = v
}

// GetIncludedLDAPStat returns the IncludedLDAPStat field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedLDAPStat() []EnumpluginIncludedLDAPStatProp {
	if o == nil || isNil(o.IncludedLDAPStat) {
		var ret []EnumpluginIncludedLDAPStatProp
		return ret
	}
	return o.IncludedLDAPStat
}

// GetIncludedLDAPStatOk returns a tuple with the IncludedLDAPStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedLDAPStatOk() ([]EnumpluginIncludedLDAPStatProp, bool) {
	if o == nil || isNil(o.IncludedLDAPStat) {
		return nil, false
	}
	return o.IncludedLDAPStat, true
}

// HasIncludedLDAPStat returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasIncludedLDAPStat() bool {
	if o != nil && !isNil(o.IncludedLDAPStat) {
		return true
	}

	return false
}

// SetIncludedLDAPStat gets a reference to the given []EnumpluginIncludedLDAPStatProp and assigns it to the IncludedLDAPStat field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetIncludedLDAPStat(v []EnumpluginIncludedLDAPStatProp) {
	o.IncludedLDAPStat = v
}

// GetIncludedResourceStat returns the IncludedResourceStat field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedResourceStat() []EnumpluginIncludedResourceStatProp {
	if o == nil || isNil(o.IncludedResourceStat) {
		var ret []EnumpluginIncludedResourceStatProp
		return ret
	}
	return o.IncludedResourceStat
}

// GetIncludedResourceStatOk returns a tuple with the IncludedResourceStat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedResourceStatOk() ([]EnumpluginIncludedResourceStatProp, bool) {
	if o == nil || isNil(o.IncludedResourceStat) {
		return nil, false
	}
	return o.IncludedResourceStat, true
}

// HasIncludedResourceStat returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasIncludedResourceStat() bool {
	if o != nil && !isNil(o.IncludedResourceStat) {
		return true
	}

	return false
}

// SetIncludedResourceStat gets a reference to the given []EnumpluginIncludedResourceStatProp and assigns it to the IncludedResourceStat field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetIncludedResourceStat(v []EnumpluginIncludedResourceStatProp) {
	o.IncludedResourceStat = v
}

// GetHistogramFormat returns the HistogramFormat field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetHistogramFormat() EnumpluginHistogramFormatProp {
	if o == nil {
		var ret EnumpluginHistogramFormatProp
		return ret
	}

	return o.HistogramFormat
}

// GetHistogramFormatOk returns a tuple with the HistogramFormat field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetHistogramFormatOk() (*EnumpluginHistogramFormatProp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistogramFormat, true
}

// SetHistogramFormat sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetHistogramFormat(v EnumpluginHistogramFormatProp) {
	o.HistogramFormat = v
}

// GetHistogramOpType returns the HistogramOpType field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetHistogramOpType() []EnumpluginHistogramOpTypeProp {
	if o == nil || isNil(o.HistogramOpType) {
		var ret []EnumpluginHistogramOpTypeProp
		return ret
	}
	return o.HistogramOpType
}

// GetHistogramOpTypeOk returns a tuple with the HistogramOpType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetHistogramOpTypeOk() ([]EnumpluginHistogramOpTypeProp, bool) {
	if o == nil || isNil(o.HistogramOpType) {
		return nil, false
	}
	return o.HistogramOpType, true
}

// HasHistogramOpType returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasHistogramOpType() bool {
	if o != nil && !isNil(o.HistogramOpType) {
		return true
	}

	return false
}

// SetHistogramOpType gets a reference to the given []EnumpluginHistogramOpTypeProp and assigns it to the HistogramOpType field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetHistogramOpType(v []EnumpluginHistogramOpTypeProp) {
	o.HistogramOpType = v
}

// GetPerApplicationLDAPStats returns the PerApplicationLDAPStats field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetPerApplicationLDAPStats() EnumpluginPerApplicationLDAPStatsProp {
	if o == nil || isNil(o.PerApplicationLDAPStats) {
		var ret EnumpluginPerApplicationLDAPStatsProp
		return ret
	}
	return *o.PerApplicationLDAPStats
}

// GetPerApplicationLDAPStatsOk returns a tuple with the PerApplicationLDAPStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetPerApplicationLDAPStatsOk() (*EnumpluginPerApplicationLDAPStatsProp, bool) {
	if o == nil || isNil(o.PerApplicationLDAPStats) {
		return nil, false
	}
	return o.PerApplicationLDAPStats, true
}

// HasPerApplicationLDAPStats returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasPerApplicationLDAPStats() bool {
	if o != nil && !isNil(o.PerApplicationLDAPStats) {
		return true
	}

	return false
}

// SetPerApplicationLDAPStats gets a reference to the given EnumpluginPerApplicationLDAPStatsProp and assigns it to the PerApplicationLDAPStats field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetPerApplicationLDAPStats(v EnumpluginPerApplicationLDAPStatsProp) {
	o.PerApplicationLDAPStats = &v
}

// GetStatusSummaryInfo returns the StatusSummaryInfo field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetStatusSummaryInfo() EnumpluginStatusSummaryInfoProp {
	if o == nil || isNil(o.StatusSummaryInfo) {
		var ret EnumpluginStatusSummaryInfoProp
		return ret
	}
	return *o.StatusSummaryInfo
}

// GetStatusSummaryInfoOk returns a tuple with the StatusSummaryInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetStatusSummaryInfoOk() (*EnumpluginStatusSummaryInfoProp, bool) {
	if o == nil || isNil(o.StatusSummaryInfo) {
		return nil, false
	}
	return o.StatusSummaryInfo, true
}

// HasStatusSummaryInfo returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasStatusSummaryInfo() bool {
	if o != nil && !isNil(o.StatusSummaryInfo) {
		return true
	}

	return false
}

// SetStatusSummaryInfo gets a reference to the given EnumpluginStatusSummaryInfoProp and assigns it to the StatusSummaryInfo field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetStatusSummaryInfo(v EnumpluginStatusSummaryInfoProp) {
	o.StatusSummaryInfo = &v
}

// GetLdapChangelogInfo returns the LdapChangelogInfo field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLdapChangelogInfo() EnumpluginLdapChangelogInfoProp {
	if o == nil || isNil(o.LdapChangelogInfo) {
		var ret EnumpluginLdapChangelogInfoProp
		return ret
	}
	return *o.LdapChangelogInfo
}

// GetLdapChangelogInfoOk returns a tuple with the LdapChangelogInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLdapChangelogInfoOk() (*EnumpluginLdapChangelogInfoProp, bool) {
	if o == nil || isNil(o.LdapChangelogInfo) {
		return nil, false
	}
	return o.LdapChangelogInfo, true
}

// HasLdapChangelogInfo returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasLdapChangelogInfo() bool {
	if o != nil && !isNil(o.LdapChangelogInfo) {
		return true
	}

	return false
}

// SetLdapChangelogInfo gets a reference to the given EnumpluginLdapChangelogInfoProp and assigns it to the LdapChangelogInfo field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetLdapChangelogInfo(v EnumpluginLdapChangelogInfoProp) {
	o.LdapChangelogInfo = &v
}

// GetGaugeInfo returns the GaugeInfo field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetGaugeInfo() EnumpluginGaugeInfoProp {
	if o == nil || isNil(o.GaugeInfo) {
		var ret EnumpluginGaugeInfoProp
		return ret
	}
	return *o.GaugeInfo
}

// GetGaugeInfoOk returns a tuple with the GaugeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetGaugeInfoOk() (*EnumpluginGaugeInfoProp, bool) {
	if o == nil || isNil(o.GaugeInfo) {
		return nil, false
	}
	return o.GaugeInfo, true
}

// HasGaugeInfo returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasGaugeInfo() bool {
	if o != nil && !isNil(o.GaugeInfo) {
		return true
	}

	return false
}

// SetGaugeInfo gets a reference to the given EnumpluginGaugeInfoProp and assigns it to the GaugeInfo field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetGaugeInfo(v EnumpluginGaugeInfoProp) {
	o.GaugeInfo = &v
}

// GetLogFileFormat returns the LogFileFormat field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFileFormat() EnumpluginLogFileFormatProp {
	if o == nil || isNil(o.LogFileFormat) {
		var ret EnumpluginLogFileFormatProp
		return ret
	}
	return *o.LogFileFormat
}

// GetLogFileFormatOk returns a tuple with the LogFileFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFileFormatOk() (*EnumpluginLogFileFormatProp, bool) {
	if o == nil || isNil(o.LogFileFormat) {
		return nil, false
	}
	return o.LogFileFormat, true
}

// HasLogFileFormat returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasLogFileFormat() bool {
	if o != nil && !isNil(o.LogFileFormat) {
		return true
	}

	return false
}

// SetLogFileFormat gets a reference to the given EnumpluginLogFileFormatProp and assigns it to the LogFileFormat field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetLogFileFormat(v EnumpluginLogFileFormatProp) {
	o.LogFileFormat = &v
}

// GetLogFile returns the LogFile field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFile
}

// GetLogFileOk returns a tuple with the LogFile field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFile, true
}

// SetLogFile sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetLogFile(v string) {
	o.LogFile = v
}

// GetLogFilePermissions returns the LogFilePermissions field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFilePermissions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFilePermissions
}

// GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLogFilePermissionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFilePermissions, true
}

// SetLogFilePermissions sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetLogFilePermissions(v string) {
	o.LogFilePermissions = v
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetAppend() bool {
	if o == nil || isNil(o.Append) {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetAppendOk() (*bool, bool) {
	if o == nil || isNil(o.Append) {
		return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasAppend() bool {
	if o != nil && !isNil(o.Append) {
		return true
	}

	return false
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetAppend(v bool) {
	o.Append = &v
}

// GetRotationPolicy returns the RotationPolicy field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetRotationPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RotationPolicy
}

// GetRotationPolicyOk returns a tuple with the RotationPolicy field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetRotationPolicyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RotationPolicy, true
}

// SetRotationPolicy sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetRotationPolicy(v []string) {
	o.RotationPolicy = v
}

// GetRotationListener returns the RotationListener field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetRotationListener() []string {
	if o == nil || isNil(o.RotationListener) {
		var ret []string
		return ret
	}
	return o.RotationListener
}

// GetRotationListenerOk returns a tuple with the RotationListener field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetRotationListenerOk() ([]string, bool) {
	if o == nil || isNil(o.RotationListener) {
		return nil, false
	}
	return o.RotationListener, true
}

// HasRotationListener returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasRotationListener() bool {
	if o != nil && !isNil(o.RotationListener) {
		return true
	}

	return false
}

// SetRotationListener gets a reference to the given []string and assigns it to the RotationListener field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetRotationListener(v []string) {
	o.RotationListener = v
}

// GetRetentionPolicy returns the RetentionPolicy field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetRetentionPolicy() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RetentionPolicy
}

// GetRetentionPolicyOk returns a tuple with the RetentionPolicy field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetRetentionPolicyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetentionPolicy, true
}

// SetRetentionPolicy sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetRetentionPolicy(v []string) {
	o.RetentionPolicy = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLoggingErrorBehavior() EnumpluginLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumpluginLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLoggingErrorBehaviorOk() (*EnumpluginLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumpluginLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetLoggingErrorBehavior(v EnumpluginLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetLocalDBBackendInfo returns the LocalDBBackendInfo field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLocalDBBackendInfo() EnumpluginLocalDBBackendInfoProp {
	if o == nil || isNil(o.LocalDBBackendInfo) {
		var ret EnumpluginLocalDBBackendInfoProp
		return ret
	}
	return *o.LocalDBBackendInfo
}

// GetLocalDBBackendInfoOk returns a tuple with the LocalDBBackendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetLocalDBBackendInfoOk() (*EnumpluginLocalDBBackendInfoProp, bool) {
	if o == nil || isNil(o.LocalDBBackendInfo) {
		return nil, false
	}
	return o.LocalDBBackendInfo, true
}

// HasLocalDBBackendInfo returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasLocalDBBackendInfo() bool {
	if o != nil && !isNil(o.LocalDBBackendInfo) {
		return true
	}

	return false
}

// SetLocalDBBackendInfo gets a reference to the given EnumpluginLocalDBBackendInfoProp and assigns it to the LocalDBBackendInfo field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetLocalDBBackendInfo(v EnumpluginLocalDBBackendInfoProp) {
	o.LocalDBBackendInfo = &v
}

// GetReplicationInfo returns the ReplicationInfo field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetReplicationInfo() EnumpluginReplicationInfoProp {
	if o == nil || isNil(o.ReplicationInfo) {
		var ret EnumpluginReplicationInfoProp
		return ret
	}
	return *o.ReplicationInfo
}

// GetReplicationInfoOk returns a tuple with the ReplicationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetReplicationInfoOk() (*EnumpluginReplicationInfoProp, bool) {
	if o == nil || isNil(o.ReplicationInfo) {
		return nil, false
	}
	return o.ReplicationInfo, true
}

// HasReplicationInfo returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasReplicationInfo() bool {
	if o != nil && !isNil(o.ReplicationInfo) {
		return true
	}

	return false
}

// SetReplicationInfo gets a reference to the given EnumpluginReplicationInfoProp and assigns it to the ReplicationInfo field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetReplicationInfo(v EnumpluginReplicationInfoProp) {
	o.ReplicationInfo = &v
}

// GetEntryCacheInfo returns the EntryCacheInfo field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetEntryCacheInfo() EnumpluginEntryCacheInfoProp {
	if o == nil || isNil(o.EntryCacheInfo) {
		var ret EnumpluginEntryCacheInfoProp
		return ret
	}
	return *o.EntryCacheInfo
}

// GetEntryCacheInfoOk returns a tuple with the EntryCacheInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetEntryCacheInfoOk() (*EnumpluginEntryCacheInfoProp, bool) {
	if o == nil || isNil(o.EntryCacheInfo) {
		return nil, false
	}
	return o.EntryCacheInfo, true
}

// HasEntryCacheInfo returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasEntryCacheInfo() bool {
	if o != nil && !isNil(o.EntryCacheInfo) {
		return true
	}

	return false
}

// SetEntryCacheInfo gets a reference to the given EnumpluginEntryCacheInfoProp and assigns it to the EntryCacheInfo field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetEntryCacheInfo(v EnumpluginEntryCacheInfoProp) {
	o.EntryCacheInfo = &v
}

// GetHostInfo returns the HostInfo field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetHostInfo() []EnumpluginHostInfoProp {
	if o == nil || isNil(o.HostInfo) {
		var ret []EnumpluginHostInfoProp
		return ret
	}
	return o.HostInfo
}

// GetHostInfoOk returns a tuple with the HostInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetHostInfoOk() ([]EnumpluginHostInfoProp, bool) {
	if o == nil || isNil(o.HostInfo) {
		return nil, false
	}
	return o.HostInfo, true
}

// HasHostInfo returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasHostInfo() bool {
	if o != nil && !isNil(o.HostInfo) {
		return true
	}

	return false
}

// SetHostInfo gets a reference to the given []EnumpluginHostInfoProp and assigns it to the HostInfo field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetHostInfo(v []EnumpluginHostInfoProp) {
	o.HostInfo = v
}

// GetIncludedLDAPApplication returns the IncludedLDAPApplication field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedLDAPApplication() []string {
	if o == nil || isNil(o.IncludedLDAPApplication) {
		var ret []string
		return ret
	}
	return o.IncludedLDAPApplication
}

// GetIncludedLDAPApplicationOk returns a tuple with the IncludedLDAPApplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetIncludedLDAPApplicationOk() ([]string, bool) {
	if o == nil || isNil(o.IncludedLDAPApplication) {
		return nil, false
	}
	return o.IncludedLDAPApplication, true
}

// HasIncludedLDAPApplication returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasIncludedLDAPApplication() bool {
	if o != nil && !isNil(o.IncludedLDAPApplication) {
		return true
	}

	return false
}

// SetIncludedLDAPApplication gets a reference to the given []string and assigns it to the IncludedLDAPApplication field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetIncludedLDAPApplication(v []string) {
	o.IncludedLDAPApplication = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddPeriodicStatsLoggerPluginRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddPeriodicStatsLoggerPluginRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddPeriodicStatsLoggerPluginRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddPeriodicStatsLoggerPluginRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddPeriodicStatsLoggerPluginRequest) SetEnabled(v bool) {
	o.Enabled = v
}

func (o AddPeriodicStatsLoggerPluginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pluginName"] = o.PluginName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["logInterval"] = o.LogInterval
	}
	if true {
		toSerialize["collectionInterval"] = o.CollectionInterval
	}
	if true {
		toSerialize["suppressIfIdle"] = o.SuppressIfIdle
	}
	if !isNil(o.HeaderPrefixPerColumn) {
		toSerialize["headerPrefixPerColumn"] = o.HeaderPrefixPerColumn
	}
	if !isNil(o.EmptyInsteadOfZero) {
		toSerialize["emptyInsteadOfZero"] = o.EmptyInsteadOfZero
	}
	if true {
		toSerialize["linesBetweenHeader"] = o.LinesBetweenHeader
	}
	if !isNil(o.IncludedLDAPStat) {
		toSerialize["includedLDAPStat"] = o.IncludedLDAPStat
	}
	if !isNil(o.IncludedResourceStat) {
		toSerialize["includedResourceStat"] = o.IncludedResourceStat
	}
	if true {
		toSerialize["histogramFormat"] = o.HistogramFormat
	}
	if !isNil(o.HistogramOpType) {
		toSerialize["histogramOpType"] = o.HistogramOpType
	}
	if !isNil(o.PerApplicationLDAPStats) {
		toSerialize["perApplicationLDAPStats"] = o.PerApplicationLDAPStats
	}
	if !isNil(o.StatusSummaryInfo) {
		toSerialize["statusSummaryInfo"] = o.StatusSummaryInfo
	}
	if !isNil(o.LdapChangelogInfo) {
		toSerialize["ldapChangelogInfo"] = o.LdapChangelogInfo
	}
	if !isNil(o.GaugeInfo) {
		toSerialize["gaugeInfo"] = o.GaugeInfo
	}
	if !isNil(o.LogFileFormat) {
		toSerialize["logFileFormat"] = o.LogFileFormat
	}
	if true {
		toSerialize["logFile"] = o.LogFile
	}
	if true {
		toSerialize["logFilePermissions"] = o.LogFilePermissions
	}
	if !isNil(o.Append) {
		toSerialize["append"] = o.Append
	}
	if true {
		toSerialize["rotationPolicy"] = o.RotationPolicy
	}
	if !isNil(o.RotationListener) {
		toSerialize["rotationListener"] = o.RotationListener
	}
	if true {
		toSerialize["retentionPolicy"] = o.RetentionPolicy
	}
	if !isNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	if !isNil(o.LocalDBBackendInfo) {
		toSerialize["localDBBackendInfo"] = o.LocalDBBackendInfo
	}
	if !isNil(o.ReplicationInfo) {
		toSerialize["replicationInfo"] = o.ReplicationInfo
	}
	if !isNil(o.EntryCacheInfo) {
		toSerialize["entryCacheInfo"] = o.EntryCacheInfo
	}
	if !isNil(o.HostInfo) {
		toSerialize["hostInfo"] = o.HostInfo
	}
	if !isNil(o.IncludedLDAPApplication) {
		toSerialize["includedLDAPApplication"] = o.IncludedLDAPApplication
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableAddPeriodicStatsLoggerPluginRequest struct {
	value *AddPeriodicStatsLoggerPluginRequest
	isSet bool
}

func (v NullableAddPeriodicStatsLoggerPluginRequest) Get() *AddPeriodicStatsLoggerPluginRequest {
	return v.value
}

func (v *NullableAddPeriodicStatsLoggerPluginRequest) Set(val *AddPeriodicStatsLoggerPluginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPeriodicStatsLoggerPluginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPeriodicStatsLoggerPluginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPeriodicStatsLoggerPluginRequest(val *AddPeriodicStatsLoggerPluginRequest) *NullableAddPeriodicStatsLoggerPluginRequest {
	return &NullableAddPeriodicStatsLoggerPluginRequest{value: val, isSet: true}
}

func (v NullableAddPeriodicStatsLoggerPluginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPeriodicStatsLoggerPluginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
