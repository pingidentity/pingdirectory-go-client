# GetPlugin200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | [**[]EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn**](EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn.md) |  | 
**Id** | **string** | Name of the Plugin | 
**MaxUpdateFrequency** | Pointer to **string** | Specifies the maximum frequency with which last access time values should be written for an entry. This may help limit the rate of internal write operations processed in the server. | [optional] 
**OperationType** | Pointer to [**[]EnumpluginOperationTypeProp**](EnumpluginOperationTypeProp.md) |  | [optional] 
**InvokeForFailedBinds** | Pointer to **bool** | Indicates whether to update the last access time for an entry targeted by a bind operation if the bind is unsuccessful. | [optional] 
**MaxSearchResultEntriesToUpdate** | Pointer to **int64** | Specifies the maximum number of entries that should be updated in a search operation. Only search result entries actually returned to the client may have their last access time updated, but because a single search operation may return a very large number of entries, the plugin will only update entries if no more than a specified number of entries are updated. | [optional] 
**RequestCriteria** | **string** | A reference to request criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | 
**InvokeForInternalOperations** | Pointer to **bool** | Indicates whether the plug-in should be invoked for internal operations. | [optional] 
**Description** | Pointer to **string** | A description for this Plugin | [optional] 
**Enabled** | **bool** | Indicates whether the plug-in is enabled for use. | 
**Meta** | Pointer to [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Urnpingidentityschemasconfigurationmessages20** | Pointer to [**MetaUrnPingidentitySchemasConfigurationMessages20**](MetaUrnPingidentitySchemasConfigurationMessages20.md) |  | [optional] 
**SampleInterval** | **string** | The duration between statistics collections. Setting this value too small can have an impact on performance. This value should be a multiple of collection-interval. | 
**CollectionInterval** | **string** | Some of the calculated statistics, such as the average and maximum queue sizes, can use multiple samples within a log interval. This value controls how often samples are gathered. It should be a multiple of the log-interval. | 
**LdapInfo** | Pointer to [**EnumpluginLdapInfoProp**](EnumpluginLdapInfoProp.md) |  | [optional] 
**ServerInfo** | Pointer to [**EnumpluginServerInfoProp**](EnumpluginServerInfoProp.md) |  | [optional] 
**PerApplicationLDAPStats** | Pointer to [**EnumpluginPeriodicStatsLoggerPerApplicationLDAPStatsProp**](EnumpluginPeriodicStatsLoggerPerApplicationLDAPStatsProp.md) |  | [optional] 
**LdapChangelogInfo** | Pointer to [**EnumpluginLdapChangelogInfoProp**](EnumpluginLdapChangelogInfoProp.md) |  | [optional] 
**StatusSummaryInfo** | Pointer to [**EnumpluginStatusSummaryInfoProp**](EnumpluginStatusSummaryInfoProp.md) |  | [optional] 
**GenerateCollectorFiles** | Pointer to **bool** | Indicates whether this plugin should store metric samples on disk for use by the Data Metrics Server. If the Stats Collector Plugin is only being used to collect metrics for one or more StatsD Monitoring Endpoints, then this can be set to false to prevent unnecessary I/O. | [optional] 
**LocalDBBackendInfo** | Pointer to [**EnumpluginLocalDBBackendInfoProp**](EnumpluginLocalDBBackendInfoProp.md) |  | [optional] 
**ReplicationInfo** | Pointer to [**EnumpluginReplicationInfoProp**](EnumpluginReplicationInfoProp.md) |  | [optional] 
**EntryCacheInfo** | Pointer to [**EnumpluginEntryCacheInfoProp**](EnumpluginEntryCacheInfoProp.md) |  | [optional] 
**HostInfo** | Pointer to [**[]EnumpluginHostInfoProp**](EnumpluginHostInfoProp.md) |  | [optional] 
**IncludedLDAPApplication** | Pointer to **[]string** | If statistics should not be included for all applications, this property names the subset of applications that should be included. | [optional] 
**TraditionalStaticGroupObjectClass** | Pointer to [**EnumpluginTraditionalStaticGroupObjectClassProp**](EnumpluginTraditionalStaticGroupObjectClassProp.md) |  | [optional] 
**MaximumMembershipUpdatesPerModify** | Pointer to **int64** | An integer property that specifies the maximum number of membership changes that will be supported in a single modify operation. A value of zero indicates that modify operations targeting the group entry should not be permitted to alter the set of members for the group. | [optional] 
**ReadOperationSupport** | Pointer to [**EnumpluginReadOperationSupportProp**](EnumpluginReadOperationSupportProp.md) |  | [optional] 
**PluginType** | [**[]EnumpluginPluginTypeProp**](EnumpluginPluginTypeProp.md) |  | 
**NumThreads** | **int64** | Specifies the number of concurrent threads that should be used to process the search operations. | 
**BaseDN** | **[]string** | Specifies a base DN within which the attribute must be unique. | 
**LowerBound** | Pointer to **int64** | Specifies the lower bound for the numeric value which will be inserted into the search filter. | [optional] 
**UpperBound** | Pointer to **int64** | Specifies the upper bound for the numeric value which will be inserted into the search filter. | [optional] 
**FilterPrefix** | **string** | Specifies a prefix which will be used in front of the randomly-selected numeric value in all search filters used. If no upper bound is defined, then this should contain the entire filter string. | 
**FilterSuffix** | Pointer to **string** | Specifies a suffix which will be used after of the randomly-selected numeric value in all search filters used. If no upper bound is defined, then this should be omitted. | [optional] 
**Filter** | **string** | Specifies the search filter to apply to determine if attribute uniqueness is enforced for the matching entries. | 
**AttributeType** | **[]string** | Specifies the attribute types for which referential integrity is to be maintained. | 
**PollingInterval** | **string** | This specifies how often the plugin should check for expired data. It also controls the offset of peer servers (see the peer-server-priority-index for more information). | 
**PeerServerPriorityIndex** | Pointer to **int64** | In a replicated environment, this determines the order in which peer servers should attempt to purge data. | [optional] 
**MaxUpdatesPerSecond** | **int64** | This setting smooths out the performance impact on the server by throttling the purging to the specified maximum number of updates per second. To avoid a large backlog, this value should be set comfortably above the average rate that expired data is generated. When purge-behavior is set to subtree-delete-entries, then deletion of the entire subtree is considered a single update for the purposes of throttling. | 
**NumDeleteThreads** | **int64** | The number of threads used to delete expired entries. | 
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
**HttpProxyExternalServer** | Pointer to **string** | A reference to an HTTP proxy server that should be used for requests sent to the PingOne service. | [optional] 
**IncludedLocalEntryBaseDN** | Pointer to **[]string** | The base DNs for the local users whose authentication attempts may be passed through to the external authentication service. | [optional] 
**ConnectionCriteria** | Pointer to **string** | A reference to connection criteria that will be used to indicate which bind requests should be passed through to the external authentication service. | [optional] 
**TryLocalBind** | **bool** | Indicates whether to attempt the bind in the local server first and only send the request to the external authentication service if the local bind attempt fails, or to only attempt the bind in the external service. | 
**OverrideLocalPassword** | **bool** | Indicates whether to attempt the authentication in the external service if the local user entry includes a password. This property will be ignored if try-local-bind is false. | 
**UpdateLocalPassword** | **bool** | Indicates whether to overwrite the user&#39;s local password if the local bind fails but the authentication attempt succeeds when attempted in the external service. This property may only be set to true if try-local-bind is also true. | 
**UpdateLocalPasswordDN** | Pointer to **string** | The DN of the authorization identity that will be used when updating the user&#39;s local password if update-local-password is true. This is primarily intended for use if the Data Sync Server will be used to synchronize passwords between the local server and the external service, and in that case, the DN used here should also be added to the ignore-changes-by-dn property in the appropriate Sync Source object in the Data Sync Server configuration. | [optional] 
**AllowLaxPassThroughAuthenticationPasswords** | Pointer to **bool** | Indicates whether to overwrite the user&#39;s local password even if the password used to authenticate to the external service would have failed validation if the user attempted to set it directly. | [optional] 
**IgnoredPasswordPolicyStateErrorCondition** | Pointer to [**[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp**](EnumpluginIgnoredPasswordPolicyStateErrorConditionProp.md) |  | [optional] 
**UserMappingLocalAttribute** | **[]string** | The names of the attributes in the local user entry whose values must match the values of the corresponding fields in the PingOne service. | 
**UserMappingRemoteJSONField** | **[]string** | The names of the fields in the PingOne service whose values must match the values of the corresponding attributes in the local user entry, as specified in the user-mapping-local-attribute property. | 
**AdditionalUserMappingSCIMFilter** | Pointer to **string** | An optional SCIM filter that will be ANDed with the filter created to identify the account in the PingOne service that corresponds to the local entry. Only the \&quot;eq\&quot;, \&quot;sw\&quot;, \&quot;and\&quot;, and \&quot;or\&quot; filter types may be used. | [optional] 
**ChangelogPasswordEncryptionKey** | Pointer to **string** | A passphrase that may be used to generate the key for encrypting passwords stored in the changelog. The same passphrase also needs to be set (either through the \&quot;changelog-password-decryption-key\&quot; property or the \&quot;changelog-password-decryption-key-passphrase-provider\&quot; property) in the Global Sync Configuration in the Data Sync Server. | [optional] 
**ChangelogPasswordEncryptionKeyPassphraseProvider** | Pointer to **string** | A passphrase provider that may be used to obtain the passphrase that will be used to generate the key for encrypting passwords stored in the changelog. The same passphrase also needs to be set (either through the \&quot;changelog-password-decryption-key\&quot; property or the \&quot;changelog-password-decryption-key-passphrase-provider\&quot; property) in the Global Sync Configuration in the Data Sync Server. | [optional] 
**HistogramCategoryBoundary** | **[]string** | Specifies the boundary values that will be used to separate the processing times into categories. Values should be specified as durations, and all values must be greater than zero. | 
**IncludeQueueTime** | Pointer to **bool** | Indicates whether operation processing times should include the time spent waiting on the work queue. This will only be available if the work queue is configured to monitor the queue time. | [optional] 
**SeparateMonitorEntryPerTrackedApplication** | Pointer to **bool** | When enabled, separate monitor entries will be included for each application defined in the Global Configuration&#39;s tracked-application property. | [optional] 
**Scope** | [**EnumpluginScopeProp**](EnumpluginScopeProp.md) |  | 
**IncludeAttribute** | Pointer to **[]string** | Specifies the name or OID of an attribute type that must be updated in order for the modifiersName and modifyTimestamp attributes to be updated in the target entry. | [optional] 
**OutputFile** | **string** | The path of an LDIF file that should be created with the results of the search. | 
**PreviousFileExtension** | Pointer to **string** | An extension that should be appended to the name of an existing output file rather than deleting it. If a file already exists with the full previous file name, then it will be deleted before the current file is renamed to become the previous file. | [optional] 
**LogInterval** | **string** | The duration between logging dumps of cn&#x3D;monitor to a file. | 
**SuppressIfIdle** | **bool** | If the server is idle during the specified interval, then do not log any output if this property is set to true. The server is idle if during the interval, no new connections were established, no operations were processed, and no operations are pending. | 
**HeaderPrefixPerColumn** | Pointer to **bool** | This property controls whether the header prefix, which applies to a group of columns, appears at the start of each column header or only the first column in a group. | [optional] 
**EmptyInsteadOfZero** | Pointer to **bool** | This property controls whether a value in the output is shown as empty if the value is zero. | [optional] 
**LinesBetweenHeader** | **int64** | The number of lines to log between logging the header line that summarizes the columns in the table. | 
**IncludedLDAPStat** | Pointer to [**[]EnumpluginIncludedLDAPStatProp**](EnumpluginIncludedLDAPStatProp.md) |  | [optional] 
**IncludedResourceStat** | Pointer to [**[]EnumpluginIncludedResourceStatProp**](EnumpluginIncludedResourceStatProp.md) |  | [optional] 
**HistogramFormat** | [**EnumpluginHistogramFormatProp**](EnumpluginHistogramFormatProp.md) |  | 
**HistogramOpType** | Pointer to [**[]EnumpluginHistogramOpTypeProp**](EnumpluginHistogramOpTypeProp.md) |  | [optional] 
**GaugeInfo** | Pointer to [**EnumpluginGaugeInfoProp**](EnumpluginGaugeInfoProp.md) |  | [optional] 
**LogFileFormat** | Pointer to [**EnumpluginLogFileFormatProp**](EnumpluginLogFileFormatProp.md) |  | [optional] 
**LogFile** | **string** | Specifies the log file location where the update records are written when the plug-in is in background-mode processing. | 
**LogFilePermissions** | **string** | The UNIX permissions of the log files created by this Monitor History Plugin. | 
**Append** | Pointer to **bool** | Specifies whether to append to existing log files. | [optional] 
**RotationPolicy** | **[]string** | The rotation policy to use for the Periodic Stats Logger Plugin . | 
**RotationListener** | Pointer to **[]string** | A listener that should be notified whenever a log file is rotated out of service. | [optional] 
**RetentionPolicy** | **[]string** | The retention policy to use for the Monitor History Plugin . | 
**LoggingErrorBehavior** | Pointer to [**EnumpluginLoggingErrorBehaviorProp**](EnumpluginLoggingErrorBehaviorProp.md) |  | [optional] 
**DatetimeAttribute** | **string** | The LDAP attribute that determines when data should be deleted. This could store the expiration time, or it could store the creation time and the expiration-offset property specifies the duration before data is deleted. | 
**DatetimeJSONField** | Pointer to **string** | The top-level JSON field within the configured datetime-attribute that determines when data should be deleted. This could store the expiration time, or it could store the creation time and the expiration-offset property specifies the duration before data is deleted. | [optional] 
**DatetimeFormat** | [**EnumpluginDatetimeFormatProp**](EnumpluginDatetimeFormatProp.md) |  | 
**CustomDatetimeFormat** | Pointer to **string** | When the datetime-format property is configured with a value of \&quot;custom\&quot;, this specifies the format (using a string compatible with the java.text.SimpleDateFormat class) that will be used to search for expired data. | [optional] 
**CustomTimezone** | Pointer to **string** | Specifies the time zone to use when generating a date string using the configured custom-datetime-format value. The provided value must be accepted by java.util.TimeZone.getTimeZone. | [optional] 
**ExpirationOffset** | **string** | Sessions whose last activity timestamp is older than this offset will be removed. | 
**PurgeBehavior** | Pointer to [**EnumpluginPurgeBehaviorProp**](EnumpluginPurgeBehaviorProp.md) |  | [optional] 
**NumMostExpensivePhasesShown** | **int64** | This controls how many of the most expensive phases are included per operation type in the monitor entry. | 
**ExtensionClass** | **string** | The fully-qualified name of the Java class providing the logic for the Third Party Plugin. | 
**ExtensionArgument** | Pointer to **[]string** | The set of arguments used to customize the behavior for the Third Party Plugin. Each configuration property should be given in the form &#39;name&#x3D;value&#39;. | [optional] 
**EncryptionSettingsDefinitionID** | Pointer to **string** | Specifies the ID of the encryption settings definition that should be used to encrypt the data. If this is not provided, the server&#39;s preferred encryption settings definition will be used. The \&quot;encryption-settings list\&quot; command can be used to obtain a list of the encryption settings definitions available in the server. | [optional] 
**Server** | **[]string** | Specifies the LDAP external server(s) to which authentication attempts should be forwarded. | 
**ServerAccessMode** | [**EnumpluginServerAccessModeProp**](EnumpluginServerAccessModeProp.md) |  | 
**DnMap** | Pointer to **[]string** | Specifies one or more DN mappings that may be used to transform bind DNs before attempting to bind to the external servers. | [optional] 
**BindDNPattern** | Pointer to **string** | A pattern to use to construct the bind DN for the simple bind request to send to the remote server. This may consist of a combination of static text and attribute values and other directives enclosed in curly braces.  For example, the value \&quot;cn&#x3D;{cn},ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot; indicates that the remote bind DN should be constructed from the text \&quot;cn&#x3D;\&quot; followed by the value of the local entry&#39;s cn attribute followed by the text \&quot;ou&#x3D;People,dc&#x3D;example,dc&#x3D;com\&quot;. If an attribute contains the value to use as the bind DN for pass-through authentication, then the pattern may simply be the name of that attribute in curly braces (e.g., if the seeAlso attribute contains the bind DN for the target user, then a bind DN pattern of \&quot;{seeAlso}\&quot; would be appropriate).  Note that a bind DN pattern can be used to construct a bind DN that is not actually a valid LDAP distinguished name. For example, if authentication is being passed through to a Microsoft Active Directory server, then a bind DN pattern could be used to construct a user principal name (UPN) as an alternative to a distinguished name. | [optional] 
**SearchBaseDN** | Pointer to **string** | The base DN to use when searching for the user entry using a filter constructed from the pattern defined in the search-filter-pattern property. If no base DN is specified, the null DN will be used as the search base DN. | [optional] 
**SearchFilterPattern** | Pointer to **string** | A pattern to use to construct a filter to use when searching an external server for the entry of the user as whom to bind. For example, \&quot;(mail&#x3D;{uid:ldapFilterEscape}@example.com)\&quot; would construct a search filter to search for a user whose entry in the local server contains a uid attribute whose value appears before \&quot;@example.com\&quot; in the mail attribute in the external server. Note that the \&quot;ldapFilterEscape\&quot; modifier should almost always be used with attributes specified in the pattern. | [optional] 
**InitialConnections** | **int64** | Specifies the initial number of connections to establish to each external server against which authentication may be attempted. | 
**MaxConnections** | **int64** | Specifies the maximum number of connections to maintain to each external server against which authentication may be attempted. This value must be greater than or equal to the value for the initial-connections property. | 
**SourceDN** | **string** | Specifies the source DN that may appear in client requests which should be remapped to the target DN. Note that the source DN must not be equal to the target DN. | 
**TargetDN** | **string** | Specifies the DN to which the source DN should be mapped. Note that the target DN must not be equal to the source DN. | 
**EnableAttributeMapping** | **bool** | Indicates whether DN mapping should be applied to the values of attributes with appropriate syntaxes. | 
**MapAttribute** | Pointer to **[]string** | Specifies a set of specific attributes for which DN mapping should be applied. This will only be applicable if the enable-attribute-mapping property has a value of \&quot;true\&quot;. Any attributes listed must be defined in the server schema with either the distinguished name syntax or the name and optional UID syntax. | [optional] 
**EnableControlMapping** | **bool** | Indicates whether mapping should be applied to attribute types that may be present in specific controls. If enabled, attribute mapping will only be applied for control types which are specifically supported by the attribute mapper plugin. | 
**AlwaysMapResponses** | **bool** | Indicates whether the target attribute in response messages should always be remapped back to the source attribute. If this is \&quot;false\&quot;, then the mapping will be performed for a response message only if one or more elements of the associated request are mapped. Otherwise, the mapping will be performed for all responses regardless of whether the mapping was applied to the request. | 
**RetainFilesSparselyByAge** | Pointer to **bool** | Retain some older files to give greater perspective on how monitoring information has changed over time. | [optional] 
**Sanitize** | Pointer to **bool** | Server monitoring data can include a small amount of personally identifiable information in the form of LDAP DNs and search filters. Setting this property to true will redact this information from the monitor files. This should only be used when necessary, as it reduces the information available in the archive and can increase the time to find the source of support issues. | [optional] 
**ReferralBaseURL** | **[]string** | Specifies the base URL to use for the referrals generated by this plugin. It should include only the scheme, address, and port to use to communicate with the target server (e.g., \&quot;ldap://server.example.com:389/\&quot;). | 
**ContextName** | Pointer to **string** | The SNMP context name for this sub-agent. The context name must not be longer than 30 ASCII characters. Each server in a topology must have a unique SNMP context name. | [optional] 
**AgentxAddress** | **string** | The hostname or IP address of the SNMP master agent. | 
**AgentxPort** | **int64** | The port number on which the SNMP master agent will be contacted. | 
**NumWorkerThreads** | Pointer to **int64** | The number of worker threads to use to handle SNMP requests. | [optional] 
**SessionTimeout** | Pointer to **string** | Specifies the maximum amount of time to wait for a session to the master agent to be established. | [optional] 
**ConnectRetryMaxWait** | Pointer to **string** | The maximum amount of time to wait between attempts to establish a connection to the master agent. | [optional] 
**PingInterval** | Pointer to **string** | The amount of time between consecutive pings sent by the sub-agent on its connection to the master agent. A value of zero disables the sending of pings by the sub-agent. | [optional] 
**AllowedRequestControl** | Pointer to **[]string** | Specifies the OIDs of the controls that are allowed to be present in operations to coalesce. These controls are passed through when the request is validated, but they will not be included when the background thread applies the coalesced modify requests. | [optional] 
**DefaultUserPasswordStorageScheme** | Pointer to **[]string** | Specifies the names of the password storage schemes to be used for encoding passwords contained in attributes with the user password syntax for entries that do not include the ds-pwp-password-policy-dn attribute specifying which password policy is to be used to govern them. | [optional] 
**DefaultAuthPasswordStorageScheme** | Pointer to **[]string** | Specifies the names of password storage schemes that to be used for encoding passwords contained in attributes with the auth password syntax for entries that do not include the ds-pwp-password-policy-dn attribute specifying which password policy should be used to govern them. | [optional] 
**ProfileSampleInterval** | **string** | Specifies the sample interval in milliseconds to be used when capturing profiling information in the server. | 
**ProfileDirectory** | **string** | Specifies the path to the directory where profile information is to be written. This path may be either an absolute path or a path that is relative to the root of the Directory Server instance. | 
**EnableProfilingOnStartup** | **bool** | Indicates whether the profiler plug-in is to start collecting data automatically when the Directory Server is started. | 
**ProfileAction** | Pointer to [**EnumpluginProfileActionProp**](EnumpluginProfileActionProp.md) |  | [optional] 
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
**ExcludeAttribute** | Pointer to **[]string** | Specifies the name or OID of an attribute type which may be updated in a modify or modify DN operation without causing the modifiersName and modifyTimestamp values to be updated for that entry. | [optional] 
**PassThroughAuthenticationHandler** | **string** | The component used to manage authentication with the external authentication service. | 
**UpdateInterval** | Pointer to **string** | Specifies the interval in seconds when referential integrity updates are made. | [optional] 
**Type** | **[]string** | Specifies the type of attributes to check for value uniqueness. | 
**MultipleAttributeBehavior** | Pointer to [**EnumpluginUniqueAttributeMultipleAttributeBehaviorProp**](EnumpluginUniqueAttributeMultipleAttributeBehaviorProp.md) |  | [optional] 
**PreventConflictsWithSoftDeletedEntries** | Pointer to **bool** | Indicates whether this Unique Attribute Plugin should reject a change that would result in one or more conflicts, even if those conflicts only exist in soft-deleted entries. | [optional] 
**PreventAddingMembersToNonexistentGroups** | Pointer to **bool** | Indicates whether the server should prevent updates to user entries that attempt to add them as a member of an inverted static group that does not exist. | [optional] 
**PreventAddingGroupsAsInvertedStaticGroupMembers** | Pointer to **bool** | Indicates whether the server should prevent attempts to add a group as a regular member of an inverted static group. If the members of another group should be considered members of an inverted static group, then the other group should be added as a nested group rather than a regular member. | [optional] 
**PreventNestingNonexistentGroups** | Pointer to **bool** | Indicates whether the server should prevent updates to inverted static groups that add references to nested groups that don&#39;t exist. | [optional] 

## Methods

### NewGetPlugin200Response

`func NewGetPlugin200Response(schemas []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn, id string, requestCriteria string, enabled bool, sampleInterval string, collectionInterval string, pluginType []EnumpluginPluginTypeProp, numThreads int64, baseDN []string, filterPrefix string, filter string, attributeType []string, pollingInterval string, maxUpdatesPerSecond int64, numDeleteThreads int64, invokeGCTimeUtc []string, apiURL string, authURL string, oAuthClientID string, environmentID string, tryLocalBind bool, overrideLocalPassword bool, updateLocalPassword bool, userMappingLocalAttribute []string, userMappingRemoteJSONField []string, histogramCategoryBoundary []string, scope EnumpluginScopeProp, outputFile string, logInterval string, suppressIfIdle bool, linesBetweenHeader int64, histogramFormat EnumpluginHistogramFormatProp, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, datetimeAttribute string, datetimeFormat EnumpluginDatetimeFormatProp, expirationOffset string, numMostExpensivePhasesShown int64, extensionClass string, server []string, serverAccessMode EnumpluginServerAccessModeProp, initialConnections int64, maxConnections int64, sourceDN string, targetDN string, enableAttributeMapping bool, enableControlMapping bool, alwaysMapResponses bool, referralBaseURL []string, agentxAddress string, agentxPort int64, profileSampleInterval string, profileDirectory string, enableProfilingOnStartup bool, valuePattern []string, sourceAttribute string, targetAttribute string, delay string, scriptClass string, passThroughAuthenticationHandler string, type_ []string, ) *GetPlugin200Response`

NewGetPlugin200Response instantiates a new GetPlugin200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlugin200ResponseWithDefaults

`func NewGetPlugin200ResponseWithDefaults() *GetPlugin200Response`

NewGetPlugin200ResponseWithDefaults instantiates a new GetPlugin200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *GetPlugin200Response) GetSchemas() []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *GetPlugin200Response) GetSchemasOk() (*[]EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *GetPlugin200Response) SetSchemas(v []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *GetPlugin200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPlugin200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPlugin200Response) SetId(v string)`

SetId sets Id field to given value.


### GetMaxUpdateFrequency

`func (o *GetPlugin200Response) GetMaxUpdateFrequency() string`

GetMaxUpdateFrequency returns the MaxUpdateFrequency field if non-nil, zero value otherwise.

### GetMaxUpdateFrequencyOk

`func (o *GetPlugin200Response) GetMaxUpdateFrequencyOk() (*string, bool)`

GetMaxUpdateFrequencyOk returns a tuple with the MaxUpdateFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdateFrequency

`func (o *GetPlugin200Response) SetMaxUpdateFrequency(v string)`

SetMaxUpdateFrequency sets MaxUpdateFrequency field to given value.

### HasMaxUpdateFrequency

`func (o *GetPlugin200Response) HasMaxUpdateFrequency() bool`

HasMaxUpdateFrequency returns a boolean if a field has been set.

### GetOperationType

`func (o *GetPlugin200Response) GetOperationType() []EnumpluginOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *GetPlugin200Response) GetOperationTypeOk() (*[]EnumpluginOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *GetPlugin200Response) SetOperationType(v []EnumpluginOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *GetPlugin200Response) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetInvokeForFailedBinds

`func (o *GetPlugin200Response) GetInvokeForFailedBinds() bool`

GetInvokeForFailedBinds returns the InvokeForFailedBinds field if non-nil, zero value otherwise.

### GetInvokeForFailedBindsOk

`func (o *GetPlugin200Response) GetInvokeForFailedBindsOk() (*bool, bool)`

GetInvokeForFailedBindsOk returns a tuple with the InvokeForFailedBinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForFailedBinds

`func (o *GetPlugin200Response) SetInvokeForFailedBinds(v bool)`

SetInvokeForFailedBinds sets InvokeForFailedBinds field to given value.

### HasInvokeForFailedBinds

`func (o *GetPlugin200Response) HasInvokeForFailedBinds() bool`

HasInvokeForFailedBinds returns a boolean if a field has been set.

### GetMaxSearchResultEntriesToUpdate

`func (o *GetPlugin200Response) GetMaxSearchResultEntriesToUpdate() int64`

GetMaxSearchResultEntriesToUpdate returns the MaxSearchResultEntriesToUpdate field if non-nil, zero value otherwise.

### GetMaxSearchResultEntriesToUpdateOk

`func (o *GetPlugin200Response) GetMaxSearchResultEntriesToUpdateOk() (*int64, bool)`

GetMaxSearchResultEntriesToUpdateOk returns a tuple with the MaxSearchResultEntriesToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSearchResultEntriesToUpdate

`func (o *GetPlugin200Response) SetMaxSearchResultEntriesToUpdate(v int64)`

SetMaxSearchResultEntriesToUpdate sets MaxSearchResultEntriesToUpdate field to given value.

### HasMaxSearchResultEntriesToUpdate

`func (o *GetPlugin200Response) HasMaxSearchResultEntriesToUpdate() bool`

HasMaxSearchResultEntriesToUpdate returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *GetPlugin200Response) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *GetPlugin200Response) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *GetPlugin200Response) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.


### GetInvokeForInternalOperations

`func (o *GetPlugin200Response) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *GetPlugin200Response) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *GetPlugin200Response) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *GetPlugin200Response) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetDescription

`func (o *GetPlugin200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPlugin200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPlugin200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPlugin200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetPlugin200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetPlugin200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetPlugin200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *GetPlugin200Response) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetPlugin200Response) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetPlugin200Response) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetPlugin200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *GetPlugin200Response) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *GetPlugin200Response) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *GetPlugin200Response) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *GetPlugin200Response) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSampleInterval

`func (o *GetPlugin200Response) GetSampleInterval() string`

GetSampleInterval returns the SampleInterval field if non-nil, zero value otherwise.

### GetSampleIntervalOk

`func (o *GetPlugin200Response) GetSampleIntervalOk() (*string, bool)`

GetSampleIntervalOk returns a tuple with the SampleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleInterval

`func (o *GetPlugin200Response) SetSampleInterval(v string)`

SetSampleInterval sets SampleInterval field to given value.


### GetCollectionInterval

`func (o *GetPlugin200Response) GetCollectionInterval() string`

GetCollectionInterval returns the CollectionInterval field if non-nil, zero value otherwise.

### GetCollectionIntervalOk

`func (o *GetPlugin200Response) GetCollectionIntervalOk() (*string, bool)`

GetCollectionIntervalOk returns a tuple with the CollectionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionInterval

`func (o *GetPlugin200Response) SetCollectionInterval(v string)`

SetCollectionInterval sets CollectionInterval field to given value.


### GetLdapInfo

`func (o *GetPlugin200Response) GetLdapInfo() EnumpluginLdapInfoProp`

GetLdapInfo returns the LdapInfo field if non-nil, zero value otherwise.

### GetLdapInfoOk

`func (o *GetPlugin200Response) GetLdapInfoOk() (*EnumpluginLdapInfoProp, bool)`

GetLdapInfoOk returns a tuple with the LdapInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapInfo

`func (o *GetPlugin200Response) SetLdapInfo(v EnumpluginLdapInfoProp)`

SetLdapInfo sets LdapInfo field to given value.

### HasLdapInfo

`func (o *GetPlugin200Response) HasLdapInfo() bool`

HasLdapInfo returns a boolean if a field has been set.

### GetServerInfo

`func (o *GetPlugin200Response) GetServerInfo() EnumpluginServerInfoProp`

GetServerInfo returns the ServerInfo field if non-nil, zero value otherwise.

### GetServerInfoOk

`func (o *GetPlugin200Response) GetServerInfoOk() (*EnumpluginServerInfoProp, bool)`

GetServerInfoOk returns a tuple with the ServerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInfo

`func (o *GetPlugin200Response) SetServerInfo(v EnumpluginServerInfoProp)`

SetServerInfo sets ServerInfo field to given value.

### HasServerInfo

`func (o *GetPlugin200Response) HasServerInfo() bool`

HasServerInfo returns a boolean if a field has been set.

### GetPerApplicationLDAPStats

`func (o *GetPlugin200Response) GetPerApplicationLDAPStats() EnumpluginPeriodicStatsLoggerPerApplicationLDAPStatsProp`

GetPerApplicationLDAPStats returns the PerApplicationLDAPStats field if non-nil, zero value otherwise.

### GetPerApplicationLDAPStatsOk

`func (o *GetPlugin200Response) GetPerApplicationLDAPStatsOk() (*EnumpluginPeriodicStatsLoggerPerApplicationLDAPStatsProp, bool)`

GetPerApplicationLDAPStatsOk returns a tuple with the PerApplicationLDAPStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerApplicationLDAPStats

`func (o *GetPlugin200Response) SetPerApplicationLDAPStats(v EnumpluginPeriodicStatsLoggerPerApplicationLDAPStatsProp)`

SetPerApplicationLDAPStats sets PerApplicationLDAPStats field to given value.

### HasPerApplicationLDAPStats

`func (o *GetPlugin200Response) HasPerApplicationLDAPStats() bool`

HasPerApplicationLDAPStats returns a boolean if a field has been set.

### GetLdapChangelogInfo

`func (o *GetPlugin200Response) GetLdapChangelogInfo() EnumpluginLdapChangelogInfoProp`

GetLdapChangelogInfo returns the LdapChangelogInfo field if non-nil, zero value otherwise.

### GetLdapChangelogInfoOk

`func (o *GetPlugin200Response) GetLdapChangelogInfoOk() (*EnumpluginLdapChangelogInfoProp, bool)`

GetLdapChangelogInfoOk returns a tuple with the LdapChangelogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapChangelogInfo

`func (o *GetPlugin200Response) SetLdapChangelogInfo(v EnumpluginLdapChangelogInfoProp)`

SetLdapChangelogInfo sets LdapChangelogInfo field to given value.

### HasLdapChangelogInfo

`func (o *GetPlugin200Response) HasLdapChangelogInfo() bool`

HasLdapChangelogInfo returns a boolean if a field has been set.

### GetStatusSummaryInfo

`func (o *GetPlugin200Response) GetStatusSummaryInfo() EnumpluginStatusSummaryInfoProp`

GetStatusSummaryInfo returns the StatusSummaryInfo field if non-nil, zero value otherwise.

### GetStatusSummaryInfoOk

`func (o *GetPlugin200Response) GetStatusSummaryInfoOk() (*EnumpluginStatusSummaryInfoProp, bool)`

GetStatusSummaryInfoOk returns a tuple with the StatusSummaryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSummaryInfo

`func (o *GetPlugin200Response) SetStatusSummaryInfo(v EnumpluginStatusSummaryInfoProp)`

SetStatusSummaryInfo sets StatusSummaryInfo field to given value.

### HasStatusSummaryInfo

`func (o *GetPlugin200Response) HasStatusSummaryInfo() bool`

HasStatusSummaryInfo returns a boolean if a field has been set.

### GetGenerateCollectorFiles

`func (o *GetPlugin200Response) GetGenerateCollectorFiles() bool`

GetGenerateCollectorFiles returns the GenerateCollectorFiles field if non-nil, zero value otherwise.

### GetGenerateCollectorFilesOk

`func (o *GetPlugin200Response) GetGenerateCollectorFilesOk() (*bool, bool)`

GetGenerateCollectorFilesOk returns a tuple with the GenerateCollectorFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCollectorFiles

`func (o *GetPlugin200Response) SetGenerateCollectorFiles(v bool)`

SetGenerateCollectorFiles sets GenerateCollectorFiles field to given value.

### HasGenerateCollectorFiles

`func (o *GetPlugin200Response) HasGenerateCollectorFiles() bool`

HasGenerateCollectorFiles returns a boolean if a field has been set.

### GetLocalDBBackendInfo

`func (o *GetPlugin200Response) GetLocalDBBackendInfo() EnumpluginLocalDBBackendInfoProp`

GetLocalDBBackendInfo returns the LocalDBBackendInfo field if non-nil, zero value otherwise.

### GetLocalDBBackendInfoOk

`func (o *GetPlugin200Response) GetLocalDBBackendInfoOk() (*EnumpluginLocalDBBackendInfoProp, bool)`

GetLocalDBBackendInfoOk returns a tuple with the LocalDBBackendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDBBackendInfo

`func (o *GetPlugin200Response) SetLocalDBBackendInfo(v EnumpluginLocalDBBackendInfoProp)`

SetLocalDBBackendInfo sets LocalDBBackendInfo field to given value.

### HasLocalDBBackendInfo

`func (o *GetPlugin200Response) HasLocalDBBackendInfo() bool`

HasLocalDBBackendInfo returns a boolean if a field has been set.

### GetReplicationInfo

`func (o *GetPlugin200Response) GetReplicationInfo() EnumpluginReplicationInfoProp`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *GetPlugin200Response) GetReplicationInfoOk() (*EnumpluginReplicationInfoProp, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *GetPlugin200Response) SetReplicationInfo(v EnumpluginReplicationInfoProp)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *GetPlugin200Response) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetEntryCacheInfo

`func (o *GetPlugin200Response) GetEntryCacheInfo() EnumpluginEntryCacheInfoProp`

GetEntryCacheInfo returns the EntryCacheInfo field if non-nil, zero value otherwise.

### GetEntryCacheInfoOk

`func (o *GetPlugin200Response) GetEntryCacheInfoOk() (*EnumpluginEntryCacheInfoProp, bool)`

GetEntryCacheInfoOk returns a tuple with the EntryCacheInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCacheInfo

`func (o *GetPlugin200Response) SetEntryCacheInfo(v EnumpluginEntryCacheInfoProp)`

SetEntryCacheInfo sets EntryCacheInfo field to given value.

### HasEntryCacheInfo

`func (o *GetPlugin200Response) HasEntryCacheInfo() bool`

HasEntryCacheInfo returns a boolean if a field has been set.

### GetHostInfo

`func (o *GetPlugin200Response) GetHostInfo() []EnumpluginHostInfoProp`

GetHostInfo returns the HostInfo field if non-nil, zero value otherwise.

### GetHostInfoOk

`func (o *GetPlugin200Response) GetHostInfoOk() (*[]EnumpluginHostInfoProp, bool)`

GetHostInfoOk returns a tuple with the HostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfo

`func (o *GetPlugin200Response) SetHostInfo(v []EnumpluginHostInfoProp)`

SetHostInfo sets HostInfo field to given value.

### HasHostInfo

`func (o *GetPlugin200Response) HasHostInfo() bool`

HasHostInfo returns a boolean if a field has been set.

### GetIncludedLDAPApplication

`func (o *GetPlugin200Response) GetIncludedLDAPApplication() []string`

GetIncludedLDAPApplication returns the IncludedLDAPApplication field if non-nil, zero value otherwise.

### GetIncludedLDAPApplicationOk

`func (o *GetPlugin200Response) GetIncludedLDAPApplicationOk() (*[]string, bool)`

GetIncludedLDAPApplicationOk returns a tuple with the IncludedLDAPApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPApplication

`func (o *GetPlugin200Response) SetIncludedLDAPApplication(v []string)`

SetIncludedLDAPApplication sets IncludedLDAPApplication field to given value.

### HasIncludedLDAPApplication

`func (o *GetPlugin200Response) HasIncludedLDAPApplication() bool`

HasIncludedLDAPApplication returns a boolean if a field has been set.

### GetTraditionalStaticGroupObjectClass

`func (o *GetPlugin200Response) GetTraditionalStaticGroupObjectClass() EnumpluginTraditionalStaticGroupObjectClassProp`

GetTraditionalStaticGroupObjectClass returns the TraditionalStaticGroupObjectClass field if non-nil, zero value otherwise.

### GetTraditionalStaticGroupObjectClassOk

`func (o *GetPlugin200Response) GetTraditionalStaticGroupObjectClassOk() (*EnumpluginTraditionalStaticGroupObjectClassProp, bool)`

GetTraditionalStaticGroupObjectClassOk returns a tuple with the TraditionalStaticGroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraditionalStaticGroupObjectClass

`func (o *GetPlugin200Response) SetTraditionalStaticGroupObjectClass(v EnumpluginTraditionalStaticGroupObjectClassProp)`

SetTraditionalStaticGroupObjectClass sets TraditionalStaticGroupObjectClass field to given value.

### HasTraditionalStaticGroupObjectClass

`func (o *GetPlugin200Response) HasTraditionalStaticGroupObjectClass() bool`

HasTraditionalStaticGroupObjectClass returns a boolean if a field has been set.

### GetMaximumMembershipUpdatesPerModify

`func (o *GetPlugin200Response) GetMaximumMembershipUpdatesPerModify() int64`

GetMaximumMembershipUpdatesPerModify returns the MaximumMembershipUpdatesPerModify field if non-nil, zero value otherwise.

### GetMaximumMembershipUpdatesPerModifyOk

`func (o *GetPlugin200Response) GetMaximumMembershipUpdatesPerModifyOk() (*int64, bool)`

GetMaximumMembershipUpdatesPerModifyOk returns a tuple with the MaximumMembershipUpdatesPerModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMembershipUpdatesPerModify

`func (o *GetPlugin200Response) SetMaximumMembershipUpdatesPerModify(v int64)`

SetMaximumMembershipUpdatesPerModify sets MaximumMembershipUpdatesPerModify field to given value.

### HasMaximumMembershipUpdatesPerModify

`func (o *GetPlugin200Response) HasMaximumMembershipUpdatesPerModify() bool`

HasMaximumMembershipUpdatesPerModify returns a boolean if a field has been set.

### GetReadOperationSupport

`func (o *GetPlugin200Response) GetReadOperationSupport() EnumpluginReadOperationSupportProp`

GetReadOperationSupport returns the ReadOperationSupport field if non-nil, zero value otherwise.

### GetReadOperationSupportOk

`func (o *GetPlugin200Response) GetReadOperationSupportOk() (*EnumpluginReadOperationSupportProp, bool)`

GetReadOperationSupportOk returns a tuple with the ReadOperationSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOperationSupport

`func (o *GetPlugin200Response) SetReadOperationSupport(v EnumpluginReadOperationSupportProp)`

SetReadOperationSupport sets ReadOperationSupport field to given value.

### HasReadOperationSupport

`func (o *GetPlugin200Response) HasReadOperationSupport() bool`

HasReadOperationSupport returns a boolean if a field has been set.

### GetPluginType

`func (o *GetPlugin200Response) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *GetPlugin200Response) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *GetPlugin200Response) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetNumThreads

`func (o *GetPlugin200Response) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *GetPlugin200Response) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *GetPlugin200Response) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.


### GetBaseDN

`func (o *GetPlugin200Response) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *GetPlugin200Response) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *GetPlugin200Response) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetLowerBound

`func (o *GetPlugin200Response) GetLowerBound() int64`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *GetPlugin200Response) GetLowerBoundOk() (*int64, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *GetPlugin200Response) SetLowerBound(v int64)`

SetLowerBound sets LowerBound field to given value.

### HasLowerBound

`func (o *GetPlugin200Response) HasLowerBound() bool`

HasLowerBound returns a boolean if a field has been set.

### GetUpperBound

`func (o *GetPlugin200Response) GetUpperBound() int64`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *GetPlugin200Response) GetUpperBoundOk() (*int64, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *GetPlugin200Response) SetUpperBound(v int64)`

SetUpperBound sets UpperBound field to given value.

### HasUpperBound

`func (o *GetPlugin200Response) HasUpperBound() bool`

HasUpperBound returns a boolean if a field has been set.

### GetFilterPrefix

`func (o *GetPlugin200Response) GetFilterPrefix() string`

GetFilterPrefix returns the FilterPrefix field if non-nil, zero value otherwise.

### GetFilterPrefixOk

`func (o *GetPlugin200Response) GetFilterPrefixOk() (*string, bool)`

GetFilterPrefixOk returns a tuple with the FilterPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPrefix

`func (o *GetPlugin200Response) SetFilterPrefix(v string)`

SetFilterPrefix sets FilterPrefix field to given value.


### GetFilterSuffix

`func (o *GetPlugin200Response) GetFilterSuffix() string`

GetFilterSuffix returns the FilterSuffix field if non-nil, zero value otherwise.

### GetFilterSuffixOk

`func (o *GetPlugin200Response) GetFilterSuffixOk() (*string, bool)`

GetFilterSuffixOk returns a tuple with the FilterSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSuffix

`func (o *GetPlugin200Response) SetFilterSuffix(v string)`

SetFilterSuffix sets FilterSuffix field to given value.

### HasFilterSuffix

`func (o *GetPlugin200Response) HasFilterSuffix() bool`

HasFilterSuffix returns a boolean if a field has been set.

### GetFilter

`func (o *GetPlugin200Response) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetPlugin200Response) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetPlugin200Response) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetAttributeType

`func (o *GetPlugin200Response) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *GetPlugin200Response) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *GetPlugin200Response) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetPollingInterval

`func (o *GetPlugin200Response) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *GetPlugin200Response) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *GetPlugin200Response) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.


### GetPeerServerPriorityIndex

`func (o *GetPlugin200Response) GetPeerServerPriorityIndex() int64`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *GetPlugin200Response) GetPeerServerPriorityIndexOk() (*int64, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *GetPlugin200Response) SetPeerServerPriorityIndex(v int64)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *GetPlugin200Response) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetMaxUpdatesPerSecond

`func (o *GetPlugin200Response) GetMaxUpdatesPerSecond() int64`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *GetPlugin200Response) GetMaxUpdatesPerSecondOk() (*int64, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *GetPlugin200Response) SetMaxUpdatesPerSecond(v int64)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.


### GetNumDeleteThreads

`func (o *GetPlugin200Response) GetNumDeleteThreads() int64`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *GetPlugin200Response) GetNumDeleteThreadsOk() (*int64, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *GetPlugin200Response) SetNumDeleteThreads(v int64)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.


### GetInvokeGCDayOfWeek

`func (o *GetPlugin200Response) GetInvokeGCDayOfWeek() []EnumpluginInvokeGCDayOfWeekProp`

GetInvokeGCDayOfWeek returns the InvokeGCDayOfWeek field if non-nil, zero value otherwise.

### GetInvokeGCDayOfWeekOk

`func (o *GetPlugin200Response) GetInvokeGCDayOfWeekOk() (*[]EnumpluginInvokeGCDayOfWeekProp, bool)`

GetInvokeGCDayOfWeekOk returns a tuple with the InvokeGCDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeGCDayOfWeek

`func (o *GetPlugin200Response) SetInvokeGCDayOfWeek(v []EnumpluginInvokeGCDayOfWeekProp)`

SetInvokeGCDayOfWeek sets InvokeGCDayOfWeek field to given value.

### HasInvokeGCDayOfWeek

`func (o *GetPlugin200Response) HasInvokeGCDayOfWeek() bool`

HasInvokeGCDayOfWeek returns a boolean if a field has been set.

### GetInvokeGCTimeUtc

`func (o *GetPlugin200Response) GetInvokeGCTimeUtc() []string`

GetInvokeGCTimeUtc returns the InvokeGCTimeUtc field if non-nil, zero value otherwise.

### GetInvokeGCTimeUtcOk

`func (o *GetPlugin200Response) GetInvokeGCTimeUtcOk() (*[]string, bool)`

GetInvokeGCTimeUtcOk returns a tuple with the InvokeGCTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeGCTimeUtc

`func (o *GetPlugin200Response) SetInvokeGCTimeUtc(v []string)`

SetInvokeGCTimeUtc sets InvokeGCTimeUtc field to given value.


### GetDelayAfterAlert

`func (o *GetPlugin200Response) GetDelayAfterAlert() string`

GetDelayAfterAlert returns the DelayAfterAlert field if non-nil, zero value otherwise.

### GetDelayAfterAlertOk

`func (o *GetPlugin200Response) GetDelayAfterAlertOk() (*string, bool)`

GetDelayAfterAlertOk returns a tuple with the DelayAfterAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayAfterAlert

`func (o *GetPlugin200Response) SetDelayAfterAlert(v string)`

SetDelayAfterAlert sets DelayAfterAlert field to given value.

### HasDelayAfterAlert

`func (o *GetPlugin200Response) HasDelayAfterAlert() bool`

HasDelayAfterAlert returns a boolean if a field has been set.

### GetDelayPostGC

`func (o *GetPlugin200Response) GetDelayPostGC() string`

GetDelayPostGC returns the DelayPostGC field if non-nil, zero value otherwise.

### GetDelayPostGCOk

`func (o *GetPlugin200Response) GetDelayPostGCOk() (*string, bool)`

GetDelayPostGCOk returns a tuple with the DelayPostGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayPostGC

`func (o *GetPlugin200Response) SetDelayPostGC(v string)`

SetDelayPostGC sets DelayPostGC field to given value.

### HasDelayPostGC

`func (o *GetPlugin200Response) HasDelayPostGC() bool`

HasDelayPostGC returns a boolean if a field has been set.

### GetApiURL

`func (o *GetPlugin200Response) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *GetPlugin200Response) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *GetPlugin200Response) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.


### GetAuthURL

`func (o *GetPlugin200Response) GetAuthURL() string`

GetAuthURL returns the AuthURL field if non-nil, zero value otherwise.

### GetAuthURLOk

`func (o *GetPlugin200Response) GetAuthURLOk() (*string, bool)`

GetAuthURLOk returns a tuple with the AuthURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthURL

`func (o *GetPlugin200Response) SetAuthURL(v string)`

SetAuthURL sets AuthURL field to given value.


### GetOAuthClientID

`func (o *GetPlugin200Response) GetOAuthClientID() string`

GetOAuthClientID returns the OAuthClientID field if non-nil, zero value otherwise.

### GetOAuthClientIDOk

`func (o *GetPlugin200Response) GetOAuthClientIDOk() (*string, bool)`

GetOAuthClientIDOk returns a tuple with the OAuthClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientID

`func (o *GetPlugin200Response) SetOAuthClientID(v string)`

SetOAuthClientID sets OAuthClientID field to given value.


### GetOAuthClientSecret

`func (o *GetPlugin200Response) GetOAuthClientSecret() string`

GetOAuthClientSecret returns the OAuthClientSecret field if non-nil, zero value otherwise.

### GetOAuthClientSecretOk

`func (o *GetPlugin200Response) GetOAuthClientSecretOk() (*string, bool)`

GetOAuthClientSecretOk returns a tuple with the OAuthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecret

`func (o *GetPlugin200Response) SetOAuthClientSecret(v string)`

SetOAuthClientSecret sets OAuthClientSecret field to given value.

### HasOAuthClientSecret

`func (o *GetPlugin200Response) HasOAuthClientSecret() bool`

HasOAuthClientSecret returns a boolean if a field has been set.

### GetOAuthClientSecretPassphraseProvider

`func (o *GetPlugin200Response) GetOAuthClientSecretPassphraseProvider() string`

GetOAuthClientSecretPassphraseProvider returns the OAuthClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetOAuthClientSecretPassphraseProviderOk

`func (o *GetPlugin200Response) GetOAuthClientSecretPassphraseProviderOk() (*string, bool)`

GetOAuthClientSecretPassphraseProviderOk returns a tuple with the OAuthClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecretPassphraseProvider

`func (o *GetPlugin200Response) SetOAuthClientSecretPassphraseProvider(v string)`

SetOAuthClientSecretPassphraseProvider sets OAuthClientSecretPassphraseProvider field to given value.

### HasOAuthClientSecretPassphraseProvider

`func (o *GetPlugin200Response) HasOAuthClientSecretPassphraseProvider() bool`

HasOAuthClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *GetPlugin200Response) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *GetPlugin200Response) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *GetPlugin200Response) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.


### GetHttpProxyExternalServer

`func (o *GetPlugin200Response) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *GetPlugin200Response) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *GetPlugin200Response) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *GetPlugin200Response) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *GetPlugin200Response) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *GetPlugin200Response) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *GetPlugin200Response) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *GetPlugin200Response) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *GetPlugin200Response) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *GetPlugin200Response) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *GetPlugin200Response) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *GetPlugin200Response) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetTryLocalBind

`func (o *GetPlugin200Response) GetTryLocalBind() bool`

GetTryLocalBind returns the TryLocalBind field if non-nil, zero value otherwise.

### GetTryLocalBindOk

`func (o *GetPlugin200Response) GetTryLocalBindOk() (*bool, bool)`

GetTryLocalBindOk returns a tuple with the TryLocalBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryLocalBind

`func (o *GetPlugin200Response) SetTryLocalBind(v bool)`

SetTryLocalBind sets TryLocalBind field to given value.


### GetOverrideLocalPassword

`func (o *GetPlugin200Response) GetOverrideLocalPassword() bool`

GetOverrideLocalPassword returns the OverrideLocalPassword field if non-nil, zero value otherwise.

### GetOverrideLocalPasswordOk

`func (o *GetPlugin200Response) GetOverrideLocalPasswordOk() (*bool, bool)`

GetOverrideLocalPasswordOk returns a tuple with the OverrideLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLocalPassword

`func (o *GetPlugin200Response) SetOverrideLocalPassword(v bool)`

SetOverrideLocalPassword sets OverrideLocalPassword field to given value.


### GetUpdateLocalPassword

`func (o *GetPlugin200Response) GetUpdateLocalPassword() bool`

GetUpdateLocalPassword returns the UpdateLocalPassword field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordOk

`func (o *GetPlugin200Response) GetUpdateLocalPasswordOk() (*bool, bool)`

GetUpdateLocalPasswordOk returns a tuple with the UpdateLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPassword

`func (o *GetPlugin200Response) SetUpdateLocalPassword(v bool)`

SetUpdateLocalPassword sets UpdateLocalPassword field to given value.


### GetUpdateLocalPasswordDN

`func (o *GetPlugin200Response) GetUpdateLocalPasswordDN() string`

GetUpdateLocalPasswordDN returns the UpdateLocalPasswordDN field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordDNOk

`func (o *GetPlugin200Response) GetUpdateLocalPasswordDNOk() (*string, bool)`

GetUpdateLocalPasswordDNOk returns a tuple with the UpdateLocalPasswordDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPasswordDN

`func (o *GetPlugin200Response) SetUpdateLocalPasswordDN(v string)`

SetUpdateLocalPasswordDN sets UpdateLocalPasswordDN field to given value.

### HasUpdateLocalPasswordDN

`func (o *GetPlugin200Response) HasUpdateLocalPasswordDN() bool`

HasUpdateLocalPasswordDN returns a boolean if a field has been set.

### GetAllowLaxPassThroughAuthenticationPasswords

`func (o *GetPlugin200Response) GetAllowLaxPassThroughAuthenticationPasswords() bool`

GetAllowLaxPassThroughAuthenticationPasswords returns the AllowLaxPassThroughAuthenticationPasswords field if non-nil, zero value otherwise.

### GetAllowLaxPassThroughAuthenticationPasswordsOk

`func (o *GetPlugin200Response) GetAllowLaxPassThroughAuthenticationPasswordsOk() (*bool, bool)`

GetAllowLaxPassThroughAuthenticationPasswordsOk returns a tuple with the AllowLaxPassThroughAuthenticationPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLaxPassThroughAuthenticationPasswords

`func (o *GetPlugin200Response) SetAllowLaxPassThroughAuthenticationPasswords(v bool)`

SetAllowLaxPassThroughAuthenticationPasswords sets AllowLaxPassThroughAuthenticationPasswords field to given value.

### HasAllowLaxPassThroughAuthenticationPasswords

`func (o *GetPlugin200Response) HasAllowLaxPassThroughAuthenticationPasswords() bool`

HasAllowLaxPassThroughAuthenticationPasswords returns a boolean if a field has been set.

### GetIgnoredPasswordPolicyStateErrorCondition

`func (o *GetPlugin200Response) GetIgnoredPasswordPolicyStateErrorCondition() []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp`

GetIgnoredPasswordPolicyStateErrorCondition returns the IgnoredPasswordPolicyStateErrorCondition field if non-nil, zero value otherwise.

### GetIgnoredPasswordPolicyStateErrorConditionOk

`func (o *GetPlugin200Response) GetIgnoredPasswordPolicyStateErrorConditionOk() (*[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp, bool)`

GetIgnoredPasswordPolicyStateErrorConditionOk returns a tuple with the IgnoredPasswordPolicyStateErrorCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredPasswordPolicyStateErrorCondition

`func (o *GetPlugin200Response) SetIgnoredPasswordPolicyStateErrorCondition(v []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp)`

SetIgnoredPasswordPolicyStateErrorCondition sets IgnoredPasswordPolicyStateErrorCondition field to given value.

### HasIgnoredPasswordPolicyStateErrorCondition

`func (o *GetPlugin200Response) HasIgnoredPasswordPolicyStateErrorCondition() bool`

HasIgnoredPasswordPolicyStateErrorCondition returns a boolean if a field has been set.

### GetUserMappingLocalAttribute

`func (o *GetPlugin200Response) GetUserMappingLocalAttribute() []string`

GetUserMappingLocalAttribute returns the UserMappingLocalAttribute field if non-nil, zero value otherwise.

### GetUserMappingLocalAttributeOk

`func (o *GetPlugin200Response) GetUserMappingLocalAttributeOk() (*[]string, bool)`

GetUserMappingLocalAttributeOk returns a tuple with the UserMappingLocalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingLocalAttribute

`func (o *GetPlugin200Response) SetUserMappingLocalAttribute(v []string)`

SetUserMappingLocalAttribute sets UserMappingLocalAttribute field to given value.


### GetUserMappingRemoteJSONField

`func (o *GetPlugin200Response) GetUserMappingRemoteJSONField() []string`

GetUserMappingRemoteJSONField returns the UserMappingRemoteJSONField field if non-nil, zero value otherwise.

### GetUserMappingRemoteJSONFieldOk

`func (o *GetPlugin200Response) GetUserMappingRemoteJSONFieldOk() (*[]string, bool)`

GetUserMappingRemoteJSONFieldOk returns a tuple with the UserMappingRemoteJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingRemoteJSONField

`func (o *GetPlugin200Response) SetUserMappingRemoteJSONField(v []string)`

SetUserMappingRemoteJSONField sets UserMappingRemoteJSONField field to given value.


### GetAdditionalUserMappingSCIMFilter

`func (o *GetPlugin200Response) GetAdditionalUserMappingSCIMFilter() string`

GetAdditionalUserMappingSCIMFilter returns the AdditionalUserMappingSCIMFilter field if non-nil, zero value otherwise.

### GetAdditionalUserMappingSCIMFilterOk

`func (o *GetPlugin200Response) GetAdditionalUserMappingSCIMFilterOk() (*string, bool)`

GetAdditionalUserMappingSCIMFilterOk returns a tuple with the AdditionalUserMappingSCIMFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserMappingSCIMFilter

`func (o *GetPlugin200Response) SetAdditionalUserMappingSCIMFilter(v string)`

SetAdditionalUserMappingSCIMFilter sets AdditionalUserMappingSCIMFilter field to given value.

### HasAdditionalUserMappingSCIMFilter

`func (o *GetPlugin200Response) HasAdditionalUserMappingSCIMFilter() bool`

HasAdditionalUserMappingSCIMFilter returns a boolean if a field has been set.

### GetChangelogPasswordEncryptionKey

`func (o *GetPlugin200Response) GetChangelogPasswordEncryptionKey() string`

GetChangelogPasswordEncryptionKey returns the ChangelogPasswordEncryptionKey field if non-nil, zero value otherwise.

### GetChangelogPasswordEncryptionKeyOk

`func (o *GetPlugin200Response) GetChangelogPasswordEncryptionKeyOk() (*string, bool)`

GetChangelogPasswordEncryptionKeyOk returns a tuple with the ChangelogPasswordEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPasswordEncryptionKey

`func (o *GetPlugin200Response) SetChangelogPasswordEncryptionKey(v string)`

SetChangelogPasswordEncryptionKey sets ChangelogPasswordEncryptionKey field to given value.

### HasChangelogPasswordEncryptionKey

`func (o *GetPlugin200Response) HasChangelogPasswordEncryptionKey() bool`

HasChangelogPasswordEncryptionKey returns a boolean if a field has been set.

### GetChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *GetPlugin200Response) GetChangelogPasswordEncryptionKeyPassphraseProvider() string`

GetChangelogPasswordEncryptionKeyPassphraseProvider returns the ChangelogPasswordEncryptionKeyPassphraseProvider field if non-nil, zero value otherwise.

### GetChangelogPasswordEncryptionKeyPassphraseProviderOk

`func (o *GetPlugin200Response) GetChangelogPasswordEncryptionKeyPassphraseProviderOk() (*string, bool)`

GetChangelogPasswordEncryptionKeyPassphraseProviderOk returns a tuple with the ChangelogPasswordEncryptionKeyPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *GetPlugin200Response) SetChangelogPasswordEncryptionKeyPassphraseProvider(v string)`

SetChangelogPasswordEncryptionKeyPassphraseProvider sets ChangelogPasswordEncryptionKeyPassphraseProvider field to given value.

### HasChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *GetPlugin200Response) HasChangelogPasswordEncryptionKeyPassphraseProvider() bool`

HasChangelogPasswordEncryptionKeyPassphraseProvider returns a boolean if a field has been set.

### GetHistogramCategoryBoundary

`func (o *GetPlugin200Response) GetHistogramCategoryBoundary() []string`

GetHistogramCategoryBoundary returns the HistogramCategoryBoundary field if non-nil, zero value otherwise.

### GetHistogramCategoryBoundaryOk

`func (o *GetPlugin200Response) GetHistogramCategoryBoundaryOk() (*[]string, bool)`

GetHistogramCategoryBoundaryOk returns a tuple with the HistogramCategoryBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramCategoryBoundary

`func (o *GetPlugin200Response) SetHistogramCategoryBoundary(v []string)`

SetHistogramCategoryBoundary sets HistogramCategoryBoundary field to given value.


### GetIncludeQueueTime

`func (o *GetPlugin200Response) GetIncludeQueueTime() bool`

GetIncludeQueueTime returns the IncludeQueueTime field if non-nil, zero value otherwise.

### GetIncludeQueueTimeOk

`func (o *GetPlugin200Response) GetIncludeQueueTimeOk() (*bool, bool)`

GetIncludeQueueTimeOk returns a tuple with the IncludeQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeQueueTime

`func (o *GetPlugin200Response) SetIncludeQueueTime(v bool)`

SetIncludeQueueTime sets IncludeQueueTime field to given value.

### HasIncludeQueueTime

`func (o *GetPlugin200Response) HasIncludeQueueTime() bool`

HasIncludeQueueTime returns a boolean if a field has been set.

### GetSeparateMonitorEntryPerTrackedApplication

`func (o *GetPlugin200Response) GetSeparateMonitorEntryPerTrackedApplication() bool`

GetSeparateMonitorEntryPerTrackedApplication returns the SeparateMonitorEntryPerTrackedApplication field if non-nil, zero value otherwise.

### GetSeparateMonitorEntryPerTrackedApplicationOk

`func (o *GetPlugin200Response) GetSeparateMonitorEntryPerTrackedApplicationOk() (*bool, bool)`

GetSeparateMonitorEntryPerTrackedApplicationOk returns a tuple with the SeparateMonitorEntryPerTrackedApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateMonitorEntryPerTrackedApplication

`func (o *GetPlugin200Response) SetSeparateMonitorEntryPerTrackedApplication(v bool)`

SetSeparateMonitorEntryPerTrackedApplication sets SeparateMonitorEntryPerTrackedApplication field to given value.

### HasSeparateMonitorEntryPerTrackedApplication

`func (o *GetPlugin200Response) HasSeparateMonitorEntryPerTrackedApplication() bool`

HasSeparateMonitorEntryPerTrackedApplication returns a boolean if a field has been set.

### GetScope

`func (o *GetPlugin200Response) GetScope() EnumpluginScopeProp`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetPlugin200Response) GetScopeOk() (*EnumpluginScopeProp, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetPlugin200Response) SetScope(v EnumpluginScopeProp)`

SetScope sets Scope field to given value.


### GetIncludeAttribute

`func (o *GetPlugin200Response) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *GetPlugin200Response) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *GetPlugin200Response) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *GetPlugin200Response) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetOutputFile

`func (o *GetPlugin200Response) GetOutputFile() string`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *GetPlugin200Response) GetOutputFileOk() (*string, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *GetPlugin200Response) SetOutputFile(v string)`

SetOutputFile sets OutputFile field to given value.


### GetPreviousFileExtension

`func (o *GetPlugin200Response) GetPreviousFileExtension() string`

GetPreviousFileExtension returns the PreviousFileExtension field if non-nil, zero value otherwise.

### GetPreviousFileExtensionOk

`func (o *GetPlugin200Response) GetPreviousFileExtensionOk() (*string, bool)`

GetPreviousFileExtensionOk returns a tuple with the PreviousFileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousFileExtension

`func (o *GetPlugin200Response) SetPreviousFileExtension(v string)`

SetPreviousFileExtension sets PreviousFileExtension field to given value.

### HasPreviousFileExtension

`func (o *GetPlugin200Response) HasPreviousFileExtension() bool`

HasPreviousFileExtension returns a boolean if a field has been set.

### GetLogInterval

`func (o *GetPlugin200Response) GetLogInterval() string`

GetLogInterval returns the LogInterval field if non-nil, zero value otherwise.

### GetLogIntervalOk

`func (o *GetPlugin200Response) GetLogIntervalOk() (*string, bool)`

GetLogIntervalOk returns a tuple with the LogInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogInterval

`func (o *GetPlugin200Response) SetLogInterval(v string)`

SetLogInterval sets LogInterval field to given value.


### GetSuppressIfIdle

`func (o *GetPlugin200Response) GetSuppressIfIdle() bool`

GetSuppressIfIdle returns the SuppressIfIdle field if non-nil, zero value otherwise.

### GetSuppressIfIdleOk

`func (o *GetPlugin200Response) GetSuppressIfIdleOk() (*bool, bool)`

GetSuppressIfIdleOk returns a tuple with the SuppressIfIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressIfIdle

`func (o *GetPlugin200Response) SetSuppressIfIdle(v bool)`

SetSuppressIfIdle sets SuppressIfIdle field to given value.


### GetHeaderPrefixPerColumn

`func (o *GetPlugin200Response) GetHeaderPrefixPerColumn() bool`

GetHeaderPrefixPerColumn returns the HeaderPrefixPerColumn field if non-nil, zero value otherwise.

### GetHeaderPrefixPerColumnOk

`func (o *GetPlugin200Response) GetHeaderPrefixPerColumnOk() (*bool, bool)`

GetHeaderPrefixPerColumnOk returns a tuple with the HeaderPrefixPerColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPrefixPerColumn

`func (o *GetPlugin200Response) SetHeaderPrefixPerColumn(v bool)`

SetHeaderPrefixPerColumn sets HeaderPrefixPerColumn field to given value.

### HasHeaderPrefixPerColumn

`func (o *GetPlugin200Response) HasHeaderPrefixPerColumn() bool`

HasHeaderPrefixPerColumn returns a boolean if a field has been set.

### GetEmptyInsteadOfZero

`func (o *GetPlugin200Response) GetEmptyInsteadOfZero() bool`

GetEmptyInsteadOfZero returns the EmptyInsteadOfZero field if non-nil, zero value otherwise.

### GetEmptyInsteadOfZeroOk

`func (o *GetPlugin200Response) GetEmptyInsteadOfZeroOk() (*bool, bool)`

GetEmptyInsteadOfZeroOk returns a tuple with the EmptyInsteadOfZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyInsteadOfZero

`func (o *GetPlugin200Response) SetEmptyInsteadOfZero(v bool)`

SetEmptyInsteadOfZero sets EmptyInsteadOfZero field to given value.

### HasEmptyInsteadOfZero

`func (o *GetPlugin200Response) HasEmptyInsteadOfZero() bool`

HasEmptyInsteadOfZero returns a boolean if a field has been set.

### GetLinesBetweenHeader

`func (o *GetPlugin200Response) GetLinesBetweenHeader() int64`

GetLinesBetweenHeader returns the LinesBetweenHeader field if non-nil, zero value otherwise.

### GetLinesBetweenHeaderOk

`func (o *GetPlugin200Response) GetLinesBetweenHeaderOk() (*int64, bool)`

GetLinesBetweenHeaderOk returns a tuple with the LinesBetweenHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesBetweenHeader

`func (o *GetPlugin200Response) SetLinesBetweenHeader(v int64)`

SetLinesBetweenHeader sets LinesBetweenHeader field to given value.


### GetIncludedLDAPStat

`func (o *GetPlugin200Response) GetIncludedLDAPStat() []EnumpluginIncludedLDAPStatProp`

GetIncludedLDAPStat returns the IncludedLDAPStat field if non-nil, zero value otherwise.

### GetIncludedLDAPStatOk

`func (o *GetPlugin200Response) GetIncludedLDAPStatOk() (*[]EnumpluginIncludedLDAPStatProp, bool)`

GetIncludedLDAPStatOk returns a tuple with the IncludedLDAPStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPStat

`func (o *GetPlugin200Response) SetIncludedLDAPStat(v []EnumpluginIncludedLDAPStatProp)`

SetIncludedLDAPStat sets IncludedLDAPStat field to given value.

### HasIncludedLDAPStat

`func (o *GetPlugin200Response) HasIncludedLDAPStat() bool`

HasIncludedLDAPStat returns a boolean if a field has been set.

### GetIncludedResourceStat

`func (o *GetPlugin200Response) GetIncludedResourceStat() []EnumpluginIncludedResourceStatProp`

GetIncludedResourceStat returns the IncludedResourceStat field if non-nil, zero value otherwise.

### GetIncludedResourceStatOk

`func (o *GetPlugin200Response) GetIncludedResourceStatOk() (*[]EnumpluginIncludedResourceStatProp, bool)`

GetIncludedResourceStatOk returns a tuple with the IncludedResourceStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedResourceStat

`func (o *GetPlugin200Response) SetIncludedResourceStat(v []EnumpluginIncludedResourceStatProp)`

SetIncludedResourceStat sets IncludedResourceStat field to given value.

### HasIncludedResourceStat

`func (o *GetPlugin200Response) HasIncludedResourceStat() bool`

HasIncludedResourceStat returns a boolean if a field has been set.

### GetHistogramFormat

`func (o *GetPlugin200Response) GetHistogramFormat() EnumpluginHistogramFormatProp`

GetHistogramFormat returns the HistogramFormat field if non-nil, zero value otherwise.

### GetHistogramFormatOk

`func (o *GetPlugin200Response) GetHistogramFormatOk() (*EnumpluginHistogramFormatProp, bool)`

GetHistogramFormatOk returns a tuple with the HistogramFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramFormat

`func (o *GetPlugin200Response) SetHistogramFormat(v EnumpluginHistogramFormatProp)`

SetHistogramFormat sets HistogramFormat field to given value.


### GetHistogramOpType

`func (o *GetPlugin200Response) GetHistogramOpType() []EnumpluginHistogramOpTypeProp`

GetHistogramOpType returns the HistogramOpType field if non-nil, zero value otherwise.

### GetHistogramOpTypeOk

`func (o *GetPlugin200Response) GetHistogramOpTypeOk() (*[]EnumpluginHistogramOpTypeProp, bool)`

GetHistogramOpTypeOk returns a tuple with the HistogramOpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramOpType

`func (o *GetPlugin200Response) SetHistogramOpType(v []EnumpluginHistogramOpTypeProp)`

SetHistogramOpType sets HistogramOpType field to given value.

### HasHistogramOpType

`func (o *GetPlugin200Response) HasHistogramOpType() bool`

HasHistogramOpType returns a boolean if a field has been set.

### GetGaugeInfo

`func (o *GetPlugin200Response) GetGaugeInfo() EnumpluginGaugeInfoProp`

GetGaugeInfo returns the GaugeInfo field if non-nil, zero value otherwise.

### GetGaugeInfoOk

`func (o *GetPlugin200Response) GetGaugeInfoOk() (*EnumpluginGaugeInfoProp, bool)`

GetGaugeInfoOk returns a tuple with the GaugeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeInfo

`func (o *GetPlugin200Response) SetGaugeInfo(v EnumpluginGaugeInfoProp)`

SetGaugeInfo sets GaugeInfo field to given value.

### HasGaugeInfo

`func (o *GetPlugin200Response) HasGaugeInfo() bool`

HasGaugeInfo returns a boolean if a field has been set.

### GetLogFileFormat

`func (o *GetPlugin200Response) GetLogFileFormat() EnumpluginLogFileFormatProp`

GetLogFileFormat returns the LogFileFormat field if non-nil, zero value otherwise.

### GetLogFileFormatOk

`func (o *GetPlugin200Response) GetLogFileFormatOk() (*EnumpluginLogFileFormatProp, bool)`

GetLogFileFormatOk returns a tuple with the LogFileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileFormat

`func (o *GetPlugin200Response) SetLogFileFormat(v EnumpluginLogFileFormatProp)`

SetLogFileFormat sets LogFileFormat field to given value.

### HasLogFileFormat

`func (o *GetPlugin200Response) HasLogFileFormat() bool`

HasLogFileFormat returns a boolean if a field has been set.

### GetLogFile

`func (o *GetPlugin200Response) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *GetPlugin200Response) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *GetPlugin200Response) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.


### GetLogFilePermissions

`func (o *GetPlugin200Response) GetLogFilePermissions() string`

GetLogFilePermissions returns the LogFilePermissions field if non-nil, zero value otherwise.

### GetLogFilePermissionsOk

`func (o *GetPlugin200Response) GetLogFilePermissionsOk() (*string, bool)`

GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFilePermissions

`func (o *GetPlugin200Response) SetLogFilePermissions(v string)`

SetLogFilePermissions sets LogFilePermissions field to given value.


### GetAppend

`func (o *GetPlugin200Response) GetAppend() bool`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *GetPlugin200Response) GetAppendOk() (*bool, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *GetPlugin200Response) SetAppend(v bool)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *GetPlugin200Response) HasAppend() bool`

HasAppend returns a boolean if a field has been set.

### GetRotationPolicy

`func (o *GetPlugin200Response) GetRotationPolicy() []string`

GetRotationPolicy returns the RotationPolicy field if non-nil, zero value otherwise.

### GetRotationPolicyOk

`func (o *GetPlugin200Response) GetRotationPolicyOk() (*[]string, bool)`

GetRotationPolicyOk returns a tuple with the RotationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPolicy

`func (o *GetPlugin200Response) SetRotationPolicy(v []string)`

SetRotationPolicy sets RotationPolicy field to given value.


### GetRotationListener

`func (o *GetPlugin200Response) GetRotationListener() []string`

GetRotationListener returns the RotationListener field if non-nil, zero value otherwise.

### GetRotationListenerOk

`func (o *GetPlugin200Response) GetRotationListenerOk() (*[]string, bool)`

GetRotationListenerOk returns a tuple with the RotationListener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationListener

`func (o *GetPlugin200Response) SetRotationListener(v []string)`

SetRotationListener sets RotationListener field to given value.

### HasRotationListener

`func (o *GetPlugin200Response) HasRotationListener() bool`

HasRotationListener returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *GetPlugin200Response) GetRetentionPolicy() []string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *GetPlugin200Response) GetRetentionPolicyOk() (*[]string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *GetPlugin200Response) SetRetentionPolicy(v []string)`

SetRetentionPolicy sets RetentionPolicy field to given value.


### GetLoggingErrorBehavior

`func (o *GetPlugin200Response) GetLoggingErrorBehavior() EnumpluginLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *GetPlugin200Response) GetLoggingErrorBehaviorOk() (*EnumpluginLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *GetPlugin200Response) SetLoggingErrorBehavior(v EnumpluginLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *GetPlugin200Response) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.

### GetDatetimeAttribute

`func (o *GetPlugin200Response) GetDatetimeAttribute() string`

GetDatetimeAttribute returns the DatetimeAttribute field if non-nil, zero value otherwise.

### GetDatetimeAttributeOk

`func (o *GetPlugin200Response) GetDatetimeAttributeOk() (*string, bool)`

GetDatetimeAttributeOk returns a tuple with the DatetimeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeAttribute

`func (o *GetPlugin200Response) SetDatetimeAttribute(v string)`

SetDatetimeAttribute sets DatetimeAttribute field to given value.


### GetDatetimeJSONField

`func (o *GetPlugin200Response) GetDatetimeJSONField() string`

GetDatetimeJSONField returns the DatetimeJSONField field if non-nil, zero value otherwise.

### GetDatetimeJSONFieldOk

`func (o *GetPlugin200Response) GetDatetimeJSONFieldOk() (*string, bool)`

GetDatetimeJSONFieldOk returns a tuple with the DatetimeJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeJSONField

`func (o *GetPlugin200Response) SetDatetimeJSONField(v string)`

SetDatetimeJSONField sets DatetimeJSONField field to given value.

### HasDatetimeJSONField

`func (o *GetPlugin200Response) HasDatetimeJSONField() bool`

HasDatetimeJSONField returns a boolean if a field has been set.

### GetDatetimeFormat

`func (o *GetPlugin200Response) GetDatetimeFormat() EnumpluginDatetimeFormatProp`

GetDatetimeFormat returns the DatetimeFormat field if non-nil, zero value otherwise.

### GetDatetimeFormatOk

`func (o *GetPlugin200Response) GetDatetimeFormatOk() (*EnumpluginDatetimeFormatProp, bool)`

GetDatetimeFormatOk returns a tuple with the DatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeFormat

`func (o *GetPlugin200Response) SetDatetimeFormat(v EnumpluginDatetimeFormatProp)`

SetDatetimeFormat sets DatetimeFormat field to given value.


### GetCustomDatetimeFormat

`func (o *GetPlugin200Response) GetCustomDatetimeFormat() string`

GetCustomDatetimeFormat returns the CustomDatetimeFormat field if non-nil, zero value otherwise.

### GetCustomDatetimeFormatOk

`func (o *GetPlugin200Response) GetCustomDatetimeFormatOk() (*string, bool)`

GetCustomDatetimeFormatOk returns a tuple with the CustomDatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDatetimeFormat

`func (o *GetPlugin200Response) SetCustomDatetimeFormat(v string)`

SetCustomDatetimeFormat sets CustomDatetimeFormat field to given value.

### HasCustomDatetimeFormat

`func (o *GetPlugin200Response) HasCustomDatetimeFormat() bool`

HasCustomDatetimeFormat returns a boolean if a field has been set.

### GetCustomTimezone

`func (o *GetPlugin200Response) GetCustomTimezone() string`

GetCustomTimezone returns the CustomTimezone field if non-nil, zero value otherwise.

### GetCustomTimezoneOk

`func (o *GetPlugin200Response) GetCustomTimezoneOk() (*string, bool)`

GetCustomTimezoneOk returns a tuple with the CustomTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimezone

`func (o *GetPlugin200Response) SetCustomTimezone(v string)`

SetCustomTimezone sets CustomTimezone field to given value.

### HasCustomTimezone

`func (o *GetPlugin200Response) HasCustomTimezone() bool`

HasCustomTimezone returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *GetPlugin200Response) GetExpirationOffset() string`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *GetPlugin200Response) GetExpirationOffsetOk() (*string, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *GetPlugin200Response) SetExpirationOffset(v string)`

SetExpirationOffset sets ExpirationOffset field to given value.


### GetPurgeBehavior

`func (o *GetPlugin200Response) GetPurgeBehavior() EnumpluginPurgeBehaviorProp`

GetPurgeBehavior returns the PurgeBehavior field if non-nil, zero value otherwise.

### GetPurgeBehaviorOk

`func (o *GetPlugin200Response) GetPurgeBehaviorOk() (*EnumpluginPurgeBehaviorProp, bool)`

GetPurgeBehaviorOk returns a tuple with the PurgeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeBehavior

`func (o *GetPlugin200Response) SetPurgeBehavior(v EnumpluginPurgeBehaviorProp)`

SetPurgeBehavior sets PurgeBehavior field to given value.

### HasPurgeBehavior

`func (o *GetPlugin200Response) HasPurgeBehavior() bool`

HasPurgeBehavior returns a boolean if a field has been set.

### GetNumMostExpensivePhasesShown

`func (o *GetPlugin200Response) GetNumMostExpensivePhasesShown() int64`

GetNumMostExpensivePhasesShown returns the NumMostExpensivePhasesShown field if non-nil, zero value otherwise.

### GetNumMostExpensivePhasesShownOk

`func (o *GetPlugin200Response) GetNumMostExpensivePhasesShownOk() (*int64, bool)`

GetNumMostExpensivePhasesShownOk returns a tuple with the NumMostExpensivePhasesShown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMostExpensivePhasesShown

`func (o *GetPlugin200Response) SetNumMostExpensivePhasesShown(v int64)`

SetNumMostExpensivePhasesShown sets NumMostExpensivePhasesShown field to given value.


### GetExtensionClass

`func (o *GetPlugin200Response) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *GetPlugin200Response) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *GetPlugin200Response) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *GetPlugin200Response) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *GetPlugin200Response) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *GetPlugin200Response) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *GetPlugin200Response) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *GetPlugin200Response) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *GetPlugin200Response) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *GetPlugin200Response) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *GetPlugin200Response) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetServer

`func (o *GetPlugin200Response) GetServer() []string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetPlugin200Response) GetServerOk() (*[]string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetPlugin200Response) SetServer(v []string)`

SetServer sets Server field to given value.


### GetServerAccessMode

`func (o *GetPlugin200Response) GetServerAccessMode() EnumpluginServerAccessModeProp`

GetServerAccessMode returns the ServerAccessMode field if non-nil, zero value otherwise.

### GetServerAccessModeOk

`func (o *GetPlugin200Response) GetServerAccessModeOk() (*EnumpluginServerAccessModeProp, bool)`

GetServerAccessModeOk returns a tuple with the ServerAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAccessMode

`func (o *GetPlugin200Response) SetServerAccessMode(v EnumpluginServerAccessModeProp)`

SetServerAccessMode sets ServerAccessMode field to given value.


### GetDnMap

`func (o *GetPlugin200Response) GetDnMap() []string`

GetDnMap returns the DnMap field if non-nil, zero value otherwise.

### GetDnMapOk

`func (o *GetPlugin200Response) GetDnMapOk() (*[]string, bool)`

GetDnMapOk returns a tuple with the DnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnMap

`func (o *GetPlugin200Response) SetDnMap(v []string)`

SetDnMap sets DnMap field to given value.

### HasDnMap

`func (o *GetPlugin200Response) HasDnMap() bool`

HasDnMap returns a boolean if a field has been set.

### GetBindDNPattern

`func (o *GetPlugin200Response) GetBindDNPattern() string`

GetBindDNPattern returns the BindDNPattern field if non-nil, zero value otherwise.

### GetBindDNPatternOk

`func (o *GetPlugin200Response) GetBindDNPatternOk() (*string, bool)`

GetBindDNPatternOk returns a tuple with the BindDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDNPattern

`func (o *GetPlugin200Response) SetBindDNPattern(v string)`

SetBindDNPattern sets BindDNPattern field to given value.

### HasBindDNPattern

`func (o *GetPlugin200Response) HasBindDNPattern() bool`

HasBindDNPattern returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *GetPlugin200Response) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *GetPlugin200Response) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *GetPlugin200Response) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.

### HasSearchBaseDN

`func (o *GetPlugin200Response) HasSearchBaseDN() bool`

HasSearchBaseDN returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *GetPlugin200Response) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *GetPlugin200Response) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *GetPlugin200Response) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *GetPlugin200Response) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetInitialConnections

`func (o *GetPlugin200Response) GetInitialConnections() int64`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *GetPlugin200Response) GetInitialConnectionsOk() (*int64, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *GetPlugin200Response) SetInitialConnections(v int64)`

SetInitialConnections sets InitialConnections field to given value.


### GetMaxConnections

`func (o *GetPlugin200Response) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *GetPlugin200Response) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *GetPlugin200Response) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.


### GetSourceDN

`func (o *GetPlugin200Response) GetSourceDN() string`

GetSourceDN returns the SourceDN field if non-nil, zero value otherwise.

### GetSourceDNOk

`func (o *GetPlugin200Response) GetSourceDNOk() (*string, bool)`

GetSourceDNOk returns a tuple with the SourceDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDN

`func (o *GetPlugin200Response) SetSourceDN(v string)`

SetSourceDN sets SourceDN field to given value.


### GetTargetDN

`func (o *GetPlugin200Response) GetTargetDN() string`

GetTargetDN returns the TargetDN field if non-nil, zero value otherwise.

### GetTargetDNOk

`func (o *GetPlugin200Response) GetTargetDNOk() (*string, bool)`

GetTargetDNOk returns a tuple with the TargetDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDN

`func (o *GetPlugin200Response) SetTargetDN(v string)`

SetTargetDN sets TargetDN field to given value.


### GetEnableAttributeMapping

`func (o *GetPlugin200Response) GetEnableAttributeMapping() bool`

GetEnableAttributeMapping returns the EnableAttributeMapping field if non-nil, zero value otherwise.

### GetEnableAttributeMappingOk

`func (o *GetPlugin200Response) GetEnableAttributeMappingOk() (*bool, bool)`

GetEnableAttributeMappingOk returns a tuple with the EnableAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAttributeMapping

`func (o *GetPlugin200Response) SetEnableAttributeMapping(v bool)`

SetEnableAttributeMapping sets EnableAttributeMapping field to given value.


### GetMapAttribute

`func (o *GetPlugin200Response) GetMapAttribute() []string`

GetMapAttribute returns the MapAttribute field if non-nil, zero value otherwise.

### GetMapAttributeOk

`func (o *GetPlugin200Response) GetMapAttributeOk() (*[]string, bool)`

GetMapAttributeOk returns a tuple with the MapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAttribute

`func (o *GetPlugin200Response) SetMapAttribute(v []string)`

SetMapAttribute sets MapAttribute field to given value.

### HasMapAttribute

`func (o *GetPlugin200Response) HasMapAttribute() bool`

HasMapAttribute returns a boolean if a field has been set.

### GetEnableControlMapping

`func (o *GetPlugin200Response) GetEnableControlMapping() bool`

GetEnableControlMapping returns the EnableControlMapping field if non-nil, zero value otherwise.

### GetEnableControlMappingOk

`func (o *GetPlugin200Response) GetEnableControlMappingOk() (*bool, bool)`

GetEnableControlMappingOk returns a tuple with the EnableControlMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableControlMapping

`func (o *GetPlugin200Response) SetEnableControlMapping(v bool)`

SetEnableControlMapping sets EnableControlMapping field to given value.


### GetAlwaysMapResponses

`func (o *GetPlugin200Response) GetAlwaysMapResponses() bool`

GetAlwaysMapResponses returns the AlwaysMapResponses field if non-nil, zero value otherwise.

### GetAlwaysMapResponsesOk

`func (o *GetPlugin200Response) GetAlwaysMapResponsesOk() (*bool, bool)`

GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysMapResponses

`func (o *GetPlugin200Response) SetAlwaysMapResponses(v bool)`

SetAlwaysMapResponses sets AlwaysMapResponses field to given value.


### GetRetainFilesSparselyByAge

`func (o *GetPlugin200Response) GetRetainFilesSparselyByAge() bool`

GetRetainFilesSparselyByAge returns the RetainFilesSparselyByAge field if non-nil, zero value otherwise.

### GetRetainFilesSparselyByAgeOk

`func (o *GetPlugin200Response) GetRetainFilesSparselyByAgeOk() (*bool, bool)`

GetRetainFilesSparselyByAgeOk returns a tuple with the RetainFilesSparselyByAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFilesSparselyByAge

`func (o *GetPlugin200Response) SetRetainFilesSparselyByAge(v bool)`

SetRetainFilesSparselyByAge sets RetainFilesSparselyByAge field to given value.

### HasRetainFilesSparselyByAge

`func (o *GetPlugin200Response) HasRetainFilesSparselyByAge() bool`

HasRetainFilesSparselyByAge returns a boolean if a field has been set.

### GetSanitize

`func (o *GetPlugin200Response) GetSanitize() bool`

GetSanitize returns the Sanitize field if non-nil, zero value otherwise.

### GetSanitizeOk

`func (o *GetPlugin200Response) GetSanitizeOk() (*bool, bool)`

GetSanitizeOk returns a tuple with the Sanitize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitize

`func (o *GetPlugin200Response) SetSanitize(v bool)`

SetSanitize sets Sanitize field to given value.

### HasSanitize

`func (o *GetPlugin200Response) HasSanitize() bool`

HasSanitize returns a boolean if a field has been set.

### GetReferralBaseURL

`func (o *GetPlugin200Response) GetReferralBaseURL() []string`

GetReferralBaseURL returns the ReferralBaseURL field if non-nil, zero value otherwise.

### GetReferralBaseURLOk

`func (o *GetPlugin200Response) GetReferralBaseURLOk() (*[]string, bool)`

GetReferralBaseURLOk returns a tuple with the ReferralBaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralBaseURL

`func (o *GetPlugin200Response) SetReferralBaseURL(v []string)`

SetReferralBaseURL sets ReferralBaseURL field to given value.


### GetContextName

`func (o *GetPlugin200Response) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *GetPlugin200Response) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *GetPlugin200Response) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *GetPlugin200Response) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### GetAgentxAddress

`func (o *GetPlugin200Response) GetAgentxAddress() string`

GetAgentxAddress returns the AgentxAddress field if non-nil, zero value otherwise.

### GetAgentxAddressOk

`func (o *GetPlugin200Response) GetAgentxAddressOk() (*string, bool)`

GetAgentxAddressOk returns a tuple with the AgentxAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxAddress

`func (o *GetPlugin200Response) SetAgentxAddress(v string)`

SetAgentxAddress sets AgentxAddress field to given value.


### GetAgentxPort

`func (o *GetPlugin200Response) GetAgentxPort() int64`

GetAgentxPort returns the AgentxPort field if non-nil, zero value otherwise.

### GetAgentxPortOk

`func (o *GetPlugin200Response) GetAgentxPortOk() (*int64, bool)`

GetAgentxPortOk returns a tuple with the AgentxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxPort

`func (o *GetPlugin200Response) SetAgentxPort(v int64)`

SetAgentxPort sets AgentxPort field to given value.


### GetNumWorkerThreads

`func (o *GetPlugin200Response) GetNumWorkerThreads() int64`

GetNumWorkerThreads returns the NumWorkerThreads field if non-nil, zero value otherwise.

### GetNumWorkerThreadsOk

`func (o *GetPlugin200Response) GetNumWorkerThreadsOk() (*int64, bool)`

GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkerThreads

`func (o *GetPlugin200Response) SetNumWorkerThreads(v int64)`

SetNumWorkerThreads sets NumWorkerThreads field to given value.

### HasNumWorkerThreads

`func (o *GetPlugin200Response) HasNumWorkerThreads() bool`

HasNumWorkerThreads returns a boolean if a field has been set.

### GetSessionTimeout

`func (o *GetPlugin200Response) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *GetPlugin200Response) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *GetPlugin200Response) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *GetPlugin200Response) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetConnectRetryMaxWait

`func (o *GetPlugin200Response) GetConnectRetryMaxWait() string`

GetConnectRetryMaxWait returns the ConnectRetryMaxWait field if non-nil, zero value otherwise.

### GetConnectRetryMaxWaitOk

`func (o *GetPlugin200Response) GetConnectRetryMaxWaitOk() (*string, bool)`

GetConnectRetryMaxWaitOk returns a tuple with the ConnectRetryMaxWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectRetryMaxWait

`func (o *GetPlugin200Response) SetConnectRetryMaxWait(v string)`

SetConnectRetryMaxWait sets ConnectRetryMaxWait field to given value.

### HasConnectRetryMaxWait

`func (o *GetPlugin200Response) HasConnectRetryMaxWait() bool`

HasConnectRetryMaxWait returns a boolean if a field has been set.

### GetPingInterval

`func (o *GetPlugin200Response) GetPingInterval() string`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *GetPlugin200Response) GetPingIntervalOk() (*string, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *GetPlugin200Response) SetPingInterval(v string)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *GetPlugin200Response) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### GetAllowedRequestControl

`func (o *GetPlugin200Response) GetAllowedRequestControl() []string`

GetAllowedRequestControl returns the AllowedRequestControl field if non-nil, zero value otherwise.

### GetAllowedRequestControlOk

`func (o *GetPlugin200Response) GetAllowedRequestControlOk() (*[]string, bool)`

GetAllowedRequestControlOk returns a tuple with the AllowedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequestControl

`func (o *GetPlugin200Response) SetAllowedRequestControl(v []string)`

SetAllowedRequestControl sets AllowedRequestControl field to given value.

### HasAllowedRequestControl

`func (o *GetPlugin200Response) HasAllowedRequestControl() bool`

HasAllowedRequestControl returns a boolean if a field has been set.

### GetDefaultUserPasswordStorageScheme

`func (o *GetPlugin200Response) GetDefaultUserPasswordStorageScheme() []string`

GetDefaultUserPasswordStorageScheme returns the DefaultUserPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDefaultUserPasswordStorageSchemeOk

`func (o *GetPlugin200Response) GetDefaultUserPasswordStorageSchemeOk() (*[]string, bool)`

GetDefaultUserPasswordStorageSchemeOk returns a tuple with the DefaultUserPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserPasswordStorageScheme

`func (o *GetPlugin200Response) SetDefaultUserPasswordStorageScheme(v []string)`

SetDefaultUserPasswordStorageScheme sets DefaultUserPasswordStorageScheme field to given value.

### HasDefaultUserPasswordStorageScheme

`func (o *GetPlugin200Response) HasDefaultUserPasswordStorageScheme() bool`

HasDefaultUserPasswordStorageScheme returns a boolean if a field has been set.

### GetDefaultAuthPasswordStorageScheme

`func (o *GetPlugin200Response) GetDefaultAuthPasswordStorageScheme() []string`

GetDefaultAuthPasswordStorageScheme returns the DefaultAuthPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDefaultAuthPasswordStorageSchemeOk

`func (o *GetPlugin200Response) GetDefaultAuthPasswordStorageSchemeOk() (*[]string, bool)`

GetDefaultAuthPasswordStorageSchemeOk returns a tuple with the DefaultAuthPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuthPasswordStorageScheme

`func (o *GetPlugin200Response) SetDefaultAuthPasswordStorageScheme(v []string)`

SetDefaultAuthPasswordStorageScheme sets DefaultAuthPasswordStorageScheme field to given value.

### HasDefaultAuthPasswordStorageScheme

`func (o *GetPlugin200Response) HasDefaultAuthPasswordStorageScheme() bool`

HasDefaultAuthPasswordStorageScheme returns a boolean if a field has been set.

### GetProfileSampleInterval

`func (o *GetPlugin200Response) GetProfileSampleInterval() string`

GetProfileSampleInterval returns the ProfileSampleInterval field if non-nil, zero value otherwise.

### GetProfileSampleIntervalOk

`func (o *GetPlugin200Response) GetProfileSampleIntervalOk() (*string, bool)`

GetProfileSampleIntervalOk returns a tuple with the ProfileSampleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileSampleInterval

`func (o *GetPlugin200Response) SetProfileSampleInterval(v string)`

SetProfileSampleInterval sets ProfileSampleInterval field to given value.


### GetProfileDirectory

`func (o *GetPlugin200Response) GetProfileDirectory() string`

GetProfileDirectory returns the ProfileDirectory field if non-nil, zero value otherwise.

### GetProfileDirectoryOk

`func (o *GetPlugin200Response) GetProfileDirectoryOk() (*string, bool)`

GetProfileDirectoryOk returns a tuple with the ProfileDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDirectory

`func (o *GetPlugin200Response) SetProfileDirectory(v string)`

SetProfileDirectory sets ProfileDirectory field to given value.


### GetEnableProfilingOnStartup

`func (o *GetPlugin200Response) GetEnableProfilingOnStartup() bool`

GetEnableProfilingOnStartup returns the EnableProfilingOnStartup field if non-nil, zero value otherwise.

### GetEnableProfilingOnStartupOk

`func (o *GetPlugin200Response) GetEnableProfilingOnStartupOk() (*bool, bool)`

GetEnableProfilingOnStartupOk returns a tuple with the EnableProfilingOnStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProfilingOnStartup

`func (o *GetPlugin200Response) SetEnableProfilingOnStartup(v bool)`

SetEnableProfilingOnStartup sets EnableProfilingOnStartup field to given value.


### GetProfileAction

`func (o *GetPlugin200Response) GetProfileAction() EnumpluginProfileActionProp`

GetProfileAction returns the ProfileAction field if non-nil, zero value otherwise.

### GetProfileActionOk

`func (o *GetPlugin200Response) GetProfileActionOk() (*EnumpluginProfileActionProp, bool)`

GetProfileActionOk returns a tuple with the ProfileAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAction

`func (o *GetPlugin200Response) SetProfileAction(v EnumpluginProfileActionProp)`

SetProfileAction sets ProfileAction field to given value.

### HasProfileAction

`func (o *GetPlugin200Response) HasProfileAction() bool`

HasProfileAction returns a boolean if a field has been set.

### GetValuePattern

`func (o *GetPlugin200Response) GetValuePattern() []string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *GetPlugin200Response) GetValuePatternOk() (*[]string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *GetPlugin200Response) SetValuePattern(v []string)`

SetValuePattern sets ValuePattern field to given value.


### GetMultipleValuePatternBehavior

`func (o *GetPlugin200Response) GetMultipleValuePatternBehavior() EnumpluginMultipleValuePatternBehaviorProp`

GetMultipleValuePatternBehavior returns the MultipleValuePatternBehavior field if non-nil, zero value otherwise.

### GetMultipleValuePatternBehaviorOk

`func (o *GetPlugin200Response) GetMultipleValuePatternBehaviorOk() (*EnumpluginMultipleValuePatternBehaviorProp, bool)`

GetMultipleValuePatternBehaviorOk returns a tuple with the MultipleValuePatternBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleValuePatternBehavior

`func (o *GetPlugin200Response) SetMultipleValuePatternBehavior(v EnumpluginMultipleValuePatternBehaviorProp)`

SetMultipleValuePatternBehavior sets MultipleValuePatternBehavior field to given value.

### HasMultipleValuePatternBehavior

`func (o *GetPlugin200Response) HasMultipleValuePatternBehavior() bool`

HasMultipleValuePatternBehavior returns a boolean if a field has been set.

### GetMultiValuedAttributeBehavior

`func (o *GetPlugin200Response) GetMultiValuedAttributeBehavior() EnumpluginMultiValuedAttributeBehaviorProp`

GetMultiValuedAttributeBehavior returns the MultiValuedAttributeBehavior field if non-nil, zero value otherwise.

### GetMultiValuedAttributeBehaviorOk

`func (o *GetPlugin200Response) GetMultiValuedAttributeBehaviorOk() (*EnumpluginMultiValuedAttributeBehaviorProp, bool)`

GetMultiValuedAttributeBehaviorOk returns a tuple with the MultiValuedAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValuedAttributeBehavior

`func (o *GetPlugin200Response) SetMultiValuedAttributeBehavior(v EnumpluginMultiValuedAttributeBehaviorProp)`

SetMultiValuedAttributeBehavior sets MultiValuedAttributeBehavior field to given value.

### HasMultiValuedAttributeBehavior

`func (o *GetPlugin200Response) HasMultiValuedAttributeBehavior() bool`

HasMultiValuedAttributeBehavior returns a boolean if a field has been set.

### GetTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *GetPlugin200Response) GetTargetAttributeExistsDuringInitialPopulationBehavior() EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp`

GetTargetAttributeExistsDuringInitialPopulationBehavior returns the TargetAttributeExistsDuringInitialPopulationBehavior field if non-nil, zero value otherwise.

### GetTargetAttributeExistsDuringInitialPopulationBehaviorOk

`func (o *GetPlugin200Response) GetTargetAttributeExistsDuringInitialPopulationBehaviorOk() (*EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp, bool)`

GetTargetAttributeExistsDuringInitialPopulationBehaviorOk returns a tuple with the TargetAttributeExistsDuringInitialPopulationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *GetPlugin200Response) SetTargetAttributeExistsDuringInitialPopulationBehavior(v EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp)`

SetTargetAttributeExistsDuringInitialPopulationBehavior sets TargetAttributeExistsDuringInitialPopulationBehavior field to given value.

### HasTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *GetPlugin200Response) HasTargetAttributeExistsDuringInitialPopulationBehavior() bool`

HasTargetAttributeExistsDuringInitialPopulationBehavior returns a boolean if a field has been set.

### GetUpdateSourceAttributeBehavior

`func (o *GetPlugin200Response) GetUpdateSourceAttributeBehavior() EnumpluginUpdateSourceAttributeBehaviorProp`

GetUpdateSourceAttributeBehavior returns the UpdateSourceAttributeBehavior field if non-nil, zero value otherwise.

### GetUpdateSourceAttributeBehaviorOk

`func (o *GetPlugin200Response) GetUpdateSourceAttributeBehaviorOk() (*EnumpluginUpdateSourceAttributeBehaviorProp, bool)`

GetUpdateSourceAttributeBehaviorOk returns a tuple with the UpdateSourceAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSourceAttributeBehavior

`func (o *GetPlugin200Response) SetUpdateSourceAttributeBehavior(v EnumpluginUpdateSourceAttributeBehaviorProp)`

SetUpdateSourceAttributeBehavior sets UpdateSourceAttributeBehavior field to given value.

### HasUpdateSourceAttributeBehavior

`func (o *GetPlugin200Response) HasUpdateSourceAttributeBehavior() bool`

HasUpdateSourceAttributeBehavior returns a boolean if a field has been set.

### GetSourceAttributeRemovalBehavior

`func (o *GetPlugin200Response) GetSourceAttributeRemovalBehavior() EnumpluginSourceAttributeRemovalBehaviorProp`

GetSourceAttributeRemovalBehavior returns the SourceAttributeRemovalBehavior field if non-nil, zero value otherwise.

### GetSourceAttributeRemovalBehaviorOk

`func (o *GetPlugin200Response) GetSourceAttributeRemovalBehaviorOk() (*EnumpluginSourceAttributeRemovalBehaviorProp, bool)`

GetSourceAttributeRemovalBehaviorOk returns a tuple with the SourceAttributeRemovalBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttributeRemovalBehavior

`func (o *GetPlugin200Response) SetSourceAttributeRemovalBehavior(v EnumpluginSourceAttributeRemovalBehaviorProp)`

SetSourceAttributeRemovalBehavior sets SourceAttributeRemovalBehavior field to given value.

### HasSourceAttributeRemovalBehavior

`func (o *GetPlugin200Response) HasSourceAttributeRemovalBehavior() bool`

HasSourceAttributeRemovalBehavior returns a boolean if a field has been set.

### GetUpdateTargetAttributeBehavior

`func (o *GetPlugin200Response) GetUpdateTargetAttributeBehavior() EnumpluginUpdateTargetAttributeBehaviorProp`

GetUpdateTargetAttributeBehavior returns the UpdateTargetAttributeBehavior field if non-nil, zero value otherwise.

### GetUpdateTargetAttributeBehaviorOk

`func (o *GetPlugin200Response) GetUpdateTargetAttributeBehaviorOk() (*EnumpluginUpdateTargetAttributeBehaviorProp, bool)`

GetUpdateTargetAttributeBehaviorOk returns a tuple with the UpdateTargetAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTargetAttributeBehavior

`func (o *GetPlugin200Response) SetUpdateTargetAttributeBehavior(v EnumpluginUpdateTargetAttributeBehaviorProp)`

SetUpdateTargetAttributeBehavior sets UpdateTargetAttributeBehavior field to given value.

### HasUpdateTargetAttributeBehavior

`func (o *GetPlugin200Response) HasUpdateTargetAttributeBehavior() bool`

HasUpdateTargetAttributeBehavior returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *GetPlugin200Response) GetIncludeBaseDN() []string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *GetPlugin200Response) GetIncludeBaseDNOk() (*[]string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *GetPlugin200Response) SetIncludeBaseDN(v []string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *GetPlugin200Response) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetExcludeBaseDN

`func (o *GetPlugin200Response) GetExcludeBaseDN() []string`

GetExcludeBaseDN returns the ExcludeBaseDN field if non-nil, zero value otherwise.

### GetExcludeBaseDNOk

`func (o *GetPlugin200Response) GetExcludeBaseDNOk() (*[]string, bool)`

GetExcludeBaseDNOk returns a tuple with the ExcludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBaseDN

`func (o *GetPlugin200Response) SetExcludeBaseDN(v []string)`

SetExcludeBaseDN sets ExcludeBaseDN field to given value.

### HasExcludeBaseDN

`func (o *GetPlugin200Response) HasExcludeBaseDN() bool`

HasExcludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *GetPlugin200Response) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *GetPlugin200Response) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *GetPlugin200Response) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *GetPlugin200Response) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *GetPlugin200Response) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *GetPlugin200Response) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *GetPlugin200Response) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *GetPlugin200Response) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *GetPlugin200Response) GetUpdatedEntryNewlyMatchesCriteriaBehavior() EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp`

GetUpdatedEntryNewlyMatchesCriteriaBehavior returns the UpdatedEntryNewlyMatchesCriteriaBehavior field if non-nil, zero value otherwise.

### GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk

`func (o *GetPlugin200Response) GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp, bool)`

GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNewlyMatchesCriteriaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *GetPlugin200Response) SetUpdatedEntryNewlyMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp)`

SetUpdatedEntryNewlyMatchesCriteriaBehavior sets UpdatedEntryNewlyMatchesCriteriaBehavior field to given value.

### HasUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *GetPlugin200Response) HasUpdatedEntryNewlyMatchesCriteriaBehavior() bool`

HasUpdatedEntryNewlyMatchesCriteriaBehavior returns a boolean if a field has been set.

### GetUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *GetPlugin200Response) GetUpdatedEntryNoLongerMatchesCriteriaBehavior() EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp`

GetUpdatedEntryNoLongerMatchesCriteriaBehavior returns the UpdatedEntryNoLongerMatchesCriteriaBehavior field if non-nil, zero value otherwise.

### GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk

`func (o *GetPlugin200Response) GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp, bool)`

GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNoLongerMatchesCriteriaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *GetPlugin200Response) SetUpdatedEntryNoLongerMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp)`

SetUpdatedEntryNoLongerMatchesCriteriaBehavior sets UpdatedEntryNoLongerMatchesCriteriaBehavior field to given value.

### HasUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *GetPlugin200Response) HasUpdatedEntryNoLongerMatchesCriteriaBehavior() bool`

HasUpdatedEntryNoLongerMatchesCriteriaBehavior returns a boolean if a field has been set.

### GetSourceAttribute

`func (o *GetPlugin200Response) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *GetPlugin200Response) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *GetPlugin200Response) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *GetPlugin200Response) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *GetPlugin200Response) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *GetPlugin200Response) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetDelay

`func (o *GetPlugin200Response) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *GetPlugin200Response) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *GetPlugin200Response) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetScriptClass

`func (o *GetPlugin200Response) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *GetPlugin200Response) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *GetPlugin200Response) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *GetPlugin200Response) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *GetPlugin200Response) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *GetPlugin200Response) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *GetPlugin200Response) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetExcludeAttribute

`func (o *GetPlugin200Response) GetExcludeAttribute() []string`

GetExcludeAttribute returns the ExcludeAttribute field if non-nil, zero value otherwise.

### GetExcludeAttributeOk

`func (o *GetPlugin200Response) GetExcludeAttributeOk() (*[]string, bool)`

GetExcludeAttributeOk returns a tuple with the ExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttribute

`func (o *GetPlugin200Response) SetExcludeAttribute(v []string)`

SetExcludeAttribute sets ExcludeAttribute field to given value.

### HasExcludeAttribute

`func (o *GetPlugin200Response) HasExcludeAttribute() bool`

HasExcludeAttribute returns a boolean if a field has been set.

### GetPassThroughAuthenticationHandler

`func (o *GetPlugin200Response) GetPassThroughAuthenticationHandler() string`

GetPassThroughAuthenticationHandler returns the PassThroughAuthenticationHandler field if non-nil, zero value otherwise.

### GetPassThroughAuthenticationHandlerOk

`func (o *GetPlugin200Response) GetPassThroughAuthenticationHandlerOk() (*string, bool)`

GetPassThroughAuthenticationHandlerOk returns a tuple with the PassThroughAuthenticationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThroughAuthenticationHandler

`func (o *GetPlugin200Response) SetPassThroughAuthenticationHandler(v string)`

SetPassThroughAuthenticationHandler sets PassThroughAuthenticationHandler field to given value.


### GetUpdateInterval

`func (o *GetPlugin200Response) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *GetPlugin200Response) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *GetPlugin200Response) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *GetPlugin200Response) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetType

`func (o *GetPlugin200Response) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPlugin200Response) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPlugin200Response) SetType(v []string)`

SetType sets Type field to given value.


### GetMultipleAttributeBehavior

`func (o *GetPlugin200Response) GetMultipleAttributeBehavior() EnumpluginUniqueAttributeMultipleAttributeBehaviorProp`

GetMultipleAttributeBehavior returns the MultipleAttributeBehavior field if non-nil, zero value otherwise.

### GetMultipleAttributeBehaviorOk

`func (o *GetPlugin200Response) GetMultipleAttributeBehaviorOk() (*EnumpluginUniqueAttributeMultipleAttributeBehaviorProp, bool)`

GetMultipleAttributeBehaviorOk returns a tuple with the MultipleAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAttributeBehavior

`func (o *GetPlugin200Response) SetMultipleAttributeBehavior(v EnumpluginUniqueAttributeMultipleAttributeBehaviorProp)`

SetMultipleAttributeBehavior sets MultipleAttributeBehavior field to given value.

### HasMultipleAttributeBehavior

`func (o *GetPlugin200Response) HasMultipleAttributeBehavior() bool`

HasMultipleAttributeBehavior returns a boolean if a field has been set.

### GetPreventConflictsWithSoftDeletedEntries

`func (o *GetPlugin200Response) GetPreventConflictsWithSoftDeletedEntries() bool`

GetPreventConflictsWithSoftDeletedEntries returns the PreventConflictsWithSoftDeletedEntries field if non-nil, zero value otherwise.

### GetPreventConflictsWithSoftDeletedEntriesOk

`func (o *GetPlugin200Response) GetPreventConflictsWithSoftDeletedEntriesOk() (*bool, bool)`

GetPreventConflictsWithSoftDeletedEntriesOk returns a tuple with the PreventConflictsWithSoftDeletedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventConflictsWithSoftDeletedEntries

`func (o *GetPlugin200Response) SetPreventConflictsWithSoftDeletedEntries(v bool)`

SetPreventConflictsWithSoftDeletedEntries sets PreventConflictsWithSoftDeletedEntries field to given value.

### HasPreventConflictsWithSoftDeletedEntries

`func (o *GetPlugin200Response) HasPreventConflictsWithSoftDeletedEntries() bool`

HasPreventConflictsWithSoftDeletedEntries returns a boolean if a field has been set.

### GetPreventAddingMembersToNonexistentGroups

`func (o *GetPlugin200Response) GetPreventAddingMembersToNonexistentGroups() bool`

GetPreventAddingMembersToNonexistentGroups returns the PreventAddingMembersToNonexistentGroups field if non-nil, zero value otherwise.

### GetPreventAddingMembersToNonexistentGroupsOk

`func (o *GetPlugin200Response) GetPreventAddingMembersToNonexistentGroupsOk() (*bool, bool)`

GetPreventAddingMembersToNonexistentGroupsOk returns a tuple with the PreventAddingMembersToNonexistentGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAddingMembersToNonexistentGroups

`func (o *GetPlugin200Response) SetPreventAddingMembersToNonexistentGroups(v bool)`

SetPreventAddingMembersToNonexistentGroups sets PreventAddingMembersToNonexistentGroups field to given value.

### HasPreventAddingMembersToNonexistentGroups

`func (o *GetPlugin200Response) HasPreventAddingMembersToNonexistentGroups() bool`

HasPreventAddingMembersToNonexistentGroups returns a boolean if a field has been set.

### GetPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *GetPlugin200Response) GetPreventAddingGroupsAsInvertedStaticGroupMembers() bool`

GetPreventAddingGroupsAsInvertedStaticGroupMembers returns the PreventAddingGroupsAsInvertedStaticGroupMembers field if non-nil, zero value otherwise.

### GetPreventAddingGroupsAsInvertedStaticGroupMembersOk

`func (o *GetPlugin200Response) GetPreventAddingGroupsAsInvertedStaticGroupMembersOk() (*bool, bool)`

GetPreventAddingGroupsAsInvertedStaticGroupMembersOk returns a tuple with the PreventAddingGroupsAsInvertedStaticGroupMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *GetPlugin200Response) SetPreventAddingGroupsAsInvertedStaticGroupMembers(v bool)`

SetPreventAddingGroupsAsInvertedStaticGroupMembers sets PreventAddingGroupsAsInvertedStaticGroupMembers field to given value.

### HasPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *GetPlugin200Response) HasPreventAddingGroupsAsInvertedStaticGroupMembers() bool`

HasPreventAddingGroupsAsInvertedStaticGroupMembers returns a boolean if a field has been set.

### GetPreventNestingNonexistentGroups

`func (o *GetPlugin200Response) GetPreventNestingNonexistentGroups() bool`

GetPreventNestingNonexistentGroups returns the PreventNestingNonexistentGroups field if non-nil, zero value otherwise.

### GetPreventNestingNonexistentGroupsOk

`func (o *GetPlugin200Response) GetPreventNestingNonexistentGroupsOk() (*bool, bool)`

GetPreventNestingNonexistentGroupsOk returns a tuple with the PreventNestingNonexistentGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventNestingNonexistentGroups

`func (o *GetPlugin200Response) SetPreventNestingNonexistentGroups(v bool)`

SetPreventNestingNonexistentGroups sets PreventNestingNonexistentGroups field to given value.

### HasPreventNestingNonexistentGroups

`func (o *GetPlugin200Response) HasPreventNestingNonexistentGroups() bool`

HasPreventNestingNonexistentGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


