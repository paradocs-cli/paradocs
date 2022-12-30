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

	directories, err := walkDirectories(s)

	terraformDirectories, err := checkTerraformDirectories(directories)
	if err != nil {
		return nil, err
	}

	return terraformDirectories, nil
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
		Final.Vars = append(Final.Vars, loadVariables(config.Variables)...)

		log.Printf(color.HiMagentaString("[INFO] checking for resources in directory: %s", v))
		Final.Resources = append(Final.Resources, loadResources(config.ManagedResources)...)

		log.Printf(color.HiMagentaString("[INFO] checking for data resources in directory: %s", v))
		Final.Datas = append(Final.Datas, loadDataResources(config.DataResources)...)

		log.Printf(color.HiMagentaString("[INFO] checking for module calls in directory: %s", v))
		Final.Modules = append(Final.Modules, loadModules(config.ModuleCalls)...)

		log.Printf(color.HiMagentaString("[INFO] checking for output calls in directory: %s", v))
		Final.Outputs = append(Final.Outputs, loadOutputs(config.Outputs)...)

		log.Printf(color.HiMagentaString("[INFO] checking for provider configs in directory: %s", v))
		Final.Providers = append(Final.Providers, loadProviders(config.ProviderConfigs)...)
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

func loadDataResources(m map[string]*tfconfig.Resource) []Data {
	var data []Data
	for _, v := range m {
		data = append(data, Data{
			DataType:               v.Type,
			Name:                   v.Name,
			ProviderName:           v.Provider.Name,
			ProviderAlias:          v.Provider.Alias,
			SourcePositionFileName: v.Pos.Filename,
			SourcePositionLine:     strconv.Itoa(v.Pos.Line),
			Link:                   DataResourceLinkBuilder(v.Provider.Name, v.Type),
		})
	}
	return data
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

func loadOutputs(m map[string]*tfconfig.Output) []Output {
	var outer []Output
	for _, v := range m {
		outer = append(outer, Output{
			Name:                   v.Name,
			Description:            v.Description,
			Sensitive:              v.Sensitive,
			SourcePositionFileName: v.Pos.Filename,
			SourcePositionLine:     strconv.Itoa(v.Pos.Line),
		})
	}
	return outer
}

func loadProviders(m map[string]*tfconfig.ProviderConfig) []Provider {
	var provider []Provider
	for _, v := range m {
		provider = append(provider, Provider{
			Name:  v.Name,
			Alias: v.Alias,
			Link:  ProviderLinkBuilder(v.Name),
		})
	}
	return provider
}

func walkDirectories(directory string) ([]string, error) {
	var dirs []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		dirs = append(dirs, path)
		return nil
	})
	fmt.Println(len(dirs))
	if err != nil {
		return nil, err
	}
	return dirs, nil
}

func checkTerraformDirectories(directories []string) ([]string, error) {
	var terraDirs []string
	for _, v := range directories {
		check := tfconfig.IsModuleDir(v)
		if check {
			terraDirs = append(terraDirs, v)
			log.Printf(color.HiMagentaString("[INFO] 😎found config files for directory: %s😎", v))
		}
	}

	if len(terraDirs) == 0 {
		return nil, fmt.Errorf(color.HiRedString("[ERROR] 🚨found no configuration files for specified directory!!!!🚨"))
	}
	return terraDirs, nil
}
