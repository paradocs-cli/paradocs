package pkr

import (
	"fmt"
	"github.com/hashicorp/hcl/v2/hclparse"
)

// parseAttributes parses an hcl file and returns a decoded block of the attributes
func parseAttributes(filename string) (PackerAttributes, error) {
	n := hclparse.NewParser()

	file, err := n.ParseHCLFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to parse hcl file with error: %s", err.Error())
	}

	atts, err := file.Body.JustAttributes()
	if err != nil {
		return nil, fmt.Errorf("failed to parse hcl file with error: %s", err.Error())
	}

	build := make(PackerAttributes, len(atts))

	for _, v := range atts {
		build = append(build, PackerAttribute{
			Name: v.Name,
		})
	}
	return build, nil

}
