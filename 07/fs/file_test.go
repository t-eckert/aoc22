package fs_test

import (
	"dayseven/fs"
	"testing"
)

func TestNewFile(t *testing.T) {
	readout := "92461 nbvnzg"
	expectedName := "nbvnzg"
	expectedSize := 92461

	actual := fs.NewFile(readout)

	if expectedName != actual.Name {
		t.Fail()
	}

	if expectedSize != actual.Size {
		t.Fail()
	}
}
