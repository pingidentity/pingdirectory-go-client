# AccessLogFieldMappingShared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumaccessLogFieldMappingSchemaUrn**](EnumaccessLogFieldMappingSchemaUrn.md) |  | 
**LogFieldTimestamp** | Pointer to **string** | The time that the operation was processed. | [optional] 
**LogFieldConnectionID** | Pointer to **string** | The connection ID assigned to the client connection. | [optional] 
**LogFieldStartupid** | Pointer to **string** | The startup ID for the Directory Server. A different value will be generated each time the server is started, and it may be used to distinguish between operations with the same connection ID and operation ID across server restarts. | [optional] 
**LogFieldProductName** | Pointer to **string** | The name for this Directory Server product, which may be used to identify which product was used to log the message if multiple products log to the same database table. | [optional] 
**LogFieldInstanceName** | Pointer to **string** | A name that uniquely identifies this Directory Server instance, which may be used to identify which instance was used to log the message if multiple server instances log to the same database table. | [optional] 
**LogFieldOperationID** | Pointer to **string** | The operation ID for the operation processed by the server. | [optional] 
**LogFieldMessageType** | Pointer to **string** | The type of log message. Message types may include \&quot;CONNECT\&quot;, \&quot;DISCONNECT\&quot;, \&quot;FORWARD\&quot;, \&quot;RESULT\&quot;, \&quot;ENTRY\&quot;, or \&quot;REFERENCE\&quot;. | [optional] 
**LogFieldOperationType** | Pointer to **string** | The type of operation that was processed. | [optional] 
**LogFieldMessageID** | Pointer to **string** | The message ID included in the client request. | [optional] 
**LogFieldResultCode** | Pointer to **string** | The numeric result code for the operation. | [optional] 
**LogFieldMessage** | Pointer to **string** | The diagnostic message for the operation. | [optional] 
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

## Methods

### NewAccessLogFieldMappingShared

`func NewAccessLogFieldMappingShared(schemas []EnumaccessLogFieldMappingSchemaUrn, ) *AccessLogFieldMappingShared`

NewAccessLogFieldMappingShared instantiates a new AccessLogFieldMappingShared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessLogFieldMappingSharedWithDefaults

`func NewAccessLogFieldMappingSharedWithDefaults() *AccessLogFieldMappingShared`

NewAccessLogFieldMappingSharedWithDefaults instantiates a new AccessLogFieldMappingShared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AccessLogFieldMappingShared) GetSchemas() []EnumaccessLogFieldMappingSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AccessLogFieldMappingShared) GetSchemasOk() (*[]EnumaccessLogFieldMappingSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AccessLogFieldMappingShared) SetSchemas(v []EnumaccessLogFieldMappingSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetLogFieldTimestamp

`func (o *AccessLogFieldMappingShared) GetLogFieldTimestamp() string`

GetLogFieldTimestamp returns the LogFieldTimestamp field if non-nil, zero value otherwise.

### GetLogFieldTimestampOk

`func (o *AccessLogFieldMappingShared) GetLogFieldTimestampOk() (*string, bool)`

GetLogFieldTimestampOk returns a tuple with the LogFieldTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTimestamp

`func (o *AccessLogFieldMappingShared) SetLogFieldTimestamp(v string)`

SetLogFieldTimestamp sets LogFieldTimestamp field to given value.

### HasLogFieldTimestamp

`func (o *AccessLogFieldMappingShared) HasLogFieldTimestamp() bool`

HasLogFieldTimestamp returns a boolean if a field has been set.

### GetLogFieldConnectionID

`func (o *AccessLogFieldMappingShared) GetLogFieldConnectionID() string`

GetLogFieldConnectionID returns the LogFieldConnectionID field if non-nil, zero value otherwise.

### GetLogFieldConnectionIDOk

`func (o *AccessLogFieldMappingShared) GetLogFieldConnectionIDOk() (*string, bool)`

GetLogFieldConnectionIDOk returns a tuple with the LogFieldConnectionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldConnectionID

`func (o *AccessLogFieldMappingShared) SetLogFieldConnectionID(v string)`

SetLogFieldConnectionID sets LogFieldConnectionID field to given value.

### HasLogFieldConnectionID

`func (o *AccessLogFieldMappingShared) HasLogFieldConnectionID() bool`

HasLogFieldConnectionID returns a boolean if a field has been set.

### GetLogFieldStartupid

`func (o *AccessLogFieldMappingShared) GetLogFieldStartupid() string`

GetLogFieldStartupid returns the LogFieldStartupid field if non-nil, zero value otherwise.

### GetLogFieldStartupidOk

`func (o *AccessLogFieldMappingShared) GetLogFieldStartupidOk() (*string, bool)`

GetLogFieldStartupidOk returns a tuple with the LogFieldStartupid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldStartupid

`func (o *AccessLogFieldMappingShared) SetLogFieldStartupid(v string)`

SetLogFieldStartupid sets LogFieldStartupid field to given value.

### HasLogFieldStartupid

`func (o *AccessLogFieldMappingShared) HasLogFieldStartupid() bool`

HasLogFieldStartupid returns a boolean if a field has been set.

### GetLogFieldProductName

`func (o *AccessLogFieldMappingShared) GetLogFieldProductName() string`

GetLogFieldProductName returns the LogFieldProductName field if non-nil, zero value otherwise.

### GetLogFieldProductNameOk

`func (o *AccessLogFieldMappingShared) GetLogFieldProductNameOk() (*string, bool)`

GetLogFieldProductNameOk returns a tuple with the LogFieldProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProductName

`func (o *AccessLogFieldMappingShared) SetLogFieldProductName(v string)`

SetLogFieldProductName sets LogFieldProductName field to given value.

### HasLogFieldProductName

`func (o *AccessLogFieldMappingShared) HasLogFieldProductName() bool`

HasLogFieldProductName returns a boolean if a field has been set.

### GetLogFieldInstanceName

`func (o *AccessLogFieldMappingShared) GetLogFieldInstanceName() string`

GetLogFieldInstanceName returns the LogFieldInstanceName field if non-nil, zero value otherwise.

### GetLogFieldInstanceNameOk

`func (o *AccessLogFieldMappingShared) GetLogFieldInstanceNameOk() (*string, bool)`

GetLogFieldInstanceNameOk returns a tuple with the LogFieldInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldInstanceName

`func (o *AccessLogFieldMappingShared) SetLogFieldInstanceName(v string)`

SetLogFieldInstanceName sets LogFieldInstanceName field to given value.

### HasLogFieldInstanceName

`func (o *AccessLogFieldMappingShared) HasLogFieldInstanceName() bool`

HasLogFieldInstanceName returns a boolean if a field has been set.

### GetLogFieldOperationID

`func (o *AccessLogFieldMappingShared) GetLogFieldOperationID() string`

GetLogFieldOperationID returns the LogFieldOperationID field if non-nil, zero value otherwise.

### GetLogFieldOperationIDOk

`func (o *AccessLogFieldMappingShared) GetLogFieldOperationIDOk() (*string, bool)`

GetLogFieldOperationIDOk returns a tuple with the LogFieldOperationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldOperationID

`func (o *AccessLogFieldMappingShared) SetLogFieldOperationID(v string)`

SetLogFieldOperationID sets LogFieldOperationID field to given value.

### HasLogFieldOperationID

`func (o *AccessLogFieldMappingShared) HasLogFieldOperationID() bool`

HasLogFieldOperationID returns a boolean if a field has been set.

### GetLogFieldMessageType

`func (o *AccessLogFieldMappingShared) GetLogFieldMessageType() string`

GetLogFieldMessageType returns the LogFieldMessageType field if non-nil, zero value otherwise.

### GetLogFieldMessageTypeOk

`func (o *AccessLogFieldMappingShared) GetLogFieldMessageTypeOk() (*string, bool)`

GetLogFieldMessageTypeOk returns a tuple with the LogFieldMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageType

`func (o *AccessLogFieldMappingShared) SetLogFieldMessageType(v string)`

SetLogFieldMessageType sets LogFieldMessageType field to given value.

### HasLogFieldMessageType

`func (o *AccessLogFieldMappingShared) HasLogFieldMessageType() bool`

HasLogFieldMessageType returns a boolean if a field has been set.

### GetLogFieldOperationType

`func (o *AccessLogFieldMappingShared) GetLogFieldOperationType() string`

GetLogFieldOperationType returns the LogFieldOperationType field if non-nil, zero value otherwise.

### GetLogFieldOperationTypeOk

`func (o *AccessLogFieldMappingShared) GetLogFieldOperationTypeOk() (*string, bool)`

GetLogFieldOperationTypeOk returns a tuple with the LogFieldOperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldOperationType

`func (o *AccessLogFieldMappingShared) SetLogFieldOperationType(v string)`

SetLogFieldOperationType sets LogFieldOperationType field to given value.

### HasLogFieldOperationType

`func (o *AccessLogFieldMappingShared) HasLogFieldOperationType() bool`

HasLogFieldOperationType returns a boolean if a field has been set.

### GetLogFieldMessageID

`func (o *AccessLogFieldMappingShared) GetLogFieldMessageID() string`

GetLogFieldMessageID returns the LogFieldMessageID field if non-nil, zero value otherwise.

### GetLogFieldMessageIDOk

`func (o *AccessLogFieldMappingShared) GetLogFieldMessageIDOk() (*string, bool)`

GetLogFieldMessageIDOk returns a tuple with the LogFieldMessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageID

`func (o *AccessLogFieldMappingShared) SetLogFieldMessageID(v string)`

SetLogFieldMessageID sets LogFieldMessageID field to given value.

### HasLogFieldMessageID

`func (o *AccessLogFieldMappingShared) HasLogFieldMessageID() bool`

HasLogFieldMessageID returns a boolean if a field has been set.

### GetLogFieldResultCode

`func (o *AccessLogFieldMappingShared) GetLogFieldResultCode() string`

GetLogFieldResultCode returns the LogFieldResultCode field if non-nil, zero value otherwise.

### GetLogFieldResultCodeOk

`func (o *AccessLogFieldMappingShared) GetLogFieldResultCodeOk() (*string, bool)`

GetLogFieldResultCodeOk returns a tuple with the LogFieldResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldResultCode

`func (o *AccessLogFieldMappingShared) SetLogFieldResultCode(v string)`

SetLogFieldResultCode sets LogFieldResultCode field to given value.

### HasLogFieldResultCode

`func (o *AccessLogFieldMappingShared) HasLogFieldResultCode() bool`

HasLogFieldResultCode returns a boolean if a field has been set.

### GetLogFieldMessage

`func (o *AccessLogFieldMappingShared) GetLogFieldMessage() string`

GetLogFieldMessage returns the LogFieldMessage field if non-nil, zero value otherwise.

### GetLogFieldMessageOk

`func (o *AccessLogFieldMappingShared) GetLogFieldMessageOk() (*string, bool)`

GetLogFieldMessageOk returns a tuple with the LogFieldMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessage

`func (o *AccessLogFieldMappingShared) SetLogFieldMessage(v string)`

SetLogFieldMessage sets LogFieldMessage field to given value.

### HasLogFieldMessage

`func (o *AccessLogFieldMappingShared) HasLogFieldMessage() bool`

HasLogFieldMessage returns a boolean if a field has been set.

### GetLogFieldOrigin

`func (o *AccessLogFieldMappingShared) GetLogFieldOrigin() string`

GetLogFieldOrigin returns the LogFieldOrigin field if non-nil, zero value otherwise.

### GetLogFieldOriginOk

`func (o *AccessLogFieldMappingShared) GetLogFieldOriginOk() (*string, bool)`

GetLogFieldOriginOk returns a tuple with the LogFieldOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldOrigin

`func (o *AccessLogFieldMappingShared) SetLogFieldOrigin(v string)`

SetLogFieldOrigin sets LogFieldOrigin field to given value.

### HasLogFieldOrigin

`func (o *AccessLogFieldMappingShared) HasLogFieldOrigin() bool`

HasLogFieldOrigin returns a boolean if a field has been set.

### GetLogFieldRequesterDN

`func (o *AccessLogFieldMappingShared) GetLogFieldRequesterDN() string`

GetLogFieldRequesterDN returns the LogFieldRequesterDN field if non-nil, zero value otherwise.

### GetLogFieldRequesterDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldRequesterDNOk() (*string, bool)`

GetLogFieldRequesterDNOk returns a tuple with the LogFieldRequesterDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequesterDN

`func (o *AccessLogFieldMappingShared) SetLogFieldRequesterDN(v string)`

SetLogFieldRequesterDN sets LogFieldRequesterDN field to given value.

### HasLogFieldRequesterDN

`func (o *AccessLogFieldMappingShared) HasLogFieldRequesterDN() bool`

HasLogFieldRequesterDN returns a boolean if a field has been set.

### GetLogFieldDisconnectReason

`func (o *AccessLogFieldMappingShared) GetLogFieldDisconnectReason() string`

GetLogFieldDisconnectReason returns the LogFieldDisconnectReason field if non-nil, zero value otherwise.

### GetLogFieldDisconnectReasonOk

`func (o *AccessLogFieldMappingShared) GetLogFieldDisconnectReasonOk() (*string, bool)`

GetLogFieldDisconnectReasonOk returns a tuple with the LogFieldDisconnectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldDisconnectReason

`func (o *AccessLogFieldMappingShared) SetLogFieldDisconnectReason(v string)`

SetLogFieldDisconnectReason sets LogFieldDisconnectReason field to given value.

### HasLogFieldDisconnectReason

`func (o *AccessLogFieldMappingShared) HasLogFieldDisconnectReason() bool`

HasLogFieldDisconnectReason returns a boolean if a field has been set.

### GetLogFieldDeleteOldRDN

`func (o *AccessLogFieldMappingShared) GetLogFieldDeleteOldRDN() string`

GetLogFieldDeleteOldRDN returns the LogFieldDeleteOldRDN field if non-nil, zero value otherwise.

### GetLogFieldDeleteOldRDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldDeleteOldRDNOk() (*string, bool)`

GetLogFieldDeleteOldRDNOk returns a tuple with the LogFieldDeleteOldRDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldDeleteOldRDN

`func (o *AccessLogFieldMappingShared) SetLogFieldDeleteOldRDN(v string)`

SetLogFieldDeleteOldRDN sets LogFieldDeleteOldRDN field to given value.

### HasLogFieldDeleteOldRDN

`func (o *AccessLogFieldMappingShared) HasLogFieldDeleteOldRDN() bool`

HasLogFieldDeleteOldRDN returns a boolean if a field has been set.

### GetLogFieldAuthenticatedUserDN

`func (o *AccessLogFieldMappingShared) GetLogFieldAuthenticatedUserDN() string`

GetLogFieldAuthenticatedUserDN returns the LogFieldAuthenticatedUserDN field if non-nil, zero value otherwise.

### GetLogFieldAuthenticatedUserDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldAuthenticatedUserDNOk() (*string, bool)`

GetLogFieldAuthenticatedUserDNOk returns a tuple with the LogFieldAuthenticatedUserDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAuthenticatedUserDN

`func (o *AccessLogFieldMappingShared) SetLogFieldAuthenticatedUserDN(v string)`

SetLogFieldAuthenticatedUserDN sets LogFieldAuthenticatedUserDN field to given value.

### HasLogFieldAuthenticatedUserDN

`func (o *AccessLogFieldMappingShared) HasLogFieldAuthenticatedUserDN() bool`

HasLogFieldAuthenticatedUserDN returns a boolean if a field has been set.

### GetLogFieldProcessingTime

`func (o *AccessLogFieldMappingShared) GetLogFieldProcessingTime() string`

GetLogFieldProcessingTime returns the LogFieldProcessingTime field if non-nil, zero value otherwise.

### GetLogFieldProcessingTimeOk

`func (o *AccessLogFieldMappingShared) GetLogFieldProcessingTimeOk() (*string, bool)`

GetLogFieldProcessingTimeOk returns a tuple with the LogFieldProcessingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProcessingTime

`func (o *AccessLogFieldMappingShared) SetLogFieldProcessingTime(v string)`

SetLogFieldProcessingTime sets LogFieldProcessingTime field to given value.

### HasLogFieldProcessingTime

`func (o *AccessLogFieldMappingShared) HasLogFieldProcessingTime() bool`

HasLogFieldProcessingTime returns a boolean if a field has been set.

### GetLogFieldRequestedAttributes

`func (o *AccessLogFieldMappingShared) GetLogFieldRequestedAttributes() string`

GetLogFieldRequestedAttributes returns the LogFieldRequestedAttributes field if non-nil, zero value otherwise.

### GetLogFieldRequestedAttributesOk

`func (o *AccessLogFieldMappingShared) GetLogFieldRequestedAttributesOk() (*string, bool)`

GetLogFieldRequestedAttributesOk returns a tuple with the LogFieldRequestedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequestedAttributes

`func (o *AccessLogFieldMappingShared) SetLogFieldRequestedAttributes(v string)`

SetLogFieldRequestedAttributes sets LogFieldRequestedAttributes field to given value.

### HasLogFieldRequestedAttributes

`func (o *AccessLogFieldMappingShared) HasLogFieldRequestedAttributes() bool`

HasLogFieldRequestedAttributes returns a boolean if a field has been set.

### GetLogFieldSASLMechanismName

`func (o *AccessLogFieldMappingShared) GetLogFieldSASLMechanismName() string`

GetLogFieldSASLMechanismName returns the LogFieldSASLMechanismName field if non-nil, zero value otherwise.

### GetLogFieldSASLMechanismNameOk

`func (o *AccessLogFieldMappingShared) GetLogFieldSASLMechanismNameOk() (*string, bool)`

GetLogFieldSASLMechanismNameOk returns a tuple with the LogFieldSASLMechanismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldSASLMechanismName

`func (o *AccessLogFieldMappingShared) SetLogFieldSASLMechanismName(v string)`

SetLogFieldSASLMechanismName sets LogFieldSASLMechanismName field to given value.

### HasLogFieldSASLMechanismName

`func (o *AccessLogFieldMappingShared) HasLogFieldSASLMechanismName() bool`

HasLogFieldSASLMechanismName returns a boolean if a field has been set.

### GetLogFieldNewRDN

`func (o *AccessLogFieldMappingShared) GetLogFieldNewRDN() string`

GetLogFieldNewRDN returns the LogFieldNewRDN field if non-nil, zero value otherwise.

### GetLogFieldNewRDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldNewRDNOk() (*string, bool)`

GetLogFieldNewRDNOk returns a tuple with the LogFieldNewRDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldNewRDN

`func (o *AccessLogFieldMappingShared) SetLogFieldNewRDN(v string)`

SetLogFieldNewRDN sets LogFieldNewRDN field to given value.

### HasLogFieldNewRDN

`func (o *AccessLogFieldMappingShared) HasLogFieldNewRDN() bool`

HasLogFieldNewRDN returns a boolean if a field has been set.

### GetLogFieldBaseDN

`func (o *AccessLogFieldMappingShared) GetLogFieldBaseDN() string`

GetLogFieldBaseDN returns the LogFieldBaseDN field if non-nil, zero value otherwise.

### GetLogFieldBaseDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldBaseDNOk() (*string, bool)`

GetLogFieldBaseDNOk returns a tuple with the LogFieldBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldBaseDN

`func (o *AccessLogFieldMappingShared) SetLogFieldBaseDN(v string)`

SetLogFieldBaseDN sets LogFieldBaseDN field to given value.

### HasLogFieldBaseDN

`func (o *AccessLogFieldMappingShared) HasLogFieldBaseDN() bool`

HasLogFieldBaseDN returns a boolean if a field has been set.

### GetLogFieldBindDN

`func (o *AccessLogFieldMappingShared) GetLogFieldBindDN() string`

GetLogFieldBindDN returns the LogFieldBindDN field if non-nil, zero value otherwise.

### GetLogFieldBindDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldBindDNOk() (*string, bool)`

GetLogFieldBindDNOk returns a tuple with the LogFieldBindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldBindDN

`func (o *AccessLogFieldMappingShared) SetLogFieldBindDN(v string)`

SetLogFieldBindDN sets LogFieldBindDN field to given value.

### HasLogFieldBindDN

`func (o *AccessLogFieldMappingShared) HasLogFieldBindDN() bool`

HasLogFieldBindDN returns a boolean if a field has been set.

### GetLogFieldMatchedDN

`func (o *AccessLogFieldMappingShared) GetLogFieldMatchedDN() string`

GetLogFieldMatchedDN returns the LogFieldMatchedDN field if non-nil, zero value otherwise.

### GetLogFieldMatchedDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldMatchedDNOk() (*string, bool)`

GetLogFieldMatchedDNOk returns a tuple with the LogFieldMatchedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMatchedDN

`func (o *AccessLogFieldMappingShared) SetLogFieldMatchedDN(v string)`

SetLogFieldMatchedDN sets LogFieldMatchedDN field to given value.

### HasLogFieldMatchedDN

`func (o *AccessLogFieldMappingShared) HasLogFieldMatchedDN() bool`

HasLogFieldMatchedDN returns a boolean if a field has been set.

### GetLogFieldRequesterIPAddress

`func (o *AccessLogFieldMappingShared) GetLogFieldRequesterIPAddress() string`

GetLogFieldRequesterIPAddress returns the LogFieldRequesterIPAddress field if non-nil, zero value otherwise.

### GetLogFieldRequesterIPAddressOk

`func (o *AccessLogFieldMappingShared) GetLogFieldRequesterIPAddressOk() (*string, bool)`

GetLogFieldRequesterIPAddressOk returns a tuple with the LogFieldRequesterIPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequesterIPAddress

`func (o *AccessLogFieldMappingShared) SetLogFieldRequesterIPAddress(v string)`

SetLogFieldRequesterIPAddress sets LogFieldRequesterIPAddress field to given value.

### HasLogFieldRequesterIPAddress

`func (o *AccessLogFieldMappingShared) HasLogFieldRequesterIPAddress() bool`

HasLogFieldRequesterIPAddress returns a boolean if a field has been set.

### GetLogFieldAuthenticationType

`func (o *AccessLogFieldMappingShared) GetLogFieldAuthenticationType() string`

GetLogFieldAuthenticationType returns the LogFieldAuthenticationType field if non-nil, zero value otherwise.

### GetLogFieldAuthenticationTypeOk

`func (o *AccessLogFieldMappingShared) GetLogFieldAuthenticationTypeOk() (*string, bool)`

GetLogFieldAuthenticationTypeOk returns a tuple with the LogFieldAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAuthenticationType

`func (o *AccessLogFieldMappingShared) SetLogFieldAuthenticationType(v string)`

SetLogFieldAuthenticationType sets LogFieldAuthenticationType field to given value.

### HasLogFieldAuthenticationType

`func (o *AccessLogFieldMappingShared) HasLogFieldAuthenticationType() bool`

HasLogFieldAuthenticationType returns a boolean if a field has been set.

### GetLogFieldNewSuperiorDN

`func (o *AccessLogFieldMappingShared) GetLogFieldNewSuperiorDN() string`

GetLogFieldNewSuperiorDN returns the LogFieldNewSuperiorDN field if non-nil, zero value otherwise.

### GetLogFieldNewSuperiorDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldNewSuperiorDNOk() (*string, bool)`

GetLogFieldNewSuperiorDNOk returns a tuple with the LogFieldNewSuperiorDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldNewSuperiorDN

`func (o *AccessLogFieldMappingShared) SetLogFieldNewSuperiorDN(v string)`

SetLogFieldNewSuperiorDN sets LogFieldNewSuperiorDN field to given value.

### HasLogFieldNewSuperiorDN

`func (o *AccessLogFieldMappingShared) HasLogFieldNewSuperiorDN() bool`

HasLogFieldNewSuperiorDN returns a boolean if a field has been set.

### GetLogFieldFilter

`func (o *AccessLogFieldMappingShared) GetLogFieldFilter() string`

GetLogFieldFilter returns the LogFieldFilter field if non-nil, zero value otherwise.

### GetLogFieldFilterOk

`func (o *AccessLogFieldMappingShared) GetLogFieldFilterOk() (*string, bool)`

GetLogFieldFilterOk returns a tuple with the LogFieldFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldFilter

`func (o *AccessLogFieldMappingShared) SetLogFieldFilter(v string)`

SetLogFieldFilter sets LogFieldFilter field to given value.

### HasLogFieldFilter

`func (o *AccessLogFieldMappingShared) HasLogFieldFilter() bool`

HasLogFieldFilter returns a boolean if a field has been set.

### GetLogFieldAlternateAuthorizationDN

`func (o *AccessLogFieldMappingShared) GetLogFieldAlternateAuthorizationDN() string`

GetLogFieldAlternateAuthorizationDN returns the LogFieldAlternateAuthorizationDN field if non-nil, zero value otherwise.

### GetLogFieldAlternateAuthorizationDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldAlternateAuthorizationDNOk() (*string, bool)`

GetLogFieldAlternateAuthorizationDNOk returns a tuple with the LogFieldAlternateAuthorizationDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAlternateAuthorizationDN

`func (o *AccessLogFieldMappingShared) SetLogFieldAlternateAuthorizationDN(v string)`

SetLogFieldAlternateAuthorizationDN sets LogFieldAlternateAuthorizationDN field to given value.

### HasLogFieldAlternateAuthorizationDN

`func (o *AccessLogFieldMappingShared) HasLogFieldAlternateAuthorizationDN() bool`

HasLogFieldAlternateAuthorizationDN returns a boolean if a field has been set.

### GetLogFieldEntryDN

`func (o *AccessLogFieldMappingShared) GetLogFieldEntryDN() string`

GetLogFieldEntryDN returns the LogFieldEntryDN field if non-nil, zero value otherwise.

### GetLogFieldEntryDNOk

`func (o *AccessLogFieldMappingShared) GetLogFieldEntryDNOk() (*string, bool)`

GetLogFieldEntryDNOk returns a tuple with the LogFieldEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldEntryDN

`func (o *AccessLogFieldMappingShared) SetLogFieldEntryDN(v string)`

SetLogFieldEntryDN sets LogFieldEntryDN field to given value.

### HasLogFieldEntryDN

`func (o *AccessLogFieldMappingShared) HasLogFieldEntryDN() bool`

HasLogFieldEntryDN returns a boolean if a field has been set.

### GetLogFieldEntriesReturned

`func (o *AccessLogFieldMappingShared) GetLogFieldEntriesReturned() string`

GetLogFieldEntriesReturned returns the LogFieldEntriesReturned field if non-nil, zero value otherwise.

### GetLogFieldEntriesReturnedOk

`func (o *AccessLogFieldMappingShared) GetLogFieldEntriesReturnedOk() (*string, bool)`

GetLogFieldEntriesReturnedOk returns a tuple with the LogFieldEntriesReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldEntriesReturned

`func (o *AccessLogFieldMappingShared) SetLogFieldEntriesReturned(v string)`

SetLogFieldEntriesReturned sets LogFieldEntriesReturned field to given value.

### HasLogFieldEntriesReturned

`func (o *AccessLogFieldMappingShared) HasLogFieldEntriesReturned() bool`

HasLogFieldEntriesReturned returns a boolean if a field has been set.

### GetLogFieldAuthenticationFailureID

`func (o *AccessLogFieldMappingShared) GetLogFieldAuthenticationFailureID() string`

GetLogFieldAuthenticationFailureID returns the LogFieldAuthenticationFailureID field if non-nil, zero value otherwise.

### GetLogFieldAuthenticationFailureIDOk

`func (o *AccessLogFieldMappingShared) GetLogFieldAuthenticationFailureIDOk() (*string, bool)`

GetLogFieldAuthenticationFailureIDOk returns a tuple with the LogFieldAuthenticationFailureID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAuthenticationFailureID

`func (o *AccessLogFieldMappingShared) SetLogFieldAuthenticationFailureID(v string)`

SetLogFieldAuthenticationFailureID sets LogFieldAuthenticationFailureID field to given value.

### HasLogFieldAuthenticationFailureID

`func (o *AccessLogFieldMappingShared) HasLogFieldAuthenticationFailureID() bool`

HasLogFieldAuthenticationFailureID returns a boolean if a field has been set.

### GetLogFieldRequestOID

`func (o *AccessLogFieldMappingShared) GetLogFieldRequestOID() string`

GetLogFieldRequestOID returns the LogFieldRequestOID field if non-nil, zero value otherwise.

### GetLogFieldRequestOIDOk

`func (o *AccessLogFieldMappingShared) GetLogFieldRequestOIDOk() (*string, bool)`

GetLogFieldRequestOIDOk returns a tuple with the LogFieldRequestOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequestOID

`func (o *AccessLogFieldMappingShared) SetLogFieldRequestOID(v string)`

SetLogFieldRequestOID sets LogFieldRequestOID field to given value.

### HasLogFieldRequestOID

`func (o *AccessLogFieldMappingShared) HasLogFieldRequestOID() bool`

HasLogFieldRequestOID returns a boolean if a field has been set.

### GetLogFieldResponseOID

`func (o *AccessLogFieldMappingShared) GetLogFieldResponseOID() string`

GetLogFieldResponseOID returns the LogFieldResponseOID field if non-nil, zero value otherwise.

### GetLogFieldResponseOIDOk

`func (o *AccessLogFieldMappingShared) GetLogFieldResponseOIDOk() (*string, bool)`

GetLogFieldResponseOIDOk returns a tuple with the LogFieldResponseOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldResponseOID

`func (o *AccessLogFieldMappingShared) SetLogFieldResponseOID(v string)`

SetLogFieldResponseOID sets LogFieldResponseOID field to given value.

### HasLogFieldResponseOID

`func (o *AccessLogFieldMappingShared) HasLogFieldResponseOID() bool`

HasLogFieldResponseOID returns a boolean if a field has been set.

### GetLogFieldTargetProtocol

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetProtocol() string`

GetLogFieldTargetProtocol returns the LogFieldTargetProtocol field if non-nil, zero value otherwise.

### GetLogFieldTargetProtocolOk

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetProtocolOk() (*string, bool)`

GetLogFieldTargetProtocolOk returns a tuple with the LogFieldTargetProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetProtocol

`func (o *AccessLogFieldMappingShared) SetLogFieldTargetProtocol(v string)`

SetLogFieldTargetProtocol sets LogFieldTargetProtocol field to given value.

### HasLogFieldTargetProtocol

`func (o *AccessLogFieldMappingShared) HasLogFieldTargetProtocol() bool`

HasLogFieldTargetProtocol returns a boolean if a field has been set.

### GetLogFieldTargetPort

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetPort() string`

GetLogFieldTargetPort returns the LogFieldTargetPort field if non-nil, zero value otherwise.

### GetLogFieldTargetPortOk

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetPortOk() (*string, bool)`

GetLogFieldTargetPortOk returns a tuple with the LogFieldTargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetPort

`func (o *AccessLogFieldMappingShared) SetLogFieldTargetPort(v string)`

SetLogFieldTargetPort sets LogFieldTargetPort field to given value.

### HasLogFieldTargetPort

`func (o *AccessLogFieldMappingShared) HasLogFieldTargetPort() bool`

HasLogFieldTargetPort returns a boolean if a field has been set.

### GetLogFieldTargetAddress

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetAddress() string`

GetLogFieldTargetAddress returns the LogFieldTargetAddress field if non-nil, zero value otherwise.

### GetLogFieldTargetAddressOk

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetAddressOk() (*string, bool)`

GetLogFieldTargetAddressOk returns a tuple with the LogFieldTargetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetAddress

`func (o *AccessLogFieldMappingShared) SetLogFieldTargetAddress(v string)`

SetLogFieldTargetAddress sets LogFieldTargetAddress field to given value.

### HasLogFieldTargetAddress

`func (o *AccessLogFieldMappingShared) HasLogFieldTargetAddress() bool`

HasLogFieldTargetAddress returns a boolean if a field has been set.

### GetLogFieldTargetAttribute

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetAttribute() string`

GetLogFieldTargetAttribute returns the LogFieldTargetAttribute field if non-nil, zero value otherwise.

### GetLogFieldTargetAttributeOk

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetAttributeOk() (*string, bool)`

GetLogFieldTargetAttributeOk returns a tuple with the LogFieldTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetAttribute

`func (o *AccessLogFieldMappingShared) SetLogFieldTargetAttribute(v string)`

SetLogFieldTargetAttribute sets LogFieldTargetAttribute field to given value.

### HasLogFieldTargetAttribute

`func (o *AccessLogFieldMappingShared) HasLogFieldTargetAttribute() bool`

HasLogFieldTargetAttribute returns a boolean if a field has been set.

### GetLogFieldTargetHost

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetHost() string`

GetLogFieldTargetHost returns the LogFieldTargetHost field if non-nil, zero value otherwise.

### GetLogFieldTargetHostOk

`func (o *AccessLogFieldMappingShared) GetLogFieldTargetHostOk() (*string, bool)`

GetLogFieldTargetHostOk returns a tuple with the LogFieldTargetHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldTargetHost

`func (o *AccessLogFieldMappingShared) SetLogFieldTargetHost(v string)`

SetLogFieldTargetHost sets LogFieldTargetHost field to given value.

### HasLogFieldTargetHost

`func (o *AccessLogFieldMappingShared) HasLogFieldTargetHost() bool`

HasLogFieldTargetHost returns a boolean if a field has been set.

### GetLogFieldProtocolVersion

`func (o *AccessLogFieldMappingShared) GetLogFieldProtocolVersion() string`

GetLogFieldProtocolVersion returns the LogFieldProtocolVersion field if non-nil, zero value otherwise.

### GetLogFieldProtocolVersionOk

`func (o *AccessLogFieldMappingShared) GetLogFieldProtocolVersionOk() (*string, bool)`

GetLogFieldProtocolVersionOk returns a tuple with the LogFieldProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProtocolVersion

`func (o *AccessLogFieldMappingShared) SetLogFieldProtocolVersion(v string)`

SetLogFieldProtocolVersion sets LogFieldProtocolVersion field to given value.

### HasLogFieldProtocolVersion

`func (o *AccessLogFieldMappingShared) HasLogFieldProtocolVersion() bool`

HasLogFieldProtocolVersion returns a boolean if a field has been set.

### GetLogFieldProtocolName

`func (o *AccessLogFieldMappingShared) GetLogFieldProtocolName() string`

GetLogFieldProtocolName returns the LogFieldProtocolName field if non-nil, zero value otherwise.

### GetLogFieldProtocolNameOk

`func (o *AccessLogFieldMappingShared) GetLogFieldProtocolNameOk() (*string, bool)`

GetLogFieldProtocolNameOk returns a tuple with the LogFieldProtocolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldProtocolName

`func (o *AccessLogFieldMappingShared) SetLogFieldProtocolName(v string)`

SetLogFieldProtocolName sets LogFieldProtocolName field to given value.

### HasLogFieldProtocolName

`func (o *AccessLogFieldMappingShared) HasLogFieldProtocolName() bool`

HasLogFieldProtocolName returns a boolean if a field has been set.

### GetLogFieldAuthenticationFailureReason

`func (o *AccessLogFieldMappingShared) GetLogFieldAuthenticationFailureReason() string`

GetLogFieldAuthenticationFailureReason returns the LogFieldAuthenticationFailureReason field if non-nil, zero value otherwise.

### GetLogFieldAuthenticationFailureReasonOk

`func (o *AccessLogFieldMappingShared) GetLogFieldAuthenticationFailureReasonOk() (*string, bool)`

GetLogFieldAuthenticationFailureReasonOk returns a tuple with the LogFieldAuthenticationFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAuthenticationFailureReason

`func (o *AccessLogFieldMappingShared) SetLogFieldAuthenticationFailureReason(v string)`

SetLogFieldAuthenticationFailureReason sets LogFieldAuthenticationFailureReason field to given value.

### HasLogFieldAuthenticationFailureReason

`func (o *AccessLogFieldMappingShared) HasLogFieldAuthenticationFailureReason() bool`

HasLogFieldAuthenticationFailureReason returns a boolean if a field has been set.

### GetLogFieldAdditionalInformation

`func (o *AccessLogFieldMappingShared) GetLogFieldAdditionalInformation() string`

GetLogFieldAdditionalInformation returns the LogFieldAdditionalInformation field if non-nil, zero value otherwise.

### GetLogFieldAdditionalInformationOk

`func (o *AccessLogFieldMappingShared) GetLogFieldAdditionalInformationOk() (*string, bool)`

GetLogFieldAdditionalInformationOk returns a tuple with the LogFieldAdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldAdditionalInformation

`func (o *AccessLogFieldMappingShared) SetLogFieldAdditionalInformation(v string)`

SetLogFieldAdditionalInformation sets LogFieldAdditionalInformation field to given value.

### HasLogFieldAdditionalInformation

`func (o *AccessLogFieldMappingShared) HasLogFieldAdditionalInformation() bool`

HasLogFieldAdditionalInformation returns a boolean if a field has been set.

### GetLogFieldUnindexed

`func (o *AccessLogFieldMappingShared) GetLogFieldUnindexed() string`

GetLogFieldUnindexed returns the LogFieldUnindexed field if non-nil, zero value otherwise.

### GetLogFieldUnindexedOk

`func (o *AccessLogFieldMappingShared) GetLogFieldUnindexedOk() (*string, bool)`

GetLogFieldUnindexedOk returns a tuple with the LogFieldUnindexed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldUnindexed

`func (o *AccessLogFieldMappingShared) SetLogFieldUnindexed(v string)`

SetLogFieldUnindexed sets LogFieldUnindexed field to given value.

### HasLogFieldUnindexed

`func (o *AccessLogFieldMappingShared) HasLogFieldUnindexed() bool`

HasLogFieldUnindexed returns a boolean if a field has been set.

### GetLogFieldScope

`func (o *AccessLogFieldMappingShared) GetLogFieldScope() string`

GetLogFieldScope returns the LogFieldScope field if non-nil, zero value otherwise.

### GetLogFieldScopeOk

`func (o *AccessLogFieldMappingShared) GetLogFieldScopeOk() (*string, bool)`

GetLogFieldScopeOk returns a tuple with the LogFieldScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldScope

`func (o *AccessLogFieldMappingShared) SetLogFieldScope(v string)`

SetLogFieldScope sets LogFieldScope field to given value.

### HasLogFieldScope

`func (o *AccessLogFieldMappingShared) HasLogFieldScope() bool`

HasLogFieldScope returns a boolean if a field has been set.

### GetLogFieldReferralUrls

`func (o *AccessLogFieldMappingShared) GetLogFieldReferralUrls() string`

GetLogFieldReferralUrls returns the LogFieldReferralUrls field if non-nil, zero value otherwise.

### GetLogFieldReferralUrlsOk

`func (o *AccessLogFieldMappingShared) GetLogFieldReferralUrlsOk() (*string, bool)`

GetLogFieldReferralUrlsOk returns a tuple with the LogFieldReferralUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldReferralUrls

`func (o *AccessLogFieldMappingShared) SetLogFieldReferralUrls(v string)`

SetLogFieldReferralUrls sets LogFieldReferralUrls field to given value.

### HasLogFieldReferralUrls

`func (o *AccessLogFieldMappingShared) HasLogFieldReferralUrls() bool`

HasLogFieldReferralUrls returns a boolean if a field has been set.

### GetLogFieldSourceAddress

`func (o *AccessLogFieldMappingShared) GetLogFieldSourceAddress() string`

GetLogFieldSourceAddress returns the LogFieldSourceAddress field if non-nil, zero value otherwise.

### GetLogFieldSourceAddressOk

`func (o *AccessLogFieldMappingShared) GetLogFieldSourceAddressOk() (*string, bool)`

GetLogFieldSourceAddressOk returns a tuple with the LogFieldSourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldSourceAddress

`func (o *AccessLogFieldMappingShared) SetLogFieldSourceAddress(v string)`

SetLogFieldSourceAddress sets LogFieldSourceAddress field to given value.

### HasLogFieldSourceAddress

`func (o *AccessLogFieldMappingShared) HasLogFieldSourceAddress() bool`

HasLogFieldSourceAddress returns a boolean if a field has been set.

### GetLogFieldMessageIDToAbandon

`func (o *AccessLogFieldMappingShared) GetLogFieldMessageIDToAbandon() string`

GetLogFieldMessageIDToAbandon returns the LogFieldMessageIDToAbandon field if non-nil, zero value otherwise.

### GetLogFieldMessageIDToAbandonOk

`func (o *AccessLogFieldMappingShared) GetLogFieldMessageIDToAbandonOk() (*string, bool)`

GetLogFieldMessageIDToAbandonOk returns a tuple with the LogFieldMessageIDToAbandon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldMessageIDToAbandon

`func (o *AccessLogFieldMappingShared) SetLogFieldMessageIDToAbandon(v string)`

SetLogFieldMessageIDToAbandon sets LogFieldMessageIDToAbandon field to given value.

### HasLogFieldMessageIDToAbandon

`func (o *AccessLogFieldMappingShared) HasLogFieldMessageIDToAbandon() bool`

HasLogFieldMessageIDToAbandon returns a boolean if a field has been set.

### GetLogFieldResponseControls

`func (o *AccessLogFieldMappingShared) GetLogFieldResponseControls() string`

GetLogFieldResponseControls returns the LogFieldResponseControls field if non-nil, zero value otherwise.

### GetLogFieldResponseControlsOk

`func (o *AccessLogFieldMappingShared) GetLogFieldResponseControlsOk() (*string, bool)`

GetLogFieldResponseControlsOk returns a tuple with the LogFieldResponseControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldResponseControls

`func (o *AccessLogFieldMappingShared) SetLogFieldResponseControls(v string)`

SetLogFieldResponseControls sets LogFieldResponseControls field to given value.

### HasLogFieldResponseControls

`func (o *AccessLogFieldMappingShared) HasLogFieldResponseControls() bool`

HasLogFieldResponseControls returns a boolean if a field has been set.

### GetLogFieldRequestControls

`func (o *AccessLogFieldMappingShared) GetLogFieldRequestControls() string`

GetLogFieldRequestControls returns the LogFieldRequestControls field if non-nil, zero value otherwise.

### GetLogFieldRequestControlsOk

`func (o *AccessLogFieldMappingShared) GetLogFieldRequestControlsOk() (*string, bool)`

GetLogFieldRequestControlsOk returns a tuple with the LogFieldRequestControls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldRequestControls

`func (o *AccessLogFieldMappingShared) SetLogFieldRequestControls(v string)`

SetLogFieldRequestControls sets LogFieldRequestControls field to given value.

### HasLogFieldRequestControls

`func (o *AccessLogFieldMappingShared) HasLogFieldRequestControls() bool`

HasLogFieldRequestControls returns a boolean if a field has been set.

### GetLogFieldIntermediateClientResult

`func (o *AccessLogFieldMappingShared) GetLogFieldIntermediateClientResult() string`

GetLogFieldIntermediateClientResult returns the LogFieldIntermediateClientResult field if non-nil, zero value otherwise.

### GetLogFieldIntermediateClientResultOk

`func (o *AccessLogFieldMappingShared) GetLogFieldIntermediateClientResultOk() (*string, bool)`

GetLogFieldIntermediateClientResultOk returns a tuple with the LogFieldIntermediateClientResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldIntermediateClientResult

`func (o *AccessLogFieldMappingShared) SetLogFieldIntermediateClientResult(v string)`

SetLogFieldIntermediateClientResult sets LogFieldIntermediateClientResult field to given value.

### HasLogFieldIntermediateClientResult

`func (o *AccessLogFieldMappingShared) HasLogFieldIntermediateClientResult() bool`

HasLogFieldIntermediateClientResult returns a boolean if a field has been set.

### GetLogFieldIntermediateClientRequest

`func (o *AccessLogFieldMappingShared) GetLogFieldIntermediateClientRequest() string`

GetLogFieldIntermediateClientRequest returns the LogFieldIntermediateClientRequest field if non-nil, zero value otherwise.

### GetLogFieldIntermediateClientRequestOk

`func (o *AccessLogFieldMappingShared) GetLogFieldIntermediateClientRequestOk() (*string, bool)`

GetLogFieldIntermediateClientRequestOk returns a tuple with the LogFieldIntermediateClientRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldIntermediateClientRequest

`func (o *AccessLogFieldMappingShared) SetLogFieldIntermediateClientRequest(v string)`

SetLogFieldIntermediateClientRequest sets LogFieldIntermediateClientRequest field to given value.

### HasLogFieldIntermediateClientRequest

`func (o *AccessLogFieldMappingShared) HasLogFieldIntermediateClientRequest() bool`

HasLogFieldIntermediateClientRequest returns a boolean if a field has been set.

### GetLogFieldReplicationChangeID

`func (o *AccessLogFieldMappingShared) GetLogFieldReplicationChangeID() string`

GetLogFieldReplicationChangeID returns the LogFieldReplicationChangeID field if non-nil, zero value otherwise.

### GetLogFieldReplicationChangeIDOk

`func (o *AccessLogFieldMappingShared) GetLogFieldReplicationChangeIDOk() (*string, bool)`

GetLogFieldReplicationChangeIDOk returns a tuple with the LogFieldReplicationChangeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFieldReplicationChangeID

`func (o *AccessLogFieldMappingShared) SetLogFieldReplicationChangeID(v string)`

SetLogFieldReplicationChangeID sets LogFieldReplicationChangeID field to given value.

### HasLogFieldReplicationChangeID

`func (o *AccessLogFieldMappingShared) HasLogFieldReplicationChangeID() bool`

HasLogFieldReplicationChangeID returns a boolean if a field has been set.

### GetDescription

`func (o *AccessLogFieldMappingShared) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessLogFieldMappingShared) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessLogFieldMappingShared) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessLogFieldMappingShared) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


