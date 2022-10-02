package generate_docs

import (
	"fmt"
	"html/template"
	"os"
	"reflect"
)

//WriteMarkdownTerra generates a README which recursively documents all terraform code
//WriteMarkdownTerra does template execution to write the file
func WriteMarkdownTerra(w Stats, dir string) error {

	for i := 0; i < reflect.ValueOf(w).Type().NumField(); i++ {
		switch reflect.ValueOf(w).Type().Field(i).Name {
		case "Vars":
			err := w.WriteMarkdownTerraVars(dir)
			if err != nil {
				return err
			}
			return nil
		case "Resources":
			err := w.WriteMarkdownTerraResources(dir)
			if err != nil {
				return err
			}
			return nil
		case "Modules":
			err := w.WriteMarkdownTerraModules(dir)
			if err != nil {
				return err
			}
			return nil
		case "Outputs":
			err := w.WriteMarkdownTerraOutputs(dir)
			if err != nil {
				return err
			}
			return nil
		case "Datas":
			err := w.WriteMarkdownTerraData(dir)
			if err != nil {
				return err
			}
			return nil
		case "Providers":
			err := w.WriteMarkdownTerraProviders(dir)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func (s *Stats) WriteMarkdownTerraVars(dir string) error {
	tmpls := template.New("md")

	template.Must(tmpls.Parse(TerraformVariables))

	f, err := os.Create(fmt.Sprintf("%s/TERRAFORM_VARS.%s", dir, tmpls.Name()))
	if err != nil {
		return err
	}

	err = tmpls.Execute(f, s.Vars)
	if err != nil {
		return err
	}

	return nil
}

func (s *Stats) WriteMarkdownTerraProviders(dir string) error {
	tmpls := template.New("md")

	template.Must(tmpls.Parse(TerraformProviders))

	f, err := os.Create(fmt.Sprintf("%s/TERRAFORM_PROVIDERS.%s", dir, tmpls.Name()))
	if err != nil {
		return err
	}

	err = tmpls.Execute(f, s.Providers)
	if err != nil {
		return err
	}

	return nil
}

func (s *Stats) WriteMarkdownTerraOutputs(dir string) error {
	tmpls := template.New("md")

	template.Must(tmpls.Parse(TerraformOutputs))

	f, err := os.Create(fmt.Sprintf("%s/TERRAFORM_OUTPUTS.%s", dir, tmpls.Name()))
	if err != nil {
		return err
	}

	err = tmpls.Execute(f, s.Outputs)
	if err != nil {
		return err
	}

	return nil
}

func (s *Stats) WriteMarkdownTerraModules(dir string) error {
	tmpls := template.New("md")

	template.Must(tmpls.Parse(TerraformModules))

	f, err := os.Create(fmt.Sprintf("%s/TERRAFORM_MODULES.%s", dir, tmpls.Name()))
	if err != nil {
		return err
	}

	err = tmpls.Execute(f, s.Modules)
	if err != nil {
		return err
	}

	return nil
}

func (s *Stats) WriteMarkdownTerraData(dir string) error {
	tmpls := template.New("md")

	template.Must(tmpls.Parse(TerraformDataSources))

	f, err := os.Create(fmt.Sprintf("%s/TERRAFORM_DATA.%s", dir, tmpls.Name()))
	if err != nil {
		return err
	}

	err = tmpls.Execute(f, s.Datas)
	if err != nil {
		return err
	}

	return nil
}

func (s *Stats) WriteMarkdownTerraResources(dir string) error {
	tmpls := template.New("md")

	template.Must(tmpls.Parse(TerraformResources))

	f, err := os.Create(fmt.Sprintf("%s/TERRAFORM_RESOURCES.%s", dir, tmpls.Name()))
	if err != nil {
		return err
	}

	err = tmpls.Execute(f, s.Resources)
	if err != nil {
		return err
	}

	return nil
}

//WriteMarkdownCloudState generates a README which recursively documents all terraform state resources from the specified backend
//WriteMarkdownCloudState does template execution to write the file
func WriteMarkdownCloudState(st CloudState) os.File {

	tmpls := template.New("md")
	template.Must(tmpls.Parse(TerraStateDocCloud))

	f, err := os.Create(fmt.Sprintf("STATE.%s", tmpls.Name()))
	if err != nil {
		if err != nil {
			fmt.Println("Unable to create file for Markdown")
			os.Exit(1)
		}
	}

	err = tmpls.Execute(f, st)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return *f
}

//WriteMarkdownTfcState generates a README which recursively documents all terraform state resources from the specified backend
//WriteMarkdownTfcState does template execution to write the file
func WriteMarkdownTfcState(st TfcState) os.File {

	tmpls := template.New("md")
	template.Must(tmpls.Parse(TerraStateDocCloud))

	f, err := os.Create(fmt.Sprintf("STATE.%s", tmpls.Name()))
	if err != nil {
		if err != nil {
			fmt.Println("Unable to create file for Markdown")
			os.Exit(1)
		}
	}

	err = tmpls.Execute(f, st)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return *f
}
