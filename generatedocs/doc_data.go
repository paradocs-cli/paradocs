package generate_docs

import (
	"fmt"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
		newVars := config.Variables
		newResources := config.ManagedResources
		newModules := config.ModuleCalls
		newOutputs := config.Outputs
		newData := config.DataResources
		newProviders := config.ProviderConfigs

		for _, z := range newVars {
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

		for _, x := range newResources {
			someResources := Resource{
				Mode:                   x.Mode.String(),
				Type:                   x.Type,
				Name:                   x.Name,
				ProviderName:           x.Provider.Name,
				ProviderAlias:          x.Provider.Alias,
				SourcePositionFileName: fmt.Sprintf("./%s", x.Pos.Filename),
				SourcePositionLine:     strconv.Itoa(x.Pos.Line),
				Link: fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/%s/latest/docs/resources/%s", x.Provider.Name, strings.TrimPrefix(x.Type, fmt.Sprintf("%s_", x.Provider.Name))),
			}
			Final.Resources = append(Final.Resources, someResources)
		}

		for _, y := range newModules {
			someModules := Module{
				Name:                   y.Name,
				ModSource:              y.Source,
				Version:                y.Version,
				SourcePositionFileName: y.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(y.Pos.Line),
			}
			Final.Modules = append(Final.Modules, someModules)
		}

		for _, w := range newOutputs {
			someOutputs := Output{
				Name:                   w.Name,
				Description:            w.Description,
				Sensitive:              w.Sensitive,
				SourcePositionFileName: w.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(w.Pos.Line),
			}
			Final.Outputs = append(Final.Outputs, someOutputs)
		}

		for _, u := range newData {
			someData := Data{
				DataType:               u.Type,
				Name:                   u.Name,
				ProviderName:           u.Provider.Name,
				ProviderAlias:          u.Provider.Alias,
				SourcePositionFileName: u.Pos.Filename,
				SourcePositionLine:     strconv.Itoa(u.Pos.Line),
				Link: DataSourceLinks(u),
			}
			Final.Datas = append(Final.Datas, someData)
		}

		for _, u := range newProviders {
			someProvider := Provider{
				Name:  u.Name,
				Alias: u.Alias,
				Link: fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/%s/latest/docs", u.Name),
			}
			Final.Providers = append(Final.Providers, someProvider)
		}
	}
	return Final, nil
}

func DataSourceLinks(r *tfconfig.Resource) string{
	if r.Type == "external" {
		return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/%s/latest/docs/data-sources/data_source", r.Provider.Name)
	}
	return fmt.Sprintf("https://registry.terraform.io/providers/hashicorp/%s/latest/docs/data-sources/%s", r.Provider.Name, strings.TrimPrefix(r.Type, fmt.Sprintf("%s_", r.Provider.Name)))
}

//GetDirData iterates through directories and returns data about each directory
func GetDirData(s string) (RepoInfo, error){
	var dirs RepoInfo
	get, err := GetDirs(s)
	if err != nil {
		return dirs, fmt.Errorf(err.Error())
	}
	for _, v := range get {
		read, err := ioutil.ReadDir(v)
		if err != nil {
			return dirs, fmt.Errorf(err.Error())
		}
		for _, z := range read {
			var theDir Dirs
				theDir.Name = z.Name()
				theDir.ModificationTime = z.ModTime()
				if tfconfig.IsModuleDir(z.Name()) {
					theDir.IsTerraDir = true
				} else {
					theDir.IsTerraDir = false
				}
				dirs.Directories = append(dirs.Directories, theDir)
		}
	}
	return dirs, nil
}

//GetFileInfo iterates through files and returns data about each file
func GetFileInfo(s string) (RepoInfo, error) {
	var files RepoInfo
	get, err := GetDirs(s)
	if err != nil {
		return files, fmt.Errorf(err.Error())
	}
	for _, v := range get {
		grep, err := ioutil.ReadDir(v)
		if err != nil {
			return files, fmt.Errorf(err.Error())
		}
		
		for _, x := range grep {
			if !x.IsDir() {
				var file File
				file.Name = x.Name()
				file.ModificationTime = x.ModTime()
				if strings.Contains(x.Name(), ".tf") {
					file.IsTfFile = true
				} else {
					file.IsTfFile = false
				}
				files.Files = append(files.Files, file)
			}
		}
	}
	return files, nil 
}

