package main_test

import (
	"os"
	"peptide-analyse/pepcli"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCheckBeforeInvalidArgs(t *testing.T) {
	app := cli.App{
		Action: pepcli.CliAction,
		Before: pepcli.CheckBefore,
		Flags:  pepcli.Flags,
	}

	// invalid mass range
	invalidArgs := [...]string{
		"", // cli flags start at index 1
		"-s",
		"1000",
		"-e",
		"900",
		"-f",
		"test_data/test.fasta",
	}

	err := app.Run(invalidArgs[:])
	if err == nil {
		t.Error("Got nil, expected error")
	}
}

func TestCheckBeforeValidArgs(t *testing.T) {
	tmpStdout := os.Stdout
	defer func() { os.Stdout = tmpStdout }()
	var ferr error
	os.Stdout, ferr = os.Open(os.DevNull)
	if ferr != nil {
		t.Errorf("Could not redirect to dev null: %s", ferr.Error())
	}

	app := cli.App{
		Action: pepcli.CliAction,
		Before: pepcli.CheckBefore,
		Flags:  pepcli.Flags,
	}

	validArgs := [...]string{
		"", // cli flags start at index 1
		"-s",
		"1000",
		"-e",
		"1200",
		"-f",
		"test_data/test.fasta",
	}

	err := app.Run(validArgs[:])
	if err != nil {
		t.Errorf("Got error, expected nil: %s", err.Error())
	}
}
