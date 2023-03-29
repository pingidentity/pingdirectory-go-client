# HighThroughputWorkQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | Pointer to [**[]EnumhighThroughputWorkQueueSchemaUrn**](EnumhighThroughputWorkQueueSchemaUrn.md) |  | [optional] 
**NumWorkerThreads** | Pointer to **int32** | Specifies the total number of worker threads that should be used within the server in order to process requested operations. The worker threads will be split evenly across all of the configured queues. | [optional] 
**NumWriteWorkerThreads** | Pointer to **int32** | Specifies the number of worker threads that should be used within the server to process write (add, delete, modify, and modify DN) operations. If this is specified, then separate sets of worker threads will be used for processing read and write operations, and the value of the num-worker-threads property will reflect the number of threads to use to process read operations. | [optional] 
**NumAdministrativeSessionWorkerThreads** | Pointer to **int32** | Specifies the number of worker threads that should be used to process operations as part of an administrative session. These threads may be reserved only for special use by management applications like dsconfig, the administration console, and other administrative tools, so that these applications may be used to diagnose problems and take any necessary corrective action even if all \&quot;normal\&quot; worker threads are busy processing other requests. | [optional] 
**NumQueues** | Pointer to **int32** | Specifies the number of blocking queues that should be maintained. A value of zero indicates that the server should attempt to automatically select an optimal value (one queue for every two worker threads). | [optional] 
**NumWriteQueues** | Pointer to **int32** | Specifies the number of blocking queues that should be maintained for write operations. This will only be used if a value is specified for the num-write-worker-threads property, in which case the num-queues property will specify the number of queues for read operations. Otherwise, all operations will be processed by a common set of worker threads and the value of the num-queues property will specify the number of queues for all types of operations. | [optional] 
**MaxWorkQueueCapacity** | Pointer to **int32** | Specifies the maximum number of pending operations that may be held in any of the queues at any given time. The total number of pending requests may be as large as this value times the total number of queues. | [optional] 
**MaxOfferTime** | Pointer to **string** | Specifies the maximum length of time that the connection handler should be allowed to wait to enqueue a request if the work queue is full. If the attempt to enqueue an operation does not succeed within this period of time, then the operation will be rejected and an error response will be returned to the client. A value of zero indicates that operations should be rejected immediately if the work queue is already at its maximum capacity. | [optional] 
**MonitorQueueTime** | Pointer to **bool** | Indicates whether the work queue should monitor the length of time that operations are held in the queue. When enabled the queue time will be included with access log messages as \&quot;qtime\&quot; in milliseconds. | [optional] 
**MaxQueueTime** | Pointer to **string** | Specifies the maximum length of time that an operation should be allowed to wait on the work queue. If an operation has been waiting on the queue longer than this period of time, then it will receive an immediate failure result rather than being processed once it has been handed off to a worker thread. A value of zero seconds indicates that there should not be any maximum queue time imposed. This setting will only be used if the monitor-queue-time property has a value of true. | [optional] 
**ExpensiveOperationCheckInterval** | Pointer to **string** | The interval that the work queue should use when checking for potentially expensive operations. If at least expensive-operation-minimum-concurrent-count worker threads are found to be processing the same operation on two consecutive polls separated by this time interval (i.e., the worker thread has been processing that operation for at least this length of time, and potentially up to twice this length of time), then a stack trace of all running threads will be written to a file for analysis to provide potentially useful information that may help better understand the reason it is taking so long. It may be that the operation is simply an expensive one to process, but there may be other external factors (e.g., a database checkpoint, a log rotation, lock contention, etc.) that could be to blame. This option is primarily intended for debugging purposes and should generally be used under the direction of Ping Identity support. | [optional] 
**ExpensiveOperationMinimumConcurrentCount** | Pointer to **int32** | The minimum number of concurrent expensive operations that should be detected to trigger dumping stack traces for all threads. If at least this number of worker threads are seen processing the same operations in two consecutive intervals, then the server will dump a stack trace of all threads to a file. This option is primarily intended for debugging purposes and should generally be used under the direction of Ping Identity support. | [optional] 
**ExpensiveOperationMinimumDumpInterval** | Pointer to **string** | The minimum length of time that should be required to pass after dumping stack trace information for all threads before the server should be allowed to create a second dump. This will help prevent the server from dumping stack traces too frequently and eventually consuming all available disk space with stack trace log output. This option is primarily intended for debugging purposes and should generally be used under the direction of Ping Identity support. | [optional] 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 

## Methods

### NewHighThroughputWorkQueueResponse

`func NewHighThroughputWorkQueueResponse() *HighThroughputWorkQueueResponse`

NewHighThroughputWorkQueueResponse instantiates a new HighThroughputWorkQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHighThroughputWorkQueueResponseWithDefaults

`func NewHighThroughputWorkQueueResponseWithDefaults() *HighThroughputWorkQueueResponse`

NewHighThroughputWorkQueueResponseWithDefaults instantiates a new HighThroughputWorkQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *HighThroughputWorkQueueResponse) GetSchemas() []EnumhighThroughputWorkQueueSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *HighThroughputWorkQueueResponse) GetSchemasOk() (*[]EnumhighThroughputWorkQueueSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *HighThroughputWorkQueueResponse) SetSchemas(v []EnumhighThroughputWorkQueueSchemaUrn)`

SetSchemas sets Schemas field to given value.

### HasSchemas

`func (o *HighThroughputWorkQueueResponse) HasSchemas() bool`

HasSchemas returns a boolean if a field has been set.

### GetNumWorkerThreads

`func (o *HighThroughputWorkQueueResponse) GetNumWorkerThreads() int32`

GetNumWorkerThreads returns the NumWorkerThreads field if non-nil, zero value otherwise.

### GetNumWorkerThreadsOk

`func (o *HighThroughputWorkQueueResponse) GetNumWorkerThreadsOk() (*int32, bool)`

GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkerThreads

`func (o *HighThroughputWorkQueueResponse) SetNumWorkerThreads(v int32)`

SetNumWorkerThreads sets NumWorkerThreads field to given value.

### HasNumWorkerThreads

`func (o *HighThroughputWorkQueueResponse) HasNumWorkerThreads() bool`

HasNumWorkerThreads returns a boolean if a field has been set.

### GetNumWriteWorkerThreads

`func (o *HighThroughputWorkQueueResponse) GetNumWriteWorkerThreads() int32`

GetNumWriteWorkerThreads returns the NumWriteWorkerThreads field if non-nil, zero value otherwise.

### GetNumWriteWorkerThreadsOk

`func (o *HighThroughputWorkQueueResponse) GetNumWriteWorkerThreadsOk() (*int32, bool)`

GetNumWriteWorkerThreadsOk returns a tuple with the NumWriteWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWriteWorkerThreads

`func (o *HighThroughputWorkQueueResponse) SetNumWriteWorkerThreads(v int32)`

SetNumWriteWorkerThreads sets NumWriteWorkerThreads field to given value.

### HasNumWriteWorkerThreads

`func (o *HighThroughputWorkQueueResponse) HasNumWriteWorkerThreads() bool`

HasNumWriteWorkerThreads returns a boolean if a field has been set.

### GetNumAdministrativeSessionWorkerThreads

`func (o *HighThroughputWorkQueueResponse) GetNumAdministrativeSessionWorkerThreads() int32`

GetNumAdministrativeSessionWorkerThreads returns the NumAdministrativeSessionWorkerThreads field if non-nil, zero value otherwise.

### GetNumAdministrativeSessionWorkerThreadsOk

`func (o *HighThroughputWorkQueueResponse) GetNumAdministrativeSessionWorkerThreadsOk() (*int32, bool)`

GetNumAdministrativeSessionWorkerThreadsOk returns a tuple with the NumAdministrativeSessionWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAdministrativeSessionWorkerThreads

`func (o *HighThroughputWorkQueueResponse) SetNumAdministrativeSessionWorkerThreads(v int32)`

SetNumAdministrativeSessionWorkerThreads sets NumAdministrativeSessionWorkerThreads field to given value.

### HasNumAdministrativeSessionWorkerThreads

`func (o *HighThroughputWorkQueueResponse) HasNumAdministrativeSessionWorkerThreads() bool`

HasNumAdministrativeSessionWorkerThreads returns a boolean if a field has been set.

### GetNumQueues

`func (o *HighThroughputWorkQueueResponse) GetNumQueues() int32`

GetNumQueues returns the NumQueues field if non-nil, zero value otherwise.

### GetNumQueuesOk

`func (o *HighThroughputWorkQueueResponse) GetNumQueuesOk() (*int32, bool)`

GetNumQueuesOk returns a tuple with the NumQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumQueues

`func (o *HighThroughputWorkQueueResponse) SetNumQueues(v int32)`

SetNumQueues sets NumQueues field to given value.

### HasNumQueues

`func (o *HighThroughputWorkQueueResponse) HasNumQueues() bool`

HasNumQueues returns a boolean if a field has been set.

### GetNumWriteQueues

`func (o *HighThroughputWorkQueueResponse) GetNumWriteQueues() int32`

GetNumWriteQueues returns the NumWriteQueues field if non-nil, zero value otherwise.

### GetNumWriteQueuesOk

`func (o *HighThroughputWorkQueueResponse) GetNumWriteQueuesOk() (*int32, bool)`

GetNumWriteQueuesOk returns a tuple with the NumWriteQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWriteQueues

`func (o *HighThroughputWorkQueueResponse) SetNumWriteQueues(v int32)`

SetNumWriteQueues sets NumWriteQueues field to given value.

### HasNumWriteQueues

`func (o *HighThroughputWorkQueueResponse) HasNumWriteQueues() bool`

HasNumWriteQueues returns a boolean if a field has been set.

### GetMaxWorkQueueCapacity

`func (o *HighThroughputWorkQueueResponse) GetMaxWorkQueueCapacity() int32`

GetMaxWorkQueueCapacity returns the MaxWorkQueueCapacity field if non-nil, zero value otherwise.

### GetMaxWorkQueueCapacityOk

`func (o *HighThroughputWorkQueueResponse) GetMaxWorkQueueCapacityOk() (*int32, bool)`

GetMaxWorkQueueCapacityOk returns a tuple with the MaxWorkQueueCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkQueueCapacity

`func (o *HighThroughputWorkQueueResponse) SetMaxWorkQueueCapacity(v int32)`

SetMaxWorkQueueCapacity sets MaxWorkQueueCapacity field to given value.

### HasMaxWorkQueueCapacity

`func (o *HighThroughputWorkQueueResponse) HasMaxWorkQueueCapacity() bool`

HasMaxWorkQueueCapacity returns a boolean if a field has been set.

### GetMaxOfferTime

`func (o *HighThroughputWorkQueueResponse) GetMaxOfferTime() string`

GetMaxOfferTime returns the MaxOfferTime field if non-nil, zero value otherwise.

### GetMaxOfferTimeOk

`func (o *HighThroughputWorkQueueResponse) GetMaxOfferTimeOk() (*string, bool)`

GetMaxOfferTimeOk returns a tuple with the MaxOfferTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOfferTime

`func (o *HighThroughputWorkQueueResponse) SetMaxOfferTime(v string)`

SetMaxOfferTime sets MaxOfferTime field to given value.

### HasMaxOfferTime

`func (o *HighThroughputWorkQueueResponse) HasMaxOfferTime() bool`

HasMaxOfferTime returns a boolean if a field has been set.

### GetMonitorQueueTime

`func (o *HighThroughputWorkQueueResponse) GetMonitorQueueTime() bool`

GetMonitorQueueTime returns the MonitorQueueTime field if non-nil, zero value otherwise.

### GetMonitorQueueTimeOk

`func (o *HighThroughputWorkQueueResponse) GetMonitorQueueTimeOk() (*bool, bool)`

GetMonitorQueueTimeOk returns a tuple with the MonitorQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorQueueTime

`func (o *HighThroughputWorkQueueResponse) SetMonitorQueueTime(v bool)`

SetMonitorQueueTime sets MonitorQueueTime field to given value.

### HasMonitorQueueTime

`func (o *HighThroughputWorkQueueResponse) HasMonitorQueueTime() bool`

HasMonitorQueueTime returns a boolean if a field has been set.

### GetMaxQueueTime

`func (o *HighThroughputWorkQueueResponse) GetMaxQueueTime() string`

GetMaxQueueTime returns the MaxQueueTime field if non-nil, zero value otherwise.

### GetMaxQueueTimeOk

`func (o *HighThroughputWorkQueueResponse) GetMaxQueueTimeOk() (*string, bool)`

GetMaxQueueTimeOk returns a tuple with the MaxQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueTime

`func (o *HighThroughputWorkQueueResponse) SetMaxQueueTime(v string)`

SetMaxQueueTime sets MaxQueueTime field to given value.

### HasMaxQueueTime

`func (o *HighThroughputWorkQueueResponse) HasMaxQueueTime() bool`

HasMaxQueueTime returns a boolean if a field has been set.

### GetExpensiveOperationCheckInterval

`func (o *HighThroughputWorkQueueResponse) GetExpensiveOperationCheckInterval() string`

GetExpensiveOperationCheckInterval returns the ExpensiveOperationCheckInterval field if non-nil, zero value otherwise.

### GetExpensiveOperationCheckIntervalOk

`func (o *HighThroughputWorkQueueResponse) GetExpensiveOperationCheckIntervalOk() (*string, bool)`

GetExpensiveOperationCheckIntervalOk returns a tuple with the ExpensiveOperationCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpensiveOperationCheckInterval

`func (o *HighThroughputWorkQueueResponse) SetExpensiveOperationCheckInterval(v string)`

SetExpensiveOperationCheckInterval sets ExpensiveOperationCheckInterval field to given value.

### HasExpensiveOperationCheckInterval

`func (o *HighThroughputWorkQueueResponse) HasExpensiveOperationCheckInterval() bool`

HasExpensiveOperationCheckInterval returns a boolean if a field has been set.

### GetExpensiveOperationMinimumConcurrentCount

`func (o *HighThroughputWorkQueueResponse) GetExpensiveOperationMinimumConcurrentCount() int32`

GetExpensiveOperationMinimumConcurrentCount returns the ExpensiveOperationMinimumConcurrentCount field if non-nil, zero value otherwise.

### GetExpensiveOperationMinimumConcurrentCountOk

`func (o *HighThroughputWorkQueueResponse) GetExpensiveOperationMinimumConcurrentCountOk() (*int32, bool)`

GetExpensiveOperationMinimumConcurrentCountOk returns a tuple with the ExpensiveOperationMinimumConcurrentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpensiveOperationMinimumConcurrentCount

`func (o *HighThroughputWorkQueueResponse) SetExpensiveOperationMinimumConcurrentCount(v int32)`

SetExpensiveOperationMinimumConcurrentCount sets ExpensiveOperationMinimumConcurrentCount field to given value.

### HasExpensiveOperationMinimumConcurrentCount

`func (o *HighThroughputWorkQueueResponse) HasExpensiveOperationMinimumConcurrentCount() bool`

HasExpensiveOperationMinimumConcurrentCount returns a boolean if a field has been set.

### GetExpensiveOperationMinimumDumpInterval

`func (o *HighThroughputWorkQueueResponse) GetExpensiveOperationMinimumDumpInterval() string`

GetExpensiveOperationMinimumDumpInterval returns the ExpensiveOperationMinimumDumpInterval field if non-nil, zero value otherwise.

### GetExpensiveOperationMinimumDumpIntervalOk

`func (o *HighThroughputWorkQueueResponse) GetExpensiveOperationMinimumDumpIntervalOk() (*string, bool)`

GetExpensiveOperationMinimumDumpIntervalOk returns a tuple with the ExpensiveOperationMinimumDumpInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpensiveOperationMinimumDumpInterval

`func (o *HighThroughputWorkQueueResponse) SetExpensiveOperationMinimumDumpInterval(v string)`

SetExpensiveOperationMinimumDumpInterval sets ExpensiveOperationMinimumDumpInterval field to given value.

### HasExpensiveOperationMinimumDumpInterval

`func (o *HighThroughputWorkQueueResponse) HasExpensiveOperationMinimumDumpInterval() bool`

HasExpensiveOperationMinimumDumpInterval returns a boolean if a field has been set.

### GetMeta

`func (o *HighThroughputWorkQueueResponse) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *HighThroughputWorkQueueResponse) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *HighThroughputWorkQueueResponse) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *HighThroughputWorkQueueResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *HighThroughputWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *HighThroughputWorkQueueResponse) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *HighThroughputWorkQueueResponse) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *HighThroughputWorkQueueResponse) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


