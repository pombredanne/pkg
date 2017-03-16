// Copyright 2017 Ruda Moura. All rights reserved.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdFiles = &Command{
	Run:       runFiles,
	UsageLine: "files [-volume path] [packages]",
	Short:     "list files from packages",
	Long: `
List files from installed packages.

The -volume flag instructs to perform operation on the specified volume.  By
default, it uses the root volume "/".

See also: pkg info, pkg which.
	`,
}

var filesVolume = "/"

func init() {
	cmdFiles.Flag.StringVar(&filesVolume, "volume", "/", "")
}

func filesPackage(pkg string, volume string) (output string, err error) {
	cmd := exec.Command("pkgutil", "--volume", volume, "--files", pkg)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func runFiles(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.Usage()
	}
	for _, pkg := range args {
		out, err := filesPackage(pkg, filesVolume)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", out)
			continue
		}
		fmt.Printf("%s", out)
	}
}
