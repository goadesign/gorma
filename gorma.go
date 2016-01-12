package gorma

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/raphael/goa/design"
	"github.com/raphael/goa/goagen/codegen"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
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
	return new(Generator), nil
}

// Generate produces the generated model files
func (g *Generator) Generate(api *design.APIDefinition) ([]string, error) {
	if api == nil {
		return nil, fmt.Errorf("missing API definition")
	}

	// RBAC is unversioned, do it first
	if err := g.generateRBAC(api); err != nil {
		return nil, err
	}
	if err := g.generateModels(api); err != nil {
		return nil, err
	}
	if err := g.generateImpls(api); err != nil {
		return nil, err
	}
	if err := g.generateMedia(api); err != nil {
		return nil, err
	}
	if err := g.generateResources(api); err != nil {
		return nil, err
	}
	return g.genfiles, nil
}

// Generate produces the implementation model files
func (g *Generator) generateImpls(api *design.APIDefinition) error {
	app := kingpin.New("Model generator", "model generator")
	codegen.RegisterFlags(app)
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	outdir := implDir()
	gopath := filepath.SplitList(os.Getenv("GOPATH"))[0]

	mainimp, err := filepath.Rel(filepath.Join(gopath, "src"), codegen.OutputDir)
	if err != nil {
		return err
	}
	mainimp = filepath.ToSlash(mainimp)
	imp := path.Join(mainimp, "app")
	imports := []*codegen.ImportSpec{
		codegen.SimpleImport(imp),
		codegen.SimpleImport("github.com/jinzhu/gorm"),
		codegen.SimpleImport("github.com/jinzhu/copier"),
		codegen.SimpleImport("time"),
	}
	// get the imports for the app packages
	api.IterateVersions(func(v *design.APIVersionDefinition) error {
		if v.IsDefault() {
			return nil
		}
		imports = append(imports, codegen.SimpleImport(imp+"/"+codegen.Goify(v.Version, false)))
		return nil
	})
	// Now generate the models, by iterating the versions
	err = api.IterateVersions(func(v *design.APIVersionDefinition) error {
		verdir := outdir
		if v.Version != "" {
			return nil
		}
		if err := os.MkdirAll(verdir, 0755); err != nil {
			return err
		}
		var outPkg string
		// going to hell for this == HELP Wanted (windows) TODO:(BJK)
		outPkg = codegen.DesignPackagePath[0:strings.LastIndex(codegen.DesignPackagePath, "/")]
		if err != nil {
			panic(err)
		}
		outPkg = strings.TrimPrefix(outPkg, "src/")

		_, cached := metaLookup(api.Metadata, "#cached")
		if cached {
			imports = append(imports, codegen.SimpleImport("github.com/patrickmn/go-cache"))
		}

		err = v.IterateUserTypes(func(res *design.UserTypeDefinition) error {
			if res.Type.IsObject() {
				title := fmt.Sprintf("%s: Models", api.Name)
				name := strings.ToLower(deModel(res.TypeName))

				err := os.MkdirAll(filepath.Join(implDir()), 0755)
				if err != nil {
					panic(err)
				}

				filename := filepath.Join(verdir, name+"_model.go")
				os.Remove(filename)
				mtw, err := NewImplWriter(filename)
				if err != nil {
					panic(err)
				}

				md := NewImplData(v.Version, res)
				for k := range md.RequiredPackages {
					imports = append(imports, codegen.SimpleImport(path.Join(mainimp, modelDir(), k)))
				}
				imports = append(imports, codegen.SimpleImport(path.Join(mainimp, "models")))
				mtw.WriteHeader(title, name, imports)
				if m, ok := metaLookup(res.Metadata, ""); ok && m == "Model" {
					err = mtw.Execute(&md)
					if err != nil {
						fmt.Println("Error executing Gorma: ", err.Error())
						g.Cleanup()
						return err
					}
				}
				if err := mtw.FormatCode(); err != nil {
					fmt.Println("Error executing Gorma: ", err.Error())
					g.Cleanup()

				}
				if err == nil {
					g.genfiles = append(g.genfiles, filename)
				}
				return nil
			}

			return nil

		})
		return nil
	})

	return err
}

// Generate produces the generated model files
func (g *Generator) generateModels(api *design.APIDefinition) error {
	app := kingpin.New("Model generator", "model generator")
	codegen.RegisterFlags(app)
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	outdir := modelDir()
	gopath := filepath.SplitList(os.Getenv("GOPATH"))[0]

	mainimp, err := filepath.Rel(filepath.Join(gopath, "src"), codegen.OutputDir)
	if err != nil {
		return err
	}
	mainimp = filepath.ToSlash(mainimp)
	imp := path.Join(mainimp, "app")
	imports := []*codegen.ImportSpec{
		codegen.SimpleImport(imp),
		codegen.SimpleImport("github.com/jinzhu/gorm"),
		codegen.SimpleImport("github.com/jinzhu/copier"),
		codegen.SimpleImport("time"),
	}
	// get the imports for the app packages
	api.IterateVersions(func(v *design.APIVersionDefinition) error {
		if v.IsDefault() {
			return nil
		}
		imports = append(imports, codegen.SimpleImport(imp+"/"+codegen.Goify(v.Version, false)))
		return nil
	})
	// Now generate the models, by iterating the versions
	err = api.IterateVersions(func(v *design.APIVersionDefinition) error {
		verdir := outdir
		if v.Version != "" {
			return nil
		}
		if err := os.MkdirAll(verdir, 0755); err != nil {
			return err
		}
		var outPkg string
		// going to hell for this == HELP Wanted (windows) TODO:(BJK)
		outPkg = codegen.DesignPackagePath[0:strings.LastIndex(codegen.DesignPackagePath, "/")]
		if err != nil {
			panic(err)
		}
		outPkg = strings.TrimPrefix(outPkg, "src/")

		_, cached := metaLookup(api.Metadata, "#cached")
		if cached {
			imports = append(imports, codegen.SimpleImport("github.com/patrickmn/go-cache"))
		}

		err = v.IterateUserTypes(func(res *design.UserTypeDefinition) error {
			if res.Type.IsObject() {
				title := fmt.Sprintf("%s: Models", api.Name)
				name := strings.ToLower(deModel(res.TypeName))

				err := os.MkdirAll(filepath.Join(modelDir()), 0755)
				if err != nil {
					panic(err)
				}

				filename := filepath.Join(verdir, name+"_genmodel.go")
				os.Remove(filename)
				mtw, err := NewModelWriter(filename)
				if err != nil {
					panic(err)
				}

				md := NewModelData(v.Version, res)
				for k, _ := range md.RequiredPackages {
					imports = append(imports, codegen.SimpleImport(path.Join(mainimp, "gorma", k)))
				}

				mtw.WriteHeader(title, name, imports)
				if m, ok := metaLookup(res.Metadata, ""); ok && m == "Model" {
					err = mtw.Execute(&md)
					if err != nil {
						fmt.Println("Error executing Gorma: ", err.Error())
						g.Cleanup()
						return err
					}
				}
				if err := mtw.FormatCode(); err != nil {
					fmt.Println("Error executing Gorma: ", err.Error())
					g.Cleanup()

				}
				if err == nil {
					g.genfiles = append(g.genfiles, filename)
				}
				return nil
			}

			return nil

		})
		return nil
	})

	return err
}

// Generate produces the generated rbac files
func (g *Generator) generateRBAC(api *design.APIDefinition) error {
	err := os.MkdirAll(modelDir(), 0755)
	if err != nil {
		panic(err)
	}
	app := kingpin.New("Model generator", "model generator")
	codegen.RegisterFlags(app)
	_, err = app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	var outPkg string
	// going to hell for this == HELP Wanted (windows) TODO:(BJK)
	outPkg = codegen.DesignPackagePath[0:strings.LastIndex(codegen.DesignPackagePath, "/")]
	if err != nil {
		panic(err)
	}
	outPkg = strings.TrimPrefix(outPkg, "src/")
	appPkg := filepath.Join(outPkg, "app")

	rbacimports := []*codegen.ImportSpec{
		codegen.SimpleImport(appPkg),
		codegen.SimpleImport("github.com/mikespook/gorbac"),
	}

	rbactitle := fmt.Sprintf("%s: RBAC", api.Name)
	_, dorbac := metaLookup(api.Metadata, "#rbac")

	if dorbac {
		rbacfilename := filepath.Join(modelDir(), "rbac_genmodel.go")
		os.Remove(rbacfilename)
		rbacw, err := NewRbacWriter(rbacfilename)
		if err != nil {
			fmt.Println("Error executing Gorma: ", err.Error())
			panic(err)
		}
		rbacw.WriteHeader(rbactitle, "models", rbacimports)
		err = rbacw.Execute(api)
		if err != nil {
			fmt.Println("Error executing Gorma: ", err.Error())
			g.Cleanup()
			return err
		}
		if err := rbacw.FormatCode(); err != nil {
			fmt.Println("Error executing Gorma: ", err.Error())
			g.Cleanup()
			return err
		}
		if err == nil {
			g.genfiles = append(g.genfiles, rbacfilename)
		}
	}

	return err
}

// Generate produces the generated media files
func (g *Generator) generateResources(api *design.APIDefinition) error {
	os.MkdirAll(modelDir(), 0755)
	app := kingpin.New("Model generator", "model generator")
	codegen.RegisterFlags(app)
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	gopath := filepath.SplitList(os.Getenv("GOPATH"))[0]

	mainimp, err := filepath.Rel(filepath.Join(gopath, "src"), codegen.OutputDir)
	if err != nil {
		return err
	}
	mainimp = filepath.ToSlash(mainimp)
	imp := path.Join(mainimp, "app")
	imports := []*codegen.ImportSpec{
		codegen.SimpleImport(imp),
		codegen.SimpleImport("github.com/jinzhu/copier"),
	}
	// get the imports for the app packages
	api.IterateVersions(func(v *design.APIVersionDefinition) error {
		if v.IsDefault() {
			return nil
		}
		imports = append(imports, codegen.SimpleImport(filepath.Join(filepath.ToSlash(imp), codegen.Goify(codegen.VersionPackage(v.Version), false))))
		return nil
	})

	title := fmt.Sprintf("%s: Media Helpers", api.Name)

	err = api.IterateVersions(func(v *design.APIVersionDefinition) error {
		err = v.IterateResources(func(res *design.ResourceDefinition) error {
			actionable := false
			err = res.IterateActions(func(ad *design.ActionDefinition) error {
				if hasUserType(ad) {
					actionable = true
				}
				return nil

			})
			if !actionable {
				return nil
			}

			if !res.SupportsVersion(v.Version) {
				return nil
			}
			prefix := "resource"
			if v.Version != "" {
				prefix = prefix + "_v" + codegen.Goify(v.Version, false)

			}
			name := strings.ToLower(codegen.Goify(res.Name, false))

			err := os.MkdirAll(filepath.Join(modelDir()), 0755)
			if err != nil {
				panic(err)
			}

			mediafilename := filepath.Join(modelDir(), name+"_"+prefix+"_genmodel.go")
			os.Remove(mediafilename)

			resw, err := NewResourceWriter(mediafilename)
			if err != nil {
				fmt.Println("Error executing Gorma: ", err.Error())
				panic(err)
			}

			rd := NewResourceData(v.Version, res)
			for k, _ := range rd.RequiredPackages {
				imports = append(imports, codegen.SimpleImport(path.Join(mainimp, "gorma", k)))
			}
			resw.WriteHeader(title, name, imports)

			err = resw.Execute(&rd)
			if err != nil {
				fmt.Println("Error executing Gorma: ", err.Error())
				g.Cleanup()
				return err
			}
			if err := resw.FormatCode(); err != nil {
				fmt.Println("Error executing Gorma: ", err.Error())
				g.Cleanup()
				return err
			}
			if err == nil {
				g.genfiles = append(g.genfiles, mediafilename)
			}
			return nil
		})
		return nil
	})
	return err
}

// Generate produces the generated media files
func (g *Generator) generateMedia(api *design.APIDefinition) error {
	os.MkdirAll(modelDir(), 0755)
	app := kingpin.New("Model generator", "model generator")
	codegen.RegisterFlags(app)
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}
	gopath := filepath.SplitList(os.Getenv("GOPATH"))[0]

	mainimp, err := filepath.Rel(filepath.Join(gopath, "src"), codegen.OutputDir)
	if err != nil {
		return err
	}
	mainimp = filepath.ToSlash(mainimp)
	imp := path.Join(mainimp, "app")
	imports := []*codegen.ImportSpec{
		codegen.SimpleImport(imp),
		codegen.SimpleImport("github.com/jinzhu/copier"),
	}
	// get the imports for the app packages
	api.IterateVersions(func(v *design.APIVersionDefinition) error {
		if v.IsDefault() {
			return nil
		}
		imports = append(imports, codegen.SimpleImport(imp+"/"+codegen.Goify(codegen.VersionPackage(v.Version), false)))
		return nil
	})

	title := fmt.Sprintf("%s: Media Helpers", api.Name)

	err = api.IterateVersions(func(v *design.APIVersionDefinition) error {
		err = v.IterateMediaTypes(func(res *design.MediaTypeDefinition) error {
			if res.Reference == nil {
				// not a mediatype that references a model
				return nil
			}
			if model, ok := res.Reference.(*design.UserTypeDefinition); ok {
				if !modelMetadata(model.Definition()) {
					return nil
				}
			}

			if res.Type.IsObject() {
				if !res.SupportsVersion(v.Version) {
					return nil
				}
				prefix := "media"
				if v.Version != "" {
					prefix = prefix + "_v" + codegen.Goify(v.Version, false)

				}
				name := strings.ToLower(codegen.Goify(res.TypeName, false))

				err := os.MkdirAll(filepath.Join(modelDir()), 0755)
				if err != nil {
					panic(err)
				}

				mediafilename := filepath.Join(modelDir(), name+"_"+prefix+"_genmodel.go")

				os.Remove(mediafilename)
				resw, err := NewMediaWriter(mediafilename)
				if err != nil {
					fmt.Println("Error executing Gorma: ", err.Error())
					panic(err)
				}

				md := NewMediaData(v.Version, res)
				for k, _ := range md.RequiredPackages {
					imports = append(imports, codegen.SimpleImport(path.Join(mainimp, "gorma", k)))
				}
				resw.WriteHeader(title, name, imports)

				err = resw.Execute(&md)
				if err != nil {
					fmt.Println("Error executing Gorma: ", err.Error())
					g.Cleanup()
					return err
				}
				if err := resw.FormatCode(); err != nil {
					fmt.Println("Error executing Gorma: ", err.Error())
					g.Cleanup()
					return err
				}
				if err == nil {
					g.genfiles = append(g.genfiles, mediafilename)
				}
			}
			return nil

		})
		return err
	})
	return err
}

// Cleanup removes all the files generated by this generator during the last invokation of Generate.
func (g *Generator) Cleanup() {
	for _, f := range g.genfiles {
		os.Remove(f)
	}
	g.genfiles = nil
}
