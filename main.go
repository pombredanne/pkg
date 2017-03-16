// Copyright 2017 Ruda Moura. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE.txt file.
//
// This source contains code from the Go language that parses commands.
// Copyright 2011 The Go Authors. All rights reserved.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type Command struct {
	Run         func(cmd *Command, args []string)
	UsageLine   string
	Short       string
	Long        string
	Flag        flag.FlagSet
	CustomFlags bool
}

func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
	os.Exit(2)
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

var commands = []*Command{
	cmdVersion,
	cmdInstall,
	cmdList,
	cmdInfo,
	cmdFiles,
	cmdWhich,
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	if args[0] == "help" {
		help(args[1:])
		return
	}
	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "pkg: unknown subcommand %q\nRun 'pkg help' for usage.\n", args[0])
	os.Exit(2)
}

var usageTemplate = `pkg is a tool for managing macOS packages.

Usage:

	pkg command [arguments]

The commands are:
{{range .}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "pkg help [command]" for more information about a command.
`

var helpTemplate = `{{if .Runnable}}usage: pkg {{.UsageLine}}

{{end}}{{.Long | trim}}
`

type errWriter struct {
	w   io.Writer
	err error
}

func (w *errWriter) Write(b []byte) (int, error) {
	n, err := w.w.Write(b)
	if err != nil {
		w.err = err
	}
	return n, err
}

func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace})
	template.Must(t.Parse(text))
	ew := &errWriter{w: w}
	err := t.Execute(ew, data)
	if ew.err != nil {
		if strings.Contains(ew.err.Error(), "pipe") {
			os.Exit(1)
		}
	}
	if err != nil {
		panic(err)
	}
}

func printUsage(w io.Writer) {
	bw := bufio.NewWriter(w)
	tmpl(bw, usageTemplate, commands)
	bw.Flush()
}

func usage() {
	printUsage(os.Stderr)
	os.Exit(2)
}

func help(args []string) {
	if len(args) == 0 {
		printUsage(os.Stdout)
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: pkg help command\n\nToo many arguments given.\n")
		os.Exit(2)
	}
	arg := args[0]
	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmpl(os.Stdout, helpTemplate, cmd)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "Unknown help topic %#q.  Run 'pkg help'.\n", arg)
	os.Exit(2)
}
