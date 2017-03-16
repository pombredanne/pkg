// Copyright 2017 Ruda Moura. All rights reserved.

package main

import (
	"fmt"
	"os"
	"os/exec"
)

var cmdWhich = &Command{
	Run:       runWhich,
	UsageLine: "which [files]",
	Short:     "display which package installed a specific file",
	Long: `
Display which package installed a specific file.

See also: pkg list, pkg files.
	`,
}

func fileInfo(file string) (output string, err error) {
	cmd := exec.Command("pkgutil", "--verbose", "--file-info", file)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func runWhich(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.Usage()
	}
	for _, file := range args {
		out, err := fileInfo(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", out)
			continue
		}
		fmt.Printf("%s", out)
	}
}
