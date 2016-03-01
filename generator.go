package gorma

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/goadesign/goa/goagen/utils"
	"github.com/spf13/cobra"
)

// Generator is the application code generator.
type Generator struct {
	genfiles []string
}

// Generate is the generator entry point called by the meta generator.
func Generate(roots []interface{}) (files []string, err error) {
	api := roots[0].(*design.APIDefinition)
	g := new(Generator)
	root := &cobra.Command{
		Use:   "gorma",
		Short: "Code generator",
		Long:  "database model code generator",
		PreRunE: func(*cobra.Command, []string) error {
			outdir := ModelOutputDir()
			g.genfiles = []string{outdir}
			err = os.MkdirAll(outdir, 0777)
			return err
		},
		Run: func(*cobra.Command, []string) { files, err = g.Generate(api) },
	}
	codegen.RegisterFlags(root)
	NewCommand().RegisterFlags(root)
	root.Execute()
	return
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
	if err := g.generateUserHelpers(outdir, api); err != nil {
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

// generateUserTypes iterates through the user types and generates the data structures and
// marshaling code.
func (g *Generator) generateUserTypes(outdir string, api *design.APIDefinition) error {
	var modelname, filename string
	err := GormaDesign.IterateStores(func(store *RelationalStoreDefinition) error {
		err := store.IterateModels(func(model *RelationalModelDefinition) error {
			modelname = strings.ToLower(codegen.Goify(model.ModelName, false))

			filename = fmt.Sprintf("%s.go", modelname)
			utFile := filepath.Join(outdir, filename)
			err := os.RemoveAll(utFile)
			if err != nil {
				fmt.Println(err)
			}
			utWr, err := NewUserTypesWriter(utFile)
			if err != nil {
				panic(err) // bug
			}
			title := fmt.Sprintf("%s: Models", api.Context())
			ap, err := AppPackagePath()
			if err != nil {
				panic(err)
			}
			imports := []*codegen.ImportSpec{
				codegen.SimpleImport(ap),
				codegen.SimpleImport("time"),
				codegen.SimpleImport("github.com/goadesign/goa"),
				codegen.SimpleImport("github.com/jinzhu/gorm"),
				codegen.SimpleImport("golang.org/x/net/context"),
				codegen.SimpleImport("golang.org/x/net/context"),
			}

			if model.Cached {
				imp := codegen.NewImport("cache", "github.com/patrickmn/go-cache")
				imports = append(imports, imp)
				imp = codegen.SimpleImport("strconv")
				imports = append(imports, imp)
			}
			utWr.WriteHeader(title, TargetPackage, imports)
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
}

// generateUserHelpers iterates through the user types and generates the data structures and
// marshaling code.
func (g *Generator) generateUserHelpers(outdir string, api *design.APIDefinition) error {
	var modelname, filename string
	err := GormaDesign.IterateStores(func(store *RelationalStoreDefinition) error {
		err := store.IterateModels(func(model *RelationalModelDefinition) error {
			modelname = strings.ToLower(codegen.Goify(model.ModelName, false))

			filename = fmt.Sprintf("%s_helper.go", modelname)
			utFile := filepath.Join(outdir, filename)
			err := os.RemoveAll(utFile)
			if err != nil {
				fmt.Println(err)
			}
			utWr, err := NewUserHelperWriter(utFile)
			if err != nil {
				panic(err) // bug
			}
			title := fmt.Sprintf("%s: Model Helpers", api.Context())
			ap, err := AppPackagePath()
			if err != nil {
				panic(err)
			}
			imports := []*codegen.ImportSpec{
				codegen.SimpleImport(ap),
				codegen.SimpleImport("time"),
				codegen.SimpleImport("github.com/goadesign/goa"),
				codegen.SimpleImport("github.com/jinzhu/gorm"),
				codegen.SimpleImport("golang.org/x/net/context"),
				codegen.SimpleImport("golang.org/x/net/context"),
			}

			if model.Cached {
				imp := codegen.NewImport("cache", "github.com/patrickmn/go-cache")
				imports = append(imports, imp)
				imp = codegen.SimpleImport("strconv")
				imports = append(imports, imp)
			}
			utWr.WriteHeader(title, TargetPackage, imports)
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
}
