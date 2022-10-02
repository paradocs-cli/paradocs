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
		return mods, fmt.Errorf(err.Error())
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
		return Final, fmt.Errorf(err.Error())
	}
	for _, v := range get {
		config, err := tfconfig.LoadModule(v)
		if err != nil {
			return Final, fmt.Errorf(err.Error())
		}
		for _, z := range config.Variables {
			someVars := Var{
				VarName:                z.Name,
				VarType:                z.Type,
				VarDefault:             z.Default,
				VarDescription:         z.Description,
				VarRequired:            z.Required,
				VarSensitive:           z.Sensitive,
				SourcePositionFileName: fmt.Sprintf("./%s", z.Pos.Filename),
				SourcePositionLine:     strconv.Itoa(z.Pos.Line),
			}

			Final.Vars = append(Final.Vars, someVars)
		}

		for _, x := range config.ManagedResources {
			someResources := Resource{
				Mode:                   x.Mode.String(),
				Type:                   x.Type,
				Name:                   x.Name,
				ProviderName:           x.Provider.Name,
				ProviderAlias:          x.Provider.Alias,
				SourcePositionFileName: fmt.Sprintf("./%s", x.Pos.Filename),
				SourcePositionLine:     strconv.Itoa(x.Pos.Line),
				Link:                   fmt.Sprintf("%s/resources/%s", LinkBuilder(x.Provider.Name), x.Type),
			}
			Final.Resources = append(Final.Resources, someResources)
		}

		for _, y := range config.ModuleCalls {
			someModules := Module{
				Name:                   y.Name,
				ModSource:              y.Source,
				Version:                y.Version,
				SourcePositionFileName: y.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(y.Pos.Line),
			}
			Final.Modules = append(Final.Modules, someModules)
		}

		for _, w := range config.Outputs {
			someOutputs := Output{
				Name:                   w.Name,
				Description:            w.Description,
				Sensitive:              w.Sensitive,
				SourcePositionFileName: w.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(w.Pos.Line),
			}
			Final.Outputs = append(Final.Outputs, someOutputs)
		}

		for _, u := range config.DataResources {
			someData := Data{
				DataType:               u.Type,
				Name:                   u.Name,
				ProviderName:           u.Provider.Name,
				ProviderAlias:          u.Provider.Alias,
				SourcePositionFileName: u.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(u.Pos.Line),
				Link:                   fmt.Sprintf("%s/data-sources/%s", LinkBuilder(u.Provider.Name), u.Type),
			}
			Final.Datas = append(Final.Datas, someData)
		}

		for _, u := range config.ProviderConfigs {
			someProvider := Provider{
				Name:  u.Name,
				Alias: u.Alias,
				Link:  LinkBuilder(u.Name),
			}
			Final.Providers = append(Final.Providers, someProvider)
		}
	}
	return Final, nil
}
