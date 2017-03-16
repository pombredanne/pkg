// Copyright 2017 Ruda Moura. All rights reserved.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdInfo = &Command{
	Run:       runInfo,
	UsageLine: "info [-volume path] [packages]",
	Short:     "get information from packages",
	Long: `
Get information from installed packages.

The -volume flag instructs to perform operation on the specified volume.  By
default, it uses the root volume "/".

See also: pkg files.
	`,
}

var infoVolume = "/"

func init() {
	cmdInfo.Flag.StringVar(&infoVolume, "volume", "/", "")
}

func infoPackage(pkg string, volume string) (output string, err error) {
	cmd := exec.Command("pkgutil", "--verbose", "--volume", volume, "--pkg-info", pkg)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func runInfo(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.Usage()
	}
	for _, pkg := range args {
		out, err := infoPackage(pkg, infoVolume)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", out)
			continue
		}
		fmt.Printf("%s", out)
	}
}
