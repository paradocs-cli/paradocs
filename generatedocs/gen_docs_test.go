package generate_docs

import "testing"

func TestCloudState_WriteMarkdownCloudState(t *testing.T) {
	type fields struct {
		Version          int
		TerraformVersion string
		Serial           int
		Lineage          string
		Outputs          interface{}
		Resources        []struct {
			Module    string `json:"module,omitempty"`
			Mode      string `json:"mode"`
			Type      string `json:"type"`
			Name      string `json:"name"`
			Provider  string `json:"provider"`
			Instances []struct {
				SchemaVersion int    `json:"schema_version"`
				IndexKey      string `json:"index_key,omitempty"`
				Attributes    struct {
					Id           string   `json:"id,omitempty"`
					Location     string   `json:"location,omitempty"`
					Name         string   `json:"name,omitempty"`
					Dependencies []string `json:"dependencies,omitempty"`
				} `json:"attributes,omitempty"`
			} `json:"instances,omitempty"`
		}
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CloudState{
				Version:          tt.fields.Version,
				TerraformVersion: tt.fields.TerraformVersion,
				Serial:           tt.fields.Serial,
				Lineage:          tt.fields.Lineage,
				Outputs:          tt.fields.Outputs,
				Resources:        tt.fields.Resources,
			}
			if err := s.WriteMarkdownCloudState(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownCloudState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStats_WriteMarkdownTerraData(t *testing.T) {
	type fields struct {
		Vars      []Var
		Resources []Resource
		Modules   []Module
		Outputs   []Output
		Datas     []Data
		Providers []Provider
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stats{
				Vars:      tt.fields.Vars,
				Resources: tt.fields.Resources,
				Modules:   tt.fields.Modules,
				Outputs:   tt.fields.Outputs,
				Datas:     tt.fields.Datas,
				Providers: tt.fields.Providers,
			}
			if err := s.WriteMarkdownTerraData(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownTerraData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStats_WriteMarkdownTerraModules(t *testing.T) {
	type fields struct {
		Vars      []Var
		Resources []Resource
		Modules   []Module
		Outputs   []Output
		Datas     []Data
		Providers []Provider
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stats{
				Vars:      tt.fields.Vars,
				Resources: tt.fields.Resources,
				Modules:   tt.fields.Modules,
				Outputs:   tt.fields.Outputs,
				Datas:     tt.fields.Datas,
				Providers: tt.fields.Providers,
			}
			if err := s.WriteMarkdownTerraModules(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownTerraModules() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStats_WriteMarkdownTerraOutputs(t *testing.T) {
	type fields struct {
		Vars      []Var
		Resources []Resource
		Modules   []Module
		Outputs   []Output
		Datas     []Data
		Providers []Provider
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stats{
				Vars:      tt.fields.Vars,
				Resources: tt.fields.Resources,
				Modules:   tt.fields.Modules,
				Outputs:   tt.fields.Outputs,
				Datas:     tt.fields.Datas,
				Providers: tt.fields.Providers,
			}
			if err := s.WriteMarkdownTerraOutputs(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownTerraOutputs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStats_WriteMarkdownTerraProviders(t *testing.T) {
	type fields struct {
		Vars      []Var
		Resources []Resource
		Modules   []Module
		Outputs   []Output
		Datas     []Data
		Providers []Provider
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stats{
				Vars:      tt.fields.Vars,
				Resources: tt.fields.Resources,
				Modules:   tt.fields.Modules,
				Outputs:   tt.fields.Outputs,
				Datas:     tt.fields.Datas,
				Providers: tt.fields.Providers,
			}
			if err := s.WriteMarkdownTerraProviders(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownTerraProviders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStats_WriteMarkdownTerraResources(t *testing.T) {
	type fields struct {
		Vars      []Var
		Resources []Resource
		Modules   []Module
		Outputs   []Output
		Datas     []Data
		Providers []Provider
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stats{
				Vars:      tt.fields.Vars,
				Resources: tt.fields.Resources,
				Modules:   tt.fields.Modules,
				Outputs:   tt.fields.Outputs,
				Datas:     tt.fields.Datas,
				Providers: tt.fields.Providers,
			}
			if err := s.WriteMarkdownTerraResources(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownTerraResources() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStats_WriteMarkdownTerraVars(t *testing.T) {
	type fields struct {
		Vars      []Var
		Resources []Resource
		Modules   []Module
		Outputs   []Output
		Datas     []Data
		Providers []Provider
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stats{
				Vars:      tt.fields.Vars,
				Resources: tt.fields.Resources,
				Modules:   tt.fields.Modules,
				Outputs:   tt.fields.Outputs,
				Datas:     tt.fields.Datas,
				Providers: tt.fields.Providers,
			}
			if err := s.WriteMarkdownTerraVars(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownTerraVars() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteMarkdownTerra(t *testing.T) {
	type args struct {
		w   Stats
		dir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteMarkdownTerra(tt.args.w, tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("WriteMarkdownTerra() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
