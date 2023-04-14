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

// checks if the JdbcBasedAccessLogPublisherResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JdbcBasedAccessLogPublisherResponse{}

// JdbcBasedAccessLogPublisherResponse struct for JdbcBasedAccessLogPublisherResponse
type JdbcBasedAccessLogPublisherResponse struct {
	// Name of the Log Publisher
	Id      string                                     `json:"id"`
	Schemas []EnumjdbcBasedAccessLogPublisherSchemaUrn `json:"schemas"`
	// The JDBC-based Database Server to use for a connection.
	Server string `json:"server"`
	// The log field mapping associates loggable fields to database column names. The table name is not part of this mapping.
	LogFieldMapping string `json:"logFieldMapping"`
	// The table name to log entries to the database server.
	LogTableName string `json:"logTableName"`
	// The maximum number of log records that can be stored in the asynchronous queue.
	QueueSize *int64 `json:"queueSize,omitempty"`
	// Indicates whether to log information about connections established to the server.
	LogConnects *bool `json:"logConnects,omitempty"`
	// Indicates whether to log information about connections that have been closed by the client or terminated by the server.
	LogDisconnects *bool `json:"logDisconnects,omitempty"`
	// Indicates whether to log information about the result of any security negotiation (e.g., SSL handshake) processing that has been performed.
	LogSecurityNegotiation *bool `json:"logSecurityNegotiation,omitempty"`
	// Indicates whether to log information about any client certificates presented to the server.
	LogClientCertificates *bool `json:"logClientCertificates,omitempty"`
	// Indicates whether to log information about requests received from clients.
	LogRequests *bool `json:"logRequests,omitempty"`
	// Indicates whether to log information about the results of client requests.
	LogResults *bool `json:"logResults,omitempty"`
	// Indicates whether to log information about search result entries sent to the client.
	LogSearchEntries *bool `json:"logSearchEntries,omitempty"`
	// Indicates whether to log information about search result references sent to the client.
	LogSearchReferences *bool `json:"logSearchReferences,omitempty"`
	// Indicates whether to log information about intermediate responses sent to the client.
	LogIntermediateResponses *bool `json:"logIntermediateResponses,omitempty"`
	// Indicates whether internal operations (for example, operations that are initiated by plugins) should be logged along with the operations that are requested by users.
	SuppressInternalOperations *bool `json:"suppressInternalOperations,omitempty"`
	// Indicates whether access messages that are generated by replication operations should be suppressed.
	SuppressReplicationOperations *bool `json:"suppressReplicationOperations,omitempty"`
	// Indicates whether to automatically log result messages for any operation in which the corresponding request was logged. In such cases, the result, entry, and reference criteria will be ignored, although the log-responses, log-search-entries, and log-search-references properties will be honored.
	CorrelateRequestsAndResults *bool `json:"correlateRequestsAndResults,omitempty"`
	// Specifies a set of connection criteria that must match the associated client connection in order for a connect, disconnect, request, or result message to be logged.
	ConnectionCriteria *string `json:"connectionCriteria,omitempty"`
	// Specifies a set of request criteria that must match the associated operation request in order for a request or result to be logged by this Access Log Publisher.
	RequestCriteria *string `json:"requestCriteria,omitempty"`
	// Specifies a set of result criteria that must match the associated operation result in order for that result to be logged by this Access Log Publisher.
	ResultCriteria *string `json:"resultCriteria,omitempty"`
	// Specifies a set of search entry criteria that must match the associated search result entry in order for that it to be logged by this Access Log Publisher.
	SearchEntryCriteria *string `json:"searchEntryCriteria,omitempty"`
	// Specifies a set of search reference criteria that must match the associated search result reference in order for that it to be logged by this Access Log Publisher.
	SearchReferenceCriteria *string `json:"searchReferenceCriteria,omitempty"`
	// A description for this Log Publisher
	Description *string `json:"description,omitempty"`
	// Indicates whether the Log Publisher is enabled for use.
	Enabled                                       bool                                               `json:"enabled"`
	LoggingErrorBehavior                          *EnumlogPublisherLoggingErrorBehaviorProp          `json:"loggingErrorBehavior,omitempty"`
	Meta                                          *MetaMeta                                          `json:"meta,omitempty"`
	Urnpingidentityschemasconfigurationmessages20 *MetaUrnPingidentitySchemasConfigurationMessages20 `json:"urn:pingidentity:schemas:configuration:messages:2.0,omitempty"`
}

// NewJdbcBasedAccessLogPublisherResponse instantiates a new JdbcBasedAccessLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJdbcBasedAccessLogPublisherResponse(id string, schemas []EnumjdbcBasedAccessLogPublisherSchemaUrn, server string, logFieldMapping string, logTableName string, enabled bool) *JdbcBasedAccessLogPublisherResponse {
	this := JdbcBasedAccessLogPublisherResponse{}
	this.Id = id
	this.Schemas = schemas
	this.Server = server
	this.LogFieldMapping = logFieldMapping
	this.LogTableName = logTableName
	this.Enabled = enabled
	return &this
}

// NewJdbcBasedAccessLogPublisherResponseWithDefaults instantiates a new JdbcBasedAccessLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJdbcBasedAccessLogPublisherResponseWithDefaults() *JdbcBasedAccessLogPublisherResponse {
	this := JdbcBasedAccessLogPublisherResponse{}
	return &this
}

// GetId returns the Id field value
func (o *JdbcBasedAccessLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JdbcBasedAccessLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *JdbcBasedAccessLogPublisherResponse) GetSchemas() []EnumjdbcBasedAccessLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumjdbcBasedAccessLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetSchemasOk() ([]EnumjdbcBasedAccessLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *JdbcBasedAccessLogPublisherResponse) SetSchemas(v []EnumjdbcBasedAccessLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetServer returns the Server field value
func (o *JdbcBasedAccessLogPublisherResponse) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *JdbcBasedAccessLogPublisherResponse) SetServer(v string) {
	o.Server = v
}

// GetLogFieldMapping returns the LogFieldMapping field value
func (o *JdbcBasedAccessLogPublisherResponse) GetLogFieldMapping() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFieldMapping
}

// GetLogFieldMappingOk returns a tuple with the LogFieldMapping field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogFieldMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFieldMapping, true
}

// SetLogFieldMapping sets field value
func (o *JdbcBasedAccessLogPublisherResponse) SetLogFieldMapping(v string) {
	o.LogFieldMapping = v
}

// GetLogTableName returns the LogTableName field value
func (o *JdbcBasedAccessLogPublisherResponse) GetLogTableName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogTableName
}

// GetLogTableNameOk returns a tuple with the LogTableName field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogTableName, true
}

// SetLogTableName sets field value
func (o *JdbcBasedAccessLogPublisherResponse) SetLogTableName(v string) {
	o.LogTableName = v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetQueueSize() int64 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int64
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetQueueSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int64 and assigns it to the QueueSize field.
func (o *JdbcBasedAccessLogPublisherResponse) SetQueueSize(v int64) {
	o.QueueSize = &v
}

// GetLogConnects returns the LogConnects field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogConnects() bool {
	if o == nil || IsNil(o.LogConnects) {
		var ret bool
		return ret
	}
	return *o.LogConnects
}

// GetLogConnectsOk returns a tuple with the LogConnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogConnectsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogConnects) {
		return nil, false
	}
	return o.LogConnects, true
}

// HasLogConnects returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogConnects() bool {
	if o != nil && !IsNil(o.LogConnects) {
		return true
	}

	return false
}

// SetLogConnects gets a reference to the given bool and assigns it to the LogConnects field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogConnects(v bool) {
	o.LogConnects = &v
}

// GetLogDisconnects returns the LogDisconnects field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogDisconnects() bool {
	if o == nil || IsNil(o.LogDisconnects) {
		var ret bool
		return ret
	}
	return *o.LogDisconnects
}

// GetLogDisconnectsOk returns a tuple with the LogDisconnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogDisconnectsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogDisconnects) {
		return nil, false
	}
	return o.LogDisconnects, true
}

// HasLogDisconnects returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogDisconnects() bool {
	if o != nil && !IsNil(o.LogDisconnects) {
		return true
	}

	return false
}

// SetLogDisconnects gets a reference to the given bool and assigns it to the LogDisconnects field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogDisconnects(v bool) {
	o.LogDisconnects = &v
}

// GetLogSecurityNegotiation returns the LogSecurityNegotiation field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogSecurityNegotiation() bool {
	if o == nil || IsNil(o.LogSecurityNegotiation) {
		var ret bool
		return ret
	}
	return *o.LogSecurityNegotiation
}

// GetLogSecurityNegotiationOk returns a tuple with the LogSecurityNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogSecurityNegotiationOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSecurityNegotiation) {
		return nil, false
	}
	return o.LogSecurityNegotiation, true
}

// HasLogSecurityNegotiation returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogSecurityNegotiation() bool {
	if o != nil && !IsNil(o.LogSecurityNegotiation) {
		return true
	}

	return false
}

// SetLogSecurityNegotiation gets a reference to the given bool and assigns it to the LogSecurityNegotiation field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogSecurityNegotiation(v bool) {
	o.LogSecurityNegotiation = &v
}

// GetLogClientCertificates returns the LogClientCertificates field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogClientCertificates() bool {
	if o == nil || IsNil(o.LogClientCertificates) {
		var ret bool
		return ret
	}
	return *o.LogClientCertificates
}

// GetLogClientCertificatesOk returns a tuple with the LogClientCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogClientCertificatesOk() (*bool, bool) {
	if o == nil || IsNil(o.LogClientCertificates) {
		return nil, false
	}
	return o.LogClientCertificates, true
}

// HasLogClientCertificates returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogClientCertificates() bool {
	if o != nil && !IsNil(o.LogClientCertificates) {
		return true
	}

	return false
}

// SetLogClientCertificates gets a reference to the given bool and assigns it to the LogClientCertificates field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogClientCertificates(v bool) {
	o.LogClientCertificates = &v
}

// GetLogRequests returns the LogRequests field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogRequests() bool {
	if o == nil || IsNil(o.LogRequests) {
		var ret bool
		return ret
	}
	return *o.LogRequests
}

// GetLogRequestsOk returns a tuple with the LogRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogRequests) {
		return nil, false
	}
	return o.LogRequests, true
}

// HasLogRequests returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogRequests() bool {
	if o != nil && !IsNil(o.LogRequests) {
		return true
	}

	return false
}

// SetLogRequests gets a reference to the given bool and assigns it to the LogRequests field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogRequests(v bool) {
	o.LogRequests = &v
}

// GetLogResults returns the LogResults field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogResults() bool {
	if o == nil || IsNil(o.LogResults) {
		var ret bool
		return ret
	}
	return *o.LogResults
}

// GetLogResultsOk returns a tuple with the LogResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogResults) {
		return nil, false
	}
	return o.LogResults, true
}

// HasLogResults returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogResults() bool {
	if o != nil && !IsNil(o.LogResults) {
		return true
	}

	return false
}

// SetLogResults gets a reference to the given bool and assigns it to the LogResults field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogResults(v bool) {
	o.LogResults = &v
}

// GetLogSearchEntries returns the LogSearchEntries field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogSearchEntries() bool {
	if o == nil || IsNil(o.LogSearchEntries) {
		var ret bool
		return ret
	}
	return *o.LogSearchEntries
}

// GetLogSearchEntriesOk returns a tuple with the LogSearchEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogSearchEntriesOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSearchEntries) {
		return nil, false
	}
	return o.LogSearchEntries, true
}

// HasLogSearchEntries returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogSearchEntries() bool {
	if o != nil && !IsNil(o.LogSearchEntries) {
		return true
	}

	return false
}

// SetLogSearchEntries gets a reference to the given bool and assigns it to the LogSearchEntries field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogSearchEntries(v bool) {
	o.LogSearchEntries = &v
}

// GetLogSearchReferences returns the LogSearchReferences field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogSearchReferences() bool {
	if o == nil || IsNil(o.LogSearchReferences) {
		var ret bool
		return ret
	}
	return *o.LogSearchReferences
}

// GetLogSearchReferencesOk returns a tuple with the LogSearchReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogSearchReferencesOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSearchReferences) {
		return nil, false
	}
	return o.LogSearchReferences, true
}

// HasLogSearchReferences returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogSearchReferences() bool {
	if o != nil && !IsNil(o.LogSearchReferences) {
		return true
	}

	return false
}

// SetLogSearchReferences gets a reference to the given bool and assigns it to the LogSearchReferences field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogSearchReferences(v bool) {
	o.LogSearchReferences = &v
}

// GetLogIntermediateResponses returns the LogIntermediateResponses field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogIntermediateResponses() bool {
	if o == nil || IsNil(o.LogIntermediateResponses) {
		var ret bool
		return ret
	}
	return *o.LogIntermediateResponses
}

// GetLogIntermediateResponsesOk returns a tuple with the LogIntermediateResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLogIntermediateResponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.LogIntermediateResponses) {
		return nil, false
	}
	return o.LogIntermediateResponses, true
}

// HasLogIntermediateResponses returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLogIntermediateResponses() bool {
	if o != nil && !IsNil(o.LogIntermediateResponses) {
		return true
	}

	return false
}

// SetLogIntermediateResponses gets a reference to the given bool and assigns it to the LogIntermediateResponses field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLogIntermediateResponses(v bool) {
	o.LogIntermediateResponses = &v
}

// GetSuppressInternalOperations returns the SuppressInternalOperations field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetSuppressInternalOperations() bool {
	if o == nil || IsNil(o.SuppressInternalOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressInternalOperations
}

// GetSuppressInternalOperationsOk returns a tuple with the SuppressInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetSuppressInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressInternalOperations) {
		return nil, false
	}
	return o.SuppressInternalOperations, true
}

// HasSuppressInternalOperations returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasSuppressInternalOperations() bool {
	if o != nil && !IsNil(o.SuppressInternalOperations) {
		return true
	}

	return false
}

// SetSuppressInternalOperations gets a reference to the given bool and assigns it to the SuppressInternalOperations field.
func (o *JdbcBasedAccessLogPublisherResponse) SetSuppressInternalOperations(v bool) {
	o.SuppressInternalOperations = &v
}

// GetSuppressReplicationOperations returns the SuppressReplicationOperations field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetSuppressReplicationOperations() bool {
	if o == nil || IsNil(o.SuppressReplicationOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressReplicationOperations
}

// GetSuppressReplicationOperationsOk returns a tuple with the SuppressReplicationOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetSuppressReplicationOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressReplicationOperations) {
		return nil, false
	}
	return o.SuppressReplicationOperations, true
}

// HasSuppressReplicationOperations returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasSuppressReplicationOperations() bool {
	if o != nil && !IsNil(o.SuppressReplicationOperations) {
		return true
	}

	return false
}

// SetSuppressReplicationOperations gets a reference to the given bool and assigns it to the SuppressReplicationOperations field.
func (o *JdbcBasedAccessLogPublisherResponse) SetSuppressReplicationOperations(v bool) {
	o.SuppressReplicationOperations = &v
}

// GetCorrelateRequestsAndResults returns the CorrelateRequestsAndResults field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetCorrelateRequestsAndResults() bool {
	if o == nil || IsNil(o.CorrelateRequestsAndResults) {
		var ret bool
		return ret
	}
	return *o.CorrelateRequestsAndResults
}

// GetCorrelateRequestsAndResultsOk returns a tuple with the CorrelateRequestsAndResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetCorrelateRequestsAndResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.CorrelateRequestsAndResults) {
		return nil, false
	}
	return o.CorrelateRequestsAndResults, true
}

// HasCorrelateRequestsAndResults returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasCorrelateRequestsAndResults() bool {
	if o != nil && !IsNil(o.CorrelateRequestsAndResults) {
		return true
	}

	return false
}

// SetCorrelateRequestsAndResults gets a reference to the given bool and assigns it to the CorrelateRequestsAndResults field.
func (o *JdbcBasedAccessLogPublisherResponse) SetCorrelateRequestsAndResults(v bool) {
	o.CorrelateRequestsAndResults = &v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetConnectionCriteria() string {
	if o == nil || IsNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasConnectionCriteria() bool {
	if o != nil && !IsNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *JdbcBasedAccessLogPublisherResponse) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *JdbcBasedAccessLogPublisherResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetResultCriteria returns the ResultCriteria field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetResultCriteria() string {
	if o == nil || IsNil(o.ResultCriteria) {
		var ret string
		return ret
	}
	return *o.ResultCriteria
}

// GetResultCriteriaOk returns a tuple with the ResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetResultCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCriteria) {
		return nil, false
	}
	return o.ResultCriteria, true
}

// HasResultCriteria returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasResultCriteria() bool {
	if o != nil && !IsNil(o.ResultCriteria) {
		return true
	}

	return false
}

// SetResultCriteria gets a reference to the given string and assigns it to the ResultCriteria field.
func (o *JdbcBasedAccessLogPublisherResponse) SetResultCriteria(v string) {
	o.ResultCriteria = &v
}

// GetSearchEntryCriteria returns the SearchEntryCriteria field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetSearchEntryCriteria() string {
	if o == nil || IsNil(o.SearchEntryCriteria) {
		var ret string
		return ret
	}
	return *o.SearchEntryCriteria
}

// GetSearchEntryCriteriaOk returns a tuple with the SearchEntryCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetSearchEntryCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.SearchEntryCriteria) {
		return nil, false
	}
	return o.SearchEntryCriteria, true
}

// HasSearchEntryCriteria returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasSearchEntryCriteria() bool {
	if o != nil && !IsNil(o.SearchEntryCriteria) {
		return true
	}

	return false
}

// SetSearchEntryCriteria gets a reference to the given string and assigns it to the SearchEntryCriteria field.
func (o *JdbcBasedAccessLogPublisherResponse) SetSearchEntryCriteria(v string) {
	o.SearchEntryCriteria = &v
}

// GetSearchReferenceCriteria returns the SearchReferenceCriteria field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetSearchReferenceCriteria() string {
	if o == nil || IsNil(o.SearchReferenceCriteria) {
		var ret string
		return ret
	}
	return *o.SearchReferenceCriteria
}

// GetSearchReferenceCriteriaOk returns a tuple with the SearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetSearchReferenceCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.SearchReferenceCriteria) {
		return nil, false
	}
	return o.SearchReferenceCriteria, true
}

// HasSearchReferenceCriteria returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasSearchReferenceCriteria() bool {
	if o != nil && !IsNil(o.SearchReferenceCriteria) {
		return true
	}

	return false
}

// SetSearchReferenceCriteria gets a reference to the given string and assigns it to the SearchReferenceCriteria field.
func (o *JdbcBasedAccessLogPublisherResponse) SetSearchReferenceCriteria(v string) {
	o.SearchReferenceCriteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *JdbcBasedAccessLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *JdbcBasedAccessLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *JdbcBasedAccessLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *JdbcBasedAccessLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || IsNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *JdbcBasedAccessLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

// GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field value if set, zero value otherwise.
func (o *JdbcBasedAccessLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20 {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		var ret MetaUrnPingidentitySchemasConfigurationMessages20
		return ret
	}
	return *o.Urnpingidentityschemasconfigurationmessages20
}

// GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JdbcBasedAccessLogPublisherResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool) {
	if o == nil || IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return nil, false
	}
	return o.Urnpingidentityschemasconfigurationmessages20, true
}

// HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.
func (o *JdbcBasedAccessLogPublisherResponse) HasUrnpingidentityschemasconfigurationmessages20() bool {
	if o != nil && !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		return true
	}

	return false
}

// SetUrnpingidentityschemasconfigurationmessages20 gets a reference to the given MetaUrnPingidentitySchemasConfigurationMessages20 and assigns it to the Urnpingidentityschemasconfigurationmessages20 field.
func (o *JdbcBasedAccessLogPublisherResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20) {
	o.Urnpingidentityschemasconfigurationmessages20 = &v
}

func (o JdbcBasedAccessLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JdbcBasedAccessLogPublisherResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["schemas"] = o.Schemas
	toSerialize["server"] = o.Server
	toSerialize["logFieldMapping"] = o.LogFieldMapping
	toSerialize["logTableName"] = o.LogTableName
	if !IsNil(o.QueueSize) {
		toSerialize["queueSize"] = o.QueueSize
	}
	if !IsNil(o.LogConnects) {
		toSerialize["logConnects"] = o.LogConnects
	}
	if !IsNil(o.LogDisconnects) {
		toSerialize["logDisconnects"] = o.LogDisconnects
	}
	if !IsNil(o.LogSecurityNegotiation) {
		toSerialize["logSecurityNegotiation"] = o.LogSecurityNegotiation
	}
	if !IsNil(o.LogClientCertificates) {
		toSerialize["logClientCertificates"] = o.LogClientCertificates
	}
	if !IsNil(o.LogRequests) {
		toSerialize["logRequests"] = o.LogRequests
	}
	if !IsNil(o.LogResults) {
		toSerialize["logResults"] = o.LogResults
	}
	if !IsNil(o.LogSearchEntries) {
		toSerialize["logSearchEntries"] = o.LogSearchEntries
	}
	if !IsNil(o.LogSearchReferences) {
		toSerialize["logSearchReferences"] = o.LogSearchReferences
	}
	if !IsNil(o.LogIntermediateResponses) {
		toSerialize["logIntermediateResponses"] = o.LogIntermediateResponses
	}
	if !IsNil(o.SuppressInternalOperations) {
		toSerialize["suppressInternalOperations"] = o.SuppressInternalOperations
	}
	if !IsNil(o.SuppressReplicationOperations) {
		toSerialize["suppressReplicationOperations"] = o.SuppressReplicationOperations
	}
	if !IsNil(o.CorrelateRequestsAndResults) {
		toSerialize["correlateRequestsAndResults"] = o.CorrelateRequestsAndResults
	}
	if !IsNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !IsNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !IsNil(o.ResultCriteria) {
		toSerialize["resultCriteria"] = o.ResultCriteria
	}
	if !IsNil(o.SearchEntryCriteria) {
		toSerialize["searchEntryCriteria"] = o.SearchEntryCriteria
	}
	if !IsNil(o.SearchReferenceCriteria) {
		toSerialize["searchReferenceCriteria"] = o.SearchReferenceCriteria
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Urnpingidentityschemasconfigurationmessages20) {
		toSerialize["urn:pingidentity:schemas:configuration:messages:2.0"] = o.Urnpingidentityschemasconfigurationmessages20
	}
	return toSerialize, nil
}

type NullableJdbcBasedAccessLogPublisherResponse struct {
	value *JdbcBasedAccessLogPublisherResponse
	isSet bool
}

func (v NullableJdbcBasedAccessLogPublisherResponse) Get() *JdbcBasedAccessLogPublisherResponse {
	return v.value
}

func (v *NullableJdbcBasedAccessLogPublisherResponse) Set(val *JdbcBasedAccessLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJdbcBasedAccessLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJdbcBasedAccessLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJdbcBasedAccessLogPublisherResponse(val *JdbcBasedAccessLogPublisherResponse) *NullableJdbcBasedAccessLogPublisherResponse {
	return &NullableJdbcBasedAccessLogPublisherResponse{value: val, isSet: true}
}

func (v NullableJdbcBasedAccessLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJdbcBasedAccessLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
