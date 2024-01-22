# PluginListResponseResourcesInner

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
**Type** | **[]string** | Specifies the type of attributes to check for value uniqueness. | 
**MultipleAttributeBehavior** | Pointer to [**EnumpluginUniqueAttributeMultipleAttributeBehaviorProp**](EnumpluginUniqueAttributeMultipleAttributeBehaviorProp.md) |  | [optional] 
**SubtreeView** | **[]string** | The subtree view(s) for which to maintain referential integrity. | 
**PreventConflictsWithSoftDeletedEntries** | Pointer to **bool** | Indicates whether this Unique Attribute Plugin should reject a change that would result in one or more conflicts, even if those conflicts only exist in soft-deleted entries. | [optional] 
**PreCommitValidation** | Pointer to [**EnumpluginPreCommitValidationProp**](EnumpluginPreCommitValidationProp.md) |  | [optional] 
**PostCommitValidation** | Pointer to [**EnumpluginPostCommitValidationProp**](EnumpluginPostCommitValidationProp.md) |  | [optional] 
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
**NumWorkerThreads** | **int64** | The number of worker threads to use to handle SNMP requests. | 
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
**ListenAddress** | **string** | The IP address on which the SNMP agent will listen for client requests. | 
**ListenPort** | **int64** | The port number on which the SNMP agent will listen for client requests. | 
**AgentxTransport** | [**EnumpluginAgentxTransportProp**](EnumpluginAgentxTransportProp.md) |  | 
**AgentxListenAddress** | **string** | The IP address on which the SNMP agent will listen for sub-agent sessions. | 
**AgentxListenPort** | **int64** | The port number on which the SNMP agent will listen for sub-agent sessions. | 
**NotificationTargetAddress** | Pointer to **string** | The IP address of the target to which the SNMP agent should send notifications (traps). | [optional] 
**NotificationTargetPort** | Pointer to **int64** | The port number of the target to which the SNMP agent should send notifications (traps). | [optional] 
**AgentSNMPVersion** | [**[]EnumpluginAgentSNMPVersionProp**](EnumpluginAgentSNMPVersionProp.md) |  | 
**CommunityName** | **string** | The name of the community to use for the SNMP agent. | 
**AgentSecurityName** | Pointer to **string** | The security name (i.e., username) to use for the SNMP agent. This must be defined if SNMPv3 is to be used. | [optional] 
**AgentSecurityLevel** | Pointer to [**EnumpluginAgentSecurityLevelProp**](EnumpluginAgentSecurityLevelProp.md) |  | [optional] 
**AgentAuthenticationProtocol** | Pointer to [**EnumpluginAgentAuthenticationProtocolProp**](EnumpluginAgentAuthenticationProtocolProp.md) |  | [optional] 
**AgentAuthenticationPassphrase** | Pointer to **string** | The authentication passphrase to use for SNMPv3 if authentication is to be performed. | [optional] 
**AgentPrivacyProtocol** | Pointer to [**EnumpluginAgentPrivacyProtocolProp**](EnumpluginAgentPrivacyProtocolProp.md) |  | [optional] 
**AgentPrivacyPassphrase** | Pointer to **string** | The privacy passphrase to use for SNMPv3 if privacy is to be used. | [optional] 
**PreventAddingMembersToNonexistentGroups** | Pointer to **bool** | Indicates whether the server should prevent updates to user entries that attempt to add them as a member of an inverted static group that does not exist. | [optional] 
**PreventAddingGroupsAsInvertedStaticGroupMembers** | Pointer to **bool** | Indicates whether the server should prevent attempts to add a group as a regular member of an inverted static group. If the members of another group should be considered members of an inverted static group, then the other group should be added as a nested group rather than a regular member. | [optional] 
**PreventNestingNonexistentGroups** | Pointer to **bool** | Indicates whether the server should prevent updates to inverted static groups that add references to nested groups that don&#39;t exist. | [optional] 

## Methods

### NewPluginListResponseResourcesInner

`func NewPluginListResponseResourcesInner(schemas []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn, id string, requestCriteria string, enabled bool, sampleInterval string, collectionInterval string, pluginType []EnumpluginPluginTypeProp, numThreads int64, baseDN []string, filterPrefix string, filter string, attributeType []string, pollingInterval string, maxUpdatesPerSecond int64, numDeleteThreads int64, invokeGCTimeUtc []string, apiURL string, authURL string, oAuthClientID string, environmentID string, tryLocalBind bool, overrideLocalPassword bool, updateLocalPassword bool, userMappingLocalAttribute []string, userMappingRemoteJSONField []string, type_ []string, subtreeView []string, histogramCategoryBoundary []string, scope EnumpluginScopeProp, outputFile string, logInterval string, suppressIfIdle bool, linesBetweenHeader int64, histogramFormat EnumpluginHistogramFormatProp, logFile string, logFilePermissions string, rotationPolicy []string, retentionPolicy []string, datetimeAttribute string, datetimeFormat EnumpluginDatetimeFormatProp, expirationOffset string, numMostExpensivePhasesShown int64, extensionClass string, server []string, serverAccessMode EnumpluginServerAccessModeProp, initialConnections int64, maxConnections int64, sourceDN string, targetDN string, enableAttributeMapping bool, enableControlMapping bool, alwaysMapResponses bool, referralBaseURL []string, agentxAddress string, agentxPort int64, numWorkerThreads int64, profileSampleInterval string, profileDirectory string, enableProfilingOnStartup bool, valuePattern []string, sourceAttribute string, targetAttribute string, delay string, scriptClass string, passThroughAuthenticationHandler string, listenAddress string, listenPort int64, agentxTransport EnumpluginAgentxTransportProp, agentxListenAddress string, agentxListenPort int64, agentSNMPVersion []EnumpluginAgentSNMPVersionProp, communityName string, ) *PluginListResponseResourcesInner`

NewPluginListResponseResourcesInner instantiates a new PluginListResponseResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginListResponseResourcesInnerWithDefaults

`func NewPluginListResponseResourcesInnerWithDefaults() *PluginListResponseResourcesInner`

NewPluginListResponseResourcesInnerWithDefaults instantiates a new PluginListResponseResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *PluginListResponseResourcesInner) GetSchemas() []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *PluginListResponseResourcesInner) GetSchemasOk() (*[]EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *PluginListResponseResourcesInner) SetSchemas(v []EnuminvertedStaticGroupReferentialIntegrityPluginSchemaUrn)`

SetSchemas sets Schemas field to given value.


### GetId

`func (o *PluginListResponseResourcesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PluginListResponseResourcesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PluginListResponseResourcesInner) SetId(v string)`

SetId sets Id field to given value.


### GetMaxUpdateFrequency

`func (o *PluginListResponseResourcesInner) GetMaxUpdateFrequency() string`

GetMaxUpdateFrequency returns the MaxUpdateFrequency field if non-nil, zero value otherwise.

### GetMaxUpdateFrequencyOk

`func (o *PluginListResponseResourcesInner) GetMaxUpdateFrequencyOk() (*string, bool)`

GetMaxUpdateFrequencyOk returns a tuple with the MaxUpdateFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdateFrequency

`func (o *PluginListResponseResourcesInner) SetMaxUpdateFrequency(v string)`

SetMaxUpdateFrequency sets MaxUpdateFrequency field to given value.

### HasMaxUpdateFrequency

`func (o *PluginListResponseResourcesInner) HasMaxUpdateFrequency() bool`

HasMaxUpdateFrequency returns a boolean if a field has been set.

### GetOperationType

`func (o *PluginListResponseResourcesInner) GetOperationType() []EnumpluginOperationTypeProp`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PluginListResponseResourcesInner) GetOperationTypeOk() (*[]EnumpluginOperationTypeProp, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PluginListResponseResourcesInner) SetOperationType(v []EnumpluginOperationTypeProp)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *PluginListResponseResourcesInner) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetInvokeForFailedBinds

`func (o *PluginListResponseResourcesInner) GetInvokeForFailedBinds() bool`

GetInvokeForFailedBinds returns the InvokeForFailedBinds field if non-nil, zero value otherwise.

### GetInvokeForFailedBindsOk

`func (o *PluginListResponseResourcesInner) GetInvokeForFailedBindsOk() (*bool, bool)`

GetInvokeForFailedBindsOk returns a tuple with the InvokeForFailedBinds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForFailedBinds

`func (o *PluginListResponseResourcesInner) SetInvokeForFailedBinds(v bool)`

SetInvokeForFailedBinds sets InvokeForFailedBinds field to given value.

### HasInvokeForFailedBinds

`func (o *PluginListResponseResourcesInner) HasInvokeForFailedBinds() bool`

HasInvokeForFailedBinds returns a boolean if a field has been set.

### GetMaxSearchResultEntriesToUpdate

`func (o *PluginListResponseResourcesInner) GetMaxSearchResultEntriesToUpdate() int64`

GetMaxSearchResultEntriesToUpdate returns the MaxSearchResultEntriesToUpdate field if non-nil, zero value otherwise.

### GetMaxSearchResultEntriesToUpdateOk

`func (o *PluginListResponseResourcesInner) GetMaxSearchResultEntriesToUpdateOk() (*int64, bool)`

GetMaxSearchResultEntriesToUpdateOk returns a tuple with the MaxSearchResultEntriesToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSearchResultEntriesToUpdate

`func (o *PluginListResponseResourcesInner) SetMaxSearchResultEntriesToUpdate(v int64)`

SetMaxSearchResultEntriesToUpdate sets MaxSearchResultEntriesToUpdate field to given value.

### HasMaxSearchResultEntriesToUpdate

`func (o *PluginListResponseResourcesInner) HasMaxSearchResultEntriesToUpdate() bool`

HasMaxSearchResultEntriesToUpdate returns a boolean if a field has been set.

### GetRequestCriteria

`func (o *PluginListResponseResourcesInner) GetRequestCriteria() string`

GetRequestCriteria returns the RequestCriteria field if non-nil, zero value otherwise.

### GetRequestCriteriaOk

`func (o *PluginListResponseResourcesInner) GetRequestCriteriaOk() (*string, bool)`

GetRequestCriteriaOk returns a tuple with the RequestCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCriteria

`func (o *PluginListResponseResourcesInner) SetRequestCriteria(v string)`

SetRequestCriteria sets RequestCriteria field to given value.


### GetInvokeForInternalOperations

`func (o *PluginListResponseResourcesInner) GetInvokeForInternalOperations() bool`

GetInvokeForInternalOperations returns the InvokeForInternalOperations field if non-nil, zero value otherwise.

### GetInvokeForInternalOperationsOk

`func (o *PluginListResponseResourcesInner) GetInvokeForInternalOperationsOk() (*bool, bool)`

GetInvokeForInternalOperationsOk returns a tuple with the InvokeForInternalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeForInternalOperations

`func (o *PluginListResponseResourcesInner) SetInvokeForInternalOperations(v bool)`

SetInvokeForInternalOperations sets InvokeForInternalOperations field to given value.

### HasInvokeForInternalOperations

`func (o *PluginListResponseResourcesInner) HasInvokeForInternalOperations() bool`

HasInvokeForInternalOperations returns a boolean if a field has been set.

### GetDescription

`func (o *PluginListResponseResourcesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PluginListResponseResourcesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PluginListResponseResourcesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PluginListResponseResourcesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *PluginListResponseResourcesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PluginListResponseResourcesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PluginListResponseResourcesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMeta

`func (o *PluginListResponseResourcesInner) GetMeta() MetaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *PluginListResponseResourcesInner) GetMetaOk() (*MetaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *PluginListResponseResourcesInner) SetMeta(v MetaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *PluginListResponseResourcesInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetUrnpingidentityschemasconfigurationmessages20

`func (o *PluginListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20() MetaUrnPingidentitySchemasConfigurationMessages20`

GetUrnpingidentityschemasconfigurationmessages20 returns the Urnpingidentityschemasconfigurationmessages20 field if non-nil, zero value otherwise.

### GetUrnpingidentityschemasconfigurationmessages20Ok

`func (o *PluginListResponseResourcesInner) GetUrnpingidentityschemasconfigurationmessages20Ok() (*MetaUrnPingidentitySchemasConfigurationMessages20, bool)`

GetUrnpingidentityschemasconfigurationmessages20Ok returns a tuple with the Urnpingidentityschemasconfigurationmessages20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrnpingidentityschemasconfigurationmessages20

`func (o *PluginListResponseResourcesInner) SetUrnpingidentityschemasconfigurationmessages20(v MetaUrnPingidentitySchemasConfigurationMessages20)`

SetUrnpingidentityschemasconfigurationmessages20 sets Urnpingidentityschemasconfigurationmessages20 field to given value.

### HasUrnpingidentityschemasconfigurationmessages20

`func (o *PluginListResponseResourcesInner) HasUrnpingidentityschemasconfigurationmessages20() bool`

HasUrnpingidentityschemasconfigurationmessages20 returns a boolean if a field has been set.

### GetSampleInterval

`func (o *PluginListResponseResourcesInner) GetSampleInterval() string`

GetSampleInterval returns the SampleInterval field if non-nil, zero value otherwise.

### GetSampleIntervalOk

`func (o *PluginListResponseResourcesInner) GetSampleIntervalOk() (*string, bool)`

GetSampleIntervalOk returns a tuple with the SampleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleInterval

`func (o *PluginListResponseResourcesInner) SetSampleInterval(v string)`

SetSampleInterval sets SampleInterval field to given value.


### GetCollectionInterval

`func (o *PluginListResponseResourcesInner) GetCollectionInterval() string`

GetCollectionInterval returns the CollectionInterval field if non-nil, zero value otherwise.

### GetCollectionIntervalOk

`func (o *PluginListResponseResourcesInner) GetCollectionIntervalOk() (*string, bool)`

GetCollectionIntervalOk returns a tuple with the CollectionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionInterval

`func (o *PluginListResponseResourcesInner) SetCollectionInterval(v string)`

SetCollectionInterval sets CollectionInterval field to given value.


### GetLdapInfo

`func (o *PluginListResponseResourcesInner) GetLdapInfo() EnumpluginLdapInfoProp`

GetLdapInfo returns the LdapInfo field if non-nil, zero value otherwise.

### GetLdapInfoOk

`func (o *PluginListResponseResourcesInner) GetLdapInfoOk() (*EnumpluginLdapInfoProp, bool)`

GetLdapInfoOk returns a tuple with the LdapInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapInfo

`func (o *PluginListResponseResourcesInner) SetLdapInfo(v EnumpluginLdapInfoProp)`

SetLdapInfo sets LdapInfo field to given value.

### HasLdapInfo

`func (o *PluginListResponseResourcesInner) HasLdapInfo() bool`

HasLdapInfo returns a boolean if a field has been set.

### GetServerInfo

`func (o *PluginListResponseResourcesInner) GetServerInfo() EnumpluginServerInfoProp`

GetServerInfo returns the ServerInfo field if non-nil, zero value otherwise.

### GetServerInfoOk

`func (o *PluginListResponseResourcesInner) GetServerInfoOk() (*EnumpluginServerInfoProp, bool)`

GetServerInfoOk returns a tuple with the ServerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInfo

`func (o *PluginListResponseResourcesInner) SetServerInfo(v EnumpluginServerInfoProp)`

SetServerInfo sets ServerInfo field to given value.

### HasServerInfo

`func (o *PluginListResponseResourcesInner) HasServerInfo() bool`

HasServerInfo returns a boolean if a field has been set.

### GetPerApplicationLDAPStats

`func (o *PluginListResponseResourcesInner) GetPerApplicationLDAPStats() EnumpluginPeriodicStatsLoggerPerApplicationLDAPStatsProp`

GetPerApplicationLDAPStats returns the PerApplicationLDAPStats field if non-nil, zero value otherwise.

### GetPerApplicationLDAPStatsOk

`func (o *PluginListResponseResourcesInner) GetPerApplicationLDAPStatsOk() (*EnumpluginPeriodicStatsLoggerPerApplicationLDAPStatsProp, bool)`

GetPerApplicationLDAPStatsOk returns a tuple with the PerApplicationLDAPStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerApplicationLDAPStats

`func (o *PluginListResponseResourcesInner) SetPerApplicationLDAPStats(v EnumpluginPeriodicStatsLoggerPerApplicationLDAPStatsProp)`

SetPerApplicationLDAPStats sets PerApplicationLDAPStats field to given value.

### HasPerApplicationLDAPStats

`func (o *PluginListResponseResourcesInner) HasPerApplicationLDAPStats() bool`

HasPerApplicationLDAPStats returns a boolean if a field has been set.

### GetLdapChangelogInfo

`func (o *PluginListResponseResourcesInner) GetLdapChangelogInfo() EnumpluginLdapChangelogInfoProp`

GetLdapChangelogInfo returns the LdapChangelogInfo field if non-nil, zero value otherwise.

### GetLdapChangelogInfoOk

`func (o *PluginListResponseResourcesInner) GetLdapChangelogInfoOk() (*EnumpluginLdapChangelogInfoProp, bool)`

GetLdapChangelogInfoOk returns a tuple with the LdapChangelogInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapChangelogInfo

`func (o *PluginListResponseResourcesInner) SetLdapChangelogInfo(v EnumpluginLdapChangelogInfoProp)`

SetLdapChangelogInfo sets LdapChangelogInfo field to given value.

### HasLdapChangelogInfo

`func (o *PluginListResponseResourcesInner) HasLdapChangelogInfo() bool`

HasLdapChangelogInfo returns a boolean if a field has been set.

### GetStatusSummaryInfo

`func (o *PluginListResponseResourcesInner) GetStatusSummaryInfo() EnumpluginStatusSummaryInfoProp`

GetStatusSummaryInfo returns the StatusSummaryInfo field if non-nil, zero value otherwise.

### GetStatusSummaryInfoOk

`func (o *PluginListResponseResourcesInner) GetStatusSummaryInfoOk() (*EnumpluginStatusSummaryInfoProp, bool)`

GetStatusSummaryInfoOk returns a tuple with the StatusSummaryInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSummaryInfo

`func (o *PluginListResponseResourcesInner) SetStatusSummaryInfo(v EnumpluginStatusSummaryInfoProp)`

SetStatusSummaryInfo sets StatusSummaryInfo field to given value.

### HasStatusSummaryInfo

`func (o *PluginListResponseResourcesInner) HasStatusSummaryInfo() bool`

HasStatusSummaryInfo returns a boolean if a field has been set.

### GetGenerateCollectorFiles

`func (o *PluginListResponseResourcesInner) GetGenerateCollectorFiles() bool`

GetGenerateCollectorFiles returns the GenerateCollectorFiles field if non-nil, zero value otherwise.

### GetGenerateCollectorFilesOk

`func (o *PluginListResponseResourcesInner) GetGenerateCollectorFilesOk() (*bool, bool)`

GetGenerateCollectorFilesOk returns a tuple with the GenerateCollectorFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateCollectorFiles

`func (o *PluginListResponseResourcesInner) SetGenerateCollectorFiles(v bool)`

SetGenerateCollectorFiles sets GenerateCollectorFiles field to given value.

### HasGenerateCollectorFiles

`func (o *PluginListResponseResourcesInner) HasGenerateCollectorFiles() bool`

HasGenerateCollectorFiles returns a boolean if a field has been set.

### GetLocalDBBackendInfo

`func (o *PluginListResponseResourcesInner) GetLocalDBBackendInfo() EnumpluginLocalDBBackendInfoProp`

GetLocalDBBackendInfo returns the LocalDBBackendInfo field if non-nil, zero value otherwise.

### GetLocalDBBackendInfoOk

`func (o *PluginListResponseResourcesInner) GetLocalDBBackendInfoOk() (*EnumpluginLocalDBBackendInfoProp, bool)`

GetLocalDBBackendInfoOk returns a tuple with the LocalDBBackendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDBBackendInfo

`func (o *PluginListResponseResourcesInner) SetLocalDBBackendInfo(v EnumpluginLocalDBBackendInfoProp)`

SetLocalDBBackendInfo sets LocalDBBackendInfo field to given value.

### HasLocalDBBackendInfo

`func (o *PluginListResponseResourcesInner) HasLocalDBBackendInfo() bool`

HasLocalDBBackendInfo returns a boolean if a field has been set.

### GetReplicationInfo

`func (o *PluginListResponseResourcesInner) GetReplicationInfo() EnumpluginReplicationInfoProp`

GetReplicationInfo returns the ReplicationInfo field if non-nil, zero value otherwise.

### GetReplicationInfoOk

`func (o *PluginListResponseResourcesInner) GetReplicationInfoOk() (*EnumpluginReplicationInfoProp, bool)`

GetReplicationInfoOk returns a tuple with the ReplicationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationInfo

`func (o *PluginListResponseResourcesInner) SetReplicationInfo(v EnumpluginReplicationInfoProp)`

SetReplicationInfo sets ReplicationInfo field to given value.

### HasReplicationInfo

`func (o *PluginListResponseResourcesInner) HasReplicationInfo() bool`

HasReplicationInfo returns a boolean if a field has been set.

### GetEntryCacheInfo

`func (o *PluginListResponseResourcesInner) GetEntryCacheInfo() EnumpluginEntryCacheInfoProp`

GetEntryCacheInfo returns the EntryCacheInfo field if non-nil, zero value otherwise.

### GetEntryCacheInfoOk

`func (o *PluginListResponseResourcesInner) GetEntryCacheInfoOk() (*EnumpluginEntryCacheInfoProp, bool)`

GetEntryCacheInfoOk returns a tuple with the EntryCacheInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCacheInfo

`func (o *PluginListResponseResourcesInner) SetEntryCacheInfo(v EnumpluginEntryCacheInfoProp)`

SetEntryCacheInfo sets EntryCacheInfo field to given value.

### HasEntryCacheInfo

`func (o *PluginListResponseResourcesInner) HasEntryCacheInfo() bool`

HasEntryCacheInfo returns a boolean if a field has been set.

### GetHostInfo

`func (o *PluginListResponseResourcesInner) GetHostInfo() []EnumpluginHostInfoProp`

GetHostInfo returns the HostInfo field if non-nil, zero value otherwise.

### GetHostInfoOk

`func (o *PluginListResponseResourcesInner) GetHostInfoOk() (*[]EnumpluginHostInfoProp, bool)`

GetHostInfoOk returns a tuple with the HostInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostInfo

`func (o *PluginListResponseResourcesInner) SetHostInfo(v []EnumpluginHostInfoProp)`

SetHostInfo sets HostInfo field to given value.

### HasHostInfo

`func (o *PluginListResponseResourcesInner) HasHostInfo() bool`

HasHostInfo returns a boolean if a field has been set.

### GetIncludedLDAPApplication

`func (o *PluginListResponseResourcesInner) GetIncludedLDAPApplication() []string`

GetIncludedLDAPApplication returns the IncludedLDAPApplication field if non-nil, zero value otherwise.

### GetIncludedLDAPApplicationOk

`func (o *PluginListResponseResourcesInner) GetIncludedLDAPApplicationOk() (*[]string, bool)`

GetIncludedLDAPApplicationOk returns a tuple with the IncludedLDAPApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPApplication

`func (o *PluginListResponseResourcesInner) SetIncludedLDAPApplication(v []string)`

SetIncludedLDAPApplication sets IncludedLDAPApplication field to given value.

### HasIncludedLDAPApplication

`func (o *PluginListResponseResourcesInner) HasIncludedLDAPApplication() bool`

HasIncludedLDAPApplication returns a boolean if a field has been set.

### GetTraditionalStaticGroupObjectClass

`func (o *PluginListResponseResourcesInner) GetTraditionalStaticGroupObjectClass() EnumpluginTraditionalStaticGroupObjectClassProp`

GetTraditionalStaticGroupObjectClass returns the TraditionalStaticGroupObjectClass field if non-nil, zero value otherwise.

### GetTraditionalStaticGroupObjectClassOk

`func (o *PluginListResponseResourcesInner) GetTraditionalStaticGroupObjectClassOk() (*EnumpluginTraditionalStaticGroupObjectClassProp, bool)`

GetTraditionalStaticGroupObjectClassOk returns a tuple with the TraditionalStaticGroupObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraditionalStaticGroupObjectClass

`func (o *PluginListResponseResourcesInner) SetTraditionalStaticGroupObjectClass(v EnumpluginTraditionalStaticGroupObjectClassProp)`

SetTraditionalStaticGroupObjectClass sets TraditionalStaticGroupObjectClass field to given value.

### HasTraditionalStaticGroupObjectClass

`func (o *PluginListResponseResourcesInner) HasTraditionalStaticGroupObjectClass() bool`

HasTraditionalStaticGroupObjectClass returns a boolean if a field has been set.

### GetMaximumMembershipUpdatesPerModify

`func (o *PluginListResponseResourcesInner) GetMaximumMembershipUpdatesPerModify() int64`

GetMaximumMembershipUpdatesPerModify returns the MaximumMembershipUpdatesPerModify field if non-nil, zero value otherwise.

### GetMaximumMembershipUpdatesPerModifyOk

`func (o *PluginListResponseResourcesInner) GetMaximumMembershipUpdatesPerModifyOk() (*int64, bool)`

GetMaximumMembershipUpdatesPerModifyOk returns a tuple with the MaximumMembershipUpdatesPerModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMembershipUpdatesPerModify

`func (o *PluginListResponseResourcesInner) SetMaximumMembershipUpdatesPerModify(v int64)`

SetMaximumMembershipUpdatesPerModify sets MaximumMembershipUpdatesPerModify field to given value.

### HasMaximumMembershipUpdatesPerModify

`func (o *PluginListResponseResourcesInner) HasMaximumMembershipUpdatesPerModify() bool`

HasMaximumMembershipUpdatesPerModify returns a boolean if a field has been set.

### GetReadOperationSupport

`func (o *PluginListResponseResourcesInner) GetReadOperationSupport() EnumpluginReadOperationSupportProp`

GetReadOperationSupport returns the ReadOperationSupport field if non-nil, zero value otherwise.

### GetReadOperationSupportOk

`func (o *PluginListResponseResourcesInner) GetReadOperationSupportOk() (*EnumpluginReadOperationSupportProp, bool)`

GetReadOperationSupportOk returns a tuple with the ReadOperationSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOperationSupport

`func (o *PluginListResponseResourcesInner) SetReadOperationSupport(v EnumpluginReadOperationSupportProp)`

SetReadOperationSupport sets ReadOperationSupport field to given value.

### HasReadOperationSupport

`func (o *PluginListResponseResourcesInner) HasReadOperationSupport() bool`

HasReadOperationSupport returns a boolean if a field has been set.

### GetPluginType

`func (o *PluginListResponseResourcesInner) GetPluginType() []EnumpluginPluginTypeProp`

GetPluginType returns the PluginType field if non-nil, zero value otherwise.

### GetPluginTypeOk

`func (o *PluginListResponseResourcesInner) GetPluginTypeOk() (*[]EnumpluginPluginTypeProp, bool)`

GetPluginTypeOk returns a tuple with the PluginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginType

`func (o *PluginListResponseResourcesInner) SetPluginType(v []EnumpluginPluginTypeProp)`

SetPluginType sets PluginType field to given value.


### GetNumThreads

`func (o *PluginListResponseResourcesInner) GetNumThreads() int64`

GetNumThreads returns the NumThreads field if non-nil, zero value otherwise.

### GetNumThreadsOk

`func (o *PluginListResponseResourcesInner) GetNumThreadsOk() (*int64, bool)`

GetNumThreadsOk returns a tuple with the NumThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumThreads

`func (o *PluginListResponseResourcesInner) SetNumThreads(v int64)`

SetNumThreads sets NumThreads field to given value.


### GetBaseDN

`func (o *PluginListResponseResourcesInner) GetBaseDN() []string`

GetBaseDN returns the BaseDN field if non-nil, zero value otherwise.

### GetBaseDNOk

`func (o *PluginListResponseResourcesInner) GetBaseDNOk() (*[]string, bool)`

GetBaseDNOk returns a tuple with the BaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDN

`func (o *PluginListResponseResourcesInner) SetBaseDN(v []string)`

SetBaseDN sets BaseDN field to given value.


### GetLowerBound

`func (o *PluginListResponseResourcesInner) GetLowerBound() int64`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *PluginListResponseResourcesInner) GetLowerBoundOk() (*int64, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *PluginListResponseResourcesInner) SetLowerBound(v int64)`

SetLowerBound sets LowerBound field to given value.

### HasLowerBound

`func (o *PluginListResponseResourcesInner) HasLowerBound() bool`

HasLowerBound returns a boolean if a field has been set.

### GetUpperBound

`func (o *PluginListResponseResourcesInner) GetUpperBound() int64`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *PluginListResponseResourcesInner) GetUpperBoundOk() (*int64, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *PluginListResponseResourcesInner) SetUpperBound(v int64)`

SetUpperBound sets UpperBound field to given value.

### HasUpperBound

`func (o *PluginListResponseResourcesInner) HasUpperBound() bool`

HasUpperBound returns a boolean if a field has been set.

### GetFilterPrefix

`func (o *PluginListResponseResourcesInner) GetFilterPrefix() string`

GetFilterPrefix returns the FilterPrefix field if non-nil, zero value otherwise.

### GetFilterPrefixOk

`func (o *PluginListResponseResourcesInner) GetFilterPrefixOk() (*string, bool)`

GetFilterPrefixOk returns a tuple with the FilterPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPrefix

`func (o *PluginListResponseResourcesInner) SetFilterPrefix(v string)`

SetFilterPrefix sets FilterPrefix field to given value.


### GetFilterSuffix

`func (o *PluginListResponseResourcesInner) GetFilterSuffix() string`

GetFilterSuffix returns the FilterSuffix field if non-nil, zero value otherwise.

### GetFilterSuffixOk

`func (o *PluginListResponseResourcesInner) GetFilterSuffixOk() (*string, bool)`

GetFilterSuffixOk returns a tuple with the FilterSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSuffix

`func (o *PluginListResponseResourcesInner) SetFilterSuffix(v string)`

SetFilterSuffix sets FilterSuffix field to given value.

### HasFilterSuffix

`func (o *PluginListResponseResourcesInner) HasFilterSuffix() bool`

HasFilterSuffix returns a boolean if a field has been set.

### GetFilter

`func (o *PluginListResponseResourcesInner) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PluginListResponseResourcesInner) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PluginListResponseResourcesInner) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetAttributeType

`func (o *PluginListResponseResourcesInner) GetAttributeType() []string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *PluginListResponseResourcesInner) GetAttributeTypeOk() (*[]string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *PluginListResponseResourcesInner) SetAttributeType(v []string)`

SetAttributeType sets AttributeType field to given value.


### GetPollingInterval

`func (o *PluginListResponseResourcesInner) GetPollingInterval() string`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *PluginListResponseResourcesInner) GetPollingIntervalOk() (*string, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *PluginListResponseResourcesInner) SetPollingInterval(v string)`

SetPollingInterval sets PollingInterval field to given value.


### GetPeerServerPriorityIndex

`func (o *PluginListResponseResourcesInner) GetPeerServerPriorityIndex() int64`

GetPeerServerPriorityIndex returns the PeerServerPriorityIndex field if non-nil, zero value otherwise.

### GetPeerServerPriorityIndexOk

`func (o *PluginListResponseResourcesInner) GetPeerServerPriorityIndexOk() (*int64, bool)`

GetPeerServerPriorityIndexOk returns a tuple with the PeerServerPriorityIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerPriorityIndex

`func (o *PluginListResponseResourcesInner) SetPeerServerPriorityIndex(v int64)`

SetPeerServerPriorityIndex sets PeerServerPriorityIndex field to given value.

### HasPeerServerPriorityIndex

`func (o *PluginListResponseResourcesInner) HasPeerServerPriorityIndex() bool`

HasPeerServerPriorityIndex returns a boolean if a field has been set.

### GetMaxUpdatesPerSecond

`func (o *PluginListResponseResourcesInner) GetMaxUpdatesPerSecond() int64`

GetMaxUpdatesPerSecond returns the MaxUpdatesPerSecond field if non-nil, zero value otherwise.

### GetMaxUpdatesPerSecondOk

`func (o *PluginListResponseResourcesInner) GetMaxUpdatesPerSecondOk() (*int64, bool)`

GetMaxUpdatesPerSecondOk returns a tuple with the MaxUpdatesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUpdatesPerSecond

`func (o *PluginListResponseResourcesInner) SetMaxUpdatesPerSecond(v int64)`

SetMaxUpdatesPerSecond sets MaxUpdatesPerSecond field to given value.


### GetNumDeleteThreads

`func (o *PluginListResponseResourcesInner) GetNumDeleteThreads() int64`

GetNumDeleteThreads returns the NumDeleteThreads field if non-nil, zero value otherwise.

### GetNumDeleteThreadsOk

`func (o *PluginListResponseResourcesInner) GetNumDeleteThreadsOk() (*int64, bool)`

GetNumDeleteThreadsOk returns a tuple with the NumDeleteThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDeleteThreads

`func (o *PluginListResponseResourcesInner) SetNumDeleteThreads(v int64)`

SetNumDeleteThreads sets NumDeleteThreads field to given value.


### GetInvokeGCDayOfWeek

`func (o *PluginListResponseResourcesInner) GetInvokeGCDayOfWeek() []EnumpluginInvokeGCDayOfWeekProp`

GetInvokeGCDayOfWeek returns the InvokeGCDayOfWeek field if non-nil, zero value otherwise.

### GetInvokeGCDayOfWeekOk

`func (o *PluginListResponseResourcesInner) GetInvokeGCDayOfWeekOk() (*[]EnumpluginInvokeGCDayOfWeekProp, bool)`

GetInvokeGCDayOfWeekOk returns a tuple with the InvokeGCDayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeGCDayOfWeek

`func (o *PluginListResponseResourcesInner) SetInvokeGCDayOfWeek(v []EnumpluginInvokeGCDayOfWeekProp)`

SetInvokeGCDayOfWeek sets InvokeGCDayOfWeek field to given value.

### HasInvokeGCDayOfWeek

`func (o *PluginListResponseResourcesInner) HasInvokeGCDayOfWeek() bool`

HasInvokeGCDayOfWeek returns a boolean if a field has been set.

### GetInvokeGCTimeUtc

`func (o *PluginListResponseResourcesInner) GetInvokeGCTimeUtc() []string`

GetInvokeGCTimeUtc returns the InvokeGCTimeUtc field if non-nil, zero value otherwise.

### GetInvokeGCTimeUtcOk

`func (o *PluginListResponseResourcesInner) GetInvokeGCTimeUtcOk() (*[]string, bool)`

GetInvokeGCTimeUtcOk returns a tuple with the InvokeGCTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvokeGCTimeUtc

`func (o *PluginListResponseResourcesInner) SetInvokeGCTimeUtc(v []string)`

SetInvokeGCTimeUtc sets InvokeGCTimeUtc field to given value.


### GetDelayAfterAlert

`func (o *PluginListResponseResourcesInner) GetDelayAfterAlert() string`

GetDelayAfterAlert returns the DelayAfterAlert field if non-nil, zero value otherwise.

### GetDelayAfterAlertOk

`func (o *PluginListResponseResourcesInner) GetDelayAfterAlertOk() (*string, bool)`

GetDelayAfterAlertOk returns a tuple with the DelayAfterAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayAfterAlert

`func (o *PluginListResponseResourcesInner) SetDelayAfterAlert(v string)`

SetDelayAfterAlert sets DelayAfterAlert field to given value.

### HasDelayAfterAlert

`func (o *PluginListResponseResourcesInner) HasDelayAfterAlert() bool`

HasDelayAfterAlert returns a boolean if a field has been set.

### GetDelayPostGC

`func (o *PluginListResponseResourcesInner) GetDelayPostGC() string`

GetDelayPostGC returns the DelayPostGC field if non-nil, zero value otherwise.

### GetDelayPostGCOk

`func (o *PluginListResponseResourcesInner) GetDelayPostGCOk() (*string, bool)`

GetDelayPostGCOk returns a tuple with the DelayPostGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayPostGC

`func (o *PluginListResponseResourcesInner) SetDelayPostGC(v string)`

SetDelayPostGC sets DelayPostGC field to given value.

### HasDelayPostGC

`func (o *PluginListResponseResourcesInner) HasDelayPostGC() bool`

HasDelayPostGC returns a boolean if a field has been set.

### GetApiURL

`func (o *PluginListResponseResourcesInner) GetApiURL() string`

GetApiURL returns the ApiURL field if non-nil, zero value otherwise.

### GetApiURLOk

`func (o *PluginListResponseResourcesInner) GetApiURLOk() (*string, bool)`

GetApiURLOk returns a tuple with the ApiURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiURL

`func (o *PluginListResponseResourcesInner) SetApiURL(v string)`

SetApiURL sets ApiURL field to given value.


### GetAuthURL

`func (o *PluginListResponseResourcesInner) GetAuthURL() string`

GetAuthURL returns the AuthURL field if non-nil, zero value otherwise.

### GetAuthURLOk

`func (o *PluginListResponseResourcesInner) GetAuthURLOk() (*string, bool)`

GetAuthURLOk returns a tuple with the AuthURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthURL

`func (o *PluginListResponseResourcesInner) SetAuthURL(v string)`

SetAuthURL sets AuthURL field to given value.


### GetOAuthClientID

`func (o *PluginListResponseResourcesInner) GetOAuthClientID() string`

GetOAuthClientID returns the OAuthClientID field if non-nil, zero value otherwise.

### GetOAuthClientIDOk

`func (o *PluginListResponseResourcesInner) GetOAuthClientIDOk() (*string, bool)`

GetOAuthClientIDOk returns a tuple with the OAuthClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientID

`func (o *PluginListResponseResourcesInner) SetOAuthClientID(v string)`

SetOAuthClientID sets OAuthClientID field to given value.


### GetOAuthClientSecret

`func (o *PluginListResponseResourcesInner) GetOAuthClientSecret() string`

GetOAuthClientSecret returns the OAuthClientSecret field if non-nil, zero value otherwise.

### GetOAuthClientSecretOk

`func (o *PluginListResponseResourcesInner) GetOAuthClientSecretOk() (*string, bool)`

GetOAuthClientSecretOk returns a tuple with the OAuthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecret

`func (o *PluginListResponseResourcesInner) SetOAuthClientSecret(v string)`

SetOAuthClientSecret sets OAuthClientSecret field to given value.

### HasOAuthClientSecret

`func (o *PluginListResponseResourcesInner) HasOAuthClientSecret() bool`

HasOAuthClientSecret returns a boolean if a field has been set.

### GetOAuthClientSecretPassphraseProvider

`func (o *PluginListResponseResourcesInner) GetOAuthClientSecretPassphraseProvider() string`

GetOAuthClientSecretPassphraseProvider returns the OAuthClientSecretPassphraseProvider field if non-nil, zero value otherwise.

### GetOAuthClientSecretPassphraseProviderOk

`func (o *PluginListResponseResourcesInner) GetOAuthClientSecretPassphraseProviderOk() (*string, bool)`

GetOAuthClientSecretPassphraseProviderOk returns a tuple with the OAuthClientSecretPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthClientSecretPassphraseProvider

`func (o *PluginListResponseResourcesInner) SetOAuthClientSecretPassphraseProvider(v string)`

SetOAuthClientSecretPassphraseProvider sets OAuthClientSecretPassphraseProvider field to given value.

### HasOAuthClientSecretPassphraseProvider

`func (o *PluginListResponseResourcesInner) HasOAuthClientSecretPassphraseProvider() bool`

HasOAuthClientSecretPassphraseProvider returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *PluginListResponseResourcesInner) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *PluginListResponseResourcesInner) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *PluginListResponseResourcesInner) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.


### GetHttpProxyExternalServer

`func (o *PluginListResponseResourcesInner) GetHttpProxyExternalServer() string`

GetHttpProxyExternalServer returns the HttpProxyExternalServer field if non-nil, zero value otherwise.

### GetHttpProxyExternalServerOk

`func (o *PluginListResponseResourcesInner) GetHttpProxyExternalServerOk() (*string, bool)`

GetHttpProxyExternalServerOk returns a tuple with the HttpProxyExternalServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyExternalServer

`func (o *PluginListResponseResourcesInner) SetHttpProxyExternalServer(v string)`

SetHttpProxyExternalServer sets HttpProxyExternalServer field to given value.

### HasHttpProxyExternalServer

`func (o *PluginListResponseResourcesInner) HasHttpProxyExternalServer() bool`

HasHttpProxyExternalServer returns a boolean if a field has been set.

### GetIncludedLocalEntryBaseDN

`func (o *PluginListResponseResourcesInner) GetIncludedLocalEntryBaseDN() []string`

GetIncludedLocalEntryBaseDN returns the IncludedLocalEntryBaseDN field if non-nil, zero value otherwise.

### GetIncludedLocalEntryBaseDNOk

`func (o *PluginListResponseResourcesInner) GetIncludedLocalEntryBaseDNOk() (*[]string, bool)`

GetIncludedLocalEntryBaseDNOk returns a tuple with the IncludedLocalEntryBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLocalEntryBaseDN

`func (o *PluginListResponseResourcesInner) SetIncludedLocalEntryBaseDN(v []string)`

SetIncludedLocalEntryBaseDN sets IncludedLocalEntryBaseDN field to given value.

### HasIncludedLocalEntryBaseDN

`func (o *PluginListResponseResourcesInner) HasIncludedLocalEntryBaseDN() bool`

HasIncludedLocalEntryBaseDN returns a boolean if a field has been set.

### GetConnectionCriteria

`func (o *PluginListResponseResourcesInner) GetConnectionCriteria() string`

GetConnectionCriteria returns the ConnectionCriteria field if non-nil, zero value otherwise.

### GetConnectionCriteriaOk

`func (o *PluginListResponseResourcesInner) GetConnectionCriteriaOk() (*string, bool)`

GetConnectionCriteriaOk returns a tuple with the ConnectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCriteria

`func (o *PluginListResponseResourcesInner) SetConnectionCriteria(v string)`

SetConnectionCriteria sets ConnectionCriteria field to given value.

### HasConnectionCriteria

`func (o *PluginListResponseResourcesInner) HasConnectionCriteria() bool`

HasConnectionCriteria returns a boolean if a field has been set.

### GetTryLocalBind

`func (o *PluginListResponseResourcesInner) GetTryLocalBind() bool`

GetTryLocalBind returns the TryLocalBind field if non-nil, zero value otherwise.

### GetTryLocalBindOk

`func (o *PluginListResponseResourcesInner) GetTryLocalBindOk() (*bool, bool)`

GetTryLocalBindOk returns a tuple with the TryLocalBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTryLocalBind

`func (o *PluginListResponseResourcesInner) SetTryLocalBind(v bool)`

SetTryLocalBind sets TryLocalBind field to given value.


### GetOverrideLocalPassword

`func (o *PluginListResponseResourcesInner) GetOverrideLocalPassword() bool`

GetOverrideLocalPassword returns the OverrideLocalPassword field if non-nil, zero value otherwise.

### GetOverrideLocalPasswordOk

`func (o *PluginListResponseResourcesInner) GetOverrideLocalPasswordOk() (*bool, bool)`

GetOverrideLocalPasswordOk returns a tuple with the OverrideLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideLocalPassword

`func (o *PluginListResponseResourcesInner) SetOverrideLocalPassword(v bool)`

SetOverrideLocalPassword sets OverrideLocalPassword field to given value.


### GetUpdateLocalPassword

`func (o *PluginListResponseResourcesInner) GetUpdateLocalPassword() bool`

GetUpdateLocalPassword returns the UpdateLocalPassword field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordOk

`func (o *PluginListResponseResourcesInner) GetUpdateLocalPasswordOk() (*bool, bool)`

GetUpdateLocalPasswordOk returns a tuple with the UpdateLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPassword

`func (o *PluginListResponseResourcesInner) SetUpdateLocalPassword(v bool)`

SetUpdateLocalPassword sets UpdateLocalPassword field to given value.


### GetUpdateLocalPasswordDN

`func (o *PluginListResponseResourcesInner) GetUpdateLocalPasswordDN() string`

GetUpdateLocalPasswordDN returns the UpdateLocalPasswordDN field if non-nil, zero value otherwise.

### GetUpdateLocalPasswordDNOk

`func (o *PluginListResponseResourcesInner) GetUpdateLocalPasswordDNOk() (*string, bool)`

GetUpdateLocalPasswordDNOk returns a tuple with the UpdateLocalPasswordDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateLocalPasswordDN

`func (o *PluginListResponseResourcesInner) SetUpdateLocalPasswordDN(v string)`

SetUpdateLocalPasswordDN sets UpdateLocalPasswordDN field to given value.

### HasUpdateLocalPasswordDN

`func (o *PluginListResponseResourcesInner) HasUpdateLocalPasswordDN() bool`

HasUpdateLocalPasswordDN returns a boolean if a field has been set.

### GetAllowLaxPassThroughAuthenticationPasswords

`func (o *PluginListResponseResourcesInner) GetAllowLaxPassThroughAuthenticationPasswords() bool`

GetAllowLaxPassThroughAuthenticationPasswords returns the AllowLaxPassThroughAuthenticationPasswords field if non-nil, zero value otherwise.

### GetAllowLaxPassThroughAuthenticationPasswordsOk

`func (o *PluginListResponseResourcesInner) GetAllowLaxPassThroughAuthenticationPasswordsOk() (*bool, bool)`

GetAllowLaxPassThroughAuthenticationPasswordsOk returns a tuple with the AllowLaxPassThroughAuthenticationPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowLaxPassThroughAuthenticationPasswords

`func (o *PluginListResponseResourcesInner) SetAllowLaxPassThroughAuthenticationPasswords(v bool)`

SetAllowLaxPassThroughAuthenticationPasswords sets AllowLaxPassThroughAuthenticationPasswords field to given value.

### HasAllowLaxPassThroughAuthenticationPasswords

`func (o *PluginListResponseResourcesInner) HasAllowLaxPassThroughAuthenticationPasswords() bool`

HasAllowLaxPassThroughAuthenticationPasswords returns a boolean if a field has been set.

### GetIgnoredPasswordPolicyStateErrorCondition

`func (o *PluginListResponseResourcesInner) GetIgnoredPasswordPolicyStateErrorCondition() []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp`

GetIgnoredPasswordPolicyStateErrorCondition returns the IgnoredPasswordPolicyStateErrorCondition field if non-nil, zero value otherwise.

### GetIgnoredPasswordPolicyStateErrorConditionOk

`func (o *PluginListResponseResourcesInner) GetIgnoredPasswordPolicyStateErrorConditionOk() (*[]EnumpluginIgnoredPasswordPolicyStateErrorConditionProp, bool)`

GetIgnoredPasswordPolicyStateErrorConditionOk returns a tuple with the IgnoredPasswordPolicyStateErrorCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredPasswordPolicyStateErrorCondition

`func (o *PluginListResponseResourcesInner) SetIgnoredPasswordPolicyStateErrorCondition(v []EnumpluginIgnoredPasswordPolicyStateErrorConditionProp)`

SetIgnoredPasswordPolicyStateErrorCondition sets IgnoredPasswordPolicyStateErrorCondition field to given value.

### HasIgnoredPasswordPolicyStateErrorCondition

`func (o *PluginListResponseResourcesInner) HasIgnoredPasswordPolicyStateErrorCondition() bool`

HasIgnoredPasswordPolicyStateErrorCondition returns a boolean if a field has been set.

### GetUserMappingLocalAttribute

`func (o *PluginListResponseResourcesInner) GetUserMappingLocalAttribute() []string`

GetUserMappingLocalAttribute returns the UserMappingLocalAttribute field if non-nil, zero value otherwise.

### GetUserMappingLocalAttributeOk

`func (o *PluginListResponseResourcesInner) GetUserMappingLocalAttributeOk() (*[]string, bool)`

GetUserMappingLocalAttributeOk returns a tuple with the UserMappingLocalAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingLocalAttribute

`func (o *PluginListResponseResourcesInner) SetUserMappingLocalAttribute(v []string)`

SetUserMappingLocalAttribute sets UserMappingLocalAttribute field to given value.


### GetUserMappingRemoteJSONField

`func (o *PluginListResponseResourcesInner) GetUserMappingRemoteJSONField() []string`

GetUserMappingRemoteJSONField returns the UserMappingRemoteJSONField field if non-nil, zero value otherwise.

### GetUserMappingRemoteJSONFieldOk

`func (o *PluginListResponseResourcesInner) GetUserMappingRemoteJSONFieldOk() (*[]string, bool)`

GetUserMappingRemoteJSONFieldOk returns a tuple with the UserMappingRemoteJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMappingRemoteJSONField

`func (o *PluginListResponseResourcesInner) SetUserMappingRemoteJSONField(v []string)`

SetUserMappingRemoteJSONField sets UserMappingRemoteJSONField field to given value.


### GetAdditionalUserMappingSCIMFilter

`func (o *PluginListResponseResourcesInner) GetAdditionalUserMappingSCIMFilter() string`

GetAdditionalUserMappingSCIMFilter returns the AdditionalUserMappingSCIMFilter field if non-nil, zero value otherwise.

### GetAdditionalUserMappingSCIMFilterOk

`func (o *PluginListResponseResourcesInner) GetAdditionalUserMappingSCIMFilterOk() (*string, bool)`

GetAdditionalUserMappingSCIMFilterOk returns a tuple with the AdditionalUserMappingSCIMFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUserMappingSCIMFilter

`func (o *PluginListResponseResourcesInner) SetAdditionalUserMappingSCIMFilter(v string)`

SetAdditionalUserMappingSCIMFilter sets AdditionalUserMappingSCIMFilter field to given value.

### HasAdditionalUserMappingSCIMFilter

`func (o *PluginListResponseResourcesInner) HasAdditionalUserMappingSCIMFilter() bool`

HasAdditionalUserMappingSCIMFilter returns a boolean if a field has been set.

### GetChangelogPasswordEncryptionKey

`func (o *PluginListResponseResourcesInner) GetChangelogPasswordEncryptionKey() string`

GetChangelogPasswordEncryptionKey returns the ChangelogPasswordEncryptionKey field if non-nil, zero value otherwise.

### GetChangelogPasswordEncryptionKeyOk

`func (o *PluginListResponseResourcesInner) GetChangelogPasswordEncryptionKeyOk() (*string, bool)`

GetChangelogPasswordEncryptionKeyOk returns a tuple with the ChangelogPasswordEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPasswordEncryptionKey

`func (o *PluginListResponseResourcesInner) SetChangelogPasswordEncryptionKey(v string)`

SetChangelogPasswordEncryptionKey sets ChangelogPasswordEncryptionKey field to given value.

### HasChangelogPasswordEncryptionKey

`func (o *PluginListResponseResourcesInner) HasChangelogPasswordEncryptionKey() bool`

HasChangelogPasswordEncryptionKey returns a boolean if a field has been set.

### GetChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *PluginListResponseResourcesInner) GetChangelogPasswordEncryptionKeyPassphraseProvider() string`

GetChangelogPasswordEncryptionKeyPassphraseProvider returns the ChangelogPasswordEncryptionKeyPassphraseProvider field if non-nil, zero value otherwise.

### GetChangelogPasswordEncryptionKeyPassphraseProviderOk

`func (o *PluginListResponseResourcesInner) GetChangelogPasswordEncryptionKeyPassphraseProviderOk() (*string, bool)`

GetChangelogPasswordEncryptionKeyPassphraseProviderOk returns a tuple with the ChangelogPasswordEncryptionKeyPassphraseProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *PluginListResponseResourcesInner) SetChangelogPasswordEncryptionKeyPassphraseProvider(v string)`

SetChangelogPasswordEncryptionKeyPassphraseProvider sets ChangelogPasswordEncryptionKeyPassphraseProvider field to given value.

### HasChangelogPasswordEncryptionKeyPassphraseProvider

`func (o *PluginListResponseResourcesInner) HasChangelogPasswordEncryptionKeyPassphraseProvider() bool`

HasChangelogPasswordEncryptionKeyPassphraseProvider returns a boolean if a field has been set.

### GetType

`func (o *PluginListResponseResourcesInner) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PluginListResponseResourcesInner) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PluginListResponseResourcesInner) SetType(v []string)`

SetType sets Type field to given value.


### GetMultipleAttributeBehavior

`func (o *PluginListResponseResourcesInner) GetMultipleAttributeBehavior() EnumpluginUniqueAttributeMultipleAttributeBehaviorProp`

GetMultipleAttributeBehavior returns the MultipleAttributeBehavior field if non-nil, zero value otherwise.

### GetMultipleAttributeBehaviorOk

`func (o *PluginListResponseResourcesInner) GetMultipleAttributeBehaviorOk() (*EnumpluginUniqueAttributeMultipleAttributeBehaviorProp, bool)`

GetMultipleAttributeBehaviorOk returns a tuple with the MultipleAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAttributeBehavior

`func (o *PluginListResponseResourcesInner) SetMultipleAttributeBehavior(v EnumpluginUniqueAttributeMultipleAttributeBehaviorProp)`

SetMultipleAttributeBehavior sets MultipleAttributeBehavior field to given value.

### HasMultipleAttributeBehavior

`func (o *PluginListResponseResourcesInner) HasMultipleAttributeBehavior() bool`

HasMultipleAttributeBehavior returns a boolean if a field has been set.

### GetSubtreeView

`func (o *PluginListResponseResourcesInner) GetSubtreeView() []string`

GetSubtreeView returns the SubtreeView field if non-nil, zero value otherwise.

### GetSubtreeViewOk

`func (o *PluginListResponseResourcesInner) GetSubtreeViewOk() (*[]string, bool)`

GetSubtreeViewOk returns a tuple with the SubtreeView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtreeView

`func (o *PluginListResponseResourcesInner) SetSubtreeView(v []string)`

SetSubtreeView sets SubtreeView field to given value.


### GetPreventConflictsWithSoftDeletedEntries

`func (o *PluginListResponseResourcesInner) GetPreventConflictsWithSoftDeletedEntries() bool`

GetPreventConflictsWithSoftDeletedEntries returns the PreventConflictsWithSoftDeletedEntries field if non-nil, zero value otherwise.

### GetPreventConflictsWithSoftDeletedEntriesOk

`func (o *PluginListResponseResourcesInner) GetPreventConflictsWithSoftDeletedEntriesOk() (*bool, bool)`

GetPreventConflictsWithSoftDeletedEntriesOk returns a tuple with the PreventConflictsWithSoftDeletedEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventConflictsWithSoftDeletedEntries

`func (o *PluginListResponseResourcesInner) SetPreventConflictsWithSoftDeletedEntries(v bool)`

SetPreventConflictsWithSoftDeletedEntries sets PreventConflictsWithSoftDeletedEntries field to given value.

### HasPreventConflictsWithSoftDeletedEntries

`func (o *PluginListResponseResourcesInner) HasPreventConflictsWithSoftDeletedEntries() bool`

HasPreventConflictsWithSoftDeletedEntries returns a boolean if a field has been set.

### GetPreCommitValidation

`func (o *PluginListResponseResourcesInner) GetPreCommitValidation() EnumpluginPreCommitValidationProp`

GetPreCommitValidation returns the PreCommitValidation field if non-nil, zero value otherwise.

### GetPreCommitValidationOk

`func (o *PluginListResponseResourcesInner) GetPreCommitValidationOk() (*EnumpluginPreCommitValidationProp, bool)`

GetPreCommitValidationOk returns a tuple with the PreCommitValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCommitValidation

`func (o *PluginListResponseResourcesInner) SetPreCommitValidation(v EnumpluginPreCommitValidationProp)`

SetPreCommitValidation sets PreCommitValidation field to given value.

### HasPreCommitValidation

`func (o *PluginListResponseResourcesInner) HasPreCommitValidation() bool`

HasPreCommitValidation returns a boolean if a field has been set.

### GetPostCommitValidation

`func (o *PluginListResponseResourcesInner) GetPostCommitValidation() EnumpluginPostCommitValidationProp`

GetPostCommitValidation returns the PostCommitValidation field if non-nil, zero value otherwise.

### GetPostCommitValidationOk

`func (o *PluginListResponseResourcesInner) GetPostCommitValidationOk() (*EnumpluginPostCommitValidationProp, bool)`

GetPostCommitValidationOk returns a tuple with the PostCommitValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCommitValidation

`func (o *PluginListResponseResourcesInner) SetPostCommitValidation(v EnumpluginPostCommitValidationProp)`

SetPostCommitValidation sets PostCommitValidation field to given value.

### HasPostCommitValidation

`func (o *PluginListResponseResourcesInner) HasPostCommitValidation() bool`

HasPostCommitValidation returns a boolean if a field has been set.

### GetHistogramCategoryBoundary

`func (o *PluginListResponseResourcesInner) GetHistogramCategoryBoundary() []string`

GetHistogramCategoryBoundary returns the HistogramCategoryBoundary field if non-nil, zero value otherwise.

### GetHistogramCategoryBoundaryOk

`func (o *PluginListResponseResourcesInner) GetHistogramCategoryBoundaryOk() (*[]string, bool)`

GetHistogramCategoryBoundaryOk returns a tuple with the HistogramCategoryBoundary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramCategoryBoundary

`func (o *PluginListResponseResourcesInner) SetHistogramCategoryBoundary(v []string)`

SetHistogramCategoryBoundary sets HistogramCategoryBoundary field to given value.


### GetIncludeQueueTime

`func (o *PluginListResponseResourcesInner) GetIncludeQueueTime() bool`

GetIncludeQueueTime returns the IncludeQueueTime field if non-nil, zero value otherwise.

### GetIncludeQueueTimeOk

`func (o *PluginListResponseResourcesInner) GetIncludeQueueTimeOk() (*bool, bool)`

GetIncludeQueueTimeOk returns a tuple with the IncludeQueueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeQueueTime

`func (o *PluginListResponseResourcesInner) SetIncludeQueueTime(v bool)`

SetIncludeQueueTime sets IncludeQueueTime field to given value.

### HasIncludeQueueTime

`func (o *PluginListResponseResourcesInner) HasIncludeQueueTime() bool`

HasIncludeQueueTime returns a boolean if a field has been set.

### GetSeparateMonitorEntryPerTrackedApplication

`func (o *PluginListResponseResourcesInner) GetSeparateMonitorEntryPerTrackedApplication() bool`

GetSeparateMonitorEntryPerTrackedApplication returns the SeparateMonitorEntryPerTrackedApplication field if non-nil, zero value otherwise.

### GetSeparateMonitorEntryPerTrackedApplicationOk

`func (o *PluginListResponseResourcesInner) GetSeparateMonitorEntryPerTrackedApplicationOk() (*bool, bool)`

GetSeparateMonitorEntryPerTrackedApplicationOk returns a tuple with the SeparateMonitorEntryPerTrackedApplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeparateMonitorEntryPerTrackedApplication

`func (o *PluginListResponseResourcesInner) SetSeparateMonitorEntryPerTrackedApplication(v bool)`

SetSeparateMonitorEntryPerTrackedApplication sets SeparateMonitorEntryPerTrackedApplication field to given value.

### HasSeparateMonitorEntryPerTrackedApplication

`func (o *PluginListResponseResourcesInner) HasSeparateMonitorEntryPerTrackedApplication() bool`

HasSeparateMonitorEntryPerTrackedApplication returns a boolean if a field has been set.

### GetScope

`func (o *PluginListResponseResourcesInner) GetScope() EnumpluginScopeProp`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PluginListResponseResourcesInner) GetScopeOk() (*EnumpluginScopeProp, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PluginListResponseResourcesInner) SetScope(v EnumpluginScopeProp)`

SetScope sets Scope field to given value.


### GetIncludeAttribute

`func (o *PluginListResponseResourcesInner) GetIncludeAttribute() []string`

GetIncludeAttribute returns the IncludeAttribute field if non-nil, zero value otherwise.

### GetIncludeAttributeOk

`func (o *PluginListResponseResourcesInner) GetIncludeAttributeOk() (*[]string, bool)`

GetIncludeAttributeOk returns a tuple with the IncludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttribute

`func (o *PluginListResponseResourcesInner) SetIncludeAttribute(v []string)`

SetIncludeAttribute sets IncludeAttribute field to given value.

### HasIncludeAttribute

`func (o *PluginListResponseResourcesInner) HasIncludeAttribute() bool`

HasIncludeAttribute returns a boolean if a field has been set.

### GetOutputFile

`func (o *PluginListResponseResourcesInner) GetOutputFile() string`

GetOutputFile returns the OutputFile field if non-nil, zero value otherwise.

### GetOutputFileOk

`func (o *PluginListResponseResourcesInner) GetOutputFileOk() (*string, bool)`

GetOutputFileOk returns a tuple with the OutputFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFile

`func (o *PluginListResponseResourcesInner) SetOutputFile(v string)`

SetOutputFile sets OutputFile field to given value.


### GetPreviousFileExtension

`func (o *PluginListResponseResourcesInner) GetPreviousFileExtension() string`

GetPreviousFileExtension returns the PreviousFileExtension field if non-nil, zero value otherwise.

### GetPreviousFileExtensionOk

`func (o *PluginListResponseResourcesInner) GetPreviousFileExtensionOk() (*string, bool)`

GetPreviousFileExtensionOk returns a tuple with the PreviousFileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousFileExtension

`func (o *PluginListResponseResourcesInner) SetPreviousFileExtension(v string)`

SetPreviousFileExtension sets PreviousFileExtension field to given value.

### HasPreviousFileExtension

`func (o *PluginListResponseResourcesInner) HasPreviousFileExtension() bool`

HasPreviousFileExtension returns a boolean if a field has been set.

### GetLogInterval

`func (o *PluginListResponseResourcesInner) GetLogInterval() string`

GetLogInterval returns the LogInterval field if non-nil, zero value otherwise.

### GetLogIntervalOk

`func (o *PluginListResponseResourcesInner) GetLogIntervalOk() (*string, bool)`

GetLogIntervalOk returns a tuple with the LogInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogInterval

`func (o *PluginListResponseResourcesInner) SetLogInterval(v string)`

SetLogInterval sets LogInterval field to given value.


### GetSuppressIfIdle

`func (o *PluginListResponseResourcesInner) GetSuppressIfIdle() bool`

GetSuppressIfIdle returns the SuppressIfIdle field if non-nil, zero value otherwise.

### GetSuppressIfIdleOk

`func (o *PluginListResponseResourcesInner) GetSuppressIfIdleOk() (*bool, bool)`

GetSuppressIfIdleOk returns a tuple with the SuppressIfIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressIfIdle

`func (o *PluginListResponseResourcesInner) SetSuppressIfIdle(v bool)`

SetSuppressIfIdle sets SuppressIfIdle field to given value.


### GetHeaderPrefixPerColumn

`func (o *PluginListResponseResourcesInner) GetHeaderPrefixPerColumn() bool`

GetHeaderPrefixPerColumn returns the HeaderPrefixPerColumn field if non-nil, zero value otherwise.

### GetHeaderPrefixPerColumnOk

`func (o *PluginListResponseResourcesInner) GetHeaderPrefixPerColumnOk() (*bool, bool)`

GetHeaderPrefixPerColumnOk returns a tuple with the HeaderPrefixPerColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderPrefixPerColumn

`func (o *PluginListResponseResourcesInner) SetHeaderPrefixPerColumn(v bool)`

SetHeaderPrefixPerColumn sets HeaderPrefixPerColumn field to given value.

### HasHeaderPrefixPerColumn

`func (o *PluginListResponseResourcesInner) HasHeaderPrefixPerColumn() bool`

HasHeaderPrefixPerColumn returns a boolean if a field has been set.

### GetEmptyInsteadOfZero

`func (o *PluginListResponseResourcesInner) GetEmptyInsteadOfZero() bool`

GetEmptyInsteadOfZero returns the EmptyInsteadOfZero field if non-nil, zero value otherwise.

### GetEmptyInsteadOfZeroOk

`func (o *PluginListResponseResourcesInner) GetEmptyInsteadOfZeroOk() (*bool, bool)`

GetEmptyInsteadOfZeroOk returns a tuple with the EmptyInsteadOfZero field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyInsteadOfZero

`func (o *PluginListResponseResourcesInner) SetEmptyInsteadOfZero(v bool)`

SetEmptyInsteadOfZero sets EmptyInsteadOfZero field to given value.

### HasEmptyInsteadOfZero

`func (o *PluginListResponseResourcesInner) HasEmptyInsteadOfZero() bool`

HasEmptyInsteadOfZero returns a boolean if a field has been set.

### GetLinesBetweenHeader

`func (o *PluginListResponseResourcesInner) GetLinesBetweenHeader() int64`

GetLinesBetweenHeader returns the LinesBetweenHeader field if non-nil, zero value otherwise.

### GetLinesBetweenHeaderOk

`func (o *PluginListResponseResourcesInner) GetLinesBetweenHeaderOk() (*int64, bool)`

GetLinesBetweenHeaderOk returns a tuple with the LinesBetweenHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesBetweenHeader

`func (o *PluginListResponseResourcesInner) SetLinesBetweenHeader(v int64)`

SetLinesBetweenHeader sets LinesBetweenHeader field to given value.


### GetIncludedLDAPStat

`func (o *PluginListResponseResourcesInner) GetIncludedLDAPStat() []EnumpluginIncludedLDAPStatProp`

GetIncludedLDAPStat returns the IncludedLDAPStat field if non-nil, zero value otherwise.

### GetIncludedLDAPStatOk

`func (o *PluginListResponseResourcesInner) GetIncludedLDAPStatOk() (*[]EnumpluginIncludedLDAPStatProp, bool)`

GetIncludedLDAPStatOk returns a tuple with the IncludedLDAPStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedLDAPStat

`func (o *PluginListResponseResourcesInner) SetIncludedLDAPStat(v []EnumpluginIncludedLDAPStatProp)`

SetIncludedLDAPStat sets IncludedLDAPStat field to given value.

### HasIncludedLDAPStat

`func (o *PluginListResponseResourcesInner) HasIncludedLDAPStat() bool`

HasIncludedLDAPStat returns a boolean if a field has been set.

### GetIncludedResourceStat

`func (o *PluginListResponseResourcesInner) GetIncludedResourceStat() []EnumpluginIncludedResourceStatProp`

GetIncludedResourceStat returns the IncludedResourceStat field if non-nil, zero value otherwise.

### GetIncludedResourceStatOk

`func (o *PluginListResponseResourcesInner) GetIncludedResourceStatOk() (*[]EnumpluginIncludedResourceStatProp, bool)`

GetIncludedResourceStatOk returns a tuple with the IncludedResourceStat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedResourceStat

`func (o *PluginListResponseResourcesInner) SetIncludedResourceStat(v []EnumpluginIncludedResourceStatProp)`

SetIncludedResourceStat sets IncludedResourceStat field to given value.

### HasIncludedResourceStat

`func (o *PluginListResponseResourcesInner) HasIncludedResourceStat() bool`

HasIncludedResourceStat returns a boolean if a field has been set.

### GetHistogramFormat

`func (o *PluginListResponseResourcesInner) GetHistogramFormat() EnumpluginHistogramFormatProp`

GetHistogramFormat returns the HistogramFormat field if non-nil, zero value otherwise.

### GetHistogramFormatOk

`func (o *PluginListResponseResourcesInner) GetHistogramFormatOk() (*EnumpluginHistogramFormatProp, bool)`

GetHistogramFormatOk returns a tuple with the HistogramFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramFormat

`func (o *PluginListResponseResourcesInner) SetHistogramFormat(v EnumpluginHistogramFormatProp)`

SetHistogramFormat sets HistogramFormat field to given value.


### GetHistogramOpType

`func (o *PluginListResponseResourcesInner) GetHistogramOpType() []EnumpluginHistogramOpTypeProp`

GetHistogramOpType returns the HistogramOpType field if non-nil, zero value otherwise.

### GetHistogramOpTypeOk

`func (o *PluginListResponseResourcesInner) GetHistogramOpTypeOk() (*[]EnumpluginHistogramOpTypeProp, bool)`

GetHistogramOpTypeOk returns a tuple with the HistogramOpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistogramOpType

`func (o *PluginListResponseResourcesInner) SetHistogramOpType(v []EnumpluginHistogramOpTypeProp)`

SetHistogramOpType sets HistogramOpType field to given value.

### HasHistogramOpType

`func (o *PluginListResponseResourcesInner) HasHistogramOpType() bool`

HasHistogramOpType returns a boolean if a field has been set.

### GetGaugeInfo

`func (o *PluginListResponseResourcesInner) GetGaugeInfo() EnumpluginGaugeInfoProp`

GetGaugeInfo returns the GaugeInfo field if non-nil, zero value otherwise.

### GetGaugeInfoOk

`func (o *PluginListResponseResourcesInner) GetGaugeInfoOk() (*EnumpluginGaugeInfoProp, bool)`

GetGaugeInfoOk returns a tuple with the GaugeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaugeInfo

`func (o *PluginListResponseResourcesInner) SetGaugeInfo(v EnumpluginGaugeInfoProp)`

SetGaugeInfo sets GaugeInfo field to given value.

### HasGaugeInfo

`func (o *PluginListResponseResourcesInner) HasGaugeInfo() bool`

HasGaugeInfo returns a boolean if a field has been set.

### GetLogFileFormat

`func (o *PluginListResponseResourcesInner) GetLogFileFormat() EnumpluginLogFileFormatProp`

GetLogFileFormat returns the LogFileFormat field if non-nil, zero value otherwise.

### GetLogFileFormatOk

`func (o *PluginListResponseResourcesInner) GetLogFileFormatOk() (*EnumpluginLogFileFormatProp, bool)`

GetLogFileFormatOk returns a tuple with the LogFileFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileFormat

`func (o *PluginListResponseResourcesInner) SetLogFileFormat(v EnumpluginLogFileFormatProp)`

SetLogFileFormat sets LogFileFormat field to given value.

### HasLogFileFormat

`func (o *PluginListResponseResourcesInner) HasLogFileFormat() bool`

HasLogFileFormat returns a boolean if a field has been set.

### GetLogFile

`func (o *PluginListResponseResourcesInner) GetLogFile() string`

GetLogFile returns the LogFile field if non-nil, zero value otherwise.

### GetLogFileOk

`func (o *PluginListResponseResourcesInner) GetLogFileOk() (*string, bool)`

GetLogFileOk returns a tuple with the LogFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFile

`func (o *PluginListResponseResourcesInner) SetLogFile(v string)`

SetLogFile sets LogFile field to given value.


### GetLogFilePermissions

`func (o *PluginListResponseResourcesInner) GetLogFilePermissions() string`

GetLogFilePermissions returns the LogFilePermissions field if non-nil, zero value otherwise.

### GetLogFilePermissionsOk

`func (o *PluginListResponseResourcesInner) GetLogFilePermissionsOk() (*string, bool)`

GetLogFilePermissionsOk returns a tuple with the LogFilePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFilePermissions

`func (o *PluginListResponseResourcesInner) SetLogFilePermissions(v string)`

SetLogFilePermissions sets LogFilePermissions field to given value.


### GetAppend

`func (o *PluginListResponseResourcesInner) GetAppend() bool`

GetAppend returns the Append field if non-nil, zero value otherwise.

### GetAppendOk

`func (o *PluginListResponseResourcesInner) GetAppendOk() (*bool, bool)`

GetAppendOk returns a tuple with the Append field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppend

`func (o *PluginListResponseResourcesInner) SetAppend(v bool)`

SetAppend sets Append field to given value.

### HasAppend

`func (o *PluginListResponseResourcesInner) HasAppend() bool`

HasAppend returns a boolean if a field has been set.

### GetRotationPolicy

`func (o *PluginListResponseResourcesInner) GetRotationPolicy() []string`

GetRotationPolicy returns the RotationPolicy field if non-nil, zero value otherwise.

### GetRotationPolicyOk

`func (o *PluginListResponseResourcesInner) GetRotationPolicyOk() (*[]string, bool)`

GetRotationPolicyOk returns a tuple with the RotationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPolicy

`func (o *PluginListResponseResourcesInner) SetRotationPolicy(v []string)`

SetRotationPolicy sets RotationPolicy field to given value.


### GetRotationListener

`func (o *PluginListResponseResourcesInner) GetRotationListener() []string`

GetRotationListener returns the RotationListener field if non-nil, zero value otherwise.

### GetRotationListenerOk

`func (o *PluginListResponseResourcesInner) GetRotationListenerOk() (*[]string, bool)`

GetRotationListenerOk returns a tuple with the RotationListener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationListener

`func (o *PluginListResponseResourcesInner) SetRotationListener(v []string)`

SetRotationListener sets RotationListener field to given value.

### HasRotationListener

`func (o *PluginListResponseResourcesInner) HasRotationListener() bool`

HasRotationListener returns a boolean if a field has been set.

### GetRetentionPolicy

`func (o *PluginListResponseResourcesInner) GetRetentionPolicy() []string`

GetRetentionPolicy returns the RetentionPolicy field if non-nil, zero value otherwise.

### GetRetentionPolicyOk

`func (o *PluginListResponseResourcesInner) GetRetentionPolicyOk() (*[]string, bool)`

GetRetentionPolicyOk returns a tuple with the RetentionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicy

`func (o *PluginListResponseResourcesInner) SetRetentionPolicy(v []string)`

SetRetentionPolicy sets RetentionPolicy field to given value.


### GetLoggingErrorBehavior

`func (o *PluginListResponseResourcesInner) GetLoggingErrorBehavior() EnumpluginLoggingErrorBehaviorProp`

GetLoggingErrorBehavior returns the LoggingErrorBehavior field if non-nil, zero value otherwise.

### GetLoggingErrorBehaviorOk

`func (o *PluginListResponseResourcesInner) GetLoggingErrorBehaviorOk() (*EnumpluginLoggingErrorBehaviorProp, bool)`

GetLoggingErrorBehaviorOk returns a tuple with the LoggingErrorBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggingErrorBehavior

`func (o *PluginListResponseResourcesInner) SetLoggingErrorBehavior(v EnumpluginLoggingErrorBehaviorProp)`

SetLoggingErrorBehavior sets LoggingErrorBehavior field to given value.

### HasLoggingErrorBehavior

`func (o *PluginListResponseResourcesInner) HasLoggingErrorBehavior() bool`

HasLoggingErrorBehavior returns a boolean if a field has been set.

### GetDatetimeAttribute

`func (o *PluginListResponseResourcesInner) GetDatetimeAttribute() string`

GetDatetimeAttribute returns the DatetimeAttribute field if non-nil, zero value otherwise.

### GetDatetimeAttributeOk

`func (o *PluginListResponseResourcesInner) GetDatetimeAttributeOk() (*string, bool)`

GetDatetimeAttributeOk returns a tuple with the DatetimeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeAttribute

`func (o *PluginListResponseResourcesInner) SetDatetimeAttribute(v string)`

SetDatetimeAttribute sets DatetimeAttribute field to given value.


### GetDatetimeJSONField

`func (o *PluginListResponseResourcesInner) GetDatetimeJSONField() string`

GetDatetimeJSONField returns the DatetimeJSONField field if non-nil, zero value otherwise.

### GetDatetimeJSONFieldOk

`func (o *PluginListResponseResourcesInner) GetDatetimeJSONFieldOk() (*string, bool)`

GetDatetimeJSONFieldOk returns a tuple with the DatetimeJSONField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeJSONField

`func (o *PluginListResponseResourcesInner) SetDatetimeJSONField(v string)`

SetDatetimeJSONField sets DatetimeJSONField field to given value.

### HasDatetimeJSONField

`func (o *PluginListResponseResourcesInner) HasDatetimeJSONField() bool`

HasDatetimeJSONField returns a boolean if a field has been set.

### GetDatetimeFormat

`func (o *PluginListResponseResourcesInner) GetDatetimeFormat() EnumpluginDatetimeFormatProp`

GetDatetimeFormat returns the DatetimeFormat field if non-nil, zero value otherwise.

### GetDatetimeFormatOk

`func (o *PluginListResponseResourcesInner) GetDatetimeFormatOk() (*EnumpluginDatetimeFormatProp, bool)`

GetDatetimeFormatOk returns a tuple with the DatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetimeFormat

`func (o *PluginListResponseResourcesInner) SetDatetimeFormat(v EnumpluginDatetimeFormatProp)`

SetDatetimeFormat sets DatetimeFormat field to given value.


### GetCustomDatetimeFormat

`func (o *PluginListResponseResourcesInner) GetCustomDatetimeFormat() string`

GetCustomDatetimeFormat returns the CustomDatetimeFormat field if non-nil, zero value otherwise.

### GetCustomDatetimeFormatOk

`func (o *PluginListResponseResourcesInner) GetCustomDatetimeFormatOk() (*string, bool)`

GetCustomDatetimeFormatOk returns a tuple with the CustomDatetimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDatetimeFormat

`func (o *PluginListResponseResourcesInner) SetCustomDatetimeFormat(v string)`

SetCustomDatetimeFormat sets CustomDatetimeFormat field to given value.

### HasCustomDatetimeFormat

`func (o *PluginListResponseResourcesInner) HasCustomDatetimeFormat() bool`

HasCustomDatetimeFormat returns a boolean if a field has been set.

### GetCustomTimezone

`func (o *PluginListResponseResourcesInner) GetCustomTimezone() string`

GetCustomTimezone returns the CustomTimezone field if non-nil, zero value otherwise.

### GetCustomTimezoneOk

`func (o *PluginListResponseResourcesInner) GetCustomTimezoneOk() (*string, bool)`

GetCustomTimezoneOk returns a tuple with the CustomTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimezone

`func (o *PluginListResponseResourcesInner) SetCustomTimezone(v string)`

SetCustomTimezone sets CustomTimezone field to given value.

### HasCustomTimezone

`func (o *PluginListResponseResourcesInner) HasCustomTimezone() bool`

HasCustomTimezone returns a boolean if a field has been set.

### GetExpirationOffset

`func (o *PluginListResponseResourcesInner) GetExpirationOffset() string`

GetExpirationOffset returns the ExpirationOffset field if non-nil, zero value otherwise.

### GetExpirationOffsetOk

`func (o *PluginListResponseResourcesInner) GetExpirationOffsetOk() (*string, bool)`

GetExpirationOffsetOk returns a tuple with the ExpirationOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationOffset

`func (o *PluginListResponseResourcesInner) SetExpirationOffset(v string)`

SetExpirationOffset sets ExpirationOffset field to given value.


### GetPurgeBehavior

`func (o *PluginListResponseResourcesInner) GetPurgeBehavior() EnumpluginPurgeBehaviorProp`

GetPurgeBehavior returns the PurgeBehavior field if non-nil, zero value otherwise.

### GetPurgeBehaviorOk

`func (o *PluginListResponseResourcesInner) GetPurgeBehaviorOk() (*EnumpluginPurgeBehaviorProp, bool)`

GetPurgeBehaviorOk returns a tuple with the PurgeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeBehavior

`func (o *PluginListResponseResourcesInner) SetPurgeBehavior(v EnumpluginPurgeBehaviorProp)`

SetPurgeBehavior sets PurgeBehavior field to given value.

### HasPurgeBehavior

`func (o *PluginListResponseResourcesInner) HasPurgeBehavior() bool`

HasPurgeBehavior returns a boolean if a field has been set.

### GetNumMostExpensivePhasesShown

`func (o *PluginListResponseResourcesInner) GetNumMostExpensivePhasesShown() int64`

GetNumMostExpensivePhasesShown returns the NumMostExpensivePhasesShown field if non-nil, zero value otherwise.

### GetNumMostExpensivePhasesShownOk

`func (o *PluginListResponseResourcesInner) GetNumMostExpensivePhasesShownOk() (*int64, bool)`

GetNumMostExpensivePhasesShownOk returns a tuple with the NumMostExpensivePhasesShown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMostExpensivePhasesShown

`func (o *PluginListResponseResourcesInner) SetNumMostExpensivePhasesShown(v int64)`

SetNumMostExpensivePhasesShown sets NumMostExpensivePhasesShown field to given value.


### GetExtensionClass

`func (o *PluginListResponseResourcesInner) GetExtensionClass() string`

GetExtensionClass returns the ExtensionClass field if non-nil, zero value otherwise.

### GetExtensionClassOk

`func (o *PluginListResponseResourcesInner) GetExtensionClassOk() (*string, bool)`

GetExtensionClassOk returns a tuple with the ExtensionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionClass

`func (o *PluginListResponseResourcesInner) SetExtensionClass(v string)`

SetExtensionClass sets ExtensionClass field to given value.


### GetExtensionArgument

`func (o *PluginListResponseResourcesInner) GetExtensionArgument() []string`

GetExtensionArgument returns the ExtensionArgument field if non-nil, zero value otherwise.

### GetExtensionArgumentOk

`func (o *PluginListResponseResourcesInner) GetExtensionArgumentOk() (*[]string, bool)`

GetExtensionArgumentOk returns a tuple with the ExtensionArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionArgument

`func (o *PluginListResponseResourcesInner) SetExtensionArgument(v []string)`

SetExtensionArgument sets ExtensionArgument field to given value.

### HasExtensionArgument

`func (o *PluginListResponseResourcesInner) HasExtensionArgument() bool`

HasExtensionArgument returns a boolean if a field has been set.

### GetEncryptionSettingsDefinitionID

`func (o *PluginListResponseResourcesInner) GetEncryptionSettingsDefinitionID() string`

GetEncryptionSettingsDefinitionID returns the EncryptionSettingsDefinitionID field if non-nil, zero value otherwise.

### GetEncryptionSettingsDefinitionIDOk

`func (o *PluginListResponseResourcesInner) GetEncryptionSettingsDefinitionIDOk() (*string, bool)`

GetEncryptionSettingsDefinitionIDOk returns a tuple with the EncryptionSettingsDefinitionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSettingsDefinitionID

`func (o *PluginListResponseResourcesInner) SetEncryptionSettingsDefinitionID(v string)`

SetEncryptionSettingsDefinitionID sets EncryptionSettingsDefinitionID field to given value.

### HasEncryptionSettingsDefinitionID

`func (o *PluginListResponseResourcesInner) HasEncryptionSettingsDefinitionID() bool`

HasEncryptionSettingsDefinitionID returns a boolean if a field has been set.

### GetServer

`func (o *PluginListResponseResourcesInner) GetServer() []string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PluginListResponseResourcesInner) GetServerOk() (*[]string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PluginListResponseResourcesInner) SetServer(v []string)`

SetServer sets Server field to given value.


### GetServerAccessMode

`func (o *PluginListResponseResourcesInner) GetServerAccessMode() EnumpluginServerAccessModeProp`

GetServerAccessMode returns the ServerAccessMode field if non-nil, zero value otherwise.

### GetServerAccessModeOk

`func (o *PluginListResponseResourcesInner) GetServerAccessModeOk() (*EnumpluginServerAccessModeProp, bool)`

GetServerAccessModeOk returns a tuple with the ServerAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAccessMode

`func (o *PluginListResponseResourcesInner) SetServerAccessMode(v EnumpluginServerAccessModeProp)`

SetServerAccessMode sets ServerAccessMode field to given value.


### GetDnMap

`func (o *PluginListResponseResourcesInner) GetDnMap() []string`

GetDnMap returns the DnMap field if non-nil, zero value otherwise.

### GetDnMapOk

`func (o *PluginListResponseResourcesInner) GetDnMapOk() (*[]string, bool)`

GetDnMapOk returns a tuple with the DnMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnMap

`func (o *PluginListResponseResourcesInner) SetDnMap(v []string)`

SetDnMap sets DnMap field to given value.

### HasDnMap

`func (o *PluginListResponseResourcesInner) HasDnMap() bool`

HasDnMap returns a boolean if a field has been set.

### GetBindDNPattern

`func (o *PluginListResponseResourcesInner) GetBindDNPattern() string`

GetBindDNPattern returns the BindDNPattern field if non-nil, zero value otherwise.

### GetBindDNPatternOk

`func (o *PluginListResponseResourcesInner) GetBindDNPatternOk() (*string, bool)`

GetBindDNPatternOk returns a tuple with the BindDNPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDNPattern

`func (o *PluginListResponseResourcesInner) SetBindDNPattern(v string)`

SetBindDNPattern sets BindDNPattern field to given value.

### HasBindDNPattern

`func (o *PluginListResponseResourcesInner) HasBindDNPattern() bool`

HasBindDNPattern returns a boolean if a field has been set.

### GetSearchBaseDN

`func (o *PluginListResponseResourcesInner) GetSearchBaseDN() string`

GetSearchBaseDN returns the SearchBaseDN field if non-nil, zero value otherwise.

### GetSearchBaseDNOk

`func (o *PluginListResponseResourcesInner) GetSearchBaseDNOk() (*string, bool)`

GetSearchBaseDNOk returns a tuple with the SearchBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDN

`func (o *PluginListResponseResourcesInner) SetSearchBaseDN(v string)`

SetSearchBaseDN sets SearchBaseDN field to given value.

### HasSearchBaseDN

`func (o *PluginListResponseResourcesInner) HasSearchBaseDN() bool`

HasSearchBaseDN returns a boolean if a field has been set.

### GetSearchFilterPattern

`func (o *PluginListResponseResourcesInner) GetSearchFilterPattern() string`

GetSearchFilterPattern returns the SearchFilterPattern field if non-nil, zero value otherwise.

### GetSearchFilterPatternOk

`func (o *PluginListResponseResourcesInner) GetSearchFilterPatternOk() (*string, bool)`

GetSearchFilterPatternOk returns a tuple with the SearchFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilterPattern

`func (o *PluginListResponseResourcesInner) SetSearchFilterPattern(v string)`

SetSearchFilterPattern sets SearchFilterPattern field to given value.

### HasSearchFilterPattern

`func (o *PluginListResponseResourcesInner) HasSearchFilterPattern() bool`

HasSearchFilterPattern returns a boolean if a field has been set.

### GetInitialConnections

`func (o *PluginListResponseResourcesInner) GetInitialConnections() int64`

GetInitialConnections returns the InitialConnections field if non-nil, zero value otherwise.

### GetInitialConnectionsOk

`func (o *PluginListResponseResourcesInner) GetInitialConnectionsOk() (*int64, bool)`

GetInitialConnectionsOk returns a tuple with the InitialConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialConnections

`func (o *PluginListResponseResourcesInner) SetInitialConnections(v int64)`

SetInitialConnections sets InitialConnections field to given value.


### GetMaxConnections

`func (o *PluginListResponseResourcesInner) GetMaxConnections() int64`

GetMaxConnections returns the MaxConnections field if non-nil, zero value otherwise.

### GetMaxConnectionsOk

`func (o *PluginListResponseResourcesInner) GetMaxConnectionsOk() (*int64, bool)`

GetMaxConnectionsOk returns a tuple with the MaxConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConnections

`func (o *PluginListResponseResourcesInner) SetMaxConnections(v int64)`

SetMaxConnections sets MaxConnections field to given value.


### GetSourceDN

`func (o *PluginListResponseResourcesInner) GetSourceDN() string`

GetSourceDN returns the SourceDN field if non-nil, zero value otherwise.

### GetSourceDNOk

`func (o *PluginListResponseResourcesInner) GetSourceDNOk() (*string, bool)`

GetSourceDNOk returns a tuple with the SourceDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDN

`func (o *PluginListResponseResourcesInner) SetSourceDN(v string)`

SetSourceDN sets SourceDN field to given value.


### GetTargetDN

`func (o *PluginListResponseResourcesInner) GetTargetDN() string`

GetTargetDN returns the TargetDN field if non-nil, zero value otherwise.

### GetTargetDNOk

`func (o *PluginListResponseResourcesInner) GetTargetDNOk() (*string, bool)`

GetTargetDNOk returns a tuple with the TargetDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDN

`func (o *PluginListResponseResourcesInner) SetTargetDN(v string)`

SetTargetDN sets TargetDN field to given value.


### GetEnableAttributeMapping

`func (o *PluginListResponseResourcesInner) GetEnableAttributeMapping() bool`

GetEnableAttributeMapping returns the EnableAttributeMapping field if non-nil, zero value otherwise.

### GetEnableAttributeMappingOk

`func (o *PluginListResponseResourcesInner) GetEnableAttributeMappingOk() (*bool, bool)`

GetEnableAttributeMappingOk returns a tuple with the EnableAttributeMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAttributeMapping

`func (o *PluginListResponseResourcesInner) SetEnableAttributeMapping(v bool)`

SetEnableAttributeMapping sets EnableAttributeMapping field to given value.


### GetMapAttribute

`func (o *PluginListResponseResourcesInner) GetMapAttribute() []string`

GetMapAttribute returns the MapAttribute field if non-nil, zero value otherwise.

### GetMapAttributeOk

`func (o *PluginListResponseResourcesInner) GetMapAttributeOk() (*[]string, bool)`

GetMapAttributeOk returns a tuple with the MapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapAttribute

`func (o *PluginListResponseResourcesInner) SetMapAttribute(v []string)`

SetMapAttribute sets MapAttribute field to given value.

### HasMapAttribute

`func (o *PluginListResponseResourcesInner) HasMapAttribute() bool`

HasMapAttribute returns a boolean if a field has been set.

### GetEnableControlMapping

`func (o *PluginListResponseResourcesInner) GetEnableControlMapping() bool`

GetEnableControlMapping returns the EnableControlMapping field if non-nil, zero value otherwise.

### GetEnableControlMappingOk

`func (o *PluginListResponseResourcesInner) GetEnableControlMappingOk() (*bool, bool)`

GetEnableControlMappingOk returns a tuple with the EnableControlMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableControlMapping

`func (o *PluginListResponseResourcesInner) SetEnableControlMapping(v bool)`

SetEnableControlMapping sets EnableControlMapping field to given value.


### GetAlwaysMapResponses

`func (o *PluginListResponseResourcesInner) GetAlwaysMapResponses() bool`

GetAlwaysMapResponses returns the AlwaysMapResponses field if non-nil, zero value otherwise.

### GetAlwaysMapResponsesOk

`func (o *PluginListResponseResourcesInner) GetAlwaysMapResponsesOk() (*bool, bool)`

GetAlwaysMapResponsesOk returns a tuple with the AlwaysMapResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysMapResponses

`func (o *PluginListResponseResourcesInner) SetAlwaysMapResponses(v bool)`

SetAlwaysMapResponses sets AlwaysMapResponses field to given value.


### GetRetainFilesSparselyByAge

`func (o *PluginListResponseResourcesInner) GetRetainFilesSparselyByAge() bool`

GetRetainFilesSparselyByAge returns the RetainFilesSparselyByAge field if non-nil, zero value otherwise.

### GetRetainFilesSparselyByAgeOk

`func (o *PluginListResponseResourcesInner) GetRetainFilesSparselyByAgeOk() (*bool, bool)`

GetRetainFilesSparselyByAgeOk returns a tuple with the RetainFilesSparselyByAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetainFilesSparselyByAge

`func (o *PluginListResponseResourcesInner) SetRetainFilesSparselyByAge(v bool)`

SetRetainFilesSparselyByAge sets RetainFilesSparselyByAge field to given value.

### HasRetainFilesSparselyByAge

`func (o *PluginListResponseResourcesInner) HasRetainFilesSparselyByAge() bool`

HasRetainFilesSparselyByAge returns a boolean if a field has been set.

### GetSanitize

`func (o *PluginListResponseResourcesInner) GetSanitize() bool`

GetSanitize returns the Sanitize field if non-nil, zero value otherwise.

### GetSanitizeOk

`func (o *PluginListResponseResourcesInner) GetSanitizeOk() (*bool, bool)`

GetSanitizeOk returns a tuple with the Sanitize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSanitize

`func (o *PluginListResponseResourcesInner) SetSanitize(v bool)`

SetSanitize sets Sanitize field to given value.

### HasSanitize

`func (o *PluginListResponseResourcesInner) HasSanitize() bool`

HasSanitize returns a boolean if a field has been set.

### GetReferralBaseURL

`func (o *PluginListResponseResourcesInner) GetReferralBaseURL() []string`

GetReferralBaseURL returns the ReferralBaseURL field if non-nil, zero value otherwise.

### GetReferralBaseURLOk

`func (o *PluginListResponseResourcesInner) GetReferralBaseURLOk() (*[]string, bool)`

GetReferralBaseURLOk returns a tuple with the ReferralBaseURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralBaseURL

`func (o *PluginListResponseResourcesInner) SetReferralBaseURL(v []string)`

SetReferralBaseURL sets ReferralBaseURL field to given value.


### GetContextName

`func (o *PluginListResponseResourcesInner) GetContextName() string`

GetContextName returns the ContextName field if non-nil, zero value otherwise.

### GetContextNameOk

`func (o *PluginListResponseResourcesInner) GetContextNameOk() (*string, bool)`

GetContextNameOk returns a tuple with the ContextName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextName

`func (o *PluginListResponseResourcesInner) SetContextName(v string)`

SetContextName sets ContextName field to given value.

### HasContextName

`func (o *PluginListResponseResourcesInner) HasContextName() bool`

HasContextName returns a boolean if a field has been set.

### GetAgentxAddress

`func (o *PluginListResponseResourcesInner) GetAgentxAddress() string`

GetAgentxAddress returns the AgentxAddress field if non-nil, zero value otherwise.

### GetAgentxAddressOk

`func (o *PluginListResponseResourcesInner) GetAgentxAddressOk() (*string, bool)`

GetAgentxAddressOk returns a tuple with the AgentxAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxAddress

`func (o *PluginListResponseResourcesInner) SetAgentxAddress(v string)`

SetAgentxAddress sets AgentxAddress field to given value.


### GetAgentxPort

`func (o *PluginListResponseResourcesInner) GetAgentxPort() int64`

GetAgentxPort returns the AgentxPort field if non-nil, zero value otherwise.

### GetAgentxPortOk

`func (o *PluginListResponseResourcesInner) GetAgentxPortOk() (*int64, bool)`

GetAgentxPortOk returns a tuple with the AgentxPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxPort

`func (o *PluginListResponseResourcesInner) SetAgentxPort(v int64)`

SetAgentxPort sets AgentxPort field to given value.


### GetNumWorkerThreads

`func (o *PluginListResponseResourcesInner) GetNumWorkerThreads() int64`

GetNumWorkerThreads returns the NumWorkerThreads field if non-nil, zero value otherwise.

### GetNumWorkerThreadsOk

`func (o *PluginListResponseResourcesInner) GetNumWorkerThreadsOk() (*int64, bool)`

GetNumWorkerThreadsOk returns a tuple with the NumWorkerThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumWorkerThreads

`func (o *PluginListResponseResourcesInner) SetNumWorkerThreads(v int64)`

SetNumWorkerThreads sets NumWorkerThreads field to given value.


### GetSessionTimeout

`func (o *PluginListResponseResourcesInner) GetSessionTimeout() string`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *PluginListResponseResourcesInner) GetSessionTimeoutOk() (*string, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *PluginListResponseResourcesInner) SetSessionTimeout(v string)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *PluginListResponseResourcesInner) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### GetConnectRetryMaxWait

`func (o *PluginListResponseResourcesInner) GetConnectRetryMaxWait() string`

GetConnectRetryMaxWait returns the ConnectRetryMaxWait field if non-nil, zero value otherwise.

### GetConnectRetryMaxWaitOk

`func (o *PluginListResponseResourcesInner) GetConnectRetryMaxWaitOk() (*string, bool)`

GetConnectRetryMaxWaitOk returns a tuple with the ConnectRetryMaxWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectRetryMaxWait

`func (o *PluginListResponseResourcesInner) SetConnectRetryMaxWait(v string)`

SetConnectRetryMaxWait sets ConnectRetryMaxWait field to given value.

### HasConnectRetryMaxWait

`func (o *PluginListResponseResourcesInner) HasConnectRetryMaxWait() bool`

HasConnectRetryMaxWait returns a boolean if a field has been set.

### GetPingInterval

`func (o *PluginListResponseResourcesInner) GetPingInterval() string`

GetPingInterval returns the PingInterval field if non-nil, zero value otherwise.

### GetPingIntervalOk

`func (o *PluginListResponseResourcesInner) GetPingIntervalOk() (*string, bool)`

GetPingIntervalOk returns a tuple with the PingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingInterval

`func (o *PluginListResponseResourcesInner) SetPingInterval(v string)`

SetPingInterval sets PingInterval field to given value.

### HasPingInterval

`func (o *PluginListResponseResourcesInner) HasPingInterval() bool`

HasPingInterval returns a boolean if a field has been set.

### GetAllowedRequestControl

`func (o *PluginListResponseResourcesInner) GetAllowedRequestControl() []string`

GetAllowedRequestControl returns the AllowedRequestControl field if non-nil, zero value otherwise.

### GetAllowedRequestControlOk

`func (o *PluginListResponseResourcesInner) GetAllowedRequestControlOk() (*[]string, bool)`

GetAllowedRequestControlOk returns a tuple with the AllowedRequestControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedRequestControl

`func (o *PluginListResponseResourcesInner) SetAllowedRequestControl(v []string)`

SetAllowedRequestControl sets AllowedRequestControl field to given value.

### HasAllowedRequestControl

`func (o *PluginListResponseResourcesInner) HasAllowedRequestControl() bool`

HasAllowedRequestControl returns a boolean if a field has been set.

### GetDefaultUserPasswordStorageScheme

`func (o *PluginListResponseResourcesInner) GetDefaultUserPasswordStorageScheme() []string`

GetDefaultUserPasswordStorageScheme returns the DefaultUserPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDefaultUserPasswordStorageSchemeOk

`func (o *PluginListResponseResourcesInner) GetDefaultUserPasswordStorageSchemeOk() (*[]string, bool)`

GetDefaultUserPasswordStorageSchemeOk returns a tuple with the DefaultUserPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserPasswordStorageScheme

`func (o *PluginListResponseResourcesInner) SetDefaultUserPasswordStorageScheme(v []string)`

SetDefaultUserPasswordStorageScheme sets DefaultUserPasswordStorageScheme field to given value.

### HasDefaultUserPasswordStorageScheme

`func (o *PluginListResponseResourcesInner) HasDefaultUserPasswordStorageScheme() bool`

HasDefaultUserPasswordStorageScheme returns a boolean if a field has been set.

### GetDefaultAuthPasswordStorageScheme

`func (o *PluginListResponseResourcesInner) GetDefaultAuthPasswordStorageScheme() []string`

GetDefaultAuthPasswordStorageScheme returns the DefaultAuthPasswordStorageScheme field if non-nil, zero value otherwise.

### GetDefaultAuthPasswordStorageSchemeOk

`func (o *PluginListResponseResourcesInner) GetDefaultAuthPasswordStorageSchemeOk() (*[]string, bool)`

GetDefaultAuthPasswordStorageSchemeOk returns a tuple with the DefaultAuthPasswordStorageScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAuthPasswordStorageScheme

`func (o *PluginListResponseResourcesInner) SetDefaultAuthPasswordStorageScheme(v []string)`

SetDefaultAuthPasswordStorageScheme sets DefaultAuthPasswordStorageScheme field to given value.

### HasDefaultAuthPasswordStorageScheme

`func (o *PluginListResponseResourcesInner) HasDefaultAuthPasswordStorageScheme() bool`

HasDefaultAuthPasswordStorageScheme returns a boolean if a field has been set.

### GetProfileSampleInterval

`func (o *PluginListResponseResourcesInner) GetProfileSampleInterval() string`

GetProfileSampleInterval returns the ProfileSampleInterval field if non-nil, zero value otherwise.

### GetProfileSampleIntervalOk

`func (o *PluginListResponseResourcesInner) GetProfileSampleIntervalOk() (*string, bool)`

GetProfileSampleIntervalOk returns a tuple with the ProfileSampleInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileSampleInterval

`func (o *PluginListResponseResourcesInner) SetProfileSampleInterval(v string)`

SetProfileSampleInterval sets ProfileSampleInterval field to given value.


### GetProfileDirectory

`func (o *PluginListResponseResourcesInner) GetProfileDirectory() string`

GetProfileDirectory returns the ProfileDirectory field if non-nil, zero value otherwise.

### GetProfileDirectoryOk

`func (o *PluginListResponseResourcesInner) GetProfileDirectoryOk() (*string, bool)`

GetProfileDirectoryOk returns a tuple with the ProfileDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileDirectory

`func (o *PluginListResponseResourcesInner) SetProfileDirectory(v string)`

SetProfileDirectory sets ProfileDirectory field to given value.


### GetEnableProfilingOnStartup

`func (o *PluginListResponseResourcesInner) GetEnableProfilingOnStartup() bool`

GetEnableProfilingOnStartup returns the EnableProfilingOnStartup field if non-nil, zero value otherwise.

### GetEnableProfilingOnStartupOk

`func (o *PluginListResponseResourcesInner) GetEnableProfilingOnStartupOk() (*bool, bool)`

GetEnableProfilingOnStartupOk returns a tuple with the EnableProfilingOnStartup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProfilingOnStartup

`func (o *PluginListResponseResourcesInner) SetEnableProfilingOnStartup(v bool)`

SetEnableProfilingOnStartup sets EnableProfilingOnStartup field to given value.


### GetProfileAction

`func (o *PluginListResponseResourcesInner) GetProfileAction() EnumpluginProfileActionProp`

GetProfileAction returns the ProfileAction field if non-nil, zero value otherwise.

### GetProfileActionOk

`func (o *PluginListResponseResourcesInner) GetProfileActionOk() (*EnumpluginProfileActionProp, bool)`

GetProfileActionOk returns a tuple with the ProfileAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAction

`func (o *PluginListResponseResourcesInner) SetProfileAction(v EnumpluginProfileActionProp)`

SetProfileAction sets ProfileAction field to given value.

### HasProfileAction

`func (o *PluginListResponseResourcesInner) HasProfileAction() bool`

HasProfileAction returns a boolean if a field has been set.

### GetValuePattern

`func (o *PluginListResponseResourcesInner) GetValuePattern() []string`

GetValuePattern returns the ValuePattern field if non-nil, zero value otherwise.

### GetValuePatternOk

`func (o *PluginListResponseResourcesInner) GetValuePatternOk() (*[]string, bool)`

GetValuePatternOk returns a tuple with the ValuePattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuePattern

`func (o *PluginListResponseResourcesInner) SetValuePattern(v []string)`

SetValuePattern sets ValuePattern field to given value.


### GetMultipleValuePatternBehavior

`func (o *PluginListResponseResourcesInner) GetMultipleValuePatternBehavior() EnumpluginMultipleValuePatternBehaviorProp`

GetMultipleValuePatternBehavior returns the MultipleValuePatternBehavior field if non-nil, zero value otherwise.

### GetMultipleValuePatternBehaviorOk

`func (o *PluginListResponseResourcesInner) GetMultipleValuePatternBehaviorOk() (*EnumpluginMultipleValuePatternBehaviorProp, bool)`

GetMultipleValuePatternBehaviorOk returns a tuple with the MultipleValuePatternBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleValuePatternBehavior

`func (o *PluginListResponseResourcesInner) SetMultipleValuePatternBehavior(v EnumpluginMultipleValuePatternBehaviorProp)`

SetMultipleValuePatternBehavior sets MultipleValuePatternBehavior field to given value.

### HasMultipleValuePatternBehavior

`func (o *PluginListResponseResourcesInner) HasMultipleValuePatternBehavior() bool`

HasMultipleValuePatternBehavior returns a boolean if a field has been set.

### GetMultiValuedAttributeBehavior

`func (o *PluginListResponseResourcesInner) GetMultiValuedAttributeBehavior() EnumpluginMultiValuedAttributeBehaviorProp`

GetMultiValuedAttributeBehavior returns the MultiValuedAttributeBehavior field if non-nil, zero value otherwise.

### GetMultiValuedAttributeBehaviorOk

`func (o *PluginListResponseResourcesInner) GetMultiValuedAttributeBehaviorOk() (*EnumpluginMultiValuedAttributeBehaviorProp, bool)`

GetMultiValuedAttributeBehaviorOk returns a tuple with the MultiValuedAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValuedAttributeBehavior

`func (o *PluginListResponseResourcesInner) SetMultiValuedAttributeBehavior(v EnumpluginMultiValuedAttributeBehaviorProp)`

SetMultiValuedAttributeBehavior sets MultiValuedAttributeBehavior field to given value.

### HasMultiValuedAttributeBehavior

`func (o *PluginListResponseResourcesInner) HasMultiValuedAttributeBehavior() bool`

HasMultiValuedAttributeBehavior returns a boolean if a field has been set.

### GetTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *PluginListResponseResourcesInner) GetTargetAttributeExistsDuringInitialPopulationBehavior() EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp`

GetTargetAttributeExistsDuringInitialPopulationBehavior returns the TargetAttributeExistsDuringInitialPopulationBehavior field if non-nil, zero value otherwise.

### GetTargetAttributeExistsDuringInitialPopulationBehaviorOk

`func (o *PluginListResponseResourcesInner) GetTargetAttributeExistsDuringInitialPopulationBehaviorOk() (*EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp, bool)`

GetTargetAttributeExistsDuringInitialPopulationBehaviorOk returns a tuple with the TargetAttributeExistsDuringInitialPopulationBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *PluginListResponseResourcesInner) SetTargetAttributeExistsDuringInitialPopulationBehavior(v EnumpluginTargetAttributeExistsDuringInitialPopulationBehaviorProp)`

SetTargetAttributeExistsDuringInitialPopulationBehavior sets TargetAttributeExistsDuringInitialPopulationBehavior field to given value.

### HasTargetAttributeExistsDuringInitialPopulationBehavior

`func (o *PluginListResponseResourcesInner) HasTargetAttributeExistsDuringInitialPopulationBehavior() bool`

HasTargetAttributeExistsDuringInitialPopulationBehavior returns a boolean if a field has been set.

### GetUpdateSourceAttributeBehavior

`func (o *PluginListResponseResourcesInner) GetUpdateSourceAttributeBehavior() EnumpluginUpdateSourceAttributeBehaviorProp`

GetUpdateSourceAttributeBehavior returns the UpdateSourceAttributeBehavior field if non-nil, zero value otherwise.

### GetUpdateSourceAttributeBehaviorOk

`func (o *PluginListResponseResourcesInner) GetUpdateSourceAttributeBehaviorOk() (*EnumpluginUpdateSourceAttributeBehaviorProp, bool)`

GetUpdateSourceAttributeBehaviorOk returns a tuple with the UpdateSourceAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateSourceAttributeBehavior

`func (o *PluginListResponseResourcesInner) SetUpdateSourceAttributeBehavior(v EnumpluginUpdateSourceAttributeBehaviorProp)`

SetUpdateSourceAttributeBehavior sets UpdateSourceAttributeBehavior field to given value.

### HasUpdateSourceAttributeBehavior

`func (o *PluginListResponseResourcesInner) HasUpdateSourceAttributeBehavior() bool`

HasUpdateSourceAttributeBehavior returns a boolean if a field has been set.

### GetSourceAttributeRemovalBehavior

`func (o *PluginListResponseResourcesInner) GetSourceAttributeRemovalBehavior() EnumpluginSourceAttributeRemovalBehaviorProp`

GetSourceAttributeRemovalBehavior returns the SourceAttributeRemovalBehavior field if non-nil, zero value otherwise.

### GetSourceAttributeRemovalBehaviorOk

`func (o *PluginListResponseResourcesInner) GetSourceAttributeRemovalBehaviorOk() (*EnumpluginSourceAttributeRemovalBehaviorProp, bool)`

GetSourceAttributeRemovalBehaviorOk returns a tuple with the SourceAttributeRemovalBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttributeRemovalBehavior

`func (o *PluginListResponseResourcesInner) SetSourceAttributeRemovalBehavior(v EnumpluginSourceAttributeRemovalBehaviorProp)`

SetSourceAttributeRemovalBehavior sets SourceAttributeRemovalBehavior field to given value.

### HasSourceAttributeRemovalBehavior

`func (o *PluginListResponseResourcesInner) HasSourceAttributeRemovalBehavior() bool`

HasSourceAttributeRemovalBehavior returns a boolean if a field has been set.

### GetUpdateTargetAttributeBehavior

`func (o *PluginListResponseResourcesInner) GetUpdateTargetAttributeBehavior() EnumpluginUpdateTargetAttributeBehaviorProp`

GetUpdateTargetAttributeBehavior returns the UpdateTargetAttributeBehavior field if non-nil, zero value otherwise.

### GetUpdateTargetAttributeBehaviorOk

`func (o *PluginListResponseResourcesInner) GetUpdateTargetAttributeBehaviorOk() (*EnumpluginUpdateTargetAttributeBehaviorProp, bool)`

GetUpdateTargetAttributeBehaviorOk returns a tuple with the UpdateTargetAttributeBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTargetAttributeBehavior

`func (o *PluginListResponseResourcesInner) SetUpdateTargetAttributeBehavior(v EnumpluginUpdateTargetAttributeBehaviorProp)`

SetUpdateTargetAttributeBehavior sets UpdateTargetAttributeBehavior field to given value.

### HasUpdateTargetAttributeBehavior

`func (o *PluginListResponseResourcesInner) HasUpdateTargetAttributeBehavior() bool`

HasUpdateTargetAttributeBehavior returns a boolean if a field has been set.

### GetIncludeBaseDN

`func (o *PluginListResponseResourcesInner) GetIncludeBaseDN() []string`

GetIncludeBaseDN returns the IncludeBaseDN field if non-nil, zero value otherwise.

### GetIncludeBaseDNOk

`func (o *PluginListResponseResourcesInner) GetIncludeBaseDNOk() (*[]string, bool)`

GetIncludeBaseDNOk returns a tuple with the IncludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBaseDN

`func (o *PluginListResponseResourcesInner) SetIncludeBaseDN(v []string)`

SetIncludeBaseDN sets IncludeBaseDN field to given value.

### HasIncludeBaseDN

`func (o *PluginListResponseResourcesInner) HasIncludeBaseDN() bool`

HasIncludeBaseDN returns a boolean if a field has been set.

### GetExcludeBaseDN

`func (o *PluginListResponseResourcesInner) GetExcludeBaseDN() []string`

GetExcludeBaseDN returns the ExcludeBaseDN field if non-nil, zero value otherwise.

### GetExcludeBaseDNOk

`func (o *PluginListResponseResourcesInner) GetExcludeBaseDNOk() (*[]string, bool)`

GetExcludeBaseDNOk returns a tuple with the ExcludeBaseDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBaseDN

`func (o *PluginListResponseResourcesInner) SetExcludeBaseDN(v []string)`

SetExcludeBaseDN sets ExcludeBaseDN field to given value.

### HasExcludeBaseDN

`func (o *PluginListResponseResourcesInner) HasExcludeBaseDN() bool`

HasExcludeBaseDN returns a boolean if a field has been set.

### GetIncludeFilter

`func (o *PluginListResponseResourcesInner) GetIncludeFilter() []string`

GetIncludeFilter returns the IncludeFilter field if non-nil, zero value otherwise.

### GetIncludeFilterOk

`func (o *PluginListResponseResourcesInner) GetIncludeFilterOk() (*[]string, bool)`

GetIncludeFilterOk returns a tuple with the IncludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilter

`func (o *PluginListResponseResourcesInner) SetIncludeFilter(v []string)`

SetIncludeFilter sets IncludeFilter field to given value.

### HasIncludeFilter

`func (o *PluginListResponseResourcesInner) HasIncludeFilter() bool`

HasIncludeFilter returns a boolean if a field has been set.

### GetExcludeFilter

`func (o *PluginListResponseResourcesInner) GetExcludeFilter() []string`

GetExcludeFilter returns the ExcludeFilter field if non-nil, zero value otherwise.

### GetExcludeFilterOk

`func (o *PluginListResponseResourcesInner) GetExcludeFilterOk() (*[]string, bool)`

GetExcludeFilterOk returns a tuple with the ExcludeFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFilter

`func (o *PluginListResponseResourcesInner) SetExcludeFilter(v []string)`

SetExcludeFilter sets ExcludeFilter field to given value.

### HasExcludeFilter

`func (o *PluginListResponseResourcesInner) HasExcludeFilter() bool`

HasExcludeFilter returns a boolean if a field has been set.

### GetUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *PluginListResponseResourcesInner) GetUpdatedEntryNewlyMatchesCriteriaBehavior() EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp`

GetUpdatedEntryNewlyMatchesCriteriaBehavior returns the UpdatedEntryNewlyMatchesCriteriaBehavior field if non-nil, zero value otherwise.

### GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk

`func (o *PluginListResponseResourcesInner) GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp, bool)`

GetUpdatedEntryNewlyMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNewlyMatchesCriteriaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *PluginListResponseResourcesInner) SetUpdatedEntryNewlyMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNewlyMatchesCriteriaBehaviorProp)`

SetUpdatedEntryNewlyMatchesCriteriaBehavior sets UpdatedEntryNewlyMatchesCriteriaBehavior field to given value.

### HasUpdatedEntryNewlyMatchesCriteriaBehavior

`func (o *PluginListResponseResourcesInner) HasUpdatedEntryNewlyMatchesCriteriaBehavior() bool`

HasUpdatedEntryNewlyMatchesCriteriaBehavior returns a boolean if a field has been set.

### GetUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *PluginListResponseResourcesInner) GetUpdatedEntryNoLongerMatchesCriteriaBehavior() EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp`

GetUpdatedEntryNoLongerMatchesCriteriaBehavior returns the UpdatedEntryNoLongerMatchesCriteriaBehavior field if non-nil, zero value otherwise.

### GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk

`func (o *PluginListResponseResourcesInner) GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk() (*EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp, bool)`

GetUpdatedEntryNoLongerMatchesCriteriaBehaviorOk returns a tuple with the UpdatedEntryNoLongerMatchesCriteriaBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *PluginListResponseResourcesInner) SetUpdatedEntryNoLongerMatchesCriteriaBehavior(v EnumpluginUpdatedEntryNoLongerMatchesCriteriaBehaviorProp)`

SetUpdatedEntryNoLongerMatchesCriteriaBehavior sets UpdatedEntryNoLongerMatchesCriteriaBehavior field to given value.

### HasUpdatedEntryNoLongerMatchesCriteriaBehavior

`func (o *PluginListResponseResourcesInner) HasUpdatedEntryNoLongerMatchesCriteriaBehavior() bool`

HasUpdatedEntryNoLongerMatchesCriteriaBehavior returns a boolean if a field has been set.

### GetSourceAttribute

`func (o *PluginListResponseResourcesInner) GetSourceAttribute() string`

GetSourceAttribute returns the SourceAttribute field if non-nil, zero value otherwise.

### GetSourceAttributeOk

`func (o *PluginListResponseResourcesInner) GetSourceAttributeOk() (*string, bool)`

GetSourceAttributeOk returns a tuple with the SourceAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAttribute

`func (o *PluginListResponseResourcesInner) SetSourceAttribute(v string)`

SetSourceAttribute sets SourceAttribute field to given value.


### GetTargetAttribute

`func (o *PluginListResponseResourcesInner) GetTargetAttribute() string`

GetTargetAttribute returns the TargetAttribute field if non-nil, zero value otherwise.

### GetTargetAttributeOk

`func (o *PluginListResponseResourcesInner) GetTargetAttributeOk() (*string, bool)`

GetTargetAttributeOk returns a tuple with the TargetAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttribute

`func (o *PluginListResponseResourcesInner) SetTargetAttribute(v string)`

SetTargetAttribute sets TargetAttribute field to given value.


### GetDelay

`func (o *PluginListResponseResourcesInner) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *PluginListResponseResourcesInner) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *PluginListResponseResourcesInner) SetDelay(v string)`

SetDelay sets Delay field to given value.


### GetScriptClass

`func (o *PluginListResponseResourcesInner) GetScriptClass() string`

GetScriptClass returns the ScriptClass field if non-nil, zero value otherwise.

### GetScriptClassOk

`func (o *PluginListResponseResourcesInner) GetScriptClassOk() (*string, bool)`

GetScriptClassOk returns a tuple with the ScriptClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptClass

`func (o *PluginListResponseResourcesInner) SetScriptClass(v string)`

SetScriptClass sets ScriptClass field to given value.


### GetScriptArgument

`func (o *PluginListResponseResourcesInner) GetScriptArgument() []string`

GetScriptArgument returns the ScriptArgument field if non-nil, zero value otherwise.

### GetScriptArgumentOk

`func (o *PluginListResponseResourcesInner) GetScriptArgumentOk() (*[]string, bool)`

GetScriptArgumentOk returns a tuple with the ScriptArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptArgument

`func (o *PluginListResponseResourcesInner) SetScriptArgument(v []string)`

SetScriptArgument sets ScriptArgument field to given value.

### HasScriptArgument

`func (o *PluginListResponseResourcesInner) HasScriptArgument() bool`

HasScriptArgument returns a boolean if a field has been set.

### GetExcludeAttribute

`func (o *PluginListResponseResourcesInner) GetExcludeAttribute() []string`

GetExcludeAttribute returns the ExcludeAttribute field if non-nil, zero value otherwise.

### GetExcludeAttributeOk

`func (o *PluginListResponseResourcesInner) GetExcludeAttributeOk() (*[]string, bool)`

GetExcludeAttributeOk returns a tuple with the ExcludeAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAttribute

`func (o *PluginListResponseResourcesInner) SetExcludeAttribute(v []string)`

SetExcludeAttribute sets ExcludeAttribute field to given value.

### HasExcludeAttribute

`func (o *PluginListResponseResourcesInner) HasExcludeAttribute() bool`

HasExcludeAttribute returns a boolean if a field has been set.

### GetPassThroughAuthenticationHandler

`func (o *PluginListResponseResourcesInner) GetPassThroughAuthenticationHandler() string`

GetPassThroughAuthenticationHandler returns the PassThroughAuthenticationHandler field if non-nil, zero value otherwise.

### GetPassThroughAuthenticationHandlerOk

`func (o *PluginListResponseResourcesInner) GetPassThroughAuthenticationHandlerOk() (*string, bool)`

GetPassThroughAuthenticationHandlerOk returns a tuple with the PassThroughAuthenticationHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassThroughAuthenticationHandler

`func (o *PluginListResponseResourcesInner) SetPassThroughAuthenticationHandler(v string)`

SetPassThroughAuthenticationHandler sets PassThroughAuthenticationHandler field to given value.


### GetUpdateInterval

`func (o *PluginListResponseResourcesInner) GetUpdateInterval() string`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *PluginListResponseResourcesInner) GetUpdateIntervalOk() (*string, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *PluginListResponseResourcesInner) SetUpdateInterval(v string)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *PluginListResponseResourcesInner) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.

### GetListenAddress

`func (o *PluginListResponseResourcesInner) GetListenAddress() string`

GetListenAddress returns the ListenAddress field if non-nil, zero value otherwise.

### GetListenAddressOk

`func (o *PluginListResponseResourcesInner) GetListenAddressOk() (*string, bool)`

GetListenAddressOk returns a tuple with the ListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddress

`func (o *PluginListResponseResourcesInner) SetListenAddress(v string)`

SetListenAddress sets ListenAddress field to given value.


### GetListenPort

`func (o *PluginListResponseResourcesInner) GetListenPort() int64`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *PluginListResponseResourcesInner) GetListenPortOk() (*int64, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *PluginListResponseResourcesInner) SetListenPort(v int64)`

SetListenPort sets ListenPort field to given value.


### GetAgentxTransport

`func (o *PluginListResponseResourcesInner) GetAgentxTransport() EnumpluginAgentxTransportProp`

GetAgentxTransport returns the AgentxTransport field if non-nil, zero value otherwise.

### GetAgentxTransportOk

`func (o *PluginListResponseResourcesInner) GetAgentxTransportOk() (*EnumpluginAgentxTransportProp, bool)`

GetAgentxTransportOk returns a tuple with the AgentxTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxTransport

`func (o *PluginListResponseResourcesInner) SetAgentxTransport(v EnumpluginAgentxTransportProp)`

SetAgentxTransport sets AgentxTransport field to given value.


### GetAgentxListenAddress

`func (o *PluginListResponseResourcesInner) GetAgentxListenAddress() string`

GetAgentxListenAddress returns the AgentxListenAddress field if non-nil, zero value otherwise.

### GetAgentxListenAddressOk

`func (o *PluginListResponseResourcesInner) GetAgentxListenAddressOk() (*string, bool)`

GetAgentxListenAddressOk returns a tuple with the AgentxListenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxListenAddress

`func (o *PluginListResponseResourcesInner) SetAgentxListenAddress(v string)`

SetAgentxListenAddress sets AgentxListenAddress field to given value.


### GetAgentxListenPort

`func (o *PluginListResponseResourcesInner) GetAgentxListenPort() int64`

GetAgentxListenPort returns the AgentxListenPort field if non-nil, zero value otherwise.

### GetAgentxListenPortOk

`func (o *PluginListResponseResourcesInner) GetAgentxListenPortOk() (*int64, bool)`

GetAgentxListenPortOk returns a tuple with the AgentxListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentxListenPort

`func (o *PluginListResponseResourcesInner) SetAgentxListenPort(v int64)`

SetAgentxListenPort sets AgentxListenPort field to given value.


### GetNotificationTargetAddress

`func (o *PluginListResponseResourcesInner) GetNotificationTargetAddress() string`

GetNotificationTargetAddress returns the NotificationTargetAddress field if non-nil, zero value otherwise.

### GetNotificationTargetAddressOk

`func (o *PluginListResponseResourcesInner) GetNotificationTargetAddressOk() (*string, bool)`

GetNotificationTargetAddressOk returns a tuple with the NotificationTargetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTargetAddress

`func (o *PluginListResponseResourcesInner) SetNotificationTargetAddress(v string)`

SetNotificationTargetAddress sets NotificationTargetAddress field to given value.

### HasNotificationTargetAddress

`func (o *PluginListResponseResourcesInner) HasNotificationTargetAddress() bool`

HasNotificationTargetAddress returns a boolean if a field has been set.

### GetNotificationTargetPort

`func (o *PluginListResponseResourcesInner) GetNotificationTargetPort() int64`

GetNotificationTargetPort returns the NotificationTargetPort field if non-nil, zero value otherwise.

### GetNotificationTargetPortOk

`func (o *PluginListResponseResourcesInner) GetNotificationTargetPortOk() (*int64, bool)`

GetNotificationTargetPortOk returns a tuple with the NotificationTargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTargetPort

`func (o *PluginListResponseResourcesInner) SetNotificationTargetPort(v int64)`

SetNotificationTargetPort sets NotificationTargetPort field to given value.

### HasNotificationTargetPort

`func (o *PluginListResponseResourcesInner) HasNotificationTargetPort() bool`

HasNotificationTargetPort returns a boolean if a field has been set.

### GetAgentSNMPVersion

`func (o *PluginListResponseResourcesInner) GetAgentSNMPVersion() []EnumpluginAgentSNMPVersionProp`

GetAgentSNMPVersion returns the AgentSNMPVersion field if non-nil, zero value otherwise.

### GetAgentSNMPVersionOk

`func (o *PluginListResponseResourcesInner) GetAgentSNMPVersionOk() (*[]EnumpluginAgentSNMPVersionProp, bool)`

GetAgentSNMPVersionOk returns a tuple with the AgentSNMPVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSNMPVersion

`func (o *PluginListResponseResourcesInner) SetAgentSNMPVersion(v []EnumpluginAgentSNMPVersionProp)`

SetAgentSNMPVersion sets AgentSNMPVersion field to given value.


### GetCommunityName

`func (o *PluginListResponseResourcesInner) GetCommunityName() string`

GetCommunityName returns the CommunityName field if non-nil, zero value otherwise.

### GetCommunityNameOk

`func (o *PluginListResponseResourcesInner) GetCommunityNameOk() (*string, bool)`

GetCommunityNameOk returns a tuple with the CommunityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityName

`func (o *PluginListResponseResourcesInner) SetCommunityName(v string)`

SetCommunityName sets CommunityName field to given value.


### GetAgentSecurityName

`func (o *PluginListResponseResourcesInner) GetAgentSecurityName() string`

GetAgentSecurityName returns the AgentSecurityName field if non-nil, zero value otherwise.

### GetAgentSecurityNameOk

`func (o *PluginListResponseResourcesInner) GetAgentSecurityNameOk() (*string, bool)`

GetAgentSecurityNameOk returns a tuple with the AgentSecurityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSecurityName

`func (o *PluginListResponseResourcesInner) SetAgentSecurityName(v string)`

SetAgentSecurityName sets AgentSecurityName field to given value.

### HasAgentSecurityName

`func (o *PluginListResponseResourcesInner) HasAgentSecurityName() bool`

HasAgentSecurityName returns a boolean if a field has been set.

### GetAgentSecurityLevel

`func (o *PluginListResponseResourcesInner) GetAgentSecurityLevel() EnumpluginAgentSecurityLevelProp`

GetAgentSecurityLevel returns the AgentSecurityLevel field if non-nil, zero value otherwise.

### GetAgentSecurityLevelOk

`func (o *PluginListResponseResourcesInner) GetAgentSecurityLevelOk() (*EnumpluginAgentSecurityLevelProp, bool)`

GetAgentSecurityLevelOk returns a tuple with the AgentSecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentSecurityLevel

`func (o *PluginListResponseResourcesInner) SetAgentSecurityLevel(v EnumpluginAgentSecurityLevelProp)`

SetAgentSecurityLevel sets AgentSecurityLevel field to given value.

### HasAgentSecurityLevel

`func (o *PluginListResponseResourcesInner) HasAgentSecurityLevel() bool`

HasAgentSecurityLevel returns a boolean if a field has been set.

### GetAgentAuthenticationProtocol

`func (o *PluginListResponseResourcesInner) GetAgentAuthenticationProtocol() EnumpluginAgentAuthenticationProtocolProp`

GetAgentAuthenticationProtocol returns the AgentAuthenticationProtocol field if non-nil, zero value otherwise.

### GetAgentAuthenticationProtocolOk

`func (o *PluginListResponseResourcesInner) GetAgentAuthenticationProtocolOk() (*EnumpluginAgentAuthenticationProtocolProp, bool)`

GetAgentAuthenticationProtocolOk returns a tuple with the AgentAuthenticationProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAuthenticationProtocol

`func (o *PluginListResponseResourcesInner) SetAgentAuthenticationProtocol(v EnumpluginAgentAuthenticationProtocolProp)`

SetAgentAuthenticationProtocol sets AgentAuthenticationProtocol field to given value.

### HasAgentAuthenticationProtocol

`func (o *PluginListResponseResourcesInner) HasAgentAuthenticationProtocol() bool`

HasAgentAuthenticationProtocol returns a boolean if a field has been set.

### GetAgentAuthenticationPassphrase

`func (o *PluginListResponseResourcesInner) GetAgentAuthenticationPassphrase() string`

GetAgentAuthenticationPassphrase returns the AgentAuthenticationPassphrase field if non-nil, zero value otherwise.

### GetAgentAuthenticationPassphraseOk

`func (o *PluginListResponseResourcesInner) GetAgentAuthenticationPassphraseOk() (*string, bool)`

GetAgentAuthenticationPassphraseOk returns a tuple with the AgentAuthenticationPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAuthenticationPassphrase

`func (o *PluginListResponseResourcesInner) SetAgentAuthenticationPassphrase(v string)`

SetAgentAuthenticationPassphrase sets AgentAuthenticationPassphrase field to given value.

### HasAgentAuthenticationPassphrase

`func (o *PluginListResponseResourcesInner) HasAgentAuthenticationPassphrase() bool`

HasAgentAuthenticationPassphrase returns a boolean if a field has been set.

### GetAgentPrivacyProtocol

`func (o *PluginListResponseResourcesInner) GetAgentPrivacyProtocol() EnumpluginAgentPrivacyProtocolProp`

GetAgentPrivacyProtocol returns the AgentPrivacyProtocol field if non-nil, zero value otherwise.

### GetAgentPrivacyProtocolOk

`func (o *PluginListResponseResourcesInner) GetAgentPrivacyProtocolOk() (*EnumpluginAgentPrivacyProtocolProp, bool)`

GetAgentPrivacyProtocolOk returns a tuple with the AgentPrivacyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPrivacyProtocol

`func (o *PluginListResponseResourcesInner) SetAgentPrivacyProtocol(v EnumpluginAgentPrivacyProtocolProp)`

SetAgentPrivacyProtocol sets AgentPrivacyProtocol field to given value.

### HasAgentPrivacyProtocol

`func (o *PluginListResponseResourcesInner) HasAgentPrivacyProtocol() bool`

HasAgentPrivacyProtocol returns a boolean if a field has been set.

### GetAgentPrivacyPassphrase

`func (o *PluginListResponseResourcesInner) GetAgentPrivacyPassphrase() string`

GetAgentPrivacyPassphrase returns the AgentPrivacyPassphrase field if non-nil, zero value otherwise.

### GetAgentPrivacyPassphraseOk

`func (o *PluginListResponseResourcesInner) GetAgentPrivacyPassphraseOk() (*string, bool)`

GetAgentPrivacyPassphraseOk returns a tuple with the AgentPrivacyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentPrivacyPassphrase

`func (o *PluginListResponseResourcesInner) SetAgentPrivacyPassphrase(v string)`

SetAgentPrivacyPassphrase sets AgentPrivacyPassphrase field to given value.

### HasAgentPrivacyPassphrase

`func (o *PluginListResponseResourcesInner) HasAgentPrivacyPassphrase() bool`

HasAgentPrivacyPassphrase returns a boolean if a field has been set.

### GetPreventAddingMembersToNonexistentGroups

`func (o *PluginListResponseResourcesInner) GetPreventAddingMembersToNonexistentGroups() bool`

GetPreventAddingMembersToNonexistentGroups returns the PreventAddingMembersToNonexistentGroups field if non-nil, zero value otherwise.

### GetPreventAddingMembersToNonexistentGroupsOk

`func (o *PluginListResponseResourcesInner) GetPreventAddingMembersToNonexistentGroupsOk() (*bool, bool)`

GetPreventAddingMembersToNonexistentGroupsOk returns a tuple with the PreventAddingMembersToNonexistentGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAddingMembersToNonexistentGroups

`func (o *PluginListResponseResourcesInner) SetPreventAddingMembersToNonexistentGroups(v bool)`

SetPreventAddingMembersToNonexistentGroups sets PreventAddingMembersToNonexistentGroups field to given value.

### HasPreventAddingMembersToNonexistentGroups

`func (o *PluginListResponseResourcesInner) HasPreventAddingMembersToNonexistentGroups() bool`

HasPreventAddingMembersToNonexistentGroups returns a boolean if a field has been set.

### GetPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *PluginListResponseResourcesInner) GetPreventAddingGroupsAsInvertedStaticGroupMembers() bool`

GetPreventAddingGroupsAsInvertedStaticGroupMembers returns the PreventAddingGroupsAsInvertedStaticGroupMembers field if non-nil, zero value otherwise.

### GetPreventAddingGroupsAsInvertedStaticGroupMembersOk

`func (o *PluginListResponseResourcesInner) GetPreventAddingGroupsAsInvertedStaticGroupMembersOk() (*bool, bool)`

GetPreventAddingGroupsAsInvertedStaticGroupMembersOk returns a tuple with the PreventAddingGroupsAsInvertedStaticGroupMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *PluginListResponseResourcesInner) SetPreventAddingGroupsAsInvertedStaticGroupMembers(v bool)`

SetPreventAddingGroupsAsInvertedStaticGroupMembers sets PreventAddingGroupsAsInvertedStaticGroupMembers field to given value.

### HasPreventAddingGroupsAsInvertedStaticGroupMembers

`func (o *PluginListResponseResourcesInner) HasPreventAddingGroupsAsInvertedStaticGroupMembers() bool`

HasPreventAddingGroupsAsInvertedStaticGroupMembers returns a boolean if a field has been set.

### GetPreventNestingNonexistentGroups

`func (o *PluginListResponseResourcesInner) GetPreventNestingNonexistentGroups() bool`

GetPreventNestingNonexistentGroups returns the PreventNestingNonexistentGroups field if non-nil, zero value otherwise.

### GetPreventNestingNonexistentGroupsOk

`func (o *PluginListResponseResourcesInner) GetPreventNestingNonexistentGroupsOk() (*bool, bool)`

GetPreventNestingNonexistentGroupsOk returns a tuple with the PreventNestingNonexistentGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventNestingNonexistentGroups

`func (o *PluginListResponseResourcesInner) SetPreventNestingNonexistentGroups(v bool)`

SetPreventNestingNonexistentGroups sets PreventNestingNonexistentGroups field to given value.

### HasPreventNestingNonexistentGroups

`func (o *PluginListResponseResourcesInner) HasPreventNestingNonexistentGroups() bool`

HasPreventNestingNonexistentGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


