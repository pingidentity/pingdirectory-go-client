# SimpleRequestCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Request Criteria | 
**Schemas** | [**[]EnumsimpleRequestCriteriaSchemaUrn**](EnumsimpleRequestCriteriaSchemaUrn.md) |  | 
**OperationType** | Pointer to [**[]EnumrequestCriteriaSimpleOperationTypeProp**](EnumrequestCriteriaSimpleOperationTypeProp.md) |  | [optional] 
**OperationOrigin** | Pointer to [**[]EnumrequestCriteriaOperationOriginProp**](EnumrequestCriteriaOperationOriginProp.md) |  | [optional] 
**ConnectionCriteria** | Pointer to **string** | Specifies a connection criteria object that must match the associated client connection for operations included in this Simple Request Criteria. | [optional] 
**AllIncludedRequestControl** | Pointer to **[]string** | Specifies the OID of a control that must be present in the request from the client for operations included in this Simple Request Criteria. If any control OIDs are provided, then the request must contain all of those controls. | [optional] 
**AnyIncludedRequestControl** | Pointer to **[]string** | Specifies the OID of a control that may be present in the request from the client for operations included in this Simple Request Criteria. If any control OIDs are provided, then the request must contain at least one of those controls. | [optional] 
**NotAllIncludedRequestControl** | Pointer to **[]string** | Specifies the OID of a control that should not be present in the request from the client for operations included in this Simple Request Criteria. If any control OIDs are provided, then the request must not contain at least one of those controls (that is, the request may contain zero or more of those controls, but not all of them). | [optional] 
**NoneIncludedRequestControl** | Pointer to **[]string** | Specifies the OID of a control that must not be present in the request from the client for operations included in this Simple Request Criteria. If any control OIDs are provided, then the request must not contain any of those controls. | [optional] 
**IncludedTargetEntryDN** | Pointer to **[]string** | Specifies a base DN below which targeted entries may exist for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. | [optional] 
**ExcludedTargetEntryDN** | Pointer to **[]string** | Specifies a base DN below which targeted entries may not exist for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. | [optional] 
**AllIncludedTargetEntryFilter** | Pointer to **[]string** | Specifies a search filter that must match the target entry for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any filters are provided, then the target entry must match all of those filters. | [optional] 
**AnyIncludedTargetEntryFilter** | Pointer to **[]string** | Specifies a search filter that may match the target entry for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any filters are provided, then the target entry must match at least one of those filters. | [optional] 
**NotAllIncludedTargetEntryFilter** | Pointer to **[]string** | Specifies a search filter that should not match the target entry for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any filters are provided, then the target entry must not match at least one of those filters (that is, the request may match zero or more of those filters, but not of all of them). | [optional] 
**NoneIncludedTargetEntryFilter** | Pointer to **[]string** | Specifies a search filter that must not match the target entry for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any filters are provided, then the target entry must not match any of those filters. | [optional] 
**AllIncludedTargetEntryGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which the user associated with the target entry must be a member for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any group DNs are provided, then the target entry must be a member of all of those groups. | [optional] 
**AnyIncludedTargetEntryGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which the user associated with the target entry may be a member for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any group DNs are provided, then the target entry must be a member of at least one of those groups. | [optional] 
**NotAllIncludedTargetEntryGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which the user associated with the target entry should not be a member for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any group DNs are provided, then the target entry must not be a member of at least one of those groups (that is, the target entry may be a member of zero or more of those groups, but not all of them). | [optional] 
**NoneIncludedTargetEntryGroupDN** | Pointer to **[]string** | Specifies the DN of a group in which the user associated with the target entry must not be a member for requests included in this Simple Request Criteria. This will only be taken into account for add, simple bind, compare, delete, modify, modify DN, and search operations. It will be ignored for abandon, SASL bind, extended, and unbind operations. If any group DNs are provided, then the target entry must not be a member of any of those groups. | [optional] 
**TargetBindType** | Pointer to [**[]EnumrequestCriteriaTargetBindTypeProp**](EnumrequestCriteriaTargetBindTypeProp.md) |  | [optional] 
**IncludedTargetSASLMechanism** | Pointer to **[]string** | Specifies the name of a SASL mechanism for bind requests included in this Simple Request Criteria. This will only be taken into account for SASL bind operations and will be ignored for other types of operations and for bind operations that do not use SASL authentication. | [optional] 
**ExcludedTargetSASLMechanism** | Pointer to **[]string** | Specifies the name of a SASL mechanism for bind requests excluded from this Simple Request Criteria. This will only be taken into account for SASL bind operations and will be ignored for other types of operations and for bind operations that do not use SASL authentication. | [optional] 
**IncludedTargetAttribute** | Pointer to **[]string** | Specifies the name or OID of an attribute type which must be targeted by requests included in this Simple Request Criteria. This will only be taken into account for add, compare, modify, modify DN, and search operations. It will be ignored for abandon, bind, delete, extended, and unbind operations. | [optional] 
**ExcludedTargetAttribute** | Pointer to **[]string** | Specifies the name or OID of an attribute type which must not be targeted by requests included in this Simple Request Criteria. This will only be taken into account for add, compare, modify, modify DN, and search operations. It will be ignored for abandon, bind, delete, extended, and unbind operations. | [optional] 
**IncludedExtendedOperationOID** | Pointer to **[]string** | Specifies the request OID for extended requests included in this Simple Request Criteria. This will only be taken into account for extended requests and will be ignored for all other types of requests. | [optional] 
**ExcludedExtendedOperationOID** | Pointer to **[]string** | Specifies the request OID for extended requests excluded from this Simple Request Criteria. This will only be taken into account for extended requests and will be ignored for all other types of requests. | [optional] 
**IncludedSearchScope** | Pointer to [**[]EnumrequestCriteriaIncludedSearchScopeProp**](EnumrequestCriteriaIncludedSearchScopeProp.md) |  | [optional] 
**UsingAdministrativeSessionWorkerThread** | Pointer to [**EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp**](EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp.md) |  | [optional] 
**IncludedApplicationName** | Pointer to **[]string** | Specifies an application name for requests included in this Simple Request Criteria. | [optional] 
**ExcludedApplicationName** | Pointer to **[]string** | Specifies an application name for requests excluded from this Simple Request Criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Request Criteria | [optional] 

## Methods

### NewSimpleRequestCriteriaResponse

`func NewSimpleRequestCriteriaResponse(id string, schemas []EnumsimpleRequestCriteriaSchemaUrn, ) *SimpleRequestCriteriaResponse`

NewSimpleRequestCriteriaResponse instantiates a new SimpleRequestCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleRequestCriteriaResponseWithDefaults

`func NewSimpleRequestCriteriaResponseWithDefaults() *SimpleRequestCriteriaResponse`

NewSimpleRequestCriteriaResponseWithDefaults instantiates a new SimpleRequestCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *SimpleRequestCriteriaResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *SimpleRequestCriteriaResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *SimpleRequestCriteriaResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *SimpleRequestCriteriaResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleRequestCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *SimpleRequestCriteriaResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleRequestCriteriaResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *SimpleRequestCriteriaResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *SimpleRequestCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleRequestCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleRequestCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SimpleRequestCriteriaResponse) GetSchemas() []EnumsimpleRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SimpleRequestCriteriaResponse) GetSchemasOk() (*[]EnumsimpleRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SimpleRequestCriteriaResponse) SetSchemas(v []EnumsimpleRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOperationType

`func (o *SimpleRequestCriteriaResponse) GetOperationType() []EnumrequestCriteriaSimpleOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *SimpleRequestCriteriaResponse) GetOperationTypeOk() (*[]EnumrequestCriteriaSimpleOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *SimpleRequestCriteriaResponse) SetOperationType(v []EnumrequestCriteriaSimpleOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *SimpleRequestCriteriaResponse) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetOperationOrigin

`func (o *SimpleRequestCriteriaResponse) GetOperationOrigin() []EnumrequestCriteriaOperationOriginProp`

GetOperationOrigin returns the OperationOrigin field if non-nil, zero value otherwise.

### GetOperationOriginOk

`func (o *SimpleRequestCriteriaResponse) GetOperationOriginOk() (*[]EnumrequestCriteriaOperationOriginProp, bool)`

GetOperationOriginOk returns a tuple with the OperationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationOrigin

`func (o *SimpleRequestCriteriaResponse) SetOperationOrigin(v []EnumrequestCriteriaOperationOriginProp)`

SetOperationOrigin sets OperationOrigin field to given value.

### HasOperationOrigin

`func (o *SimpleRequestCriteriaResponse) HasOperationOrigin() bool`

HasOperationOrigin returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *SimpleRequestCriteriaResponse) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *SimpleRequestCriteriaResponse) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *SimpleRequestCriteriaResponse) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *SimpleRequestCriteriaResponse) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetAllIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) GetAllIncludedRequestControl() []string`

GetAllIncludedRequestControl returns the AllIncludedRequestControl field if non-nil, zero value otherwise.

### GetAllIncludedRequestControlOk

`func (o *SimpleRequestCriteriaResponse) GetAllIncludedRequestControlOk() (*[]string, bool)`

GetAllIncludedRequestControlOk returns a tuple with the AllIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) SetAllIncludedRequestControl(v []string)`

SetAllIncludedRequestControl sets AllIncludedRequestControl field to given value.

### HasAllIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) HasAllIncludedRequestControl() bool`

HasAllIncludedRequestControl returns a boolean if a field has been set.

### GetAnyIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) GetAnyIncludedRequestControl() []string`

GetAnyIncludedRequestControl returns the AnyIncludedRequestControl field if non-nil, zero value otherwise.

### GetAnyIncludedRequestControlOk

`func (o *SimpleRequestCriteriaResponse) GetAnyIncludedRequestControlOk() (*[]string, bool)`

GetAnyIncludedRequestControlOk returns a tuple with the AnyIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) SetAnyIncludedRequestControl(v []string)`

SetAnyIncludedRequestControl sets AnyIncludedRequestControl field to given value.

### HasAnyIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) HasAnyIncludedRequestControl() bool`

HasAnyIncludedRequestControl returns a boolean if a field has been set.

### GetNotAllIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) GetNotAllIncludedRequestControl() []string`

GetNotAllIncludedRequestControl returns the NotAllIncludedRequestControl field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestControlOk

`func (o *SimpleRequestCriteriaResponse) GetNotAllIncludedRequestControlOk() (*[]string, bool)`

GetNotAllIncludedRequestControlOk returns a tuple with the NotAllIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) SetNotAllIncludedRequestControl(v []string)`

SetNotAllIncludedRequestControl sets NotAllIncludedRequestControl field to given value.

### HasNotAllIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) HasNotAllIncludedRequestControl() bool`

HasNotAllIncludedRequestControl returns a boolean if a field has been set.

### GetNoneIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) GetNoneIncludedRequestControl() []string`

GetNoneIncludedRequestControl returns the NoneIncludedRequestControl field if non-nil, zero value otherwise.

### GetNoneIncludedRequestControlOk

`func (o *SimpleRequestCriteriaResponse) GetNoneIncludedRequestControlOk() (*[]string, bool)`

GetNoneIncludedRequestControlOk returns a tuple with the NoneIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) SetNoneIncludedRequestControl(v []string)`

SetNoneIncludedRequestControl sets NoneIncludedRequestControl field to given value.

### HasNoneIncludedRequestControl

`func (o *SimpleRequestCriteriaResponse) HasNoneIncludedRequestControl() bool`

HasNoneIncludedRequestControl returns a boolean if a field has been set.

### GetIncludedTargetEntryDN

`func (o *SimpleRequestCriteriaResponse) GetIncludedTargetEntryDN() []string`

GetIncludedTargetEntryDN returns the IncludedTargetEntryDN field if non-nil, zero value otherwise.

### GetIncludedTargetEntryDNOk

`func (o *SimpleRequestCriteriaResponse) GetIncludedTargetEntryDNOk() (*[]string, bool)`

GetIncludedTargetEntryDNOk returns a tuple with the IncludedTargetEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetEntryDN

`func (o *SimpleRequestCriteriaResponse) SetIncludedTargetEntryDN(v []string)`

SetIncludedTargetEntryDN sets IncludedTargetEntryDN field to given value.

### HasIncludedTargetEntryDN

`func (o *SimpleRequestCriteriaResponse) HasIncludedTargetEntryDN() bool`

HasIncludedTargetEntryDN returns a boolean if a field has been set.

### GetExcludedTargetEntryDN

`func (o *SimpleRequestCriteriaResponse) GetExcludedTargetEntryDN() []string`

GetExcludedTargetEntryDN returns the ExcludedTargetEntryDN field if non-nil, zero value otherwise.

### GetExcludedTargetEntryDNOk

`func (o *SimpleRequestCriteriaResponse) GetExcludedTargetEntryDNOk() (*[]string, bool)`

GetExcludedTargetEntryDNOk returns a tuple with the ExcludedTargetEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetEntryDN

`func (o *SimpleRequestCriteriaResponse) SetExcludedTargetEntryDN(v []string)`

SetExcludedTargetEntryDN sets ExcludedTargetEntryDN field to given value.

### HasExcludedTargetEntryDN

`func (o *SimpleRequestCriteriaResponse) HasExcludedTargetEntryDN() bool`

HasExcludedTargetEntryDN returns a boolean if a field has been set.

### GetAllIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) GetAllIncludedTargetEntryFilter() []string`

GetAllIncludedTargetEntryFilter returns the AllIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetAllIncludedTargetEntryFilterOk

`func (o *SimpleRequestCriteriaResponse) GetAllIncludedTargetEntryFilterOk() (*[]string, bool)`

GetAllIncludedTargetEntryFilterOk returns a tuple with the AllIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) SetAllIncludedTargetEntryFilter(v []string)`

SetAllIncludedTargetEntryFilter sets AllIncludedTargetEntryFilter field to given value.

### HasAllIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) HasAllIncludedTargetEntryFilter() bool`

HasAllIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetAnyIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) GetAnyIncludedTargetEntryFilter() []string`

GetAnyIncludedTargetEntryFilter returns the AnyIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetAnyIncludedTargetEntryFilterOk

`func (o *SimpleRequestCriteriaResponse) GetAnyIncludedTargetEntryFilterOk() (*[]string, bool)`

GetAnyIncludedTargetEntryFilterOk returns a tuple with the AnyIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) SetAnyIncludedTargetEntryFilter(v []string)`

SetAnyIncludedTargetEntryFilter sets AnyIncludedTargetEntryFilter field to given value.

### HasAnyIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) HasAnyIncludedTargetEntryFilter() bool`

HasAnyIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetNotAllIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) GetNotAllIncludedTargetEntryFilter() []string`

GetNotAllIncludedTargetEntryFilter returns the NotAllIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedTargetEntryFilterOk

`func (o *SimpleRequestCriteriaResponse) GetNotAllIncludedTargetEntryFilterOk() (*[]string, bool)`

GetNotAllIncludedTargetEntryFilterOk returns a tuple with the NotAllIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) SetNotAllIncludedTargetEntryFilter(v []string)`

SetNotAllIncludedTargetEntryFilter sets NotAllIncludedTargetEntryFilter field to given value.

### HasNotAllIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) HasNotAllIncludedTargetEntryFilter() bool`

HasNotAllIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetNoneIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) GetNoneIncludedTargetEntryFilter() []string`

GetNoneIncludedTargetEntryFilter returns the NoneIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetNoneIncludedTargetEntryFilterOk

`func (o *SimpleRequestCriteriaResponse) GetNoneIncludedTargetEntryFilterOk() (*[]string, bool)`

GetNoneIncludedTargetEntryFilterOk returns a tuple with the NoneIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) SetNoneIncludedTargetEntryFilter(v []string)`

SetNoneIncludedTargetEntryFilter sets NoneIncludedTargetEntryFilter field to given value.

### HasNoneIncludedTargetEntryFilter

`func (o *SimpleRequestCriteriaResponse) HasNoneIncludedTargetEntryFilter() bool`

HasNoneIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetAllIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) GetAllIncludedTargetEntryGroupDN() []string`

GetAllIncludedTargetEntryGroupDN returns the AllIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedTargetEntryGroupDNOk

`func (o *SimpleRequestCriteriaResponse) GetAllIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetAllIncludedTargetEntryGroupDNOk returns a tuple with the AllIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) SetAllIncludedTargetEntryGroupDN(v []string)`

SetAllIncludedTargetEntryGroupDN sets AllIncludedTargetEntryGroupDN field to given value.

### HasAllIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) HasAllIncludedTargetEntryGroupDN() bool`

HasAllIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetAnyIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) GetAnyIncludedTargetEntryGroupDN() []string`

GetAnyIncludedTargetEntryGroupDN returns the AnyIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedTargetEntryGroupDNOk

`func (o *SimpleRequestCriteriaResponse) GetAnyIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetAnyIncludedTargetEntryGroupDNOk returns a tuple with the AnyIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) SetAnyIncludedTargetEntryGroupDN(v []string)`

SetAnyIncludedTargetEntryGroupDN sets AnyIncludedTargetEntryGroupDN field to given value.

### HasAnyIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) HasAnyIncludedTargetEntryGroupDN() bool`

HasAnyIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) GetNotAllIncludedTargetEntryGroupDN() []string`

GetNotAllIncludedTargetEntryGroupDN returns the NotAllIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedTargetEntryGroupDNOk

`func (o *SimpleRequestCriteriaResponse) GetNotAllIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetNotAllIncludedTargetEntryGroupDNOk returns a tuple with the NotAllIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) SetNotAllIncludedTargetEntryGroupDN(v []string)`

SetNotAllIncludedTargetEntryGroupDN sets NotAllIncludedTargetEntryGroupDN field to given value.

### HasNotAllIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) HasNotAllIncludedTargetEntryGroupDN() bool`

HasNotAllIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetNoneIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) GetNoneIncludedTargetEntryGroupDN() []string`

GetNoneIncludedTargetEntryGroupDN returns the NoneIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedTargetEntryGroupDNOk

`func (o *SimpleRequestCriteriaResponse) GetNoneIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetNoneIncludedTargetEntryGroupDNOk returns a tuple with the NoneIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) SetNoneIncludedTargetEntryGroupDN(v []string)`

SetNoneIncludedTargetEntryGroupDN sets NoneIncludedTargetEntryGroupDN field to given value.

### HasNoneIncludedTargetEntryGroupDN

`func (o *SimpleRequestCriteriaResponse) HasNoneIncludedTargetEntryGroupDN() bool`

HasNoneIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetTargetBindType

`func (o *SimpleRequestCriteriaResponse) GetTargetBindType() []EnumrequestCriteriaTargetBindTypeProp`

GetTargetBindType returns the TargetBindType field if non-nil, zero value otherwise.

### GetTargetBindTypeOk

`func (o *SimpleRequestCriteriaResponse) GetTargetBindTypeOk() (*[]EnumrequestCriteriaTargetBindTypeProp, bool)`

GetTargetBindTypeOk returns a tuple with the TargetBindType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBindType

`func (o *SimpleRequestCriteriaResponse) SetTargetBindType(v []EnumrequestCriteriaTargetBindTypeProp)`

SetTargetBindType sets TargetBindType field to given value.

### HasTargetBindType

`func (o *SimpleRequestCriteriaResponse) HasTargetBindType() bool`

HasTargetBindType returns a boolean if a field has been set.

### GetIncludedTargetSASLMechanism

`func (o *SimpleRequestCriteriaResponse) GetIncludedTargetSASLMechanism() []string`

GetIncludedTargetSASLMechanism returns the IncludedTargetSASLMechanism field if non-nil, zero value otherwise.

### GetIncludedTargetSASLMechanismOk

`func (o *SimpleRequestCriteriaResponse) GetIncludedTargetSASLMechanismOk() (*[]string, bool)`

GetIncludedTargetSASLMechanismOk returns a tuple with the IncludedTargetSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetSASLMechanism

`func (o *SimpleRequestCriteriaResponse) SetIncludedTargetSASLMechanism(v []string)`

SetIncludedTargetSASLMechanism sets IncludedTargetSASLMechanism field to given value.

### HasIncludedTargetSASLMechanism

`func (o *SimpleRequestCriteriaResponse) HasIncludedTargetSASLMechanism() bool`

HasIncludedTargetSASLMechanism returns a boolean if a field has been set.

### GetExcludedTargetSASLMechanism

`func (o *SimpleRequestCriteriaResponse) GetExcludedTargetSASLMechanism() []string`

GetExcludedTargetSASLMechanism returns the ExcludedTargetSASLMechanism field if non-nil, zero value otherwise.

### GetExcludedTargetSASLMechanismOk

`func (o *SimpleRequestCriteriaResponse) GetExcludedTargetSASLMechanismOk() (*[]string, bool)`

GetExcludedTargetSASLMechanismOk returns a tuple with the ExcludedTargetSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetSASLMechanism

`func (o *SimpleRequestCriteriaResponse) SetExcludedTargetSASLMechanism(v []string)`

SetExcludedTargetSASLMechanism sets ExcludedTargetSASLMechanism field to given value.

### HasExcludedTargetSASLMechanism

`func (o *SimpleRequestCriteriaResponse) HasExcludedTargetSASLMechanism() bool`

HasExcludedTargetSASLMechanism returns a boolean if a field has been set.

### GetIncludedTargetAttribute

`func (o *SimpleRequestCriteriaResponse) GetIncludedTargetAttribute() []string`

GetIncludedTargetAttribute returns the IncludedTargetAttribute field if non-nil, zero value otherwise.

### GetIncludedTargetAttributeOk

`func (o *SimpleRequestCriteriaResponse) GetIncludedTargetAttributeOk() (*[]string, bool)`

GetIncludedTargetAttributeOk returns a tuple with the IncludedTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetAttribute

`func (o *SimpleRequestCriteriaResponse) SetIncludedTargetAttribute(v []string)`

SetIncludedTargetAttribute sets IncludedTargetAttribute field to given value.

### HasIncludedTargetAttribute

`func (o *SimpleRequestCriteriaResponse) HasIncludedTargetAttribute() bool`

HasIncludedTargetAttribute returns a boolean if a field has been set.

### GetExcludedTargetAttribute

`func (o *SimpleRequestCriteriaResponse) GetExcludedTargetAttribute() []string`

GetExcludedTargetAttribute returns the ExcludedTargetAttribute field if non-nil, zero value otherwise.

### GetExcludedTargetAttributeOk

`func (o *SimpleRequestCriteriaResponse) GetExcludedTargetAttributeOk() (*[]string, bool)`

GetExcludedTargetAttributeOk returns a tuple with the ExcludedTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetAttribute

`func (o *SimpleRequestCriteriaResponse) SetExcludedTargetAttribute(v []string)`

SetExcludedTargetAttribute sets ExcludedTargetAttribute field to given value.

### HasExcludedTargetAttribute

`func (o *SimpleRequestCriteriaResponse) HasExcludedTargetAttribute() bool`

HasExcludedTargetAttribute returns a boolean if a field has been set.

### GetIncludedExtendedOperationOID

`func (o *SimpleRequestCriteriaResponse) GetIncludedExtendedOperationOID() []string`

GetIncludedExtendedOperationOID returns the IncludedExtendedOperationOID field if non-nil, zero value otherwise.

### GetIncludedExtendedOperationOIDOk

`func (o *SimpleRequestCriteriaResponse) GetIncludedExtendedOperationOIDOk() (*[]string, bool)`

GetIncludedExtendedOperationOIDOk returns a tuple with the IncludedExtendedOperationOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedExtendedOperationOID

`func (o *SimpleRequestCriteriaResponse) SetIncludedExtendedOperationOID(v []string)`

SetIncludedExtendedOperationOID sets IncludedExtendedOperationOID field to given value.

### HasIncludedExtendedOperationOID

`func (o *SimpleRequestCriteriaResponse) HasIncludedExtendedOperationOID() bool`

HasIncludedExtendedOperationOID returns a boolean if a field has been set.

### GetExcludedExtendedOperationOID

`func (o *SimpleRequestCriteriaResponse) GetExcludedExtendedOperationOID() []string`

GetExcludedExtendedOperationOID returns the ExcludedExtendedOperationOID field if non-nil, zero value otherwise.

### GetExcludedExtendedOperationOIDOk

`func (o *SimpleRequestCriteriaResponse) GetExcludedExtendedOperationOIDOk() (*[]string, bool)`

GetExcludedExtendedOperationOIDOk returns a tuple with the ExcludedExtendedOperationOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedExtendedOperationOID

`func (o *SimpleRequestCriteriaResponse) SetExcludedExtendedOperationOID(v []string)`

SetExcludedExtendedOperationOID sets ExcludedExtendedOperationOID field to given value.

### HasExcludedExtendedOperationOID

`func (o *SimpleRequestCriteriaResponse) HasExcludedExtendedOperationOID() bool`

HasExcludedExtendedOperationOID returns a boolean if a field has been set.

### GetIncludedSearchScope

`func (o *SimpleRequestCriteriaResponse) GetIncludedSearchScope() []EnumrequestCriteriaIncludedSearchScopeProp`

GetIncludedSearchScope returns the IncludedSearchScope field if non-nil, zero value otherwise.

### GetIncludedSearchScopeOk

`func (o *SimpleRequestCriteriaResponse) GetIncludedSearchScopeOk() (*[]EnumrequestCriteriaIncludedSearchScopeProp, bool)`

GetIncludedSearchScopeOk returns a tuple with the IncludedSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSearchScope

`func (o *SimpleRequestCriteriaResponse) SetIncludedSearchScope(v []EnumrequestCriteriaIncludedSearchScopeProp)`

SetIncludedSearchScope sets IncludedSearchScope field to given value.

### HasIncludedSearchScope

`func (o *SimpleRequestCriteriaResponse) HasIncludedSearchScope() bool`

HasIncludedSearchScope returns a boolean if a field has been set.

### GetUsingAdministrativeSessionWorkerThread

`func (o *SimpleRequestCriteriaResponse) GetUsingAdministrativeSessionWorkerThread() EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp`

GetUsingAdministrativeSessionWorkerThread returns the UsingAdministrativeSessionWorkerThread field if non-nil, zero value otherwise.

### GetUsingAdministrativeSessionWorkerThreadOk

`func (o *SimpleRequestCriteriaResponse) GetUsingAdministrativeSessionWorkerThreadOk() (*EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp, bool)`

GetUsingAdministrativeSessionWorkerThreadOk returns a tuple with the UsingAdministrativeSessionWorkerThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingAdministrativeSessionWorkerThread

`func (o *SimpleRequestCriteriaResponse) SetUsingAdministrativeSessionWorkerThread(v EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp)`

SetUsingAdministrativeSessionWorkerThread sets UsingAdministrativeSessionWorkerThread field to given value.

### HasUsingAdministrativeSessionWorkerThread

`func (o *SimpleRequestCriteriaResponse) HasUsingAdministrativeSessionWorkerThread() bool`

HasUsingAdministrativeSessionWorkerThread returns a boolean if a field has been set.

### GetIncludedApplicationName

`func (o *SimpleRequestCriteriaResponse) GetIncludedApplicationName() []string`

GetIncludedApplicationName returns the IncludedApplicationName field if non-nil, zero value otherwise.

### GetIncludedApplicationNameOk

`func (o *SimpleRequestCriteriaResponse) GetIncludedApplicationNameOk() (*[]string, bool)`

GetIncludedApplicationNameOk returns a tuple with the IncludedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedApplicationName

`func (o *SimpleRequestCriteriaResponse) SetIncludedApplicationName(v []string)`

SetIncludedApplicationName sets IncludedApplicationName field to given value.

### HasIncludedApplicationName

`func (o *SimpleRequestCriteriaResponse) HasIncludedApplicationName() bool`

HasIncludedApplicationName returns a boolean if a field has been set.

### GetExcludedApplicationName

`func (o *SimpleRequestCriteriaResponse) GetExcludedApplicationName() []string`

GetExcludedApplicationName returns the ExcludedApplicationName field if non-nil, zero value otherwise.

### GetExcludedApplicationNameOk

`func (o *SimpleRequestCriteriaResponse) GetExcludedApplicationNameOk() (*[]string, bool)`

GetExcludedApplicationNameOk returns a tuple with the ExcludedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedApplicationName

`func (o *SimpleRequestCriteriaResponse) SetExcludedApplicationName(v []string)`

SetExcludedApplicationName sets ExcludedApplicationName field to given value.

### HasExcludedApplicationName

`func (o *SimpleRequestCriteriaResponse) HasExcludedApplicationName() bool`

HasExcludedApplicationName returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleRequestCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleRequestCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleRequestCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleRequestCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


