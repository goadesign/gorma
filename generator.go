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

// AppOutputDir returns the directory containing the generated files.
func AppOutputDir() string {
	return filepath.Join(codegen.OutputDir, AppPackage)
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

// AppPackagePath returns the Go package path to the generated package.
func AppPackagePath() (string, error) {
	outputDir := AppOutputDir()
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

	if err := g.generateUserTypes(outdir, api); err != nil {
		return g.genfiles, err
	}

	return g.genfiles, nil
}

// Cleanup removes the entire "app" directory if it was created by this generator.
func (g *Generator) Cleanup() {
	if len(g.genfiles) == 0 {
		return
	}
	//os.RemoveAll(ModelOutputDir())
	g.genfiles = nil
}

// Generated package name for resources supporting the given version.
func packageName(version *design.APIVersionDefinition) (pack string) {
	pack = AppPackage
	if version.Version != "" {
		pack = codegen.Goify(codegen.VersionPackage(version.Version), false)
	}
	return
}

// generateUserTypes iterates through the user types and generates the data structures and
// marshaling code.
func (g *Generator) generateUserTypes(outdir string, api *design.APIDefinition) error {
	fmt.Println("Generating user types")
	err := api.IterateVersions(func(version *design.APIVersionDefinition) error {
		if version.Version != "" {
			fmt.Println("Skipping version")
			return nil
		}
		fmt.Println("Not skipping verson")
		var modelname, filename string
		err := GormaDesign.IterateStores(func(store *RelationalStoreDefinition) error {
			err := store.IterateModels(func(model *RelationalModelDefinition) error {
				fmt.Println("Iterating models")
				modelname = strings.ToLower(codegen.Goify(model.Name, false))
				modeldir := filepath.Join(outdir, codegen.Goify(model.Name, false))

				if err := os.MkdirAll(modeldir, 0755); err != nil {
					return nil
				}

				filename = fmt.Sprintf("%s_gen.go", modelname)
				utFile := filepath.Join(modeldir, filename)
				fmt.Println(filename)
				fmt.Println(utFile)
				err := os.RemoveAll(utFile)
				if err != nil {
					fmt.Println(err)
				}
				utWr, err := NewUserTypesWriter(utFile)
				if err != nil {
					panic(err) // bug
				}
				title := fmt.Sprintf("%s: Models", version.Context())
				imports := []*codegen.ImportSpec{
					codegen.SimpleImport("github.com/jinzhu/gorm"),
					codegen.SimpleImport("golang.org/x/net/context"),
				}

				utWr.WriteHeader(title, modelname, imports)
				data := &UserTypeTemplateData{
					APIDefinition: api,
					UserType:      model,
					DefaultPkg:    TargetPackage,
					AppPkg:        AppPackage,
				}
				err = utWr.Execute(data)
				g.genfiles = append(g.genfiles, utFile)
				if err != nil {
					fmt.Println(err)
					return err
				}
				err = utWr.FormatCode()
				if err != nil {
					fmt.Println(err)
				}
				return err

			})
			return err
		})
		return err
	})
	return err
}

/*
	err = version.IterateUserTypes(func(t *design.UserTypeDefinition) error {

	})
*/
