package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestManifestStorage(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "checksum")
	check(err)
	defer os.RemoveAll(tempDir)

	s := NewManifestStorage(tempDir)
	fmt.Println(s)
}
