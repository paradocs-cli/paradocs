package cmd

import "testing"

func TestTerraformCmd(t *testing.T) {
	cmd := NewTerraformCmd(".")
	err := cmd.Execute()
	if err != nil {
		t.Fail()
		t.Logf(err.Error())
	} else {
		t.Logf("Test passed for terraform sub command!")
	}
}

func TestCodeCmd(t *testing.T) {
	cmd := NewCodeCommand()
	err := cmd.Execute()
	if err != nil {
		t.Fail()
		t.Logf(err.Error())
	} else {
		t.Logf("Test passed for code sub command!")
	}
}