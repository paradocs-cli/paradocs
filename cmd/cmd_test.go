package cmd

import "testing"

func TestTerraformCmd(t *testing.T) {
	cmd := NewTerraformCmd(".")
	err := cmd.Execute()
	if err != nil {
		t.Fail()
		t.Logf(err.Error())
	}
}

