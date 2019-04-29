package main

import (
	"fmt"
	"github.com/vbatoufflet/go-livestatus"
	"time"

	"os"
)

func main() {
	start := time.Now()

	count := 30
	then := start.Add(time.Duration(-count) * time.Second).Unix()
    timeFilter := fmt.Sprintf("last_check > %d AND ", then)

	// c := livestatus.NewClient("tcp", "localhost:6557")
	c := livestatus.NewClient("unix", "/opt/monitor/var/rw/live")
	defer c.Close()

	q := livestatus.NewQuery("services")
	q.Columns("acknowledged, acknowledgement_type, check_command, check_freshness, check_type, custom_variable_names, custom_variable_values, custom_variables, description, display_name, downtimes_with_info, execution_time, groups, host_name, host_acknowledged, host_acknowledgement_type, host_address, host_alias, host_check_freshness, host_check_type, host_childs, host_custom_variable_names, host_custom_variable_values, host_custom_variables, host_display_name, host_downtimes_with_info, host_groups, host_hard_state, host_is_flapping, host_last_check, host_last_hard_state_change, host_last_hard_state, host_last_state_change, host_last_state, host_last_time_down, host_last_time_unreachable, host_last_time_up, host_latency, host_long_plugin_output, host_num_services_crit, host_num_services_hard_crit, host_num_services_hard_ok, host_num_services_hard_unknown, host_num_services_hard_warn, host_num_services_ok, host_num_services_pending, host_num_services_unknown, host_num_services_warn, host_num_services, host_parents, host_pending_flex_downtime, host_percent_state_change, host_perf_data, host_plugin_output, host_scheduled_downtime_depth, host_state_type, host_state, host_total_services, host_worst_service_hard_state, host_worst_service_state, is_flapping, last_check, last_hard_state_change, last_hard_state, last_state_change, last_state, last_time_critical, last_time_ok, last_time_unknown, last_time_warning, latency, long_plugin_output, percent_state_change, perf_data, plugin_output, scheduled_downtime_depth, state_type, state")
	q.Filter(timeFilter)
	q.Filter("state = 0")

	resp, err := c.Exec(q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}

	for _, r := range resp.Records {
		host_name, err := r.GetString("host_name")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostNumServices, err := r.GetInt("host_num_services")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		description, err := r.GetString("description")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		state, err := r.GetInt("state")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostIsFlapping, err := r.GetBool("host_is_flapping")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		lastCheck, err := r.GetTime("last_check")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostGroups, err := r.GetSlice("host_groups")
		if err != nil{
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}


		fmt.Printf("Host: %s has %d services.\n    The Service %s is currently in state: %d,\n    Is the host flapping? %t\n    Last Service Check was performed at %s\n    The host is in the following hostgroups: %v\n", host_name, hostNumServices, description, state, hostIsFlapping, lastCheck, hostGroups)
	}
}
