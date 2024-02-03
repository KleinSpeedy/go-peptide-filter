package main

import (
	"fmt"
	"os"
	"peptide-analyse/pepcli"
	"sort"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:   "peptide-analyse",
		Usage:  "analyse peptides and filter them according to a mass range",
		Action: pepcli.CliAction,
		Before: pepcli.CheckBefore,
		Flags:  pepcli.Flags,
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		ts := time.Now().Format(time.DateTime)
		fmt.Println("Program exited with errors - ", ts)
		fmt.Print(err)
	}
}
