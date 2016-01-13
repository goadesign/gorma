package gorma

import (
	"github.com/raphael/goa/goagen/codegen"
	"github.com/raphael/goa/goagen/meta"
)

var (
	// TargetPackage is the name of the generated Go package.
	TargetPackage string
)

// Command is the gorma model code generator command line data structure.
type Command struct {
	*codegen.BaseCommand
}

// NewCommand instantiates a new command.
func NewCommand() *Command {
	base := codegen.NewBaseCommand("gorma", "Generate model code")
	return &Command{BaseCommand: base}
}

// RegisterFlags registers the command line flags with the given registry.
func (c *Command) RegisterFlags(r codegen.FlagRegistry) {
	r.Flag("pkg", "Name of generated Go package containing models").
		Default("models").StringVar(&TargetPackage)
}

// Run simply calls the meta generator.
func (c *Command) Run() ([]string, error) {
	flags := map[string]string{"pkg": TargetPackage}
	gen := meta.NewGenerator(
		"gorma.Generate",
		[]*codegen.ImportSpec{codegen.SimpleImport("github.com/bketelsen/gorma")},
		flags,
	)
	return gen.Generate()
}
