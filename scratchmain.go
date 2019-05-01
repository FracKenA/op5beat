package main

import (
	"fmt"
	"github.com/vbatoufflet/go-livestatus"
	"time"

	"os"
)

func main() {
	start := time.Now()

	count := 300
	then := start.Add(time.Duration(-count) * time.Second).Unix()
    timeFilter := fmt.Sprintf("last_check > %d AND ", then)

	// c := livestatus.NewClient("tcp", "localhost:6557")
	c := livestatus.NewClient("unix", "/opt/monitor/var/rw/live")
	defer c.Close()

	q := livestatus.NewQuery("services")
	q.Columns("description", "host_name", "host_is_flapping", "last_check", "host_groups", "state", "host_num_services")
	q.Filter(timeFilter)
	q.Filter("state = 0")

	resp, err := c.Exec(q)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}

	for _, r := range resp.Records {
		host_name_plain, err := r.Get("host_name")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostNumServices_plain, err := r.Get("host_num_services")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		description_plain, err := r.Get("description")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		state_plain, err := r.Get("state")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostIsFlapping_plain, err := r.Get("host_is_flapping")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		lastCheck_plain, err := r.Get("last_check")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostGroups_plain, err := r.Get("host_groups")
		if err != nil{
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}
		host_name_string, err := r.GetString("host_name")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostNumServices_int, err := r.GetInt("host_num_services")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		description_string, err := r.GetString("description")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		state_int, err := r.GetInt("state")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostIsFlapping_bool, err := r.GetBool("host_is_flapping")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		lastCheck_time, err := r.GetTime("last_check")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}

		hostGroups_slice, err := r.GetSlice("host_groups")
		if err != nil{
			fmt.Fprintf(os.Stderr, "Warning: %s", err)
		}


		fmt.Printf("\n\nUsing Plain Get Method\nHostname: %s\n\tHostNumServices: %f\n\tDescription: %s\n\tState: %f\n\tFlapping: %f\n\tLastCheck: %f\n\tHostGroups: %v\n", host_name_plain, hostNumServices_plain, description_plain, state_plain, hostIsFlapping_plain, lastCheck_plain, hostGroups_plain)
		fmt.Printf("\n\nUsing Specific Get Method\nHostname: %s\n\tHostNumServices: %d\n\tDescription: %s\n\tState: %d\n\tFlapping: %t\n\tLastCheck: %s\n\tHostGroups: %v\n", host_name_string, hostNumServices_int, description_string, state_int, hostIsFlapping_bool, lastCheck_time, hostGroups_slice)

	}
}
