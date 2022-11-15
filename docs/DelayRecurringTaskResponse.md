# DelayRecurringTaskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Name of the Recurring Task | 
**Schemas** | [**[]EnumdelayRecurringTaskSchemaUrn**](EnumdelayRecurringTaskSchemaUrn.md) |  | 
**SleepDuration** | Pointer to **string** | The length of time to sleep before the task completes. | [optional] 
**DurationToWaitForWorkQueueIdle** | Pointer to **string** | Indicates that task should wait for up to the specified length of time for the work queue to report that all worker threads are idle and there are no pending operations. Note that this primarily monitors operations that use worker threads, which does not include internal operations (for example, those invoked by extensions), and may not include requests from non-LDAP clients (for example, HTTP-based clients). | [optional] 
**LdapURLForSearchExpectedToReturnEntries** | Pointer to **[]string** |  | [optional] 
**SearchInterval** | Pointer to **string** | The length of time the server should sleep between searches performed using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**SearchTimeLimit** | Pointer to **string** | The length of time that the server will wait for a response to each internal search performed using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**DurationToWaitForSearchToReturnEntries** | Pointer to **string** | The maximum length of time that the server will continue to perform internal searches using the criteria from the ldap-url-for-search-expected-to-return-entries property. | [optional] 
**TaskReturnStateIfTimeoutIsEncountered** | Pointer to [**EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp**](EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp.md) |  | [optional] 
**Description** | Pointer to **string** | A description for this Recurring Task | [optional] 
**CancelOnTaskDependencyFailure** | Pointer to **bool** | Indicates whether an instance of this Recurring Task should be canceled if the task immediately before it in the recurring task chain fails to complete successfully (including if it is canceled by an administrator before it starts or while it is running). | [optional] 
**EmailOnStart** | Pointer to **[]string** |  | [optional] 
**EmailOnSuccess** | Pointer to **[]string** |  | [optional] 
**EmailOnFailure** | Pointer to **[]string** |  | [optional] 
**AlertOnStart** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task starts running. | [optional] 
**AlertOnSuccess** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task completes successfully. | [optional] 
**AlertOnFailure** | Pointer to **bool** | Indicates whether the server should generate an administrative alert whenever an instance of this Recurring Task fails to complete successfully. | [optional] 

## Methods

### NewDelayRecurringTaskResponse

`func NewDelayRecurringTaskResponse(id string, schemas []EnumdelayRecurringTaskSchemaUrn, ) *DelayRecurringTaskResponse`

NewDelayRecurringTaskResponse instantiates a new DelayRecurringTaskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelayRecurringTaskResponseWithDefaults

`func NewDelayRecurringTaskResponseWithDefaults() *DelayRecurringTaskResponse`

NewDelayRecurringTaskResponseWithDefaults instantiates a new DelayRecurringTaskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DelayRecurringTaskResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DelayRecurringTaskResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DelayRecurringTaskResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSchemas

`func (o *DelayRecurringTaskResponse) GetSchemas() []EnumdelayRecurringTaskSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *DelayRecurringTaskResponse) GetSchemasOk() (*[]EnumdelayRecurringTaskSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *DelayRecurringTaskResponse) SetSchemas(v []EnumdelayRecurringTaskSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetSleepDuration

`func (o *DelayRecurringTaskResponse) GetSleepDuration() string`

GetSleepDuration returns the SleepDuration field if non-nil, zero value otherwise.

### GetSleepDurationOk

`func (o *DelayRecurringTaskResponse) GetSleepDurationOk() (*string, bool)`

GetSleepDurationOk returns a tuple with the SleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepDuration

`func (o *DelayRecurringTaskResponse) SetSleepDuration(v string)`

SetSleepDuration sets SleepDuration field to given value.

### HasSleepDuration

`func (o *DelayRecurringTaskResponse) HasSleepDuration() bool`

HasSleepDuration returns a boolean if a field has been set.

### GetDurationToWaitForWorkQueueIdle

`func (o *DelayRecurringTaskResponse) GetDurationToWaitForWorkQueueIdle() string`

GetDurationToWaitForWorkQueueIdle returns the DurationToWaitForWorkQueueIdle field if non-nil, zero value otherwise.

### GetDurationToWaitForWorkQueueIdleOk

`func (o *DelayRecurringTaskResponse) GetDurationToWaitForWorkQueueIdleOk() (*string, bool)`

GetDurationToWaitForWorkQueueIdleOk returns a tuple with the DurationToWaitForWorkQueueIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToWaitForWorkQueueIdle

`func (o *DelayRecurringTaskResponse) SetDurationToWaitForWorkQueueIdle(v string)`

SetDurationToWaitForWorkQueueIdle sets DurationToWaitForWorkQueueIdle field to given value.

### HasDurationToWaitForWorkQueueIdle

`func (o *DelayRecurringTaskResponse) HasDurationToWaitForWorkQueueIdle() bool`

HasDurationToWaitForWorkQueueIdle returns a boolean if a field has been set.

### GetLdapURLForSearchExpectedToReturnEntries

`func (o *DelayRecurringTaskResponse) GetLdapURLForSearchExpectedToReturnEntries() []string`

GetLdapURLForSearchExpectedToReturnEntries returns the LdapURLForSearchExpectedToReturnEntries field if non-nil, zero value otherwise.

### GetLdapURLForSearchExpectedToReturnEntriesOk

`func (o *DelayRecurringTaskResponse) GetLdapURLForSearchExpectedToReturnEntriesOk() (*[]string, bool)`

GetLdapURLForSearchExpectedToReturnEntriesOk returns a tuple with the LdapURLForSearchExpectedToReturnEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapURLForSearchExpectedToReturnEntries

`func (o *DelayRecurringTaskResponse) SetLdapURLForSearchExpectedToReturnEntries(v []string)`

SetLdapURLForSearchExpectedToReturnEntries sets LdapURLForSearchExpectedToReturnEntries field to given value.

### HasLdapURLForSearchExpectedToReturnEntries

`func (o *DelayRecurringTaskResponse) HasLdapURLForSearchExpectedToReturnEntries() bool`

HasLdapURLForSearchExpectedToReturnEntries returns a boolean if a field has been set.

### GetSearchInterval

`func (o *DelayRecurringTaskResponse) GetSearchInterval() string`

GetSearchInterval returns the SearchInterval field if non-nil, zero value otherwise.

### GetSearchIntervalOk

`func (o *DelayRecurringTaskResponse) GetSearchIntervalOk() (*string, bool)`

GetSearchIntervalOk returns a tuple with the SearchInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchInterval

`func (o *DelayRecurringTaskResponse) SetSearchInterval(v string)`

SetSearchInterval sets SearchInterval field to given value.

### HasSearchInterval

`func (o *DelayRecurringTaskResponse) HasSearchInterval() bool`

HasSearchInterval returns a boolean if a field has been set.

### GetSearchTimeLimit

`func (o *DelayRecurringTaskResponse) GetSearchTimeLimit() string`

GetSearchTimeLimit returns the SearchTimeLimit field if non-nil, zero value otherwise.

### GetSearchTimeLimitOk

`func (o *DelayRecurringTaskResponse) GetSearchTimeLimitOk() (*string, bool)`

GetSearchTimeLimitOk returns a tuple with the SearchTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTimeLimit

`func (o *DelayRecurringTaskResponse) SetSearchTimeLimit(v string)`

SetSearchTimeLimit sets SearchTimeLimit field to given value.

### HasSearchTimeLimit

`func (o *DelayRecurringTaskResponse) HasSearchTimeLimit() bool`

HasSearchTimeLimit returns a boolean if a field has been set.

### GetDurationToWaitForSearchToReturnEntries

`func (o *DelayRecurringTaskResponse) GetDurationToWaitForSearchToReturnEntries() string`

GetDurationToWaitForSearchToReturnEntries returns the DurationToWaitForSearchToReturnEntries field if non-nil, zero value otherwise.

### GetDurationToWaitForSearchToReturnEntriesOk

`func (o *DelayRecurringTaskResponse) GetDurationToWaitForSearchToReturnEntriesOk() (*string, bool)`

GetDurationToWaitForSearchToReturnEntriesOk returns a tuple with the DurationToWaitForSearchToReturnEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationToWaitForSearchToReturnEntries

`func (o *DelayRecurringTaskResponse) SetDurationToWaitForSearchToReturnEntries(v string)`

SetDurationToWaitForSearchToReturnEntries sets DurationToWaitForSearchToReturnEntries field to given value.

### HasDurationToWaitForSearchToReturnEntries

`func (o *DelayRecurringTaskResponse) HasDurationToWaitForSearchToReturnEntries() bool`

HasDurationToWaitForSearchToReturnEntries returns a boolean if a field has been set.

### GetTaskReturnStateIfTimeoutIsEncountered

`func (o *DelayRecurringTaskResponse) GetTaskReturnStateIfTimeoutIsEncountered() EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp`

GetTaskReturnStateIfTimeoutIsEncountered returns the TaskReturnStateIfTimeoutIsEncountered field if non-nil, zero value otherwise.

### GetTaskReturnStateIfTimeoutIsEncounteredOk

`func (o *DelayRecurringTaskResponse) GetTaskReturnStateIfTimeoutIsEncounteredOk() (*EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp, bool)`

GetTaskReturnStateIfTimeoutIsEncounteredOk returns a tuple with the TaskReturnStateIfTimeoutIsEncountered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskReturnStateIfTimeoutIsEncountered

`func (o *DelayRecurringTaskResponse) SetTaskReturnStateIfTimeoutIsEncountered(v EnumrecurringTaskTaskReturnStateIfTimeoutIsEncounteredProp)`

SetTaskReturnStateIfTimeoutIsEncountered sets TaskReturnStateIfTimeoutIsEncountered field to given value.

### HasTaskReturnStateIfTimeoutIsEncountered

`func (o *DelayRecurringTaskResponse) HasTaskReturnStateIfTimeoutIsEncountered() bool`

HasTaskReturnStateIfTimeoutIsEncountered returns a boolean if a field has been set.

### GetDescription

`func (o *DelayRecurringTaskResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DelayRecurringTaskResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DelayRecurringTaskResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DelayRecurringTaskResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCancelOnTaskDependencyFailure

`func (o *DelayRecurringTaskResponse) GetCancelOnTaskDependencyFailure() bool`

GetCancelOnTaskDependencyFailure returns the CancelOnTaskDependencyFailure field if non-nil, zero value otherwise.

### GetCancelOnTaskDependencyFailureOk

`func (o *DelayRecurringTaskResponse) GetCancelOnTaskDependencyFailureOk() (*bool, bool)`

GetCancelOnTaskDependencyFailureOk returns a tuple with the CancelOnTaskDependencyFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOnTaskDependencyFailure

`func (o *DelayRecurringTaskResponse) SetCancelOnTaskDependencyFailure(v bool)`

SetCancelOnTaskDependencyFailure sets CancelOnTaskDependencyFailure field to given value.

### HasCancelOnTaskDependencyFailure

`func (o *DelayRecurringTaskResponse) HasCancelOnTaskDependencyFailure() bool`

HasCancelOnTaskDependencyFailure returns a boolean if a field has been set.

### GetEmailOnStart

`func (o *DelayRecurringTaskResponse) GetEmailOnStart() []string`

GetEmailOnStart returns the EmailOnStart field if non-nil, zero value otherwise.

### GetEmailOnStartOk

`func (o *DelayRecurringTaskResponse) GetEmailOnStartOk() (*[]string, bool)`

GetEmailOnStartOk returns a tuple with the EmailOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnStart

`func (o *DelayRecurringTaskResponse) SetEmailOnStart(v []string)`

SetEmailOnStart sets EmailOnStart field to given value.

### HasEmailOnStart

`func (o *DelayRecurringTaskResponse) HasEmailOnStart() bool`

HasEmailOnStart returns a boolean if a field has been set.

### GetEmailOnSuccess

`func (o *DelayRecurringTaskResponse) GetEmailOnSuccess() []string`

GetEmailOnSuccess returns the EmailOnSuccess field if non-nil, zero value otherwise.

### GetEmailOnSuccessOk

`func (o *DelayRecurringTaskResponse) GetEmailOnSuccessOk() (*[]string, bool)`

GetEmailOnSuccessOk returns a tuple with the EmailOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnSuccess

`func (o *DelayRecurringTaskResponse) SetEmailOnSuccess(v []string)`

SetEmailOnSuccess sets EmailOnSuccess field to given value.

### HasEmailOnSuccess

`func (o *DelayRecurringTaskResponse) HasEmailOnSuccess() bool`

HasEmailOnSuccess returns a boolean if a field has been set.

### GetEmailOnFailure

`func (o *DelayRecurringTaskResponse) GetEmailOnFailure() []string`

GetEmailOnFailure returns the EmailOnFailure field if non-nil, zero value otherwise.

### GetEmailOnFailureOk

`func (o *DelayRecurringTaskResponse) GetEmailOnFailureOk() (*[]string, bool)`

GetEmailOnFailureOk returns a tuple with the EmailOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailOnFailure

`func (o *DelayRecurringTaskResponse) SetEmailOnFailure(v []string)`

SetEmailOnFailure sets EmailOnFailure field to given value.

### HasEmailOnFailure

`func (o *DelayRecurringTaskResponse) HasEmailOnFailure() bool`

HasEmailOnFailure returns a boolean if a field has been set.

### GetAlertOnStart

`func (o *DelayRecurringTaskResponse) GetAlertOnStart() bool`

GetAlertOnStart returns the AlertOnStart field if non-nil, zero value otherwise.

### GetAlertOnStartOk

`func (o *DelayRecurringTaskResponse) GetAlertOnStartOk() (*bool, bool)`

GetAlertOnStartOk returns a tuple with the AlertOnStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnStart

`func (o *DelayRecurringTaskResponse) SetAlertOnStart(v bool)`

SetAlertOnStart sets AlertOnStart field to given value.

### HasAlertOnStart

`func (o *DelayRecurringTaskResponse) HasAlertOnStart() bool`

HasAlertOnStart returns a boolean if a field has been set.

### GetAlertOnSuccess

`func (o *DelayRecurringTaskResponse) GetAlertOnSuccess() bool`

GetAlertOnSuccess returns the AlertOnSuccess field if non-nil, zero value otherwise.

### GetAlertOnSuccessOk

`func (o *DelayRecurringTaskResponse) GetAlertOnSuccessOk() (*bool, bool)`

GetAlertOnSuccessOk returns a tuple with the AlertOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnSuccess

`func (o *DelayRecurringTaskResponse) SetAlertOnSuccess(v bool)`

SetAlertOnSuccess sets AlertOnSuccess field to given value.

### HasAlertOnSuccess

`func (o *DelayRecurringTaskResponse) HasAlertOnSuccess() bool`

HasAlertOnSuccess returns a boolean if a field has been set.

### GetAlertOnFailure

`func (o *DelayRecurringTaskResponse) GetAlertOnFailure() bool`

GetAlertOnFailure returns the AlertOnFailure field if non-nil, zero value otherwise.

### GetAlertOnFailureOk

`func (o *DelayRecurringTaskResponse) GetAlertOnFailureOk() (*bool, bool)`

GetAlertOnFailureOk returns a tuple with the AlertOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertOnFailure

`func (o *DelayRecurringTaskResponse) SetAlertOnFailure(v bool)`

SetAlertOnFailure sets AlertOnFailure field to given value.

### HasAlertOnFailure

`func (o *DelayRecurringTaskResponse) HasAlertOnFailure() bool`

HasAlertOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


