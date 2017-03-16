// Copyright 2017 Ruda Moura. All rights reserved.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdInstall = &Command{
	Run:       runInstall,
	UsageLine: "install [-volume path] [packages]",
	Short:     "install packages",
	Long: `
Install macOS packages to a volume.

The -volume flag instructs to perform operation on the specified volume.  By
default, it uses the root volume "/".

See also: pkg list.
	`,
}

var installVolume = "/"

func init() {
	cmdInstall.Flag.StringVar(&installVolume, "volume", "/", "")
}

func installPackage(pkg string, volume string) (output string, err error) {
	cmd := exec.Command("installer", "-pkg", pkg, "-target", volume)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func runInstall(cmd *Command, args []string) {
	verbose := true
	if len(args) == 0 {
		cmd.Usage()
	}
	for _, pkg := range args {
		fmt.Printf("Installing %s...\n", pkg)
		out, err := installPackage(pkg, installVolume)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s", out)
			continue
		}
		if verbose {
			fmt.Printf("%s", out)
		}
	}
}
