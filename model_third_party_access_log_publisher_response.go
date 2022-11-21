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

// ThirdPartyAccessLogPublisherResponse struct for ThirdPartyAccessLogPublisherResponse
type ThirdPartyAccessLogPublisherResponse struct {
	// Name of the Log Publisher
	Id string `json:"id"`
	Schemas []EnumthirdPartyAccessLogPublisherSchemaUrn `json:"schemas"`
	// The fully-qualified name of the Java class providing the logic for the Third Party Access Log Publisher.
	ExtensionClass string `json:"extensionClass"`
	// The set of arguments used to customize the behavior for the Third Party Access Log Publisher. Each configuration property should be given in the form 'name=value'.
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
	Meta *MetaMeta `json:"meta,omitempty"`
}

// NewThirdPartyAccessLogPublisherResponse instantiates a new ThirdPartyAccessLogPublisherResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartyAccessLogPublisherResponse(id string, schemas []EnumthirdPartyAccessLogPublisherSchemaUrn, extensionClass string, enabled bool) *ThirdPartyAccessLogPublisherResponse {
	this := ThirdPartyAccessLogPublisherResponse{}
	this.Id = id
	this.Schemas = schemas
	this.ExtensionClass = extensionClass
	this.Enabled = enabled
	return &this
}

// NewThirdPartyAccessLogPublisherResponseWithDefaults instantiates a new ThirdPartyAccessLogPublisherResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartyAccessLogPublisherResponseWithDefaults() *ThirdPartyAccessLogPublisherResponse {
	this := ThirdPartyAccessLogPublisherResponse{}
	return &this
}

// GetId returns the Id field value
func (o *ThirdPartyAccessLogPublisherResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ThirdPartyAccessLogPublisherResponse) SetId(v string) {
	o.Id = v
}

// GetSchemas returns the Schemas field value
func (o *ThirdPartyAccessLogPublisherResponse) GetSchemas() []EnumthirdPartyAccessLogPublisherSchemaUrn {
	if o == nil {
		var ret []EnumthirdPartyAccessLogPublisherSchemaUrn
		return ret
	}

	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetSchemasOk() ([]EnumthirdPartyAccessLogPublisherSchemaUrn, bool) {
	if o == nil {
    return nil, false
	}
	return o.Schemas, true
}

// SetSchemas sets field value
func (o *ThirdPartyAccessLogPublisherResponse) SetSchemas(v []EnumthirdPartyAccessLogPublisherSchemaUrn) {
	o.Schemas = v
}

// GetExtensionClass returns the ExtensionClass field value
func (o *ThirdPartyAccessLogPublisherResponse) GetExtensionClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExtensionClass
}

// GetExtensionClassOk returns a tuple with the ExtensionClass field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetExtensionClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ExtensionClass, true
}

// SetExtensionClass sets field value
func (o *ThirdPartyAccessLogPublisherResponse) SetExtensionClass(v string) {
	o.ExtensionClass = v
}

// GetExtensionArgument returns the ExtensionArgument field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetExtensionArgument() []string {
	if o == nil || isNil(o.ExtensionArgument) {
		var ret []string
		return ret
	}
	return o.ExtensionArgument
}

// GetExtensionArgumentOk returns a tuple with the ExtensionArgument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetExtensionArgumentOk() ([]string, bool) {
	if o == nil || isNil(o.ExtensionArgument) {
    return nil, false
	}
	return o.ExtensionArgument, true
}

// HasExtensionArgument returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasExtensionArgument() bool {
	if o != nil && !isNil(o.ExtensionArgument) {
		return true
	}

	return false
}

// SetExtensionArgument gets a reference to the given []string and assigns it to the ExtensionArgument field.
func (o *ThirdPartyAccessLogPublisherResponse) SetExtensionArgument(v []string) {
	o.ExtensionArgument = v
}

// GetLogConnects returns the LogConnects field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogConnects() bool {
	if o == nil || isNil(o.LogConnects) {
		var ret bool
		return ret
	}
	return *o.LogConnects
}

// GetLogConnectsOk returns a tuple with the LogConnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogConnectsOk() (*bool, bool) {
	if o == nil || isNil(o.LogConnects) {
    return nil, false
	}
	return o.LogConnects, true
}

// HasLogConnects returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogConnects() bool {
	if o != nil && !isNil(o.LogConnects) {
		return true
	}

	return false
}

// SetLogConnects gets a reference to the given bool and assigns it to the LogConnects field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogConnects(v bool) {
	o.LogConnects = &v
}

// GetLogDisconnects returns the LogDisconnects field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogDisconnects() bool {
	if o == nil || isNil(o.LogDisconnects) {
		var ret bool
		return ret
	}
	return *o.LogDisconnects
}

// GetLogDisconnectsOk returns a tuple with the LogDisconnects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogDisconnectsOk() (*bool, bool) {
	if o == nil || isNil(o.LogDisconnects) {
    return nil, false
	}
	return o.LogDisconnects, true
}

// HasLogDisconnects returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogDisconnects() bool {
	if o != nil && !isNil(o.LogDisconnects) {
		return true
	}

	return false
}

// SetLogDisconnects gets a reference to the given bool and assigns it to the LogDisconnects field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogDisconnects(v bool) {
	o.LogDisconnects = &v
}

// GetLogSecurityNegotiation returns the LogSecurityNegotiation field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogSecurityNegotiation() bool {
	if o == nil || isNil(o.LogSecurityNegotiation) {
		var ret bool
		return ret
	}
	return *o.LogSecurityNegotiation
}

// GetLogSecurityNegotiationOk returns a tuple with the LogSecurityNegotiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogSecurityNegotiationOk() (*bool, bool) {
	if o == nil || isNil(o.LogSecurityNegotiation) {
    return nil, false
	}
	return o.LogSecurityNegotiation, true
}

// HasLogSecurityNegotiation returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogSecurityNegotiation() bool {
	if o != nil && !isNil(o.LogSecurityNegotiation) {
		return true
	}

	return false
}

// SetLogSecurityNegotiation gets a reference to the given bool and assigns it to the LogSecurityNegotiation field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogSecurityNegotiation(v bool) {
	o.LogSecurityNegotiation = &v
}

// GetLogClientCertificates returns the LogClientCertificates field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogClientCertificates() bool {
	if o == nil || isNil(o.LogClientCertificates) {
		var ret bool
		return ret
	}
	return *o.LogClientCertificates
}

// GetLogClientCertificatesOk returns a tuple with the LogClientCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogClientCertificatesOk() (*bool, bool) {
	if o == nil || isNil(o.LogClientCertificates) {
    return nil, false
	}
	return o.LogClientCertificates, true
}

// HasLogClientCertificates returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogClientCertificates() bool {
	if o != nil && !isNil(o.LogClientCertificates) {
		return true
	}

	return false
}

// SetLogClientCertificates gets a reference to the given bool and assigns it to the LogClientCertificates field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogClientCertificates(v bool) {
	o.LogClientCertificates = &v
}

// GetLogRequests returns the LogRequests field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogRequests() bool {
	if o == nil || isNil(o.LogRequests) {
		var ret bool
		return ret
	}
	return *o.LogRequests
}

// GetLogRequestsOk returns a tuple with the LogRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogRequestsOk() (*bool, bool) {
	if o == nil || isNil(o.LogRequests) {
    return nil, false
	}
	return o.LogRequests, true
}

// HasLogRequests returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogRequests() bool {
	if o != nil && !isNil(o.LogRequests) {
		return true
	}

	return false
}

// SetLogRequests gets a reference to the given bool and assigns it to the LogRequests field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogRequests(v bool) {
	o.LogRequests = &v
}

// GetLogResults returns the LogResults field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogResults() bool {
	if o == nil || isNil(o.LogResults) {
		var ret bool
		return ret
	}
	return *o.LogResults
}

// GetLogResultsOk returns a tuple with the LogResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogResultsOk() (*bool, bool) {
	if o == nil || isNil(o.LogResults) {
    return nil, false
	}
	return o.LogResults, true
}

// HasLogResults returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogResults() bool {
	if o != nil && !isNil(o.LogResults) {
		return true
	}

	return false
}

// SetLogResults gets a reference to the given bool and assigns it to the LogResults field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogResults(v bool) {
	o.LogResults = &v
}

// GetLogSearchEntries returns the LogSearchEntries field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogSearchEntries() bool {
	if o == nil || isNil(o.LogSearchEntries) {
		var ret bool
		return ret
	}
	return *o.LogSearchEntries
}

// GetLogSearchEntriesOk returns a tuple with the LogSearchEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogSearchEntriesOk() (*bool, bool) {
	if o == nil || isNil(o.LogSearchEntries) {
    return nil, false
	}
	return o.LogSearchEntries, true
}

// HasLogSearchEntries returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogSearchEntries() bool {
	if o != nil && !isNil(o.LogSearchEntries) {
		return true
	}

	return false
}

// SetLogSearchEntries gets a reference to the given bool and assigns it to the LogSearchEntries field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogSearchEntries(v bool) {
	o.LogSearchEntries = &v
}

// GetLogSearchReferences returns the LogSearchReferences field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogSearchReferences() bool {
	if o == nil || isNil(o.LogSearchReferences) {
		var ret bool
		return ret
	}
	return *o.LogSearchReferences
}

// GetLogSearchReferencesOk returns a tuple with the LogSearchReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogSearchReferencesOk() (*bool, bool) {
	if o == nil || isNil(o.LogSearchReferences) {
    return nil, false
	}
	return o.LogSearchReferences, true
}

// HasLogSearchReferences returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogSearchReferences() bool {
	if o != nil && !isNil(o.LogSearchReferences) {
		return true
	}

	return false
}

// SetLogSearchReferences gets a reference to the given bool and assigns it to the LogSearchReferences field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogSearchReferences(v bool) {
	o.LogSearchReferences = &v
}

// GetLogIntermediateResponses returns the LogIntermediateResponses field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogIntermediateResponses() bool {
	if o == nil || isNil(o.LogIntermediateResponses) {
		var ret bool
		return ret
	}
	return *o.LogIntermediateResponses
}

// GetLogIntermediateResponsesOk returns a tuple with the LogIntermediateResponses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLogIntermediateResponsesOk() (*bool, bool) {
	if o == nil || isNil(o.LogIntermediateResponses) {
    return nil, false
	}
	return o.LogIntermediateResponses, true
}

// HasLogIntermediateResponses returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLogIntermediateResponses() bool {
	if o != nil && !isNil(o.LogIntermediateResponses) {
		return true
	}

	return false
}

// SetLogIntermediateResponses gets a reference to the given bool and assigns it to the LogIntermediateResponses field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLogIntermediateResponses(v bool) {
	o.LogIntermediateResponses = &v
}

// GetSuppressInternalOperations returns the SuppressInternalOperations field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetSuppressInternalOperations() bool {
	if o == nil || isNil(o.SuppressInternalOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressInternalOperations
}

// GetSuppressInternalOperationsOk returns a tuple with the SuppressInternalOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetSuppressInternalOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.SuppressInternalOperations) {
    return nil, false
	}
	return o.SuppressInternalOperations, true
}

// HasSuppressInternalOperations returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasSuppressInternalOperations() bool {
	if o != nil && !isNil(o.SuppressInternalOperations) {
		return true
	}

	return false
}

// SetSuppressInternalOperations gets a reference to the given bool and assigns it to the SuppressInternalOperations field.
func (o *ThirdPartyAccessLogPublisherResponse) SetSuppressInternalOperations(v bool) {
	o.SuppressInternalOperations = &v
}

// GetSuppressReplicationOperations returns the SuppressReplicationOperations field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetSuppressReplicationOperations() bool {
	if o == nil || isNil(o.SuppressReplicationOperations) {
		var ret bool
		return ret
	}
	return *o.SuppressReplicationOperations
}

// GetSuppressReplicationOperationsOk returns a tuple with the SuppressReplicationOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetSuppressReplicationOperationsOk() (*bool, bool) {
	if o == nil || isNil(o.SuppressReplicationOperations) {
    return nil, false
	}
	return o.SuppressReplicationOperations, true
}

// HasSuppressReplicationOperations returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasSuppressReplicationOperations() bool {
	if o != nil && !isNil(o.SuppressReplicationOperations) {
		return true
	}

	return false
}

// SetSuppressReplicationOperations gets a reference to the given bool and assigns it to the SuppressReplicationOperations field.
func (o *ThirdPartyAccessLogPublisherResponse) SetSuppressReplicationOperations(v bool) {
	o.SuppressReplicationOperations = &v
}

// GetCorrelateRequestsAndResults returns the CorrelateRequestsAndResults field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetCorrelateRequestsAndResults() bool {
	if o == nil || isNil(o.CorrelateRequestsAndResults) {
		var ret bool
		return ret
	}
	return *o.CorrelateRequestsAndResults
}

// GetCorrelateRequestsAndResultsOk returns a tuple with the CorrelateRequestsAndResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetCorrelateRequestsAndResultsOk() (*bool, bool) {
	if o == nil || isNil(o.CorrelateRequestsAndResults) {
    return nil, false
	}
	return o.CorrelateRequestsAndResults, true
}

// HasCorrelateRequestsAndResults returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasCorrelateRequestsAndResults() bool {
	if o != nil && !isNil(o.CorrelateRequestsAndResults) {
		return true
	}

	return false
}

// SetCorrelateRequestsAndResults gets a reference to the given bool and assigns it to the CorrelateRequestsAndResults field.
func (o *ThirdPartyAccessLogPublisherResponse) SetCorrelateRequestsAndResults(v bool) {
	o.CorrelateRequestsAndResults = &v
}

// GetConnectionCriteria returns the ConnectionCriteria field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetConnectionCriteria() string {
	if o == nil || isNil(o.ConnectionCriteria) {
		var ret string
		return ret
	}
	return *o.ConnectionCriteria
}

// GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetConnectionCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ConnectionCriteria) {
    return nil, false
	}
	return o.ConnectionCriteria, true
}

// HasConnectionCriteria returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasConnectionCriteria() bool {
	if o != nil && !isNil(o.ConnectionCriteria) {
		return true
	}

	return false
}

// SetConnectionCriteria gets a reference to the given string and assigns it to the ConnectionCriteria field.
func (o *ThirdPartyAccessLogPublisherResponse) SetConnectionCriteria(v string) {
	o.ConnectionCriteria = &v
}

// GetRequestCriteria returns the RequestCriteria field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetRequestCriteria() string {
	if o == nil || isNil(o.RequestCriteria) {
		var ret string
		return ret
	}
	return *o.RequestCriteria
}

// GetRequestCriteriaOk returns a tuple with the RequestCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetRequestCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.RequestCriteria) {
    return nil, false
	}
	return o.RequestCriteria, true
}

// HasRequestCriteria returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasRequestCriteria() bool {
	if o != nil && !isNil(o.RequestCriteria) {
		return true
	}

	return false
}

// SetRequestCriteria gets a reference to the given string and assigns it to the RequestCriteria field.
func (o *ThirdPartyAccessLogPublisherResponse) SetRequestCriteria(v string) {
	o.RequestCriteria = &v
}

// GetResultCriteria returns the ResultCriteria field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetResultCriteria() string {
	if o == nil || isNil(o.ResultCriteria) {
		var ret string
		return ret
	}
	return *o.ResultCriteria
}

// GetResultCriteriaOk returns a tuple with the ResultCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetResultCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.ResultCriteria) {
    return nil, false
	}
	return o.ResultCriteria, true
}

// HasResultCriteria returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasResultCriteria() bool {
	if o != nil && !isNil(o.ResultCriteria) {
		return true
	}

	return false
}

// SetResultCriteria gets a reference to the given string and assigns it to the ResultCriteria field.
func (o *ThirdPartyAccessLogPublisherResponse) SetResultCriteria(v string) {
	o.ResultCriteria = &v
}

// GetSearchEntryCriteria returns the SearchEntryCriteria field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetSearchEntryCriteria() string {
	if o == nil || isNil(o.SearchEntryCriteria) {
		var ret string
		return ret
	}
	return *o.SearchEntryCriteria
}

// GetSearchEntryCriteriaOk returns a tuple with the SearchEntryCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetSearchEntryCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.SearchEntryCriteria) {
    return nil, false
	}
	return o.SearchEntryCriteria, true
}

// HasSearchEntryCriteria returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasSearchEntryCriteria() bool {
	if o != nil && !isNil(o.SearchEntryCriteria) {
		return true
	}

	return false
}

// SetSearchEntryCriteria gets a reference to the given string and assigns it to the SearchEntryCriteria field.
func (o *ThirdPartyAccessLogPublisherResponse) SetSearchEntryCriteria(v string) {
	o.SearchEntryCriteria = &v
}

// GetSearchReferenceCriteria returns the SearchReferenceCriteria field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetSearchReferenceCriteria() string {
	if o == nil || isNil(o.SearchReferenceCriteria) {
		var ret string
		return ret
	}
	return *o.SearchReferenceCriteria
}

// GetSearchReferenceCriteriaOk returns a tuple with the SearchReferenceCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetSearchReferenceCriteriaOk() (*string, bool) {
	if o == nil || isNil(o.SearchReferenceCriteria) {
    return nil, false
	}
	return o.SearchReferenceCriteria, true
}

// HasSearchReferenceCriteria returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasSearchReferenceCriteria() bool {
	if o != nil && !isNil(o.SearchReferenceCriteria) {
		return true
	}

	return false
}

// SetSearchReferenceCriteria gets a reference to the given string and assigns it to the SearchReferenceCriteria field.
func (o *ThirdPartyAccessLogPublisherResponse) SetSearchReferenceCriteria(v string) {
	o.SearchReferenceCriteria = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ThirdPartyAccessLogPublisherResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *ThirdPartyAccessLogPublisherResponse) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ThirdPartyAccessLogPublisherResponse) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLoggingErrorBehavior returns the LoggingErrorBehavior field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetLoggingErrorBehavior() EnumlogPublisherLoggingErrorBehaviorProp {
	if o == nil || isNil(o.LoggingErrorBehavior) {
		var ret EnumlogPublisherLoggingErrorBehaviorProp
		return ret
	}
	return *o.LoggingErrorBehavior
}

// GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetLoggingErrorBehaviorOk() (*EnumlogPublisherLoggingErrorBehaviorProp, bool) {
	if o == nil || isNil(o.LoggingErrorBehavior) {
    return nil, false
	}
	return o.LoggingErrorBehavior, true
}

// HasLoggingErrorBehavior returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasLoggingErrorBehavior() bool {
	if o != nil && !isNil(o.LoggingErrorBehavior) {
		return true
	}

	return false
}

// SetLoggingErrorBehavior gets a reference to the given EnumlogPublisherLoggingErrorBehaviorProp and assigns it to the LoggingErrorBehavior field.
func (o *ThirdPartyAccessLogPublisherResponse) SetLoggingErrorBehavior(v EnumlogPublisherLoggingErrorBehaviorProp) {
	o.LoggingErrorBehavior = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ThirdPartyAccessLogPublisherResponse) GetMeta() MetaMeta {
	if o == nil || isNil(o.Meta) {
		var ret MetaMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartyAccessLogPublisherResponse) GetMetaOk() (*MetaMeta, bool) {
	if o == nil || isNil(o.Meta) {
    return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ThirdPartyAccessLogPublisherResponse) HasMeta() bool {
	if o != nil && !isNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given MetaMeta and assigns it to the Meta field.
func (o *ThirdPartyAccessLogPublisherResponse) SetMeta(v MetaMeta) {
	o.Meta = &v
}

func (o ThirdPartyAccessLogPublisherResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
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
	if !isNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableThirdPartyAccessLogPublisherResponse struct {
	value *ThirdPartyAccessLogPublisherResponse
	isSet bool
}

func (v NullableThirdPartyAccessLogPublisherResponse) Get() *ThirdPartyAccessLogPublisherResponse {
	return v.value
}

func (v *NullableThirdPartyAccessLogPublisherResponse) Set(val *ThirdPartyAccessLogPublisherResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartyAccessLogPublisherResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartyAccessLogPublisherResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartyAccessLogPublisherResponse(val *ThirdPartyAccessLogPublisherResponse) *NullableThirdPartyAccessLogPublisherResponse {
	return &NullableThirdPartyAccessLogPublisherResponse{value: val, isSet: true}
}

func (v NullableThirdPartyAccessLogPublisherResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartyAccessLogPublisherResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


