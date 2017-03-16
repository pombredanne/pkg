// Copyright 2017 Ruda Moura. All rights reserved.

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

const version = "2017.3"

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print pkg version",
	Long:      `Version prints the pkg version.`,
}

type dict struct {
	Keys   []string `xml:"dict>key"`
	Values []string `xml:"dict>string"`
}

func macOSVersion() string {
	data, err := ioutil.ReadFile("/System/Library/CoreServices/SystemVersion.plist")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return "?"
	}
	var dict dict
	err = xml.Unmarshal(data, &dict)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		return "?"
	}
	if len(dict.Keys) == len(dict.Values) {
		for pos, key := range dict.Keys {
			if key == "ProductVersion" {
				return dict.Values[pos]
			}
		}
	}
	return "?"
}

func runVersion(cmd *Command, args []string) {
	fmt.Printf("pkg version %s on macOS %s (%s/%s)\n",
		version,
		macOSVersion(),
		runtime.GOOS,
		runtime.GOARCH)
}
