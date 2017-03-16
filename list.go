// Copyright 2017 Ruda Moura. All rights reserved.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdList = &Command{
	Run:       runList,
	UsageLine: "list [-volume path]",
	Short:     "list packages",
	Long: `
List all installed package.

The -volume flag instructs to perform operation on the specified volume.  By
default, it uses the root volume "/".

See also: pkg info, pkg files, pkg which.
	`,
}

var listVolume = "/"

func init() {
	cmdList.Flag.StringVar(&listVolume, "volume", "/", "")
}

func listPackages(volume string) (output string, err error) {
	cmd := exec.Command("pkgutil", "--volume", volume, "--pkgs")
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func runList(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	out, err := listPackages(listVolume)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", out)
		os.Exit(2)
	}
	fmt.Printf("%s", out)
}
