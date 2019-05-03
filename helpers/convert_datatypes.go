package helpers

import (
	"fmt"
	"github.com/vbatoufflet/go-livestatus"
)

func GetWithCorrectDataType(r livestatus.Record, c string) (interface{}, error) {
	functionMatches := map[string]string{"accept_passive_checks": "GetBool", "accept_passive_host_checks": "GetBool", "accept_passive_service_checks": "GetBool", "acknowledged": "GetBool", "active_checks_enabled": "GetBool", "can_submit_commands": "GetBool", "check_external_commands": "GetBool", "check_flapping_recovery_notification": "GetBool", "check_freshness": "GetBool", "check_host_freshness": "GetBool", "check_options": "GetBool", "check_service_freshness": "GetBool", "checks_enabled": "GetBool", "current_contact_can_submit_commands": "GetBool", "current_contact_host_notifications_enabled": "GetBool", "current_contact_in_host_notification_period": "GetBool", "current_contact_in_service_notification_period": "GetBool", "current_contact_service_notifications_enabled": "GetBool", "current_host_accept_passive_checks": "GetBool", "current_host_acknowledged": "GetBool", "current_host_active_checks_enabled": "GetBool", "current_host_check_flapping_recovery_notification": "GetBool", "current_host_check_freshness": "GetBool", "current_host_checks_enabled": "GetBool", "current_host_event_handler_enabled": "GetBool", "current_host_flap_detection_enabled": "GetBool", "current_host_has_been_checked": "GetBool", "current_host_in_check_period": "GetBool", "current_host_in_notification_period": "GetBool", "current_host_is_executing": "GetBool", "current_host_is_flapping": "GetBool", "current_host_no_more_notifications": "GetBool", "current_host_notifications_enabled": "GetBool", "current_host_obsess": "GetBool", "current_host_obsess_over_host": "GetBool", "current_host_pending_flex_downtime": "GetBool", "current_host_pnpgraph_present": "GetBool", "current_host_process_performance_data": "GetBool", "current_host_should_be_scheduled": "GetBool", "current_service_accept_passive_checks": "GetBool", "current_service_acknowledged": "GetBool", "current_service_active_checks_enabled": "GetBool", "current_service_check_freshness": "GetBool", "current_service_check_options": "GetBool", "current_service_checks_enabled": "GetBool", "current_service_event_handler_enabled": "GetBool", "current_service_flap_detection_enabled": "GetBool", "current_service_has_been_checked": "GetBool", "current_service_in_check_period": "GetBool", "current_service_in_notification_period": "GetBool", "current_service_is_executing": "GetBool", "current_service_is_flapping": "GetBool", "current_service_no_more_notifications": "GetBool", "current_service_notifications_enabled": "GetBool", "current_service_obsess": "GetBool", "current_service_obsess_over_service": "GetBool", "current_service_pnpgraph_present": "GetBool", "current_service_process_performance_data": "GetBool", "current_service_should_be_scheduled": "GetBool", "enable_event_handlers": "GetBool", "enable_flap_detection": "GetBool", "enable_notifications": "GetBool", "event_handler_enabled": "GetBool", "execute_host_checks": "GetBool", "execute_service_checks": "GetBool", "flap_detection_enabled": "GetBool", "has_been_checked": "GetBool", "host_accept_passive_checks": "GetBool", "host_acknowledged": "GetBool", "host_active_checks_enabled": "GetBool", "host_check_flapping_recovery_notification": "GetBool", "host_check_freshness": "GetBool", "host_checks_enabled": "GetBool", "host_event_handler_enabled": "GetBool", "host_flap_detection_enabled": "GetBool", "host_has_been_checked": "GetBool", "host_in_check_period": "GetBool", "host_in_notification_period": "GetBool", "host_is_executing": "GetBool", "host_is_flapping": "GetBool", "host_no_more_notifications": "GetBool", "host_notifications_enabled": "GetBool", "host_obsess": "GetBool", "host_obsess_over_host": "GetBool", "host_pending_flex_downtime": "GetBool", "host_pnpgraph_present": "GetBool", "host_process_performance_data": "GetBool", "host_should_be_scheduled": "GetBool", "in": "GetBool", "in_check_period": "GetBool", "in_host_notification_period": "GetBool", "in_notification_period": "GetBool", "in_service_notification_period": "GetBool", "is_executing": "GetBool", "is_flapping": "GetBool", "no_more_notifications": "GetBool", "notifications_enabled": "GetBool", "obsess": "GetBool", "obsess_over_host": "GetBool", "obsess_over_hosts": "GetBool", "obsess_over_service": "GetBool", "obsess_over_services": "GetBool", "pending_flex_downtime": "GetBool", "persistent": "GetBool", "pnpgraph_present": "GetBool", "process_performance_data": "GetBool", "service_accept_passive_checks": "GetBool", "service_acknowledged": "GetBool", "service_active_checks_enabled": "GetBool", "service_check_freshness": "GetBool", "service_check_options": "GetBool", "service_checks_enabled": "GetBool", "service_event_handler_enabled": "GetBool", "service_flap_detection_enabled": "GetBool", "service_has_been_checked": "GetBool", "service_in_check_period": "GetBool", "service_in_notification_period": "GetBool", "service_is_executing": "GetBool", "service_is_flapping": "GetBool", "service_no_more_notifications": "GetBool", "service_notifications_enabled": "GetBool", "service_obsess": "GetBool", "service_obsess_over_service": "GetBool", "service_pnpgraph_present": "GetBool", "service_process_performance_data": "GetBool", "service_should_be_scheduled": "GetBool", "should_be_scheduled": "GetBool", "current_host_last_check": "GetTime", "current_host_last_hard_state_change": "GetTime", "current_host_last_notification": "GetTime", "current_host_last_state_change": "GetTime", "current_host_last_time_down": "GetTime", "current_host_last_time_unreachable": "GetTime", "current_host_last_time_up": "GetTime", "current_host_next_check": "GetTime", "current_host_next_notification": "GetTime", "current_service_last_check": "GetTime", "current_service_last_hard_state_change": "GetTime", "current_service_last_notification": "GetTime", "current_service_last_state_change": "GetTime", "current_service_last_time_critical": "GetTime", "current_service_last_time_ok": "GetTime", "current_service_last_time_unknown": "GetTime", "current_service_last_time_warning": "GetTime", "current_service_next_check": "GetTime", "current_service_next_notification": "GetTime", "end_time": "GetTime", "entry_time": "GetTime", "expire_time": "GetTime", "host_last_check": "GetTime", "host_last_hard_state_change": "GetTime", "host_last_notification": "GetTime", "host_last_state_change": "GetTime", "host_last_time_down": "GetTime", "host_last_time_unreachable": "GetTime", "host_last_time_up": "GetTime", "host_next_check": "GetTime", "host_next_notification": "GetTime", "last_check": "GetTime", "last_command_check": "GetTime", "last_hard_state_change": "GetTime", "last_log_rotation": "GetTime", "last_notification": "GetTime", "last_state_change": "GetTime", "last_time_critical": "GetTime", "last_time_down": "GetTime", "last_time_ok": "GetTime", "last_time_unknown": "GetTime", "last_time_unreachable": "GetTime", "last_time_up": "GetTime", "last_time_warning": "GetTime", "next_check": "GetTime", "next_notification": "GetTime", "program_start": "GetTime", "service_last_check": "GetTime", "service_last_hard_state_change": "GetTime", "service_last_notification": "GetTime", "service_last_state_change": "GetTime", "service_last_time_critical": "GetTime", "service_last_time_ok": "GetTime", "service_last_time_unknown": "GetTime", "service_last_time_warning": "GetTime", "service_next_check": "GetTime", "service_next_notification": "GetTime", "start_time": "GetTime", "time": "GetTime", "check_interval": "GetFloat", "connections_rate": "GetFloat", "current_host_check_interval": "GetFloat", "current_host_execution_time": "GetFloat", "current_host_first_notification_delay": "GetFloat", "current_host_high_flap_threshold": "GetFloat", "current_host_latency": "GetFloat", "current_host_low_flap_threshold": "GetFloat", "current_host_notification_interval": "GetFloat", "current_host_percent_state_change": "GetFloat", "current_host_retry_interval": "GetFloat", "current_host_staleness": "GetFloat", "current_host_x_3d": "GetFloat", "current_host_y_3d": "GetFloat", "current_host_z_3d": "GetFloat", "current_service_check_interval": "GetFloat", "current_service_execution_time": "GetFloat", "current_service_first_notification_delay": "GetFloat", "current_service_high_flap_threshold": "GetFloat", "current_service_latency": "GetFloat", "current_service_low_flap_threshold": "GetFloat", "current_service_notification_interval": "GetFloat", "current_service_percent_state_change": "GetFloat", "current_service_retry_interval": "GetFloat", "current_service_staleness": "GetFloat", "execution_time": "GetFloat", "first_notification_delay": "GetFloat", "forks_rate": "GetFloat", "high_flap_threshold": "GetFloat", "host_check_interval": "GetFloat", "host_checks_rate": "GetFloat", "host_execution_time": "GetFloat", "host_first_notification_delay": "GetFloat", "host_high_flap_threshold": "GetFloat", "host_latency": "GetFloat", "host_low_flap_threshold": "GetFloat", "host_notification_interval": "GetFloat", "host_percent_state_change": "GetFloat", "host_retry_interval": "GetFloat", "host_staleness": "GetFloat", "host_x_3d": "GetFloat", "host_y_3d": "GetFloat", "host_z_3d": "GetFloat", "latency": "GetFloat", "livecheck_overflows_rate": "GetFloat", "livechecks_rate": "GetFloat", "log_messages_rate": "GetFloat", "low_flap_threshold": "GetFloat", "neb_callbacks_rate": "GetFloat", "notification_interval": "GetFloat", "percent_state_change": "GetFloat", "requests_rate": "GetFloat", "retry_interval": "GetFloat", "service_check_interval": "GetFloat", "service_checks_rate": "GetFloat", "service_execution_time": "GetFloat", "service_first_notification_delay": "GetFloat", "service_high_flap_threshold": "GetFloat", "service_latency": "GetFloat", "service_low_flap_threshold": "GetFloat", "service_notification_interval": "GetFloat", "service_percent_state_change": "GetFloat", "service_retry_interval": "GetFloat", "service_staleness": "GetFloat", "staleness": "GetFloat", "x_3d": "GetFloat", "y_3d": "GetFloat", "z_3d": "GetFloat", "acknowledgement_type": "GetInt", "attempt": "GetInt", "cached_log_messages": "GetInt", "check_type": "GetInt", "class": "GetInt", "connections": "GetInt", "current_attempt": "GetInt", "current_command_id": "GetInt", "current_contact_id": "GetInt", "current_contact_modified_attributes": "GetInt", "current_host_acknowledgement_type": "GetInt", "current_host_check_options": "GetInt", "current_host_check_type": "GetInt", "current_host_current_attempt": "GetInt", "current_host_current_notification_number": "GetInt", "current_host_hard_state": "GetInt", "current_host_hourly_value": "GetInt", "current_host_id": "GetInt", "current_host_initial_state": "GetInt", "current_host_last_hard_state": "GetInt", "current_host_last_state": "GetInt", "current_host_max_check_attempts": "GetInt", "current_host_modified_attributes": "GetInt", "current_host_notified_on": "GetInt", "current_host_num_services": "GetInt", "current_host_num_services_crit": "GetInt", "current_host_num_services_hard_crit": "GetInt", "current_host_num_services_hard_ok": "GetInt", "current_host_num_services_hard_unknown": "GetInt", "current_host_num_services_hard_warn": "GetInt", "current_host_num_services_ok": "GetInt", "current_host_num_services_pending": "GetInt", "current_host_num_services_unknown": "GetInt", "current_host_num_services_warn": "GetInt", "current_host_scheduled_downtime_depth": "GetInt", "current_host_state": "GetInt", "current_host_state_type": "GetInt", "current_host_total_services": "GetInt", "current_host_worst_service_hard_state": "GetInt", "current_host_worst_service_state": "GetInt", "current_notification_number": "GetInt", "current_service_acknowledgement_type": "GetInt", "current_service_check_type": "GetInt", "current_service_current_attempt": "GetInt", "current_service_current_notification_number": "GetInt", "current_service_hourly_value": "GetInt", "current_service_id": "GetInt", "current_service_initial_state": "GetInt", "current_service_last_hard_state": "GetInt", "current_service_last_state": "GetInt", "current_service_max_check_attempts": "GetInt", "current_service_modified_attributes": "GetInt", "current_service_notified_on": "GetInt", "current_service_scheduled_downtime_depth": "GetInt", "current_service_state": "GetInt", "current_service_state_type": "GetInt", "duration": "GetInt", "entry_type": "GetInt", "expires": "GetInt", "fixed": "GetInt", "forks": "GetInt", "hard_state": "GetInt", "host_acknowledgement_type": "GetInt", "host_check_options": "GetInt", "host_check_type": "GetInt", "host_checks": "GetInt", "host_current_attempt": "GetInt", "host_current_notification_number": "GetInt", "host_hard_state": "GetInt", "host_hourly_value": "GetInt", "host_id": "GetInt", "host_initial_state": "GetInt", "host_last_hard_state": "GetInt", "host_last_state": "GetInt", "host_max_check_attempts": "GetInt", "host_modified_attributes": "GetInt", "host_notified_on": "GetInt", "host_num_services": "GetInt", "host_num_services_crit": "GetInt", "host_num_services_hard_crit": "GetInt", "host_num_services_hard_ok": "GetInt", "host_num_services_hard_unknown": "GetInt", "host_num_services_hard_warn": "GetInt", "host_num_services_ok": "GetInt", "host_num_services_pending": "GetInt", "host_num_services_unknown": "GetInt", "host_num_services_warn": "GetInt", "host_scheduled_downtime_depth": "GetInt", "host_state": "GetInt", "host_state_type": "GetInt", "host_total_services": "GetInt", "host_worst_service_hard_state": "GetInt", "host_worst_service_state": "GetInt", "hostgroup_id": "GetInt", "hostgroup_num_hosts": "GetInt", "hostgroup_num_hosts_down": "GetInt", "hostgroup_num_hosts_pending": "GetInt", "hostgroup_num_hosts_unreach": "GetInt", "hostgroup_num_hosts_up": "GetInt", "hostgroup_num_services": "GetInt", "hostgroup_num_services_crit": "GetInt", "hostgroup_num_services_hard_crit": "GetInt", "hostgroup_num_services_hard_ok": "GetInt", "hostgroup_num_services_hard_unknown": "GetInt", "hostgroup_num_services_hard_warn": "GetInt", "hostgroup_num_services_ok": "GetInt", "hostgroup_num_services_pending": "GetInt", "hostgroup_num_services_unknown": "GetInt", "hostgroup_num_services_warn": "GetInt", "hostgroup_worst_host_state": "GetInt", "hostgroup_worst_service_hard_state": "GetInt", "hostgroup_worst_service_state": "GetInt", "hourly_value": "GetInt", "id": "GetInt", "initial_state": "GetInt", "interval_length": "GetInt", "is_service": "GetInt", "last_hard_state": "GetInt", "last_state": "GetInt", "lineno": "GetInt", "livecheck_overflows": "GetInt", "livechecks": "GetInt", "log_messages": "GetInt", "max_check_attempts": "GetInt", "modified_attributes": "GetInt", "nagios_pid": "GetInt", "neb_callbacks": "GetInt", "notified_on": "GetInt", "num_hosts": "GetInt", "num_hosts_down": "GetInt", "num_hosts_pending": "GetInt", "num_hosts_unreach": "GetInt", "num_hosts_up": "GetInt", "num_services": "GetInt", "num_services_crit": "GetInt", "num_services_hard_crit": "GetInt", "num_services_hard_ok": "GetInt", "num_services_hard_unknown": "GetInt", "num_services_hard_warn": "GetInt", "num_services_ok": "GetInt", "num_services_pending": "GetInt", "num_services_unknown": "GetInt", "num_services_warn": "GetInt", "requests": "GetInt", "scheduled_downtime_depth": "GetInt", "service_acknowledgement_type": "GetInt", "service_check_type": "GetInt", "service_checks": "GetInt", "service_current_attempt": "GetInt", "service_current_notification_number": "GetInt", "service_hourly_value": "GetInt", "service_id": "GetInt", "service_initial_state": "GetInt", "service_last_hard_state": "GetInt", "service_last_state": "GetInt", "service_max_check_attempts": "GetInt", "service_modified_attributes": "GetInt", "service_notified_on": "GetInt", "service_scheduled_downtime_depth": "GetInt", "service_state": "GetInt", "service_state_type": "GetInt", "servicegroup_id": "GetInt", "servicegroup_num_services": "GetInt", "servicegroup_num_services_crit": "GetInt", "servicegroup_num_services_hard_crit": "GetInt", "servicegroup_num_services_hard_ok": "GetInt", "servicegroup_num_services_hard_unknown": "GetInt", "servicegroup_num_services_hard_warn": "GetInt", "servicegroup_num_services_ok": "GetInt", "servicegroup_num_services_pending": "GetInt", "servicegroup_num_services_unknown": "GetInt", "servicegroup_num_services_warn": "GetInt", "servicegroup_worst_service_state": "GetInt", "source": "GetInt", "state": "GetInt", "state_type": "GetInt", "total_services": "GetInt", "triggered_by": "GetInt", "type": "GetInt", "worst_host_state": "GetInt", "worst_service_hard_state": "GetInt", "worst_service_state": "GetInt", "address": "GetString", "alias": "GetString", "check_command": "GetString", "check_source": "GetString", "command_name": "GetString", "contact_name": "GetString", "current_command_name": "GetString", "current_contact_alias": "GetString", "current_contact_name": "GetString", "current_host_address": "GetString", "current_host_alias": "GetString", "current_host_check_command": "GetString", "current_host_check_source": "GetString", "current_host_display_name": "GetString", "current_host_long_plugin_output": "GetString", "current_host_name": "GetString", "current_host_plugin_output": "GetString", "current_service_check_command": "GetString", "current_service_description": "GetString", "current_service_display_name": "GetString", "current_service_long_plugin_output": "GetString", "current_service_notification_period": "GetString", "current_service_plugin_output": "GetString", "description": "GetString", "display_name": "GetString", "event_handler": "GetString", "host_alias": "GetString", "host_check_command": "GetString", "host_display_name": "GetString", "host_long_plugin_output": "GetString", "host_name": "GetString", "host_plugin_output": "GetString", "hostgroup_alias": "GetString", "hostgroup_name": "GetString", "long_plugin_output": "GetString", "name": "GetString", "plugin_output": "GetString", "service_check_command": "GetString", "service_description": "GetString", "service_display_name": "GetString", "service_long_plugin_output": "GetString", "service_plugin_output": "GetString", "servicegroup_alias": "GetString", "servicegroup_name": "GetString", "childs": "GetSlice", "comments": "GetSlice", "comments_with_info": "GetSlice", "contact_groups": "GetSlice", "contacts": "GetSlice", "current_contact_custom_variable_names": "GetSlice", "current_contact_custom_variable_values": "GetSlice", "current_contact_custom_variables": "GetSlice", "current_contact_modified_attributes_list": "GetSlice", "current_host_childs": "GetSlice", "current_host_comments": "GetSlice", "current_host_comments_with_info": "GetSlice", "current_host_contact_groups": "GetSlice", "current_host_contacts": "GetSlice", "current_host_custom_variable_names": "GetSlice", "current_host_custom_variable_values": "GetSlice", "current_host_custom_variables": "GetSlice", "current_host_downtimes": "GetSlice", "current_host_downtimes_with_info": "GetSlice", "current_host_groups": "GetSlice", "current_host_modified_attributes_list": "GetSlice", "current_host_parents": "GetSlice", "current_host_services": "GetSlice", "current_host_services_with_info": "GetSlice", "current_host_services_with_state": "GetSlice", "current_service_comments": "GetSlice", "current_service_comments_with_info": "GetSlice", "current_service_contact_groups": "GetSlice", "current_service_contacts": "GetSlice", "current_service_custom_variable_names": "GetSlice", "current_service_custom_variable_values": "GetSlice", "current_service_custom_variables": "GetSlice", "current_service_downtimes": "GetSlice", "current_service_downtimes_with_info": "GetSlice", "current_service_groups": "GetSlice", "current_service_modified_attributes_list": "GetSlice", "custom_variable_names": "GetSlice", "custom_variables": "GetSlice", "days": "GetSlice", "downtimes": "GetSlice", "downtimes_with_info": "GetSlice", "exceptions_calendar_dates": "GetSlice", "exceptions_month_date": "GetSlice", "exceptions_month_day": "GetSlice", "exceptions_month_week_day": "GetSlice", "exceptions_week_day": "GetSlice", "exclusions": "GetSlice", "groups": "GetSlice", "host_childs": "GetSlice", "host_comments": "GetSlice", "host_comments_with_info": "GetSlice", "host_contact_groups": "GetSlice", "host_contacts": "GetSlice", "host_custom_variable_names": "GetSlice", "host_custom_variable_values": "GetSlice", "host_custom_variables": "GetSlice", "host_downtimes": "GetSlice", "host_downtimes_with_info": "GetSlice", "host_groups": "GetSlice", "host_modified_attributes_list": "GetSlice", "host_parents": "GetSlice", "host_services": "GetSlice", "host_services_with_info": "GetSlice", "host_services_with_state": "GetSlice", "hostgroup_members": "GetSlice", "hostgroup_members_with_state": "GetSlice", "members": "GetSlice", "members_with_state": "GetSlice", "modified_attributes_list": "GetSlice", "parents": "GetSlice", "service_comments": "GetSlice", "service_comments_with_info": "GetSlice", "service_contact_groups": "GetSlice", "service_contacts": "GetSlice", "service_custom_variable_names": "GetSlice", "service_custom_variable_values": "GetSlice", "service_custom_variables": "GetSlice", "service_downtimes": "GetSlice", "service_downtimes_with_info": "GetSlice", "service_groups": "GetSlice", "service_modified_attributes_list": "GetSlice", "servicegroup_members": "GetSlice", "servicegroup_members_with_state": "GetSlice", "services": "GetSlice", "services_with_info": "GetSlice", "services_with_state": "GetSlice", "action_url": "GetString", "action_url_expanded": "GetString", "address1": "GetString", "address2": "GetString", "address3": "GetString", "address4": "GetString", "address5": "GetString", "address6": "GetString", "author": "GetString", "check_period": "GetString", "comment": "GetString", "current_command_line": "GetString", "current_contact_address1": "GetString", "current_contact_address2": "GetString", "current_contact_address3": "GetString", "current_contact_address4": "GetString", "current_contact_address5": "GetString", "current_contact_address6": "GetString", "current_contact_email": "GetString", "current_contact_host_notification_period": "GetString", "current_contact_pager": "GetString", "current_contact_service_notification_period": "GetString", "current_host_action_url": "GetString", "current_host_action_url_expanded": "GetString", "current_host_check_period": "GetString", "current_host_event_handler": "GetString", "current_host_filename": "GetString", "current_host_icon_image": "GetString", "current_host_icon_image_alt": "GetString", "current_host_icon_image_expanded": "GetString", "current_host_notes": "GetString", "current_host_notes_expanded": "GetString", "current_host_notes_url": "GetString", "current_host_notes_url_expanded": "GetString", "current_host_notification_period": "GetString", "current_host_perf_data": "GetString", "current_host_statusmap_image": "GetString", "current_service_action_url": "GetString", "current_service_action_url_expanded": "GetString", "current_service_check_period": "GetString", "current_service_check_source": "GetString", "current_service_event_handler": "GetString", "current_service_icon_image": "GetString", "current_service_icon_image_alt": "GetString", "current_service_icon_image_expanded": "GetString", "current_service_notes": "GetString", "current_service_notes_expanded": "GetString", "current_service_notes_url": "GetString", "current_service_notes_url_expanded": "GetString", "current_service_perf_data": "GetString", "email": "GetString", "filename": "GetString", "host_action_url": "GetString", "host_action_url_expanded": "GetString", "host_address": "GetString", "host_check_period": "GetString", "host_check_source": "GetString", "host_event_handler": "GetString", "host_filename": "GetString", "host_icon_image": "GetString", "host_icon_image_alt": "GetString", "host_icon_image_expanded": "GetString", "host_notes": "GetString", "host_notes_expanded": "GetString", "host_notes_url": "GetString", "host_notes_url_expanded": "GetString", "host_notification_period": "GetString", "host_perf_data": "GetString", "host_statusmap_image": "GetString", "hostgroup_action_url": "GetString", "hostgroup_notes": "GetString", "hostgroup_notes_url": "GetString", "icon_image": "GetString", "icon_image_alt": "GetString", "icon_image_expanded": "GetString", "line": "GetString", "livestatus_version": "GetString", "message": "GetString", "notes": "GetString", "notes_expanded": "GetString", "notes_url": "GetString", "notes_url_expanded": "GetString", "notification_period": "GetString", "options": "GetString", "pager": "GetString", "perf_data": "GetString", "program_version": "GetString", "service_action_url": "GetString", "service_action_url_expanded": "GetString", "service_check_period": "GetString", "service_check_source": "GetString", "service_event_handler": "GetString", "service_icon_image": "GetString", "service_icon_image_alt": "GetString", "service_icon_image_expanded": "GetString", "service_notes": "GetString", "service_notes_expanded": "GetString", "service_notes_url": "GetString", "service_notes_url_expanded": "GetString", "service_notification_period": "GetString", "service_perf_data": "GetString", "servicegroup_action_url": "GetString", "servicegroup_notes": "GetString", "servicegroup_notes_url": "GetString", "statusmap_image": "GetString", "table": "GetString"}

	if r == nil || c == "" {
		return "", fmt.Errorf("empty parameter")
	}
	switch functionMatches[c] {
	case "GetString":
		return r.GetString(c)
	case "GetBool":
		return r.GetBool(c)
	case "GetTime":
		return r.GetTime(c)
	case "GetFloat":
		return r.GetFloat(c)
	case "GetInt":
		return r.GetInt(c)
	case "GetSlice":
		return r.GetSlice(c)
	default:
		return r.Get(c)
	}
	return r, fmt.Errorf("empty parameter")
}
