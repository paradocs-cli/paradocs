package generate_docs

import "testing"

func TestDirectoryWalk(t *testing.T) {
	directories, err := walkDirectories(".")
	if err != nil {
		t.Logf("TestDirectoryWalk failed with error: %s", err.Error())
	} else {
		t.Logf("TestDirectoryWalk succeeded with directories: %v", directories)
	}
}
