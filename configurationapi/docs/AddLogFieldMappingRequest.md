# AddLogFieldMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MappingName** | **string** | Name of the new Log Field Mapping | 
**Schemas** | [**[]EnumerrorLogFieldMappingSchemaUrn**](EnumerrorLogFieldMappingSchemaUrn.md) |  | 
**LogFieldTimestamp** | Pointer to **string** | The time that the log message was generated. | [optional] 
**LogFieldConnectionID** | Pointer to **string** | The connection ID assigned to the client connection. | [optional] 
**LogFieldStartupid** | Pointer to **string** | The startup ID for the Directory Server. A different value will be generated each time the server is started. | [optional] 
**LogFieldProductName** | Pointer to **string** | The name for this Directory Server product, which may be used to identify which product was used to log the message if multiple products log to the same database table. | [optional] 
**LogFieldInstanceName** | Pointer to **string** | A name that uniquely identifies this Directory Server instance, which may be used to identify which instance was used to log the message if multiple server instances log to the same database table. | [optional] 
**LogFieldOperationID** | Pointer to **string** | The operation ID for the operation processed by the server. | [optional] 
**LogFieldMessageType** | Pointer to **string** | The type of log message. Message types may include \&quot;CONNECT\&quot;, \&quot;DISCONNECT\&quot;, \&quot;FORWARD\&quot;, \&quot;RESULT\&quot;, \&quot;ENTRY\&quot;, or \&quot;REFERENCE\&quot;. | [optional] 
**LogFieldOperationType** | Pointer to **string** | The type of operation that was processed. | [optional] 
**LogFieldMessageID** | Pointer to **string** | The numeric value which uniquely identifies the type of message. | [optional] 
**LogFieldResultCode** | Pointer to **string** | The numeric result code for the operation. | [optional] 
**LogFieldMessage** | Pointer to **string** | The text of the log message. | [optional] 
**LogFieldOrigin** | Pointer to **string** | The origin for the operation. Values may include \&quot;replication\&quot; (if the operation was received via replication), \&quot;internal\&quot; (if it was an internal operation processed by a server component), or \&quot;external\&quot; (if it was a request from a client). | [optional] 
**LogFieldRequesterDN** | Pointer to **string** | The DN of the user that requested the operation. | [optional] 
**LogFieldDisconnectReason** | Pointer to **string** | The reason that the client connection was closed. | [optional] 
**LogFieldDeleteOldRDN** | Pointer to **string** | Indicates whether the old RDN values should be removed from an entry while processing a modify DN operation. | [optional] 
**LogFieldAuthenticatedUserDN** | Pointer to **string** | The DN of the user that authenticated to the server. | [optional] 
**LogFieldProcessingTime** | Pointer to **string** | The length of time (in milliseconds with microsecond accuracy) required to process the operation. | [optional] 
**LogFieldRequestedAttributes** | Pointer to **string** | The set of requested attributes for the search operation. | [optional] 
**LogFieldSASLMechanismName** | Pointer to **string** | The name of the SASL mechanism used to authenticate. | [optional] 
**LogFieldNewRDN** | Pointer to **string** | The new RDN to use for a modify DN operation. | [optional] 
**LogFieldBaseDN** | Pointer to **string** | The base DN for a search operation. | [optional] 
**LogFieldBindDN** | Pointer to **string** | The bind DN for a bind operation. | [optional] 
**LogFieldMatchedDN** | Pointer to **string** | The DN of the superior entry closest to the DN specified by the client. | [optional] 
**LogFieldRequesterIPAddress** | Pointer to **string** | The IP address of the client that requested the operation. | [optional] 
**LogFieldAuthenticationType** | Pointer to **string** | The type of authentication requested by the client. | [optional] 
**LogFieldNewSuperiorDN** | Pointer to **string** | The new superior DN from a modify DN operation. | [optional] 
**LogFieldFilter** | Pointer to **string** | The filter from a search operation. | [optional] 
**LogFieldAlternateAuthorizationDN** | Pointer to **string** | The DN of the alternate authorization identity used when processing the operation. | [optional] 
**LogFieldEntryDN** | Pointer to **string** | The DN of the entry targeted by the operation. | [optional] 
**LogFieldEntriesReturned** | Pointer to **string** | The number of search result entries returned to the client. | [optional] 
**LogFieldAuthenticationFailureID** | Pointer to **string** | The numeric identifier for the authentication failure reason. | [optional] 
**LogFieldRequestOID** | Pointer to **string** | The OID of an extended request. | [optional] 
**LogFieldResponseOID** | Pointer to **string** | The OID of an extended response. | [optional] 
**LogFieldTargetProtocol** | Pointer to **string** | The protocol used when forwarding the request to a backend server. | [optional] 
**LogFieldTargetPort** | Pointer to **string** | The network port of the Directory Server to which the client connection has been established, or of the backend server to which the request has been forwarded. | [optional] 
**LogFieldTargetAddress** | Pointer to **string** | The network address of the Directory Server to which the client connection has been established. | [optional] 
**LogFieldTargetAttribute** | Pointer to **string** | The name of the attribute targeted by a compare operation. | [optional] 
**LogFieldTargetHost** | Pointer to **string** | The address of the server to which the request has been forwarded. | [optional] 
**LogFieldProtocolVersion** | Pointer to **string** | The protocol version used by the client when communicating with the Directory Server. | [optional] 
**LogFieldProtocolName** | Pointer to **string** | The name of the protocol the client is using to communicate with the Directory Server. | [optional] 
**LogFieldAuthenticationFailureReason** | Pointer to **string** | A message explaining the reason that the authentication attempt failed. | [optional] 
**LogFieldAdditionalInformation** | Pointer to **string** | Additional information about the operation that was processed which was not returned to the client. | [optional] 
**LogFieldUnindexed** | Pointer to **string** | Indicates whether the requested search operation was unindexed. | [optional] 
**LogFieldScope** | Pointer to **string** | The scope for the search operation. | [optional] 
**LogFieldReferralUrls** | Pointer to **string** | The referral URLs returned to the client. | [optional] 
**LogFieldSourceAddress** | Pointer to **string** | The address of the client from which the connection was established. | [optional] 
**LogFieldMessageIDToAbandon** | Pointer to **string** | The message ID of the operation to be abandoned. | [optional] 
**LogFieldResponseControls** | Pointer to **string** | The OIDs of the response controls returned to the client. | [optional] 
**LogFieldRequestControls** | Pointer to **string** | The OIDs of the request controls returned to the client. | [optional] 
**LogFieldIntermediateClientResult** | Pointer to **string** | The contents of the intermediate client response control returned to the client. | [optional] 
**LogFieldIntermediateClientRequest** | Pointer to **string** | The contents of the intermediate client request control provided by the client. | [optional] 
**LogFieldReplicationChangeID** | Pointer to **string** | The replication change ID. | [optional] 
**Description** | Pointer to **string** | A description for this Log Field Mapping | [optional] 
**LogFieldCategory** | Pointer to **string** | The category for the log message. | [optional] 
**LogFieldSeverity** | Pointer to **string** | The severity for the log message. | [optional] 

## Methods

### NewAddLogFieldMappingRequest

`func NewAddLogFieldMappingRequest(mappingName string, schemas []EnumerrorLogFieldMappingSchemaUrn, ) *AddLogFieldMappingRequest`

NewAddLogFieldMappingRequest instantiates a new AddLogFieldMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLogFieldMappingRequestWithDefaults

`func NewAddLogFieldMappingRequestWithDefaults() *AddLogFieldMappingRequest`

NewAddLogFieldMappingRequestWithDefaults instantiates a new AddLogFieldMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappingName

`func (o *AddLogFieldMappingRequest) GetMappingName() string`

GetMappingName returns the MappingName field if non-nil, zero value otherwise.

### GetMappingNameOk

`func (o *AddLogFieldMappingRequest) GetMappingNameOk() (*string, bool)`

GetMappingNameOk returns a tuple with the MappingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingName

`func (o *AddLogFieldMappingRequest) SetMappingName(v string)`

SetMappingName sets MappingName field to given value.


### GetSchemas

`func (o *AddLogFieldMappingRequest) GetSchemas() []EnumerrorLogFieldMappingSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddLogFieldMappingRequest) GetSchemasOk() (*[]EnumerrorLogFieldMappingSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddLogFieldMappingRequest) SetSchemas(v []EnumerrorLogFieldMappingSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogFieldTimestamp

`func (o *AddLogFieldMappingRequest) GetLogFieldTimestamp() string`

GetLogFieldTimestamp returns the LogFieldTimestamp field if non-nil, zero value otherwise.

### GetLogFieldTimestampOk

`func (o *AddLogFieldMappingRequest) GetLogFieldTimestampOk() (*string, bool)`

GetLogFieldTimestampOk returns a tuple with the LogFieldTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTimestamp

`func (o *AddLogFieldMappingRequest) SetLogFieldTimestamp(v string)`

SetLogFieldTimestamp sets LogFieldTimestamp field to given value.

### HasLogFieldTimestamp

`func (o *AddLogFieldMappingRequest) HasLogFieldTimestamp() bool`

HasLogFieldTimestamp returns a boolean if a field has been set.

### GetLogFieldConnectionID

`func (o *AddLogFieldMappingRequest) GetLogFieldConnectionID() string`

GetLogFieldConnectionID returns the LogFieldConnectionID field if non-nil, zero value otherwise.

### GetLogFieldConnectionIDOk

`func (o *AddLogFieldMappingRequest) GetLogFieldConnectionIDOk() (*string, bool)`

GetLogFieldConnectionIDOk returns a tuple with the LogFieldConnectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldConnectionID

`func (o *AddLogFieldMappingRequest) SetLogFieldConnectionID(v string)`

SetLogFieldConnectionID sets LogFieldConnectionID field to given value.

### HasLogFieldConnectionID

`func (o *AddLogFieldMappingRequest) HasLogFieldConnectionID() bool`

HasLogFieldConnectionID returns a boolean if a field has been set.

### GetLogFieldStartupid

`func (o *AddLogFieldMappingRequest) GetLogFieldStartupid() string`

GetLogFieldStartupid returns the LogFieldStartupid field if non-nil, zero value otherwise.

### GetLogFieldStartupidOk

`func (o *AddLogFieldMappingRequest) GetLogFieldStartupidOk() (*string, bool)`

GetLogFieldStartupidOk returns a tuple with the LogFieldStartupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldStartupid

`func (o *AddLogFieldMappingRequest) SetLogFieldStartupid(v string)`

SetLogFieldStartupid sets LogFieldStartupid field to given value.

### HasLogFieldStartupid

`func (o *AddLogFieldMappingRequest) HasLogFieldStartupid() bool`

HasLogFieldStartupid returns a boolean if a field has been set.

### GetLogFieldProductName

`func (o *AddLogFieldMappingRequest) GetLogFieldProductName() string`

GetLogFieldProductName returns the LogFieldProductName field if non-nil, zero value otherwise.

### GetLogFieldProductNameOk

`func (o *AddLogFieldMappingRequest) GetLogFieldProductNameOk() (*string, bool)`

GetLogFieldProductNameOk returns a tuple with the LogFieldProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProductName

`func (o *AddLogFieldMappingRequest) SetLogFieldProductName(v string)`

SetLogFieldProductName sets LogFieldProductName field to given value.

### HasLogFieldProductName

`func (o *AddLogFieldMappingRequest) HasLogFieldProductName() bool`

HasLogFieldProductName returns a boolean if a field has been set.

### GetLogFieldInstanceName

`func (o *AddLogFieldMappingRequest) GetLogFieldInstanceName() string`

GetLogFieldInstanceName returns the LogFieldInstanceName field if non-nil, zero value otherwise.

### GetLogFieldInstanceNameOk

`func (o *AddLogFieldMappingRequest) GetLogFieldInstanceNameOk() (*string, bool)`

GetLogFieldInstanceNameOk returns a tuple with the LogFieldInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldInstanceName

`func (o *AddLogFieldMappingRequest) SetLogFieldInstanceName(v string)`

SetLogFieldInstanceName sets LogFieldInstanceName field to given value.

### HasLogFieldInstanceName

`func (o *AddLogFieldMappingRequest) HasLogFieldInstanceName() bool`

HasLogFieldInstanceName returns a boolean if a field has been set.

### GetLogFieldOperationID

`func (o *AddLogFieldMappingRequest) GetLogFieldOperationID() string`

GetLogFieldOperationID returns the LogFieldOperationID field if non-nil, zero value otherwise.

### GetLogFieldOperationIDOk

`func (o *AddLogFieldMappingRequest) GetLogFieldOperationIDOk() (*string, bool)`

GetLogFieldOperationIDOk returns a tuple with the LogFieldOperationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldOperationID

`func (o *AddLogFieldMappingRequest) SetLogFieldOperationID(v string)`

SetLogFieldOperationID sets LogFieldOperationID field to given value.

### HasLogFieldOperationID

`func (o *AddLogFieldMappingRequest) HasLogFieldOperationID() bool`

HasLogFieldOperationID returns a boolean if a field has been set.

### GetLogFieldMessageType

`func (o *AddLogFieldMappingRequest) GetLogFieldMessageType() string`

GetLogFieldMessageType returns the LogFieldMessageType field if non-nil, zero value otherwise.

### GetLogFieldMessageTypeOk

`func (o *AddLogFieldMappingRequest) GetLogFieldMessageTypeOk() (*string, bool)`

GetLogFieldMessageTypeOk returns a tuple with the LogFieldMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageType

`func (o *AddLogFieldMappingRequest) SetLogFieldMessageType(v string)`

SetLogFieldMessageType sets LogFieldMessageType field to given value.

### HasLogFieldMessageType

`func (o *AddLogFieldMappingRequest) HasLogFieldMessageType() bool`

HasLogFieldMessageType returns a boolean if a field has been set.

### GetLogFieldOperationType

`func (o *AddLogFieldMappingRequest) GetLogFieldOperationType() string`

GetLogFieldOperationType returns the LogFieldOperationType field if non-nil, zero value otherwise.

### GetLogFieldOperationTypeOk

`func (o *AddLogFieldMappingRequest) GetLogFieldOperationTypeOk() (*string, bool)`

GetLogFieldOperationTypeOk returns a tuple with the LogFieldOperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldOperationType

`func (o *AddLogFieldMappingRequest) SetLogFieldOperationType(v string)`

SetLogFieldOperationType sets LogFieldOperationType field to given value.

### HasLogFieldOperationType

`func (o *AddLogFieldMappingRequest) HasLogFieldOperationType() bool`

HasLogFieldOperationType returns a boolean if a field has been set.

### GetLogFieldMessageID

`func (o *AddLogFieldMappingRequest) GetLogFieldMessageID() string`

GetLogFieldMessageID returns the LogFieldMessageID field if non-nil, zero value otherwise.

### GetLogFieldMessageIDOk

`func (o *AddLogFieldMappingRequest) GetLogFieldMessageIDOk() (*string, bool)`

GetLogFieldMessageIDOk returns a tuple with the LogFieldMessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageID

`func (o *AddLogFieldMappingRequest) SetLogFieldMessageID(v string)`

SetLogFieldMessageID sets LogFieldMessageID field to given value.

### HasLogFieldMessageID

`func (o *AddLogFieldMappingRequest) HasLogFieldMessageID() bool`

HasLogFieldMessageID returns a boolean if a field has been set.

### GetLogFieldResultCode

`func (o *AddLogFieldMappingRequest) GetLogFieldResultCode() string`

GetLogFieldResultCode returns the LogFieldResultCode field if non-nil, zero value otherwise.

### GetLogFieldResultCodeOk

`func (o *AddLogFieldMappingRequest) GetLogFieldResultCodeOk() (*string, bool)`

GetLogFieldResultCodeOk returns a tuple with the LogFieldResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldResultCode

`func (o *AddLogFieldMappingRequest) SetLogFieldResultCode(v string)`

SetLogFieldResultCode sets LogFieldResultCode field to given value.

### HasLogFieldResultCode

`func (o *AddLogFieldMappingRequest) HasLogFieldResultCode() bool`

HasLogFieldResultCode returns a boolean if a field has been set.

### GetLogFieldMessage

`func (o *AddLogFieldMappingRequest) GetLogFieldMessage() string`

GetLogFieldMessage returns the LogFieldMessage field if non-nil, zero value otherwise.

### GetLogFieldMessageOk

`func (o *AddLogFieldMappingRequest) GetLogFieldMessageOk() (*string, bool)`

GetLogFieldMessageOk returns a tuple with the LogFieldMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessage

`func (o *AddLogFieldMappingRequest) SetLogFieldMessage(v string)`

SetLogFieldMessage sets LogFieldMessage field to given value.

### HasLogFieldMessage

`func (o *AddLogFieldMappingRequest) HasLogFieldMessage() bool`

HasLogFieldMessage returns a boolean if a field has been set.

### GetLogFieldOrigin

`func (o *AddLogFieldMappingRequest) GetLogFieldOrigin() string`

GetLogFieldOrigin returns the LogFieldOrigin field if non-nil, zero value otherwise.

### GetLogFieldOriginOk

`func (o *AddLogFieldMappingRequest) GetLogFieldOriginOk() (*string, bool)`

GetLogFieldOriginOk returns a tuple with the LogFieldOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldOrigin

`func (o *AddLogFieldMappingRequest) SetLogFieldOrigin(v string)`

SetLogFieldOrigin sets LogFieldOrigin field to given value.

### HasLogFieldOrigin

`func (o *AddLogFieldMappingRequest) HasLogFieldOrigin() bool`

HasLogFieldOrigin returns a boolean if a field has been set.

### GetLogFieldRequesterDN

`func (o *AddLogFieldMappingRequest) GetLogFieldRequesterDN() string`

GetLogFieldRequesterDN returns the LogFieldRequesterDN field if non-nil, zero value otherwise.

### GetLogFieldRequesterDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldRequesterDNOk() (*string, bool)`

GetLogFieldRequesterDNOk returns a tuple with the LogFieldRequesterDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequesterDN

`func (o *AddLogFieldMappingRequest) SetLogFieldRequesterDN(v string)`

SetLogFieldRequesterDN sets LogFieldRequesterDN field to given value.

### HasLogFieldRequesterDN

`func (o *AddLogFieldMappingRequest) HasLogFieldRequesterDN() bool`

HasLogFieldRequesterDN returns a boolean if a field has been set.

### GetLogFieldDisconnectReason

`func (o *AddLogFieldMappingRequest) GetLogFieldDisconnectReason() string`

GetLogFieldDisconnectReason returns the LogFieldDisconnectReason field if non-nil, zero value otherwise.

### GetLogFieldDisconnectReasonOk

`func (o *AddLogFieldMappingRequest) GetLogFieldDisconnectReasonOk() (*string, bool)`

GetLogFieldDisconnectReasonOk returns a tuple with the LogFieldDisconnectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldDisconnectReason

`func (o *AddLogFieldMappingRequest) SetLogFieldDisconnectReason(v string)`

SetLogFieldDisconnectReason sets LogFieldDisconnectReason field to given value.

### HasLogFieldDisconnectReason

`func (o *AddLogFieldMappingRequest) HasLogFieldDisconnectReason() bool`

HasLogFieldDisconnectReason returns a boolean if a field has been set.

### GetLogFieldDeleteOldRDN

`func (o *AddLogFieldMappingRequest) GetLogFieldDeleteOldRDN() string`

GetLogFieldDeleteOldRDN returns the LogFieldDeleteOldRDN field if non-nil, zero value otherwise.

### GetLogFieldDeleteOldRDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldDeleteOldRDNOk() (*string, bool)`

GetLogFieldDeleteOldRDNOk returns a tuple with the LogFieldDeleteOldRDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldDeleteOldRDN

`func (o *AddLogFieldMappingRequest) SetLogFieldDeleteOldRDN(v string)`

SetLogFieldDeleteOldRDN sets LogFieldDeleteOldRDN field to given value.

### HasLogFieldDeleteOldRDN

`func (o *AddLogFieldMappingRequest) HasLogFieldDeleteOldRDN() bool`

HasLogFieldDeleteOldRDN returns a boolean if a field has been set.

### GetLogFieldAuthenticatedUserDN

`func (o *AddLogFieldMappingRequest) GetLogFieldAuthenticatedUserDN() string`

GetLogFieldAuthenticatedUserDN returns the LogFieldAuthenticatedUserDN field if non-nil, zero value otherwise.

### GetLogFieldAuthenticatedUserDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldAuthenticatedUserDNOk() (*string, bool)`

GetLogFieldAuthenticatedUserDNOk returns a tuple with the LogFieldAuthenticatedUserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAuthenticatedUserDN

`func (o *AddLogFieldMappingRequest) SetLogFieldAuthenticatedUserDN(v string)`

SetLogFieldAuthenticatedUserDN sets LogFieldAuthenticatedUserDN field to given value.

### HasLogFieldAuthenticatedUserDN

`func (o *AddLogFieldMappingRequest) HasLogFieldAuthenticatedUserDN() bool`

HasLogFieldAuthenticatedUserDN returns a boolean if a field has been set.

### GetLogFieldProcessingTime

`func (o *AddLogFieldMappingRequest) GetLogFieldProcessingTime() string`

GetLogFieldProcessingTime returns the LogFieldProcessingTime field if non-nil, zero value otherwise.

### GetLogFieldProcessingTimeOk

`func (o *AddLogFieldMappingRequest) GetLogFieldProcessingTimeOk() (*string, bool)`

GetLogFieldProcessingTimeOk returns a tuple with the LogFieldProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProcessingTime

`func (o *AddLogFieldMappingRequest) SetLogFieldProcessingTime(v string)`

SetLogFieldProcessingTime sets LogFieldProcessingTime field to given value.

### HasLogFieldProcessingTime

`func (o *AddLogFieldMappingRequest) HasLogFieldProcessingTime() bool`

HasLogFieldProcessingTime returns a boolean if a field has been set.

### GetLogFieldRequestedAttributes

`func (o *AddLogFieldMappingRequest) GetLogFieldRequestedAttributes() string`

GetLogFieldRequestedAttributes returns the LogFieldRequestedAttributes field if non-nil, zero value otherwise.

### GetLogFieldRequestedAttributesOk

`func (o *AddLogFieldMappingRequest) GetLogFieldRequestedAttributesOk() (*string, bool)`

GetLogFieldRequestedAttributesOk returns a tuple with the LogFieldRequestedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequestedAttributes

`func (o *AddLogFieldMappingRequest) SetLogFieldRequestedAttributes(v string)`

SetLogFieldRequestedAttributes sets LogFieldRequestedAttributes field to given value.

### HasLogFieldRequestedAttributes

`func (o *AddLogFieldMappingRequest) HasLogFieldRequestedAttributes() bool`

HasLogFieldRequestedAttributes returns a boolean if a field has been set.

### GetLogFieldSASLMechanismName

`func (o *AddLogFieldMappingRequest) GetLogFieldSASLMechanismName() string`

GetLogFieldSASLMechanismName returns the LogFieldSASLMechanismName field if non-nil, zero value otherwise.

### GetLogFieldSASLMechanismNameOk

`func (o *AddLogFieldMappingRequest) GetLogFieldSASLMechanismNameOk() (*string, bool)`

GetLogFieldSASLMechanismNameOk returns a tuple with the LogFieldSASLMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldSASLMechanismName

`func (o *AddLogFieldMappingRequest) SetLogFieldSASLMechanismName(v string)`

SetLogFieldSASLMechanismName sets LogFieldSASLMechanismName field to given value.

### HasLogFieldSASLMechanismName

`func (o *AddLogFieldMappingRequest) HasLogFieldSASLMechanismName() bool`

HasLogFieldSASLMechanismName returns a boolean if a field has been set.

### GetLogFieldNewRDN

`func (o *AddLogFieldMappingRequest) GetLogFieldNewRDN() string`

GetLogFieldNewRDN returns the LogFieldNewRDN field if non-nil, zero value otherwise.

### GetLogFieldNewRDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldNewRDNOk() (*string, bool)`

GetLogFieldNewRDNOk returns a tuple with the LogFieldNewRDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldNewRDN

`func (o *AddLogFieldMappingRequest) SetLogFieldNewRDN(v string)`

SetLogFieldNewRDN sets LogFieldNewRDN field to given value.

### HasLogFieldNewRDN

`func (o *AddLogFieldMappingRequest) HasLogFieldNewRDN() bool`

HasLogFieldNewRDN returns a boolean if a field has been set.

### GetLogFieldBaseDN

`func (o *AddLogFieldMappingRequest) GetLogFieldBaseDN() string`

GetLogFieldBaseDN returns the LogFieldBaseDN field if non-nil, zero value otherwise.

### GetLogFieldBaseDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldBaseDNOk() (*string, bool)`

GetLogFieldBaseDNOk returns a tuple with the LogFieldBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldBaseDN

`func (o *AddLogFieldMappingRequest) SetLogFieldBaseDN(v string)`

SetLogFieldBaseDN sets LogFieldBaseDN field to given value.

### HasLogFieldBaseDN

`func (o *AddLogFieldMappingRequest) HasLogFieldBaseDN() bool`

HasLogFieldBaseDN returns a boolean if a field has been set.

### GetLogFieldBindDN

`func (o *AddLogFieldMappingRequest) GetLogFieldBindDN() string`

GetLogFieldBindDN returns the LogFieldBindDN field if non-nil, zero value otherwise.

### GetLogFieldBindDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldBindDNOk() (*string, bool)`

GetLogFieldBindDNOk returns a tuple with the LogFieldBindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldBindDN

`func (o *AddLogFieldMappingRequest) SetLogFieldBindDN(v string)`

SetLogFieldBindDN sets LogFieldBindDN field to given value.

### HasLogFieldBindDN

`func (o *AddLogFieldMappingRequest) HasLogFieldBindDN() bool`

HasLogFieldBindDN returns a boolean if a field has been set.

### GetLogFieldMatchedDN

`func (o *AddLogFieldMappingRequest) GetLogFieldMatchedDN() string`

GetLogFieldMatchedDN returns the LogFieldMatchedDN field if non-nil, zero value otherwise.

### GetLogFieldMatchedDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldMatchedDNOk() (*string, bool)`

GetLogFieldMatchedDNOk returns a tuple with the LogFieldMatchedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMatchedDN

`func (o *AddLogFieldMappingRequest) SetLogFieldMatchedDN(v string)`

SetLogFieldMatchedDN sets LogFieldMatchedDN field to given value.

### HasLogFieldMatchedDN

`func (o *AddLogFieldMappingRequest) HasLogFieldMatchedDN() bool`

HasLogFieldMatchedDN returns a boolean if a field has been set.

### GetLogFieldRequesterIPAddress

`func (o *AddLogFieldMappingRequest) GetLogFieldRequesterIPAddress() string`

GetLogFieldRequesterIPAddress returns the LogFieldRequesterIPAddress field if non-nil, zero value otherwise.

### GetLogFieldRequesterIPAddressOk

`func (o *AddLogFieldMappingRequest) GetLogFieldRequesterIPAddressOk() (*string, bool)`

GetLogFieldRequesterIPAddressOk returns a tuple with the LogFieldRequesterIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequesterIPAddress

`func (o *AddLogFieldMappingRequest) SetLogFieldRequesterIPAddress(v string)`

SetLogFieldRequesterIPAddress sets LogFieldRequesterIPAddress field to given value.

### HasLogFieldRequesterIPAddress

`func (o *AddLogFieldMappingRequest) HasLogFieldRequesterIPAddress() bool`

HasLogFieldRequesterIPAddress returns a boolean if a field has been set.

### GetLogFieldAuthenticationType

`func (o *AddLogFieldMappingRequest) GetLogFieldAuthenticationType() string`

GetLogFieldAuthenticationType returns the LogFieldAuthenticationType field if non-nil, zero value otherwise.

### GetLogFieldAuthenticationTypeOk

`func (o *AddLogFieldMappingRequest) GetLogFieldAuthenticationTypeOk() (*string, bool)`

GetLogFieldAuthenticationTypeOk returns a tuple with the LogFieldAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAuthenticationType

`func (o *AddLogFieldMappingRequest) SetLogFieldAuthenticationType(v string)`

SetLogFieldAuthenticationType sets LogFieldAuthenticationType field to given value.

### HasLogFieldAuthenticationType

`func (o *AddLogFieldMappingRequest) HasLogFieldAuthenticationType() bool`

HasLogFieldAuthenticationType returns a boolean if a field has been set.

### GetLogFieldNewSuperiorDN

`func (o *AddLogFieldMappingRequest) GetLogFieldNewSuperiorDN() string`

GetLogFieldNewSuperiorDN returns the LogFieldNewSuperiorDN field if non-nil, zero value otherwise.

### GetLogFieldNewSuperiorDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldNewSuperiorDNOk() (*string, bool)`

GetLogFieldNewSuperiorDNOk returns a tuple with the LogFieldNewSuperiorDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldNewSuperiorDN

`func (o *AddLogFieldMappingRequest) SetLogFieldNewSuperiorDN(v string)`

SetLogFieldNewSuperiorDN sets LogFieldNewSuperiorDN field to given value.

### HasLogFieldNewSuperiorDN

`func (o *AddLogFieldMappingRequest) HasLogFieldNewSuperiorDN() bool`

HasLogFieldNewSuperiorDN returns a boolean if a field has been set.

### GetLogFieldFilter

`func (o *AddLogFieldMappingRequest) GetLogFieldFilter() string`

GetLogFieldFilter returns the LogFieldFilter field if non-nil, zero value otherwise.

### GetLogFieldFilterOk

`func (o *AddLogFieldMappingRequest) GetLogFieldFilterOk() (*string, bool)`

GetLogFieldFilterOk returns a tuple with the LogFieldFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldFilter

`func (o *AddLogFieldMappingRequest) SetLogFieldFilter(v string)`

SetLogFieldFilter sets LogFieldFilter field to given value.

### HasLogFieldFilter

`func (o *AddLogFieldMappingRequest) HasLogFieldFilter() bool`

HasLogFieldFilter returns a boolean if a field has been set.

### GetLogFieldAlternateAuthorizationDN

`func (o *AddLogFieldMappingRequest) GetLogFieldAlternateAuthorizationDN() string`

GetLogFieldAlternateAuthorizationDN returns the LogFieldAlternateAuthorizationDN field if non-nil, zero value otherwise.

### GetLogFieldAlternateAuthorizationDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldAlternateAuthorizationDNOk() (*string, bool)`

GetLogFieldAlternateAuthorizationDNOk returns a tuple with the LogFieldAlternateAuthorizationDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAlternateAuthorizationDN

`func (o *AddLogFieldMappingRequest) SetLogFieldAlternateAuthorizationDN(v string)`

SetLogFieldAlternateAuthorizationDN sets LogFieldAlternateAuthorizationDN field to given value.

### HasLogFieldAlternateAuthorizationDN

`func (o *AddLogFieldMappingRequest) HasLogFieldAlternateAuthorizationDN() bool`

HasLogFieldAlternateAuthorizationDN returns a boolean if a field has been set.

### GetLogFieldEntryDN

`func (o *AddLogFieldMappingRequest) GetLogFieldEntryDN() string`

GetLogFieldEntryDN returns the LogFieldEntryDN field if non-nil, zero value otherwise.

### GetLogFieldEntryDNOk

`func (o *AddLogFieldMappingRequest) GetLogFieldEntryDNOk() (*string, bool)`

GetLogFieldEntryDNOk returns a tuple with the LogFieldEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldEntryDN

`func (o *AddLogFieldMappingRequest) SetLogFieldEntryDN(v string)`

SetLogFieldEntryDN sets LogFieldEntryDN field to given value.

### HasLogFieldEntryDN

`func (o *AddLogFieldMappingRequest) HasLogFieldEntryDN() bool`

HasLogFieldEntryDN returns a boolean if a field has been set.

### GetLogFieldEntriesReturned

`func (o *AddLogFieldMappingRequest) GetLogFieldEntriesReturned() string`

GetLogFieldEntriesReturned returns the LogFieldEntriesReturned field if non-nil, zero value otherwise.

### GetLogFieldEntriesReturnedOk

`func (o *AddLogFieldMappingRequest) GetLogFieldEntriesReturnedOk() (*string, bool)`

GetLogFieldEntriesReturnedOk returns a tuple with the LogFieldEntriesReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldEntriesReturned

`func (o *AddLogFieldMappingRequest) SetLogFieldEntriesReturned(v string)`

SetLogFieldEntriesReturned sets LogFieldEntriesReturned field to given value.

### HasLogFieldEntriesReturned

`func (o *AddLogFieldMappingRequest) HasLogFieldEntriesReturned() bool`

HasLogFieldEntriesReturned returns a boolean if a field has been set.

### GetLogFieldAuthenticationFailureID

`func (o *AddLogFieldMappingRequest) GetLogFieldAuthenticationFailureID() string`

GetLogFieldAuthenticationFailureID returns the LogFieldAuthenticationFailureID field if non-nil, zero value otherwise.

### GetLogFieldAuthenticationFailureIDOk

`func (o *AddLogFieldMappingRequest) GetLogFieldAuthenticationFailureIDOk() (*string, bool)`

GetLogFieldAuthenticationFailureIDOk returns a tuple with the LogFieldAuthenticationFailureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAuthenticationFailureID

`func (o *AddLogFieldMappingRequest) SetLogFieldAuthenticationFailureID(v string)`

SetLogFieldAuthenticationFailureID sets LogFieldAuthenticationFailureID field to given value.

### HasLogFieldAuthenticationFailureID

`func (o *AddLogFieldMappingRequest) HasLogFieldAuthenticationFailureID() bool`

HasLogFieldAuthenticationFailureID returns a boolean if a field has been set.

### GetLogFieldRequestOID

`func (o *AddLogFieldMappingRequest) GetLogFieldRequestOID() string`

GetLogFieldRequestOID returns the LogFieldRequestOID field if non-nil, zero value otherwise.

### GetLogFieldRequestOIDOk

`func (o *AddLogFieldMappingRequest) GetLogFieldRequestOIDOk() (*string, bool)`

GetLogFieldRequestOIDOk returns a tuple with the LogFieldRequestOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequestOID

`func (o *AddLogFieldMappingRequest) SetLogFieldRequestOID(v string)`

SetLogFieldRequestOID sets LogFieldRequestOID field to given value.

### HasLogFieldRequestOID

`func (o *AddLogFieldMappingRequest) HasLogFieldRequestOID() bool`

HasLogFieldRequestOID returns a boolean if a field has been set.

### GetLogFieldResponseOID

`func (o *AddLogFieldMappingRequest) GetLogFieldResponseOID() string`

GetLogFieldResponseOID returns the LogFieldResponseOID field if non-nil, zero value otherwise.

### GetLogFieldResponseOIDOk

`func (o *AddLogFieldMappingRequest) GetLogFieldResponseOIDOk() (*string, bool)`

GetLogFieldResponseOIDOk returns a tuple with the LogFieldResponseOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldResponseOID

`func (o *AddLogFieldMappingRequest) SetLogFieldResponseOID(v string)`

SetLogFieldResponseOID sets LogFieldResponseOID field to given value.

### HasLogFieldResponseOID

`func (o *AddLogFieldMappingRequest) HasLogFieldResponseOID() bool`

HasLogFieldResponseOID returns a boolean if a field has been set.

### GetLogFieldTargetProtocol

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetProtocol() string`

GetLogFieldTargetProtocol returns the LogFieldTargetProtocol field if non-nil, zero value otherwise.

### GetLogFieldTargetProtocolOk

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetProtocolOk() (*string, bool)`

GetLogFieldTargetProtocolOk returns a tuple with the LogFieldTargetProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetProtocol

`func (o *AddLogFieldMappingRequest) SetLogFieldTargetProtocol(v string)`

SetLogFieldTargetProtocol sets LogFieldTargetProtocol field to given value.

### HasLogFieldTargetProtocol

`func (o *AddLogFieldMappingRequest) HasLogFieldTargetProtocol() bool`

HasLogFieldTargetProtocol returns a boolean if a field has been set.

### GetLogFieldTargetPort

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetPort() string`

GetLogFieldTargetPort returns the LogFieldTargetPort field if non-nil, zero value otherwise.

### GetLogFieldTargetPortOk

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetPortOk() (*string, bool)`

GetLogFieldTargetPortOk returns a tuple with the LogFieldTargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetPort

`func (o *AddLogFieldMappingRequest) SetLogFieldTargetPort(v string)`

SetLogFieldTargetPort sets LogFieldTargetPort field to given value.

### HasLogFieldTargetPort

`func (o *AddLogFieldMappingRequest) HasLogFieldTargetPort() bool`

HasLogFieldTargetPort returns a boolean if a field has been set.

### GetLogFieldTargetAddress

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetAddress() string`

GetLogFieldTargetAddress returns the LogFieldTargetAddress field if non-nil, zero value otherwise.

### GetLogFieldTargetAddressOk

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetAddressOk() (*string, bool)`

GetLogFieldTargetAddressOk returns a tuple with the LogFieldTargetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetAddress

`func (o *AddLogFieldMappingRequest) SetLogFieldTargetAddress(v string)`

SetLogFieldTargetAddress sets LogFieldTargetAddress field to given value.

### HasLogFieldTargetAddress

`func (o *AddLogFieldMappingRequest) HasLogFieldTargetAddress() bool`

HasLogFieldTargetAddress returns a boolean if a field has been set.

### GetLogFieldTargetAttribute

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetAttribute() string`

GetLogFieldTargetAttribute returns the LogFieldTargetAttribute field if non-nil, zero value otherwise.

### GetLogFieldTargetAttributeOk

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetAttributeOk() (*string, bool)`

GetLogFieldTargetAttributeOk returns a tuple with the LogFieldTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetAttribute

`func (o *AddLogFieldMappingRequest) SetLogFieldTargetAttribute(v string)`

SetLogFieldTargetAttribute sets LogFieldTargetAttribute field to given value.

### HasLogFieldTargetAttribute

`func (o *AddLogFieldMappingRequest) HasLogFieldTargetAttribute() bool`

HasLogFieldTargetAttribute returns a boolean if a field has been set.

### GetLogFieldTargetHost

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetHost() string`

GetLogFieldTargetHost returns the LogFieldTargetHost field if non-nil, zero value otherwise.

### GetLogFieldTargetHostOk

`func (o *AddLogFieldMappingRequest) GetLogFieldTargetHostOk() (*string, bool)`

GetLogFieldTargetHostOk returns a tuple with the LogFieldTargetHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetHost

`func (o *AddLogFieldMappingRequest) SetLogFieldTargetHost(v string)`

SetLogFieldTargetHost sets LogFieldTargetHost field to given value.

### HasLogFieldTargetHost

`func (o *AddLogFieldMappingRequest) HasLogFieldTargetHost() bool`

HasLogFieldTargetHost returns a boolean if a field has been set.

### GetLogFieldProtocolVersion

`func (o *AddLogFieldMappingRequest) GetLogFieldProtocolVersion() string`

GetLogFieldProtocolVersion returns the LogFieldProtocolVersion field if non-nil, zero value otherwise.

### GetLogFieldProtocolVersionOk

`func (o *AddLogFieldMappingRequest) GetLogFieldProtocolVersionOk() (*string, bool)`

GetLogFieldProtocolVersionOk returns a tuple with the LogFieldProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProtocolVersion

`func (o *AddLogFieldMappingRequest) SetLogFieldProtocolVersion(v string)`

SetLogFieldProtocolVersion sets LogFieldProtocolVersion field to given value.

### HasLogFieldProtocolVersion

`func (o *AddLogFieldMappingRequest) HasLogFieldProtocolVersion() bool`

HasLogFieldProtocolVersion returns a boolean if a field has been set.

### GetLogFieldProtocolName

`func (o *AddLogFieldMappingRequest) GetLogFieldProtocolName() string`

GetLogFieldProtocolName returns the LogFieldProtocolName field if non-nil, zero value otherwise.

### GetLogFieldProtocolNameOk

`func (o *AddLogFieldMappingRequest) GetLogFieldProtocolNameOk() (*string, bool)`

GetLogFieldProtocolNameOk returns a tuple with the LogFieldProtocolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProtocolName

`func (o *AddLogFieldMappingRequest) SetLogFieldProtocolName(v string)`

SetLogFieldProtocolName sets LogFieldProtocolName field to given value.

### HasLogFieldProtocolName

`func (o *AddLogFieldMappingRequest) HasLogFieldProtocolName() bool`

HasLogFieldProtocolName returns a boolean if a field has been set.

### GetLogFieldAuthenticationFailureReason

`func (o *AddLogFieldMappingRequest) GetLogFieldAuthenticationFailureReason() string`

GetLogFieldAuthenticationFailureReason returns the LogFieldAuthenticationFailureReason field if non-nil, zero value otherwise.

### GetLogFieldAuthenticationFailureReasonOk

`func (o *AddLogFieldMappingRequest) GetLogFieldAuthenticationFailureReasonOk() (*string, bool)`

GetLogFieldAuthenticationFailureReasonOk returns a tuple with the LogFieldAuthenticationFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAuthenticationFailureReason

`func (o *AddLogFieldMappingRequest) SetLogFieldAuthenticationFailureReason(v string)`

SetLogFieldAuthenticationFailureReason sets LogFieldAuthenticationFailureReason field to given value.

### HasLogFieldAuthenticationFailureReason

`func (o *AddLogFieldMappingRequest) HasLogFieldAuthenticationFailureReason() bool`

HasLogFieldAuthenticationFailureReason returns a boolean if a field has been set.

### GetLogFieldAdditionalInformation

`func (o *AddLogFieldMappingRequest) GetLogFieldAdditionalInformation() string`

GetLogFieldAdditionalInformation returns the LogFieldAdditionalInformation field if non-nil, zero value otherwise.

### GetLogFieldAdditionalInformationOk

`func (o *AddLogFieldMappingRequest) GetLogFieldAdditionalInformationOk() (*string, bool)`

GetLogFieldAdditionalInformationOk returns a tuple with the LogFieldAdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAdditionalInformation

`func (o *AddLogFieldMappingRequest) SetLogFieldAdditionalInformation(v string)`

SetLogFieldAdditionalInformation sets LogFieldAdditionalInformation field to given value.

### HasLogFieldAdditionalInformation

`func (o *AddLogFieldMappingRequest) HasLogFieldAdditionalInformation() bool`

HasLogFieldAdditionalInformation returns a boolean if a field has been set.

### GetLogFieldUnindexed

`func (o *AddLogFieldMappingRequest) GetLogFieldUnindexed() string`

GetLogFieldUnindexed returns the LogFieldUnindexed field if non-nil, zero value otherwise.

### GetLogFieldUnindexedOk

`func (o *AddLogFieldMappingRequest) GetLogFieldUnindexedOk() (*string, bool)`

GetLogFieldUnindexedOk returns a tuple with the LogFieldUnindexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldUnindexed

`func (o *AddLogFieldMappingRequest) SetLogFieldUnindexed(v string)`

SetLogFieldUnindexed sets LogFieldUnindexed field to given value.

### HasLogFieldUnindexed

`func (o *AddLogFieldMappingRequest) HasLogFieldUnindexed() bool`

HasLogFieldUnindexed returns a boolean if a field has been set.

### GetLogFieldScope

`func (o *AddLogFieldMappingRequest) GetLogFieldScope() string`

GetLogFieldScope returns the LogFieldScope field if non-nil, zero value otherwise.

### GetLogFieldScopeOk

`func (o *AddLogFieldMappingRequest) GetLogFieldScopeOk() (*string, bool)`

GetLogFieldScopeOk returns a tuple with the LogFieldScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldScope

`func (o *AddLogFieldMappingRequest) SetLogFieldScope(v string)`

SetLogFieldScope sets LogFieldScope field to given value.

### HasLogFieldScope

`func (o *AddLogFieldMappingRequest) HasLogFieldScope() bool`

HasLogFieldScope returns a boolean if a field has been set.

### GetLogFieldReferralUrls

`func (o *AddLogFieldMappingRequest) GetLogFieldReferralUrls() string`

GetLogFieldReferralUrls returns the LogFieldReferralUrls field if non-nil, zero value otherwise.

### GetLogFieldReferralUrlsOk

`func (o *AddLogFieldMappingRequest) GetLogFieldReferralUrlsOk() (*string, bool)`

GetLogFieldReferralUrlsOk returns a tuple with the LogFieldReferralUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldReferralUrls

`func (o *AddLogFieldMappingRequest) SetLogFieldReferralUrls(v string)`

SetLogFieldReferralUrls sets LogFieldReferralUrls field to given value.

### HasLogFieldReferralUrls

`func (o *AddLogFieldMappingRequest) HasLogFieldReferralUrls() bool`

HasLogFieldReferralUrls returns a boolean if a field has been set.

### GetLogFieldSourceAddress

`func (o *AddLogFieldMappingRequest) GetLogFieldSourceAddress() string`

GetLogFieldSourceAddress returns the LogFieldSourceAddress field if non-nil, zero value otherwise.

### GetLogFieldSourceAddressOk

`func (o *AddLogFieldMappingRequest) GetLogFieldSourceAddressOk() (*string, bool)`

GetLogFieldSourceAddressOk returns a tuple with the LogFieldSourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldSourceAddress

`func (o *AddLogFieldMappingRequest) SetLogFieldSourceAddress(v string)`

SetLogFieldSourceAddress sets LogFieldSourceAddress field to given value.

### HasLogFieldSourceAddress

`func (o *AddLogFieldMappingRequest) HasLogFieldSourceAddress() bool`

HasLogFieldSourceAddress returns a boolean if a field has been set.

### GetLogFieldMessageIDToAbandon

`func (o *AddLogFieldMappingRequest) GetLogFieldMessageIDToAbandon() string`

GetLogFieldMessageIDToAbandon returns the LogFieldMessageIDToAbandon field if non-nil, zero value otherwise.

### GetLogFieldMessageIDToAbandonOk

`func (o *AddLogFieldMappingRequest) GetLogFieldMessageIDToAbandonOk() (*string, bool)`

GetLogFieldMessageIDToAbandonOk returns a tuple with the LogFieldMessageIDToAbandon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageIDToAbandon

`func (o *AddLogFieldMappingRequest) SetLogFieldMessageIDToAbandon(v string)`

SetLogFieldMessageIDToAbandon sets LogFieldMessageIDToAbandon field to given value.

### HasLogFieldMessageIDToAbandon

`func (o *AddLogFieldMappingRequest) HasLogFieldMessageIDToAbandon() bool`

HasLogFieldMessageIDToAbandon returns a boolean if a field has been set.

### GetLogFieldResponseControls

`func (o *AddLogFieldMappingRequest) GetLogFieldResponseControls() string`

GetLogFieldResponseControls returns the LogFieldResponseControls field if non-nil, zero value otherwise.

### GetLogFieldResponseControlsOk

`func (o *AddLogFieldMappingRequest) GetLogFieldResponseControlsOk() (*string, bool)`

GetLogFieldResponseControlsOk returns a tuple with the LogFieldResponseControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldResponseControls

`func (o *AddLogFieldMappingRequest) SetLogFieldResponseControls(v string)`

SetLogFieldResponseControls sets LogFieldResponseControls field to given value.

### HasLogFieldResponseControls

`func (o *AddLogFieldMappingRequest) HasLogFieldResponseControls() bool`

HasLogFieldResponseControls returns a boolean if a field has been set.

### GetLogFieldRequestControls

`func (o *AddLogFieldMappingRequest) GetLogFieldRequestControls() string`

GetLogFieldRequestControls returns the LogFieldRequestControls field if non-nil, zero value otherwise.

### GetLogFieldRequestControlsOk

`func (o *AddLogFieldMappingRequest) GetLogFieldRequestControlsOk() (*string, bool)`

GetLogFieldRequestControlsOk returns a tuple with the LogFieldRequestControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequestControls

`func (o *AddLogFieldMappingRequest) SetLogFieldRequestControls(v string)`

SetLogFieldRequestControls sets LogFieldRequestControls field to given value.

### HasLogFieldRequestControls

`func (o *AddLogFieldMappingRequest) HasLogFieldRequestControls() bool`

HasLogFieldRequestControls returns a boolean if a field has been set.

### GetLogFieldIntermediateClientResult

`func (o *AddLogFieldMappingRequest) GetLogFieldIntermediateClientResult() string`

GetLogFieldIntermediateClientResult returns the LogFieldIntermediateClientResult field if non-nil, zero value otherwise.

### GetLogFieldIntermediateClientResultOk

`func (o *AddLogFieldMappingRequest) GetLogFieldIntermediateClientResultOk() (*string, bool)`

GetLogFieldIntermediateClientResultOk returns a tuple with the LogFieldIntermediateClientResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldIntermediateClientResult

`func (o *AddLogFieldMappingRequest) SetLogFieldIntermediateClientResult(v string)`

SetLogFieldIntermediateClientResult sets LogFieldIntermediateClientResult field to given value.

### HasLogFieldIntermediateClientResult

`func (o *AddLogFieldMappingRequest) HasLogFieldIntermediateClientResult() bool`

HasLogFieldIntermediateClientResult returns a boolean if a field has been set.

### GetLogFieldIntermediateClientRequest

`func (o *AddLogFieldMappingRequest) GetLogFieldIntermediateClientRequest() string`

GetLogFieldIntermediateClientRequest returns the LogFieldIntermediateClientRequest field if non-nil, zero value otherwise.

### GetLogFieldIntermediateClientRequestOk

`func (o *AddLogFieldMappingRequest) GetLogFieldIntermediateClientRequestOk() (*string, bool)`

GetLogFieldIntermediateClientRequestOk returns a tuple with the LogFieldIntermediateClientRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldIntermediateClientRequest

`func (o *AddLogFieldMappingRequest) SetLogFieldIntermediateClientRequest(v string)`

SetLogFieldIntermediateClientRequest sets LogFieldIntermediateClientRequest field to given value.

### HasLogFieldIntermediateClientRequest

`func (o *AddLogFieldMappingRequest) HasLogFieldIntermediateClientRequest() bool`

HasLogFieldIntermediateClientRequest returns a boolean if a field has been set.

### GetLogFieldReplicationChangeID

`func (o *AddLogFieldMappingRequest) GetLogFieldReplicationChangeID() string`

GetLogFieldReplicationChangeID returns the LogFieldReplicationChangeID field if non-nil, zero value otherwise.

### GetLogFieldReplicationChangeIDOk

`func (o *AddLogFieldMappingRequest) GetLogFieldReplicationChangeIDOk() (*string, bool)`

GetLogFieldReplicationChangeIDOk returns a tuple with the LogFieldReplicationChangeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldReplicationChangeID

`func (o *AddLogFieldMappingRequest) SetLogFieldReplicationChangeID(v string)`

SetLogFieldReplicationChangeID sets LogFieldReplicationChangeID field to given value.

### HasLogFieldReplicationChangeID

`func (o *AddLogFieldMappingRequest) HasLogFieldReplicationChangeID() bool`

HasLogFieldReplicationChangeID returns a boolean if a field has been set.

### GetDescription

`func (o *AddLogFieldMappingRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLogFieldMappingRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLogFieldMappingRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLogFieldMappingRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLogFieldCategory

`func (o *AddLogFieldMappingRequest) GetLogFieldCategory() string`

GetLogFieldCategory returns the LogFieldCategory field if non-nil, zero value otherwise.

### GetLogFieldCategoryOk

`func (o *AddLogFieldMappingRequest) GetLogFieldCategoryOk() (*string, bool)`

GetLogFieldCategoryOk returns a tuple with the LogFieldCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldCategory

`func (o *AddLogFieldMappingRequest) SetLogFieldCategory(v string)`

SetLogFieldCategory sets LogFieldCategory field to given value.

### HasLogFieldCategory

`func (o *AddLogFieldMappingRequest) HasLogFieldCategory() bool`

HasLogFieldCategory returns a boolean if a field has been set.

### GetLogFieldSeverity

`func (o *AddLogFieldMappingRequest) GetLogFieldSeverity() string`

GetLogFieldSeverity returns the LogFieldSeverity field if non-nil, zero value otherwise.

### GetLogFieldSeverityOk

`func (o *AddLogFieldMappingRequest) GetLogFieldSeverityOk() (*string, bool)`

GetLogFieldSeverityOk returns a tuple with the LogFieldSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldSeverity

`func (o *AddLogFieldMappingRequest) SetLogFieldSeverity(v string)`

SetLogFieldSeverity sets LogFieldSeverity field to given value.

### HasLogFieldSeverity

`func (o *AddLogFieldMappingRequest) HasLogFieldSeverity() bool`

HasLogFieldSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


