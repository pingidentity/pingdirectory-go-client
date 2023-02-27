# AddRequestCriteria200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**Id** | **string** | Name of the Request Criteria | 
**Schemas** | [**[]EnumthirdPartyRequestCriteriaSchemaUrn**](EnumthirdPartyRequestCriteriaSchemaUrn.md) |  | 
**OperationType** | Pointer to [**[]EnumrequestCriteriaSimpleOperationTypeProp**](EnumrequestCriteriaSimpleOperationTypeProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Request Criteria | [optional] 
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
**AllIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that must match the associated operation request in order to match the aggregate request criteria. If one or more all-included request criteria objects are provided, then an operation request must match all of them in order to match the aggregate request criteria. | [optional] 
**AnyIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that may match the associated operation request in order to the this aggregate request criteria. If one or more any-included request criteria objects are provided, then an operation request must match at least one of them in order to match the aggregate request criteria. | [optional] 
**NotAllIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that should not match the associated operation request in order to match the aggregate request criteria. If one or more not-all-included request criteria objects are provided, then an operation request must not match all of them (that is, it may match zero or more of them, but it must not match all of them) in order to match the aggregate request criteria. | [optional] 
**NoneIncludedRequestCriteria** | Pointer to **[]string** | Specifies a request criteria object that must not match the associated operation request in order to match the aggregate request criteria. If one or more none-included request criteria objects are provided, then an operation request must not match any of them in order to match the aggregate request criteria. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Request Criteria. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Request Criteria. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 

## Methods

### NewAddRequestCriteria200Response

`func NewAddRequestCriteria200Response(id string, schemas []EnumthirdPartyRequestCriteriaSchemaUrn, extensionClass string, ) *AddRequestCriteria200Response`

NewAddRequestCriteria200Response instantiates a new AddRequestCriteria200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRequestCriteria200ResponseWithDefaults

`func NewAddRequestCriteria200ResponseWithDefaults() *AddRequestCriteria200Response`

NewAddRequestCriteria200ResponseWithDefaults instantiates a new AddRequestCriteria200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *AddRequestCriteria200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AddRequestCriteria200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AddRequestCriteria200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AddRequestCriteria200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *AddRequestCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *AddRequestCriteria200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *AddRequestCriteria200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *AddRequestCriteria200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetId

`func (o *AddRequestCriteria200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddRequestCriteria200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddRequestCriteria200Response) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *AddRequestCriteria200Response) GetSchemas() []EnumthirdPartyRequestCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddRequestCriteria200Response) GetSchemasOk() (*[]EnumthirdPartyRequestCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddRequestCriteria200Response) SetSchemas(v []EnumthirdPartyRequestCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetOperationType

`func (o *AddRequestCriteria200Response) GetOperationType() []EnumrequestCriteriaSimpleOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *AddRequestCriteria200Response) GetOperationTypeOk() (*[]EnumrequestCriteriaSimpleOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *AddRequestCriteria200Response) SetOperationType(v []EnumrequestCriteriaSimpleOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *AddRequestCriteria200Response) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetDescription

`func (o *AddRequestCriteria200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRequestCriteria200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRequestCriteria200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRequestCriteria200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOperationOrigin

`func (o *AddRequestCriteria200Response) GetOperationOrigin() []EnumrequestCriteriaOperationOriginProp`

GetOperationOrigin returns the OperationOrigin field if non-nil, zero value otherwise.

### GetOperationOriginOk

`func (o *AddRequestCriteria200Response) GetOperationOriginOk() (*[]EnumrequestCriteriaOperationOriginProp, bool)`

GetOperationOriginOk returns a tuple with the OperationOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationOrigin

`func (o *AddRequestCriteria200Response) SetOperationOrigin(v []EnumrequestCriteriaOperationOriginProp)`

SetOperationOrigin sets OperationOrigin field to given value.

### HasOperationOrigin

`func (o *AddRequestCriteria200Response) HasOperationOrigin() bool`

HasOperationOrigin returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddRequestCriteria200Response) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddRequestCriteria200Response) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddRequestCriteria200Response) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddRequestCriteria200Response) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetAllIncludedRequestControl

`func (o *AddRequestCriteria200Response) GetAllIncludedRequestControl() []string`

GetAllIncludedRequestControl returns the AllIncludedRequestControl field if non-nil, zero value otherwise.

### GetAllIncludedRequestControlOk

`func (o *AddRequestCriteria200Response) GetAllIncludedRequestControlOk() (*[]string, bool)`

GetAllIncludedRequestControlOk returns a tuple with the AllIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestControl

`func (o *AddRequestCriteria200Response) SetAllIncludedRequestControl(v []string)`

SetAllIncludedRequestControl sets AllIncludedRequestControl field to given value.

### HasAllIncludedRequestControl

`func (o *AddRequestCriteria200Response) HasAllIncludedRequestControl() bool`

HasAllIncludedRequestControl returns a boolean if a field has been set.

### GetAnyIncludedRequestControl

`func (o *AddRequestCriteria200Response) GetAnyIncludedRequestControl() []string`

GetAnyIncludedRequestControl returns the AnyIncludedRequestControl field if non-nil, zero value otherwise.

### GetAnyIncludedRequestControlOk

`func (o *AddRequestCriteria200Response) GetAnyIncludedRequestControlOk() (*[]string, bool)`

GetAnyIncludedRequestControlOk returns a tuple with the AnyIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestControl

`func (o *AddRequestCriteria200Response) SetAnyIncludedRequestControl(v []string)`

SetAnyIncludedRequestControl sets AnyIncludedRequestControl field to given value.

### HasAnyIncludedRequestControl

`func (o *AddRequestCriteria200Response) HasAnyIncludedRequestControl() bool`

HasAnyIncludedRequestControl returns a boolean if a field has been set.

### GetNotAllIncludedRequestControl

`func (o *AddRequestCriteria200Response) GetNotAllIncludedRequestControl() []string`

GetNotAllIncludedRequestControl returns the NotAllIncludedRequestControl field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestControlOk

`func (o *AddRequestCriteria200Response) GetNotAllIncludedRequestControlOk() (*[]string, bool)`

GetNotAllIncludedRequestControlOk returns a tuple with the NotAllIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestControl

`func (o *AddRequestCriteria200Response) SetNotAllIncludedRequestControl(v []string)`

SetNotAllIncludedRequestControl sets NotAllIncludedRequestControl field to given value.

### HasNotAllIncludedRequestControl

`func (o *AddRequestCriteria200Response) HasNotAllIncludedRequestControl() bool`

HasNotAllIncludedRequestControl returns a boolean if a field has been set.

### GetNoneIncludedRequestControl

`func (o *AddRequestCriteria200Response) GetNoneIncludedRequestControl() []string`

GetNoneIncludedRequestControl returns the NoneIncludedRequestControl field if non-nil, zero value otherwise.

### GetNoneIncludedRequestControlOk

`func (o *AddRequestCriteria200Response) GetNoneIncludedRequestControlOk() (*[]string, bool)`

GetNoneIncludedRequestControlOk returns a tuple with the NoneIncludedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestControl

`func (o *AddRequestCriteria200Response) SetNoneIncludedRequestControl(v []string)`

SetNoneIncludedRequestControl sets NoneIncludedRequestControl field to given value.

### HasNoneIncludedRequestControl

`func (o *AddRequestCriteria200Response) HasNoneIncludedRequestControl() bool`

HasNoneIncludedRequestControl returns a boolean if a field has been set.

### GetIncludedTargetEntryDN

`func (o *AddRequestCriteria200Response) GetIncludedTargetEntryDN() []string`

GetIncludedTargetEntryDN returns the IncludedTargetEntryDN field if non-nil, zero value otherwise.

### GetIncludedTargetEntryDNOk

`func (o *AddRequestCriteria200Response) GetIncludedTargetEntryDNOk() (*[]string, bool)`

GetIncludedTargetEntryDNOk returns a tuple with the IncludedTargetEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetEntryDN

`func (o *AddRequestCriteria200Response) SetIncludedTargetEntryDN(v []string)`

SetIncludedTargetEntryDN sets IncludedTargetEntryDN field to given value.

### HasIncludedTargetEntryDN

`func (o *AddRequestCriteria200Response) HasIncludedTargetEntryDN() bool`

HasIncludedTargetEntryDN returns a boolean if a field has been set.

### GetExcludedTargetEntryDN

`func (o *AddRequestCriteria200Response) GetExcludedTargetEntryDN() []string`

GetExcludedTargetEntryDN returns the ExcludedTargetEntryDN field if non-nil, zero value otherwise.

### GetExcludedTargetEntryDNOk

`func (o *AddRequestCriteria200Response) GetExcludedTargetEntryDNOk() (*[]string, bool)`

GetExcludedTargetEntryDNOk returns a tuple with the ExcludedTargetEntryDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetEntryDN

`func (o *AddRequestCriteria200Response) SetExcludedTargetEntryDN(v []string)`

SetExcludedTargetEntryDN sets ExcludedTargetEntryDN field to given value.

### HasExcludedTargetEntryDN

`func (o *AddRequestCriteria200Response) HasExcludedTargetEntryDN() bool`

HasExcludedTargetEntryDN returns a boolean if a field has been set.

### GetAllIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) GetAllIncludedTargetEntryFilter() []string`

GetAllIncludedTargetEntryFilter returns the AllIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetAllIncludedTargetEntryFilterOk

`func (o *AddRequestCriteria200Response) GetAllIncludedTargetEntryFilterOk() (*[]string, bool)`

GetAllIncludedTargetEntryFilterOk returns a tuple with the AllIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) SetAllIncludedTargetEntryFilter(v []string)`

SetAllIncludedTargetEntryFilter sets AllIncludedTargetEntryFilter field to given value.

### HasAllIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) HasAllIncludedTargetEntryFilter() bool`

HasAllIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetAnyIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) GetAnyIncludedTargetEntryFilter() []string`

GetAnyIncludedTargetEntryFilter returns the AnyIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetAnyIncludedTargetEntryFilterOk

`func (o *AddRequestCriteria200Response) GetAnyIncludedTargetEntryFilterOk() (*[]string, bool)`

GetAnyIncludedTargetEntryFilterOk returns a tuple with the AnyIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) SetAnyIncludedTargetEntryFilter(v []string)`

SetAnyIncludedTargetEntryFilter sets AnyIncludedTargetEntryFilter field to given value.

### HasAnyIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) HasAnyIncludedTargetEntryFilter() bool`

HasAnyIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetNotAllIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) GetNotAllIncludedTargetEntryFilter() []string`

GetNotAllIncludedTargetEntryFilter returns the NotAllIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedTargetEntryFilterOk

`func (o *AddRequestCriteria200Response) GetNotAllIncludedTargetEntryFilterOk() (*[]string, bool)`

GetNotAllIncludedTargetEntryFilterOk returns a tuple with the NotAllIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) SetNotAllIncludedTargetEntryFilter(v []string)`

SetNotAllIncludedTargetEntryFilter sets NotAllIncludedTargetEntryFilter field to given value.

### HasNotAllIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) HasNotAllIncludedTargetEntryFilter() bool`

HasNotAllIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetNoneIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) GetNoneIncludedTargetEntryFilter() []string`

GetNoneIncludedTargetEntryFilter returns the NoneIncludedTargetEntryFilter field if non-nil, zero value otherwise.

### GetNoneIncludedTargetEntryFilterOk

`func (o *AddRequestCriteria200Response) GetNoneIncludedTargetEntryFilterOk() (*[]string, bool)`

GetNoneIncludedTargetEntryFilterOk returns a tuple with the NoneIncludedTargetEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) SetNoneIncludedTargetEntryFilter(v []string)`

SetNoneIncludedTargetEntryFilter sets NoneIncludedTargetEntryFilter field to given value.

### HasNoneIncludedTargetEntryFilter

`func (o *AddRequestCriteria200Response) HasNoneIncludedTargetEntryFilter() bool`

HasNoneIncludedTargetEntryFilter returns a boolean if a field has been set.

### GetAllIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) GetAllIncludedTargetEntryGroupDN() []string`

GetAllIncludedTargetEntryGroupDN returns the AllIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedTargetEntryGroupDNOk

`func (o *AddRequestCriteria200Response) GetAllIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetAllIncludedTargetEntryGroupDNOk returns a tuple with the AllIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) SetAllIncludedTargetEntryGroupDN(v []string)`

SetAllIncludedTargetEntryGroupDN sets AllIncludedTargetEntryGroupDN field to given value.

### HasAllIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) HasAllIncludedTargetEntryGroupDN() bool`

HasAllIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetAnyIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) GetAnyIncludedTargetEntryGroupDN() []string`

GetAnyIncludedTargetEntryGroupDN returns the AnyIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedTargetEntryGroupDNOk

`func (o *AddRequestCriteria200Response) GetAnyIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetAnyIncludedTargetEntryGroupDNOk returns a tuple with the AnyIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) SetAnyIncludedTargetEntryGroupDN(v []string)`

SetAnyIncludedTargetEntryGroupDN sets AnyIncludedTargetEntryGroupDN field to given value.

### HasAnyIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) HasAnyIncludedTargetEntryGroupDN() bool`

HasAnyIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) GetNotAllIncludedTargetEntryGroupDN() []string`

GetNotAllIncludedTargetEntryGroupDN returns the NotAllIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedTargetEntryGroupDNOk

`func (o *AddRequestCriteria200Response) GetNotAllIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetNotAllIncludedTargetEntryGroupDNOk returns a tuple with the NotAllIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) SetNotAllIncludedTargetEntryGroupDN(v []string)`

SetNotAllIncludedTargetEntryGroupDN sets NotAllIncludedTargetEntryGroupDN field to given value.

### HasNotAllIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) HasNotAllIncludedTargetEntryGroupDN() bool`

HasNotAllIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetNoneIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) GetNoneIncludedTargetEntryGroupDN() []string`

GetNoneIncludedTargetEntryGroupDN returns the NoneIncludedTargetEntryGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedTargetEntryGroupDNOk

`func (o *AddRequestCriteria200Response) GetNoneIncludedTargetEntryGroupDNOk() (*[]string, bool)`

GetNoneIncludedTargetEntryGroupDNOk returns a tuple with the NoneIncludedTargetEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) SetNoneIncludedTargetEntryGroupDN(v []string)`

SetNoneIncludedTargetEntryGroupDN sets NoneIncludedTargetEntryGroupDN field to given value.

### HasNoneIncludedTargetEntryGroupDN

`func (o *AddRequestCriteria200Response) HasNoneIncludedTargetEntryGroupDN() bool`

HasNoneIncludedTargetEntryGroupDN returns a boolean if a field has been set.

### GetTargetBindType

`func (o *AddRequestCriteria200Response) GetTargetBindType() []EnumrequestCriteriaTargetBindTypeProp`

GetTargetBindType returns the TargetBindType field if non-nil, zero value otherwise.

### GetTargetBindTypeOk

`func (o *AddRequestCriteria200Response) GetTargetBindTypeOk() (*[]EnumrequestCriteriaTargetBindTypeProp, bool)`

GetTargetBindTypeOk returns a tuple with the TargetBindType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBindType

`func (o *AddRequestCriteria200Response) SetTargetBindType(v []EnumrequestCriteriaTargetBindTypeProp)`

SetTargetBindType sets TargetBindType field to given value.

### HasTargetBindType

`func (o *AddRequestCriteria200Response) HasTargetBindType() bool`

HasTargetBindType returns a boolean if a field has been set.

### GetIncludedTargetSASLMechanism

`func (o *AddRequestCriteria200Response) GetIncludedTargetSASLMechanism() []string`

GetIncludedTargetSASLMechanism returns the IncludedTargetSASLMechanism field if non-nil, zero value otherwise.

### GetIncludedTargetSASLMechanismOk

`func (o *AddRequestCriteria200Response) GetIncludedTargetSASLMechanismOk() (*[]string, bool)`

GetIncludedTargetSASLMechanismOk returns a tuple with the IncludedTargetSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetSASLMechanism

`func (o *AddRequestCriteria200Response) SetIncludedTargetSASLMechanism(v []string)`

SetIncludedTargetSASLMechanism sets IncludedTargetSASLMechanism field to given value.

### HasIncludedTargetSASLMechanism

`func (o *AddRequestCriteria200Response) HasIncludedTargetSASLMechanism() bool`

HasIncludedTargetSASLMechanism returns a boolean if a field has been set.

### GetExcludedTargetSASLMechanism

`func (o *AddRequestCriteria200Response) GetExcludedTargetSASLMechanism() []string`

GetExcludedTargetSASLMechanism returns the ExcludedTargetSASLMechanism field if non-nil, zero value otherwise.

### GetExcludedTargetSASLMechanismOk

`func (o *AddRequestCriteria200Response) GetExcludedTargetSASLMechanismOk() (*[]string, bool)`

GetExcludedTargetSASLMechanismOk returns a tuple with the ExcludedTargetSASLMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetSASLMechanism

`func (o *AddRequestCriteria200Response) SetExcludedTargetSASLMechanism(v []string)`

SetExcludedTargetSASLMechanism sets ExcludedTargetSASLMechanism field to given value.

### HasExcludedTargetSASLMechanism

`func (o *AddRequestCriteria200Response) HasExcludedTargetSASLMechanism() bool`

HasExcludedTargetSASLMechanism returns a boolean if a field has been set.

### GetIncludedTargetAttribute

`func (o *AddRequestCriteria200Response) GetIncludedTargetAttribute() []string`

GetIncludedTargetAttribute returns the IncludedTargetAttribute field if non-nil, zero value otherwise.

### GetIncludedTargetAttributeOk

`func (o *AddRequestCriteria200Response) GetIncludedTargetAttributeOk() (*[]string, bool)`

GetIncludedTargetAttributeOk returns a tuple with the IncludedTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTargetAttribute

`func (o *AddRequestCriteria200Response) SetIncludedTargetAttribute(v []string)`

SetIncludedTargetAttribute sets IncludedTargetAttribute field to given value.

### HasIncludedTargetAttribute

`func (o *AddRequestCriteria200Response) HasIncludedTargetAttribute() bool`

HasIncludedTargetAttribute returns a boolean if a field has been set.

### GetExcludedTargetAttribute

`func (o *AddRequestCriteria200Response) GetExcludedTargetAttribute() []string`

GetExcludedTargetAttribute returns the ExcludedTargetAttribute field if non-nil, zero value otherwise.

### GetExcludedTargetAttributeOk

`func (o *AddRequestCriteria200Response) GetExcludedTargetAttributeOk() (*[]string, bool)`

GetExcludedTargetAttributeOk returns a tuple with the ExcludedTargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedTargetAttribute

`func (o *AddRequestCriteria200Response) SetExcludedTargetAttribute(v []string)`

SetExcludedTargetAttribute sets ExcludedTargetAttribute field to given value.

### HasExcludedTargetAttribute

`func (o *AddRequestCriteria200Response) HasExcludedTargetAttribute() bool`

HasExcludedTargetAttribute returns a boolean if a field has been set.

### GetIncludedExtendedOperationOID

`func (o *AddRequestCriteria200Response) GetIncludedExtendedOperationOID() []string`

GetIncludedExtendedOperationOID returns the IncludedExtendedOperationOID field if non-nil, zero value otherwise.

### GetIncludedExtendedOperationOIDOk

`func (o *AddRequestCriteria200Response) GetIncludedExtendedOperationOIDOk() (*[]string, bool)`

GetIncludedExtendedOperationOIDOk returns a tuple with the IncludedExtendedOperationOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedExtendedOperationOID

`func (o *AddRequestCriteria200Response) SetIncludedExtendedOperationOID(v []string)`

SetIncludedExtendedOperationOID sets IncludedExtendedOperationOID field to given value.

### HasIncludedExtendedOperationOID

`func (o *AddRequestCriteria200Response) HasIncludedExtendedOperationOID() bool`

HasIncludedExtendedOperationOID returns a boolean if a field has been set.

### GetExcludedExtendedOperationOID

`func (o *AddRequestCriteria200Response) GetExcludedExtendedOperationOID() []string`

GetExcludedExtendedOperationOID returns the ExcludedExtendedOperationOID field if non-nil, zero value otherwise.

### GetExcludedExtendedOperationOIDOk

`func (o *AddRequestCriteria200Response) GetExcludedExtendedOperationOIDOk() (*[]string, bool)`

GetExcludedExtendedOperationOIDOk returns a tuple with the ExcludedExtendedOperationOID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedExtendedOperationOID

`func (o *AddRequestCriteria200Response) SetExcludedExtendedOperationOID(v []string)`

SetExcludedExtendedOperationOID sets ExcludedExtendedOperationOID field to given value.

### HasExcludedExtendedOperationOID

`func (o *AddRequestCriteria200Response) HasExcludedExtendedOperationOID() bool`

HasExcludedExtendedOperationOID returns a boolean if a field has been set.

### GetIncludedSearchScope

`func (o *AddRequestCriteria200Response) GetIncludedSearchScope() []EnumrequestCriteriaIncludedSearchScopeProp`

GetIncludedSearchScope returns the IncludedSearchScope field if non-nil, zero value otherwise.

### GetIncludedSearchScopeOk

`func (o *AddRequestCriteria200Response) GetIncludedSearchScopeOk() (*[]EnumrequestCriteriaIncludedSearchScopeProp, bool)`

GetIncludedSearchScopeOk returns a tuple with the IncludedSearchScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedSearchScope

`func (o *AddRequestCriteria200Response) SetIncludedSearchScope(v []EnumrequestCriteriaIncludedSearchScopeProp)`

SetIncludedSearchScope sets IncludedSearchScope field to given value.

### HasIncludedSearchScope

`func (o *AddRequestCriteria200Response) HasIncludedSearchScope() bool`

HasIncludedSearchScope returns a boolean if a field has been set.

### GetUsingAdministrativeSessionWorkerThread

`func (o *AddRequestCriteria200Response) GetUsingAdministrativeSessionWorkerThread() EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp`

GetUsingAdministrativeSessionWorkerThread returns the UsingAdministrativeSessionWorkerThread field if non-nil, zero value otherwise.

### GetUsingAdministrativeSessionWorkerThreadOk

`func (o *AddRequestCriteria200Response) GetUsingAdministrativeSessionWorkerThreadOk() (*EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp, bool)`

GetUsingAdministrativeSessionWorkerThreadOk returns a tuple with the UsingAdministrativeSessionWorkerThread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingAdministrativeSessionWorkerThread

`func (o *AddRequestCriteria200Response) SetUsingAdministrativeSessionWorkerThread(v EnumrequestCriteriaUsingAdministrativeSessionWorkerThreadProp)`

SetUsingAdministrativeSessionWorkerThread sets UsingAdministrativeSessionWorkerThread field to given value.

### HasUsingAdministrativeSessionWorkerThread

`func (o *AddRequestCriteria200Response) HasUsingAdministrativeSessionWorkerThread() bool`

HasUsingAdministrativeSessionWorkerThread returns a boolean if a field has been set.

### GetIncludedApplicationName

`func (o *AddRequestCriteria200Response) GetIncludedApplicationName() []string`

GetIncludedApplicationName returns the IncludedApplicationName field if non-nil, zero value otherwise.

### GetIncludedApplicationNameOk

`func (o *AddRequestCriteria200Response) GetIncludedApplicationNameOk() (*[]string, bool)`

GetIncludedApplicationNameOk returns a tuple with the IncludedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedApplicationName

`func (o *AddRequestCriteria200Response) SetIncludedApplicationName(v []string)`

SetIncludedApplicationName sets IncludedApplicationName field to given value.

### HasIncludedApplicationName

`func (o *AddRequestCriteria200Response) HasIncludedApplicationName() bool`

HasIncludedApplicationName returns a boolean if a field has been set.

### GetExcludedApplicationName

`func (o *AddRequestCriteria200Response) GetExcludedApplicationName() []string`

GetExcludedApplicationName returns the ExcludedApplicationName field if non-nil, zero value otherwise.

### GetExcludedApplicationNameOk

`func (o *AddRequestCriteria200Response) GetExcludedApplicationNameOk() (*[]string, bool)`

GetExcludedApplicationNameOk returns a tuple with the ExcludedApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedApplicationName

`func (o *AddRequestCriteria200Response) SetExcludedApplicationName(v []string)`

SetExcludedApplicationName sets ExcludedApplicationName field to given value.

### HasExcludedApplicationName

`func (o *AddRequestCriteria200Response) HasExcludedApplicationName() bool`

HasExcludedApplicationName returns a boolean if a field has been set.

### GetAllIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) GetAllIncludedRequestCriteria() []string`

GetAllIncludedRequestCriteria returns the AllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAllIncludedRequestCriteriaOk

`func (o *AddRequestCriteria200Response) GetAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetAllIncludedRequestCriteriaOk returns a tuple with the AllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) SetAllIncludedRequestCriteria(v []string)`

SetAllIncludedRequestCriteria sets AllIncludedRequestCriteria field to given value.

### HasAllIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) HasAllIncludedRequestCriteria() bool`

HasAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetAnyIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) GetAnyIncludedRequestCriteria() []string`

GetAnyIncludedRequestCriteria returns the AnyIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetAnyIncludedRequestCriteriaOk

`func (o *AddRequestCriteria200Response) GetAnyIncludedRequestCriteriaOk() (*[]string, bool)`

GetAnyIncludedRequestCriteriaOk returns a tuple with the AnyIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) SetAnyIncludedRequestCriteria(v []string)`

SetAnyIncludedRequestCriteria sets AnyIncludedRequestCriteria field to given value.

### HasAnyIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) HasAnyIncludedRequestCriteria() bool`

HasAnyIncludedRequestCriteria returns a boolean if a field has been set.

### GetNotAllIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) GetNotAllIncludedRequestCriteria() []string`

GetNotAllIncludedRequestCriteria returns the NotAllIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNotAllIncludedRequestCriteriaOk

`func (o *AddRequestCriteria200Response) GetNotAllIncludedRequestCriteriaOk() (*[]string, bool)`

GetNotAllIncludedRequestCriteriaOk returns a tuple with the NotAllIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) SetNotAllIncludedRequestCriteria(v []string)`

SetNotAllIncludedRequestCriteria sets NotAllIncludedRequestCriteria field to given value.

### HasNotAllIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) HasNotAllIncludedRequestCriteria() bool`

HasNotAllIncludedRequestCriteria returns a boolean if a field has been set.

### GetNoneIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) GetNoneIncludedRequestCriteria() []string`

GetNoneIncludedRequestCriteria returns the NoneIncludedRequestCriteria field if non-nil, zero value otherwise.

### GetNoneIncludedRequestCriteriaOk

`func (o *AddRequestCriteria200Response) GetNoneIncludedRequestCriteriaOk() (*[]string, bool)`

GetNoneIncludedRequestCriteriaOk returns a tuple with the NoneIncludedRequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) SetNoneIncludedRequestCriteria(v []string)`

SetNoneIncludedRequestCriteria sets NoneIncludedRequestCriteria field to given value.

### HasNoneIncludedRequestCriteria

`func (o *AddRequestCriteria200Response) HasNoneIncludedRequestCriteria() bool`

HasNoneIncludedRequestCriteria returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddRequestCriteria200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddRequestCriteria200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddRequestCriteria200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddRequestCriteria200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddRequestCriteria200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddRequestCriteria200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddRequestCriteria200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


