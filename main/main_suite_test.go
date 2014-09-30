package main_test

import (
	"os"
	"os/exec"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("go", "build", "-o", path.Join(dir, "..", "fixtures", "plugins", "test_1"), path.Join(dir, "..", "fixtures", "plugins", "test_1.go"))
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	cmd = exec.Command("go", "build", "-o", path.Join(dir, "..", "fixtures", "plugins", "test_2"), path.Join(dir, "..", "fixtures", "plugins", "test_2.go"))
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	RunSpecs(t, "Main Suite")
}
