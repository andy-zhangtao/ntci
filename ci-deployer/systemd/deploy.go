package main

import (
	"fmt"
)

func restartUnit(name string) (err error) {

	result := make(chan string)
	_, err = conn.TryRestartUnit(name, "replace", result)
	if err != nil {
		return err
	}

	fmt.Println(<-result)
	return nil
}
