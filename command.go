package gorma

import (
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/goadesign/goa/goagen/meta"
)

var (
	// TargetPackage is the name of the generated Go package.
	TargetPackage string
	AppPackage    string
)

// Command is the goa application code generator command line data structure.
type Command struct {
	*codegen.BaseCommand
}

// NewCommand instantiates a new command.
func NewCommand() *Command {
	base := codegen.NewBaseCommand("gorma", "Generate Models")
	return &Command{BaseCommand: base}
}

// RegisterFlags registers the command line flags with the given registry.
func (c *Command) RegisterFlags(r codegen.FlagRegistry) {
	r.Flag("pkg", "Name of generated Go package containing models").
		Default("models").StringVar(&TargetPackage)
	r.Flag("app", "Name of goa generated Go package containing controllers, etc.").
		Default("app").StringVar(&AppPackage)
}

// Run simply calls the meta generator.
func (c *Command) Run() ([]string, error) {
	flags := map[string]string{"pkg": TargetPackage}
	gen := meta.NewGenerator(
		"modelgen.Generate",
		[]*codegen.ImportSpec{codegen.SimpleImport("github.com/bketelsen/gorma/modelgen")},
		flags,
	)
	return gen.Generate()
}
