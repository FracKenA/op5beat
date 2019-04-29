package main

import (
	"fmt"
	"time"
)

func main() {
//	count := -30
//	currentTime := time.Now().Unix()
//	then := currentTime.
//	fmt.Println("Current Unix Time:", currentTime)
//	fmt.Println("30 seconds ago in Unix Time:", currentTime)

	// now := time.Now()
	currentTime := time.Now().Unix()
	fmt.Println("now:", currentTime)

	count := 30
	then := time.Now().Add(time.Duration(-count) * time.Second).Unix()
		// if we had fix number of units to subtract, we can use following line instead fo above 2 lines. It does type convertion automatically.
		// then := now.Add(-10 * time.Minute)


	fmt.Println("30 seconds ago:", then)
	}
