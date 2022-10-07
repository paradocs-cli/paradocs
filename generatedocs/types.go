package generate_docs

import "time"


type Stats struct {
	Vars      []Var      `json:"variables,omitempty"`
	Resources []Resource `json:"resources,omitempty"`
	Modules   []Module   `json:"modules,omitempty"`
	Outputs   []Output   `json:"outputs,omitempty"`
	Datas     []Data     `json:"datas,omitempty"`
	Providers []Provider `json:"providers,omitempty"`
}
type Var struct {
	VarName                string      `json:"name,omitempty"`
	VarType                string      `json:"type,omitempty"`
	VarDescription         string      `json:"description,omitempty"`
	VarDefault             interface{} `json:"default_value,omitempty"`
	VarRequired            bool        `json:"required,omitempty"`
	VarSensitive           bool        `json:"sensitive,omitempty"`
	SourcePositionFileName string      `json:"source_file_name,omitempty"`
	SourcePositionLine     string      `json:"source_position_line,omitempty"`
}
type Resource struct {
	Mode                   string `json:"mode,omitempty"`
	Type                   string `json:"type,omitempty"`
	Name                   string `json:"name,omitempty"`
	ProviderName           string `json:"provider_name,omitempty"`
	ProviderAlias          string `json:"provider_alias,omitempty"`
	SourcePositionFileName string `json:"source_file_name,omitempty"`
	SourcePositionLine     string `json:"source_position_line,omitempty"`
	Link string
}

type Module struct {
	Name                   string `json:"name,omitempty"`
	ModSource              string `json:"module_source,omitempty"`
	Version                string `json:"module_version,omitempty"`
	SourcePositionFileName string `json:"source_file_name,omitempty"`
	SourcePositionLine     string `json:"source_position_line,omitempty"`
}

type Output struct {
	Name                   string `json:"name,omitempty"`
	Description            string `json:"description,omitempty"`
	Sensitive              bool   `json:"module_version,omitempty"`
	SourcePositionFileName string `json:"source_file_name,omitempty"`
	SourcePositionLine     string `json:"source_position_line,omitempty"`
	Link string
}

type Data struct {
	Name                   string `json:"name,omitempty"`
	DataType               string `json:"data_type,omitempty"`
	ProviderName           string `json:"provider_name,omitempty"`
	ProviderAlias          string `json:"provider_alias,omitempty"`
	SourcePositionFileName string `json:"source_file_name,omitempty"`
	SourcePositionLine     string `json:"source_position_line,omitempty"`
	Link string
}

type Provider struct {
	Name  string `json:"name,omitempty"`
	Alias string `json:"alias,omitempty"`
	Link  string
}

type Dirs struct {
	Name string
	ModificationTime time.Time
	IsTerraDir bool
}

type File struct {
	Name string
	ModificationTime time.Time
	IsTfFile bool
}

type RepoInfo struct {
	Directories []Dirs
	Files []File
}

type TfcState struct {
	Data struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			CreatedAt              time.Time `json:"created-at"`
			Size                   int       `json:"size"`
			HostedStateDownloadUrl string    `json:"hosted-state-download-url"`
			Modules                struct {
			} `json:"modules"`
			Providers struct {
			} `json:"providers"`
			Resources          []struct{
				Module string `json:"module,omitempty"`
				Mode string `json:"mode"`
				Type string `json:"type"`
				Name string `json:"name"`
				Provider string `json:"provider"`
				Instances []struct{
					SchemaVersion int `json:"schema_version"`
					IndexKey string `json:"index_key,omitempty"`
					Attributes struct{
						Id string `json:"id,omitempty"`
						Location string `json:"location,omitempty"`
						Name string `json:"name,omitempty"`
						Dependencies []string `json:"dependencies,omitempty"`
					} `json:"attributes,omitempty"`
				} `json:"instances,omitempty"`
			} `json:"resources"`
			ResourcesProcessed bool          `json:"resources-processed"`
			Serial             int           `json:"serial"`
			StateVersion       int           `json:"state-version"`
			TerraformVersion   string        `json:"terraform-version"`
			VcsCommitUrl       interface{}   `json:"vcs-commit-url"`
			VcsCommitSha       interface{}   `json:"vcs-commit-sha"`
		} `json:"attributes"`
		Relationships struct {
			Run struct {
				Data interface{} `json:"data"`
			} `json:"run"`
			CreatedBy struct {
				Data struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
				Links struct {
					Self string `json:"self"`
				} `json:"links"`
			} `json:"created-by"`
			Workspace struct {
				Data struct {
					Id   string `json:"id"`
					Type string `json:"type"`
				} `json:"data"`
			} `json:"workspace"`
			Outputs struct {
				Data  []interface{} `json:"data"`
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"outputs"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}

type CloudState struct {
	Version          int         `json:"version"`
	TerraformVersion string      `json:"terraform_version"`
	Serial           int         `json:"serial"`
	Lineage          string      `json:"lineage"`
	Outputs          interface{} `json:"outputs"`
	Resources        []struct{
		Module string `json:"module,omitempty"`
		Mode string `json:"mode"`
		Type string `json:"type"`
		Name string `json:"name"`
		Provider string `json:"provider"`
		Instances []struct{
			SchemaVersion int `json:"schema_version"`
			IndexKey string `json:"index_key,omitempty"`
			Attributes struct{
				Id string `json:"id,omitempty"`
				Location string `json:"location,omitempty"`
				Name string `json:"name,omitempty"`
				Dependencies []string `json:"dependencies,omitempty"`
			} `json:"attributes,omitempty"`
		} `json:"instances,omitempty"`
	} `json:"resources"`
}

type StateProviders struct {
	Azure struct {
		StorageAccountName string
		ContainerName string
		BlobName string
		SasToken string
	}
	Aws struct {
		AccessKey string
		SecretAccessKey string
		SessionToken string
		BucketName string
		Object string
		Region string
	}
	TerraformCloud struct {
		ApiToken string
		WorkspaceId string
	}
	GoogleCloud struct {
		BucketName string
		ObjectName string
		Oauth2Token string
	}
}