/*
PingData Config - OpenAPI 3.0

This is the PingData Configuration API

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// EnumalertHandlerEnabledAlertTypeProp Specifies the names of the alert types that are enabled for this alert handler.
type EnumalertHandlerEnabledAlertTypeProp string

// List of Enumalert-handler-enabledAlertTypeProp
const (
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCESS_CONTROL_CHANGE EnumalertHandlerEnabledAlertTypeProp = "access-control-change"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCESS_CONTROL_DISABLED EnumalertHandlerEnabledAlertTypeProp = "access-control-disabled"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCESS_CONTROL_ENABLED EnumalertHandlerEnabledAlertTypeProp = "access-control-enabled"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCESS_CONTROL_PARSE_FAILURE EnumalertHandlerEnabledAlertTypeProp = "access-control-parse-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCESS_LOG_CRITERIA_MATCHED EnumalertHandlerEnabledAlertTypeProp = "access-log-criteria-matched"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ALARM_CLEARED EnumalertHandlerEnabledAlertTypeProp = "alarm-cleared"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ALARM_CRITICAL EnumalertHandlerEnabledAlertTypeProp = "alarm-critical"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ALARM_MAJOR EnumalertHandlerEnabledAlertTypeProp = "alarm-major"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ALARM_MINOR EnumalertHandlerEnabledAlertTypeProp = "alarm-minor"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ALARM_WARNING EnumalertHandlerEnabledAlertTypeProp = "alarm-warning"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_BACKEND_DISABLED EnumalertHandlerEnabledAlertTypeProp = "backend-disabled"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_BACKEND_INITIALIZATION_FAILED EnumalertHandlerEnabledAlertTypeProp = "backend-initialization-failed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_BACKUP_FAILED EnumalertHandlerEnabledAlertTypeProp = "backup-failed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_ACQUIRE_SHARED_BACKEND_LOCK EnumalertHandlerEnabledAlertTypeProp = "cannot-acquire-shared-backend-lock"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_COPY_SCHEMA_FILES EnumalertHandlerEnabledAlertTypeProp = "cannot-copy-schema-files"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_DECODE_ENTRY EnumalertHandlerEnabledAlertTypeProp = "cannot-decode-entry"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_FIND_RECURRING_TASK EnumalertHandlerEnabledAlertTypeProp = "cannot-find-recurring-task"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_REGISTER_BACKEND EnumalertHandlerEnabledAlertTypeProp = "cannot-register-backend"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_RELEASE_SHARED_BACKEND_LOCK EnumalertHandlerEnabledAlertTypeProp = "cannot-release-shared-backend-lock"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_RENAME_CURRENT_TASK_FILE EnumalertHandlerEnabledAlertTypeProp = "cannot-rename-current-task-file"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_RENAME_NEW_TASK_FILE EnumalertHandlerEnabledAlertTypeProp = "cannot-rename-new-task-file"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_RESTORE_BACKUP EnumalertHandlerEnabledAlertTypeProp = "cannot-restore-backup"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_SCHEDULE_RECURRING_TASK_ITERATION EnumalertHandlerEnabledAlertTypeProp = "cannot-schedule-recurring-task-iteration"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_WRITE_CONFIGURATION EnumalertHandlerEnabledAlertTypeProp = "cannot-write-configuration"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_WRITE_NEW_SCHEMA_FILES EnumalertHandlerEnabledAlertTypeProp = "cannot-write-new-schema-files"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_WRITE_SERVER_STATE_FILE EnumalertHandlerEnabledAlertTypeProp = "cannot-write-server-state-file"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CANNOT_WRITE_TASK_BACKING_FILE EnumalertHandlerEnabledAlertTypeProp = "cannot-write-task-backing-file"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CONFIG_CHANGE EnumalertHandlerEnabledAlertTypeProp = "config-change"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CONSOLE_LOGGER_WITHOUT_NO_DETACH EnumalertHandlerEnabledAlertTypeProp = "console-logger-without-no-detach"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CRYPTO_MANAGER_ERROR EnumalertHandlerEnabledAlertTypeProp = "crypto-manager-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_CONTINUOUS_GARBAGE_COLLECTION_DETECTED EnumalertHandlerEnabledAlertTypeProp = "continuous-garbage-collection-detected"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_DEADLOCK_DETECTED EnumalertHandlerEnabledAlertTypeProp = "deadlock-detected"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_DEBUG_LOGGING_ENABLED EnumalertHandlerEnabledAlertTypeProp = "debug-logging-enabled"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_DELEGATED_ADMIN_CONFIGURATION_ERRORS EnumalertHandlerEnabledAlertTypeProp = "delegated-admin-configuration-errors"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_DUPLICATE_ALERTS_SUPPRESSED EnumalertHandlerEnabledAlertTypeProp = "duplicate-alerts-suppressed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_DUPLICATE_ERROR_ALERTS_SUPPRESSED EnumalertHandlerEnabledAlertTypeProp = "duplicate-error-alerts-suppressed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_DUPLICATE_FATAL_ALERTS_SUPPRESSED EnumalertHandlerEnabledAlertTypeProp = "duplicate-fatal-alerts-suppressed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_DUPLICATE_INFO_ALERTS_SUPPRESSED EnumalertHandlerEnabledAlertTypeProp = "duplicate-info-alerts-suppressed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_DUPLICATE_WARNING_ALERTS_SUPPRESSED EnumalertHandlerEnabledAlertTypeProp = "duplicate-warning-alerts-suppressed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_EMBEDDED_POSTGRESQL_UNAVAILABLE EnumalertHandlerEnabledAlertTypeProp = "embedded-postgresql-unavailable"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ENTERING_LOCKDOWN_MODE EnumalertHandlerEnabledAlertTypeProp = "entering-lockdown-mode"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ENTRY_REFERENCES_REMOVED_ATTRIBUTE_TYPE EnumalertHandlerEnabledAlertTypeProp = "entry-references-removed-attribute-type"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_EXEC_TASK_LAUNCHING_COMMAND EnumalertHandlerEnabledAlertTypeProp = "exec-task-launching-command"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_EXPLODED_INDEX_BACKGROUND_DELETE_CLEANUP_FAILED EnumalertHandlerEnabledAlertTypeProp = "exploded-index-background-delete-cleanup-failed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_EXPLODED_INDEX_BACKGROUND_DELETE_FAILED EnumalertHandlerEnabledAlertTypeProp = "exploded-index-background-delete-failed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_EXTERNAL_CONFIG_FILE_EDIT_HANDLED EnumalertHandlerEnabledAlertTypeProp = "external-config-file-edit-handled"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_EXTERNAL_CONFIG_FILE_EDIT_LOST EnumalertHandlerEnabledAlertTypeProp = "external-config-file-edit-lost"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_EXTERNAL_SERVER_INITIALIZATION_FAILED EnumalertHandlerEnabledAlertTypeProp = "external-server-initialization-failed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_FAILED_TO_APPLY_MIRRORED_CONFIGURATION EnumalertHandlerEnabledAlertTypeProp = "failed-to-apply-mirrored-configuration"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_FILE_RETENTION_TASK_DELETE_FAILURE EnumalertHandlerEnabledAlertTypeProp = "file-retention-task-delete-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_FORCE_GC_COMPLETE EnumalertHandlerEnabledAlertTypeProp = "force-gc-complete"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_FORCE_GC_STARTING EnumalertHandlerEnabledAlertTypeProp = "force-gc-starting"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_GLOBAL_INDEX_INSUFFICIENT_DISK_SPACE_ERROR EnumalertHandlerEnabledAlertTypeProp = "global-index-insufficient-disk-space-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_GLOBAL_INDEX_PERSISTENCE_ERROR EnumalertHandlerEnabledAlertTypeProp = "global-index-persistence-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_GLOBAL_INDEX_READ_ERROR EnumalertHandlerEnabledAlertTypeProp = "global-index-read-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_GLOBAL_REFERENTIAL_INTEGRITY_UPDATE_FAILURE EnumalertHandlerEnabledAlertTypeProp = "global-referential-integrity-update-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_GLOBALLY_UNIQUE_ATTRIBUTE_CONFLICT EnumalertHandlerEnabledAlertTypeProp = "globally-unique-attribute-conflict"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_HEALTH_CHECK_AVAILABLE_TO_DEGRADED EnumalertHandlerEnabledAlertTypeProp = "health-check-available-to-degraded"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_HEALTH_CHECK_AVAILABLE_TO_UNAVAILABLE EnumalertHandlerEnabledAlertTypeProp = "health-check-available-to-unavailable"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_HEALTH_CHECK_DEGRADED_TO_AVAILABLE EnumalertHandlerEnabledAlertTypeProp = "health-check-degraded-to-available"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_HEALTH_CHECK_DEGRADED_TO_UNAVAILABLE EnumalertHandlerEnabledAlertTypeProp = "health-check-degraded-to-unavailable"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_HEALTH_CHECK_UNAVAILABLE_TO_AVAILABLE EnumalertHandlerEnabledAlertTypeProp = "health-check-unavailable-to-available"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_HEALTH_CHECK_UNAVAILABLE_TO_DEGRADED EnumalertHandlerEnabledAlertTypeProp = "health-check-unavailable-to-degraded"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_HTTP_CONNECTION_HANDLER_DUPLICATE_CONTEXT_PATH EnumalertHandlerEnabledAlertTypeProp = "http-connection-handler-duplicate-context-path"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_HTTP_CONNECTION_HANDLER_DUPLICATE_SERVLET_EXTENSION EnumalertHandlerEnabledAlertTypeProp = "http-connection-handler-duplicate-servlet-extension"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_INDEX_CORRUPT EnumalertHandlerEnabledAlertTypeProp = "index-corrupt"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_INDEX_DEGRADED EnumalertHandlerEnabledAlertTypeProp = "index-degraded"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_INDEX_REBUILD_COMPLETED EnumalertHandlerEnabledAlertTypeProp = "index-rebuild-completed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_INDEX_REBUILD_IN_PROGRESS EnumalertHandlerEnabledAlertTypeProp = "index-rebuild-in-progress"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_INSECURE_ACCESS_TOKEN_VALIDATOR_ENABLED EnumalertHandlerEnabledAlertTypeProp = "insecure-access-token-validator-enabled"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_INVALID_PRIVILEGE EnumalertHandlerEnabledAlertTypeProp = "invalid-privilege"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_JE_BACKGROUND_SYNC_FAILED EnumalertHandlerEnabledAlertTypeProp = "je-background-sync-failed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_JE_CLEANER_DISABLED EnumalertHandlerEnabledAlertTypeProp = "je-cleaner-disabled"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_JE_DAEMON_THREAD_EXCEPTION EnumalertHandlerEnabledAlertTypeProp = "je-daemon-thread-exception"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_JE_ENVIRONMENT_NOT_CLOSED_CLEANLY EnumalertHandlerEnabledAlertTypeProp = "je-environment-not-closed-cleanly"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_JE_RECOVERY_REQUIRED EnumalertHandlerEnabledAlertTypeProp = "je-recovery-required"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_JVM_MISCONFIGURATION EnumalertHandlerEnabledAlertTypeProp = "jvm-misconfiguration"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LARGE_ATTRIBUTE_UPDATE_FAILURE EnumalertHandlerEnabledAlertTypeProp = "large-attribute-update-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LBA_NO_AVAILABLE_SERVERS EnumalertHandlerEnabledAlertTypeProp = "lba-no-available-servers"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LDAP_CONNECTION_HANDLER_CANNOT_LISTEN EnumalertHandlerEnabledAlertTypeProp = "ldap-connection-handler-cannot-listen"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LDAP_CONNECTION_HANDLER_CONSECUTIVE_FAILURES EnumalertHandlerEnabledAlertTypeProp = "ldap-connection-handler-consecutive-failures"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LDAP_CONNECTION_HANDLER_UNCAUGHT_ERROR EnumalertHandlerEnabledAlertTypeProp = "ldap-connection-handler-uncaught-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LDIF_BACKEND_CANNOT_WRITE EnumalertHandlerEnabledAlertTypeProp = "ldif-backend-cannot-write"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LDIF_CONNECTION_HANDLER_PARSE_ERROR EnumalertHandlerEnabledAlertTypeProp = "ldif-connection-handler-parse-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LDIF_CONNECTION_HANDLER_IO_ERROR EnumalertHandlerEnabledAlertTypeProp = "ldif-connection-handler-io-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LEAVING_LOCKDOWN_MODE EnumalertHandlerEnabledAlertTypeProp = "leaving-lockdown-mode"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LOG_FILE_ROTATION_LISTENER_INVOKE_ERROR EnumalertHandlerEnabledAlertTypeProp = "log-file-rotation-listener-invoke-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LOG_FILE_ROTATION_LISTENER_PROCESSING_TAKES_TOO_LONG EnumalertHandlerEnabledAlertTypeProp = "log-file-rotation-listener-processing-takes-too-long"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LOGGING_ERROR EnumalertHandlerEnabledAlertTypeProp = "logging-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LOW_DISK_SPACE_ERROR EnumalertHandlerEnabledAlertTypeProp = "low-disk-space-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_LOW_DISK_SPACE_WARNING EnumalertHandlerEnabledAlertTypeProp = "low-disk-space-warning"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MIRRORED_SUBTREE_MANAGER_FORCED_AS_MASTER_ERROR EnumalertHandlerEnabledAlertTypeProp = "mirrored-subtree-manager-forced-as-master-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MIRRORED_SUBTREE_MANAGER_FORCED_AS_MASTER_WARNING EnumalertHandlerEnabledAlertTypeProp = "mirrored-subtree-manager-forced-as-master-warning"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MIRRORED_SUBTREE_MANAGER_NO_MASTER_FOUND EnumalertHandlerEnabledAlertTypeProp = "mirrored-subtree-manager-no-master-found"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MIRRORED_SUBTREE_SERVER_NOT_IN_TOPOLOGY EnumalertHandlerEnabledAlertTypeProp = "mirrored-subtree-server-not-in-topology"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MIRRORED_SUBTREE_MANAGER_OPERATION_ERROR EnumalertHandlerEnabledAlertTypeProp = "mirrored-subtree-manager-operation-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MIRRORED_SUBTREE_MANAGER_FAILED_OUTBOUND_CONNECTION EnumalertHandlerEnabledAlertTypeProp = "mirrored-subtree-manager-failed-outbound-connection"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MIRRORED_SUBTREE_MANAGER_CONNECTION_ASYMMETRY EnumalertHandlerEnabledAlertTypeProp = "mirrored-subtree-manager-connection-asymmetry"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MISSING_SCHEMA_ELEMENTS_REFERENCED_BY_BACKEND EnumalertHandlerEnabledAlertTypeProp = "missing-schema-elements-referenced-by-backend"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MONITORING_ENDPOINT_UNABLE_TO_CONNECT EnumalertHandlerEnabledAlertTypeProp = "monitoring-endpoint-unable-to-connect"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_NO_ENABLED_ALERT_HANDLERS EnumalertHandlerEnabledAlertTypeProp = "no-enabled-alert-handlers"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_OFFLINE_CONFIG_CHANGE_DETECTED EnumalertHandlerEnabledAlertTypeProp = "offline-config-change-detected"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_OUT_OF_DISK_SPACE_ERROR EnumalertHandlerEnabledAlertTypeProp = "out-of-disk-space-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PDP_UNAVAILABLE EnumalertHandlerEnabledAlertTypeProp = "pdp-unavailable"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PDP_TRUST_FRAMEWORK_VERSION_DEPRECATED EnumalertHandlerEnabledAlertTypeProp = "pdp-trust-framework-version-deprecated"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PROXY_ENTRY_BALANCING_OPERATION_FAILURE EnumalertHandlerEnabledAlertTypeProp = "proxy-entry-balancing-operation-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PROXY_ENTRY_BALANCING_ERROR_MULTIPLE_OPERATIONS_SUCCEEDED EnumalertHandlerEnabledAlertTypeProp = "proxy-entry-balancing-error-multiple-operations-succeeded"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PROXY_ENTRY_REBALANCING_ADMIN_ACTION_REQUIRED EnumalertHandlerEnabledAlertTypeProp = "proxy-entry-rebalancing-admin-action-required"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_BACKLOGGED EnumalertHandlerEnabledAlertTypeProp = "replication-backlogged"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_METADATA_DECODE_FAILURE EnumalertHandlerEnabledAlertTypeProp = "replication-metadata-decode-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_MISSING_CHANGES EnumalertHandlerEnabledAlertTypeProp = "replication-missing-changes"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_MONITOR_DATA_UNAVAILABLE EnumalertHandlerEnabledAlertTypeProp = "replication-monitor-data-unavailable"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_PLUGIN_MESSAGE_SERIALIZATION_FAILURE EnumalertHandlerEnabledAlertTypeProp = "replication-plugin-message-serialization-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_SERVER_CHANGELOG_FAILURE EnumalertHandlerEnabledAlertTypeProp = "replication-server-changelog-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_SERVER_LISTEN_FAILURE EnumalertHandlerEnabledAlertTypeProp = "replication-server-listen-failure"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_UNRESOLVED_CONFLICT EnumalertHandlerEnabledAlertTypeProp = "replication-unresolved-conflict"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_UNSENT_CHANGES EnumalertHandlerEnabledAlertTypeProp = "replication-unsent-changes"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_REPLAY_FAILED EnumalertHandlerEnabledAlertTypeProp = "replication-replay-failed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_RESTART_REQUIRED EnumalertHandlerEnabledAlertTypeProp = "restart-required"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SCHEMA_CHECKING_DISABLED EnumalertHandlerEnabledAlertTypeProp = "schema-checking-disabled"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SCIM_LOOKTHROUGH_LIMIT_EXCEEDED EnumalertHandlerEnabledAlertTypeProp = "scim-lookthrough-limit-exceeded"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_RESTRICTED_SUBTREE_ACCESSIBILITY EnumalertHandlerEnabledAlertTypeProp = "restricted-subtree-accessibility"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SERVER_SHUTTING_DOWN EnumalertHandlerEnabledAlertTypeProp = "server-shutting-down"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SERVER_STARTING EnumalertHandlerEnabledAlertTypeProp = "server-starting"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SERVER_STARTED EnumalertHandlerEnabledAlertTypeProp = "server-started"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SUBTREE_DELETE_INTERRUPTED EnumalertHandlerEnabledAlertTypeProp = "subtree-delete-interrupted"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SYNC_RESOURCE_CONNECTION_ERROR EnumalertHandlerEnabledAlertTypeProp = "sync-resource-connection-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SYNC_RESOURCE_OPERATION_ERROR EnumalertHandlerEnabledAlertTypeProp = "sync-resource-operation-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SYNC_PIPE_INITIALIZATION_ERROR EnumalertHandlerEnabledAlertTypeProp = "sync-pipe-initialization-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SYNC_PIPE_BACKLOG_ABOVE_THRESHOLD EnumalertHandlerEnabledAlertTypeProp = "sync-pipe-backlog-above-threshold"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SYNC_PIPE_BACKLOG_BELOW_THRESHOLD EnumalertHandlerEnabledAlertTypeProp = "sync-pipe-backlog-below-threshold"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SYSTEM_NANOTIME_STOPPED EnumalertHandlerEnabledAlertTypeProp = "system-nanotime-stopped"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SYSTEM_CURRENT_TIME_SHIFTED EnumalertHandlerEnabledAlertTypeProp = "system-current-time-shifted"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_TASK_STARTED EnumalertHandlerEnabledAlertTypeProp = "task-started"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_TASK_COMPLETED EnumalertHandlerEnabledAlertTypeProp = "task-completed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_TASK_FAILED EnumalertHandlerEnabledAlertTypeProp = "task-failed"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_THIRD_PARTY_EXTENSION_EXCEPTION EnumalertHandlerEnabledAlertTypeProp = "third-party-extension-exception"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_THREAD_EXIT_HOLDING_LOCK EnumalertHandlerEnabledAlertTypeProp = "thread-exit-holding-lock"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_THRESHOLD_WARNING_ENTRY EnumalertHandlerEnabledAlertTypeProp = "threshold-warning-entry"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_THRESHOLD_WARNING_EXIT EnumalertHandlerEnabledAlertTypeProp = "threshold-warning-exit"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_THRESHOLD_CRITICAL_ENTRY EnumalertHandlerEnabledAlertTypeProp = "threshold-critical-entry"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_THRESHOLD_CRITICAL_EXIT EnumalertHandlerEnabledAlertTypeProp = "threshold-critical-exit"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_TOPOLOGY_REGISTRY_SECRET_KEY_DELETED EnumalertHandlerEnabledAlertTypeProp = "topology-registry-secret-key-deleted"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_UNCAUGHT_EXCEPTION EnumalertHandlerEnabledAlertTypeProp = "uncaught-exception"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_UNINDEXED_INTERNAL_SEARCH EnumalertHandlerEnabledAlertTypeProp = "unindexed-internal-search"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_UNIQUE_ATTRIBUTE_SYNC_CONFLICT EnumalertHandlerEnabledAlertTypeProp = "unique-attribute-sync-conflict"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_UNIQUE_ATTRIBUTE_SYNC_ERROR EnumalertHandlerEnabledAlertTypeProp = "unique-attribute-sync-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_UNIQUENESS_CONTROL_POST_COMMIT_CONFLICT EnumalertHandlerEnabledAlertTypeProp = "uniqueness-control-post-commit-conflict"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_UNLICENSED_PRODUCT EnumalertHandlerEnabledAlertTypeProp = "unlicensed-product"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_UNRECOGNIZED_ALERT_TYPE EnumalertHandlerEnabledAlertTypeProp = "unrecognized-alert-type"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_USER_DEFINED_ERROR EnumalertHandlerEnabledAlertTypeProp = "user-defined-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_USER_DEFINED_FATAL EnumalertHandlerEnabledAlertTypeProp = "user-defined-fatal"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_USER_DEFINED_INFO EnumalertHandlerEnabledAlertTypeProp = "user-defined-info"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_USER_DEFINED_WARNING EnumalertHandlerEnabledAlertTypeProp = "user-defined-warning"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_WORKER_THREAD_CAUGHT_ERROR EnumalertHandlerEnabledAlertTypeProp = "worker-thread-caught-error"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_WORK_QUEUE_BACKLOGGED EnumalertHandlerEnabledAlertTypeProp = "work-queue-backlogged"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_WORK_QUEUE_FULL EnumalertHandlerEnabledAlertTypeProp = "work-queue-full"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_WORK_QUEUE_NO_THREADS_REMAINING EnumalertHandlerEnabledAlertTypeProp = "work-queue-no-threads-remaining"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SERVER_JVM_PAUSED EnumalertHandlerEnabledAlertTypeProp = "server-jvm-paused"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_SENSITIVE_TRACE_DATA_LOGGED_WARNING EnumalertHandlerEnabledAlertTypeProp = "sensitive-trace-data-logged-warning"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLICATION_GENERATION_ID_MISMATCH EnumalertHandlerEnabledAlertTypeProp = "replication-generation-id-mismatch"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_TEMPORARILY_LOCKED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-temporarily-locked-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_PERMANENTLY_LOCKED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-permanently-locked-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_UNLOCKED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-unlocked-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_IDLE_LOCKED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-idle-locked-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_RESET_LOCKED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-reset-locked-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_DISABLED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-disabled-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_ENABLED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-enabled-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_NOT_YET_ACTIVE_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-not-yet-active-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_EXPIRED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-expired-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PASSWORD_EXPIRED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "password-expired-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PASSWORD_EXPIRING_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "password-expiring-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PASSWORD_RESET_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "password-reset-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PASSWORD_CHANGED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "password-changed-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_CREATED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-created-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_ACCOUNT_UPDATED_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "account-updated-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_BIND_PASSWORD_FAILED_VALIDATION_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "bind-password-failed-validation-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_MUST_CHANGE_PASSWORD_ACCOUNT_STATUS_NOTIFICATION EnumalertHandlerEnabledAlertTypeProp = "must-change-password-account-status-notification"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_PRIVILEGE_ASSIGNED EnumalertHandlerEnabledAlertTypeProp = "privilege-assigned"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_INSECURE_REQUEST_REJECTED EnumalertHandlerEnabledAlertTypeProp = "insecure-request-rejected"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_INCREMENTAL_BACKUPS_DEPRECATED EnumalertHandlerEnabledAlertTypeProp = "incremental-backups-deprecated"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLACE_CERTIFICATE_SUCCEEDED EnumalertHandlerEnabledAlertTypeProp = "replace-certificate-succeeded"
	ENUMALERTHANDLERENABLEDALERTTYPEPROP_REPLACE_CERTIFICATE_FAILED EnumalertHandlerEnabledAlertTypeProp = "replace-certificate-failed"
)

// All allowed values of EnumalertHandlerEnabledAlertTypeProp enum
var AllowedEnumalertHandlerEnabledAlertTypePropEnumValues = []EnumalertHandlerEnabledAlertTypeProp{
	"access-control-change",
	"access-control-disabled",
	"access-control-enabled",
	"access-control-parse-failure",
	"access-log-criteria-matched",
	"alarm-cleared",
	"alarm-critical",
	"alarm-major",
	"alarm-minor",
	"alarm-warning",
	"backend-disabled",
	"backend-initialization-failed",
	"backup-failed",
	"cannot-acquire-shared-backend-lock",
	"cannot-copy-schema-files",
	"cannot-decode-entry",
	"cannot-find-recurring-task",
	"cannot-register-backend",
	"cannot-release-shared-backend-lock",
	"cannot-rename-current-task-file",
	"cannot-rename-new-task-file",
	"cannot-restore-backup",
	"cannot-schedule-recurring-task-iteration",
	"cannot-write-configuration",
	"cannot-write-new-schema-files",
	"cannot-write-server-state-file",
	"cannot-write-task-backing-file",
	"config-change",
	"console-logger-without-no-detach",
	"crypto-manager-error",
	"continuous-garbage-collection-detected",
	"deadlock-detected",
	"debug-logging-enabled",
	"delegated-admin-configuration-errors",
	"duplicate-alerts-suppressed",
	"duplicate-error-alerts-suppressed",
	"duplicate-fatal-alerts-suppressed",
	"duplicate-info-alerts-suppressed",
	"duplicate-warning-alerts-suppressed",
	"embedded-postgresql-unavailable",
	"entering-lockdown-mode",
	"entry-references-removed-attribute-type",
	"exec-task-launching-command",
	"exploded-index-background-delete-cleanup-failed",
	"exploded-index-background-delete-failed",
	"external-config-file-edit-handled",
	"external-config-file-edit-lost",
	"external-server-initialization-failed",
	"failed-to-apply-mirrored-configuration",
	"file-retention-task-delete-failure",
	"force-gc-complete",
	"force-gc-starting",
	"global-index-insufficient-disk-space-error",
	"global-index-persistence-error",
	"global-index-read-error",
	"global-referential-integrity-update-failure",
	"globally-unique-attribute-conflict",
	"health-check-available-to-degraded",
	"health-check-available-to-unavailable",
	"health-check-degraded-to-available",
	"health-check-degraded-to-unavailable",
	"health-check-unavailable-to-available",
	"health-check-unavailable-to-degraded",
	"http-connection-handler-duplicate-context-path",
	"http-connection-handler-duplicate-servlet-extension",
	"index-corrupt",
	"index-degraded",
	"index-rebuild-completed",
	"index-rebuild-in-progress",
	"insecure-access-token-validator-enabled",
	"invalid-privilege",
	"je-background-sync-failed",
	"je-cleaner-disabled",
	"je-daemon-thread-exception",
	"je-environment-not-closed-cleanly",
	"je-recovery-required",
	"jvm-misconfiguration",
	"large-attribute-update-failure",
	"lba-no-available-servers",
	"ldap-connection-handler-cannot-listen",
	"ldap-connection-handler-consecutive-failures",
	"ldap-connection-handler-uncaught-error",
	"ldif-backend-cannot-write",
	"ldif-connection-handler-parse-error",
	"ldif-connection-handler-io-error",
	"leaving-lockdown-mode",
	"log-file-rotation-listener-invoke-error",
	"log-file-rotation-listener-processing-takes-too-long",
	"logging-error",
	"low-disk-space-error",
	"low-disk-space-warning",
	"mirrored-subtree-manager-forced-as-master-error",
	"mirrored-subtree-manager-forced-as-master-warning",
	"mirrored-subtree-manager-no-master-found",
	"mirrored-subtree-server-not-in-topology",
	"mirrored-subtree-manager-operation-error",
	"mirrored-subtree-manager-failed-outbound-connection",
	"mirrored-subtree-manager-connection-asymmetry",
	"missing-schema-elements-referenced-by-backend",
	"monitoring-endpoint-unable-to-connect",
	"no-enabled-alert-handlers",
	"offline-config-change-detected",
	"out-of-disk-space-error",
	"pdp-unavailable",
	"pdp-trust-framework-version-deprecated",
	"proxy-entry-balancing-operation-failure",
	"proxy-entry-balancing-error-multiple-operations-succeeded",
	"proxy-entry-rebalancing-admin-action-required",
	"replication-backlogged",
	"replication-metadata-decode-failure",
	"replication-missing-changes",
	"replication-monitor-data-unavailable",
	"replication-plugin-message-serialization-failure",
	"replication-server-changelog-failure",
	"replication-server-listen-failure",
	"replication-unresolved-conflict",
	"replication-unsent-changes",
	"replication-replay-failed",
	"restart-required",
	"schema-checking-disabled",
	"scim-lookthrough-limit-exceeded",
	"restricted-subtree-accessibility",
	"server-shutting-down",
	"server-starting",
	"server-started",
	"subtree-delete-interrupted",
	"sync-resource-connection-error",
	"sync-resource-operation-error",
	"sync-pipe-initialization-error",
	"sync-pipe-backlog-above-threshold",
	"sync-pipe-backlog-below-threshold",
	"system-nanotime-stopped",
	"system-current-time-shifted",
	"task-started",
	"task-completed",
	"task-failed",
	"third-party-extension-exception",
	"thread-exit-holding-lock",
	"threshold-warning-entry",
	"threshold-warning-exit",
	"threshold-critical-entry",
	"threshold-critical-exit",
	"topology-registry-secret-key-deleted",
	"uncaught-exception",
	"unindexed-internal-search",
	"unique-attribute-sync-conflict",
	"unique-attribute-sync-error",
	"uniqueness-control-post-commit-conflict",
	"unlicensed-product",
	"unrecognized-alert-type",
	"user-defined-error",
	"user-defined-fatal",
	"user-defined-info",
	"user-defined-warning",
	"worker-thread-caught-error",
	"work-queue-backlogged",
	"work-queue-full",
	"work-queue-no-threads-remaining",
	"server-jvm-paused",
	"sensitive-trace-data-logged-warning",
	"replication-generation-id-mismatch",
	"account-temporarily-locked-account-status-notification",
	"account-permanently-locked-account-status-notification",
	"account-unlocked-account-status-notification",
	"account-idle-locked-account-status-notification",
	"account-reset-locked-account-status-notification",
	"account-disabled-account-status-notification",
	"account-enabled-account-status-notification",
	"account-not-yet-active-account-status-notification",
	"account-expired-account-status-notification",
	"password-expired-account-status-notification",
	"password-expiring-account-status-notification",
	"password-reset-account-status-notification",
	"password-changed-account-status-notification",
	"account-created-account-status-notification",
	"account-updated-account-status-notification",
	"bind-password-failed-validation-account-status-notification",
	"must-change-password-account-status-notification",
	"privilege-assigned",
	"insecure-request-rejected",
	"incremental-backups-deprecated",
	"replace-certificate-succeeded",
	"replace-certificate-failed",
}

func (v *EnumalertHandlerEnabledAlertTypeProp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumalertHandlerEnabledAlertTypeProp(value)
	for _, existing := range AllowedEnumalertHandlerEnabledAlertTypePropEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumalertHandlerEnabledAlertTypeProp", value)
}

// NewEnumalertHandlerEnabledAlertTypePropFromValue returns a pointer to a valid EnumalertHandlerEnabledAlertTypeProp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumalertHandlerEnabledAlertTypePropFromValue(v string) (*EnumalertHandlerEnabledAlertTypeProp, error) {
	ev := EnumalertHandlerEnabledAlertTypeProp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumalertHandlerEnabledAlertTypeProp: valid values are %v", v, AllowedEnumalertHandlerEnabledAlertTypePropEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumalertHandlerEnabledAlertTypeProp) IsValid() bool {
	for _, existing := range AllowedEnumalertHandlerEnabledAlertTypePropEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Enumalert-handler-enabledAlertTypeProp value
func (v EnumalertHandlerEnabledAlertTypeProp) Ptr() *EnumalertHandlerEnabledAlertTypeProp {
	return &v
}

type NullableEnumalertHandlerEnabledAlertTypeProp struct {
	value *EnumalertHandlerEnabledAlertTypeProp
	isSet bool
}

func (v NullableEnumalertHandlerEnabledAlertTypeProp) Get() *EnumalertHandlerEnabledAlertTypeProp {
	return v.value
}

func (v *NullableEnumalertHandlerEnabledAlertTypeProp) Set(val *EnumalertHandlerEnabledAlertTypeProp) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumalertHandlerEnabledAlertTypeProp) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumalertHandlerEnabledAlertTypeProp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumalertHandlerEnabledAlertTypeProp(val *EnumalertHandlerEnabledAlertTypeProp) *NullableEnumalertHandlerEnabledAlertTypeProp {
	return &NullableEnumalertHandlerEnabledAlertTypeProp{value: val, isSet: true}
}

func (v NullableEnumalertHandlerEnabledAlertTypeProp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumalertHandlerEnabledAlertTypeProp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

