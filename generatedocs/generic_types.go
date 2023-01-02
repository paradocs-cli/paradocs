package generate_docs

type HclVariables map[string]HclVariable

type HclVariable struct {
	Name        string
	Default     interface{}
	Type        interface{}
	Description string
}

type HclBlock struct {
	Name        string
	Default     interface{}
	Type        interface{}
	Description string
}
