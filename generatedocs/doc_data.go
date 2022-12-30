package generate_docs

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

//GetDirs takes a string argument and returns a slice of string of directories that containt terraform files
func GetDirs(s string) ([]string, error) {
	if s == "." {
		log.Printf(color.HiMagentaString("[INFO] reading directory info for: root"))
	} else {
		log.Printf(color.HiMagentaString("[INFO] reading directory info for: %s", s))
	}

	var mods []string
	var dirs []string
	err := filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		dirs = append(dirs, path)
		return nil
	})
	if err != nil {
		return mods, err
	}
	for _, v := range dirs {
		check := tfconfig.IsModuleDir(v)
		if check {
			mods = append(mods, v)
			log.Printf(color.HiMagentaString("[INFO] 😎found config files for directory: %s😎", v))
		}
	}
	if len(mods) == 0 {
		log.Fatalf(color.HiRedString("[ERROR] 🚨found no configuration files for specified directory!!!!🚨"))
	}
	return mods, nil
}

//GetData takes a slice of sting and creates JSON objects from data retrieved such as variables, resources, modules, etc
func GetData(s string) (Stats, error) {
	var Final Stats
	get, err := GetDirs(s)
	if err != nil {
		return Stats{}, err
	}
	log.Printf(color.HiMagentaString("[INFO] loading data for directories: %s", s))
	for _, v := range get {
		config, errMod := tfconfig.LoadModule(v)
		if errMod != nil {
			return Stats{}, err
		}
		log.Printf(color.HiMagentaString("[INFO] checking for variables in directory: %s", v))
		Final.Vars = loadVariables(config.Variables)

		log.Printf(color.HiMagentaString("[INFO] checking for resources in directory: %s", v))
		Final.Resources = loadResources(config.ManagedResources)

		log.Printf(color.HiMagentaString("[INFO] checking for module calls in directory: %s", v))
		Final.Modules = loadModules(config.ModuleCalls)

		log.Printf(color.HiMagentaString("[INFO] checking for output calls in directory: %s", v))
		for _, w := range config.Outputs {
			Final.Outputs = append(Final.Outputs, Output{
				Name:                   w.Name,
				Description:            w.Description,
				Sensitive:              w.Sensitive,
				SourcePositionFileName: w.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(w.Pos.Line),
			})
		}
		log.Printf(color.HiMagentaString("[INFO] checking for data resources in directory: %s", v))
		for _, u := range config.DataResources {
			Final.Datas = append(Final.Datas, Data{
				DataType:               u.Type,
				Name:                   u.Name,
				ProviderName:           u.Provider.Name,
				ProviderAlias:          u.Provider.Alias,
				SourcePositionFileName: u.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(u.Pos.Line),
				Link:                   DataResourceLinkBuilder(u.Provider.Name, u.Type),
			})
		}
		log.Printf(color.HiMagentaString("[INFO] checking for provider configs in directory: %s", v))
		for _, u := range config.ProviderConfigs {
			Final.Providers = append(Final.Providers, Provider{
				Name:  u.Name,
				Alias: u.Alias,
				Link:  ProviderLinkBuilder(u.Name),
			})
		}
	}
	return Final, nil
}

func loadVariables(m map[string]*tfconfig.Variable) []Var {
	var vars []Var
	for _, v := range m {
		vars = append(vars, Var{
			VarName:                v.Name,
			VarType:                v.Type,
			VarDefault:             v.Default,
			VarDescription:         v.Description,
			VarRequired:            v.Required,
			VarSensitive:           v.Sensitive,
			SourcePositionFileName: fmt.Sprintf("./%s", v.Pos.Filename),
			SourcePositionLine:     strconv.Itoa(v.Pos.Line),
		})
	}
	return vars
}

func loadResources(m map[string]*tfconfig.Resource) []Resource {
	var res []Resource
	for _, v := range m {
		res = append(res, Resource{
			Mode:                   v.Mode.String(),
			Type:                   v.Type,
			Name:                   v.Name,
			ProviderName:           v.Provider.Name,
			ProviderAlias:          v.Provider.Alias,
			SourcePositionFileName: fmt.Sprintf("./%s", v.Pos.Filename),
			SourcePositionLine:     strconv.Itoa(v.Pos.Line),
			Link:                   ResourceLinkBuilder(v.Provider.Name, v.Type),
		})
	}
	return res
}

func loadModules(m map[string]*tfconfig.ModuleCall) []Module {
	var mod []Module
	for _, v := range m {
		mod = append(mod, Module{
			Name:                   v.Name,
			ModSource:              v.Source,
			Version:                v.Version,
			SourcePositionFileName: v.Pos.Filename,
			SourcePositionLine:     strconv.Itoa(v.Pos.Line),
		})
	}
	return mod
}
