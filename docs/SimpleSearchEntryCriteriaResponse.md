# SimpleSearchEntryCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Search Entry Criteria | 
**Schemas** | [**[]EnumsimpleSearchEntryCriteriaSchemaUrn**](EnumsimpleSearchEntryCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for entries included in this Simple Search Entry Criteria. of them. | [optional] 
**AllIncludedEntryControl** | Pointer to **[]string** |  | [optional] 
**AnyIncludedEntryControl** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedEntryControl** | Pointer to **[]string** |  | [optional] 
**NoneIncludedEntryControl** | Pointer to **[]string** |  | [optional] 
**IncludedEntryBaseDN** | Pointer to **[]string** |  | [optional] 
**ExcludedEntryBaseDN** | Pointer to **[]string** |  | [optional] 
**AllIncludedEntryFilter** | Pointer to **[]string** |  | [optional] 
**AnyIncludedEntryFilter** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedEntryFilter** | Pointer to **[]string** |  | [optional] 
**NoneIncludedEntryFilter** | Pointer to **[]string** |  | [optional] 
**AllIncludedEntryGroupDN** | Pointer to **[]string** |  | [optional] 
**AnyIncludedEntryGroupDN** | Pointer to **[]string** |  | [optional] 
**NotAllIncludedEntryGroupDN** | Pointer to **[]string** |  | [optional] 
**NoneIncludedEntryGroupDN** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** | A description for this Search Entry Criteria | [optional] 

## Methods

### NewSimpleSearchEntryCriteriaResponse

`func NewSimpleSearchEntryCriteriaResponse(id string, schemas []EnumsimpleSearchEntryCriteriaSchemaUrn, ) *SimpleSearchEntryCriteriaResponse`

NewSimpleSearchEntryCriteriaResponse instantiates a new SimpleSearchEntryCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleSearchEntryCriteriaResponseWithDefaults

`func NewSimpleSearchEntryCriteriaResponseWithDefaults() *SimpleSearchEntryCriteriaResponse`

NewSimpleSearchEntryCriteriaResponseWithDefaults instantiates a new SimpleSearchEntryCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleSearchEntryCriteriaResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleSearchEntryCriteriaResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleSearchEntryCriteriaResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *SimpleSearchEntryCriteriaResponse) GetSchemas() []EnumsimpleSearchEntryCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *SimpleSearchEntryCriteriaResponse) GetSchemasOk() (*[]EnumsimpleSearchEntryCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *SimpleSearchEntryCriteriaResponse) SetSchemas(v []EnumsimpleSearchEntryCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *SimpleSearchEntryCriteriaResponse) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *SimpleSearchEntryCriteriaResponse) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *SimpleSearchEntryCriteriaResponse) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *SimpleSearchEntryCriteriaResponse) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetAllIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) GetAllIncludedEntryControl() []string`

GetAllIncludedEntryControl returns the AllIncludedEntryControl field if non-nil, zero value otherwise.

### GetAllIncludedEntryControlOk

`func (o *SimpleSearchEntryCriteriaResponse) GetAllIncludedEntryControlOk() (*[]string, bool)`

GetAllIncludedEntryControlOk returns a tuple with the AllIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) SetAllIncludedEntryControl(v []string)`

SetAllIncludedEntryControl sets AllIncludedEntryControl field to given value.

### HasAllIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) HasAllIncludedEntryControl() bool`

HasAllIncludedEntryControl returns a boolean if a field has been set.

### GetAnyIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) GetAnyIncludedEntryControl() []string`

GetAnyIncludedEntryControl returns the AnyIncludedEntryControl field if non-nil, zero value otherwise.

### GetAnyIncludedEntryControlOk

`func (o *SimpleSearchEntryCriteriaResponse) GetAnyIncludedEntryControlOk() (*[]string, bool)`

GetAnyIncludedEntryControlOk returns a tuple with the AnyIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) SetAnyIncludedEntryControl(v []string)`

SetAnyIncludedEntryControl sets AnyIncludedEntryControl field to given value.

### HasAnyIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) HasAnyIncludedEntryControl() bool`

HasAnyIncludedEntryControl returns a boolean if a field has been set.

### GetNotAllIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) GetNotAllIncludedEntryControl() []string`

GetNotAllIncludedEntryControl returns the NotAllIncludedEntryControl field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryControlOk

`func (o *SimpleSearchEntryCriteriaResponse) GetNotAllIncludedEntryControlOk() (*[]string, bool)`

GetNotAllIncludedEntryControlOk returns a tuple with the NotAllIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) SetNotAllIncludedEntryControl(v []string)`

SetNotAllIncludedEntryControl sets NotAllIncludedEntryControl field to given value.

### HasNotAllIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) HasNotAllIncludedEntryControl() bool`

HasNotAllIncludedEntryControl returns a boolean if a field has been set.

### GetNoneIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) GetNoneIncludedEntryControl() []string`

GetNoneIncludedEntryControl returns the NoneIncludedEntryControl field if non-nil, zero value otherwise.

### GetNoneIncludedEntryControlOk

`func (o *SimpleSearchEntryCriteriaResponse) GetNoneIncludedEntryControlOk() (*[]string, bool)`

GetNoneIncludedEntryControlOk returns a tuple with the NoneIncludedEntryControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) SetNoneIncludedEntryControl(v []string)`

SetNoneIncludedEntryControl sets NoneIncludedEntryControl field to given value.

### HasNoneIncludedEntryControl

`func (o *SimpleSearchEntryCriteriaResponse) HasNoneIncludedEntryControl() bool`

HasNoneIncludedEntryControl returns a boolean if a field has been set.

### GetIncludedEntryBaseDN

`func (o *SimpleSearchEntryCriteriaResponse) GetIncludedEntryBaseDN() []string`

GetIncludedEntryBaseDN returns the IncludedEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedEntryBaseDNOk

`func (o *SimpleSearchEntryCriteriaResponse) GetIncludedEntryBaseDNOk() (*[]string, bool)`

GetIncludedEntryBaseDNOk returns a tuple with the IncludedEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedEntryBaseDN

`func (o *SimpleSearchEntryCriteriaResponse) SetIncludedEntryBaseDN(v []string)`

SetIncludedEntryBaseDN sets IncludedEntryBaseDN field to given value.

### HasIncludedEntryBaseDN

`func (o *SimpleSearchEntryCriteriaResponse) HasIncludedEntryBaseDN() bool`

HasIncludedEntryBaseDN returns a boolean if a field has been set.

### GetExcludedEntryBaseDN

`func (o *SimpleSearchEntryCriteriaResponse) GetExcludedEntryBaseDN() []string`

GetExcludedEntryBaseDN returns the ExcludedEntryBaseDN field if non-nil, zero value otherwise.

### GetExcludedEntryBaseDNOk

`func (o *SimpleSearchEntryCriteriaResponse) GetExcludedEntryBaseDNOk() (*[]string, bool)`

GetExcludedEntryBaseDNOk returns a tuple with the ExcludedEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedEntryBaseDN

`func (o *SimpleSearchEntryCriteriaResponse) SetExcludedEntryBaseDN(v []string)`

SetExcludedEntryBaseDN sets ExcludedEntryBaseDN field to given value.

### HasExcludedEntryBaseDN

`func (o *SimpleSearchEntryCriteriaResponse) HasExcludedEntryBaseDN() bool`

HasExcludedEntryBaseDN returns a boolean if a field has been set.

### GetAllIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) GetAllIncludedEntryFilter() []string`

GetAllIncludedEntryFilter returns the AllIncludedEntryFilter field if non-nil, zero value otherwise.

### GetAllIncludedEntryFilterOk

`func (o *SimpleSearchEntryCriteriaResponse) GetAllIncludedEntryFilterOk() (*[]string, bool)`

GetAllIncludedEntryFilterOk returns a tuple with the AllIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) SetAllIncludedEntryFilter(v []string)`

SetAllIncludedEntryFilter sets AllIncludedEntryFilter field to given value.

### HasAllIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) HasAllIncludedEntryFilter() bool`

HasAllIncludedEntryFilter returns a boolean if a field has been set.

### GetAnyIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) GetAnyIncludedEntryFilter() []string`

GetAnyIncludedEntryFilter returns the AnyIncludedEntryFilter field if non-nil, zero value otherwise.

### GetAnyIncludedEntryFilterOk

`func (o *SimpleSearchEntryCriteriaResponse) GetAnyIncludedEntryFilterOk() (*[]string, bool)`

GetAnyIncludedEntryFilterOk returns a tuple with the AnyIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) SetAnyIncludedEntryFilter(v []string)`

SetAnyIncludedEntryFilter sets AnyIncludedEntryFilter field to given value.

### HasAnyIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) HasAnyIncludedEntryFilter() bool`

HasAnyIncludedEntryFilter returns a boolean if a field has been set.

### GetNotAllIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) GetNotAllIncludedEntryFilter() []string`

GetNotAllIncludedEntryFilter returns the NotAllIncludedEntryFilter field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryFilterOk

`func (o *SimpleSearchEntryCriteriaResponse) GetNotAllIncludedEntryFilterOk() (*[]string, bool)`

GetNotAllIncludedEntryFilterOk returns a tuple with the NotAllIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) SetNotAllIncludedEntryFilter(v []string)`

SetNotAllIncludedEntryFilter sets NotAllIncludedEntryFilter field to given value.

### HasNotAllIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) HasNotAllIncludedEntryFilter() bool`

HasNotAllIncludedEntryFilter returns a boolean if a field has been set.

### GetNoneIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) GetNoneIncludedEntryFilter() []string`

GetNoneIncludedEntryFilter returns the NoneIncludedEntryFilter field if non-nil, zero value otherwise.

### GetNoneIncludedEntryFilterOk

`func (o *SimpleSearchEntryCriteriaResponse) GetNoneIncludedEntryFilterOk() (*[]string, bool)`

GetNoneIncludedEntryFilterOk returns a tuple with the NoneIncludedEntryFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) SetNoneIncludedEntryFilter(v []string)`

SetNoneIncludedEntryFilter sets NoneIncludedEntryFilter field to given value.

### HasNoneIncludedEntryFilter

`func (o *SimpleSearchEntryCriteriaResponse) HasNoneIncludedEntryFilter() bool`

HasNoneIncludedEntryFilter returns a boolean if a field has been set.

### GetAllIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) GetAllIncludedEntryGroupDN() []string`

GetAllIncludedEntryGroupDN returns the AllIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetAllIncludedEntryGroupDNOk

`func (o *SimpleSearchEntryCriteriaResponse) GetAllIncludedEntryGroupDNOk() (*[]string, bool)`

GetAllIncludedEntryGroupDNOk returns a tuple with the AllIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) SetAllIncludedEntryGroupDN(v []string)`

SetAllIncludedEntryGroupDN sets AllIncludedEntryGroupDN field to given value.

### HasAllIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) HasAllIncludedEntryGroupDN() bool`

HasAllIncludedEntryGroupDN returns a boolean if a field has been set.

### GetAnyIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) GetAnyIncludedEntryGroupDN() []string`

GetAnyIncludedEntryGroupDN returns the AnyIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetAnyIncludedEntryGroupDNOk

`func (o *SimpleSearchEntryCriteriaResponse) GetAnyIncludedEntryGroupDNOk() (*[]string, bool)`

GetAnyIncludedEntryGroupDNOk returns a tuple with the AnyIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnyIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) SetAnyIncludedEntryGroupDN(v []string)`

SetAnyIncludedEntryGroupDN sets AnyIncludedEntryGroupDN field to given value.

### HasAnyIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) HasAnyIncludedEntryGroupDN() bool`

HasAnyIncludedEntryGroupDN returns a boolean if a field has been set.

### GetNotAllIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) GetNotAllIncludedEntryGroupDN() []string`

GetNotAllIncludedEntryGroupDN returns the NotAllIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetNotAllIncludedEntryGroupDNOk

`func (o *SimpleSearchEntryCriteriaResponse) GetNotAllIncludedEntryGroupDNOk() (*[]string, bool)`

GetNotAllIncludedEntryGroupDNOk returns a tuple with the NotAllIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) SetNotAllIncludedEntryGroupDN(v []string)`

SetNotAllIncludedEntryGroupDN sets NotAllIncludedEntryGroupDN field to given value.

### HasNotAllIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) HasNotAllIncludedEntryGroupDN() bool`

HasNotAllIncludedEntryGroupDN returns a boolean if a field has been set.

### GetNoneIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) GetNoneIncludedEntryGroupDN() []string`

GetNoneIncludedEntryGroupDN returns the NoneIncludedEntryGroupDN field if non-nil, zero value otherwise.

### GetNoneIncludedEntryGroupDNOk

`func (o *SimpleSearchEntryCriteriaResponse) GetNoneIncludedEntryGroupDNOk() (*[]string, bool)`

GetNoneIncludedEntryGroupDNOk returns a tuple with the NoneIncludedEntryGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoneIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) SetNoneIncludedEntryGroupDN(v []string)`

SetNoneIncludedEntryGroupDN sets NoneIncludedEntryGroupDN field to given value.

### HasNoneIncludedEntryGroupDN

`func (o *SimpleSearchEntryCriteriaResponse) HasNoneIncludedEntryGroupDN() bool`

HasNoneIncludedEntryGroupDN returns a boolean if a field has been set.

### GetDescription

`func (o *SimpleSearchEntryCriteriaResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimpleSearchEntryCriteriaResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimpleSearchEntryCriteriaResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimpleSearchEntryCriteriaResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


