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

// AddGroovyScriptedAccessLogPublisherRequest struct for AddGroovyScriptedAccessLogPublisherRequest
type AddGroovyScriptedAccessLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName string `json:"publisherName"`
	Schemas []EnumgroovyScriptedAccessLogPublisherSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Access Log Publisher.
	ScriptClass string `json:"scriptClass"`
	// The set of arguments used to customize the behavior for the Scripted Access Log Publisher. Each configuration property should be given in the form 'name=value'.
	ScriptArgument []string `json:"scriptArgument,omitempty"`
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
	Enabled bool `json:"enabled"`
	LoggingErrorBehavior *EnumlogPublisherLoggingErrorBehaviorProp `json:"loggingErrorBehavior,omitempty"`
}

// NewAddGroovyScriptedAccessLogPublisherRequest instantiates a new AddGroovyScriptedAccessLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddGroovyScriptedAccessLogPublisherRequest(publisherName string, schemas []EnumgroovyScriptedAccessLogPublisherSchemaUrn, scriptClass string, enabled bool) *AddGroovyScriptedAccessLogPublisherRequest {
	this := AddGroovyScriptedAccessLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.ScriptClass = scriptClass
	this.Enabled = enabled
	return &this
}

// NewAddGroovyScriptedAccessLogPublisherRequestWithDefaults instantiates a new AddGroovyScriptedAccessLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddGroovyScriptedAccessLogPublisherRequestWithDefaults() *AddGroovyScriptedAccessLogPublisherRequest {
	this := AddGroovyScriptedAccessLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSchemas() []EnumgroovyScriptedAccessLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumgroovyScriptedAccessLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSchemasOk() ([]EnumgroovyScriptedAccessLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetSchemas(v []EnumgroovyScriptedAccessLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetScriptClass returns the ScriptClass field value
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetScriptClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScriptClass
}

// GetScriptClassOk returns a tuple with the ScriptClass field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetScriptClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ScriptClass, true
}

// SetScriptClass sets field value
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetScriptClass(v string) {
	o.ScriptClass = v
}

// GetScriptArgument returns the ScriptArgument field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetScriptArgument() []string {
	if o == nil || isNil(o.ScriptArgument) {
		var ret []string
		return ret
	}
	return o.ScriptArgument
}

// GetScriptArgumentOk returns a tuple with the ScriptArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetScriptArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ScriptArgument) {
    return nil, false
	}
	return o.ScriptArgument, true
}

// HasScriptArgument returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasScriptArgument() bool {
	if o != nil && !isNil(o.ScriptArgument) {
		return true
	}

	return false
}

// SetScriptArgument gets a reference to the given []string and assigns it to the ScriptArgument field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetScriptArgument(v []string) {
	o.ScriptArgument = v
}

// GetLogConnects returns the LogConnects field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogConnects() bool {
	if o == nil || isNil(o.LogConnects) {
		var ret bool
		return ret
	}
	return *o.LogConnects
}

// GetLogConnectsOk returns a tuple with the LogConnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogConnectsOk() (*bool, bool) {
	if o == nil || isNil(o.LogConnects) {
    return nil, false
	}
	return o.LogConnects, true
}

// HasLogConnects returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogConnects() bool {
	if o != nil && !isNil(o.LogConnects) {
		return true
	}

	return false
}

// SetLogConnects gets a reference to the given bool and assigns it to the LogConnects field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogConnects(v bool) {
	o.LogConnects = &v
}

// GetLogDisconnects returns the LogDisconnects field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogDisconnects() bool {
	if o == nil || isNil(o.LogDisconnects) {
		var ret bool
		return ret
	}
	return *o.LogDisconnects
}

// GetLogDisconnectsOk returns a tuple with the LogDisconnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogDisconnectsOk() (*bool, bool) {
	if o == nil || isNil(o.LogDisconnects) {
    return nil, false
	}
	return o.LogDisconnects, true
}

// HasLogDisconnects returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogDisconnects() bool {
	if o != nil && !isNil(o.LogDisconnects) {
		return true
	}

	return false
}

// SetLogDisconnects gets a reference to the given bool and assigns it to the LogDisconnects field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogDisconnects(v bool) {
	o.LogDisconnects = &v
}

// GetLogSecurityNegotiation returns the LogSecurityNegotiation field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogSecurityNegotiation() bool {
	if o == nil || isNil(o.LogSecurityNegotiation) {
		var ret bool
		return ret
	}
	return *o.LogSecurityNegotiation
}

// GetLogSecurityNegotiationOk returns a tuple with the LogSecurityNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogSecurityNegotiationOk() (*bool, bool) {
	if o == nil || isNil(o.LogSecurityNegotiation) {
    return nil, false
	}
	return o.LogSecurityNegotiation, true
}

// HasLogSecurityNegotiation returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogSecurityNegotiation() bool {
	if o != nil && !isNil(o.LogSecurityNegotiation) {
		return true
	}

	return false
}

// SetLogSecurityNegotiation gets a reference to the given bool and assigns it to the LogSecurityNegotiation field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogSecurityNegotiation(v bool) {
	o.LogSecurityNegotiation = &v
}

// GetLogClientCertificates returns the LogClientCertificates field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogClientCertificates() bool {
	if o == nil || isNil(o.LogClientCertificates) {
		var ret bool
		return ret
	}
	return *o.LogClientCertificates
}

// GetLogClientCertificatesOk returns a tuple with the LogClientCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogClientCertificatesOk() (*bool, bool) {
	if o == nil || isNil(o.LogClientCertificates) {
    return nil, false
	}
	return o.LogClientCertificates, true
}

// HasLogClientCertificates returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogClientCertificates() bool {
	if o != nil && !isNil(o.LogClientCertificates) {
		return true
	}

	return false
}

// SetLogClientCertificates gets a reference to the given bool and assigns it to the LogClientCertificates field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogClientCertificates(v bool) {
	o.LogClientCertificates = &v
}

// GetLogRequests returns the LogRequests field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogRequests() bool {
	if o == nil || isNil(o.LogRequests) {
		var ret bool
		return ret
	}
	return *o.LogRequests
}

// GetLogRequestsOk returns a tuple with the LogRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogRequestsOk() (*bool, bool) {
	if o == nil || isNil(o.LogRequests) {
    return nil, false
	}
	return o.LogRequests, true
}

// HasLogRequests returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogRequests() bool {
	if o != nil && !isNil(o.LogRequests) {
		return true
	}

	return false
}

// SetLogRequests gets a reference to the given bool and assigns it to the LogRequests field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogRequests(v bool) {
	o.LogRequests = &v
}

// GetLogResults returns the LogResults field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogResults() bool {
	if o == nil || isNil(o.LogResults) {
		var ret bool
		return ret
	}
	return *o.LogResults
}

// GetLogResultsOk returns a tuple with the LogResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogResultsOk() (*bool, bool) {
	if o == nil || isNil(o.LogResults) {
    return nil, false
	}
	return o.LogResults, true
}

// HasLogResults returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogResults() bool {
	if o != nil && !isNil(o.LogResults) {
		return true
	}

	return false
}

// SetLogResults gets a reference to the given bool and assigns it to the LogResults field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogResults(v bool) {
	o.LogResults = &v
}

// GetLogSearchEntries returns the LogSearchEntries field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogSearchEntries() bool {
	if o == nil || isNil(o.LogSearchEntries) {
		var ret bool
		return ret
	}
	return *o.LogSearchEntries
}

// GetLogSearchEntriesOk returns a tuple with the LogSearchEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogSearchEntriesOk() (*bool, bool) {
	if o == nil || isNil(o.LogSearchEntries) {
    return nil, false
	}
	return o.LogSearchEntries, true
}

// HasLogSearchEntries returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogSearchEntries() bool {
	if o != nil && !isNil(o.LogSearchEntries) {
		return true
	}

	return false
}

// SetLogSearchEntries gets a reference to the given bool and assigns it to the LogSearchEntries field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogSearchEntries(v bool) {
	o.LogSearchEntries = &v
}

// GetLogSearchReferences returns the LogSearchReferences field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogSearchReferences() bool {
	if o == nil || isNil(o.LogSearchReferences) {
		var ret bool
		return ret
	}
	return *o.LogSearchReferences
}

// GetLogSearchReferencesOk returns a tuple with the LogSearchReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogSearchReferencesOk() (*bool, bool) {
	if o == nil || isNil(o.LogSearchReferences) {
    return nil, false
	}
	return o.LogSearchReferences, true
}

// HasLogSearchReferences returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogSearchReferences() bool {
	if o != nil && !isNil(o.LogSearchReferences) {
		return true
	}

	return false
}

// SetLogSearchReferences gets a reference to the given bool and assigns it to the LogSearchReferences field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogSearchReferences(v bool) {
	o.LogSearchReferences = &v
}

// GetLogIntermediateResponses returns the LogIntermediateResponses field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogIntermediateResponses() bool {
	if o == nil || isNil(o.LogIntermediateResponses) {
		var ret bool
		return ret
	}
	return *o.LogIntermediateResponses
}

// GetLogIntermediateResponsesOk returns a tuple with the LogIntermediateResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLogIntermediateResponsesOk() (*bool, bool) {
	if o == nil || isNil(o.LogIntermediateResponses) {
    return nil, false
	}
	return o.LogIntermediateResponses, true
}

// HasLogIntermediateResponses returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLogIntermediateResponses() bool {
	if o != nil && !isNil(o.LogIntermediateResponses) {
		return true
	}

	return false
}

// SetLogIntermediateResponses gets a reference to the given bool and assigns it to the LogIntermediateResponses field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLogIntermediateResponses(v bool) {
	o.LogIntermediateResponses = &v
}

// GetSuppressInternalOperations returns the SuppressInternalOperations field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSuppressInternalOperations() bool {
	if o == nil || isNil(o.SuppressInternalOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressInternalOperations
}

// GetSuppressInternalOperationsOk returns a tuple with the SuppressInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSuppressInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.SuppressInternalOperations) {
    return nil, false
	}
	return o.SuppressInternalOperations, true
}

// HasSuppressInternalOperations returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasSuppressInternalOperations() bool {
	if o != nil && !isNil(o.SuppressInternalOperations) {
		return true
	}

	return false
}

// SetSuppressInternalOperations gets a reference to the given bool and assigns it to the SuppressInternalOperations field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetSuppressInternalOperations(v bool) {
	o.SuppressInternalOperations = &v
}

// GetSuppressReplicationOperations returns the SuppressReplicationOperations field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSuppressReplicationOperations() bool {
	if o == nil || isNil(o.SuppressReplicationOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressReplicationOperations
}

// GetSuppressReplicationOperationsOk returns a tuple with the SuppressReplicationOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSuppressReplicationOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.SuppressReplicationOperations) {
    return nil, false
	}
	return o.SuppressReplicationOperations, true
}

// HasSuppressReplicationOperations returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasSuppressReplicationOperations() bool {
	if o != nil && !isNil(o.SuppressReplicationOperations) {
		return true
	}

	return false
}

// SetSuppressReplicationOperations gets a reference to the given bool and assigns it to the SuppressReplicationOperations field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetSuppressReplicationOperations(v bool) {
	o.SuppressReplicationOperations = &v
}

// GetCorrelateRequestsAndResults returns the CorrelateRequestsAndResults field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetCorrelateRequestsAndResults() bool {
	if o == nil || isNil(o.CorrelateRequestsAndResults) {
		var ret bool
		return ret
	}
	return *o.CorrelateRequestsAndResults
}

// GetCorrelateRequestsAndResultsOk returns a tuple with the CorrelateRequestsAndResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetCorrelateRequestsAndResultsOk() (*bool, bool) {
	if o == nil || isNil(o.CorrelateRequestsAndResults) {
    return nil, false
	}
	return o.CorrelateRequestsAndResults, true
}

// HasCorrelateRequestsAndResults returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasCorrelateRequestsAndResults() bool {
	if o != nil && !isNil(o.CorrelateRequestsAndResults) {
		return true
	}

	return false
}

// SetCorrelateRequestsAndResults gets a reference to the given bool and assigns it to the CorrelateRequestsAndResults field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetCorrelateRequestsAndResults(v bool) {
	o.CorrelateRequestsAndResults = &v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetConnectionCriteria() string {
	if o == nil || isNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionCriteria) {
    return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasConnectionCriteria() bool {
	if o != nil && !isNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
    return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetResultCriteria returns the ResultCriteria field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetResultCriteria() string {
	if o == nil || isNil(o.ResultCriteria) {
		var ret string
		return ret
	}
	return *o.ResultCriteria
}

// GetResultCriteriaOk returns a tuple with the ResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetResultCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ResultCriteria) {
    return nil, false
	}
	return o.ResultCriteria, true
}

// HasResultCriteria returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasResultCriteria() bool {
	if o != nil && !isNil(o.ResultCriteria) {
		return true
	}

	return false
}

// SetResultCriteria gets a reference to the given string and assigns it to the ResultCriteria field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetResultCriteria(v string) {
	o.ResultCriteria = &v
}

// GetSearchEntryCriteria returns the SearchEntryCriteria field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSearchEntryCriteria() string {
	if o == nil || isNil(o.SearchEntryCriteria) {
		var ret string
		return ret
	}
	return *o.SearchEntryCriteria
}

// GetSearchEntryCriteriaOk returns a tuple with the SearchEntryCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSearchEntryCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.SearchEntryCriteria) {
    return nil, false
	}
	return o.SearchEntryCriteria, true
}

// HasSearchEntryCriteria returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasSearchEntryCriteria() bool {
	if o != nil && !isNil(o.SearchEntryCriteria) {
		return true
	}

	return false
}

// SetSearchEntryCriteria gets a reference to the given string and assigns it to the SearchEntryCriteria field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetSearchEntryCriteria(v string) {
	o.SearchEntryCriteria = &v
}

// GetSearchReferenceCriteria returns the SearchReferenceCriteria field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSearchReferenceCriteria() string {
	if o == nil || isNil(o.SearchReferenceCriteria) {
		var ret string
		return ret
	}
	return *o.SearchReferenceCriteria
}

// GetSearchReferenceCriteriaOk returns a tuple with the SearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetSearchReferenceCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.SearchReferenceCriteria) {
    return nil, false
	}
	return o.SearchReferenceCriteria, true
}

// HasSearchReferenceCriteria returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasSearchReferenceCriteria() bool {
	if o != nil && !isNil(o.SearchReferenceCriteria) {
		return true
	}

	return false
}

// SetSearchReferenceCriteria gets a reference to the given string and assigns it to the SearchReferenceCriteria field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetSearchReferenceCriteria(v string) {
	o.SearchReferenceCriteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddGroovyScriptedAccessLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddGroovyScriptedAccessLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddGroovyScriptedAccessLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publisherName"] = o.PublisherName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["scriptClass"] = o.ScriptClass
	}
	if !isNil(o.ScriptArgument) {
		toSerialize["scriptArgument"] = o.ScriptArgument
	}
	if !isNil(o.LogConnects) {
		toSerialize["logConnects"] = o.LogConnects
	}
	if !isNil(o.LogDisconnects) {
		toSerialize["logDisconnects"] = o.LogDisconnects
	}
	if !isNil(o.LogSecurityNegotiation) {
		toSerialize["logSecurityNegotiation"] = o.LogSecurityNegotiation
	}
	if !isNil(o.LogClientCertificates) {
		toSerialize["logClientCertificates"] = o.LogClientCertificates
	}
	if !isNil(o.LogRequests) {
		toSerialize["logRequests"] = o.LogRequests
	}
	if !isNil(o.LogResults) {
		toSerialize["logResults"] = o.LogResults
	}
	if !isNil(o.LogSearchEntries) {
		toSerialize["logSearchEntries"] = o.LogSearchEntries
	}
	if !isNil(o.LogSearchReferences) {
		toSerialize["logSearchReferences"] = o.LogSearchReferences
	}
	if !isNil(o.LogIntermediateResponses) {
		toSerialize["logIntermediateResponses"] = o.LogIntermediateResponses
	}
	if !isNil(o.SuppressInternalOperations) {
		toSerialize["suppressInternalOperations"] = o.SuppressInternalOperations
	}
	if !isNil(o.SuppressReplicationOperations) {
		toSerialize["suppressReplicationOperations"] = o.SuppressReplicationOperations
	}
	if !isNil(o.CorrelateRequestsAndResults) {
		toSerialize["correlateRequestsAndResults"] = o.CorrelateRequestsAndResults
	}
	if !isNil(o.ConnectionCriteria) {
		toSerialize["connectionCriteria"] = o.ConnectionCriteria
	}
	if !isNil(o.RequestCriteria) {
		toSerialize["requestCriteria"] = o.RequestCriteria
	}
	if !isNil(o.ResultCriteria) {
		toSerialize["resultCriteria"] = o.ResultCriteria
	}
	if !isNil(o.SearchEntryCriteria) {
		toSerialize["searchEntryCriteria"] = o.SearchEntryCriteria
	}
	if !isNil(o.SearchReferenceCriteria) {
		toSerialize["searchReferenceCriteria"] = o.SearchReferenceCriteria
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.LoggingErrorBehavior) {
		toSerialize["loggingErrorBehavior"] = o.LoggingErrorBehavior
	}
	return json.Marshal(toSerialize)
}

type NullableAddGroovyScriptedAccessLogPublisherRequest struct {
	value *AddGroovyScriptedAccessLogPublisherRequest
	isSet bool
}

func (v NullableAddGroovyScriptedAccessLogPublisherRequest) Get() *AddGroovyScriptedAccessLogPublisherRequest {
	return v.value
}

func (v *NullableAddGroovyScriptedAccessLogPublisherRequest) Set(val *AddGroovyScriptedAccessLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddGroovyScriptedAccessLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddGroovyScriptedAccessLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddGroovyScriptedAccessLogPublisherRequest(val *AddGroovyScriptedAccessLogPublisherRequest) *NullableAddGroovyScriptedAccessLogPublisherRequest {
	return &NullableAddGroovyScriptedAccessLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddGroovyScriptedAccessLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddGroovyScriptedAccessLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


