// Copyright 2016 Canonical Ltd.  All rights reserved.

package main

import (
	"fmt"
	"os"

	"github.com/juju/cmd"

	tcmd "github.com/juju/terms-client/v2/cmd"
)

func main() {
	ctx, err := cmd.DefaultContext()
	if err != nil {
		fmt.Printf("failed to get command context: %v\n", err)
		os.Exit(2)
	}
	c := tcmd.NewReleaseTermCommand()
	args := os.Args
	os.Exit(cmd.Main(c, ctx, args[1:]))
}
