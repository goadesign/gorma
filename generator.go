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

var storageGroup *StorageGroup

// Generator is the application code generator.
type Generator struct {
	*codegen.GoGenerator
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
	if err = os.MkdirAll(outdir, 0777); err != nil {
		return nil, err
	}
	return &Generator{
		GoGenerator: codegen.NewGoGenerator(outdir),
		genfiles:    []string{outdir},
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
	storageGroup, err = NewStorageGroup(api)
	if err != nil {
		fmt.Println("Error Parsing API: ", err)
	}
	// create the output directory
	outdir := ModelOutputDir()
	if err := os.MkdirAll(outdir, 0755); err != nil {
		return g.genfiles, err
	}

	// models are unversioned - outside the loop
	if err := g.generateUserTypes(outdir, api); err != nil {
		return g.genfiles, err
	}

	err = api.IterateVersions(func(v *design.APIVersionDefinition) error {
		if err := g.generatePayloadHelpers(outdir, storageGroup, api, v); err != nil {
			return err
		}
		if err := g.generateMediaTypes(outdir, storageGroup, v); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return g.genfiles, nil

}

// Cleanup doesn't do anything here.  Move along
func (g *Generator) Cleanup() {
	if len(g.genfiles) == 0 {
		return
	}
	g.genfiles = nil
}

// Generated package name for resources supporting the given version.
func packageName(version *design.APIVersionDefinition) (pack string) {
	pack = TargetPackage
	if version.Version != "" {
		pack = codegen.Goify(codegen.VersionPackage(version.Version), false)
	}
	return
}

// generateContexts iterates through the version resources and actions and generates the action
// contexts.
func (g *Generator) generatePayloadHelpers(verdir string, sg *StorageGroup, api *design.APIDefinition, version *design.APIVersionDefinition) error {

	err := version.IterateResources(func(r *design.ResourceDefinition) error {
		actionable := false
		err := r.IterateActions(func(ad *design.ActionDefinition) error {
			if hasUserType(ad) {
				fmt.Printf("Version %s has actions\n", version.Version)
				actionable = true
			}
			return nil

		})
		if !actionable {
			return nil
		}
		if !r.SupportsVersion(version.Version) {
			return nil
		}
		prefix := "_resource"
		if version.Version != "" {
			prefix = prefix + "_" + codegen.Goify(version.Version, false)

		}
		name := strings.ToLower(codegen.Goify(r.Name, false))
		ctxFile := filepath.Join(verdir, name, name+prefix+"_gen.go")
		ctxWr, err := NewContextsWriter(ctxFile)
		if err != nil {
			panic(err) // bug
		}
		title := fmt.Sprintf("%s: Resource and Payload Helpers", version.Context())
		imports := []*codegen.ImportSpec{
			codegen.SimpleImport("fmt"),
			codegen.SimpleImport("strconv"),
			codegen.SimpleImport("github.com/raphael/goa"),
		}
		if !version.IsDefault() {
			appPkg, err := ModelPackagePath()
			if err != nil {
				return err
			}
			imports = append(imports, codegen.SimpleImport(appPkg))
		}
		os.Remove(ctxFile)
		fmt.Println("writing header")
		ctxWr.WriteHeader(title, name, imports)
		ctxData := NewConversionData(version, r)
		err = ctxWr.Execute(&ctxData)
		if err != nil {
			return err
		}
		g.genfiles = append(g.genfiles, ctxFile)
		return ctxWr.FormatCode()
	})
	return err
}

// generateMediaTypes iterates through the media types and generate the data structures and
// marshaling code.
func (g *Generator) generateMediaTypes(verdir string, sg *StorageGroup, version *design.APIVersionDefinition) error {
	err := version.IterateMediaTypes(func(mt *design.MediaTypeDefinition) error {
		if !mt.SupportsVersion(version.Version) {
			return nil
		}
		prefix := "_media"
		if version.Version != "" {
			prefix = prefix + "_" + codegen.Goify(version.Version, false)

		}
		name := strings.ToLower(codegen.Goify(mt.TypeName, false))
		dirname := strings.Replace(name, "collection", "", -1)
		err := os.MkdirAll(filepath.Join(verdir, dirname), 0755)
		if err != nil {
			panic(err)
		}
		mtFile := filepath.Join(verdir, dirname, name+prefix+"_gen.go")
		os.Remove(mtFile)
		mtWr, err := NewMediaTypesWriter(mtFile)
		if err != nil {
			panic(err) // bug
		}
		title := fmt.Sprintf("%s: Application Media Helpers", version.Context())
		imports := []*codegen.ImportSpec{
			codegen.SimpleImport("github.com/raphael/goa"),
			codegen.SimpleImport("fmt"),
		}
		mtWr.WriteHeader(title, dirname, imports)
		data := &MediaTypeTemplateData{
			MediaType:  mt,
			Versioned:  version.Version != "",
			DefaultPkg: TargetPackage,
		}
		if mt.Type.IsObject() || mt.Type.IsArray() {
			err = mtWr.Execute(data)
		}
		if err != nil {
			return err
		}
		g.genfiles = append(g.genfiles, mtFile)
		err = mtWr.FormatCode()
		return err
	})
	return err
}

// generateUserTypes iterates through the user types and generates the data structures and
// marshaling code.
func (g *Generator) generateUserTypes(verdir string, api *design.APIDefinition) error {
	err := storageGroup.RelationalStore.IterateModels(func(m *RelationalModel) error {

		pkgName := strings.ToLower(m.Name)
		err := os.MkdirAll(filepath.Join(verdir, pkgName), 0755)
		if err != nil {
			return err
		}
		filename := fmt.Sprintf("%s_gen.go", codegen.Goify(m.Name, false))
		_ = os.Remove(filepath.Join(verdir, pkgName, filename))
		utFile := filepath.Join(verdir, pkgName, filename)
		utWr, err := NewUserTypesWriter(utFile)
		if err != nil {
			panic(err) // bug
		}
		title := fmt.Sprintf("Generated Models")
		imports := []*codegen.ImportSpec{
			codegen.SimpleImport("github.com/raphael/goa"),
			codegen.SimpleImport("github.com/patrickmn/go-cache"),
			codegen.SimpleImport("fmt"),
		}

		var utd *design.UserTypeDefinition
		// find the right UserTypeDefinition for this RelationalModel
		err = api.IterateUserTypes(func(ut *design.UserTypeDefinition) error {
			fmt.Println(ut.TypeName, m.Name)
			if deModel(ut.TypeName) == m.Name {
				utd = ut
			}
			return nil
		})
		utWr.WriteHeader(title, codegen.Goify(m.Name, false), imports)
		data := &UserTypeTemplateData{
			UserType:   m,
			BaseType:   utd,
			DefaultPkg: TargetPackage,
		}
		err = utWr.Execute(data)
		if err != nil {
			return err
		}
		g.genfiles = append(g.genfiles, utFile)
		//return err
		return utWr.FormatCode()

	})

	return err
}
