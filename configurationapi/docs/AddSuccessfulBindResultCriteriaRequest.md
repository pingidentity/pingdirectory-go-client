# AddSuccessfulBindResultCriteriaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnumsuccessfulBindResultCriteriaSchemaUrn**](EnumsuccessfulBindResultCriteriaSchemaUrn.md) |  | 
**RequestCriteria** | Pointer to **string** | Specifies a request criteria object that must match the associated request for operations included in this Successful Bind Result Criteria. | [optional] 
**IncludeAnonymousBinds** | Pointer to **bool** | Indicates whether this criteria will be permitted to match bind operations that resulted in anonymous authentication. | [optional] 
**IncludedUserBaseDN** | Pointer to **[]string** | A set of base DNs for authenticated users that will be permitted to match this criteria. | [optional] 
**ExcludedUserBaseDN** | Pointer to **[]string** | A set of base DNs for authenticated users that will not be permitted to match this criteria. | [optional] 
**IncludedUserFilter** | Pointer to **[]string** | A set of filters that may be used to identify entries for authenticated users that will be permitted to match this criteria. | [optional] 
**ExcludedUserFilter** | Pointer to **[]string** | A set of filters that may be used to identify entries for authenticated users that will not be permitted to match this criteria. | [optional] 
**IncludedUserGroupDN** | Pointer to **[]string** | The DNs of the groups whose members will be permitted to match this criteria. | [optional] 
**ExcludedUserGroupDN** | Pointer to **[]string** | The DNs of the groups whose members will not be permitted to match this criteria. | [optional] 
**Description** | Pointer to **string** | A description for this Result Criteria | [optional] 
**CriteriaName** | **string** | Name of the new Result Criteria | 

## Methods

### NewAddSuccessfulBindResultCriteriaRequest

`func NewAddSuccessfulBindResultCriteriaRequest(schemas []EnumsuccessfulBindResultCriteriaSchemaUrn, criteriaName string, ) *AddSuccessfulBindResultCriteriaRequest`

NewAddSuccessfulBindResultCriteriaRequest instantiates a new AddSuccessfulBindResultCriteriaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSuccessfulBindResultCriteriaRequestWithDefaults

`func NewAddSuccessfulBindResultCriteriaRequestWithDefaults() *AddSuccessfulBindResultCriteriaRequest`

NewAddSuccessfulBindResultCriteriaRequestWithDefaults instantiates a new AddSuccessfulBindResultCriteriaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *AddSuccessfulBindResultCriteriaRequest) GetSchemas() []EnumsuccessfulBindResultCriteriaSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetSchemasOk() (*[]EnumsuccessfulBindResultCriteriaSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddSuccessfulBindResultCriteriaRequest) SetSchemas(v []EnumsuccessfulBindResultCriteriaSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetRequestCriteria

`func (o *AddSuccessfulBindResultCriteriaRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddSuccessfulBindResultCriteriaRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddSuccessfulBindResultCriteriaRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetIncludeAnonymousBinds

`func (o *AddSuccessfulBindResultCriteriaRequest) GetIncludeAnonymousBinds() bool`

GetIncludeAnonymousBinds returns the IncludeAnonymousBinds field if non-nil, zero value otherwise.

### GetIncludeAnonymousBindsOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetIncludeAnonymousBindsOk() (*bool, bool)`

GetIncludeAnonymousBindsOk returns a tuple with the IncludeAnonymousBinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnonymousBinds

`func (o *AddSuccessfulBindResultCriteriaRequest) SetIncludeAnonymousBinds(v bool)`

SetIncludeAnonymousBinds sets IncludeAnonymousBinds field to given value.

### HasIncludeAnonymousBinds

`func (o *AddSuccessfulBindResultCriteriaRequest) HasIncludeAnonymousBinds() bool`

HasIncludeAnonymousBinds returns a boolean if a field has been set.

### GetIncludedUserBaseDN

`func (o *AddSuccessfulBindResultCriteriaRequest) GetIncludedUserBaseDN() []string`

GetIncludedUserBaseDN returns the IncludedUserBaseDN field if non-nil, zero value otherwise.

### GetIncludedUserBaseDNOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetIncludedUserBaseDNOk() (*[]string, bool)`

GetIncludedUserBaseDNOk returns a tuple with the IncludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserBaseDN

`func (o *AddSuccessfulBindResultCriteriaRequest) SetIncludedUserBaseDN(v []string)`

SetIncludedUserBaseDN sets IncludedUserBaseDN field to given value.

### HasIncludedUserBaseDN

`func (o *AddSuccessfulBindResultCriteriaRequest) HasIncludedUserBaseDN() bool`

HasIncludedUserBaseDN returns a boolean if a field has been set.

### GetExcludedUserBaseDN

`func (o *AddSuccessfulBindResultCriteriaRequest) GetExcludedUserBaseDN() []string`

GetExcludedUserBaseDN returns the ExcludedUserBaseDN field if non-nil, zero value otherwise.

### GetExcludedUserBaseDNOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetExcludedUserBaseDNOk() (*[]string, bool)`

GetExcludedUserBaseDNOk returns a tuple with the ExcludedUserBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserBaseDN

`func (o *AddSuccessfulBindResultCriteriaRequest) SetExcludedUserBaseDN(v []string)`

SetExcludedUserBaseDN sets ExcludedUserBaseDN field to given value.

### HasExcludedUserBaseDN

`func (o *AddSuccessfulBindResultCriteriaRequest) HasExcludedUserBaseDN() bool`

HasExcludedUserBaseDN returns a boolean if a field has been set.

### GetIncludedUserFilter

`func (o *AddSuccessfulBindResultCriteriaRequest) GetIncludedUserFilter() []string`

GetIncludedUserFilter returns the IncludedUserFilter field if non-nil, zero value otherwise.

### GetIncludedUserFilterOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetIncludedUserFilterOk() (*[]string, bool)`

GetIncludedUserFilterOk returns a tuple with the IncludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserFilter

`func (o *AddSuccessfulBindResultCriteriaRequest) SetIncludedUserFilter(v []string)`

SetIncludedUserFilter sets IncludedUserFilter field to given value.

### HasIncludedUserFilter

`func (o *AddSuccessfulBindResultCriteriaRequest) HasIncludedUserFilter() bool`

HasIncludedUserFilter returns a boolean if a field has been set.

### GetExcludedUserFilter

`func (o *AddSuccessfulBindResultCriteriaRequest) GetExcludedUserFilter() []string`

GetExcludedUserFilter returns the ExcludedUserFilter field if non-nil, zero value otherwise.

### GetExcludedUserFilterOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetExcludedUserFilterOk() (*[]string, bool)`

GetExcludedUserFilterOk returns a tuple with the ExcludedUserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserFilter

`func (o *AddSuccessfulBindResultCriteriaRequest) SetExcludedUserFilter(v []string)`

SetExcludedUserFilter sets ExcludedUserFilter field to given value.

### HasExcludedUserFilter

`func (o *AddSuccessfulBindResultCriteriaRequest) HasExcludedUserFilter() bool`

HasExcludedUserFilter returns a boolean if a field has been set.

### GetIncludedUserGroupDN

`func (o *AddSuccessfulBindResultCriteriaRequest) GetIncludedUserGroupDN() []string`

GetIncludedUserGroupDN returns the IncludedUserGroupDN field if non-nil, zero value otherwise.

### GetIncludedUserGroupDNOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetIncludedUserGroupDNOk() (*[]string, bool)`

GetIncludedUserGroupDNOk returns a tuple with the IncludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUserGroupDN

`func (o *AddSuccessfulBindResultCriteriaRequest) SetIncludedUserGroupDN(v []string)`

SetIncludedUserGroupDN sets IncludedUserGroupDN field to given value.

### HasIncludedUserGroupDN

`func (o *AddSuccessfulBindResultCriteriaRequest) HasIncludedUserGroupDN() bool`

HasIncludedUserGroupDN returns a boolean if a field has been set.

### GetExcludedUserGroupDN

`func (o *AddSuccessfulBindResultCriteriaRequest) GetExcludedUserGroupDN() []string`

GetExcludedUserGroupDN returns the ExcludedUserGroupDN field if non-nil, zero value otherwise.

### GetExcludedUserGroupDNOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetExcludedUserGroupDNOk() (*[]string, bool)`

GetExcludedUserGroupDNOk returns a tuple with the ExcludedUserGroupDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserGroupDN

`func (o *AddSuccessfulBindResultCriteriaRequest) SetExcludedUserGroupDN(v []string)`

SetExcludedUserGroupDN sets ExcludedUserGroupDN field to given value.

### HasExcludedUserGroupDN

`func (o *AddSuccessfulBindResultCriteriaRequest) HasExcludedUserGroupDN() bool`

HasExcludedUserGroupDN returns a boolean if a field has been set.

### GetDescription

`func (o *AddSuccessfulBindResultCriteriaRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSuccessfulBindResultCriteriaRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSuccessfulBindResultCriteriaRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCriteriaName

`func (o *AddSuccessfulBindResultCriteriaRequest) GetCriteriaName() string`

GetCriteriaName returns the CriteriaName field if non-nil, zero value otherwise.

### GetCriteriaNameOk

`func (o *AddSuccessfulBindResultCriteriaRequest) GetCriteriaNameOk() (*string, bool)`

GetCriteriaNameOk returns a tuple with the CriteriaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteriaName

`func (o *AddSuccessfulBindResultCriteriaRequest) SetCriteriaName(v string)`

SetCriteriaName sets CriteriaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


