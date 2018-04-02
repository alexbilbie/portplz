package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/alexbilbie/ecs-task-ports"
)

var badArgErr = "Expected first argument to be a port number"

func main() {
	if len(os.Args) != 2 {
		fmt.Println(badArgErr)
		os.Exit(1)
	}

	portArg := os.Args[1]

	portArgInt, err := strconv.Atoi(portArg)
	if err != nil {
		fmt.Println(badArgErr)
		os.Exit(1)
	}

	addresses, err := ecstaskports.Discover()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, addr := range addresses {
		if addr.ContainerPort == portArgInt {
			fmt.Printf("%d", addr.HostPort)
		}
	}

	os.Exit(2)
}
