# AddPluginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | **string** | Name of the new Plugin | 
**Schemas** | [**[]EnumuniqueAttributePluginSchemaUrn**](EnumuniqueAttributePluginSchemaUrn.md) |  | 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**NumThreads** | Pointer to **int32** | Specifies the number of concurrent threads that should be used to process the search operations. | [optional] 
**BaseDN** | **[]string** | Specifies a base DN within which the attribute must be unique. | 
**LowerBound** | Pointer to **int32** | Specifies the lower bound for the numeric value which will be inserted into the search filter. | [optional] 
**UpperBound** | Pointer to **int32** | Specifies the upper bound for the numeric value which will be inserted into the search filter. | [optional] 
**FilterPrefix** | **string** | Specifies a prefix which will be used in front of the randomly-selected numeric value in all search filters used. If no upper bound is defined, then this should contain the entire filter string. | 
**FilterSuffix** | Pointer to **string** | Specifies a suffix which will be used after of the randomly-selected numeric value in all search filters used. If no upper bound is defined, then this should be omitted. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Filter** | **string** | Specifies the search filter to apply to determine if attribute uniqueness is enforced for the matching entries. | 
**AttributeType** | **[]string** | Specifies the attribute types for which referential integrity is to be maintained. | 
**PollingInterval** | Pointer to **string** | This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information). | [optional] 
**PeerServerPriorityIndex** | Pointer to **int32** | In a replicated environment, this determines the order in which peer servers should attempt to purge data. | [optional] 
**MaxUpdatesPerSecond** | Pointer to **int32** | This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling. | [optional] 
**NumDeleteThreads** | Pointer to **int32** | The number of threads used to delete expired entries. | [optional] 
**InvokeGCDayOfWeek** | Pointer to [**[]EnumpluginInvokeGCDayOfWeekProp**](EnumpluginInvokeGCDayOfWeekProp.md) |  | [optional] 
**InvokeGCTimeUtc** | **[]string** | Specifies the times of the day at which garbage collection may be explicitly invoked. The times should be specified in \&quot;HH:MM\&quot; format, with \&quot;HH\&quot; as a two-digit numeric value between 00 and 23 representing the hour of the day, and MM as a two-digit numeric value between 00 and 59 representing the minute of the hour. All times will be interpreted in the UTC time zone. | 
**DelayAfterAlert** | Pointer to **string** | Specifies the length of time that the Directory Server should wait after sending the \&quot;force-gc-starting\&quot; administrative alert before actually invoking the garbage collection processing. | [optional] 
**DelayPostGC** | Pointer to **string** | Specifies the length of time that the Directory Server should wait after successfully completing the garbage collection processing, before removing the \&quot;force-gc-starting\&quot; administrative alert, which marks the server as unavailable. | [optional] 
**ApiURL** | **string** | Specifies the API endpoint for the PingOne web service. | 
**AuthURL** | **string** | Specifies the API endpoint for the PingOne authentication service. | 
**OAuthClientID** | **string** | Specifies the OAuth Client ID used to authenticate connections to the PingOne API. | 
**OAuthClientSecret** | Pointer to **string** | Specifies the OAuth Client Secret used to authenticate connections to the PingOne API. | [optional] 
**OAuthClientSecretPassphraseProvider** | Pointer to **string** | Specifies a passphrase provider that can be used to obtain the OAuth Client Secret used to authenticate connections to the PingOne API. | [optional] 
**EnvironmentID** | **string** | Specifies the PingOne Environment that will be associated with this PingOne Pass Through Authentication Plugin. | 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to the external authentication service. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**RequestCriteria** | Pointer to **string** | A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**TryLocalBind** | Pointer to **bool** | Indicates whether to attempt the bind in the local server first and only send the request to the external authentication service if the local bind attempt fails, or to only attempt the bind in the external service. | [optional] 
**OverrideLocalPassword** | Pointer to **bool** | Indicates whether to attempt the authentication in the external service if the local user entry includes a password. This property will be ignored if try-local-bind is false. | [optional] 
**UpdateLocalPassword** | Pointer to **bool** | Indicates whether to overwrite the user&#39;s local password if the local bind fails but the authentication attempt succeeds when attempted in the external service. This property may only be set to true if try-local-bind is also true. | [optional] 
**UpdateLocalPasswordDN** | Pointer to **string** | The DN of the authorization identity that will be used when updating the user&#39;s local password if update-local-password is true. This is primarily intended for use if the Data Sync Server will be used to synchronize passwords between the local server and the external service, and in that case, the DN used here should also be added to the ignore-changes-by-dn property in the appropriate Sync Source object in the Data Sync Server configuration. | [optional] 
**AllowLaxPassThroughAuthenticationPasswords** | Pointer to **bool** | Indicates whether to overwrite the user&#39;s local password even if the password used to authenticate to the external service would have failed validation if the user attempted to set it directly. | [optional] 
**IgnoredPasswordPolicyStateErrorCondition** | Pointer to [**[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp**](EnumpluginIgnoredPasswordPolicyStateErrorConditionProp.md) |  | [optional] 
**UserMappingLocalAttribute** | **[]string** | The names of the attributes in the local user entry whose values must match the values of the corresponding fields in the PingOne service. | 
**UserMappingRemoteJSONField** | **[]string** | The names of the fields in the PingOne service whose values must match the values of the corresponding attributes in the local user entry, as specified in the user-mapping-local-attribute property. | 
**AdditionalUserMappingSCIMFilter** | Pointer to **string** | An optional SCIM filter that will be ANDed with the filter created to identify the account in the PingOne service that corresponds to the local entry. Only the \&quot;eq\&quot;, \&quot;sw\&quot;, \&quot;and\&quot;, and \&quot;or\&quot; filter types may be used. | [optional] 
**Scope** | [**EnumpluginScopeProp**](EnumpluginScopeProp.md) |  | 
**IncludeAttribute** | Pointer to **[]string** | The name of an attribute that should be included in the results. This may include any token which is allowed as a requested attribute in search requests, including the name of an attribute, an asterisk (to indicate all user attributes), a plus sign (to indicate all operational attributes), an object class name preceded with an at symbol (to indicate all attributes associated with that object class), an attribute name preceded by a caret (to indicate that attribute should be excluded), or an object class name preceded by a caret and an at symbol (to indicate that all attributes associated with that object class should be excluded). | [optional] 
**OutputFile** | **string** | The path of an LDIF file that should be created with the results of the search. | 
**PreviousFileExtension** | Pointer to **string** | An extension that should be appended to the name of an existing output file rather than deleting it. If a file already exists with the full previous file name, then it will be deleted before the current file is renamed to become the previous file. | [optional] 
**LogInterval** | Pointer to **string** | The duration between statistics collection and logging. A new line is logged to the output for each interval. Setting this value too small can have an impact on performance. | [optional] 
**CollectionInterval** | Pointer to **string** | Some of the calculated statistics, such as the average and maximum queue sizes, can use multiple samples within a log interval. This value controls how often samples are gathered. It should be a multiple of the log-interval. | [optional] 
**SuppressIfIdle** | Pointer to **bool** | If the server is idle during the specified interval, then do not log any output if this property is set to true. The server is idle if during the interval, no new connections were established, no operations were processed, and no operations are pending. | [optional] 
**HeaderPrefixPerColumn** | Pointer to **bool** | This property controls whether the header prefix, which applies to a group of columns, appears at the start of each column header or only the first column in a group. | [optional] 
**EmptyInsteadOfZero** | Pointer to **bool** | This property controls whether a value in the output is shown as empty if the value is zero. | [optional] 
**LinesBetweenHeader** | Pointer to **int32** | The number of lines to log between logging the header line that summarizes the columns in the table. | [optional] 
**IncludedLDAPStat** | Pointer to [**[]EnumpluginIncludedLDAPStatProp**](EnumpluginIncludedLDAPStatProp.md) |  | [optional] 
**IncludedResourceStat** | Pointer to [**[]EnumpluginIncludedResourceStatProp**](EnumpluginIncludedResourceStatProp.md) |  | [optional] 
**HistogramFormat** | Pointer to [**EnumpluginHistogramFormatProp**](EnumpluginHistogramFormatProp.md) |  | [optional] 
**HistogramOpType** | Pointer to [**[]EnumpluginHistogramOpTypeProp**](EnumpluginHistogramOpTypeProp.md) |  | [optional] 
**PerApplicationLDAPStats** | Pointer to [**EnumpluginPerApplicationLDAPStatsProp**](EnumpluginPerApplicationLDAPStatsProp.md) |  | [optional] 
**StatusSummaryInfo** | Pointer to [**EnumpluginStatusSummaryInfoProp**](EnumpluginStatusSummaryInfoProp.md) |  | [optional] 
**LdapChangelogInfo** | Pointer to [**EnumpluginLdapChangelogInfoProp**](EnumpluginLdapChangelogInfoProp.md) |  | [optional] 
**GaugeInfo** | Pointer to [**EnumpluginGaugeInfoProp**](EnumpluginGaugeInfoProp.md) |  | [optional] 
**LogFileFormat** | Pointer to [**EnumpluginLogFileFormatProp**](EnumpluginLogFileFormatProp.md) |  | [optional] 
**LogFile** | **string** | Specifies the log file location where the update records are written when the plug-in is in background-mode processing. | 
**LogFilePermissions** | Pointer to **string** | The UNIX permissions of the log files created by this Periodic Stats Logger Plugin. | [optional] 
**Append** | Pointer to **bool** | Specifies whether to append to existing log files. | [optional] 
**RotationPolicy** | Pointer to **[]string** | The rotation policy to use for the Periodic Stats Logger Plugin . | [optional] 
**RotationListener** | Pointer to **[]string** | A listener that should be notified whenever a log file is rotated out of service. | [optional] 
**RetentionPolicy** | Pointer to **[]string** | The retention policy to use for the Periodic Stats Logger Plugin . | [optional] 
**LoggingErrorBehavior** | Pointer to [**EnumpluginLoggingErrorBehaviorProp**](EnumpluginLoggingErrorBehaviorProp.md) |  | [optional] 
**LocalDBBackendInfo** | Pointer to [**EnumpluginLocalDBBackendInfoProp**](EnumpluginLocalDBBackendInfoProp.md) |  | [optional] 
**ReplicationInfo** | Pointer to [**EnumpluginReplicationInfoProp**](EnumpluginReplicationInfoProp.md) |  | [optional] 
**EntryCacheInfo** | Pointer to [**EnumpluginEntryCacheInfoProp**](EnumpluginEntryCacheInfoProp.md) |  | [optional] 
**HostInfo** | Pointer to [**[]EnumpluginHostInfoProp**](EnumpluginHostInfoProp.md) |  | [optional] 
**IncludedLDAPApplication** | Pointer to **[]string** | If statistics should not be included for all applications, this property names the subset of applications that should be included. | [optional] 
**DatetimeAttribute** | **string** | The LDAP attribute that determines when data should be deleted. This could store the expiration time, or it could store the creation time and the expiration-offset property specifies the duration before data is deleted. | 
**DatetimeJSONField** | Pointer to **string** | The top-level JSON field within the configured datetime-attribute that determines when data should be deleted. This could store the expiration time, or it could store the creation time and the expiration-offset property specifies the duration before data is deleted. | [optional] 
**DatetimeFormat** | Pointer to [**EnumpluginDatetimeFormatProp**](EnumpluginDatetimeFormatProp.md) |  | [optional] 
**CustomDatetimeFormat** | Pointer to **string** | When the datetime-format property is configured with a value of \&quot;custom\&quot;, this specifies the format (using a string compatible with the java.text.SimpleDateFormat class) that will be used to search for expired data. | [optional] 
**CustomTimezone** | Pointer to **string** | Specifies the time zone to use when generating a date string using the configured custom-datetime-format value. The provided value must be accepted by java.util.TimeZone.getTimeZone. | [optional] 
**ExpirationOffset** | **string** | Sessions whose last activity timestamp is older than this offset will be removed. | 
**PurgeBehavior** | Pointer to [**EnumpluginPurgeBehaviorProp**](EnumpluginPurgeBehaviorProp.md) |  | [optional] 
**NumMostExpensivePhasesShown** | Pointer to **int32** | This controls how many of the most expensive phases are included per operation type in the monitor entry. | [optional] 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Plugin. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Plugin. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**Server** | **[]string** | Specifies the LDAP external server(s) to which authentication attempts should be forwarded. | 
**ServerAccessMode** | Pointer to [**EnumpluginServerAccessModeProp**](EnumpluginServerAccessModeProp.md) |  | [optional] 
**DnMap** | Pointer to **[]string** | Specifies one or more DN mappings that may be used to transform bind DNs before attempting to bind to the external servers. | [optional] 
**BindDNPattern** | Pointer to **string** | A pattern to use to construct the bind DN for the simple bind request to send to the remote server. This may consist of a combination of static text and attribute values and other directives enclosed in curly braces.  For example, the value \&quot;cn&#x3D;{cn},ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot; indicates that the remote bind DN should be constructed from the text \&quot;cn&#x3D;\&quot; followed by the value of the local entry&#39;s cn attribute followed by the text \&quot;ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot;. If an attribute contains the value to use as the bind DN for pass-through authentication, then the pattern may simply be the name of that attribute in curly braces (e.g., if the seeAlso attribute contains the bind DN for the target user, then a bind DN pattern of \&quot;{seeAlso}\&quot; would be appropriate).  Note that a bind DN pattern can be used to construct a bind DN that is not actually a valid LDAP distinguished name. For example, if authentication is being passed through to a Microsoft Active Directory server, then a bind DN pattern could be used to construct a user principal name (UPN) as an alternative to a distinguished name. | [optional] 
**SearchBaseDN** | Pointer to **string** | The base DN to use when searching for the user entry using a filter constructed from the pattern defined in the search-filter-pattern property. If no base DN is specified, the null DN will be used as the search base DN. | [optional] 
**SearchFilterPattern** | Pointer to **string** | A pattern to use to construct a filter to use when searching an external server for the entry of the user as whom to bind. For example, \&quot;(mail&#x3D;{uid:ldapFilterEscape}@example.com)\&quot; would construct a search filter to search for a user whose entry in the local server contains a uid attribute whose value appears before \&quot;@example.com\&quot; in the mail attribute in the external server. Note that the \&quot;ldapFilterEscape\&quot; modifier should almost always be used with attributes specified in the pattern. | [optional] 
**InitialConnections** | Pointer to **int32** | Specifies the initial number of connections to establish to each external server against which authentication may be attempted. | [optional] 
**MaxConnections** | Pointer to **int32** | Specifies the maximum number of connections to maintain to each external server against which authentication may be attempted. This value must be greater than or equal to the value for the initial-connections property. | [optional] 
**SourceDN** | **string** | Specifies the source DN that may appear in client requests which should be remapped to the target DN. Note that the source DN must not be equal to the target DN. | 
**TargetDN** | **string** | Specifies the DN to which the source DN should be mapped. Note that the target DN must not be equal to the source DN. | 
**EnableAttributeMapping** | Pointer to **bool** | Indicates whether DN mapping should be applied to the values of attributes with appropriate syntaxes. | [optional] 
**MapAttribute** | Pointer to **[]string** | Specifies a set of specific attributes for which DN mapping should be applied. This will only be applicable if the enable-attribute-mapping property has a value of \&quot;true\&quot;. Any attributes listed must be defined in the server schema with either the distinguished name syntax or the name and optional UID syntax. | [optional] 
**EnableControlMapping** | Pointer to **bool** | Indicates whether mapping should be applied to attribute types that may be present in specific controls. If enabled, attribute mapping will only be applied for control types which are specifically supported by the attribute mapper plugin. | [optional] 
**AlwaysMapResponses** | Pointer to **bool** | Indicates whether the target attribute in response messages should always be remapped back to the source attribute. If this is \&quot;false\&quot;, then the mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request. | [optional] 
**ReferralBaseURL** | **[]string** | Specifies the base URL to use for the referrals generated by this plugin. It should include only the scheme, address, and port to use to communicate with the target server (e.g., \&quot;ldap://server.example.com:389/\&quot;). | 
**ContextName** | Pointer to **string** | The SNMP context name for this sub-agent. The context name must not be longer than 30 ASCII characters. Each server in a topology must have a unique SNMP context name. | [optional] 
**AgentxAddress** | Pointer to **string** | The hostname or IP address of the SNMP master agent. | [optional] 
**AgentxPort** | Pointer to **int32** | The port number on which the SNMP master agent will be contacted. | [optional] 
**NumWorkerThreads** | Pointer to **int32** | The number of worker threads to use to handle SNMP requests. | [optional] 
**SessionTimeout** | Pointer to **string** | Specifies the maximum amount of time to wait for a session to the master agent to be established. | [optional] 
**ConnectRetryMaxWait** | Pointer to **string** | The maximum amount of time to wait between attempts to establish a connection to the master agent. | [optional] 
**PingInterval** | Pointer to **string** | The amount of time between consecutive pings sent by the sub-agent on its connection to the master agent. A value of zero disables the sending of pings by the sub-agent. | [optional] 
**ValuePattern** | **[]string** | Specifies a pattern for constructing the values to use for the target attribute type. | 
**MultipleValuePatternBehavior** | Pointer to [**EnumpluginMultipleValuePatternBehaviorProp**](EnumpluginMultipleValuePatternBehaviorProp.md) |  | [optional] 
**MultiValuedAttributeBehavior** | Pointer to [**EnumpluginMultiValuedAttributeBehaviorProp**](EnumpluginMultiValuedAttributeBehaviorProp.md) |  | [optional] 
**TargetAttributeExistsDuringInitialPopulationBehavior** | Pointer to [**EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp**](EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp.md) |  | [optional] 
**UpdateSourceAttributeBehavior** | Pointer to [**EnumpluginUpdateSourceAttributeBehaviorProp**](EnumpluginUpdateSourceAttributeBehaviorProp.md) |  | [optional] 
**SourceAttributeRemovalBehavior** | Pointer to [**EnumpluginSourceAttributeRemovalBehaviorProp**](EnumpluginSourceAttributeRemovalBehaviorProp.md) |  | [optional] 
**UpdateTargetAttributeBehavior** | Pointer to [**EnumpluginUpdateTargetAttributeBehaviorProp**](EnumpluginUpdateTargetAttributeBehaviorProp.md) |  | [optional] 
**IncludeBaseDN** | Pointer to **[]string** | The set of base DNs below which composed values may be generated. | [optional] 
**ExcludeBaseDN** | Pointer to **[]string** | The set of base DNs below which composed values will not be generated. | [optional] 
**IncludeFilter** | Pointer to **[]string** | The set of search filters that identify entries for which composed values may be generated. | [optional] 
**ExcludeFilter** | Pointer to **[]string** | The set of search filters that identify entries for which composed values will not be generated. | [optional] 
**UpdatedEntryNewlyMatchesCriteriaBehavior** | Pointer to [**EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp**](EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp.md) |  | [optional] 
**UpdatedEntryNoLongerMatchesCriteriaBehavior** | Pointer to [**EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp**](EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp.md) |  | [optional] 
**SourceAttribute** | **string** | Specifies the source attribute type that may appear in client requests which should be remapped to the target attribute. Note that the source attribute type must be defined in the server schema and must not be equal to the target attribute type. | 
**TargetAttribute** | **string** | Specifies the target attribute type to which the source attribute type should be mapped. Note that the target attribute type must be defined in the server schema and must not be equal to the source attribute type. | 
**Delay** | **string** | The delay to inject for operations matching the associated criteria. | 
**ScriptClass** | **string** | The fully-qualified name of the Groovy class providing the logic for the Groovy Scripted Plugin. | 
**ScriptArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Scripted Plugin. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**PassThroughAuthenticationHandler** | **string** | The component used to manage authentication with the external authentication service. | 
**UpdateInterval** | Pointer to **string** | Specifies the interval in seconds when referential integrity updates are made. | [optional] 
**Type** | **[]string** | Specifies the type of attributes to check for value uniqueness. | 
**MultipleAttributeBehavior** | Pointer to [**EnumpluginMultipleAttributeBehaviorProp**](EnumpluginMultipleAttributeBehaviorProp.md) |  | [optional] 
**PreventConflictsWithSoftDeletedEntries** | Pointer to **bool** | Indicates whether this Unique Attribute Plugin should reject a change that would result in one or more conflicts, even if those conflicts only exist in soft-deleted entries. | [optional] 

## Methods

### NewAddPluginRequest

`func NewAddPluginRequest(pluginName string, schemas []EnumuniqueAttributePluginSchemaUrn, pluginType []EnumpluginPluginTypeProp, baseDN []string, filterPrefix string, enabled bool, filter string, attributeType []string, invokeGCTimeUtc []string, apiURL string, authURL string, oAuthClientID string, environmentID string, userMappingLocalAttribute []string, userMappingRemoteJSONField []string, scope EnumpluginScopeProp, outputFile string, logFile string, datetimeAttribute string, expirationOffset string, extensionClass string, server []string, sourceDN string, targetDN string, referralBaseURL []string, valuePattern []string, sourceAttribute string, targetAttribute string, delay string, scriptClass string, passThroughAuthenticationHandler string, type_ []string, ) *AddPluginRequest`

NewAddPluginRequest instantiates a new AddPluginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPluginRequestWithDefaults

`func NewAddPluginRequestWithDefaults() *AddPluginRequest`

NewAddPluginRequestWithDefaults instantiates a new AddPluginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *AddPluginRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *AddPluginRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *AddPluginRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetSchemas

`func (o *AddPluginRequest) GetSchemas() []EnumuniqueAttributePluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *AddPluginRequest) GetSchemasOk() (*[]EnumuniqueAttributePluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *AddPluginRequest) SetSchemas(v []EnumuniqueAttributePluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetPluginType

`func (o *AddPluginRequest) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *AddPluginRequest) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *AddPluginRequest) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetNumThreads

`func (o *AddPluginRequest) GetNumThreads() int32`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *AddPluginRequest) GetNumThreadsOk() (*int32, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *AddPluginRequest) SetNumThreads(v int32)`

SetNumThreads sets NumThreads field to given value.

### HasNumThreads

`func (o *AddPluginRequest) HasNumThreads() bool`

HasNumThreads returns a boolean if a field has been set.

### GetBaseDN

`func (o *AddPluginRequest) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *AddPluginRequest) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *AddPluginRequest) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetLowerBound

`func (o *AddPluginRequest) GetLowerBound() int32`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *AddPluginRequest) GetLowerBoundOk() (*int32, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *AddPluginRequest) SetLowerBound(v int32)`

SetLowerBound sets LowerBound field to given value.

### HasLowerBound

`func (o *AddPluginRequest) HasLowerBound() bool`

HasLowerBound returns a boolean if a field has been set.

### GetUpperBound

`func (o *AddPluginRequest) GetUpperBound() int32`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *AddPluginRequest) GetUpperBoundOk() (*int32, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *AddPluginRequest) SetUpperBound(v int32)`

SetUpperBound sets UpperBound field to given value.

### HasUpperBound

`func (o *AddPluginRequest) HasUpperBound() bool`

HasUpperBound returns a boolean if a field has been set.

### GetFilterPrefix

`func (o *AddPluginRequest) GetFilterPrefix() string`

GetFilterPrefix returns the FilterPrefix field if non-nil, zero value otherwise.

### GetFilterPrefixOk

`func (o *AddPluginRequest) GetFilterPrefixOk() (*string, bool)`

GetFilterPrefixOk returns a tuple with the FilterPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPrefix

`func (o *AddPluginRequest) SetFilterPrefix(v string)`

SetFilterPrefix sets FilterPrefix field to given value.


### GetFilterSuffix

`func (o *AddPluginRequest) GetFilterSuffix() string`

GetFilterSuffix returns the FilterSuffix field if non-nil, zero value otherwise.

### GetFilterSuffixOk

`func (o *AddPluginRequest) GetFilterSuffixOk() (*string, bool)`

GetFilterSuffixOk returns a tuple with the FilterSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSuffix

`func (o *AddPluginRequest) SetFilterSuffix(v string)`

SetFilterSuffix sets FilterSuffix field to given value.

### HasFilterSuffix

`func (o *AddPluginRequest) HasFilterSuffix() bool`

HasFilterSuffix returns a boolean if a field has been set.

### GetDescription

`func (o *AddPluginRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPluginRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPluginRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPluginRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPluginRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPluginRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPluginRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetInvokeForInternalOperations

`func (o *AddPluginRequest) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *AddPluginRequest) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *AddPluginRequest) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *AddPluginRequest) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetFilter

`func (o *AddPluginRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AddPluginRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AddPluginRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetAttributeType

`func (o *AddPluginRequest) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *AddPluginRequest) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *AddPluginRequest) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetPollingInterval

`func (o *AddPluginRequest) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *AddPluginRequest) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *AddPluginRequest) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *AddPluginRequest) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetPeerServerPriorityIndex

`func (o *AddPluginRequest) GetPeerServerPriorityIndex() int32`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *AddPluginRequest) GetPeerServerPriorityIndexOk() (*int32, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *AddPluginRequest) SetPeerServerPriorityIndex(v int32)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *AddPluginRequest) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetMaxUpdatesPerSecond

`func (o *AddPluginRequest) GetMaxUpdatesPerSecond() int32`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *AddPluginRequest) GetMaxUpdatesPerSecondOk() (*int32, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *AddPluginRequest) SetMaxUpdatesPerSecond(v int32)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.

### HasMaxUpdatesPerSecond

`func (o *AddPluginRequest) HasMaxUpdatesPerSecond() bool`

HasMaxUpdatesPerSecond returns a boolean if a field has been set.

### GetNumDeleteThreads

`func (o *AddPluginRequest) GetNumDeleteThreads() int32`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *AddPluginRequest) GetNumDeleteThreadsOk() (*int32, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *AddPluginRequest) SetNumDeleteThreads(v int32)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.

### HasNumDeleteThreads

`func (o *AddPluginRequest) HasNumDeleteThreads() bool`

HasNumDeleteThreads returns a boolean if a field has been set.

### GetInvokeGCDayOfWeek

`func (o *AddPluginRequest) GetInvokeGCDayOfWeek() []EnumpluginInvokeGCDayOfWeekProp`

GetInvokeGCDayOfWeek returns the InvokeGCDayOfWeek field if non-nil, zero value otherwise.

### GetInvokeGCDayOfWeekOk

`func (o *AddPluginRequest) GetInvokeGCDayOfWeekOk() (*[]EnumpluginInvokeGCDayOfWeekProp, bool)`

GetInvokeGCDayOfWeekOk returns a tuple with the InvokeGCDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeGCDayOfWeek

`func (o *AddPluginRequest) SetInvokeGCDayOfWeek(v []EnumpluginInvokeGCDayOfWeekProp)`

SetInvokeGCDayOfWeek sets InvokeGCDayOfWeek field to given value.

### HasInvokeGCDayOfWeek

`func (o *AddPluginRequest) HasInvokeGCDayOfWeek() bool`

HasInvokeGCDayOfWeek returns a boolean if a field has been set.

### GetInvokeGCTimeUtc

`func (o *AddPluginRequest) GetInvokeGCTimeUtc() []string`

GetInvokeGCTimeUtc returns the InvokeGCTimeUtc field if non-nil, zero value otherwise.

### GetInvokeGCTimeUtcOk

`func (o *AddPluginRequest) GetInvokeGCTimeUtcOk() (*[]string, bool)`

GetInvokeGCTimeUtcOk returns a tuple with the InvokeGCTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeGCTimeUtc

`func (o *AddPluginRequest) SetInvokeGCTimeUtc(v []string)`

SetInvokeGCTimeUtc sets InvokeGCTimeUtc field to given value.


### GetDelayAfterAlert

`func (o *AddPluginRequest) GetDelayAfterAlert() string`

GetDelayAfterAlert returns the DelayAfterAlert field if non-nil, zero value otherwise.

### GetDelayAfterAlertOk

`func (o *AddPluginRequest) GetDelayAfterAlertOk() (*string, bool)`

GetDelayAfterAlertOk returns a tuple with the DelayAfterAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayAfterAlert

`func (o *AddPluginRequest) SetDelayAfterAlert(v string)`

SetDelayAfterAlert sets DelayAfterAlert field to given value.

### HasDelayAfterAlert

`func (o *AddPluginRequest) HasDelayAfterAlert() bool`

HasDelayAfterAlert returns a boolean if a field has been set.

### GetDelayPostGC

`func (o *AddPluginRequest) GetDelayPostGC() string`

GetDelayPostGC returns the DelayPostGC field if non-nil, zero value otherwise.

### GetDelayPostGCOk

`func (o *AddPluginRequest) GetDelayPostGCOk() (*string, bool)`

GetDelayPostGCOk returns a tuple with the DelayPostGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayPostGC

`func (o *AddPluginRequest) SetDelayPostGC(v string)`

SetDelayPostGC sets DelayPostGC field to given value.

### HasDelayPostGC

`func (o *AddPluginRequest) HasDelayPostGC() bool`

HasDelayPostGC returns a boolean if a field has been set.

### GetApiURL

`func (o *AddPluginRequest) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *AddPluginRequest) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *AddPluginRequest) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.


### GetAuthURL

`func (o *AddPluginRequest) GetAuthURL() string`

GetAuthURL returns the AuthURL field if non-nil, zero value otherwise.

### GetAuthURLOk

`func (o *AddPluginRequest) GetAuthURLOk() (*string, bool)`

GetAuthURLOk returns a tuple with the AuthURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthURL

`func (o *AddPluginRequest) SetAuthURL(v string)`

SetAuthURL sets AuthURL field to given value.


### GetOAuthClientID

`func (o *AddPluginRequest) GetOAuthClientID() string`

GetOAuthClientID returns the OAuthClientID field if non-nil, zero value otherwise.

### GetOAuthClientIDOk

`func (o *AddPluginRequest) GetOAuthClientIDOk() (*string, bool)`

GetOAuthClientIDOk returns a tuple with the OAuthClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientID

`func (o *AddPluginRequest) SetOAuthClientID(v string)`

SetOAuthClientID sets OAuthClientID field to given value.


### GetOAuthClientSecret

`func (o *AddPluginRequest) GetOAuthClientSecret() string`

GetOAuthClientSecret returns the OAuthClientSecret field if non-nil, zero value otherwise.

### GetOAuthClientSecretOk

`func (o *AddPluginRequest) GetOAuthClientSecretOk() (*string, bool)`

GetOAuthClientSecretOk returns a tuple with the OAuthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecret

`func (o *AddPluginRequest) SetOAuthClientSecret(v string)`

SetOAuthClientSecret sets OAuthClientSecret field to given value.

### HasOAuthClientSecret

`func (o *AddPluginRequest) HasOAuthClientSecret() bool`

HasOAuthClientSecret returns a boolean if a field has been set.

### GetOAuthClientSecretPassphraseProvider

`func (o *AddPluginRequest) GetOAuthClientSecretPassphraseProvider() string`

GetOAuthClientSecretPassphraseProvider returns the OAuthClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetOAuthClientSecretPassphraseProviderOk

`func (o *AddPluginRequest) GetOAuthClientSecretPassphraseProviderOk() (*string, bool)`

GetOAuthClientSecretPassphraseProviderOk returns a tuple with the OAuthClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecretPassphraseProvider

`func (o *AddPluginRequest) SetOAuthClientSecretPassphraseProvider(v string)`

SetOAuthClientSecretPassphraseProvider sets OAuthClientSecretPassphraseProvider field to given value.

### HasOAuthClientSecretPassphraseProvider

`func (o *AddPluginRequest) HasOAuthClientSecretPassphraseProvider() bool`

HasOAuthClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *AddPluginRequest) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *AddPluginRequest) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *AddPluginRequest) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.


### GetIncludedLocalEntryBaseDN

`func (o *AddPluginRequest) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *AddPluginRequest) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *AddPluginRequest) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *AddPluginRequest) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *AddPluginRequest) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *AddPluginRequest) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *AddPluginRequest) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *AddPluginRequest) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *AddPluginRequest) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *AddPluginRequest) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *AddPluginRequest) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.

### HasRequestCriteria

`func (o *AddPluginRequest) HasRequestCriteria() bool`

HasRequestCriteria returns a boolean if a field has been set.

### GetTryLocalBind

`func (o *AddPluginRequest) GetTryLocalBind() bool`

GetTryLocalBind returns the TryLocalBind field if non-nil, zero value otherwise.

### GetTryLocalBindOk

`func (o *AddPluginRequest) GetTryLocalBindOk() (*bool, bool)`

GetTryLocalBindOk returns a tuple with the TryLocalBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryLocalBind

`func (o *AddPluginRequest) SetTryLocalBind(v bool)`

SetTryLocalBind sets TryLocalBind field to given value.

### HasTryLocalBind

`func (o *AddPluginRequest) HasTryLocalBind() bool`

HasTryLocalBind returns a boolean if a field has been set.

### GetOverrideLocalPassword

`func (o *AddPluginRequest) GetOverrideLocalPassword() bool`

GetOverrideLocalPassword returns the OverrideLocalPassword field if non-nil, zero value otherwise.

### GetOverrideLocalPasswordOk

`func (o *AddPluginRequest) GetOverrideLocalPasswordOk() (*bool, bool)`

GetOverrideLocalPasswordOk returns a tuple with the OverrideLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLocalPassword

`func (o *AddPluginRequest) SetOverrideLocalPassword(v bool)`

SetOverrideLocalPassword sets OverrideLocalPassword field to given value.

### HasOverrideLocalPassword

`func (o *AddPluginRequest) HasOverrideLocalPassword() bool`

HasOverrideLocalPassword returns a boolean if a field has been set.

### GetUpdateLocalPassword

`func (o *AddPluginRequest) GetUpdateLocalPassword() bool`

GetUpdateLocalPassword returns the UpdateLocalPassword field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordOk

`func (o *AddPluginRequest) GetUpdateLocalPasswordOk() (*bool, bool)`

GetUpdateLocalPasswordOk returns a tuple with the UpdateLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPassword

`func (o *AddPluginRequest) SetUpdateLocalPassword(v bool)`

SetUpdateLocalPassword sets UpdateLocalPassword field to given value.

### HasUpdateLocalPassword

`func (o *AddPluginRequest) HasUpdateLocalPassword() bool`

HasUpdateLocalPassword returns a boolean if a field has been set.

### GetUpdateLocalPasswordDN

`func (o *AddPluginRequest) GetUpdateLocalPasswordDN() string`

GetUpdateLocalPasswordDN returns the UpdateLocalPasswordDN field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordDNOk

`func (o *AddPluginRequest) GetUpdateLocalPasswordDNOk() (*string, bool)`

GetUpdateLocalPasswordDNOk returns a tuple with the UpdateLocalPasswordDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPasswordDN

`func (o *AddPluginRequest) SetUpdateLocalPasswordDN(v string)`

SetUpdateLocalPasswordDN sets UpdateLocalPasswordDN field to given value.

### HasUpdateLocalPasswordDN

`func (o *AddPluginRequest) HasUpdateLocalPasswordDN() bool`

HasUpdateLocalPasswordDN returns a boolean if a field has been set.

### GetAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPluginRequest) GetAllowLaxPassThroughAuthenticationPasswords() bool`

GetAllowLaxPassThroughAuthenticationPasswords returns the AllowLaxPassThroughAuthenticationPasswords field if non-nil, zero value otherwise.

### GetAllowLaxPassThroughAuthenticationPasswordsOk

`func (o *AddPluginRequest) GetAllowLaxPassThroughAuthenticationPasswordsOk() (*bool, bool)`

GetAllowLaxPassThroughAuthenticationPasswordsOk returns a tuple with the AllowLaxPassThroughAuthenticationPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPluginRequest) SetAllowLaxPassThroughAuthenticationPasswords(v bool)`

SetAllowLaxPassThroughAuthenticationPasswords sets AllowLaxPassThroughAuthenticationPasswords field to given value.

### HasAllowLaxPassThroughAuthenticationPasswords

`func (o *AddPluginRequest) HasAllowLaxPassThroughAuthenticationPasswords() bool`

HasAllowLaxPassThroughAuthenticationPasswords returns a boolean if a field has been set.

### GetIgnoredPasswordPolicyStateErrorCondition

`func (o *AddPluginRequest) GetIgnoredPasswordPolicyStateErrorCondition() []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp`

GetIgnoredPasswordPolicyStateErrorCondition returns the IgnoredPasswordPolicyStateErrorCondition field if non-nil, zero value otherwise.

### GetIgnoredPasswordPolicyStateErrorConditionOk

`func (o *AddPluginRequest) GetIgnoredPasswordPolicyStateErrorConditionOk() (*[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp, bool)`

GetIgnoredPasswordPolicyStateErrorConditionOk returns a tuple with the IgnoredPasswordPolicyStateErrorCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredPasswordPolicyStateErrorCondition

`func (o *AddPluginRequest) SetIgnoredPasswordPolicyStateErrorCondition(v []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp)`

SetIgnoredPasswordPolicyStateErrorCondition sets IgnoredPasswordPolicyStateErrorCondition field to given value.

### HasIgnoredPasswordPolicyStateErrorCondition

`func (o *AddPluginRequest) HasIgnoredPasswordPolicyStateErrorCondition() bool`

HasIgnoredPasswordPolicyStateErrorCondition returns a boolean if a field has been set.

### GetUserMappingLocalAttribute

`func (o *AddPluginRequest) GetUserMappingLocalAttribute() []string`

GetUserMappingLocalAttribute returns the UserMappingLocalAttribute field if non-nil, zero value otherwise.

### GetUserMappingLocalAttributeOk

`func (o *AddPluginRequest) GetUserMappingLocalAttributeOk() (*[]string, bool)`

GetUserMappingLocalAttributeOk returns a tuple with the UserMappingLocalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingLocalAttribute

`func (o *AddPluginRequest) SetUserMappingLocalAttribute(v []string)`

SetUserMappingLocalAttribute sets UserMappingLocalAttribute field to given value.


### GetUserMappingRemoteJSONField

`func (o *AddPluginRequest) GetUserMappingRemoteJSONField() []string`

GetUserMappingRemoteJSONField returns the UserMappingRemoteJSONField field if non-nil, zero value otherwise.

### GetUserMappingRemoteJSONFieldOk

`func (o *AddPluginRequest) GetUserMappingRemoteJSONFieldOk() (*[]string, bool)`

GetUserMappingRemoteJSONFieldOk returns a tuple with the UserMappingRemoteJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingRemoteJSONField

`func (o *AddPluginRequest) SetUserMappingRemoteJSONField(v []string)`

SetUserMappingRemoteJSONField sets UserMappingRemoteJSONField field to given value.


### GetAdditionalUserMappingSCIMFilter

`func (o *AddPluginRequest) GetAdditionalUserMappingSCIMFilter() string`

GetAdditionalUserMappingSCIMFilter returns the AdditionalUserMappingSCIMFilter field if non-nil, zero value otherwise.

### GetAdditionalUserMappingSCIMFilterOk

`func (o *AddPluginRequest) GetAdditionalUserMappingSCIMFilterOk() (*string, bool)`

GetAdditionalUserMappingSCIMFilterOk returns a tuple with the AdditionalUserMappingSCIMFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserMappingSCIMFilter

`func (o *AddPluginRequest) SetAdditionalUserMappingSCIMFilter(v string)`

SetAdditionalUserMappingSCIMFilter sets AdditionalUserMappingSCIMFilter field to given value.

### HasAdditionalUserMappingSCIMFilter

`func (o *AddPluginRequest) HasAdditionalUserMappingSCIMFilter() bool`

HasAdditionalUserMappingSCIMFilter returns a boolean if a field has been set.

### GetScope

`func (o *AddPluginRequest) GetScope() EnumpluginScopeProp`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AddPluginRequest) GetScopeOk() (*EnumpluginScopeProp, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AddPluginRequest) SetScope(v EnumpluginScopeProp)`

SetScope sets Scope field to given value.


### GetIncludeAttribute

`func (o *AddPluginRequest) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *AddPluginRequest) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *AddPluginRequest) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *AddPluginRequest) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetOutputFile

`func (o *AddPluginRequest) GetOutputFile() string`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *AddPluginRequest) GetOutputFileOk() (*string, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *AddPluginRequest) SetOutputFile(v string)`

SetOutputFile sets OutputFile field to given value.


### GetPreviousFileExtension

`func (o *AddPluginRequest) GetPreviousFileExtension() string`

GetPreviousFileExtension returns the PreviousFileExtension field if non-nil, zero value otherwise.

### GetPreviousFileExtensionOk

`func (o *AddPluginRequest) GetPreviousFileExtensionOk() (*string, bool)`

GetPreviousFileExtensionOk returns a tuple with the PreviousFileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousFileExtension

`func (o *AddPluginRequest) SetPreviousFileExtension(v string)`

SetPreviousFileExtension sets PreviousFileExtension field to given value.

### HasPreviousFileExtension

`func (o *AddPluginRequest) HasPreviousFileExtension() bool`

HasPreviousFileExtension returns a boolean if a field has been set.

### GetLogInterval

`func (o *AddPluginRequest) GetLogInterval() string`

GetLogInterval returns the LogInterval field if non-nil, zero value otherwise.

### GetLogIntervalOk

`func (o *AddPluginRequest) GetLogIntervalOk() (*string, bool)`

GetLogIntervalOk returns a tuple with the LogInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogInterval

`func (o *AddPluginRequest) SetLogInterval(v string)`

SetLogInterval sets LogInterval field to given value.

### HasLogInterval

`func (o *AddPluginRequest) HasLogInterval() bool`

HasLogInterval returns a boolean if a field has been set.

### GetCollectionInterval

`func (o *AddPluginRequest) GetCollectionInterval() string`

GetCollectionInterval returns the CollectionInterval field if non-nil, zero value otherwise.

### GetCollectionIntervalOk

`func (o *AddPluginRequest) GetCollectionIntervalOk() (*string, bool)`

GetCollectionIntervalOk returns a tuple with the CollectionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionInterval

`func (o *AddPluginRequest) SetCollectionInterval(v string)`

SetCollectionInterval sets CollectionInterval field to given value.

### HasCollectionInterval

`func (o *AddPluginRequest) HasCollectionInterval() bool`

HasCollectionInterval returns a boolean if a field has been set.

### GetSuppressIfIdle

`func (o *AddPluginRequest) GetSuppressIfIdle() bool`

GetSuppressIfIdle returns the SuppressIfIdle field if non-nil, zero value otherwise.

### GetSuppressIfIdleOk

`func (o *AddPluginRequest) GetSuppressIfIdleOk() (*bool, bool)`

GetSuppressIfIdleOk returns a tuple with the SuppressIfIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressIfIdle

`func (o *AddPluginRequest) SetSuppressIfIdle(v bool)`

SetSuppressIfIdle sets SuppressIfIdle field to given value.

### HasSuppressIfIdle

`func (o *AddPluginRequest) HasSuppressIfIdle() bool`

HasSuppressIfIdle returns a boolean if a field has been set.

### GetHeaderPrefixPerColumn

`func (o *AddPluginRequest) GetHeaderPrefixPerColumn() bool`

GetHeaderPrefixPerColumn returns the HeaderPrefixPerColumn field if non-nil, zero value otherwise.

### GetHeaderPrefixPerColumnOk

`func (o *AddPluginRequest) GetHeaderPrefixPerColumnOk() (*bool, bool)`

GetHeaderPrefixPerColumnOk returns a tuple with the HeaderPrefixPerColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPrefixPerColumn

`func (o *AddPluginRequest) SetHeaderPrefixPerColumn(v bool)`

SetHeaderPrefixPerColumn sets HeaderPrefixPerColumn field to given value.

### HasHeaderPrefixPerColumn

`func (o *AddPluginRequest) HasHeaderPrefixPerColumn() bool`

HasHeaderPrefixPerColumn returns a boolean if a field has been set.

### GetEmptyInsteadOfZero

`func (o *AddPluginRequest) GetEmptyInsteadOfZero() bool`

GetEmptyInsteadOfZero returns the EmptyInsteadOfZero field if non-nil, zero value otherwise.

### GetEmptyInsteadOfZeroOk

`func (o *AddPluginRequest) GetEmptyInsteadOfZeroOk() (*bool, bool)`

GetEmptyInsteadOfZeroOk returns a tuple with the EmptyInsteadOfZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyInsteadOfZero

`func (o *AddPluginRequest) SetEmptyInsteadOfZero(v bool)`

SetEmptyInsteadOfZero sets EmptyInsteadOfZero field to given value.

### HasEmptyInsteadOfZero

`func (o *AddPluginRequest) HasEmptyInsteadOfZero() bool`

HasEmptyInsteadOfZero returns a boolean if a field has been set.

### GetLinesBetweenHeader

`func (o *AddPluginRequest) GetLinesBetweenHeader() int32`

GetLinesBetweenHeader returns the LinesBetweenHeader field if non-nil, zero value otherwise.

### GetLinesBetweenHeaderOk

`func (o *AddPluginRequest) GetLinesBetweenHeaderOk() (*int32, bool)`

GetLinesBetweenHeaderOk returns a tuple with the LinesBetweenHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesBetweenHeader

`func (o *AddPluginRequest) SetLinesBetweenHeader(v int32)`

SetLinesBetweenHeader sets LinesBetweenHeader field to given value.

### HasLinesBetweenHeader

`func (o *AddPluginRequest) HasLinesBetweenHeader() bool`

HasLinesBetweenHeader returns a boolean if a field has been set.

### GetIncludedLDAPStat

`func (o *AddPluginRequest) GetIncludedLDAPStat() []EnumpluginIncludedLDAPStatProp`

GetIncludedLDAPStat returns the IncludedLDAPStat field if non-nil, zero value otherwise.

### GetIncludedLDAPStatOk

`func (o *AddPluginRequest) GetIncludedLDAPStatOk() (*[]EnumpluginIncludedLDAPStatProp, bool)`

GetIncludedLDAPStatOk returns a tuple with the IncludedLDAPStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPStat

`func (o *AddPluginRequest) SetIncludedLDAPStat(v []EnumpluginIncludedLDAPStatProp)`

SetIncludedLDAPStat sets IncludedLDAPStat field to given value.

### HasIncludedLDAPStat

`func (o *AddPluginRequest) HasIncludedLDAPStat() bool`

HasIncludedLDAPStat returns a boolean if a field has been set.

### GetIncludedResourceStat

`func (o *AddPluginRequest) GetIncludedResourceStat() []EnumpluginIncludedResourceStatProp`

GetIncludedResourceStat returns the IncludedResourceStat field if non-nil, zero value otherwise.

### GetIncludedResourceStatOk

`func (o *AddPluginRequest) GetIncludedResourceStatOk() (*[]EnumpluginIncludedResourceStatProp, bool)`

GetIncludedResourceStatOk returns a tuple with the IncludedResourceStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedResourceStat

`func (o *AddPluginRequest) SetIncludedResourceStat(v []EnumpluginIncludedResourceStatProp)`

SetIncludedResourceStat sets IncludedResourceStat field to given value.

### HasIncludedResourceStat

`func (o *AddPluginRequest) HasIncludedResourceStat() bool`

HasIncludedResourceStat returns a boolean if a field has been set.

### GetHistogramFormat

`func (o *AddPluginRequest) GetHistogramFormat() EnumpluginHistogramFormatProp`

GetHistogramFormat returns the HistogramFormat field if non-nil, zero value otherwise.

### GetHistogramFormatOk

`func (o *AddPluginRequest) GetHistogramFormatOk() (*EnumpluginHistogramFormatProp, bool)`

GetHistogramFormatOk returns a tuple with the HistogramFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramFormat

`func (o *AddPluginRequest) SetHistogramFormat(v EnumpluginHistogramFormatProp)`

SetHistogramFormat sets HistogramFormat field to given value.

### HasHistogramFormat

`func (o *AddPluginRequest) HasHistogramFormat() bool`

HasHistogramFormat returns a boolean if a field has been set.

### GetHistogramOpType

`func (o *AddPluginRequest) GetHistogramOpType() []EnumpluginHistogramOpTypeProp`

GetHistogramOpType returns the HistogramOpType field if non-nil, zero value otherwise.

### GetHistogramOpTypeOk

`func (o *AddPluginRequest) GetHistogramOpTypeOk() (*[]EnumpluginHistogramOpTypeProp, bool)`

GetHistogramOpTypeOk returns a tuple with the HistogramOpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramOpType

`func (o *AddPluginRequest) SetHistogramOpType(v []EnumpluginHistogramOpTypeProp)`

SetHistogramOpType sets HistogramOpType field to given value.

### HasHistogramOpType

`func (o *AddPluginRequest) HasHistogramOpType() bool`

HasHistogramOpType returns a boolean if a field has been set.

### GetPerApplicationLDAPStats

`func (o *AddPluginRequest) GetPerApplicationLDAPStats() EnumpluginPerApplicationLDAPStatsProp`

GetPerApplicationLDAPStats returns the PerApplicationLDAPStats field if non-nil, zero value otherwise.

### GetPerApplicationLDAPStatsOk

`func (o *AddPluginRequest) GetPerApplicationLDAPStatsOk() (*EnumpluginPerApplicationLDAPStatsProp, bool)`

GetPerApplicationLDAPStatsOk returns a tuple with the PerApplicationLDAPStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerApplicationLDAPStats

`func (o *AddPluginRequest) SetPerApplicationLDAPStats(v EnumpluginPerApplicationLDAPStatsProp)`

SetPerApplicationLDAPStats sets PerApplicationLDAPStats field to given value.

### HasPerApplicationLDAPStats

`func (o *AddPluginRequest) HasPerApplicationLDAPStats() bool`

HasPerApplicationLDAPStats returns a boolean if a field has been set.

### GetStatusSummaryInfo

`func (o *AddPluginRequest) GetStatusSummaryInfo() EnumpluginStatusSummaryInfoProp`

GetStatusSummaryInfo returns the StatusSummaryInfo field if non-nil, zero value otherwise.

### GetStatusSummaryInfoOk

`func (o *AddPluginRequest) GetStatusSummaryInfoOk() (*EnumpluginStatusSummaryInfoProp, bool)`

GetStatusSummaryInfoOk returns a tuple with the StatusSummaryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSummaryInfo

`func (o *AddPluginRequest) SetStatusSummaryInfo(v EnumpluginStatusSummaryInfoProp)`

SetStatusSummaryInfo sets StatusSummaryInfo field to given value.

### HasStatusSummaryInfo

`func (o *AddPluginRequest) HasStatusSummaryInfo() bool`

HasStatusSummaryInfo returns a boolean if a field has been set.

### GetLdapChangelogInfo

`func (o *AddPluginRequest) GetLdapChangelogInfo() EnumpluginLdapChangelogInfoProp`

GetLdapChangelogInfo returns the LdapChangelogInfo field if non-nil, zero value otherwise.

### GetLdapChangelogInfoOk

`func (o *AddPluginRequest) GetLdapChangelogInfoOk() (*EnumpluginLdapChangelogInfoProp, bool)`

GetLdapChangelogInfoOk returns a tuple with the LdapChangelogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapChangelogInfo

`func (o *AddPluginRequest) SetLdapChangelogInfo(v EnumpluginLdapChangelogInfoProp)`

SetLdapChangelogInfo sets LdapChangelogInfo field to given value.

### HasLdapChangelogInfo

`func (o *AddPluginRequest) HasLdapChangelogInfo() bool`

HasLdapChangelogInfo returns a boolean if a field has been set.

### GetGaugeInfo

`func (o *AddPluginRequest) GetGaugeInfo() EnumpluginGaugeInfoProp`

GetGaugeInfo returns the GaugeInfo field if non-nil, zero value otherwise.

### GetGaugeInfoOk

`func (o *AddPluginRequest) GetGaugeInfoOk() (*EnumpluginGaugeInfoProp, bool)`

GetGaugeInfoOk returns a tuple with the GaugeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeInfo

`func (o *AddPluginRequest) SetGaugeInfo(v EnumpluginGaugeInfoProp)`

SetGaugeInfo sets GaugeInfo field to given value.

### HasGaugeInfo

`func (o *AddPluginRequest) HasGaugeInfo() bool`

HasGaugeInfo returns a boolean if a field has been set.

### GetLogFileFormat

`func (o *AddPluginRequest) GetLogFileFormat() EnumpluginLogFileFormatProp`

GetLogFileFormat returns the LogFileFormat field if non-nil, zero value otherwise.

### GetLogFileFormatOk

`func (o *AddPluginRequest) GetLogFileFormatOk() (*EnumpluginLogFileFormatProp, bool)`

GetLogFileFormatOk returns a tuple with the LogFileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileFormat

`func (o *AddPluginRequest) SetLogFileFormat(v EnumpluginLogFileFormatProp)`

SetLogFileFormat sets LogFileFormat field to given value.

### HasLogFileFormat

`func (o *AddPluginRequest) HasLogFileFormat() bool`

HasLogFileFormat returns a boolean if a field has been set.

### GetLogFile

`func (o *AddPluginRequest) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *AddPluginRequest) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *AddPluginRequest) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.


### GetLogFilePermissions

`func (o *AddPluginRequest) GetLogFilePermissions() string`

GetLogFilePermissions returns the LogFilePermissions field if non-nil, zero value otherwise.

### GetLogFilePermissionsOk

`func (o *AddPluginRequest) GetLogFilePermissionsOk() (*string, bool)`

GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFilePermissions

`func (o *AddPluginRequest) SetLogFilePermissions(v string)`

SetLogFilePermissions sets LogFilePermissions field to given value.

### HasLogFilePermissions

`func (o *AddPluginRequest) HasLogFilePermissions() bool`

HasLogFilePermissions returns a boolean if a field has been set.

### GetAppend

`func (o *AddPluginRequest) GetAppend() bool`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *AddPluginRequest) GetAppendOk() (*bool, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *AddPluginRequest) SetAppend(v bool)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *AddPluginRequest) HasAppend() bool`

HasAppend returns a boolean if a field has been set.

### GetRotationPolicy

`func (o *AddPluginRequest) GetRotationPolicy() []string`

GetRotationPolicy returns the RotationPolicy field if non-nil, zero value otherwise.

### GetRotationPolicyOk

`func (o *AddPluginRequest) GetRotationPolicyOk() (*[]string, bool)`

GetRotationPolicyOk returns a tuple with the RotationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPolicy

`func (o *AddPluginRequest) SetRotationPolicy(v []string)`

SetRotationPolicy sets RotationPolicy field to given value.

### HasRotationPolicy

`func (o *AddPluginRequest) HasRotationPolicy() bool`

HasRotationPolicy returns a boolean if a field has been set.

### GetRotationListener

`func (o *AddPluginRequest) GetRotationListener() []string`

GetRotationListener returns the RotationListener field if non-nil, zero value otherwise.

### GetRotationListenerOk

`func (o *AddPluginRequest) GetRotationListenerOk() (*[]string, bool)`

GetRotationListenerOk returns a tuple with the RotationListener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationListener

`func (o *AddPluginRequest) SetRotationListener(v []string)`

SetRotationListener sets RotationListener field to given value.

### HasRotationListener

`func (o *AddPluginRequest) HasRotationListener() bool`

HasRotationListener returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *AddPluginRequest) GetRetentionPolicy() []string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *AddPluginRequest) GetRetentionPolicyOk() (*[]string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *AddPluginRequest) SetRetentionPolicy(v []string)`

SetRetentionPolicy sets RetentionPolicy field to given value.

### HasRetentionPolicy

`func (o *AddPluginRequest) HasRetentionPolicy() bool`

HasRetentionPolicy returns a boolean if a field has been set.

### GetLoggingErrorBehavior

`func (o *AddPluginRequest) GetLoggingErrorBehavior() EnumpluginLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *AddPluginRequest) GetLoggingErrorBehaviorOk() (*EnumpluginLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *AddPluginRequest) SetLoggingErrorBehavior(v EnumpluginLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *AddPluginRequest) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.

### GetLocalDBBackendInfo

`func (o *AddPluginRequest) GetLocalDBBackendInfo() EnumpluginLocalDBBackendInfoProp`

GetLocalDBBackendInfo returns the LocalDBBackendInfo field if non-nil, zero value otherwise.

### GetLocalDBBackendInfoOk

`func (o *AddPluginRequest) GetLocalDBBackendInfoOk() (*EnumpluginLocalDBBackendInfoProp, bool)`

GetLocalDBBackendInfoOk returns a tuple with the LocalDBBackendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDBBackendInfo

`func (o *AddPluginRequest) SetLocalDBBackendInfo(v EnumpluginLocalDBBackendInfoProp)`

SetLocalDBBackendInfo sets LocalDBBackendInfo field to given value.

### HasLocalDBBackendInfo

`func (o *AddPluginRequest) HasLocalDBBackendInfo() bool`

HasLocalDBBackendInfo returns a boolean if a field has been set.

### GetReplicationInfo

`func (o *AddPluginRequest) GetReplicationInfo() EnumpluginReplicationInfoProp`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *AddPluginRequest) GetReplicationInfoOk() (*EnumpluginReplicationInfoProp, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *AddPluginRequest) SetReplicationInfo(v EnumpluginReplicationInfoProp)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *AddPluginRequest) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetEntryCacheInfo

`func (o *AddPluginRequest) GetEntryCacheInfo() EnumpluginEntryCacheInfoProp`

GetEntryCacheInfo returns the EntryCacheInfo field if non-nil, zero value otherwise.

### GetEntryCacheInfoOk

`func (o *AddPluginRequest) GetEntryCacheInfoOk() (*EnumpluginEntryCacheInfoProp, bool)`

GetEntryCacheInfoOk returns a tuple with the EntryCacheInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCacheInfo

`func (o *AddPluginRequest) SetEntryCacheInfo(v EnumpluginEntryCacheInfoProp)`

SetEntryCacheInfo sets EntryCacheInfo field to given value.

### HasEntryCacheInfo

`func (o *AddPluginRequest) HasEntryCacheInfo() bool`

HasEntryCacheInfo returns a boolean if a field has been set.

### GetHostInfo

`func (o *AddPluginRequest) GetHostInfo() []EnumpluginHostInfoProp`

GetHostInfo returns the HostInfo field if non-nil, zero value otherwise.

### GetHostInfoOk

`func (o *AddPluginRequest) GetHostInfoOk() (*[]EnumpluginHostInfoProp, bool)`

GetHostInfoOk returns a tuple with the HostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfo

`func (o *AddPluginRequest) SetHostInfo(v []EnumpluginHostInfoProp)`

SetHostInfo sets HostInfo field to given value.

### HasHostInfo

`func (o *AddPluginRequest) HasHostInfo() bool`

HasHostInfo returns a boolean if a field has been set.

### GetIncludedLDAPApplication

`func (o *AddPluginRequest) GetIncludedLDAPApplication() []string`

GetIncludedLDAPApplication returns the IncludedLDAPApplication field if non-nil, zero value otherwise.

### GetIncludedLDAPApplicationOk

`func (o *AddPluginRequest) GetIncludedLDAPApplicationOk() (*[]string, bool)`

GetIncludedLDAPApplicationOk returns a tuple with the IncludedLDAPApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPApplication

`func (o *AddPluginRequest) SetIncludedLDAPApplication(v []string)`

SetIncludedLDAPApplication sets IncludedLDAPApplication field to given value.

### HasIncludedLDAPApplication

`func (o *AddPluginRequest) HasIncludedLDAPApplication() bool`

HasIncludedLDAPApplication returns a boolean if a field has been set.

### GetDatetimeAttribute

`func (o *AddPluginRequest) GetDatetimeAttribute() string`

GetDatetimeAttribute returns the DatetimeAttribute field if non-nil, zero value otherwise.

### GetDatetimeAttributeOk

`func (o *AddPluginRequest) GetDatetimeAttributeOk() (*string, bool)`

GetDatetimeAttributeOk returns a tuple with the DatetimeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeAttribute

`func (o *AddPluginRequest) SetDatetimeAttribute(v string)`

SetDatetimeAttribute sets DatetimeAttribute field to given value.


### GetDatetimeJSONField

`func (o *AddPluginRequest) GetDatetimeJSONField() string`

GetDatetimeJSONField returns the DatetimeJSONField field if non-nil, zero value otherwise.

### GetDatetimeJSONFieldOk

`func (o *AddPluginRequest) GetDatetimeJSONFieldOk() (*string, bool)`

GetDatetimeJSONFieldOk returns a tuple with the DatetimeJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeJSONField

`func (o *AddPluginRequest) SetDatetimeJSONField(v string)`

SetDatetimeJSONField sets DatetimeJSONField field to given value.

### HasDatetimeJSONField

`func (o *AddPluginRequest) HasDatetimeJSONField() bool`

HasDatetimeJSONField returns a boolean if a field has been set.

### GetDatetimeFormat

`func (o *AddPluginRequest) GetDatetimeFormat() EnumpluginDatetimeFormatProp`

GetDatetimeFormat returns the DatetimeFormat field if non-nil, zero value otherwise.

### GetDatetimeFormatOk

`func (o *AddPluginRequest) GetDatetimeFormatOk() (*EnumpluginDatetimeFormatProp, bool)`

GetDatetimeFormatOk returns a tuple with the DatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeFormat

`func (o *AddPluginRequest) SetDatetimeFormat(v EnumpluginDatetimeFormatProp)`

SetDatetimeFormat sets DatetimeFormat field to given value.

### HasDatetimeFormat

`func (o *AddPluginRequest) HasDatetimeFormat() bool`

HasDatetimeFormat returns a boolean if a field has been set.

### GetCustomDatetimeFormat

`func (o *AddPluginRequest) GetCustomDatetimeFormat() string`

GetCustomDatetimeFormat returns the CustomDatetimeFormat field if non-nil, zero value otherwise.

### GetCustomDatetimeFormatOk

`func (o *AddPluginRequest) GetCustomDatetimeFormatOk() (*string, bool)`

GetCustomDatetimeFormatOk returns a tuple with the CustomDatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDatetimeFormat

`func (o *AddPluginRequest) SetCustomDatetimeFormat(v string)`

SetCustomDatetimeFormat sets CustomDatetimeFormat field to given value.

### HasCustomDatetimeFormat

`func (o *AddPluginRequest) HasCustomDatetimeFormat() bool`

HasCustomDatetimeFormat returns a boolean if a field has been set.

### GetCustomTimezone

`func (o *AddPluginRequest) GetCustomTimezone() string`

GetCustomTimezone returns the CustomTimezone field if non-nil, zero value otherwise.

### GetCustomTimezoneOk

`func (o *AddPluginRequest) GetCustomTimezoneOk() (*string, bool)`

GetCustomTimezoneOk returns a tuple with the CustomTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimezone

`func (o *AddPluginRequest) SetCustomTimezone(v string)`

SetCustomTimezone sets CustomTimezone field to given value.

### HasCustomTimezone

`func (o *AddPluginRequest) HasCustomTimezone() bool`

HasCustomTimezone returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *AddPluginRequest) GetExpirationOffset() string`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *AddPluginRequest) GetExpirationOffsetOk() (*string, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *AddPluginRequest) SetExpirationOffset(v string)`

SetExpirationOffset sets ExpirationOffset field to given value.


### GetPurgeBehavior

`func (o *AddPluginRequest) GetPurgeBehavior() EnumpluginPurgeBehaviorProp`

GetPurgeBehavior returns the PurgeBehavior field if non-nil, zero value otherwise.

### GetPurgeBehaviorOk

`func (o *AddPluginRequest) GetPurgeBehaviorOk() (*EnumpluginPurgeBehaviorProp, bool)`

GetPurgeBehaviorOk returns a tuple with the PurgeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeBehavior

`func (o *AddPluginRequest) SetPurgeBehavior(v EnumpluginPurgeBehaviorProp)`

SetPurgeBehavior sets PurgeBehavior field to given value.

### HasPurgeBehavior

`func (o *AddPluginRequest) HasPurgeBehavior() bool`

HasPurgeBehavior returns a boolean if a field has been set.

### GetNumMostExpensivePhasesShown

`func (o *AddPluginRequest) GetNumMostExpensivePhasesShown() int32`

GetNumMostExpensivePhasesShown returns the NumMostExpensivePhasesShown field if non-nil, zero value otherwise.

### GetNumMostExpensivePhasesShownOk

`func (o *AddPluginRequest) GetNumMostExpensivePhasesShownOk() (*int32, bool)`

GetNumMostExpensivePhasesShownOk returns a tuple with the NumMostExpensivePhasesShown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMostExpensivePhasesShown

`func (o *AddPluginRequest) SetNumMostExpensivePhasesShown(v int32)`

SetNumMostExpensivePhasesShown sets NumMostExpensivePhasesShown field to given value.

### HasNumMostExpensivePhasesShown

`func (o *AddPluginRequest) HasNumMostExpensivePhasesShown() bool`

HasNumMostExpensivePhasesShown returns a boolean if a field has been set.

### GetExtensionClass

`func (o *AddPluginRequest) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *AddPluginRequest) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *AddPluginRequest) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *AddPluginRequest) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *AddPluginRequest) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *AddPluginRequest) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *AddPluginRequest) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetServer

`func (o *AddPluginRequest) GetServer() []string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddPluginRequest) GetServerOk() (*[]string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddPluginRequest) SetServer(v []string)`

SetServer sets Server field to given value.


### GetServerAccessMode

`func (o *AddPluginRequest) GetServerAccessMode() EnumpluginServerAccessModeProp`

GetServerAccessMode returns the ServerAccessMode field if non-nil, zero value otherwise.

### GetServerAccessModeOk

`func (o *AddPluginRequest) GetServerAccessModeOk() (*EnumpluginServerAccessModeProp, bool)`

GetServerAccessModeOk returns a tuple with the ServerAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAccessMode

`func (o *AddPluginRequest) SetServerAccessMode(v EnumpluginServerAccessModeProp)`

SetServerAccessMode sets ServerAccessMode field to given value.

### HasServerAccessMode

`func (o *AddPluginRequest) HasServerAccessMode() bool`

HasServerAccessMode returns a boolean if a field has been set.

### GetDnMap

`func (o *AddPluginRequest) GetDnMap() []string`

GetDnMap returns the DnMap field if non-nil, zero value otherwise.

### GetDnMapOk

`func (o *AddPluginRequest) GetDnMapOk() (*[]string, bool)`

GetDnMapOk returns a tuple with the DnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnMap

`func (o *AddPluginRequest) SetDnMap(v []string)`

SetDnMap sets DnMap field to given value.

### HasDnMap

`func (o *AddPluginRequest) HasDnMap() bool`

HasDnMap returns a boolean if a field has been set.

### GetBindDNPattern

`func (o *AddPluginRequest) GetBindDNPattern() string`

GetBindDNPattern returns the BindDNPattern field if non-nil, zero value otherwise.

### GetBindDNPatternOk

`func (o *AddPluginRequest) GetBindDNPatternOk() (*string, bool)`

GetBindDNPatternOk returns a tuple with the BindDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDNPattern

`func (o *AddPluginRequest) SetBindDNPattern(v string)`

SetBindDNPattern sets BindDNPattern field to given value.

### HasBindDNPattern

`func (o *AddPluginRequest) HasBindDNPattern() bool`

HasBindDNPattern returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *AddPluginRequest) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *AddPluginRequest) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *AddPluginRequest) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.

### HasSearchBaseDN

`func (o *AddPluginRequest) HasSearchBaseDN() bool`

HasSearchBaseDN returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *AddPluginRequest) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *AddPluginRequest) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *AddPluginRequest) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *AddPluginRequest) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetInitialConnections

`func (o *AddPluginRequest) GetInitialConnections() int32`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *AddPluginRequest) GetInitialConnectionsOk() (*int32, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *AddPluginRequest) SetInitialConnections(v int32)`

SetInitialConnections sets InitialConnections field to given value.

### HasInitialConnections

`func (o *AddPluginRequest) HasInitialConnections() bool`

HasInitialConnections returns a boolean if a field has been set.

### GetMaxConnections

`func (o *AddPluginRequest) GetMaxConnections() int32`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *AddPluginRequest) GetMaxConnectionsOk() (*int32, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *AddPluginRequest) SetMaxConnections(v int32)`

SetMaxConnections sets MaxConnections field to given value.

### HasMaxConnections

`func (o *AddPluginRequest) HasMaxConnections() bool`

HasMaxConnections returns a boolean if a field has been set.

### GetSourceDN

`func (o *AddPluginRequest) GetSourceDN() string`

GetSourceDN returns the SourceDN field if non-nil, zero value otherwise.

### GetSourceDNOk

`func (o *AddPluginRequest) GetSourceDNOk() (*string, bool)`

GetSourceDNOk returns a tuple with the SourceDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDN

`func (o *AddPluginRequest) SetSourceDN(v string)`

SetSourceDN sets SourceDN field to given value.


### GetTargetDN

`func (o *AddPluginRequest) GetTargetDN() string`

GetTargetDN returns the TargetDN field if non-nil, zero value otherwise.

### GetTargetDNOk

`func (o *AddPluginRequest) GetTargetDNOk() (*string, bool)`

GetTargetDNOk returns a tuple with the TargetDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDN

`func (o *AddPluginRequest) SetTargetDN(v string)`

SetTargetDN sets TargetDN field to given value.


### GetEnableAttributeMapping

`func (o *AddPluginRequest) GetEnableAttributeMapping() bool`

GetEnableAttributeMapping returns the EnableAttributeMapping field if non-nil, zero value otherwise.

### GetEnableAttributeMappingOk

`func (o *AddPluginRequest) GetEnableAttributeMappingOk() (*bool, bool)`

GetEnableAttributeMappingOk returns a tuple with the EnableAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAttributeMapping

`func (o *AddPluginRequest) SetEnableAttributeMapping(v bool)`

SetEnableAttributeMapping sets EnableAttributeMapping field to given value.

### HasEnableAttributeMapping

`func (o *AddPluginRequest) HasEnableAttributeMapping() bool`

HasEnableAttributeMapping returns a boolean if a field has been set.

### GetMapAttribute

`func (o *AddPluginRequest) GetMapAttribute() []string`

GetMapAttribute returns the MapAttribute field if non-nil, zero value otherwise.

### GetMapAttributeOk

`func (o *AddPluginRequest) GetMapAttributeOk() (*[]string, bool)`

GetMapAttributeOk returns a tuple with the MapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAttribute

`func (o *AddPluginRequest) SetMapAttribute(v []string)`

SetMapAttribute sets MapAttribute field to given value.

### HasMapAttribute

`func (o *AddPluginRequest) HasMapAttribute() bool`

HasMapAttribute returns a boolean if a field has been set.

### GetEnableControlMapping

`func (o *AddPluginRequest) GetEnableControlMapping() bool`

GetEnableControlMapping returns the EnableControlMapping field if non-nil, zero value otherwise.

### GetEnableControlMappingOk

`func (o *AddPluginRequest) GetEnableControlMappingOk() (*bool, bool)`

GetEnableControlMappingOk returns a tuple with the EnableControlMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableControlMapping

`func (o *AddPluginRequest) SetEnableControlMapping(v bool)`

SetEnableControlMapping sets EnableControlMapping field to given value.

### HasEnableControlMapping

`func (o *AddPluginRequest) HasEnableControlMapping() bool`

HasEnableControlMapping returns a boolean if a field has been set.

### GetAlwaysMapResponses

`func (o *AddPluginRequest) GetAlwaysMapResponses() bool`

GetAlwaysMapResponses returns the AlwaysMapResponses field if non-nil, zero value otherwise.

### GetAlwaysMapResponsesOk

`func (o *AddPluginRequest) GetAlwaysMapResponsesOk() (*bool, bool)`

GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysMapResponses

`func (o *AddPluginRequest) SetAlwaysMapResponses(v bool)`

SetAlwaysMapResponses sets AlwaysMapResponses field to given value.

### HasAlwaysMapResponses

`func (o *AddPluginRequest) HasAlwaysMapResponses() bool`

HasAlwaysMapResponses returns a boolean if a field has been set.

### GetReferralBaseURL

`func (o *AddPluginRequest) GetReferralBaseURL() []string`

GetReferralBaseURL returns the ReferralBaseURL field if non-nil, zero value otherwise.

### GetReferralBaseURLOk

`func (o *AddPluginRequest) GetReferralBaseURLOk() (*[]string, bool)`

GetReferralBaseURLOk returns a tuple with the ReferralBaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralBaseURL

`func (o *AddPluginRequest) SetReferralBaseURL(v []string)`

SetReferralBaseURL sets ReferralBaseURL field to given value.


### GetContextName

`func (o *AddPluginRequest) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *AddPluginRequest) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *AddPluginRequest) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *AddPluginRequest) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### GetAgentxAddress

`func (o *AddPluginRequest) GetAgentxAddress() string`

GetAgentxAddress returns the AgentxAddress field if non-nil, zero value otherwise.

### GetAgentxAddressOk

`func (o *AddPluginRequest) GetAgentxAddressOk() (*string, bool)`

GetAgentxAddressOk returns a tuple with the AgentxAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxAddress

`func (o *AddPluginRequest) SetAgentxAddress(v string)`

SetAgentxAddress sets AgentxAddress field to given value.

### HasAgentxAddress

`func (o *AddPluginRequest) HasAgentxAddress() bool`

HasAgentxAddress returns a boolean if a field has been set.

### GetAgentxPort

`func (o *AddPluginRequest) GetAgentxPort() int32`

GetAgentxPort returns the AgentxPort field if non-nil, zero value otherwise.

### GetAgentxPortOk

`func (o *AddPluginRequest) GetAgentxPortOk() (*int32, bool)`

GetAgentxPortOk returns a tuple with the AgentxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxPort

`func (o *AddPluginRequest) SetAgentxPort(v int32)`

SetAgentxPort sets AgentxPort field to given value.

### HasAgentxPort

`func (o *AddPluginRequest) HasAgentxPort() bool`

HasAgentxPort returns a boolean if a field has been set.

### GetNumWorkerThreads

`func (o *AddPluginRequest) GetNumWorkerThreads() int32`

GetNumWorkerThreads returns the NumWorkerThreads field if non-nil, zero value otherwise.

### GetNumWorkerThreadsOk

`func (o *AddPluginRequest) GetNumWorkerThreadsOk() (*int32, bool)`

GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkerThreads

`func (o *AddPluginRequest) SetNumWorkerThreads(v int32)`

SetNumWorkerThreads sets NumWorkerThreads field to given value.

### HasNumWorkerThreads

`func (o *AddPluginRequest) HasNumWorkerThreads() bool`

HasNumWorkerThreads returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *AddPluginRequest) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *AddPluginRequest) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *AddPluginRequest) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *AddPluginRequest) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetConnectRetryMaxWait

`func (o *AddPluginRequest) GetConnectRetryMaxWait() string`

GetConnectRetryMaxWait returns the ConnectRetryMaxWait field if non-nil, zero value otherwise.

### GetConnectRetryMaxWaitOk

`func (o *AddPluginRequest) GetConnectRetryMaxWaitOk() (*string, bool)`

GetConnectRetryMaxWaitOk returns a tuple with the ConnectRetryMaxWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectRetryMaxWait

`func (o *AddPluginRequest) SetConnectRetryMaxWait(v string)`

SetConnectRetryMaxWait sets ConnectRetryMaxWait field to given value.

### HasConnectRetryMaxWait

`func (o *AddPluginRequest) HasConnectRetryMaxWait() bool`

HasConnectRetryMaxWait returns a boolean if a field has been set.

### GetPingInterval

`func (o *AddPluginRequest) GetPingInterval() string`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *AddPluginRequest) GetPingIntervalOk() (*string, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *AddPluginRequest) SetPingInterval(v string)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *AddPluginRequest) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### GetValuePattern

`func (o *AddPluginRequest) GetValuePattern() []string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *AddPluginRequest) GetValuePatternOk() (*[]string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *AddPluginRequest) SetValuePattern(v []string)`

SetValuePattern sets ValuePattern field to given value.


### GetMultipleValuePatternBehavior

`func (o *AddPluginRequest) GetMultipleValuePatternBehavior() EnumpluginMultipleValuePatternBehaviorProp`

GetMultipleValuePatternBehavior returns the MultipleValuePatternBehavior field if non-nil, zero value otherwise.

### GetMultipleValuePatternBehaviorOk

`func (o *AddPluginRequest) GetMultipleValuePatternBehaviorOk() (*EnumpluginMultipleValuePatternBehaviorProp, bool)`

GetMultipleValuePatternBehaviorOk returns a tuple with the MultipleValuePatternBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleValuePatternBehavior

`func (o *AddPluginRequest) SetMultipleValuePatternBehavior(v EnumpluginMultipleValuePatternBehaviorProp)`

SetMultipleValuePatternBehavior sets MultipleValuePatternBehavior field to given value.

### HasMultipleValuePatternBehavior

`func (o *AddPluginRequest) HasMultipleValuePatternBehavior() bool`

HasMultipleValuePatternBehavior returns a boolean if a field has been set.

### GetMultiValuedAttributeBehavior

`func (o *AddPluginRequest) GetMultiValuedAttributeBehavior() EnumpluginMultiValuedAttributeBehaviorProp`

GetMultiValuedAttributeBehavior returns the MultiValuedAttributeBehavior field if non-nil, zero value otherwise.

### GetMultiValuedAttributeBehaviorOk

`func (o *AddPluginRequest) GetMultiValuedAttributeBehaviorOk() (*EnumpluginMultiValuedAttributeBehaviorProp, bool)`

GetMultiValuedAttributeBehaviorOk returns a tuple with the MultiValuedAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValuedAttributeBehavior

`func (o *AddPluginRequest) SetMultiValuedAttributeBehavior(v EnumpluginMultiValuedAttributeBehaviorProp)`

SetMultiValuedAttributeBehavior sets MultiValuedAttributeBehavior field to given value.

### HasMultiValuedAttributeBehavior

`func (o *AddPluginRequest) HasMultiValuedAttributeBehavior() bool`

HasMultiValuedAttributeBehavior returns a boolean if a field has been set.

### GetTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *AddPluginRequest) GetTargetAttributeExistsDuringInitialPopulationBehavior() EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp`

GetTargetAttributeExistsDuringInitialPopulationBehavior returns the TargetAttributeExistsDuringInitialPopulationBehavior field if non-nil, zero value otherwise.

### GetTargetAttributeExistsDuringInitialPopulationBehaviorOk

`func (o *AddPluginRequest) GetTargetAttributeExistsDuringInitialPopulationBehaviorOk() (*EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp, bool)`

GetTargetAttributeExistsDuringInitialPopulationBehaviorOk returns a tuple with the TargetAttributeExistsDuringInitialPopulationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *AddPluginRequest) SetTargetAttributeExistsDuringInitialPopulationBehavior(v EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp)`

SetTargetAttributeExistsDuringInitialPopulationBehavior sets TargetAttributeExistsDuringInitialPopulationBehavior field to given value.

### HasTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *AddPluginRequest) HasTargetAttributeExistsDuringInitialPopulationBehavior() bool`

HasTargetAttributeExistsDuringInitialPopulationBehavior returns a boolean if a field has been set.

### GetUpdateSourceAttributeBehavior

`func (o *AddPluginRequest) GetUpdateSourceAttributeBehavior() EnumpluginUpdateSourceAttributeBehaviorProp`

GetUpdateSourceAttributeBehavior returns the UpdateSourceAttributeBehavior field if non-nil, zero value otherwise.

### GetUpdateSourceAttributeBehaviorOk

`func (o *AddPluginRequest) GetUpdateSourceAttributeBehaviorOk() (*EnumpluginUpdateSourceAttributeBehaviorProp, bool)`

GetUpdateSourceAttributeBehaviorOk returns a tuple with the UpdateSourceAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSourceAttributeBehavior

`func (o *AddPluginRequest) SetUpdateSourceAttributeBehavior(v EnumpluginUpdateSourceAttributeBehaviorProp)`

SetUpdateSourceAttributeBehavior sets UpdateSourceAttributeBehavior field to given value.

### HasUpdateSourceAttributeBehavior

`func (o *AddPluginRequest) HasUpdateSourceAttributeBehavior() bool`

HasUpdateSourceAttributeBehavior returns a boolean if a field has been set.

### GetSourceAttributeRemovalBehavior

`func (o *AddPluginRequest) GetSourceAttributeRemovalBehavior() EnumpluginSourceAttributeRemovalBehaviorProp`

GetSourceAttributeRemovalBehavior returns the SourceAttributeRemovalBehavior field if non-nil, zero value otherwise.

### GetSourceAttributeRemovalBehaviorOk

`func (o *AddPluginRequest) GetSourceAttributeRemovalBehaviorOk() (*EnumpluginSourceAttributeRemovalBehaviorProp, bool)`

GetSourceAttributeRemovalBehaviorOk returns a tuple with the SourceAttributeRemovalBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttributeRemovalBehavior

`func (o *AddPluginRequest) SetSourceAttributeRemovalBehavior(v EnumpluginSourceAttributeRemovalBehaviorProp)`

SetSourceAttributeRemovalBehavior sets SourceAttributeRemovalBehavior field to given value.

### HasSourceAttributeRemovalBehavior

`func (o *AddPluginRequest) HasSourceAttributeRemovalBehavior() bool`

HasSourceAttributeRemovalBehavior returns a boolean if a field has been set.

### GetUpdateTargetAttributeBehavior

`func (o *AddPluginRequest) GetUpdateTargetAttributeBehavior() EnumpluginUpdateTargetAttributeBehaviorProp`

GetUpdateTargetAttributeBehavior returns the UpdateTargetAttributeBehavior field if non-nil, zero value otherwise.

### GetUpdateTargetAttributeBehaviorOk

`func (o *AddPluginRequest) GetUpdateTargetAttributeBehaviorOk() (*EnumpluginUpdateTargetAttributeBehaviorProp, bool)`

GetUpdateTargetAttributeBehaviorOk returns a tuple with the UpdateTargetAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTargetAttributeBehavior

`func (o *AddPluginRequest) SetUpdateTargetAttributeBehavior(v EnumpluginUpdateTargetAttributeBehaviorProp)`

SetUpdateTargetAttributeBehavior sets UpdateTargetAttributeBehavior field to given value.

### HasUpdateTargetAttributeBehavior

`func (o *AddPluginRequest) HasUpdateTargetAttributeBehavior() bool`

HasUpdateTargetAttributeBehavior returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *AddPluginRequest) GetIncludeBaseDN() []string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *AddPluginRequest) GetIncludeBaseDNOk() (*[]string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *AddPluginRequest) SetIncludeBaseDN(v []string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *AddPluginRequest) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetExcludeBaseDN

`func (o *AddPluginRequest) GetExcludeBaseDN() []string`

GetExcludeBaseDN returns the ExcludeBaseDN field if non-nil, zero value otherwise.

### GetExcludeBaseDNOk

`func (o *AddPluginRequest) GetExcludeBaseDNOk() (*[]string, bool)`

GetExcludeBaseDNOk returns a tuple with the ExcludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBaseDN

`func (o *AddPluginRequest) SetExcludeBaseDN(v []string)`

SetExcludeBaseDN sets ExcludeBaseDN field to given value.

### HasExcludeBaseDN

`func (o *AddPluginRequest) HasExcludeBaseDN() bool`

HasExcludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *AddPluginRequest) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *AddPluginRequest) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *AddPluginRequest) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *AddPluginRequest) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *AddPluginRequest) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *AddPluginRequest) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *AddPluginRequest) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *AddPluginRequest) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *AddPluginRequest) GetUpdatedEntryNewlyMatchesCriteriaBehavior() EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp`

GetUpdatedEntryNewlyMatchesCriteriaBehavior returns the UpdatedEntryNewlyMatchesCriteriaBehavior field if non-nil, zero value otherwise.

### GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk

`func (o *AddPluginRequest) GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp, bool)`

GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNewlyMatchesCriteriaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *AddPluginRequest) SetUpdatedEntryNewlyMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp)`

SetUpdatedEntryNewlyMatchesCriteriaBehavior sets UpdatedEntryNewlyMatchesCriteriaBehavior field to given value.

### HasUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *AddPluginRequest) HasUpdatedEntryNewlyMatchesCriteriaBehavior() bool`

HasUpdatedEntryNewlyMatchesCriteriaBehavior returns a boolean if a field has been set.

### GetUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *AddPluginRequest) GetUpdatedEntryNoLongerMatchesCriteriaBehavior() EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp`

GetUpdatedEntryNoLongerMatchesCriteriaBehavior returns the UpdatedEntryNoLongerMatchesCriteriaBehavior field if non-nil, zero value otherwise.

### GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk

`func (o *AddPluginRequest) GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp, bool)`

GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNoLongerMatchesCriteriaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *AddPluginRequest) SetUpdatedEntryNoLongerMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp)`

SetUpdatedEntryNoLongerMatchesCriteriaBehavior sets UpdatedEntryNoLongerMatchesCriteriaBehavior field to given value.

### HasUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *AddPluginRequest) HasUpdatedEntryNoLongerMatchesCriteriaBehavior() bool`

HasUpdatedEntryNoLongerMatchesCriteriaBehavior returns a boolean if a field has been set.

### GetSourceAttribute

`func (o *AddPluginRequest) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *AddPluginRequest) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *AddPluginRequest) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *AddPluginRequest) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *AddPluginRequest) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *AddPluginRequest) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetDelay

`func (o *AddPluginRequest) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *AddPluginRequest) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *AddPluginRequest) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetScriptClass

`func (o *AddPluginRequest) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *AddPluginRequest) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *AddPluginRequest) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *AddPluginRequest) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *AddPluginRequest) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *AddPluginRequest) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *AddPluginRequest) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetPassThroughAuthenticationHandler

`func (o *AddPluginRequest) GetPassThroughAuthenticationHandler() string`

GetPassThroughAuthenticationHandler returns the PassThroughAuthenticationHandler field if non-nil, zero value otherwise.

### GetPassThroughAuthenticationHandlerOk

`func (o *AddPluginRequest) GetPassThroughAuthenticationHandlerOk() (*string, bool)`

GetPassThroughAuthenticationHandlerOk returns a tuple with the PassThroughAuthenticationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThroughAuthenticationHandler

`func (o *AddPluginRequest) SetPassThroughAuthenticationHandler(v string)`

SetPassThroughAuthenticationHandler sets PassThroughAuthenticationHandler field to given value.


### GetUpdateInterval

`func (o *AddPluginRequest) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *AddPluginRequest) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *AddPluginRequest) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *AddPluginRequest) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetType

`func (o *AddPluginRequest) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddPluginRequest) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddPluginRequest) SetType(v []string)`

SetType sets Type field to given value.


### GetMultipleAttributeBehavior

`func (o *AddPluginRequest) GetMultipleAttributeBehavior() EnumpluginMultipleAttributeBehaviorProp`

GetMultipleAttributeBehavior returns the MultipleAttributeBehavior field if non-nil, zero value otherwise.

### GetMultipleAttributeBehaviorOk

`func (o *AddPluginRequest) GetMultipleAttributeBehaviorOk() (*EnumpluginMultipleAttributeBehaviorProp, bool)`

GetMultipleAttributeBehaviorOk returns a tuple with the MultipleAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAttributeBehavior

`func (o *AddPluginRequest) SetMultipleAttributeBehavior(v EnumpluginMultipleAttributeBehaviorProp)`

SetMultipleAttributeBehavior sets MultipleAttributeBehavior field to given value.

### HasMultipleAttributeBehavior

`func (o *AddPluginRequest) HasMultipleAttributeBehavior() bool`

HasMultipleAttributeBehavior returns a boolean if a field has been set.

### GetPreventConflictsWithSoftDeletedEntries

`func (o *AddPluginRequest) GetPreventConflictsWithSoftDeletedEntries() bool`

GetPreventConflictsWithSoftDeletedEntries returns the PreventConflictsWithSoftDeletedEntries field if non-nil, zero value otherwise.

### GetPreventConflictsWithSoftDeletedEntriesOk

`func (o *AddPluginRequest) GetPreventConflictsWithSoftDeletedEntriesOk() (*bool, bool)`

GetPreventConflictsWithSoftDeletedEntriesOk returns a tuple with the PreventConflictsWithSoftDeletedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventConflictsWithSoftDeletedEntries

`func (o *AddPluginRequest) SetPreventConflictsWithSoftDeletedEntries(v bool)`

SetPreventConflictsWithSoftDeletedEntries sets PreventConflictsWithSoftDeletedEntries field to given value.

### HasPreventConflictsWithSoftDeletedEntries

`func (o *AddPluginRequest) HasPreventConflictsWithSoftDeletedEntries() bool`

HasPreventConflictsWithSoftDeletedEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


