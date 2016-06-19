package gorma

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	"github.com/goadesign/goa/goagen/utils"
)

// Generator is the application code generator.
type Generator struct {
	genfiles   []string // Generated files
	outDir     string   // Absolute path to output directory
	target     string   // Target package name - "models" by default
	appPkg     string   // Generated goa app package name - "app" by default
	appPkgPath string   // Generated goa app package import path
}

// Generate is the generator entry point called by the meta generator.
func Generate() (files []string, err error) {
	var outDir, target, appPkg, ver string

	set := flag.NewFlagSet("gorma", flag.PanicOnError)
	set.String("design", "", "")
	set.StringVar(&outDir, "out", "", "")
	set.StringVar(&ver, "version", "", "")
	set.StringVar(&target, "pkg", "models", "")
	set.StringVar(&appPkg, "app", "app", "")
	set.Parse(os.Args[2:])

	// First check compatibility
	if err := codegen.CheckVersion(ver); err != nil {
		return nil, err
	}

	// Now proceed
	appPkgPath, err := codegen.PackagePath(filepath.Join(outDir, appPkg))
	if err != nil {
		return nil, fmt.Errorf("invalid app package: %s", err)
	}
	outDir = filepath.Join(outDir, target)

	g := &Generator{outDir: outDir, target: target, appPkg: appPkg, appPkgPath: appPkgPath}

	return g.Generate(design.Design)
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
	if err := os.MkdirAll(g.outDir, 0755); err != nil {
		return nil, err
	}

	if err := g.generateUserTypes(g.outDir, api); err != nil {
		return g.genfiles, err
	}
	if err := g.generateUserHelpers(g.outDir, api); err != nil {
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
			imports := []*codegen.ImportSpec{
				codegen.SimpleImport(g.appPkgPath),
				codegen.SimpleImport("time"),
				codegen.SimpleImport("github.com/goadesign/goa"),
				codegen.SimpleImport("github.com/jinzhu/gorm"),
				codegen.SimpleImport("golang.org/x/net/context"),
				codegen.SimpleImport("golang.org/x/net/context"),
				codegen.SimpleImport("github.com/goadesign/goa/uuid"),
			}

			if model.Cached {
				imp := codegen.NewImport("cache", "github.com/patrickmn/go-cache")
				imports = append(imports, imp)
				imp = codegen.SimpleImport("strconv")
				imports = append(imports, imp)
			}
			utWr.WriteHeader(title, g.target, imports)
			data := &UserTypeTemplateData{
				APIDefinition: api,
				UserType:      model,
				DefaultPkg:    g.target,
				AppPkg:        g.appPkgPath,
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
			imports := []*codegen.ImportSpec{
				codegen.SimpleImport(g.appPkgPath),
				codegen.SimpleImport("time"),
				codegen.SimpleImport("github.com/goadesign/goa"),
				codegen.SimpleImport("github.com/jinzhu/gorm"),
				codegen.SimpleImport("golang.org/x/net/context"),
				codegen.SimpleImport("golang.org/x/net/context"),
				codegen.SimpleImport("github.com/goadesign/goa/uuid"),
			}

			if model.Cached {
				imp := codegen.NewImport("cache", "github.com/patrickmn/go-cache")
				imports = append(imports, imp)
				imp = codegen.SimpleImport("strconv")
				imports = append(imports, imp)
			}
			utWr.WriteHeader(title, g.target, imports)
			data := &UserTypeTemplateData{
				APIDefinition: api,
				UserType:      model,
				DefaultPkg:    g.target,
				AppPkg:        g.appPkgPath,
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
