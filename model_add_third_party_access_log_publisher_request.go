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

// AddThirdPartyAccessLogPublisherRequest struct for AddThirdPartyAccessLogPublisherRequest
type AddThirdPartyAccessLogPublisherRequest struct {
	// Name of the new Log Publisher
	PublisherName string `json:"publisherName"`
	Schemas []EnumthirdPartyAccessLogPublisherSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Access Log Publisher.
	ExtensionClass string `json:"extensionClass"`
	ExtensionArgument []string `json:"extensionArgument,omitempty"`
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

// NewAddThirdPartyAccessLogPublisherRequest instantiates a new AddThirdPartyAccessLogPublisherRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddThirdPartyAccessLogPublisherRequest(publisherName string, schemas []EnumthirdPartyAccessLogPublisherSchemaUrn, extensionClass string, enabled bool) *AddThirdPartyAccessLogPublisherRequest {
	this := AddThirdPartyAccessLogPublisherRequest{}
	this.PublisherName = publisherName
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewAddThirdPartyAccessLogPublisherRequestWithDefaults instantiates a new AddThirdPartyAccessLogPublisherRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddThirdPartyAccessLogPublisherRequestWithDefaults() *AddThirdPartyAccessLogPublisherRequest {
	this := AddThirdPartyAccessLogPublisherRequest{}
	return &this
}

// GetPublisherName returns the PublisherName field value
func (o *AddThirdPartyAccessLogPublisherRequest) GetPublisherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublisherName
}

// GetPublisherNameOk returns a tuple with the PublisherName field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetPublisherNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PublisherName, true
}

// SetPublisherName sets field value
func (o *AddThirdPartyAccessLogPublisherRequest) SetPublisherName(v string) {
	o.PublisherName = v
}

// GetSchemas returns the Schemas field value
func (o *AddThirdPartyAccessLogPublisherRequest) GetSchemas() []EnumthirdPartyAccessLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyAccessLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSchemasOk() ([]EnumthirdPartyAccessLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *AddThirdPartyAccessLogPublisherRequest) SetSchemas(v []EnumthirdPartyAccessLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *AddThirdPartyAccessLogPublisherRequest) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *AddThirdPartyAccessLogPublisherRequest) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetLogConnects returns the LogConnects field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogConnects() bool {
	if o == nil || isNil(o.LogConnects) {
		var ret bool
		return ret
	}
	return *o.LogConnects
}

// GetLogConnectsOk returns a tuple with the LogConnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogConnectsOk() (*bool, bool) {
	if o == nil || isNil(o.LogConnects) {
    return nil, false
	}
	return o.LogConnects, true
}

// HasLogConnects returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogConnects() bool {
	if o != nil && !isNil(o.LogConnects) {
		return true
	}

	return false
}

// SetLogConnects gets a reference to the given bool and assigns it to the LogConnects field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogConnects(v bool) {
	o.LogConnects = &v
}

// GetLogDisconnects returns the LogDisconnects field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogDisconnects() bool {
	if o == nil || isNil(o.LogDisconnects) {
		var ret bool
		return ret
	}
	return *o.LogDisconnects
}

// GetLogDisconnectsOk returns a tuple with the LogDisconnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogDisconnectsOk() (*bool, bool) {
	if o == nil || isNil(o.LogDisconnects) {
    return nil, false
	}
	return o.LogDisconnects, true
}

// HasLogDisconnects returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogDisconnects() bool {
	if o != nil && !isNil(o.LogDisconnects) {
		return true
	}

	return false
}

// SetLogDisconnects gets a reference to the given bool and assigns it to the LogDisconnects field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogDisconnects(v bool) {
	o.LogDisconnects = &v
}

// GetLogSecurityNegotiation returns the LogSecurityNegotiation field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogSecurityNegotiation() bool {
	if o == nil || isNil(o.LogSecurityNegotiation) {
		var ret bool
		return ret
	}
	return *o.LogSecurityNegotiation
}

// GetLogSecurityNegotiationOk returns a tuple with the LogSecurityNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogSecurityNegotiationOk() (*bool, bool) {
	if o == nil || isNil(o.LogSecurityNegotiation) {
    return nil, false
	}
	return o.LogSecurityNegotiation, true
}

// HasLogSecurityNegotiation returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogSecurityNegotiation() bool {
	if o != nil && !isNil(o.LogSecurityNegotiation) {
		return true
	}

	return false
}

// SetLogSecurityNegotiation gets a reference to the given bool and assigns it to the LogSecurityNegotiation field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogSecurityNegotiation(v bool) {
	o.LogSecurityNegotiation = &v
}

// GetLogClientCertificates returns the LogClientCertificates field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogClientCertificates() bool {
	if o == nil || isNil(o.LogClientCertificates) {
		var ret bool
		return ret
	}
	return *o.LogClientCertificates
}

// GetLogClientCertificatesOk returns a tuple with the LogClientCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogClientCertificatesOk() (*bool, bool) {
	if o == nil || isNil(o.LogClientCertificates) {
    return nil, false
	}
	return o.LogClientCertificates, true
}

// HasLogClientCertificates returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogClientCertificates() bool {
	if o != nil && !isNil(o.LogClientCertificates) {
		return true
	}

	return false
}

// SetLogClientCertificates gets a reference to the given bool and assigns it to the LogClientCertificates field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogClientCertificates(v bool) {
	o.LogClientCertificates = &v
}

// GetLogRequests returns the LogRequests field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogRequests() bool {
	if o == nil || isNil(o.LogRequests) {
		var ret bool
		return ret
	}
	return *o.LogRequests
}

// GetLogRequestsOk returns a tuple with the LogRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogRequestsOk() (*bool, bool) {
	if o == nil || isNil(o.LogRequests) {
    return nil, false
	}
	return o.LogRequests, true
}

// HasLogRequests returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogRequests() bool {
	if o != nil && !isNil(o.LogRequests) {
		return true
	}

	return false
}

// SetLogRequests gets a reference to the given bool and assigns it to the LogRequests field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogRequests(v bool) {
	o.LogRequests = &v
}

// GetLogResults returns the LogResults field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogResults() bool {
	if o == nil || isNil(o.LogResults) {
		var ret bool
		return ret
	}
	return *o.LogResults
}

// GetLogResultsOk returns a tuple with the LogResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogResultsOk() (*bool, bool) {
	if o == nil || isNil(o.LogResults) {
    return nil, false
	}
	return o.LogResults, true
}

// HasLogResults returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogResults() bool {
	if o != nil && !isNil(o.LogResults) {
		return true
	}

	return false
}

// SetLogResults gets a reference to the given bool and assigns it to the LogResults field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogResults(v bool) {
	o.LogResults = &v
}

// GetLogSearchEntries returns the LogSearchEntries field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogSearchEntries() bool {
	if o == nil || isNil(o.LogSearchEntries) {
		var ret bool
		return ret
	}
	return *o.LogSearchEntries
}

// GetLogSearchEntriesOk returns a tuple with the LogSearchEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogSearchEntriesOk() (*bool, bool) {
	if o == nil || isNil(o.LogSearchEntries) {
    return nil, false
	}
	return o.LogSearchEntries, true
}

// HasLogSearchEntries returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogSearchEntries() bool {
	if o != nil && !isNil(o.LogSearchEntries) {
		return true
	}

	return false
}

// SetLogSearchEntries gets a reference to the given bool and assigns it to the LogSearchEntries field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogSearchEntries(v bool) {
	o.LogSearchEntries = &v
}

// GetLogSearchReferences returns the LogSearchReferences field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogSearchReferences() bool {
	if o == nil || isNil(o.LogSearchReferences) {
		var ret bool
		return ret
	}
	return *o.LogSearchReferences
}

// GetLogSearchReferencesOk returns a tuple with the LogSearchReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogSearchReferencesOk() (*bool, bool) {
	if o == nil || isNil(o.LogSearchReferences) {
    return nil, false
	}
	return o.LogSearchReferences, true
}

// HasLogSearchReferences returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogSearchReferences() bool {
	if o != nil && !isNil(o.LogSearchReferences) {
		return true
	}

	return false
}

// SetLogSearchReferences gets a reference to the given bool and assigns it to the LogSearchReferences field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogSearchReferences(v bool) {
	o.LogSearchReferences = &v
}

// GetLogIntermediateResponses returns the LogIntermediateResponses field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogIntermediateResponses() bool {
	if o == nil || isNil(o.LogIntermediateResponses) {
		var ret bool
		return ret
	}
	return *o.LogIntermediateResponses
}

// GetLogIntermediateResponsesOk returns a tuple with the LogIntermediateResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLogIntermediateResponsesOk() (*bool, bool) {
	if o == nil || isNil(o.LogIntermediateResponses) {
    return nil, false
	}
	return o.LogIntermediateResponses, true
}

// HasLogIntermediateResponses returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLogIntermediateResponses() bool {
	if o != nil && !isNil(o.LogIntermediateResponses) {
		return true
	}

	return false
}

// SetLogIntermediateResponses gets a reference to the given bool and assigns it to the LogIntermediateResponses field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLogIntermediateResponses(v bool) {
	o.LogIntermediateResponses = &v
}

// GetSuppressInternalOperations returns the SuppressInternalOperations field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSuppressInternalOperations() bool {
	if o == nil || isNil(o.SuppressInternalOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressInternalOperations
}

// GetSuppressInternalOperationsOk returns a tuple with the SuppressInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSuppressInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.SuppressInternalOperations) {
    return nil, false
	}
	return o.SuppressInternalOperations, true
}

// HasSuppressInternalOperations returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasSuppressInternalOperations() bool {
	if o != nil && !isNil(o.SuppressInternalOperations) {
		return true
	}

	return false
}

// SetSuppressInternalOperations gets a reference to the given bool and assigns it to the SuppressInternalOperations field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetSuppressInternalOperations(v bool) {
	o.SuppressInternalOperations = &v
}

// GetSuppressReplicationOperations returns the SuppressReplicationOperations field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSuppressReplicationOperations() bool {
	if o == nil || isNil(o.SuppressReplicationOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressReplicationOperations
}

// GetSuppressReplicationOperationsOk returns a tuple with the SuppressReplicationOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSuppressReplicationOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.SuppressReplicationOperations) {
    return nil, false
	}
	return o.SuppressReplicationOperations, true
}

// HasSuppressReplicationOperations returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasSuppressReplicationOperations() bool {
	if o != nil && !isNil(o.SuppressReplicationOperations) {
		return true
	}

	return false
}

// SetSuppressReplicationOperations gets a reference to the given bool and assigns it to the SuppressReplicationOperations field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetSuppressReplicationOperations(v bool) {
	o.SuppressReplicationOperations = &v
}

// GetCorrelateRequestsAndResults returns the CorrelateRequestsAndResults field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetCorrelateRequestsAndResults() bool {
	if o == nil || isNil(o.CorrelateRequestsAndResults) {
		var ret bool
		return ret
	}
	return *o.CorrelateRequestsAndResults
}

// GetCorrelateRequestsAndResultsOk returns a tuple with the CorrelateRequestsAndResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetCorrelateRequestsAndResultsOk() (*bool, bool) {
	if o == nil || isNil(o.CorrelateRequestsAndResults) {
    return nil, false
	}
	return o.CorrelateRequestsAndResults, true
}

// HasCorrelateRequestsAndResults returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasCorrelateRequestsAndResults() bool {
	if o != nil && !isNil(o.CorrelateRequestsAndResults) {
		return true
	}

	return false
}

// SetCorrelateRequestsAndResults gets a reference to the given bool and assigns it to the CorrelateRequestsAndResults field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetCorrelateRequestsAndResults(v bool) {
	o.CorrelateRequestsAndResults = &v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetConnectionCriteria() string {
	if o == nil || isNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionCriteria) {
    return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasConnectionCriteria() bool {
	if o != nil && !isNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
    return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetResultCriteria returns the ResultCriteria field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetResultCriteria() string {
	if o == nil || isNil(o.ResultCriteria) {
		var ret string
		return ret
	}
	return *o.ResultCriteria
}

// GetResultCriteriaOk returns a tuple with the ResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetResultCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ResultCriteria) {
    return nil, false
	}
	return o.ResultCriteria, true
}

// HasResultCriteria returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasResultCriteria() bool {
	if o != nil && !isNil(o.ResultCriteria) {
		return true
	}

	return false
}

// SetResultCriteria gets a reference to the given string and assigns it to the ResultCriteria field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetResultCriteria(v string) {
	o.ResultCriteria = &v
}

// GetSearchEntryCriteria returns the SearchEntryCriteria field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSearchEntryCriteria() string {
	if o == nil || isNil(o.SearchEntryCriteria) {
		var ret string
		return ret
	}
	return *o.SearchEntryCriteria
}

// GetSearchEntryCriteriaOk returns a tuple with the SearchEntryCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSearchEntryCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.SearchEntryCriteria) {
    return nil, false
	}
	return o.SearchEntryCriteria, true
}

// HasSearchEntryCriteria returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasSearchEntryCriteria() bool {
	if o != nil && !isNil(o.SearchEntryCriteria) {
		return true
	}

	return false
}

// SetSearchEntryCriteria gets a reference to the given string and assigns it to the SearchEntryCriteria field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetSearchEntryCriteria(v string) {
	o.SearchEntryCriteria = &v
}

// GetSearchReferenceCriteria returns the SearchReferenceCriteria field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSearchReferenceCriteria() string {
	if o == nil || isNil(o.SearchReferenceCriteria) {
		var ret string
		return ret
	}
	return *o.SearchReferenceCriteria
}

// GetSearchReferenceCriteriaOk returns a tuple with the SearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetSearchReferenceCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.SearchReferenceCriteria) {
    return nil, false
	}
	return o.SearchReferenceCriteria, true
}

// HasSearchReferenceCriteria returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasSearchReferenceCriteria() bool {
	if o != nil && !isNil(o.SearchReferenceCriteria) {
		return true
	}

	return false
}

// SetSearchReferenceCriteria gets a reference to the given string and assigns it to the SearchReferenceCriteria field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetSearchReferenceCriteria(v string) {
	o.SearchReferenceCriteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *AddThirdPartyAccessLogPublisherRequest) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AddThirdPartyAccessLogPublisherRequest) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *AddThirdPartyAccessLogPublisherRequest) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *AddThirdPartyAccessLogPublisherRequest) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

func (o AddThirdPartyAccessLogPublisherRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["publisherName"] = o.PublisherName
	}
	if true {
		toSerialize["schemas"] = o.Schemas
	}
	if true {
		toSerialize["extensionClass"] = o.ExtensionClass
	}
	if !isNil(o.ExtensionArgument) {
		toSerialize["extensionArgument"] = o.ExtensionArgument
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

type NullableAddThirdPartyAccessLogPublisherRequest struct {
	value *AddThirdPartyAccessLogPublisherRequest
	isSet bool
}

func (v NullableAddThirdPartyAccessLogPublisherRequest) Get() *AddThirdPartyAccessLogPublisherRequest {
	return v.value
}

func (v *NullableAddThirdPartyAccessLogPublisherRequest) Set(val *AddThirdPartyAccessLogPublisherRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddThirdPartyAccessLogPublisherRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddThirdPartyAccessLogPublisherRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddThirdPartyAccessLogPublisherRequest(val *AddThirdPartyAccessLogPublisherRequest) *NullableAddThirdPartyAccessLogPublisherRequest {
	return &NullableAddThirdPartyAccessLogPublisherRequest{value: val, isSet: true}
}

func (v NullableAddThirdPartyAccessLogPublisherRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddThirdPartyAccessLogPublisherRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


