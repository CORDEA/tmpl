package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"path"
	"tmpl/config"
)

const (
	baseDir  = "templates"
	fileName = "template.yaml"
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
	set := f.Arg(0)
	conf := config.Parse(path.Join(baseDir, set, fileName))
	args := conf.MapArgs(n.name, n.args)
	generator := generator{
		templatePath: conf.TemplatePath,
		args:         args,
	}
	generator.generate()
	return subcommands.ExitSuccess
}
