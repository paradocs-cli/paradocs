package pipelines

type AzureDevops struct {
	Trigger   interface{} `yaml:"trigger"`
	Variables []struct {
		Group    string `yaml:"group,omitempty"`
		Template string `yaml:"template,omitempty"`
		Name     string `yaml:"name,omitempty"`
		Value    string `yaml:"value,omitempty"`
	} `yaml:"variables"`
	Name   string `yaml:"name,omitempty"`
	Stages []struct {
		Stage       string `yaml:"stage,omitempty"`
		Name        string `yaml:"name,omitempty"`
		DisplayName string `yaml:"displayName,omitempty"`
		Jobs        []struct {
			Job         string `yaml:"job"`
			DisplayName string `yaml:"displayName"`
			Pool        struct {
				VMImage string `yaml:"vmImage"`
			} `yaml:"pool"`
			Steps []struct {
				Script      string `yaml:"script,omitempty"`
				DisplayName string `yaml:"displayName"`
				Task        string `yaml:"task,omitempty"`
			} `yaml:"steps"`
		} `yaml:"jobs,omitempty"`
		DependsOn string `yaml:"dependsOn,omitempty"`
		Condition string `yaml:"condition,omitempty"`
	} `yaml:"stages"`
}
