package gorma

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
	"github.com/raphael/goa/goagen/utils"

	"gopkg.in/alecthomas/kingpin.v2"
)

// Generator is the application code generator.
type Generator struct {
	genfiles []string
}

// Generate is the generator entry point called by the meta generator.
func Generate(api *design.APIDefinition) ([]string, error) {
	g, err := NewGenerator()
	if err != nil {
		return nil, err
	}
	return g.Generate(api)
}

// NewGenerator returns the application code generator.
func NewGenerator() (*Generator, error) {
	app := kingpin.New("Code generator", "application code generator")
	codegen.RegisterFlags(app)
	NewCommand().RegisterFlags(app)
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		return nil, fmt.Errorf(`invalid command line: %s. Command line was "%s"`,
			err, strings.Join(os.Args, " "))
	}
	outdir := ModelOutputDir()
	//os.RemoveAll(outdir)
	if err = os.MkdirAll(outdir, 0777); err != nil {
		return nil, err
	}
	return &Generator{
		genfiles: []string{outdir},
	}, nil
}

// ModelOutputDir returns the directory containing the generated files.
func ModelOutputDir() string {
	return filepath.Join(codegen.OutputDir, TargetPackage)
}

// ModelPackagePath returns the Go package path to the generated package.
func ModelPackagePath() (string, error) {
	outputDir := ModelOutputDir()
	gopaths := filepath.SplitList(os.Getenv("GOPATH"))
	for _, gopath := range gopaths {
		if strings.HasPrefix(outputDir, gopath) {
			path, err := filepath.Rel(filepath.Join(gopath, "src"), outputDir)
			if err != nil {
				return "", err
			}
			return filepath.ToSlash(path), nil
		}
	}
	return "", fmt.Errorf("output directory outside of Go workspace, make sure to define GOPATH correctly or change output directory")
}

// Generate the application code, implement codegen.Generator.
func (g *Generator) Generate(api *design.APIDefinition) (_ []string, err error) {
	if api == nil {
		return nil, fmt.Errorf("missing API definition, make sure design.Design is properly initialized")
	}
	go utils.Catch(nil, func() { g.Cleanup() })
	defer func() {
		if err != nil {
			g.Cleanup()
		}
	}()
	outdir := ModelOutputDir()
	if err := os.MkdirAll(outdir, 0755); err != nil {
		return nil, err
	}

	err = api.IterateVersions(func(v *design.APIVersionDefinition) error {

		// only generate for the base unnamed version
		// because user types are unversioned
		if v.Version != "" {
			return nil
		}
		if err := g.generateUserTypes(outdir, v); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return g.genfiles, nil
}

// Cleanup removes the entire "app" directory if it was created by this generator.
func (g *Generator) Cleanup() {
	if len(g.genfiles) == 0 {
		return
	}
	os.RemoveAll(ModelOutputDir())
	g.genfiles = nil
}

// MergeResponses merge the response maps overriding the first argument map entries with the
// second argument map entries in case of collision.
func MergeResponses(l, r map[string]*design.ResponseDefinition) map[string]*design.ResponseDefinition {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	for n, r := range r {
		l[n] = r
	}
	return l
}

// Generated package name for resources supporting the given version.
func packageName(version *design.APIVersionDefinition) (pack string) {
	pack = TargetPackage
	if version.Version != "" {
		pack = codegen.Goify(codegen.VersionPackage(version.Version), false)
	}
	return
}

// generateUserTypes iterates through the user types and generates the data structures and
// marshaling code.
func (g *Generator) generateUserTypes(outdir string, version *design.APIVersionDefinition) error {

	var modelname, filename string
	err := GormaDesign.IterateStores(func(store *RelationalStoreDefinition) error {
		store.IterateModels(func(model *RelationalModelDefinition) error {
			modelname = strings.ToLower(codegen.Goify(model.Name, false))
			modeldir := filepath.Join(outdir, "models", model.Name)
			filename = fmt.Sprintf("%s_gen.go", modelname)
			utFile := filepath.Join(modeldir, filename)
			utWr, err := NewUserTypesWriter(utFile)
			if err != nil {
				panic(err) // bug
			}
			title := fmt.Sprintf("%s: Models", version.Context())
			imports := []*codegen.ImportSpec{
				codegen.SimpleImport("fmt"),
			}

			utWr.WriteHeader(title, packageName(version), imports)
			data := &UserTypeTemplateData{
				UserType:   nil,
				Versioned:  version.Version != "",
				DefaultPkg: TargetPackage,
			}
			err = utWr.Execute(data)
			g.genfiles = append(g.genfiles, utFile)
			if err != nil {
				return err
			}
			return utWr.FormatCode()

		})
		return nil
	})
	return err
}

/*
	err = version.IterateUserTypes(func(t *design.UserTypeDefinition) error {

	})
*/
