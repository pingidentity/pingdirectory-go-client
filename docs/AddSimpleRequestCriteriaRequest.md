# AddSimpleRequestCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriteriaName** | **string** | Name of the new Request Criteria | 
**Schemas** | [**[]EnumsimpleRequestCriteriaSchemaUrn**](EnumsimpleRequestCriteriaSchemaUrn.md) |  | 
**OperationType** | Pointer to [**[]EnumrequestCriteriaOperationTypeProp**](EnumrequestCriteriaOperationTypeProp.md) |  | [optional] 
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

### NewAddSimpleRequestCriteriaRequest

`func NewAddSimpleRequestCriteriaRequest(criteriaName string, schemas []EnumsimpleRequestCriteriaSchemaUrn, ) *AddSimpleRequestCriteriaRequest`

NewAddSimpleRequestCriteriaRequest instantiates a new AddSimpleRequestCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSimpleRequestCriteriaRequestWithDefaults

`func NewAddSimpleRequestCriteriaRequestWithDefaults() *AddSimpleRequestCriteriaRequest`

NewAddSimpleRequestCriteriaRequestWithDefaults instantiates a new AddSimpleRequestCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteriaName

`func (o *AddSimpleRequestCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddSimpleRequestCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddSimpleRequestCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.


### GetSchemas

`func (o *AddSimpleRequestCriteriaRequest) GetSchemas() []EnumsimpleRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSimpleRequestCriteriaRequest) GetSchemasOk() (*[]EnumsimpleRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSimpleRequestCriteriaRequest) SetSchemas(v []EnumsimpleRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOperationType

`func (o *AddSimpleRequestCriteriaRequest) GetOperationType() []EnumrequestCriteriaOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *AddSimpleRequestCriteriaRequest) GetOperationTypeOk() (*[]EnumrequestCriteriaOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *AddSimpleRequestCriteriaRequest) SetOperationType(v []EnumrequestCriteriaOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *AddSimpleRequestCriteriaRequest) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetOperationOrigin

`func (o *AddSimpleRequestCriteriaRequest) GetOperationOrigin() []EnumrequestCriteriaOperationOriginProp`

GetOperationOrigin returns the OperationOrigin field if non-nil, zero value otherwise.

### GetOperationOriginOk

`func (o *AddSimpleRequestCriteriaRequest) GetOperationOriginOk() (*[]EnumrequestCriteriaOperationOriginProp, bool)`

GetOperationOriginOk returns a tuple with the OperationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationOrigin

`func (o *AddSimpleRequestCriteriaRequest) SetOperationOrigin(v []EnumrequestCriteriaOperationOriginProp)`

SetOperationOrigin sets OperationOrigin field to given value.

### HasOperationOrigin

`func (o *AddSimpleRequestCriteriaRequest) HasOperationOrigin() bool`

HasOperationOrigin returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddSimpleRequestCriteriaRequest) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddSimpleRequestCriteriaRequest) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddSimpleRequestCriteriaRequest) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddSimpleRequestCriteriaRequest) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetAllIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) GetAllIncludedRequestControl() []string`

GetAllIncludedRequestControl returns the AllIncludedRequestControl field if non-nil, zero value otherwise.

### GetAllIncludedRequestControlOk

`func (o *AddSimpleRequestCriteriaRequest) GetAllIncludedRequestControlOk() (*[]string, bool)`

GetAllIncludedRequestControlOk returns a tuple with the AllIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) SetAllIncludedRequestControl(v []string)`

SetAllIncludedRequestControl sets AllIncludedRequestControl field to given value.

### HasAllIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) HasAllIncludedRequestControl() bool`

HasAllIncludedRequestControl returns a boolean if a field has been set.

### GetAnyIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) GetAnyIncludedRequestControl() []string`

GetAnyIncludedRequestControl returns the AnyIncludedRequestControl field if non-nil, zero value otherwise.

### GetAnyIncludedRequestControlOk

`func (o *AddSimpleRequestCriteriaRequest) GetAnyIncludedRequestControlOk() (*[]string, bool)`

GetAnyIncludedRequestControlOk returns a tuple with the AnyIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) SetAnyIncludedRequestControl(v []string)`

SetAnyIncludedRequestControl sets AnyIncludedRequestControl field to given value.

### HasAnyIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) HasAnyIncludedRequestControl() bool`

HasAnyIncludedRequestControl returns a boolean if a field has been set.

### GetNotAllIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) GetNotAllIncludedRequestControl() []string`

GetNotAllIncludedRequestControl returns the NotAllIncludedRequestControl field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestControlOk

`func (o *AddSimpleRequestCriteriaRequest) GetNotAllIncludedRequestControlOk() (*[]string, bool)`

GetNotAllIncludedRequestControlOk returns a tuple with the NotAllIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) SetNotAllIncludedRequestControl(v []string)`

SetNotAllIncludedRequestControl sets NotAllIncludedRequestControl field to given value.

### HasNotAllIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) HasNotAllIncludedRequestControl() bool`

HasNotAllIncludedRequestControl returns a boolean if a field has been set.

### GetNoneIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) GetNoneIncludedRequestControl() []string`

GetNoneIncludedRequestControl returns the NoneIncludedRequestControl field if non-nil, zero value otherwise.

### GetNoneIncludedRequestControlOk

`func (o *AddSimpleRequestCriteriaRequest) GetNoneIncludedRequestControlOk() (*[]string, bool)`

GetNoneIncludedRequestControlOk returns a tuple with the NoneIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) SetNoneIncludedRequestControl(v []string)`

SetNoneIncludedRequestControl sets NoneIncludedRequestControl field to given value.

### HasNoneIncludedRequestControl

`func (o *AddSimpleRequestCriteriaRequest) HasNoneIncludedRequestControl() bool`

HasNoneIncludedRequestControl returns a boolean if a field has been set.

### GetIncludedTargetEntryDN

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedTargetEntryDN() []string`

GetIncludedTargetEntryDN returns the IncludedTargetEntryDN field if non-nil, zero value otherwise.

### GetIncludedTargetEntryDNOk

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedTargetEntryDNOk() (*[]string, bool)`

GetIncludedTargetEntryDNOk returns a tuple with the IncludedTargetEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetEntryDN

`func (o *AddSimpleRequestCriteriaRequest) SetIncludedTargetEntryDN(v []string)`

SetIncludedTargetEntryDN sets IncludedTargetEntryDN field to given value.

### HasIncludedTargetEntryDN

`func (o *AddSimpleRequestCriteriaRequest) HasIncludedTargetEntryDN() bool`

HasIncludedTargetEntryDN returns a boolean if a field has been set.

### GetExcludedTargetEntryDN

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedTargetEntryDN() []string`

GetExcludedTargetEntryDN returns the ExcludedTargetEntryDN field if non-nil, zero value otherwise.

### GetExcludedTargetEntryDNOk

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedTargetEntryDNOk() (*[]string, bool)`

GetExcludedTargetEntryDNOk returns a tuple with the ExcludedTargetEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetEntryDN

`func (o *AddSimpleRequestCriteriaRequest) SetExcludedTargetEntryDN(v []string)`

SetExcludedTargetEntryDN sets ExcludedTargetEntryDN field to given value.

### HasExcludedTargetEntryDN

`func (o *AddSimpleRequestCriteriaRequest) HasExcludedTargetEntryDN() bool`

HasExcludedTargetEntryDN returns a boolean if a field has been set.

### GetAllIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) GetAllIncludedTargetEntryFilter() []string`

GetAllIncludedTargetEntryFilter returns the AllIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetAllIncludedTargetEntryFilterOk

`func (o *AddSimpleRequestCriteriaRequest) GetAllIncludedTargetEntryFilterOk() (*[]string, bool)`

GetAllIncludedTargetEntryFilterOk returns a tuple with the AllIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) SetAllIncludedTargetEntryFilter(v []string)`

SetAllIncludedTargetEntryFilter sets AllIncludedTargetEntryFilter field to given value.

### HasAllIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) HasAllIncludedTargetEntryFilter() bool`

HasAllIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetAnyIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) GetAnyIncludedTargetEntryFilter() []string`

GetAnyIncludedTargetEntryFilter returns the AnyIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetAnyIncludedTargetEntryFilterOk

`func (o *AddSimpleRequestCriteriaRequest) GetAnyIncludedTargetEntryFilterOk() (*[]string, bool)`

GetAnyIncludedTargetEntryFilterOk returns a tuple with the AnyIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) SetAnyIncludedTargetEntryFilter(v []string)`

SetAnyIncludedTargetEntryFilter sets AnyIncludedTargetEntryFilter field to given value.

### HasAnyIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) HasAnyIncludedTargetEntryFilter() bool`

HasAnyIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetNotAllIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) GetNotAllIncludedTargetEntryFilter() []string`

GetNotAllIncludedTargetEntryFilter returns the NotAllIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedTargetEntryFilterOk

`func (o *AddSimpleRequestCriteriaRequest) GetNotAllIncludedTargetEntryFilterOk() (*[]string, bool)`

GetNotAllIncludedTargetEntryFilterOk returns a tuple with the NotAllIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) SetNotAllIncludedTargetEntryFilter(v []string)`

SetNotAllIncludedTargetEntryFilter sets NotAllIncludedTargetEntryFilter field to given value.

### HasNotAllIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) HasNotAllIncludedTargetEntryFilter() bool`

HasNotAllIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetNoneIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) GetNoneIncludedTargetEntryFilter() []string`

GetNoneIncludedTargetEntryFilter returns the NoneIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetNoneIncludedTargetEntryFilterOk

`func (o *AddSimpleRequestCriteriaRequest) GetNoneIncludedTargetEntryFilterOk() (*[]string, bool)`

GetNoneIncludedTargetEntryFilterOk returns a tuple with the NoneIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) SetNoneIncludedTargetEntryFilter(v []string)`

SetNoneIncludedTargetEntryFilter sets NoneIncludedTargetEntryFilter field to given value.

### HasNoneIncludedTargetEntryFilter

`func (o *AddSimpleRequestCriteriaRequest) HasNoneIncludedTargetEntryFilter() bool`

HasNoneIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetAllIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) GetAllIncludedTargetEntryGroupDN() []string`

GetAllIncludedTargetEntryGroupDN returns the AllIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedTargetEntryGroupDNOk

`func (o *AddSimpleRequestCriteriaRequest) GetAllIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetAllIncludedTargetEntryGroupDNOk returns a tuple with the AllIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) SetAllIncludedTargetEntryGroupDN(v []string)`

SetAllIncludedTargetEntryGroupDN sets AllIncludedTargetEntryGroupDN field to given value.

### HasAllIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) HasAllIncludedTargetEntryGroupDN() bool`

HasAllIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetAnyIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) GetAnyIncludedTargetEntryGroupDN() []string`

GetAnyIncludedTargetEntryGroupDN returns the AnyIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedTargetEntryGroupDNOk

`func (o *AddSimpleRequestCriteriaRequest) GetAnyIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetAnyIncludedTargetEntryGroupDNOk returns a tuple with the AnyIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) SetAnyIncludedTargetEntryGroupDN(v []string)`

SetAnyIncludedTargetEntryGroupDN sets AnyIncludedTargetEntryGroupDN field to given value.

### HasAnyIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) HasAnyIncludedTargetEntryGroupDN() bool`

HasAnyIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) GetNotAllIncludedTargetEntryGroupDN() []string`

GetNotAllIncludedTargetEntryGroupDN returns the NotAllIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedTargetEntryGroupDNOk

`func (o *AddSimpleRequestCriteriaRequest) GetNotAllIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetNotAllIncludedTargetEntryGroupDNOk returns a tuple with the NotAllIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) SetNotAllIncludedTargetEntryGroupDN(v []string)`

SetNotAllIncludedTargetEntryGroupDN sets NotAllIncludedTargetEntryGroupDN field to given value.

### HasNotAllIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) HasNotAllIncludedTargetEntryGroupDN() bool`

HasNotAllIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetNoneIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) GetNoneIncludedTargetEntryGroupDN() []string`

GetNoneIncludedTargetEntryGroupDN returns the NoneIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedTargetEntryGroupDNOk

`func (o *AddSimpleRequestCriteriaRequest) GetNoneIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetNoneIncludedTargetEntryGroupDNOk returns a tuple with the NoneIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) SetNoneIncludedTargetEntryGroupDN(v []string)`

SetNoneIncludedTargetEntryGroupDN sets NoneIncludedTargetEntryGroupDN field to given value.

### HasNoneIncludedTargetEntryGroupDN

`func (o *AddSimpleRequestCriteriaRequest) HasNoneIncludedTargetEntryGroupDN() bool`

HasNoneIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetTargetBindType

`func (o *AddSimpleRequestCriteriaRequest) GetTargetBindType() []EnumrequestCriteriaTargetBindTypeProp`

GetTargetBindType returns the TargetBindType field if non-nil, zero value otherwise.

### GetTargetBindTypeOk

`func (o *AddSimpleRequestCriteriaRequest) GetTargetBindTypeOk() (*[]EnumrequestCriteriaTargetBindTypeProp, bool)`

GetTargetBindTypeOk returns a tuple with the TargetBindType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBindType

`func (o *AddSimpleRequestCriteriaRequest) SetTargetBindType(v []EnumrequestCriteriaTargetBindTypeProp)`

SetTargetBindType sets TargetBindType field to given value.

### HasTargetBindType

`func (o *AddSimpleRequestCriteriaRequest) HasTargetBindType() bool`

HasTargetBindType returns a boolean if a field has been set.

### GetIncludedTargetSASLMechanism

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedTargetSASLMechanism() []string`

GetIncludedTargetSASLMechanism returns the IncludedTargetSASLMechanism field if non-nil, zero value otherwise.

### GetIncludedTargetSASLMechanismOk

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedTargetSASLMechanismOk() (*[]string, bool)`

GetIncludedTargetSASLMechanismOk returns a tuple with the IncludedTargetSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetSASLMechanism

`func (o *AddSimpleRequestCriteriaRequest) SetIncludedTargetSASLMechanism(v []string)`

SetIncludedTargetSASLMechanism sets IncludedTargetSASLMechanism field to given value.

### HasIncludedTargetSASLMechanism

`func (o *AddSimpleRequestCriteriaRequest) HasIncludedTargetSASLMechanism() bool`

HasIncludedTargetSASLMechanism returns a boolean if a field has been set.

### GetExcludedTargetSASLMechanism

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedTargetSASLMechanism() []string`

GetExcludedTargetSASLMechanism returns the ExcludedTargetSASLMechanism field if non-nil, zero value otherwise.

### GetExcludedTargetSASLMechanismOk

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedTargetSASLMechanismOk() (*[]string, bool)`

GetExcludedTargetSASLMechanismOk returns a tuple with the ExcludedTargetSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetSASLMechanism

`func (o *AddSimpleRequestCriteriaRequest) SetExcludedTargetSASLMechanism(v []string)`

SetExcludedTargetSASLMechanism sets ExcludedTargetSASLMechanism field to given value.

### HasExcludedTargetSASLMechanism

`func (o *AddSimpleRequestCriteriaRequest) HasExcludedTargetSASLMechanism() bool`

HasExcludedTargetSASLMechanism returns a boolean if a field has been set.

### GetIncludedTargetAttribute

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedTargetAttribute() []string`

GetIncludedTargetAttribute returns the IncludedTargetAttribute field if non-nil, zero value otherwise.

### GetIncludedTargetAttributeOk

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedTargetAttributeOk() (*[]string, bool)`

GetIncludedTargetAttributeOk returns a tuple with the IncludedTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetAttribute

`func (o *AddSimpleRequestCriteriaRequest) SetIncludedTargetAttribute(v []string)`

SetIncludedTargetAttribute sets IncludedTargetAttribute field to given value.

### HasIncludedTargetAttribute

`func (o *AddSimpleRequestCriteriaRequest) HasIncludedTargetAttribute() bool`

HasIncludedTargetAttribute returns a boolean if a field has been set.

### GetExcludedTargetAttribute

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedTargetAttribute() []string`

GetExcludedTargetAttribute returns the ExcludedTargetAttribute field if non-nil, zero value otherwise.

### GetExcludedTargetAttributeOk

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedTargetAttributeOk() (*[]string, bool)`

GetExcludedTargetAttributeOk returns a tuple with the ExcludedTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetAttribute

`func (o *AddSimpleRequestCriteriaRequest) SetExcludedTargetAttribute(v []string)`

SetExcludedTargetAttribute sets ExcludedTargetAttribute field to given value.

### HasExcludedTargetAttribute

`func (o *AddSimpleRequestCriteriaRequest) HasExcludedTargetAttribute() bool`

HasExcludedTargetAttribute returns a boolean if a field has been set.

### GetIncludedExtendedOperationOID

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedExtendedOperationOID() []string`

GetIncludedExtendedOperationOID returns the IncludedExtendedOperationOID field if non-nil, zero value otherwise.

### GetIncludedExtendedOperationOIDOk

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedExtendedOperationOIDOk() (*[]string, bool)`

GetIncludedExtendedOperationOIDOk returns a tuple with the IncludedExtendedOperationOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedExtendedOperationOID

`func (o *AddSimpleRequestCriteriaRequest) SetIncludedExtendedOperationOID(v []string)`

SetIncludedExtendedOperationOID sets IncludedExtendedOperationOID field to given value.

### HasIncludedExtendedOperationOID

`func (o *AddSimpleRequestCriteriaRequest) HasIncludedExtendedOperationOID() bool`

HasIncludedExtendedOperationOID returns a boolean if a field has been set.

### GetExcludedExtendedOperationOID

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedExtendedOperationOID() []string`

GetExcludedExtendedOperationOID returns the ExcludedExtendedOperationOID field if non-nil, zero value otherwise.

### GetExcludedExtendedOperationOIDOk

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedExtendedOperationOIDOk() (*[]string, bool)`

GetExcludedExtendedOperationOIDOk returns a tuple with the ExcludedExtendedOperationOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedExtendedOperationOID

`func (o *AddSimpleRequestCriteriaRequest) SetExcludedExtendedOperationOID(v []string)`

SetExcludedExtendedOperationOID sets ExcludedExtendedOperationOID field to given value.

### HasExcludedExtendedOperationOID

`func (o *AddSimpleRequestCriteriaRequest) HasExcludedExtendedOperationOID() bool`

HasExcludedExtendedOperationOID returns a boolean if a field has been set.

### GetIncludedSearchScope

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedSearchScope() []EnumrequestCriteriaIncludedSearchScopeProp`

GetIncludedSearchScope returns the IncludedSearchScope field if non-nil, zero value otherwise.

### GetIncludedSearchScopeOk

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedSearchScopeOk() (*[]EnumrequestCriteriaIncludedSearchScopeProp, bool)`

GetIncludedSearchScopeOk returns a tuple with the IncludedSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSearchScope

`func (o *AddSimpleRequestCriteriaRequest) SetIncludedSearchScope(v []EnumrequestCriteriaIncludedSearchScopeProp)`

SetIncludedSearchScope sets IncludedSearchScope field to given value.

### HasIncludedSearchScope

`func (o *AddSimpleRequestCriteriaRequest) HasIncludedSearchScope() bool`

HasIncludedSearchScope returns a boolean if a field has been set.

### GetUsingAdministrativeSessionWorkerThread

`func (o *AddSimpleRequestCriteriaRequest) GetUsingAdministrativeSessionWorkerThread() EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp`

GetUsingAdministrativeSessionWorkerThread returns the UsingAdministrativeSessionWorkerThread field if non-nil, zero value otherwise.

### GetUsingAdministrativeSessionWorkerThreadOk

`func (o *AddSimpleRequestCriteriaRequest) GetUsingAdministrativeSessionWorkerThreadOk() (*EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp, bool)`

GetUsingAdministrativeSessionWorkerThreadOk returns a tuple with the UsingAdministrativeSessionWorkerThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingAdministrativeSessionWorkerThread

`func (o *AddSimpleRequestCriteriaRequest) SetUsingAdministrativeSessionWorkerThread(v EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp)`

SetUsingAdministrativeSessionWorkerThread sets UsingAdministrativeSessionWorkerThread field to given value.

### HasUsingAdministrativeSessionWorkerThread

`func (o *AddSimpleRequestCriteriaRequest) HasUsingAdministrativeSessionWorkerThread() bool`

HasUsingAdministrativeSessionWorkerThread returns a boolean if a field has been set.

### GetIncludedApplicationName

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedApplicationName() []string`

GetIncludedApplicationName returns the IncludedApplicationName field if non-nil, zero value otherwise.

### GetIncludedApplicationNameOk

`func (o *AddSimpleRequestCriteriaRequest) GetIncludedApplicationNameOk() (*[]string, bool)`

GetIncludedApplicationNameOk returns a tuple with the IncludedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedApplicationName

`func (o *AddSimpleRequestCriteriaRequest) SetIncludedApplicationName(v []string)`

SetIncludedApplicationName sets IncludedApplicationName field to given value.

### HasIncludedApplicationName

`func (o *AddSimpleRequestCriteriaRequest) HasIncludedApplicationName() bool`

HasIncludedApplicationName returns a boolean if a field has been set.

### GetExcludedApplicationName

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedApplicationName() []string`

GetExcludedApplicationName returns the ExcludedApplicationName field if non-nil, zero value otherwise.

### GetExcludedApplicationNameOk

`func (o *AddSimpleRequestCriteriaRequest) GetExcludedApplicationNameOk() (*[]string, bool)`

GetExcludedApplicationNameOk returns a tuple with the ExcludedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedApplicationName

`func (o *AddSimpleRequestCriteriaRequest) SetExcludedApplicationName(v []string)`

SetExcludedApplicationName sets ExcludedApplicationName field to given value.

### HasExcludedApplicationName

`func (o *AddSimpleRequestCriteriaRequest) HasExcludedApplicationName() bool`

HasExcludedApplicationName returns a boolean if a field has been set.

### GetDescription

`func (o *AddSimpleRequestCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSimpleRequestCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSimpleRequestCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSimpleRequestCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


