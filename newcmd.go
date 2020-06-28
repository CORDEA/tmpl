package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

type newCmd struct {
	name string
	args string
}

func (*newCmd) Name() string {
	return "new"
}

func (*newCmd) Synopsis() string {
	return ""
}

func (*newCmd) Usage() string {
	return ""
}

func (n *newCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&n.name, "name", "", "name")
	f.StringVar(&n.args, "args", "", "args")
}

func (n *newCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	return subcommands.ExitSuccess
}