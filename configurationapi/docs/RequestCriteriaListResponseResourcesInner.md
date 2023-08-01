# RequestCriteriaListResponseResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Request Criteria | 
**Schemas** | [**[]EnumthirdPartyRequestCriteriaSchemaUrn**](EnumthirdPartyRequestCriteriaSchemaUrn.md) |  | 
**OperationType** | Pointer to [**[]EnumrequestCriteriaRootDseOperationTypeProp**](EnumrequestCriteriaRootDseOperationTypeProp.md) |  | [optional] 
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
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**AllIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that must match the associated operation request in order to match the aggregate request criteria. If one or more all-included request criteria objects are provided, then an operation request must match all of them in order to match the aggregate request criteria. | [optional] 
**AnyIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that may match the associated operation request in order to the this aggregate request criteria. If one or more any-included request criteria objects are provided, then an operation request must match at least one of them in order to match the aggregate request criteria. | [optional] 
**NotAllIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that should not match the associated operation request in order to match the aggregate request criteria. If one or more not-all-included request criteria objects are provided, then an operation request must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate request criteria. | [optional] 
**NoneIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that must not match the associated operation request in order to match the aggregate request criteria. If one or more none-included request criteria objects are provided, then an operation request must not match any of them in order to match the aggregate request criteria. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Request Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Request Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewRequestCriteriaListResponseResourcesInner

`func NewRequestCriteriaListResponseResourcesInner(id string, schemas []EnumthirdPartyRequestCriteriaSchemaUrn, extensionClass string, ) *RequestCriteriaListResponseResourcesInner`

NewRequestCriteriaListResponseResourcesInner instantiates a new RequestCriteriaListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCriteriaListResponseResourcesInnerWithDefaults

`func NewRequestCriteriaListResponseResourcesInnerWithDefaults() *RequestCriteriaListResponseResourcesInner`

NewRequestCriteriaListResponseResourcesInnerWithDefaults instantiates a new RequestCriteriaListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestCriteriaListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestCriteriaListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestCriteriaListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *RequestCriteriaListResponseResourcesInner) GetSchemas() []EnumthirdPartyRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *RequestCriteriaListResponseResourcesInner) GetSchemasOk() (*[]EnumthirdPartyRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *RequestCriteriaListResponseResourcesInner) SetSchemas(v []EnumthirdPartyRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOperationType

`func (o *RequestCriteriaListResponseResourcesInner) GetOperationType() []EnumrequestCriteriaRootDseOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *RequestCriteriaListResponseResourcesInner) GetOperationTypeOk() (*[]EnumrequestCriteriaRootDseOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *RequestCriteriaListResponseResourcesInner) SetOperationType(v []EnumrequestCriteriaRootDseOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *RequestCriteriaListResponseResourcesInner) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetOperationOrigin

`func (o *RequestCriteriaListResponseResourcesInner) GetOperationOrigin() []EnumrequestCriteriaOperationOriginProp`

GetOperationOrigin returns the OperationOrigin field if non-nil, zero value otherwise.

### GetOperationOriginOk

`func (o *RequestCriteriaListResponseResourcesInner) GetOperationOriginOk() (*[]EnumrequestCriteriaOperationOriginProp, bool)`

GetOperationOriginOk returns a tuple with the OperationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationOrigin

`func (o *RequestCriteriaListResponseResourcesInner) SetOperationOrigin(v []EnumrequestCriteriaOperationOriginProp)`

SetOperationOrigin sets OperationOrigin field to given value.

### HasOperationOrigin

`func (o *RequestCriteriaListResponseResourcesInner) HasOperationOrigin() bool`

HasOperationOrigin returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *RequestCriteriaListResponseResourcesInner) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *RequestCriteriaListResponseResourcesInner) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *RequestCriteriaListResponseResourcesInner) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *RequestCriteriaListResponseResourcesInner) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetAllIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) GetAllIncludedRequestControl() []string`

GetAllIncludedRequestControl returns the AllIncludedRequestControl field if non-nil, zero value otherwise.

### GetAllIncludedRequestControlOk

`func (o *RequestCriteriaListResponseResourcesInner) GetAllIncludedRequestControlOk() (*[]string, bool)`

GetAllIncludedRequestControlOk returns a tuple with the AllIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) SetAllIncludedRequestControl(v []string)`

SetAllIncludedRequestControl sets AllIncludedRequestControl field to given value.

### HasAllIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) HasAllIncludedRequestControl() bool`

HasAllIncludedRequestControl returns a boolean if a field has been set.

### GetAnyIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) GetAnyIncludedRequestControl() []string`

GetAnyIncludedRequestControl returns the AnyIncludedRequestControl field if non-nil, zero value otherwise.

### GetAnyIncludedRequestControlOk

`func (o *RequestCriteriaListResponseResourcesInner) GetAnyIncludedRequestControlOk() (*[]string, bool)`

GetAnyIncludedRequestControlOk returns a tuple with the AnyIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) SetAnyIncludedRequestControl(v []string)`

SetAnyIncludedRequestControl sets AnyIncludedRequestControl field to given value.

### HasAnyIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) HasAnyIncludedRequestControl() bool`

HasAnyIncludedRequestControl returns a boolean if a field has been set.

### GetNotAllIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) GetNotAllIncludedRequestControl() []string`

GetNotAllIncludedRequestControl returns the NotAllIncludedRequestControl field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestControlOk

`func (o *RequestCriteriaListResponseResourcesInner) GetNotAllIncludedRequestControlOk() (*[]string, bool)`

GetNotAllIncludedRequestControlOk returns a tuple with the NotAllIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) SetNotAllIncludedRequestControl(v []string)`

SetNotAllIncludedRequestControl sets NotAllIncludedRequestControl field to given value.

### HasNotAllIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) HasNotAllIncludedRequestControl() bool`

HasNotAllIncludedRequestControl returns a boolean if a field has been set.

### GetNoneIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) GetNoneIncludedRequestControl() []string`

GetNoneIncludedRequestControl returns the NoneIncludedRequestControl field if non-nil, zero value otherwise.

### GetNoneIncludedRequestControlOk

`func (o *RequestCriteriaListResponseResourcesInner) GetNoneIncludedRequestControlOk() (*[]string, bool)`

GetNoneIncludedRequestControlOk returns a tuple with the NoneIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) SetNoneIncludedRequestControl(v []string)`

SetNoneIncludedRequestControl sets NoneIncludedRequestControl field to given value.

### HasNoneIncludedRequestControl

`func (o *RequestCriteriaListResponseResourcesInner) HasNoneIncludedRequestControl() bool`

HasNoneIncludedRequestControl returns a boolean if a field has been set.

### GetIncludedTargetEntryDN

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedTargetEntryDN() []string`

GetIncludedTargetEntryDN returns the IncludedTargetEntryDN field if non-nil, zero value otherwise.

### GetIncludedTargetEntryDNOk

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedTargetEntryDNOk() (*[]string, bool)`

GetIncludedTargetEntryDNOk returns a tuple with the IncludedTargetEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetEntryDN

`func (o *RequestCriteriaListResponseResourcesInner) SetIncludedTargetEntryDN(v []string)`

SetIncludedTargetEntryDN sets IncludedTargetEntryDN field to given value.

### HasIncludedTargetEntryDN

`func (o *RequestCriteriaListResponseResourcesInner) HasIncludedTargetEntryDN() bool`

HasIncludedTargetEntryDN returns a boolean if a field has been set.

### GetExcludedTargetEntryDN

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedTargetEntryDN() []string`

GetExcludedTargetEntryDN returns the ExcludedTargetEntryDN field if non-nil, zero value otherwise.

### GetExcludedTargetEntryDNOk

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedTargetEntryDNOk() (*[]string, bool)`

GetExcludedTargetEntryDNOk returns a tuple with the ExcludedTargetEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetEntryDN

`func (o *RequestCriteriaListResponseResourcesInner) SetExcludedTargetEntryDN(v []string)`

SetExcludedTargetEntryDN sets ExcludedTargetEntryDN field to given value.

### HasExcludedTargetEntryDN

`func (o *RequestCriteriaListResponseResourcesInner) HasExcludedTargetEntryDN() bool`

HasExcludedTargetEntryDN returns a boolean if a field has been set.

### GetAllIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) GetAllIncludedTargetEntryFilter() []string`

GetAllIncludedTargetEntryFilter returns the AllIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetAllIncludedTargetEntryFilterOk

`func (o *RequestCriteriaListResponseResourcesInner) GetAllIncludedTargetEntryFilterOk() (*[]string, bool)`

GetAllIncludedTargetEntryFilterOk returns a tuple with the AllIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) SetAllIncludedTargetEntryFilter(v []string)`

SetAllIncludedTargetEntryFilter sets AllIncludedTargetEntryFilter field to given value.

### HasAllIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) HasAllIncludedTargetEntryFilter() bool`

HasAllIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetAnyIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) GetAnyIncludedTargetEntryFilter() []string`

GetAnyIncludedTargetEntryFilter returns the AnyIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetAnyIncludedTargetEntryFilterOk

`func (o *RequestCriteriaListResponseResourcesInner) GetAnyIncludedTargetEntryFilterOk() (*[]string, bool)`

GetAnyIncludedTargetEntryFilterOk returns a tuple with the AnyIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) SetAnyIncludedTargetEntryFilter(v []string)`

SetAnyIncludedTargetEntryFilter sets AnyIncludedTargetEntryFilter field to given value.

### HasAnyIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) HasAnyIncludedTargetEntryFilter() bool`

HasAnyIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetNotAllIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) GetNotAllIncludedTargetEntryFilter() []string`

GetNotAllIncludedTargetEntryFilter returns the NotAllIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedTargetEntryFilterOk

`func (o *RequestCriteriaListResponseResourcesInner) GetNotAllIncludedTargetEntryFilterOk() (*[]string, bool)`

GetNotAllIncludedTargetEntryFilterOk returns a tuple with the NotAllIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) SetNotAllIncludedTargetEntryFilter(v []string)`

SetNotAllIncludedTargetEntryFilter sets NotAllIncludedTargetEntryFilter field to given value.

### HasNotAllIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) HasNotAllIncludedTargetEntryFilter() bool`

HasNotAllIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetNoneIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) GetNoneIncludedTargetEntryFilter() []string`

GetNoneIncludedTargetEntryFilter returns the NoneIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetNoneIncludedTargetEntryFilterOk

`func (o *RequestCriteriaListResponseResourcesInner) GetNoneIncludedTargetEntryFilterOk() (*[]string, bool)`

GetNoneIncludedTargetEntryFilterOk returns a tuple with the NoneIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) SetNoneIncludedTargetEntryFilter(v []string)`

SetNoneIncludedTargetEntryFilter sets NoneIncludedTargetEntryFilter field to given value.

### HasNoneIncludedTargetEntryFilter

`func (o *RequestCriteriaListResponseResourcesInner) HasNoneIncludedTargetEntryFilter() bool`

HasNoneIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetAllIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) GetAllIncludedTargetEntryGroupDN() []string`

GetAllIncludedTargetEntryGroupDN returns the AllIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedTargetEntryGroupDNOk

`func (o *RequestCriteriaListResponseResourcesInner) GetAllIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetAllIncludedTargetEntryGroupDNOk returns a tuple with the AllIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) SetAllIncludedTargetEntryGroupDN(v []string)`

SetAllIncludedTargetEntryGroupDN sets AllIncludedTargetEntryGroupDN field to given value.

### HasAllIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) HasAllIncludedTargetEntryGroupDN() bool`

HasAllIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetAnyIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) GetAnyIncludedTargetEntryGroupDN() []string`

GetAnyIncludedTargetEntryGroupDN returns the AnyIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedTargetEntryGroupDNOk

`func (o *RequestCriteriaListResponseResourcesInner) GetAnyIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetAnyIncludedTargetEntryGroupDNOk returns a tuple with the AnyIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) SetAnyIncludedTargetEntryGroupDN(v []string)`

SetAnyIncludedTargetEntryGroupDN sets AnyIncludedTargetEntryGroupDN field to given value.

### HasAnyIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) HasAnyIncludedTargetEntryGroupDN() bool`

HasAnyIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) GetNotAllIncludedTargetEntryGroupDN() []string`

GetNotAllIncludedTargetEntryGroupDN returns the NotAllIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedTargetEntryGroupDNOk

`func (o *RequestCriteriaListResponseResourcesInner) GetNotAllIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetNotAllIncludedTargetEntryGroupDNOk returns a tuple with the NotAllIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) SetNotAllIncludedTargetEntryGroupDN(v []string)`

SetNotAllIncludedTargetEntryGroupDN sets NotAllIncludedTargetEntryGroupDN field to given value.

### HasNotAllIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) HasNotAllIncludedTargetEntryGroupDN() bool`

HasNotAllIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetNoneIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) GetNoneIncludedTargetEntryGroupDN() []string`

GetNoneIncludedTargetEntryGroupDN returns the NoneIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedTargetEntryGroupDNOk

`func (o *RequestCriteriaListResponseResourcesInner) GetNoneIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetNoneIncludedTargetEntryGroupDNOk returns a tuple with the NoneIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) SetNoneIncludedTargetEntryGroupDN(v []string)`

SetNoneIncludedTargetEntryGroupDN sets NoneIncludedTargetEntryGroupDN field to given value.

### HasNoneIncludedTargetEntryGroupDN

`func (o *RequestCriteriaListResponseResourcesInner) HasNoneIncludedTargetEntryGroupDN() bool`

HasNoneIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetTargetBindType

`func (o *RequestCriteriaListResponseResourcesInner) GetTargetBindType() []EnumrequestCriteriaTargetBindTypeProp`

GetTargetBindType returns the TargetBindType field if non-nil, zero value otherwise.

### GetTargetBindTypeOk

`func (o *RequestCriteriaListResponseResourcesInner) GetTargetBindTypeOk() (*[]EnumrequestCriteriaTargetBindTypeProp, bool)`

GetTargetBindTypeOk returns a tuple with the TargetBindType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBindType

`func (o *RequestCriteriaListResponseResourcesInner) SetTargetBindType(v []EnumrequestCriteriaTargetBindTypeProp)`

SetTargetBindType sets TargetBindType field to given value.

### HasTargetBindType

`func (o *RequestCriteriaListResponseResourcesInner) HasTargetBindType() bool`

HasTargetBindType returns a boolean if a field has been set.

### GetIncludedTargetSASLMechanism

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedTargetSASLMechanism() []string`

GetIncludedTargetSASLMechanism returns the IncludedTargetSASLMechanism field if non-nil, zero value otherwise.

### GetIncludedTargetSASLMechanismOk

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedTargetSASLMechanismOk() (*[]string, bool)`

GetIncludedTargetSASLMechanismOk returns a tuple with the IncludedTargetSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetSASLMechanism

`func (o *RequestCriteriaListResponseResourcesInner) SetIncludedTargetSASLMechanism(v []string)`

SetIncludedTargetSASLMechanism sets IncludedTargetSASLMechanism field to given value.

### HasIncludedTargetSASLMechanism

`func (o *RequestCriteriaListResponseResourcesInner) HasIncludedTargetSASLMechanism() bool`

HasIncludedTargetSASLMechanism returns a boolean if a field has been set.

### GetExcludedTargetSASLMechanism

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedTargetSASLMechanism() []string`

GetExcludedTargetSASLMechanism returns the ExcludedTargetSASLMechanism field if non-nil, zero value otherwise.

### GetExcludedTargetSASLMechanismOk

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedTargetSASLMechanismOk() (*[]string, bool)`

GetExcludedTargetSASLMechanismOk returns a tuple with the ExcludedTargetSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetSASLMechanism

`func (o *RequestCriteriaListResponseResourcesInner) SetExcludedTargetSASLMechanism(v []string)`

SetExcludedTargetSASLMechanism sets ExcludedTargetSASLMechanism field to given value.

### HasExcludedTargetSASLMechanism

`func (o *RequestCriteriaListResponseResourcesInner) HasExcludedTargetSASLMechanism() bool`

HasExcludedTargetSASLMechanism returns a boolean if a field has been set.

### GetIncludedTargetAttribute

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedTargetAttribute() []string`

GetIncludedTargetAttribute returns the IncludedTargetAttribute field if non-nil, zero value otherwise.

### GetIncludedTargetAttributeOk

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedTargetAttributeOk() (*[]string, bool)`

GetIncludedTargetAttributeOk returns a tuple with the IncludedTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetAttribute

`func (o *RequestCriteriaListResponseResourcesInner) SetIncludedTargetAttribute(v []string)`

SetIncludedTargetAttribute sets IncludedTargetAttribute field to given value.

### HasIncludedTargetAttribute

`func (o *RequestCriteriaListResponseResourcesInner) HasIncludedTargetAttribute() bool`

HasIncludedTargetAttribute returns a boolean if a field has been set.

### GetExcludedTargetAttribute

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedTargetAttribute() []string`

GetExcludedTargetAttribute returns the ExcludedTargetAttribute field if non-nil, zero value otherwise.

### GetExcludedTargetAttributeOk

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedTargetAttributeOk() (*[]string, bool)`

GetExcludedTargetAttributeOk returns a tuple with the ExcludedTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetAttribute

`func (o *RequestCriteriaListResponseResourcesInner) SetExcludedTargetAttribute(v []string)`

SetExcludedTargetAttribute sets ExcludedTargetAttribute field to given value.

### HasExcludedTargetAttribute

`func (o *RequestCriteriaListResponseResourcesInner) HasExcludedTargetAttribute() bool`

HasExcludedTargetAttribute returns a boolean if a field has been set.

### GetIncludedExtendedOperationOID

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedExtendedOperationOID() []string`

GetIncludedExtendedOperationOID returns the IncludedExtendedOperationOID field if non-nil, zero value otherwise.

### GetIncludedExtendedOperationOIDOk

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedExtendedOperationOIDOk() (*[]string, bool)`

GetIncludedExtendedOperationOIDOk returns a tuple with the IncludedExtendedOperationOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedExtendedOperationOID

`func (o *RequestCriteriaListResponseResourcesInner) SetIncludedExtendedOperationOID(v []string)`

SetIncludedExtendedOperationOID sets IncludedExtendedOperationOID field to given value.

### HasIncludedExtendedOperationOID

`func (o *RequestCriteriaListResponseResourcesInner) HasIncludedExtendedOperationOID() bool`

HasIncludedExtendedOperationOID returns a boolean if a field has been set.

### GetExcludedExtendedOperationOID

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedExtendedOperationOID() []string`

GetExcludedExtendedOperationOID returns the ExcludedExtendedOperationOID field if non-nil, zero value otherwise.

### GetExcludedExtendedOperationOIDOk

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedExtendedOperationOIDOk() (*[]string, bool)`

GetExcludedExtendedOperationOIDOk returns a tuple with the ExcludedExtendedOperationOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedExtendedOperationOID

`func (o *RequestCriteriaListResponseResourcesInner) SetExcludedExtendedOperationOID(v []string)`

SetExcludedExtendedOperationOID sets ExcludedExtendedOperationOID field to given value.

### HasExcludedExtendedOperationOID

`func (o *RequestCriteriaListResponseResourcesInner) HasExcludedExtendedOperationOID() bool`

HasExcludedExtendedOperationOID returns a boolean if a field has been set.

### GetIncludedSearchScope

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedSearchScope() []EnumrequestCriteriaIncludedSearchScopeProp`

GetIncludedSearchScope returns the IncludedSearchScope field if non-nil, zero value otherwise.

### GetIncludedSearchScopeOk

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedSearchScopeOk() (*[]EnumrequestCriteriaIncludedSearchScopeProp, bool)`

GetIncludedSearchScopeOk returns a tuple with the IncludedSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSearchScope

`func (o *RequestCriteriaListResponseResourcesInner) SetIncludedSearchScope(v []EnumrequestCriteriaIncludedSearchScopeProp)`

SetIncludedSearchScope sets IncludedSearchScope field to given value.

### HasIncludedSearchScope

`func (o *RequestCriteriaListResponseResourcesInner) HasIncludedSearchScope() bool`

HasIncludedSearchScope returns a boolean if a field has been set.

### GetUsingAdministrativeSessionWorkerThread

`func (o *RequestCriteriaListResponseResourcesInner) GetUsingAdministrativeSessionWorkerThread() EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp`

GetUsingAdministrativeSessionWorkerThread returns the UsingAdministrativeSessionWorkerThread field if non-nil, zero value otherwise.

### GetUsingAdministrativeSessionWorkerThreadOk

`func (o *RequestCriteriaListResponseResourcesInner) GetUsingAdministrativeSessionWorkerThreadOk() (*EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp, bool)`

GetUsingAdministrativeSessionWorkerThreadOk returns a tuple with the UsingAdministrativeSessionWorkerThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingAdministrativeSessionWorkerThread

`func (o *RequestCriteriaListResponseResourcesInner) SetUsingAdministrativeSessionWorkerThread(v EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp)`

SetUsingAdministrativeSessionWorkerThread sets UsingAdministrativeSessionWorkerThread field to given value.

### HasUsingAdministrativeSessionWorkerThread

`func (o *RequestCriteriaListResponseResourcesInner) HasUsingAdministrativeSessionWorkerThread() bool`

HasUsingAdministrativeSessionWorkerThread returns a boolean if a field has been set.

### GetIncludedApplicationName

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedApplicationName() []string`

GetIncludedApplicationName returns the IncludedApplicationName field if non-nil, zero value otherwise.

### GetIncludedApplicationNameOk

`func (o *RequestCriteriaListResponseResourcesInner) GetIncludedApplicationNameOk() (*[]string, bool)`

GetIncludedApplicationNameOk returns a tuple with the IncludedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedApplicationName

`func (o *RequestCriteriaListResponseResourcesInner) SetIncludedApplicationName(v []string)`

SetIncludedApplicationName sets IncludedApplicationName field to given value.

### HasIncludedApplicationName

`func (o *RequestCriteriaListResponseResourcesInner) HasIncludedApplicationName() bool`

HasIncludedApplicationName returns a boolean if a field has been set.

### GetExcludedApplicationName

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedApplicationName() []string`

GetExcludedApplicationName returns the ExcludedApplicationName field if non-nil, zero value otherwise.

### GetExcludedApplicationNameOk

`func (o *RequestCriteriaListResponseResourcesInner) GetExcludedApplicationNameOk() (*[]string, bool)`

GetExcludedApplicationNameOk returns a tuple with the ExcludedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedApplicationName

`func (o *RequestCriteriaListResponseResourcesInner) SetExcludedApplicationName(v []string)`

SetExcludedApplicationName sets ExcludedApplicationName field to given value.

### HasExcludedApplicationName

`func (o *RequestCriteriaListResponseResourcesInner) HasExcludedApplicationName() bool`

HasExcludedApplicationName returns a boolean if a field has been set.

### GetDescription

`func (o *RequestCriteriaListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestCriteriaListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestCriteriaListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestCriteriaListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *RequestCriteriaListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RequestCriteriaListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RequestCriteriaListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RequestCriteriaListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *RequestCriteriaListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *RequestCriteriaListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *RequestCriteriaListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *RequestCriteriaListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetAllIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) GetAllIncludedRequestCriteria() []string`

GetAllIncludedRequestCriteria returns the AllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAllIncludedRequestCriteriaOk

`func (o *RequestCriteriaListResponseResourcesInner) GetAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetAllIncludedRequestCriteriaOk returns a tuple with the AllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) SetAllIncludedRequestCriteria(v []string)`

SetAllIncludedRequestCriteria sets AllIncludedRequestCriteria field to given value.

### HasAllIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) HasAllIncludedRequestCriteria() bool`

HasAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetAnyIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) GetAnyIncludedRequestCriteria() []string`

GetAnyIncludedRequestCriteria returns the AnyIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedRequestCriteriaOk

`func (o *RequestCriteriaListResponseResourcesInner) GetAnyIncludedRequestCriteriaOk() (*[]string, bool)`

GetAnyIncludedRequestCriteriaOk returns a tuple with the AnyIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) SetAnyIncludedRequestCriteria(v []string)`

SetAnyIncludedRequestCriteria sets AnyIncludedRequestCriteria field to given value.

### HasAnyIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) HasAnyIncludedRequestCriteria() bool`

HasAnyIncludedRequestCriteria returns a boolean if a field has been set.

### GetNotAllIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) GetNotAllIncludedRequestCriteria() []string`

GetNotAllIncludedRequestCriteria returns the NotAllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestCriteriaOk

`func (o *RequestCriteriaListResponseResourcesInner) GetNotAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetNotAllIncludedRequestCriteriaOk returns a tuple with the NotAllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) SetNotAllIncludedRequestCriteria(v []string)`

SetNotAllIncludedRequestCriteria sets NotAllIncludedRequestCriteria field to given value.

### HasNotAllIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) HasNotAllIncludedRequestCriteria() bool`

HasNotAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetNoneIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) GetNoneIncludedRequestCriteria() []string`

GetNoneIncludedRequestCriteria returns the NoneIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedRequestCriteriaOk

`func (o *RequestCriteriaListResponseResourcesInner) GetNoneIncludedRequestCriteriaOk() (*[]string, bool)`

GetNoneIncludedRequestCriteriaOk returns a tuple with the NoneIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) SetNoneIncludedRequestCriteria(v []string)`

SetNoneIncludedRequestCriteria sets NoneIncludedRequestCriteria field to given value.

### HasNoneIncludedRequestCriteria

`func (o *RequestCriteriaListResponseResourcesInner) HasNoneIncludedRequestCriteria() bool`

HasNoneIncludedRequestCriteria returns a boolean if a field has been set.

### GetExtensionClass

`func (o *RequestCriteriaListResponseResourcesInner) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *RequestCriteriaListResponseResourcesInner) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *RequestCriteriaListResponseResourcesInner) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *RequestCriteriaListResponseResourcesInner) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *RequestCriteriaListResponseResourcesInner) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *RequestCriteriaListResponseResourcesInner) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *RequestCriteriaListResponseResourcesInner) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


