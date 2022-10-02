package generate_docs

import (
	"fmt"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"os"
	"path/filepath"
	"strconv"
)

//GetDirs takes a string argument and returns a slice of string of directories that containt terraform files
func GetDirs(s string) ([]string, error) {
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
		}
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
	for _, v := range get {
		config, errMod := tfconfig.LoadModule(v)
		if errMod != nil {
			return Stats{}, err
		}
		for _, z := range config.Variables {
			Final.Vars = append(Final.Vars, Var{
				VarName:                z.Name,
				VarType:                z.Type,
				VarDefault:             z.Default,
				VarDescription:         z.Description,
				VarRequired:            z.Required,
				VarSensitive:           z.Sensitive,
				SourcePositionFileName: fmt.Sprintf("./%s", z.Pos.Filename),
				SourcePositionLine:     strconv.Itoa(z.Pos.Line),
			})
		}

		for _, x := range config.ManagedResources {
			Final.Resources = append(Final.Resources, Resource{
				Mode:                   x.Mode.String(),
				Type:                   x.Type,
				Name:                   x.Name,
				ProviderName:           x.Provider.Name,
				ProviderAlias:          x.Provider.Alias,
				SourcePositionFileName: fmt.Sprintf("./%s", x.Pos.Filename),
				SourcePositionLine:     strconv.Itoa(x.Pos.Line),
				Link:                   fmt.Sprintf("%s/resources/%s", LinkBuilder(x.Provider.Name), x.Type),
			})
		}

		for _, y := range config.ModuleCalls {
			Final.Modules = append(Final.Modules, Module{
				Name:                   y.Name,
				ModSource:              y.Source,
				Version:                y.Version,
				SourcePositionFileName: y.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(y.Pos.Line),
			})
		}

		for _, w := range config.Outputs {
			Final.Outputs = append(Final.Outputs, Output{
				Name:                   w.Name,
				Description:            w.Description,
				Sensitive:              w.Sensitive,
				SourcePositionFileName: w.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(w.Pos.Line),
			})
		}

		for _, u := range config.DataResources {
			Final.Datas = append(Final.Datas, Data{
				DataType:               u.Type,
				Name:                   u.Name,
				ProviderName:           u.Provider.Name,
				ProviderAlias:          u.Provider.Alias,
				SourcePositionFileName: u.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(u.Pos.Line),
				Link:                   fmt.Sprintf("%s/data-sources/%s", LinkBuilder(u.Provider.Name), u.Type),
			})
		}

		for _, u := range config.ProviderConfigs {
			Final.Providers = append(Final.Providers, Provider{
				Name:  u.Name,
				Alias: u.Alias,
				Link:  LinkBuilder(u.Name),
			})
		}
	}
	return Final, nil
}
