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

// checks if the AddJdbcBasedAccessLogPublisherRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddJdbcBasedAccessLogPublisherRequest{}

// AddJdbcBasedAccessLogPublisherRequest struct for AddJdbcBasedAccessLogPublisherRequest
type AddJdbcBasedAccessLogPublisherRequest struct {
	Schemas []EnumjdbcBasedAccessLogPublisherSchemaUrn `json:"schemas"`
	// The JDBC-based Database Server to use for a connection.
	Server string `json:"server"`
	// The log field mapping associates loggable fields to database column names. The table name is not part of this mapping.
	LogFieldMapping string `json:"logFieldMapping"`
	// The table name to log entries to the database server.
	LogTableName *string `json:"logTableName,omitempty"`
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
	Enabled              bool                                      `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
	// Name of the new Log Publisher
	PublisherName string `json:"publisherName"`
}

// NewAddJdbcBasedAccessLogPublisherRequest instantiates a new AddJdbcBasedAccessLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddJdbcBasedAccessLogPublisherRequest(schemas []EnumjdbcBasedAccessLogPublisherSchemaUrn, server string, logFieldMapping string, enabled bool, publisherName string) *AddJdbcBasedAccessLogPublisherRequest {
	this := AddJdbcBasedAccessLogPublisherRequest{}
	this.Schemas = schemas
	this.Server = server
	this.LogFieldMapping = logFieldMapping
	this.Enabled = enabled
	this.PublisherName = publisherName
	return &this
}

// NewAddJdbcBasedAccessLogPublisherRequestWithDefaults instantiates a new AddJdbcBasedAccessLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddJdbcBasedAccessLogPublisherRequestWithDefaults() *AddJdbcBasedAccessLogPublisherRequest {
	this := AddJdbcBasedAccessLogPublisherRequest{}
	return &this
}

// GetSchemas returns the Schemas field value
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSchemas() []EnumjdbcBasedAccessLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumjdbcBasedAccessLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSchemasOk() ([]EnumjdbcBasedAccessLogPublisherSchemaUrn, bool) {
	if o == nil {
		return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddJdbcBasedAccessLogPublisherRequest) SetSchemas(v []EnumjdbcBasedAccessLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetServer returns the Server field value
func (o *AddJdbcBasedAccessLogPublisherRequest) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *AddJdbcBasedAccessLogPublisherRequest) SetServer(v string) {
	o.Server = v
}

// GetLogFieldMapping returns the LogFieldMapping field value
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogFieldMapping() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogFieldMapping
}

// GetLogFieldMappingOk returns a tuple with the LogFieldMapping field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogFieldMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogFieldMapping, true
}

// SetLogFieldMapping sets field value
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogFieldMapping(v string) {
	o.LogFieldMapping = v
}

// GetLogTableName returns the LogTableName field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogTableName() string {
	if o == nil || IsNil(o.LogTableName) {
		var ret string
		return ret
	}
	return *o.LogTableName
}

// GetLogTableNameOk returns a tuple with the LogTableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogTableNameOk() (*string, bool) {
	if o == nil || IsNil(o.LogTableName) {
		return nil, false
	}
	return o.LogTableName, true
}

// HasLogTableName returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogTableName() bool {
	if o != nil && !IsNil(o.LogTableName) {
		return true
	}

	return false
}

// SetLogTableName gets a reference to the given string and assigns it to the LogTableName field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogTableName(v string) {
	o.LogTableName = &v
}

// GetQueueSize returns the QueueSize field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetQueueSize() int64 {
	if o == nil || IsNil(o.QueueSize) {
		var ret int64
		return ret
	}
	return *o.QueueSize
}

// GetQueueSizeOk returns a tuple with the QueueSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetQueueSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.QueueSize) {
		return nil, false
	}
	return o.QueueSize, true
}

// HasQueueSize returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasQueueSize() bool {
	if o != nil && !IsNil(o.QueueSize) {
		return true
	}

	return false
}

// SetQueueSize gets a reference to the given int64 and assigns it to the QueueSize field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetQueueSize(v int64) {
	o.QueueSize = &v
}

// GetLogConnects returns the LogConnects field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogConnects() bool {
	if o == nil || IsNil(o.LogConnects) {
		var ret bool
		return ret
	}
	return *o.LogConnects
}

// GetLogConnectsOk returns a tuple with the LogConnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogConnectsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogConnects) {
		return nil, false
	}
	return o.LogConnects, true
}

// HasLogConnects returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogConnects() bool {
	if o != nil && !IsNil(o.LogConnects) {
		return true
	}

	return false
}

// SetLogConnects gets a reference to the given bool and assigns it to the LogConnects field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogConnects(v bool) {
	o.LogConnects = &v
}

// GetLogDisconnects returns the LogDisconnects field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogDisconnects() bool {
	if o == nil || IsNil(o.LogDisconnects) {
		var ret bool
		return ret
	}
	return *o.LogDisconnects
}

// GetLogDisconnectsOk returns a tuple with the LogDisconnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogDisconnectsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogDisconnects) {
		return nil, false
	}
	return o.LogDisconnects, true
}

// HasLogDisconnects returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogDisconnects() bool {
	if o != nil && !IsNil(o.LogDisconnects) {
		return true
	}

	return false
}

// SetLogDisconnects gets a reference to the given bool and assigns it to the LogDisconnects field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogDisconnects(v bool) {
	o.LogDisconnects = &v
}

// GetLogSecurityNegotiation returns the LogSecurityNegotiation field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogSecurityNegotiation() bool {
	if o == nil || IsNil(o.LogSecurityNegotiation) {
		var ret bool
		return ret
	}
	return *o.LogSecurityNegotiation
}

// GetLogSecurityNegotiationOk returns a tuple with the LogSecurityNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogSecurityNegotiationOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSecurityNegotiation) {
		return nil, false
	}
	return o.LogSecurityNegotiation, true
}

// HasLogSecurityNegotiation returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogSecurityNegotiation() bool {
	if o != nil && !IsNil(o.LogSecurityNegotiation) {
		return true
	}

	return false
}

// SetLogSecurityNegotiation gets a reference to the given bool and assigns it to the LogSecurityNegotiation field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogSecurityNegotiation(v bool) {
	o.LogSecurityNegotiation = &v
}

// GetLogClientCertificates returns the LogClientCertificates field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogClientCertificates() bool {
	if o == nil || IsNil(o.LogClientCertificates) {
		var ret bool
		return ret
	}
	return *o.LogClientCertificates
}

// GetLogClientCertificatesOk returns a tuple with the LogClientCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogClientCertificatesOk() (*bool, bool) {
	if o == nil || IsNil(o.LogClientCertificates) {
		return nil, false
	}
	return o.LogClientCertificates, true
}

// HasLogClientCertificates returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogClientCertificates() bool {
	if o != nil && !IsNil(o.LogClientCertificates) {
		return true
	}

	return false
}

// SetLogClientCertificates gets a reference to the given bool and assigns it to the LogClientCertificates field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogClientCertificates(v bool) {
	o.LogClientCertificates = &v
}

// GetLogRequests returns the LogRequests field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogRequests() bool {
	if o == nil || IsNil(o.LogRequests) {
		var ret bool
		return ret
	}
	return *o.LogRequests
}

// GetLogRequestsOk returns a tuple with the LogRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogRequestsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogRequests) {
		return nil, false
	}
	return o.LogRequests, true
}

// HasLogRequests returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogRequests() bool {
	if o != nil && !IsNil(o.LogRequests) {
		return true
	}

	return false
}

// SetLogRequests gets a reference to the given bool and assigns it to the LogRequests field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogRequests(v bool) {
	o.LogRequests = &v
}

// GetLogResults returns the LogResults field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogResults() bool {
	if o == nil || IsNil(o.LogResults) {
		var ret bool
		return ret
	}
	return *o.LogResults
}

// GetLogResultsOk returns a tuple with the LogResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.LogResults) {
		return nil, false
	}
	return o.LogResults, true
}

// HasLogResults returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogResults() bool {
	if o != nil && !IsNil(o.LogResults) {
		return true
	}

	return false
}

// SetLogResults gets a reference to the given bool and assigns it to the LogResults field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogResults(v bool) {
	o.LogResults = &v
}

// GetLogSearchEntries returns the LogSearchEntries field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogSearchEntries() bool {
	if o == nil || IsNil(o.LogSearchEntries) {
		var ret bool
		return ret
	}
	return *o.LogSearchEntries
}

// GetLogSearchEntriesOk returns a tuple with the LogSearchEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogSearchEntriesOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSearchEntries) {
		return nil, false
	}
	return o.LogSearchEntries, true
}

// HasLogSearchEntries returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogSearchEntries() bool {
	if o != nil && !IsNil(o.LogSearchEntries) {
		return true
	}

	return false
}

// SetLogSearchEntries gets a reference to the given bool and assigns it to the LogSearchEntries field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogSearchEntries(v bool) {
	o.LogSearchEntries = &v
}

// GetLogSearchReferences returns the LogSearchReferences field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogSearchReferences() bool {
	if o == nil || IsNil(o.LogSearchReferences) {
		var ret bool
		return ret
	}
	return *o.LogSearchReferences
}

// GetLogSearchReferencesOk returns a tuple with the LogSearchReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogSearchReferencesOk() (*bool, bool) {
	if o == nil || IsNil(o.LogSearchReferences) {
		return nil, false
	}
	return o.LogSearchReferences, true
}

// HasLogSearchReferences returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogSearchReferences() bool {
	if o != nil && !IsNil(o.LogSearchReferences) {
		return true
	}

	return false
}

// SetLogSearchReferences gets a reference to the given bool and assigns it to the LogSearchReferences field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogSearchReferences(v bool) {
	o.LogSearchReferences = &v
}

// GetLogIntermediateResponses returns the LogIntermediateResponses field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogIntermediateResponses() bool {
	if o == nil || IsNil(o.LogIntermediateResponses) {
		var ret bool
		return ret
	}
	return *o.LogIntermediateResponses
}

// GetLogIntermediateResponsesOk returns a tuple with the LogIntermediateResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLogIntermediateResponsesOk() (*bool, bool) {
	if o == nil || IsNil(o.LogIntermediateResponses) {
		return nil, false
	}
	return o.LogIntermediateResponses, true
}

// HasLogIntermediateResponses returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLogIntermediateResponses() bool {
	if o != nil && !IsNil(o.LogIntermediateResponses) {
		return true
	}

	return false
}

// SetLogIntermediateResponses gets a reference to the given bool and assigns it to the LogIntermediateResponses field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLogIntermediateResponses(v bool) {
	o.LogIntermediateResponses = &v
}

// GetSuppressInternalOperations returns the SuppressInternalOperations field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSuppressInternalOperations() bool {
	if o == nil || IsNil(o.SuppressInternalOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressInternalOperations
}

// GetSuppressInternalOperationsOk returns a tuple with the SuppressInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSuppressInternalOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressInternalOperations) {
		return nil, false
	}
	return o.SuppressInternalOperations, true
}

// HasSuppressInternalOperations returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasSuppressInternalOperations() bool {
	if o != nil && !IsNil(o.SuppressInternalOperations) {
		return true
	}

	return false
}

// SetSuppressInternalOperations gets a reference to the given bool and assigns it to the SuppressInternalOperations field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetSuppressInternalOperations(v bool) {
	o.SuppressInternalOperations = &v
}

// GetSuppressReplicationOperations returns the SuppressReplicationOperations field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSuppressReplicationOperations() bool {
	if o == nil || IsNil(o.SuppressReplicationOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressReplicationOperations
}

// GetSuppressReplicationOperationsOk returns a tuple with the SuppressReplicationOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSuppressReplicationOperationsOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressReplicationOperations) {
		return nil, false
	}
	return o.SuppressReplicationOperations, true
}

// HasSuppressReplicationOperations returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasSuppressReplicationOperations() bool {
	if o != nil && !IsNil(o.SuppressReplicationOperations) {
		return true
	}

	return false
}

// SetSuppressReplicationOperations gets a reference to the given bool and assigns it to the SuppressReplicationOperations field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetSuppressReplicationOperations(v bool) {
	o.SuppressReplicationOperations = &v
}

// GetCorrelateRequestsAndResults returns the CorrelateRequestsAndResults field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetCorrelateRequestsAndResults() bool {
	if o == nil || IsNil(o.CorrelateRequestsAndResults) {
		var ret bool
		return ret
	}
	return *o.CorrelateRequestsAndResults
}

// GetCorrelateRequestsAndResultsOk returns a tuple with the CorrelateRequestsAndResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetCorrelateRequestsAndResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.CorrelateRequestsAndResults) {
		return nil, false
	}
	return o.CorrelateRequestsAndResults, true
}

// HasCorrelateRequestsAndResults returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasCorrelateRequestsAndResults() bool {
	if o != nil && !IsNil(o.CorrelateRequestsAndResults) {
		return true
	}

	return false
}

// SetCorrelateRequestsAndResults gets a reference to the given bool and assigns it to the CorrelateRequestsAndResults field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetCorrelateRequestsAndResults(v bool) {
	o.CorrelateRequestsAndResults = &v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetConnectionCriteria() string {
	if o == nil || IsNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionCriteria) {
		return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasConnectionCriteria() bool {
	if o != nil && !IsNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetRequestCriteria() string {
	if o == nil || IsNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.RequestCriteria) {
		return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasRequestCriteria() bool {
	if o != nil && !IsNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetResultCriteria returns the ResultCriteria field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetResultCriteria() string {
	if o == nil || IsNil(o.ResultCriteria) {
		var ret string
		return ret
	}
	return *o.ResultCriteria
}

// GetResultCriteriaOk returns a tuple with the ResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetResultCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.ResultCriteria) {
		return nil, false
	}
	return o.ResultCriteria, true
}

// HasResultCriteria returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasResultCriteria() bool {
	if o != nil && !IsNil(o.ResultCriteria) {
		return true
	}

	return false
}

// SetResultCriteria gets a reference to the given string and assigns it to the ResultCriteria field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetResultCriteria(v string) {
	o.ResultCriteria = &v
}

// GetSearchEntryCriteria returns the SearchEntryCriteria field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSearchEntryCriteria() string {
	if o == nil || IsNil(o.SearchEntryCriteria) {
		var ret string
		return ret
	}
	return *o.SearchEntryCriteria
}

// GetSearchEntryCriteriaOk returns a tuple with the SearchEntryCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSearchEntryCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.SearchEntryCriteria) {
		return nil, false
	}
	return o.SearchEntryCriteria, true
}

// HasSearchEntryCriteria returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasSearchEntryCriteria() bool {
	if o != nil && !IsNil(o.SearchEntryCriteria) {
		return true
	}

	return false
}

// SetSearchEntryCriteria gets a reference to the given string and assigns it to the SearchEntryCriteria field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetSearchEntryCriteria(v string) {
	o.SearchEntryCriteria = &v
}

// GetSearchReferenceCriteria returns the SearchReferenceCriteria field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSearchReferenceCriteria() string {
	if o == nil || IsNil(o.SearchReferenceCriteria) {
		var ret string
		return ret
	}
	return *o.SearchReferenceCriteria
}

// GetSearchReferenceCriteriaOk returns a tuple with the SearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetSearchReferenceCriteriaOk() (*string, bool) {
	if o == nil || IsNil(o.SearchReferenceCriteria) {
		return nil, false
	}
	return o.SearchReferenceCriteria, true
}

// HasSearchReferenceCriteria returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasSearchReferenceCriteria() bool {
	if o != nil && !IsNil(o.SearchReferenceCriteria) {
		return true
	}

	return false
}

// SetSearchReferenceCriteria gets a reference to the given string and assigns it to the SearchReferenceCriteria field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetSearchReferenceCriteria(v string) {
	o.SearchReferenceCriteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddJdbcBasedAccessLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddJdbcBasedAccessLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || IsNil(o.LoggingErrorBehavior) {
		return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !IsNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddJdbcBasedAccessLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetPublisherName returns the PublisherName field value
func (o *AddJdbcBasedAccessLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddJdbcBasedAccessLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddJdbcBasedAccessLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

func (o AddJdbcBasedAccessLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddJdbcBasedAccessLogPublisherRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["schemas"] = o.Schemas
	toSerialize["server"] = o.Server
	toSerialize["logFieldMapping"] = o.LogFieldMapping
	if !IsNil(o.LogTableName) {
		toSerialize["logTableName"] = o.LogTableName
	}
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
	toSerialize["publisherName"] = o.PublisherName
	return toSerialize, nil
}

type NullableAddJdbcBasedAccessLogPublisherRequest struct {
	value *AddJdbcBasedAccessLogPublisherRequest
	isSet bool
}

func (v NullableAddJdbcBasedAccessLogPublisherRequest) Get() *AddJdbcBasedAccessLogPublisherRequest {
	return v.value
}

func (v *NullableAddJdbcBasedAccessLogPublisherRequest) Set(val *AddJdbcBasedAccessLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddJdbcBasedAccessLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddJdbcBasedAccessLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddJdbcBasedAccessLogPublisherRequest(val *AddJdbcBasedAccessLogPublisherRequest) *NullableAddJdbcBasedAccessLogPublisherRequest {
	return &NullableAddJdbcBasedAccessLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddJdbcBasedAccessLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddJdbcBasedAccessLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
